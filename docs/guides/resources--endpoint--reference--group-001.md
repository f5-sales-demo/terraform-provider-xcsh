---
page_title: "xcsh_endpoint reference"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_endpoint reference."
---

# xcsh_endpoint reference

<a id="canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200321211222313-3100102001232132-0300300232010013-2302002303233120-2310313101223201-1220222303021233-2211202300032222-0101201122122022"></a>

## Property reference — Property reference / 130211022331 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- Property reference

<a id="canonical-0312122001012120-0031311223300212-0213112010300333-2202233002303220-1022310010200032-1121322131233232-1322232222123002-1002131010312002"></a>

## Direct properties — Property reference / 130211022331 / 3

<a id="canonical-3022213103212211-3211100210113213-1001132313102022-3213002022130203-2031223022013312-3100110002012003-0023321322233203-2102331002300123"></a>

<a id="canonical-0210021323101323-2113102011022032-3100202111311120-2331201311003320-3321131021213210-2113112311130123-3130331200131123-0023130220011222"></a>

## annotations property — Property reference / 130211022331 / 4

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

<a id="canonical-0020110230020100-3212020120032003-1021113021012310-2031211110302210-0201303331000013-2223101320022030-1022031110220012-1322223011332032"></a>

<a id="canonical-0132303200301201-0320110302303231-2330220000223213-0233003011221133-0330330003332203-2113333202203101-2013331123003121-3301203113202330"></a>

## description property — Property reference / 130211022331 / 5

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

<a id="canonical-2301133030310102-0102333231312101-1332133300211313-0223320232023331-2322032233311202-3333122120312310-1110132010013121-3212101310320201"></a>

<a id="canonical-3001233003202220-3331123210211133-3120301321311102-0211223030223122-0333112303201000-3232301103133311-0213031233202012-0220333030212212"></a>

## disable property — Property reference / 130211022331 / 6

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

<a id="canonical-3310001122220332-0202000331110201-0232321203231113-0120300030233312-0113032030210230-2312012330002213-3031331213112311-2333313101003001"></a>

<a id="canonical-3231322020220202-3022010333031301-3303101320000312-2032002333322011-1003333332331011-1211210301103022-1033010100203211-2333021211211130"></a>

## dns_name property — Property reference / 130211022331 / 7

Type: `"string"`. Optional, Computed.

\[OneOf: DNS\_name, DNS\_name\_advanced, ip, service\_info\] Exclusive with \[DNS\_name\_advanced IP
service\_info\] Endpoint's IP address is discovered using DNS name resolution. The name given here
is fully qualified domain name.

Upstream description:

Exclusive with \[DNS\_name\_advanced IP service\_info\] Endpoint's IP address is discovered using
DNS name resolution. The name given here is fully qualified domain name.

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
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

OneOf alternatives in this subsection:

- [dns_name](resources--endpoint--reference--group-001.md#canonical-3310001122220332-0202000331110201-0232321203231113-0120300030233312-0113032030210230-2312012330002213-3031331213112311-2333313101003001)
- [dns_name_advanced](resources--endpoint--reference--group-001.md#canonical-1222333211123011-2110202013032122-1223233320230320-1332023111220012-0311320012210232-3130003011332222-1111330222233012-0223100100103233)
- [ip](resources--endpoint--reference--group-001.md#canonical-1222010013000233-1220213020102112-0301323233211122-2113201230020310-2003022033023133-0311110300200203-3232302220123332-0220120112130333)
- [service_info](resources--endpoint--reference--group-001.md#canonical-1020021200332333-3133001231313101-2300200301010001-1303232011230100-3022313222320021-3330012321133332-1310320212032210-2301232013132221)

Select alternatives according to the provider validators above.

- [dns_name_advanced](resources--endpoint--reference--group-001.md#canonical-3031010133132331-3222110131233333-0211132022132131-2121101130201212-0013311333303003-1112332032030033-1130323022230323-2033313033231121): complete subsection reference.

<a id="canonical-3102203101021222-0123123003303101-0310200102032321-1323002313022230-3022030212021131-1001013022203231-0133222313313313-2313101031201322"></a>

<a id="canonical-2113130102003233-1331222212233112-2310021301123232-3100232033221002-3011230132202112-2330212011000013-2212101222211210-2020320213002221"></a>

## health_check_port property — Property reference / 130211022331 / 8

Type: `"number"`. Optional, Computed.

By default the health check port of an endpoint is the same as the endpoint’s port. This option
provides an alternative health check port. Setting this with a non-zero value allows an endpoint to
have different health check port.

Upstream description:

By default the health check port of an endpoint is the same as the endpoint’s port. This option
provides an alternative health check port. Setting this with a non-zero value allows an endpoint to
have different health check port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(65535),
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2303101220233213-2311330130112323-2300300030233333-2313322021220033-2003000232302312-2001210302032231-3133131233303023-0220211330001300"></a>

<a id="canonical-0213002110120103-0132330213221011-3031221120132013-0001002020222233-3211100031310200-2300021030221133-0302230301131312-3032301210200122"></a>

## ID property — Property reference / 130211022331 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1222010013000233-1220213020102112-0301323233211122-2113201230020310-2003022033023133-0311110300200203-3232302220123332-0220120112130333"></a>

<a id="canonical-2323122312131131-0333111220233212-2123011020231010-2112213230323101-1102302031233200-0032003313121112-1120032200322212-2012133010120001"></a>

## ip property — Property reference / 130211022331 / 10

Type: `"string"`. Optional, Computed.

Exclusive with \[DNS\_name DNS\_name\_advanced service\_info\] Endpoint is reachable at the given
IPv4/IPv6 address.

Upstream description:

Exclusive with \[DNS\_name DNS\_name\_advanced service\_info\] Endpoint is reachable at the given
IPv4/IPv6 address.

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-0200212111212302-0031113123203201-0023301320132001-0320011131000021-1333000222011223-1332211333111200-2300212022131201-1020120002222121"></a>

<a id="canonical-3301213023133120-3312330203012221-2133210220032213-1023302023201112-3130311000202021-1112023202101011-0311121013233203-3302113103013011"></a>

## labels property — Property reference / 130211022331 / 11

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

<a id="canonical-3221330012032211-2133133320330120-3310310200310222-1022212233222130-3331211012211100-1231332330012003-1221223031221113-0102333311023212"></a>

<a id="canonical-2222313300030301-2200122331303201-0202031023211022-2001221130101121-1113311333303221-0102233103100012-2223110221020031-2310210223233231"></a>

## name property — Property reference / 130211022331 / 12

Type: `"string"`. Required.

Name of the Endpoint. Must be unique within the namespace.

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

<a id="canonical-2011321231012020-0213221330111230-1300122203302332-0130221012221210-1112200310032113-0200100210310131-3113312012310302-1110010030211001"></a>

<a id="canonical-2032011313212310-1220300331120313-3102322103321222-3330301320031320-2233203002311021-0010213222023123-3233020310013233-3010221101231101"></a>

## namespace property — Property reference / 130211022331 / 13

Type: `"string"`. Required.

Namespace where the Endpoint is created.

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

<a id="canonical-3031321210213211-2001323231020220-1032011013233012-3010022122313030-2001101120210210-3122210020221110-0013023111000210-1323132222133203"></a>

<a id="canonical-1210333330010032-3330223132030131-0220203230100202-1311003332131312-0210233220122221-1303311320232123-3110231101320112-3233102031021131"></a>

## port property — Property reference / 130211022331 / 14

Type: `"number"`. Optional, Computed.

Endpoint service is available on this port.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1202110322213202-0313001323310021-2233221212221312-3222123231113100-1102113111220221-1011233000121210-3311211202313212-1213120030212223"></a>

<a id="canonical-3000000022130002-0301301330102323-0211310010112012-1112313033113023-1012232231331110-2002131022303311-0321011001002001-3120201020213030"></a>

## protocol property — Property reference / 130211022331 / 15

Type: `"string"`. Optional, Computed.

\[Enum: TCP|UDP\] Protocol. Endpoint protocol. Default is TCP. Both TCP and UDP protocols are
supported. Possible values are \`TCP\`, \`UDP\`.

Upstream description:

Endpoint protocol. Default is TCP. Both TCP and UDP protocols are supported.

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
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  }
}
```

- [service_info](resources--endpoint--reference--group-001.md#canonical-3023332231310113-2111321122221101-1300032112031023-0133010100210121-2002200120323201-3322322203231112-3301330002231202-3022103031323110): complete subsection reference.

- [snat_pool](resources--endpoint--reference--group-001.md#canonical-2320200211122210-3211302200313303-3020303021320320-1031000323132322-3113303203020011-3030011120020011-1330211210211133-1310330111120320): complete subsection reference.

- [timeouts](resources--endpoint--reference--group-001.md#canonical-2020321312111333-1210011130112021-2133220101210000-1123022133023223-3111320002101211-3011311301303201-2123222213011020-2210110223212230): complete subsection reference.

- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012): complete subsection reference.

<a id="canonical-2233130212220300-2132132123313122-0113032121230310-1102200133210012-0123100020012232-3203310112113121-0011212021131211-0111233002302232"></a>

## All schema paths — Property reference / 130211022331 / 16

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--endpoint--reference--group-001.md#canonical-3022213103212211-3211100210113213-1001132313102022-3213002022130203-2031223022013312-3100110002012003-0023321322233203-2102331002300123) |
| `description` | [description](resources--endpoint--reference--group-001.md#canonical-0020110230020100-3212020120032003-1021113021012310-2031211110302210-0201303331000013-2223101320022030-1022031110220012-1322223011332032) |
| `disable` | [disable](resources--endpoint--reference--group-001.md#canonical-2301133030310102-0102333231312101-1332133300211313-0223320232023331-2322032233311202-3333122120312310-1110132010013121-3212101310320201) |
| `dns_name` | [dns_name](resources--endpoint--reference--group-001.md#canonical-3310001122220332-0202000331110201-0232321203231113-0120300030233312-0113032030210230-2312012330002213-3031331213112311-2333313101003001) |
| `dns_name_advanced` | [dns_name_advanced](resources--endpoint--reference--group-001.md#canonical-1222333211123011-2110202013032122-1223233320230320-1332023111220012-0311320012210232-3130003011332222-1111330222233012-0223100100103233) |
| `dns_name_advanced.name` | [dns_name_advanced.name](resources--endpoint--reference--group-001.md#canonical-1200101121121311-0320130200120331-2003313331100102-2112011122323230-0120122312121001-0022322000333012-2333120230011030-0033200313333302) |
| `dns_name_advanced.refresh_interval` | [dns_name_advanced.refresh_interval](resources--endpoint--reference--group-001.md#canonical-2101022100023103-3331201102130222-3212130111333300-0033130311330013-3213103233100001-0233122133200122-1221030132022321-2322121222312120) |
| `health_check_port` | [health_check_port](resources--endpoint--reference--group-001.md#canonical-3102203101021222-0123123003303101-0310200102032321-1323002313022230-3022030212021131-1001013022203231-0133222313313313-2313101031201322) |
| `id` | [ID](resources--endpoint--reference--group-001.md#canonical-2303101220233213-2311330130112323-2300300030233333-2313322021220033-2003000232302312-2001210302032231-3133131233303023-0220211330001300) |
| `ip` | [ip](resources--endpoint--reference--group-001.md#canonical-1222010013000233-1220213020102112-0301323233211122-2113201230020310-2003022033023133-0311110300200203-3232302220123332-0220120112130333) |
| `labels` | [labels](resources--endpoint--reference--group-001.md#canonical-0200212111212302-0031113123203201-0023301320132001-0320011131000021-1333000222011223-1332211333111200-2300212022131201-1020120002222121) |
| `name` | [name](resources--endpoint--reference--group-001.md#canonical-3221330012032211-2133133320330120-3310310200310222-1022212233222130-3331211012211100-1231332330012003-1221223031221113-0102333311023212) |
| `namespace` | [namespace](resources--endpoint--reference--group-001.md#canonical-2011321231012020-0213221330111230-1300122203302332-0130221012221210-1112200310032113-0200100210310131-3113312012310302-1110010030211001) |
| `port` | [port](resources--endpoint--reference--group-001.md#canonical-3031321210213211-2001323231020220-1032011013233012-3010022122313030-2001101120210210-3122210020221110-0013023111000210-1323132222133203) |
| `protocol` | [protocol](resources--endpoint--reference--group-001.md#canonical-1202110322213202-0313001323310021-2233221212221312-3222123231113100-1102113111220221-1011233000121210-3311211202313212-1213120030212223) |
| `service_info` | [service_info](resources--endpoint--reference--group-001.md#canonical-1020021200332333-3133001231313101-2300200301010001-1303232011230100-3022313222320021-3330012321133332-1310320212032210-2301232013132221) |
| `service_info.discovery_type` | [service_info.discovery_type](resources--endpoint--reference--group-001.md#canonical-0013000321032320-0212101300333233-3121221012112122-1310210013133012-2000200201312200-0000331331303110-2312222210220101-2331333021012313) |
| `service_info.service_name` | [service_info.service_name](resources--endpoint--reference--group-001.md#canonical-3231123022120020-2111030211032123-0023011333113230-0111122313220022-2031111133233212-1313300011033300-1030110232212212-2023010201031331) |
| `service_info.service_selector` | [service_info.service_selector](resources--endpoint--reference--group-001.md#canonical-0332000022303100-2203110221300303-1230000001020310-0211011033331102-0213001011332323-1220132301032232-0020200133033302-2332120010010000) |
| `service_info.service_selector.expressions` | [service_info.service_selector.expressions](resources--endpoint--reference--group-001.md#canonical-0110310321321232-0313101222002123-2011302111003312-0031000001202201-0311322003230222-1320021001313313-3103321312332101-2312103030300131) |
| `snat_pool` | [snat_pool](resources--endpoint--reference--group-001.md#canonical-2311032201033313-3002230300310211-2312022012133320-1201203303011122-3333013313110312-1310212130123212-1310310012130201-0203330123310112) |
| `snat_pool.no_snat_pool` | [snat_pool.no_snat_pool](resources--endpoint--reference--group-001.md#canonical-0300113022102312-3030123332131002-3303100123003002-2121013330320231-2030231210113301-2022020121200022-2331310323111000-0212030130210001) |
| `snat_pool.snat_pool` | [snat_pool.snat_pool](resources--endpoint--reference--group-001.md#canonical-0001123111332322-3332130110312023-2332003003130332-3130111030111233-2320003311230031-3231332022011330-1332103123031300-1310032200100220) |
| `snat_pool.snat_pool.prefixes` | [snat_pool.snat_pool.prefixes](resources--endpoint--reference--group-001.md#canonical-0230131223311010-3232212122232302-3312311101232200-3123332221033011-1010103101001112-0300331101321223-2133113203001130-2101221301203313) |
| `timeouts` | [timeouts](resources--endpoint--reference--group-001.md#canonical-0322230301332211-0303021300212220-0033210212002002-3002022012333123-2101303201012102-0210020000111310-1320230312130301-0100033330332212) |
| `timeouts.create` | [timeouts.create](resources--endpoint--reference--group-001.md#canonical-0003321030302221-1201123230321313-1222303012122221-3331222332310302-3013101222003322-0333210013133223-3133131323303021-1223120301113221) |
| `timeouts.delete` | [timeouts.delete](resources--endpoint--reference--group-001.md#canonical-3002123011132210-3331212303322033-2001333223113000-0300033030322320-0312113023200323-0323112223212110-1210100213020323-1022001211102322) |
| `timeouts.read` | [timeouts.read](resources--endpoint--reference--group-001.md#canonical-3303303221111013-3010023233132121-2102232000121020-2111212101013012-0231003013022113-0330303321020103-0203233331110023-3330003203133131) |
| `timeouts.update` | [timeouts.update](resources--endpoint--reference--group-001.md#canonical-3212223331301120-1012012312223323-3122302210302003-2013233033312132-1003120020210102-2313220121130233-3222323310322002-1030312033010130) |
| `where` | [where](resources--endpoint--reference--group-001.md#canonical-0122200201312221-0003212222313100-2220000103332131-3120303320201131-2032011322200200-0002121020123233-1132222002003000-3213111111323100) |
| `where.site` | [where.site](resources--endpoint--reference--group-001.md#canonical-3302103303320011-1200210023013031-2310303321211233-3002130301113222-0330103232120121-1010003113320002-0033023031121203-2211120313133033) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](resources--endpoint--reference--group-001.md#canonical-2233303320000310-1113133130103002-3310102232213101-0122331201032311-3021330313222122-1131120302313133-3131213333133323-0332300310130212) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](resources--endpoint--reference--group-001.md#canonical-0021030201231113-2023010223312112-0110203201212212-3323202331002022-2302130320303023-3320021302123003-1021130133332302-3303002332032021) |
| `where.site.network_type` | [where.site.network_type](resources--endpoint--reference--group-001.md#canonical-3103011300312322-3201103133231312-0123223312023303-1201213221120312-3130111202010031-2101233321130020-0023121310012123-1020000220230200) |
| `where.site.ref` | [where.site.ref](resources--endpoint--reference--group-001.md#canonical-0013020030111213-3333101000002103-0112023033330312-1233122121322213-0300210100123010-3202302121123132-2201310230231203-1110300220121232) |
| `where.site.ref.kind` | [where.site.ref.kind](resources--endpoint--reference--group-001.md#canonical-1211112031022130-1130322211333233-0301020120222303-1201201000222302-2200210332023031-3113121132011220-3122331302131111-0210212001102301) |
| `where.site.ref.name` | [where.site.ref.name](resources--endpoint--reference--group-001.md#canonical-1222311300020333-0323112011223213-2320013212233020-0201210203111031-1221311100132222-0312200131033101-1003003031331302-3123212220320202) |
| `where.site.ref.namespace` | [where.site.ref.namespace](resources--endpoint--reference--group-001.md#canonical-1002123332331330-0113100330111111-1030213301302022-0213023301111133-2010101300101201-1132120202332210-0200212303230322-0222033030010132) |
| `where.site.ref.tenant` | [where.site.ref.tenant](resources--endpoint--reference--group-001.md#canonical-0201031301100100-1200012333333310-2203031320021120-2000210133101101-2000222330312002-0233201130032312-0131100303231201-2320330121022013) |
| `where.site.ref.uid` | [where.site.ref.uid](resources--endpoint--reference--group-001.md#canonical-1011132230231011-2330220233032213-2313022212200120-2213330011231122-0101220003200200-2331123323333201-1111031231023132-3103001222203300) |
| `where.virtual_network` | [where.virtual_network](resources--endpoint--reference--group-001.md#canonical-2101202101001133-1003001112233203-2131221102310310-3102012003122211-0032101202222110-1013122303310320-0223333203101103-2133300123020220) |
| `where.virtual_network.ref` | [where.virtual_network.ref](resources--endpoint--reference--group-001.md#canonical-1230333333110313-1132212310033213-0333013220231331-1130032301003011-2121223220133312-1010020121330333-3011230121311221-3021113211202012) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](resources--endpoint--reference--group-001.md#canonical-3023000123210133-0132101201032020-0000013331011031-2011101101033220-3030032020233001-2010213110031220-1023321210121302-3011333102121000) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](resources--endpoint--reference--group-001.md#canonical-0123230210210312-0310301022013210-1113113003223230-1112322002011202-2113023123302002-3201110012312201-1300232303313232-3300331122222113) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](resources--endpoint--reference--group-001.md#canonical-1130123200311020-2222323233122123-0021233122113020-3021331320103013-2111310132203002-0221333032103213-2321131132013301-0212030110100210) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](resources--endpoint--reference--group-001.md#canonical-2312132103321023-2202330133022032-3011203311131321-0300011233032313-1131101111213121-2113233011030012-3321123111132013-2113210322203100) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](resources--endpoint--reference--group-001.md#canonical-3321103222100033-2300031103122232-1211202002303301-0012003122112013-3123210322103323-0210011102100132-3330302100320332-0022331230220212) |
| `where.virtual_site` | [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-3122120320111113-3232032030132310-0131123131102023-1002012022331232-0003112220113021-0010200121333000-0021222033010212-2033031322020111) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](resources--endpoint--reference--group-001.md#canonical-2232322021000322-2323200033303311-3013132331130012-2113221003021030-1002221202330300-2221002211121102-1320311300011223-1330323113000100) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](resources--endpoint--reference--group-001.md#canonical-3323210101231133-3102220113001112-0230030133302110-2303313310112133-1231123010120223-1020030022121010-2013001230001133-1003303220222200) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](resources--endpoint--reference--group-001.md#canonical-3321021121011020-3222121203000230-3300111022111311-3132121011113001-3122312231111310-3323010232220230-2112103322200223-1133011212321313) |
| `where.virtual_site.ref` | [where.virtual_site.ref](resources--endpoint--reference--group-001.md#canonical-3010232111333021-3302331002321313-0333132230130020-3200302032021021-2010111221012322-3331003021101332-0120220230220113-1300032212331322) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](resources--endpoint--reference--group-001.md#canonical-1013220213110331-3132022112000300-1032002230233100-3203210233121213-3313303012102210-2333113212233310-2103112303102211-2302223031302012) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](resources--endpoint--reference--group-001.md#canonical-2223132011120231-3211131210103112-1012100102231011-3310103322031232-3302321113212103-1202213232232103-0000123320300103-1332300032131132) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](resources--endpoint--reference--group-001.md#canonical-2030033221122030-3220122311333223-0302233222312322-0012201211110121-2011030230231101-1212332332222012-3300113131122123-3001100112233323) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](resources--endpoint--reference--group-001.md#canonical-0103022030223130-3201221203030300-3123311302131221-2202100012011021-1233311021001201-1321210200211233-0000223123333120-1003232303030132) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](resources--endpoint--reference--group-001.md#canonical-3001101111122202-3010120303201001-3111003210020221-2113313100111323-1221213232131312-1301330022133321-2132003311002221-1300000020122221) |

<a id="canonical-1202222122032302-3313101000023221-1212300323131103-3222223112211331-3310311233113231-2223222200200303-0033123121322020-1210212030303001"></a>

## Next pages — Property reference / 130211022331 / 17

- [dns_name_advanced](resources--endpoint--reference--group-001.md#canonical-3031010133132331-3222110131233333-0211132022132131-2121101130201212-0013311333303003-1112332032030033-1130323022230323-2033313033231121)
- [service_info](resources--endpoint--reference--group-001.md#canonical-3023332231310113-2111321122221101-1300032112031023-0133010100210121-2002200120323201-3322322203231112-3301330002231202-3022103031323110)
- [snat_pool](resources--endpoint--reference--group-001.md#canonical-2320200211122210-3211302200313303-3020303021320320-1031000323132322-3113303203020011-3030011120020011-1330211210211133-1310330111120320)
- [timeouts](resources--endpoint--reference--group-001.md#canonical-2020321312111333-1210011130112021-2133220101210000-1123022133023223-3111320002101211-3011311301303201-2123222213011020-2210110223212230)
- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-3031010133132331-3222110131233333-0211132022132131-2121101130201212-0013311333303003-1112332032030033-1130323022230323-2033313033231121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000110223220011-0320323332020321-2330313121030330-3213210110203312-0202113033101320-2003332213201310-2200010203001231-2300301312300221"></a>

## dns_name_advanced — dns_name_advanced / 101233201323 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- dns_name_advanced

<a id="canonical-1222333211123011-2110202013032122-1223233320230320-1332023111220012-0311320012210232-3130003011332222-1111330222233012-0223100100103233"></a>

Type: `"object"`. single nested block, Optional.

Specifies name and TTL used for DNS resolution.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ttl_choice": "[\"refresh_interval\"]"
}
```

Terraform syntax:

```terraform
dns_name_advanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122210333130313-2002330121312211-2022023230321133-1123313301012102-3023322322030013-2211111230311102-3101122011312323-0202200013001012"></a>

## Direct properties — dns_name_advanced / 101233201323 / 3

<a id="canonical-1200101121121311-0320130200120331-2003313331100102-2112011122323230-0120122312121001-0022322000333012-2333120230011030-0033200313333302"></a>

<a id="canonical-3222121023122232-3100030320330231-0322332300230232-3003312123100121-0230103133330230-2022231003002230-3013121132022012-3210000022132202"></a>

## name property — dns_name_advanced / 101233201323 / 4

Type: `"string"`. Optional.

Endpoint's IP address is discovered using DNS name resolution. The name given here is fully
qualified domain name.

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
    "format": "hostname",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2101022100023103-3331201102130222-3212130111333300-0033130311330013-3213103233100001-0233122133200122-1221030132022321-2322121222312120"></a>

<a id="canonical-2302033123120112-3003130313203220-0013101112003123-0120221221230032-0221023113231012-1021112111103011-2012210232032122-2022022202012002"></a>

## refresh_interval property — dns_name_advanced / 101233201323 / 5

Type: `"number"`. Optional.

Exclusive with \[\] Interval for DNS refresh in seconds.

Upstream description:

Exclusive with \[\] Interval for DNS refresh in seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(10, 604800),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "604800"
  }
}
```

<a id="canonical-0122020222103013-2030132101103031-0023202300112303-0113000333100311-2323133020311020-0312010301021002-1112012220211321-3301102330022303"></a>

## Next pages — dns_name_advanced / 101233201323 / 6

- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-3023332231310113-2111321122221101-1300032112031023-0133010100210121-2002200120323201-3322322203231112-3301330002231202-3022103031323110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030230230013121-0203203010021023-3330222010132101-3303311022033122-0011002032031033-3103320201333332-2132213110222201-2312313102100111"></a>

## service_info — service_info / 120033102100 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- service_info

<a id="canonical-1020021200332333-3133001231313101-2300200301010001-1303232011230100-3022313222320021-3330012321133332-1310320212032210-2301232013132221"></a>

Type: `"object"`. single nested block, Optional.

Specifies whether endpoint service is discovered by name or labels.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("service_name",
    "service_selector")}
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
  "x-ves-oneof-field-service_info": "[\"service_name\",\"service_selector\"]"
}
```

Terraform syntax:

```terraform
service_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212311311311300-2032022131103132-3133310023001323-1211112211301001-3120312303202330-0111003302010023-0330132322100003-3331131102303120"></a>

## Direct properties — service_info / 120033102100 / 3

<a id="canonical-0013000321032320-0212101300333233-3121221012112122-1310210013133012-2000200201312200-0000331331303110-2312222210220101-2331333021012313"></a>

<a id="canonical-1131010133300212-1302230123000303-3202010010013313-3132023213123011-0121131303130130-0221201000311202-1022231313111013-1123201332131032"></a>

## discovery_type property — service_info / 120033102100 / 4

Type: `"string"`. Optional.

\[Enum: INVALID\_DISCOVERY|K8S|CONSUL|CLASSIC\_BIGIP|THIRD\_PARTY|NGINX\_ONE\] Specifies the type of
discovery Invalid Discovery mechanism Discover from Kubernetes cluster Discover from Consul service
Discover from Classic BIG-IP Clusters Discover for Third Party Application Discover from NGINX One.
Possible values are \`INVALID\_DISCOVERY\`, \`K8S\`, \`CONSUL\`, \`CLASSIC\_BIGIP\`,
\`THIRD\_PARTY\`, \`NGINX\_ONE\`. Defaults to \`INVALID\_DISCOVERY\`.

Upstream description:

Specifies the type of discovery

Invalid Discovery mechanism Discover from Kubernetes cluster Discover from Consul service Discover
from Classic BIG-IP Clusters Discover for Third Party Application Discover from NGINX One.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INVALID_DISCOVERY",
    "K8S",
    "CONSUL",
    "CLASSIC_BIGIP",
    "THIRD_PARTY",
    "NGINX_ONE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALID_DISCOVERY",
  "enum": [
    "INVALID_DISCOVERY",
    "K8S",
    "CONSUL",
    "CLASSIC_BIGIP",
    "THIRD_PARTY",
    "NGINX_ONE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3231123022120020-2111030211032123-0023011333113230-0111122313220022-2031111133233212-1313300011033300-1030110232212212-2023010201031331"></a>

<a id="canonical-3120312131333210-1220222201011131-3023002112020333-3030223022023010-0003223012120102-2312200133213111-1111301212330003-2100313233133230"></a>

## service_name property — service_info / 120033102100 / 5

Type: `"string"`. Optional.

Exclusive with \[service\_selector\] Name of the service to discover with an optional namespace and
cluster identifier. The format is service\_name.namespace\_name:cluster\_identifier for K8s and
service\_name:cluster\_identifier for Consul Endpoint will be discovered in all discovery objects
where the..

Upstream description:

Exclusive with \[service\_selector\] Name of the service to discover with an optional namespace and
cluster identifier. The format is service\_name.namespace\_name:cluster\_identifier for K8s and
service\_name:cluster\_identifier for Consul Endpoint will be discovered in all discovery objects
where the cluster identifier matches. If cluster identifier is not specified then discovery will be
done in all discovery objects of the site.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [service_selector](resources--endpoint--reference--group-001.md#canonical-2021012311300000-3320101133322230-1100212322210332-3313011103020113-0030020001022231-0023200000210133-0220021221312010-1230001003310230): complete subsection reference.

<a id="canonical-3020232131220330-0020022200333013-1000220112010110-1131321210011102-3030332132210302-0032001032023313-0311022012001313-0101001020031023"></a>

## Next pages — service_info / 120033102100 / 6

- [service_info.service_selector](resources--endpoint--reference--group-001.md#canonical-2021012311300000-3320101133322230-1100212322210332-3313011103020113-0030020001022231-0023200000210133-0220021221312010-1230001003310230)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-2021012311300000-3320101133322230-1100212322210332-3313011103020113-0030020001022231-0023200000210133-0220021221312010-1230001003310230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331032233222011-3301133302121001-2000321101120303-3010311322021201-3213310122301011-1100001323120111-1111030223211011-2222002111012331"></a>

## service_info.service_selector — service_selector / 112212133332 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [service_info](resources--endpoint--reference--group-001.md#canonical-3023332231310113-2111321122221101-1300032112031023-0133010100210121-2002200120323201-3322322203231112-3301330002231202-3022103031323110)
- service_info.service_selector

<a id="canonical-0332000022303100-2203110221300303-1230000001020310-0211011033331102-0213001011332323-1220132301032232-0020200133033302-2332120010010000"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
service_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-1012011300112302-2301210103102130-3023332322312131-1101111222031132-3033101003230311-3011310003022012-2231133312222202-3102332100103332"></a>

## Direct properties — service_selector / 112212133332 / 3

<a id="canonical-0110310321321232-0313101222002123-2011302111003312-0031000001202201-0311322003230222-1320021001313313-3103321312332101-2312103030300131"></a>

<a id="canonical-3230220022030022-3213311110312333-3100022221232200-1230020230131321-0300213231022133-3200122213031322-2211313300200131-1212001210001230"></a>

## expressions property — service_selector / 112212133332 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3212330111003203-0310200100030201-1122322313302220-3302202023120212-1030333303103021-0033110322211010-0212111132113200-1322323113231213"></a>

## Next pages — service_selector / 112212133332 / 5

- [service_info](resources--endpoint--reference--group-001.md#canonical-3023332231310113-2111321122221101-1300032112031023-0133010100210121-2002200120323201-3322322203231112-3301330002231202-3022103031323110)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-2320200211122210-3211302200313303-3020303021320320-1031000323132322-3113303203020011-3030011120020011-1330211210211133-1310330111120320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312113302210030-2200321000101221-2030330120300201-3023311132301330-3212032101331030-1013201001222212-0122103212001231-2012233010101301"></a>

## snat_pool — snat_pool / 212221002332 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- snat_pool

<a id="canonical-2311032201033313-3002230300310211-2312022012133320-1201203303011122-3333013313110312-1310212130123212-1310310012130201-0203330123310112"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121122331332231-0322212330002203-2101230102021021-1202001333023230-2333230232331200-3322211033331013-0131120133100033-2011221012300220"></a>

## Direct properties — snat_pool / 212221002332 / 3

- [no_snat_pool](resources--endpoint--reference--group-001.md#canonical-2220221221301300-0010220322312023-2102300102101000-3101021221201033-1120021010113302-1101312221123121-2022101121133301-0232103001330100): complete subsection reference.

- [snat_pool](resources--endpoint--reference--group-001.md#canonical-3111213201100002-0211301331321122-0020202102011002-3313121001230300-1210221230210033-0002002022223133-2012000010023123-2222031331101001): complete subsection reference.

<a id="canonical-1230012221033323-1002330301102223-3033321303032223-1022222222101031-2220201023020312-2230303233330201-0201320310021101-1003131202231011"></a>

## Next pages — snat_pool / 212221002332 / 4

- [snat_pool.no_snat_pool](resources--endpoint--reference--group-001.md#canonical-2220221221301300-0010220322312023-2102300102101000-3101021221201033-1120021010113302-1101312221123121-2022101121133301-0232103001330100)
- [snat_pool.snat_pool](resources--endpoint--reference--group-001.md#canonical-3111213201100002-0211301331321122-0020202102011002-3313121001230300-1210221230210033-0002002022223133-2012000010023123-2222031331101001)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-2220221221301300-0010220322312023-2102300102101000-3101021221201033-1120021010113302-1101312221123121-2022101121133301-0232103001330100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311210003202010-0133110221033222-2023323333013113-1030112200330231-3223120202031123-2303303131020102-0203030333021303-1030111011230123"></a>

## snat_pool.no_snat_pool — no_snat_pool / 232300230021 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [snat_pool](resources--endpoint--reference--group-001.md#canonical-2320200211122210-3211302200313303-3020303021320320-1031000323132322-3113303203020011-3030011120020011-1330211210211133-1310330111120320)
- snat_pool.no_snat_pool

<a id="canonical-0300113022102312-3030123332131002-3303100123003002-2121013330320231-2030231210113301-2022020121200022-2331310323111000-0212030130210001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

<a id="canonical-2330120321233001-1123213003330221-0321031223202103-0313132321122132-1032131013200220-3320202211310013-0300013312230312-3013131030200011"></a>

## Direct properties — no_snat_pool / 232300230021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301321000101020-1320000130222303-3122032222212102-2312321102222031-1230000002221333-0220311121010130-3212012133122311-2011013312032023"></a>

## Next pages — no_snat_pool / 232300230021 / 4

- [snat_pool](resources--endpoint--reference--group-001.md#canonical-2320200211122210-3211302200313303-3020303021320320-1031000323132322-3113303203020011-3030011120020011-1330211210211133-1310330111120320)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-3111213201100002-0211301331321122-0020202102011002-3313121001230300-1210221230210033-0002002022223133-2012000010023123-2222031331101001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111013110103103-2231331313002323-2133003123213112-0031312231212323-0000031032200120-2322220112233313-3021332301323312-3322000123232213"></a>

## snat_pool.snat_pool — snat_pool / 012321222213 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [snat_pool](resources--endpoint--reference--group-001.md#canonical-2320200211122210-3211302200313303-3020303021320320-1031000323132322-3113303203020011-3030011120020011-1330211210211133-1310330111120320)
- snat_pool.snat_pool

<a id="canonical-0001123111332322-3332130110312023-2332003003130332-3130111030111233-2320003311230031-3231332022011330-1332103123031300-1310032200100220"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213101003003322-0111333321231212-2322132221133333-1213122313333022-1232131213232202-0020112112301020-0001103333300322-3131022333313000"></a>

## Direct properties — snat_pool / 012321222213 / 3

<a id="canonical-0230131223311010-3232212122232302-3312311101232200-3123332221033011-1010103101001112-0300331101321223-2133113203001130-2101221301203313"></a>

<a id="canonical-1002021213321323-2213120321311220-1300030303113012-0233312031232200-1311222333333130-1223112222012303-3120203302332302-1300102133310001"></a>

## prefixes property — snat_pool / 012321222213 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1302113011023210-2122000002333033-1110211210311112-0133012230102320-2333222330030001-3332013033032112-2002010013131320-2100330211103210"></a>

## Next pages — snat_pool / 012321222213 / 5

- [snat_pool](resources--endpoint--reference--group-001.md#canonical-2320200211122210-3211302200313303-3020303021320320-1031000323132322-3113303203020011-3030011120020011-1330211210211133-1310330111120320)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-2020321312111333-1210011130112021-2133220101210000-1123022133023223-3111320002101211-3011311301303201-2123222213011020-2210110223212230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102221002331021-2320032103323223-2001002030321310-3101013332013321-1203123013033313-0000031333311313-2211132321212110-0231102123221231"></a>

## timeouts — timeouts / 112022232321 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- timeouts

<a id="canonical-0322230301332211-0303021300212220-0033210212002002-3002022012333123-2101303201012102-0210020000111310-1320230312130301-0100033330332212"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200322203130101-3331003102010131-2020221120230002-3233001033321113-2223322010323113-3233332303023200-2231312231013033-3200332232003233"></a>

## Direct properties — timeouts / 112022232321 / 3

<a id="canonical-0003321030302221-1201123230321313-1222303012122221-3331222332310302-3013101222003322-0333210013133223-3133131323303021-1223120301113221"></a>

<a id="canonical-2000310001123321-2032231022202231-2113211313111020-1312023100011133-1121211030333211-1102303233222123-3322320223002112-3102331312323222"></a>

## create property — timeouts / 112022232321 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3002123011132210-3331212303322033-2001333223113000-0300033030322320-0312113023200323-0323112223212110-1210100213020323-1022001211102322"></a>

<a id="canonical-1200220102030223-1212131003300333-1010211030331312-3233300122023100-2320130003203313-1133132023111113-2130103122003321-1121030311001020"></a>

## delete property — timeouts / 112022232321 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3303303221111013-3010023233132121-2102232000121020-2111212101013012-0231003013022113-0330303321020103-0203233331110023-3330003203133131"></a>

<a id="canonical-2123212110013103-3110101021002011-3333212002033202-0200020030330313-1230332230113032-2032301331123032-0213320202232032-2220222233301320"></a>

## read property — timeouts / 112022232321 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3212223331301120-1012012312223323-3122302210302003-2013233033312132-1003120020210102-2313220121130233-3222323310322002-1030312033010130"></a>

<a id="canonical-0100032231230023-0312121111133020-2030300301132312-0300322310301101-3311100030110313-0220032023010222-3230332023322103-3122002222312302"></a>

## update property — timeouts / 112022232321 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2130000331322302-3031333311111103-1200322112333301-3330233101310002-2112030302302321-3212323031200213-1021302122301201-3022020210203330"></a>

## Next pages — timeouts / 112022232321 / 8

- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032203311302223-0221032201110221-2000233200231333-1221112121030330-3301221113300310-1022002003200133-1332302202231220-0002213211312310"></a>

## where — where / 133231103220 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- where

<a id="canonical-0122200201312221-0003212222313100-2220000103332131-3120303320201131-2032011322200200-0002121020123233-1132222002003000-3213111111323100"></a>

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

<a id="canonical-3232312111110320-3010200032033210-2031231121200001-2213001131110010-2323321013220123-3310232111023013-0310012110302203-3101232332123231"></a>

## Direct properties — where / 133231103220 / 3

- [site](resources--endpoint--reference--group-001.md#canonical-1132210311020321-0331233220230203-3322122311103220-3322211130002220-3032233123333310-1310130032000101-3203021211020111-0102203012101311): complete subsection reference.

- [virtual_network](resources--endpoint--reference--group-001.md#canonical-3311203232122023-0113222203021311-0202021021330303-3213222312112201-0130301123302213-1013122012131220-0102110131200300-1210321123212101): complete subsection reference.

- [virtual_site](resources--endpoint--reference--group-001.md#canonical-0233032123312002-0020322311112123-1313223231330103-0310321302231323-1220111232221310-2001231130313210-3110123101002233-2111000100202022): complete subsection reference.

<a id="canonical-2202330302103231-0233331311232303-2321230333012232-3222321133203312-0101220032323120-2312030200232233-3131212020012112-0023320011122123"></a>

## Next pages — where / 133231103220 / 4

- [where.site](resources--endpoint--reference--group-001.md#canonical-1132210311020321-0331233220230203-3322122311103220-3322211130002220-3032233123333310-1310130032000101-3203021211020111-0102203012101311)
- [where.virtual_network](resources--endpoint--reference--group-001.md#canonical-3311203232122023-0113222203021311-0202021021330303-3213222312112201-0130301123302213-1013122012131220-0102110131200300-1210321123212101)
- [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-0233032123312002-0020322311112123-1313223231330103-0310321302231323-1220111232221310-2001231130313210-3110123101002233-2111000100202022)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-1132210311020321-0331233220230203-3322122311103220-3322211130002220-3032233123333310-1310130032000101-3203021211020111-0102203012101311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301312112310101-1002232331000021-0031011311312202-3033222201030313-3101302203222220-3323232300213223-1300203211121212-0132233303213332"></a>

## where.site — site / 303231213111 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012)
- where.site

<a id="canonical-3302103303320011-1200210023013031-2310303321211233-3002130301113222-0330103232120121-1010003113320002-0033023031121203-2211120313133033"></a>

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

<a id="canonical-3013032003010121-3303302322113323-1220111102210113-1013210023011323-3301221320211320-3301212303321301-2103212230200012-2122011223022120"></a>

## Direct properties — site / 303231213111 / 3

- [disable_internet_vip](resources--endpoint--reference--group-001.md#canonical-1212323120222031-2200021131301203-0231111101331132-1212013113020311-2100103201123020-1230223302110332-3130100333232010-2112120221103010): complete subsection reference.

- [enable_internet_vip](resources--endpoint--reference--group-001.md#canonical-3012201123010002-2212330330333300-2333003223022130-3111000330103130-1301332023000233-3303220302210220-1332223210131012-1220211231310132): complete subsection reference.

<a id="canonical-3103011300312322-3201103133231312-0123223312023303-1201213221120312-3130111202010031-2101233321130020-0023121310012123-1020000220230200"></a>

<a id="canonical-0233000333101333-2113200020013200-2310210200322133-3022002332012101-0113330230123010-0113332112123303-3031310201023000-1312110213203233"></a>

## network_type property — site / 303231213111 / 4

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

- [ref](resources--endpoint--reference--group-001.md#canonical-1110122033032020-0300203112133100-1010133211331233-0322323113010213-2332123111032311-3001200102302010-1331200302032320-2103133032003211): complete subsection reference.

<a id="canonical-0003020020201213-0333202330313333-3201312230111101-3021013310331203-2133323031321313-1333333022011313-3032230002031313-2102102322133112"></a>

## Next pages — site / 303231213111 / 5

- [where.site.disable_internet_vip](resources--endpoint--reference--group-001.md#canonical-1212323120222031-2200021131301203-0231111101331132-1212013113020311-2100103201123020-1230223302110332-3130100333232010-2112120221103010)
- [where.site.enable_internet_vip](resources--endpoint--reference--group-001.md#canonical-3012201123010002-2212330330333300-2333003223022130-3111000330103130-1301332023000233-3303220302210220-1332223210131012-1220211231310132)
- [where.site.ref](resources--endpoint--reference--group-001.md#canonical-1110122033032020-0300203112133100-1010133211331233-0322323113010213-2332123111032311-3001200102302010-1331200302032320-2103133032003211)
- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-1212323120222031-2200021131301203-0231111101331132-1212013113020311-2100103201123020-1230223302110332-3130100333232010-2112120221103010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200031013110001-2200201330001200-3122302013231232-0221020333211232-1301231131203130-2001123000110301-3222203323032213-2003033301301103"></a>

## where.site.disable_internet_vip — disable_internet_vip / 311300000223 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012)
- [where.site](resources--endpoint--reference--group-001.md#canonical-1132210311020321-0331233220230203-3322122311103220-3322211130002220-3032233123333310-1310130032000101-3203021211020111-0102203012101311)
- where.site.disable_internet_vip

<a id="canonical-2233303320000310-1113133130103002-3310102232213101-0122331201032311-3021330313222122-1131120302313133-3131213333133323-0332300310130212"></a>

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

<a id="canonical-3323021301303113-1020120010003011-0330012303302230-2101113110210111-3131030322032020-3012202132230223-3302211212331022-0220333313300213"></a>

## Direct properties — disable_internet_vip / 311300000223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101020012300130-1123303031022030-1223103322332333-3313131133220300-3031232011302113-1001312223210012-2231220301321131-2133222321131023"></a>

## Next pages — disable_internet_vip / 311300000223 / 4

- [where.site](resources--endpoint--reference--group-001.md#canonical-1132210311020321-0331233220230203-3322122311103220-3322211130002220-3032233123333310-1310130032000101-3203021211020111-0102203012101311)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-3012201123010002-2212330330333300-2333003223022130-3111000330103130-1301332023000233-3303220302210220-1332223210131012-1220211231310132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231002010323122-3310120122132120-3223232300212132-3202010320101001-0131120011013230-0230002203230212-2321133003333230-2002122002132033"></a>

## where.site.enable_internet_vip — enable_internet_vip / 000223121100 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012)
- [where.site](resources--endpoint--reference--group-001.md#canonical-1132210311020321-0331233220230203-3322122311103220-3322211130002220-3032233123333310-1310130032000101-3203021211020111-0102203012101311)
- where.site.enable_internet_vip

<a id="canonical-0021030201231113-2023010223312112-0110203201212212-3323202331002022-2302130320303023-3320021302123003-1021130133332302-3303002332032021"></a>

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

<a id="canonical-1303001010020103-3230321032201122-2130201231132000-3111322232000313-0320102000120223-1120322001113022-0102331303221133-0001132200301123"></a>

## Direct properties — enable_internet_vip / 000223121100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021310000203311-0200123032023011-0032233123213013-1303020022201210-2002103220231320-1200112023113313-2021330111011012-3123031212102310"></a>

## Next pages — enable_internet_vip / 000223121100 / 4

- [where.site](resources--endpoint--reference--group-001.md#canonical-1132210311020321-0331233220230203-3322122311103220-3322211130002220-3032233123333310-1310130032000101-3203021211020111-0102203012101311)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-1110122033032020-0300203112133100-1010133211331233-0322323113010213-2332123111032311-3001200102302010-1331200302032320-2103133032003211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211222310103021-1320003301001020-2321202232320311-2233223020011100-2212103032232111-1133000210021100-2111101333311200-1110110322233220"></a>

## where.site.ref — ref / 201031321332 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012)
- [where.site](resources--endpoint--reference--group-001.md#canonical-1132210311020321-0331233220230203-3322122311103220-3322211130002220-3032233123333310-1310130032000101-3203021211020111-0102203012101311)
- where.site.ref

<a id="canonical-0013020030111213-3333101000002103-0112023033330312-1233122121322213-0300210100123010-3202302121123132-2201310230231203-1110300220121232"></a>

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031112121233030-1100033310133330-0201323031122133-0020133201030000-3101311203032033-3030123130010212-0233211033332231-2121312010002101"></a>

## Direct properties — ref / 201031321332 / 3

<a id="canonical-1211112031022130-1130322211333233-0301020120222303-1201201000222302-2200210332023031-3113121132011220-3122331302131111-0210212001102301"></a>

<a id="canonical-3201112211333221-1021013023213302-3220221220003131-3022210333210102-2131033032120011-0310033013210232-2023200330132102-3123103213323201"></a>

## kind property — ref / 201031321332 / 4

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

<a id="canonical-1222311300020333-0323112011223213-2320013212233020-0201210203111031-1221311100132222-0312200131033101-1003003031331302-3123212220320202"></a>

<a id="canonical-0330231030132001-2033331212101213-0102101032013103-3103203100301312-0002020312200031-3311321200200302-3323012002022131-0012300211123002"></a>

## name property — ref / 201031321332 / 5

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

<a id="canonical-1002123332331330-0113100330111111-1030213301302022-0213023301111133-2010101300101201-1132120202332210-0200212303230322-0222033030010132"></a>

<a id="canonical-2010230313210303-2130220120213232-0111202001331131-0010013130302021-2332331030113101-1121232002303312-0012222113333313-1032230301120120"></a>

## namespace property — ref / 201031321332 / 6

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

<a id="canonical-0201031301100100-1200012333333310-2203031320021120-2000210133101101-2000222330312002-0233201130032312-0131100303231201-2320330121022013"></a>

<a id="canonical-3230223021000303-1321303303130002-1121001201033133-2003000320233233-2133210023212321-1131223101033200-2310011311020112-3111221300010232"></a>

## tenant property — ref / 201031321332 / 7

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

<a id="canonical-1011132230231011-2330220233032213-2313022212200120-2213330011231122-0101220003200200-2331123323333201-1111031231023132-3103001222203300"></a>

<a id="canonical-3000000311322202-3211311010310033-0033123333131023-1231310000210112-1213201312333220-3120321131200121-2302033321020111-1203321032300303"></a>

## uid property — ref / 201031321332 / 8

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

<a id="canonical-3333031200202231-3103213211213012-0113330010323002-3002331132230303-0210031033323213-2001323233322010-0103321232313021-0232201112303221"></a>

## Next pages — ref / 201031321332 / 9

- [where.site](resources--endpoint--reference--group-001.md#canonical-1132210311020321-0331233220230203-3322122311103220-3322211130002220-3032233123333310-1310130032000101-3203021211020111-0102203012101311)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-3311203232122023-0113222203021311-0202021021330303-3213222312112201-0130301123302213-1013122012131220-0102110131200300-1210321123212101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133223200000322-1021020131312022-0323330103123023-3021230020220320-2013023211011110-0113231031333121-2022303212322322-3021212012311323"></a>

## where.virtual_network — virtual_network / 020300101013 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012)
- where.virtual_network

<a id="canonical-2101202101001133-1003001112233203-2131221102310310-3102012003122211-0032101202222110-1013122303310320-0223333203101103-2133300123020220"></a>

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

<a id="canonical-3230030010110101-2203232121131123-0002133113223321-1003011330103232-2130002003302221-0112301311033103-0012231300223223-3311221123022032"></a>

## Direct properties — virtual_network / 020300101013 / 3

- [ref](resources--endpoint--reference--group-001.md#canonical-2130021033102320-1002112123230023-2111213002011210-2202020231123102-0123033323011232-1232232231302120-2101311212231023-1001032320300301): complete subsection reference.

<a id="canonical-0023300230132331-1020120203023111-3210130301303022-3102222202221310-2313303211013023-1101333323211122-0210001131313032-1001013210112232"></a>

## Next pages — virtual_network / 020300101013 / 4

- [where.virtual_network.ref](resources--endpoint--reference--group-001.md#canonical-2130021033102320-1002112123230023-2111213002011210-2202020231123102-0123033323011232-1232232231302120-2101311212231023-1001032320300301)
- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-2130021033102320-1002112123230023-2111213002011210-2202020231123102-0123033323011232-1232232231302120-2101311212231023-1001032320300301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101012201203001-1230121320321233-1122212023333321-2001201212200312-1000221130303100-1222122013022211-0321113123121232-3311101133022210"></a>

## where.virtual_network.ref — ref / 201112003101 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012)
- [where.virtual_network](resources--endpoint--reference--group-001.md#canonical-3311203232122023-0113222203021311-0202021021330303-3213222312112201-0130301123302213-1013122012131220-0102110131200300-1210321123212101)
- where.virtual_network.ref

<a id="canonical-1230333333110313-1132212310033213-0333013220231331-1130032301003011-2121223220133312-1010020121330333-3011230121311221-3021113211202012"></a>

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221311010302103-3012032332000313-1010231201311033-0212110110210333-3220233001000122-2320101330103133-2030313231320203-2311020030010122"></a>

## Direct properties — ref / 201112003101 / 3

<a id="canonical-3023000123210133-0132101201032020-0000013331011031-2011101101033220-3030032020233001-2010213110031220-1023321210121302-3011333102121000"></a>

<a id="canonical-0030303011132332-2011333111203110-2201131101221303-2221103100021123-0231330232311133-2120121022322202-1200301313203120-3330231133232023"></a>

## kind property — ref / 201112003101 / 4

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

<a id="canonical-0123230210210312-0310301022013210-1113113003223230-1112322002011202-2113023123302002-3201110012312201-1300232303313232-3300331122222113"></a>

<a id="canonical-0112120201231020-0320220001210213-3321013300123113-1033313021302102-3303201322220311-0332031030110121-2122302322330013-2223333230332020"></a>

## name property — ref / 201112003101 / 5

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

<a id="canonical-1130123200311020-2222323233122123-0021233122113020-3021331320103013-2111310132203002-0221333032103213-2321131132013301-0212030110100210"></a>

<a id="canonical-1330310002210332-0110113032202031-1133123231233033-0010311211303320-3031000031022310-2330100221030230-3003122031110001-3130213232133020"></a>

## namespace property — ref / 201112003101 / 6

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

<a id="canonical-2312132103321023-2202330133022032-3011203311131321-0300011233032313-1131101111213121-2113233011030012-3321123111132013-2113210322203100"></a>

<a id="canonical-2231332020012232-3320122022203201-0120013211323332-0231231231010021-0013003232112133-3221211012002321-2300312011100323-3023231013203233"></a>

## tenant property — ref / 201112003101 / 7

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

<a id="canonical-3321103222100033-2300031103122232-1211202002303301-0012003122112013-3123210322103323-0210011102100132-3330302100320332-0022331230220212"></a>

<a id="canonical-1201101202222201-3121103011013303-1111010311023002-1313203022030011-1203303211332212-1201121302201221-3002121013130201-1022101333300100"></a>

## uid property — ref / 201112003101 / 8

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

<a id="canonical-0200031220330021-2011112233133303-0221101001330032-1300222010013210-3102010312020012-3222001200231103-2101133303201101-3221233033012100"></a>

## Next pages — ref / 201112003101 / 9

- [where.virtual_network](resources--endpoint--reference--group-001.md#canonical-3311203232122023-0113222203021311-0202021021330303-3213222312112201-0130301123302213-1013122012131220-0102110131200300-1210321123212101)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-0233032123312002-0020322311112123-1313223231330103-0310321302231323-1220111232221310-2001231130313210-3110123101002233-2111000100202022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222131313302102-0200322221033300-0012020113302330-2212110122001232-2233302013101323-3032132123201032-2310102003301333-1102333220021201"></a>

## where.virtual_site — virtual_site / 022031010221 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012)
- where.virtual_site

<a id="canonical-3122120320111113-3232032030132310-0131123131102023-1002012022331232-0003112220113021-0010200121333000-0021222033010212-2033031322020111"></a>

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

<a id="canonical-2132101321120320-2030031223000023-1310033231211103-0000131101311123-2110330120133301-3002100132302100-0232230202222300-0300033120210222"></a>

## Direct properties — virtual_site / 022031010221 / 3

- [disable_internet_vip](resources--endpoint--reference--group-001.md#canonical-0321313020021211-2232003203110323-2200133123320100-3103231201001313-1021103202102332-0223223222112102-3221333020303003-2111211323022001): complete subsection reference.

- [enable_internet_vip](resources--endpoint--reference--group-001.md#canonical-3023111311332101-2023203220133100-0100210332330111-2333231031011001-2202132302113122-2131203000203210-0102212011023200-2001023202221332): complete subsection reference.

<a id="canonical-3321021121011020-3222121203000230-3300111022111311-3132121011113001-3122312231111310-3323010232220230-2112103322200223-1133011212321313"></a>

<a id="canonical-2102321303110332-0223121013112113-1013033122300000-0202321320111120-3012012101130023-2021331112130333-0010132102202120-0122132003011102"></a>

## network_type property — virtual_site / 022031010221 / 4

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

- [ref](resources--endpoint--reference--group-001.md#canonical-0332232303331132-3032020031033111-0303021212313131-3211210102121310-3020102312311123-0101121300030200-0022101222012332-1120102221331212): complete subsection reference.

<a id="canonical-0201032300201121-2113032221023210-3220112221120330-0233003221121201-1111113331103321-1132013102230321-1130201002200030-3011203200102211"></a>

## Next pages — virtual_site / 022031010221 / 5

- [where.virtual_site.disable_internet_vip](resources--endpoint--reference--group-001.md#canonical-0321313020021211-2232003203110323-2200133123320100-3103231201001313-1021103202102332-0223223222112102-3221333020303003-2111211323022001)
- [where.virtual_site.enable_internet_vip](resources--endpoint--reference--group-001.md#canonical-3023111311332101-2023203220133100-0100210332330111-2333231031011001-2202132302113122-2131203000203210-0102212011023200-2001023202221332)
- [where.virtual_site.ref](resources--endpoint--reference--group-001.md#canonical-0332232303331132-3032020031033111-0303021212313131-3211210102121310-3020102312311123-0101121300030200-0022101222012332-1120102221331212)
- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-0321313020021211-2232003203110323-2200133123320100-3103231201001313-1021103202102332-0223223222112102-3221333020303003-2111211323022001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222233220222122-3021113333330021-3013322030322110-3121113201132331-2220123110103331-0012111332032133-3011310033123300-3022232310023012"></a>

## where.virtual_site.disable_internet_vip — disable_internet_vip / 202212331003 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012)
- [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-0233032123312002-0020322311112123-1313223231330103-0310321302231323-1220111232221310-2001231130313210-3110123101002233-2111000100202022)
- where.virtual_site.disable_internet_vip

<a id="canonical-2232322021000322-2323200033303311-3013132331130012-2113221003021030-1002221202330300-2221002211121102-1320311300011223-1330323113000100"></a>

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

<a id="canonical-0230223000203030-1200230220203003-3012312320023213-0022112021012110-2312210021121312-2232100321133113-0210012222012113-2303000013232110"></a>

## Direct properties — disable_internet_vip / 202212331003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333100003011013-1333332112200120-2222302330102013-0201013101102100-1120332320132030-0001122001101102-0022210101131020-3030233320202102"></a>

## Next pages — disable_internet_vip / 202212331003 / 4

- [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-0233032123312002-0020322311112123-1313223231330103-0310321302231323-1220111232221310-2001231130313210-3110123101002233-2111000100202022)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-3023111311332101-2023203220133100-0100210332330111-2333231031011001-2202132302113122-2131203000203210-0102212011023200-2001023202221332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212331110003312-2130201113213300-1003111233010322-2211030310301121-3212331111102000-1330020013201001-1332230301203123-1120303332130120"></a>

## where.virtual_site.enable_internet_vip — enable_internet_vip / 002302101333 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012)
- [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-0233032123312002-0020322311112123-1313223231330103-0310321302231323-1220111232221310-2001231130313210-3110123101002233-2111000100202022)
- where.virtual_site.enable_internet_vip

<a id="canonical-3323210101231133-3102220113001112-0230030133302110-2303313310112133-1231123010120223-1020030022121010-2013001230001133-1003303220222200"></a>

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

<a id="canonical-1030011323120021-3222131232113113-3322103313032203-2022223331333232-0121120022032123-1222120303231322-0112302222002223-0020222310201110"></a>

## Direct properties — enable_internet_vip / 002302101333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033031111322200-2211021003212131-2210220130323332-1213030331100210-2132130001121230-1233120021003113-1300001120232030-1001002233110102"></a>

## Next pages — enable_internet_vip / 002302101333 / 4

- [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-0233032123312002-0020322311112123-1313223231330103-0310321302231323-1220111232221310-2001231130313210-3110123101002233-2111000100202022)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)

<a id="canonical-0332232303331132-3032020031033111-0303021212313131-3211210102121310-3020102312311123-0101121300030200-0022101222012332-1120102221331212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022323313111222-3001303312300103-3123023300011132-2210012012101301-0303231320332332-0103200120002121-2302332101313012-0003331220221112"></a>

## where.virtual_site.ref — ref / 111013111020 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [where](resources--endpoint--reference--group-001.md#canonical-0220301202000013-0110122032311022-3013221012200233-1311123310310110-2121130123330031-2023311031113102-3212231210212101-0322131210230012)
- [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-0233032123312002-0020322311112123-1313223231330103-0310321302231323-1220111232221310-2001231130313210-3110123101002233-2111000100202022)
- where.virtual_site.ref

<a id="canonical-3010232111333021-3302331002321313-0333132230130020-3200302032021021-2010111221012322-3331003021101332-0120220230220113-1300032212331322"></a>

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130232320203210-1110100303123322-3001023120122120-0330121312331013-0200122231330233-2203233133100333-1002120110012221-1002133301201333"></a>

## Direct properties — ref / 111013111020 / 3

<a id="canonical-1013220213110331-3132022112000300-1032002230233100-3203210233121213-3313303012102210-2333113212233310-2103112303102211-2302223031302012"></a>

<a id="canonical-3102123211133131-2102332110123312-3231233212122303-1020001130203331-3120202002200120-0021120230102232-3213230101033223-1333033210323022"></a>

## kind property — ref / 111013111020 / 4

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

<a id="canonical-2223132011120231-3211131210103112-1012100102231011-3310103322031232-3302321113212103-1202213232232103-0000123320300103-1332300032131132"></a>

<a id="canonical-1112210221033000-2013321111232001-3000231230202102-3032131320020031-2330021002220101-1121321100310213-2310310212230231-2223221120221111"></a>

## name property — ref / 111013111020 / 5

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

<a id="canonical-2030033221122030-3220122311333223-0302233222312322-0012201211110121-2011030230231101-1212332332222012-3300113131122123-3001100112233323"></a>

<a id="canonical-2001013222203131-0100011232101030-0313003022112200-3031311332321112-3031333221230010-2302113133133311-3300230212132202-0011032312313222"></a>

## namespace property — ref / 111013111020 / 6

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

<a id="canonical-0103022030223130-3201221203030300-3123311302131221-2202100012011021-1233311021001201-1321210200211233-0000223123333120-1003232303030132"></a>

<a id="canonical-1010203030301231-0120102220002230-2121220131011120-3103212201023212-1103023133110331-2201230021320013-3300311120032212-0200313110000211"></a>

## tenant property — ref / 111013111020 / 7

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

<a id="canonical-3001101111122202-3010120303201001-3111003210020221-2113313100111323-1221213232131312-1301330022133321-2132003311002221-1300000020122221"></a>

<a id="canonical-3322103122311320-2331012103232031-1103022312021211-3233001233223103-2010303111201300-1031121120303212-3303323133001221-0120123123311300"></a>

## uid property — ref / 111013111020 / 8

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

<a id="canonical-0113221101322101-2333003031010123-3203101312000223-2330022223333222-1002123001001231-2021100113101301-0331011011032332-0203020031022031"></a>

## Next pages — ref / 111013111020 / 9

- [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-0233032123312002-0020322311112123-1313223231330103-0310321302231323-1220111232221310-2001231130313210-3110123101002233-2111000100202022)
- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
