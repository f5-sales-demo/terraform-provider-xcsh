---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-2322323222020011-0201203223111000-0013011001112001-2013202130102112-0020122321110000-0322230202101022-3323031132113121-1332231330120032"></a>

## Next pages — buffer_policy / 300213133121 / 6

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3030320121132033-0113101012333213-2333033121300031-3102221202303210-3222022210330032-0103332033311333-0233333333103003-0313010030101011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200123123002102-0101133303213301-0210302020213322-3100011331220220-1001010300323312-2321011010211132-0120121233332111-3323310312132022"></a>

## captcha_challenge — captcha_challenge / 233330331131 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- captcha_challenge

<a id="canonical-0323333321111013-0223212113102021-0320213033023313-3002301101200322-2122100233212230-0130313033323212-2101021122013300-0032223213002013"></a>

Type: `"single"`. Computed.

\[OneOf: captcha\_challenge, js\_challenge, no\_challenge; Default: no\_challenge\] Enables
loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With
this feature enabled, only clients that pass the captcha challenge will be allowed to complete the
HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect..

Upstream description:

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

OneOf alternatives in this subsection:

- [captcha_challenge](data-sources--virtual_host--reference--group-002.md#canonical-0323333321111013-0223212113102021-0320213033023313-3002301101200322-2122100233212230-0130313033323212-2101021122013300-0032223213002013)
- [js_challenge](data-sources--virtual_host--reference--group-002.md#canonical-1233322231003110-0102032003010110-2120003213121113-1012123200113212-2002113320322202-0133131333003033-3333000122301333-0101331330331000)
- [no_challenge](data-sources--virtual_host--reference--group-002.md#canonical-0333002331231113-1021121120232222-1021202203033121-0100301010000210-0311202222011232-1313320122120233-3202220223102123-3332002100110330)

Select alternatives according to the provider validators above.

<a id="canonical-3333010131331312-3103203131312233-3102300333110122-2113013220203212-0223203223310110-1101231311331233-0230231132321332-2302133200311110"></a>

## Direct properties — captcha_challenge / 233330331131 / 3

<a id="canonical-2313032031202002-2323131312322312-2012100013033132-0213022300212313-1011321123333012-0020300030213123-3101311313011100-2031113312310121"></a>

<a id="canonical-1013210301200212-1223223132100112-1123030020220330-2331321201330202-2333032300132130-1212223103000010-3322230301332123-1013123132133110"></a>

## cookie_expiry property — captcha_challenge / 233330331131 / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1023303000202202-2310101312302113-1011223101331311-3010012200021300-3013231231232321-0113313113231313-0012331330211003-2203131202210321"></a>

<a id="canonical-3303230231232321-2132232220020201-3201233201113323-0111320202000200-0020023023332132-0220312031220201-2022313202230211-2133121013111120"></a>

## custom_page property — captcha_challenge / 233330331131 / 5

Type: `"string"`. Computed.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format.

Upstream description:

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2310102211330032-0001033110322312-3110210200330033-1120100200033312-2230003310300120-0332030221003200-3023000221003231-3100322210031223"></a>

## Next pages — captcha_challenge / 233330331131 / 6

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1311333121020231-2023022200000300-2230331133331031-2300321203030302-1000030032223302-0132000331210013-0110211201102210-0011003111130011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302130130222213-3301231323231330-1010112321001003-1323313231100321-2130000101122011-2301101202220023-2231031322111313-0031112223001312"></a>

## coalescing_options — coalescing_options / 130100020323 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- coalescing_options

<a id="canonical-3003023221300230-2101220002010101-2301202002003033-0013123000032321-0331000322101331-2323033312000302-0001110123233331-2133220013333201"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

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

<a id="canonical-1233331302200033-1031121322122132-3111011200300102-0023011133021013-3011202100233322-3232131032233011-1303310312111201-2201010012110121"></a>

## Direct properties — coalescing_options / 130100020323 / 3

- [default_coalescing](data-sources--virtual_host--reference--group-002.md#canonical-0010203330110110-0201312323332103-2232013332220003-3221103010000312-1321332213331300-3110323223033130-1302200201011301-2033122320032020): complete subsection reference.

- [strict_coalescing](data-sources--virtual_host--reference--group-002.md#canonical-1033223010220211-3000021102020113-3003031222230013-2132001300002312-0030131223331001-0211113233311210-1011102112032202-0323301122011323): complete subsection reference.

<a id="canonical-2330200231320122-0010211203313112-0030313110121213-1033210202012330-3311310000233312-3101322130311231-1303322321031003-3023033220030321"></a>

## Next pages — coalescing_options / 130100020323 / 4

- [coalescing_options.default_coalescing](data-sources--virtual_host--reference--group-002.md#canonical-0010203330110110-0201312323332103-2232013332220003-3221103010000312-1321332213331300-3110323223033130-1302200201011301-2033122320032020)
- [coalescing_options.strict_coalescing](data-sources--virtual_host--reference--group-002.md#canonical-1033223010220211-3000021102020113-3003031222230013-2132001300002312-0030131223331001-0211113233311210-1011102112032202-0323301122011323)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0010203330110110-0201312323332103-2232013332220003-3221103010000312-1321332213331300-3110323223033130-1302200201011301-2033122320032020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320332021222030-1110232023331020-0311121233111231-0202012112200102-1003321310122313-2311120200220333-2313322330313103-2002021213130313"></a>

## coalescing_options.default_coalescing — default_coalescing / 102331213130 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [coalescing_options](data-sources--virtual_host--reference--group-002.md#canonical-1311333121020231-2023022200000300-2230331133331031-2300321203030302-1000030032223302-0132000331210013-0110211201102210-0011003111130011)
- coalescing_options.default_coalescing

<a id="canonical-0332011001302102-2220230311100233-1122210310102113-0012003211300111-1232222231321010-3302112321222210-0023112101131001-0203210131211001"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-2213000200101010-2022210302022223-2302010033012220-2103103302312120-0212210111003203-2320222022113311-0103103103012130-3112320231332020"></a>

## Direct properties — default_coalescing / 102331213130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131211300101210-0013220221020131-2122001222331330-1211203120021200-2110301300221131-0011020302222111-2110313221312120-2321111230321103"></a>

## Next pages — default_coalescing / 102331213130 / 4

- [coalescing_options](data-sources--virtual_host--reference--group-002.md#canonical-1311333121020231-2023022200000300-2230331133331031-2300321203030302-1000030032223302-0132000331210013-0110211201102210-0011003111130011)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1033223010220211-3000021102020113-3003031222230013-2132001300002312-0030131223331001-0211113233311210-1011102112032202-0323301122011323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023203021211211-1032110011223131-3100010033310301-0112101030113200-2030331012330222-2012202032303022-2330110210033232-1030102003033211"></a>

## coalescing_options.strict_coalescing — strict_coalescing / 122130110020 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [coalescing_options](data-sources--virtual_host--reference--group-002.md#canonical-1311333121020231-2023022200000300-2230331133331031-2300321203030302-1000030032223302-0132000331210013-0110211201102210-0011003111130011)
- coalescing_options.strict_coalescing

<a id="canonical-3321032220211300-0111031200213322-2313223213103101-1032101010301202-3002220201113200-2023000233010323-0302101333021132-0221330300233232"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-3301313110023001-3010110111222233-0232012021030011-1031100101101222-2322302133322211-1131012230101122-2010302230110130-1203321231102131"></a>

## Direct properties — strict_coalescing / 122130110020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212332013300122-2102333102111112-2131130212102121-0323211132231101-2331223333101333-1203203020013120-1100010032031311-1002211232131231"></a>

## Next pages — strict_coalescing / 122130110020 / 4

- [coalescing_options](data-sources--virtual_host--reference--group-002.md#canonical-1311333121020231-2023022200000300-2230331133331031-2300321203030302-1000030032223302-0132000331210013-0110211201102210-0011003111130011)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0113113212120330-2002000130223321-2000320201233130-1100212010220012-2220231112033113-3100133020203133-3032232101303231-1021031112133031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230113130020002-1013030230221032-3230002030201100-3212011013232121-3102123021112102-1111020312113232-0332011113130022-3302313203002201"></a>

## compression_params — compression_params / 021330332010 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- compression_params

<a id="canonical-0313011002003132-2220330320132201-0102332000101123-0212201213122331-2011031320011022-0110321130112110-3222230213210003-0110003130320121"></a>

Type: `"single"`. Computed.

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

Upstream description:

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

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1202012130220020-1012211033031230-0020321213021332-0110131130321312-3010120111321212-0323322130201321-3322222112112331-1010313202003013"></a>

## Direct properties — compression_params / 021330332010 / 3

<a id="canonical-1202210022321230-1213002133010133-0121100320230320-3331121231100010-0312230133230221-0313302333001030-0331332333022221-3222031303020203"></a>

<a id="canonical-1002233320023000-2000200220010220-2003033233331011-3322302011132212-1031101213022021-1001221202320001-0103032212320131-0232333102002322"></a>

## content_length property — compression_params / 021330332010 / 4

Type: `"number"`. Computed.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Upstream description:

Minimum response length, in bytes, which will trigger compression. The default value is 30.

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

<a id="canonical-3121131330031112-1322032200130030-0212211311133211-0103231132331330-2100221331131013-0102212223002310-1323231202330123-0200202203320101"></a>

<a id="canonical-0132222201133223-0020101030020333-1322332320011233-2023023101101201-0011002220030101-2130033121030023-2133032323131333-1113221012112010"></a>

## content_type property — compression_params / 021330332010 / 5

Type: `["list", "string"]`. Computed.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/JavaScript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Upstream description:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/JavaScript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

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

<a id="canonical-2031200203032112-3133012330210022-3100123300021302-3011131022223210-3203020021022123-3223220210213213-1112231132011221-1213331201332110"></a>

<a id="canonical-3022133300111001-2013013233031033-0102101020231113-3232030110210101-3012001332001133-1033003232202013-2322233221303212-2011130231023023"></a>

## disable_on_etag_header property — compression_params / 021330332010 / 6

Type: `"bool"`. Computed.

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

Upstream description:

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

<a id="canonical-1320100232322033-1231232022213011-3103030313201103-2303111100022201-2333213322112210-0001233223333201-0302312320202322-3113200001121313"></a>

<a id="canonical-3132112032303201-1021013000320032-1331132111221031-2101032012111112-3232320332021113-1031010010203123-2232323213112210-0312203323332212"></a>

## remove_accept_encoding_header property — compression_params / 021330332010 / 7

Type: `"bool"`. Computed.

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

Upstream description:

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

<a id="canonical-1033203202121211-0223203223122033-3102302022300332-0000200331013302-0302330222213002-0202013032112031-1120311120231102-1122223312012010"></a>

## Next pages — compression_params / 021330332010 / 8

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2210323031131321-0313202301131023-3213332203010030-0103323031221230-3302100002020020-3202222112020312-2232021222011023-1201032103301123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011333001312200-2310011023203222-2212131033101221-1331220031201323-0001122022001023-0112203013021123-3223122331223311-2120333001211300"></a>

## cors_policy — cors_policy / 331233303100 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- cors_policy

<a id="canonical-2000022133131221-1002001101221112-1210213110223313-3033230021132322-2202310223211101-1203010003300113-1323332122301323-3130222330120000"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2032312002001330-0010012113230022-0303230212223102-1331023122020300-3120323203003003-3312011101203022-0201033010123330-2210312231021132"></a>

## Direct properties — cors_policy / 331233303100 / 3

<a id="canonical-2010200221011321-0203030121013302-3023130313322200-2232231133003020-1332300123322122-1133123321010323-2103000101330133-1303113313322032"></a>

<a id="canonical-0330121021111113-0320202010310120-1201110302302202-3212020323003332-3100222223130303-1111002312102301-0102131231231133-0230221100223332"></a>

## allow_credentials property — cors_policy / 331233303100 / 4

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

<a id="canonical-3333110322201200-3200111132322112-0213032322321223-2310300310303011-3023203102213301-1001320310010221-3022002232220332-3030220000313110"></a>

<a id="canonical-2201022220313133-0010212121221002-2311121031003202-0212000221310023-1231120200010201-0322300232112233-2221022230103013-1322000201112300"></a>

## allow_headers property — cors_policy / 331233303100 / 5

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

<a id="canonical-3011003321332333-0003303003333202-0033110031131202-3121033102213331-0121221303133310-2030231133132302-2032223103323030-1011110230100323"></a>

<a id="canonical-0213020022200313-3233012103032002-1200233210102223-1322033001002122-0313001331020132-0222213132110112-0301233130322111-2202333321120221"></a>

## allow_methods property — cors_policy / 331233303100 / 6

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

<a id="canonical-1023231122322210-2011322313323333-3133330000030100-0303221213033123-3000021202232021-2122103210321021-2320000101230322-0011010101200310"></a>

<a id="canonical-1001122100331201-0223321130231001-3120320012311120-2221123210023323-0200002302000100-1303333331013303-1331220303300000-3232002312213331"></a>

## allow_origin property — cors_policy / 331233303100 / 7

Type: `["list", "string"]`. Computed.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Upstream description:

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

<a id="canonical-1030231302230221-1232012120110010-2120211300100330-3211122203310210-3022232320020232-1213032111031330-3113300230022222-0011230012111032"></a>

<a id="canonical-3222310100130320-0122112020123112-0020232303023033-1300201010223013-3221231130213001-1110231313030022-0021110300212301-3231131210131330"></a>

## allow_origin_regex property — cors_policy / 331233303100 / 8

Type: `["list", "string"]`. Computed.

Specifies regular expression patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Upstream description:

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

<a id="canonical-3210132113233112-2311110231332203-0103231031113223-2002311233031201-1033100200303322-0110202200232223-0221013132200002-1222202221100111"></a>

<a id="canonical-1131003021120101-1200113203312131-1303131012030323-0000110000000120-0033211332332203-0011301030110131-2030231012313213-3322002133011123"></a>

## disabled property — cors_policy / 331233303100 / 9

Type: `"bool"`. Computed.

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

<a id="canonical-2330012220221131-3312310133002203-3230013303010331-0013301102323103-1232021133021311-1031111203332201-3210010231330322-1102211311321110"></a>

<a id="canonical-2322013320303131-2323311003320222-1122121320302133-2103120311113332-0003213103301023-2310000123330103-3233031322112203-3010313112322101"></a>

## expose_headers property — cors_policy / 331233303100 / 10

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

<a id="canonical-3322022203032032-1330310120220232-0131333002211321-3311032213120212-3102323303202321-1322113211133310-2010201123013202-3313203021212120"></a>

<a id="canonical-0021011101211220-3201210003213100-2222102222030322-2222023103111311-3033303021020333-2213001231322130-0322222302220130-2100310132031010"></a>

## maximum_age property — cors_policy / 331233303100 / 11

Type: `"number"`. Computed.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Upstream description:

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours)

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

<a id="canonical-3021330333123003-0233210000321012-0113122313121002-3301112301323200-2121210102320002-1331203233231000-0003211000031030-1302001000331132"></a>

## Next pages — cors_policy / 331233303100 / 12

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1032321323223020-2022231132220113-3011022121331033-0131132302111121-0112122202013132-2121330303333120-2133303323113131-3121123003303012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200301131130012-3030022111303130-3001223310102221-3320221131322332-2321333120130033-3002012332321210-0311000113021320-2031333100231313"></a>

## csrf_policy — csrf_policy / 123321012120 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- csrf_policy

<a id="canonical-3031223133200111-0120310302032202-1030211123102211-1000333030031220-1303103321101333-1132110213133312-2111201110102013-2320231323100000"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3211121021200310-0122231012131031-2131321220030310-3220030021220332-1300011333132120-3033312011010122-1222002111222020-2000101021332330"></a>

## Direct properties — csrf_policy / 123321012120 / 3

- [all_load_balancer_domains](data-sources--virtual_host--reference--group-002.md#canonical-3223131022012023-2310333212121120-0113330123323000-2232100122302321-2202310021100022-3221232033313333-3121132031030012-3112220233012010): complete subsection reference.

- [custom_domain_list](data-sources--virtual_host--reference--group-002.md#canonical-3030123002112223-0132013220102110-2111213112222322-0230132133330032-1201203113221330-0111313331112302-1121111033101120-0212232121222303): complete subsection reference.

- [disabled](data-sources--virtual_host--reference--group-002.md#canonical-2000203202003022-2322223203300201-3102123003332322-0222031232103110-2033331033032230-0023103321011232-3012033203221010-3012322300301120): complete subsection reference.

<a id="canonical-0001211003011221-3322321010331232-2222223332021232-2200203300001123-0023311333222300-3102212111032311-3223313020100220-0230313013000213"></a>

## Next pages — csrf_policy / 123321012120 / 4

- [csrf_policy.all_load_balancer_domains](data-sources--virtual_host--reference--group-002.md#canonical-3223131022012023-2310333212121120-0113330123323000-2232100122302321-2202310021100022-3221232033313333-3121132031030012-3112220233012010)
- [csrf_policy.custom_domain_list](data-sources--virtual_host--reference--group-002.md#canonical-3030123002112223-0132013220102110-2111213112222322-0230132133330032-1201203113221330-0111313331112302-1121111033101120-0212232121222303)
- [csrf_policy.disabled](data-sources--virtual_host--reference--group-002.md#canonical-2000203202003022-2322223203300201-3102123003332322-0222031232103110-2033331033032230-0023103321011232-3012033203221010-3012322300301120)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3223131022012023-2310333212121120-0113330123323000-2232100122302321-2202310021100022-3221232033313333-3121132031030012-3112220233012010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121002220100113-3100201130311101-2100332321003331-2033002011330130-2103333032303012-2101223231311132-0322203322102011-2000333330012311"></a>

## csrf_policy.all_load_balancer_domains — all_load_balancer_domains / 012023233200 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-1032321323223020-2022231132220113-3011022121331033-0131132302111121-0112122202013132-2121330303333120-2133303323113131-3121123003303012)
- csrf_policy.all_load_balancer_domains

<a id="canonical-1122301000310102-2230013332210023-3030030331210002-3012113120120033-0233302122233333-3103011203121201-0123033322112323-2000331113021021"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3102120120312221-3333331312323221-0313022300101203-2311233130121030-0313132101202320-0021121321011311-1120031200213123-3211033221300112"></a>

## Direct properties — all_load_balancer_domains / 012023233200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013331100330022-3113010231000131-1230333020300321-1132330202002313-3031310320233210-2233230302131300-2301101321111111-1323101312211032"></a>

## Next pages — all_load_balancer_domains / 012023233200 / 4

- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-1032321323223020-2022231132220113-3011022121331033-0131132302111121-0112122202013132-2121330303333120-2133303323113131-3121123003303012)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3030123002112223-0132013220102110-2111213112222322-0230132133330032-1201203113221330-0111313331112302-1121111033101120-0212232121222303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231102123202232-2102201203320312-0122002000130002-2110333110201230-1120203213310331-1222133013300231-3023311213320321-2212030022323003"></a>

## csrf_policy.custom_domain_list — custom_domain_list / 132201113300 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-1032321323223020-2022231132220113-3011022121331033-0131132302111121-0112122202013132-2121330303333120-2133303323113131-3121123003303012)
- csrf_policy.custom_domain_list

<a id="canonical-1313202213002001-3233011212130111-1030101220013023-1000133331130223-0002131231033230-0302210322330312-3313232230220110-2110323021230121"></a>

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

<a id="canonical-3113323131301030-1230223221321022-3133210231313300-1300210001032011-2011133031121321-3023120202103110-0003303101332320-2101111323301322"></a>

## Direct properties — custom_domain_list / 132201113300 / 3

<a id="canonical-0113300021331300-3220133103110033-0202002232331212-3030220321203022-0113221222011331-1321133230001221-3111111200302030-0312301120222132"></a>

<a id="canonical-3230023210012222-2323200232002212-0213033301233310-2233112002022021-1123222323122033-0223133012210231-3230222320020013-3101002202333300"></a>

## domains property — custom_domain_list / 132201113300 / 4

Type: `["list", "string"]`. Computed.

List of domain names that will be matched to loadbalancer. These domains are not used for SNI match.
Wildcard names are supported in the suffix or prefix form.

Upstream description:

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

<a id="canonical-1321320133013131-3221130332230301-0101021110012210-1023002320300320-1002030101133030-3111130113023121-1332321313030202-1333200110133020"></a>

## Next pages — custom_domain_list / 132201113300 / 5

- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-1032321323223020-2022231132220113-3011022121331033-0131132302111121-0112122202013132-2121330303333120-2133303323113131-3121123003303012)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2000203202003022-2322223203300201-3102123003332322-0222031232103110-2033331033032230-0023103321011232-3012033203221010-3012322300301120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112211230233120-1323210211213123-2200132101002310-3310003003131303-3112203003011220-0222022120312221-1203330200303331-3312122010321003"></a>

## csrf_policy.disabled — disabled / 310113022203 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-1032321323223020-2022231132220113-3011022121331033-0131132302111121-0112122202013132-2121330303333120-2133303323113131-3121123003303012)
- csrf_policy.disabled

<a id="canonical-0323033223122213-0133300131232223-2202110010030003-2332331012323133-0333133123231311-1223122031332132-3022301020231310-0320133003013332"></a>

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

<a id="canonical-1210202320221300-3120020303300100-0132231212232212-0032010110030100-3132321102020032-3102203123330323-3130012030211032-1112221012031033"></a>

## Direct properties — disabled / 310113022203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231221210010001-0022232023132203-3302103210100030-3221112322111211-3213323130120001-0330332210330313-0131310130133311-1332232230030330"></a>

## Next pages — disabled / 310113022203 / 4

- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-1032321323223020-2022231132220113-3011022121331033-0131132302111121-0112122202013132-2121330303333120-2133303323113131-3121123003303012)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0223221131230033-2031313102132311-2030022100211130-0021201230212110-3110310230000301-0000220012010000-1002213210103013-0320220200333210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222320233323133-3221331111131002-0312333031221103-2311312121310211-1203311020131213-1022130320321302-3302001103122320-3302012213020021"></a>

## default_header — default_header / 110322110133 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- default_header

<a id="canonical-0131231003301103-0321312203033301-3002222033101103-1222201231020220-1232210300213310-0131212300130011-2002101030031123-0001030211230102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-2320211202002312-3233210211212033-2120212210302123-1320203012120021-1302203302210231-2303010112103120-0312210111233112-2333120302100222"></a>

## Direct properties — default_header / 110322110133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312323211012110-3233003300001112-3200303322212000-1021322222222022-2112101312212033-2032123222223103-3000010001122323-1233201021000310"></a>

## Next pages — default_header / 110322110133 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2332201222012331-2020200003133331-3003021002300031-1121110112222301-0202020120202032-2121220121022313-1310230211023310-1000212330230231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130223222210233-3221002222313213-2133132310110210-2102102213310000-3221100112020330-0111033201121303-2031220202023033-3323132023320021"></a>

## default_loadbalancer — default_loadbalancer / 001112223021 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- default_loadbalancer

<a id="canonical-0203210323302221-1200031011123211-0101202233211001-3323232022101020-0000321121303303-2333220023222121-3300321333323100-1032021333100320"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_loadbalancer, non\_default\_loadbalancer; Default: default\_loadbalancer\]
Configuration parameter for default loadbalancer.

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

OneOf alternatives in this subsection:

- [default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-0203210323302221-1200031011123211-0101202233211001-3323232022101020-0000321121303303-2333220023222121-3300321333323100-1032021333100320)
- [non_default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-1030023303220330-1233033230000310-3231300000123020-3023333313100103-3300223002022021-2032203120322203-1313212202312012-3321111231213003)

Select alternatives according to the provider validators above.

<a id="canonical-0001013311122222-0331332003220312-0222132321131203-3010130130301030-1031011112231331-3200333002020030-0231102032020222-1131331300103033"></a>

## Direct properties — default_loadbalancer / 001112223021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133013223111110-2221301223303121-2230202210001301-1033121332031212-3033022311303131-0233321313331121-1100202222200022-1232031011311132"></a>

## Next pages — default_loadbalancer / 001112223021 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3323313313333101-1022230003023322-1121321233003231-3321002002120031-3231313020123102-0332221122301230-2301200331132022-2132211330222112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333312021103300-3021123310310100-0311213100231333-0030021310212233-0201022003211211-0032001022301012-1122321203001111-2131012300003202"></a>

## disable_path_normalize — disable_path_normalize / 211120100310 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- disable_path_normalize

<a id="canonical-3301102111312121-0031222233000230-3002223322002213-3111133201103210-3210213120101312-2112223132033332-1123130311310222-0210010110330230"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_path\_normalize, enable\_path\_normalize; Default: disable\_path\_normalize\]
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

OneOf alternatives in this subsection:

- [disable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-3301102111312121-0031222233000230-3002223322002213-3111133201103210-3210213120101312-2112223132033332-1123130311310222-0210010110330230)
- [enable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-0201022300021320-1303303003222003-0131032130300103-2101300200200320-1102022211110023-3002001212313220-0232120011002133-0313322031310312)

Select alternatives according to the provider validators above.

<a id="canonical-1222233232023333-3300332101010120-0301310021000123-2312101321002230-0300001130033303-3332302300202230-3111222033102330-0230132113111020"></a>

## Direct properties — disable_path_normalize / 211120100310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203021230312023-0020021310032032-1030122210011100-1323332023203202-2113230102210102-2103001233310131-2312010130103020-3021303101330120"></a>

## Next pages — disable_path_normalize / 211120100310 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2320322030133330-2313030201332123-3133200313313203-0300032102221333-1001120222131112-1330331333122030-2312013001023021-0010002323330332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321312220100011-2310312333000123-2030121003101222-0303213020033222-3011223310311222-1030030301110330-2032101100233113-1030021122121200"></a>

## dynamic_reverse_proxy — dynamic_reverse_proxy / 020210231110 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- dynamic_reverse_proxy

<a id="canonical-3312001003220201-1021111332030323-2320100230313010-0231102023313230-3231131000121133-2230033123002221-1131102320231311-0211223000310232"></a>

Type: `"single"`. Computed.

In this mode of proxy, virtual host will resolve the destination endpoint dynamically. The dynamic
resolution is done using a predefined field in the request. This predefined field depends on the
ProxyType configured on the Virtual Host.

Upstream description:

In this mode of proxy, virtual host will resolve the destination endpoint dynamically.

The dynamic resolution is done using a predefined field in the request. This predefined field
depends on the ProxyType configured on the Virtual Host.

For HTTP traffic, i.e. With ProxyType as HTTP\_PROXY or HTTPS\_PROXY, virtual host will use the
"HOST" HTTP header from the request and perform DNS resolution to select destination endpoint.

For TCP traffic with SNI, (If the ProxyType is TCP\_PROXY\_WITH\_SNI), virtual host will perform DNS
resolution using the SNI.

The DNS resolution is performed in the virtual network specified in outside\_network\_type or
outside\_network

In both modes of operation(either using Host header or SNI), the DNS resolution could return
multiple addresses. First IPv4 address from such returned list is used as endpoint for the request.
The DNS response is cached for 60s by default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3023031220313310-0032232222202113-0200301231333032-1130330110220331-2123131012022012-0100331200302221-0022111201100133-1013213333202101"></a>

## Direct properties — dynamic_reverse_proxy / 020210231110 / 3

<a id="canonical-3211220132132303-3112133120120133-3202311011213200-1002000110020003-2320230021021132-1310002121231023-1132132011202020-0331230332012130"></a>

<a id="canonical-1230020031330311-0103103110123132-1331131312113312-1202333312023322-3202220313130112-3233213111203222-3102000023122023-1330230003322100"></a>

## connection_timeout property — dynamic_reverse_proxy / 020210231110 / 4

Type: `"number"`. Computed.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Upstream description:

The timeout for new network connections to upstream server. This is specified in milliseconds. The
default value is 2000 (2 seconds)

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

- [resolution_network](data-sources--virtual_host--reference--group-002.md#canonical-3300300230003011-1333312102330113-0310322301112003-1300032321101023-2110132310122301-1033000103331012-2033333123221333-2230311031113123): complete subsection reference.

<a id="canonical-3010130003111302-2130210213012133-2133122313331220-2021130112120331-0021231020013123-3121011111201010-2203011113302120-0333322103130113"></a>

<a id="canonical-3121210113211210-0121003010033000-2022130230131030-2003302130123230-1130220131131322-1233030130033223-2112123223321110-3031312103130232"></a>

## resolution_network_type property — dynamic_reverse_proxy / 020210231110 / 5

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2111003331333012-0102002000130221-1313311222331131-3203333123223321-1212303201102123-2220132000103222-3130113303221032-1200202002000221"></a>

<a id="canonical-2330101311120323-0021230230111323-0100223130322013-1130321130111131-2000121332120311-1322033110001330-3221033120110013-2002011031131020"></a>

## resolve_endpoint_dynamically property — dynamic_reverse_proxy / 020210231110 / 6

Type: `"bool"`. Computed.

X-example : true In this mode of proxy, virtual host will resolve the destination endpoint
dynamically. The dynamic resolution is done using a predefined field in the request. This predefined
field depends on the ProxyType configured on the Virtual Host.

Upstream description:

X-example : true In this mode of proxy, virtual host will resolve the destination endpoint
dynamically.

The dynamic resolution is done using a predefined field in the request. This predefined field
depends on the ProxyType configured on the Virtual Host.

For HTTP traffic, i.e. With ProxyType as HTTP\_PROXY or HTTPS\_PROXY, virtual host will use the
"HOST" HTTP header from the request and perform DNS resolution to select destination endpoint.

For TCP traffic with SNI, (If the ProxyType is TCP\_PROXY\_WITH\_SNI), virtual host will perform DNS
resolution using the SNI.

The DNS resolution is performed in the virtual network specified in outside\_network\_type or
outside\_network

In both modes of operation(either using Host header or SNI), the DNS resolution could return
multiple addresses. First IPv4 address from such returned list is used as endpoint for the request.
The DNS response is cached for 60s by default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0032333022133333-3033210203101001-0302002121310212-2003311231310221-3000222111303222-3230102231112001-0012212130302200-2012132302203021"></a>

## Next pages — dynamic_reverse_proxy / 020210231110 / 7

- [dynamic_reverse_proxy.resolution_network](data-sources--virtual_host--reference--group-002.md#canonical-3300300230003011-1333312102330113-0310322301112003-1300032321101023-2110132310122301-1033000103331012-2033333123221333-2230311031113123)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3300300230003011-1333312102330113-0310322301112003-1300032321101023-2110132310122301-1033000103331012-2033333123221333-2230311031113123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222101212101033-3320211102323202-0311110123120312-2121020232022031-1313100323021202-3312122132010331-3103202001101330-3221303130232303"></a>

## dynamic_reverse_proxy.resolution_network — resolution_network / 320002213320 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [dynamic_reverse_proxy](data-sources--virtual_host--reference--group-002.md#canonical-2320322030133330-2313030201332123-3133200313313203-0300032102221333-1001120222131112-1330331333122030-2312013001023021-0010002323330332)
- dynamic_reverse_proxy.resolution_network

<a id="canonical-3033013100002213-2212012332232022-0101032322103101-3203023313232302-2003311011303113-2011101113132023-1103120111033031-0211210301122103"></a>

Type: `"list"`. Computed.

Reference to virtual network where the endpoint is resolved. Reference is valid only when the
network type is VIRTUAL\_NETWORK\_PER\_SITE or VIRTUAL\_NETWORK\_GLOBAL. It is ignored for all other
network types.

Upstream description:

Reference to virtual network where the endpoint is resolved. Reference is valid only when the
network type is VIRTUAL\_NETWORK\_PER\_SITE or VIRTUAL\_NETWORK\_GLOBAL. It is ignored for all other
network types.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3300203103323013-0021213302321000-2110003313230132-3132133111331233-3331222221300220-0230131230331222-1000121330310031-0111033013101121"></a>

## Direct properties — resolution_network / 320002213320 / 3

<a id="canonical-3300130110213322-2033023121303131-0322022120321010-2031100303233130-0033131101331112-3100331302112101-0123032100001131-1103221033132010"></a>

<a id="canonical-0113302112103030-1133012232232102-1010000332322122-1201023031323021-3330000210300212-2200102111101012-3321110232120130-0111332333332132"></a>

## kind property — resolution_network / 320002213320 / 4

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

<a id="canonical-0332300122302032-2201103333100232-3201320201321221-3033020123112300-0102213011320201-2020331130323130-2331001000332332-3311230233032221"></a>

<a id="canonical-1201223212022131-1313012313132030-3033322001303333-1200200132303110-0301010232311013-3021202012210201-0102331103331132-1110300001101131"></a>

## name property — resolution_network / 320002213320 / 5

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

<a id="canonical-3030031013223202-3132003021232202-3032020020210201-1213002033221110-2222101222221300-1122002222231100-2030001102222002-3102113221212030"></a>

<a id="canonical-0330302331322231-0111133332110133-3012313031222313-1311100113012020-0320111030230103-1023202200122320-2312320223032212-1222033122233022"></a>

## namespace property — resolution_network / 320002213320 / 6

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

<a id="canonical-2221121323012333-2030311132131000-0121232223200221-2111301111231320-1202323202320031-3011112301121100-1321301331102200-2303313311123303"></a>

<a id="canonical-3023202133311300-3130030321301011-0130112331300102-1221112213221100-0203221310010212-0320210333330222-0331100113231330-2312123333202000"></a>

## tenant property — resolution_network / 320002213320 / 7

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

<a id="canonical-1132023131100102-0101122333332123-2320101021133230-2230123020222330-1022222131313332-3002001110221222-3111330122331220-3122100003321120"></a>

<a id="canonical-1220212031302013-1030230102311333-2231101220223130-3003000212030303-2102101112330302-2331113310301100-1333222333232200-2003113323031332"></a>

## uid property — resolution_network / 320002213320 / 8

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

<a id="canonical-1033033333002301-2130331230313113-0101333023212120-0303022012232112-1031233002230221-2310012330313133-0313223332320031-0221032333221211"></a>

## Next pages — resolution_network / 320002213320 / 9

- [dynamic_reverse_proxy](data-sources--virtual_host--reference--group-002.md#canonical-2320322030133330-2313030201332123-3133200313313203-0300032102221333-1001120222131112-1330331333122030-2312013001023021-0010002323330332)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0101130103132132-0220121022213231-2333332112030130-0121011300022021-3031032103223222-2102320203302331-2033233303102321-3213203032133202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002023002210322-1130001102033102-0122321310100333-3310021002022302-1013321311132120-3101031201001001-1311330231110102-2002123022323231"></a>

## enable_path_normalize — enable_path_normalize / 013301020200 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- enable_path_normalize

<a id="canonical-0201022300021320-1303303003222003-0131032130300103-2101300200200320-1102022211110023-3002001212313220-0232120011002133-0313322031310312"></a>

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

<a id="canonical-1123300111333113-0012122021010211-2203221210231300-2031011130211211-0321023231001102-2233031002031202-2332030302320112-1122023033021023"></a>

## Direct properties — enable_path_normalize / 013301020200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322033300030320-2201131221220223-0330100213021322-2320221113220013-3000112332211010-1033223001213330-2120020001113103-1120203321123013"></a>

## Next pages — enable_path_normalize / 013301020200 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322002001320020-3232021032121201-2203133033110103-1303032210212200-3013001001013010-1121232112020002-1112233002132012-3101030103032312"></a>

## http_protocol_options — http_protocol_options / 300002321200 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- http_protocol_options

<a id="canonical-2310113310130010-3202002120331033-0311323332301131-0002201322002231-2330330123123022-1120101032001030-2313321211221113-1133330133022113"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

<a id="canonical-2020230020120031-3223310321223002-1020332202021310-0213221301321010-2133120020202123-3232222122200030-0230200023230133-3133021210201133"></a>

## Direct properties — http_protocol_options / 300002321200 / 3

- [http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--virtual_host--reference--group-002.md#canonical-3210120111312310-2121102031133313-3130022301322021-2012120212332123-0113011331121132-3312213132131213-2113212202323121-3011331111311333): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--virtual_host--reference--group-002.md#canonical-1032310021331000-0032112130231321-1331130221313133-1120003323332031-0032200222220100-0213231100001120-1230112322102122-0321132010132111): complete subsection reference.

<a id="canonical-2023003100003023-0033333110330131-2213300002021322-1232021003230033-0000130003301110-3002002100010311-2223232302101310-3223331310021022"></a>

## Next pages — http_protocol_options / 300002321200 / 4

- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203)
- [http_protocol_options.http_protocol_enable_v1_v2](data-sources--virtual_host--reference--group-002.md#canonical-3210120111312310-2121102031133313-3130022301322021-2012120212332123-0113011331121132-3312213132131213-2113212202323121-3011331111311333)
- [http_protocol_options.http_protocol_enable_v2_only](data-sources--virtual_host--reference--group-002.md#canonical-1032310021331000-0032112130231321-1331130221313133-1120003323332031-0032200222220100-0213231100001120-1230112322102122-0321132010132111)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332203000222033-0111322002210233-0000101312222213-0213030133132033-3211310202212023-2112131233013231-1230323110221323-3300010223100333"></a>

## http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 102100311000 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-2120030322212322-2013100321103031-2132112213110203-2310201230132310-0312211110131133-3322120012323020-3003321110133113-0310310012222201"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1200322331130000-1310110321033322-1100001020122312-2220321323313213-3302100301332211-0132110020332013-2033333223030003-1311001111201013"></a>

## Direct properties — http_protocol_enable_v1_only / 102100311000 / 3

- [header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130323112003123-0132103001001033-3311111001110313-3311020202100020-1302112201113113-1000033210213210-0130111202301032-1221110113013200): complete subsection reference.

<a id="canonical-2110303130203231-3313232210112131-1311233011333333-3011201231231202-2020031033323012-1001130020303122-1212233022112210-1000210011013233"></a>

## Next pages — http_protocol_enable_v1_only / 102100311000 / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130323112003123-0132103001001033-3311111001110313-3311020202100020-1302112201113113-1000033210213210-0130111202301032-1221110113013200)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0130323112003123-0132103001001033-3311111001110313-3311020202100020-1302112201113113-1000033210213210-0130111202301032-1221110113013200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110031133303212-2321203112001102-2322212300131320-0310032102220201-0122310210123222-0312132001333321-1130002333013111-3131332032330213"></a>

## http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 220000130013 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0120332231310121-0113131120312310-1211221010210013-0031020333330210-1333013003123023-1113211132212222-0002003313032023-2310310221333231"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

<a id="canonical-3032310022222112-2322101023330002-0311133101303020-0120021031030220-3011032311302113-3112330321111013-2331120121321331-2201001103102111"></a>

## Direct properties — header_transformation / 220000130013 / 3

- [default_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1001302320233011-3002221201303200-0303303333210232-1101222111030322-0032202021201323-3231122113203201-0110201220101310-3011113120321300): complete subsection reference.

- [preserve_case_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130003111302320-1202220031111320-1200302232213332-0132330021113312-1000322121123001-0232210021233200-0021031133201322-3233010012213020): complete subsection reference.

- [proper_case_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1200133213310011-2013210300031221-0230122333221032-2223133010020231-1323212213101322-3010133032131223-2311201102100320-2013310023320012): complete subsection reference.

<a id="canonical-2033003133330021-2010232111233021-2301023112032022-1131322311320203-0230001310113102-3132211332110023-2100331232013033-2313321131103000"></a>

## Next pages — header_transformation / 220000130013 / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1001302320233011-3002221201303200-0303303333210232-1101222111030322-0032202021201323-3231122113203201-0110201220101310-3011113120321300)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130003111302320-1202220031111320-1200302232213332-0132330021113312-1000322121123001-0232210021233200-0021031133201322-3233010012213020)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1200133213310011-2013210300031221-0230122333221032-2223133010020231-1323212213101322-3010133032131223-2311201102100320-2013310023320012)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1001302320233011-3002221201303200-0303303333210232-1101222111030322-0032202021201323-3231122113203201-0110201220101310-3011113120321300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232022330020020-1023212003112002-0330012120203203-2320001212212120-1100321101002213-3300323321120232-2121300201003303-0032001313222033"></a>

## http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 013110312310 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130323112003123-0132103001001033-3311111001110313-3311020202100020-1302112201113113-1000033210213210-0130111202301032-1221110113013200)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-0030121220133201-0213311121113220-2100110103320320-1221333110111101-0213031001333223-3202303333332302-3010213113200331-0211003311131303"></a>

Type: `["object", {}]`. Computed.

Use the platform's current default HTTP header transformation behavior.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3232332221030320-3223331123232200-0022130032322320-0233302020132331-2132123321131022-2311032220020033-2022001232313232-2033032130302033"></a>

## Direct properties — default_header_transformation / 013110312310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030123112323013-0102223233232100-2012310323123303-0112011231221032-0031330000310010-1310130111232212-2232130020312000-2302133332302132"></a>

## Next pages — default_header_transformation / 013110312310 / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130323112003123-0132103001001033-3311111001110313-3311020202100020-1302112201113113-1000033210213210-0130111202301032-1221110113013200)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0130003111302320-1202220031111320-1200302232213332-0132330021113312-1000322121123001-0232210021233200-0021031133201322-3233010012213020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233123303100112-0010102210002321-3030221011301021-2323011013100100-2213223203213010-3301121023131101-1123212010102301-0023233002301031"></a>

## http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 003021103030 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130323112003123-0132103001001033-3311111001110313-3311020202100020-1302112201113113-1000033210213210-0130111202301032-1221110113013200)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2233331223213101-3023030222201313-0211012203010211-2131230000231202-1031111220030233-1211332231333121-1010123001311312-1222312101210023"></a>

Type: `["object", {}]`. Computed.

Preserve HTTP header-name case when upstream case must remain unchanged.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3221102202302010-0321330121013310-0333230231020010-1110132033033001-0031100103111033-1231223102220111-3220021011123323-3101130021302133"></a>

## Direct properties — preserve_case_header_transformation / 003021103030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003120122200230-2223012212200130-3011001131332303-2100101021322131-3000221111121131-2001303030031100-0001311311333300-1222200000132300"></a>

## Next pages — preserve_case_header_transformation / 003021103030 / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130323112003123-0132103001001033-3311111001110313-3311020202100020-1302112201113113-1000033210213210-0130111202301032-1221110113013200)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1200133213310011-2013210300031221-0230122333221032-2223133010020231-1323212213101322-3010133032131223-2311201102100320-2013310023320012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332300222331030-2310133301120201-3020210213320202-1232233210031331-0110102120212212-1201012301112013-1201233200322231-2022110023022130"></a>

## http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 120312000213 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130323112003123-0132103001001033-3311111001110313-3311020202100020-1302112201113113-1000033210213210-0130111202301032-1221110113013200)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1230200223331302-2002002213010322-3013103231233030-3120303222000121-2023110213010322-1303313210001222-0101332311313131-3302301200120123"></a>

Type: `["object", {}]`. Computed.

Transform HTTP header names to proper case when explicit transformation is required.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1013331111200000-3133210201102033-0110201231130010-1222222010022323-3131121200310321-2221201110312032-3031220210012230-0210213132300122"></a>

## Direct properties — proper_case_header_transformation / 120312000213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231131333311201-0021021300000032-2220302333213220-0313132113030220-0011012100323021-1201013220001020-0303333201322310-1233131133330023"></a>

## Next pages — proper_case_header_transformation / 120312000213 / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130323112003123-0132103001001033-3311111001110313-3311020202100020-1302112201113113-1000033210213210-0130111202301032-1221110113013200)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3210120111312310-2121102031133313-3130022301322021-2012120212332123-0113011331121132-3312213132131213-2113212202323121-3011331111311333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220103221002310-1103200111110013-0033221200032032-1300220131013301-3132323201131230-3130333132130013-2203323113231322-2332021012312320"></a>

## http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 210031213131 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-3303201210330112-1103310130233322-1233213311213131-3233203100332101-3013230011313001-0311010121312000-1212302030323210-1111201323321210"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

<a id="canonical-0231103033101323-1031301310102122-1122021013123023-2122223233013120-2101300220210120-3202100203333000-2300302231033112-2023122300021111"></a>

## Direct properties — http_protocol_enable_v1_v2 / 210031213131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133012212323010-0220201232201010-1200220102021122-0223233021230302-3102310032202133-2303030120010100-1121130123233012-2211223312300120"></a>

## Next pages — http_protocol_enable_v1_v2 / 210031213131 / 4

- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1032310021331000-0032112130231321-1331130221313133-1120003323332031-0032200222220100-0213231100001120-1230112322102122-0321132010132111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020101211032112-0213122201301232-2030330331302032-0130033103101201-1011331223221221-0301301113023131-1230130202110231-0011232120233333"></a>

## http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 100211223331 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-2211100030121313-1121323330100333-3211112023122100-0030201230331100-3103030202000302-1111213101120220-1001021233020211-2022321222210223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-1003300001120221-0032330310232103-3213211121030301-1012220211032132-2201310303331313-1202330213231130-1023013132210001-3231031133311013"></a>

## Direct properties — http_protocol_enable_v2_only / 100211223331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210311001122203-3003011113001301-3200330323231133-3333223113221120-0002120122320212-2022023203231103-2103133103200110-1003022013302311"></a>

## Next pages — http_protocol_enable_v2_only / 100211223331 / 4

- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0311311032210221-1213332233330222-2103321233000202-2020120320331211-3131003020011333-1211120300112103-3021112313203310-2031301312202110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002000030212230-1213310133211201-1221111002331021-3233300320021310-0020021323023212-2313112113002330-0120310102100321-0011302132020331"></a>

## js_challenge — js_challenge / 221122010113 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- js_challenge

<a id="canonical-1233322231003110-0102032003010110-2120003213121113-1012123200113212-2002113320322202-0133131333003033-3333000122301333-0101331330331000"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript. With this feature enabled, only clients that are capable of executing JavaScript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

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

<a id="canonical-3012012323020111-0301012231123103-3301333132230312-1002011100211010-1023221230111110-2313223023002332-2232111201003020-0001032203102231"></a>

## Direct properties — js_challenge / 221122010113 / 3

<a id="canonical-2313113333022330-1020033223233003-2130201233133003-1002002123131033-0213020231321222-2033002121320310-1133123211230323-1011301023100201"></a>

<a id="canonical-0303330020201110-2020123113202303-1222231313321013-3212221003310020-1131223203203302-0221233312133013-3233220130202132-0003132203021332"></a>

## cookie_expiry property — js_challenge / 221122010113 / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1302230100332332-2202313211121220-0112313320203313-1211131320323102-3310132321221203-1310313010201213-2113100130022103-0010311132323211"></a>

<a id="canonical-2320012321132322-3220323011000133-2213213131120331-2132211032320312-3123321033012132-3312311013231201-0213220201303001-2012313231103131"></a>

## custom_page property — js_challenge / 221122010113 / 5

Type: `"string"`. Computed.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format.

Upstream description:

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0103120220231031-1132002321101122-0112010200210321-3303312133213213-1302000012302123-0010032133330111-1122032303302303-1112323322232230"></a>

<a id="canonical-1100332230102221-3130312110100203-3100011022031132-2023303112321313-0103202120101130-2022321110012032-1103033202033121-2221021200000033"></a>

## js_script_delay property — js_challenge / 221122010113 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2232000232202210-2003203322021331-3300122210213013-1313233302200332-2310023323032111-2103131233332223-3010002012010133-3300023310202131"></a>

## Next pages — js_challenge / 221122010113 / 7

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1233312122313321-2011202230001010-3102133312211200-1130003122002233-2223112232011102-0131310120111220-3120020032023010-0313002202102123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232212121210333-3230003220120211-3102012322103030-0002211303023332-3030230013131021-0023303131111300-1203203213103233-1202120313011133"></a>

## no_authentication — no_authentication / 010033123001 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- no_authentication

<a id="canonical-1102100302002300-1313133201331100-3113202203002110-2131030023211011-2113223023332022-0203202123003020-0213012011103301-3232331311333110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no authentication.

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

<a id="canonical-1033210230211000-3010011022030312-0022030113312130-1012213202011221-0203231032331222-2033211123113020-1312103022320020-3202112321222210"></a>

## Direct properties — no_authentication / 010033123001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030000210100011-0020121003201133-3303032232202021-2212011113230101-0013101012030301-1323133033200303-2320101013033302-0131310212003330"></a>

## Next pages — no_authentication / 010033123001 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3110101102113030-0110100111222000-3312320302213100-0201010223313300-1213232220022000-2133222310110000-2331222120121232-2212130222311233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030113012100102-0211102000003200-3333130003022033-0212322221200200-0202100302220002-1221131130001122-2312133333032203-0211333330120133"></a>

## no_challenge — no_challenge / 013001121033 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- no_challenge

<a id="canonical-0333002331231113-1021121120232222-1021202203033121-0100301010000210-0311202222011232-1313320122120233-3202220223102123-3332002100110330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no challenge.

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

<a id="canonical-1213000121221031-2110331330111033-2212200031211102-1301321323200200-3131313220102233-0031113110303011-0123002213112230-1113302003201132"></a>

## Direct properties — no_challenge / 013001121033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311010013022313-1130330320321033-0332001111103203-1203331201033101-3012003332302013-0333121103330010-3302231321013012-1032131230202002"></a>

## Next pages — no_challenge / 013001121033 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1110112230220221-2303010133322201-1101300033010303-0103123220222322-3010312001321310-0123120221200321-1022302003222231-0020320101133303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331002213122013-1010230323112012-0030100122230312-1030311110311320-2110312000121323-1311332222210313-2030313123111131-0121201220321313"></a>

## no_request_limit_per_connection — no_request_limit_per_connection / 000002030212 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- no_request_limit_per_connection

<a id="canonical-2222231010030030-3130021202122031-3100103020130033-3121222000100032-1102001222302230-2111131103111212-2133112303230032-3021011203213203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no request limit per connection.

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

<a id="canonical-1313132323332231-0313002211031101-1323330022031012-1123123030303302-0321000301200120-3103300121100303-3012002001031211-2330322123322213"></a>

## Direct properties — no_request_limit_per_connection / 000002030212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110123302312331-3002231103321311-1312233302310310-0112213212310003-0001113012130101-0200332132102113-3233221101202212-3021003132320131"></a>

## Next pages — no_request_limit_per_connection / 000002030212 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2023112320112123-3300121231232030-1100122000331201-1003213223120203-1222330212012300-3302102223112001-3102332220100221-1121233221011031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203110111302320-1222310213013223-3332201022113300-0133333201320201-1212022102101220-3100312210301323-2122312122111102-1133033002223331"></a>

## non_default_loadbalancer — non_default_loadbalancer / 322022330330 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- non_default_loadbalancer

<a id="canonical-1030023303220330-1233033230000310-3231300000123020-3023333313100103-3300223002022021-2032203120322203-1313212202312012-3321111231213003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

<a id="canonical-3021313210020303-1322300120113030-2133102012333121-2332321022013032-0030303000323012-2302222322103313-0020123313202100-0103302233233033"></a>

## Direct properties — non_default_loadbalancer / 322022330330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231131002233111-3111220301302332-1312303023301021-0232111031333102-3011202021210103-1203231102101033-2030221002330323-0120033221002323"></a>

## Next pages — non_default_loadbalancer / 322022330330 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1202312331000133-2122223121103310-2100301313303013-2310311201301110-2210333202010230-0211311013211222-0101120011203120-3303021112101311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320003231001103-2001111323113312-3230021223213022-3221232122022021-0110203113133031-0202210113333212-2312131113033232-0023211000130002"></a>

## pass_through — pass_through / 101233230113 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- pass_through

<a id="canonical-2011230100101232-2231130122231011-0101023133113300-2212113001232301-2201101303031211-2222001202121303-1333000311103303-2303303020030321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

<a id="canonical-2132102301223012-3133311012103000-3330303221023202-1100223103303020-0330212012321302-1313330210321130-0022011333012220-0123201321123023"></a>

## Direct properties — pass_through / 101233230113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221001131020302-1022231103003202-3102220131122220-0201221323232300-0030321013123123-3132131110310030-2222131202102101-2031021132212322"></a>

## Next pages — pass_through / 101233230113 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1133310002220332-1103221000011023-3032223311312122-2321113131121222-2110333100201023-1011110232221212-1300111103011301-1303120213023022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330013312312021-0212021232033212-2211312220233022-2320211031333212-2113120300103331-2322323300201232-1030030122023123-1213310321312130"></a>

## rate_limiter_allowed_prefixes — rate_limiter_allowed_prefixes / 113130002001 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- rate_limiter_allowed_prefixes

<a id="canonical-0331330100322200-1002100120300333-2120120320132113-3002220323202030-1000321101320011-1133201200033030-1322111110211323-3011210232003110"></a>

Type: `"list"`. Computed.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Upstream description:

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-0223031322321011-3001230200123303-0200120003302311-1331201120311203-0102223001020100-1031031130223311-2113000303323013-1000033103320202"></a>

## Direct properties — rate_limiter_allowed_prefixes / 113130002001 / 3

<a id="canonical-3003022201012312-1020122303201010-1221030322123200-3032321102213223-2011202001121100-2210003123231310-3001133022303023-3030233101233031"></a>

<a id="canonical-3002131221132102-3102013001222303-3020133111222120-2112210003313102-3131111323100130-1001230202200130-3021312233223213-2103112300230010"></a>

## kind property — rate_limiter_allowed_prefixes / 113130002001 / 4

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

<a id="canonical-2300030130022113-2002103122312233-0332110112313112-3202313123320002-1230311322013002-1031321111000223-2100220012333200-3003023301123303"></a>

<a id="canonical-1030300203111022-2302121210033102-0211030033023203-2323022300233323-0131302233331310-0111132223203012-3130201322223021-0131300130330301"></a>

## name property — rate_limiter_allowed_prefixes / 113130002001 / 5

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

<a id="canonical-0123332313122333-2233322111222220-2232000033013202-0202221023100030-0320330123210302-0121323021303203-1211002111011002-3010202021201023"></a>

<a id="canonical-0023320222000003-0100312310221003-0221021031113120-2300201110122003-3101211130313333-2000210332212321-1313011013233223-2010323212013221"></a>

## namespace property — rate_limiter_allowed_prefixes / 113130002001 / 6

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

<a id="canonical-0332132321331221-0010120222322311-0232113213223032-0302323220221311-3203021021121303-2301030303122210-0201302331321202-2311133033020220"></a>

<a id="canonical-2331032200132303-0132021022102101-2122200100100231-0012333221033212-1103211333112013-0310020222103321-3221122211130021-3333103010002321"></a>

## tenant property — rate_limiter_allowed_prefixes / 113130002001 / 7

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

<a id="canonical-2003130203331213-2112231202131123-0011312113200000-3123033332122320-0200100301022012-0103211130201011-0332102312300101-0322031302120323"></a>

<a id="canonical-0200331210211331-1330113031133220-3303222202313303-0021333032222020-0020210321232323-0110312103022103-0233132121022112-0232202321130011"></a>

## uid property — rate_limiter_allowed_prefixes / 113130002001 / 8

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

<a id="canonical-2103013122120023-3121333021130102-3133220131311120-2313132101002200-0101013213010230-2133101000201001-2223000310332213-0332102010302202"></a>

## Next pages — rate_limiter_allowed_prefixes / 113130002001 / 9

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3012001313211032-0121110012120332-1211212103012331-1030201012110112-3012311230231023-0312310313121023-3122223302301012-1011130302002022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232203312321211-0001013312201130-0210233022031300-2112120130110230-3302012223333132-0213033301111300-2202030101212032-0030201223113012"></a>

## request_cookies_to_add — request_cookies_to_add / 303111032112 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- request_cookies_to_add

<a id="canonical-1313120312033223-3223302312021001-2210111031213102-0032030313222120-1111023302121322-0321022003120100-1002030012222312-3203122220311202"></a>

Type: `"list"`. Computed.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

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

<a id="canonical-3201101321100102-1320232233031022-1321001023332200-2220133000033001-0201112312211113-3232131231120131-0001302133130322-2010321003320100"></a>

## Direct properties — request_cookies_to_add / 303111032112 / 3

<a id="canonical-3210111320021210-0032132032121020-3302313300112212-0032120213000131-3303212001021130-0210100012323310-2201200303232120-3330020121311200"></a>

<a id="canonical-2012121113321111-1311213120323030-1320312231221323-1302100200021023-3121112203323102-3113010112100101-2332023330330203-1333132103123032"></a>

## name property — request_cookies_to_add / 303111032112 / 4

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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

<a id="canonical-3123011230320133-2030230232131002-2132301112010200-3000221223231033-2313112322222022-1020210231223000-2300113110033321-3213101011220211"></a>

<a id="canonical-2321210312030022-0032032322022111-3332330222003103-2323020211123320-1130230332201120-3030201310022213-0301322123131111-2221123002122030"></a>

## overwrite property — request_cookies_to_add / 303111032112 / 5

Type: `"bool"`. Computed.

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

- [secret_value](data-sources--virtual_host--reference--group-002.md#canonical-2221012012102111-2312110102322101-0103322330302302-1323220122123302-3032133202210302-3121200131100310-1322210101030202-2123321301200213): complete subsection reference.

<a id="canonical-3022111331213332-0111133010110023-1121102122000323-3311001313022000-3321032110313012-0123130321032020-2320111232303303-0231121300231003"></a>

<a id="canonical-3303223103010032-3300110022022320-0031220000330122-0213011020013311-1232121312023220-3103223023001122-2211031021222033-3020300220213210"></a>

## value property — request_cookies_to_add / 303111032112 / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

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

<a id="canonical-0011323012302013-1302221020131030-1312321232313213-3221012022120111-3302321323301310-2101323122322220-1322111003310132-2032010331123312"></a>

## Next pages — request_cookies_to_add / 303111032112 / 7

- [request_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-2221012012102111-2312110102322101-0103322330302302-1323220122123302-3032133202210302-3121200131100310-1322210101030202-2123321301200213)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2221012012102111-2312110102322101-0103322330302302-1323220122123302-3032133202210302-3121200131100310-1322210101030202-2123321301200213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100332313033011-0021311232331002-0333212203102130-2231203132101232-0212112001003300-3201122211201210-0323011020121202-1122323303102211"></a>

## request_cookies_to_add.secret_value — secret_value / 300300011212 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3012001313211032-0121110012120332-1211212103012331-1030201012110112-3012311230231023-0312310313121023-3122223302301012-1011130302002022)
- request_cookies_to_add.secret_value

<a id="canonical-0322303201132233-3331133033102320-3102210333112212-2121003031010100-2022302021021031-0203120133110221-0322311031233013-0123233230031222"></a>

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

<a id="canonical-3321333221100230-3021121033021213-1212223300122112-0312310323230331-3023003312112100-2332133322332200-0321112221110123-0123020313222033"></a>

## Direct properties — secret_value / 300300011212 / 3

- [blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-3322003220000330-2001322312113202-3022120321102120-3112112220322111-3303332001032122-2231023301212002-2120202311103032-0101303213310030): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-1021000123131222-3033303221301233-2130203001030221-2132332321310010-1013300122330020-0310222323300021-1012200331010301-1200011230231132): complete subsection reference.

<a id="canonical-3201021003320021-1311030130000233-3203312301012113-2102301131321010-1023210312123311-0300331013010303-3030332010001311-0322222000320231"></a>

## Next pages — secret_value / 300300011212 / 4

- [request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-3322003220000330-2001322312113202-3022120321102120-3112112220322111-3303332001032122-2231023301212002-2120202311103032-0101303213310030)
- [request_cookies_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-1021000123131222-3033303221301233-2130203001030221-2132332321310010-1013300122330020-0310222323300021-1012200331010301-1200011230231132)
- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3012001313211032-0121110012120332-1211212103012331-1030201012110112-3012311230231023-0312310313121023-3122223302301012-1011130302002022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3322003220000330-2001322312113202-3022120321102120-3112112220322111-3303332001032122-2231023301212002-2120202311103032-0101303213310030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330103002321031-1032323002310331-1123113123102133-2333333130003033-2302101030001133-0300100203210322-1321122103033300-0013022000103020"></a>

## request_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 301110132122 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3012001313211032-0121110012120332-1211212103012331-1030201012110112-3012311230231023-0312310313121023-3122223302301012-1011130302002022)
- [request_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-2221012012102111-2312110102322101-0103322330302302-1323220122123302-3032133202210302-3121200131100310-1322210101030202-2123321301200213)
- request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-1102332111131111-2312210102230132-2102301032003113-1222233223112201-0103032111122220-3020133233303122-0123020330232333-1122130322203001"></a>

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

<a id="canonical-2200213323131213-0230330220113120-2132230232023010-1222023332031330-1013023211312132-1031111330313210-3221100011032032-1220012013001333"></a>

## Direct properties — blindfold_secret_info / 301110132122 / 3

<a id="canonical-0011112001231020-0202022121232302-1313002012223121-3332000123100131-1013300013130302-1030223302012111-3100330301110020-3321032102301203"></a>

<a id="canonical-1122303231333222-1002002133221030-0233222023130130-0021121320203120-2002132312210300-1122130033222303-0131300233201313-3013001212010222"></a>

## decryption_provider property — blindfold_secret_info / 301110132122 / 4

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

<a id="canonical-3203200133231320-3031120000321011-2221130111333310-0200203230030211-2231210200203130-0000201320023200-2012003002032200-1212202200213233"></a>

<a id="canonical-3023231312320111-3200103233101000-2133220003102031-3120022213220113-1231303010213323-2220000010002020-1303232220110113-0221222030102333"></a>

## location property — blindfold_secret_info / 301110132122 / 5

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

<a id="canonical-0331331030323300-0132232222123003-2033300320120322-1032320321132331-1000223130031122-1221113010222033-3133002212133132-0200230301210200"></a>

<a id="canonical-3133122120130111-3021022223001202-2212230031200211-1300212301312223-1200013303210230-3121023122321132-2033200231010311-2211233100111330"></a>

## store_provider property — blindfold_secret_info / 301110132122 / 6

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

<a id="canonical-0211133311323230-3030003221000031-2000311302021121-0111213113221233-2001120132301213-3222023300030002-3302100002022031-3313111131200122"></a>

## Next pages — blindfold_secret_info / 301110132122 / 7

- [request_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-2221012012102111-2312110102322101-0103322330302302-1323220122123302-3032133202210302-3121200131100310-1322210101030202-2123321301200213)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1021000123131222-3033303221301233-2130203001030221-2132332321310010-1013300122330020-0310222323300021-1012200331010301-1200011230231132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000100111031102-0221330232002030-0322310111132001-3002233200002000-3232311302311010-2012321000130000-3323333023112122-2213232232300321"></a>

## request_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 213022331132 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3012001313211032-0121110012120332-1211212103012331-1030201012110112-3012311230231023-0312310313121023-3122223302301012-1011130302002022)
- [request_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-2221012012102111-2312110102322101-0103322330302302-1323220122123302-3032133202210302-3121200131100310-1322210101030202-2123321301200213)
- request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-3020312203131200-0020022130002033-0321323213310022-2313111212203031-2030211221112111-0213103303011130-1323002013112332-3320202133102220"></a>

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

<a id="canonical-0130221101230002-1001033220313220-1221020200032210-2202033233212003-1330013222301211-0101001203021310-0213233221103230-3321212111000300"></a>

## Direct properties — clear_secret_info / 213022331132 / 3

<a id="canonical-3232213000032311-3301210102010031-3222232013311122-0233112321032130-2120322312202132-0330000323231203-2202231213101220-1003102131001330"></a>

<a id="canonical-2331132132103002-1100320022301103-1312311232210121-2222313000311311-3132120023330033-2033032110211122-1100202033231323-3231112222303203"></a>

## provider_ref property — clear_secret_info / 213022331132 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0300321231110031-2331012232231000-0123331022213032-3030302211320333-0113230030130112-1201122031033032-0211323200112230-3233131303313211"></a>

<a id="canonical-1100002003031010-3332212301110230-0302232313312022-2222130123330221-0211023122132031-1001033033100212-1330123100110302-3030023010020220"></a>

## URL property — clear_secret_info / 213022331132 / 5

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

<a id="canonical-0213331310032231-2013221102223213-1310103332101132-0133010032231130-3323311033132202-0311302211323033-2010132100102332-0000332230133201"></a>

## Next pages — clear_secret_info / 213022331132 / 6

- [request_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-2221012012102111-2312110102322101-0103322330302302-1323220122123302-3032133202210302-3121200131100310-1322210101030202-2123321301200213)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3303213013223123-0333021210322333-3000030200231312-1002303202333220-2113133012232131-0030321011113311-3012010300230300-3131003321220113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233320103310223-2103332003102333-0210313011033032-0313122303230002-3211332202322133-3221110201022210-1120313223033313-0300320132203010"></a>

## request_headers_to_add — request_headers_to_add / 221013131232 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- request_headers_to_add

<a id="canonical-0212331111220023-1110003001131213-2230212300133010-0032203310311221-3132120120212033-2213331220010100-0122232301011223-1113212132012332"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

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

<a id="canonical-2222112303000101-1210102322011201-2131211033303122-1112113202230320-1301103330302122-2020111100233231-1203233300320130-1130222033120133"></a>

## Direct properties — request_headers_to_add / 221013131232 / 3

<a id="canonical-0111003322111311-0210311322203231-0033231212032013-0120201220320101-3322301111220313-0130312231102133-1010101013000110-0323013230030323"></a>

<a id="canonical-0130313333222233-3211333100101302-0123212020331222-2001132013313111-3133012131133300-0121032211210200-0001010302330310-1201233130332232"></a>

## append property — request_headers_to_add / 221013131232 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-2013010322133123-3312231203130101-0001201323033122-0313203013300222-2002002231011321-0320010312120320-1311220121011202-3021332101220200"></a>

<a id="canonical-0300212230301231-0132111033202123-2301123023030200-1033330223202211-1103202100033021-2301221121332332-1002321010030210-2201300210201031"></a>

## name property — request_headers_to_add / 221013131232 / 5

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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

- [secret_value](data-sources--virtual_host--reference--group-002.md#canonical-1330032123203232-2013202012321330-0033330131300212-1213132002032112-0122223201331312-1123132323101310-2132002113033311-1331011333020313): complete subsection reference.

<a id="canonical-0002221320121320-1033122001021012-2021121302103102-0303312323133110-2231122220223301-3120003322302203-0211231223130313-2203011201111113"></a>

<a id="canonical-3030003000000331-3132302313020100-3011031003300133-3330102021221222-1310122120220123-3100101133021012-0312221303103201-0212131202233200"></a>

## value property — request_headers_to_add / 221013131232 / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

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

<a id="canonical-3300013032113322-3332132201023120-2001221012100023-2323320121331010-0003321112211302-3230332210320323-1301112211031311-0312113011123122"></a>

## Next pages — request_headers_to_add / 221013131232 / 7

- [request_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-1330032123203232-2013202012321330-0033330131300212-1213132002032112-0122223201331312-1123132323101310-2132002113033311-1331011333020313)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1330032123203232-2013202012321330-0033330131300212-1213132002032112-0122223201331312-1123132323101310-2132002113033311-1331011333020313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320213120022321-3321110102332230-2122323223332121-2223110313223022-0231310103122030-1031212000311111-1022202221110130-2213332212000320"></a>

## request_headers_to_add.secret_value — secret_value / 210111120221 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3303213013223123-0333021210322333-3000030200231312-1002303202333220-2113133012232131-0030321011113311-3012010300230300-3131003321220113)
- request_headers_to_add.secret_value

<a id="canonical-1113100003100331-1012223330010101-2311201233313310-2202132323023113-1002120103022112-2322022213110322-1002212202302302-0022023101201032"></a>

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

<a id="canonical-2320020212313103-1111233111232210-1333132211011001-2331310120302320-0312332302220200-0031330202233102-3333310320232110-1113022231223232"></a>

## Direct properties — secret_value / 210111120221 / 3

- [blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-2033131003100211-0303122221321010-0333312111121321-0012031121300212-1333011022011210-3200220232000331-3020131301233330-2223303200330022): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-3011313210203032-1222222310212103-3111133313132122-1222112203233121-0322012013002212-2101111133232210-0101031020303223-0023103231210223): complete subsection reference.

<a id="canonical-2210301120231311-0212130033031311-3131012201101200-3001220320302233-1302020202321331-1122313032203311-3232300302231002-1110202023312012"></a>

## Next pages — secret_value / 210111120221 / 4

- [request_headers_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-2033131003100211-0303122221321010-0333312111121321-0012031121300212-1333011022011210-3200220232000331-3020131301233330-2223303200330022)
- [request_headers_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-3011313210203032-1222222310212103-3111133313132122-1222112203233121-0322012013002212-2101111133232210-0101031020303223-0023103231210223)
- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3303213013223123-0333021210322333-3000030200231312-1002303202333220-2113133012232131-0030321011113311-3012010300230300-3131003321220113)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2033131003100211-0303122221321010-0333312111121321-0012031121300212-1333011022011210-3200220232000331-3020131301233330-2223303200330022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303033311110331-3101002100233020-2333232030110011-2012300023311301-0011332231303323-1133232003213210-1202032213320312-3212310333003211"></a>

## request_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 202322310323 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3303213013223123-0333021210322333-3000030200231312-1002303202333220-2113133012232131-0030321011113311-3012010300230300-3131003321220113)
- [request_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-1330032123203232-2013202012321330-0033330131300212-1213132002032112-0122223201331312-1123132323101310-2132002113033311-1331011333020313)
- request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-1110020211333301-0201002333112132-1123323201210102-1122032113112330-0033210303230021-3311212110111200-2322201312102122-0103010302332210"></a>

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

<a id="canonical-0101302220203013-0301130213030323-0022210000130300-2310100202023122-3231131103320022-2302002212213233-1210100030221102-1022001331323303"></a>

## Direct properties — blindfold_secret_info / 202322310323 / 3

<a id="canonical-3221300001313201-2202013112112012-3210001201320011-0211221022210112-1022212232103113-1310001122321233-3210230123221120-3323312303010110"></a>

<a id="canonical-2201322220232320-3221030212333111-3333221121120003-3230132133130331-3313030031302031-3330320323331032-3331321323003303-2333011331123210"></a>

## decryption_provider property — blindfold_secret_info / 202322310323 / 4

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

<a id="canonical-1130010110321230-3102120121003322-1010221321200311-1132112033100303-0001301202313013-3130212230131221-1310230033220310-3200222213021312"></a>

<a id="canonical-1232232102223112-2302311202201010-3130330012332203-1233101002203320-1212023210223023-1122200310031133-2321130032111301-3331302101020011"></a>

## location property — blindfold_secret_info / 202322310323 / 5

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

<a id="canonical-3303102130111003-1333003000312221-0023001320311002-2133322032102330-1230230222003021-0100332210120222-2321030033213221-0130122233033021"></a>

<a id="canonical-0311021223110022-2122111220030311-1201130213032202-2100133023112203-0102002200102133-0312302201312100-2013202311220022-2321132302231101"></a>

## store_provider property — blindfold_secret_info / 202322310323 / 6

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

<a id="canonical-0330131131222233-2323213030313011-0222323023110033-2230003232021230-0023031311212210-1132030233310000-1231102121123012-1110321232120230"></a>

## Next pages — blindfold_secret_info / 202322310323 / 7

- [request_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-1330032123203232-2013202012321330-0033330131300212-1213132002032112-0122223201331312-1123132323101310-2132002113033311-1331011333020313)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3011313210203032-1222222310212103-3111133313132122-1222112203233121-0322012013002212-2101111133232210-0101031020303223-0023103231210223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031023132332012-1033112131021213-1313333322113232-3032312232131100-0002113221022012-1312222123033130-1213221123013020-3021213331121102"></a>

## request_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 122022321013 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3303213013223123-0333021210322333-3000030200231312-1002303202333220-2113133012232131-0030321011113311-3012010300230300-3131003321220113)
- [request_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-1330032123203232-2013202012321330-0033330131300212-1213132002032112-0122223201331312-1123132323101310-2132002113033311-1331011333020313)
- request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-3121111213002323-1231101112112231-3101333032202332-0200013011002332-2223020023231333-1133033202321323-2031003023020022-1103101121113301"></a>

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

<a id="canonical-1233223330222022-2320013133221132-0003101220231331-2120030320202001-1211021111321331-2103302131311130-3330132102211031-2021012213303211"></a>

## Direct properties — clear_secret_info / 122022321013 / 3

<a id="canonical-1200322010133001-2100232100010330-1133302301320120-3201131331201221-0221220222211133-0120320101303330-3211133301010023-2230331232012120"></a>

<a id="canonical-0332312012131233-3021210333201022-1302130333201332-2223323123303120-0100103011323222-2203112310302200-1003110032230322-3023011203201332"></a>

## provider_ref property — clear_secret_info / 122022321013 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2330021311131111-1132011222001031-0010222130001213-1112132200302132-0322123330122212-1220113330112203-1210310130002200-2211211201220201"></a>

<a id="canonical-3300330012203232-0012121200132211-3331300303022120-3231130023300121-1001213011003023-1200301101133010-3022231012122010-1300321330013001"></a>

## URL property — clear_secret_info / 122022321013 / 5

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

<a id="canonical-2130113000112330-3333120010323110-0110113232312122-0210230222112201-3201210330123311-2320102323032213-0101311123322002-3031302301001000"></a>

## Next pages — clear_secret_info / 122022321013 / 6

- [request_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-1330032123203232-2013202012321330-0033330131300212-1213132002032112-0122223201331312-1123132323101310-2132002113033311-1331011333020313)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003023123311110-1320212332103301-1212120102323032-1212323323322320-2011101120013013-2222003030301331-2232110331003233-0202113023001011"></a>

## response_cookies_to_add — response_cookies_to_add / 203333102001 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- response_cookies_to_add

<a id="canonical-0113212001232011-3133322010021302-3102230131212330-2020322230330311-0130213133032111-3122133022112220-3200232202122232-3312300203123230"></a>

Type: `"list"`. Computed.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Upstream description:

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

<a id="canonical-3011222323203103-1210102001333311-3132020122120220-0202023100112033-1101331211121122-2330100111032113-3020102010312101-2223230000303010"></a>

## Direct properties — response_cookies_to_add / 203333102001 / 3

<a id="canonical-0001011020132012-1232111312113132-0330300223301231-1011002313230031-3232220300331003-2300110310311011-3121013011113200-0013021222011120"></a>

<a id="canonical-0103113301313211-2200022000302323-2223113201303330-1001330230012012-3132210230330002-1210112302033323-1300112221210032-1131003120111010"></a>

## add_domain property — response_cookies_to_add / 203333102001 / 4

Type: `"string"`. Computed.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

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

<a id="canonical-0323331213312223-3021003123002211-1202213220023132-3010013212200011-1022332133021023-3323301323122213-0303012032122120-2201212310221002"></a>

<a id="canonical-0030012303333322-0130233313322120-3332222230311321-2310203120023331-0131031123033123-1032220132122012-3201023301222123-0110320232033033"></a>

## add_expiry property — response_cookies_to_add / 203333102001 / 5

Type: `"string"`. Computed.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

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

- [add_httponly](data-sources--virtual_host--reference--group-002.md#canonical-3133121010112233-3322313331011132-2132333100021313-1103100203002133-0322310213022112-3130232120302020-1130301033103101-2131030113031102): complete subsection reference.

- [add_partitioned](data-sources--virtual_host--reference--group-002.md#canonical-1300012000233302-2313330232300323-2330201210102101-2032131201102111-1012033012111001-0303101333222210-3103312110002222-3110030211130311): complete subsection reference.

<a id="canonical-0211031222022221-1210232233232213-0123111010113301-1210201123011222-3321203311111332-0020311311131231-2211233232100232-0022110201032030"></a>

<a id="canonical-1320232132100333-2302301130200021-0132102203023023-3001322002320130-3031122222313113-3130000210210221-0022110303103110-1020221201003133"></a>

## add_path property — response_cookies_to_add / 203333102001 / 6

Type: `"string"`. Computed.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

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

- [add_secure](data-sources--virtual_host--reference--group-002.md#canonical-1131111231223133-2310003201221100-1031203233212002-1112222321112123-1322013301002233-2203202100133322-0013112323023003-2211003331110111): complete subsection reference.

- [ignore_domain](data-sources--virtual_host--reference--group-002.md#canonical-3130310133231002-1131000313113312-2030022200200233-2011302002303013-3303011201121220-3202203133303302-2011203131131013-2332211220103303): complete subsection reference.

- [ignore_expiry](data-sources--virtual_host--reference--group-002.md#canonical-2030201030022210-2032320011321111-0112121020213023-0333200011123312-3333331012331211-1023111003112233-3010101100331300-0211322231032200): complete subsection reference.

- [ignore_httponly](data-sources--virtual_host--reference--group-002.md#canonical-2113311312232221-0013310131132212-3010132032113032-1201123320023131-0122300301100323-1020233312020223-2231033013201101-2032022103010000): complete subsection reference.

- [ignore_max_age](data-sources--virtual_host--reference--group-002.md#canonical-1213131003133130-1132100030000101-1312020211100300-1023102133112033-3323332101222100-0231120312322020-1322000302113023-3023223103021013): complete subsection reference.

- [ignore_partitioned](data-sources--virtual_host--reference--group-002.md#canonical-3013010331312000-2120221322022202-0012313203013322-1023230122220300-1011300220110232-2010213110003132-3110210010212011-2133323103202010): complete subsection reference.

- [ignore_path](data-sources--virtual_host--reference--group-002.md#canonical-1102211312201202-3111210321011033-3221202120032131-1310122310121202-1332103031102001-2210203322320113-3110230213230232-2212122010333220): complete subsection reference.

- [ignore_samesite](data-sources--virtual_host--reference--group-002.md#canonical-0330232211331123-0102120023332301-3231132200322201-3312221110102011-1310012301122323-1000113222013221-3330022122030133-0312230303222303): complete subsection reference.

- [ignore_secure](data-sources--virtual_host--reference--group-002.md#canonical-1203333011321131-3020123131103020-0121032013213000-1323211132232130-3202000322303300-1330102133201011-1212021303032023-0100122223001102): complete subsection reference.

- [ignore_value](data-sources--virtual_host--reference--group-002.md#canonical-0103130323203132-3111033320232102-1202231000313232-0102103013333002-2332301022003020-2020200133102202-3103033030130030-3331001202111002): complete subsection reference.

<a id="canonical-1223332230031002-3021221300023122-1202320001120111-1301132121100020-0021200302320200-1031203103230123-1212100133020002-1023133010220030"></a>

<a id="canonical-2301003110113133-0223202100221100-3031003300232132-0103333121303323-2112203302200000-1130211310010020-2002100002022000-2013313331130313"></a>

## max_age_value property — response_cookies_to_add / 203333102001 / 7

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

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

<a id="canonical-0033210331200122-0220202000132121-0133233333200332-0202313322332222-3132303020321000-1030002033013000-2100222230022023-1011132021323013"></a>

<a id="canonical-3202032203301331-2323110330323131-2321331212332111-1001012303123012-2200122023012321-0223232211231232-0332132310212001-2212322013001101"></a>

## name property — response_cookies_to_add / 203333102001 / 8

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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

<a id="canonical-0131110323003203-3300311130312322-0010021032303102-3033001202211300-3133130232132202-3203131033323323-2310212232320232-0322100132320110"></a>

<a id="canonical-3201310220010301-1301321021310303-2211210222200220-1123233320102213-2333320201100001-2311231120013320-3110113300133113-0001300031130211"></a>

## overwrite property — response_cookies_to_add / 203333102001 / 9

Type: `"bool"`. Computed.

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

- [samesite_lax](data-sources--virtual_host--reference--group-002.md#canonical-1032311232333023-0123210200322102-2332221303132113-0311311123312313-0130113003223133-1222002122120001-1301232232200131-2113321032033121): complete subsection reference.

- [samesite_none](data-sources--virtual_host--reference--group-002.md#canonical-2233112002222232-2223030333212333-3201132101110203-0211330111331112-0221230222022203-0010121333033210-0122232002122312-1310233330303223): complete subsection reference.

- [samesite_strict](data-sources--virtual_host--reference--group-002.md#canonical-1322201213133222-1023323300033031-0231030302203233-0133010022003131-3302220101130320-3210200311111031-3001003010311323-1132010111102000): complete subsection reference.

- [secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3313221231300020-1321102321213000-1322132323323113-1300023011223200-2300033321130201-1012102200310331-3101122201123203-1330201223331132): complete subsection reference.

<a id="canonical-0132212221332200-0112030120132231-0303203332110213-0221133212132012-3112222223320130-3311011121331322-2300303301330321-0232012032110113"></a>

<a id="canonical-3201133110222001-3232031233212023-2110231200202000-1000002203031322-3330221200213311-2121022231222123-2303231123132111-1002231211033223"></a>

## value property — response_cookies_to_add / 203333102001 / 10

Type: `"string"`. Computed.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

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

<a id="canonical-3300122021032030-2212111030110333-0321203322322333-3232111210322122-3010102103302132-0012233013000222-2032301002110222-0102302312202113"></a>

## Next pages — response_cookies_to_add / 203333102001 / 11

- [response_cookies_to_add.add_httponly](data-sources--virtual_host--reference--group-002.md#canonical-3133121010112233-3322313331011132-2132333100021313-1103100203002133-0322310213022112-3130232120302020-1130301033103101-2131030113031102)
- [response_cookies_to_add.add_partitioned](data-sources--virtual_host--reference--group-002.md#canonical-1300012000233302-2313330232300323-2330201210102101-2032131201102111-1012033012111001-0303101333222210-3103312110002222-3110030211130311)
- [response_cookies_to_add.add_secure](data-sources--virtual_host--reference--group-002.md#canonical-1131111231223133-2310003201221100-1031203233212002-1112222321112123-1322013301002233-2203202100133322-0013112323023003-2211003331110111)
- [response_cookies_to_add.ignore_domain](data-sources--virtual_host--reference--group-002.md#canonical-3130310133231002-1131000313113312-2030022200200233-2011302002303013-3303011201121220-3202203133303302-2011203131131013-2332211220103303)
- [response_cookies_to_add.ignore_expiry](data-sources--virtual_host--reference--group-002.md#canonical-2030201030022210-2032320011321111-0112121020213023-0333200011123312-3333331012331211-1023111003112233-3010101100331300-0211322231032200)
- [response_cookies_to_add.ignore_httponly](data-sources--virtual_host--reference--group-002.md#canonical-2113311312232221-0013310131132212-3010132032113032-1201123320023131-0122300301100323-1020233312020223-2231033013201101-2032022103010000)
- [response_cookies_to_add.ignore_max_age](data-sources--virtual_host--reference--group-002.md#canonical-1213131003133130-1132100030000101-1312020211100300-1023102133112033-3323332101222100-0231120312322020-1322000302113023-3023223103021013)
- [response_cookies_to_add.ignore_partitioned](data-sources--virtual_host--reference--group-002.md#canonical-3013010331312000-2120221322022202-0012313203013322-1023230122220300-1011300220110232-2010213110003132-3110210010212011-2133323103202010)
- [response_cookies_to_add.ignore_path](data-sources--virtual_host--reference--group-002.md#canonical-1102211312201202-3111210321011033-3221202120032131-1310122310121202-1332103031102001-2210203322320113-3110230213230232-2212122010333220)
- [response_cookies_to_add.ignore_samesite](data-sources--virtual_host--reference--group-002.md#canonical-0330232211331123-0102120023332301-3231132200322201-3312221110102011-1310012301122323-1000113222013221-3330022122030133-0312230303222303)
- [response_cookies_to_add.ignore_secure](data-sources--virtual_host--reference--group-002.md#canonical-1203333011321131-3020123131103020-0121032013213000-1323211132232130-3202000322303300-1330102133201011-1212021303032023-0100122223001102)
- [response_cookies_to_add.ignore_value](data-sources--virtual_host--reference--group-002.md#canonical-0103130323203132-3111033320232102-1202231000313232-0102103013333002-2332301022003020-2020200133102202-3103033030130030-3331001202111002)
- [response_cookies_to_add.samesite_lax](data-sources--virtual_host--reference--group-002.md#canonical-1032311232333023-0123210200322102-2332221303132113-0311311123312313-0130113003223133-1222002122120001-1301232232200131-2113321032033121)
- [response_cookies_to_add.samesite_none](data-sources--virtual_host--reference--group-002.md#canonical-2233112002222232-2223030333212333-3201132101110203-0211330111331112-0221230222022203-0010121333033210-0122232002122312-1310233330303223)
- [response_cookies_to_add.samesite_strict](data-sources--virtual_host--reference--group-002.md#canonical-1322201213133222-1023323300033031-0231030302203233-0133010022003131-3302220101130320-3210200311111031-3001003010311323-1132010111102000)
- [response_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3313221231300020-1321102321213000-1322132323323113-1300023011223200-2300033321130201-1012102200310331-3101122201123203-1330201223331132)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3133121010112233-3322313331011132-2132333100021313-1103100203002133-0322310213022112-3130232120302020-1130301033103101-2131030113031102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333122302010120-0133103001313201-3003323002033221-3313220210322222-0022003012223100-1212230130032011-0231000013201223-1113010023312100"></a>

## response_cookies_to_add.add_httponly — add_httponly / 330131302203 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.add_httponly

<a id="canonical-3331200311131232-2111312122313330-3311022112022321-0210222232203233-2322120132222221-0330300102033233-0002221112121012-2231003222320123"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2021221003210313-3130103310033213-2030112100130101-3321300330302211-2302311001022011-3103310320023302-2211213013303121-0033213323111011"></a>

## Direct properties — add_httponly / 330131302203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220330120202300-3013122030222330-2312300330311321-3101322031031122-2210230100322323-1123202301220320-1030320233022122-2012123131333213"></a>

## Next pages — add_httponly / 330131302203 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1300012000233302-2313330232300323-2330201210102101-2032131201102111-1012033012111001-0303101333222210-3103312110002222-3110030211130311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100321130233020-1002113100031331-0012230330320333-1120200100100031-3223203120023113-1103302220001300-0220302122012101-0113303013131301"></a>

## response_cookies_to_add.add_partitioned — add_partitioned / 301201233220 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.add_partitioned

<a id="canonical-2030311302202101-1021020302223000-2330003013102031-1031233032103232-3210302012231012-3212130200300211-1000222233330032-0202323320031200"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3000233032030103-0032020132010013-0030100200121113-1132313033311031-2331013033312020-2323110300222310-3321031033130233-1313311212312311"></a>

## Direct properties — add_partitioned / 301201233220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320331012231310-2232011202000002-2112100101033311-1131133301111031-2130022203321001-2123231111023010-1120221122332013-3203132103211211"></a>

## Next pages — add_partitioned / 301201233220 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1131111231223133-2310003201221100-1031203233212002-1112222321112123-1322013301002233-2203202100133322-0013112323023003-2211003331110111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230310231320113-3100222100302301-2131333112002003-1122300132131311-1130013030332103-3321022022021222-1331320231221312-2231321000122120"></a>

## response_cookies_to_add.add_secure — add_secure / 120022131200 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.add_secure

<a id="canonical-3003210300322131-1210230110002132-0111332020330200-2212311130233302-0330011012323322-1132211132132020-3220101313231112-1023130121100010"></a>

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

<a id="canonical-0030012031103120-2323032202133133-3211200303202321-1013223230122101-1030231032302010-1322331311001200-3302331030033222-2210121111023023"></a>

## Direct properties — add_secure / 120022131200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220210200023221-1122011201113133-3331121211123132-3231320113102203-0320002330303133-0301220101203223-1200032223122121-0321032330000203"></a>

## Next pages — add_secure / 120022131200 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3130310133231002-1131000313113312-2030022200200233-2011302002303013-3303011201121220-3202203133303302-2011203131131013-2332211220103303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111130311232011-0300120133022031-1210200101000113-2021023202323020-3311103333222202-2201120210121323-2333220100100223-2012331012033032"></a>

## response_cookies_to_add.ignore_domain — ignore_domain / 012121330110 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_domain

<a id="canonical-2300121232212121-0121210302102312-2113030112230210-3213202010320220-3000113302302223-2020031013213103-3000003131023032-3302121313313112"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3013103033320223-0120002310020032-0213301132223333-0311332330002123-3030303133100202-3223223120212203-1330311132300210-2212003202021311"></a>

## Direct properties — ignore_domain / 012121330110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010133211111222-3322112322213132-3221100300331131-2331230323303110-0132022133012001-2123310320121113-3211310213101332-3230103212231330"></a>

## Next pages — ignore_domain / 012121330110 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2030201030022210-2032320011321111-0112121020213023-0333200011123312-3333331012331211-1023111003112233-3010101100331300-0211322231032200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232011120103112-3031020021322311-0210012332200321-0022120013133331-2231133033110000-3001133003103131-2011133211133201-1232101200312312"></a>

## response_cookies_to_add.ignore_expiry — ignore_expiry / 033102122303 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_expiry

<a id="canonical-0131302033331002-2302213311003033-0320130131231012-2312100333202102-0023100123222031-1213230122321313-3121023211212321-3101233013232101"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3033203221111202-2210022103223100-3230133113211033-0133013322030300-0333332223120000-3211130132322113-0331313233132333-0303132022111113"></a>

## Direct properties — ignore_expiry / 033102122303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002332131212021-0203203030131211-0203330230301030-0303230231320003-0220302233320333-1200320331210221-0321310103333220-0033230301330000"></a>

## Next pages — ignore_expiry / 033102122303 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2113311312232221-0013310131132212-3010132032113032-1201123320023131-0122300301100323-1020233312020223-2231033013201101-2032022103010000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222001302133100-3013000102322031-2030000130132011-1022123323231030-1103332121131103-2033222222122000-0121321313120013-1321131022002332"></a>

## response_cookies_to_add.ignore_httponly — ignore_httponly / 031000231203 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_httponly

<a id="canonical-0301030232121022-0112212033023121-2032223012003223-2230102331131002-3330102031321322-3333222213312032-3120310303312112-0330111332132312"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3202321322031231-1202021233030030-0333203223203202-1012130031122210-2232033201311312-3300312310103323-0220322133300211-1233330103022212"></a>

## Direct properties — ignore_httponly / 031000231203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120010123232120-2110030312310121-3311000103021013-0333321003302121-1300132002302001-3231003022132110-1303013302010222-0312332302203003"></a>

## Next pages — ignore_httponly / 031000231203 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1213131003133130-1132100030000101-1312020211100300-1023102133112033-3323332101222100-0231120312322020-1322000302113023-3023223103021013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110032311310030-2020222011311131-3222331321123312-1220323231223233-3103320031002302-1030030220323312-3320311311021201-2321001130130100"></a>

## response_cookies_to_add.ignore_max_age — ignore_max_age / 212333202223 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_max_age

<a id="canonical-1312333330322012-1230201311201103-2333211023331033-2201220330121222-1321331313012310-1310302302331010-0333121011001310-1121031003120230"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2031011111001113-1100030302321010-0111020121332221-1000130012231113-3130013103023021-3220300023233212-0003313120300121-0022121032032230"></a>

## Direct properties — ignore_max_age / 212333202223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001200201210320-3200022022103202-3023123122212311-2301232022211300-1303020120012133-2122023003032313-0111321033200011-2123312003132332"></a>

## Next pages — ignore_max_age / 212333202223 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3013010331312000-2120221322022202-0012313203013322-1023230122220300-1011300220110232-2010213110003132-3110210010212011-2133323103202010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001001230200301-3112203203211330-2131232230030210-3331210110131233-2121221023200223-3330130022223321-0203023020102210-3233303323301011"></a>

## response_cookies_to_add.ignore_partitioned — ignore_partitioned / 002010222323 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_partitioned

<a id="canonical-3110203131101201-2120323010230203-0202113001231330-3013301310123112-2002133221302221-0101033122100131-1331130012322133-2301222312322232"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1320310301123200-0030223213330131-3010313020213131-2323331113001023-3323203220203133-1200100223233021-0320231130000200-1130212020031123"></a>

## Direct properties — ignore_partitioned / 002010222323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021012002121003-2010201102203010-3233110001222201-1212020110030121-2210211102103320-3201030212322002-1330131332203331-1310303210302130"></a>

## Next pages — ignore_partitioned / 002010222323 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1102211312201202-3111210321011033-3221202120032131-1310122310121202-1332103031102001-2210203322320113-3110230213230232-2212122010333220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222311033033120-3210013221212302-0203222021121030-2002132233320110-2331310121130202-0013312213001033-0033113130022113-0200011112300210"></a>

## response_cookies_to_add.ignore_path — ignore_path / 112231201233 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_path

<a id="canonical-2130320121310303-1332032310213100-2222113113013213-2303023222113201-1312310230031200-0122023033133131-1230233332231203-3230131012023332"></a>

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

<a id="canonical-0311103123210130-2232130313323030-2310210323130133-2031232020123122-3322232013310001-3033301310000302-1202331033030021-1233232112031220"></a>

## Direct properties — ignore_path / 112231201233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222032211321002-2112310111311030-3113230231333221-3320023323033230-0003321223332131-1323122233210310-2200100031213001-1030332322233313"></a>

## Next pages — ignore_path / 112231201233 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0330232211331123-0102120023332301-3231132200322201-3312221110102011-1310012301122323-1000113222013221-3330022122030133-0312230303222303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232031212222303-0221332201133022-1010033001312011-1012223222213330-3122122011020333-2212332022030300-3110020111021333-1222313223103101"></a>

## response_cookies_to_add.ignore_samesite — ignore_samesite / 213321321212 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_samesite

<a id="canonical-3021320121212200-3213213331333010-3110311312301132-0231220113203222-2131102100213123-1320001210231333-3131200320223223-1112322321131001"></a>

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

<a id="canonical-2330022111200110-0222002230333310-2203202032101033-0310121032122111-0202220232021320-3113323232122111-2123123112230302-0032233331111320"></a>

## Direct properties — ignore_samesite / 213321321212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020123302001110-0202122121030130-1131230131232213-1301303103233111-2221221311033212-2001323311303233-2022030013322101-2330002300331210"></a>

## Next pages — ignore_samesite / 213321321212 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1203333011321131-3020123131103020-0121032013213000-1323211132232130-3202000322303300-1330102133201011-1212021303032023-0100122223001102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113100110002120-3211112211030111-1100220111312321-0132032300231303-0333201331323300-1131230000102320-0032323103333021-2100302022233303"></a>

## response_cookies_to_add.ignore_secure — ignore_secure / 133033132303 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_secure

<a id="canonical-2302001212203030-1130010112322122-2033200121331232-3200023020330333-1102221111111232-0221131320122112-2011203322210322-1030202221211032"></a>

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

<a id="canonical-0102021211232030-2033101022320213-1033112032301320-1130321002210003-1332300022020211-0120203130333300-3012122302332311-3200031313031231"></a>

## Direct properties — ignore_secure / 133033132303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332032120211333-3323103333130332-2202233121101123-1221020230010122-3022323011312003-3200111200110201-1031300201031203-2113312111013230"></a>

## Next pages — ignore_secure / 133033132303 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0103130323203132-3111033320232102-1202231000313232-0102103013333002-2332301022003020-2020200133102202-3103033030130030-3331001202111002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320301232312212-3020023202331202-3222000110003203-3013021302131002-2212310213021210-1311000132002310-2233231031301330-2322023210032130"></a>

## response_cookies_to_add.ignore_value — ignore_value / 220222131111 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_value

<a id="canonical-3222313001302022-0111130312012223-0231010310010311-1112112223103321-2201303303000332-3322220033221033-1301013100210230-0020213102300331"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3023030112301033-1022201300112312-2222233210233113-3112002213012310-0310132002120012-2210211303210022-0200111122223211-0223213003201132"></a>

## Direct properties — ignore_value / 220222131111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300133202111133-2012000312100031-0011003330203212-0100231331302302-3103222223021112-0120303300002331-1131202031011133-3321010200211313"></a>

## Next pages — ignore_value / 220222131111 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1032311232333023-0123210200322102-2332221303132113-0311311123312313-0130113003223133-1222002122120001-1301232232200131-2113321032033121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213110031122130-2330333133002130-3200003312312233-0133221212201203-2120111332112210-1323202020222230-1302302211301331-3310010103330302"></a>

## response_cookies_to_add.samesite_lax — samesite_lax / 231121121013 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.samesite_lax

<a id="canonical-2033121212113103-0210022033031232-2333211101213203-2320033333020211-0031301210202211-1103111012210301-0213003303311011-0321333020001302"></a>

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

<a id="canonical-2201211022321123-3011021012103112-1011103310133332-2221022120012202-2311320300212122-1230202020100003-0121101033023011-0130121113212203"></a>

## Direct properties — samesite_lax / 231121121013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121112230201333-0011012133032122-3330030012123320-0233121022013232-0111312202103111-1121123203110301-3130123013023331-1012331220103233"></a>

## Next pages — samesite_lax / 231121121013 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2233112002222232-2223030333212333-3201132101110203-0211330111331112-0221230222022203-0010121333033210-0122232002122312-1310233330303223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211020320100200-3201302112003321-1331030010310302-2301323012002313-3220330020312222-3030332230102320-2122003003011113-2110313011110311"></a>

## response_cookies_to_add.samesite_none — samesite_none / 022200132111 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.samesite_none

<a id="canonical-3220320102032323-0102112031021131-2101130101011021-3133023100130231-0301332213303331-0232020023003022-0200212031300200-3202111010023000"></a>

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

<a id="canonical-1110133122110130-2131111033130132-0000230102002333-1210100113302102-0122122322130111-0110312333213101-3203101001032122-0200233120001103"></a>

## Direct properties — samesite_none / 022200132111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112002222013103-1213222002120313-1010100130030113-0320012213221032-0100021023212221-1312211232321232-2302032300133032-0022030103103310"></a>

## Next pages — samesite_none / 022200132111 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1322201213133222-1023323300033031-0231030302203233-0133010022003131-3302220101130320-3210200311111031-3001003010311323-1132010111102000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223213201210232-2221200123032001-3100011121121310-0302301120013303-1332230332010001-0302233013012132-2100012213112303-3320013031202322"></a>

## response_cookies_to_add.samesite_strict — samesite_strict / 102201310203 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.samesite_strict

<a id="canonical-0020132200132200-2312023030212103-0101121332113211-0300223301032223-3131310123120132-1030202201211311-3033223330103300-1001010231133231"></a>

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

<a id="canonical-2110311133032120-1130110330203212-1301203022023300-3101232010320313-2311120320122231-3310211100322133-0201030322122222-2102001002303322"></a>

## Direct properties — samesite_strict / 102201310203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022130310002321-0121100003313331-2323331011130031-3302102121211331-2013310210001001-3301202220203030-0022312111211023-3123023020322330"></a>

## Next pages — samesite_strict / 102201310203 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3313221231300020-1321102321213000-1322132323323113-1300023011223200-2300033321130201-1012102200310331-3101122201123203-1330201223331132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213032330233322-0221111301230031-3120000031213100-1220002330321002-2232121001120102-0230233013001300-0020232002220123-2202321100032132"></a>

## response_cookies_to_add.secret_value — secret_value / 022011003132 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.secret_value

<a id="canonical-0230220011133210-1121201122001221-2000212031123303-2011300012231322-0222223121102113-1020232220032232-2213211003211031-3233111203002212"></a>

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

<a id="canonical-0203200310113113-2001210111102031-1323313201312222-1333001112223312-3003023220131021-1200223331220331-2303013233102232-2303301220323021"></a>

## Direct properties — secret_value / 022011003132 / 3

- [blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-3300132111031211-1300013102022033-3332200320011200-2012102022101122-2001121123201220-1320332302132110-3321311012311203-1330022030301103): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-0102213233221021-3120100212310301-0312202123311301-3333333110333131-3030200311231130-1130102230300212-3133223031012311-2112320110222021): complete subsection reference.

<a id="canonical-1210011233221022-0221110100311202-2001130033121222-3113101021120322-1011111122130000-3001333311322311-1302210202100211-0230102300222100"></a>

## Next pages — secret_value / 022011003132 / 4

- [response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-3300132111031211-1300013102022033-3332200320011200-2012102022101122-2001121123201220-1320332302132110-3321311012311203-1330022030301103)
- [response_cookies_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-0102213233221021-3120100212310301-0312202123311301-3333333110333131-3030200311231130-1130102230300212-3133223031012311-2112320110222021)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3300132111031211-1300013102022033-3332200320011200-2012102022101122-2001121123201220-1320332302132110-3321311012311203-1330022030301103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111203033332113-0031131121233122-3210203323012320-3100020022123300-1210332200202012-0130313320020232-0122012331111210-0000120033232021"></a>

## response_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 330223111102 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [response_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3313221231300020-1321102321213000-1322132323323113-1300023011223200-2300033321130201-1012102200310331-3101122201123203-1330201223331132)
- response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-0121002103233003-2010103032030132-2022032223201103-3322111330020103-2222021232310220-1223233302111221-0230220220201122-1022002331131202"></a>

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

<a id="canonical-1221330212203313-0120312121103303-1020200330111213-0211222210130013-0121211211022012-3230323212030131-0023212233300301-0310302100031003"></a>

## Direct properties — blindfold_secret_info / 330223111102 / 3

<a id="canonical-0012113123020121-1200013002032230-3330322213112113-1020322132300102-0110011223303332-2113013103003112-2300312331121330-3233223133103002"></a>

<a id="canonical-2322312100210300-3122321311010321-1212331320111003-1201302210131101-0131033101322133-0220111210313112-1030321220310332-3123211023100011"></a>

## decryption_provider property — blindfold_secret_info / 330223111102 / 4

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

<a id="canonical-0302122213101131-1201112003300311-0100202202232231-3311121103123012-2323231121121112-2100023230202130-0333031121231000-1332322330021013"></a>

<a id="canonical-1133003011112213-3322103003030110-3220110333031012-0333303023322331-0212321133131222-3010301212033322-2300311312012110-2311333113312023"></a>

## location property — blindfold_secret_info / 330223111102 / 5

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

<a id="canonical-1020312200013313-2020300101100203-0232321123331320-3121113012310200-1231310002031023-0332120303101003-0112111203122001-0123002323122201"></a>

<a id="canonical-1001333001200212-3121030311223100-3211322130021323-2013322331302011-2023212332003013-0001310323003323-0012232320322130-1201313120003032"></a>

## store_provider property — blindfold_secret_info / 330223111102 / 6

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

<a id="canonical-3131223113222001-3111102210301221-0031103020222231-0312311213231301-0310223232203311-0133113200013100-0113030122211312-2130233211030311"></a>

## Next pages — blindfold_secret_info / 330223111102 / 7

- [response_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3313221231300020-1321102321213000-1322132323323113-1300023011223200-2300033321130201-1012102200310331-3101122201123203-1330201223331132)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0102213233221021-3120100212310301-0312202123311301-3333333110333131-3030200311231130-1130102230300212-3133223031012311-2112320110222021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122121313322003-3123000032222123-2313330212221313-2311211313302133-1011010022333202-0311003200133212-1201011203122200-3312221220232223"></a>

## response_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 101310100222 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [response_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3313221231300020-1321102321213000-1322132323323113-1300023011223200-2300033321130201-1012102200310331-3101122201123203-1330201223331132)
- response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-3313221201232231-0032112231301020-0000102331131233-2230002130021321-1300331221033122-2331012012100222-3333311331212011-1213221010013302"></a>

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

<a id="canonical-2123033101123213-3023313301231120-0023022130100200-3303210022131210-3210230230201033-3320310222313102-2131121232200012-1103332330103122"></a>

## Direct properties — clear_secret_info / 101310100222 / 3

<a id="canonical-2130100100231230-0111333213301122-2223013210030123-0210223112111132-1332013333011322-0330223010022203-0101320132230022-3033120023232331"></a>

<a id="canonical-0000323222102331-3231301222331002-3110101302302310-0220303311013030-2211211310211122-3123200323002110-0003303000333121-0210130310233320"></a>

## provider_ref property — clear_secret_info / 101310100222 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2322300220130300-2033203000011122-2111200122010030-2331010010200213-2223120132233322-3103001213000023-3221102222232333-0212231210102202"></a>

<a id="canonical-3313300113113023-1331010322233003-0301121230102210-3100021321333133-0001211302202023-1021121300220230-2002113030222132-1231223302020100"></a>

## URL property — clear_secret_info / 101310100222 / 5

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

<a id="canonical-3330203112021003-2122110010313211-2011301321030030-3232303101030030-0211202323031031-1133131310020103-0120100221203000-2032103011033222"></a>

## Next pages — clear_secret_info / 101310100222 / 6

- [response_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3313221231300020-1321102321213000-1322132323323113-1300023011223200-2300033321130201-1012102200310331-3101122201123203-1330201223331132)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1122301030300332-0033230331023011-1102233120003323-3100320132003322-0030000111220313-3302203220201033-0222213312322000-2123201002112113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231000331231111-0013100102201301-2033030010110133-3331100332323110-0232233111030101-2232212202323320-0023030002331100-1112210233033102"></a>

## response_headers_to_add — response_headers_to_add / 300210112002 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- response_headers_to_add

<a id="canonical-0101202001030202-2332212330311113-0131001010330313-3313032330010131-2020200230233013-3311312012021031-2120012311323110-1223013022101202"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

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

<a id="canonical-3112322012231303-3032101212300301-0312220332001112-1200100310301111-2031003022310311-2031231023123211-3302021101033303-0011103003310110"></a>

## Direct properties — response_headers_to_add / 300210112002 / 3

<a id="canonical-3303003312131233-3030200320213123-3100123110201120-3131321300002023-2233310232001113-2010231121122122-3320330023303211-3030031322132203"></a>

<a id="canonical-3000222000311012-2221012332310321-1020211213331120-2122222011332313-2322233123320301-2023312203301111-1023211022101013-1121200333211010"></a>

## append property — response_headers_to_add / 300210112002 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-2003113000313011-3210313331203312-2212323303332232-3131111332022133-1231002120203303-2020000303233020-2210103001332230-3332312130233310"></a>

<a id="canonical-0033133031123232-0030022103011231-2302232101301102-3123233333231112-3011232301022001-2030231332121303-0132102312010031-3000122333023312"></a>

## name property — response_headers_to_add / 300210112002 / 5

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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

- [secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3101230321222032-3323212212032112-0000210100132300-2200202302113313-3202033223320201-3113222013230022-1120310312103300-3330002320302120): complete subsection reference.

<a id="canonical-3322312333311133-0031013222000232-2233131202221213-1100201210300331-3301011233011300-3011131010300113-1200103303323130-0210113003213303"></a>

<a id="canonical-3223013103022300-0003210133031122-0220302312332103-1121133333332112-1033300232201020-3230321120223033-3212230322003001-1010013110112100"></a>

## value property — response_headers_to_add / 300210112002 / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

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

<a id="canonical-0203013133321311-2310301330011332-2223332220223202-3130030120031220-2001313330010231-2103330110101012-1011121112201313-0120120232032211"></a>

## Next pages — response_headers_to_add / 300210112002 / 7

- [response_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3101230321222032-3323212212032112-0000210100132300-2200202302113313-3202033223320201-3113222013230022-1120310312103300-3330002320302120)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3101230321222032-3323212212032112-0000210100132300-2200202302113313-3202033223320201-3113222013230022-1120310312103300-3330002320302120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122010230223010-1311203312310322-3030130112002223-0020231303113331-0013102013323300-2023033311211001-0230011331233312-2301300310023122"></a>

## response_headers_to_add.secret_value — secret_value / 333011323220 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-1122301030300332-0033230331023011-1102233120003323-3100320132003322-0030000111220313-3302203220201033-0222213312322000-2123201002112113)
- response_headers_to_add.secret_value

<a id="canonical-3100032201022001-2001120211311323-2303132030311023-1223333122211033-2111333132122321-0032110321210032-1303101313331301-0220323001111102"></a>

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

<a id="canonical-3113202200000111-3120110322212303-3232300312022323-2103111322331233-2012321210013300-3300103320201103-3112230300122021-0213202012331021"></a>

## Direct properties — secret_value / 333011323220 / 3

- [blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-2012122311123233-1323213122212312-1301311230331021-0011110120223113-2303130021033010-1120320122031200-2120330332002223-3103211030113022): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-0310123303110031-0303203333000323-2223131103332331-0200031312112020-1120223122110221-2312231113313233-0010021002321202-0100213233303303): complete subsection reference.

<a id="canonical-2013002233310000-0002002230101003-3310100333323300-2322233231233323-1211003020121322-0210312311230332-0021030103112203-0122133002201332"></a>

## Next pages — secret_value / 333011323220 / 4

- [response_headers_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-2012122311123233-1323213122212312-1301311230331021-0011110120223113-2303130021033010-1120320122031200-2120330332002223-3103211030113022)
- [response_headers_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-0310123303110031-0303203333000323-2223131103332331-0200031312112020-1120223122110221-2312231113313233-0010021002321202-0100213233303303)
- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-1122301030300332-0033230331023011-1102233120003323-3100320132003322-0030000111220313-3302203220201033-0222213312322000-2123201002112113)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2012122311123233-1323213122212312-1301311230331021-0011110120223113-2303130021033010-1120320122031200-2120330332002223-3103211030113022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133112222113302-2012103301212233-3222221223213201-0302310032220003-2101312031210123-1211203033303101-2001011330330210-0203130022223313"></a>

## response_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 100101203321 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-1122301030300332-0033230331023011-1102233120003323-3100320132003322-0030000111220313-3302203220201033-0222213312322000-2123201002112113)
- [response_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3101230321222032-3323212212032112-0000210100132300-2200202302113313-3202033223320201-3113222013230022-1120310312103300-3330002320302120)
- response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0320122313112313-3331203313022000-0020013010012332-0133303123120012-0301223323113003-1323202202311113-0103112102302020-3111312210311233"></a>

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

<a id="canonical-1003100132033010-3200322100000231-0221332030330211-1311031102003130-3233000222033303-3113030100331020-3002233101013203-2020230120301112"></a>

## Direct properties — blindfold_secret_info / 100101203321 / 3

<a id="canonical-2321220203333331-1210133320202232-0003321312003130-3020101102233113-2330202230100231-2103013212232033-2220000302030330-1233011223003331"></a>

<a id="canonical-0302032300303031-1232202202320323-2023323023012213-2010031113121203-1000021013303232-0303220201330000-1223302003231221-2001221221303000"></a>

## decryption_provider property — blindfold_secret_info / 100101203321 / 4

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

<a id="canonical-3000332200020302-1200112232231112-2111231332323331-0213223100200310-0203001300120302-2032202130030311-1201102110132112-1032123023322310"></a>
