---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- Property reference

<a id="canonical-1002220000132112-3011220112112231-1122122333133112-1113130200001130-3110310333230231-0320011201002030-2330002212000113-3000310011313132"></a>

### Direct properties for `xcsh_virtual_host`

<a id="canonical-0013323330322123-3300201000211210-3132102322303022-3332302201101122-0222110233301011-1223323133333310-2102212100121313-0313003110112032"></a>

#### `add_location` property

Type: `"bool"`. Optional, Computed.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses. This configuration is ignored on CE sites.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [advertise_policies](resources--virtual_host--reference--group-001.md#canonical-3101233030000031-0230120000021132-0100111102033122-2131330301131003-3313323011111033-1133231010010122-2231321131221000-2223022121212113): complete subsection reference.

<a id="canonical-2110200310323201-2210212033231133-1111022333112202-3033102201230032-2202310030031320-1131121321033313-3223203220233002-3003132201313211"></a>

<a id="canonical-2302010110022001-2321103311203312-1210121312103131-1020032002201200-0033223003221200-3220001030331220-1111321000220321-0321001221012111"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-2032132033332212-2303220021231303-0323013021013231-2113200332010333-0210210201220233-1300002321020233-0003012010320111-0122302230222333"></a>

<a id="canonical-1020200012133100-3200301102300123-2033320220010011-1122220200010322-1000233333011221-1023330203111000-1220103213213121-3010122110100230"></a>

#### `append_server_name` property

Type: `"string"`. Optional, Computed.

\[OneOf: append\_server\_name, default\_header, pass\_through, server\_name; Default:
default\_header\] Exclusive with \[default\_header pass\_through server\_name\] Specifies the value
to be used for Server header if it is not already present. If Server Header is already present it is
not overwritten. It is just passed.

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

OneOf alternatives in this subsection:

- [append_server_name](resources--virtual_host--reference--group-001.md#canonical-2032132033332212-2303220021231303-0323013021013231-2113200332010333-0210210201220233-1300002321020233-0003012010320111-0122302230222333)
- [default_header](resources--virtual_host--reference--group-002.md#canonical-3103110113233101-2311030300030002-3133031023300010-1132231000122223-1300222101112311-3130112321310012-0333210202330121-0301221023322011)
- [pass_through](resources--virtual_host--reference--group-002.md#canonical-0132112102332232-0123332032222330-0301020323101110-3123022022131303-1232332032232110-0303301323030131-0113120332013210-1113310021333202)
- [server_name](resources--virtual_host--reference--group-001.md#canonical-3202022102332232-2133012230301013-1001202130203031-0012102010333000-0231302310211120-1101011033112211-3303110301000033-1300011110102320)

Select alternatives according to the provider validators above.

- [authentication](resources--virtual_host--reference--group-001.md#canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123): complete subsection reference.

- [buffer_policy](resources--virtual_host--reference--group-001.md#canonical-3232012230212003-3221232011020313-0202200302301332-0002323023102302-2322221023103201-3002223222331300-0233232032310121-2021012312213213): complete subsection reference.

- [captcha_challenge](resources--virtual_host--reference--group-001.md#canonical-1020312231213113-0223120010100113-3031300332310212-2210330020033311-2212131302230022-1310032130223312-2033030030001301-2312002230010203): complete subsection reference.

- [coalescing_options](resources--virtual_host--reference--group-001.md#canonical-1133103011022022-2203323101231031-2333022101312310-0311232221002132-1112210132302130-3101211220023233-3011302023021202-3033100231101311): complete subsection reference.

- [compression_params](resources--virtual_host--reference--group-001.md#canonical-1033033121221113-3201013233022330-2031033001101131-2231201001233133-3212322221010320-2231221131030022-1003330210110233-1333033203013221): complete subsection reference.

<a id="canonical-3032310110300222-2000212021201213-0223203311330221-0020101110013232-3303120030031001-2010133132130000-1323033031102122-3112201220102032"></a>

<a id="canonical-1312313012202212-1211222323103311-3303311233120220-0213013013203132-2301121211313233-2230303312113012-3313120213212132-2111201220202011"></a>

#### `connection_idle_timeout` property

Type: `"number"`. Optional, Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

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

- [cors_policy](resources--virtual_host--reference--group-001.md#canonical-3232131223002131-1100030322001210-3211031031233033-0203330020202031-0110112203211203-1312221320131310-3120133001121223-0213203230311301): complete subsection reference.

- [csrf_policy](resources--virtual_host--reference--group-001.md#canonical-0032333010121002-0332331033300212-2002033020321220-2332202013303113-2313133231100331-0003222333003100-3133200232230203-0121213102202020): complete subsection reference.

<a id="canonical-0311101123310131-0113032000322132-2330012100310121-1000103023313213-3220131033131221-2323000031310313-1313130231210122-0131211113000313"></a>

<a id="canonical-3120121031023322-1323300302221100-1020301220103133-1232323332323233-1213322220300032-2213323003022010-2101211121202303-3211023313232330"></a>

#### `custom_errors` property

Type: `["map", "string"]`. Optional.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value is the URI\_ref. Currently supported URL schemes
is string:///. For string:/// scheme, message needs to be encoded in base64 format. You can specify
this message as base64 encoded plain text message e.g. "Access Denied" or it can be HTML paragraph
or a body string encoded as base64 string E.g. "&lt;p&gt; Access Denied &lt;/p&gt;". base64 encoded
string for this HTML is "PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==" Specific response code takes preference
when both response code and response code class matches for a request.

The configured custom errors are only applicable for loadbalancer generated errors. Errors returned
from upstream server is propagated as is.

F5XC provides default error pages for the errors generated by the loadbalancer. Content of these
pages are not editable. User has an option to disable the use of default F5XC error pages.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{
  validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maximum\":599,\"minimum\":3,\"type\":\"uint32-string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.uint32.gte\":\"3\",\"ves.io.schema.rules.map.keys.uint32.lte\":\"599\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"65536\",\"ves.io.schema.rules.map.values.string.uri_ref\":\"true\"},\"values\":{\"format\":\"uri-reference\",\"maxLength\":65536,\"type\":\"string\"}}"),
}
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
      "maximum": 599,
      "minimum": 3,
      "type": "uint32-string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.uint32.gte": "3",
      "ves.io.schema.rules.map.keys.uint32.lte": "599",
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
    "ves.io.schema.rules.map.keys.uint32.gte": "3",
    "ves.io.schema.rules.map.keys.uint32.lte": "599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.gte": "3",
    "ves.io.schema.rules.map.keys.uint32.lte": "599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  }
}
```

- [default_header](resources--virtual_host--reference--group-002.md#canonical-1332133000210111-0223310312103001-1010233010211033-1102013020202323-2130113222000300-0232202222221101-2023022020231220-2012203012032313): complete subsection reference.

- [default_loadbalancer](resources--virtual_host--reference--group-002.md#canonical-3121022200033310-1220301130221312-2021030331023200-1312300232120133-1231101321130312-2222300001330221-1303311113210322-3100130111300312): complete subsection reference.

<a id="canonical-0030002233332331-3231212021013021-3012013122113301-2221302213223021-0230122310220033-0031202333020303-1130023120032301-1201010010332021"></a>

<a id="canonical-0322212321303232-1100031020002212-2200221013120111-3332301120310210-3022002213013113-3210311020123312-3121223231303120-3322222000003023"></a>

#### `description` property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-2003222121102200-3023011233323121-0102220130201310-3000330003100122-2203301130331101-3303230221312121-3312020121123203-1110212000202232"></a>

<a id="canonical-2112001113330321-2211000020333313-3001032303210020-1333022121100202-1112113132322130-0211231112302333-3330131013300030-2102123311003131"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1001322231001000-0232313010110102-0202220213223333-0220001233312033-3123230200011123-2301021113221020-1330331101010221-2020113001111310"></a>

<a id="canonical-2133010103321221-3101123032031230-3210230333313231-2010001221031112-1033103210011110-0211202321123223-3031203113212102-3320002300211010"></a>

#### `disable_default_error_pages` property

Type: `"bool"`. Optional, Computed.

An option to specify whether to disable using default F5XC error pages.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0013130302202231-2302310310120022-1101121131101321-3101012121213221-0311233303300201-3230202310203023-0012233011211200-3132021120331021"></a>

<a id="canonical-1002101213230303-2330331312203212-1202330021331033-3100200220331111-1110123302030003-1023112003322030-3010002302222313-2333323331303122"></a>

#### `disable_dns_resolve` property

Type: `"bool"`. Optional, Computed.

Disable DNS resolution for domains specified in the virtual host When the virtual host is configured
as Dynamive Resolve Proxy (DRP), disable DNS resolution for domains configured. This configuration
is suitable for HTTP CONNECT proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_path_normalize](resources--virtual_host--reference--group-002.md#canonical-0002023231210102-0232311230110330-2231021201011211-3000222000033100-3121230313301113-1303121121223033-1322331301232132-0220030301200000): complete subsection reference.

<a id="canonical-3212000311330033-3332230300113221-1321030133310313-2310133120230230-2113011130011110-0212031123103320-2300100230130123-2000321302332032"></a>

<a id="canonical-1012132011323212-1322112030221133-2113021330202222-1103320220033103-0233200002220010-3001221131233100-0123232031231020-0033102131100223"></a>

#### `domains` property

Type: `["list", "string"]`. Optional.

List of domain names matched to this virtual host for routing incoming requests. Supports wildcard
patterns like \*.example.com for subdomain matching.

Additional upstream details:

A list of Domains (host/authority header) that will be matched to this Virtual Host. Wildcard hosts
are supported in the suffix or prefix form

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com. Not supported Domains: &#8203;- Just a Wildcard:
\* &#8203;- A Wildcard and TLD with no root Domain: \*.com. &#8203;- A Wildcard not matching a whole
DNS label. E.g. \*.example.com and \*.bar.example.com are valid Wildcards however \*bar.example.com,
\*-bar.example.com, and bar\*.example.com are all invalid. Additional notes: A Wildcard will not
match empty string. E.g. \*.example.com will match bar.example.com and baz-bar.example.com but not
.example.com. The longest Wildcards match first. Only a single virtual host in the entire route
configuration can match on \*. Also a Domain must be unique across all virtual hosts within an
advertise policy. Domains are also used for SNI matching if the virtual host proxy type is
TCP\_PROXY\_WITH\_SNI/HTTPS\_PROXY Domains also indicate the list of names for which DNS resolution
will be automatically resolved to IP addresses by the system.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 33),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 33,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 33,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "33",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "33",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [dynamic_reverse_proxy](resources--virtual_host--reference--group-002.md#canonical-2133031213301110-2320211003030001-3300321001112212-0020212013122303-0311232110311233-2131100131212002-2313301132131310-2201010033313212): complete subsection reference.

- [enable_path_normalize](resources--virtual_host--reference--group-002.md#canonical-1222113320121333-2000121121321100-2112033320330111-1023323321010321-0122103013112310-0102102320310031-1303320233230330-2311203122001130): complete subsection reference.

- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-0212202230011113-2203132220233101-1321131330330010-2030212021032322-3133023002222200-2133232330210133-3312230332103333-3123111320010133): complete subsection reference.

<a id="canonical-2100122103113131-1300210011133013-2332123120112031-0211300012113033-1002303312231332-2211231310201032-0320120311000222-2031132103332113"></a>

<a id="canonical-0102000320303013-0012211101313110-2020121010313130-1213301221321031-3123131223111121-3121221202331230-3000101002022313-1001122213121112"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2130311100021022-0231313123020200-2011301202100101-0212022023311022-2110321111203231-3133303221213013-1220001320321030-3112232011003003"></a>

<a id="canonical-3123121030222200-1010323220121201-0033302111333010-2231330233022202-0323222021123223-0120320331102001-3312200300220112-2301332303122012"></a>

#### `idle_timeout` property

Type: `"number"`. Optional, Computed.

Idle timeout is the amount of time that the loadbalancer will allow a stream to exist with no
upstream or downstream activity.

Idle timeout and Proxy Type:

HTTP\_PROXY, HTTPS\_PROXY: Idle timer is started when the first byte is received on the connection.
Each time an encode/decode event for headers or data is processed for the stream, the timer will be
reset. If the timeout fires, the stream is terminated with a 504 (Gateway Timeout) error code if no
upstream response header has been received, otherwise a stream reset occurs. The default idle
timeout is 30 seconds

TCP PROXY, TCP\_PROXY\_WITH\_SNI, SMA\_PROXY: The idle timeout is defined as the period in which
there are no bytes sent or received on either the upstream or downstream connection. The default
idle timeout is 1 hour.

UDP PROXY: The idle timeout for sessions. Idle timeout is defined as the period in which there are
no datagrams sent or received on the session. The default if not specified is 1 minute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [js_challenge](resources--virtual_host--reference--group-002.md#canonical-0101231302233001-2011223111023233-0001132321323232-0221122312210111-3200231102333020-3311221202021012-2021103121020012-0121132032203123): complete subsection reference.

<a id="canonical-0030111302020011-3233132301130311-2223020332120120-0022233112232323-1103020111123112-2103100113321030-2003122112003113-0203121201121120"></a>

<a id="canonical-2020311031202220-2030012221231011-1111001321300012-2232200221331200-3211230220233310-2002311320031331-0221303111202110-1322212203231321"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2301221233022002-1033000200231312-3302301110130013-2320013313321302-2103022203021132-3322332032100230-0300002000030230-1311310010111130"></a>

<a id="canonical-3212020200131301-2122333102020320-1133102031203132-1330133131003300-2221012011220311-1022002213303321-3103230103212031-3101113000200311"></a>

#### `max_request_header_size` property

Type: `"number"`. Optional, Computed.

The maximum request header size in KiB for incoming connections.

If un-configured, the default max request headers allowed is 60 KiB.

Requests that exceed this limit will receive a 431 response.

The max configurable limit is 96 KiB, based on current implementation constraints.

Note: a. This configuration parameter is applicable only for HTTP\_PROXY and HTTPS\_PROXY b. When
multiple HTTP\_PROXY virtual hosts share the same advertise policy, the effective "maximum request
header size" for such virtual hosts is the highest value configured on any of the virtual hosts.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3302203010032022-0312123123321211-3211010021230223-3131033301001120-0312112002101330-3323100120222113-1112211233210202-3301310201101123"></a>

<a id="canonical-0201321321330300-1003310102000300-2321023310203300-1310321020310012-1000032222303102-2001001302113032-1201221011020033-1201312110001000"></a>

#### `max_requests_per_connection` property

Type: `"number"`. Optional, Computed.

\[OneOf: max\_requests\_per\_connection, no\_request\_limit\_per\_connection; Default:
no\_request\_limit\_per\_connection\] Exclusive with \[no\_request\_limit\_per\_connection\] Sets
the maximum number of requests a downstream client can send over a single connection to Envoy. Enter
a value &gt;=1 to define the request limit per connection.

Provider validators and defaults (from schema source):

```go
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

OneOf alternatives in this subsection:

- [max_requests_per_connection](resources--virtual_host--reference--group-001.md#canonical-3302203010032022-0312123123321211-3211010021230223-3131033301001120-0312112002101330-3323100120222113-1112211233210202-3301310201101123)
- [no_request_limit_per_connection](resources--virtual_host--reference--group-002.md#canonical-1220233000321320-2201210110230233-3022023021301313-3131222033332030-3021323101221112-2320200132210013-0112101301202311-0300202201122113)

Select alternatives according to the provider validators above.

<a id="canonical-3232022013210020-0331232120212123-1221211300320201-3022330213113112-1022203033321213-1322213313333121-1220232230301012-2101322111230011"></a>

<a id="canonical-2000012102313322-3013132303231031-2000011101202103-1302211022232210-3303222022330003-2232321202300312-3231131020333020-2312013011333103"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Virtual Host. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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

<a id="canonical-2323100200332233-2011102133130001-1222113033131122-2120201122213103-0313232011233313-0101303300103130-2313002000332200-1312121100313031"></a>

<a id="canonical-2130020102023213-1201310220022233-2011221203031101-2003202022012111-0200231300012132-1130131310110213-1010202232313130-1300330332133001"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Virtual Host is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [no_authentication](resources--virtual_host--reference--group-002.md#canonical-3013320031330300-3111311312212033-2213131023230113-1110310112213020-3030011311320203-2011102330101110-1110113223010031-3121032131222123): complete subsection reference.

- [no_challenge](resources--virtual_host--reference--group-002.md#canonical-0032230233200032-3113230021023123-1002001030322103-2001220130000333-3010022102100323-0113113322202322-2132212011203032-2132231103032230): complete subsection reference.

- [no_request_limit_per_connection](resources--virtual_host--reference--group-002.md#canonical-1201021232211320-1113031321300211-0230223223322233-2013130232110130-1300111221111201-1102111321313123-0013222103000031-1210301000110230): complete subsection reference.

- [non_default_loadbalancer](resources--virtual_host--reference--group-002.md#canonical-2103112131222102-3010030322132022-3211222333113012-3301232022132313-0320122322331301-1020131023201302-2322122130131023-1210231202212021): complete subsection reference.

- [pass_through](resources--virtual_host--reference--group-002.md#canonical-0323002221202331-0013320123000331-2220112010231031-3313320111013331-3201130132130201-2323102121111212-0000301210022202-0212030313323103): complete subsection reference.

<a id="canonical-0221100123330010-0213131020311100-0132311121102322-1031202000000222-2303221011223212-3232320210123212-3231022110202221-3223312233230221"></a>

<a id="canonical-1312112203232121-2120001212203032-0311322031023111-1110013212120023-0131312010103123-2211213022333011-1220312113230302-0120111311033333"></a>

#### `proxy` property

Type: `"string"`. Optional, Computed.

\[Enum:
UDP\_PROXY|SMA\_PROXY|DNS\_PROXY|ZTNA\_PROXY|UZTNA\_PROXY|TMM\_HTTP\_PROXY|TMM\_HTTPS\_PROXY|TMM\_TCP\_PROXY|TMM\_UDP\_PROXY|TMM\_QUIC\_PROXY\]
ProxyType tells the type of proxy to install for the virtual host. Only the following combination of
VirtualHosts within same AdvertisePolicy is permitted (None of them should have '\*' in domains when
used with other VirtualHosts in same AdvertisePolicy) 1. Multiple TCP\_PROXY\_WITH\_SNI and..
Possible values are \`UDP\_PROXY\`, \`SMA\_PROXY\`, \`DNS\_PROXY\`, \`ZTNA\_PROXY\`,
\`UZTNA\_PROXY\`, \`TMM\_HTTP\_PROXY\`, \`TMM\_HTTPS\_PROXY\`, \`TMM\_TCP\_PROXY\`,
\`TMM\_UDP\_PROXY\`, \`TMM\_QUIC\_PROXY\`.

Additional upstream details:

ProxyType tells the type of proxy to install for the virtual host. Only the following combination of
VirtualHosts within same AdvertisePolicy is permitted (None of them should have "\*" in domains when
used with other VirtualHosts in same AdvertisePolicy) &#8203;1. Multiple TCP\_PROXY\_WITH\_SNI and
multiple HTTPS\_PROXY &#8203;2. Multiple HTTP\_PROXY &#8203;3. Multiple HTTPS\_PROXY &#8203;4.
Multiple TCP\_PROXY\_WITH\_SNI

HTTPS\_PROXY without TLS parameters is not permitted
HTTP\_PROXY/HTTPS\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY with empty domains is not permitted
TCP\_PROXY\_WITH\_SNI/SMA\_PROXY should not have "\*" in domains

&#8203;- HTTP\_PROXY: HTTP\_PROXY

Install HTTP proxy. HTTP Proxy is the default proxy installed. &#8203;- TCP\_PROXY: TCP\_PROXY

Install TCP proxy &#8203;- TCP\_PROXY\_WITH\_SNI: TCP\_PROXY\_WITH\_SNI

Install TCP proxy with SNI Routing &#8203;- TLS\_TCP\_PROXY: TCP\_PROXY

Install TCP proxy &#8203;- TLS\_TCP\_PROXY\_WITH\_SNI: TCP\_PROXY\_WITH\_SNI

Install TCP proxy with SNI Routing &#8203;- HTTPS\_PROXY: HTTPS\_PROXY

Install HTTPS proxy &#8203;- UDP\_PROXY: UDP\_PROXY

Install UDP proxy &#8203;- SMA\_PROXY: SMA\_PROXY

Install Secret Management Access proxy &#8203;- DNS\_PROXY: DNS\_PROXY

Install DNS proxy &#8203;- ZTNA\_PROXY: ZTNA\_PROXY

Install ZTNA proxy.this is going to be deprecated with UZTNA\_PROXY. &#8203;- UZTNA\_PROXY:
UZTNA\_PROXY

Install UZTNA proxy &#8203;- TMM\_HTTP\_PROXY: TMM\_HTTP\_PROXY

Install TMM HTTP proxy for HTTP/1.1 and HTTP/2 traffic. Used by TMM proxy type. &#8203;-
TMM\_HTTPS\_PROXY: TMM\_HTTPS\_PROXY

Install TMM HTTPS proxy for HTTP/1.1 and HTTP/2 traffic. Used by TMM proxy type. &#8203;-
TMM\_TCP\_PROXY: TMM\_TCP\_PROXY

Install TMM TCP proxy for TCP traffic. Used by TMM proxy type. &#8203;- TMM\_UDP\_PROXY:
TMM\_UDP\_PROXY

Install TMM UDP proxy for UDP traffic. Used by TMM proxy type. &#8203;- TMM\_QUIC\_PROXY:
TMM\_QUIC\_PROXY

Install TMM QUIC proxy for HTTP/3 traffic. Used by TMM proxy type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UDP_PROXY",
    "SMA_PROXY",
    "DNS_PROXY",
    "ZTNA_PROXY",
    "UZTNA_PROXY",
    "TMM_HTTP_PROXY",
    "TMM_HTTPS_PROXY",
    "TMM_TCP_PROXY",
    "TMM_UDP_PROXY",
    "TMM_QUIC_PROXY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "UDP_PROXY",
    "SMA_PROXY",
    "DNS_PROXY",
    "ZTNA_PROXY",
    "UZTNA_PROXY",
    "TMM_HTTP_PROXY",
    "TMM_HTTPS_PROXY",
    "TMM_TCP_PROXY",
    "TMM_UDP_PROXY",
    "TMM_QUIC_PROXY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [rate_limiter_allowed_prefixes](resources--virtual_host--reference--group-002.md#canonical-1012012232023120-0032122201003333-2110302001323221-0232313320231102-2031333000031122-0111010022332321-0120313223113103-2321021201002203): complete subsection reference.

- [request_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-0111131022031311-0102023300100223-2120311331022200-0220230223131102-3320022000312001-1021212300331101-1201203011120311-1323012103110312): complete subsection reference.

<a id="canonical-1133330212300212-3331103111021021-2222211121303313-1132022022331223-0012010222330130-1323313333101330-3122002312312112-2222110220100202"></a>

<a id="canonical-3011001012013312-2323121031031213-1331031332103210-2331003211100233-0023223002002331-1330233232122013-0313303231100000-3300010321332133"></a>

#### `request_cookies_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
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

- [request_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-3023100112011132-2302030100320013-2030100013010220-0112333330230033-0110121200123102-2012130203132311-0221233210230001-2212303220211221): complete subsection reference.

<a id="canonical-1100302332102030-2323230022003003-1330220321313332-2322002113222021-0002102022112000-1201023131232203-0031030110230320-2100220033303302"></a>

<a id="canonical-3022112030122201-3301310122001213-3110023232323201-1232233311323023-2230023202321132-0312331031322130-0021223302300003-2031202311100302"></a>

#### `request_headers_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313): complete subsection reference.

<a id="canonical-0021022033103122-1202330200232333-3201001113021130-0212302323332013-0231120333031230-0200300003332213-2122022002232013-1132203333330221"></a>

<a id="canonical-0202003132003133-2231123130030003-3231302203032212-2312303002323200-1230002122032112-3331301212211023-1311113302101310-3121322321233300"></a>

#### `response_cookies_to_remove` property

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Provider validators and defaults (from schema source):

```go
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

- [response_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-2031023100033202-2330100210303322-3113323303313133-2030111112021212-1302210011133231-0221213311321333-0210033213203311-0133030211030022): complete subsection reference.

<a id="canonical-2333020201013203-3202312230231313-0302112333030221-1020023111131203-1102321222121331-3100222012333100-0123033103232131-1330302200313013"></a>

<a id="canonical-2020113303113130-3032021331200211-0032323221122101-1312330010101032-1321132000331220-0231300003010130-3221133130123103-2333331302133220"></a>

#### `response_headers_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [retry_policy](resources--virtual_host--reference--group-002.md#canonical-3120210313011012-3220120012303030-2320221131003020-2003111301323011-0131032003311130-3311130033201323-1332312222032333-0100213232033200): complete subsection reference.

- [routes](resources--virtual_host--reference--group-002.md#canonical-1230120312310211-1032010120013201-1323220112101321-0101321330220132-1310202220130210-3031233111323031-2123332330332300-2030220320303223): complete subsection reference.

- [sensitive_data_policy](resources--virtual_host--reference--group-002.md#canonical-1113203023200020-3223013230231323-2212021220030131-0310020233201101-0121323110323203-3122333302133331-0211123332303330-1130022122020210): complete subsection reference.

<a id="canonical-3202022102332232-2133012230301013-1001202130203031-0012102010333000-0231302310211120-1101011033112211-3303110301000033-1300011110102320"></a>

<a id="canonical-1322211202330213-1211330210012311-1002122210303023-1003221001021113-3110131313111311-3210101022122132-2223332132232032-2300103022201302"></a>

#### `server_name` property

Type: `"string"`. Optional, Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Specifies the value to be used
for Server header inserted in responses. This will overwrite existing values if any for Server
Header.

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

- [slow_ddos_mitigation](resources--virtual_host--reference--group-002.md#canonical-2333312030122121-1100330320331032-1120232213211221-3131120330330231-2221003131100222-1230211011000221-1033321202131202-1012333211023301): complete subsection reference.

- [timeouts](resources--virtual_host--reference--group-002.md#canonical-2133311002322212-0233231012302121-3111302012013321-0301321213223223-0023120022022300-3003323023033302-0131213332220332-1033313111212310): complete subsection reference.

- [tls_cert_params](resources--virtual_host--reference--group-002.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331): complete subsection reference.

- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323): complete subsection reference.

- [user_identification](resources--virtual_host--reference--group-003.md#canonical-1012201010031002-2232101323213000-1022110300231232-0230000221021101-1203021223001011-2301313113020120-2002021332022003-0311230002232310): complete subsection reference.

- [waf_type](resources--virtual_host--reference--group-003.md#canonical-2322100101202101-2331131111122123-0200313013302211-3200032010130132-1310103000220023-2103022322122223-2311300301002030-1022212031023331): complete subsection reference.

<a id="canonical-0211231123110012-2220212020311322-2000110303212202-0100111102201022-0121201110212331-1021130332221030-3331112213023330-0332231121200131"></a>

### All schema paths for `xcsh_virtual_host`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `add_location` | [add_location](resources--virtual_host--reference--group-001.md#canonical-0013323330322123-3300201000211210-3132102322303022-3332302201101122-0222110233301011-1223323133333310-2102212100121313-0313003110112032) |
| `advertise_policies` | [advertise_policies](resources--virtual_host--reference--group-001.md#canonical-0320003333012113-1133301102110223-2031213123331300-3111030302111233-2110111202320311-3330220130131102-0130210011213222-0331122101131101) |
| `advertise_policies.kind` | [advertise_policies.kind](resources--virtual_host--reference--group-001.md#canonical-2032121022320021-1130023200022111-3012032332203210-3332000323131031-2212022222003131-1232333200210103-1230223313301002-3200210003130330) |
| `advertise_policies.name` | [advertise_policies.name](resources--virtual_host--reference--group-001.md#canonical-0033032333302201-1023211013311110-3121210233302220-1223102031110021-0310230013223321-3311121033101203-3122100312101311-0332210230230133) |
| `advertise_policies.namespace` | [advertise_policies.namespace](resources--virtual_host--reference--group-001.md#canonical-3333301100211030-2032331211212030-2130131202003121-2323323010310200-3323220110122011-3323010111230321-2101203331133210-1222010221303233) |
| `advertise_policies.tenant` | [advertise_policies.tenant](resources--virtual_host--reference--group-001.md#canonical-3121020321100200-1132203212213010-0222103313322323-0330303001022302-1023332230221213-2322032220021213-2133111232103103-0313112013201112) |
| `advertise_policies.uid` | [advertise_policies.uid](resources--virtual_host--reference--group-001.md#canonical-3303232113332101-0322313303002201-0123111220331110-2221132010211312-1133233110302002-0211110131101133-3202030300120330-1123022332210210) |
| `annotations` | [annotations](resources--virtual_host--reference--group-001.md#canonical-2110200310323201-2210212033231133-1111022333112202-3033102201230032-2202310030031320-1131121321033313-3223203220233002-3003132201313211) |
| `append_server_name` | [append_server_name](resources--virtual_host--reference--group-001.md#canonical-2032132033332212-2303220021231303-0323013021013231-2113200332010333-0210210201220233-1300002321020233-0003012010320111-0122302230222333) |
| `authentication` | [authentication](resources--virtual_host--reference--group-001.md#canonical-1211203032111231-3201133300013320-1000022233332122-2130221123031222-0121201132233320-3033233323321303-1211323003132102-1332111023313133) |
| `authentication.auth_config` | [authentication.auth_config](resources--virtual_host--reference--group-001.md#canonical-3222333202210002-1332002110001222-2321131000322220-0021330121013322-1012321113002202-1201022312333312-0301332332020211-0300331331113111) |
| `authentication.auth_config.kind` | [authentication.auth_config.kind](resources--virtual_host--reference--group-001.md#canonical-3233213030322030-0121020202133233-0101100210310322-0210203012223101-0311122012310122-2330303130323210-0333110300120000-0221130232133000) |
| `authentication.auth_config.name` | [authentication.auth_config.name](resources--virtual_host--reference--group-001.md#canonical-2210323323232133-0111330013012131-3203130303203230-1212100121113023-0233032210103000-3301111130123022-2011323020222211-2301023312103223) |
| `authentication.auth_config.namespace` | [authentication.auth_config.namespace](resources--virtual_host--reference--group-001.md#canonical-2332312331312121-2211210330311130-3012301003100222-2200110222100223-1221300331111011-0022100212130003-3330003202032322-3311233213112122) |
| `authentication.auth_config.tenant` | [authentication.auth_config.tenant](resources--virtual_host--reference--group-001.md#canonical-3110233221101300-1302130133021311-0330000101300320-0112320303202201-0333032332103102-2220321211213233-3131200222321323-1131121312012211) |
| `authentication.auth_config.uid` | [authentication.auth_config.uid](resources--virtual_host--reference--group-001.md#canonical-2132011022102102-1031310123330301-1012112212032323-0313031223000223-0121322221121120-1012001233012132-1302021001023032-2210033000323001) |
| `authentication.cookie_params` | [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-2200302311133111-0102132023022310-1301212120132130-3221133002312210-3300320202132332-2331020133223300-1003323112320120-1123220213023321) |
| `authentication.cookie_params.auth_hmac` | [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-1022310232123201-1222012330120313-0312323333230122-0222212201233030-3220103112103323-0312232123331002-3213312013312002-1220201303101120) |
| `authentication.cookie_params.auth_hmac.prim_key` | [authentication.cookie_params.auth_hmac.prim_key](resources--virtual_host--reference--group-001.md#canonical-0213112010330011-2030120013020000-0113322020302031-0331312000133020-2231003100331300-2231301101310120-2303212311030003-2020010132012332) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info](resources--virtual_host--reference--group-001.md#canonical-3010333000322033-1223023321231232-1131023033120132-1132211111223001-2110021030112022-1022000131130320-0020220010332133-3201203012112012) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider](resources--virtual_host--reference--group-001.md#canonical-2002332211202102-2102212131110320-2330101012213032-2132000311313130-3110201012211030-3322222200202223-0033231301201220-2320011032010323) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.location` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.location](resources--virtual_host--reference--group-001.md#canonical-0030220120311200-2331002111032301-2003200113132032-1022311232130320-1110330000203122-1111302100222331-3003200001333303-0020232302210131) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider](resources--virtual_host--reference--group-001.md#canonical-3113112131010121-1002000231210110-3200012001322231-1300013223330312-3321023210123311-1312101110130222-3223202200002110-2122321322111021) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info](resources--virtual_host--reference--group-001.md#canonical-1200122123003332-0032032231222311-0130120303312313-2231300032201233-2301102000021020-3322300323231233-3202321311221122-3230222321232010) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref](resources--virtual_host--reference--group-001.md#canonical-0301130100122013-1221203201203110-2013220110132012-3221312100022211-0223320222201322-3231013132223020-2113332020013000-2110023122133233) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.url` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.url](resources--virtual_host--reference--group-001.md#canonical-1110002222212002-2100020303301232-2211231002020130-1102231203300321-2202021023012330-1122111313232322-0323020023231101-0112032320213112) |
| `authentication.cookie_params.auth_hmac.prim_key_expiry` | [authentication.cookie_params.auth_hmac.prim_key_expiry](resources--virtual_host--reference--group-001.md#canonical-1320121102131231-2003222332232201-2332202010001321-2211201210210330-2221031031123032-3321001321210303-0120022213132213-1123213131020123) |
| `authentication.cookie_params.auth_hmac.sec_key` | [authentication.cookie_params.auth_hmac.sec_key](resources--virtual_host--reference--group-001.md#canonical-0221121003332310-2133111123021002-3111013321013131-3121223101011301-2213330003013103-3321112212021021-1223012220233231-2303212300031130) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info](resources--virtual_host--reference--group-001.md#canonical-3021101023312221-2333311213130231-3220231221130022-1012023321110310-0222313020200321-3301332022333032-2303232232131122-1121211332311203) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider](resources--virtual_host--reference--group-001.md#canonical-0200013301031002-3322001010032320-2233312022102303-2130303131110132-1310121302233302-3033232302013002-2322012101010000-3311313333200233) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.location` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.location](resources--virtual_host--reference--group-001.md#canonical-3311123302003210-2202021332312110-2030321000320112-0301311000131102-0302020131033330-1101121003020033-0130210101100001-0031232311300022) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider](resources--virtual_host--reference--group-001.md#canonical-1221203101000022-1201320132321031-0211013300202103-3233302332221333-3132302333220000-2323012302123103-0030031202311203-3031030121313330) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info](resources--virtual_host--reference--group-001.md#canonical-0130223210022330-3133123331332110-0132113210021321-1312032101032211-0300220210002203-3133021023110203-0130221322221322-3202013221133103) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref](resources--virtual_host--reference--group-001.md#canonical-0220310313202223-3022132133103130-2231000231300131-1200202032001211-3021113331210133-1020130013031112-3312130103303010-3312102123222221) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.url` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.url](resources--virtual_host--reference--group-001.md#canonical-1212011022331301-2032012333210310-1101301101203132-2231021320123300-0132012321133131-3232012311220100-1302033323220201-3103032313223311) |
| `authentication.cookie_params.auth_hmac.sec_key_expiry` | [authentication.cookie_params.auth_hmac.sec_key_expiry](resources--virtual_host--reference--group-001.md#canonical-3013313201001112-1332312332333222-2302121022112300-0333130100301133-0131121031321110-3033013003133222-1331303313311022-2220310230333302) |
| `authentication.cookie_params.cookie_expiry` | [authentication.cookie_params.cookie_expiry](resources--virtual_host--reference--group-001.md#canonical-3332332303002333-0000320313200310-0021133232332301-1202123301110113-3333110120102220-3002230121221301-0323010100133330-0033002310102333) |
| `authentication.cookie_params.cookie_refresh_interval` | [authentication.cookie_params.cookie_refresh_interval](resources--virtual_host--reference--group-001.md#canonical-1000232101101000-1030122230120331-3033120323031312-1122212323210023-1132131030113231-0300220222132003-0113213210032300-0032301330002303) |
| `authentication.cookie_params.kms_key_hmac` | [authentication.cookie_params.kms_key_hmac](resources--virtual_host--reference--group-001.md#canonical-1321333011213012-2012333300202103-3313300131313230-2031022333132100-1000323212331013-1121300333313313-0111213223021001-3230332330100111) |
| `authentication.cookie_params.session_expiry` | [authentication.cookie_params.session_expiry](resources--virtual_host--reference--group-001.md#canonical-1212232133020110-1303101001003022-3321020102020130-3232203001330210-0121103020311110-0033112022301121-2020210030003333-3333022203322200) |
| `authentication.redirect_dynamic` | [authentication.redirect_dynamic](resources--virtual_host--reference--group-001.md#canonical-3213102103320201-1032210310101101-1212233222123001-3120123321130021-1013220321200123-2200302302223310-3111023123021303-1320003001230120) |
| `authentication.redirect_url` | [authentication.redirect_url](resources--virtual_host--reference--group-001.md#canonical-2222102131030220-3310022313321300-3032123000123010-1213222111222003-1312131313032203-3211203102033213-3012020012200110-1031100112230312) |
| `authentication.use_auth_object_config` | [authentication.use_auth_object_config](resources--virtual_host--reference--group-001.md#canonical-0210222011020122-3312120333231130-0123030202102320-0131112000013232-1001111002202030-0031210011002221-3011120221110231-0230321030201102) |
| `buffer_policy` | [buffer_policy](resources--virtual_host--reference--group-001.md#canonical-3011223012232200-0033131112212212-0311032201323010-1230133111002012-1210000023123031-1120322321013011-3010101320312231-0021000020120100) |
| `buffer_policy.disabled` | [buffer_policy.disabled](resources--virtual_host--reference--group-001.md#canonical-1121202302021000-1122211120311322-0323030302130113-1212010331301331-0332310323322023-2231003333031010-2210031112120000-2102301030203023) |
| `buffer_policy.max_request_bytes` | [buffer_policy.max_request_bytes](resources--virtual_host--reference--group-001.md#canonical-1002323221030301-3223121130300123-0323212120223032-0131122231312320-2000021101211122-0031223322313323-1030133030303003-3203230222023012) |
| `captcha_challenge` | [captcha_challenge](resources--virtual_host--reference--group-001.md#canonical-1301312132130220-0131231003121113-3122120112223111-3123000301131222-1330033022112332-1210203020031031-3310220022023231-0120030021010211) |
| `captcha_challenge.cookie_expiry` | [captcha_challenge.cookie_expiry](resources--virtual_host--reference--group-001.md#canonical-0320001301322310-1312032101020021-0320302330321213-0330303223323021-1330102102330330-3220302130112121-1323332330102010-3323213100100103) |
| `captcha_challenge.custom_page` | [captcha_challenge.custom_page](resources--virtual_host--reference--group-001.md#canonical-3211020102201311-2010132132102223-2033110322132322-0313311233122022-3000211323330210-2013021010201102-2132120101020303-0111330130133121) |
| `coalescing_options` | [coalescing_options](resources--virtual_host--reference--group-001.md#canonical-3120121021011122-3011303310310313-1322223303233223-1001000133333132-3332113200330200-2022103003121300-0132333332101302-1122132031033100) |
| `coalescing_options.default_coalescing` | [coalescing_options.default_coalescing](resources--virtual_host--reference--group-001.md#canonical-2022310332331113-2010031031010312-3311022321032023-3113223221113202-2022330302223001-2112132123113031-3032213302301111-0231100312011013) |
| `coalescing_options.strict_coalescing` | [coalescing_options.strict_coalescing](resources--virtual_host--reference--group-001.md#canonical-2132210221002320-3201211031212203-3032331033323023-3333321331100021-2000323000001233-1230010301133312-2303310330203210-1113123032112131) |
| `compression_params` | [compression_params](resources--virtual_host--reference--group-001.md#canonical-0031211031111203-2132322102121311-2321132201121110-3323300211311302-2310031121311103-3210323003130202-2310301301000301-2121120102321013) |
| `compression_params.content_length` | [compression_params.content_length](resources--virtual_host--reference--group-001.md#canonical-0301021121011203-2323321100122010-0312130302110121-0322111010002312-1200320133020202-2000031122322120-2033201132323320-2303300123022312) |
| `compression_params.content_type` | [compression_params.content_type](resources--virtual_host--reference--group-001.md#canonical-0110132312321221-1110032302010002-2312121020310022-2320101313301030-3022032013003330-0331320020113310-1200113122101131-0231122212310200) |
| `compression_params.disable_on_etag_header` | [compression_params.disable_on_etag_header](resources--virtual_host--reference--group-001.md#canonical-3111303120201001-3203301320332302-1000312300113310-1323111201220011-0032232010122031-0301220301003023-3110013322121101-1112330020122213) |
| `compression_params.remove_accept_encoding_header` | [compression_params.remove_accept_encoding_header](resources--virtual_host--reference--group-001.md#canonical-0121010310031311-1130223200113313-1221313013121223-0220202022102322-1103303312311231-3323221000133302-0013202321201021-0212102313223101) |
| `connection_idle_timeout` | [connection_idle_timeout](resources--virtual_host--reference--group-001.md#canonical-3032310110300222-2000212021201213-0223203311330221-0020101110013232-3303120030031001-2010133132130000-1323033031102122-3112201220102032) |
| `cors_policy` | [cors_policy](resources--virtual_host--reference--group-001.md#canonical-1130200103032202-0123310333002032-1000010331032332-1022233233012013-1223232030010332-2121121130102002-2200003223011003-3232203131113100) |
| `cors_policy.allow_credentials` | [cors_policy.allow_credentials](resources--virtual_host--reference--group-001.md#canonical-3131233213331001-1113002311133000-0010121022010303-2000233123120112-2033312130133131-0022202203023011-2001221323202301-3111120112130010) |
| `cors_policy.allow_headers` | [cors_policy.allow_headers](resources--virtual_host--reference--group-001.md#canonical-0122013231310310-0212110121303122-3003000212303012-3002020212131011-2332310330202111-1121203201103010-0301130031310031-2231213130130230) |
| `cors_policy.allow_methods` | [cors_policy.allow_methods](resources--virtual_host--reference--group-001.md#canonical-2101211112322121-2300103100213220-2111121132223013-3210022133013130-0310031211321202-3213231212032112-1121303133030202-3302211322201300) |
| `cors_policy.allow_origin` | [cors_policy.allow_origin](resources--virtual_host--reference--group-001.md#canonical-1130202231311301-3330300031020110-1122101230020033-2002013310233100-2003011302022013-3102021130120221-3200021330312310-3000323213123331) |
| `cors_policy.allow_origin_regex` | [cors_policy.allow_origin_regex](resources--virtual_host--reference--group-001.md#canonical-1022012010132013-3331001120023021-0133303321210303-1333133201123131-0233203312230322-3033022002111032-0021313003212012-1233133033133100) |
| `cors_policy.disabled` | [cors_policy.disabled](resources--virtual_host--reference--group-001.md#canonical-1221122011312203-1232112122311221-3102003313302112-2002103210122312-1033322011200322-1212132211213233-1301010312302333-1213220111310230) |
| `cors_policy.expose_headers` | [cors_policy.expose_headers](resources--virtual_host--reference--group-001.md#canonical-0011322302112011-1223033011121000-0330231302300223-2000132022002001-3132212033231120-0030300313032222-0013223333123311-1032322202030233) |
| `cors_policy.maximum_age` | [cors_policy.maximum_age](resources--virtual_host--reference--group-001.md#canonical-1211222333020113-1100002102312331-3121031203111333-2111001231213323-0303133212213102-0111112133311003-3002203200101121-3300300122112302) |
| `csrf_policy` | [csrf_policy](resources--virtual_host--reference--group-001.md#canonical-3000331031023322-3131110331302003-2320301022032003-1010232010300203-1103201100200333-0321313303002220-3023320100320230-1120102030310133) |
| `csrf_policy.all_load_balancer_domains` | [csrf_policy.all_load_balancer_domains](resources--virtual_host--reference--group-001.md#canonical-3133102211130012-2123322312022330-0023111120020320-1231211020211130-3320222020302212-3221122332320201-0312001003320213-0121333032231313) |
| `csrf_policy.custom_domain_list` | [csrf_policy.custom_domain_list](resources--virtual_host--reference--group-002.md#canonical-3022331111332211-0031131112222202-3131030223330302-0110001000321200-3232202013212100-1320302022121232-3001233213210200-2200323031112011) |
| `csrf_policy.custom_domain_list.domains` | [csrf_policy.custom_domain_list.domains](resources--virtual_host--reference--group-002.md#canonical-2101300102020301-2101200313113331-0220321031323312-2230003102122130-3301001111312030-3033323303300033-0031233233321032-0321210100033133) |
| `csrf_policy.disabled` | [csrf_policy.disabled](resources--virtual_host--reference--group-002.md#canonical-3121103112021120-0011030112000311-2302123123110223-3223313113231001-2131230123131131-0202132333333211-1022120220221321-1323331310320122) |
| `custom_errors` | [custom_errors](resources--virtual_host--reference--group-001.md#canonical-0311101123310131-0113032000322132-2330012100310121-1000103023313213-3220131033131221-2323000031310313-1313130231210122-0131211113000313) |
| `default_header` | [default_header](resources--virtual_host--reference--group-002.md#canonical-3103110113233101-2311030300030002-3133031023300010-1132231000122223-1300222101112311-3130112321310012-0333210202330121-0301221023322011) |
| `default_loadbalancer` | [default_loadbalancer](resources--virtual_host--reference--group-002.md#canonical-3000212102323033-3010302213033131-1311032222332011-3102311311323303-3231233111020000-3032233222331012-3132021303003223-3203122122312100) |
| `description` | [description](resources--virtual_host--reference--group-001.md#canonical-0030002233332331-3231212021013021-3012013122113301-2221302213223021-0230122310220033-0031202333020303-1130023120032301-1201010010332021) |
| `disable` | [disable](resources--virtual_host--reference--group-001.md#canonical-2003222121102200-3023011233323121-0102220130201310-3000330003100122-2203301130331101-3303230221312121-3312020121123203-1110212000202232) |
| `disable_default_error_pages` | [disable_default_error_pages](resources--virtual_host--reference--group-001.md#canonical-1001322231001000-0232313010110102-0202220213223333-0220001233312033-3123230200011123-2301021113221020-1330331101010221-2020113001111310) |
| `disable_dns_resolve` | [disable_dns_resolve](resources--virtual_host--reference--group-001.md#canonical-0013130302202231-2302310310120022-1101121131101321-3101012121213221-0311233303300201-3230202310203023-0012233011211200-3132021120331021) |
| `disable_path_normalize` | [disable_path_normalize](resources--virtual_host--reference--group-002.md#canonical-2301113032002302-0100100301320003-2232331201100022-0101013130330333-0310113103302030-1201021302033333-3100320010022221-2210331012103001) |
| `domains` | [domains](resources--virtual_host--reference--group-001.md#canonical-3212000311330033-3332230300113221-1321030133310313-2310133120230230-2113011130011110-0212031123103320-2300100230130123-2000321302332032) |
| `dynamic_reverse_proxy` | [dynamic_reverse_proxy](resources--virtual_host--reference--group-002.md#canonical-1313301200300110-2211133100220020-0323222013011332-0203113310011312-0030120202123032-1323133201003113-1312022023331003-1322222111213110) |
| `dynamic_reverse_proxy.connection_timeout` | [dynamic_reverse_proxy.connection_timeout](resources--virtual_host--reference--group-002.md#canonical-0223112013223232-2011132122110100-2330212121311333-3121002322123130-0322113120121013-3302013032003110-2032032123132223-2101203222023123) |
| `dynamic_reverse_proxy.resolution_network` | [dynamic_reverse_proxy.resolution_network](resources--virtual_host--reference--group-002.md#canonical-3102311130233210-3021002322331210-2111323113033013-1100321012211322-0312100303210020-3000301012221103-0123130030223233-0233312230220133) |
| `dynamic_reverse_proxy.resolution_network.kind` | [dynamic_reverse_proxy.resolution_network.kind](resources--virtual_host--reference--group-002.md#canonical-0232123323033222-3031320122332211-2103111102110222-1032122111301213-0211132120111012-1020103311022210-1220303022120122-0131031330002013) |
| `dynamic_reverse_proxy.resolution_network.name` | [dynamic_reverse_proxy.resolution_network.name](resources--virtual_host--reference--group-002.md#canonical-2002233033212032-3120023223122312-3310023130223322-0232202111123311-2320232110300220-1132220213123200-1220303103003112-0133033021121131) |
| `dynamic_reverse_proxy.resolution_network.namespace` | [dynamic_reverse_proxy.resolution_network.namespace](resources--virtual_host--reference--group-002.md#canonical-0031013311021113-2201033233210000-1132220231323112-3213101233000310-2121131123122230-2321103230333332-2230301213131221-1320101332330333) |
| `dynamic_reverse_proxy.resolution_network.tenant` | [dynamic_reverse_proxy.resolution_network.tenant](resources--virtual_host--reference--group-002.md#canonical-0130313230221120-1322312032310331-1032310333133310-0133200231233123-0311212132232230-0220200001130201-0020012000313113-3012312212002231) |
| `dynamic_reverse_proxy.resolution_network.uid` | [dynamic_reverse_proxy.resolution_network.uid](resources--virtual_host--reference--group-002.md#canonical-0011332222022203-1321313000323223-2001001020220102-2103303333333333-1213213201131200-3122223123123202-1233112131013321-1002302010312332) |
| `dynamic_reverse_proxy.resolution_network_type` | [dynamic_reverse_proxy.resolution_network_type](resources--virtual_host--reference--group-002.md#canonical-2000023100203002-2312211232023010-0113221312023022-3002021230113223-3203303333112002-0222023102001002-1310331300111101-0030332222312033) |
| `dynamic_reverse_proxy.resolve_endpoint_dynamically` | [dynamic_reverse_proxy.resolve_endpoint_dynamically](resources--virtual_host--reference--group-002.md#canonical-0121231000330131-1303300020303311-0302201331312312-3002032021323221-2220113112220200-3031100122322131-1333323003310103-1222020232110200) |
| `enable_path_normalize` | [enable_path_normalize](resources--virtual_host--reference--group-002.md#canonical-3220103012323132-3123012003332321-1010200201202130-0001120230103122-2231023300220233-0201200010231213-0212213100033130-1203010112203200) |
| `http_protocol_options` | [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-0121200200323330-3000031332102221-0220123010203231-2022233221110203-0012032001020321-0320130130030032-1312231312302031-0233031323201121) |
| `http_protocol_options.http_protocol_enable_v1_only` | [http_protocol_options.http_protocol_enable_v1_only](resources--virtual_host--reference--group-002.md#canonical-3330312023321112-3330313332212130-0120133011211001-3103312021201100-3000332101020231-3033021210323010-0131023230331013-1322332220302110) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--reference--group-002.md#canonical-2223100312013213-0332203100322230-1123022310201220-1312022210333331-3121310103120000-1122130003323012-1112222301332130-3123123011311123) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--virtual_host--reference--group-002.md#canonical-1211112003010222-1132210302020312-3200213211203023-3223132223223323-3131312011333031-1323012200202331-3020233322032303-0321212032132302) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--virtual_host--reference--group-002.md#canonical-0112122310122312-0130121212231300-0303120201203002-0203203023010322-2302320323320123-0312003322102021-1030102032322312-3000331302121301) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--virtual_host--reference--group-002.md#canonical-3333100110223012-0101032322023230-0302033322210012-3200111120013222-1030121321300203-2223323133303102-1203012331311232-3122113001220111) |
| `http_protocol_options.http_protocol_enable_v1_v2` | [http_protocol_options.http_protocol_enable_v1_v2](resources--virtual_host--reference--group-002.md#canonical-0003033122202032-3012112010322112-3130100200211113-0222303121320312-1132233323223312-2202230312233131-0033223332200231-2223233132201301) |
| `http_protocol_options.http_protocol_enable_v2_only` | [http_protocol_options.http_protocol_enable_v2_only](resources--virtual_host--reference--group-002.md#canonical-2300022212233023-0322210221310110-0301000331312331-3001203310203123-2133232023321231-2001123210011021-1101212313022103-1123233110110001) |
| `id` | [ID](resources--virtual_host--reference--group-001.md#canonical-2100122103113131-1300210011133013-2332123120112031-0211300012113033-1002303312231332-2211231310201032-0320120311000222-2031132103332113) |
| `idle_timeout` | [idle_timeout](resources--virtual_host--reference--group-001.md#canonical-2130311100021022-0231313123020200-2011301202100101-0212022023311022-2110321111203231-3133303221213013-1220001320321030-3112232011003003) |
| `js_challenge` | [js_challenge](resources--virtual_host--reference--group-002.md#canonical-2301110332123321-2201100001322312-3310232321120302-2100202013231333-1032213203023203-3311131310311320-2313310010120313-0030132221110033) |
| `js_challenge.cookie_expiry` | [js_challenge.cookie_expiry](resources--virtual_host--reference--group-002.md#canonical-2200112100333330-2303233000113213-3212132031201320-0212303202022102-2120120101323201-2003211321300313-3332323023023231-1100212112031211) |
| `js_challenge.custom_page` | [js_challenge.custom_page](resources--virtual_host--reference--group-002.md#canonical-2030230111000310-0320101310031231-0321202033312222-0333120230011323-2102311023321321-1012110011313222-3113231302011121-3120110323033332) |
| `js_challenge.js_script_delay` | [js_challenge.js_script_delay](resources--virtual_host--reference--group-002.md#canonical-0123013001000222-2133130113221231-3300212132013200-3020131221210312-2323031211333202-3113331323203212-3002023122101303-2300231001110222) |
| `labels` | [labels](resources--virtual_host--reference--group-001.md#canonical-0030111302020011-3233132301130311-2223020332120120-0022233112232323-1103020111123112-2103100113321030-2003122112003113-0203121201121120) |
| `max_request_header_size` | [max_request_header_size](resources--virtual_host--reference--group-001.md#canonical-2301221233022002-1033000200231312-3302301110130013-2320013313321302-2103022203021132-3322332032100230-0300002000030230-1311310010111130) |
| `max_requests_per_connection` | [max_requests_per_connection](resources--virtual_host--reference--group-001.md#canonical-3302203010032022-0312123123321211-3211010021230223-3131033301001120-0312112002101330-3323100120222113-1112211233210202-3301310201101123) |
| `name` | [name](resources--virtual_host--reference--group-001.md#canonical-3232022013210020-0331232120212123-1221211300320201-3022330213113112-1022203033321213-1322213313333121-1220232230301012-2101322111230011) |
| `namespace` | [namespace](resources--virtual_host--reference--group-001.md#canonical-2323100200332233-2011102133130001-1222113033131122-2120201122213103-0313232011233313-0101303300103130-2313002000332200-1312121100313031) |
| `no_authentication` | [no_authentication](resources--virtual_host--reference--group-002.md#canonical-0103231132201100-0003112003220201-1232112230200001-0013023021010100-2023323122300101-0222212120020221-3323111231232231-2003131212032321) |
| `no_challenge` | [no_challenge](resources--virtual_host--reference--group-002.md#canonical-1023032113232030-0020323313122121-0322301020011111-0223320231112323-1001102230211202-1023032220132002-0312223223230123-1312013321231030) |
| `no_request_limit_per_connection` | [no_request_limit_per_connection](resources--virtual_host--reference--group-002.md#canonical-1220233000321320-2201210110230233-3022023021301313-3131222033332030-3021323101221112-2320200132210013-0112101301202311-0300202201122113) |
| `non_default_loadbalancer` | [non_default_loadbalancer](resources--virtual_host--reference--group-002.md#canonical-0203302010201112-0331211102220230-0300132232103103-1210332110032101-1332230330310132-1011232020203331-2233222103311321-3203201030110132) |
| `pass_through` | [pass_through](resources--virtual_host--reference--group-002.md#canonical-0132112102332232-0123332032222330-0301020323101110-3123022022131303-1232332032232110-0303301323030131-0113120332013210-1113310021333202) |
| `proxy` | [proxy](resources--virtual_host--reference--group-001.md#canonical-0221100123330010-0213131020311100-0132311121102322-1031202000000222-2303221011223212-3232320210123212-3231022110202221-3223312233230221) |
| `rate_limiter_allowed_prefixes` | [rate_limiter_allowed_prefixes](resources--virtual_host--reference--group-002.md#canonical-3301131211121202-3332111111230013-2113212301003120-2021213232103210-0002020320300330-0100112121211312-3010223232330201-2112231233011130) |
| `rate_limiter_allowed_prefixes.kind` | [rate_limiter_allowed_prefixes.kind](resources--virtual_host--reference--group-002.md#canonical-1112011231300102-3300100010132003-1031313112202220-2020211133203122-2003031213203312-3303012123313231-3202003000300233-0332112203130000) |
| `rate_limiter_allowed_prefixes.name` | [rate_limiter_allowed_prefixes.name](resources--virtual_host--reference--group-002.md#canonical-0132122021311320-3023233222230102-1001101200123310-3002333121301330-2231132101133302-0021131122121120-0213022303302030-1022231032333103) |
| `rate_limiter_allowed_prefixes.namespace` | [rate_limiter_allowed_prefixes.namespace](resources--virtual_host--reference--group-002.md#canonical-2133312030102020-0200220310130122-2232012120132030-0303233321113312-0032022203331210-2212203323100320-3020313131330011-1132200300002223) |
| `rate_limiter_allowed_prefixes.tenant` | [rate_limiter_allowed_prefixes.tenant](resources--virtual_host--reference--group-002.md#canonical-3110002303210333-2033011203312313-0310330021233230-3120311200230003-3102103203211023-0013200012012103-3201030312132213-3332213321003223) |
| `rate_limiter_allowed_prefixes.uid` | [rate_limiter_allowed_prefixes.uid](resources--virtual_host--reference--group-002.md#canonical-1213110300200123-3100001102333012-1202001002031132-3231313223311102-1231032121032110-0110301321212033-1323110220122300-0110121002301231) |
| `request_cookies_to_add` | [request_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1132200202122012-0120231300202311-1120202301232030-0133023100121320-0313102032300112-1303032313321032-3233123332123220-2123322212001013) |
| `request_cookies_to_add.name` | [request_cookies_to_add.name](resources--virtual_host--reference--group-002.md#canonical-3110333310310322-2332330123133021-2321112123103312-1321100111332112-2102130131001310-0312330021023323-1013220111001332-1122210101331122) |
| `request_cookies_to_add.overwrite` | [request_cookies_to_add.overwrite](resources--virtual_host--reference--group-002.md#canonical-0332100333321211-3331133232201230-3112220021022313-3331213102300221-2220310001201232-1231223310231213-3211103223022210-3013002000333003) |
| `request_cookies_to_add.secret_value` | [request_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-3000303221003030-1131111221011200-0032123200103013-2232133132011313-1333231331130232-1030300222313003-3023233002321200-0213112001202101) |
| `request_cookies_to_add.secret_value.blindfold_secret_info` | [request_cookies_to_add.secret_value.blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-1132212111133203-0110230032011110-3230012121010213-1131230013030212-3213011322011102-2001131102120123-2302302323212231-1110002113213111) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--virtual_host--reference--group-002.md#canonical-1323002121230133-3211001201313111-1200023313022123-0221110121332201-1033230220331201-3022333002032100-3101100031012130-1223210312012030) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.location` | [request_cookies_to_add.secret_value.blindfold_secret_info.location](resources--virtual_host--reference--group-002.md#canonical-3222233101030000-2321122212102221-2222320323313020-0231102101102102-1330331310210210-1300111222021132-2200023033212333-0120032002313231) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--virtual_host--reference--group-002.md#canonical-1121322202213333-1003103022211131-1212033133310101-1031220003213123-3113023000332220-0213230132202311-2000131003133121-2020310112000112) |
| `request_cookies_to_add.secret_value.clear_secret_info` | [request_cookies_to_add.secret_value.clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-3111133101200302-3112022120333133-0220331201122321-0021122331010131-0130001312201013-2111031300222023-3201000233321301-0010021301321323) |
| `request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [request_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--virtual_host--reference--group-002.md#canonical-0333121232013311-2020101201033302-1201032132322101-1100022322333120-1013110233310021-1030232131233112-3111201000110130-1300210102121211) |
| `request_cookies_to_add.secret_value.clear_secret_info.url` | [request_cookies_to_add.secret_value.clear_secret_info.url](resources--virtual_host--reference--group-002.md#canonical-2203223313011111-0301002302020021-3320331302111210-0210020020130003-2021100300132200-0333321211301200-0210331002221233-2303301120212313) |
| `request_cookies_to_add.value` | [request_cookies_to_add.value](resources--virtual_host--reference--group-002.md#canonical-1312300232101312-3321320301021013-2303102103100002-0010110101010002-3210223332310320-1010222010031032-1330200230130322-2003000230110203) |
| `request_cookies_to_remove` | [request_cookies_to_remove](resources--virtual_host--reference--group-001.md#canonical-1133330212300212-3331103111021021-2222211121303313-1132022022331223-0012010222330130-1323313333101330-3122002312312112-2222110220100202) |
| `request_headers_to_add` | [request_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-3012313223321220-3121120103002220-3002111110010303-2131002021322310-1130100033030200-0333222300022220-2303003031131332-1233011000012013) |
| `request_headers_to_add.append` | [request_headers_to_add.append](resources--virtual_host--reference--group-002.md#canonical-2010323003001201-0231101202111211-2120113301310303-3312133100101001-2201303020221012-2332033001021213-3302203123233222-1012220033201003) |
| `request_headers_to_add.name` | [request_headers_to_add.name](resources--virtual_host--reference--group-002.md#canonical-1203233002121312-3211321130122001-3200301223231123-0110230220213120-3013022120013201-2222213300031321-3120301131121223-1322013112130100) |
| `request_headers_to_add.secret_value` | [request_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-3232230202330303-2130001230123102-1132113011220232-2023301202100301-2323121322221323-0330032213301110-0300212303011010-2201232102001011) |
| `request_headers_to_add.secret_value.blindfold_secret_info` | [request_headers_to_add.secret_value.blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-2313103131023113-1011002230021132-2213101022023113-0021213311030232-3113002303232212-0220101233120032-1032311110333022-3030202211221233) |
| `request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--virtual_host--reference--group-002.md#canonical-1233302001023212-2322221231211312-2010222210202121-3220033013323333-1321313201012031-2112023201132333-0302101121303103-0233102201333131) |
| `request_headers_to_add.secret_value.blindfold_secret_info.location` | [request_headers_to_add.secret_value.blindfold_secret_info.location](resources--virtual_host--reference--group-002.md#canonical-3310302302211221-0020102100103113-3303122133103013-3230221202313123-1120112010232310-3321333321233303-0233133320122202-2333312310331112) |
| `request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [request_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--virtual_host--reference--group-002.md#canonical-3102211033132301-0100133322312313-2023311210012200-0010000120013332-1211032020311010-0212303322100201-2003030003113300-2122302300012130) |
| `request_headers_to_add.secret_value.clear_secret_info` | [request_headers_to_add.secret_value.clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-1221212101301212-0112121012232333-2033300101312312-0022201332120231-0112223312102120-0233302322133312-1223312131120002-3332313032111230) |
| `request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [request_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--virtual_host--reference--group-002.md#canonical-2112002133311301-0323013311233013-3133110323300322-0212131031201223-1233022211001013-2200122210030320-3122001012311202-1023321311312030) |
| `request_headers_to_add.secret_value.clear_secret_info.url` | [request_headers_to_add.secret_value.clear_secret_info.url](resources--virtual_host--reference--group-002.md#canonical-2323211222111331-2300102300130302-2011102332122120-0211211221221322-0132001101030031-3220313130030123-0201132001203312-0130012210133203) |
| `request_headers_to_add.value` | [request_headers_to_add.value](resources--virtual_host--reference--group-002.md#canonical-0013233321213220-2212201230202132-3232211100113332-0322133131011032-2233300322320021-1321333312010100-2333100103200222-1021322122132023) |
| `request_headers_to_remove` | [request_headers_to_remove](resources--virtual_host--reference--group-001.md#canonical-1100302332102030-2323230022003003-1330220321313332-2322002113222021-0002102022112000-1201023131232203-0031030110230320-2100220033303302) |
| `response_cookies_to_add` | [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-2222003011301030-3202222031130133-1300311132321213-0002212031300032-1232213100222100-1111332011333020-0302312231231200-1310300303220332) |
| `response_cookies_to_add.add_domain` | [response_cookies_to_add.add_domain](resources--virtual_host--reference--group-002.md#canonical-2021011121031211-3210023310302303-2020203012111211-0221120301100103-3212101023203021-2130102121211231-2322331312031022-3030102330312111) |
| `response_cookies_to_add.add_expiry` | [response_cookies_to_add.add_expiry](resources--virtual_host--reference--group-002.md#canonical-1332223231213232-1112200333232311-3210301020332000-3220212222132230-2010311223203223-0120130320131032-2233120110111203-1210320033102300) |
| `response_cookies_to_add.add_httponly` | [response_cookies_to_add.add_httponly](resources--virtual_host--reference--group-002.md#canonical-3120210223030310-3030233311200121-3021100300230021-3132333020111033-0003022231312221-1012310033020022-1133301002130011-0120202210231212) |
| `response_cookies_to_add.add_partitioned` | [response_cookies_to_add.add_partitioned](resources--virtual_host--reference--group-002.md#canonical-3002222321133222-1211333333013321-2111122313120123-3332021132331033-3002132122211310-2311123013230003-2113321000322202-0302201010332331) |
| `response_cookies_to_add.add_path` | [response_cookies_to_add.add_path](resources--virtual_host--reference--group-002.md#canonical-1100032020320211-1220111000033233-2222030022020113-1120131203023020-1021320121010212-0230031013010310-1213312003212021-1212013200113203) |
| `response_cookies_to_add.add_secure` | [response_cookies_to_add.add_secure](resources--virtual_host--reference--group-002.md#canonical-0221210230121010-1030110221213031-1212220021223203-3021301321222013-2303320130322233-2311033332322112-0313020111313022-1202131101300122) |
| `response_cookies_to_add.ignore_domain` | [response_cookies_to_add.ignore_domain](resources--virtual_host--reference--group-002.md#canonical-3231320210312130-0220031111132303-0330101001001033-2001223300313031-1302330022010033-2022210301031031-2212103221031231-1220030203203223) |
| `response_cookies_to_add.ignore_expiry` | [response_cookies_to_add.ignore_expiry](resources--virtual_host--reference--group-002.md#canonical-1222320221003333-2312002000233333-3222002211220222-2313012020122000-0111312231201212-3321220110201123-3231110220122321-1232123003133010) |
| `response_cookies_to_add.ignore_httponly` | [response_cookies_to_add.ignore_httponly](resources--virtual_host--reference--group-002.md#canonical-2033020133203330-2012130130232022-2021120112011101-1332122031131330-1332112230133333-3111001210302103-1232303021030320-3110121002021100) |
| `response_cookies_to_add.ignore_max_age` | [response_cookies_to_add.ignore_max_age](resources--virtual_host--reference--group-002.md#canonical-2332113113003020-1322101231212120-0132020311110113-2123000110131322-3022313133210310-1203232330112103-0210110133103031-2212133323322313) |
| `response_cookies_to_add.ignore_partitioned` | [response_cookies_to_add.ignore_partitioned](resources--virtual_host--reference--group-002.md#canonical-0011210320120210-2232220310231120-0000311010223303-0330111213202302-2321101301222001-0320231101333332-1120100000130120-1120310301102203) |
| `response_cookies_to_add.ignore_path` | [response_cookies_to_add.ignore_path](resources--virtual_host--reference--group-002.md#canonical-3210022222033021-2202133320331330-1003301111223131-2010103012122113-2200222222002232-3133210131103312-0230321211002201-2000231230213131) |
| `response_cookies_to_add.ignore_samesite` | [response_cookies_to_add.ignore_samesite](resources--virtual_host--reference--group-002.md#canonical-2011020332310022-0201003300002320-1211123211021232-3313203000233302-2033210131020030-2131230221222332-2103102131103313-2321012230230303) |
| `response_cookies_to_add.ignore_secure` | [response_cookies_to_add.ignore_secure](resources--virtual_host--reference--group-002.md#canonical-2101322000003222-1012302202233233-1001013031032132-2000212121201330-2011220311133030-3202003220100023-1122132213000220-2023130133102313) |
| `response_cookies_to_add.ignore_value` | [response_cookies_to_add.ignore_value](resources--virtual_host--reference--group-002.md#canonical-0213002033003332-0111222102033220-1311130033221331-3132230322210022-0130101102330101-2120002032320202-3321011113130031-3032111330112323) |
| `response_cookies_to_add.max_age_value` | [response_cookies_to_add.max_age_value](resources--virtual_host--reference--group-002.md#canonical-3012301010312230-1222133313103133-0112000321231213-3013031310000203-3113002232133110-1301020220111203-0232113120001010-2332011101102002) |
| `response_cookies_to_add.name` | [response_cookies_to_add.name](resources--virtual_host--reference--group-002.md#canonical-3221033320123302-3333110200003230-2130020301100001-1323103313233310-2200203320223331-3113333032202330-2112201333001131-2001113300303032) |
| `response_cookies_to_add.overwrite` | [response_cookies_to_add.overwrite](resources--virtual_host--reference--group-002.md#canonical-3013021010000112-2133312130123322-1002313233033200-0013222210323001-3221111203332023-3221100221313022-3303231030010021-0311011101233013) |
| `response_cookies_to_add.samesite_lax` | [response_cookies_to_add.samesite_lax](resources--virtual_host--reference--group-002.md#canonical-0132233221213211-0232200202330301-1022100132323301-0331221033100233-1311231131203122-0000300022002000-2320201011123101-3023233201132130) |
| `response_cookies_to_add.samesite_none` | [response_cookies_to_add.samesite_none](resources--virtual_host--reference--group-002.md#canonical-3122300331321103-2100331310330333-1201132103301031-1311222322233233-0322303010101011-3200310021101203-0033222003020020-1222301120320110) |
| `response_cookies_to_add.samesite_strict` | [response_cookies_to_add.samesite_strict](resources--virtual_host--reference--group-002.md#canonical-0300322121023121-3203023200020310-2302121113011232-3232231320202203-3231003330201110-1132233331313202-0013100301202300-0100022110313211) |
| `response_cookies_to_add.secret_value` | [response_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-0210333030013002-2101002202331002-2011133222322133-0210132303330022-0212321333231213-1030321310103123-1000302123000020-1102231123032101) |
| `response_cookies_to_add.secret_value.blindfold_secret_info` | [response_cookies_to_add.secret_value.blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-3201302200000310-0311222231330013-1231313303010133-2301001233012010-1303321020231113-3022120110230123-0211031323212100-3102220201320122) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--virtual_host--reference--group-002.md#canonical-3310132220002331-3030312200210132-1132100302302032-0313012103112312-0320113033300302-0020030022123311-1331333021021310-2011210310030210) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.location` | [response_cookies_to_add.secret_value.blindfold_secret_info.location](resources--virtual_host--reference--group-002.md#canonical-1002013121110101-0211100012202320-2112301112231220-0222333310122230-3200010313011113-3310213213213013-0220011122131010-1220300210322031) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--virtual_host--reference--group-002.md#canonical-0032332111130133-3321030330002122-2032023012133030-0221122002222210-3201101021300313-1031020232323333-0320221302332311-1112003320121300) |
| `response_cookies_to_add.secret_value.clear_secret_info` | [response_cookies_to_add.secret_value.clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-2201120200032020-3231311213302112-1301111132002220-0202333213133220-0021121202211221-3030301001202113-1022222200202312-3003001322031221) |
| `response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [response_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--virtual_host--reference--group-002.md#canonical-0320023110210030-0132120000023010-2003233200032232-1033033333102032-2302232311011110-2201231232310331-0220033101211311-1023031221330303) |
| `response_cookies_to_add.secret_value.clear_secret_info.url` | [response_cookies_to_add.secret_value.clear_secret_info.url](resources--virtual_host--reference--group-002.md#canonical-0122102121020011-2032120130332333-2200123132330033-0220030130232103-0130313312010111-3220231110302223-3331313222023220-1130131333330003) |
| `response_cookies_to_add.value` | [response_cookies_to_add.value](resources--virtual_host--reference--group-002.md#canonical-1330120322003022-3120321110121302-2111132010331030-1130313121303303-3010330130001001-2233020120121030-3113023322023311-2312322200303101) |
| `response_cookies_to_remove` | [response_cookies_to_remove](resources--virtual_host--reference--group-001.md#canonical-0021022033103122-1202330200232333-3201001113021130-0212302323332013-0231120333031230-0200300003332213-2122022002232013-1132203333330221) |
| `response_headers_to_add` | [response_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-1332201100013111-2211122001031301-3021122213021002-3103112322011321-1203121332113223-1233013010332233-2313001300333001-1313003003202120) |
| `response_headers_to_add.append` | [response_headers_to_add.append](resources--virtual_host--reference--group-002.md#canonical-1002223013013103-2102120223002030-1132001311122000-0310111231212011-0030132021203103-1200302310031220-3320213123333230-2020120032130001) |
| `response_headers_to_add.name` | [response_headers_to_add.name](resources--virtual_host--reference--group-002.md#canonical-1110020330011320-0301111023102030-0003130000010223-2001003023103133-0000032033331303-1132211333120122-2310221231211001-3013002233202111) |
| `response_headers_to_add.secret_value` | [response_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-1022221223111302-1022022232321022-1120002030211230-1101301311000311-3111213102210232-2132233222132130-1231321212320312-3133213103111301) |
| `response_headers_to_add.secret_value.blindfold_secret_info` | [response_headers_to_add.secret_value.blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-0130000101120133-3330110310003133-1030200231211102-0332201203201230-3102013302233322-1212300113012200-2213203321130132-2232031310101020) |
| `response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--virtual_host--reference--group-002.md#canonical-1111111033210122-1101122203100302-0221001130032300-3102022022333022-2330121333133333-2232210230332203-2113123220013313-3313222123232310) |
| `response_headers_to_add.secret_value.blindfold_secret_info.location` | [response_headers_to_add.secret_value.blindfold_secret_info.location](resources--virtual_host--reference--group-002.md#canonical-1111130000031302-0133021123222020-2320332132300230-3313201201030202-2221201333312132-2323012303313133-3133323121221232-1232032332310330) |
| `response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [response_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--virtual_host--reference--group-002.md#canonical-0022001000331301-1333032031222220-3320311321313032-0030313223332031-3003300301330300-1103101031131120-1132111233220002-0112201201312321) |
| `response_headers_to_add.secret_value.clear_secret_info` | [response_headers_to_add.secret_value.clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-3201300033031132-3102213312230110-3200121221130002-1023200120331300-3230013111301132-2030330103012213-0301003133123000-0132311302103310) |
| `response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [response_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--virtual_host--reference--group-002.md#canonical-0301030202012021-1031112013213332-0020010013030001-0132201123200003-0213013101010120-3011122021023131-0011023032220332-3011103001102211) |
| `response_headers_to_add.secret_value.clear_secret_info.url` | [response_headers_to_add.secret_value.clear_secret_info.url](resources--virtual_host--reference--group-002.md#canonical-1320102133121300-3203103220333200-3010331312220321-1133333132303110-1111330002332321-2220023213320031-3003310101013301-0130323201220222) |
| `response_headers_to_add.value` | [response_headers_to_add.value](resources--virtual_host--reference--group-002.md#canonical-1020202333133332-0012231300003001-1303210322300100-2030022011210310-2101011223100332-0321110122110112-0100302200222313-2100011130210121) |
| `response_headers_to_remove` | [response_headers_to_remove](resources--virtual_host--reference--group-001.md#canonical-2333020201013203-3202312230231313-0302112333030221-1020023111131203-1102321222121331-3100222012333100-0123033103232131-1330302200313013) |
| `retry_policy` | [retry_policy](resources--virtual_host--reference--group-002.md#canonical-2030200310111133-3030032003003120-3002230020001330-3301223321122321-1312122333313220-0210331300030232-3323011033030330-1311130010323222) |
| `retry_policy.back_off` | [retry_policy.back_off](resources--virtual_host--reference--group-002.md#canonical-2022202132020122-0313220232010302-3323220003131000-2031030233333230-3131220223033201-0011222012332112-0112320230011102-2023131003303230) |
| `retry_policy.back_off.base_interval` | [retry_policy.back_off.base_interval](resources--virtual_host--reference--group-002.md#canonical-0000230233313210-2020102122303230-3321200120031222-1303022020310110-1013313323231311-1333003213221020-3001311030132301-3011222031210332) |
| `retry_policy.back_off.max_interval` | [retry_policy.back_off.max_interval](resources--virtual_host--reference--group-002.md#canonical-3231131112331230-2022002123220210-3113011203301202-2302330310302201-1333332112023330-2310323300012111-1332023211003102-3032211120133331) |
| `retry_policy.num_retries` | [retry_policy.num_retries](resources--virtual_host--reference--group-002.md#canonical-3321101220111111-0210303233310200-3322021320321323-3330223120110123-1012032303123103-2222330310020030-1321031311333220-2031130212212310) |
| `retry_policy.per_try_timeout` | [retry_policy.per_try_timeout](resources--virtual_host--reference--group-002.md#canonical-0330120133021002-2200030001111231-1010010302211330-1110133103332012-2212201003221211-0120331032312010-1021120021131332-3013112231103000) |
| `retry_policy.retriable_status_codes` | [retry_policy.retriable_status_codes](resources--virtual_host--reference--group-002.md#canonical-2103233032202121-2131010003130032-2113013210000222-1111110310120113-0210013031311000-3323331321011201-1030332120032200-3003202201132113) |
| `retry_policy.retry_condition` | [retry_policy.retry_condition](resources--virtual_host--reference--group-002.md#canonical-2313310213313002-3233220321023030-0133101030030101-0310300223013100-1123210020031310-0133131313012233-3330200322311330-2011031110102212) |
| `routes` | [routes](resources--virtual_host--reference--group-002.md#canonical-1001231111002131-3010222122130310-3000212310212132-2332312133330201-3310103133323231-3122230203211313-0231120223222120-1013010013030320) |
| `routes.kind` | [routes.kind](resources--virtual_host--reference--group-002.md#canonical-1332302203002203-1323131112012302-3212130321233131-2222020231232012-0112202202120210-0003100001213322-2223322311313002-2012200003301002) |
| `routes.name` | [routes.name](resources--virtual_host--reference--group-002.md#canonical-3233200101223303-0331333031002211-3222223003320121-1322231321313012-0121030223213223-0322132220323222-3021200203131133-1120233203031200) |
| `routes.namespace` | [routes.namespace](resources--virtual_host--reference--group-002.md#canonical-2333233100103023-3300021033112020-3031120032113020-0122110223123013-0002132123121023-3033013230301313-3211201223013131-0220133132133222) |
| `routes.tenant` | [routes.tenant](resources--virtual_host--reference--group-002.md#canonical-0103130202322010-0213301300021223-3112001333123233-0103211230330111-1321103321202003-0301023010232321-1303032013202001-2032030210101102) |
| `routes.uid` | [routes.uid](resources--virtual_host--reference--group-002.md#canonical-2102332203002001-0323322233101200-1331003332313110-3223012300133232-1213122100320301-2020213310212303-2200001301312132-1311101333021310) |
| `sensitive_data_policy` | [sensitive_data_policy](resources--virtual_host--reference--group-002.md#canonical-0222123113222031-2013231001121032-1101203330320312-2232211013122133-3202312112000010-0012121230102013-2101110323013332-3013221200021231) |
| `sensitive_data_policy.kind` | [sensitive_data_policy.kind](resources--virtual_host--reference--group-002.md#canonical-0131000121302011-2211210111203202-3202112101330022-3313330121012122-2021031221310010-3111111301310100-0110013310300230-3222333202322002) |
| `sensitive_data_policy.name` | [sensitive_data_policy.name](resources--virtual_host--reference--group-002.md#canonical-3233322232232023-0203210130330331-0303200101211000-2132112013233123-3320231200012123-2210220313133102-2232031332322121-3100231132323232) |
| `sensitive_data_policy.namespace` | [sensitive_data_policy.namespace](resources--virtual_host--reference--group-002.md#canonical-3110012223312303-1313130030112020-0331100232120323-3321001203301003-0200233222200113-2323133010211300-1320001100121123-2233313303101100) |
| `sensitive_data_policy.tenant` | [sensitive_data_policy.tenant](resources--virtual_host--reference--group-002.md#canonical-0113201221033133-3130332221120333-1032123330203122-2122101012020032-3103202022233021-1210012302020032-0202132120312133-2231220202012031) |
| `sensitive_data_policy.uid` | [sensitive_data_policy.uid](resources--virtual_host--reference--group-002.md#canonical-2110331321310100-1220302312211220-2313213200133332-1303221103013133-1201323220013221-1202021020312132-2200123111300130-0103030110123330) |
| `server_name` | [server_name](resources--virtual_host--reference--group-001.md#canonical-3202022102332232-2133012230301013-1001202130203031-0012102010333000-0231302310211120-1101011033112211-3303110301000033-1300011110102320) |
| `slow_ddos_mitigation` | [slow_ddos_mitigation](resources--virtual_host--reference--group-002.md#canonical-0213112013313021-0210021331121222-0031003332330302-2013322310222003-2302001202130331-0023232133221022-2330101022202311-2132022201020312) |
| `slow_ddos_mitigation.disable_request_timeout` | [slow_ddos_mitigation.disable_request_timeout](resources--virtual_host--reference--group-002.md#canonical-0032120312201033-3200032031011122-0011133112303101-1033331011003302-3232321223332332-3022232232310020-1232233310013321-0201211221133333) |
| `slow_ddos_mitigation.request_headers_timeout` | [slow_ddos_mitigation.request_headers_timeout](resources--virtual_host--reference--group-002.md#canonical-0310121203030120-2202111030132111-3230013223020011-0232201133232112-3210322310111002-1020201320121213-2211132302123003-0003003120012112) |
| `slow_ddos_mitigation.request_timeout` | [slow_ddos_mitigation.request_timeout](resources--virtual_host--reference--group-002.md#canonical-1010331231022331-1132221033201011-3230133101222221-2013121321230003-3031003203021123-1310310002323232-2012301111230332-3011023231213112) |
| `timeouts` | [timeouts](resources--virtual_host--reference--group-002.md#canonical-1223303000101110-3011111313112031-1002101201030211-1001300332202033-0331102311120200-0233023330322221-2003333130122211-3310101221230210) |
| `timeouts.create` | [timeouts.create](resources--virtual_host--reference--group-002.md#canonical-1113332021000102-1022012011032030-2313220010330013-3333310312000023-1122113130313220-1013111221012123-0032002330101330-3012130212102132) |
| `timeouts.delete` | [timeouts.delete](resources--virtual_host--reference--group-002.md#canonical-3030233233112021-1210231122222201-0022210333223313-0013233323323122-0011321101322012-1120210233132011-3303000120023020-2313103233221213) |
| `timeouts.read` | [timeouts.read](resources--virtual_host--reference--group-002.md#canonical-1302101300113101-2232013120221032-2130001111123003-1011023213020302-0311120111212333-1223211000133333-0210003110220230-3230132211103030) |
| `timeouts.update` | [timeouts.update](resources--virtual_host--reference--group-002.md#canonical-1310103030131222-1332220101011001-1012020331030322-1323110131211203-1231000312020102-1331323032132213-1133012111210022-3001023220230031) |
| `tls_cert_params` | [tls_cert_params](resources--virtual_host--reference--group-002.md#canonical-2123330110100202-3132023320131302-1212121211011102-3110211230311121-1121003310302011-2120012132033323-2102333031220302-2103213231223331) |
| `tls_cert_params.certificates` | [tls_cert_params.certificates](resources--virtual_host--reference--group-002.md#canonical-1002330313231012-3132011231022303-2330322112320020-3333223010230220-2230020323030221-3301101101021312-0001221121123213-3222202330100022) |
| `tls_cert_params.certificates.kind` | [tls_cert_params.certificates.kind](resources--virtual_host--reference--group-002.md#canonical-2030112331132002-3113310132003032-2331232100230000-0200230011323213-0022103130310033-1103130211221313-2003301233203132-0300301222330133) |
| `tls_cert_params.certificates.name` | [tls_cert_params.certificates.name](resources--virtual_host--reference--group-002.md#canonical-2000100313113101-1301313210100011-2200312233122122-2221320303323322-0102032301030311-0030302132331230-2332202022231223-0112023233333321) |
| `tls_cert_params.certificates.namespace` | [tls_cert_params.certificates.namespace](resources--virtual_host--reference--group-002.md#canonical-2320202202302103-3203100131231231-2202033310030023-3323131020333232-2202213202302212-3210222330301000-2320300010333030-1013121122130111) |
| `tls_cert_params.certificates.tenant` | [tls_cert_params.certificates.tenant](resources--virtual_host--reference--group-002.md#canonical-0210103321000312-3101022020303002-0222032131002212-0301330303330320-0110201223120300-2323233303233003-1221103312002120-1200112220101233) |
| `tls_cert_params.certificates.uid` | [tls_cert_params.certificates.uid](resources--virtual_host--reference--group-002.md#canonical-0033103333210122-3131033102132023-2103000002320311-2201223212010302-2003002203313230-2311333021012310-3010020111030332-3212120020301103) |
| `tls_cert_params.cipher_suites` | [tls_cert_params.cipher_suites](resources--virtual_host--reference--group-002.md#canonical-3332232212210020-0203212331130212-1300121313102121-1221223322011312-3333022100321011-2130003130102020-3332020313022023-0321302322211331) |
| `tls_cert_params.client_certificate_optional` | [tls_cert_params.client_certificate_optional](resources--virtual_host--reference--group-002.md#canonical-1223302233121202-2213311121111220-0322222120322002-2333330312333021-0301101302320000-1331020233022201-0301232202203023-1322033023121103) |
| `tls_cert_params.client_certificate_required` | [tls_cert_params.client_certificate_required](resources--virtual_host--reference--group-002.md#canonical-2102010012200310-3201101313122222-3023032020132310-0123202022323332-0220210320300102-0230223003110030-3100321002312200-0300000220221331) |
| `tls_cert_params.maximum_protocol_version` | [tls_cert_params.maximum_protocol_version](resources--virtual_host--reference--group-002.md#canonical-3202210313202032-2103210032220213-2333313310011023-3320000022101131-0032313032232223-0101023100203130-2320333202112301-0003222302323300) |
| `tls_cert_params.minimum_protocol_version` | [tls_cert_params.minimum_protocol_version](resources--virtual_host--reference--group-002.md#canonical-3113032321331133-1323121012031212-3213020313001013-3302110032002301-2213313313322121-0230022201023101-2332323023210222-2013122300311010) |
| `tls_cert_params.no_client_certificate` | [tls_cert_params.no_client_certificate](resources--virtual_host--reference--group-002.md#canonical-3310213000123010-3132121002122102-3103230132120112-3301100200202131-3310211001131130-2332131103111121-0010012310330323-1132323202112103) |
| `tls_cert_params.validation_params` | [tls_cert_params.validation_params](resources--virtual_host--reference--group-002.md#canonical-1032202302300211-3001333330201033-2112023021131021-2121022110300031-2233123033220131-3032133101222000-3022121132022232-2331211111102331) |
| `tls_cert_params.validation_params.skip_hostname_verification` | [tls_cert_params.validation_params.skip_hostname_verification](resources--virtual_host--reference--group-002.md#canonical-0110332132111203-2133212021330330-3212100333320203-0022213121203020-0233000003322130-3211110100121032-0312120230130232-1323222112100233) |
| `tls_cert_params.validation_params.trusted_ca` | [tls_cert_params.validation_params.trusted_ca](resources--virtual_host--reference--group-002.md#canonical-1002033131101220-0211131331011102-3222212131332030-3030112200103030-2121330011020101-1333203122110120-0322323002313131-1131020202321102) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list](resources--virtual_host--reference--group-002.md#canonical-1111101121102001-0321110223122321-2230131322101220-0303201223120133-3123230230020330-3232201232300022-0001002321302333-1002213333312003) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.kind](resources--virtual_host--reference--group-002.md#canonical-1333112023322221-3313100101333133-1002320321022103-0220322010211232-3200303231013330-3201333310012122-0331001331123100-0232021113032202) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.name](resources--virtual_host--reference--group-002.md#canonical-3323131031113030-3113210102313220-3103232022202130-3303030002230221-3323133211013021-3130223000030013-3011223210330211-0123211231100323) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.namespace](resources--virtual_host--reference--group-002.md#canonical-3121030130012330-2210330211031333-1132310012230213-1103113012210102-2213032002020302-0210123121221030-2110311331000333-1002331332001011) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.tenant](resources--virtual_host--reference--group-002.md#canonical-1100201012230233-0213301233120301-3033323121112133-1023222100113222-0002201103130021-1100313113220113-1212201220111220-1203033102002311) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.uid](resources--virtual_host--reference--group-002.md#canonical-0233133023123231-3003132311323220-2110120101023033-3323303310130133-1120030212121311-1232200123221003-1121130203203202-0010210011200322) |
| `tls_cert_params.validation_params.trusted_ca_url` | [tls_cert_params.validation_params.trusted_ca_url](resources--virtual_host--reference--group-002.md#canonical-0102232033030023-0113332012201000-1312102323333021-2031133211031031-3313020313320333-1300232331030302-3302013232323323-1011120210323103) |
| `tls_cert_params.validation_params.verify_subject_alt_names` | [tls_cert_params.validation_params.verify_subject_alt_names](resources--virtual_host--reference--group-002.md#canonical-1330331320133131-3230022131231000-1320110033311002-0013002210032220-3201200210123011-3210101321101131-3020300122020130-3002301301100201) |
| `tls_cert_params.xfcc_header_elements` | [tls_cert_params.xfcc_header_elements](resources--virtual_host--reference--group-002.md#canonical-3200233311230203-0032131020020012-0003320011300330-2100210331110120-2023023332113213-2013120223101120-0313131322110120-3022030102321311) |
| `tls_parameters` | [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0221010032112232-0002311102011103-1002323023221012-1223320011202010-1100310130023232-0301233120212133-0021101133201301-2201231323313333) |
| `tls_parameters.client_certificate_optional` | [tls_parameters.client_certificate_optional](resources--virtual_host--reference--group-002.md#canonical-0333121111220133-0323020310031210-3333011101131031-3201121301201312-3211211020311333-0123102320230100-3032302023333002-3010000221101202) |
| `tls_parameters.client_certificate_required` | [tls_parameters.client_certificate_required](resources--virtual_host--reference--group-002.md#canonical-3130232032303203-0232011333223321-0030113131322311-0313010120132021-1001000100003301-2121210321312131-2013311200100131-1313302322110000) |
| `tls_parameters.common_params` | [tls_parameters.common_params](resources--virtual_host--reference--group-002.md#canonical-0303120133133233-0132220103133322-1022302231330312-1302033300200132-0301031130022023-3013200203121330-3031202120031130-0010023113130121) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](resources--virtual_host--reference--group-002.md#canonical-2310321300133003-0230000103032103-1231003120202213-3113000213210122-1131113331001110-3222122220323311-1212222013202333-0213312132100201) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](resources--virtual_host--reference--group-002.md#canonical-3133332330120023-2322121212320313-2131311120313023-2222221001230020-3221112322200303-3231220021213000-0022100122003132-2333333332323110) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](resources--virtual_host--reference--group-002.md#canonical-0213203310111320-3203301022001113-0203122220203000-3200302123110222-3231321112312320-1303333303122323-3232311230032312-0323311013031321) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-002.md#canonical-1011323231323101-2001121311000010-1021211103001120-2203300311211232-0231231033120330-0102330323332022-1001313320010300-1222312131021200) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](resources--virtual_host--reference--group-002.md#canonical-3302211131233331-1203311302023312-0302121312102232-2033201220322003-0033302011231110-2211111233000312-0111112211001122-2012210123233002) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--virtual_host--reference--group-002.md#canonical-1301331000212220-0123310303202130-1202122222101313-0021222302101313-3330133233100022-3013103103331121-2002100321130302-1031320201012133) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--virtual_host--reference--group-002.md#canonical-2101210120001021-3312301120323023-1100101002233120-3322221303231330-3320000311132031-1330223130103001-3302010013211200-3331031002333311) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](resources--virtual_host--reference--group-002.md#canonical-1011010332212323-1313312303231200-0033031122132330-3033330310233021-3030002010001233-2123323233323022-2120221301330301-2322220212223200) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--virtual_host--reference--group-002.md#canonical-0101302010000031-3023011131112013-2233103023132011-2112211322201222-0333030013211102-1112301112212301-1323133003022312-1222332133003033) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--reference--group-002.md#canonical-2302212021121312-0000312312213033-0102223312230203-1231320211020201-1302122302102123-3130030311202313-1210120133033102-3012110331121213) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-3231101123122320-1201100232102023-2010212231312012-1001113320211332-3213000223222222-2013031102003301-2131313102300223-3023321110213033) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--virtual_host--reference--group-002.md#canonical-3030230232232312-3303112101112320-3212030010310032-1010013100021230-3031122122133033-2220333111231020-2003122101201021-0213213332311220) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](resources--virtual_host--reference--group-002.md#canonical-3100331112122311-3220232233111123-0321301322130101-2002001331113332-2202302003200313-3322032002010310-2322303201033033-2113030033012023) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--virtual_host--reference--group-002.md#canonical-0303011010303232-3221003323310310-3110111121213110-0123330223321133-2323313211311122-1233331330121300-0032333223022030-0232030103311203) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-0301020201000130-1303100322320331-2223120202001311-0022210010221201-3120021103322203-3002120231130230-2002132332210033-0030221312122011) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](resources--virtual_host--reference--group-002.md#canonical-3313010330101223-3000011030300100-1313321023120311-1230220310221010-2011111133113232-1330302100232100-0131211033333122-0102000321130332) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](resources--virtual_host--reference--group-002.md#canonical-1022201333211033-3102322311210000-3323103323031202-1213302020101211-2330103303103001-0202000111131313-3223221132332232-1031323102102212) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--virtual_host--reference--group-003.md#canonical-3233303122201020-3302201102012022-2002230132202021-1022313312101123-2320113202122202-1223303101311133-2323210201310212-0132132113032203) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-2031103232112023-1301131322123321-3322010330112122-2333130021322233-1130323231323001-3333003021210333-3103131313200010-2310321101122310) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](resources--virtual_host--reference--group-003.md#canonical-2300021102323300-0221100330002100-3333213200321211-2031123132031202-3113301010112003-0223102333201011-3303311102310130-2213332022212001) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-2102002130033323-2330120222232233-0302302223102230-2321122000230332-2002221220220110-0222020122210222-0310213113000300-1221310213031203) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--virtual_host--reference--group-003.md#canonical-3012330211123302-1022033223131232-3113120133013000-0000221203203123-3210202130332330-3210013003001012-2202232013131322-1232222121013122) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](resources--virtual_host--reference--group-003.md#canonical-3212002232220301-3313010002121021-2003113122131223-2111300310233020-0022220303032011-2230022111021101-0120003113132301-2130223012221233) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](resources--virtual_host--reference--group-003.md#canonical-2313110201103221-0131000112102213-1323231013021300-1022212013303212-0201001231110210-1231003332302301-1131020200010130-2012331321330100) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](resources--virtual_host--reference--group-003.md#canonical-1023023233122030-3330233203231330-0132100103311113-3031010113003331-0312120100203022-1201031103120330-0111221121300102-2320213000002233) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](resources--virtual_host--reference--group-003.md#canonical-3031210003122301-1210322301302202-0212313003122220-2113312132330000-2020123000222303-1121011311110210-2103232030313003-3001122103023322) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](resources--virtual_host--reference--group-003.md#canonical-3323031231011300-2221021132130121-0033000333033011-3100330102112110-1332121030013310-3113232123233002-2023332321103121-1120321321121200) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](resources--virtual_host--reference--group-003.md#canonical-2201330001003130-1103011331220233-0222132032022310-0202130110120123-3313013301131232-2100012011220310-2311130330033130-0310300011033013) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](resources--virtual_host--reference--group-003.md#canonical-2322030231212021-2033120211323221-0300330333313200-2030313333202300-0101103133202310-2102020323122011-3120212302200022-0213121113110013) |
| `tls_parameters.no_client_certificate` | [tls_parameters.no_client_certificate](resources--virtual_host--reference--group-003.md#canonical-2323112302311330-1200313313203222-0323031221031333-3111020121011033-0233231302132131-3303333311000121-0000010300221333-2213020222010031) |
| `tls_parameters.xfcc_header_elements` | [tls_parameters.xfcc_header_elements](resources--virtual_host--reference--group-002.md#canonical-0213200310122130-0101300331002010-3102233222333003-2122132330021030-3201110012210001-3122300102120002-0031012231012200-2300100230221000) |
| `user_identification` | [user_identification](resources--virtual_host--reference--group-003.md#canonical-3123231102032223-3202112001213220-2323200123223203-3121212301113113-1210231333210211-1123211210111022-3223021311311112-0221103332030102) |
| `user_identification.kind` | [user_identification.kind](resources--virtual_host--reference--group-003.md#canonical-0120222111310222-0320333030121310-1102300020313231-1030130301033323-3221101111030120-1212213333213300-1203123010003020-0211031003300033) |
| `user_identification.name` | [user_identification.name](resources--virtual_host--reference--group-003.md#canonical-2120122131223012-0223301301103212-3223131020130310-0232232310300032-2221002122032323-3232211101011112-0231223121212011-3211100213131103) |
| `user_identification.namespace` | [user_identification.namespace](resources--virtual_host--reference--group-003.md#canonical-0203221000130021-1111003301123201-1103331311111102-2111332110120302-3102113223330131-1213123230022303-1312333332002220-3203320022133102) |
| `user_identification.tenant` | [user_identification.tenant](resources--virtual_host--reference--group-003.md#canonical-1101222132001233-2111020200313002-1001210010013223-0010312130211201-0113130311031323-0311031031232233-2133011202013233-2210300023310220) |
| `user_identification.uid` | [user_identification.uid](resources--virtual_host--reference--group-003.md#canonical-0322121032330201-3201023101112312-2101102133132312-3332212232010300-2201212200202233-1002222210023032-2021100121233103-0332130232003102) |
| `waf_type` | [waf_type](resources--virtual_host--reference--group-003.md#canonical-3230013003023000-1033211112313122-3023101303102010-3313122122303000-1202210222021131-1031330201322010-0303332312111121-3302321220210320) |
| `waf_type.app_firewall` | [waf_type.app_firewall](resources--virtual_host--reference--group-003.md#canonical-3213022012121222-0002330131002030-3310033233203333-1033330211110330-2120011000001330-2201021020201233-0332123000320320-2212020330301333) |
| `waf_type.app_firewall.app_firewall` | [waf_type.app_firewall.app_firewall](resources--virtual_host--reference--group-003.md#canonical-2123223303123331-0321031133303331-2022221030110212-0301330033003311-1012333203333123-3230110301020030-1221011102120001-0210201321321333) |
| `waf_type.app_firewall.app_firewall.kind` | [waf_type.app_firewall.app_firewall.kind](resources--virtual_host--reference--group-003.md#canonical-0010032012013331-0303102231333212-3201133131010023-3133001021110200-2202200133000320-0130230102311131-1323223223130000-2101322301322200) |
| `waf_type.app_firewall.app_firewall.name` | [waf_type.app_firewall.app_firewall.name](resources--virtual_host--reference--group-003.md#canonical-1200031333031012-0023302213113102-1121212310312203-3310320333231022-0321110113222010-2030302020120332-1303232230333001-1010322323123100) |
| `waf_type.app_firewall.app_firewall.namespace` | [waf_type.app_firewall.app_firewall.namespace](resources--virtual_host--reference--group-003.md#canonical-2031123111321331-3332320323132012-3232332013331211-1212331111033220-0003212320322001-1121312120331231-1222212232220321-2313321203022032) |
| `waf_type.app_firewall.app_firewall.tenant` | [waf_type.app_firewall.app_firewall.tenant](resources--virtual_host--reference--group-003.md#canonical-3101322222023123-0001313303001031-0033203333000220-0003233001000020-2212202003310100-1312320202023303-1201312112203202-0131222003332121) |
| `waf_type.app_firewall.app_firewall.uid` | [waf_type.app_firewall.app_firewall.uid](resources--virtual_host--reference--group-003.md#canonical-1123303212111101-0231123003302020-1031321303000110-3122101313300322-0030111332130222-1221123331013331-2012120220012013-3332131100221002) |
| `waf_type.disable_waf` | [waf_type.disable_waf](resources--virtual_host--reference--group-003.md#canonical-0302310212030013-0311333130001103-0122210310011201-1201023311102112-0230033112133120-3001212100112231-3123023021221320-0220210212310210) |
| `waf_type.inherit_waf` | [waf_type.inherit_waf](resources--virtual_host--reference--group-003.md#canonical-0011131120230110-2322330020322312-1011301301130303-1223322210131003-0221300000132021-1302213300322220-0231320323010231-2321222020130331) |

<a id="canonical-3101233030000031-0230120000021132-0100111102033122-2131330301131003-3313323011111033-1133231010010122-2231321131221000-2223022121212113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_policies` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- advertise_policies

<a id="canonical-0320003333012113-1133301102110223-2031213123331300-3111030302111233-2110111202320311-3330220130131102-0130210011213222-0331122101131101"></a>

Type: `"object"`. list nested block, Optional.

Advertise Policy allows you to define networks or sites where you want a VIP for this virtual host
to be advertised. Each Policy rule can have different parameters, like TLS configuration, ports,
optionally IP address to be used for VIP. If advertise policy is not specified then no VIP is
assigned for this virtual host.

Receipt-pinned upstream constraints:

```json
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
advertise_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033132303303230-1023031332322201-3222112202220233-3323000011301202-1331010210220003-1203303133011000-2200200320311131-1323111313302030"></a>

### Direct properties for `advertise_policies`

<a id="canonical-2032121022320021-1130023200022111-3012032332203210-3332000323131031-2212022222003131-1232333200210103-1230223313301002-3200210003130330"></a>

#### `advertise_policies.kind` property

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

<a id="canonical-0033032333302201-1023211013311110-3121210233302220-1223102031110021-0310230013223321-3311121033101203-3122100312101311-0332210230230133"></a>

<a id="canonical-1211211132232330-3123010220010303-1022232223212313-3113030013033113-0233022110032213-2021330130012220-2230022013213012-3330012100301201"></a>

#### `advertise_policies.name` property

Type: `"string"`. Optional.

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

<a id="canonical-3333301100211030-2032331211212030-2130131202003121-2323323010310200-3323220110122011-3323010111230321-2101203331133210-1222010221303233"></a>

<a id="canonical-1200032210111123-0113000332130312-2102113000301031-3322203230103033-1200230211331303-2223123011220003-3300123302121121-1300330110230012"></a>

#### `advertise_policies.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-3121020321100200-1132203212213010-0222103313322323-0330303001022302-1023332230221213-2322032220021213-2133111232103103-0313112013201112"></a>

<a id="canonical-0220123201110330-1202302021120012-2100133023222133-2311002312332312-2213102221020303-0033223102301310-1210300212033031-0122102000320022"></a>

#### `advertise_policies.tenant` property

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

<a id="canonical-3303232113332101-0322313303002201-0123111220331110-2221132010211312-1133233110302002-0211110131101133-3202030300120330-1123022332210210"></a>

<a id="canonical-0331020133102122-2320001311023232-2233203233232001-0222332010031230-3122323020121301-0223332223130220-3312223322332031-0323311033120220"></a>

#### `advertise_policies.uid` property

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

<a id="canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `authentication` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- authentication

<a id="canonical-1211203032111231-3201133300013320-1000022233332122-2130221123031222-0121201132233320-3033233323321303-1211323003132102-1332111023313133"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: authentication, no\_authentication; Default: no\_authentication\] Authentication related
information. This allows to configure the URL to redirect after the authentication Authentication
Object Reference, configuration of cookie params etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("auth_config"),
  validators.ConflictingObjectAttributes("cookie_params",
    "use_auth_object_config"),
  validators.ConflictingObjectAttributes("redirect_dynamic",
    "redirect_url")}
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
  "x-ves-oneof-field-cookie_params_choice": "[\"cookie_params\",\"use_auth_object_config\"]",
  "x-ves-oneof-field-redirect_url_choice": "[\"redirect_dynamic\",\"redirect_url\"]"
}
```

OneOf alternatives in this subsection:

- [authentication](resources--virtual_host--reference--group-001.md#canonical-1211203032111231-3201133300013320-1000022233332122-2130221123031222-0121201132233320-3033233323321303-1211323003132102-1332111023313133)
- [no_authentication](resources--virtual_host--reference--group-002.md#canonical-0103231132201100-0003112003220201-1232112230200001-0013023021010100-2023323122300101-0222212120020221-3323111231232231-2003131212032321)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
authentication {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030320001303002-0030222031232022-1120200210231022-2233220111000020-0012320300033030-0222223311100033-0321210100213022-3121230212322323"></a>

### Direct properties for `authentication`

- [auth_config](resources--virtual_host--reference--group-001.md#canonical-1131122000023211-0323333123322310-2001311023102220-0321330310211031-1320221233022310-3013233100022332-0123001323131130-3112213333031223): complete subsection reference.

- [cookie_params](resources--virtual_host--reference--group-001.md#canonical-1131131321312220-2012123000002311-0103102321232020-3331020223222330-3002023112000031-1310331020230210-2120111121232302-2012200122312121): complete subsection reference.

- [redirect_dynamic](resources--virtual_host--reference--group-001.md#canonical-0313011103023132-2300102103031013-0232223211222030-2120102003032023-2110331132332321-1013021312110232-1200201333111012-2132031021223100): complete subsection reference.

<a id="canonical-2222102131030220-3310022313321300-3032123000123010-1213222111222003-1312131313032203-3211203102033213-3012020012200110-1031100112230312"></a>

<a id="canonical-2333230030013201-3322012233301132-0220002220011003-1113110210303223-2133110333103002-0112133313011031-1233322300021010-2220101112002000"></a>

#### `authentication.redirect_url` property

Type: `"string"`. Optional.

Exclusive with \[redirect\_dynamic\] user can provide a URL for e.g https&#58;//abc.xyz.com where
user gets redirected. This URL configured here must match with the redirect URL configured with the
OIDC provider.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [use_auth_object_config](resources--virtual_host--reference--group-001.md#canonical-1302322221332101-3012223212012310-0220333220112221-3102201210323222-3130311012220201-1131210331201223-3021101112131200-3001112123033302): complete subsection reference.

<a id="canonical-1131122000023211-0323333123322310-2001311023102220-0321330310211031-1320221233022310-3013233100022332-0123001323131130-3112213333031223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `authentication.auth_config` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123)
- authentication.auth_config

<a id="canonical-3222333202210002-1332002110001222-2321131000322220-0021330121013322-1012321113002202-1201022312333312-0301332332020211-0300331331113111"></a>

Type: `"object"`. list nested block, Optional.

Reference to Authentication Config Object.

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

Terraform syntax:

```terraform
auth_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103221313130103-1002333203221333-0223312223321033-3122203303330020-2311113102100032-0211033100301102-2032220200221103-0010013123203312"></a>

### Direct properties for `authentication.auth_config`

<a id="canonical-3233213030322030-0121020202133233-0101100210310322-0210203012223101-0311122012310122-2330303130323210-0333110300120000-0221130232133000"></a>

#### `authentication.auth_config.kind` property

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

<a id="canonical-2210323323232133-0111330013012131-3203130303203230-1212100121113023-0233032210103000-3301111130123022-2011323020222211-2301023312103223"></a>

<a id="canonical-2131110332011320-1013130033133131-2310300332022310-0023300333021131-1232200030020312-2221132223320331-1033002233320101-2122333002320013"></a>

#### `authentication.auth_config.name` property

Type: `"string"`. Optional.

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

<a id="canonical-2332312331312121-2211210330311130-3012301003100222-2200110222100223-1221300331111011-0022100212130003-3330003202032322-3311233213112122"></a>

<a id="canonical-3133300312013320-0311213223323020-3012111132133201-2131020020220303-2123001231333301-2330102021133030-1311223201332100-0023223013121022"></a>

#### `authentication.auth_config.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-3110233221101300-1302130133021311-0330000101300320-0112320303202201-0333032332103102-2220321211213233-3131200222321323-1131121312012211"></a>

<a id="canonical-2021311330320111-1131333121212111-2300322113133330-1312232132233112-1321032033320031-1313123202212322-2132133130120200-0003221300101123"></a>

#### `authentication.auth_config.tenant` property

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

<a id="canonical-2132011022102102-1031310123330301-1012112212032323-0313031223000223-0121322221121120-1012001233012132-1302021001023032-2210033000323001"></a>

<a id="canonical-1230320001001333-0211220330001321-1111302020330032-3220121221233001-1132121032213101-2023021311010031-3231021303022123-2230132113323333"></a>

#### `authentication.auth_config.uid` property

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

<a id="canonical-1131131321312220-2012123000002311-0103102321232020-3331020223222330-3002023112000031-1310331020230210-2120111121232302-2012200122312121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `authentication.cookie_params` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123)
- authentication.cookie_params

<a id="canonical-2200302311133111-0102132023022310-1301212120132130-3221133002312210-3300320202132332-2331020133223300-1003323112320120-1123220213023321"></a>

Type: `"object"`. single nested block, Optional.

Specifies different cookie related config parameters for authentication.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auth_hmac",
    "kms_key_hmac")}
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
  "x-ves-oneof-field-secret_choice": "[\"auth_hmac\",\"kms_key_hmac\"]"
}
```

Terraform syntax:

```terraform
cookie_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1103131031312023-3330332231030311-3011320303032301-0303031230030023-2233222211233302-1200131321222030-0130311031113231-1330322220110323"></a>

### Direct properties for `authentication.cookie_params`

- [auth_hmac](resources--virtual_host--reference--group-001.md#canonical-0112101301310321-2022233203002122-3032312032133321-3013000233012113-0013123320312230-0321123320333001-3102301133231301-3002231230330232): complete subsection reference.

<a id="canonical-3332332303002333-0000320313200310-0021133232332301-1202123301110113-3333110120102220-3002230121221301-0323010100133330-0033002310102333"></a>

<a id="canonical-0210010132313223-2323122212021221-0213311110013030-1112231301300320-2331230200312113-2333333332023121-1311301231110310-0212201100103002"></a>

#### `authentication.cookie_params.cookie_expiry` property

Type: `"number"`. Optional.

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client-side after which client will not
be setting the cookie as part of the request. Default cookie expiry is 3600 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(86400),
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-1000232101101000-1030122230120331-3033120323031312-1122212323210023-1132131030113231-0300220222132003-0113213210032300-0032301330002303"></a>

<a id="canonical-1112122333012301-2033033123120121-0321100233120103-1211032033110233-0123023211203013-3202333022123320-1212021211223302-0120022230333030"></a>

#### `authentication.cookie_params.cookie_refresh_interval` property

Type: `"number"`. Optional.

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session expiry. Default refresh interval is 3000 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(86400),
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

- [kms_key_hmac](resources--virtual_host--reference--group-001.md#canonical-2230033003322103-2333330222332211-1312331010131331-3103000120330311-3302023123120232-3232131132303011-0011333001101110-3203320201201122): complete subsection reference.

<a id="canonical-1212232133020110-1303101001003022-3321020102020130-3232203001330210-0121103020311110-0033112022301121-2020210030003333-3333022203322200"></a>

<a id="canonical-1213123003130311-3300032120223232-2332330331102203-0130013023223310-2031032330120010-3113113131010100-0323233222033330-1222110100203133"></a>

#### `authentication.cookie_params.session_expiry` property

Type: `"number"`. Optional.

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1296000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1296000,
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
    "ves.io.schema.rules.uint32.lte": "1296000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1296000"
  }
}
```

<a id="canonical-0112101301310321-2022233203002122-3032312032133321-3013000233012113-0013123320312230-0321123320333001-3102301133231301-3002231230330232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `authentication.cookie_params.auth_hmac` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-1131131321312220-2012123000002311-0103102321232020-3331020223222330-3002023112000031-1310331020230210-2120111121232302-2012200122312121)
- authentication.cookie_params.auth_hmac

<a id="canonical-1022310232123201-1222012330120313-0312323333230122-0222212201233030-3220103112103323-0312232123331002-3213312013312002-1220201303101120"></a>

Type: `"object"`. single nested block, Optional.

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prim_key_expiry",
    "sec_key_expiry")}
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
auth_hmac {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310333310122333-0332303123230333-3210102211101113-0200010202200013-1010110201200113-0021321010312331-2120301131322133-3002303212112120"></a>

### Direct properties for `authentication.cookie_params.auth_hmac`

- [prim_key](resources--virtual_host--reference--group-001.md#canonical-3200010310233203-1210331030223203-1011131031030201-0000131311322302-3132300210320223-0302302113112202-3012111312112213-3221330010132221): complete subsection reference.

<a id="canonical-1320121102131231-2003222332232201-2332202010001321-2211201210210330-2221031031123032-3321001321210303-0120022213132213-1123213131020123"></a>

<a id="canonical-2032121211220130-0012111312102123-2033000003012310-3111122321012112-1201013332002100-3123103302020000-3002131030211200-3002200130210230"></a>

#### `authentication.cookie_params.auth_hmac.prim_key_expiry` property

Type: `"string"`. Optional.

HMAC Primary Key Expiry. Primary HMAC Key Expiry time.

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

- [sec_key](resources--virtual_host--reference--group-001.md#canonical-1020112220213021-2013010221201110-1201030001102333-3203321222232303-1312012322330201-2030010101010333-0233031002023333-0022232113021301): complete subsection reference.

<a id="canonical-3013313201001112-1332312332333222-2302121022112300-0333130100301133-0131121031321110-3033013003133222-1331303313311022-2220310230333302"></a>

<a id="canonical-2313021032322013-3122013321213103-0021200012332030-1131211021020321-0010321330113112-0010201223230320-3012110130113002-1323010212123132"></a>

#### `authentication.cookie_params.auth_hmac.sec_key_expiry` property

Type: `"string"`. Optional.

HMAC Secondary Key Expiry. Secondary HMAC Key Expiry time.

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

<a id="canonical-3200010310233203-1210331030223203-1011131031030201-0000131311322302-3132300210320223-0302302113112202-3012111312112213-3221330010132221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `authentication.cookie_params.auth_hmac.prim_key` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-1131131321312220-2012123000002311-0103102321232020-3331020223222330-3002023112000031-1310331020230210-2120111121232302-2012200122312121)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-0112101301310321-2022233203002122-3032312032133321-3013000233012113-0013123320312230-0321123320333001-3102301133231301-3002231230330232)
- authentication.cookie_params.auth_hmac.prim_key

<a id="canonical-0213112010330011-2030120013020000-0113322020302031-0331312000133020-2231003100331300-2231301101310120-2303212311030003-2020010132012332"></a>

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
prim_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322312121013233-3313330010112123-3032031111203331-2303100011212120-2002233102030313-1032230031122313-1101011032302201-3331000033203221"></a>

### Direct properties for `authentication.cookie_params.auth_hmac.prim_key`

- [blindfold_secret_info](resources--virtual_host--reference--group-001.md#canonical-1203203310302023-1202221202223330-2103231310322213-1203323003120213-2210312203020110-2112030003320333-1033103222120303-1023112100110300): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-001.md#canonical-1312203201220310-0110323230102332-3300201310110122-0030123202333000-0330230313213330-1002230331333323-2033212010133230-2021031201101021): complete subsection reference.

<a id="canonical-1203203310302023-1202221202223330-2103231310322213-1203323003120213-2210312203020110-2112030003320333-1033103222120303-1023112100110300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-1131131321312220-2012123000002311-0103102321232020-3331020223222330-3002023112000031-1310331020230210-2120111121232302-2012200122312121)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-0112101301310321-2022233203002122-3032312032133321-3013000233012113-0013123320312230-0321123320333001-3102301133231301-3002231230330232)
- [authentication.cookie_params.auth_hmac.prim_key](resources--virtual_host--reference--group-001.md#canonical-3200010310233203-1210331030223203-1011131031030201-0000131311322302-3132300210320223-0302302113112202-3012111312112213-3221330010132221)
- authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info

<a id="canonical-3010333000322033-1223023321231232-1131023033120132-1132211111223001-2110021030112022-1022000131130320-0020220010332133-3201203012112012"></a>

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

<a id="canonical-3210221012200132-1101113311320032-2300231213300103-3311100132011130-3212130122333120-2233232303120300-1302112030021133-2310330330032032"></a>

### Direct properties for `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info`

<a id="canonical-2002332211202102-2102212131110320-2330101012213032-2132000311313130-3110201012211030-3322222200202223-0033231301201220-2320011032010323"></a>

#### `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0030220120311200-2331002111032301-2003200113132032-1022311232130320-1110330000203122-1111302100222331-3003200001333303-0020232302210131"></a>

<a id="canonical-0231021312202100-2123300313112200-0222130233302122-0023133113010023-1201122130130100-1233111023213030-2223121100132030-0101321213023330"></a>

#### `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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

<a id="canonical-3113112131010121-1002000231210110-3200012001322231-1300013223330312-3321023210123311-1312101110130222-3223202200002110-2122321322111021"></a>

<a id="canonical-1321121230202122-3210031210023210-0003030022013213-2122100303333222-3212331322122323-0103132111122201-3210200010032233-3232103022010131"></a>

#### `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-1312203201220310-0110323230102332-3300201310110122-0030123202333000-0330230313213330-1002230331333323-2033212010133230-2021031201101021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-1131131321312220-2012123000002311-0103102321232020-3331020223222330-3002023112000031-1310331020230210-2120111121232302-2012200122312121)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-0112101301310321-2022233203002122-3032312032133321-3013000233012113-0013123320312230-0321123320333001-3102301133231301-3002231230330232)
- [authentication.cookie_params.auth_hmac.prim_key](resources--virtual_host--reference--group-001.md#canonical-3200010310233203-1210331030223203-1011131031030201-0000131311322302-3132300210320223-0302302113112202-3012111312112213-3221330010132221)
- authentication.cookie_params.auth_hmac.prim_key.clear_secret_info

<a id="canonical-1200122123003332-0032032231222311-0130120303312313-2231300032201233-2301102000021020-3322300323231233-3202321311221122-3230222321232010"></a>

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

<a id="canonical-2031113332303012-3012123332133032-1010130020021312-2301213303120122-1212312210131021-3131102333211112-1002333301213201-1100312320300231"></a>

### Direct properties for `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info`

<a id="canonical-0301130100122013-1221203201203110-2013220110132012-3221312100022211-0223320222201322-3231013132223020-2113332020013000-2110023122133233"></a>

#### `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1110002222212002-2100020303301232-2211231002020130-1102231203300321-2202021023012330-1122111313232322-0323020023231101-0112032320213112"></a>

<a id="canonical-2103111111011003-3301013120111022-0200321003321012-1301223331102323-2023201003321321-1133103112330212-3210101021200020-1023230301321010"></a>

#### `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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

<a id="canonical-1020112220213021-2013010221201110-1201030001102333-3203321222232303-1312012322330201-2030010101010333-0233031002023333-0022232113021301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `authentication.cookie_params.auth_hmac.sec_key` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-1131131321312220-2012123000002311-0103102321232020-3331020223222330-3002023112000031-1310331020230210-2120111121232302-2012200122312121)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-0112101301310321-2022233203002122-3032312032133321-3013000233012113-0013123320312230-0321123320333001-3102301133231301-3002231230330232)
- authentication.cookie_params.auth_hmac.sec_key

<a id="canonical-0221121003332310-2133111123021002-3111013321013131-3121223101011301-2213330003013103-3321112212021021-1223012220233231-2303212300031130"></a>

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
sec_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130310110311203-2210233320120300-3100231301332020-3002201100331103-3101321002320020-0001302002223130-3103303032023203-2201211123303030"></a>

### Direct properties for `authentication.cookie_params.auth_hmac.sec_key`

- [blindfold_secret_info](resources--virtual_host--reference--group-001.md#canonical-3113131201300102-2023030131011101-3132321133033312-2211303211003103-2233313203010322-1112021303203003-3103330300320111-3331000131123222): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-001.md#canonical-2323223221310303-3223203232303030-3202113200031133-1033322321332131-2301000332102220-0103002132021210-3301303213233013-0022223213011132): complete subsection reference.

<a id="canonical-3113131201300102-2023030131011101-3132321133033312-2211303211003103-2233313203010322-1112021303203003-3103330300320111-3331000131123222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-1131131321312220-2012123000002311-0103102321232020-3331020223222330-3002023112000031-1310331020230210-2120111121232302-2012200122312121)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-0112101301310321-2022233203002122-3032312032133321-3013000233012113-0013123320312230-0321123320333001-3102301133231301-3002231230330232)
- [authentication.cookie_params.auth_hmac.sec_key](resources--virtual_host--reference--group-001.md#canonical-1020112220213021-2013010221201110-1201030001102333-3203321222232303-1312012322330201-2030010101010333-0233031002023333-0022232113021301)
- authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info

<a id="canonical-3021101023312221-2333311213130231-3220231221130022-1012023321110310-0222313020200321-3301332022333032-2303232232131122-1121211332311203"></a>

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

<a id="canonical-2002112331030221-2221203123232010-1313230223233220-3212111310220022-2302211120202303-1210302202122320-2110300313013201-3001223322000213"></a>

### Direct properties for `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info`

<a id="canonical-0200013301031002-3322001010032320-2233312022102303-2130303131110132-1310121302233302-3033232302013002-2322012101010000-3311313333200233"></a>

#### `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3311123302003210-2202021332312110-2030321000320112-0301311000131102-0302020131033330-1101121003020033-0130210101100001-0031232311300022"></a>

<a id="canonical-2300221121121301-1130013030100310-2131013130331103-2000011202133331-2212220330311122-1213100113122021-0301232223123120-2132113303312233"></a>

#### `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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

<a id="canonical-1221203101000022-1201320132321031-0211013300202103-3233302332221333-3132302333220000-2323012302123103-0030031202311203-3031030121313330"></a>

<a id="canonical-0212213201220232-2231331302032033-3201222300023013-2031232302123302-3322102023012020-1100313032300012-3032121022223011-0122220033210102"></a>

#### `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-2323223221310303-3223203232303030-3202113200031133-1033322321332131-2301000332102220-0103002132021210-3301303213233013-0022223213011132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-1131131321312220-2012123000002311-0103102321232020-3331020223222330-3002023112000031-1310331020230210-2120111121232302-2012200122312121)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-0112101301310321-2022233203002122-3032312032133321-3013000233012113-0013123320312230-0321123320333001-3102301133231301-3002231230330232)
- [authentication.cookie_params.auth_hmac.sec_key](resources--virtual_host--reference--group-001.md#canonical-1020112220213021-2013010221201110-1201030001102333-3203321222232303-1312012322330201-2030010101010333-0233031002023333-0022232113021301)
- authentication.cookie_params.auth_hmac.sec_key.clear_secret_info

<a id="canonical-0130223210022330-3133123331332110-0132113210021321-1312032101032211-0300220210002203-3133021023110203-0130221322221322-3202013221133103"></a>

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

<a id="canonical-0330021013110232-0003211312311310-2011122012010303-1013121211031221-2302101010321110-0331112101122310-3113033030303001-3213322313002221"></a>

### Direct properties for `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info`

<a id="canonical-0220310313202223-3022132133103130-2231000231300131-1200202032001211-3021113331210133-1020130013031112-3312130103303010-3312102123222221"></a>

#### `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1212011022331301-2032012333210310-1101301101203132-2231021320123300-0132012321133131-3232012311220100-1302033323220201-3103032313223311"></a>

<a id="canonical-2000101333311010-2220112132013012-1220330113101332-2231133202013322-0022030223111002-3113002130120113-0303230000012103-1301230120131013"></a>

#### `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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

<a id="canonical-2230033003322103-2333330222332211-1312331010131331-3103000120330311-3302023123120232-3232131132303011-0011333001101110-3203320201201122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `authentication.cookie_params.kms_key_hmac` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-1131131321312220-2012123000002311-0103102321232020-3331020223222330-3002023112000031-1310331020230210-2120111121232302-2012200122312121)
- authentication.cookie_params.kms_key_hmac

<a id="canonical-1321333011213012-2012333300202103-3313300131313230-2031022333132100-1000323212331013-1121300333313313-0111213223021001-3230332330100111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for kms key hmac.

Additional upstream details:

Reference to KMS Key Object.

Receipt-pinned upstream constraints:

```json
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
kms_key_hmac = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313011103023132-2300102103031013-0232223211222030-2120102003032023-2110331132332321-1013021312110232-1200201333111012-2132031021223100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `authentication.redirect_dynamic` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123)
- authentication.redirect_dynamic

<a id="canonical-3213102103320201-1032210310101101-1212233222123001-3120123321130021-1013220321200123-2200302302223310-3111023123021303-1320003001230120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for redirect dynamic.

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
redirect_dynamic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302322221332101-3012223212012310-0220333220112221-3102201210323222-3130311012220201-1131210331201223-3021101112131200-3001112123033302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `authentication.use_auth_object_config` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123)
- authentication.use_auth_object_config

<a id="canonical-0210222011020122-3312120333231130-0123030202102320-0131112000013232-1001111002202030-0031210011002221-3011120221110231-0230321030201102"></a>

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
use_auth_object_config = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3232012230212003-3221232011020313-0202200302301332-0002323023102302-2322221023103201-3002223222331300-0233232032310121-2021012312213213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `buffer_policy` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- buffer_policy

<a id="canonical-3011223012232200-0033131112212212-0311032201323010-1230133111002012-1210000023123031-1120322321013011-3010101320312231-0021000020120100"></a>

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

<a id="canonical-1232133130130120-0110120302112113-1003230223220100-3322300213322003-0223022320130221-3020301113120333-3211323022032310-2210123321312312"></a>

### Direct properties for `buffer_policy`

<a id="canonical-1121202302021000-1122211120311322-0323030302130113-1212010331301331-0332310323322023-2231003333031010-2210031112120000-2102301030203023"></a>

#### `buffer_policy.disabled` property

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

<a id="canonical-1002323221030301-3223121130300123-0323212120223032-0131122231312320-2000021101211122-0031223322313323-1030133030303003-3203230222023012"></a>

<a id="canonical-3030010320310032-0023232220000023-2010331233023312-3302103322310003-2301110332121200-1100212213222000-1231032211233332-2201101301123111"></a>

#### `buffer_policy.max_request_bytes` property

Type: `"number"`. Optional.

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

<a id="canonical-1020312231213113-0223120010100113-3031300332310212-2210330020033311-2212131302230022-1310032130223312-2033030030001301-2312002230010203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `captcha_challenge` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- captcha_challenge

<a id="canonical-1301312132130220-0131231003121113-3122120112223111-3123000301131222-1330033022112332-1210203020031031-3310220022023231-0120030021010211"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: captcha\_challenge, js\_challenge, no\_challenge; Default: no\_challenge\] Enables
loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With
this feature enabled, only clients that pass the captcha challenge will be allowed to complete the
HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect..

Additional upstream details:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha. When loadbalancer is configured to do Captcha
Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page
will have captcha challenge embedded in it. Client will be allowed to make the request only if the
captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid
Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only
human users can pass through. Generally, computers or bots are not capable of solving a captcha. You
can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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

OneOf alternatives in this subsection:

- [captcha_challenge](resources--virtual_host--reference--group-001.md#canonical-1301312132130220-0131231003121113-3122120112223111-3123000301131222-1330033022112332-1210203020031031-3310220022023231-0120030021010211)
- [js_challenge](resources--virtual_host--reference--group-002.md#canonical-2301110332123321-2201100001322312-3310232321120302-2100202013231333-1032213203023203-3311131310311320-2313310010120313-0030132221110033)
- [no_challenge](resources--virtual_host--reference--group-002.md#canonical-1023032113232030-0020323313122121-0322301020011111-0223320231112323-1001102230211202-1023032220132002-0312223223230123-1312013321231030)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
captcha_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-3033001131302312-0203103010332113-0113320133322033-1020303121310112-0032232222110233-0000212021001031-1332223323033202-0003301331032330"></a>

### Direct properties for `captcha_challenge`

<a id="canonical-0320001301322310-1312032101020021-0320302330321213-0330303223323021-1330102102330330-3220302130112121-1323332330102010-3323213100100103"></a>

#### `captcha_challenge.cookie_expiry` property

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
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

<a id="canonical-3211020102201311-2010132132102223-2033110322132322-0313311233122022-3000211323330210-2013021010201102-2132120101020303-0111330130133121"></a>

<a id="canonical-3202020321032131-2230031023020021-3202123333101032-3023033023031330-1120120003200223-0200301113331232-0211210311101031-3201131101020123"></a>

#### `captcha_challenge.custom_page` property

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

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

<a id="canonical-1133103011022022-2203323101231031-2333022101312310-0311232221002132-1112210132302130-3101211220023233-3011302023021202-3033100231101311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `coalescing_options` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- coalescing_options

<a id="canonical-3120121021011122-3011303310310313-1322223303233223-1001000133333132-3332113200330200-2022103003121300-0132333332101302-1122132031033100"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210231101321130-2211033202020312-0033112222130302-0032202332332112-0200122133223001-0203021201112031-2322203311022120-3211212210313121"></a>

### Direct properties for `coalescing_options`

- [default_coalescing](resources--virtual_host--reference--group-001.md#canonical-2012121230122021-0023032213020033-1233303132320033-3122023311001213-2010021213012031-1222000331003313-1020200120230101-2012231101333203): complete subsection reference.

- [strict_coalescing](resources--virtual_host--reference--group-001.md#canonical-2212313210111032-2123210112200322-0120220112031321-2131321210130301-1101211101323033-3322033022302121-3313210301020200-0003230232313030): complete subsection reference.

<a id="canonical-2012121230122021-0023032213020033-1233303132320033-3122023311001213-2010021213012031-1222000331003313-1020200120230101-2012231101333203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [coalescing_options](resources--virtual_host--reference--group-001.md#canonical-1133103011022022-2203323101231031-2333022101312310-0311232221002132-1112210132302130-3101211220023233-3011302023021202-3033100231101311)
- coalescing_options.default_coalescing

<a id="canonical-2022310332331113-2010031031010312-3311022321032023-3113223221113202-2022330302223001-2112132123113031-3032213302301111-0231100312011013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

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
default_coalescing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212313210111032-2123210112200322-0120220112031321-2131321210130301-1101211101323033-3322033022302121-3313210301020200-0003230232313030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [coalescing_options](resources--virtual_host--reference--group-001.md#canonical-1133103011022022-2203323101231031-2333022101312310-0311232221002132-1112210132302130-3101211220023233-3011302023021202-3033100231101311)
- coalescing_options.strict_coalescing

<a id="canonical-2132210221002320-3201211031212203-3032331033323023-3333321331100021-2000323000001233-1230010301133312-2303310330203210-1113123032112131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

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
strict_coalescing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033033121221113-3201013233022330-2031033001101131-2231201001233133-3212322221010320-2231221131030022-1003330210110233-1333033203013221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `compression_params` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- compression_params

<a id="canonical-0031211031111203-2132322102121311-2321132201121110-3323300211311302-2310031121311103-3210323003130202-2310301301000301-2121120102321013"></a>

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

<a id="canonical-1000011020300310-1112030003313203-1033132113231003-0212112200031021-2002210221103100-0003132013203103-0123213311110110-0303331102013232"></a>

### Direct properties for `compression_params`

<a id="canonical-0301021121011203-2323321100122010-0312130302110121-0322111010002312-1200320133020202-2000031122322120-2033201132323320-2303300123022312"></a>

#### `compression_params.content_length` property

Type: `"number"`. Optional.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Additional upstream details:

The default value is 30.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0110132312321221-1110032302010002-2312121020310022-2320101313301030-3022032013003330-0331320020113310-1200113122101131-0231122212310200"></a>

<a id="canonical-3223022223101101-0021112131302032-3110200312200111-1013201030302113-2312332203123212-3000201303203211-3222002203112232-1112231320020200"></a>

#### `compression_params.content_type` property

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

<a id="canonical-3111303120201001-3203301320332302-1000312300113310-1323111201220011-0032232010122031-0301220301003023-3110013322121101-1112330020122213"></a>

<a id="canonical-1222211031033301-1231033011222001-3332232211021203-2211103013220222-1112103012213211-1303131313220012-0130220312030201-1112200112332320"></a>

#### `compression_params.disable_on_etag_header` property

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

<a id="canonical-0121010310031311-1130223200113313-1221313013121223-0220202022102322-1103303312311231-3323221000133302-0013202321201021-0212102313223101"></a>

<a id="canonical-2302210001020001-1012103131310333-3302122121132222-1231003111321113-1133333001230112-3120103000333003-2121000230132200-0211020210033322"></a>

#### `compression_params.remove_accept_encoding_header` property

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

<a id="canonical-3232131223002131-1100030322001210-3211031031233033-0203330020202031-0110112203211203-1312221320131310-3120133001121223-0213203230311301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cors_policy` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- cors_policy

<a id="canonical-1130200103032202-0123310333002032-1000010331032332-1022233233012013-1223232030010332-2121121130102002-2200003223011003-3232203131113100"></a>

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

<a id="canonical-2112300312002210-2032203032313112-0011311223100033-1302022010131011-1231230333120021-3323003200211202-1113020331223020-2122020011030113"></a>

### Direct properties for `cors_policy`

<a id="canonical-3131233213331001-1113002311133000-0010121022010303-2000233123120112-2033312130133131-0022202203023011-2001221323202301-3111120112130010"></a>

#### `cors_policy.allow_credentials` property

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

<a id="canonical-0122013231310310-0212110121303122-3003000212303012-3002020212131011-2332310330202111-1121203201103010-0301130031310031-2231213130130230"></a>

<a id="canonical-0021111202322331-3102300003110102-3220010210323030-3102122201012230-0120103030121030-3312200311023022-0103023323131301-3211001223103010"></a>

#### `cors_policy.allow_headers` property

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

<a id="canonical-2101211112322121-2300103100213220-2111121132223013-3210022133013130-0310031211321202-3213231212032112-1121303133030202-3302211322201300"></a>

<a id="canonical-3202211223230032-0013203113203302-0131232323213220-2312011302223320-3021312103003330-3000003322233100-0133223013313201-1123112103020223"></a>

#### `cors_policy.allow_methods` property

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

<a id="canonical-1130202231311301-3330300031020110-1122101230020033-2002013310233100-2003011302022013-3102021130120221-3200021330312310-3000323213123331"></a>

<a id="canonical-3332321233212122-0110031212023102-0103300220202332-0120121032033123-2023123102312321-0102121012010023-2021132322202210-2201333131122310"></a>

#### `cors_policy.allow_origin` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-1022012010132013-3331001120023021-0133303321210303-1333133201123131-0233203312230322-3033022002111032-0021313003212012-1233133033133100"></a>

<a id="canonical-0330103122223212-3310223110130110-1301003112002003-0313311313100102-1211231322133332-2323332032233203-3300000100200122-1132012323112322"></a>

#### `cors_policy.allow_origin_regex` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-1221122011312203-1232112122311221-3102003313302112-2002103210122312-1033322011200322-1212132211213233-1301010312302333-1213220111310230"></a>

<a id="canonical-1132110232112222-3210130120220320-0223331322021322-2031122033200321-3022001123300000-2023101100232333-0333211211200002-0020033321000113"></a>

#### `cors_policy.disabled` property

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

<a id="canonical-0011322302112011-1223033011121000-0330231302300223-2000132022002001-3132212033231120-0030300313032222-0013223333123311-1032322202030233"></a>

<a id="canonical-0133123303212121-2320301301203011-1003331103111321-1230022012000331-1202203322311302-0203111200303021-1110031010321121-0121021200202111"></a>

#### `cors_policy.expose_headers` property

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

<a id="canonical-1211222333020113-1100002102312331-3121031203111333-2111001231213323-0303133212213102-0111112133311003-3002203200101121-3300300122112302"></a>

<a id="canonical-1202322331201020-0001220222221222-1120233120330332-1120003320103100-2133312231332312-1310131002322030-2230100221332313-0332201023120333"></a>

#### `cors_policy.maximum_age` property

Type: `"number"`. Optional.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

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

<a id="canonical-0032333010121002-0332331033300212-2002033020321220-2332202013303113-2313133231100331-0003222333003100-3133200232230203-0121213102202020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- csrf_policy

<a id="canonical-3000331031023322-3131110331302003-2320301022032003-1010232010300203-1103201100200333-0321313303002220-3023320100320230-1120102030310133"></a>

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

<a id="canonical-0223112302012100-1030233033310030-3101103101130100-1011020032312122-3321121302021133-1322010211011211-0321133100233111-3130031131331013"></a>

### Direct properties for `csrf_policy`

- [all_load_balancer_domains](resources--virtual_host--reference--group-001.md#canonical-0110122323300011-3202101110312103-1212210232132200-3123012203220303-0012003220313200-1100012030223133-3011220003003221-2131002333120213): complete subsection reference.

- [custom_domain_list](resources--virtual_host--reference--group-002.md#canonical-0210302321013100-1300220003111020-1330033301211103-3021331002113120-2311332011003021-2100020010022101-3320121011300333-3030122322323001): complete subsection reference.

- [disabled](resources--virtual_host--reference--group-002.md#canonical-1002103013312222-1311000201003223-0013030330132112-0012122333221101-2120131022311012-0030223010122223-0301322123230103-2210301232001223): complete subsection reference.

<a id="canonical-0110122323300011-3202101110312103-1212210232132200-3123012203220303-0012003220313200-1100012030223133-3011220003003221-2131002333120213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy.all_load_balancer_domains` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [csrf_policy](resources--virtual_host--reference--group-001.md#canonical-0032333010121002-0332331033300212-2002033020321220-2332202013303113-2313133231100331-0003222333003100-3133200232230203-0121213102202020)
- csrf_policy.all_load_balancer_domains

<a id="canonical-3133102211130012-2123322312022330-0023111120020320-1231211020211130-3320222020302212-3221122332320201-0312001003320213-0121333032231313"></a>

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
