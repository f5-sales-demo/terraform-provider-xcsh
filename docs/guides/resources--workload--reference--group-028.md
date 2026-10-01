---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-0332231033012210-2223111013021011-3312210133012320-1333303212033130-3331210032310021-1001022111121230-3022310310010321-2000331031210020"></a>

## Next pages — redirect_route / 030030011103 / 5

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers](resources--workload--reference--group-028.md#canonical-0032131122323001-1033332220221013-0310210230202112-2321033223002123-2113001013123333-2232332230200021-3130102220022032-3130121230320302)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-028.md#canonical-1303023022032323-0011211310021103-2302310103100323-1020331133202010-1211220133311022-0313230133233210-0330011000132000-0200302302220132)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path](resources--workload--reference--group-028.md#canonical-3012300121100130-2123213313031300-1111323022000031-3200021200320130-3002333023012223-3010202011010023-0133110113330213-0102003333203322)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-028.md#canonical-3302110132323220-0130312011310200-3213310121101213-2332012011222213-2322213101231122-2221221001002212-3303021101022230-2212202010103100)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0032131122323001-1033332220221013-0310210230202112-2321033223002123-2113001013123333-2232332230200021-3130102220022032-3130121230320302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020303101001301-3331321233103313-1030323311232210-2330121303202021-1303022013020123-0110222311323331-0203100121122030-2021210030013333"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers — headers / 202330001223 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-027.md#canonical-0330013222233302-2210122030021330-0131201133002112-2132231131331102-2322210303201110-1121010222301222-3302213332031013-0122333212301102)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-0133122232220013-3102021303211211-1022320323111012-2211033032122133-3023323021301102-0222010003123322-2123302032010122-1121331232201200"></a>

Type: `"object"`. list nested block, Optional.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
    "regex")}
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
    "minItems": 0,
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013233310021323-1303201001212110-1032133233212312-1113312123031100-0203320010001202-1213230231131011-0213033311002301-3123102022110233"></a>

## Direct properties — headers / 202330001223 / 3

<a id="canonical-0232200101013201-2312000100322330-3132010332122101-2130312030101100-1331023011021031-0312021310302233-3311321131302202-1120023210310023"></a>

<a id="canonical-1020031313121121-1130332311122020-3112112231222310-1010000130223203-2322222121321123-2212033123112131-2120333321133130-1022021210001213"></a>

## exact property — headers / 202330001223 / 4

Type: `"string"`. Optional.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

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
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-3132111112033021-1023333321313211-1323220113130221-3211120002231010-2033130030033202-1100121233232301-2313032033300031-3302032230313003"></a>

<a id="canonical-0223031210032111-0213231331010010-0103021133131312-0311032302033010-1113120333201303-0321310132030320-2313013021223033-0332003002210233"></a>

## invert_match property — headers / 202330001223 / 5

Type: `"bool"`. Optional.

Invert the result of the match to detect missing header or non-matching value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1200302003123100-2101212032022033-2221020033131322-2333220130013313-3312223111030101-2103312001300202-2030300301210302-2031333322233103"></a>

<a id="canonical-3223112203021233-0131112303221011-3012020111322233-3311010302322132-0032032212223312-0011030212110222-3222333000231320-1233123332210003"></a>

## name property — headers / 202330001223 / 6

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3323210012220212-3311330133022211-0112223102102121-2122313021302313-3002112201313110-2032303002033200-3221011112311031-0212300332203103"></a>

<a id="canonical-1211123200033120-0110223321012130-0023002113101222-3221322111332120-1222021133310211-0322312132210003-1130111032220012-2203210133030003"></a>

## presence property — headers / 202330001223 / 7

Type: `"bool"`. Optional.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0203231332213103-3112000333322303-0231332320221122-2032212321211022-3320300100231033-2002300333010022-0320121013002223-0231120101302220"></a>

<a id="canonical-3111022321123123-2002030132201311-2303200220322301-0001320322220220-2211202033101033-2320320301303311-1331013211021011-3323211332300213"></a>

## regex property — headers / 202330001223 / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

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
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1033220113330323-1232320122301032-0220211112303302-0112021131302030-2221200320330010-0100013020333022-2321122112321332-2112302031023020"></a>

## Next pages — headers / 202330001223 / 9

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-027.md#canonical-0330013222233302-2210122030021330-0131201133002112-2132231131331102-2322210303201110-1121010222301222-3302213332031013-0122333212301102)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1303023022032323-0011211310021103-2302310103100323-1020331133202010-1211220133311022-0313230133233210-0330011000132000-0200302302220132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031333321131020-2121100131322012-3222310232010123-3020101001032032-2221123001323103-3222221023133231-3133310011310023-2121032010031302"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port — incoming_port / 301101101120 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-027.md#canonical-0330013222233302-2210122030021330-0131201133002112-2132231131331102-2322210303201110-1121010222301222-3302213332031013-0122333212301102)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-1313022113032010-2002202301320122-3232122132100210-1323100131033233-0323121031212313-3022013133200211-0023112320312330-3301310103203311"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
incoming_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230132223221100-2320112223231013-3101102110232300-1113323302003223-3021013013311121-2321302211003112-1300010311021032-3311300201012213"></a>

## Direct properties — incoming_port / 301101101120 / 3

- [no_port_match](resources--workload--reference--group-028.md#canonical-0021233132001133-2331002131121230-3112332120113032-0333113012011200-0320100031220123-2213021110230101-0112202030302213-3330201022331222): complete subsection reference.

<a id="canonical-0210302232010132-0203333032013311-1003310032303232-1230210131122012-1012111233311323-1102123011131211-3203322223100232-1022020202211031"></a>

<a id="canonical-2132022131031120-0220013231123012-0310011302220100-0201303123300111-2203223023001301-3230013332131002-0230311122113013-3232220303330032"></a>

## port property — incoming_port / 301101101120 / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2331330103002021-3313221203010001-1230001111002222-3022112320010202-3110032021221100-2010202332112100-2011011222121321-1011033200300022"></a>

<a id="canonical-3123103232203202-1203012101211301-1212321332122232-3123322112032110-0331131021232313-1012233221323312-0021312012103113-2031320211010211"></a>

## port_ranges property — incoming_port / 301101101120 / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-1220333231331332-1312113013333001-0031133213010131-2030220130012131-1123010020120003-2223302220311023-2120210311102132-1202012332310121"></a>

## Next pages — incoming_port / 301101101120 / 6

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match](resources--workload--reference--group-028.md#canonical-0021233132001133-2331002131121230-3112332120113032-0333113012011200-0320100031220123-2213021110230101-0112202030302213-3330201022331222)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-027.md#canonical-0330013222233302-2210122030021330-0131201133002112-2132231131331102-2322210303201110-1121010222301222-3302213332031013-0122333212301102)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0021233132001133-2331002131121230-3112332120113032-0333113012011200-0320100031220123-2213021110230101-0112202030302213-3330201022331222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313232023211321-2030102021311313-0320211232300223-1032002113223020-0201102331323201-2211322211012022-0310201230301020-0100010010220013"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match — no_port_match / 001303002311 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-027.md#canonical-0330013222233302-2210122030021330-0131201133002112-2132231131331102-2322210303201110-1121010222301222-3302213332031013-0122333212301102)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-028.md#canonical-1303023022032323-0011211310021103-2302310103100323-1020331133202010-1211220133311022-0313230133233210-0330011000132000-0200302302220132)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-1123121003313212-1132231122100332-3133222110313303-2013222122012212-1131312032300120-0122022031231012-2313100302303011-3130320120031210"></a>

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
no_port_match = {}
```

<a id="canonical-3202123321222030-0302013211010311-2023331032311132-3022103322112131-1032322021012031-1200023030211023-1130330202123101-0101313213301213"></a>

## Direct properties — no_port_match / 001303002311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330322222102310-3033100012210332-2133322332102300-2123333030012130-1311320021212212-3103320221122110-0220301210021000-3033312332203233"></a>

## Next pages — no_port_match / 001303002311 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-028.md#canonical-1303023022032323-0011211310021103-2302310103100323-1020331133202010-1211220133311022-0313230133233210-0330011000132000-0200302302220132)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3012300121100130-2123213313031300-1111323022000031-3200021200320130-3002333023012223-3010202011010023-0133110113330213-0102003333203322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302010203011102-0231200322303312-0300321210003111-0222212022331322-3111030030222202-0312201022133100-2003020101101232-2032212323111222"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path — path / 130233112111 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-027.md#canonical-0330013222233302-2210122030021330-0131201133002112-2132231131331102-2322210303201110-1121010222301222-3302213332031013-0122333212301102)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-2110002022213231-0110200102113003-2213030123311011-3031131131130221-1000123200332201-2333033313333100-1323200012300312-2110331002310100"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-1100110122212332-1313112011111111-3000001322321013-3023003211021230-0223333113111101-3230021300132203-2201022102120002-0220320201002211"></a>

## Direct properties — path / 130233112111 / 3

<a id="canonical-2303221123011010-0222301203312020-3033110103330321-3113203023202201-2231301012122111-1133231100033000-0122020330030201-2210021211011001"></a>

<a id="canonical-0231320223222323-0330221332313201-1012313130010001-3300300002323332-2200131122011021-0023221122201033-2200033122233323-2312303222311131"></a>

## path property — path / 130233112111 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-3010331133321002-0110130233002230-2200203111303312-3220221123222012-0231303230101302-1033213001310312-1232232303200020-1233031103233322"></a>

<a id="canonical-2023203132202020-2231130121313120-2002203213021122-3000201313111200-0121121222222301-0020300312310021-3233332020003100-2012013032331131"></a>

## prefix property — path / 130233112111 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-2220031032200302-2221001210331022-0013223322322033-2103333112220012-0301120020333230-1220131210032312-3011301010301112-0310132012001111"></a>

<a id="canonical-0100302010032321-0020003322211113-3133121010130210-1022032011311003-1233302203030203-3120333031110130-1302210231310232-2010113001201311"></a>

## regex property — path / 130233112111 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0030322322121031-3032003213310203-1333320110130220-3110013003303010-1100201300301230-2033320010002221-3021333312202213-1222223113121023"></a>

## Next pages — path / 130233112111 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-027.md#canonical-0330013222233302-2210122030021330-0131201133002112-2132231131331102-2322210303201110-1121010222301222-3302213332031013-0122333212301102)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3302110132323220-0130312011310200-3213310121101213-2332012011222213-2322213101231122-2221221001002212-3303021101022230-2212202010103100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132020020130023-2132101112233201-0010103112303230-1231100330302222-1220110120302321-0223133220321113-2130230333311032-0123010122020031"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect — route_redirect / 222212013203 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-027.md#canonical-0330013222233302-2210122030021330-0131201133002112-2132231131331102-2322210303201110-1121010222301222-3302213332031013-0122333212301102)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-3201011132131102-2001032013323030-2120130232113221-0101013131320111-0213210111310213-2311112201022333-1133012320212021-0300102003301200"></a>

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

<a id="canonical-0103000203030211-2002303221210210-0232022121102022-0202222121300101-3000223220011020-0203020222202100-0220022031001331-0113312122201032"></a>

## Direct properties — route_redirect / 222212013203 / 3

<a id="canonical-3301331312010010-3322322101112113-1303103200002121-3212022213323220-3023001003230233-0312332032033220-2211010122110012-1330223302033310"></a>

<a id="canonical-0321221211101311-0010032102003032-3202002331311220-1113030313330033-3123222120103230-2132202331111020-0023031212221323-2300203032020012"></a>

## host_redirect property — route_redirect / 222212013203 / 4

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

<a id="canonical-1121100030102212-3311020332001021-2112332101232232-0110023301021013-2320000332333213-1312201223323033-3100233200101110-2210000113300222"></a>

<a id="canonical-0133121123030203-1230032100230132-3313222221000301-3200213200133302-3011100012103121-0212023321001220-0002201111123133-0223101102023300"></a>

## path_redirect property — route_redirect / 222212013203 / 5

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

<a id="canonical-3332022021311002-2111313231002022-0202120021301100-3221323332212212-3033223332302110-0123011000200001-1030321022321311-2330230022331201"></a>

<a id="canonical-3001320013311031-0012122310232031-3032011211311200-0002120023020212-0131033130120103-3233101030023130-2322213303001120-1003011332323130"></a>

## prefix_rewrite property — route_redirect / 222212013203 / 6

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

<a id="canonical-1023113220303232-2031112320333020-2203023101020031-0102221030013010-2022131002023030-1100220133101121-0230312302322030-2121231111023102"></a>

<a id="canonical-3031321212213021-0220200220020111-1111203013200310-0323220211022123-3132221313222103-0113022320002102-1120232003030031-3000300231231201"></a>

## proto_redirect property — route_redirect / 222212013203 / 7

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

- [remove_all_params](resources--workload--reference--group-028.md#canonical-3022302201200233-3230121112131102-3033302012210102-1132033121003310-0321123121011023-2220011131102301-3221003232313002-0110223211023003): complete subsection reference.

<a id="canonical-1031020012213321-0023302333123012-0031022320102213-2313030310300222-0002333110311022-0000000113001103-2031302120322313-2031311221203120"></a>

<a id="canonical-3200220022222132-1212330113221201-2023121032131120-3210330122023133-2201100331300202-0203120312132002-2123210020322232-1331210322301130"></a>

## replace_params property — route_redirect / 222212013203 / 8

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

<a id="canonical-1233330323311023-3320122010110320-3013103311212001-1110233023301013-2023302001213110-2132102331130101-1000332210302310-1220231301130211"></a>

<a id="canonical-3033301011303220-1132211312031011-0111021113000203-0113330003000232-1013201201322110-0331331202021132-1200231123333221-1132000111303210"></a>

## response_code property — route_redirect / 222212013203 / 9

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

- [retain_all_params](resources--workload--reference--group-028.md#canonical-2111112002121012-0111132023300111-1303021012101121-0320331000102103-0112020202203322-0130322022020300-1213223100221333-2330202231122301): complete subsection reference.

<a id="canonical-1011221020132133-0221111321033002-1310310131223112-2230103122201103-0233333033322102-2002123231221323-1223101102123211-3221310100212200"></a>

## Next pages — route_redirect / 222212013203 / 10

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params](resources--workload--reference--group-028.md#canonical-3022302201200233-3230121112131102-3033302012210102-1132033121003310-0321123121011023-2220011131102301-3221003232313002-0110223211023003)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params](resources--workload--reference--group-028.md#canonical-2111112002121012-0111132023300111-1303021012101121-0320331000102103-0112020202203322-0130322022020300-1213223100221333-2330202231122301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-027.md#canonical-0330013222233302-2210122030021330-0131201133002112-2132231131331102-2322210303201110-1121010222301222-3302213332031013-0122333212301102)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3022302201200233-3230121112131102-3033302012210102-1132033121003310-0321123121011023-2220011131102301-3221003232313002-0110223211023003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321113110013313-2130301120231231-1211032010002030-1201121130233220-3212221333010103-0200333022131130-0222131102103022-3011231301001013"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params — remove_all_params / 311211121300 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-027.md#canonical-0330013222233302-2210122030021330-0131201133002112-2132231131331102-2322210303201110-1121010222301222-3302213332031013-0122333212301102)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-028.md#canonical-3302110132323220-0130312011310200-3213310121101213-2332012011222213-2322213101231122-2221221001002212-3303021101022230-2212202010103100)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-2110223021323332-2310230023302000-2112231032130332-3202312201201321-1023023332212203-1132323103101310-3301002131301003-1012122130013021"></a>

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

<a id="canonical-1333030032212001-3021200023010021-1131102233323221-0022320000201302-2323020201300231-1100323011311332-1320333032102021-1133021203220300"></a>

## Direct properties — remove_all_params / 311211121300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003120013001320-3212002012002020-0223230120001013-1010230201331023-1323220022011221-3323322211013202-0021230100302130-3011312111221020"></a>

## Next pages — remove_all_params / 311211121300 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-028.md#canonical-3302110132323220-0130312011310200-3213310121101213-2332012011222213-2322213101231122-2221221001002212-3303021101022230-2212202010103100)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2111112002121012-0111132023300111-1303021012101121-0320331000102103-0112020202203322-0130322022020300-1213223100221333-2330202231122301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301222321302133-0031101212030323-3330330321131102-1112120020232122-0310002022200201-3200310101023221-2012123312210011-2210220301010032"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params — retain_all_params / 212001000001 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-027.md#canonical-0330013222233302-2210122030021330-0131201133002112-2132231131331102-2322210303201110-1121010222301222-3302213332031013-0122333212301102)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-028.md#canonical-3302110132323220-0130312011310200-3213310121101213-2332012011222213-2322213101231122-2221221001002212-3303021101022230-2212202010103100)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-2203012131333332-2001010102202011-1111113223000212-0322011222320312-3213122332302031-1330111030303002-0211011020033222-0302201302230020"></a>

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

<a id="canonical-0111233202211132-0022322211011212-0313332033132031-0320031020313102-3120033110323231-1120321012013113-1100232221313221-0020031023331100"></a>

## Direct properties — retain_all_params / 212001000001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010022110020321-1102023132311323-1323022021123011-3011223101230132-0113213211212113-1323210302312001-0112212123211311-0200330010032230"></a>

## Next pages — retain_all_params / 212001000001 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-028.md#canonical-3302110132323220-0130312011310200-3213310121101213-2332012011222213-2322213101231122-2221221001002212-3303021101022230-2212202010103100)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3223313232022132-1300023030303322-3120001113102132-2101020221132002-0011333231033012-1032011333210012-0303212231230122-1132022333212200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013012213313132-0312023312123100-3101001333313333-3001222021011032-3010100221323121-0101202330120101-0121133330230213-2012002030302200"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route — simple_route / 123223003300 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-3220131130020110-3002022132020200-2213123020123211-0003001323331100-3202312021330332-3202301101313313-0213131010311033-2313320123133321"></a>

Type: `"object"`. single nested block, Optional.

Simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

Upstream description:

A simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
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
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
simple_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323120033201211-0122212123223300-2322130302320221-2220233201221021-2302010320112030-1033113022031110-2123222303033231-2310120110001312"></a>

## Direct properties — simple_route / 123223003300 / 3

- [auto_host_rewrite](resources--workload--reference--group-028.md#canonical-0212220221111023-1000023213211100-3010222331203110-2120131200321123-0312032321201232-1111033102332123-1111121221230001-1321212321130012): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-028.md#canonical-3201103303230233-1311003111323331-3201331333211200-2033111332103310-0101023203110110-3332121123222232-1103002012213120-0010321001223021): complete subsection reference.

<a id="canonical-1301130322320332-3300200300121100-1131331121133210-1312323212010130-2212210332002213-2333022100232112-0021300203310112-0032210330231023"></a>

<a id="canonical-1310133202102013-0033223222322201-2132133221033113-0103303032013320-3201132300120302-0112231333011111-1013201200320203-2110133102322321"></a>

## host_rewrite property — simple_route / 123223003300 / 4

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-1013133322133001-1001022202113023-0101102033022211-3003213322303123-2111110020110132-0032311121032120-0101021333310232-1131102200030110"></a>

<a id="canonical-1223120333111333-2222032200333232-3120033323013201-0000321331333302-2221132310331112-2021100203103212-3321012322010201-0222102021333222"></a>

## http_method property — simple_route / 123223003300 / 5

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [path](resources--workload--reference--group-028.md#canonical-3021333212112313-2101331311213202-0232212101132113-3023302321322000-1220322010120102-3322031001132322-3000132030301332-0020122313213010): complete subsection reference.

<a id="canonical-1230330010313221-2302230131300120-2301001103021013-2231303212101033-0002230201003313-3333200220000033-2001133211101313-1020020021102121"></a>

## Next pages — simple_route / 123223003300 / 6

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](resources--workload--reference--group-028.md#canonical-0212220221111023-1000023213211100-3010222331203110-2120131200321123-0312032321201232-1111033102332123-1111121221230001-1321212321130012)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](resources--workload--reference--group-028.md#canonical-3201103303230233-1311003111323331-3201331333211200-2033111332103310-0101023203110110-3332121123222232-1103002012213120-0010321001223021)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path](resources--workload--reference--group-028.md#canonical-3021333212112313-2101331311213202-0232212101132113-3023302321322000-1220322010120102-3322031001132322-3000132030301332-0020122313213010)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0212220221111023-1000023213211100-3010222331203110-2120131200321123-0312032321201232-1111033102332123-1111121221230001-1321212321130012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331023332012323-2133030310003223-1231010331223031-1232203012321131-2321222121233333-2001333111120012-2303032120122123-2312100021120203"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite — auto_host_rewrite / 321233132130 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-028.md#canonical-3223313232022132-1300023030303322-3120001113102132-2101020221132002-0011333231033012-1032011333210012-0303212231230122-1132022333212200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-0221103303322032-1323133333203112-0203312122100112-1331121213202101-0131231203232113-0320021322101130-3203213311112311-1323103302322001"></a>

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
auto_host_rewrite = {}
```

<a id="canonical-2123101302110003-2111133310313330-0201221320002012-3311221332302010-1302001001330020-0110021000231011-2210332222202001-3021021100202221"></a>

## Direct properties — auto_host_rewrite / 321233132130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233200322222132-1311011331232013-2110012011020001-2102313013322021-3330110130010011-3211201201101211-1310133332233303-1100300313000110"></a>

## Next pages — auto_host_rewrite / 321233132130 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-028.md#canonical-3223313232022132-1300023030303322-3120001113102132-2101020221132002-0011333231033012-1032011333210012-0303212231230122-1132022333212200)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3201103303230233-1311003111323331-3201331333211200-2033111332103310-0101023203110110-3332121123222232-1103002012213120-0010321001223021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133221332203203-0203102121301233-2102023022010001-2121210223300000-1313221232221221-0121113110301231-0020001303110230-1120103133330232"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite — disable_host_rewrite / 322101212231 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-028.md#canonical-3223313232022132-1300023030303322-3120001113102132-2101020221132002-0011333231033012-1032011333210012-0303212231230122-1132022333212200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-1302022122012203-0131322123323220-2330210323202011-3202011210100032-3022131310303133-2131123120003130-2010113112121212-0003111123213230"></a>

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
disable_host_rewrite = {}
```

<a id="canonical-2322131122210123-2101112001022122-3323023213013221-0131122123123023-2120303310302133-0123332303222313-3001132311020132-0000321021222312"></a>

## Direct properties — disable_host_rewrite / 322101212231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203012320223301-2021000031230223-1222202131120300-1023032211001301-2220112020222100-0100333113012322-3112210301200300-0233330331121222"></a>

## Next pages — disable_host_rewrite / 322101212231 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-028.md#canonical-3223313232022132-1300023030303322-3120001113102132-2101020221132002-0011333231033012-1032011333210012-0303212231230122-1132022333212200)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3021333212112313-2101331311213202-0232212101132113-3023302321322000-1220322010120102-3322031001132322-3000132030301332-0020122313213010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100201222022320-2213330032311103-2201321303112002-2023132222302210-3332201233221210-0332103231301101-1132330232022310-2333220312211131"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path — path / 320201032031 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-028.md#canonical-3223313232022132-1300023030303322-3120001113102132-2101020221132002-0011333231033012-1032011333210012-0303212231230122-1132022333212200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-0310013323322201-1202302121113021-1300200201013120-0201210320313022-0101012213312010-1122312131132302-1122103100233223-3022001010113202"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001322312032001-1223013230323030-3203232332323221-1131330332223013-0311100303300321-0310331201221110-0101132212121022-0111100312103020"></a>

## Direct properties — path / 320201032031 / 3

<a id="canonical-0102203013100130-1223100310023132-1022212223310223-3020020301221022-0331331102333010-0331023000220221-0103202012031303-2301301222102020"></a>

<a id="canonical-0213212200003202-1023311131101330-0302103132233013-3303132130120210-1003113232321003-0113001103222013-0022000002221022-1322030202232331"></a>

## path property — path / 320201032031 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-3130020022112022-1001330102122313-0220332113232103-1031123101021300-2010000000101110-3123101010013010-0220332332113232-2102011303121323"></a>

<a id="canonical-1113320103020111-3130012123000200-1022013230202000-2102233112320131-3222322313003323-0020333013111300-2112032120002131-2232111220212223"></a>

## prefix property — path / 320201032031 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-0321332322203211-0122323313332021-0210000311013321-1322322213232203-0312032212201321-1212322211303333-3321311122102030-3103312102321220"></a>

<a id="canonical-3120133311130213-2323200223112313-1312202033032111-2111001010200020-0132200330130000-0322001200320211-0033002123110231-1232230011123131"></a>

## regex property — path / 320201032031 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2300021312211233-2313303323232122-3302023131231321-0312320123010320-3012021132120113-1131211030010030-2020033221123000-2132303101021102"></a>

## Next pages — path / 320201032031 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-028.md#canonical-3223313232022132-1300023030303322-3120001113102132-2101020221132002-0011333231033012-1032011333210012-0303212231230122-1132022333212200)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2323101323000230-3130303013113131-1312013231212223-1212222320213220-2013000033112020-2320301032133202-3111123300030002-2013131210032122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223212202130331-3323003120002300-0221133230321230-0300100020333323-3300330012011022-3000103231100001-1211103212310023-3223122202011000"></a>

## stateful_service.advertise_options.advertise_on_public.port.port — port / 022120030322 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- stateful_service.advertise_options.advertise_on_public.port.port

<a id="canonical-1131031200011330-2003320130233000-2310101133021033-2302211303123330-2002100332333100-0222332012221222-3213002133233322-2130331123333212"></a>

Type: `"object"`. single nested block, Optional.

Port. Single port.

Upstream description:

Single port.

Receipt-pinned upstream constraints:

```json
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
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201321213221201-0311103133211033-0332322312323131-1321210113022223-0231002312203023-0013313213021210-3121133010011301-1111323103303233"></a>

## Direct properties — port / 022120030322 / 3

- [info](resources--workload--reference--group-028.md#canonical-2322021320231203-3320131000003231-1103323320202131-3000213021202011-0332120303031122-2330331300030221-3032131102120330-2202233022300200): complete subsection reference.

<a id="canonical-0102220213310133-0023030313332122-2032331231103312-2320212110201130-2301322103300230-1223121030131000-1101003221231300-3221011232131330"></a>

## Next pages — port / 022120030322 / 4

- [stateful_service.advertise_options.advertise_on_public.port.port.info](resources--workload--reference--group-028.md#canonical-2322021320231203-3320131000003231-1103323320202131-3000213021202011-0332120303031122-2330331300030221-3032131102120330-2202233022300200)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2322021320231203-3320131000003231-1103323320202131-3000213021202011-0332120303031122-2330331300030221-3032131102120330-2202233022300200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221011213200023-1102333303300313-0221020212300200-0231331132000233-2103033221031221-0021121112232031-3023011220010313-3220030122023131"></a>

## stateful_service.advertise_options.advertise_on_public.port.port.info — info / 211321121330 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-028.md#canonical-2323101323000230-3130303013113131-1312013231212223-1212222320213220-2013000033112020-2320301032133202-3111123300030002-2013131210032122)
- stateful_service.advertise_options.advertise_on_public.port.port.info

<a id="canonical-1003200323012333-3123333100232332-0221302030111010-1122030020112123-2111323020001212-2332020222010003-0320231121331300-0202201002111130"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

Upstream description:

Port information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port"),
  validators.ConflictingObjectAttributes("same_as_port",
    "target_port")}
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
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

Terraform syntax:

```terraform
info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013100022101233-0031003110301321-2021333303030210-2312131300102301-0331231130323312-3312122010133130-1221333301030233-1330200132101332"></a>

## Direct properties — info / 211321121330 / 3

<a id="canonical-2130233332120323-2122131012111103-1210300211000022-3332230100113013-2002312130330033-1221120303203303-1022130221321032-0033332212131113"></a>

<a id="canonical-3022320030102103-0020031210033332-3111322333231101-0003323213313330-0213001110213030-0133013020333032-1000233002103002-1013000202303033"></a>

## port property — info / 211321121330 / 4

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2032230310010231-0012200323011202-0220320313023101-0003102032213212-3321003002333110-3111231212211022-3112121103121001-3313223020131202"></a>

<a id="canonical-0001112321033233-1302100033320230-3001021203222113-1222102132101111-3123130323213000-0213112202111231-3013101222300202-0313330220210122"></a>

## protocol property — info / 211321121330 / 5

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

TCP &#8203;- PROTOCOL\_HTTP: HTTP

HTTP &#8203;- PROTOCOL\_HTTP2: HTTP2

HTTP2 &#8203;- PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI

TLS with SNI &#8203;- PROTOCOL\_UDP: UDP

UDP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
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

- [same_as_port](resources--workload--reference--group-028.md#canonical-2321001300010003-1102133001310130-2102300131332133-0012202022311302-1023033220032013-1331030230213323-3132030033000222-0302013110321223): complete subsection reference.

<a id="canonical-1302320102002022-0131012232103000-1203033133022133-2113012321323021-2321310223313232-3121333222033301-1201232130313301-1313303013230301"></a>

<a id="canonical-0100110111102001-0222133301321111-1211010320003020-0211200102202312-3103133212030000-3200121203003232-3300110222230011-3212001130300333"></a>

## target_port property — info / 211321121330 / 6

Type: `"number"`. Optional.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1100010300230323-3213312222121213-1230002323222102-1302300013023331-0002020113221223-0010001132002323-0223133100321132-2301223331120222"></a>

## Next pages — info / 211321121330 / 7

- [stateful_service.advertise_options.advertise_on_public.port.port.info.same_as_port](resources--workload--reference--group-028.md#canonical-2321001300010003-1102133001310130-2102300131332133-0012202022311302-1023033220032013-1331030230213323-3132030033000222-0302013110321223)
- [stateful_service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-028.md#canonical-2323101323000230-3130303013113131-1312013231212223-1212222320213220-2013000033112020-2320301032133202-3111123300030002-2013131210032122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2321001300010003-1102133001310130-2102300131332133-0012202022311302-1023033220032013-1331030230213323-3132030033000222-0302013110321223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001022212102210-1003021123233031-0212110223320203-0013112023222130-0113033321100312-3020032002022333-1311310201120023-2210220102112333"></a>

## stateful_service.advertise_options.advertise_on_public.port.port.info.same_as_port — same_as_port / 301002103303 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-028.md#canonical-2323101323000230-3130303013113131-1312013231212223-1212222320213220-2013000033112020-2320301032133202-3111123300030002-2013131210032122)
- [stateful_service.advertise_options.advertise_on_public.port.port.info](resources--workload--reference--group-028.md#canonical-2322021320231203-3320131000003231-1103323320202131-3000213021202011-0332120303031122-2330331300030221-3032131102120330-2202233022300200)
- stateful_service.advertise_options.advertise_on_public.port.port.info.same_as_port

<a id="canonical-2313130303311000-3211112201320300-2210211200333123-0000002212022323-0002031002133011-0023021323032210-3003321203121000-0213203130030321"></a>

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
same_as_port = {}
```

<a id="canonical-2321322003322133-1113301001323132-3313001233031110-3312322332210313-1133030230323012-1200320031310131-1321303133230121-0003102321110233"></a>

## Direct properties — same_as_port / 301002103303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201210101112132-2130312022021220-2101313033202312-3133020322311203-2011020023231201-2031033220320000-1132312320100201-1231313212002330"></a>

## Next pages — same_as_port / 301002103303 / 4

- [stateful_service.advertise_options.advertise_on_public.port.port.info](resources--workload--reference--group-028.md#canonical-2322021320231203-3320131000003231-1103323320202131-3000213021202011-0332120303031122-2330331300030221-3032131102120330-2202233022300200)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0201031010010211-0110122211331300-2100213123230032-2210131211202000-0123012121332012-2132032030312023-0111231211021031-2320223020033130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032133002213222-3122123222023331-3203211322122220-0103212233021002-0230102222123301-1022323033030113-3022011032123201-3033021121021230"></a>

## stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer — tcp_loadbalancer / 012221302320 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer

<a id="canonical-2002323122321000-2112022202131013-0020210332321210-2110030111322210-0230333322111022-2303230321311131-3000022230202211-0112030103210210"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tcp loadbalancer.

Upstream description:

TCP loadbalancer.

Receipt-pinned upstream constraints:

```json
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
tcp_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331112030003310-0220123303200133-2131001123103321-1132332133223030-2100333233000013-1231312331023111-3213220033131132-3300033332021300"></a>

## Direct properties — tcp_loadbalancer / 012221302320 / 3

<a id="canonical-1301230131031211-1002122223323222-1020200132030311-1012113121300110-3031230310310131-3331030332311000-0331332220001120-3030322330112030"></a>

<a id="canonical-0001001000032331-3123010301100301-0201022331202310-3110223301200302-3032213021211001-1013232102033100-0030233312230322-3032001221130011"></a>

## domains property — tcp_loadbalancer / 012221302320 / 4

Type: `["list", "string"]`. Optional.

List of additional domains (host/authority header) that will be matched to this loadbalancer.
Domains are also used for SNI matching if the is true Domains also indicate the list of names for
which DNS resolution will be done by VER.

Upstream description:

A list of additional domains (host/authority header) that will be matched to this loadbalancer.

Domains are also used for SNI matching if the \`with\_sni\` is true Domains also indicate the list
of names for which DNS resolution will be done by VER.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2320313100122030-1230022111301233-0102031311332303-3002013303111302-2313303230010022-1323112332003230-3012230300101132-0020201133031222"></a>

<a id="canonical-0133210102100221-0121210310311221-2210021013322031-1030322000003301-2030233021131021-2230202211112112-1010000233200333-3033222122322202"></a>

## with_sni property — tcp_loadbalancer / 012221302320 / 5

Type: `"bool"`. Optional.

Set to true to enable TCP loadbalancer with SNI.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3302000201113231-3013310113220021-3311032220021123-2020222002122203-0211112111211331-1221023121123311-1132013231213130-3100130210131202"></a>

## Next pages — tcp_loadbalancer / 012221302320 / 6

- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3202321232103323-1000231221221030-0001122222031132-1003133321010023-1002201331103010-1230130022220311-0032100201113003-0333121223233010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031203200120102-1323233301221231-0033233301303003-3021022231333222-1012200213222331-2012203312330010-3002100230102012-2023230021312132"></a>

## stateful_service.advertise_options.do_not_advertise — do_not_advertise / 302132320322 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- stateful_service.advertise_options.do_not_advertise

<a id="canonical-3002301133331320-1322203131133310-1112020032333111-0133302112022120-0233211032202220-0122333032100300-2110130011122200-1221220222220303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise.

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
do_not_advertise = {}
```

<a id="canonical-0301312023303131-2103013013003232-0322032202230133-3211101030321030-1123321002330312-3222213223033210-2013330300310332-0211002231130312"></a>

## Direct properties — do_not_advertise / 302132320322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232111321130313-0112022030320010-2233113213011133-0230200322030102-3233203113321230-2132120133201031-1320022213012032-1332012321200113"></a>

## Next pages — do_not_advertise / 302132320322 / 4

- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3203012310220003-2210330300331110-2103223003030222-2222201323310103-2300110002113011-0130333021003321-0201020033321220-1102301103203230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101202312333220-2211200333103323-3221321230223013-3003313332022001-1200012223110010-0030111101023122-3322333232202020-0322003100320102"></a>

## stateful_service.configuration — configuration / 103313011113 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- stateful_service.configuration

<a id="canonical-1201322303303202-0221320010021012-1303000331110131-0101010002332012-2322111220213321-1101133003320203-1013112320231110-0212330121232013"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameters of the workload.

Receipt-pinned upstream constraints:

```json
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
configuration {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220103012311213-3102022020310201-3311200211112010-1233322213123203-3222010100031302-0033122122310120-0031112022132332-2210021303322233"></a>

## Direct properties — configuration / 103313011113 / 3

- [parameters](resources--workload--reference--group-028.md#canonical-1302213033220212-3202132112103233-0311330233023033-0321030031223222-2102010112112002-2320312322333220-0032233021200100-1022102313033123): complete subsection reference.

<a id="canonical-3013012011300323-1322121012230321-2201122230001322-2322033030211322-2320303021103001-2332201230100232-1030330023110310-0313003333303201"></a>

## Next pages — configuration / 103313011113 / 4

- [stateful_service.configuration.parameters](resources--workload--reference--group-028.md#canonical-1302213033220212-3202132112103233-0311330233023033-0321030031223222-2102010112112002-2320312322333220-0032233021200100-1022102313033123)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1302213033220212-3202132112103233-0311330233023033-0321030031223222-2102010112112002-2320312322333220-0032233021200100-1022102313033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332110010001010-2013313311220303-3202133223200100-2202303301013330-0102032213323300-2000021002321300-1023233113132200-2213210011233132"></a>

## stateful_service.configuration.parameters — parameters / 212030030032 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.configuration](resources--workload--reference--group-028.md#canonical-3203012310220003-2210330300331110-2103223003030222-2222201323310103-2300110002113011-0130333021003321-0201020033321220-1102301103203230)
- stateful_service.configuration.parameters

<a id="canonical-0333303303203013-2021330333311330-0211211331210301-1220230230230102-1013220331210112-1110310111112203-1213212300210113-0003113111231111"></a>

Type: `"object"`. list nested block, Optional.

Parameters. Parameters for the workload.

Upstream description:

Parameters for the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("env_var",
    "file")}
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
parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220312113113113-2001212200203312-2221013332032330-2023033010012312-0321032111220232-0031331223101321-1210111110012202-3131002013022203"></a>

## Direct properties — parameters / 212030030032 / 3

- [env_var](resources--workload--reference--group-028.md#canonical-1323013220320120-3113201120222231-3323230231223312-3132031213122030-0230203302303023-0002322101130002-2012310022211131-2033101310312030): complete subsection reference.

- [file](resources--workload--reference--group-028.md#canonical-3111133103101303-0323302220303200-2031333233121200-0222313132030022-3332112322312012-1320033010330232-0302211321120013-3033112100233210): complete subsection reference.

<a id="canonical-2021031223103233-3331302332023213-3202211312023120-0232023103033131-3313022112321303-3201002310200000-3212012302023201-1002323002003323"></a>

## Next pages — parameters / 212030030032 / 4

- [stateful_service.configuration.parameters.env_var](resources--workload--reference--group-028.md#canonical-1323013220320120-3113201120222231-3323230231223312-3132031213122030-0230203302303023-0002322101130002-2012310022211131-2033101310312030)
- [stateful_service.configuration.parameters.file](resources--workload--reference--group-028.md#canonical-3111133103101303-0323302220303200-2031333233121200-0222313132030022-3332112322312012-1320033010330232-0302211321120013-3033112100233210)
- [stateful_service.configuration](resources--workload--reference--group-028.md#canonical-3203012310220003-2210330300331110-2103223003030222-2222201323310103-2300110002113011-0130333021003321-0201020033321220-1102301103203230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1323013220320120-3113201120222231-3323230231223312-3132031213122030-0230203302303023-0002322101130002-2012310022211131-2033101310312030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020312123011131-2300331300210112-0112330203310230-1222203312201120-1301313110223211-3013223220201110-2332000031220123-2123232231000001"></a>

## stateful_service.configuration.parameters.env_var — env_var / 100332012112 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.configuration](resources--workload--reference--group-028.md#canonical-3203012310220003-2210330300331110-2103223003030222-2222201323310103-2300110002113011-0130333021003321-0201020033321220-1102301103203230)
- [stateful_service.configuration.parameters](resources--workload--reference--group-028.md#canonical-1302213033220212-3202132112103233-0311330233023033-0321030031223222-2102010112112002-2320312322333220-0032233021200100-1022102313033123)
- stateful_service.configuration.parameters.env_var

<a id="canonical-0222323200322103-1001330330310020-2223213220122033-0312213222000102-2220311212203301-1313212103302002-0322202321313222-0202323113132010"></a>

Type: `"object"`. single nested block, Optional.

Environment Variable. Environment Variable.

Upstream description:

Environment Variable.

Receipt-pinned upstream constraints:

```json
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
env_var {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310112321302022-0322131203330230-3303032032233223-3313020000000201-0031020131301020-1101003112103321-3203012310201033-2200000231211302"></a>

## Direct properties — env_var / 100332012112 / 3

<a id="canonical-3222100201300110-1130032302211021-2102200003012220-0013002332221001-1231100211012313-3131321100320010-1221203201011112-0022301132202233"></a>

<a id="canonical-2223010302321130-1303103000233222-3002323322000301-1333320130110221-3123022113322001-0131130130023012-0030323120202322-1231320133232331"></a>

## name property — env_var / 100332012112 / 4

Type: `"string"`. Optional.

Name. Name of Environment Variable.

Upstream description:

Name of Environment Variable.

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

<a id="canonical-0332132212021231-2033021132011311-0000202101331333-3303302033331310-2013233012021300-1011230132310033-0202212131213213-2110112201302011"></a>

<a id="canonical-0100203122012100-3223011011112022-3032312113012300-3033230021110321-2323122113101112-0201021103121001-3303003010013131-3012002330010120"></a>

## value property — env_var / 100332012112 / 5

Type: `"string"`. Optional.

Value. Value of Environment Variable.

Upstream description:

Value of Environment Variable.

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

<a id="canonical-2121303301001021-3221222101213133-2312023222023300-3203113311201332-1212013202201332-3132012210121202-1101300130033210-1232212011131010"></a>

## Next pages — env_var / 100332012112 / 6

- [stateful_service.configuration.parameters](resources--workload--reference--group-028.md#canonical-1302213033220212-3202132112103233-0311330233023033-0321030031223222-2102010112112002-2320312322333220-0032233021200100-1022102313033123)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3111133103101303-0323302220303200-2031333233121200-0222313132030022-3332112322312012-1320033010330232-0302211321120013-3033112100233210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121301021132111-3203011123333011-3332331121322311-3131210110221130-1311210122122222-0310111212232122-2002011302332002-0230322022322002"></a>

## stateful_service.configuration.parameters.file — file / 003102002222 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.configuration](resources--workload--reference--group-028.md#canonical-3203012310220003-2210330300331110-2103223003030222-2222201323310103-2300110002113011-0130333021003321-0201020033321220-1102301103203230)
- [stateful_service.configuration.parameters](resources--workload--reference--group-028.md#canonical-1302213033220212-3202132112103233-0311330233023033-0321030031223222-2102010112112002-2320312322333220-0032233021200100-1022102313033123)
- stateful_service.configuration.parameters.file

<a id="canonical-0020013323122130-3121213332131131-1201313310110320-3100321010323213-0131121232132012-3131122213103311-0321012310213221-2011021002111231"></a>

Type: `"object"`. single nested block, Optional.

Configuration File. Configuration File for the workload.

Upstream description:

Configuration File for the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name",
    "volume_name")}
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
file {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000303021211022-3301133000112123-2231230102132211-3112101010011111-3200220123132203-1012312333201122-1012303330012313-0133031333010202"></a>

## Direct properties — file / 003102002222 / 3

<a id="canonical-0203223030221121-2121202222200002-1233201032223132-2033103132103213-1203001310112123-0211231132221030-2131003000033103-2032110113000032"></a>

<a id="canonical-0333130101010120-2011033110033132-1302022222131223-3101202122101221-2203123132300301-1032021023300221-0023131320010120-1223100001022321"></a>

## data property — file / 003102002222 / 4

Type: `"string"`. Optional.

Data. File data

Upstream description:

File data

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(16384),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16384,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 16384,
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
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [mount](resources--workload--reference--group-028.md#canonical-0233000201230203-0013120130001113-1020123313011012-3331231213030323-1031331300313332-1022000203000000-1200013002010223-1333333023011110): complete subsection reference.

<a id="canonical-0112000310120033-3010033122211331-2221033233201121-2000303021222011-0003121003123032-3230030011331330-3020223203233221-0200322123023311"></a>

<a id="canonical-0101312011030013-1332003303001233-3332013230333123-2322100321203333-0130031122102131-2320021122223120-1212123220121113-3332302303103113"></a>

## name property — file / 003102002222 / 5

Type: `"string"`. Optional.

Name. Name of the file.

Upstream description:

Name of the file.

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

<a id="canonical-1110122130212220-3323121330233113-0323300323301321-2231021321210230-0001320213231033-3220033211102002-3001221000023113-3033222101211100"></a>

<a id="canonical-2222221012303330-3331311220031323-0002101011311101-3212213200232021-2031231302330203-0133032112030111-2121232201330121-1120032231100000"></a>

## volume_name property — file / 003102002222 / 6

Type: `"string"`. Optional.

Volume Name. Name of the Volume.

Upstream description:

Name of the Volume.

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

<a id="canonical-0013110021310312-3223311220230331-2212130233223030-3331031230330222-0323301330331100-0123030331132322-3020110111312311-0203131033131033"></a>

## Next pages — file / 003102002222 / 7

- [stateful_service.configuration.parameters.file.mount](resources--workload--reference--group-028.md#canonical-0233000201230203-0013120130001113-1020123313011012-3331231213030323-1031331300313332-1022000203000000-1200013002010223-1333333023011110)
- [stateful_service.configuration.parameters](resources--workload--reference--group-028.md#canonical-1302213033220212-3202132112103233-0311330233023033-0321030031223222-2102010112112002-2320312322333220-0032233021200100-1022102313033123)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0233000201230203-0013120130001113-1020123313011012-3331231213030323-1031331300313332-1022000203000000-1200013002010223-1333333023011110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133111230301231-1112301120323003-1302001200203301-1332001011201012-0210232323021222-3201131123130113-1233130200102232-1202010312222211"></a>

## stateful_service.configuration.parameters.file.mount — mount / 022301111002 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.configuration](resources--workload--reference--group-028.md#canonical-3203012310220003-2210330300331110-2103223003030222-2222201323310103-2300110002113011-0130333021003321-0201020033321220-1102301103203230)
- [stateful_service.configuration.parameters](resources--workload--reference--group-028.md#canonical-1302213033220212-3202132112103233-0311330233023033-0321030031223222-2102010112112002-2320312322333220-0032233021200100-1022102313033123)
- [stateful_service.configuration.parameters.file](resources--workload--reference--group-028.md#canonical-3111133103101303-0323302220303200-2031333233121200-0222313132030022-3332112322312012-1320033010330232-0302211321120013-3033112100233210)
- stateful_service.configuration.parameters.file.mount

<a id="canonical-3222300012011013-3222112022030031-0333003331130000-2003211133012212-1331102322213200-0221012030311100-3120102101322322-1213110130012231"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010210003102222-0102222213323010-2100322233300121-1311323232231303-2330012330303101-0230133222321222-2333011232113322-1001212231011322"></a>

## Direct properties — mount / 022301111002 / 3

<a id="canonical-2322311313103131-0101221031322302-3232233102110110-1001010021312101-0212133310022003-3301312102022112-2002122131021111-1210230132133110"></a>

<a id="canonical-2022200001002102-3032002203003311-2010220331111232-0113112021331110-1121333111203120-2010300020032301-0210130003220110-2212023023311203"></a>

## mode property — mount / 022301111002 / 4

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3031321122023211-2211202131311023-0223203332202131-3032112021230313-1212110323111112-1012200221120121-2312323132212233-1230310101222201"></a>

<a id="canonical-2321221131002033-1233013111021223-3313113030012010-0302123030310001-1223232100301121-2212331023103031-3232011311220322-0231211101000203"></a>

## mount_path property — mount / 022301111002 / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
    },
    "pattern": "^[^:]*$"
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
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-3312322121030111-2300020310330113-1113101332223203-2112310233010112-1112301303032300-2110222221021323-2230233323030312-3300300103123303"></a>

<a id="canonical-2021232313213230-0231110123312132-3223301000123002-3100033011320333-2033110021102012-1201022231131023-0023301312033121-2203211123200330"></a>

## sub_path property — mount / 022301111002 / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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

<a id="canonical-2312311312212121-1002331131122323-0121001130230221-3001121220030101-1112102021222101-0121010311302000-1200213113301120-1220211311112202"></a>

## Next pages — mount / 022301111002 / 7

- [stateful_service.configuration.parameters.file](resources--workload--reference--group-028.md#canonical-3111133103101303-0323302220303200-2031333233121200-0222313132030022-3332112322312012-1320033010330232-0302211321120013-3033112100233210)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020003202220121-3212302322020012-1333223330210231-2001323200032233-0332202313310032-0011330300332133-3110131212031121-3100123012133313"></a>

## stateful_service.containers — containers / 231213010032 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- stateful_service.containers

<a id="canonical-1113030003011000-0010130220333121-0021121013302333-2232320020302313-0033322320320313-0222330100312121-0130123102002030-2300120133302233"></a>

Type: `"object"`. list nested block, Optional.

Containers. Containers to use for service.

Upstream description:

Containers to use for service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("custom_flavor",
    "default_flavor"),
  validators.ConflictingListObjectAttributes("custom_flavor",
    "flavor"),
  validators.ConflictingListObjectAttributes("default_flavor",
    "flavor")}
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

Terraform syntax:

```terraform
containers {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011123332120023-0103333322013300-1331201323113113-3022332310003321-2300023203123100-1301212313101203-0331112111001102-3202311201112102"></a>

## Direct properties — containers / 231213010032 / 3

<a id="canonical-0302200203121231-3303000323310310-1102102212320233-2010201021110301-1221000222221312-2120310110302112-2233223023110002-1131122322321320"></a>

<a id="canonical-2013330122001002-2112313231131333-2311112213011022-1000103021110213-1233221332121103-2030302232200302-1113110333300033-1131130310120323"></a>

## args property — containers / 231213010032 / 4

Type: `["list", "string"]`. Optional.

Arguments to the entrypoint. Overrides the Docker image's CMD.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-0210003003322332-0203311023330301-0220013100311120-1223131303213312-2100000110232023-2003100103013110-3311020230120200-0201001010103321"></a>

<a id="canonical-0131322230132320-3121221122212301-1011003032022033-1213100001233030-2331121331122001-0110000130120312-1022111103210013-2101112133102000"></a>

## command property — containers / 231213010032 / 5

Type: `["list", "string"]`. Optional.

Command to execute. Overrides the Docker image's ENTRYPOINT.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [custom_flavor](resources--workload--reference--group-028.md#canonical-3213302013021002-0200100013303102-2133231113010230-0012320031302120-0111021223212331-2223102120311311-1102330223321001-0011001011331122): complete subsection reference.

- [default_flavor](resources--workload--reference--group-028.md#canonical-2220323323131313-3213223223212102-1023010101002123-1010210210130022-3213010130010023-0000113022111132-1222132313331300-2003202303001101): complete subsection reference.

<a id="canonical-3021233321033321-3101113132213002-2310100123313111-0020132221001310-0130100310023230-2100001201220301-2121322320201111-1133020322102031"></a>

<a id="canonical-0333013021102313-0100201232132321-0301113132032013-2303120333003101-0222023033222222-2202320233030022-3211220011123312-1031030003001312"></a>

## flavor property — containers / 231213010032 / 6

Type: `"string"`. Optional.

\[Enum:
CONTAINER\_FLAVOR\_TYPE\_TINY|CONTAINER\_FLAVOR\_TYPE\_MEDIUM|CONTAINER\_FLAVOR\_TYPE\_LARGE\]
Container Flavor type - CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny Tiny containers have limit of 0.1 vCPU
and 256 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium Medium containers have limit
of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_LARGE: Large Large containers
have.. Possible values are \`CONTAINER\_FLAVOR\_TYPE\_TINY\`, \`CONTAINER\_FLAVOR\_TYPE\_MEDIUM\`,
\`CONTAINER\_FLAVOR\_TYPE\_LARGE\`. Defaults to \`CONTAINER\_FLAVOR\_TYPE\_TINY\`.

Upstream description:

Container Flavor type

&#8203;- CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny

Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium

Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_LARGE: Large

Large containers have limit of 1 vCPU and 2048 MiB (mebibyte) memory.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTAINER_FLAVOR_TYPE_TINY",
    "CONTAINER_FLAVOR_TYPE_MEDIUM",
    "CONTAINER_FLAVOR_TYPE_LARGE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTAINER_FLAVOR_TYPE_TINY",
  "enum": [
    "CONTAINER_FLAVOR_TYPE_TINY",
    "CONTAINER_FLAVOR_TYPE_MEDIUM",
    "CONTAINER_FLAVOR_TYPE_LARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [image](resources--workload--reference--group-028.md#canonical-0210130211131331-2131110001320022-1333130131033002-1332001300303301-1003303222231001-1031320230211023-3113003031002110-0122201210203302): complete subsection reference.

<a id="canonical-0330321120110120-2302311101022233-3302032333232112-0302123232213202-2200202311200113-3220232030131123-3132130110230333-3331011120111102"></a>

<a id="canonical-0212302302133002-2213133122232121-2330300111032101-2001100301232300-1122031330020321-0033012033230320-1130000003331000-3200222011200213"></a>

## init_container property — containers / 231213010032 / 7

Type: `"bool"`. Optional.

Specialized container that runs before application container and runs to completion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111): complete subsection reference.

<a id="canonical-3233232112033033-3233003233311113-3110313330213300-2232123200001302-1330311112103130-3030221002203301-2002303301333022-2321333121132212"></a>

<a id="canonical-2333210201020131-2132132113013220-2202331333221212-3221202031020202-0323010110133321-1122302333222010-2033101300131311-0112032001203012"></a>

## name property — containers / 231213010032 / 8

Type: `"string"`. Optional.

Name. Name of the container.

Upstream description:

Name of the container.

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

- [readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321): complete subsection reference.

<a id="canonical-2230332001231012-3300011120011010-0200331320000031-0333210221221223-2203322221332031-3012100213333010-0310220130031211-1313201333301331"></a>

## Next pages — containers / 231213010032 / 9

- [stateful_service.containers.custom_flavor](resources--workload--reference--group-028.md#canonical-3213302013021002-0200100013303102-2133231113010230-0012320031302120-0111021223212331-2223102120311311-1102330223321001-0011001011331122)
- [stateful_service.containers.default_flavor](resources--workload--reference--group-028.md#canonical-2220323323131313-3213223223212102-1023010101002123-1010210210130022-3213010130010023-0000113022111132-1222132313331300-2003202303001101)
- [stateful_service.containers.image](resources--workload--reference--group-028.md#canonical-0210130211131331-2131110001320022-1333130131033002-1332001300303301-1003303222231001-1031320230211023-3113003031002110-0122201210203302)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3213302013021002-0200100013303102-2133231113010230-0012320031302120-0111021223212331-2223102120311311-1102330223321001-0011001011331122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111231221210110-0131131321002333-2130321201202030-2031102121122221-3300113330033210-3123311100320203-3202312131032200-3132110112313011"></a>

## stateful_service.containers.custom_flavor — custom_flavor / 231303103220 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- stateful_service.containers.custom_flavor

<a id="canonical-3033231233130120-1132332211121203-1332322321023230-0321020313200330-1332100013131120-0000223100131203-1133313320211013-3331212112000001"></a>

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
custom_flavor {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300100101122211-2311020311223323-1213011012223322-3212003203020033-0323312333100112-1330102103312320-2310122200033122-0121131320013212"></a>

## Direct properties — custom_flavor / 231303103220 / 3

<a id="canonical-3130313112031001-3332132002101332-1001112020200220-1012232221301123-1111031200120222-1001111121330233-1302131223022300-0200002333101232"></a>

<a id="canonical-2121221000211101-1201220202003120-1130303113333210-0210000232102132-1033032233110132-1222122313201233-3023003223303130-0031100111222300"></a>

## name property — custom_flavor / 231303103220 / 4

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

<a id="canonical-3110232020213111-1022113322121023-2121012112310223-3113201233001103-3323011033222003-1010213333010033-3133133320313110-2121300131323010"></a>

<a id="canonical-1003003112000322-0003101100020303-1333323203213112-1231333011301331-2121033322102320-2322310113233100-0011323111100230-2132122130112101"></a>

## namespace property — custom_flavor / 231303103220 / 5

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

<a id="canonical-2322213101110213-3223320233320031-2320333221121011-2100311012133311-2213202022300311-1312010303311022-3103033323210023-2202133022031122"></a>

<a id="canonical-0202212132110231-2223323310310312-3201222322211010-2103020330113213-0020021322020100-1123132312203202-2102000330232331-0223222233331132"></a>

## tenant property — custom_flavor / 231303103220 / 6

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

<a id="canonical-2231330123333333-0003303011331322-2002203100323111-3210102322311130-2111133331311311-2202033211212131-3202132333103113-0203002030201010"></a>

## Next pages — custom_flavor / 231303103220 / 7

- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2220323323131313-3213223223212102-1023010101002123-1010210210130022-3213010130010023-0000113022111132-1222132313331300-2003202303001101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233332113131120-1212113230031300-0301223011131001-0132230112131030-0033111030012301-3200321010332130-1010031201303233-2020023302221030"></a>

## stateful_service.containers.default_flavor — default_flavor / 323031232322 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- stateful_service.containers.default_flavor

<a id="canonical-3233211301023200-1303330031010210-2303303020321200-3120032331013331-3330302300023332-1230201322232003-3012120300021030-0130110333103310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default flavor.

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
default_flavor = {}
```

<a id="canonical-1311221310333012-2223000322101313-3130102111332310-3303111101202201-3032332320222312-2112031232310103-0213230210000012-2200132030011120"></a>

## Direct properties — default_flavor / 323031232322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221023020330221-0011022022230121-1011312231311101-1112203121003103-0202320312211132-0003031202303321-1213023132010321-3122003110212332"></a>

## Next pages — default_flavor / 323031232322 / 4

- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0210130211131331-2131110001320022-1333130131033002-1332001300303301-1003303222231001-1031320230211023-3113003031002110-0122201210203302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030212333300221-2210113203303132-0031223331120313-3322131001300122-0132132331022210-2120320232310031-1332033103132111-3330200001032212"></a>

## stateful_service.containers.image — image / 203000101333 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- stateful_service.containers.image

<a id="canonical-3001023131330212-3030133103033233-3102010310130023-1233001101222100-2211212033323210-1223220120212231-3320001030021301-0330311222111212"></a>

Type: `"object"`. single nested block, Optional.

ImageType configures the image to use, how to pull the image, and the associated secrets to use if
any.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name"),
  validators.ConflictingObjectAttributes("container_registry",
    "public")}
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
  "x-ves-oneof-field-registry_choice": "[\"container_registry\",\"public\"]"
}
```

Terraform syntax:

```terraform
image {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222032330310231-3021012313231310-2310330313000111-0331130121312131-3330022122120001-3201031121120031-1021222000013012-0103132211121322"></a>

## Direct properties — image / 203000101333 / 3

- [container_registry](resources--workload--reference--group-028.md#canonical-1310113331123012-2113223012123210-3020033210132330-2130201100002301-3023312013303313-3102233310233203-1331211112031221-1100202031131313): complete subsection reference.

<a id="canonical-1313132001301020-3131313223113330-1201013132303011-0113123120230310-3212200222323303-2330120222123310-1202010132330130-0212111203232232"></a>

<a id="canonical-3033120102133220-0313221121130231-2313321322001311-3012331312223201-1130201031103002-1222223021003311-1331021032213232-3011002031331023"></a>

## name property — image / 203000101333 / 4

Type: `"string"`. Optional.

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed.

Upstream description:

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed.

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

- [public](resources--workload--reference--group-028.md#canonical-1010202101123303-0233303332120030-3031220123331031-2201301010020310-2120002112132330-0230230121311201-3221211223232113-0231313011210133): complete subsection reference.

<a id="canonical-3210102203303223-1212113132033300-3203323112012213-2200320232331210-1300023312202232-0220132122221030-2020030331230222-1302333101212122"></a>

<a id="canonical-2333103301030203-3023200132211013-3223322233323202-0322200132330020-2201120212321131-1311021321202231-0310233230121212-0301020000212210"></a>

## pull_policy property — image / 203000101333 / 5

Type: `"string"`. Optional.

\[Enum:
IMAGE\_PULL\_POLICY\_DEFAULT|IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT|IMAGE\_PULL\_POLICY\_ALWAYS|IMAGE\_PULL\_POLICY\_NEVER\]
Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload - IMAGE\_PULL\_POLICY\_DEFAULT: Default Default will always pull image if :latest tag
is specified in image name. If :latest tag is not specified in image name, it will pull image only..
Possible values are \`IMAGE\_PULL\_POLICY\_DEFAULT\`, \`IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT\`,
\`IMAGE\_PULL\_POLICY\_ALWAYS\`, \`IMAGE\_PULL\_POLICY\_NEVER\`. Defaults to
\`IMAGE\_PULL\_POLICY\_DEFAULT\`.

Upstream description:

Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload

&#8203;- IMAGE\_PULL\_POLICY\_DEFAULT: Default

Default will always pull image if :latest tag is specified in image name. If :latest tag is not
specified in image name, it will pull image only if it does not already exist on the node &#8203;-
IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT: IfNotPresent

Only pull the image if it does not already exist on the node &#8203;- IMAGE\_PULL\_POLICY\_ALWAYS:
Always

Always pull the image &#8203;- IMAGE\_PULL\_POLICY\_NEVER: Never

Never pull the image.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("IMAGE_PULL_POLICY_DEFAULT",
    "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
    "IMAGE_PULL_POLICY_ALWAYS",
    "IMAGE_PULL_POLICY_NEVER"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "IMAGE_PULL_POLICY_DEFAULT",
  "enum": [
    "IMAGE_PULL_POLICY_DEFAULT",
    "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
    "IMAGE_PULL_POLICY_ALWAYS",
    "IMAGE_PULL_POLICY_NEVER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2022320021213002-3122212021031121-3211330111013102-1213211131300130-0020233111233231-2013120002310213-3320213110020231-3303131010231221"></a>

## Next pages — image / 203000101333 / 6

- [stateful_service.containers.image.container_registry](resources--workload--reference--group-028.md#canonical-1310113331123012-2113223012123210-3020033210132330-2130201100002301-3023312013303313-3102233310233203-1331211112031221-1100202031131313)
- [stateful_service.containers.image.public](resources--workload--reference--group-028.md#canonical-1010202101123303-0233303332120030-3031220123331031-2201301010020310-2120002112132330-0230230121311201-3221211223232113-0231313011210133)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1310113331123012-2113223012123210-3020033210132330-2130201100002301-3023312013303313-3102233310233203-1331211112031221-1100202031131313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122101001221030-1321001302203021-1331031131031021-3232130211203213-0212033132313200-1023010333302121-3020023111313013-1200203210220312"></a>

## stateful_service.containers.image.container_registry — container_registry / 133132023001 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.image](resources--workload--reference--group-028.md#canonical-0210130211131331-2131110001320022-1333130131033002-1332001300303301-1003303222231001-1031320230211023-3113003031002110-0122201210203302)
- stateful_service.containers.image.container_registry

<a id="canonical-3101103113313320-3033213030112022-0311203332333223-0210201103123211-1231213013131010-3332331011020231-0230213102133110-3113232101000123"></a>

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
container_registry {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313022230111132-2331232100021210-1021202020112022-3303211131220023-2103311332333021-0330313232203132-0322012231311002-1322001303212002"></a>

## Direct properties — container_registry / 133132023001 / 3

<a id="canonical-1221113233112033-0133110213203212-1313232030311232-1213330021332302-0021113133200011-3320011330003003-0311322033202333-2310221201123002"></a>

<a id="canonical-1020021310032220-1023321121000300-0232123022201003-0113103113202012-1103223311303100-0303023100310320-2132301010122113-1103031323310302"></a>

## name property — container_registry / 133132023001 / 4

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

<a id="canonical-0312131233211133-1330101321013023-0223010130023313-3301212002100233-0233210102011231-1232200321132320-2132010213013321-0313011220013220"></a>

<a id="canonical-1002331322333310-1100220011120302-3332111121333111-0120320010010132-2122022230220101-2221313220112222-1311122310311203-3123022221211332"></a>

## namespace property — container_registry / 133132023001 / 5

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

<a id="canonical-0221012130201021-3232113212001123-3303332333000322-1130321120232221-1203323332222121-1320023012110322-1320220323200201-0023021302313012"></a>

<a id="canonical-0300230333121331-1123201210210232-1233210301213211-2001102111123003-3121220322322223-2033202003301322-1221333322121100-3321033222001103"></a>

## tenant property — container_registry / 133132023001 / 6

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

<a id="canonical-1301132130132132-0223321323103030-1212002010113332-0012302003320102-3032221131322110-3112020133201013-0011021011101011-1110002100231210"></a>

## Next pages — container_registry / 133132023001 / 7

- [stateful_service.containers.image](resources--workload--reference--group-028.md#canonical-0210130211131331-2131110001320022-1333130131033002-1332001300303301-1003303222231001-1031320230211023-3113003031002110-0122201210203302)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1010202101123303-0233303332120030-3031220123331031-2201301010020310-2120002112132330-0230230121311201-3221211223232113-0231313011210133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111031120300302-3131202223230301-2323133300001032-0311202033032221-3300122023332131-3321230030322313-2212011110211111-1201221022302131"></a>

## stateful_service.containers.image.public — public / 330331000301 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.image](resources--workload--reference--group-028.md#canonical-0210130211131331-2131110001320022-1333130131033002-1332001300303301-1003303222231001-1031320230211023-3113003031002110-0122201210203302)
- stateful_service.containers.image.public

<a id="canonical-0210003223122312-0333203231302013-3100003211123231-0112320200321100-2112331000110030-0131313112220211-2332101212113130-2322332322112131"></a>

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
public = {}
```

<a id="canonical-0231113301211022-1012321220202100-2002331311122320-1133331133313302-1330023312102121-1311033233033121-2132123300110310-1023221110002212"></a>

## Direct properties — public / 330331000301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313213011220023-2033101222132022-3231012211230013-2313033231301233-3230301223330013-3302210233333230-2233312212011110-3132132111032201"></a>

## Next pages — public / 330331000301 / 4

- [stateful_service.containers.image](resources--workload--reference--group-028.md#canonical-0210130211131331-2131110001320022-1333130131033002-1332001300303301-1003303222231001-1031320230211023-3113003031002110-0122201210203302)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033300213223020-2212111221110023-2230313301120320-1123213333310111-1012213301303210-0002221232000232-3322210331220011-1130011303120012"></a>

## stateful_service.containers.liveness_check — liveness_check / 222012220131 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- stateful_service.containers.liveness_check

<a id="canonical-0320023333012002-2311211021333230-3011110221210203-1110030132210302-1101021123023103-1312200232210212-1332232213330211-2201331123001213"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "http_health_check"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "tcp_health_check"),
  validators.ConflictingObjectAttributes("http_health_check",
    "tcp_health_check")}
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
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

Terraform syntax:

```terraform
liveness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211001002311011-0123222122113231-2020110011233212-3302021112101310-0200131002311233-0130232330312212-3333101332210300-3010123333313200"></a>

## Direct properties — liveness_check / 222012220131 / 3

- [exec_health_check](resources--workload--reference--group-028.md#canonical-3202211021213013-0010223010220301-3100011320332220-1133312022220330-3102212031212130-0300013331012011-3320111013133111-2110023110323311): complete subsection reference.

<a id="canonical-3322303113312020-0223121313300111-0010233030022022-3002011111202131-1302230030023122-1303001322310031-0311001303032323-1030331220023000"></a>

<a id="canonical-2101001222001031-1111310020200210-0313322031321022-2210012320211001-2001300332233312-1232233221003133-3222023111200133-3213213202330021"></a>

## healthy_threshold property — liveness_check / 222012220131 / 4

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](resources--workload--reference--group-028.md#canonical-0110313303111000-0201320232102000-3002132013033323-0000210321132302-0130313122020010-1120231201101213-2331033302000122-2030120120202000): complete subsection reference.

<a id="canonical-0123111122003333-1010233333030220-1210120211023102-2033320031232111-3231102313222201-0121133000021320-2210010112023331-2122031333130032"></a>

<a id="canonical-1331102303011101-1213133122132312-3311320132032330-0030221102123003-0220031121001030-0222203311132002-1321211311022020-1100012022122133"></a>

## initial_delay property — liveness_check / 222012220131 / 5

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-1030313032132003-0111222133222120-1011000201323130-2323232013303030-2311310033201302-1310003333222332-3021233202231130-0320103223200132"></a>

<a id="canonical-1122103011132331-1122320221231330-2120102102331311-3030203023300321-0002310133130111-0130320021211012-3223102212123020-2132010230011103"></a>

## interval property — liveness_check / 222012220131 / 6

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](resources--workload--reference--group-028.md#canonical-2211202333323310-1203311303331120-2033130120013202-0200033312120121-3033212311030220-3331130220320002-0311323221310111-0032011220101332): complete subsection reference.

<a id="canonical-1102232021212222-1032032020013311-2031203321102101-1130033330102223-3110323200210002-0131303302032030-3013113213211011-2321221322322311"></a>

<a id="canonical-2103333202221110-1110123302011031-0102210110101023-0331021322013311-1202120002332002-3232323333320301-3200213020130131-3120102311111223"></a>

## timeout property — liveness_check / 222012220131 / 7

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-2130032331112332-0303112021320233-2121303012233223-2220213313020121-0100323102213231-1222211201232112-3011233203331223-0000133200233232"></a>

<a id="canonical-2330322321031301-2031031113201001-2031130313221213-1201200332022033-3002122331323323-3312100000321331-0021310303111021-3311321121012320"></a>

## unhealthy_threshold property — liveness_check / 222012220131 / 8

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-2321221032012130-0322312111113110-3123223322332300-1011320220230111-3311011130303032-2332011332303033-0111113311311211-1230202121130020"></a>

## Next pages — liveness_check / 222012220131 / 9

- [stateful_service.containers.liveness_check.exec_health_check](resources--workload--reference--group-028.md#canonical-3202211021213013-0010223010220301-3100011320332220-1133312022220330-3102212031212130-0300013331012011-3320111013133111-2110023110323311)
- [stateful_service.containers.liveness_check.http_health_check](resources--workload--reference--group-028.md#canonical-0110313303111000-0201320232102000-3002132013033323-0000210321132302-0130313122020010-1120231201101213-2331033302000122-2030120120202000)
- [stateful_service.containers.liveness_check.tcp_health_check](resources--workload--reference--group-028.md#canonical-2211202333323310-1203311303331120-2033130120013202-0200033312120121-3033212311030220-3331130220320002-0311323221310111-0032011220101332)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3202211021213013-0010223010220301-3100011320332220-1133312022220330-3102212031212130-0300013331012011-3320111013133111-2110023110323311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312222310321021-2232022320021122-2001012220022332-3002101102000230-3311233312220201-1030212032232310-2002032232023130-3323021003033212"></a>

## stateful_service.containers.liveness_check.exec_health_check — exec_health_check / 231002011011 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111)
- stateful_service.containers.liveness_check.exec_health_check

<a id="canonical-0223120101113202-3032331300013333-0223123233121021-2310311233213303-0322011202031033-0311102221130112-0021013130312300-3232313322213123"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("command")}
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
exec_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013123022200203-2022003213013221-1030033103121131-0220030230201021-1332011123022002-3002023310301133-1221011220102100-3122103133332023"></a>

## Direct properties — exec_health_check / 231002011011 / 3

<a id="canonical-0232202132022102-0301223332331101-0311222232322313-2213301310101101-1230013021110003-1121121131330011-3231132121031233-0302220220210030"></a>

<a id="canonical-1033120130132300-0103202220312320-2101201022030201-3321313000002210-1012311332102200-3013213320302120-3211102101301220-2203131301121023"></a>

## command property — exec_health_check / 231002011011 / 4

Type: `["list", "string"]`. Optional.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2330333003330032-3210131223100012-1001300311102100-2120000023122022-2230032312211303-3020132300223330-1002212010023033-2101111132211321"></a>

## Next pages — exec_health_check / 231002011011 / 5

- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0110313303111000-0201320232102000-3002132013033323-0000210321132302-0130313122020010-1120231201101213-2331033302000122-2030120120202000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233322133120130-0030233311301321-1230320102111100-2123230022002203-2211233121111331-2132031203121303-2320110130020300-0021312122333132"></a>

## stateful_service.containers.liveness_check.http_health_check — http_health_check / 223131320031 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111)
- stateful_service.containers.liveness_check.http_health_check

<a id="canonical-0211101021013333-2301021331102013-0301213132000120-2111100322220323-2323121311200200-2312211223202023-3011202132210013-0311003301333103"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-1110020212233201-2002112202210023-1302113112021113-2323311011002023-1211121003030022-3230332202003231-0213010301320020-2302221003303101"></a>

## Direct properties — http_health_check / 223131320031 / 3

<a id="canonical-0220003000022122-3212233013312302-0020311323032101-3212102012023232-1211001113032230-1001223213002212-2232233202322113-0312333200322332"></a>

<a id="canonical-2013233223122111-0133311320202210-0200001210210202-0031131323201133-1023213001123323-2311101223311310-3130023301020020-1321311110121112"></a>

## headers property — http_health_check / 223131320031 / 4

Type: `["map", "string"]`. Optional.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-3201003023133333-0203021131021100-2113010312212101-0201231132033221-3331030122010303-3101012121213010-1133122320103103-0120223110001232"></a>

<a id="canonical-1010130222102022-2211302212313012-3321312221201120-2032110201202011-1101333032310031-1301112000232221-1313210212120123-1132123330230303"></a>

## host_header property — http_health_check / 223131320031 / 5

Type: `"string"`. Optional.

The value of the host header in the HTTP health check request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-1113001212213102-0203301231003132-1320003333310121-3322223013310200-0020121000032030-0210312011330312-3332303230201302-1110132222230102"></a>

<a id="canonical-1230203220102201-1133203203322330-1100202010032223-0130011123321303-0213223330331212-2313013120120303-2001232313011220-1230232000132310"></a>

## path property — http_health_check / 223131320031 / 6

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](resources--workload--reference--group-028.md#canonical-0303322310300133-0233313213223100-1012003101331013-0010311203033303-2102103210121101-1012100013320123-3311232311313232-2011322020221233): complete subsection reference.

<a id="canonical-0032310022332021-3230110013101301-0333110322101332-3022110123313011-2130211330102210-2130231330300230-1210123313132320-1321211202210312"></a>

## Next pages — http_health_check / 223131320031 / 7

- [stateful_service.containers.liveness_check.http_health_check.port](resources--workload--reference--group-028.md#canonical-0303322310300133-0233313213223100-1012003101331013-0010311203033303-2102103210121101-1012100013320123-3311232311313232-2011322020221233)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0303322310300133-0233313213223100-1012003101331013-0010311203033303-2102103210121101-1012100013320123-3311232311313232-2011322020221233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110121123321303-3112101132111121-2200113332110321-1233033122020223-0222022102003120-2212132111121230-0012321031113101-3233223322123001"></a>

## stateful_service.containers.liveness_check.http_health_check.port — port / 312333203300 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111)
- [stateful_service.containers.liveness_check.http_health_check](resources--workload--reference--group-028.md#canonical-0110313303111000-0201320232102000-3002132013033323-0000210321132302-0130313122020010-1120231201101213-2331033302000122-2030120120202000)
- stateful_service.containers.liveness_check.http_health_check.port

<a id="canonical-1001013021330012-0222020302120323-2322330332331013-1223133222120110-0221201111100031-2201221013332322-3112020313021130-1011210022031002"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111131110130202-1303203030220221-1111311232332030-2133313332313110-0112322333210312-1120303021310011-2313000322310332-0210212133121221"></a>

## Direct properties — port / 312333203300 / 3

<a id="canonical-0320200322102100-2010132323022323-0330131102013213-1110221020301112-1313301313033023-0103201020023302-2013100133123101-2300100011303322"></a>

<a id="canonical-3123311022330302-2310001023232301-1331000022003230-3220233232202112-0022103113113313-2131310200031212-0012303333013222-0030202002031103"></a>

## name property — port / 312333203300 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-2321322102312113-0022130022201311-0023302001123111-0331022313320113-2310232033231203-3320301221033302-3100310230221113-3330213233302030"></a>

<a id="canonical-3322312232311232-1310212113120032-0011203031032120-1121103001120232-1021211303232030-3311030330333313-2322011121332212-0223310032002301"></a>

## num property — port / 312333203300 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0113102112330023-1120322332311101-2303103221123113-1132031030322221-3123212110201023-2201221002202102-0332311333232221-2331312331101010"></a>

## Next pages — port / 312333203300 / 6

- [stateful_service.containers.liveness_check.http_health_check](resources--workload--reference--group-028.md#canonical-0110313303111000-0201320232102000-3002132013033323-0000210321132302-0130313122020010-1120231201101213-2331033302000122-2030120120202000)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2211202333323310-1203311303331120-2033130120013202-0200033312120121-3033212311030220-3331130220320002-0311323221310111-0032011220101332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110122010310333-0101302301002201-3032010002223210-0213123202022103-2303131101111031-2021033211001132-1130033311122130-0302033230103303"></a>

## stateful_service.containers.liveness_check.tcp_health_check — tcp_health_check / 321203321312 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111)
- stateful_service.containers.liveness_check.tcp_health_check

<a id="canonical-0031212212331002-3110330332021301-2021021233110110-2120230113313213-0300020223200232-2033033032333333-2122223312221301-3020002222202330"></a>

Type: `"object"`. single nested block, Optional.

TCPHealthCheckType describes a health check based on opening a TCP connection.

Receipt-pinned upstream constraints:

```json
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
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311222233110121-0123030112323212-3102011211012003-1030211000303230-3310002232032130-2220300011323033-2300330102003102-3101103021113122"></a>

## Direct properties — tcp_health_check / 321203321312 / 3

- [port](resources--workload--reference--group-028.md#canonical-1023303303332203-1232211220000333-2212302213222012-0123311331133103-3030230132332133-1332333111132332-0111333123031020-0130101220000030): complete subsection reference.

<a id="canonical-2111231313001113-3132010330000111-3113232112003200-2320200010112232-2001333031210203-1213023233012120-0103101320020011-0021322102313013"></a>

## Next pages — tcp_health_check / 321203321312 / 4

- [stateful_service.containers.liveness_check.tcp_health_check.port](resources--workload--reference--group-028.md#canonical-1023303303332203-1232211220000333-2212302213222012-0123311331133103-3030230132332133-1332333111132332-0111333123031020-0130101220000030)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1023303303332203-1232211220000333-2212302213222012-0123311331133103-3030230132332133-1332333111132332-0111333123031020-0130101220000030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322011322212301-3131033232321323-0120131233201333-3100331012110212-2101312331303232-0100213313323031-3302002300221213-1320213130123213"></a>

## stateful_service.containers.liveness_check.tcp_health_check.port — port / 202013131313 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111)
- [stateful_service.containers.liveness_check.tcp_health_check](resources--workload--reference--group-028.md#canonical-2211202333323310-1203311303331120-2033130120013202-0200033312120121-3033212311030220-3331130220320002-0311323221310111-0032011220101332)
- stateful_service.containers.liveness_check.tcp_health_check.port

<a id="canonical-2202201102011123-2023313313010221-0102033000003123-1033300322012313-0131103000332200-0301332023303121-0330233231112122-1221023021021301"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-0312020311322330-2002322230201002-1200121223233021-3032202210313230-0030332231021232-3313130310203300-0221210333120003-2231010122323001"></a>

## Direct properties — port / 202013131313 / 3

<a id="canonical-3303120011233211-1023012031123231-0001203322220000-1102230303123110-1131133332301212-2222233012001232-1032233210312033-3322133110111133"></a>

<a id="canonical-1330202122221330-2021101302010303-0113111032311020-3110030030112012-3131123102300210-2010130002322322-2110002011100321-0002222333112230"></a>

## name property — port / 202013131313 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-2021201200212211-3312221121211321-3113110330110001-0031203212111223-3220002213230233-1120303110213221-1131021003130131-3023212311002323"></a>

<a id="canonical-3002332032003322-0213203320013022-1202230112010120-1133032222202312-3302131102221111-3330031031231332-0112322011000323-2112311001220020"></a>

## num property — port / 202013131313 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2031133033033010-0021310221320011-3020011113010111-2122022320011031-2300322310223012-1200231013312113-0132002012210023-2102101212011033"></a>

## Next pages — port / 202013131313 / 6

- [stateful_service.containers.liveness_check.tcp_health_check](resources--workload--reference--group-028.md#canonical-2211202333323310-1203311303331120-2033130120013202-0200033312120121-3033212311030220-3331130220320002-0311323221310111-0032011220101332)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012003011110000-3030322122122011-3003300123011100-0123202120132222-2022021210202023-3101123002032122-0110320133133212-1031031201201302"></a>

## stateful_service.containers.readiness_check — readiness_check / 323122100101 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- stateful_service.containers.readiness_check

<a id="canonical-0323220310301030-3133212133211033-3332322000133113-2120302312100102-2312200300213201-2120220022300011-3202202202331300-1320200230223221"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "http_health_check"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "tcp_health_check"),
  validators.ConflictingObjectAttributes("http_health_check",
    "tcp_health_check")}
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
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

Terraform syntax:

```terraform
readiness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231122303200333-1002213010231023-1010120213013130-3010133323000211-3020032202020120-1221211113230200-0030012132230022-2312220312132013"></a>

## Direct properties — readiness_check / 323122100101 / 3

- [exec_health_check](resources--workload--reference--group-028.md#canonical-2120310202200113-3301220212111331-0022001202232223-3112320131020100-3232021031113102-2200212230203103-2200201020131302-1323310100003320): complete subsection reference.

<a id="canonical-2200202110212210-0010130300032112-1300012231113002-2013220302232202-3203311311103032-0123033202012022-3330011321023321-0200002200020001"></a>

<a id="canonical-2310211223211031-0201210302222320-1100303101002232-2100023222311133-2302322211222033-2201023031301231-0010312310133303-0310310021021113"></a>

## healthy_threshold property — readiness_check / 323122100101 / 4

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](resources--workload--reference--group-028.md#canonical-3103313332112223-2130033020321112-0322301232002211-2223322202123210-2331330121022122-1311321200133031-0301021310022213-3322311200230001): complete subsection reference.

<a id="canonical-2032323233022000-2111301312233212-3220202111202203-2321033120131000-3313310033332030-3311230232200223-0111123032232123-1111211103230331"></a>

<a id="canonical-1013010000101333-1103120312232122-2210311011120003-0230130113331310-0102321213032312-2330233032121112-2130113023230233-0301101123030313"></a>

## initial_delay property — readiness_check / 323122100101 / 5

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-3111111033122312-3203113301120013-3031100222000030-1311010023233033-3111023131321220-2021023130023003-0211113012332210-0232102320332321"></a>

<a id="canonical-0211321201021130-3203023122013313-0033111111202330-1002203231323230-0330322112200132-1201210313120310-3102022110231121-1023130023120321"></a>

## interval property — readiness_check / 323122100101 / 6

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](resources--workload--reference--group-028.md#canonical-2300221210013310-2003113111031132-1030311021321222-3322010001000232-1230010311010331-2111102232333312-1002231210303301-0322323223330032): complete subsection reference.

<a id="canonical-1233010033201322-1122030301233232-0202011331311310-0211001302311300-0013301100330102-3230022213333220-0312333021121123-0322233333303313"></a>

<a id="canonical-3121303122302223-0002002213231003-0012030000132323-1312310322230330-2212101112212312-3013301020113032-3120122110311220-3021002121312002"></a>

## timeout property — readiness_check / 323122100101 / 7

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-3222202101002313-0221320333112330-2020022233313320-0200323033322023-0203233000103112-2313133133221033-0031220312013033-3232013200230132"></a>

<a id="canonical-3311223023302220-1111101210003300-2101101202322113-1303021232301001-1000220130202032-0221222010231132-2321200221303302-1322212310001013"></a>

## unhealthy_threshold property — readiness_check / 323122100101 / 8

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-2323111313003221-1321123213212213-1201120003211221-2310321232112330-3010302313330033-3133121332321303-1320223110311132-0213133123032103"></a>

## Next pages — readiness_check / 323122100101 / 9

- [stateful_service.containers.readiness_check.exec_health_check](resources--workload--reference--group-028.md#canonical-2120310202200113-3301220212111331-0022001202232223-3112320131020100-3232021031113102-2200212230203103-2200201020131302-1323310100003320)
- [stateful_service.containers.readiness_check.http_health_check](resources--workload--reference--group-028.md#canonical-3103313332112223-2130033020321112-0322301232002211-2223322202123210-2331330121022122-1311321200133031-0301021310022213-3322311200230001)
- [stateful_service.containers.readiness_check.tcp_health_check](resources--workload--reference--group-028.md#canonical-2300221210013310-2003113111031132-1030311021321222-3322010001000232-1230010311010331-2111102232333312-1002231210303301-0322323223330032)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2120310202200113-3301220212111331-0022001202232223-3112320131020100-3232021031113102-2200212230203103-2200201020131302-1323310100003320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102102202022032-2121230101013321-1112333110132013-2303120110123332-0320003211000302-0012222033312301-0020002303020011-3210121323031300"></a>

## stateful_service.containers.readiness_check.exec_health_check — exec_health_check / 312121022231 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321)
- stateful_service.containers.readiness_check.exec_health_check

<a id="canonical-1021330120223030-2122321022222031-2220202113013110-2132123213112332-1222301211310121-3211102000203223-3100122121320031-2002020302130300"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("command")}
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
exec_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-0312013201210300-2002222213210213-0313122203311203-3032302202110033-2230213020103012-1231210302130232-2112212302130201-1203321033121022"></a>

## Direct properties — exec_health_check / 312121022231 / 3

<a id="canonical-1311101110021022-0030212232231232-1300320032323013-0011030312230013-1123031320001100-0022102101330330-1010222302212102-3132121021312211"></a>

<a id="canonical-3210221301002001-1123322101000023-3100100230321220-1103200333301103-3210121211221013-0300110203113103-1020131300131301-1310023312020011"></a>

## command property — exec_health_check / 312121022231 / 4

Type: `["list", "string"]`. Optional.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3112203231030313-1011010323112223-3312201223123023-1233323323222003-3020210222322301-1103130302220333-1320223100213023-0203033332312311"></a>

## Next pages — exec_health_check / 312121022231 / 5

- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3103313332112223-2130033020321112-0322301232002211-2223322202123210-2331330121022122-1311321200133031-0301021310022213-3322311200230001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331121022331023-2300233320011202-3330133013310130-3333300003203312-1232132230112230-0310020003311110-3230301233030032-2213102111232200"></a>

## stateful_service.containers.readiness_check.http_health_check — http_health_check / 011111022021 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321)
- stateful_service.containers.readiness_check.http_health_check

<a id="canonical-2003032130003110-0133333131320013-3310311320222313-3333131330211120-0023231130312103-0302021303013003-2132102130203133-0213303213023131"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112330010010332-0222133121120232-1303211112312010-0021302332130333-1210221323313311-2320333001102112-0311001020301003-1312032132230223"></a>

## Direct properties — http_health_check / 011111022021 / 3

<a id="canonical-0331301032010221-1213233122200002-0203012323313201-1020312223233211-0002331320312201-2233212131030220-2013300313033212-1003212200220200"></a>

<a id="canonical-2330332233221032-1022113122303110-2231123211321223-3213033200321110-0202202102013231-3111302121133303-1232022133210311-1213330201322131"></a>

## headers property — http_health_check / 011111022021 / 4

Type: `["map", "string"]`. Optional.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0300100120103133-0311203130330020-3112003301012300-1203110331121101-0320003301131123-3222310032120033-0210112233232323-2233012103110320"></a>

<a id="canonical-2121131203232013-1323003211011202-2311131001031330-0123200132110232-2232023222301113-3302032111213023-3001320121303313-1010232010333202"></a>

## host_header property — http_health_check / 011111022021 / 5

Type: `"string"`. Optional.

The value of the host header in the HTTP health check request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-2112331031223133-3123002102002110-3321001131130010-1131002002110101-1030230113302111-3313023011112300-0230300312131330-0032131023013112"></a>

<a id="canonical-1230103300210001-3030030210300330-2303203233100332-0311011331020011-2121003232113033-2212312000100120-2111330223332332-3000000210201302"></a>

## path property — http_health_check / 011111022021 / 6

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](resources--workload--reference--group-028.md#canonical-3231130010320201-0320033112203300-3112302023120132-2000032130130302-0101323200333230-1121220101133011-2332020330112110-0222103200100130): complete subsection reference.

<a id="canonical-0201300322320030-0310002312320020-1213320212223002-0220020221032331-3232212133333303-1021321122112131-2320210103021111-3211223203031220"></a>

## Next pages — http_health_check / 011111022021 / 7

- [stateful_service.containers.readiness_check.http_health_check.port](resources--workload--reference--group-028.md#canonical-3231130010320201-0320033112203300-3112302023120132-2000032130130302-0101323200333230-1121220101133011-2332020330112110-0222103200100130)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3231130010320201-0320033112203300-3112302023120132-2000032130130302-0101323200333230-1121220101133011-2332020330112110-0222103200100130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030322033221210-2101012132010232-0031112031102020-0003013231312123-2231221300113331-2002031130133121-3030333301210130-0111202301303303"></a>

## stateful_service.containers.readiness_check.http_health_check.port — port / 211021210023 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321)
- [stateful_service.containers.readiness_check.http_health_check](resources--workload--reference--group-028.md#canonical-3103313332112223-2130033020321112-0322301232002211-2223322202123210-2331330121022122-1311321200133031-0301021310022213-3322311200230001)
- stateful_service.containers.readiness_check.http_health_check.port

<a id="canonical-0221112331321211-2220201111002102-1330010110320311-0010313111110001-3223332331100130-0211230333023300-3101311222031101-2221333132022123"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220102101313010-3030330031322133-1012310311020330-2233111120133130-2101022033032323-1321001110220221-0120210202030000-2032232011303113"></a>

## Direct properties — port / 211021210023 / 3

<a id="canonical-1302120101221101-1322011210031210-0102220213022222-3220002031033310-2313233331233121-2023000011011320-0202320133213211-0222333222202010"></a>

<a id="canonical-2023030230111322-3220311112110311-1001000001300003-0101332111332012-2320133310202021-2131023321133302-2112320323330000-3130332132303331"></a>

## name property — port / 211021210023 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-1212332103010031-0230002021032223-0203012010230032-3010212301210233-3300022012321113-3203212011030031-3320323030312100-0330322022232100"></a>

<a id="canonical-3302022201020131-2023030021032303-1130312222102232-1103301311130030-1132333033031123-1120133230022321-0003312300201322-0330113132212301"></a>

## num property — port / 211021210023 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3301223030313213-0231332003113202-0322113113301012-1222212323021020-0211200011232202-0213111010312033-2130033001032300-1120232300312103"></a>

## Next pages — port / 211021210023 / 6

- [stateful_service.containers.readiness_check.http_health_check](resources--workload--reference--group-028.md#canonical-3103313332112223-2130033020321112-0322301232002211-2223322202123210-2331330121022122-1311321200133031-0301021310022213-3322311200230001)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2300221210013310-2003113111031132-1030311021321222-3322010001000232-1230010311010331-2111102232333312-1002231210303301-0322323223330032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
