---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- http_proxy.more_option

<a id="canonical-2221001223210112-1100323203103002-1221130002010130-3330211303220210-3321303213313233-3133102003321312-1003210223123213-2133113122313131"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to define a route.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("max_requests_per_connection",
    "no_request_limit_per_connection")}
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
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-strict_sni_host_header_check_choice": "[]"
}
```

Terraform syntax:

```terraform
more_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020232023301010-0233323002131131-2120010131020320-3120003132102001-3332222110312330-3022300300110012-1213123031021000-1101021210002230"></a>

### Direct properties for `http_proxy.more_option`

- [buffer_policy](resources--proxy--reference--group-004.md#canonical-0300111113000322-0301122303123033-2331031131312202-2232013311303303-2130110102131033-3332231021130301-2213011301023020-3333033311323303): complete subsection reference.

- [compression_params](resources--proxy--reference--group-004.md#canonical-3320023111221201-1200202123231322-0232100223203033-2300211013131021-2111033013110111-1031222122321032-3331001111222321-1223223332010113): complete subsection reference.

<a id="canonical-0313022031120112-0330201322330102-2202030133331203-0310301212223021-0011100232010303-2121001310020313-3100201212011233-0003031130000313"></a>

<a id="canonical-0311123313000323-3032100313330001-0132233323100122-3333122000010220-3211001000312030-3211220003330100-3012103002332101-0201013130220301"></a>

#### `http_proxy.more_option.custom_errors` property

Type: `["map", "string"]`. Optional.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP
responses. Specific response code takes preference when both response code and response code class
matches for a request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"ranges\":[[3,3],[4,4],[5,5],[300,599]],\"type\":\"uint32-string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.uint32.ranges\":\"3,4,5,300-599\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"65536\",\"ves.io.schema.rules.map.values.string.uri_ref\":\"true\"},\"values\":{\"format\":\"uri-reference\",\"maxLength\":65536,\"type\":\"string\"}}")}
```

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
      "ranges": [
        [
          3,
          3
        ],
        [
          4,
          4
        ],
        [
          5,
          5
        ],
        [
          300,
          599
        ]
      ],
      "type": "uint32-string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "65536",
      "ves.io.schema.rules.map.values.string.uri_ref": "true"
    },
    "values": {
      "format": "uri-reference",
      "maxLength": 65536,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  }
}
```

<a id="canonical-3231320000303110-1012233021103302-2102122331023303-0231021323111200-3100110000332321-3223322102131121-3110020330303030-3210220023021330"></a>

<a id="canonical-2312122322202212-2010001030012110-1122201233312313-0031210302303223-2013320113122023-2332212032310331-1122122020320112-3110130113203311"></a>

#### `http_proxy.more_option.disable_default_error_pages` property

Type: `"bool"`. Optional.

Disable the use of default F5XC error pages.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_path_normalize](resources--proxy--reference--group-004.md#canonical-0233013301333121-0003002030332233-0221001332302133-0113132323023221-0313010200131210-2102233002103312-3131300033223113-1102011222311112): complete subsection reference.

- [enable_path_normalize](resources--proxy--reference--group-004.md#canonical-2301321023223101-2012123133321310-3120133122312330-3322110201211220-2120110022201321-0230303210303012-3111320112022320-2300212111030323): complete subsection reference.

<a id="canonical-0112113002130301-2002113131330321-0132330230130300-3133311110212013-3200102133232212-3103213231203200-2121130310333101-3332133011113012"></a>

<a id="canonical-1002022303133302-2201102200333221-0033233313010001-3133312131103010-1101302103030303-3102001313123021-3212230121123203-0302003301000003"></a>

#### `http_proxy.more_option.idle_timeout` property

Type: `"number"`. Optional.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with an HTTP 504 (Gateway Timeout) error code if no upstream response
header has been received, otherwise the stream is reset.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(3600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3600000,
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
    "ves.io.schema.rules.uint32.lte": "3600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  }
}
```

<a id="canonical-1310120112213001-2201330230223221-0020012310332112-2012013230030010-3312203223100101-3030122030323330-3130333110003021-1321322201333201"></a>

<a id="canonical-3202302321000332-0330030232033300-2033031013330210-1002323221110200-2332331330221103-3130301112203130-3103323101033120-2301321313212121"></a>

#### `http_proxy.more_option.max_request_header_size` property

Type: `"number"`. Optional.

The maximum request header size for downstream connections, in KiB. An HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 96,
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
    "ves.io.schema.rules.uint32.lte": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  }
}
```

<a id="canonical-1002130211130201-2101032201222021-1120321121103302-1300220103010130-3100122213330011-2133111200121022-1121310331111122-2202110311000302"></a>

<a id="canonical-0021022200103000-0233111202012332-2213032011231112-2322223003233322-3321231002133300-0301200213202320-1122011133022022-1233123103003013"></a>

#### `http_proxy.more_option.max_requests_per_connection` property

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [no_request_limit_per_connection](resources--proxy--reference--group-004.md#canonical-1003122322301321-1333110123011112-3221222213203212-1023220213321033-2023002003031031-3010120133222233-1001030333011203-3301303002212231): complete subsection reference.

- [request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-1311033121201011-3223001021303303-1330100313013302-1300333130322122-3121002120320021-0010212303011003-1032011123031100-3230213212313312): complete subsection reference.

<a id="canonical-0120232230030030-2321001331333320-2021302031113212-3132332013031220-3332323000313301-1212230310331123-3013320321310111-3312000112021333"></a>

<a id="canonical-3320300021320033-0223230010301033-3320102220233211-1320313110112313-1023233220221132-3302320121133213-0220302022112301-0032312022303332"></a>

#### `http_proxy.more_option.request_cookies_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](resources--proxy--reference--group-004.md#canonical-2310002312312021-3131313220231323-3103000310102012-1232012231303310-1022012331001022-1220023212200233-2312330020030233-3132033101221100): complete subsection reference.

<a id="canonical-0123101101130310-1130221002002033-2220012002122330-3223202111300120-0333333200033323-3003032122220102-0221101233303313-0210333023012213"></a>

<a id="canonical-2020210032101331-3030222131020211-1032112201130312-3123203022221213-2112302000322013-0133112320131332-0333202030333303-3122011203330103"></a>

#### `http_proxy.more_option.request_headers_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311): complete subsection reference.

<a id="canonical-2100130031301331-1211022003131011-2331122102111202-1033032130122033-0002322110223220-1221031322133213-0311213110022020-3030322110233211"></a>

<a id="canonical-1131203301130313-0002201130103010-3123331213303001-3213011003223223-0300110123131032-0131221210130032-1133123231000301-1222103110023003"></a>

#### `http_proxy.more_option.response_cookies_to_remove` property

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](resources--proxy--reference--group-005.md#canonical-0321302212120332-1003310003320101-2331211322200330-1231022123102030-0001322133103022-1202202120331013-0312112103131001-3233331200001032): complete subsection reference.

<a id="canonical-2002112100110131-0103300002133100-0010332111230100-1233301232121321-2232103211230233-3002313103012231-1011131310023122-0323031120030323"></a>

<a id="canonical-2113033321113331-1331200212230331-2223220222331033-3231330232301212-3121130100331100-1113213310023031-1011113133213203-3033320100213102"></a>

#### `http_proxy.more_option.response_headers_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0300111113000322-0301122303123033-2331031131312202-2232013311303303-2130110102131033-3332231021130301-2213011301023020-3333033311323303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.buffer_policy` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.buffer_policy

<a id="canonical-2030331223113202-1301222121300032-1333023200302230-2223210310123030-2103212211031020-3020122131311010-3212011120111103-0222130023203220"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3112303211310331-3220110013112010-1100220323203020-1330203003033100-3321313222013032-2231002130103211-1101002321022120-1310021013320012"></a>

### Direct properties for `http_proxy.more_option.buffer_policy`

<a id="canonical-3030213210301303-3333030201220321-0230101011202023-2003022131230110-2101302003303110-3001120103222012-3333331030333213-3200333203112001"></a>

#### `http_proxy.more_option.buffer_policy.disabled` property

Type: `"bool"`. Optional.

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

<a id="canonical-3022321331221031-3101331230333301-0021011112311301-0221312013023013-2132031333301013-2331310300322220-0031213220302100-1133310220022010"></a>

<a id="canonical-2100101132131302-2030303133100103-0111200131013131-2101112003223213-2210203313231331-0031211031000103-1112333301212103-0103022212322332"></a>

#### `http_proxy.more_option.buffer_policy.max_request_bytes` property

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3320023111221201-1200202123231322-0232100223203033-2300211013131021-2111033013110111-1031222122321032-3331001111222321-1223223332010113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.compression_params` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.compression_params

<a id="canonical-0320033203010212-1211021031113132-2000000320222000-3110300022022233-1020031200101233-0020321320102303-1032023200311321-1231210233003101"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

By default compression will be skipped when:

A request does NOT contain accept-encoding header. A request includes accept-encoding header, but it
does not contain “gzip” or “\*”. A request includes accept-encoding with “gzip” or “\*” with the
weight “q=0”. Note that the “gzip” will have a higher weight then “\*”. For example, if
accept-encoding is “gzip;q=0,\*;q=1”, the filter will not compress. But if the header is set to
“\*;q=0,gzip;q=1”, the filter will compress. A request whose accept-encoding header includes
“identity”. A response contains a content-encoding header. A response contains a cache-control
header whose value includes “no-transform”. A response contains a transfer-encoding header whose
value includes “gzip”. A response does not contain a content-type value that matches one of the
selected mime-types, which default to application/JavaScript, application/JSON,
application/xhtml+XML, image/svg+XML, text/CSS, text/HTML, text/plain, text/XML. Neither
content-length nor transfer-encoding headers are present in the response. Response size is smaller
than 30 bytes (only applicable when transfer-encoding is not chunked).

When compression is applied:

The content-length is removed from response headers. Response headers contain “transfer-encoding:
chunked” and do not contain “content-encoding” header. The “vary: accept-encoding” header is
inserted on every response.

GZIP Compression Level:

A value which is optimal balance between speed of compression and amount of compression is chosen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("content_length")}
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
compression_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313313113021002-0003002123021332-2011123130032132-0231313031022302-2130210103110220-3010121001301031-2321311001011333-3200132331111111"></a>

### Direct properties for `http_proxy.more_option.compression_params`

<a id="canonical-2231331100321200-3031101132033022-3311333302312010-0302120312200032-0303300133131332-1202231322023211-0320033210230112-1230101110200303"></a>

#### `http_proxy.more_option.compression_params.content_length` property

Type: `"number"`. Optional.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Additional upstream details:

The default value is 30.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(30),
}
```

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
    "minimum": 30
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  }
}
```

<a id="canonical-0132030320221310-0320321031021113-0212032212130010-1101103000110030-0001123021221321-3003332333033103-3001333000010201-1132113220310322"></a>

<a id="canonical-0003132031320301-0122101322220022-2211313013002133-3100120331320301-0110310203301210-1102302203210200-0020102102111201-2320010322103031"></a>

#### `http_proxy.more_option.compression_params.content_type` property

Type: `["list", "string"]`. Optional.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/JavaScript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Additional upstream details:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/JavaScript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
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
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1011211133321111-1303000003000102-1123222321310302-1200221131213310-3120023031123200-0120200020000312-3031232212120311-3212222213031032"></a>

<a id="canonical-2002000113021303-2033023101000000-3112300011221101-0231221211001101-2213311121223321-1220112003310202-3330112332031002-2331320112300023"></a>

#### `http_proxy.more_option.compression_params.disable_on_etag_header` property

Type: `"bool"`. Optional.

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2132023021032303-3122203000311123-0313311221113132-1120321210130112-1200101120313313-1232323222000222-0321031221010231-2230333022231131"></a>

<a id="canonical-3220033112112102-1231102320303002-2012103322133020-3133133213002120-1101301001102010-0312003233001203-0201322323113001-3100001221333232"></a>

#### `http_proxy.more_option.compression_params.remove_accept_encoding_header` property

Type: `"bool"`. Optional.

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0233013301333121-0003002030332233-0221001332302133-0113132323023221-0313010200131210-2102233002103312-3131300033223113-1102011222311112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.disable_path_normalize

<a id="canonical-0112221233100331-0113012122102211-3011200303303011-1031032220303131-1100323202123211-2233030223201213-1331122233020302-2000333322130033"></a>

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
disable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301321023223101-2012123133321310-3120133122312330-3322110201211220-2120110022201321-0230303210303012-3111320112022320-2300212111030323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.enable_path_normalize

<a id="canonical-0203101203312302-2312202110311310-2220302001001003-0320023221131300-3001112322010233-1212200111231300-3222203231201223-2301302213220111"></a>

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
enable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003122322301321-1333110123011112-3221222213203212-1023220213321033-2023002003031031-3010120133222233-1001030333011203-3301303002212231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.no_request_limit_per_connection` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.no_request_limit_per_connection

<a id="canonical-2013003000303112-2201123200223002-3133201302102000-3033010231020010-0210033222213323-0301020130232020-0101322331132030-1301132223330003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

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
no_request_limit_per_connection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311033121201011-3223001021303303-1330100313013302-1300333130322122-3121002120320021-0010212303011003-1032011123031100-3230213212313312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_cookies_to_add` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.request_cookies_to_add

<a id="canonical-1210000331112233-3112003330210012-0230201313032113-0123320223110231-2301003312222303-3032133310110022-2302212333333231-3033003101121103"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0011112113222330-1332331123121201-0330121302230122-3330201023122220-2302322030223300-3300333213220112-3330000313130233-0102022110213233"></a>

### Direct properties for `http_proxy.more_option.request_cookies_to_add`

<a id="canonical-3102202110110102-0012310312222312-3333030331110311-3302131230122302-0211312030300201-3313231310101332-0313231013210101-1211101233123310"></a>

#### `http_proxy.more_option.request_cookies_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0203322032102113-3202110212331020-2201301320303113-3203320100333101-2121131123132332-0221202001110130-0222030111200311-3113022000101310"></a>

<a id="canonical-2313302220122213-3130321100102202-3301133033231231-2032200031210112-3203311111131231-2023001101003233-1330030330031202-0203102121223330"></a>

#### `http_proxy.more_option.request_cookies_to_add.overwrite` property

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

- [secret_value](resources--proxy--reference--group-004.md#canonical-3010332010112330-1233131123133131-3000223203333231-2330012221023020-2212031000021203-0123020021231110-0201332132110201-0113212212212331): complete subsection reference.

<a id="canonical-2111203323003033-0030313313112322-2201132003320111-3201033321221330-3023101110133332-1002300320131011-0210301200013301-2002113232202012"></a>

<a id="canonical-2333322323123331-0033320032221210-3123203023032201-1200322021202030-3220321100220220-3200300222011023-2203232313233200-2321023113020120"></a>

#### `http_proxy.more_option.request_cookies_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-3010332010112330-1233131123133131-3000223203333231-2330012221023020-2212031000021203-0123020021231110-0201332132110201-0113212212212331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-1311033121201011-3223001021303303-1330100313013302-1300333130322122-3121002120320021-0010212303011003-1032011123031100-3230213212313312)
- http_proxy.more_option.request_cookies_to_add.secret_value

<a id="canonical-0032302311112022-3220210031202120-2100222202123221-3110333132333112-1110031030233223-3233100312232031-0032130020101300-1000002031133020"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2321331231312333-0233323000302113-2220312222320332-3201033131120313-0132321330021001-0100303130112303-2123001233011102-1210310012022022"></a>

### Direct properties for `http_proxy.more_option.request_cookies_to_add.secret_value`

- [blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-2300202030130121-3302312131303030-1320112132301301-0023333100030123-0123113131313031-1220202201322003-3022213331231121-3122002202132222): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-004.md#canonical-3032322112310300-0330132331133331-3112033012013211-0003331023313132-0320021013213232-1012022031023230-2033220310022330-2322301222330131): complete subsection reference.

<a id="canonical-2300202030130121-3302312131303030-1320112132301301-0023333100030123-0123113131313031-1220202201322003-3022213331231121-3122002202132222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-1311033121201011-3223001021303303-1330100313013302-1300333130322122-3121002120320021-0010212303011003-1032011123031100-3230213212313312)
- [http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-3010332010112330-1233131123133131-3000223203333231-2330012221023020-2212031000021203-0123020021231110-0201332132110201-0113212212212331)
- http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-1010010223100223-0203100100122031-2300201230110212-0332010211032022-2111001132323003-2122122113130210-1112021023210010-1331200210011110"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1111013003010020-2113232111103323-3230332102023022-3323220113000201-2111012331313130-1311313301113013-1121321231111230-3230113232221112"></a>

### Direct properties for `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3331331233222231-0021031113232222-2111002202000012-2112002322323321-2320000303231120-2323112312132033-3112210201332202-1001011312302103"></a>

#### `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1010032322301332-3211332133013000-1303332212023012-3320331120001233-1003020302001033-1221120001200103-1000012133013212-0112011030212002"></a>

<a id="canonical-3110233012302110-2011322012220311-2231020212312231-3333030333330030-1223130102011023-0003310233230331-2122130223122112-2221120332220200"></a>

#### `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3310200010310031-0330112202113002-3203031210003220-0233121203200023-1013132132323131-0323001231331010-1200332010123130-1010000210300011"></a>

<a id="canonical-3021131013123112-2000002023201033-2301313110100331-2011213213130313-2130322011312330-2100331220003211-0131021121201321-2310321311030302"></a>

#### `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-3032322112310300-0330132331133331-3112033012013211-0003331023313132-0320021013213232-1012022031023230-2033220310022330-2322301222330131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-1311033121201011-3223001021303303-1330100313013302-1300333130322122-3121002120320021-0010212303011003-1032011123031100-3230213212313312)
- [http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-3010332010112330-1233131123133131-3000223203333231-2330012221023020-2212031000021203-0123020021231110-0201332132110201-0113212212212331)
- http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-3021131101030110-2220002212111122-2111220200213300-1201102330123030-0112213023022100-3232202333202031-3211222223310122-2013100302113113"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3130131220121130-1300322333013130-2303331130201033-2333120231331032-0020132302302131-3312022231112120-0121202000132301-0333012013213103"></a>

### Direct properties for `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-3220111300202111-1331000010032330-3032021200011123-0303133110102211-1200123231213011-0012011202002103-1333302102320303-1313321003132013"></a>

#### `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2211233212012220-2333132133213302-0321220200333222-3302300010320200-2133103300132103-1222023032333011-0330311130002031-3322022202002202"></a>

<a id="canonical-2302303002302000-0200311232030103-0112322200130130-1213013033302232-2023300302333110-2233012213201003-3332033033011122-2100000112233022"></a>

#### `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2310002312312021-3131313220231323-3103000310102012-1232012231303310-1022012331001022-1220023212200233-2312330020030233-3132033101221100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_headers_to_add` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.request_headers_to_add

<a id="canonical-1230012112232311-2322032002321033-2203021003032312-3332130230101301-3110110320203223-1023021203021101-0233010213021222-1302103230000110"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0202100030023123-2202311110013300-2232220211203113-2002302310220230-2220331331323120-2032002000020233-3120102122210013-3000022231323102"></a>

### Direct properties for `http_proxy.more_option.request_headers_to_add`

<a id="canonical-3013220321331202-1030322313102032-2203300013232001-0112212112332122-0223101002302202-2221313003312030-1302223201031223-0212323003301231"></a>

#### `http_proxy.more_option.request_headers_to_add.append` property

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

<a id="canonical-1203012110003233-1113331323000333-3130231333031022-0102322122211300-2020331011200310-0333301012303131-2120333330313030-3102321001231111"></a>

<a id="canonical-0000230013332011-1202100302130211-3000233003331313-1000221231112132-2223021010003300-3002331010130223-2303131323011321-0113102300312122"></a>

#### `http_proxy.more_option.request_headers_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](resources--proxy--reference--group-004.md#canonical-1221020210303123-2323332201130110-1022223200200331-3021132112233032-3232210111232213-0313113330212103-3321130301022111-3333302221133320): complete subsection reference.

<a id="canonical-2003232011312130-3313211102303323-0310323312310323-0230112133100213-2033323010231020-2131132121133023-0331010232001333-3111100201002021"></a>

<a id="canonical-2223011012010320-1133212003101321-2103221011321030-2021212312231022-3300301203100203-0323133021033332-3121312020001322-2331020133311303"></a>

#### `http_proxy.more_option.request_headers_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-1221020210303123-2323332201130110-1022223200200331-3021132112233032-3232210111232213-0313113330212103-3321130301022111-3333302221133320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-2310002312312021-3131313220231323-3103000310102012-1232012231303310-1022012331001022-1220023212200233-2312330020030233-3132033101221100)
- http_proxy.more_option.request_headers_to_add.secret_value

<a id="canonical-2300231120000222-2210203002022322-3020322312201133-2232223001000320-1202201021113203-2021320210113032-2130000233333201-3232300313103313"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0211331033303320-3000102002002020-2101100012232311-3031301030000203-1202021223230111-2010101121322212-1223001103133232-3220110001113212"></a>

### Direct properties for `http_proxy.more_option.request_headers_to_add.secret_value`

- [blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-2110202002210301-0033222003032113-0131101202201201-3313202031333333-2312133023131102-1210110303231232-2333321320233002-3130112013012302): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-004.md#canonical-1200031203213032-3231313111013212-0203302011121021-2310312111230312-0121113312123120-1010000110101111-2132311300210013-2002030213131330): complete subsection reference.

<a id="canonical-2110202002210301-0033222003032113-0131101202201201-3313202031333333-2312133023131102-1210110303231232-2333321320233002-3130112013012302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-2310002312312021-3131313220231323-3103000310102012-1232012231303310-1022012331001022-1220023212200233-2312330020030233-3132033101221100)
- [http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-1221020210303123-2323332201130110-1022223200200331-3021132112233032-3232210111232213-0313113330212103-3321130301022111-3333302221133320)
- http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-1031322130332031-2313202200012232-3312032121110312-1213002022212332-1112332000003313-1031100100223332-0201222221111003-2320233232230333"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3023112213210312-2322131321101301-0131033213103311-3323323133213300-2023332032120101-3310022033333322-1301112300023021-3102131013332022"></a>

### Direct properties for `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3113013301200113-2021031312203012-0022322200302100-0020312111223002-1031120031200211-1232020302301121-2012302332031111-2220030203211113"></a>

#### `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3231111013321022-1000303002003201-2100312113112020-0211213311002030-3330223130233302-3231112232331213-3222331202001300-2120311122110320"></a>

<a id="canonical-0310121222113312-2000100121232323-1010102302103103-3201232033200303-2301303322122323-0221021203032201-0203201132002003-1210313320131310"></a>

#### `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2013303323012021-2131220223212323-3120221321111313-2313020222111130-0221003320201211-1032000122202230-1230312022310210-3022000233320020"></a>

<a id="canonical-0100323202221330-0310321123002011-2201001213000122-0232300111013321-3002121200003021-2022111110300002-0131222100103321-2020333113033310"></a>

#### `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-1200031203213032-3231313111013212-0203302011121021-2310312111230312-0121113312123120-1010000110101111-2132311300210013-2002030213131330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-2310002312312021-3131313220231323-3103000310102012-1232012231303310-1022012331001022-1220023212200233-2312330020030233-3132033101221100)
- [http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-1221020210303123-2323332201130110-1022223200200331-3021132112233032-3232210111232213-0313113330212103-3321130301022111-3333302221133320)
- http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-2311311000202303-3023321133211030-0112113010311230-1332312133230231-2130111202112000-3123031200113022-3122200313203010-2001000200300322"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2311030031322230-2232311322200231-2031223201320000-2323113300032112-1113030203010320-3103113033333122-3010103201200330-2201231202113301"></a>

### Direct properties for `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-0012113100031233-3332000210230320-2303202321030130-1223201322011211-1330121030311232-1222103331130032-2313211013211210-2020021013101101"></a>

#### `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0310201103103212-1101311330010203-3332331312231012-0013322012220211-0213232212022301-0322200213103303-0003330120312123-1010132020302110"></a>

<a id="canonical-2222301312213231-0330111301021312-3100012121103132-2130013113230313-2002111311102213-0303223233023231-0212003313130313-0021330311221012"></a>

#### `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.response_cookies_to_add

<a id="canonical-1110113233031013-0303232013113321-1111012232011023-0231212322222302-2311002212132112-3102031331333331-1020030303221002-3023103221320021"></a>

Type: `"object"`. list nested block, Optional.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2201011313210012-1131320231100330-3223323301200010-0322121330030130-3220131311101012-0010023021212021-1332321223002212-1121003320302330"></a>

### Direct properties for `http_proxy.more_option.response_cookies_to_add`

<a id="canonical-3013001312222112-1301311032110131-3120013310333030-2121033123320011-2023201201220213-3320201321221022-1113122020132030-2022003033110113"></a>

#### `http_proxy.more_option.response_cookies_to_add.add_domain` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3232101222113000-0201002030012012-2131211312020101-0123201313133123-0130311313032102-2230001130101220-3122310131012033-3013210000333232"></a>

<a id="canonical-1213023112022122-2110212011101003-2121033303323021-2320201031201221-3112031211102322-3023231000221203-1323223311132031-1310130212200033"></a>

#### `http_proxy.more_option.response_cookies_to_add.add_expiry` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

- [add_httponly](resources--proxy--reference--group-004.md#canonical-0322013202332010-2030231103122231-3231303232203303-2032203302113112-2031030000332031-3132330213131202-0013233332130021-1131221202111120): complete subsection reference.

- [add_partitioned](resources--proxy--reference--group-004.md#canonical-3120333300200020-3220011211122220-0210232103101022-0111313302133333-2033110320012003-3203013202211011-3003222221201211-1031032310222321): complete subsection reference.

<a id="canonical-0023311121131322-3210023031020001-1011201132313100-0312212132010010-2123131221122001-3120323223130033-1332310300323022-2000232301231002"></a>

<a id="canonical-2111233303213332-0210013031010102-2001021110330333-1123013121002321-1123332323330213-0101121330231121-2001203202101200-0030331033212212"></a>

#### `http_proxy.more_option.response_cookies_to_add.add_path` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

- [add_secure](resources--proxy--reference--group-004.md#canonical-1120113031102033-3311223300132202-0003111130112133-1031302321113210-2001313331103213-3022322310000020-3321122301112321-3112111301012012): complete subsection reference.

- [ignore_domain](resources--proxy--reference--group-004.md#canonical-0230111310231220-0022230232320033-2001112110322132-0021320021301300-3012120001021120-3231132113231133-2000102033301312-0220221221123323): complete subsection reference.

- [ignore_expiry](resources--proxy--reference--group-004.md#canonical-2033222002030110-3233202113213202-2113331203022231-0320121110012031-2332232223012113-0203021121211312-3210003103210011-1032211210022333): complete subsection reference.

- [ignore_httponly](resources--proxy--reference--group-004.md#canonical-0320001110232130-3312320130031211-2122113003123303-2212001322310122-3201312232320101-2302221233320111-1023331312003220-0013033130300121): complete subsection reference.

- [ignore_max_age](resources--proxy--reference--group-004.md#canonical-3020333013311201-2021311030003000-1033033230122330-2011120333313213-0003132311113332-1033032010010130-3210101132310300-1100233001021330): complete subsection reference.

- [ignore_partitioned](resources--proxy--reference--group-004.md#canonical-0123111011200201-0302223113101031-3212002000133333-0111011212322021-1103212133121322-0212013322333033-0230213123202330-0121300001003122): complete subsection reference.

- [ignore_path](resources--proxy--reference--group-004.md#canonical-0321323213312032-0011221302101101-3201100322103203-2132110001122333-1103210003202103-2321100301321001-3300312221201033-3322021330302231): complete subsection reference.

- [ignore_samesite](resources--proxy--reference--group-004.md#canonical-2111200001020121-3110202000221230-0023303100300132-1122030032011310-1021120313102101-2102302300232131-3333011132032220-3212120331330203): complete subsection reference.

- [ignore_secure](resources--proxy--reference--group-004.md#canonical-0032120130021302-0201013033000220-0202332223110330-2031013200331123-1100100021332110-3133030202102031-2213210331220022-3030222301221230): complete subsection reference.

- [ignore_value](resources--proxy--reference--group-004.md#canonical-1031222023100232-3133320301231330-0312312111112012-1010303031233111-3003200000002033-0020112101021130-1221203331012312-1120103113021320): complete subsection reference.

<a id="canonical-2310032332322001-3220200203003332-1022010100011112-3220231322321131-1212130200012222-1000130121302203-3332313100123320-1002111212310323"></a>

<a id="canonical-0323302121323131-3332203201220321-0210022201123310-1212112311223332-0031221021112022-1122220323022000-1130322112313112-3233331030030121"></a>

#### `http_proxy.more_option.response_cookies_to_add.max_age_value` property

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-2032122312200131-2111332013212203-1021331031300313-2313311211103332-3302230322232122-3022012230002121-0331330113000212-1233110320112012"></a>

<a id="canonical-2213321020012310-2202320131201032-0110123231300121-1303321113310021-1010033232011100-3000100332301231-3201323012303131-3300031012002121"></a>

#### `http_proxy.more_option.response_cookies_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3123302332311112-0231212123321032-3300030101300110-0210112033110321-1130233333233013-3210230323301131-3320313300313120-2012330330210030"></a>

<a id="canonical-3132223110202331-0203123120003020-3000231313122320-2312120121001200-0020313231301310-0312303112021122-3311233022213022-1110201100022000"></a>

#### `http_proxy.more_option.response_cookies_to_add.overwrite` property

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

- [samesite_lax](resources--proxy--reference--group-004.md#canonical-1212103011120333-0223001021231133-1321303222211111-2100302322210312-0221333311012000-3223203131002122-2332032200121331-0211031130111321): complete subsection reference.

- [samesite_none](resources--proxy--reference--group-004.md#canonical-3132332212330322-0202310233313222-2003022020022302-3113230213300121-0200202011212113-1020310103120000-1323101010233002-3312222223212020): complete subsection reference.

- [samesite_strict](resources--proxy--reference--group-004.md#canonical-2222032212220010-0332200103321102-2131011023001102-3310033321132213-0103322202210010-0220312231211301-3310330222200123-3320210000213321): complete subsection reference.

- [secret_value](resources--proxy--reference--group-004.md#canonical-0032223220331321-2122103333322322-0112302321320233-1222120311123120-3122320100100323-2000101331022231-0320013223311032-0110033301203222): complete subsection reference.

<a id="canonical-2313220302022013-3032113230332112-3111022103330310-1100300002332301-1102033220111211-1020122231321022-0100032223213231-0133023300113320"></a>

<a id="canonical-2230021032221100-2121022330332122-2333103101100003-1110110020231110-3311203033302323-2032221100320320-1121130203232021-3021122032303033"></a>

#### `http_proxy.more_option.response_cookies_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-0322013202332010-2030231103122231-3231303232203303-2032203302113112-2031030000332031-3132330213131202-0013233332130021-1131221202111120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.add_httponly` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.add_httponly

<a id="canonical-2300221132220032-1212303011022123-0102011133303233-2212221022123222-0003001022000033-2000130100210121-1022013010231130-2213221221301330"></a>

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

<a id="canonical-3120333300200020-3220011211122220-0210232103101022-0111313302133333-2033110320012003-3203013202211011-3003222221201211-1031032310222321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.add_partitioned` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.add_partitioned

<a id="canonical-0000220000131020-0112002310222230-1312332010332001-3220230022310100-2332132203133132-1012330313310301-0311322023311130-0303301332131221"></a>

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

<a id="canonical-1120113031102033-3311223300132202-0003111130112133-1031302321113210-2001313331103213-3022322310000020-3321122301112321-3112111301012012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.add_secure` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.add_secure

<a id="canonical-1333133032032232-2020023020031330-2002000023320230-1303000230212332-0120301210200303-1121000321113200-1311023211220102-1231131312230222"></a>

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

<a id="canonical-0230111310231220-0022230232320033-2001112110322132-0021320021301300-3012120001021120-3231132113231133-2000102033301312-0220221221123323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_domain` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_domain

<a id="canonical-0221201230023312-0221323211011323-3332120332233110-2332022022123031-0122021132331132-2111021020131301-1102230013332100-0311030103022130"></a>

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

<a id="canonical-2033222002030110-3233202113213202-2113331203022231-0320121110012031-2332232223012113-0203021121211312-3210003103210011-1032211210022333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_expiry` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-3301221122213231-1331033131130203-1323002020111003-3031230221203231-0311102313120022-2001212033133011-2000121011022112-1133021311323230"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_expiry = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320001110232130-3312320130031211-2122113003123303-2212001322310122-3201312232320101-2302221233320111-1023331312003220-0013033130300121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_httponly` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-3021201202113032-3001303132200023-1000220310322101-0031023123032102-3231233121232202-2333023031120023-3120213203232020-2303132202332332"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020333013311201-2021311030003000-1033033230122330-2011120333313213-0003132311113332-1033032010010130-3210101132310300-1100233001021330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_max_age` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-1220101022102300-1003133022313203-2003121220233210-0212233101212032-1221100331213001-1131010230211301-3033132101002103-3022323303022301"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_max_age = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123111011200201-0302223113101031-3212002000133333-0111011212322021-1103212133121322-0212013322333033-0230213123202330-0121300001003122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_partitioned` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-2100101233323032-1130213211301030-0001300120122332-1310030330221022-2220012210311133-1330210133022130-0231030201313003-3000211013303303"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_partitioned = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321323213312032-0011221302101101-3201100322103203-2132110001122333-1103210003202103-2321100301321001-3300312221201033-3322021330302231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_path` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_path

<a id="canonical-3130232232002203-0101102011220030-0223203312101321-2133223023012313-1201330320002233-3231221133000232-2231230233021000-0220222132200313"></a>

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
ignore_path = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111200001020121-3110202000221230-0023303100300132-1122030032011310-1021120313102101-2102302300232131-3333011132032220-3212120331330203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_samesite` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-3133123012001302-3003231110111001-2013202131123130-2110021011013110-0131123300023112-1131230123222203-2013312012310022-3223111003033310"></a>

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
ignore_samesite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032120130021302-0201013033000220-0202332223110330-2031013200331123-1100100021332110-3133030202102031-2213210331220022-3030222301221230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_secure` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_secure

<a id="canonical-1233333310313010-1110313202123232-3022121322303321-1321330103210013-2200223232312120-1000103133110022-3123132130001131-3210030332122332"></a>

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
ignore_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031222023100232-3133320301231330-0312312111112012-1010303031233111-3003200000002033-0020112101021130-1221203331012312-1120103113021320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_value` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_value

<a id="canonical-2000011001011323-1300131221132110-3303122103213333-1011011021303021-2012212223232201-0003232211233331-1132133013322310-0020230200220121"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_value = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212103011120333-0223001021231133-1321303222211111-2100302322210312-0221333311012000-3223203131002122-2332032200121331-0211031130111321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.samesite_lax` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.samesite_lax

<a id="canonical-0323000230312231-2003211031102031-1103001100233000-1202311101333212-3020232303303111-2003230311033002-2012113111100201-2122332202210211"></a>

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
samesite_lax = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132332212330322-0202310233313222-2003022020022302-3113230213300121-0200202011212113-1020310103120000-1323101010233002-3312222223212020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.samesite_none` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.samesite_none

<a id="canonical-0333030220112210-2021023111232103-0122120132213233-2033202133101020-1110312213101233-3331210333132111-3133213211111101-0130332211011103"></a>

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
samesite_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222032212220010-0332200103321102-2131011023001102-3310033321132213-0103322202210010-0220312231211301-3310330222200123-3320210000213321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.samesite_strict` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.samesite_strict

<a id="canonical-1111222130212203-0131223230132311-3010311232233220-3111231010232322-3301221010120130-3211003012131223-0303220300123312-0302102210001222"></a>

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
samesite_strict = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032223220331321-2122103333322322-0112302321320233-1222120311123120-3122320100100323-2000101331022231-0320013223311032-0110033301203222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.secret_value

<a id="canonical-1300211320131301-0203023131301210-1303101312013001-1112303003212202-0120331100100003-2001013102011323-2100133111233203-3101011302130220"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2001100202131202-2322300211233123-2332223313213312-3202323030203030-3132213010023203-0102002233302011-3201102123210101-0223233200002201"></a>

### Direct properties for `http_proxy.more_option.response_cookies_to_add.secret_value`

- [blindfold_secret_info](resources--proxy--reference--group-005.md#canonical-1210113233221230-1203323122300112-2012110322330310-2123210112123301-2110120032232111-0200223332121310-3122221101333302-1132231101001031): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-005.md#canonical-2111320201010202-1001012323322201-0233312100110032-2123030212002221-3332123331003122-0122100302330333-2233101000313130-2011200001003302): complete subsection reference.
