---
page_title: "xcsh_advertise_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_advertise_policy reference."
---

# xcsh_advertise_policy reference

<a id="canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030213213221331-3313122003220232-1322101121110023-2112300000300030-0233301221013213-1000331220111102-0322312022222333-0332123330203201"></a>

## Property reference — Property reference / 212102001203 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- Property reference

<a id="canonical-0322310033330211-0003202212212022-1210332012310303-0302001133223021-3321131201333323-3200003323311103-1000213230211013-3033322301010301"></a>

## Direct properties — Property reference / 212102001203 / 3

<a id="canonical-3300220122332320-1330000311132032-3221213103313222-2000123233220132-3232333200021322-0302110012302032-2230233312112110-2331030001212130"></a>

<a id="canonical-1332102333030323-3021220103211110-2023123300200330-0001033201031222-1312302112221000-0002301230203321-3211131023200312-1032331010102012"></a>

## address property — Property reference / 212102001203 / 4

Type: `"string"`. Optional, Computed.

Optional. VIP to advertise. This VIP can be either V4/V6 address You can not specify this if where
contains a site or virtual site of type REGIONAL\_EDGE or public network If not specified and
'where' is specified with site or virtual site option, inside\_vip or outside\_vip specified in the
site..

Upstream description:

Optional. VIP to advertise. This VIP can be either V4/V6 address You can not specify this if where
contains a site or virtual site of type REGIONAL\_EDGE or public network If not specified and
"where" is specified with site or virtual site option, inside\_vip or outside\_vip specified in the
site object will be used based on the network type. If inside\_vip/outside\_vip is not configured in
the site object, system use interface IP in the respected networks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2320333032333210-0213123022233323-0023003302233123-3312122012203300-1221021311000121-3311000020320000-1130222322100302-0303220323322131"></a>

<a id="canonical-0322320110210111-1333202122132122-0103230111021103-2211313203332323-3122111010311223-2011230121300222-0120222231300222-2300311222031010"></a>

## annotations property — Property reference / 212102001203 / 5

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

<a id="canonical-0233131232113220-1113233120233133-3022313310100023-1022313123021233-0200032133310232-2210010232130033-3220021010300331-3033311322210023"></a>

<a id="canonical-1122323001031211-2133321300101121-2323000211203000-2123120210131233-1131030200332321-2000120002022313-1330033023001003-0312221223223212"></a>

## description property — Property reference / 212102001203 / 6

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

<a id="canonical-3222332331121200-0231333310133331-0023123010322010-2123002203100202-2001201213123100-1201331121022130-1130111303231230-2102022222301213"></a>

<a id="canonical-3213210101021031-2111032230010310-3033122022000220-0113001133012221-0213201113133102-3120301010122000-3322213010123100-2221001002132222"></a>

## disable property — Property reference / 212102001203 / 7

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

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

- [dualstack](resources--advertise_policy--reference--group-001.md#canonical-3130123120101222-1220203221123220-1312323010232302-0132230022313313-3223133102023311-2231101311030311-0233132213000003-3213202311001321): complete subsection reference.

<a id="canonical-3033011111131000-0121333013111300-2313220310101010-0130120200001330-0021023010103213-1312122212312132-0210120212120110-0100033022130232"></a>

<a id="canonical-1031110002210130-0210320331111002-3012301130012122-3320213313103113-2101010202122111-1102311213002103-0010321031212212-2031002110122000"></a>

## ID property — Property reference / 212102001203 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipv4](resources--advertise_policy--reference--group-001.md#canonical-0302023220132221-3313300011221002-2010113021102131-3030302011113300-1100111101221011-2012021103323130-0311003330022212-1112201103102030): complete subsection reference.

- [ipv6](resources--advertise_policy--reference--group-001.md#canonical-2120203212121232-3132123301211012-1002302311333123-1022322131230020-2232103220003320-2132221031000330-0222013101313223-3222110233021222): complete subsection reference.

<a id="canonical-0302202033122013-0320102200220303-2021112320231002-3220113112032033-2011202333132033-3002321012110320-2112123300323100-2101033330000333"></a>

<a id="canonical-1132112021331001-1202121220302302-1101333011322321-2121310210310003-2212233131121201-2022031112231030-2012023132303132-0222032022000212"></a>

## labels property — Property reference / 212102001203 / 9

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-0120301302003230-0320000202102103-1031303203211303-2302010022213231-0130102231300021-3131012120223300-0012323110000210-0011322112131331"></a>

<a id="canonical-3321211313102312-0232132013231111-0202332120231333-0313023320030332-2031022022313112-1212032000330110-1013210300033210-2002221203131122"></a>

## name property — Property reference / 212102001203 / 10

Type: `"string"`. Required.

Name of the Advertise Policy. Must be unique within the namespace.

Upstream description:

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

<a id="canonical-0132222013013111-1320311333101100-1331202302331223-3331200303201123-3202313111021021-2312230223133120-0223120220123232-1200313212101211"></a>

<a id="canonical-3123121331022210-0320030133323000-1020032230111023-2213203133121201-1012021310030120-2220122303031231-0102022303023001-0120312100311332"></a>

## namespace property — Property reference / 212102001203 / 11

Type: `"string"`. Required.

Namespace where the Advertise Policy is created.

Upstream description:

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

<a id="canonical-1031220002103320-3313300103012130-2030030330002030-0233231301331201-2131200311200133-1132202032303113-1113201020211012-0313032230032103"></a>

<a id="canonical-0000232133103330-3133010021213011-3030023110020020-2011100231032202-3122111312310031-1323302102313121-2211301003323232-3331131333230133"></a>

## port property — Property reference / 212102001203 / 12

Type: `"number"`. Optional, Computed.

\[OneOf: port, port\_ranges\] Exclusive with \[port\_ranges\] Port to advertise.

Upstream description:

Exclusive with \[port\_ranges\] Port to advertise.

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

OneOf alternatives in this subsection:

- [port](resources--advertise_policy--reference--group-001.md#canonical-1031220002103320-3313300103012130-2030030330002030-0233231301331201-2131200311200133-1132202032303113-1113201020211012-0313032230032103)
- [port_ranges](resources--advertise_policy--reference--group-001.md#canonical-3110102230213031-2103233312131302-3303323323230000-0211003201202321-0100230222020311-2122232102033213-3122310310333233-0101023321122310)

Select alternatives according to the provider validators above.

<a id="canonical-3110102230213031-2103233312131302-3303323323230000-0211003201202321-0100230222020311-2122232102033213-3122310310333233-0101023321122310"></a>

<a id="canonical-2013003333012312-2201323032330102-0302330031002233-1222210213302032-3200321111311223-1220131323110212-2022030223123112-0023323103301221"></a>

## port_ranges property — Property reference / 212102001203 / 13

Type: `"string"`. Optional, Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-1320030031303210-2012123231311022-2311221300211011-0122303020301310-1113002001301133-1001212231220001-0022330223112023-2220211112231023"></a>

<a id="canonical-0032112101022202-0013203313310210-0301022101231113-3213131311231212-1122010100302112-0032020210022122-2023113220021310-3112233110012012"></a>

## protocol property — Property reference / 212102001203 / 14

Type: `"string"`. Optional, Computed.

\[Enum: TCP|UDP\] Protocol. Protocol to advertise. Possible values are \`TCP\`, \`UDP\`.

Upstream description:

Protocol to advertise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TCP",
    "UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "TCP",
    "UDP"
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
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  }
}
```

- [public_ip](resources--advertise_policy--reference--group-001.md#canonical-2303301110202311-2222032220031020-3121303333031321-2103131332222133-3111111130312021-0320322113112130-0301212122230111-2223203233322012): complete subsection reference.

<a id="canonical-3120000130133200-2311110310000011-0100201301103131-3022312200220313-0313013000000112-2021233110200212-2000031113010200-2323132333222113"></a>

<a id="canonical-0323320102300013-0020333310221312-1102003111003032-1321323102013310-1122311313013333-1121300112303011-2132221033030312-3103310132130232"></a>

## skip_xff_append property — Property reference / 212102001203 / 15

Type: `"bool"`. Optional, Computed.

If set, the loadbalancer will not append the remote address to the x-forwarded-for HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [timeouts](resources--advertise_policy--reference--group-001.md#canonical-1003330112013233-3322213102311100-1200212333211211-1232231112331330-2101230211113313-1332123122003021-2020332033122232-2303333311102233): complete subsection reference.

- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330): complete subsection reference.

- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101): complete subsection reference.

<a id="canonical-2131223130302323-2212102221231130-1220332212010131-0221100210033223-2312232333101103-1303332111211223-0200220001003211-2210130112112021"></a>

## All schema paths — Property reference / 212102001203 / 16

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](resources--advertise_policy--reference--group-001.md#canonical-3300220122332320-1330000311132032-3221213103313222-2000123233220132-3232333200021322-0302110012302032-2230233312112110-2331030001212130) |
| `annotations` | [annotations](resources--advertise_policy--reference--group-001.md#canonical-2320333032333210-0213123022233323-0023003302233123-3312122012203300-1221021311000121-3311000020320000-1130222322100302-0303220323322131) |
| `description` | [description](resources--advertise_policy--reference--group-001.md#canonical-0233131232113220-1113233120233133-3022313310100023-1022313123021233-0200032133310232-2210010232130033-3220021010300331-3033311322210023) |
| `disable` | [disable](resources--advertise_policy--reference--group-001.md#canonical-3222332331121200-0231333310133331-0023123010322010-2123002203100202-2001201213123100-1201331121022130-1130111303231230-2102022222301213) |
| `dualstack` | [dualstack](resources--advertise_policy--reference--group-001.md#canonical-2223302203320301-0102300203330312-3000030102112333-0320010331122313-2113323302220233-3231303220233001-2230223020111011-1022220002122002) |
| `id` | [ID](resources--advertise_policy--reference--group-001.md#canonical-3033011111131000-0121333013111300-2313220310101010-0130120200001330-0021023010103213-1312122212312132-0210120212120110-0100033022130232) |
| `ipv4` | [ipv4](resources--advertise_policy--reference--group-001.md#canonical-3333033012100020-1022313002032303-0010300211113202-0003220022013101-3320331330122012-0110132233201121-1331303222102311-2220300002112100) |
| `ipv6` | [ipv6](resources--advertise_policy--reference--group-001.md#canonical-2000120313121022-0211201113123303-2301001001301321-1213011100021321-3113101302110023-3312133120223213-3333301232021003-0102122102101133) |
| `labels` | [labels](resources--advertise_policy--reference--group-001.md#canonical-0302202033122013-0320102200220303-2021112320231002-3220113112032033-2011202333132033-3002321012110320-2112123300323100-2101033330000333) |
| `name` | [name](resources--advertise_policy--reference--group-001.md#canonical-0120301302003230-0320000202102103-1031303203211303-2302010022213231-0130102231300021-3131012120223300-0012323110000210-0011322112131331) |
| `namespace` | [namespace](resources--advertise_policy--reference--group-001.md#canonical-0132222013013111-1320311333101100-1331202302331223-3331200303201123-3202313111021021-2312230223133120-0223120220123232-1200313212101211) |
| `port` | [port](resources--advertise_policy--reference--group-001.md#canonical-1031220002103320-3313300103012130-2030030330002030-0233231301331201-2131200311200133-1132202032303113-1113201020211012-0313032230032103) |
| `port_ranges` | [port_ranges](resources--advertise_policy--reference--group-001.md#canonical-3110102230213031-2103233312131302-3303323323230000-0211003201202321-0100230222020311-2122232102033213-3122310310333233-0101023321122310) |
| `protocol` | [protocol](resources--advertise_policy--reference--group-001.md#canonical-1320030031303210-2012123231311022-2311221300211011-0122303020301310-1113002001301133-1001212231220001-0022330223112023-2220211112231023) |
| `public_ip` | [public_ip](resources--advertise_policy--reference--group-001.md#canonical-0322231233100123-2310333311123033-1230000102300232-1330033213330222-1333232003131230-0002123003130331-3322000212211202-3123230103200103) |
| `public_ip.kind` | [public_ip.kind](resources--advertise_policy--reference--group-001.md#canonical-2000031232210111-2320023101323210-3202332133111210-2222012112013111-3113201033020211-1300110011333033-2233312221203311-3333103321200202) |
| `public_ip.name` | [public_ip.name](resources--advertise_policy--reference--group-001.md#canonical-2110230131231330-0330202031113012-3100200203102321-1202222032221300-0133103000032311-2000110003220310-1222232233311311-2021311121120100) |
| `public_ip.namespace` | [public_ip.namespace](resources--advertise_policy--reference--group-001.md#canonical-2333100103121203-2002013121310213-3203101313203212-2312123211321321-0131013101322233-1033311230310231-0101033123231000-2213031122100030) |
| `public_ip.tenant` | [public_ip.tenant](resources--advertise_policy--reference--group-001.md#canonical-1123311232120211-0333023110231212-2112310200213221-2102031110133323-2013302303333212-1201331102301331-3113212000210011-0223032210302120) |
| `public_ip.uid` | [public_ip.uid](resources--advertise_policy--reference--group-001.md#canonical-0312031303223112-3210302111310300-2310322313030301-3220131210300300-3130030320113231-0321111132220331-3323312200221100-1302111312123022) |
| `skip_xff_append` | [skip_xff_append](resources--advertise_policy--reference--group-001.md#canonical-3120000130133200-2311110310000011-0100201301103131-3022312200220313-0313013000000112-2021233110200212-2000031113010200-2323132333222113) |
| `timeouts` | [timeouts](resources--advertise_policy--reference--group-001.md#canonical-2002313031012202-1130333220100233-3002210202311010-1313030223133020-0132023132330100-0333223032310022-3320103121213032-2321221322313232) |
| `timeouts.create` | [timeouts.create](resources--advertise_policy--reference--group-001.md#canonical-0231312131331030-2032313323320303-2333201002323331-1313222211100121-2031010302001203-2131213302101003-2303212021323033-2113011031102013) |
| `timeouts.delete` | [timeouts.delete](resources--advertise_policy--reference--group-001.md#canonical-1110301110210023-2121312300022313-2030122103032202-1001311013311032-0121200031332331-2322311302020322-0023201132210113-0331103133310302) |
| `timeouts.read` | [timeouts.read](resources--advertise_policy--reference--group-001.md#canonical-0130020130001332-3022220313203022-1210323210113320-3101132023023300-2231333302300120-2322223032003220-1310013001210103-2302002001103310) |
| `timeouts.update` | [timeouts.update](resources--advertise_policy--reference--group-001.md#canonical-0021220131112032-3120022020101220-2020010230100230-3032122123003002-0202102031123330-0321232123301112-0330101210010331-2310023132222100) |
| `tls_parameters` | [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-2322033321212012-0123203312022323-2311133313230210-1033021100003012-2323312323020011-3131030023300302-0000211313232320-2013001101333330) |
| `tls_parameters.client_certificate_optional` | [tls_parameters.client_certificate_optional](resources--advertise_policy--reference--group-001.md#canonical-2230112232310221-2332321301033302-0332100122213210-0320020223221302-2300102231110201-0100002133300011-1101111132030322-3233310121200230) |
| `tls_parameters.client_certificate_required` | [tls_parameters.client_certificate_required](resources--advertise_policy--reference--group-001.md#canonical-3301133022122203-1333010231102033-2010121202133013-0202213202301012-2010311232010101-1200312200312231-3331000231223023-1113230022133033) |
| `tls_parameters.common_params` | [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-1320111103102012-2221131312212233-0031003033120221-0230120322000220-1001331320310323-2102333100202203-2331000213110202-2132122123013301) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](resources--advertise_policy--reference--group-001.md#canonical-1010113101221011-3121222102323122-1110011100120222-0311003212301023-0333001301310211-3221202300122321-0132120230311333-2331320203222321) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](resources--advertise_policy--reference--group-001.md#canonical-1131000230001312-3132132103111000-1133323003211013-0101311231130201-2202110333312112-3232322021011122-2112020011133322-1120001213011011) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](resources--advertise_policy--reference--group-001.md#canonical-2331021023331300-1220033213013301-3333211312033102-3301221203122012-1101122223212202-0010113022303221-2213102331102211-2130301313000003) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-0003122211010203-0123031121123201-2202223100123223-3023020210022033-3113122100212013-2233010323300220-1202332020332010-3133033101102230) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](resources--advertise_policy--reference--group-001.md#canonical-3213023000300032-1310031010033332-0301213222023123-2233323013303233-3220332201001201-2010230232333132-0331011033002030-3022001313311321) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--advertise_policy--reference--group-001.md#canonical-2330202322013132-2011000323203210-3223131220323130-1022230313222313-1132101302302123-0320321222110230-2023333112310220-1231021202120310) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--advertise_policy--reference--group-001.md#canonical-3010333002010132-2012131321021310-3000223320002233-1322103033131000-0331312321323333-0000112133003120-2020100020230301-1021111031013030) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](resources--advertise_policy--reference--group-001.md#canonical-1012132323002121-0223122230120300-0022210200012032-0011230021232301-0133221030221112-3103213133323211-1131001331130131-0022310221123010) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--advertise_policy--reference--group-001.md#canonical-3101320302030032-2301301101223203-2231131301020021-2321001212230332-2211333302222232-2320023203002113-2322022112031030-2333121000230030) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](resources--advertise_policy--reference--group-001.md#canonical-1211331231120110-0121133221133113-0133113130201132-3333213330020022-1013231203221012-1033011130313300-2112232312232021-1102000212132001) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--advertise_policy--reference--group-001.md#canonical-0020203220332313-2320022102201012-1122112033301323-0311213031122013-1213132110020103-1310011120123030-1122320210322103-2202030100131011) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--advertise_policy--reference--group-001.md#canonical-0003101032201110-0133032222033023-0330303033203120-3202300312301130-2132122322020203-3231313230311121-3012123200332000-3321023101313200) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](resources--advertise_policy--reference--group-001.md#canonical-3203000102112331-0223011000311313-2033223123333301-1012023113332200-2000202322111123-3013103210010300-0230211101330211-2023013203111221) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--advertise_policy--reference--group-001.md#canonical-1323230120000130-3223232322231230-1010232310233120-3232232300011222-1201001112320112-1133330100030311-1232102111310102-1002132320001123) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--advertise_policy--reference--group-001.md#canonical-0010130132221212-3101213232230032-2302122333313122-2003121010313103-1010322213223232-1211310321112311-1123130321311133-3002112031323120) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](resources--advertise_policy--reference--group-001.md#canonical-0231221333201302-3222220123310302-2002130300020111-1321322122020232-1332302202002033-0020320211210100-0123331321203230-2331010133103231) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](resources--advertise_policy--reference--group-001.md#canonical-0020323313003030-3303101303123303-0303123033330200-3010203333112030-3123013021032233-3023330021303211-0102202220021121-0303113102130303) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--advertise_policy--reference--group-001.md#canonical-2122011212003312-3010223102213213-3033113230010301-1011013110112030-1121203131232113-0300312301121111-1221132313012302-2333231013113111) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](resources--advertise_policy--reference--group-001.md#canonical-0023311131032113-2023102313002030-1312021330311012-1103230212233201-3330123232001131-3321020222132111-0112022322321033-0100131022222012) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](resources--advertise_policy--reference--group-001.md#canonical-2101303300331022-0212022132111333-0003213021330111-3221200121002300-2133111103301133-2002003132022211-1123300011022212-0313223231230222) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](resources--advertise_policy--reference--group-001.md#canonical-0230311120310233-0331213013013011-3223010101133021-2220233232000201-3222330330113032-3233111100003120-0012100332000222-0302213300330103) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--advertise_policy--reference--group-001.md#canonical-3133013331232120-1113311313030300-0103123012122331-0121012320021013-1202330112001233-1130300112310033-0111122000222223-2211033120233301) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](resources--advertise_policy--reference--group-001.md#canonical-1223112012110021-3123301123021033-0211033313032231-3222200222102003-1001332103112013-2131123232120303-2321112331121020-1102003101022101) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](resources--advertise_policy--reference--group-001.md#canonical-3133213323300222-3201221220103012-2313330001330231-1101123121103122-0311210110332012-1120322120212023-1212230130032231-3020132302020330) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](resources--advertise_policy--reference--group-001.md#canonical-2120300333312032-2012212012221312-0133211303121211-1002020113330221-1101103230012321-1111111103111203-3032020321113131-2231002131210010) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](resources--advertise_policy--reference--group-001.md#canonical-1210302201120303-3023132001001011-3321203201233201-0320001313032222-0332231021212031-1122233131211123-2023021033030313-0233130222333232) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](resources--advertise_policy--reference--group-001.md#canonical-2223201230332113-3333212213200232-3002012023222031-2022213012002003-3111202223223200-1010231333111303-2031201210022233-3223031001113012) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](resources--advertise_policy--reference--group-001.md#canonical-3100013313021322-0331210001033200-1030220301112333-3210303301003211-0001220311310002-0111302021020232-0223033303211132-1300113232100022) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](resources--advertise_policy--reference--group-001.md#canonical-1020000103010302-3100220002203322-3231301330113330-1130332232323311-3101012233221213-2113302321300022-3100331023130220-1203022033222332) |
| `tls_parameters.no_client_certificate` | [tls_parameters.no_client_certificate](resources--advertise_policy--reference--group-001.md#canonical-3201021233032203-1131121110130212-3130222021212132-1313003032021023-3300030122131123-3103133220202332-1032301233313312-1133201001313001) |
| `tls_parameters.xfcc_header_elements` | [tls_parameters.xfcc_header_elements](resources--advertise_policy--reference--group-001.md#canonical-1233320031133133-0133310112210120-0331233330223122-1201133030131311-0310202202100303-0102223210100322-0203100010020011-0322102202210200) |
| `where` | [where](resources--advertise_policy--reference--group-001.md#canonical-2312133200212030-0033010301322001-2210313312222202-0112112201322210-3201200100320030-1230333102333202-3002122110312113-1313332322213203) |
| `where.site` | [where.site](resources--advertise_policy--reference--group-001.md#canonical-1310001132210023-1202201231300231-3111301133233310-1012030312122010-0301300102130001-2121022313313202-3202210113220000-0323333030031033) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-2220213230232012-1323231312313000-2222113003130033-0030303303301111-3213130222032121-3213030233332300-1313001133133321-0333110230321302) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-2032010021233201-0013323200302020-2203110030031121-3212110003110223-0133223113112321-0323122220213313-2302031001301021-1301132103312103) |
| `where.site.network_type` | [where.site.network_type](resources--advertise_policy--reference--group-001.md#canonical-1220222222233121-1202332002311222-0202322303123012-0303331320313322-3120111012030132-1310132221113232-0133222221031231-3101010222103331) |
| `where.site.ref` | [where.site.ref](resources--advertise_policy--reference--group-001.md#canonical-3030103310312132-3200000222223211-1011022113120203-1200212113013300-2012333213120223-2313331321000303-3133232203202303-1302323021010223) |
| `where.site.ref.kind` | [where.site.ref.kind](resources--advertise_policy--reference--group-001.md#canonical-1233003131220223-0100303002020101-3221100013111222-2201013200332113-1222322003332111-2132201320200232-1133201012131211-0332233202300022) |
| `where.site.ref.name` | [where.site.ref.name](resources--advertise_policy--reference--group-001.md#canonical-1311123313333221-2111023213021102-2332303013203010-2221123313301323-0101003201203001-2122001122213223-2011202021330003-2022332112203101) |
| `where.site.ref.namespace` | [where.site.ref.namespace](resources--advertise_policy--reference--group-001.md#canonical-3313223230232311-3111111110030112-3012110213200131-3030202323122020-1230030111013033-3103132232211232-3303202100022120-1220203021221012) |
| `where.site.ref.tenant` | [where.site.ref.tenant](resources--advertise_policy--reference--group-001.md#canonical-0012220102112303-2022220330010323-3113210011012030-3223023211311103-1223223213121100-2333112030300310-1322311130130211-1111110211122101) |
| `where.site.ref.uid` | [where.site.ref.uid](resources--advertise_policy--reference--group-001.md#canonical-3112202112132001-1200332300222002-0202311103122021-1312011113113030-3131220330021110-1230122210311101-2323020323203023-2131200020113131) |
| `where.virtual_network` | [where.virtual_network](resources--advertise_policy--reference--group-001.md#canonical-2021323333110122-3000213021021112-0321232311300030-2222233203023223-1020203211232213-2101113320333231-1012013111121211-3001232130321101) |
| `where.virtual_network.ref` | [where.virtual_network.ref](resources--advertise_policy--reference--group-001.md#canonical-0023133322232003-2223032020333012-1123000211033333-3030310131033333-3222220210133331-3031111120202000-3120300321033021-2331003313002131) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](resources--advertise_policy--reference--group-001.md#canonical-2112000320023012-0332312113113012-3031130311111300-1103022032203321-1030101233332022-3233021100113032-2310131202230231-0022222012332133) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](resources--advertise_policy--reference--group-001.md#canonical-2011223333132110-0130122330012221-2330001012212210-1030033013321210-1111132133322100-2123010112201020-0223120322003231-3113012232132212) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](resources--advertise_policy--reference--group-001.md#canonical-2113022213330022-2020321001333312-2201203321320133-0310121133201121-1231011203232312-1222320120330220-3222010213011200-0312213133112233) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](resources--advertise_policy--reference--group-001.md#canonical-3112002233321301-0110102110002322-1220032202200002-1110301032033322-3231203202221013-0211230003322200-1100312310231302-2223301322220320) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](resources--advertise_policy--reference--group-001.md#canonical-1312201131210212-2233023303133013-3120003003012303-0033213223003332-0003011311202212-1223202113011311-1330123000101210-3230003103133122) |
| `where.virtual_site` | [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-3111123332111000-0330013222030012-1331331132031302-1011013021230202-3003023000213103-3031123321011033-1210302232311122-2322313310300323) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-0031221001130031-2303110231300122-0122100003002322-1102230132020300-2203222310111011-3211030030121121-3301131323012200-0301002313311130) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-0302232010313012-0030212122113323-2111132133120133-1231321010000301-3321301322333323-2033100232300220-3200332120031333-1203033230210312) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](resources--advertise_policy--reference--group-001.md#canonical-1212223010033230-2233030212033332-2132221231323012-1101312003330333-1333101003112033-1320333302021003-1212202010301120-1130313330123223) |
| `where.virtual_site.ref` | [where.virtual_site.ref](resources--advertise_policy--reference--group-001.md#canonical-2230002001211311-1123223202012320-3130101232012110-1312133223233010-0321333301132132-3032113203102201-0030322013232320-0010312001210020) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](resources--advertise_policy--reference--group-001.md#canonical-3211131022023000-1033222101131022-2000013312233000-3230102012132031-0110321031001001-2203021222203101-1021133103221030-1010113300010013) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](resources--advertise_policy--reference--group-001.md#canonical-1303130221131111-2131133022111033-3222023311011332-0033023212112001-3031321310113021-3130223321103201-1122122323002211-1213202303323220) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](resources--advertise_policy--reference--group-001.md#canonical-2133222033113033-0133103330313332-3202202010301212-3102210012301001-3200332013121200-0100103032310222-2301111103132020-1321000311003311) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](resources--advertise_policy--reference--group-001.md#canonical-3122022030012320-0110111001311122-2331302333023302-3332120332103023-0311123111310202-3132330223213301-1210122022022333-3201120301230313) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](resources--advertise_policy--reference--group-001.md#canonical-1313013321313013-3022110021320313-3001223112333003-1132330013021022-1022223230133002-1130111202232312-1300210330201202-2301100213120013) |

<a id="canonical-0002323320132122-3231101010211221-0310111120031121-3110321201310130-3103013322033331-2121122311221101-3112303111332220-0213303010031323"></a>

## Next pages — Property reference / 212102001203 / 17

- [dualstack](resources--advertise_policy--reference--group-001.md#canonical-3130123120101222-1220203221123220-1312323010232302-0132230022313313-3223133102023311-2231101311030311-0233132213000003-3213202311001321)
- [ipv4](resources--advertise_policy--reference--group-001.md#canonical-0302023220132221-3313300011221002-2010113021102131-3030302011113300-1100111101221011-2012021103323130-0311003330022212-1112201103102030)
- [ipv6](resources--advertise_policy--reference--group-001.md#canonical-2120203212121232-3132123301211012-1002302311333123-1022322131230020-2232103220003320-2132221031000330-0222013101313223-3222110233021222)
- [public_ip](resources--advertise_policy--reference--group-001.md#canonical-2303301110202311-2222032220031020-3121303333031321-2103131332222133-3111111130312021-0320322113112130-0301212122230111-2223203233322012)
- [timeouts](resources--advertise_policy--reference--group-001.md#canonical-1003330112013233-3322213102311100-1200212333211211-1232231112331330-2101230211113313-1332123122003021-2020332033122232-2303333311102233)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-3130123120101222-1220203221123220-1312323010232302-0132230022313313-3223133102023311-2231101311030311-0233132213000003-3213202311001321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221101122212322-3220031230311110-0110302311333111-1021122111123301-2220112103300032-2122132210001323-1213300321122311-0122212131300231"></a>

## dualstack — dualstack / 131303201301 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- dualstack

<a id="canonical-2223302203320301-0102300203330312-3000030102112333-0320010331122313-2113323302220233-3231303220233001-2230223020111011-1022220002122002"></a>

Type: `["object", {}]`. Optional.

\[OneOf: dualstack, IPv4, IPv6\] Enable this option

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

- [dualstack](resources--advertise_policy--reference--group-001.md#canonical-2223302203320301-0102300203330312-3000030102112333-0320010331122313-2113323302220233-3231303220233001-2230223020111011-1022220002122002)
- [ipv4](resources--advertise_policy--reference--group-001.md#canonical-3333033012100020-1022313002032303-0010300211113202-0003220022013101-3320331330122012-0110132233201121-1331303222102311-2220300002112100)
- [ipv6](resources--advertise_policy--reference--group-001.md#canonical-2000120313121022-0211201113123303-2301001001301321-1213011100021321-3113101302110023-3312133120223213-3333301232021003-0102122102101133)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dualstack = {}
```

<a id="canonical-3310222022203002-0102231123232123-0213312212001331-3310232021203322-0222132012113201-3323031130120112-0333330103031131-3021103331102000"></a>

## Direct properties — dualstack / 131303201301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212101303022113-1231101131223133-0130110130100121-1223332130111120-2103313001320111-0323002113202231-3333310020313032-2113100010202002"></a>

## Next pages — dualstack / 131303201301 / 4

- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-0302023220132221-3313300011221002-2010113021102131-3030302011113300-1100111101221011-2012021103323130-0311003330022212-1112201103102030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022232121130322-3330110311210110-0013100133301220-3032311013303312-0011320313110033-3133320211220022-0021200221232313-2313011131201312"></a>

## IPv4 — IPv4 / 330322132213 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- IPv4

<a id="canonical-3333033012100020-1022313002032303-0010300211113202-0003220022013101-3320331330122012-0110132233201121-1331303222102311-2220300002112100"></a>

Type: `["object", {}]`. Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

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
ipv4 = {}
```

<a id="canonical-3222210102112103-1031022131210321-0202123302332203-1002103033300310-3131230131012200-0213301123020220-3222302232211122-2102202123332021"></a>

## Direct properties — IPv4 / 330322132213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322032031021002-0012311221133332-3203320022303111-2202311320322123-2232332212003321-3123331301300123-2320131032301023-3231101310200322"></a>

## Next pages — IPv4 / 330322132213 / 4

- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-2120203212121232-3132123301211012-1002302311333123-1022322131230020-2232103220003320-2132221031000330-0222013101313223-3222110233021222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210320223332000-1220002000313222-1233123111221201-0021212110130303-1222211330202122-3032323133222223-2103132012003221-2103203333222321"></a>

## IPv6 — IPv6 / 331100233030 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- IPv6

<a id="canonical-2000120313121022-0211201113123303-2301001001301321-1213011100021321-3113101302110023-3312133120223213-3333301232021003-0102122102101133"></a>

Type: `["object", {}]`. Optional.

IPv6 address in colon-separated hexadecimal format.

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
ipv6 = {}
```

<a id="canonical-0030122331230302-2303011011320211-2110020233333103-3310103230033331-0002233002111011-1322130133201323-2333203202112333-0111300311032011"></a>

## Direct properties — IPv6 / 331100233030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231223000030223-0032302330302123-0131301211013130-0333200233233131-1131102123200220-0201010221101222-3303101331033121-2122112311230211"></a>

## Next pages — IPv6 / 331100233030 / 4

- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-2303301110202311-2222032220031020-3121303333031321-2103131332222133-3111111130312021-0320322113112130-0301212122230111-2223203233322012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110020213100303-0022002030020303-3201021110310223-2033220321333231-1113123130210210-2121310202123302-0020002121033221-3333102213211001"></a>

## public_ip — public_ip / 302010123321 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- public_ip

<a id="canonical-0322231233100123-2310333311123033-1230000102300232-1330033213330222-1333232003131230-0002123003130331-3322000212211202-3123230103200103"></a>

Type: `"object"`. list nested block, Optional.

Optional. Public VIP to advertise This field is mutually exclusive with where and address fields.

Upstream description:

Optional. Public VIP to advertise This field is mutually exclusive with where and address fields.

Receipt-pinned upstream constraints:

```json
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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3331210233331333-0230112003020100-3113030201311013-3330213111330231-0132120012003012-0032323323210101-0111230323221310-0130312210103130"></a>

## Direct properties — public_ip / 302010123321 / 3

<a id="canonical-2000031232210111-2320023101323210-3202332133111210-2222012112013111-3113201033020211-1300110011333033-2233312221203311-3333103321200202"></a>

<a id="canonical-0110132202302113-0123311002120113-0312100223323133-1000230101321000-0013323113022031-3021103330202033-0203312213020301-1123210312322033"></a>

## kind property — public_ip / 302010123321 / 4

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

<a id="canonical-2110230131231330-0330202031113012-3100200203102321-1202222032221300-0133103000032311-2000110003220310-1222232233311311-2021311121120100"></a>

<a id="canonical-2220032031210012-3231312213212330-1123213120200313-2033012313311102-2132111313020313-0313011230131203-3203313212320131-0323200013203301"></a>

## name property — public_ip / 302010123321 / 5

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

<a id="canonical-2333100103121203-2002013121310213-3203101313203212-2312123211321321-0131013101322233-1033311230310231-0101033123231000-2213031122100030"></a>

<a id="canonical-2330330203302111-3232033102331203-0131113122103000-1303330303233000-0332133223213132-0000113323002311-2313200132311332-0201013033030331"></a>

## namespace property — public_ip / 302010123321 / 6

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

<a id="canonical-1123311232120211-0333023110231212-2112310200213221-2102031110133323-2013302303333212-1201331102301331-3113212000210011-0223032210302120"></a>

<a id="canonical-2013303220312310-3331000300231223-2002121320223033-2021313003122333-1213230130213110-0330100010201223-1033300000220320-1321112003101112"></a>

## tenant property — public_ip / 302010123321 / 7

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

<a id="canonical-0312031303223112-3210302111310300-2310322313030301-3220131210300300-3130030320113231-0321111132220331-3323312200221100-1302111312123022"></a>

<a id="canonical-2200200222001303-1313013032321023-1022002332003003-0100031332222201-1020231230103030-2030122213110203-3120223000123223-0202311300011333"></a>

## uid property — public_ip / 302010123321 / 8

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

<a id="canonical-0230232302010020-0130033031102030-0102121231023101-0232201313201300-3231330301031200-1320133113003312-2202020330223122-2121012002210013"></a>

## Next pages — public_ip / 302010123321 / 9

- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-1003330112013233-3322213102311100-1200212333211211-1232231112331330-2101230211113313-1332123122003021-2020332033122232-2303333311102233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223013221123113-3011120133310320-2132323300233033-1111311103230332-2113111322301013-3330031032322223-0012101323131202-2231311002322310"></a>

## timeouts — timeouts / 222033123313 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- timeouts

<a id="canonical-2002313031012202-1130333220100233-3002210202311010-1313030223133020-0132023132330100-0333223032310022-3320103121213032-2321221322313232"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310123100102333-1301200113321102-3023002233030122-1201222231232221-2331121321323331-3102201102302333-1131300322120132-2321323202102222"></a>

## Direct properties — timeouts / 222033123313 / 3

<a id="canonical-0231312131331030-2032313323320303-2333201002323331-1313222211100121-2031010302001203-2131213302101003-2303212021323033-2113011031102013"></a>

<a id="canonical-1001121131122033-1021232030013220-1203220001233301-3331300221032031-0031323310132203-3233232102001233-2221330233312333-0133101102212111"></a>

## create property — timeouts / 222033123313 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1110301110210023-2121312300022313-2030122103032202-1001311013311032-0121200031332331-2322311302020322-0023201132210113-0331103133310302"></a>

<a id="canonical-1030221020220122-3311200311231011-1330212212032010-3210111230221212-0320030221210022-3212002313133133-3232311123321333-0030011220112222"></a>

## delete property — timeouts / 222033123313 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0130020130001332-3022220313203022-1210323210113320-3101132023023300-2231333302300120-2322223032003220-1310013001210103-2302002001103310"></a>

<a id="canonical-0112303130220231-3103222102232013-1003303120222210-0132232220000031-1023302302311030-1111331131320330-2202233130203203-2332033131220300"></a>

## read property — timeouts / 222033123313 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0021220131112032-3120022020101220-2020010230100230-3032122123003002-0202102031123330-0321232123301112-0330101210010331-2310023132222100"></a>

<a id="canonical-0201233212303122-0130123323302033-2303011221023313-2022300103302020-0332202030013001-0012031331012002-0310310001131231-1222132122033203"></a>

## update property — timeouts / 222033123313 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3100312310133121-1333010300022123-2333201122112033-3011213322232312-1031023000220003-0013111102002311-2033301302222302-3321113111233310"></a>

## Next pages — timeouts / 222033123313 / 8

- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000030211111101-0121012230302113-0022100021320011-2033323223013022-2321211021213333-0010001031133222-2103122230003032-2222332323303030"></a>

## tls_parameters — tls_parameters / 221023012213 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- tls_parameters

<a id="canonical-2322033321212012-0123203312022323-2311133313230210-1033021100003012-2323312323020011-3131030023300302-0000211313232320-2013001101333330"></a>

Type: `"object"`. single nested block, Optional.

TLS configuration for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("client_certificate_optional",
    "client_certificate_required"),
  validators.ConflictingObjectAttributes("client_certificate_optional",
    "no_client_certificate"),
  validators.ConflictingObjectAttributes("client_certificate_required",
    "no_client_certificate")}
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
  "x-ves-oneof-field-client_certificate_verify_choice": "[\"client_certificate_optional\",\"client_certificate_required\",\"no_client_certificate\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-1120321323100232-0033310021233020-0023002111311022-2001132133030020-1132011002312033-0133322302033231-3033132302012133-0102213222033331"></a>

## Direct properties — tls_parameters / 221023012213 / 3

- [client_certificate_optional](resources--advertise_policy--reference--group-001.md#canonical-3011211000321121-1203211002310022-1301212313121113-1320112310012002-2011003113031221-3102303113123230-1103303232223311-2132302212012103): complete subsection reference.

- [client_certificate_required](resources--advertise_policy--reference--group-001.md#canonical-2123200310020220-1031310312303233-3202210101131333-2123123110203212-3020312111221120-3303011312012110-3231132302220231-2032131302002002): complete subsection reference.

- [common_params](resources--advertise_policy--reference--group-001.md#canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030): complete subsection reference.

- [no_client_certificate](resources--advertise_policy--reference--group-001.md#canonical-1123113020213121-0131331120202031-3312010123111102-1332331131322320-3133313322220110-3103120010013310-2201332023123110-2000112122133021): complete subsection reference.

<a id="canonical-1233320031133133-0133310112210120-0331233330223122-1201133030131311-0310202202100303-0102223210100322-0203100010020011-0322102202210200"></a>

<a id="canonical-3223123102012200-1020132102113333-3001011320112200-3311211032202220-1133002232022032-0331113101023110-1001102001332000-3232331202232231"></a>

## xfcc_header_elements property — tls_parameters / 221023012213 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-2300013231002021-3110012132222132-2332232201003003-3312313131312322-1220231002221211-1333113110033013-2212031131003232-2201201203021030"></a>

## Next pages — tls_parameters / 221023012213 / 5

- [tls_parameters.client_certificate_optional](resources--advertise_policy--reference--group-001.md#canonical-3011211000321121-1203211002310022-1301212313121113-1320112310012002-2011003113031221-3102303113123230-1103303232223311-2132302212012103)
- [tls_parameters.client_certificate_required](resources--advertise_policy--reference--group-001.md#canonical-2123200310020220-1031310312303233-3202210101131333-2123123110203212-3020312111221120-3303011312012110-3231132302220231-2032131302002002)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030)
- [tls_parameters.no_client_certificate](resources--advertise_policy--reference--group-001.md#canonical-1123113020213121-0131331120202031-3312010123111102-1332331131322320-3133313322220110-3103120010013310-2201332023123110-2000112122133021)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-3011211000321121-1203211002310022-1301212313121113-1320112310012002-2011003113031221-3102303113123230-1103303232223311-2132302212012103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201201122211200-1030023112201012-3002001213320201-0000012122020203-2203031331313233-3033111332331131-1023311222110010-2300221122321011"></a>

## tls_parameters.client_certificate_optional — client_certificate_optional / 221022001230 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- tls_parameters.client_certificate_optional

<a id="canonical-2230112232310221-2332321301033302-0332100122213210-0320020223221302-2300102231110201-0100002133300011-1101111132030322-3233310121200230"></a>

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
client_certificate_optional = {}
```

<a id="canonical-2300202303131013-3032312010321331-2310021333210212-3201301312300132-2120102132321000-0133203013200331-0313013233211033-3013222102033333"></a>

## Direct properties — client_certificate_optional / 221022001230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130112332233020-0223320112012223-0113211132223112-0313201202323130-1330231101122323-1310222201123333-1033320300213221-3010203113133033"></a>

## Next pages — client_certificate_optional / 221022001230 / 4

- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-2123200310020220-1031310312303233-3202210101131333-2123123110203212-3020312111221120-3303011312012110-3231132302220231-2032131302002002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013031133212131-3031301203312213-3012110312322021-3330013122321313-2223133121123300-0332213120210201-3233120003123121-0212232132203322"></a>

## tls_parameters.client_certificate_required — client_certificate_required / 132330333200 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- tls_parameters.client_certificate_required

<a id="canonical-3301133022122203-1333010231102033-2010121202133013-0202213202301012-2010311232010101-1200312200312231-3331000231223023-1113230022133033"></a>

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
client_certificate_required = {}
```

<a id="canonical-0310131220033311-2132121221122023-1111310132120311-1002111113312112-3133330320333003-0332120320301103-3101132122132010-3231311001310222"></a>

## Direct properties — client_certificate_required / 132330333200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213200231223303-3002201021110110-0031121333022131-0030002321032333-1332230300111222-3220332021133201-1132201321301110-1002100222102231"></a>

## Next pages — client_certificate_required / 132330333200 / 4

- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203302121011032-2100111103322333-3331221320331231-2320120121332303-3101222021311001-1123323000130023-3033322121312133-2103100013233230"></a>

## tls_parameters.common_params — common_params / 221320121303 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- tls_parameters.common_params

<a id="canonical-1320111103102012-2221131312212233-0031003033120221-0230120322000220-1001331320310323-2102333100202203-2331000213110202-2132122123013301"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
common_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233023020231023-1011121122322302-3303300211321323-3111202332101011-2222131233013331-0010330212203200-2203112210210032-1312220203232233"></a>

## Direct properties — common_params / 221320121303 / 3

<a id="canonical-1010113101221011-3121222102323122-1110011100120222-0311003212301023-0333001301310211-3221202300122321-0132120230311333-2331320203222321"></a>

<a id="canonical-2130132333002201-2023122110010121-2133201212303231-3312202202210102-1311033301320123-3101023022322003-0321002303213301-1210013001323101"></a>

## cipher_suites property — common_params / 221320121303 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-1131000230001312-3132132103111000-1133323003211013-0101311231130201-2202110333312112-3232322021011122-2112020011133322-1120001213011011"></a>

<a id="canonical-1112112201003331-3111020200210133-2002320010013313-1203303001012112-3001301223202122-1021221203203201-0312033220112312-2013033301201110"></a>

## maximum_protocol_version property — common_params / 221320121303 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-2331021023331300-1220033213013301-3333211312033102-3301221203122012-1101122223212202-0010113022303221-2213102331102211-2130301313000003"></a>

<a id="canonical-3000120123033013-0100101301221201-2020330122203320-0112230331122301-2211103000133021-1301331310212323-1111230013212322-2320013202331312"></a>

## minimum_protocol_version property — common_params / 221320121303 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

- [tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-1221101232003200-2111301222311033-3321222332231322-2331002022323230-0200030311030103-0222032212012220-2133123120332003-2000021213101133): complete subsection reference.

- [validation_params](resources--advertise_policy--reference--group-001.md#canonical-1120302121321011-0010323001200020-2112303031220030-2230131220032213-0330011102333312-3132321111013003-3223110311012022-0213313210313002): complete subsection reference.

<a id="canonical-2223112033103023-1111212233221102-3303101201002300-3032032113200331-3301232220301023-0132033103323330-1331221222100130-2130333311003012"></a>

## Next pages — common_params / 221320121303 / 7

- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-1221101232003200-2111301222311033-3321222332231322-2331002022323230-0200030311030103-0222032212012220-2133123120332003-2000021213101133)
- [tls_parameters.common_params.validation_params](resources--advertise_policy--reference--group-001.md#canonical-1120302121321011-0010323001200020-2112303031220030-2230131220032213-0330011102333312-3132321111013003-3223110311012022-0213313210313002)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-1221101232003200-2111301222311033-3321222332231322-2331002022323230-0200030311030103-0222032212012220-2133123120332003-2000021213101133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210203022100222-2131111113001101-0202320221333010-0201331320312213-2320033320012130-1013332211130220-1231212102131111-0220100331210311"></a>

## tls_parameters.common_params.tls_certificates — tls_certificates / 020303323020 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030)
- tls_parameters.common_params.tls_certificates

<a id="canonical-0003122211010203-0123031121123201-2202223100123223-3023020210022033-3113122100212013-2233010323300220-1202332020332010-3133033101102230"></a>

Type: `"object"`. list nested block, Optional.

TLS Certificates. Set of TLS certificates.

Upstream description:

Set of TLS certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203030200230213-1323120120230012-1101300122123002-0011121033110020-0020212332313301-2032303011003101-3211313332100223-1021111330311212"></a>

## Direct properties — tls_certificates / 020303323020 / 3

<a id="canonical-3213023000300032-1310031010033332-0301213222023123-2233323013303233-3220332201001201-2010230232333132-0331011033002030-3022001313311321"></a>

<a id="canonical-1301232200322322-2121222311211301-0211110233321232-3120313012321111-3321003201330313-3232311221300020-1310033122300132-1222311030321213"></a>

## certificate_url property — tls_certificates / 020303323020 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](resources--advertise_policy--reference--group-001.md#canonical-3332322303321230-3131133313232313-1312301101013031-1020112000320130-1312221110111122-3121033333331110-1022331331033112-2320313331213133): complete subsection reference.

<a id="canonical-1012132323002121-0223122230120300-0022210200012032-0011230021232301-0133221030221112-3103213133323211-1131001331130131-0022310221123010"></a>

<a id="canonical-3220020020200302-1312322220300102-1103000211210122-2120222003111013-0333211330131131-1001032121320111-3110002003023030-3230030301011232"></a>

## description_spec property — tls_certificates / 020303323020 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--advertise_policy--reference--group-001.md#canonical-1202331000100033-1020121133121110-2302332233223000-1121022331001132-0212212320203113-3301322222011312-1031231113220210-3230012320320310): complete subsection reference.

- [private_key](resources--advertise_policy--reference--group-001.md#canonical-0101122331121313-1011022321120120-1211301212301131-2133130302011133-1231000333133202-2332001321021330-1000310130030230-1121112221132202): complete subsection reference.

- [use_system_defaults](resources--advertise_policy--reference--group-001.md#canonical-1100113010020100-0120101113323322-1322320201022021-3331332012322122-2002021222301231-1003000102013223-2323210121330021-0223021013300112): complete subsection reference.

<a id="canonical-1331202331101013-3322213003330100-2232201231131222-2330302210021121-2222023332230233-3332231311123131-2220331000130133-1300221002221200"></a>

## Next pages — tls_certificates / 020303323020 / 6

- [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--advertise_policy--reference--group-001.md#canonical-3332322303321230-3131133313232313-1312301101013031-1020112000320130-1312221110111122-3121033333331110-1022331331033112-2320313331213133)
- [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--advertise_policy--reference--group-001.md#canonical-1202331000100033-1020121133121110-2302332233223000-1121022331001132-0212212320203113-3301322222011312-1031231113220210-3230012320320310)
- [tls_parameters.common_params.tls_certificates.private_key](resources--advertise_policy--reference--group-001.md#canonical-0101122331121313-1011022321120120-1211301212301131-2133130302011133-1231000333133202-2332001321021330-1000310130030230-1121112221132202)
- [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--advertise_policy--reference--group-001.md#canonical-1100113010020100-0120101113323322-1322320201022021-3331332012322122-2002021222301231-1003000102013223-2323210121330021-0223021013300112)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-3332322303321230-3131133313232313-1312301101013031-1020112000320130-1312221110111122-3121033333331110-1022331331033112-2320313331213133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301032030001021-3110212303210300-0330312130323300-1213113032301030-3110131200312310-1213021203031113-3322301023032223-3230322131023113"></a>

## tls_parameters.common_params.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 210222021200 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-1221101232003200-2111301222311033-3321222332231322-2331002022323230-0200030311030103-0222032212012220-2133123120332003-2000021213101133)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-2330202322013132-2011000323203210-3223131220323130-1022230313222313-1132101302302123-0320321222110230-2023333112310220-1231021202120310"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230211220313331-0321220011331000-1332022011210122-2113231103332320-1031123121122203-3013133333201210-2222033110022200-1323020123000223"></a>

## Direct properties — custom_hash_algorithms / 210222021200 / 3

<a id="canonical-3010333002010132-2012131321021310-3000223320002233-1322103033131000-0331312321323333-0000112133003120-2020100020230301-1021111031013030"></a>

<a id="canonical-0032331332232033-0011323012021021-3310131230031022-2313131313003120-0012233122320200-2110321113030310-2203310033213103-0021213000203323"></a>

## hash_algorithms property — custom_hash_algorithms / 210222021200 / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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

<a id="canonical-3023103103000212-2303010230002322-3221221000112300-3021013021011200-1112232112210212-0023032023031200-3021300110332312-0202322001222312"></a>

## Next pages — custom_hash_algorithms / 210222021200 / 5

- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-1221101232003200-2111301222311033-3321222332231322-2331002022323230-0200030311030103-0222032212012220-2133123120332003-2000021213101133)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-1202331000100033-1020121133121110-2302332233223000-1121022331001132-0212212320203113-3301322222011312-1031231113220210-3230012320320310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231320131033012-0331100011312203-1333132221311113-2120032102331233-1022223203322310-3130223210103000-1013012301010020-2233030030232002"></a>

## tls_parameters.common_params.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 103313002232 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-1221101232003200-2111301222311033-3321222332231322-2331002022323230-0200030311030103-0222032212012220-2133123120332003-2000021213101133)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-3101320302030032-2301301101223203-2231131301020021-2321001212230332-2211333302222232-2320023203002113-2322022112031030-2333121000230030"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

<a id="canonical-1031111300220033-0002030102210002-3223020212130020-2333033202322322-1331222120131310-3221132331231022-0130323121300202-1231300002130301"></a>

## Direct properties — disable_ocsp_stapling / 103313002232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222013131110231-2322210313212032-2010313130230013-2311032120012112-1012120322220001-3231132131323232-3231201001303311-2330022032003203"></a>

## Next pages — disable_ocsp_stapling / 103313002232 / 4

- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-1221101232003200-2111301222311033-3321222332231322-2331002022323230-0200030311030103-0222032212012220-2133123120332003-2000021213101133)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-0101122331121313-1011022321120120-1211301212301131-2133130302011133-1231000333133202-2332001321021330-1000310130030230-1121112221132202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103311130033132-0030332112222003-2123313100002222-1232113012020021-2332310020321222-0231320211123112-2210313330233231-0303022000301110"></a>

## tls_parameters.common_params.tls_certificates.private_key — private_key / 102222011100 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-1221101232003200-2111301222311033-3321222332231322-2331002022323230-0200030311030103-0222032212012220-2133123120332003-2000021213101133)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-1211331231120110-0121133221133113-0133113130201132-3333213330020022-1013231203221012-1033011130313300-2112232312232021-1102000212132001"></a>

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230012310200232-0333302132102013-1122103321200213-1101132102312112-2101000112032022-0320000322312013-3200113011230232-3112020030100021"></a>

## Direct properties — private_key / 102222011100 / 3

- [blindfold_secret_info](resources--advertise_policy--reference--group-001.md#canonical-2123131112031212-2022030331030321-3330313030013123-3312333332222110-2132331320320331-0333310023230110-2013022010121303-1301022022323110): complete subsection reference.

- [clear_secret_info](resources--advertise_policy--reference--group-001.md#canonical-0311233030101120-3301012122101211-1001211010310033-3202222101031320-1130011011323302-2333131322033000-3222330121202300-2023020322211001): complete subsection reference.

<a id="canonical-3332303021332131-1213133110011322-3100223303003303-0113300100130233-2120323323331030-1110201031323030-3102310001210311-3200012301010203"></a>

## Next pages — private_key / 102222011100 / 4

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--advertise_policy--reference--group-001.md#canonical-2123131112031212-2022030331030321-3330313030013123-3312333332222110-2132331320320331-0333310023230110-2013022010121303-1301022022323110)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--advertise_policy--reference--group-001.md#canonical-0311233030101120-3301012122101211-1001211010310033-3202222101031320-1130011011323302-2333131322033000-3222330121202300-2023020322211001)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-1221101232003200-2111301222311033-3321222332231322-2331002022323230-0200030311030103-0222032212012220-2133123120332003-2000021213101133)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-2123131112031212-2022030331030321-3330313030013123-3312333332222110-2132331320320331-0333310023230110-2013022010121303-1301022022323110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302320311020130-1132313221310321-1113200332220033-0030112131010312-0201031011332011-0300333212212201-2311110130331110-3323002201212130"></a>

## tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 010223021333 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-1221101232003200-2111301222311033-3321222332231322-2331002022323230-0200030311030103-0222032212012220-2133123120332003-2000021213101133)
- [tls_parameters.common_params.tls_certificates.private_key](resources--advertise_policy--reference--group-001.md#canonical-0101122331121313-1011022321120120-1211301212301131-2133130302011133-1231000333133202-2332001321021330-1000310130030230-1121112221132202)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0020203220332313-2320022102201012-1122112033301323-0311213031122013-1213132110020103-1310011120123030-1122320210322103-2202030100131011"></a>

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

<a id="canonical-2302023330311311-0200132213023232-3120321321303203-1330320231110122-3220001100013120-3313022220222131-3320203133112002-3000032001211133"></a>

## Direct properties — blindfold_secret_info / 010223021333 / 3

<a id="canonical-0003101032201110-0133032222033023-0330303033203120-3202300312301130-2132122322020203-3231313230311121-3012123200332000-3321023101313200"></a>

<a id="canonical-0120330320031102-3211202332321012-3211133220012203-2033022103200121-0033123221310332-0031132111311112-0031303203200113-0202131312011032"></a>

## decryption_provider property — blindfold_secret_info / 010223021333 / 4

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

<a id="canonical-3203000102112331-0223011000311313-2033223123333301-1012023113332200-2000202322111123-3013103210010300-0230211101330211-2023013203111221"></a>

<a id="canonical-0130202122320000-0123311310201112-3331223220210001-3120133200023132-1113231323001002-1213032111301310-0112131131001122-2331003311031330"></a>

## location property — blindfold_secret_info / 010223021333 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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

<a id="canonical-1323230120000130-3223232322231230-1010232310233120-3232232300011222-1201001112320112-1133330100030311-1232102111310102-1002132320001123"></a>

<a id="canonical-0312332011002313-1312030002300120-1123222332031100-3000221311332232-1000010023010113-3123330030220230-1203131013120320-0011101102013130"></a>

## store_provider property — blindfold_secret_info / 010223021333 / 6

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

<a id="canonical-2132233311110223-0022010130200312-1201120311202323-2230131321311210-2230032332220011-3021203230203130-1001113201311221-0131213331023323"></a>

## Next pages — blindfold_secret_info / 010223021333 / 7

- [tls_parameters.common_params.tls_certificates.private_key](resources--advertise_policy--reference--group-001.md#canonical-0101122331121313-1011022321120120-1211301212301131-2133130302011133-1231000333133202-2332001321021330-1000310130030230-1121112221132202)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-0311233030101120-3301012122101211-1001211010310033-3202222101031320-1130011011323302-2333131322033000-3222330121202300-2023020322211001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023030132310220-3312330233312022-2230330302200231-2131110101230133-2230132230100022-2230101333311231-3303003000221302-1003012323203201"></a>

## tls_parameters.common_params.tls_certificates.private_key.clear_secret_info — clear_secret_info / 323003003122 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-1221101232003200-2111301222311033-3321222332231322-2331002022323230-0200030311030103-0222032212012220-2133123120332003-2000021213101133)
- [tls_parameters.common_params.tls_certificates.private_key](resources--advertise_policy--reference--group-001.md#canonical-0101122331121313-1011022321120120-1211301212301131-2133130302011133-1231000333133202-2332001321021330-1000310130030230-1121112221132202)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-0010130132221212-3101213232230032-2302122333313122-2003121010313103-1010322213223232-1211310321112311-1123130321311133-3002112031323120"></a>

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

<a id="canonical-0103321203221103-0003101230332003-2323220023320210-1100010120130122-3132131331020132-3220122130002133-0121101232313332-0113231120122103"></a>

## Direct properties — clear_secret_info / 323003003122 / 3

<a id="canonical-0231221333201302-3222220123310302-2002130300020111-1321322122020232-1332302202002033-0020320211210100-0123331321203230-2331010133103231"></a>

<a id="canonical-3312330311200330-2333130212233231-0323120031213222-0101132301133113-0200133213111313-1302230032220222-0133333132311121-1120120202201232"></a>

## provider_ref property — clear_secret_info / 323003003122 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0020323313003030-3303101303123303-0303123033330200-3010203333112030-3123013021032233-3023330021303211-0102202220021121-0303113102130303"></a>

<a id="canonical-0201333222031200-1233000121232330-1113121102220322-3221221220210013-1112020230301221-3302002023011022-0211100011120310-0302033012231301"></a>

## URL property — clear_secret_info / 323003003122 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-3011111033121011-1001123010232212-3323203323103131-1100333102131332-2012333103000222-3113002330323010-3200300000133300-3122003010002031"></a>

## Next pages — clear_secret_info / 323003003122 / 6

- [tls_parameters.common_params.tls_certificates.private_key](resources--advertise_policy--reference--group-001.md#canonical-0101122331121313-1011022321120120-1211301212301131-2133130302011133-1231000333133202-2332001321021330-1000310130030230-1121112221132202)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-1100113010020100-0120101113323322-1322320201022021-3331332012322122-2002021222301231-1003000102013223-2323210121330021-0223021013300112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002030322200233-1001100023201120-0030312130022100-0222331032012320-0002033312113322-0011120303331301-0001032130003230-2330031233303130"></a>

## tls_parameters.common_params.tls_certificates.use_system_defaults — use_system_defaults / 010002021032 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-1221101232003200-2111301222311033-3321222332231322-2331002022323230-0200030311030103-0222032212012220-2133123120332003-2000021213101133)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-2122011212003312-3010223102213213-3033113230010301-1011013110112030-1121203131232113-0300312301121111-1221132313012302-2333231013113111"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_system_defaults = {}
```

<a id="canonical-0302012010023021-3113221001203223-0102330330103333-1103031213311300-3002313202232221-3220013031013102-1023231003022022-0013012123121320"></a>

## Direct properties — use_system_defaults / 010002021032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300010130013031-3112223201202100-3031231033211301-0120003100230200-0130332220003330-2030223031221213-1031200211233132-3120130203101112"></a>

## Next pages — use_system_defaults / 010002021032 / 4

- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-1221101232003200-2111301222311033-3321222332231322-2331002022323230-0200030311030103-0222032212012220-2133123120332003-2000021213101133)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-1120302121321011-0010323001200020-2112303031220030-2230131220032213-0330011102333312-3132321111013003-3223110311012022-0213313210313002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320333131233101-0210220133032202-2310231122311120-2231012120303103-2212012221012012-3232101112322013-2313033300320002-1313113012230012"></a>

## tls_parameters.common_params.validation_params — validation_params / 023112103122 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030)
- tls_parameters.common_params.validation_params

<a id="canonical-0023311131032113-2023102313002030-1312021330311012-1103230212233201-3330123232001131-3321020222132111-0112022322321033-0100131022222012"></a>

Type: `"object"`. single nested block, Optional.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010222010003310-0011123021002101-1321001123312012-1322001303012001-3010011233301303-3122322301312100-2020301222230322-2223303323320002"></a>

## Direct properties — validation_params / 023112103122 / 3

<a id="canonical-2101303300331022-0212022132111333-0003213021330111-3221200121002300-2133111103301133-2002003132022211-1123300011022212-0313223231230222"></a>

<a id="canonical-2123133311012011-1032123022213223-3211101122112033-0333322022301330-3113121101021202-3220213110002311-0100003020122033-0033223111003022"></a>

## skip_hostname_verification property — validation_params / 023112103122 / 4

Type: `"bool"`. Optional.

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

- [trusted_ca](resources--advertise_policy--reference--group-001.md#canonical-0233330101330001-1132312012321002-1103300322201031-2211112112212211-3120032310213123-2303211230121330-1220132133003132-1001031231101202): complete subsection reference.

<a id="canonical-3100013313021322-0331210001033200-1030220301112333-3210303301003211-0001220311310002-0111302021020232-0223033303211132-1300113232100022"></a>

<a id="canonical-0012331130323021-2221223333003212-1331223312230211-2023303113102231-1230131103113323-1303111303003001-1301001232323103-3323211200130311"></a>

## trusted_ca_url property — validation_params / 023112103122 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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

<a id="canonical-1020000103010302-3100220002203322-3231301330113330-1130332232323311-3101012233221213-2113302321300022-3100331023130220-1203022033222332"></a>

<a id="canonical-3201001200131013-1312011320132330-3232301221322101-0131232300133300-1203033003332220-3220023312121033-2100100130212101-0212230330231200"></a>

## verify_subject_alt_names property — validation_params / 023112103122 / 6

Type: `["list", "string"]`. Optional.

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

<a id="canonical-0212230022232002-2221222303032103-3323320021322222-1220310203102221-3301023000022320-0112231002233120-0032203011101112-0220302310332122"></a>

## Next pages — validation_params / 023112103122 / 7

- [tls_parameters.common_params.validation_params.trusted_ca](resources--advertise_policy--reference--group-001.md#canonical-0233330101330001-1132312012321002-1103300322201031-2211112112212211-3120032310213123-2303211230121330-1220132133003132-1001031231101202)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-0233330101330001-1132312012321002-1103300322201031-2211112112212211-3120032310213123-2303211230121330-1220132133003132-1001031231101202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201020002231002-3220030233221113-2211033333311232-1132332213333111-2331132033213230-2212221122021110-3030331302120122-0002121223311323"></a>

## tls_parameters.common_params.validation_params.trusted_ca — trusted_ca / 233203032030 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030)
- [tls_parameters.common_params.validation_params](resources--advertise_policy--reference--group-001.md#canonical-1120302121321011-0010323001200020-2112303031220030-2230131220032213-0330011102333312-3132321111013003-3223110311012022-0213313210313002)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-0230311120310233-0331213013013011-3223010101133021-2220233232000201-3222330330113032-3233111100003120-0012100332000222-0302213300330103"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233132231120022-1110231301021312-1330132212000201-0011331112320111-0211012232030120-0133212100132012-1102003223200322-1132133301110100"></a>

## Direct properties — trusted_ca / 233203032030 / 3

- [trusted_ca_list](resources--advertise_policy--reference--group-001.md#canonical-0032331212103230-1031220333120210-0133210331033101-1220300001201202-3102201000223101-2313011310330002-0212011312132223-0202011131112302): complete subsection reference.

<a id="canonical-0200333123313330-0100222323213221-0232322233032112-0210322132321310-1301012310321002-0322323301201311-1331010311303131-0200322323023331"></a>

## Next pages — trusted_ca / 233203032030 / 4

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--advertise_policy--reference--group-001.md#canonical-0032331212103230-1031220333120210-0133210331033101-1220300001201202-3102201000223101-2313011310330002-0212011312132223-0202011131112302)
- [tls_parameters.common_params.validation_params](resources--advertise_policy--reference--group-001.md#canonical-1120302121321011-0010323001200020-2112303031220030-2230131220032213-0330011102333312-3132321111013003-3223110311012022-0213313210313002)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-0032331212103230-1031220333120210-0133210331033101-1220300001201202-3102201000223101-2313011310330002-0212011312132223-0202011131112302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023011323003213-0123213330300232-3131220220202130-3320131222013230-2112333230313133-2113031302223303-0233113331310310-0212311133011002"></a>

## tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list — trusted_ca_list / 331200230232 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-0313021030032323-0110121003001310-2011311011311100-1001212322011312-2111223033130030-2203003310213020-0032233221222130-1033202003312030)
- [tls_parameters.common_params.validation_params](resources--advertise_policy--reference--group-001.md#canonical-1120302121321011-0010323001200020-2112303031220030-2230131220032213-0330011102333312-3132321111013003-3223110311012022-0213313210313002)
- [tls_parameters.common_params.validation_params.trusted_ca](resources--advertise_policy--reference--group-001.md#canonical-0233330101330001-1132312012321002-1103300322201031-2211112112212211-3120032310213123-2303211230121330-1220132133003132-1001031231101202)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-3133013331232120-1113311313030300-0103123012122331-0121012320021013-1202330112001233-1130300112310033-0111122000222223-2211033120233301"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100021312020020-0232132202013212-2020031113333101-0103123300331321-0221331302331302-0112113121032122-2221221031303212-1122001100311201"></a>

## Direct properties — trusted_ca_list / 331200230232 / 3

<a id="canonical-1223112012110021-3123301123021033-0211033313032231-3222200222102003-1001332103112013-2131123232120303-2321112331121020-1102003101022101"></a>

<a id="canonical-1223001010212202-0223011211000113-1213322233331102-2331232223130001-1100031110032223-1230230121130032-0311203211321111-3103102003230313"></a>

## kind property — trusted_ca_list / 331200230232 / 4

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

<a id="canonical-3133213323300222-3201221220103012-2313330001330231-1101123121103122-0311210110332012-1120322120212023-1212230130032231-3020132302020330"></a>

<a id="canonical-1231030032021020-1200300003320011-0221210101130012-0213101012313202-0101300232322122-3213133122310123-1120023012311323-2032002302220303"></a>

## name property — trusted_ca_list / 331200230232 / 5

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

<a id="canonical-2120300333312032-2012212012221312-0133211303121211-1002020113330221-1101103230012321-1111111103111203-3032020321113131-2231002131210010"></a>

<a id="canonical-3233002120303100-1102200300122220-2302300300222022-3001021211010030-2021003003001201-2102223312113332-2031120111213130-0232100201311212"></a>

## namespace property — trusted_ca_list / 331200230232 / 6

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

<a id="canonical-1210302201120303-3023132001001011-3321203201233201-0320001313032222-0332231021212031-1122233131211123-2023021033030313-0233130222333232"></a>

<a id="canonical-2011232023212220-1030202102322111-3023310330103220-1333300332101230-2302032301033301-0231033222313302-1113232031033330-1330201100230323"></a>

## tenant property — trusted_ca_list / 331200230232 / 7

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

<a id="canonical-2223201230332113-3333212213200232-3002012023222031-2022213012002003-3111202223223200-1010231333111303-2031201210022233-3223031001113012"></a>

<a id="canonical-1231130130201200-0001020110203200-1000303023201032-1333331101013132-2122302031320130-3123230022020332-2012310211221133-0013222202131230"></a>

## uid property — trusted_ca_list / 331200230232 / 8

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

<a id="canonical-2321312322202311-0321022333320003-3132322333023313-3232100211233032-3232210003110223-3122232023213210-2130010200331310-0331130321311311"></a>

## Next pages — trusted_ca_list / 331200230232 / 9

- [tls_parameters.common_params.validation_params.trusted_ca](resources--advertise_policy--reference--group-001.md#canonical-0233330101330001-1132312012321002-1103300322201031-2211112112212211-3120032310213123-2303211230121330-1220132133003132-1001031231101202)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-1123113020213121-0131331120202031-3312010123111102-1332331131322320-3133313322220110-3103120010013310-2201332023123110-2000112122133021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120313012301011-3121133123131222-3313120301231010-1102301021223201-0222301203031012-2120031032330001-2132201111120223-3301021032132130"></a>

## tls_parameters.no_client_certificate — no_client_certificate / 011230100232 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- tls_parameters.no_client_certificate

<a id="canonical-3201021233032203-1131121110130212-3130222021212132-1313003032021023-3300030122131123-3103133220202332-1032301233313312-1133201001313001"></a>

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
no_client_certificate = {}
```

<a id="canonical-1303013101201200-3212223311303103-3212032220102111-3003210101132201-0112201303123121-2113101213010212-0323223203033230-1200333002030213"></a>

## Direct properties — no_client_certificate / 011230100232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100233033231101-1033121133220310-2033323230023010-1101023200320103-0010322012121132-0213320220233113-2232200201223322-0303200200310311"></a>

## Next pages — no_client_certificate / 011230100232 / 4

- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203121320302010-0100211103021212-1030230230001100-3332002001310012-1113130100133032-2310101210112032-2013331202013333-0313103223011321"></a>

## where — where / 100110212303 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- where

<a id="canonical-2312133200212030-0033010301322001-2210313312222202-0112112201322210-3201200100320030-1230333102333202-3002122110312113-1313332322213203"></a>

Type: `"object"`. single nested block, Optional.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local..

Upstream description:

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingObjectAttributes("virtual_network",
    "virtual_site")}
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
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
where {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330002221213332-1022232320231232-2323110112101233-0331210300200322-3311112321132300-0230113303213211-0022130122000011-1022331120321000"></a>

## Direct properties — where / 100110212303 / 3

- [site](resources--advertise_policy--reference--group-001.md#canonical-0222101210300200-0323002322120223-1311001311133220-1332231310102120-2023232323202303-1231311002212302-2202201032210023-3331230223123020): complete subsection reference.

- [virtual_network](resources--advertise_policy--reference--group-001.md#canonical-3122020011221101-1023022311233310-1101232302222111-0103003032303222-0323233201231323-3133132331222122-0301313303200101-1121023221321331): complete subsection reference.

- [virtual_site](resources--advertise_policy--reference--group-001.md#canonical-3001022020113000-1323122102210321-0123111102031203-2122012110331032-1230222231110302-2322320110032111-2033032202032030-0000113201220002): complete subsection reference.

<a id="canonical-1132101311033103-2331231121301220-0321130320021131-3203312120021210-1310121310023200-1000130321213103-1202103012112122-3222213202300222"></a>

## Next pages — where / 100110212303 / 4

- [where.site](resources--advertise_policy--reference--group-001.md#canonical-0222101210300200-0323002322120223-1311001311133220-1332231310102120-2023232323202303-1231311002212302-2202201032210023-3331230223123020)
- [where.virtual_network](resources--advertise_policy--reference--group-001.md#canonical-3122020011221101-1023022311233310-1101232302222111-0103003032303222-0323233201231323-3133132331222122-0301313303200101-1121023221321331)
- [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-3001022020113000-1323122102210321-0123111102031203-2122012110331032-1230222231110302-2322320110032111-2033032202032030-0000113201220002)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-0222101210300200-0323002322120223-1311001311133220-1332231310102120-2023232323202303-1231311002212302-2202201032210023-3331230223123020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233321230331131-2203322103021331-1032132310110110-0003033312113203-1130103000332333-2133233200110003-1303201320332200-2033333030321221"></a>

## where.site — site / 332203213010 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101)
- where.site

<a id="canonical-1310001132210023-1202201231300231-3111301133233310-1012030312122010-0301300102130001-2121022313313202-3202210113220000-0323333030031033"></a>

Type: `"object"`. single nested block, Optional.

Specifies a direct reference to a site configuration object.

Upstream description:

This specifies a direct reference to a site configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201201322010221-1013231033022303-0231201222000211-3001010231110320-3212230032311112-1310003101221021-2030321010230030-0203013113213200"></a>

## Direct properties — site / 332203213010 / 3

- [disable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-1320002122002132-3332122222121311-3331310122113332-0110301122000222-2000320302200120-1330303001201010-2032321201033131-2313102010000122): complete subsection reference.

- [enable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-2330210001212231-3011112101111032-1032201303001301-0213220033303121-1002110123331330-2003130213231012-3332103222133001-3012312020012332): complete subsection reference.

<a id="canonical-1220222222233121-1202332002311222-0202322303123012-0303331320313322-3120111012030132-1310132221113232-0133222221031231-3101010222103331"></a>

<a id="canonical-1100033001223011-0130303101002322-0320000232330233-3123213122233123-2220010023301313-1113000313103003-2232102312303331-2320121110301201"></a>

## network_type property — site / 332203213010 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
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
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

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

- [ref](resources--advertise_policy--reference--group-001.md#canonical-3231123131130331-2122200032111122-1320131300022101-3210333203223112-0232302012133032-2001011033110030-3030101123133002-2130130331031321): complete subsection reference.

<a id="canonical-2333303321012330-1213011020021023-1102112100111033-1132003220313311-1320311123220321-2210231301231121-0233300222333300-2102220121232333"></a>

## Next pages — site / 332203213010 / 5

- [where.site.disable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-1320002122002132-3332122222121311-3331310122113332-0110301122000222-2000320302200120-1330303001201010-2032321201033131-2313102010000122)
- [where.site.enable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-2330210001212231-3011112101111032-1032201303001301-0213220033303121-1002110123331330-2003130213231012-3332103222133001-3012312020012332)
- [where.site.ref](resources--advertise_policy--reference--group-001.md#canonical-3231123131130331-2122200032111122-1320131300022101-3210333203223112-0232302012133032-2001011033110030-3030101123133002-2130130331031321)
- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-1320002122002132-3332122222121311-3331310122113332-0110301122000222-2000320302200120-1330303001201010-2032321201033131-2313102010000122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023231230302333-1231133323301233-1023030001222222-1010101133222020-1200131032221013-1130212301312031-0122030223302313-2223303332102112"></a>

## where.site.disable_internet_vip — disable_internet_vip / 122312301101 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101)
- [where.site](resources--advertise_policy--reference--group-001.md#canonical-0222101210300200-0323002322120223-1311001311133220-1332231310102120-2023232323202303-1231311002212302-2202201032210023-3331230223123020)
- where.site.disable_internet_vip

<a id="canonical-2220213230232012-1323231312313000-2222113003130033-0030303303301111-3213130222032121-3213030233332300-1313001133133321-0333110230321302"></a>

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
disable_internet_vip = {}
```

<a id="canonical-0303111131222201-2201301301311200-0323203020321231-1003020221333320-3012131133201222-0010100221332202-2013323313233112-1003201232101000"></a>

## Direct properties — disable_internet_vip / 122312301101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221002201001120-2213310102010023-3033300311022203-1002002003231011-1031011232330210-0232021100311030-2022302013121031-2003222122021202"></a>

## Next pages — disable_internet_vip / 122312301101 / 4

- [where.site](resources--advertise_policy--reference--group-001.md#canonical-0222101210300200-0323002322120223-1311001311133220-1332231310102120-2023232323202303-1231311002212302-2202201032210023-3331230223123020)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-2330210001212231-3011112101111032-1032201303001301-0213220033303121-1002110123331330-2003130213231012-3332103222133001-3012312020012332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211321032112230-1012001231003231-0120123021023311-2013221223202010-1100210321313320-2202200211023111-1102112131210310-1113300021123331"></a>

## where.site.enable_internet_vip — enable_internet_vip / 002330211211 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101)
- [where.site](resources--advertise_policy--reference--group-001.md#canonical-0222101210300200-0323002322120223-1311001311133220-1332231310102120-2023232323202303-1231311002212302-2202201032210023-3331230223123020)
- where.site.enable_internet_vip

<a id="canonical-2032010021233201-0013323200302020-2203110030031121-3212110003110223-0133223113112321-0323122220213313-2302031001301021-1301132103312103"></a>

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
enable_internet_vip = {}
```

<a id="canonical-2030212220301232-0122112203211032-3003011131313123-3323321211211111-1213102210230201-3202101222003011-0312230323031111-0300113222013031"></a>

## Direct properties — enable_internet_vip / 002330211211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133203203131331-2300111331023100-3001212110311110-0021102130101323-1133333203121112-3221213200122030-0012221101022003-2232200112121303"></a>

## Next pages — enable_internet_vip / 002330211211 / 4

- [where.site](resources--advertise_policy--reference--group-001.md#canonical-0222101210300200-0323002322120223-1311001311133220-1332231310102120-2023232323202303-1231311002212302-2202201032210023-3331230223123020)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-3231123131130331-2122200032111122-1320131300022101-3210333203223112-0232302012133032-2001011033110030-3030101123133002-2130130331031321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032222302313001-2333021031200001-3203132011033100-0130220030311310-0002132303110230-0103001123100313-2201133110021022-1221311313202020"></a>

## where.site.ref — ref / 113012301220 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101)
- [where.site](resources--advertise_policy--reference--group-001.md#canonical-0222101210300200-0323002322120223-1311001311133220-1332231310102120-2023232323202303-1231311002212302-2202201032210023-3331230223123020)
- where.site.ref

<a id="canonical-3030103310312132-3200000222223211-1011022113120203-1200212113013300-2012333213120223-2313331321000303-3133232203202303-1302323021010223"></a>

Type: `"object"`. list nested block, Optional.

Reference. A site direct reference.

Upstream description:

A site direct reference.

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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310030001323110-0012033121132332-0030003020000002-3021313220230330-2321223322013231-1011001230320321-0213000300300033-3020031233111203"></a>

## Direct properties — ref / 113012301220 / 3

<a id="canonical-1233003131220223-0100303002020101-3221100013111222-2201013200332113-1222322003332111-2132201320200232-1133201012131211-0332233202300022"></a>

<a id="canonical-1100200201131312-2312313010201032-1022232111021233-3223130110000030-1331122320012310-1121023130313300-3030222132210333-1003030031213020"></a>

## kind property — ref / 113012301220 / 4

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

<a id="canonical-1311123313333221-2111023213021102-2332303013203010-2221123313301323-0101003201203001-2122001122213223-2011202021330003-2022332112203101"></a>

<a id="canonical-0311112333331021-3120222013210223-2220022303313023-0322303111130121-3211022103320201-1313231022233232-2302120231202232-0120230000231310"></a>

## name property — ref / 113012301220 / 5

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

<a id="canonical-3313223230232311-3111111110030112-3012110213200131-3030202323122020-1230030111013033-3103132232211232-3303202100022120-1220203021221012"></a>

<a id="canonical-3321013101101002-0321033130002100-2030123220331101-3101112013101113-3213233122333201-1310331033222113-2301212023332122-3303120301020313"></a>

## namespace property — ref / 113012301220 / 6

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

<a id="canonical-0012220102112303-2022220330010323-3113210011012030-3223023211311103-1223223213121100-2333112030300310-1322311130130211-1111110211122101"></a>

<a id="canonical-1322023100300103-3022032121023312-2032231223001022-0031331011012021-0030320002130112-1130233021203133-0033320120031012-0011313021112232"></a>

## tenant property — ref / 113012301220 / 7

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

<a id="canonical-3112202112132001-1200332300222002-0202311103122021-1312011113113030-3131220330021110-1230122210311101-2323020323203023-2131200020113131"></a>

<a id="canonical-2211233113010101-2131303331012333-0003023233000120-0310122013131300-0113110130122231-3031113133132301-1222030200132023-1001301002210132"></a>

## uid property — ref / 113012301220 / 8

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

<a id="canonical-2300122320003012-2131322003221323-2031101303130310-1003323120332200-2103101313121221-0322220301331110-0212120112223222-0212031312331202"></a>

## Next pages — ref / 113012301220 / 9

- [where.site](resources--advertise_policy--reference--group-001.md#canonical-0222101210300200-0323002322120223-1311001311133220-1332231310102120-2023232323202303-1231311002212302-2202201032210023-3331230223123020)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-3122020011221101-1023022311233310-1101232302222111-0103003032303222-0323233201231323-3133132331222122-0301313303200101-1121023221321331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102031031222103-0121313101331320-3001132010232311-0322010111231033-0003101121032001-0301230323202322-3201032312313023-3000323000312130"></a>

## where.virtual_network — virtual_network / 313312213000 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101)
- where.virtual_network

<a id="canonical-2021323333110122-3000213021021112-0321232311300030-2222233203023223-1020203211232213-2101113320333231-1012013111121211-3001232130321101"></a>

Type: `"object"`. single nested block, Optional.

Specifies a direct reference to a network configuration object.

Upstream description:

This specifies a direct reference to a network configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref")}
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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-1002331133133033-2032232301220111-1202101111010113-2023303130320133-2110330013033103-3110301003022230-2023010000113301-3233120122031223"></a>

## Direct properties — virtual_network / 313312213000 / 3

- [ref](resources--advertise_policy--reference--group-001.md#canonical-1320300130123330-1221122030023102-3210312132212222-1302022121221323-2223101301311012-3022101003211323-1101301330331223-2231131221000202): complete subsection reference.

<a id="canonical-2121322103112230-0023013220002012-3312102333132201-0233313332122121-0122300200132020-2003333122002201-0332032222300233-0210331221121321"></a>

## Next pages — virtual_network / 313312213000 / 4

- [where.virtual_network.ref](resources--advertise_policy--reference--group-001.md#canonical-1320300130123330-1221122030023102-3210312132212222-1302022121221323-2223101301311012-3022101003211323-1101301330331223-2231131221000202)
- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-1320300130123330-1221122030023102-3210312132212222-1302022121221323-2223101301311012-3022101003211323-1101301330331223-2231131221000202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330013133120131-1221322010223003-3100220330011332-2300230230111112-0021123213312103-2110012210132311-0212320212333010-3302213020011310"></a>

## where.virtual_network.ref — ref / 003322213002 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101)
- [where.virtual_network](resources--advertise_policy--reference--group-001.md#canonical-3122020011221101-1023022311233310-1101232302222111-0103003032303222-0323233201231323-3133132331222122-0301313303200101-1121023221321331)
- where.virtual_network.ref

<a id="canonical-0023133322232003-2223032020333012-1123000211033333-3030310131033333-3222220210133331-3031111120202000-3120300321033021-2331003313002131"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual network direct reference.

Upstream description:

A virtual network direct reference.

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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030201112010231-1130000202310000-3211312132131001-0001310031333313-0232132133023023-2130023202333200-1130213123322031-1000101321032201"></a>

## Direct properties — ref / 003322213002 / 3

<a id="canonical-2112000320023012-0332312113113012-3031130311111300-1103022032203321-1030101233332022-3233021100113032-2310131202230231-0022222012332133"></a>

<a id="canonical-3230322100111302-2211111203330200-0321200303223123-2210330113222313-0023002023101330-1222001101100020-2133300021302132-0123202011212030"></a>

## kind property — ref / 003322213002 / 4

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

<a id="canonical-2011223333132110-0130122330012221-2330001012212210-1030033013321210-1111132133322100-2123010112201020-0223120322003231-3113012232132212"></a>

<a id="canonical-2301023312333313-0120313201232321-0133111302012023-0333313112102020-3320200211102120-1121233213021221-2210321130120223-2101100302011122"></a>

## name property — ref / 003322213002 / 5

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

<a id="canonical-2113022213330022-2020321001333312-2201203321320133-0310121133201121-1231011203232312-1222320120330220-3222010213011200-0312213133112233"></a>

<a id="canonical-1010320333122212-1022232001131303-2003322031033303-1332320001313222-0010301103220010-1310031322310102-2120322033202210-2222112213102133"></a>

## namespace property — ref / 003322213002 / 6

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

<a id="canonical-3112002233321301-0110102110002322-1220032202200002-1110301032033322-3231203202221013-0211230003322200-1100312310231302-2223301322220320"></a>

<a id="canonical-0303022111121130-1233132002120100-1323200220121301-2120331111120330-0132322030112331-2311311120211313-3220210200312121-0201030223331322"></a>

## tenant property — ref / 003322213002 / 7

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

<a id="canonical-1312201131210212-2233023303133013-3120003003012303-0033213223003332-0003011311202212-1223202113011311-1330123000101210-3230003103133122"></a>

<a id="canonical-1102333011121202-1232111133301203-1130001321232031-1202100312103110-0210030333012312-1032232032200012-0111133112012022-2333230131223032"></a>

## uid property — ref / 003322213002 / 8

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

<a id="canonical-0223133330120320-2301100330032133-3232210223210230-2232201020322123-1132301101012232-0303320310023011-0122210103200230-3102320032031223"></a>

## Next pages — ref / 003322213002 / 9

- [where.virtual_network](resources--advertise_policy--reference--group-001.md#canonical-3122020011221101-1023022311233310-1101232302222111-0103003032303222-0323233201231323-3133132331222122-0301313303200101-1121023221321331)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-3001022020113000-1323122102210321-0123111102031203-2122012110331032-1230222231110302-2322320110032111-2033032202032030-0000113201220002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133132102321012-2113311112110033-3301033032112330-2113333030232120-2112223300223010-3031310210312233-3213000012200200-1131032221133303"></a>

## where.virtual_site — virtual_site / 113110101232 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101)
- where.virtual_site

<a id="canonical-3111123332111000-0330013222030012-1331331132031302-1011013021230202-3003023000213103-3031123321011033-1210302232311122-2322313310300323"></a>

Type: `"object"`. single nested block, Optional.

Virtual Site. A reference to virtual\_site object.

Upstream description:

A reference to virtual\_site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113022121233032-3011311213102331-0230221121212000-3201030113111131-2310113013022320-1323021020032200-1223322102032121-2223313131223122"></a>

## Direct properties — virtual_site / 113110101232 / 3

- [disable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-1333302102211023-0220022132011031-2211302022331310-0031120020233331-2200020020312122-0001132013211120-0000100120102033-2203231213212213): complete subsection reference.

- [enable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-3331210201121020-3101332102331202-2230102112333212-2102200020003313-0302102211120033-2033321012301110-3221133032031123-1302201222301230): complete subsection reference.

<a id="canonical-1212223010033230-2233030212033332-2132221231323012-1101312003330333-1333101003112033-1320333302021003-1212202010301120-1130313330123223"></a>

<a id="canonical-2102030012101023-3232231112120223-1322213123132303-2002031113100003-1013220102020202-2101113300222312-2101201023101320-3112233132110311"></a>

## network_type property — virtual_site / 113110101232 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
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
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

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

- [ref](resources--advertise_policy--reference--group-001.md#canonical-0330132212201313-2100222231210122-3010202302323202-1232213220231310-3223022012311001-2102011101300123-2331110231122113-3121322330202313): complete subsection reference.

<a id="canonical-1202232113111203-2322002202023220-3303322202333031-0013122112302011-2321310330231132-2333231201300102-1120301131223030-3331120130332300"></a>

## Next pages — virtual_site / 113110101232 / 5

- [where.virtual_site.disable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-1333302102211023-0220022132011031-2211302022331310-0031120020233331-2200020020312122-0001132013211120-0000100120102033-2203231213212213)
- [where.virtual_site.enable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-3331210201121020-3101332102331202-2230102112333212-2102200020003313-0302102211120033-2033321012301110-3221133032031123-1302201222301230)
- [where.virtual_site.ref](resources--advertise_policy--reference--group-001.md#canonical-0330132212201313-2100222231210122-3010202302323202-1232213220231310-3223022012311001-2102011101300123-2331110231122113-3121322330202313)
- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-1333302102211023-0220022132011031-2211302022331310-0031120020233331-2200020020312122-0001132013211120-0000100120102033-2203231213212213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013113133222232-2203103311110101-1010333300102330-1100020031013030-1101021331213233-3220301011310101-0233221123032130-1003103213201012"></a>

## where.virtual_site.disable_internet_vip — disable_internet_vip / 303213211100 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101)
- [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-3001022020113000-1323122102210321-0123111102031203-2122012110331032-1230222231110302-2322320110032111-2033032202032030-0000113201220002)
- where.virtual_site.disable_internet_vip

<a id="canonical-0031221001130031-2303110231300122-0122100003002322-1102230132020300-2203222310111011-3211030030121121-3301131323012200-0301002313311130"></a>

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
disable_internet_vip = {}
```

<a id="canonical-3130122111123010-0212112131003213-3130231221311223-2213313202323022-0130122021001031-3033230203221231-1103331113310011-3023101223303113"></a>

## Direct properties — disable_internet_vip / 303213211100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113000013211320-1111333333333020-0201030333113202-0311013201201001-1100333021233133-3333332320213033-0012033032202311-1131023112100111"></a>

## Next pages — disable_internet_vip / 303213211100 / 4

- [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-3001022020113000-1323122102210321-0123111102031203-2122012110331032-1230222231110302-2322320110032111-2033032202032030-0000113201220002)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-3331210201121020-3101332102331202-2230102112333212-2102200020003313-0302102211120033-2033321012301110-3221133032031123-1302201222301230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301122231233232-3333101212222332-1112012302000230-0110223200311220-0100002102101002-3301223201001310-1223200111332221-3311022131321221"></a>

## where.virtual_site.enable_internet_vip — enable_internet_vip / 211022221333 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101)
- [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-3001022020113000-1323122102210321-0123111102031203-2122012110331032-1230222231110302-2322320110032111-2033032202032030-0000113201220002)
- where.virtual_site.enable_internet_vip

<a id="canonical-0302232010313012-0030212122113323-2111132133120133-1231321010000301-3321301322333323-2033100232300220-3200332120031333-1203033230210312"></a>

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
enable_internet_vip = {}
```

<a id="canonical-2001323021320302-3201123003010303-0223023131023200-1232222330212023-0130022333110120-1010223112103010-1033312331313211-3132010222200332"></a>

## Direct properties — enable_internet_vip / 211022221333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030230033113132-3322333100220321-3113223200221033-2313033000232211-3232010120021222-3223300102201302-0333020002321010-1302011030002102"></a>

## Next pages — enable_internet_vip / 211022221333 / 4

- [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-3001022020113000-1323122102210321-0123111102031203-2122012110331032-1230222231110302-2322320110032111-2033032202032030-0000113201220002)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)

<a id="canonical-0330132212201313-2100222231210122-3010202302323202-1232213220231310-3223022012311001-2102011101300123-2331110231122113-3121322330202313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303131321123211-1021010011330231-1301120013012012-2101131301311310-2111122201323112-2012133210133320-1332110133132000-0301022221030201"></a>

## where.virtual_site.ref — ref / 132101322023 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [where](resources--advertise_policy--reference--group-001.md#canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101)
- [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-3001022020113000-1323122102210321-0123111102031203-2122012110331032-1230222231110302-2322320110032111-2033032202032030-0000113201220002)
- where.virtual_site.ref

<a id="canonical-2230002001211311-1123223202012320-3130101232012110-1312133223233010-0321333301132132-3032113203102201-0030322013232320-0010312001210020"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual\_site direct reference.

Upstream description:

A virtual\_site direct reference.

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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130022332232332-0022222213220003-0002122202212330-0310233111121120-1032010001312020-0233322223310122-3303211131221303-3130133030100002"></a>

## Direct properties — ref / 132101322023 / 3

<a id="canonical-3211131022023000-1033222101131022-2000013312233000-3230102012132031-0110321031001001-2203021222203101-1021133103221030-1010113300010013"></a>

<a id="canonical-1333011322022112-1021330201233020-0113110103111212-1312213000100023-3311230000020211-2133211212202312-1221203320110232-1101110331320333"></a>

## kind property — ref / 132101322023 / 4

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

<a id="canonical-1303130221131111-2131133022111033-3222023311011332-0033023212112001-3031321310113021-3130223321103201-1122122323002211-1213202303323220"></a>

<a id="canonical-2013012032030101-2212211121223123-3233110202020110-0210213113310113-3201300110313210-1300113102100001-1202032202002101-3322232030332023"></a>

## name property — ref / 132101322023 / 5

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

<a id="canonical-2133222033113033-0133103330313332-3202202010301212-3102210012301001-3200332013121200-0100103032310222-2301111103132020-1321000311003311"></a>

<a id="canonical-1220022323300311-1310022332302032-2000233120311313-3310013120321201-1033122103323011-0030023230023213-2020100030110310-0101303121200033"></a>

## namespace property — ref / 132101322023 / 6

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

<a id="canonical-3122022030012320-0110111001311122-2331302333023302-3332120332103023-0311123111310202-3132330223213301-1210122022022333-3201120301230313"></a>

<a id="canonical-2111103312130222-1122003203223203-2202131120321003-2021323232020100-2013232231222211-3302011232312000-2313000202333322-3023323000023222"></a>

## tenant property — ref / 132101322023 / 7

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

<a id="canonical-1313013321313013-3022110021320313-3001223112333003-1132330013021022-1022223230133002-1130111202232312-1300210330201202-2301100213120013"></a>

<a id="canonical-0102301131300320-3002202013001213-0312301332303311-0321023013220230-2221310230223203-3202220323033021-0103210321223102-0203001310000321"></a>

## uid property — ref / 132101322023 / 8

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

<a id="canonical-3223230333010130-1202123211103332-3021121100113220-2100032320202230-0312333221302033-0220210032020212-2133312233213303-0123330113330221"></a>

## Next pages — ref / 132101322023 / 9

- [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-3001022020113000-1323122102210321-0123111102031203-2122012110331032-1230222231110302-2322320110032111-2033032202032030-0000113201220002)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
