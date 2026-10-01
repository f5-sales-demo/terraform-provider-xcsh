---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310220101331131-1332101033202130-2222131022002103-0220233012133201-0022010121033022-1103102033130013-3310010222021201-1133202203322202"></a>

## Property reference — Property reference / 002321300333 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- Property reference

<a id="canonical-2131220032103200-2333231331303311-0013111113112323-1022322211201212-2300331031033221-3322111231122222-1112222032132223-0331222001101132"></a>

## Direct properties — Property reference / 002321300333 / 3

<a id="canonical-1030131231003202-0332223001323000-3303213301132120-1322203002303212-3022000033211203-2201331120311013-0212231332111200-0003221011321112"></a>

<a id="canonical-2101313022001001-0313310233212310-0113321230333102-0200121100100110-0212223210010221-3221213213300121-1100003001001222-2113103233132121"></a>

## add_location property — Property reference / 002321300333 / 4

Type: `"bool"`. Computed.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses. This configuration is ignored on CE sites.

Upstream description:

X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; in responses.
This configuration is ignored on CE sites.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [advertise_policies](data-sources--virtual_host--reference--group-001.md#canonical-0100020130333221-2323111112123131-3120132302111122-1332303033101021-0310121032121322-1122013301131232-0312121331100133-1033121331110112): complete subsection reference.

<a id="canonical-2210320232301302-1231203000111220-2222222000321132-2101023221203301-3102210110333102-2232300312100112-2032133301000131-3021312023200233"></a>

<a id="canonical-0220113121033322-0112210302001231-2300300021203002-2213213223230120-3001020102202212-3213202332013121-3133330022330020-3222203321010123"></a>

## annotations property — Property reference / 002321300333 / 5

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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

<a id="canonical-0323113201300223-2222333123230213-2303032122303012-1102010201023100-1233332111220321-3010303110000300-2113333030020321-1122222213033203"></a>

<a id="canonical-0100030200131021-0122002230120321-3112332331023200-3001303303213230-0123001300221123-3123313310321333-0013033300332322-0133330000223133"></a>

## append_server_name property — Property reference / 002321300333 / 6

Type: `"string"`. Computed.

\[OneOf: append\_server\_name, default\_header, pass\_through, server\_name; Default:
default\_header\] Exclusive with \[default\_header pass\_through server\_name\] Specifies the value
to be used for Server header if it is not already present. If Server Header is already present it is
not overwritten. It is just passed.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Specifies the value to be used for
Server header if it is not already present. If Server Header is already present it is not
overwritten. It is just passed.

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

OneOf alternatives in this subsection:

- [append_server_name](data-sources--virtual_host--reference--group-001.md#canonical-0323113201300223-2222333123230213-2303032122303012-1102010201023100-1233332111220321-3010303110000300-2113333030020321-1122222213033203)
- [default_header](data-sources--virtual_host--reference--group-002.md#canonical-0131231003301103-0321312203033301-3002222033101103-1222201231020220-1232210300213310-0131212300130011-2002101030031123-0001030211230102)
- [pass_through](data-sources--virtual_host--reference--group-002.md#canonical-2011230100101232-2231130122231011-0101023133113300-2212113001232301-2201101303031211-2222001202121303-1333000311103303-2303303020030321)
- [server_name](data-sources--virtual_host--reference--group-001.md#canonical-2000133212001013-1130012132023121-0030132122333222-0303323311000221-0303100033122013-2021210110033131-2331333211132010-0220310022330223)

Select alternatives according to the provider validators above.

- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200): complete subsection reference.

- [buffer_policy](data-sources--virtual_host--reference--group-001.md#canonical-1213301223231000-1113301202120100-0322200112212300-3233130103220223-3123033012131112-0321102133021131-3201032002203320-3120102012111033): complete subsection reference.

- [captcha_challenge](data-sources--virtual_host--reference--group-002.md#canonical-3030320121132033-0113101012333213-2333033121300031-3102221202303210-3222022210330032-0103332033311333-0233333333103003-0313010030101011): complete subsection reference.

- [coalescing_options](data-sources--virtual_host--reference--group-002.md#canonical-1311333121020231-2023022200000300-2230331133331031-2300321203030302-1000030032223302-0132000331210013-0110211201102210-0011003111130011): complete subsection reference.

- [compression_params](data-sources--virtual_host--reference--group-002.md#canonical-0113113212120330-2002000130223321-2000320201233130-1100212010220012-2220231112033113-3100133020203133-3032232101303231-1021031112133031): complete subsection reference.

<a id="canonical-1131000302300201-1330312202211222-2320023031230032-3113330111131123-3033013120011233-0101101312010000-1332302023100210-3232111201303012"></a>

<a id="canonical-2332331211033110-1333302210002220-0201012312322133-0021212102132211-1203101300302312-3230100103301333-2221212132102322-0210102130031132"></a>

## connection_idle_timeout property — Property reference / 002321300333 / 7

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

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

- [cors_policy](data-sources--virtual_host--reference--group-002.md#canonical-2210323031131321-0313202301131023-3213332203010030-0103323031221230-3302100002020020-3202222112020312-2232021222011023-1201032103301123): complete subsection reference.

- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-1032321323223020-2022231132220113-3011022121331033-0131132302111121-0112122202013132-2121330303333120-2133303323113131-3121123003303012): complete subsection reference.

<a id="canonical-0311333110202211-3220301301033201-1301030023233331-3123331111231031-0112332302032310-1000122332013012-1333321330121302-2022332202001221"></a>

<a id="canonical-1320111202100200-0003200222131002-3331013020223022-3330331110010233-2300312202213323-0102112122000031-0301011121203110-3100020232220200"></a>

## custom_errors property — Property reference / 002321300333 / 8

Type: `["map", "string"]`. Computed.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx..

Upstream description:

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

- [default_header](data-sources--virtual_host--reference--group-002.md#canonical-0223221131230033-2031313102132311-2030022100211130-0021201230212110-3110310230000301-0000220012010000-1002213210103013-0320220200333210): complete subsection reference.

- [default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-2332201222012331-2020200003133331-3003021002300031-1121110112222301-0202020120202032-2121220121022313-1310230211023310-1000212330230231): complete subsection reference.

<a id="canonical-1012211020332002-2000223221113203-0231031301220101-1302320230213121-2230310120301232-0001000002221030-2323033233231311-0211012023030023"></a>

<a id="canonical-1010033212103102-3333022302021222-0123310312313200-0322301110310012-3011001012131022-3332313221112120-3113000130112020-1230221001232010"></a>

## description property — Property reference / 002321300333 / 9

Type: `"string"`. Computed.

Description of the VirtualHost.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3301123013022232-0323303003312030-2230113313201003-0112003031301202-0321130302211230-2301230021222100-2201032031012132-3003103101133302"></a>

<a id="canonical-0303021230010220-1313330122223313-3202303101031110-1113232233233221-2100102123133330-0320223222310032-1101332133233121-1312033302223201"></a>

## disable_default_error_pages property — Property reference / 002321300333 / 10

Type: `"bool"`. Computed.

Option to specify whether to disable using default F5XC error pages.

Upstream description:

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

<a id="canonical-2121220132020120-1333002003213320-0200320333211300-2311230133202100-2230201100031121-1101202121211010-1032202133302213-3101101320223213"></a>

<a id="canonical-0032333001332201-3033022201133022-3200233302031233-0021033222303103-1321200310320002-2112330321322320-0330132102132331-1302322323032233"></a>

## disable_dns_resolve property — Property reference / 002321300333 / 11

Type: `"bool"`. Computed.

Disable DNS resolution for domains specified in the virtual host When the virtual host is configured
as Dynamive Resolve Proxy (DRP), disable DNS resolution for domains configured. This configuration
is suitable for HTTP CONNECT proxy.

Upstream description:

Disable DNS resolution for domains specified in the virtual host

When the virtual host is configured as Dynamive Resolve Proxy (DRP), disable DNS resolution for
domains configured. This configuration is suitable for HTTP CONNECT proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-3323313313333101-1022230003023322-1121321233003231-3321002002120031-3231313020123102-0332221122301230-2301200331132022-2132211330222112): complete subsection reference.

<a id="canonical-2001102223333120-3121303313103013-1233202010221303-0130333220131020-2021201303210221-3023220023120302-1231120123312032-3212112103120301"></a>

<a id="canonical-2130112131201323-0323121302301012-2210000230201230-1023111120000221-0230201102310112-0231020310113223-2201123022131123-3030030221113032"></a>

## domains property — Property reference / 002321300333 / 12

Type: `["list", "string"]`. Computed.

List of domain names matched to this virtual host for routing incoming requests. Supports wildcard
patterns like \*.example.com for subdomain matching.

Upstream description:

A list of Domains (host/authority header) that will be matched to this Virtual Host. Wildcard hosts
are supported in the suffix or prefix form

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if the virtual host proxy type is
TCP\_PROXY\_WITH\_SNI/HTTPS\_PROXY Domains also indicate the list of names for which DNS resolution
will be automatically resolved to IP addresses by the system.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [dynamic_reverse_proxy](data-sources--virtual_host--reference--group-002.md#canonical-2320322030133330-2313030201332123-3133200313313203-0300032102221333-1001120222131112-1330331333122030-2312013001023021-0010002323330332): complete subsection reference.

- [enable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-0101130103132132-0220121022213231-2333332112030130-0121011300022021-3031032103223222-2102320203302331-2033233303102321-3213203032133202): complete subsection reference.

- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022): complete subsection reference.

<a id="canonical-0123200020210113-3331101011100011-2131023231221200-1302132213002302-0321212210132202-2111232220131010-1213010301331123-3313220022231303"></a>

<a id="canonical-0020013333130123-0020301201333233-2301111202131112-2212100312210232-2001220213213330-2321333033133303-2001332002031002-1210331020131030"></a>

## ID property — Property reference / 002321300333 / 13

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2032022311020320-2103200213220300-2103320332232231-2232202010112301-2022222223222103-3121300333203131-2033003232203122-0000321330323223"></a>

<a id="canonical-0103312123023333-2002311331311131-0223231033311300-1123320031310020-1001013022102000-0110131013112300-0201113020023012-1020303130330211"></a>

## idle_timeout property — Property reference / 002321300333 / 14

Type: `"number"`. Computed.

Idle timeout is the amount of time that the loadbalancer will allow a stream to exist with no
upstream or downstream activity. Idle timeout and Proxy Type: HTTP\_PROXY, HTTPS\_PROXY: Idle timer
is started when the first byte is received on the connection. Each time an encode/decode event for..

Upstream description:

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

- [js_challenge](data-sources--virtual_host--reference--group-002.md#canonical-0311311032210221-1213332233330222-2103321233000202-2020120320331211-3131003020011333-1211120300112103-3021112313203310-2031301312202110): complete subsection reference.

<a id="canonical-2321100313100033-1030011101330323-2210010211320131-3111302003012300-2223223121121020-3202023131211212-0321222301012010-3101201022333011"></a>

<a id="canonical-1023232130230130-0212121320023313-0321220101301310-3113033001112031-0003220231130300-3213103111223031-3332011301033020-3020033012100101"></a>

## labels property — Property reference / 002321300333 / 15

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

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

<a id="canonical-1112301012203222-0033322101102100-2102113011010203-2202321311313220-3123122012020121-1302131003123102-3213032133120033-2103202330022033"></a>

<a id="canonical-0010001111220122-2222013332310213-2332210013333012-2332010120030132-0003002213022300-1233013300002311-2231021300101001-1032322230031000"></a>

## max_request_header_size property — Property reference / 002321300333 / 16

Type: `"number"`. Computed.

The maximum request header size in KiB for incoming connections. If un-configured, the default max
request headers allowed is 60 KiB. Requests that exceed this limit will receive a 431 response.

Upstream description:

The maximum request header size in KiB for incoming connections.

If un-configured, the default max request headers allowed is 60 KiB.

Requests that exceed this limit will receive a 431 response.

The max configurable limit is 96 KiB, based on current implementation constraints.

Note: a. This configuration parameter is applicable only for HTTP\_PROXY and HTTPS\_PROXY b. When
multiple HTTP\_PROXY virtual hosts share the same advertise policy, the effective "maximum request
header size" for such virtual hosts is the highest value configured on any of the virtual hosts.

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
    "ves.io.schema.rules.uint32.lte": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  }
}
```

<a id="canonical-2120031230002123-2101102332011300-0002133320120210-1203232121221321-0312032203130200-3003330211201232-0112030211002123-3313310312112033"></a>

<a id="canonical-0313112332212010-0230331311120231-3101331133222031-0110310032231112-2110310000211201-1102210323131023-3331132110020113-3231123213133103"></a>

## max_requests_per_connection property — Property reference / 002321300333 / 17

Type: `"number"`. Computed.

\[OneOf: max\_requests\_per\_connection, no\_request\_limit\_per\_connection; Default:
no\_request\_limit\_per\_connection\] Exclusive with \[no\_request\_limit\_per\_connection\] Sets
the maximum number of requests a downstream client can send over a single connection to Envoy. Enter
a value &gt;=1 to define the request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

OneOf alternatives in this subsection:

- [max_requests_per_connection](data-sources--virtual_host--reference--group-001.md#canonical-2120031230002123-2101102332011300-0002133320120210-1203232121221321-0312032203130200-3003330211201232-0112030211002123-3313310312112033)
- [no_request_limit_per_connection](data-sources--virtual_host--reference--group-002.md#canonical-2222231010030030-3130021202122031-3100103020130033-3121222000100032-1102001222302230-2111131103111212-2133112303230032-3021011203213203)

Select alternatives according to the provider validators above.

<a id="canonical-0133010022031202-0132230130110103-2120123002212031-3123123320112333-3320310312003111-0103000111012002-1233322113303120-3332322030320101"></a>

<a id="canonical-3121113101101202-1021103223032022-3221122120322332-3200111011003301-0232101003323122-1100002333021333-1101103122323232-1120002303202002"></a>

## name property — Property reference / 002321300333 / 18

Type: `"string"`. Required.

Name of the VirtualHost.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1103220003212213-2311300012331233-3223311312302210-3030301132231022-3102213320310013-3210121011303201-0032023010313300-0333211213131213"></a>

<a id="canonical-0033332022330133-3131033231322231-2110033003313231-1321120210232011-2223300202131001-1123112002232311-0301022330103021-1011313332210223"></a>

## namespace property — Property reference / 002321300333 / 19

Type: `"string"`. Required.

Namespace where the VirtualHost exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [no_authentication](data-sources--virtual_host--reference--group-002.md#canonical-1233312122313321-2011202230001010-3102133312211200-1130003122002233-2223112232011102-0131310120111220-3120020032023010-0313002202102123): complete subsection reference.

- [no_challenge](data-sources--virtual_host--reference--group-002.md#canonical-3110101102113030-0110100111222000-3312320302213100-0201010223313300-1213232220022000-2133222310110000-2331222120121232-2212130222311233): complete subsection reference.

- [no_request_limit_per_connection](data-sources--virtual_host--reference--group-002.md#canonical-1110112230220221-2303010133322201-1101300033010303-0103123220222322-3010312001321310-0123120221200321-1022302003222231-0020320101133303): complete subsection reference.

- [non_default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-2023112320112123-3300121231232030-1100122000331201-1003213223120203-1222330212012300-3302102223112001-3102332220100221-1121233221011031): complete subsection reference.

- [pass_through](data-sources--virtual_host--reference--group-002.md#canonical-1202312331000133-2122223121103310-2100301313303013-2310311201301110-2210333202010230-0211311013211222-0101120011203120-3303021112101311): complete subsection reference.

<a id="canonical-2033111332122010-1200133213031022-1130303333101001-2032010113303131-0203113033130133-2222303300032300-0333000030331023-2020302012321331"></a>

<a id="canonical-3201333300310233-1112323300103012-2302200112202322-0300031123202033-3123103332200313-3023110230003131-0202202000020321-1010323222020120"></a>

## proxy property — Property reference / 002321300333 / 20

Type: `"string"`. Computed.

\[Enum:
UDP\_PROXY|SMA\_PROXY|DNS\_PROXY|ZTNA\_PROXY|UZTNA\_PROXY|TMM\_HTTP\_PROXY|TMM\_HTTPS\_PROXY|TMM\_TCP\_PROXY|TMM\_UDP\_PROXY|TMM\_QUIC\_PROXY\]
ProxyType tells the type of proxy to install for the virtual host. Only the following combination of
VirtualHosts within same AdvertisePolicy is permitted (None of them should have '\*' in domains when
used with other VirtualHosts in same AdvertisePolicy) 1. Multiple TCP\_PROXY\_WITH\_SNI and..
Possible values are \`UDP\_PROXY\`, \`SMA\_PROXY\`, \`DNS\_PROXY\`, \`ZTNA\_PROXY\`,
\`UZTNA\_PROXY\`, \`TMM\_HTTP\_PROXY\`, \`TMM\_HTTPS\_PROXY\`, \`TMM\_TCP\_PROXY\`,
\`TMM\_UDP\_PROXY\`, \`TMM\_QUIC\_PROXY\`.

Upstream description:

ProxyType tells the type of proxy to install for the virtual host.

Only the following combination of VirtualHosts within same AdvertisePolicy is permitted (None of
them should have "\*" in domains when used with other VirtualHosts in same AdvertisePolicy)
&#8203;1. Multiple TCP\_PROXY\_WITH\_SNI and multiple HTTPS\_PROXY &#8203;2. Multiple HTTP\_PROXY
&#8203;3. Multiple HTTPS\_PROXY &#8203;4. Multiple TCP\_PROXY\_WITH\_SNI

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

- [rate_limiter_allowed_prefixes](data-sources--virtual_host--reference--group-002.md#canonical-1133310002220332-1103221000011023-3032223311312122-2321113131121222-2110333100201023-1011110232221212-1300111103011301-1303120213023022): complete subsection reference.

- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3012001313211032-0121110012120332-1211212103012331-1030201012110112-3012311230231023-0312310313121023-3122223302301012-1011130302002022): complete subsection reference.

<a id="canonical-3121332030132111-1033032022010131-0231011222333033-0011333330301000-0011130210122033-0121203201322112-0010100331231132-3331101133320302"></a>

<a id="canonical-3100113201320130-3110032030120021-2303100300302110-3332022310121013-2102021233023331-2013200230010203-3220122033323322-0330121300002220"></a>

## request_cookies_to_remove property — Property reference / 002321300333 / 21

Type: `["list", "string"]`. Computed.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

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

- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3303213013223123-0333021210322333-3000030200231312-1002303202333220-2113133012232131-0030321011113311-3012010300230300-3131003321220113): complete subsection reference.

<a id="canonical-0132230023120312-0323110133220311-2200100113101302-3302003030102231-1200223330013012-3012320331030001-1002033032102022-2121031120023110"></a>

<a id="canonical-0320000120223110-3303230101130313-1001120222122022-1113223301021202-0000230111112110-0201203211203000-1212321321110333-1232033121310011"></a>

## request_headers_to_remove property — Property reference / 002321300333 / 22

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022): complete subsection reference.

<a id="canonical-2012011312030331-0033023101223210-1313013112102210-1321032113231022-2003120322020032-1332310021111120-0130022121211301-3322320320103000"></a>

<a id="canonical-0123033220003013-0303102031221222-3310101000312301-2112330320301200-2031331132221132-3001303330011333-3012322103021301-0202312131032111"></a>

## response_cookies_to_remove property — Property reference / 002321300333 / 23

Type: `["list", "string"]`. Computed.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

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

- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-1122301030300332-0033230331023011-1102233120003323-3100320132003322-0030000111220313-3302203220201033-0222213312322000-2123201002112113): complete subsection reference.

<a id="canonical-0033332310122132-2102211023030310-3331313201111222-0320231303022031-0232220110330333-2233111210213332-1311230032111033-1030311110013200"></a>

<a id="canonical-3121320100030123-2233022023202021-2033223013022133-1021013332131200-1110201330022011-1030100303231302-3111330103103023-3030222201022021"></a>

## response_headers_to_remove property — Property reference / 002321300333 / 24

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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

- [retry_policy](data-sources--virtual_host--reference--group-003.md#canonical-1021031030330102-2100323023122311-3003021113211121-2312223333303303-2002222022211201-0131100013001013-2031022313222112-2203022023311021): complete subsection reference.

- [routes](data-sources--virtual_host--reference--group-003.md#canonical-3211011200123230-0330211200101113-1112022302023320-3120120232213300-3223122013100230-0003102201231003-1303322122132321-0212222131122110): complete subsection reference.

- [sensitive_data_policy](data-sources--virtual_host--reference--group-003.md#canonical-1111130000112110-1102321322100210-0013310223010100-3213130303222101-3131013201101301-1000021000133131-0023212331022302-0120103003321211): complete subsection reference.

<a id="canonical-2000133212001013-1130012132023121-0030132122333222-0303323311000221-0303100033122013-2021210110033131-2331333211132010-0220310022330223"></a>

<a id="canonical-2312300132331310-3313131310003003-1023123302012111-3221220012123210-3201010223233301-1203131101013231-1312001321002210-3222223111310132"></a>

## server_name property — Property reference / 002321300333 / 25

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Specifies the value to be used
for Server header inserted in responses. This will overwrite existing values if any for Server
Header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Specifies the value to be used
for Server header inserted in responses. This will overwrite existing values if any for Server
Header.

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

- [slow_ddos_mitigation](data-sources--virtual_host--reference--group-003.md#canonical-2020302310011113-0110101030032202-1233233331203200-0321332220211010-1220330200213312-3321332320232003-2113021203322331-0330021133123211): complete subsection reference.

- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311): complete subsection reference.

- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001): complete subsection reference.

- [user_identification](data-sources--virtual_host--reference--group-003.md#canonical-1002302233122000-1323022011111121-3132010100020022-1310010213200321-0312233032021223-1020313001300323-0020033031030302-1102203102302301): complete subsection reference.

- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001): complete subsection reference.

<a id="canonical-1033232213223220-3321111330320100-3223010230101030-1000300031332322-3101020223012103-1111021220011101-2132322323100321-2222110131000321"></a>

## All schema paths — Property reference / 002321300333 / 26

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `add_location` | [add_location](data-sources--virtual_host--reference--group-001.md#canonical-1030131231003202-0332223001323000-3303213301132120-1322203002303212-3022000033211203-2201331120311013-0212231332111200-0003221011321112) |
| `advertise_policies` | [advertise_policies](data-sources--virtual_host--reference--group-001.md#canonical-0011102110223101-0212323101310331-2023220110030200-1001323010012021-3113203200202221-0312301121213031-3231020023122123-2203011222220211) |
| `advertise_policies.kind` | [advertise_policies.kind](data-sources--virtual_host--reference--group-001.md#canonical-3013230210303330-2021000312330303-2111023102333112-0101003022211221-0230000133232203-2313100222122212-1223223112231333-3221131121210220) |
| `advertise_policies.name` | [advertise_policies.name](data-sources--virtual_host--reference--group-001.md#canonical-3002303023302122-2022213221232200-3011013102312130-0310103331100030-2001203123220121-3323323130200331-1130133211313122-1110000232002132) |
| `advertise_policies.namespace` | [advertise_policies.namespace](data-sources--virtual_host--reference--group-001.md#canonical-3121322320123212-0111001202122000-3032010001211200-0330111311101313-1220011021311202-3023212003303133-3130003023100203-0010002233212033) |
| `advertise_policies.tenant` | [advertise_policies.tenant](data-sources--virtual_host--reference--group-001.md#canonical-3313020203303211-1003300231330230-0111100022203200-0020230132122230-3132120222230331-0330102130120332-3333302332100233-3333233220000011) |
| `advertise_policies.uid` | [advertise_policies.uid](data-sources--virtual_host--reference--group-001.md#canonical-3021122302001010-3012210120022230-1131231113211310-0022300311330332-3100313013231322-0012311233213112-3012202222331030-1120332233220223) |
| `annotations` | [annotations](data-sources--virtual_host--reference--group-001.md#canonical-2210320232301302-1231203000111220-2222222000321132-2101023221203301-3102210110333102-2232300312100112-2032133301000131-3021312023200233) |
| `append_server_name` | [append_server_name](data-sources--virtual_host--reference--group-001.md#canonical-0323113201300223-2222333123230213-2303032122303012-1102010201023100-1233332111220321-3010303110000300-2113333030020321-1122222213033203) |
| `authentication` | [authentication](data-sources--virtual_host--reference--group-001.md#canonical-1110113100202123-2230111331000202-3131210333023303-2221020221222112-0131013231123111-0123202113213213-0313312033121100-1332022222100032) |
| `authentication.auth_config` | [authentication.auth_config](data-sources--virtual_host--reference--group-001.md#canonical-1003023202213212-1301232202222102-1230002021210033-0113311211002030-0102100110110032-1232122301102200-1220020000203202-0231131131033021) |
| `authentication.auth_config.kind` | [authentication.auth_config.kind](data-sources--virtual_host--reference--group-001.md#canonical-0223022003300122-0131031023021230-1112002002133221-1213233221232131-1310021331132312-3023230122021110-2032213331113122-3301322013321320) |
| `authentication.auth_config.name` | [authentication.auth_config.name](data-sources--virtual_host--reference--group-001.md#canonical-2103020112302120-0110211313000110-1310302023230013-0230023213310231-2010002321102030-2001310110232111-0002110311303231-1230320120230131) |
| `authentication.auth_config.namespace` | [authentication.auth_config.namespace](data-sources--virtual_host--reference--group-001.md#canonical-3223331333203110-3321120012301122-2112100001213321-2303012111202310-2321102233211222-3210210131201111-1313111121020313-2031232020033231) |
| `authentication.auth_config.tenant` | [authentication.auth_config.tenant](data-sources--virtual_host--reference--group-001.md#canonical-3313200301300210-0023011110220222-2113213013321023-1002123200213203-0223212332231201-3102130020123100-0101201111232222-3111310010203133) |
| `authentication.auth_config.uid` | [authentication.auth_config.uid](data-sources--virtual_host--reference--group-001.md#canonical-1303100022102322-0021022022321111-0133012000201220-2131022021030013-1201010011112133-2333010323111210-2211223331330020-1323011001023113) |
| `authentication.cookie_params` | [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-0333203001003223-3030333220013012-0323210203130022-1220023201223030-0111220111100013-1231200223322021-3020002110121123-1212331033231233) |
| `authentication.cookie_params.auth_hmac` | [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-2232331311211131-1203133120300113-2133103133131233-2023313222112101-0123130120210311-1200011130311133-3232330222120303-2303003230133233) |
| `authentication.cookie_params.auth_hmac.prim_key` | [authentication.cookie_params.auth_hmac.prim_key](data-sources--virtual_host--reference--group-001.md#canonical-1323313131100023-3010310200121023-3233030312213031-1302212130220110-1112200130331212-3211003232222120-1133022013003323-0012120133103210) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-3000031001311211-3010133210120033-1121122220002101-0213003120100000-3010132232132123-3303101011110333-3033030313201102-3033301100001330) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider](data-sources--virtual_host--reference--group-001.md#canonical-1030113313133101-2211312331032313-3021133120333322-1313020001301301-3233313011322111-0323113300033313-2033333332220300-1211110003233100) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.location` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.location](data-sources--virtual_host--reference--group-001.md#canonical-0001202000002000-0320221121001102-3110211232210212-1120113100301230-3212101201031012-3230323022132011-3010121003311032-1310302301100301) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider](data-sources--virtual_host--reference--group-001.md#canonical-3022210310033223-2321013123230313-2213201100212210-2313212320030213-0030322201121020-2332102012111231-1131321123011233-1230013011223120) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-0021031111120101-3002023003230033-2122320032211313-0223032121101202-0232203123211203-3112123213222123-1103002322112002-0333031021133211) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref](data-sources--virtual_host--reference--group-001.md#canonical-0311101021101013-0003231132021013-1231131311230210-0301321030021002-3000332021022333-2111313112023023-0131332033301132-3223223203002123) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.url` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.url](data-sources--virtual_host--reference--group-001.md#canonical-3010020312320220-1221323300010322-3323202203113311-3233112022003101-2303130122223200-1311323101122220-0223330222111212-1003130211022020) |
| `authentication.cookie_params.auth_hmac.prim_key_expiry` | [authentication.cookie_params.auth_hmac.prim_key_expiry](data-sources--virtual_host--reference--group-001.md#canonical-3301012231300102-1221303020220301-2112330313331210-3313032222311230-1231033211202313-3221000301201330-2211213132320232-3313032202301110) |
| `authentication.cookie_params.auth_hmac.sec_key` | [authentication.cookie_params.auth_hmac.sec_key](data-sources--virtual_host--reference--group-001.md#canonical-2221210233223010-2031120323303311-0020031110022022-3320303003213211-0121222302322221-1232022212320100-3011333231212031-0001110021023313) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-0030222023020032-3120110332102222-2211110100210210-0213333112313103-3303222230322222-1102000131303333-1323321001103001-1123311320113103) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider](data-sources--virtual_host--reference--group-001.md#canonical-0201223103132200-1113031222322302-1312212022213211-2103133330303300-0313320032212332-3131232231333230-3122202311303032-0232111111132213) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.location` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.location](data-sources--virtual_host--reference--group-001.md#canonical-3013210100030223-1202130021010013-3211321100020222-0111033111320220-0232101033123100-2302033210013221-1011233020000331-1120300331222301) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider](data-sources--virtual_host--reference--group-001.md#canonical-2302231320122132-1103202002301213-2333122010320110-0023221222023210-3003332202131210-3030130310213303-2321222111231230-1110311300222122) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-1331201132303122-1220312323210231-0310330122030233-3010201020122300-1202323102100100-0012013131022033-0010221113233123-2323102302131002) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref](data-sources--virtual_host--reference--group-001.md#canonical-3310302021212331-3031031303323330-3002212011131123-2301021013322313-1200122301032101-1113312220221022-0302010113203113-1030023030131321) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.url` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.url](data-sources--virtual_host--reference--group-001.md#canonical-0320233330233232-1333300323122202-3210032212121233-1010331031021100-2232302103100323-1130100023023221-2231221001031323-3311001101113301) |
| `authentication.cookie_params.auth_hmac.sec_key_expiry` | [authentication.cookie_params.auth_hmac.sec_key_expiry](data-sources--virtual_host--reference--group-001.md#canonical-3111130303030311-0000331211223202-2022101222022032-2310120103001220-3330211013103111-1211100233110020-1332102031231230-1301301213311223) |
| `authentication.cookie_params.cookie_expiry` | [authentication.cookie_params.cookie_expiry](data-sources--virtual_host--reference--group-001.md#canonical-0003102311330132-0313132122120030-1323302011113011-1013001013212110-0133223320313310-3130220332033010-2030302330100302-2003201320331323) |
| `authentication.cookie_params.cookie_refresh_interval` | [authentication.cookie_params.cookie_refresh_interval](data-sources--virtual_host--reference--group-001.md#canonical-0012120012110132-3231303313212201-0212223002222301-1022221230122030-0001001312122302-0202332210213232-0120023011200032-3310230101021323) |
| `authentication.cookie_params.kms_key_hmac` | [authentication.cookie_params.kms_key_hmac](data-sources--virtual_host--reference--group-001.md#canonical-0023330230021020-0220320301213000-2133120211033030-1330302000221101-0000200020123303-0020112121033310-0312113022323110-3231331313033201) |
| `authentication.cookie_params.session_expiry` | [authentication.cookie_params.session_expiry](data-sources--virtual_host--reference--group-001.md#canonical-3323201110032133-3233133013200122-2101112001021030-2020012121100001-0012333323130311-0123100121012130-3200213313112123-0312121011131323) |
| `authentication.redirect_dynamic` | [authentication.redirect_dynamic](data-sources--virtual_host--reference--group-001.md#canonical-0203111121110120-0310232303132311-3012213303010111-1112012301203313-1330100301302313-1032130022011020-1001000120001220-1023030311203201) |
| `authentication.redirect_url` | [authentication.redirect_url](data-sources--virtual_host--reference--group-001.md#canonical-0021332021332103-1221101003101101-3130101310320333-0011113322032121-2303312333123202-1023131110330322-3232332301021323-1131033030122313) |
| `authentication.use_auth_object_config` | [authentication.use_auth_object_config](data-sources--virtual_host--reference--group-001.md#canonical-3210031121102023-1030002222002333-1220122030001210-1221212223232231-0102131303200000-3033112230303022-2313220122111000-0323301323103211) |
| `buffer_policy` | [buffer_policy](data-sources--virtual_host--reference--group-001.md#canonical-1113013312331001-3120013012321211-3201312010013030-3223213111323031-3010323000321112-0003300310113023-2231303133033022-1212033301322100) |
| `buffer_policy.disabled` | [buffer_policy.disabled](data-sources--virtual_host--reference--group-001.md#canonical-2230303132301303-2312003132031010-2033023210021112-0310223003331100-2233202323312113-2303300121200203-1122003302131031-1333200122002112) |
| `buffer_policy.max_request_bytes` | [buffer_policy.max_request_bytes](data-sources--virtual_host--reference--group-001.md#canonical-2232003130112123-0003030132220332-0213212130032112-2222133321201230-2203211023033210-0213023033300032-2032301333021201-0310321032100313) |
| `captcha_challenge` | [captcha_challenge](data-sources--virtual_host--reference--group-002.md#canonical-0323333321111013-0223212113102021-0320213033023313-3002301101200322-2122100233212230-0130313033323212-2101021122013300-0032223213002013) |
| `captcha_challenge.cookie_expiry` | [captcha_challenge.cookie_expiry](data-sources--virtual_host--reference--group-002.md#canonical-2313032031202002-2323131312322312-2012100013033132-0213022300212313-1011321123333012-0020300030213123-3101311313011100-2031113312310121) |
| `captcha_challenge.custom_page` | [captcha_challenge.custom_page](data-sources--virtual_host--reference--group-002.md#canonical-1023303000202202-2310101312302113-1011223101331311-3010012200021300-3013231231232321-0113313113231313-0012331330211003-2203131202210321) |
| `coalescing_options` | [coalescing_options](data-sources--virtual_host--reference--group-002.md#canonical-3003023221300230-2101220002010101-2301202002003033-0013123000032321-0331000322101331-2323033312000302-0001110123233331-2133220013333201) |
| `coalescing_options.default_coalescing` | [coalescing_options.default_coalescing](data-sources--virtual_host--reference--group-002.md#canonical-0332011001302102-2220230311100233-1122210310102113-0012003211300111-1232222231321010-3302112321222210-0023112101131001-0203210131211001) |
| `coalescing_options.strict_coalescing` | [coalescing_options.strict_coalescing](data-sources--virtual_host--reference--group-002.md#canonical-3321032220211300-0111031200213322-2313223213103101-1032101010301202-3002220201113200-2023000233010323-0302101333021132-0221330300233232) |
| `compression_params` | [compression_params](data-sources--virtual_host--reference--group-002.md#canonical-0313011002003132-2220330320132201-0102332000101123-0212201213122331-2011031320011022-0110321130112110-3222230213210003-0110003130320121) |
| `compression_params.content_length` | [compression_params.content_length](data-sources--virtual_host--reference--group-002.md#canonical-1202210022321230-1213002133010133-0121100320230320-3331121231100010-0312230133230221-0313302333001030-0331332333022221-3222031303020203) |
| `compression_params.content_type` | [compression_params.content_type](data-sources--virtual_host--reference--group-002.md#canonical-3121131330031112-1322032200130030-0212211311133211-0103231132331330-2100221331131013-0102212223002310-1323231202330123-0200202203320101) |
| `compression_params.disable_on_etag_header` | [compression_params.disable_on_etag_header](data-sources--virtual_host--reference--group-002.md#canonical-2031200203032112-3133012330210022-3100123300021302-3011131022223210-3203020021022123-3223220210213213-1112231132011221-1213331201332110) |
| `compression_params.remove_accept_encoding_header` | [compression_params.remove_accept_encoding_header](data-sources--virtual_host--reference--group-002.md#canonical-1320100232322033-1231232022213011-3103030313201103-2303111100022201-2333213322112210-0001233223333201-0302312320202322-3113200001121313) |
| `connection_idle_timeout` | [connection_idle_timeout](data-sources--virtual_host--reference--group-001.md#canonical-1131000302300201-1330312202211222-2320023031230032-3113330111131123-3033013120011233-0101101312010000-1332302023100210-3232111201303012) |
| `cors_policy` | [cors_policy](data-sources--virtual_host--reference--group-002.md#canonical-2000022133131221-1002001101221112-1210213110223313-3033230021132322-2202310223211101-1203010003300113-1323332122301323-3130222330120000) |
| `cors_policy.allow_credentials` | [cors_policy.allow_credentials](data-sources--virtual_host--reference--group-002.md#canonical-2010200221011321-0203030121013302-3023130313322200-2232231133003020-1332300123322122-1133123321010323-2103000101330133-1303113313322032) |
| `cors_policy.allow_headers` | [cors_policy.allow_headers](data-sources--virtual_host--reference--group-002.md#canonical-3333110322201200-3200111132322112-0213032322321223-2310300310303011-3023203102213301-1001320310010221-3022002232220332-3030220000313110) |
| `cors_policy.allow_methods` | [cors_policy.allow_methods](data-sources--virtual_host--reference--group-002.md#canonical-3011003321332333-0003303003333202-0033110031131202-3121033102213331-0121221303133310-2030231133132302-2032223103323030-1011110230100323) |
| `cors_policy.allow_origin` | [cors_policy.allow_origin](data-sources--virtual_host--reference--group-002.md#canonical-1023231122322210-2011322313323333-3133330000030100-0303221213033123-3000021202232021-2122103210321021-2320000101230322-0011010101200310) |
| `cors_policy.allow_origin_regex` | [cors_policy.allow_origin_regex](data-sources--virtual_host--reference--group-002.md#canonical-1030231302230221-1232012120110010-2120211300100330-3211122203310210-3022232320020232-1213032111031330-3113300230022222-0011230012111032) |
| `cors_policy.disabled` | [cors_policy.disabled](data-sources--virtual_host--reference--group-002.md#canonical-3210132113233112-2311110231332203-0103231031113223-2002311233031201-1033100200303322-0110202200232223-0221013132200002-1222202221100111) |
| `cors_policy.expose_headers` | [cors_policy.expose_headers](data-sources--virtual_host--reference--group-002.md#canonical-2330012220221131-3312310133002203-3230013303010331-0013301102323103-1232021133021311-1031111203332201-3210010231330322-1102211311321110) |
| `cors_policy.maximum_age` | [cors_policy.maximum_age](data-sources--virtual_host--reference--group-002.md#canonical-3322022203032032-1330310120220232-0131333002211321-3311032213120212-3102323303202321-1322113211133310-2010201123013202-3313203021212120) |
| `csrf_policy` | [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-3031223133200111-0120310302032202-1030211123102211-1000333030031220-1303103321101333-1132110213133312-2111201110102013-2320231323100000) |
| `csrf_policy.all_load_balancer_domains` | [csrf_policy.all_load_balancer_domains](data-sources--virtual_host--reference--group-002.md#canonical-1122301000310102-2230013332210023-3030030331210002-3012113120120033-0233302122233333-3103011203121201-0123033322112323-2000331113021021) |
| `csrf_policy.custom_domain_list` | [csrf_policy.custom_domain_list](data-sources--virtual_host--reference--group-002.md#canonical-1313202213002001-3233011212130111-1030101220013023-1000133331130223-0002131231033230-0302210322330312-3313232230220110-2110323021230121) |
| `csrf_policy.custom_domain_list.domains` | [csrf_policy.custom_domain_list.domains](data-sources--virtual_host--reference--group-002.md#canonical-0113300021331300-3220133103110033-0202002232331212-3030220321203022-0113221222011331-1321133230001221-3111111200302030-0312301120222132) |
| `csrf_policy.disabled` | [csrf_policy.disabled](data-sources--virtual_host--reference--group-002.md#canonical-0323033223122213-0133300131232223-2202110010030003-2332331012323133-0333133123231311-1223122031332132-3022301020231310-0320133003013332) |
| `custom_errors` | [custom_errors](data-sources--virtual_host--reference--group-001.md#canonical-0311333110202211-3220301301033201-1301030023233331-3123331111231031-0112332302032310-1000122332013012-1333321330121302-2022332202001221) |
| `default_header` | [default_header](data-sources--virtual_host--reference--group-002.md#canonical-0131231003301103-0321312203033301-3002222033101103-1222201231020220-1232210300213310-0131212300130011-2002101030031123-0001030211230102) |
| `default_loadbalancer` | [default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-0203210323302221-1200031011123211-0101202233211001-3323232022101020-0000321121303303-2333220023222121-3300321333323100-1032021333100320) |
| `description` | [description](data-sources--virtual_host--reference--group-001.md#canonical-1012211020332002-2000223221113203-0231031301220101-1302320230213121-2230310120301232-0001000002221030-2323033233231311-0211012023030023) |
| `disable_default_error_pages` | [disable_default_error_pages](data-sources--virtual_host--reference--group-001.md#canonical-3301123013022232-0323303003312030-2230113313201003-0112003031301202-0321130302211230-2301230021222100-2201032031012132-3003103101133302) |
| `disable_dns_resolve` | [disable_dns_resolve](data-sources--virtual_host--reference--group-001.md#canonical-2121220132020120-1333002003213320-0200320333211300-2311230133202100-2230201100031121-1101202121211010-1032202133302213-3101101320223213) |
| `disable_path_normalize` | [disable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-3301102111312121-0031222233000230-3002223322002213-3111133201103210-3210213120101312-2112223132033332-1123130311310222-0210010110330230) |
| `domains` | [domains](data-sources--virtual_host--reference--group-001.md#canonical-2001102223333120-3121303313103013-1233202010221303-0130333220131020-2021201303210221-3023220023120302-1231120123312032-3212112103120301) |
| `dynamic_reverse_proxy` | [dynamic_reverse_proxy](data-sources--virtual_host--reference--group-002.md#canonical-3312001003220201-1021111332030323-2320100230313010-0231102023313230-3231131000121133-2230033123002221-1131102320231311-0211223000310232) |
| `dynamic_reverse_proxy.connection_timeout` | [dynamic_reverse_proxy.connection_timeout](data-sources--virtual_host--reference--group-002.md#canonical-3211220132132303-3112133120120133-3202311011213200-1002000110020003-2320230021021132-1310002121231023-1132132011202020-0331230332012130) |
| `dynamic_reverse_proxy.resolution_network` | [dynamic_reverse_proxy.resolution_network](data-sources--virtual_host--reference--group-002.md#canonical-3033013100002213-2212012332232022-0101032322103101-3203023313232302-2003311011303113-2011101113132023-1103120111033031-0211210301122103) |
| `dynamic_reverse_proxy.resolution_network.kind` | [dynamic_reverse_proxy.resolution_network.kind](data-sources--virtual_host--reference--group-002.md#canonical-3300130110213322-2033023121303131-0322022120321010-2031100303233130-0033131101331112-3100331302112101-0123032100001131-1103221033132010) |
| `dynamic_reverse_proxy.resolution_network.name` | [dynamic_reverse_proxy.resolution_network.name](data-sources--virtual_host--reference--group-002.md#canonical-0332300122302032-2201103333100232-3201320201321221-3033020123112300-0102213011320201-2020331130323130-2331001000332332-3311230233032221) |
| `dynamic_reverse_proxy.resolution_network.namespace` | [dynamic_reverse_proxy.resolution_network.namespace](data-sources--virtual_host--reference--group-002.md#canonical-3030031013223202-3132003021232202-3032020020210201-1213002033221110-2222101222221300-1122002222231100-2030001102222002-3102113221212030) |
| `dynamic_reverse_proxy.resolution_network.tenant` | [dynamic_reverse_proxy.resolution_network.tenant](data-sources--virtual_host--reference--group-002.md#canonical-2221121323012333-2030311132131000-0121232223200221-2111301111231320-1202323202320031-3011112301121100-1321301331102200-2303313311123303) |
| `dynamic_reverse_proxy.resolution_network.uid` | [dynamic_reverse_proxy.resolution_network.uid](data-sources--virtual_host--reference--group-002.md#canonical-1132023131100102-0101122333332123-2320101021133230-2230123020222330-1022222131313332-3002001110221222-3111330122331220-3122100003321120) |
| `dynamic_reverse_proxy.resolution_network_type` | [dynamic_reverse_proxy.resolution_network_type](data-sources--virtual_host--reference--group-002.md#canonical-3010130003111302-2130210213012133-2133122313331220-2021130112120331-0021231020013123-3121011111201010-2203011113302120-0333322103130113) |
| `dynamic_reverse_proxy.resolve_endpoint_dynamically` | [dynamic_reverse_proxy.resolve_endpoint_dynamically](data-sources--virtual_host--reference--group-002.md#canonical-2111003331333012-0102002000130221-1313311222331131-3203333123223321-1212303201102123-2220132000103222-3130113303221032-1200202002000221) |
| `enable_path_normalize` | [enable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-0201022300021320-1303303003222003-0131032130300103-2101300200200320-1102022211110023-3002001212313220-0232120011002133-0313322031310312) |
| `http_protocol_options` | [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-2310113310130010-3202002120331033-0311323332301131-0002201322002231-2330330123123022-1120101032001030-2313321211221113-1133330133022113) |
| `http_protocol_options.http_protocol_enable_v1_only` | [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-2120030322212322-2013100321103031-2132112213110203-2310201230132310-0312211110131133-3322120012323020-3003321110133113-0310310012222201) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0120332231310121-0113131120312310-1211221010210013-0031020333330210-1333013003123023-1113211132212222-0002003313032023-2310310221333231) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0030121220133201-0213311121113220-2100110103320320-1221333110111101-0213031001333223-3202303333332302-3010213113200331-0211003311131303) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-2233331223213101-3023030222201313-0211012203010211-2131230000231202-1031111220030233-1211332231333121-1010123001311312-1222312101210023) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1230200223331302-2002002213010322-3013103231233030-3120303222000121-2023110213010322-1303313210001222-0101332311313131-3302301200120123) |
| `http_protocol_options.http_protocol_enable_v1_v2` | [http_protocol_options.http_protocol_enable_v1_v2](data-sources--virtual_host--reference--group-002.md#canonical-3303201210330112-1103310130233322-1233213311213131-3233203100332101-3013230011313001-0311010121312000-1212302030323210-1111201323321210) |
| `http_protocol_options.http_protocol_enable_v2_only` | [http_protocol_options.http_protocol_enable_v2_only](data-sources--virtual_host--reference--group-002.md#canonical-2211100030121313-1121323330100333-3211112023122100-0030201230331100-3103030202000302-1111213101120220-1001021233020211-2022321222210223) |
| `id` | [ID](data-sources--virtual_host--reference--group-001.md#canonical-0123200020210113-3331101011100011-2131023231221200-1302132213002302-0321212210132202-2111232220131010-1213010301331123-3313220022231303) |
| `idle_timeout` | [idle_timeout](data-sources--virtual_host--reference--group-001.md#canonical-2032022311020320-2103200213220300-2103320332232231-2232202010112301-2022222223222103-3121300333203131-2033003232203122-0000321330323223) |
| `js_challenge` | [js_challenge](data-sources--virtual_host--reference--group-002.md#canonical-1233322231003110-0102032003010110-2120003213121113-1012123200113212-2002113320322202-0133131333003033-3333000122301333-0101331330331000) |
| `js_challenge.cookie_expiry` | [js_challenge.cookie_expiry](data-sources--virtual_host--reference--group-002.md#canonical-2313113333022330-1020033223233003-2130201233133003-1002002123131033-0213020231321222-2033002121320310-1133123211230323-1011301023100201) |
| `js_challenge.custom_page` | [js_challenge.custom_page](data-sources--virtual_host--reference--group-002.md#canonical-1302230100332332-2202313211121220-0112313320203313-1211131320323102-3310132321221203-1310313010201213-2113100130022103-0010311132323211) |
| `js_challenge.js_script_delay` | [js_challenge.js_script_delay](data-sources--virtual_host--reference--group-002.md#canonical-0103120220231031-1132002321101122-0112010200210321-3303312133213213-1302000012302123-0010032133330111-1122032303302303-1112323322232230) |
| `labels` | [labels](data-sources--virtual_host--reference--group-001.md#canonical-2321100313100033-1030011101330323-2210010211320131-3111302003012300-2223223121121020-3202023131211212-0321222301012010-3101201022333011) |
| `max_request_header_size` | [max_request_header_size](data-sources--virtual_host--reference--group-001.md#canonical-1112301012203222-0033322101102100-2102113011010203-2202321311313220-3123122012020121-1302131003123102-3213032133120033-2103202330022033) |
| `max_requests_per_connection` | [max_requests_per_connection](data-sources--virtual_host--reference--group-001.md#canonical-2120031230002123-2101102332011300-0002133320120210-1203232121221321-0312032203130200-3003330211201232-0112030211002123-3313310312112033) |
| `name` | [name](data-sources--virtual_host--reference--group-001.md#canonical-0133010022031202-0132230130110103-2120123002212031-3123123320112333-3320310312003111-0103000111012002-1233322113303120-3332322030320101) |
| `namespace` | [namespace](data-sources--virtual_host--reference--group-001.md#canonical-1103220003212213-2311300012331233-3223311312302210-3030301132231022-3102213320310013-3210121011303201-0032023010313300-0333211213131213) |
| `no_authentication` | [no_authentication](data-sources--virtual_host--reference--group-002.md#canonical-1102100302002300-1313133201331100-3113202203002110-2131030023211011-2113223023332022-0203202123003020-0213012011103301-3232331311333110) |
| `no_challenge` | [no_challenge](data-sources--virtual_host--reference--group-002.md#canonical-0333002331231113-1021121120232222-1021202203033121-0100301010000210-0311202222011232-1313320122120233-3202220223102123-3332002100110330) |
| `no_request_limit_per_connection` | [no_request_limit_per_connection](data-sources--virtual_host--reference--group-002.md#canonical-2222231010030030-3130021202122031-3100103020130033-3121222000100032-1102001222302230-2111131103111212-2133112303230032-3021011203213203) |
| `non_default_loadbalancer` | [non_default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-1030023303220330-1233033230000310-3231300000123020-3023333313100103-3300223002022021-2032203120322203-1313212202312012-3321111231213003) |
| `pass_through` | [pass_through](data-sources--virtual_host--reference--group-002.md#canonical-2011230100101232-2231130122231011-0101023133113300-2212113001232301-2201101303031211-2222001202121303-1333000311103303-2303303020030321) |
| `proxy` | [proxy](data-sources--virtual_host--reference--group-001.md#canonical-2033111332122010-1200133213031022-1130303333101001-2032010113303131-0203113033130133-2222303300032300-0333000030331023-2020302012321331) |
| `rate_limiter_allowed_prefixes` | [rate_limiter_allowed_prefixes](data-sources--virtual_host--reference--group-002.md#canonical-0331330100322200-1002100120300333-2120120320132113-3002220323202030-1000321101320011-1133201200033030-1322111110211323-3011210232003110) |
| `rate_limiter_allowed_prefixes.kind` | [rate_limiter_allowed_prefixes.kind](data-sources--virtual_host--reference--group-002.md#canonical-3003022201012312-1020122303201010-1221030322123200-3032321102213223-2011202001121100-2210003123231310-3001133022303023-3030233101233031) |
| `rate_limiter_allowed_prefixes.name` | [rate_limiter_allowed_prefixes.name](data-sources--virtual_host--reference--group-002.md#canonical-2300030130022113-2002103122312233-0332110112313112-3202313123320002-1230311322013002-1031321111000223-2100220012333200-3003023301123303) |
| `rate_limiter_allowed_prefixes.namespace` | [rate_limiter_allowed_prefixes.namespace](data-sources--virtual_host--reference--group-002.md#canonical-0123332313122333-2233322111222220-2232000033013202-0202221023100030-0320330123210302-0121323021303203-1211002111011002-3010202021201023) |
| `rate_limiter_allowed_prefixes.tenant` | [rate_limiter_allowed_prefixes.tenant](data-sources--virtual_host--reference--group-002.md#canonical-0332132321331221-0010120222322311-0232113213223032-0302323220221311-3203021021121303-2301030303122210-0201302331321202-2311133033020220) |
| `rate_limiter_allowed_prefixes.uid` | [rate_limiter_allowed_prefixes.uid](data-sources--virtual_host--reference--group-002.md#canonical-2003130203331213-2112231202131123-0011312113200000-3123033332122320-0200100301022012-0103211130201011-0332102312300101-0322031302120323) |
| `request_cookies_to_add` | [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-1313120312033223-3223302312021001-2210111031213102-0032030313222120-1111023302121322-0321022003120100-1002030012222312-3203122220311202) |
| `request_cookies_to_add.name` | [request_cookies_to_add.name](data-sources--virtual_host--reference--group-002.md#canonical-3210111320021210-0032132032121020-3302313300112212-0032120213000131-3303212001021130-0210100012323310-2201200303232120-3330020121311200) |
| `request_cookies_to_add.overwrite` | [request_cookies_to_add.overwrite](data-sources--virtual_host--reference--group-002.md#canonical-3123011230320133-2030230232131002-2132301112010200-3000221223231033-2313112322222022-1020210231223000-2300113110033321-3213101011220211) |
| `request_cookies_to_add.secret_value` | [request_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-0322303201132233-3331133033102320-3102210333112212-2121003031010100-2022302021021031-0203120133110221-0322311031233013-0123233230031222) |
| `request_cookies_to_add.secret_value.blindfold_secret_info` | [request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-1102332111131111-2312210102230132-2102301032003113-1222233223112201-0103032111122220-3020133233303122-0123020330232333-1122130322203001) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--virtual_host--reference--group-002.md#canonical-0011112001231020-0202022121232302-1313002012223121-3332000123100131-1013300013130302-1030223302012111-3100330301110020-3321032102301203) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.location` | [request_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--virtual_host--reference--group-002.md#canonical-3203200133231320-3031120000321011-2221130111333310-0200203230030211-2231210200203130-0000201320023200-2012003002032200-1212202200213233) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--virtual_host--reference--group-002.md#canonical-0331331030323300-0132232222123003-2033300320120322-1032320321132331-1000223130031122-1221113010222033-3133002212133132-0200230301210200) |
| `request_cookies_to_add.secret_value.clear_secret_info` | [request_cookies_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-3020312203131200-0020022130002033-0321323213310022-2313111212203031-2030211221112111-0213103303011130-1323002013112332-3320202133102220) |
| `request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [request_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--virtual_host--reference--group-002.md#canonical-3232213000032311-3301210102010031-3222232013311122-0233112321032130-2120322312202132-0330000323231203-2202231213101220-1003102131001330) |
| `request_cookies_to_add.secret_value.clear_secret_info.url` | [request_cookies_to_add.secret_value.clear_secret_info.url](data-sources--virtual_host--reference--group-002.md#canonical-0300321231110031-2331012232231000-0123331022213032-3030302211320333-0113230030130112-1201122031033032-0211323200112230-3233131303313211) |
| `request_cookies_to_add.value` | [request_cookies_to_add.value](data-sources--virtual_host--reference--group-002.md#canonical-3022111331213332-0111133010110023-1121102122000323-3311001313022000-3321032110313012-0123130321032020-2320111232303303-0231121300231003) |
| `request_cookies_to_remove` | [request_cookies_to_remove](data-sources--virtual_host--reference--group-001.md#canonical-3121332030132111-1033032022010131-0231011222333033-0011333330301000-0011130210122033-0121203201322112-0010100331231132-3331101133320302) |
| `request_headers_to_add` | [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-0212331111220023-1110003001131213-2230212300133010-0032203310311221-3132120120212033-2213331220010100-0122232301011223-1113212132012332) |
| `request_headers_to_add.append` | [request_headers_to_add.append](data-sources--virtual_host--reference--group-002.md#canonical-0111003322111311-0210311322203231-0033231212032013-0120201220320101-3322301111220313-0130312231102133-1010101013000110-0323013230030323) |
| `request_headers_to_add.name` | [request_headers_to_add.name](data-sources--virtual_host--reference--group-002.md#canonical-2013010322133123-3312231203130101-0001201323033122-0313203013300222-2002002231011321-0320010312120320-1311220121011202-3021332101220200) |
| `request_headers_to_add.secret_value` | [request_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-1113100003100331-1012223330010101-2311201233313310-2202132323023113-1002120103022112-2322022213110322-1002212202302302-0022023101201032) |
| `request_headers_to_add.secret_value.blindfold_secret_info` | [request_headers_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-1110020211333301-0201002333112132-1123323201210102-1122032113112330-0033210303230021-3311212110111200-2322201312102122-0103010302332210) |
| `request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--virtual_host--reference--group-002.md#canonical-3221300001313201-2202013112112012-3210001201320011-0211221022210112-1022212232103113-1310001122321233-3210230123221120-3323312303010110) |
| `request_headers_to_add.secret_value.blindfold_secret_info.location` | [request_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--virtual_host--reference--group-002.md#canonical-1130010110321230-3102120121003322-1010221321200311-1132112033100303-0001301202313013-3130212230131221-1310230033220310-3200222213021312) |
| `request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [request_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--virtual_host--reference--group-002.md#canonical-3303102130111003-1333003000312221-0023001320311002-2133322032102330-1230230222003021-0100332210120222-2321030033213221-0130122233033021) |
| `request_headers_to_add.secret_value.clear_secret_info` | [request_headers_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-3121111213002323-1231101112112231-3101333032202332-0200013011002332-2223020023231333-1133033202321323-2031003023020022-1103101121113301) |
| `request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [request_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--virtual_host--reference--group-002.md#canonical-1200322010133001-2100232100010330-1133302301320120-3201131331201221-0221220222211133-0120320101303330-3211133301010023-2230331232012120) |
| `request_headers_to_add.secret_value.clear_secret_info.url` | [request_headers_to_add.secret_value.clear_secret_info.url](data-sources--virtual_host--reference--group-002.md#canonical-2330021311131111-1132011222001031-0010222130001213-1112132200302132-0322123330122212-1220113330112203-1210310130002200-2211211201220201) |
| `request_headers_to_add.value` | [request_headers_to_add.value](data-sources--virtual_host--reference--group-002.md#canonical-0002221320121320-1033122001021012-2021121302103102-0303312323133110-2231122220223301-3120003322302203-0211231223130313-2203011201111113) |
| `request_headers_to_remove` | [request_headers_to_remove](data-sources--virtual_host--reference--group-001.md#canonical-0132230023120312-0323110133220311-2200100113101302-3302003030102231-1200223330013012-3012320331030001-1002033032102022-2121031120023110) |
| `response_cookies_to_add` | [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-0113212001232011-3133322010021302-3102230131212330-2020322230330311-0130213133032111-3122133022112220-3200232202122232-3312300203123230) |
| `response_cookies_to_add.add_domain` | [response_cookies_to_add.add_domain](data-sources--virtual_host--reference--group-002.md#canonical-0001011020132012-1232111312113132-0330300223301231-1011002313230031-3232220300331003-2300110310311011-3121013011113200-0013021222011120) |
| `response_cookies_to_add.add_expiry` | [response_cookies_to_add.add_expiry](data-sources--virtual_host--reference--group-002.md#canonical-0323331213312223-3021003123002211-1202213220023132-3010013212200011-1022332133021023-3323301323122213-0303012032122120-2201212310221002) |
| `response_cookies_to_add.add_httponly` | [response_cookies_to_add.add_httponly](data-sources--virtual_host--reference--group-002.md#canonical-3331200311131232-2111312122313330-3311022112022321-0210222232203233-2322120132222221-0330300102033233-0002221112121012-2231003222320123) |
| `response_cookies_to_add.add_partitioned` | [response_cookies_to_add.add_partitioned](data-sources--virtual_host--reference--group-002.md#canonical-2030311302202101-1021020302223000-2330003013102031-1031233032103232-3210302012231012-3212130200300211-1000222233330032-0202323320031200) |
| `response_cookies_to_add.add_path` | [response_cookies_to_add.add_path](data-sources--virtual_host--reference--group-002.md#canonical-0211031222022221-1210232233232213-0123111010113301-1210201123011222-3321203311111332-0020311311131231-2211233232100232-0022110201032030) |
| `response_cookies_to_add.add_secure` | [response_cookies_to_add.add_secure](data-sources--virtual_host--reference--group-002.md#canonical-3003210300322131-1210230110002132-0111332020330200-2212311130233302-0330011012323322-1132211132132020-3220101313231112-1023130121100010) |
| `response_cookies_to_add.ignore_domain` | [response_cookies_to_add.ignore_domain](data-sources--virtual_host--reference--group-002.md#canonical-2300121232212121-0121210302102312-2113030112230210-3213202010320220-3000113302302223-2020031013213103-3000003131023032-3302121313313112) |
| `response_cookies_to_add.ignore_expiry` | [response_cookies_to_add.ignore_expiry](data-sources--virtual_host--reference--group-002.md#canonical-0131302033331002-2302213311003033-0320130131231012-2312100333202102-0023100123222031-1213230122321313-3121023211212321-3101233013232101) |
| `response_cookies_to_add.ignore_httponly` | [response_cookies_to_add.ignore_httponly](data-sources--virtual_host--reference--group-002.md#canonical-0301030232121022-0112212033023121-2032223012003223-2230102331131002-3330102031321322-3333222213312032-3120310303312112-0330111332132312) |
| `response_cookies_to_add.ignore_max_age` | [response_cookies_to_add.ignore_max_age](data-sources--virtual_host--reference--group-002.md#canonical-1312333330322012-1230201311201103-2333211023331033-2201220330121222-1321331313012310-1310302302331010-0333121011001310-1121031003120230) |
| `response_cookies_to_add.ignore_partitioned` | [response_cookies_to_add.ignore_partitioned](data-sources--virtual_host--reference--group-002.md#canonical-3110203131101201-2120323010230203-0202113001231330-3013301310123112-2002133221302221-0101033122100131-1331130012322133-2301222312322232) |
| `response_cookies_to_add.ignore_path` | [response_cookies_to_add.ignore_path](data-sources--virtual_host--reference--group-002.md#canonical-2130320121310303-1332032310213100-2222113113013213-2303023222113201-1312310230031200-0122023033133131-1230233332231203-3230131012023332) |
| `response_cookies_to_add.ignore_samesite` | [response_cookies_to_add.ignore_samesite](data-sources--virtual_host--reference--group-002.md#canonical-3021320121212200-3213213331333010-3110311312301132-0231220113203222-2131102100213123-1320001210231333-3131200320223223-1112322321131001) |
| `response_cookies_to_add.ignore_secure` | [response_cookies_to_add.ignore_secure](data-sources--virtual_host--reference--group-002.md#canonical-2302001212203030-1130010112322122-2033200121331232-3200023020330333-1102221111111232-0221131320122112-2011203322210322-1030202221211032) |
| `response_cookies_to_add.ignore_value` | [response_cookies_to_add.ignore_value](data-sources--virtual_host--reference--group-002.md#canonical-3222313001302022-0111130312012223-0231010310010311-1112112223103321-2201303303000332-3322220033221033-1301013100210230-0020213102300331) |
| `response_cookies_to_add.max_age_value` | [response_cookies_to_add.max_age_value](data-sources--virtual_host--reference--group-002.md#canonical-1223332230031002-3021221300023122-1202320001120111-1301132121100020-0021200302320200-1031203103230123-1212100133020002-1023133010220030) |
| `response_cookies_to_add.name` | [response_cookies_to_add.name](data-sources--virtual_host--reference--group-002.md#canonical-0033210331200122-0220202000132121-0133233333200332-0202313322332222-3132303020321000-1030002033013000-2100222230022023-1011132021323013) |
| `response_cookies_to_add.overwrite` | [response_cookies_to_add.overwrite](data-sources--virtual_host--reference--group-002.md#canonical-0131110323003203-3300311130312322-0010021032303102-3033001202211300-3133130232132202-3203131033323323-2310212232320232-0322100132320110) |
| `response_cookies_to_add.samesite_lax` | [response_cookies_to_add.samesite_lax](data-sources--virtual_host--reference--group-002.md#canonical-2033121212113103-0210022033031232-2333211101213203-2320033333020211-0031301210202211-1103111012210301-0213003303311011-0321333020001302) |
| `response_cookies_to_add.samesite_none` | [response_cookies_to_add.samesite_none](data-sources--virtual_host--reference--group-002.md#canonical-3220320102032323-0102112031021131-2101130101011021-3133023100130231-0301332213303331-0232020023003022-0200212031300200-3202111010023000) |
| `response_cookies_to_add.samesite_strict` | [response_cookies_to_add.samesite_strict](data-sources--virtual_host--reference--group-002.md#canonical-0020132200132200-2312023030212103-0101121332113211-0300223301032223-3131310123120132-1030202201211311-3033223330103300-1001010231133231) |
| `response_cookies_to_add.secret_value` | [response_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-0230220011133210-1121201122001221-2000212031123303-2011300012231322-0222223121102113-1020232220032232-2213211003211031-3233111203002212) |
| `response_cookies_to_add.secret_value.blindfold_secret_info` | [response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-0121002103233003-2010103032030132-2022032223201103-3322111330020103-2222021232310220-1223233302111221-0230220220201122-1022002331131202) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--virtual_host--reference--group-002.md#canonical-0012113123020121-1200013002032230-3330322213112113-1020322132300102-0110011223303332-2113013103003112-2300312331121330-3233223133103002) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.location` | [response_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--virtual_host--reference--group-002.md#canonical-0302122213101131-1201112003300311-0100202202232231-3311121103123012-2323231121121112-2100023230202130-0333031121231000-1332322330021013) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--virtual_host--reference--group-002.md#canonical-1020312200013313-2020300101100203-0232321123331320-3121113012310200-1231310002031023-0332120303101003-0112111203122001-0123002323122201) |
| `response_cookies_to_add.secret_value.clear_secret_info` | [response_cookies_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-3313221201232231-0032112231301020-0000102331131233-2230002130021321-1300331221033122-2331012012100222-3333311331212011-1213221010013302) |
| `response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [response_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--virtual_host--reference--group-002.md#canonical-2130100100231230-0111333213301122-2223013210030123-0210223112111132-1332013333011322-0330223010022203-0101320132230022-3033120023232331) |
| `response_cookies_to_add.secret_value.clear_secret_info.url` | [response_cookies_to_add.secret_value.clear_secret_info.url](data-sources--virtual_host--reference--group-002.md#canonical-2322300220130300-2033203000011122-2111200122010030-2331010010200213-2223120132233322-3103001213000023-3221102222232333-0212231210102202) |
| `response_cookies_to_add.value` | [response_cookies_to_add.value](data-sources--virtual_host--reference--group-002.md#canonical-0132212221332200-0112030120132231-0303203332110213-0221133212132012-3112222223320130-3311011121331322-2300303301330321-0232012032110113) |
| `response_cookies_to_remove` | [response_cookies_to_remove](data-sources--virtual_host--reference--group-001.md#canonical-2012011312030331-0033023101223210-1313013112102210-1321032113231022-2003120322020032-1332310021111120-0130022121211301-3322320320103000) |
| `response_headers_to_add` | [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-0101202001030202-2332212330311113-0131001010330313-3313032330010131-2020200230233013-3311312012021031-2120012311323110-1223013022101202) |
| `response_headers_to_add.append` | [response_headers_to_add.append](data-sources--virtual_host--reference--group-002.md#canonical-3303003312131233-3030200320213123-3100123110201120-3131321300002023-2233310232001113-2010231121122122-3320330023303211-3030031322132203) |
| `response_headers_to_add.name` | [response_headers_to_add.name](data-sources--virtual_host--reference--group-002.md#canonical-2003113000313011-3210313331203312-2212323303332232-3131111332022133-1231002120203303-2020000303233020-2210103001332230-3332312130233310) |
| `response_headers_to_add.secret_value` | [response_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3100032201022001-2001120211311323-2303132030311023-1223333122211033-2111333132122321-0032110321210032-1303101313331301-0220323001111102) |
| `response_headers_to_add.secret_value.blindfold_secret_info` | [response_headers_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-0320122313112313-3331203313022000-0020013010012332-0133303123120012-0301223323113003-1323202202311113-0103112102302020-3111312210311233) |
| `response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--virtual_host--reference--group-002.md#canonical-2321220203333331-1210133320202232-0003321312003130-3020101102233113-2330202230100231-2103013212232033-2220000302030330-1233011223003331) |
| `response_headers_to_add.secret_value.blindfold_secret_info.location` | [response_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--virtual_host--reference--group-002.md#canonical-3000332200020302-1200112232231112-2111231332323331-0213223100200310-0203001300120302-2032202130030311-1201102110132112-1032123023322310) |
| `response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [response_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--virtual_host--reference--group-003.md#canonical-1102310020100012-3102100223203311-3110012320302010-0123100022202212-0033032111133122-0103313212300302-3202100320123303-2101331100023100) |
| `response_headers_to_add.secret_value.clear_secret_info` | [response_headers_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-3130022021012020-2111013303010033-1333320203310213-1113300030300221-3010301303313032-0131000023311021-1311011231100022-0033220130331233) |
| `response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [response_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--virtual_host--reference--group-003.md#canonical-3130311230122321-1302312201303300-2201333323000103-2221203213132300-1221320233200232-0202033103310201-1130033232123211-1020312201022013) |
| `response_headers_to_add.secret_value.clear_secret_info.url` | [response_headers_to_add.secret_value.clear_secret_info.url](data-sources--virtual_host--reference--group-003.md#canonical-3021103000001322-1332201012301002-2323332233021010-3223020132221322-3321333002212312-0002112103120230-0002122102112030-2000112103213322) |
| `response_headers_to_add.value` | [response_headers_to_add.value](data-sources--virtual_host--reference--group-002.md#canonical-3322312333311133-0031013222000232-2233131202221213-1100201210300331-3301011233011300-3011131010300113-1200103303323130-0210113003213303) |
| `response_headers_to_remove` | [response_headers_to_remove](data-sources--virtual_host--reference--group-001.md#canonical-0033332310122132-2102211023030310-3331313201111222-0320231303022031-0232220110330333-2233111210213332-1311230032111033-1030311110013200) |
| `retry_policy` | [retry_policy](data-sources--virtual_host--reference--group-003.md#canonical-3121300022323212-3032100030021200-3220123111011103-0022232031022323-2212331012001311-2303233320201201-1331103322001123-2010123121001101) |
| `retry_policy.back_off` | [retry_policy.back_off](data-sources--virtual_host--reference--group-003.md#canonical-2100010313230133-0131312130011120-3021201113200123-0111123210210102-2313130302102213-3321231130312130-2130210232001211-0132213311231233) |
| `retry_policy.back_off.base_interval` | [retry_policy.back_off.base_interval](data-sources--virtual_host--reference--group-003.md#canonical-2203132023322013-0121130211022203-3232303101100320-3232300022223013-3121331123220320-2021300332201203-3201002020331213-1130003023311201) |
| `retry_policy.back_off.max_interval` | [retry_policy.back_off.max_interval](data-sources--virtual_host--reference--group-003.md#canonical-2230122033122210-3301212221320032-2101022132012321-0313311303320300-2030013311022231-1011313330031023-0122310203312220-2101300302001102) |
| `retry_policy.num_retries` | [retry_policy.num_retries](data-sources--virtual_host--reference--group-003.md#canonical-1220312222312301-1032103101122112-0131120022103202-3313321111013110-0222113311201112-0130211121333200-1110212201330031-1102103301033132) |
| `retry_policy.per_try_timeout` | [retry_policy.per_try_timeout](data-sources--virtual_host--reference--group-003.md#canonical-0320211020202131-1110221323012302-0333100001031313-0322023302330120-0201222331000302-2030031122110001-1330003133012202-3121210000302323) |
| `retry_policy.retriable_status_codes` | [retry_policy.retriable_status_codes](data-sources--virtual_host--reference--group-003.md#canonical-2121022023022001-3000132120321230-2023030231120210-3023100021111221-1021123323210330-2002210323013301-3110301013223023-0133022312311100) |
| `retry_policy.retry_condition` | [retry_policy.retry_condition](data-sources--virtual_host--reference--group-003.md#canonical-1330131133121001-0213001123313202-0333010330032300-2200023131011201-3032233303321102-0322311222232033-3011202221002323-1111222331132312) |
| `routes` | [routes](data-sources--virtual_host--reference--group-003.md#canonical-2002112323120121-1202312220122120-3322103020012232-0302122212212332-0222332110320333-1213010203011123-0000002223313331-0233310011312332) |
| `routes.kind` | [routes.kind](data-sources--virtual_host--reference--group-003.md#canonical-1223020302332210-2023312112223323-0221101131313303-3023200210310011-2031033021232212-1012021110032002-0130032212230020-3222220230002003) |
| `routes.name` | [routes.name](data-sources--virtual_host--reference--group-003.md#canonical-2202300320101110-0113011022002232-2110123233001222-1131030303001313-3232302022013011-3223233110113210-2012021133120233-1111023032031320) |
| `routes.namespace` | [routes.namespace](data-sources--virtual_host--reference--group-003.md#canonical-0210112210122000-0322130213000103-2333231131123030-2030311131121213-3031222313311233-3123322202222310-2211113012232110-1123030213303122) |
| `routes.tenant` | [routes.tenant](data-sources--virtual_host--reference--group-003.md#canonical-3320303210103110-2113232323122122-2230323000030202-1112313303333200-1000223103110013-1213132102221031-2133102330130322-1031112033033212) |
| `routes.uid` | [routes.uid](data-sources--virtual_host--reference--group-003.md#canonical-3123331032100112-0230032013120120-0232012012302122-3120122102110232-3332303313212120-2022020010212033-1322322303230002-2123032322023220) |
| `sensitive_data_policy` | [sensitive_data_policy](data-sources--virtual_host--reference--group-003.md#canonical-2113012213211213-0001300300111232-0000020003302202-0323031131213222-2010031320310213-2023110002103000-0233002333102213-0010303110201000) |
| `sensitive_data_policy.kind` | [sensitive_data_policy.kind](data-sources--virtual_host--reference--group-003.md#canonical-1131233313023311-2110020033320000-0020311223322131-0211222333323230-0030120022330022-1210321330303210-2322023311132123-1021330022311000) |
| `sensitive_data_policy.name` | [sensitive_data_policy.name](data-sources--virtual_host--reference--group-003.md#canonical-3323132320003002-0020103200100221-1223211013120003-0123101131101131-0102111300011232-1020000320333220-0322022111123232-0233331300323131) |
| `sensitive_data_policy.namespace` | [sensitive_data_policy.namespace](data-sources--virtual_host--reference--group-003.md#canonical-1230101132333321-0231110301231201-3311323123113322-0231331010203132-3310020311101021-2013203010312033-0300220112033322-2000021123301300) |
| `sensitive_data_policy.tenant` | [sensitive_data_policy.tenant](data-sources--virtual_host--reference--group-003.md#canonical-1023212213333333-2233313120032322-2311231322131322-2113102110330011-1320113113102000-3031303012020202-2232220233102131-1131030333123211) |
| `sensitive_data_policy.uid` | [sensitive_data_policy.uid](data-sources--virtual_host--reference--group-003.md#canonical-2212202010100222-2331323201130111-2313301113223320-2323302021000202-1120203311123113-3330231210332120-1023131010203020-0330312323130012) |
| `server_name` | [server_name](data-sources--virtual_host--reference--group-001.md#canonical-2000133212001013-1130012132023121-0030132122333222-0303323311000221-0303100033122013-2021210110033131-2331333211132010-0220310022330223) |
| `slow_ddos_mitigation` | [slow_ddos_mitigation](data-sources--virtual_host--reference--group-003.md#canonical-1220330111031211-1120011223033003-2123202222131321-1201302331203110-1333233222300203-1200012010311001-1323233031232331-0100332123023100) |
| `slow_ddos_mitigation.disable_request_timeout` | [slow_ddos_mitigation.disable_request_timeout](data-sources--virtual_host--reference--group-003.md#canonical-3202112331010000-2313202003002100-0032311010100012-3120320312113322-2121103100232313-1010200233110313-3312323002331313-2323013233203122) |
| `slow_ddos_mitigation.request_headers_timeout` | [slow_ddos_mitigation.request_headers_timeout](data-sources--virtual_host--reference--group-003.md#canonical-2201022000222211-2233111121033301-2331123120010132-0323013203231112-2102022322113110-2313110030333310-0011100210121223-0210320030232231) |
| `slow_ddos_mitigation.request_timeout` | [slow_ddos_mitigation.request_timeout](data-sources--virtual_host--reference--group-003.md#canonical-3230302201331122-3001013330302310-3223123011023012-3101202230323030-3121033310132003-3001132133233333-3011121011033001-0311321300131312) |
| `tls_cert_params` | [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-3032031001222133-3201022301111333-2323000130331230-3210112110333013-1013322112101202-3313113011233020-3302022013222012-2000301331222210) |
| `tls_cert_params.certificates` | [tls_cert_params.certificates](data-sources--virtual_host--reference--group-003.md#canonical-3221233032310032-1212121233210100-0122010102321131-0033111303303211-1221232112012301-0301101013322103-0011110230221111-0213323331330132) |
| `tls_cert_params.certificates.kind` | [tls_cert_params.certificates.kind](data-sources--virtual_host--reference--group-003.md#canonical-2331312322233001-3121000113203011-2233010221331311-2322002221201221-2110010232321003-2233300323021321-3302121113321032-1333213011131300) |
| `tls_cert_params.certificates.name` | [tls_cert_params.certificates.name](data-sources--virtual_host--reference--group-003.md#canonical-0011133212323013-2121300212212320-1130201002031120-2301132023313210-1112312113321130-1132330032332313-1013222223323111-3123320231112223) |
| `tls_cert_params.certificates.namespace` | [tls_cert_params.certificates.namespace](data-sources--virtual_host--reference--group-003.md#canonical-2010313321322010-2113320011031333-0202230332321300-0312330122133123-1003202103330303-3020311211200213-1030130312102012-3030113303203323) |
| `tls_cert_params.certificates.tenant` | [tls_cert_params.certificates.tenant](data-sources--virtual_host--reference--group-003.md#canonical-2302323213233122-0002011332320311-3000323113031301-3132302330131222-2311233210310120-1333323030112103-3113001012102333-3003003202101332) |
| `tls_cert_params.certificates.uid` | [tls_cert_params.certificates.uid](data-sources--virtual_host--reference--group-003.md#canonical-3311121320100133-1131023023331100-1312203313001132-3101201202113022-3123230123323222-3011311301200213-1203022211211210-3003213132012223) |
| `tls_cert_params.cipher_suites` | [tls_cert_params.cipher_suites](data-sources--virtual_host--reference--group-003.md#canonical-3300013223123001-2110202103320120-1200121330023013-3332001132021311-0220123222132132-2203303231001013-0331100132323300-1331320322002033) |
| `tls_cert_params.client_certificate_optional` | [tls_cert_params.client_certificate_optional](data-sources--virtual_host--reference--group-003.md#canonical-0021020200323331-1101002130221300-3220121322010201-3023233303300022-0013223200023131-0221100021322011-1033001331022321-2131111100333022) |
| `tls_cert_params.client_certificate_required` | [tls_cert_params.client_certificate_required](data-sources--virtual_host--reference--group-003.md#canonical-1310331110032112-3002212333220332-0011213300131032-1200013023013220-3212021223313301-0302122231011101-1321020103130121-2013232113023233) |
| `tls_cert_params.maximum_protocol_version` | [tls_cert_params.maximum_protocol_version](data-sources--virtual_host--reference--group-003.md#canonical-1131110330012331-2221211112212321-3110313301221231-2330203202132300-3323300020312333-2122033010223003-3112221320102130-2323031130220211) |
| `tls_cert_params.minimum_protocol_version` | [tls_cert_params.minimum_protocol_version](data-sources--virtual_host--reference--group-003.md#canonical-0101030003010310-1113103230302113-3230113031002121-1100010200200232-2003020012321302-0223211012123202-1001212101020032-3201222102003001) |
| `tls_cert_params.no_client_certificate` | [tls_cert_params.no_client_certificate](data-sources--virtual_host--reference--group-003.md#canonical-0210111112322023-2302233322300033-2211201100101112-3302123300112212-0022310311221102-1112120223020133-1232222230213123-2012112021333133) |
| `tls_cert_params.validation_params` | [tls_cert_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-2300302222301320-1013220320220122-2011112013302323-0003012133101110-2211222202003322-3003231101333301-1321221232232330-1330130311101123) |
| `tls_cert_params.validation_params.skip_hostname_verification` | [tls_cert_params.validation_params.skip_hostname_verification](data-sources--virtual_host--reference--group-003.md#canonical-3320123120020312-2330230202230021-3222031121223112-0033231310023012-0121101333011131-3223211331322313-3003333221300120-1220202221101312) |
| `tls_cert_params.validation_params.trusted_ca` | [tls_cert_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-2120312133211320-0213010030203123-2212023213213332-3301130020302332-2300100101203032-2002001113213200-1120233211121010-0122112323012112) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list](data-sources--virtual_host--reference--group-003.md#canonical-2322201223331011-1022202131013311-0012313030331113-1201110100110101-0000101131113030-2333112220310021-3010212122021220-0000300221301301) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.kind](data-sources--virtual_host--reference--group-003.md#canonical-2322133201223202-0211300020200333-3112010321130301-1023303210332031-3130203131022211-3133123200323202-1301003231003211-2103110113320300) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.name](data-sources--virtual_host--reference--group-003.md#canonical-0203111133022121-0012200212331321-3333021230121123-2012003010301332-1033013301210213-0202330201300022-0231302103310221-1220120102102220) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--virtual_host--reference--group-003.md#canonical-0020100332200221-2113321020230013-0120130211110000-2022223031332200-3103132231131033-2301330123211133-1010301332010300-1020313311200332) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--virtual_host--reference--group-003.md#canonical-3210031221311021-0313012001202012-1220110122213101-2321102223232113-0223032000021133-1002301012003020-3311232103100303-0022010230332100) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.uid](data-sources--virtual_host--reference--group-003.md#canonical-3301020111002020-3320012021321331-0131013011312100-1002220220322213-0212131123100033-0320001330223211-1230002230201103-1200233010300100) |
| `tls_cert_params.validation_params.trusted_ca_url` | [tls_cert_params.validation_params.trusted_ca_url](data-sources--virtual_host--reference--group-003.md#canonical-3121111203320320-3330323231303011-3031120310021332-2112002002112231-0030031211132121-3312023201203212-0233001313311131-0011311310103012) |
| `tls_cert_params.validation_params.verify_subject_alt_names` | [tls_cert_params.validation_params.verify_subject_alt_names](data-sources--virtual_host--reference--group-003.md#canonical-0202011000311301-2131110123021013-3132112300332303-2121210000202202-1221313221210323-3223123000000313-2122012321133201-3030321130313212) |
| `tls_cert_params.xfcc_header_elements` | [tls_cert_params.xfcc_header_elements](data-sources--virtual_host--reference--group-003.md#canonical-3310302001020000-3312322010233310-2000321003231101-3322201300010222-2300321120032212-3012303220211021-0313022203333323-0001021210021231) |
| `tls_parameters` | [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-0013030212110331-2322313130021202-1331130021032333-0333100022023122-2301331212332300-0032103210001031-1030010233201111-2231131010212101) |
| `tls_parameters.client_certificate_optional` | [tls_parameters.client_certificate_optional](data-sources--virtual_host--reference--group-003.md#canonical-3221212102110032-2212333010322130-1110301031211112-0233300112310032-1220113321013000-3223321211103002-1021230021203211-3100311011212110) |
| `tls_parameters.client_certificate_required` | [tls_parameters.client_certificate_required](data-sources--virtual_host--reference--group-003.md#canonical-1322121022021001-1211013131003133-2233122232330113-2203121213300003-3132332233300311-2211003133210313-0312020003120012-2020002303032033) |
| `tls_parameters.common_params` | [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-1122233013102130-2021122230220223-3002031003230030-3020321300000222-1322010202110003-2122231302323200-2203220302313131-2323311001200032) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](data-sources--virtual_host--reference--group-003.md#canonical-0210103132203322-1000023122230330-1203201210212103-2301100121230113-0221302200221113-2212212100332303-0021112331001121-3310301321303211) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](data-sources--virtual_host--reference--group-003.md#canonical-3311002010103122-1013132310120330-0311130211112101-1233333323023213-2023300300221103-2123210233300320-1331211100323023-0123321210110212) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](data-sources--virtual_host--reference--group-003.md#canonical-2010113220100203-3122102002113033-1302302001303030-1203203331122233-1122311133000133-1033311000100001-0302011121113000-0322333003012230) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-3122211203131331-1110033333310310-1013132323031321-2023210123002111-2321322320000113-3321322313230001-2322021021212302-0321000021133001) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](data-sources--virtual_host--reference--group-003.md#canonical-0130131200233131-3102001332233101-3111311001312113-0021333333233323-3201133301032220-3131301230200012-1213301010211021-3300123211231210) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](data-sources--virtual_host--reference--group-003.md#canonical-1013100311311122-3133332231121321-0011003022010013-0022130200031322-1013213110331100-1302203032133203-2132003020022003-2203222120223211) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--virtual_host--reference--group-003.md#canonical-0333323320231131-3112231210321031-1101203220011032-1030113210033133-0032022113023211-3302110133201303-1031332300002001-1221300003002300) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](data-sources--virtual_host--reference--group-003.md#canonical-1112213221331013-1033133021011102-3230011200010333-1130220130001201-1201303333121131-0110032021011112-2132112133001020-0221123210030120) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](data-sources--virtual_host--reference--group-003.md#canonical-1032000020301032-1220002203030322-3011131031012231-2021333010023032-1333201221002002-2120120103230222-3022023201322203-2202221221002310) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](data-sources--virtual_host--reference--group-003.md#canonical-0132213332013300-3330132103202111-1123122112012003-2220302132012101-2032311122122211-2212023200322032-2310023213313333-3320300322330310) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-1122221313200001-1233111222132323-2102220222033021-3301220113113232-2223011220201030-0020102100120322-0331001311223303-3303202102232133) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--virtual_host--reference--group-003.md#canonical-2330012103000311-3301320210020113-1030130102020231-1300322021030020-1003010232212003-1002333231011310-2311231013322011-1322010331122310) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](data-sources--virtual_host--reference--group-003.md#canonical-0123303221223230-0110023130232133-1122313022202321-1321230303231133-2131013033210032-1100301301221121-1113300011322213-1131023312022123) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--virtual_host--reference--group-003.md#canonical-2231210311332313-1112022311211300-0200211232013110-1313032332000221-1211331230110022-3312311202120110-3022202310100322-3313201302021230) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-0102230201001020-3303302322012223-3222113230100120-0223231012231200-0303203121101011-1131310323021032-3030033333011100-3333000032000002) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--virtual_host--reference--group-003.md#canonical-3332122332302233-2300202203032313-0113122321220012-2122120032220312-2103212212222201-2001123100233031-1232222020320132-0331300103212101) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](data-sources--virtual_host--reference--group-003.md#canonical-0323302110030333-1013121220021123-1301121113031233-3131020213232021-3102311233211130-2311211201200132-1100031320002011-3221332202201013) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](data-sources--virtual_host--reference--group-003.md#canonical-2110303302121311-1330010201013331-1133321110131010-2033121213220212-3012002021330332-0133311012231022-2233202131232210-3203230332212322) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-2212012212101320-3112001032323123-2033201010100210-2133103003303313-3033022003211212-0213131112032203-2233331211102110-2032033120003310) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](data-sources--virtual_host--reference--group-003.md#canonical-2030113320233213-0220133300211011-3122221331030231-1012202013002210-1200223021303102-3103213121121100-0323123213331301-0202123200321110) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-0110222322332222-3321211030330203-0222013221030130-2231232312323102-1203032313103202-0303220312110330-0330122230200302-0230033112222122) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--virtual_host--reference--group-003.md#canonical-1033000020332202-0023201321102003-0021322013210030-0121310110201111-0200220132013113-3211000033320120-1110013120222300-1321100223231013) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](data-sources--virtual_host--reference--group-003.md#canonical-3213102212203012-1011323021323312-2133323212032212-1110122023202100-3202131112120123-3220230301330321-3033312132230303-3133002212001320) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](data-sources--virtual_host--reference--group-003.md#canonical-3321030310313130-0123031000211330-3011131020210111-0112032230020130-0113212300003233-1203020021031122-1101202303111221-3022122322311321) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--virtual_host--reference--group-003.md#canonical-3331123303122121-2031001311331303-2232012031321123-1012101132310002-3201200202011233-1201111103301222-0021112030303200-1333320310013212) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--virtual_host--reference--group-003.md#canonical-2120122322022312-2330012223232210-0321010113112312-1203331221102222-0221032020003301-0132030133322001-1232000023320122-2202312001130002) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](data-sources--virtual_host--reference--group-003.md#canonical-1322130113222233-3221102020013002-2311301131232222-2231011212310112-2133022103213020-0022223130120321-1020223010310320-0213033333200123) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](data-sources--virtual_host--reference--group-003.md#canonical-2221202123013033-0301222212222230-0002103322012211-1223210100330122-3332211023131310-0300223232202130-1310311203231300-2301201132223210) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](data-sources--virtual_host--reference--group-003.md#canonical-2030332013012000-2133110033200303-2303032112332220-0031023223132301-2031331013200132-3121311113000130-3113133002323002-3230013220202200) |
| `tls_parameters.no_client_certificate` | [tls_parameters.no_client_certificate](data-sources--virtual_host--reference--group-003.md#canonical-2003121333303301-2111312111203022-1320320130233311-3033000220012333-0201022301310033-1312123300120330-3201231231111210-3213012001011200) |
| `tls_parameters.xfcc_header_elements` | [tls_parameters.xfcc_header_elements](data-sources--virtual_host--reference--group-003.md#canonical-3200332110312223-1011311221201220-1310300302322102-3213122331332330-1330232002303333-0010023311221110-3221003021203101-1030330303010122) |
| `user_identification` | [user_identification](data-sources--virtual_host--reference--group-003.md#canonical-3332310012322303-2301321002210033-3200311030002213-3100332131002300-3021122202111201-1333223133120302-0111201333302303-2100033020001110) |
| `user_identification.kind` | [user_identification.kind](data-sources--virtual_host--reference--group-003.md#canonical-1211032333222130-1113032311112331-1322112023200233-2221210102302032-0111132202230233-0313311101113121-2231112101110013-2110120213032300) |
| `user_identification.name` | [user_identification.name](data-sources--virtual_host--reference--group-003.md#canonical-1302023000122200-1233200213212012-0310002020132131-3212003200033311-1103203002212221-0012230220321313-0101102200232223-2111101323101303) |
| `user_identification.namespace` | [user_identification.namespace](data-sources--virtual_host--reference--group-003.md#canonical-1222313323133111-2012203111100120-2111302113221303-3130221231213012-0033000012131020-0301120010010010-1321330030230133-0220333320021122) |
| `user_identification.tenant` | [user_identification.tenant](data-sources--virtual_host--reference--group-003.md#canonical-1131031121132022-3200223321130120-3101323011322120-1220023033333233-3030131103211311-0001230112030013-0203121130311021-2311300302211020) |
| `user_identification.uid` | [user_identification.uid](data-sources--virtual_host--reference--group-003.md#canonical-0233310231303031-2221213101031123-0033311303301102-1023330030323230-2012313013130211-2322013030310211-2332122002033020-2220002133001131) |
| `waf_type` | [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-0212131312112320-3010013102333133-1003210023231021-0022011032133223-0310111333100300-0332102210001333-3113231121103003-0013001311203303) |
| `waf_type.app_firewall` | [waf_type.app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-3000222221111110-0300100323231110-3213130110203022-3102100032020310-1302332210112312-3031031231222333-0210011003021121-0131313212222203) |
| `waf_type.app_firewall.app_firewall` | [waf_type.app_firewall.app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-3031320302232103-0313103102301121-3310310322030220-1331222233331322-3323312112300120-0132330303223100-3230033103103121-1302130133000100) |
| `waf_type.app_firewall.app_firewall.kind` | [waf_type.app_firewall.app_firewall.kind](data-sources--virtual_host--reference--group-003.md#canonical-2122301313000122-3213122132202331-0303320313121310-1013233322231122-2033012022001310-1012311223231011-1122320130002211-0023012113203300) |
| `waf_type.app_firewall.app_firewall.name` | [waf_type.app_firewall.app_firewall.name](data-sources--virtual_host--reference--group-003.md#canonical-2210021022020201-1101003131010303-1033011311003033-3313112311133020-2212001121211021-1100213001032311-3233011301030120-1112313221311113) |
| `waf_type.app_firewall.app_firewall.namespace` | [waf_type.app_firewall.app_firewall.namespace](data-sources--virtual_host--reference--group-003.md#canonical-3033002313132221-2232211031303012-3110113302303333-3021112021230110-3321323113210103-0223221211001323-0302100133213212-1021312212203133) |
| `waf_type.app_firewall.app_firewall.tenant` | [waf_type.app_firewall.app_firewall.tenant](data-sources--virtual_host--reference--group-003.md#canonical-3331131032302210-1323121002333201-1230311011223103-0121133132223223-0130311111330201-0301032302331131-3303320103113200-0311132221113133) |
| `waf_type.app_firewall.app_firewall.uid` | [waf_type.app_firewall.app_firewall.uid](data-sources--virtual_host--reference--group-003.md#canonical-1013211020013021-0011321100123322-0231222213313331-2021120130333111-3211303102212321-2133230011012132-0202012022103100-1012310300010120) |
| `waf_type.disable_waf` | [waf_type.disable_waf](data-sources--virtual_host--reference--group-003.md#canonical-0013120130201131-1002011002120203-2223320033320023-3202121202131112-3200222111030331-3131000003002201-2333113123200223-2321301120031311) |
| `waf_type.inherit_waf` | [waf_type.inherit_waf](data-sources--virtual_host--reference--group-003.md#canonical-3010301221131302-2113032200130303-1223220212210300-0101211011121200-2020201320122112-3013320110313203-2213221213333033-2120022000130002) |

<a id="canonical-1221122011200000-1213210223103213-2122002220131110-1102331210100000-0210211113331020-0210321202301000-1311303312013133-3132111321202031"></a>

## Next pages — Property reference / 002321300333 / 27

- [advertise_policies](data-sources--virtual_host--reference--group-001.md#canonical-0100020130333221-2323111112123131-3120132302111122-1332303033101021-0310121032121322-1122013301131232-0312121331100133-1033121331110112)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- [buffer_policy](data-sources--virtual_host--reference--group-001.md#canonical-1213301223231000-1113301202120100-0322200112212300-3233130103220223-3123033012131112-0321102133021131-3201032002203320-3120102012111033)
- [captcha_challenge](data-sources--virtual_host--reference--group-002.md#canonical-3030320121132033-0113101012333213-2333033121300031-3102221202303210-3222022210330032-0103332033311333-0233333333103003-0313010030101011)
- [coalescing_options](data-sources--virtual_host--reference--group-002.md#canonical-1311333121020231-2023022200000300-2230331133331031-2300321203030302-1000030032223302-0132000331210013-0110211201102210-0011003111130011)
- [compression_params](data-sources--virtual_host--reference--group-002.md#canonical-0113113212120330-2002000130223321-2000320201233130-1100212010220012-2220231112033113-3100133020203133-3032232101303231-1021031112133031)
- [cors_policy](data-sources--virtual_host--reference--group-002.md#canonical-2210323031131321-0313202301131023-3213332203010030-0103323031221230-3302100002020020-3202222112020312-2232021222011023-1201032103301123)
- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-1032321323223020-2022231132220113-3011022121331033-0131132302111121-0112122202013132-2121330303333120-2133303323113131-3121123003303012)
- [default_header](data-sources--virtual_host--reference--group-002.md#canonical-0223221131230033-2031313102132311-2030022100211130-0021201230212110-3110310230000301-0000220012010000-1002213210103013-0320220200333210)
- [default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-2332201222012331-2020200003133331-3003021002300031-1121110112222301-0202020120202032-2121220121022313-1310230211023310-1000212330230231)
- [disable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-3323313313333101-1022230003023322-1121321233003231-3321002002120031-3231313020123102-0332221122301230-2301200331132022-2132211330222112)
- [dynamic_reverse_proxy](data-sources--virtual_host--reference--group-002.md#canonical-2320322030133330-2313030201332123-3133200313313203-0300032102221333-1001120222131112-1330331333122030-2312013001023021-0010002323330332)
- [enable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-0101130103132132-0220121022213231-2333332112030130-0121011300022021-3031032103223222-2102320203302331-2033233303102321-3213203032133202)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- [js_challenge](data-sources--virtual_host--reference--group-002.md#canonical-0311311032210221-1213332233330222-2103321233000202-2020120320331211-3131003020011333-1211120300112103-3021112313203310-2031301312202110)
- [no_authentication](data-sources--virtual_host--reference--group-002.md#canonical-1233312122313321-2011202230001010-3102133312211200-1130003122002233-2223112232011102-0131310120111220-3120020032023010-0313002202102123)
- [no_challenge](data-sources--virtual_host--reference--group-002.md#canonical-3110101102113030-0110100111222000-3312320302213100-0201010223313300-1213232220022000-2133222310110000-2331222120121232-2212130222311233)
- [no_request_limit_per_connection](data-sources--virtual_host--reference--group-002.md#canonical-1110112230220221-2303010133322201-1101300033010303-0103123220222322-3010312001321310-0123120221200321-1022302003222231-0020320101133303)
- [non_default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-2023112320112123-3300121231232030-1100122000331201-1003213223120203-1222330212012300-3302102223112001-3102332220100221-1121233221011031)
- [pass_through](data-sources--virtual_host--reference--group-002.md#canonical-1202312331000133-2122223121103310-2100301313303013-2310311201301110-2210333202010230-0211311013211222-0101120011203120-3303021112101311)
- [rate_limiter_allowed_prefixes](data-sources--virtual_host--reference--group-002.md#canonical-1133310002220332-1103221000011023-3032223311312122-2321113131121222-2110333100201023-1011110232221212-1300111103011301-1303120213023022)
- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3012001313211032-0121110012120332-1211212103012331-1030201012110112-3012311230231023-0312310313121023-3122223302301012-1011130302002022)
- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3303213013223123-0333021210322333-3000030200231312-1002303202333220-2113133012232131-0030321011113311-3012010300230300-3131003321220113)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-1122301030300332-0033230331023011-1102233120003323-3100320132003322-0030000111220313-3302203220201033-0222213312322000-2123201002112113)
- [retry_policy](data-sources--virtual_host--reference--group-003.md#canonical-1021031030330102-2100323023122311-3003021113211121-2312223333303303-2002222022211201-0131100013001013-2031022313222112-2203022023311021)
- [routes](data-sources--virtual_host--reference--group-003.md#canonical-3211011200123230-0330211200101113-1112022302023320-3120120232213300-3223122013100230-0003102201231003-1303322122132321-0212222131122110)
- [sensitive_data_policy](data-sources--virtual_host--reference--group-003.md#canonical-1111130000112110-1102321322100210-0013310223010100-3213130303222101-3131013201101301-1000021000133131-0023212331022302-0120103003321211)
- [slow_ddos_mitigation](data-sources--virtual_host--reference--group-003.md#canonical-2020302310011113-0110101030032202-1233233331203200-0321332220211010-1220330200213312-3321332320232003-2113021203322331-0330021133123211)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [user_identification](data-sources--virtual_host--reference--group-003.md#canonical-1002302233122000-1323022011111121-3132010100020022-1310010213200321-0312233032021223-1020313001300323-0020033031030302-1102203102302301)
- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0100020130333221-2323111112123131-3120132302111122-1332303033101021-0310121032121322-1122013301131232-0312121331100133-1033121331110112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330012203323123-1110312021311121-3111300220210322-3011011002100321-2222212002000031-1100323033333003-0030021320332110-1010220212111122"></a>

## advertise_policies — advertise_policies / 120013110333 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- advertise_policies

<a id="canonical-0011102110223101-0212323101310331-2023220110030200-1001323010012021-3113203200202221-0312301121213031-3231020023122123-2203011222220211"></a>

Type: `"list"`. Computed.

Advertise Policy allows you to define networks or sites where you want a VIP for this virtual host
to be advertised. Each Policy rule can have different parameters, like TLS configuration, ports,
optionally IP address to be used for VIP. If advertise policy is not specified then no VIP is..

Upstream description:

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

<a id="canonical-1010233030100300-2011033101002033-3121032113021223-1311122113310111-2021233110323303-2112201030031311-1100302303232100-2230121303201113"></a>

## Direct properties — advertise_policies / 120013110333 / 3

<a id="canonical-3013230210303330-2021000312330303-2111023102333112-0101003022211221-0230000133232203-2313100222122212-1223223112231333-3221131121210220"></a>

<a id="canonical-2033000133111322-2333330023010323-3331012033111233-3202203013310332-0233120223021020-0332211330300310-1331321332302310-1220231303133103"></a>

## kind property — advertise_policies / 120013110333 / 4

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

<a id="canonical-3002303023302122-2022213221232200-3011013102312130-0310103331100030-2001203123220121-3323323130200331-1130133211313122-1110000232002132"></a>

<a id="canonical-0220131312122323-3311013211333110-1132220103111203-2113301202030203-3002310100013032-2301022233021010-2321331320301001-0003233021300201"></a>

## name property — advertise_policies / 120013110333 / 5

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

<a id="canonical-3121322320123212-0111001202122000-3032010001211200-0330111311101313-1220011021311202-3023212003303133-3130003023100203-0010002233212033"></a>

<a id="canonical-3123211023232032-2030210222200120-1310200300211022-1120332301233112-3130223200321133-3130323303100033-3010201221120333-3103011120300113"></a>

## namespace property — advertise_policies / 120013110333 / 6

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

<a id="canonical-3313020203303211-1003300231330230-0111100022203200-0020230132122230-3132120222230331-0330102130120332-3333302332100233-3333233220000011"></a>

<a id="canonical-0033323313233203-1301220112110002-1102231333000003-0320233000101332-0302100200030300-3112211331312110-0211223333131020-2211320102331133"></a>

## tenant property — advertise_policies / 120013110333 / 7

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

<a id="canonical-3021122302001010-3012210120022230-1131231113211310-0022300311330332-3100313013231322-0012311233213112-3012202222331030-1120332233220223"></a>

<a id="canonical-2012330130000032-0012123312023132-1003312013130333-2230131012112303-0323321320200122-3232333102221311-0021220102023033-3132002113203333"></a>

## uid property — advertise_policies / 120013110333 / 8

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

<a id="canonical-1310332023222202-1122230130130012-1010130013222330-3221112003122020-0301011211020223-0033030321233322-1020311132030230-3003120032112221"></a>

## Next pages — advertise_policies / 120013110333 / 9

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203101211002113-3213220231101032-1012112312212110-1221021221011301-1102320202022002-3301000130312323-3013030330310123-1231131100300112"></a>

## authentication — authentication / 010230003121 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- authentication

<a id="canonical-1110113100202123-2230111331000202-3131210333023303-2221020221222112-0131013231123111-0123202113213213-0313312033121100-1332022222100032"></a>

Type: `"single"`. Computed.

\[OneOf: authentication, no\_authentication; Default: no\_authentication\] Authentication related
information. This allows to configure the URL to redirect after the authentication Authentication
Object Reference, configuration of cookie params etc.

Upstream description:

Authentication related information. This allows to configure the URL to redirect after the
authentication Authentication Object Reference, configuration of cookie params etc.

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

- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-1110113100202123-2230111331000202-3131210333023303-2221020221222112-0131013231123111-0123202113213213-0313312033121100-1332022222100032)
- [no_authentication](data-sources--virtual_host--reference--group-002.md#canonical-1102100302002300-1313133201331100-3113202203002110-2131030023211011-2113223023332022-0203202123003020-0213012011103301-3232331311333110)

Select alternatives according to the provider validators above.

<a id="canonical-0302233300133000-0012213021103100-0300312030310121-2022033120231100-1310330122303032-0313210120220322-3303012222122012-1303233311131110"></a>

## Direct properties — authentication / 010230003121 / 3

- [auth_config](data-sources--virtual_host--reference--group-001.md#canonical-1123112032100310-1210203102021323-0022120331211113-3131013132130113-1222310013320231-1311100131130310-3313232032012223-1002300202131222): complete subsection reference.

- [cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-2303003000301221-3111010311223023-1111031330330013-2131331132220311-2210233002000103-0130120011102211-3303031132200120-1023232302002330): complete subsection reference.

- [redirect_dynamic](data-sources--virtual_host--reference--group-001.md#canonical-1113230233030320-2122313120223033-2133232131033220-0123132323312120-2112121323200101-0011023123013013-0323203211103331-3113011110132312): complete subsection reference.

<a id="canonical-0021332021332103-1221101003101101-3130101310320333-0011113322032121-2303312333123202-1023131110330322-3232332301021323-1131033030122313"></a>

<a id="canonical-2121013301222103-3312230222213333-2320130021223323-0232111110222102-1100303030301110-2032223312230212-1031200022110001-3300013313300232"></a>

## redirect_url property — authentication / 010230003121 / 4

Type: `"string"`. Computed.

Exclusive with \[redirect\_dynamic\] user can provide a URL for e.g https&#58;//abc.xyz.com where
user gets redirected. This URL configured here must match with the redirect URL configured with the
OIDC provider.

Upstream description:

Exclusive with \[redirect\_dynamic\]

user can provide a URL for e.g https&#58;//abc.xyz.com where user gets redirected. This URL
configured here must match with the redirect URL configured with the OIDC provider.

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

- [use_auth_object_config](data-sources--virtual_host--reference--group-001.md#canonical-0221300013323210-3210001133323033-2011102333111333-1231031321131010-0000100232231231-2230121131331331-1300100003222313-1133220102301101): complete subsection reference.

<a id="canonical-0123231021100032-1021222132210121-0330130121113122-1202331211103122-0220123210102312-3300030000030102-1132301002121100-0131032202220313"></a>

## Next pages — authentication / 010230003121 / 5

- [authentication.auth_config](data-sources--virtual_host--reference--group-001.md#canonical-1123112032100310-1210203102021323-0022120331211113-3131013132130113-1222310013320231-1311100131130310-3313232032012223-1002300202131222)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-2303003000301221-3111010311223023-1111031330330013-2131331132220311-2210233002000103-0130120011102211-3303031132200120-1023232302002330)
- [authentication.redirect_dynamic](data-sources--virtual_host--reference--group-001.md#canonical-1113230233030320-2122313120223033-2133232131033220-0123132323312120-2112121323200101-0011023123013013-0323203211103331-3113011110132312)
- [authentication.use_auth_object_config](data-sources--virtual_host--reference--group-001.md#canonical-0221300013323210-3210001133323033-2011102333111333-1231031321131010-0000100232231231-2230121131331331-1300100003222313-1133220102301101)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1123112032100310-1210203102021323-0022120331211113-3131013132130113-1222310013320231-1311100131130310-3313232032012223-1002300202131222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213230223123100-0321123012001102-3112012201003220-0102013000112203-0331120211120032-3323033331120201-0200130002022203-3310013323123220"></a>

## authentication.auth_config — auth_config / 013221011112 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- authentication.auth_config

<a id="canonical-1003023202213212-1301232202222102-1230002021210033-0113311211002030-0102100110110032-1232122301102200-1220020000203202-0231131131033021"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3302110230130312-2020112132113121-2132112001100030-1332003212332033-1030233032103322-2222120311012022-0120022210201133-3203202223031110"></a>

## Direct properties — auth_config / 013221011112 / 3

<a id="canonical-0223022003300122-0131031023021230-1112002002133221-1213233221232131-1310021331132312-3023230122021110-2032213331113122-3301322013321320"></a>

<a id="canonical-0131211300311101-1002201202232223-3232310233130200-1312012323130002-3000113102323100-0130211300033230-1310310203210203-0100010200332030"></a>

## kind property — auth_config / 013221011112 / 4

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

<a id="canonical-2103020112302120-0110211313000110-1310302023230013-0230023213310231-2010002321102030-2001310110232111-0002110311303231-1230320120230131"></a>

<a id="canonical-1233100201002300-2323102230130211-2321323113311322-1110322122222333-3212101231221311-1333232110011222-1120020320302032-0310110023331121"></a>

## name property — auth_config / 013221011112 / 5

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

<a id="canonical-3223331333203110-3321120012301122-2112100001213321-2303012111202310-2321102233211222-3210210131201111-1313111121020313-2031232020033231"></a>

<a id="canonical-3000033323110000-2323031201000032-3310011302001230-3222210112130013-1223300311100121-2031230223132023-2122331221222030-0132001220202310"></a>

## namespace property — auth_config / 013221011112 / 6

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

<a id="canonical-3313200301300210-0023011110220222-2113213013321023-1002123200213203-0223212332231201-3102130020123100-0101201111232222-3111310010203133"></a>

<a id="canonical-1013132332022233-1123100312330221-2332023032031223-0300212122020110-0113021203131131-3232331103132102-1013013201123320-0302002211123221"></a>

## tenant property — auth_config / 013221011112 / 7

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

<a id="canonical-1303100022102322-0021022022321111-0133012000201220-2131022021030013-1201010011112133-2333010323111210-2211223331330020-1323011001023113"></a>

<a id="canonical-1311221030323011-2313301320231020-0213322212011310-1302123031001011-0011323031203313-1330022030022210-1323132002232001-3023003001102320"></a>

## uid property — auth_config / 013221011112 / 8

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

<a id="canonical-3012313212033323-1123001310333332-3100120321012231-1032101103120030-2011022300303202-3131020313023211-0331310333320102-0030332220303112"></a>

## Next pages — auth_config / 013221011112 / 9

- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2303003000301221-3111010311223023-1111031330330013-2131331132220311-2210233002000103-0130120011102211-3303031132200120-1023232302002330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303130330220222-3110122213312003-0202310313330331-2033331130011323-0033033221032322-0131122210133032-1113121311112210-2031311020022001"></a>

## authentication.cookie_params — cookie_params / 033220211200 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- authentication.cookie_params

<a id="canonical-0333203001003223-3030333220013012-0323210203130022-1220023201223030-0111220111100013-1231200223322021-3020002110121123-1212331033231233"></a>

Type: `"single"`. Computed.

Specifies different cookie related config parameters for authentication.

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

<a id="canonical-1133011022120320-2200113031111311-0121022103333123-2210312032012212-1212133311221021-0230010220212202-3313023113211122-1020101130212131"></a>

## Direct properties — cookie_params / 033220211200 / 3

- [auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-1323112121023303-0203300100112323-0303321331222110-1321011120113201-0323331013013303-0110101011200203-0321033133320310-0300000013130122): complete subsection reference.

<a id="canonical-0003102311330132-0313132122120030-1323302011113011-1013001013212110-0133223320313310-3130220332033010-2030302330100302-2003201320331323"></a>

<a id="canonical-0130132302100233-0102221112010113-2030223201221030-1202223203303133-2011102311020132-0020122210032303-2022112001100100-2302030031000303"></a>

## cookie_expiry property — cookie_params / 033220211200 / 4

Type: `"number"`. Computed.

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client-side after which client will not
be setting the cookie as part of the request.

Upstream description:

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client-side after which client will not
be setting the cookie as part of the request. Default cookie expiry is 3600 seconds.

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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0012120012110132-3231303313212201-0212223002222301-1022221230122030-0001001312122302-0202332210213232-0120023011200032-3310230101021323"></a>

<a id="canonical-3002302122123033-2131121211313323-0232123221301300-2233100302233002-0202222333333322-2330232121200013-1233102320022111-0203003303213222"></a>

## cookie_refresh_interval property — cookie_params / 033220211200 / 5

Type: `"number"`. Computed.

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session..

Upstream description:

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session expiry. Default refresh interval is 3000 seconds.

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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

- [kms_key_hmac](data-sources--virtual_host--reference--group-001.md#canonical-3310021112030201-0031131211202322-2133131213000101-1010122133011302-2132212001303011-0213113310233213-2322012011303002-3120321021213033): complete subsection reference.

<a id="canonical-3323201110032133-3233133013200122-2101112001021030-2020012121100001-0012333323130311-0123100121012130-3200213313112123-0312121011131323"></a>

<a id="canonical-2023113212103223-2013023302113100-3020030203331301-3200332103110132-0003201330310031-3101012211221122-0210201300032210-2132130331322032"></a>

## session_expiry property — cookie_params / 033220211200 / 6

Type: `"number"`. Computed.

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Upstream description:

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

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
    "ves.io.schema.rules.uint32.lte": "1296000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1296000"
  }
}
```

<a id="canonical-2022020023122303-2101120233020332-3220221220002223-0233302112310330-1003231110222011-2010010200202230-1032133012320301-2331312123230213"></a>

## Next pages — cookie_params / 033220211200 / 7

- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-1323112121023303-0203300100112323-0303321331222110-1321011120113201-0323331013013303-0110101011200203-0321033133320310-0300000013130122)
- [authentication.cookie_params.kms_key_hmac](data-sources--virtual_host--reference--group-001.md#canonical-3310021112030201-0031131211202322-2133131213000101-1010122133011302-2132212001303011-0213113310233213-2322012011303002-3120321021213033)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1323112121023303-0203300100112323-0303321331222110-1321011120113201-0323331013013303-0110101011200203-0321033133320310-0300000013130122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131111121313013-2120031331303320-0231112113130212-2112323231302030-3333002103133111-3201223203103302-2200300313130021-1020222102323003"></a>

## authentication.cookie_params.auth_hmac — auth_hmac / 331120003331 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-2303003000301221-3111010311223023-1111031330330013-2131331132220311-2210233002000103-0130120011102211-3303031132200120-1023232302002330)
- authentication.cookie_params.auth_hmac

<a id="canonical-2232331311211131-1203133120300113-2133103133131233-2023313222112101-0123130120210311-1200011130311133-3232330222120303-2303003230133233"></a>

Type: `"single"`. Computed.

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Upstream description:

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0032032312130302-3003222202112301-3003313022230300-1310202021112310-3313230010302332-0233003002211210-0102202233230200-3101133130203323"></a>

## Direct properties — auth_hmac / 331120003331 / 3

- [prim_key](data-sources--virtual_host--reference--group-001.md#canonical-2112130022321310-2221023202330200-0221111013131012-2110110031303200-2230323113101121-2203212303300102-3022112320110021-2003021013101111): complete subsection reference.

<a id="canonical-3301012231300102-1221303020220301-2112330313331210-3313032222311230-1231033211202313-3221000301201330-2211213132320232-3313032202301110"></a>

<a id="canonical-1223202322100301-1103333220212120-1011010212333233-3320002231101223-1211030113032330-1010003012001310-1322003100321301-1301121202102210"></a>

## prim_key_expiry property — auth_hmac / 331120003331 / 4

Type: `"string"`. Computed.

HMAC Primary Key Expiry. Primary HMAC Key Expiry time.

Upstream description:

Primary HMAC Key Expiry time.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [sec_key](data-sources--virtual_host--reference--group-001.md#canonical-2221113110210023-3230111101032030-2031121221233332-2133132102103132-2231321211013001-2020200331032202-0220020122311222-1230212231302311): complete subsection reference.

<a id="canonical-3111130303030311-0000331211223202-2022101222022032-2310120103001220-3330211013103111-1211100233110020-1332102031231230-1301301213311223"></a>

<a id="canonical-0002213003012102-2111003331223031-2310112002313031-3210020111331321-3232211232121333-0103010313020203-1123323321133023-3031000002010303"></a>

## sec_key_expiry property — auth_hmac / 331120003331 / 5

Type: `"string"`. Computed.

HMAC Secondary Key Expiry. Secondary HMAC Key Expiry time.

Upstream description:

Secondary HMAC Key Expiry time.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1222132112222331-3110100320212031-2012223110131303-0200203310000230-0000001333003023-1222220022303312-2221322022223233-3301230002003111"></a>

## Next pages — auth_hmac / 331120003331 / 6

- [authentication.cookie_params.auth_hmac.prim_key](data-sources--virtual_host--reference--group-001.md#canonical-2112130022321310-2221023202330200-0221111013131012-2110110031303200-2230323113101121-2203212303300102-3022112320110021-2003021013101111)
- [authentication.cookie_params.auth_hmac.sec_key](data-sources--virtual_host--reference--group-001.md#canonical-2221113110210023-3230111101032030-2031121221233332-2133132102103132-2231321211013001-2020200331032202-0220020122311222-1230212231302311)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-2303003000301221-3111010311223023-1111031330330013-2131331132220311-2210233002000103-0130120011102211-3303031132200120-1023232302002330)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2112130022321310-2221023202330200-0221111013131012-2110110031303200-2230323113101121-2203212303300102-3022112320110021-2003021013101111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313311332001130-1313012021120103-1031331011233301-2003003023100021-3032332030223030-0331112023212112-3131032220203033-1231332002123110"></a>

## authentication.cookie_params.auth_hmac.prim_key — prim_key / 333000000230 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-2303003000301221-3111010311223023-1111031330330013-2131331132220311-2210233002000103-0130120011102211-3303031132200120-1023232302002330)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-1323112121023303-0203300100112323-0303321331222110-1321011120113201-0323331013013303-0110101011200203-0321033133320310-0300000013130122)
- authentication.cookie_params.auth_hmac.prim_key

<a id="canonical-1323313131100023-3010310200121023-3233030312213031-1302212130220110-1112200130331212-3211003232222120-1133022013003323-0012120133103210"></a>

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

<a id="canonical-0222331032321131-2203113032030301-2132103112001200-3111030133013212-1110210301122313-0010211323331201-1332313023300111-2123122302300000"></a>

## Direct properties — prim_key / 333000000230 / 3

- [blindfold_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-1211012020303022-1311113103133030-2011233330332013-1313100132233330-2023102212123003-3021321133312112-2001023332200021-1323331120103320): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-1203133133100120-0323231121322012-2020323221022002-3011323000122313-2310133330031202-3221113332221201-0102010212320133-2313200101110323): complete subsection reference.

<a id="canonical-3300020302012023-3203030020123011-0202011103300110-1123321123130311-3313333002331211-0323022003230312-0023000323330103-1133122132010322"></a>

## Next pages — prim_key / 333000000230 / 4

- [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-1211012020303022-1311113103133030-2011233330332013-1313100132233330-2023102212123003-3021321133312112-2001023332200021-1323331120103320)
- [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-1203133133100120-0323231121322012-2020323221022002-3011323000122313-2310133330031202-3221113332221201-0102010212320133-2313200101110323)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-1323112121023303-0203300100112323-0303321331222110-1321011120113201-0323331013013303-0110101011200203-0321033133320310-0300000013130122)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1211012020303022-1311113103133030-2011233330332013-1313100132233330-2023102212123003-3021321133312112-2001023332200021-1323331120103320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213031210000021-2312200322032020-1122033120032123-1331020300303213-1011322232022323-3232321323203230-1113230200211001-2233211322123022"></a>

## authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info — blindfold_secret_info / 333310312200 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-2303003000301221-3111010311223023-1111031330330013-2131331132220311-2210233002000103-0130120011102211-3303031132200120-1023232302002330)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-1323112121023303-0203300100112323-0303321331222110-1321011120113201-0323331013013303-0110101011200203-0321033133320310-0300000013130122)
- [authentication.cookie_params.auth_hmac.prim_key](data-sources--virtual_host--reference--group-001.md#canonical-2112130022321310-2221023202330200-0221111013131012-2110110031303200-2230323113101121-2203212303300102-3022112320110021-2003021013101111)
- authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info

<a id="canonical-3000031001311211-3010133210120033-1121122220002101-0213003120100000-3010132232132123-3303101011110333-3033030313201102-3033301100001330"></a>

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

<a id="canonical-2101102112221101-1132331011300222-0211310333222221-2100210331333321-2311000000010223-0031330133103111-0211111222023311-3311201011031320"></a>

## Direct properties — blindfold_secret_info / 333310312200 / 3

<a id="canonical-1030113313133101-2211312331032313-3021133120333322-1313020001301301-3233313011322111-0323113300033313-2033333332220300-1211110003233100"></a>

<a id="canonical-2121300333021202-1020303002120102-3102123023121012-3323120223110113-1320303022222022-2130223300203002-0013311131013000-0001203131222021"></a>

## decryption_provider property — blindfold_secret_info / 333310312200 / 4

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

<a id="canonical-0001202000002000-0320221121001102-3110211232210212-1120113100301230-3212101201031012-3230323022132011-3010121003311032-1310302301100301"></a>

<a id="canonical-1310332333210023-2133101223121213-3000213202323323-0013302233103232-1232022123013130-2111212333122002-0213211023223302-2300110021222030"></a>

## location property — blindfold_secret_info / 333310312200 / 5

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

<a id="canonical-3022210310033223-2321013123230313-2213201100212210-2313212320030213-0030322201121020-2332102012111231-1131321123011233-1230013011223120"></a>

<a id="canonical-1210200202021022-0120020332032313-2220020130122333-1121330113232022-1103001020223233-1200223012133033-0300233213331000-0331231303332232"></a>

## store_provider property — blindfold_secret_info / 333310312200 / 6

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

<a id="canonical-0311202312331000-3001323011212233-2321200212032232-1300233213313023-1133203000100312-0322221321232112-2032312322033032-0131102101123231"></a>

## Next pages — blindfold_secret_info / 333310312200 / 7

- [authentication.cookie_params.auth_hmac.prim_key](data-sources--virtual_host--reference--group-001.md#canonical-2112130022321310-2221023202330200-0221111013131012-2110110031303200-2230323113101121-2203212303300102-3022112320110021-2003021013101111)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1203133133100120-0323231121322012-2020323221022002-3011323000122313-2310133330031202-3221113332221201-0102010212320133-2313200101110323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203210013312301-1312212112202013-0020133102331203-2121011122332200-0033331303321023-2113233223232222-0000000331313112-1020112322200202"></a>

## authentication.cookie_params.auth_hmac.prim_key.clear_secret_info — clear_secret_info / 210203223102 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-2303003000301221-3111010311223023-1111031330330013-2131331132220311-2210233002000103-0130120011102211-3303031132200120-1023232302002330)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-1323112121023303-0203300100112323-0303321331222110-1321011120113201-0323331013013303-0110101011200203-0321033133320310-0300000013130122)
- [authentication.cookie_params.auth_hmac.prim_key](data-sources--virtual_host--reference--group-001.md#canonical-2112130022321310-2221023202330200-0221111013131012-2110110031303200-2230323113101121-2203212303300102-3022112320110021-2003021013101111)
- authentication.cookie_params.auth_hmac.prim_key.clear_secret_info

<a id="canonical-0021031111120101-3002023003230033-2122320032211313-0223032121101202-0232203123211203-3112123213222123-1103002322112002-0333031021133211"></a>

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

<a id="canonical-0233330002200211-0231331101001032-0103021023122231-2303102123203202-0203100230212312-2122223032033131-2300032003300312-0121312112212303"></a>

## Direct properties — clear_secret_info / 210203223102 / 3

<a id="canonical-0311101021101013-0003231132021013-1231131311230210-0301321030021002-3000332021022333-2111313112023023-0131332033301132-3223223203002123"></a>

<a id="canonical-3233101202220013-0232303202121031-1130121202321020-3103230231032202-0311223033231212-0000332221232233-0222230210020301-0321032000102301"></a>

## provider_ref property — clear_secret_info / 210203223102 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3010020312320220-1221323300010322-3323202203113311-3233112022003101-2303130122223200-1311323101122220-0223330222111212-1003130211022020"></a>

<a id="canonical-2231020310211201-2113023231022102-0202303013331320-0200033310101330-3302132102230302-0031312102332223-1302100001301310-0320210322320311"></a>

## URL property — clear_secret_info / 210203223102 / 5

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

<a id="canonical-1200033030220001-3101322211221021-1202323022112321-3311300113111332-2120213310011222-1321122321012131-0201203220022012-0122202021032320"></a>

## Next pages — clear_secret_info / 210203223102 / 6

- [authentication.cookie_params.auth_hmac.prim_key](data-sources--virtual_host--reference--group-001.md#canonical-2112130022321310-2221023202330200-0221111013131012-2110110031303200-2230323113101121-2203212303300102-3022112320110021-2003021013101111)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2221113110210023-3230111101032030-2031121221233332-2133132102103132-2231321211013001-2020200331032202-0220020122311222-1230212231302311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202131212200203-0312200321111011-0203211200322010-0113112113003030-2330133222121112-3101123232330130-3212203132030323-2103333103120121"></a>

## authentication.cookie_params.auth_hmac.sec_key — sec_key / 221132012210 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-2303003000301221-3111010311223023-1111031330330013-2131331132220311-2210233002000103-0130120011102211-3303031132200120-1023232302002330)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-1323112121023303-0203300100112323-0303321331222110-1321011120113201-0323331013013303-0110101011200203-0321033133320310-0300000013130122)
- authentication.cookie_params.auth_hmac.sec_key

<a id="canonical-2221210233223010-2031120323303311-0020031110022022-3320303003213211-0121222302322221-1232022212320100-3011333231212031-0001110021023313"></a>

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

<a id="canonical-1103302213021002-1330012112003103-3330231003302323-2133200310312110-3221330123112012-3103100102023010-0221323020033122-3011131311231303"></a>

## Direct properties — sec_key / 221132012210 / 3

- [blindfold_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-1103321311031313-3111320123012301-2113103211020002-0312110111333130-1101211011321332-2000110123011300-2221031111131020-1113231110301232): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-1011332223002212-0333332122130102-3222031111303010-3321222220232020-3131130031123222-0111010303322321-2210031022220132-1123111001301221): complete subsection reference.

<a id="canonical-3123122333013332-0002003122203220-0012011321003211-2202211222331020-3132323210323311-3220001011012002-3010300220120110-0020233330200301"></a>

## Next pages — sec_key / 221132012210 / 4

- [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-1103321311031313-3111320123012301-2113103211020002-0312110111333130-1101211011321332-2000110123011300-2221031111131020-1113231110301232)
- [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-1011332223002212-0333332122130102-3222031111303010-3321222220232020-3131130031123222-0111010303322321-2210031022220132-1123111001301221)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-1323112121023303-0203300100112323-0303321331222110-1321011120113201-0323331013013303-0110101011200203-0321033133320310-0300000013130122)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1103321311031313-3111320123012301-2113103211020002-0312110111333130-1101211011321332-2000110123011300-2221031111131020-1113231110301232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032303312331311-2322001010200311-1232123233133131-1231132031003312-1332103333302120-3232320333203032-1322322211333221-2010112220330032"></a>

## authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info — blindfold_secret_info / 302013000031 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-2303003000301221-3111010311223023-1111031330330013-2131331132220311-2210233002000103-0130120011102211-3303031132200120-1023232302002330)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-1323112121023303-0203300100112323-0303321331222110-1321011120113201-0323331013013303-0110101011200203-0321033133320310-0300000013130122)
- [authentication.cookie_params.auth_hmac.sec_key](data-sources--virtual_host--reference--group-001.md#canonical-2221113110210023-3230111101032030-2031121221233332-2133132102103132-2231321211013001-2020200331032202-0220020122311222-1230212231302311)
- authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info

<a id="canonical-0030222023020032-3120110332102222-2211110100210210-0213333112313103-3303222230322222-1102000131303333-1323321001103001-1123311320113103"></a>

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

<a id="canonical-0322223322200010-3233232010320131-3233130030130320-0312002103320003-0101000132301122-3012101333331330-1121332033000001-2303031333013212"></a>

## Direct properties — blindfold_secret_info / 302013000031 / 3

<a id="canonical-0201223103132200-1113031222322302-1312212022213211-2103133330303300-0313320032212332-3131232231333230-3122202311303032-0232111111132213"></a>

<a id="canonical-3132023323323022-1332131022203210-2032031201102311-1110330333012312-2301302233322121-0231111131221020-0113121012020202-2002133301333121"></a>

## decryption_provider property — blindfold_secret_info / 302013000031 / 4

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

<a id="canonical-3013210100030223-1202130021010013-3211321100020222-0111033111320220-0232101033123100-2302033210013221-1011233020000331-1120300331222301"></a>

<a id="canonical-0021133010001112-2031100233310233-0202221231201022-2011301102310200-1123032320111330-0203220300101231-3131310030312333-1101031331021113"></a>

## location property — blindfold_secret_info / 302013000031 / 5

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

<a id="canonical-2302231320122132-1103202002301213-2333122010320110-0023221222023210-3003332202131210-3030130310213303-2321222111231230-1110311300222122"></a>

<a id="canonical-2112002330221111-1313200110300000-2010002010103202-0113202312311313-2103123111102321-3330110103000000-3000303111110221-3002020113113023"></a>

## store_provider property — blindfold_secret_info / 302013000031 / 6

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

<a id="canonical-3333230111001221-0022033311223232-2022133132003023-1221313010200023-0013310320201023-2231131131023222-3311323103320000-3320111332131332"></a>

## Next pages — blindfold_secret_info / 302013000031 / 7

- [authentication.cookie_params.auth_hmac.sec_key](data-sources--virtual_host--reference--group-001.md#canonical-2221113110210023-3230111101032030-2031121221233332-2133132102103132-2231321211013001-2020200331032202-0220020122311222-1230212231302311)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1011332223002212-0333332122130102-3222031111303010-3321222220232020-3131130031123222-0111010303322321-2210031022220132-1123111001301221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301300120032112-1202021113103333-3220333232312231-3223203330101322-0301000303022201-1320312303221021-0303302103322003-3100112022013001"></a>

## authentication.cookie_params.auth_hmac.sec_key.clear_secret_info — clear_secret_info / 223032222123 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-2303003000301221-3111010311223023-1111031330330013-2131331132220311-2210233002000103-0130120011102211-3303031132200120-1023232302002330)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-1323112121023303-0203300100112323-0303321331222110-1321011120113201-0323331013013303-0110101011200203-0321033133320310-0300000013130122)
- [authentication.cookie_params.auth_hmac.sec_key](data-sources--virtual_host--reference--group-001.md#canonical-2221113110210023-3230111101032030-2031121221233332-2133132102103132-2231321211013001-2020200331032202-0220020122311222-1230212231302311)
- authentication.cookie_params.auth_hmac.sec_key.clear_secret_info

<a id="canonical-1331201132303122-1220312323210231-0310330122030233-3010201020122300-1202323102100100-0012013131022033-0010221113233123-2323102302131002"></a>

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

<a id="canonical-2302020100210002-2112330330132203-3100200121000322-2320103202002131-0333202301123311-0131022122130330-2021102200302123-3312032313112202"></a>

## Direct properties — clear_secret_info / 223032222123 / 3

<a id="canonical-3310302021212331-3031031303323330-3002212011131123-2301021013322313-1200122301032101-1113312220221022-0302010113203113-1030023030131321"></a>

<a id="canonical-3112203321212301-3103303233330212-2112303000320221-2120313213110321-3231301031010303-3201002300000030-2012203132323300-2300003021001121"></a>

## provider_ref property — clear_secret_info / 223032222123 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0320233330233232-1333300323122202-3210032212121233-1010331031021100-2232302103100323-1130100023023221-2231221001031323-3311001101113301"></a>

<a id="canonical-1202203000003323-2312010222020212-3022022010332202-3013132123220113-2232300331133102-2231333012303301-0312233003003032-3333202302033301"></a>

## URL property — clear_secret_info / 223032222123 / 5

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

<a id="canonical-2320003303131032-3000133001032331-3113330131133233-2302302332313333-3000311120323220-0223002031101133-3133231213022221-2303102200331021"></a>

## Next pages — clear_secret_info / 223032222123 / 6

- [authentication.cookie_params.auth_hmac.sec_key](data-sources--virtual_host--reference--group-001.md#canonical-2221113110210023-3230111101032030-2031121221233332-2133132102103132-2231321211013001-2020200331032202-0220020122311222-1230212231302311)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3310021112030201-0031131211202322-2133131213000101-1010122133011302-2132212001303011-0213113310233213-2322012011303002-3120321021213033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321031310232202-2013302333303002-1301001120203031-1231020010230101-0113013322223323-2033002310011333-2121121030233221-0300022123332000"></a>

## authentication.cookie_params.kms_key_hmac — kms_key_hmac / 303213120300 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-2303003000301221-3111010311223023-1111031330330013-2131331132220311-2210233002000103-0130120011102211-3303031132200120-1023232302002330)
- authentication.cookie_params.kms_key_hmac

<a id="canonical-0023330230021020-0220320301213000-2133120211033030-1330302000221101-0000200020123303-0020112121033310-0312113022323110-3231331313033201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for kms key hmac.

Upstream description:

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

<a id="canonical-1311311012222311-2211012330302103-0110232010201332-3222132020002332-2101212002311301-2331120311021130-3221003302301122-2121032032223002"></a>

## Direct properties — kms_key_hmac / 303213120300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331230210101203-1222022210302301-2212212233212300-0300103312213300-2011222131031312-3003221001333133-0011000312231232-1232310203210002"></a>

## Next pages — kms_key_hmac / 303213120300 / 4

- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-2303003000301221-3111010311223023-1111031330330013-2131331132220311-2210233002000103-0130120011102211-3303031132200120-1023232302002330)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1113230233030320-2122313120223033-2133232131033220-0123132323312120-2112121323200101-0011023123013013-0323203211103331-3113011110132312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102311333211332-0322300332113120-3022300231113101-1032031310101202-2123103102033013-3322102132323010-0123220311322200-0003232001033010"></a>

## authentication.redirect_dynamic — redirect_dynamic / 213333131301 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- authentication.redirect_dynamic

<a id="canonical-0203111121110120-0310232303132311-3012213303010111-1112012301203313-1330100301302313-1032130022011020-1001000120001220-1023030311203201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for redirect dynamic.

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

<a id="canonical-2033232020110230-1122333200120222-3102223002101130-3123313031022223-3203031122233030-1123201200122101-3330202033332211-2321312330011021"></a>

## Direct properties — redirect_dynamic / 213333131301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322112211022001-2332313203103231-3012022200021202-1012202110133012-1120301012012123-2012031221213220-2003201032103311-3301322022102030"></a>

## Next pages — redirect_dynamic / 213333131301 / 4

- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0221300013323210-3210001133323033-2011102333111333-1231031321131010-0000100232231231-2230121131331331-1300100003222313-1133220102301101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233123231320200-0120201022013233-3232022120132122-3321331021223001-2300301331010302-1131122322131303-0022011133102203-1203130012313302"></a>

## authentication.use_auth_object_config — use_auth_object_config / 230130202231 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- authentication.use_auth_object_config

<a id="canonical-3210031121102023-1030002222002333-1220122030001210-1221212223232231-0102131303200000-3033112230303022-2313220122111000-0323301323103211"></a>

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

<a id="canonical-1323230103022121-0303030200331313-2302003003330202-2131101033222003-1002321223212203-0323103031012321-3031030032230120-2300010133123303"></a>

## Direct properties — use_auth_object_config / 230130202231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212301210000012-1120210030310233-0131033221201331-2023010320001300-3231130223331203-0213330213232201-0003211313313000-2301023123323231"></a>

## Next pages — use_auth_object_config / 230130202231 / 4

- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1213301223231000-1113301202120100-0322200112212300-3233130103220223-3123033012131112-0321102133021131-3201032002203320-3120102012111033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322100012210100-3332332310333223-3120333000023120-2102111322131123-3101223323320223-2030020023213001-2102330211022110-1233232321010102"></a>

## buffer_policy — buffer_policy / 300213133121 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- buffer_policy

<a id="canonical-1113013312331001-3120013012321211-3201312010013030-3223213111323031-3010323000321112-0003300310113023-2231303133033022-1212033301322100"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3021021303013303-2303231131330331-0213222013213011-2203332323110001-2123303113123301-2330321200120302-0120223132313123-3213232113321332"></a>

## Direct properties — buffer_policy / 300213133121 / 3

<a id="canonical-2230303132301303-2312003132031010-2033023210021112-0310223003331100-2233202323312113-2303300121200203-1122003302131031-1333200122002112"></a>

<a id="canonical-1233301121123112-2200210311322023-3333120002110300-0222212112133131-1033133011202311-1312000300321102-3120000130230022-0221202333000233"></a>

## disabled property — buffer_policy / 300213133121 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-2232003130112123-0003030132220332-0213212130032112-2222133321201230-2203211023033210-0213023033300032-2032301333021201-0310321032100313"></a>

<a id="canonical-1021020120233213-1010232000223231-2222010333103231-0110212300030122-1323303133131322-1301103110121000-0030230313000022-2013003013321200"></a>

## max_request_bytes property — buffer_policy / 300213133121 / 5

Type: `"number"`. Computed.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Upstream description:

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
