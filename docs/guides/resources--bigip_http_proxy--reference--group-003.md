---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-3020111323231331-0022202122030313-0313230112033301-0300021003303233-1131302111303033-2230231121321332-2010023131020102-0321123202212221"></a>

## name property — virtual_site / 201202231311 / 4

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

<a id="canonical-2011132310020120-0103301113032110-2233212221323013-1030022200322211-1100120121130111-2100302221021131-3233201230301322-3123332120032111"></a>

<a id="canonical-1131000023121131-2002113231313002-0232121320113130-0122100002020110-0101200011301102-1013200122122302-1122000320021023-2123003323131102"></a>

## namespace property — virtual_site / 201202231311 / 5

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

<a id="canonical-2103301231100202-2331001223133322-3110313310200333-2212021103313101-0102333210223031-2210313310233021-1212221121221320-3101100333231030"></a>

<a id="canonical-1110003312001210-0032132332111232-2122030122001233-3131121333021232-3331223321101030-1322302212202003-1233201030321300-0323333321003012"></a>

## tenant property — virtual_site / 201202231311 / 6

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

<a id="canonical-2103330203201013-0200332001310221-2322023201111323-2230210120203021-3333231330222012-0321210201313013-0321130030133111-0332103022132313"></a>

## Next pages — virtual_site / 201202231311 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-2210222003023302-3322232233330320-1133232011210322-2302301032121031-3332022221001312-0223333301221201-1212020213132130-1210122103301303)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3201012210311222-0013322302303000-0033333010112020-1110200303031321-2310032110213210-2223222313100032-1222203010020321-1021221212131032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023112010011332-3302211322231133-3331033221323130-2131211103311202-2222332010033021-0203223331221211-2330122332010132-2212301302311212"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service — vk8s_service / 131211033002 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service

<a id="canonical-3301003031203220-1330003030032321-0233023302121333-1101120101330111-3003011300032322-3032233230133012-2022332121321220-0020132330021030"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
vk8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111101211123211-1101001001320323-2231212102313122-2121302002131012-0110303210120010-1011322210031312-0232311111321320-1122223030120312"></a>

## Direct properties — vk8s_service / 131211033002 / 3

- [site](resources--bigip_http_proxy--reference--group-003.md#canonical-3032002303300103-3232203103200230-3331001123323011-3131201300332120-2220320123203201-1301232230000001-2113021333321222-2010000031311032): complete subsection reference.

- [virtual_site](resources--bigip_http_proxy--reference--group-003.md#canonical-2201030210312032-3310031013032100-2321330132010000-2312323032101330-3202100123013002-1200320032230213-3121002102010103-1032302201021202): complete subsection reference.

<a id="canonical-1112202021300132-0323231300330122-2231131030030211-3313213030223003-1332120322313133-0232021110111122-1311320013323031-2331130110303002"></a>

## Next pages — vk8s_service / 131211033002 / 4

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](resources--bigip_http_proxy--reference--group-003.md#canonical-3032002303300103-3232203103200230-3331001123323011-3131201300332120-2220320123203201-1301232230000001-2113021333321222-2010000031311032)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--bigip_http_proxy--reference--group-003.md#canonical-2201030210312032-3310031013032100-2321330132010000-2312323032101330-3202100123013002-1200320032230213-3121002102010103-1032302201021202)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3032002303300103-3232203103200230-3331001123323011-3131201300332120-2220320123203201-1301232230000001-2113021333321222-2010000031311032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230001231301113-3301203221010223-3022023212213330-3331320333020133-3000110030220202-2011331220320120-3121002120032010-3213101121013111"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site — site / 332213002111 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--bigip_http_proxy--reference--group-003.md#canonical-3201012210311222-0013322302303000-0033333010112020-1110200303031321-2310032110213210-2223222313100032-1222203010020321-1021221212131032)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-0310012031112203-2223132312203203-0001211313230131-2222231102010101-2132220320120333-3330331232132100-3220122023022102-1000233332011310"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230220022032022-1013320331120302-3222312133232113-0133032320212212-1212301123111102-0032132320121223-2123123023130122-1333020123103221"></a>

## Direct properties — site / 332213002111 / 3

<a id="canonical-3332121322001001-1102133321203012-1332213032333031-2102201023102331-3123101321021122-0211003010021032-1021011023200320-1301112130122321"></a>

<a id="canonical-0023013212333102-2333110011100213-1132220111331202-0123222222012013-0112210221012333-3022013300220220-1220231113113303-3031123311311221"></a>

## name property — site / 332213002111 / 4

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

<a id="canonical-2202222301231101-2110321011200232-1201322323100130-3231011230322033-0321322320103212-2220201033301013-1301112010311102-3332112121000232"></a>

<a id="canonical-3313232131121200-1223302011121130-3312220032132223-3121300201121320-3211231111010213-1232000231110122-1310130221230111-0110330103002032"></a>

## namespace property — site / 332213002111 / 5

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

<a id="canonical-0231132121332322-2101133002010123-2212333310313321-1031203100103133-2020300221011210-2003202232322121-1220211311303132-0131130333021313"></a>

<a id="canonical-3220100230101103-0312131110233300-2232323112200330-2102101302132232-0111323230233133-3100122212100210-3230202200001013-0211100123111330"></a>

## tenant property — site / 332213002111 / 6

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

<a id="canonical-3012132332311330-2323313232300002-3012133133102023-3221003232030213-2311222211133301-1013201020312312-2101211020312321-1321332131320212"></a>

## Next pages — site / 332213002111 / 7

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--bigip_http_proxy--reference--group-003.md#canonical-3201012210311222-0013322302303000-0033333010112020-1110200303031321-2310032110213210-2223222313100032-1222203010020321-1021221212131032)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2201030210312032-3310031013032100-2321330132010000-2312323032101330-3202100123013002-1200320032230213-3121002102010103-1032302201021202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202230231233123-1012001330233102-2131030021020232-1302100000121020-3020121323301332-2222111022202000-3330220201313001-1030102020002120"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site — virtual_site / 031213313003 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--bigip_http_proxy--reference--group-003.md#canonical-3201012210311222-0013322302303000-0033333010112020-1110200303031321-2310032110213210-2223222313100032-1222203010020321-1021221212131032)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-2211312030113003-1121203021221302-3130311300032212-0201233220321002-3013003132322011-1300133231011223-0300201001113010-1212112001000130"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133203300013300-1332103311123031-2131222032212222-2310323201223223-1333212302233220-0122210330030131-3120100103203102-1102220321223311"></a>

## Direct properties — virtual_site / 031213313003 / 3

<a id="canonical-2023311033103001-1233202311221312-1310221033223232-0013213213010321-2312321102101120-0121201102100012-0031121210113011-0302200210230322"></a>

<a id="canonical-2202002221331002-0011020301211332-2213330121212113-0311220320123200-2202303321023202-3203311210121100-2101320002303120-2023122021210211"></a>

## name property — virtual_site / 031213313003 / 4

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

<a id="canonical-2012122122320120-2312112011002133-3003321120300111-1213230023213320-3200301012202232-0323022100001201-0122230300122200-0232020001312013"></a>

<a id="canonical-0033220212321002-3300313113100032-0202121110132312-3023113210020120-2102223210011023-3211231100013330-2112120221303302-3112120013120001"></a>

## namespace property — virtual_site / 031213313003 / 5

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

<a id="canonical-0320212002300213-2133232013121212-2203101331310233-1203021332113311-2320321100210010-2212332232032331-0232202222103133-1333213203032212"></a>

<a id="canonical-3011320120320110-1210312312001011-1221321120230033-2333330220022121-2222031011010020-1103001202303323-2312031010122033-3023220321010012"></a>

## tenant property — virtual_site / 031213313003 / 6

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

<a id="canonical-1123110320222233-2331131303200211-3133020122013233-3110023211030211-2332112311011330-0312310211030002-1022201011212300-3033130300031102"></a>

## Next pages — virtual_site / 031213313003 / 7

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--bigip_http_proxy--reference--group-003.md#canonical-3201012210311222-0013322302303000-0033333010112020-1110200303031321-2310032110213210-2223222313100032-1222203010020321-1021221212131032)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1332033121310100-2011203031301030-3100012202102111-0231031031001013-0100201312332130-2311031012211232-0303100202100230-3112212033310122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030103111030300-1131132322210130-2033232332211032-3120110022203210-2211012101121221-2031103022232123-0203001111223132-1113022133013203"></a>

## proxy_advertisement.do_not_advertise — do_not_advertise / 002322121032 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- proxy_advertisement.do_not_advertise

<a id="canonical-1020302001020012-1111200022023222-0013230011313120-3011031001012211-3220221112212112-3313110121011001-3330011201311013-3230123323013210"></a>

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

<a id="canonical-3121130121200001-2311110312121201-1103313021130113-3033233200022011-1031330312220313-0131123021030301-3133111201031222-3021031202121311"></a>

## Direct properties — do_not_advertise / 002322121032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133333222003210-3023122023001213-0021203321103313-1011132322003013-0331000210032221-2022010110122332-1121300033102010-0230330220021111"></a>

## Next pages — do_not_advertise / 002322121032 / 4

- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010212002320300-0002323133310113-2121011332201010-3113102030332203-0110221010021223-2211230110111200-0130112030221013-1133223021202030"></a>

## proxy_config — proxy_config / 120023032301 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- proxy_config

<a id="canonical-0321032001132330-3220030230232212-1130321320202323-1011112321033102-0312012203223033-1331130030003002-0021310112303213-3332210012102121"></a>

Type: `"object"`. single nested block, Optional.

HTTP/HTTPS Load Balancer. HTTP/HTTPS Load balancer.

Upstream description:

HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "https_auto_cert"),
  validators.ConflictingObjectAttributes("https",
    "https_auto_cert")}
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
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]"
}
```

Terraform syntax:

```terraform
proxy_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332213233232120-2131210011211033-3212231133121301-0201300113003320-3102300033100222-2202023101103131-1021112321033212-3001022131211101"></a>

## Direct properties — proxy_config / 120023032301 / 3

<a id="canonical-0323031323301101-1123202220310110-0213023010203111-1211132003303011-3333123002012113-0013302202301133-1033002132222030-1301222000000021"></a>

<a id="canonical-3032111200212101-0032321131121320-1321332010200201-1320103113312330-3110201102223023-0133200222110113-0123002220313100-3210100203211301"></a>

## domains property — proxy_config / 120023032301 / 4

Type: `["list", "string"]`. Optional.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\`.

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*-bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*-bar.example.com\`\` will match
\`\`baz-bar.example.com\`\` but not \`\`-bar.example.com\`\`. The longest wildcards match first.

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

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

- [http](resources--bigip_http_proxy--reference--group-003.md#canonical-3120122013321203-0002100130200113-1320200332112232-3110210102320231-1333133322131232-0010022221131321-2221331010133313-3113033222022112): complete subsection reference.

- [https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111): complete subsection reference.

- [https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312): complete subsection reference.

<a id="canonical-2022300302211030-2121133223330221-1213322020000103-3102030130232113-1111231321303031-0230102301013020-2322222120221300-3101332313232210"></a>

## Next pages — proxy_config / 120023032301 / 5

- [proxy_config.http](resources--bigip_http_proxy--reference--group-003.md#canonical-3120122013321203-0002100130200113-1320200332112232-3110210102320231-1333133322131232-0010022221131321-2221331010133313-3113033222022112)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-2220212302213032-1131121012011221-0303230121313331-2110012310212223-2113332302210032-1330011112130303-0322200200222303-0213332211313312)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3120122013321203-0002100130200113-1320200332112232-3110210102320231-1333133322131232-0010022221131321-2221331010133313-3113033222022112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003210200013012-3013033232303122-2021132213020310-2131133130201201-0033202130022130-1232221000333331-2202312012011002-2032330300212222"></a>

## proxy_config.http — http / 101003113331 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- proxy_config.http

<a id="canonical-2033312303200203-3102110023031002-2221300310210132-2211131003033322-0130333102110000-3122333132321131-3203232331122012-3212113300332313"></a>

Type: `"object"`. single nested block, Optional.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("port",
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
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232311213031101-0131213230112033-3332302023110231-3022022210223333-1122030020312201-3021000132111312-0012031130133022-3030102131002100"></a>

## Direct properties — http / 101003113331 / 3

<a id="canonical-1130111333031312-1312121322022202-1303200330122021-0020232133222311-2221010102201131-1232333002230100-1030213112103200-1103232223132113"></a>

<a id="canonical-3303112110233102-3201100333302121-3133233103101320-1303321101031120-0322101130100120-3302130220033113-2330122002323303-2200300210010320"></a>

## dns_volterra_managed property — http / 101003113331 / 4

Type: `"bool"`. Optional.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0032331110112333-3232212233310031-2332113220323330-1120310101333010-3200332132330002-0203322311233013-0012102322202033-3323223212202131"></a>

<a id="canonical-3332310122002113-2132033322013312-2020221020122132-2020231000331221-1123102232022331-2321111221323011-1302330120331011-2231200110010211"></a>

## port property — http / 101003113331 / 5

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

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

<a id="canonical-2112000220313120-2212203110023100-0123222211230130-1301120032301320-2231320311330033-0100212222110111-2020232032111320-1031003200012233"></a>

<a id="canonical-3011102201303133-3210001221031212-2021330031113132-1221200233110003-3021122321110203-3201323030113001-1233302213233031-2101033232130023"></a>

## port_ranges property — http / 101003113331 / 6

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-1213302132122113-0112233231222030-1032312000122110-1322302103230010-2213332222221232-0000210120300321-3012301020001223-1323312023100222"></a>

## Next pages — http / 101003113331 / 7

- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233130103303123-1001313130120220-1302000222100033-0322333330123001-0311023033120102-3311221003302232-0202121013223300-0123213023033111"></a>

## proxy_config.https — https / 110230321002 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- proxy_config.https

<a id="canonical-3113232202311223-0130121132011112-2020022212012100-1232012102222301-0330223002132010-2011010113123323-3332121013003100-1231031310202301"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_parameters")}
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
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203112122111013-3301123301333000-3321020232323110-2000022023101020-0030202201232131-1303213223201223-3300012321331113-3201001101230322"></a>

## Direct properties — https / 110230321002 / 3

<a id="canonical-0011233003023213-2013302103113212-3200223313232323-1112231002120013-0011123120133300-0230333100230231-1130131101222100-2003103121000020"></a>

<a id="canonical-2220101221203011-1023012323130003-1321030213011010-3230030121131112-1002302211031121-2112220313012210-3132102332210222-0002130021222111"></a>

## add_hsts property — https / 110230321002 / 4

Type: `"bool"`. Optional.

Add HTTP Strict-Transport-Security response header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0330000322012021-1210300011103221-1220023212022000-3322223331213330-0213003023103331-1113030322231233-0120123113111112-2000030130300023"></a>

<a id="canonical-2112103012313300-1323103101330011-3221233112230201-0031320303321303-0212032100211301-0132002112002232-3113202230122222-2023201033212323"></a>

## append_server_name property — https / 110230321002 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

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

- [coalescing_options](resources--bigip_http_proxy--reference--group-003.md#canonical-0220013003110000-1011123121320123-0211100123130311-3133200103222013-0102323300313331-3332301003131101-2003012100230103-3101321311200333): complete subsection reference.

<a id="canonical-3231121300212122-0311230003213132-0312312112230131-1311211332001300-2230120211002231-2310120211102232-2031013311201322-2101002122000030"></a>

<a id="canonical-0320201322322010-1200000331203132-3320120132230020-1311212030232102-2302021333031130-2010130300313112-2301200333210032-1222021001012311"></a>

## connection_idle_timeout property — https / 110230321002 / 6

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

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

- [default_header](resources--bigip_http_proxy--reference--group-003.md#canonical-2113133303201003-2332010203301203-0101132133223020-2103011222121331-3331032000330032-1122031221022100-1320003323032023-2032210102100232): complete subsection reference.

- [default_loadbalancer](resources--bigip_http_proxy--reference--group-003.md#canonical-1100001333201312-1023331023100332-0212202323021102-1101331122232322-0231020133013033-3313332200103113-0010022200133210-3030203301323022): complete subsection reference.

- [disable_path_normalize](resources--bigip_http_proxy--reference--group-003.md#canonical-1021220202312122-1322031221010112-0102301300133130-2021113023103020-3112011122321213-0132002221323021-0120121133211323-0132013313102131): complete subsection reference.

- [enable_path_normalize](resources--bigip_http_proxy--reference--group-003.md#canonical-1202213120122303-3301230233233302-1112111001003131-3220033321332232-1303003103222301-3223033131003333-1231200112123320-1300123301103120): complete subsection reference.

- [http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113): complete subsection reference.

<a id="canonical-1011230322303111-2001310102000020-1123033332310231-1030321003330211-1030210002133002-0001102121322320-0333122011231201-2012213213032003"></a>

<a id="canonical-2302230020233202-2100103120301332-2230111231112311-1230323323321030-2013131212332103-3321131332233223-1223100001101012-1111221002012202"></a>

## http_redirect property — https / 110230321002 / 7

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [non_default_loadbalancer](resources--bigip_http_proxy--reference--group-003.md#canonical-0122020120213001-1002322011131121-1002302020013313-1333031232020323-3221231022002002-2330001003133021-3333221020001033-3220323333010003): complete subsection reference.

- [pass_through](resources--bigip_http_proxy--reference--group-003.md#canonical-2310222311302101-2121110111210313-3033122233131333-3320311202312031-1102212112100311-2120323112302121-3313233100202102-3132232232313210): complete subsection reference.

<a id="canonical-1313221102122210-2000201233200131-1302220023200122-0200100012021023-0021013233310222-0123230221201223-3131303013212313-3133031223012320"></a>

<a id="canonical-1301213320222020-2002132333230033-2111001130321313-0223303313323020-3320212011101133-0332312323022122-3311110311331010-2011130011212331"></a>

## port property — https / 110230321002 / 8

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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

<a id="canonical-3130020100020311-1330013303003301-3133013022113132-2220101323220121-0330111311122220-2003002013300131-0100322302210330-1201013301313220"></a>

<a id="canonical-0211032123003130-1113302132012000-2121123312231201-2210210313020021-0122021301102002-1010313121302013-0333222110331302-0112301223103322"></a>

## port_ranges property — https / 110230321002 / 9

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-3031332221013331-1010110220203210-1200013201210120-2101113123120020-0123210311133203-2322230202311301-0223130020200313-1010221032232133"></a>

<a id="canonical-1103233312023001-1110202023120121-3203203133310213-1231300012302002-1231100220330332-0013121301033032-1010312333222123-0333123022002202"></a>

## server_name property — https / 110230321002 / 10

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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

- [tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102): complete subsection reference.

- [tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013): complete subsection reference.

<a id="canonical-3010303322201130-1130012310232200-1031120312032112-3330212223103032-2203110211320322-3331010132321133-0221112202210211-3210102003131332"></a>

## Next pages — https / 110230321002 / 11

- [proxy_config.https.coalescing_options](resources--bigip_http_proxy--reference--group-003.md#canonical-0220013003110000-1011123121320123-0211100123130311-3133200103222013-0102323300313331-3332301003131101-2003012100230103-3101321311200333)
- [proxy_config.https.default_header](resources--bigip_http_proxy--reference--group-003.md#canonical-2113133303201003-2332010203301203-0101132133223020-2103011222121331-3331032000330032-1122031221022100-1320003323032023-2032210102100232)
- [proxy_config.https.default_loadbalancer](resources--bigip_http_proxy--reference--group-003.md#canonical-1100001333201312-1023331023100332-0212202323021102-1101331122232322-0231020133013033-3313332200103113-0010022200133210-3030203301323022)
- [proxy_config.https.disable_path_normalize](resources--bigip_http_proxy--reference--group-003.md#canonical-1021220202312122-1322031221010112-0102301300133130-2021113023103020-3112011122321213-0132002221323021-0120121133211323-0132013313102131)
- [proxy_config.https.enable_path_normalize](resources--bigip_http_proxy--reference--group-003.md#canonical-1202213120122303-3301230233233302-1112111001003131-3220033321332232-1303003103222301-3223033131003333-1231200112123320-1300123301103120)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113)
- [proxy_config.https.non_default_loadbalancer](resources--bigip_http_proxy--reference--group-003.md#canonical-0122020120213001-1002322011131121-1002302020013313-1333031232020323-3221231022002002-2330001003133021-3333221020001033-3220323333010003)
- [proxy_config.https.pass_through](resources--bigip_http_proxy--reference--group-003.md#canonical-2310222311302101-2121110111210313-3033122233131333-3320311202312031-1102212112100311-2120323112302121-3313233100202102-3132232232313210)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0220013003110000-1011123121320123-0211100123130311-3133200103222013-0102323300313331-3332301003131101-2003012100230103-3101321311200333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021110203310021-1332231331012133-2233131233310002-2333303301022033-3232230031123212-2331222012321301-0212301010302013-0202021131200223"></a>

## proxy_config.https.coalescing_options — coalescing_options / 220210002333 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.coalescing_options

<a id="canonical-1123120103122012-3002213212320332-1213303222100133-2313330212013122-2130201213310200-2023213132130031-3212033231211013-3303232011310133"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

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

<a id="canonical-2001122333212210-3302223003323301-2111231200210010-3233132312003331-0012113232300203-1300000121021333-0333000020001012-2021323320120110"></a>

## Direct properties — coalescing_options / 220210002333 / 3

- [default_coalescing](resources--bigip_http_proxy--reference--group-003.md#canonical-1312103010121233-2211320312212203-3322211300123333-0300311130132213-0332021111032223-0022202002033213-0323030203230320-3023032113212232): complete subsection reference.

- [strict_coalescing](resources--bigip_http_proxy--reference--group-003.md#canonical-2121023322022333-0122312212302220-0132301323100121-0233323321200233-0123003322211030-2203311233131231-1330133230121010-1131133012213112): complete subsection reference.

<a id="canonical-1222111232032310-2023230211311220-3023321310121002-0110123111302232-2111320001100212-3330111232211303-3331332211203301-1010300111322121"></a>

## Next pages — coalescing_options / 220210002333 / 4

- [proxy_config.https.coalescing_options.default_coalescing](resources--bigip_http_proxy--reference--group-003.md#canonical-1312103010121233-2211320312212203-3322211300123333-0300311130132213-0332021111032223-0022202002033213-0323030203230320-3023032113212232)
- [proxy_config.https.coalescing_options.strict_coalescing](resources--bigip_http_proxy--reference--group-003.md#canonical-2121023322022333-0122312212302220-0132301323100121-0233323321200233-0123003322211030-2203311233131231-1330133230121010-1131133012213112)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1312103010121233-2211320312212203-3322211300123333-0300311130132213-0332021111032223-0022202002033213-0323030203230320-3023032113212232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020320302003331-2231130013231031-0311323013210221-3312303131113311-1002201332300022-3233213311302323-0130310312100030-2033322331213302"></a>

## proxy_config.https.coalescing_options.default_coalescing — default_coalescing / 022133223202 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.coalescing_options](resources--bigip_http_proxy--reference--group-003.md#canonical-0220013003110000-1011123121320123-0211100123130311-3133200103222013-0102323300313331-3332301003131101-2003012100230103-3101321311200333)
- proxy_config.https.coalescing_options.default_coalescing

<a id="canonical-1113201100203023-0322202001122031-0002201311000232-1120010130333102-0212120212200132-0102302130320210-3003212313102322-2031131020201233"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_coalescing = {}
```

<a id="canonical-0013212312121333-1223203012312320-1023100112023303-0222001333322031-3111300213220203-0102121231311232-2203022001301322-1020101120311003"></a>

## Direct properties — default_coalescing / 022133223202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123202102130003-0330021110130323-0303210201130210-1332202200013320-2030100231313203-2110002111112321-1123212031320301-3012332021021300"></a>

## Next pages — default_coalescing / 022133223202 / 4

- [proxy_config.https.coalescing_options](resources--bigip_http_proxy--reference--group-003.md#canonical-0220013003110000-1011123121320123-0211100123130311-3133200103222013-0102323300313331-3332301003131101-2003012100230103-3101321311200333)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2121023322022333-0122312212302220-0132301323100121-0233323321200233-0123003322211030-2203311233131231-1330133230121010-1131133012213112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323033213020221-2332123002330301-2211030103131321-1223003002231130-0322123303032130-3111322332101332-0132010122031000-3110022020101302"></a>

## proxy_config.https.coalescing_options.strict_coalescing — strict_coalescing / 033221110023 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.coalescing_options](resources--bigip_http_proxy--reference--group-003.md#canonical-0220013003110000-1011123121320123-0211100123130311-3133200103222013-0102323300313331-3332301003131101-2003012100230103-3101321311200333)
- proxy_config.https.coalescing_options.strict_coalescing

<a id="canonical-1303222220230232-1301121333201113-3211221020011002-0022133222100322-3223331321210320-0032332321110223-1010023311230313-1331133032001101"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
strict_coalescing = {}
```

<a id="canonical-2313313303003301-3220131003321213-0213320031301020-0121313233102312-2110011033201120-1331113021332333-1322112020322032-2332130232120123"></a>

## Direct properties — strict_coalescing / 033221110023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131302301022031-2203001033230003-2013122332301222-0020331211300021-1330131113122023-2221022032131111-0211110233100120-1302001001221110"></a>

## Next pages — strict_coalescing / 033221110023 / 4

- [proxy_config.https.coalescing_options](resources--bigip_http_proxy--reference--group-003.md#canonical-0220013003110000-1011123121320123-0211100123130311-3133200103222013-0102323300313331-3332301003131101-2003012100230103-3101321311200333)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2113133303201003-2332010203301203-0101132133223020-2103011222121331-3331032000330032-1122031221022100-1320003323032023-2032210102100232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201012010202102-2022110312021332-0123101131000202-1003013123102132-1322023130021010-2320122132330331-3312211111233103-2301001233123200"></a>

## proxy_config.https.default_header — default_header / 233200012223 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.default_header

<a id="canonical-2333033300201021-1330001010030110-3013123320212003-2231113231011030-3020112302212202-1123103312231003-1203013012330222-1320301312330110"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_header = {}
```

<a id="canonical-1001213311111101-3330332222332232-2101301313013231-2010333103021332-0023223111120222-1222100211222303-0210101122120133-0213312232300130"></a>

## Direct properties — default_header / 233200012223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131012032212331-2101130030331301-0332011011202220-0001320101023023-2001332002131101-3023230230002332-0331231230202310-2012322110010310"></a>

## Next pages — default_header / 233200012223 / 4

- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1100001333201312-1023331023100332-0212202323021102-1101331122232322-0231020133013033-3313332200103113-0010022200133210-3030203301323022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302333222020020-2001312323311000-1130230322100222-2310320013123311-1002001002210200-1021002303133012-3212201321121320-3220013302031022"></a>

## proxy_config.https.default_loadbalancer — default_loadbalancer / 310320333020 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.default_loadbalancer

<a id="canonical-0311203212033031-3312023023011121-0320122331211221-1111313321230321-0000211133211113-2102031332230220-1011301333220202-1233110311011322"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_loadbalancer = {}
```

<a id="canonical-3203000021231320-1103112323302321-2302113233013203-2313002111021100-3023221230211201-3012230300102212-3322111300301000-1113100031312221"></a>

## Direct properties — default_loadbalancer / 310320333020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223321010202210-3123321201011310-2332112233113131-3311221210201223-1101102320133321-3203211132132112-3321122123212311-0031202213303200"></a>

## Next pages — default_loadbalancer / 310320333020 / 4

- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1021220202312122-1322031221010112-0102301300133130-2021113023103020-3112011122321213-0132002221323021-0120121133211323-0132013313102131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011333003333023-1311023230213022-3302232211023212-0331212200003331-3203010120102202-0012133203323002-2303222332320321-1211312123112033"></a>

## proxy_config.https.disable_path_normalize — disable_path_normalize / 121022110030 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.disable_path_normalize

<a id="canonical-3223230200123312-3320312110313232-3113113123100232-2330300313111023-0223001001200313-2013300220100200-1330312303221131-0033232210120210"></a>

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
disable_path_normalize = {}
```

<a id="canonical-3031021011003223-0332102203011323-3321222211130312-1110020333031200-2122000130213033-0100121113302321-3222333132000301-1211233201121302"></a>

## Direct properties — disable_path_normalize / 121022110030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211013221333111-3330023310211113-2002322112130331-0233133322322300-0231223120003123-3022100002122022-2130322222113032-3212003001313023"></a>

## Next pages — disable_path_normalize / 121022110030 / 4

- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1202213120122303-3301230233233302-1112111001003131-3220033321332232-1303003103222301-3223033131003333-1231200112123320-1300123301103120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123302101122100-3233313311020101-0223202211310110-2232002301301030-3323002011330233-2023300013112201-2120201201020200-0330312020233233"></a>

## proxy_config.https.enable_path_normalize — enable_path_normalize / 300232202100 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.enable_path_normalize

<a id="canonical-3002223111210111-3303032202311333-2003003230033302-0300003121103233-1321000030203010-1022233010011000-1101211102211102-1223322100200222"></a>

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
enable_path_normalize = {}
```

<a id="canonical-1213213312310211-2203131011122331-0013000020031120-3202303002103122-2113232313120212-0213100200233311-2001000013233113-0223030122110332"></a>

## Direct properties — enable_path_normalize / 300232202100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021023330211002-0023222103012010-0000133112321100-1113230113102202-0323213112322012-0332200220032223-2102012212020001-2121103023031020"></a>

## Next pages — enable_path_normalize / 300232202100 / 4

- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203312220323030-2102113331323223-2123322220302103-0122112011100110-1122000203301032-0030230203002220-2032213002010222-1020331300102211"></a>

## proxy_config.https.http_protocol_options — http_protocol_options / 001031023020 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.http_protocol_options

<a id="canonical-2002102330332111-1023211222312101-2023003021223313-3202222210223131-0212333233320211-2010132210230002-1310020003302130-3112301200103303"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
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
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232133211230120-2231101220010131-1023012320130321-1022200010300011-2331200330323302-0101113221222221-2130311211021212-2013122320031203"></a>

## Direct properties — http_protocol_options / 001031023020 / 3

- [http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-0313013020312011-2221103131030112-0220030323120120-0110023302321311-1222312200331132-1201132212330331-2232221200202231-1223233021132232): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--bigip_http_proxy--reference--group-003.md#canonical-0210030223123132-2223301031131002-0012120003033233-1001211032112301-2302300030200311-0010200230211001-2320333023331210-0313312310021121): complete subsection reference.

- [http_protocol_enable_v2_only](resources--bigip_http_proxy--reference--group-003.md#canonical-3111003033233120-2303101110200112-0001132101323022-2310111232301102-0210131222023233-1012002332220302-0010103003323133-1320201033200231): complete subsection reference.

<a id="canonical-0330023020013330-0230312120322012-3022032332010133-3300120032221332-1301130132000133-0231311300332323-0200110123122001-3003021330130201"></a>

## Next pages — http_protocol_options / 001031023020 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-0313013020312011-2221103131030112-0220030323120120-0110023302321311-1222312200331132-1201132212330331-2232221200202231-1223233021132232)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2](resources--bigip_http_proxy--reference--group-003.md#canonical-0210030223123132-2223301031131002-0012120003033233-1001211032112301-2302300030200311-0010200230211001-2320333023331210-0313312310021121)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v2_only](resources--bigip_http_proxy--reference--group-003.md#canonical-3111003033233120-2303101110200112-0001132101323022-2310111232301102-0210131222023233-1012002332220302-0010103003323133-1320201033200231)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0313013020312011-2221103131030112-0220030323120120-0110023302321311-1222312200331132-1201132212330331-2232221200202231-1223233021132232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322011332313113-3022301222110301-2301322223010123-3320132321120200-2101001002030121-0033210101311011-1312022323002231-3110110312101202"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 232020103023 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-2130021322001000-3311020211311310-0012122121310011-2230133301010102-2022231302112302-2220300310312013-0031130313222333-3330030031031133"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-3002210021111301-3003120030300122-2023210120012322-0003010132203021-3032003333311231-3220120111101112-3020203332201113-2001222023320202"></a>

## Direct properties — http_protocol_enable_v1_only / 232020103023 / 3

- [header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-2232331113101101-1303111213011013-2133032203200333-2113230322210220-2102030302011300-3312003131232323-2010132203213013-3302123331122130): complete subsection reference.

<a id="canonical-1103010010100313-1022113330203322-3221001101323302-3022201301201030-0010301013021100-1123213131220013-3012101002021123-3030000321302111"></a>

## Next pages — http_protocol_enable_v1_only / 232020103023 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-2232331113101101-1303111213011013-2133032203200333-2113230322210220-2102030302011300-3312003131232323-2010132203213013-3302123331122130)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2232331113101101-1303111213011013-2133032203200333-2113230322210220-2102030302011300-3312003131232323-2010132203213013-3302123331122130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310001103331331-2010110111123022-0213120111030011-0022312302303013-1313032322010321-1012123002201112-1303201323312212-0113120011030130"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 200002330023 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-0313013020312011-2221103131030112-0220030323120120-0110023302321311-1222312200331132-1201132212330331-2232221200202231-1223233021132232)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0000331132112301-2110210032122110-3003122112112200-0021302031233212-0322232122222122-0231010112223113-1201030012333122-0323030110113022"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313132230102131-0230033301220323-3133233120030202-2121303032221311-0301203332131121-0133222323221323-1021120210220202-1102211213313120"></a>

## Direct properties — header_transformation / 200002330023 / 3

- [default_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-3030022221012231-1301333130221002-2033210310110333-1110133302020231-0311233323201002-1022302033132330-1212223111033022-0113221120230231): complete subsection reference.

- [preserve_case_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-1113112123230132-0030023231003131-1002312111103200-1201310020331322-3331320023222121-1023233100000022-3303112021101212-0001131011330003): complete subsection reference.

- [proper_case_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-1223120320313001-3313200202233210-2022130230120103-3123133211022313-3212303101131131-1033023101212103-1113323033202101-2031121100101303): complete subsection reference.

<a id="canonical-1010331012131120-1312120333213103-2112002310010333-3233110213333133-3133323321102201-3310001101130033-2333010000101320-3022210132223201"></a>

## Next pages — header_transformation / 200002330023 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-3030022221012231-1301333130221002-2033210310110333-1110133302020231-0311233323201002-1022302033132330-1212223111033022-0113221120230231)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-1113112123230132-0030023231003131-1002312111103200-1201310020331322-3331320023222121-1023233100000022-3303112021101212-0001131011330003)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-1223120320313001-3313200202233210-2022130230120103-3123133211022313-3212303101131131-1033023101212103-1113323033202101-2031121100101303)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-0313013020312011-2221103131030112-0220030323120120-0110023302321311-1222312200331132-1201132212330331-2232221200202231-1223233021132232)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3030022221012231-1301333130221002-2033210310110333-1110133302020231-0311233323201002-1022302033132330-1212223111033022-0113221120230231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221233312302000-3113222201033013-1302310302103211-3311123033313111-2232321330203111-3220010100201221-3333313212120311-0202201332210232"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 112322012023 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-0313013020312011-2221103131030112-0220030323120120-0110023302321311-1222312200331132-1201132212330331-2232221200202231-1223233021132232)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-2232331113101101-1303111213011013-2133032203200333-2113230322210220-2102030302011300-3312003131232323-2010132203213013-3302123331122130)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-3021122201333311-0023112203333120-1300032212002031-0200112321201320-1111330323113331-2012133102322001-1303011200213330-3232030301012322"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_header_transformation = {}
```

<a id="canonical-2002321122311331-3031200001031210-1230312300101132-3002101031222013-2020323121000013-0320131103003100-2310121320330303-2311333123130022"></a>

## Direct properties — default_header_transformation / 112322012023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010331013033320-1021203202123331-0032321233002032-2202232121303203-1013122001323121-3322202112132102-0311021202120223-3101023033333232"></a>

## Next pages — default_header_transformation / 112322012023 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-2232331113101101-1303111213011013-2133032203200333-2113230322210220-2102030302011300-3312003131232323-2010132203213013-3302123331122130)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1113112123230132-0030023231003131-1002312111103200-1201310020331322-3331320023222121-1023233100000022-3303112021101212-0001131011330003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333011321222300-3022231331301022-3103231120133323-0021123303032100-0021021032313022-0111113120331033-2130233133302100-3111231303121133"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 221233003113 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-0313013020312011-2221103131030112-0220030323120120-0110023302321311-1222312200331132-1201132212330331-2232221200202231-1223233021132232)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-2232331113101101-1303111213011013-2133032203200333-2113230322210220-2102030302011300-3312003131232323-2010132203213013-3302123331122130)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-3333013020131220-1102232030031232-3113130221123203-2223001202022111-1221101202101003-3120302212201231-2303030222021032-3101121211132332"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
preserve_case_header_transformation = {}
```

<a id="canonical-3312321303302213-1230011213031311-0223133332032310-2000232200333020-0010333010100202-3201321022012321-1130322220312022-3110300333031101"></a>

## Direct properties — preserve_case_header_transformation / 221233003113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120000302211113-3322202213233113-0321021301222222-0031120211000022-0001103021102202-0110331313303020-2120122303301303-2111033320230322"></a>

## Next pages — preserve_case_header_transformation / 221233003113 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-2232331113101101-1303111213011013-2133032203200333-2113230322210220-2102030302011300-3312003131232323-2010132203213013-3302123331122130)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1223120320313001-3313200202233210-2022130230120103-3123133211022313-3212303101131131-1033023101212103-1113323033202101-2031121100101303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231332133110211-1321311020203002-0110200113013112-0312311201213123-3331111031133023-0102113223003013-3310213220021021-1131220321232001"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 321113330032 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-0313013020312011-2221103131030112-0220030323120120-0110023302321311-1222312200331132-1201132212330331-2232221200202231-1223233021132232)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-2232331113101101-1303111213011013-2133032203200333-2113230322210220-2102030302011300-3312003131232323-2010132203213013-3302123331122130)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-0020002331120333-3232223011311013-3012321202321013-3123020233233201-0103330222232212-1130002013010311-3112322122012320-2331032231210202"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
proper_case_header_transformation = {}
```

<a id="canonical-0013200003211221-2230303030331322-2300020112300120-3222131011332023-3223210100000121-1011023003232302-1223331033023102-0201103331021112"></a>

## Direct properties — proper_case_header_transformation / 321113330032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100110213302030-2103301010203300-0130322130102233-3101010303110120-0303113123223302-0301103201023000-2302303300032131-3032230312111000"></a>

## Next pages — proper_case_header_transformation / 321113330032 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-2232331113101101-1303111213011013-2133032203200333-2113230322210220-2102030302011300-3312003131232323-2010132203213013-3302123331122130)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0210030223123132-2223301031131002-0012120003033233-1001211032112301-2302300030200311-0010200230211001-2320333023331210-0313312310021121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313001312231023-0211123312201110-1300332000303231-1012310200122310-0323323223321332-2030231312302021-0330031023001303-1130110223312012"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 000010100310 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-1322113201320321-0020111120120212-2020032310012201-1111001122020210-0023312133022113-1021222223000331-0113131010233132-1102003322310212"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v1_v2 = {}
```

<a id="canonical-0302313130122202-3020121213312033-0331012113311121-1012312023232212-3113103322320023-1003002130033031-0123222210330022-3321010220212022"></a>

## Direct properties — http_protocol_enable_v1_v2 / 000010100310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212101300213323-3100022110330201-2303302103300131-1022231201010202-1131001121201012-0030021020110222-2121010313213002-3233112123231211"></a>

## Next pages — http_protocol_enable_v1_v2 / 000010100310 / 4

- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3111003033233120-2303101110200112-0001132101323022-2310111232301102-0210131222023233-1012002332220302-0010103003323133-1320201033200231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220012203200031-1031000001131110-3310022323020133-1113301113123011-1313000323211222-1002030013331203-1033013033321013-1033102101120021"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 133022122322 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113)
- proxy_config.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-0301101221222222-3321312321330131-2012103322002321-0322010332023030-3233030101100230-1020313210212313-1200100231023010-2211212032130031"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v2_only = {}
```

<a id="canonical-2220011300113221-3330011232022023-1122222333033233-0012021233013113-1102133223031130-2200211032021220-3013220321000101-0332111320030010"></a>

## Direct properties — http_protocol_enable_v2_only / 133022122322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120002322221003-0012002112010031-1022003332220033-0110311020031302-0101123333021323-2222221031132103-3123130103230321-0023010110231010"></a>

## Next pages — http_protocol_enable_v2_only / 133022122322 / 4

- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0122020120213001-1002322011131121-1002302020013313-1333031232020323-3221231022002002-2330001003133021-3333221020001033-3220323333010003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200302130100301-3211221131020000-1312121103133222-1312203001212222-0230101201132002-3132321022111312-0113131330020100-1332111232020111"></a>

## proxy_config.https.non_default_loadbalancer — non_default_loadbalancer / 002213300312 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.non_default_loadbalancer

<a id="canonical-3121011321013220-0322322003203110-0113320333110021-3212330003310130-0203122002230033-0021210220203030-0121110333330103-3300020120032122"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
non_default_loadbalancer = {}
```

<a id="canonical-2001030220330201-3032020020100002-2322103321322332-1232132113203002-0100232333133111-1231000110112100-2200313123101101-0032223310213233"></a>

## Direct properties — non_default_loadbalancer / 002213300312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201331032323101-3103333203022320-3333021032232110-2020322223310222-3113210333003033-2012021102232220-0231223000101230-3033223331303231"></a>

## Next pages — non_default_loadbalancer / 002213300312 / 4

- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2310222311302101-2121110111210313-3033122233131333-3320311202312031-1102212112100311-2120323112302121-3313233100202102-3132232232313210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323011021322003-3200331303331001-0103212122132222-3021010220133110-3100021302103122-3121100010320122-2102101130030031-0002123032301333"></a>

## proxy_config.https.pass_through — pass_through / 200033133302 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.pass_through

<a id="canonical-2100303113221032-3211313031330102-1010121010103101-3302131111021313-1302300103212302-2101103000311001-0223310330002100-0301201121003131"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
pass_through = {}
```

<a id="canonical-3013303013133311-0122003010031313-0332100010303303-0130223103333323-3210300030032032-2313003000213011-3330133320210202-1101331133022021"></a>

## Direct properties — pass_through / 200033133302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133013301201313-0330123033332002-0003220212101033-1100130230331122-1023020132312020-3023331202130332-2022103033113003-1121100031001303"></a>

## Next pages — pass_through / 200033133302 / 4

- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101333103320331-3011133101031300-3221301201323210-0003010012313102-1322320202133023-0300031122310012-2133031030203212-3203133132132102"></a>

## proxy_config.https.tls_cert_params — tls_cert_params / 232012232120 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.tls_cert_params

<a id="canonical-0233332000132102-1313202110002220-1210132103232310-0011130031230122-2100211121001120-0330212132020020-2331202011301010-0110310320220330"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123330031123200-2103213331301120-1010023230300330-1231023230100032-3313311223311132-2300120013301333-2233312210101132-3230023132300003"></a>

## Direct properties — tls_cert_params / 232012232120 / 3

- [certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-1013202012320033-1110130002111103-2003230201332002-3003000103011220-3103220201133133-2310231333132013-1330012122312111-0032201130223000): complete subsection reference.

- [no_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-3023033212320010-1200320302302111-3310101022321311-2102311011100020-0201320111112200-3331113213211131-2031222302101333-2133101233032223): complete subsection reference.

- [tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320): complete subsection reference.

- [use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310): complete subsection reference.

<a id="canonical-1113313121012101-3033233133023111-0123021113333021-2322001003012132-3002010301213031-3122013132331000-2331011330123212-1321223231100212"></a>

## Next pages — tls_cert_params / 232012232120 / 4

- [proxy_config.https.tls_cert_params.certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-1013202012320033-1110130002111103-2003230201332002-3003000103011220-3103220201133133-2310231333132013-1330012122312111-0032201130223000)
- [proxy_config.https.tls_cert_params.no_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-3023033212320010-1200320302302111-3310101022321311-2102311011100020-0201320111112200-3331113213211131-2031222302101333-2133101233032223)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1013202012320033-1110130002111103-2003230201332002-3003000103011220-3103220201133133-2310231333132013-1330012122312111-0032201130223000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211310121000321-0200133030230213-0203301203232021-1231303102201311-2102102002203223-0101310211102303-1223303100310212-2331113323213113"></a>

## proxy_config.https.tls_cert_params.certificates — certificates / 303002023030 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- proxy_config.https.tls_cert_params.certificates

<a id="canonical-2331102121213232-3113200110102120-2130311031232123-0023131022032013-3310110201230102-3121330232000222-1112023203131030-2233312113200011"></a>

Type: `"object"`. list nested block, Optional.

Select one or more certificates with any domain names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230120001012310-3032113230032033-1320212023131132-1001030331221022-0333301002200022-1010211203203103-3012133200032103-1010110021322220"></a>

## Direct properties — certificates / 303002023030 / 3

<a id="canonical-3020130213211231-0110022032110102-0103110201032332-3130011011102112-0000032222200003-0123132112301022-0321313310003111-2101002230132021"></a>

<a id="canonical-2023232133032220-3023210301011231-3103303111123333-1012210230321130-2010120231121212-2112012311022203-0301321312121011-3312223021223131"></a>

## name property — certificates / 303002023030 / 4

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

<a id="canonical-2221332102220321-1233103023203000-1132122302301322-0032233310020101-1333331321231302-1120223010000332-3232003120111120-1213133331301100"></a>

<a id="canonical-2232330111203031-0122011321311202-1223023001013001-0131130231202322-3332320212332331-3213130303303313-0133101202331111-0022212121130223"></a>

## namespace property — certificates / 303002023030 / 5

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

<a id="canonical-1310220210131220-0001132111131111-0320020301231232-3001111231312233-1003032012100032-2020230333123220-1220301033122022-2131101123313200"></a>

<a id="canonical-3231303010033232-2301321113232003-0201012110331331-1020230221202112-3133320203033332-2322321330231313-1132101202103331-3231011310002130"></a>

## tenant property — certificates / 303002023030 / 6

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

<a id="canonical-3330221131022310-1100220123333102-3301311203232330-1202321330332131-1222011010110122-3003013203220203-3333301122112201-2232130222031120"></a>

## Next pages — certificates / 303002023030 / 7

- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3023033212320010-1200320302302111-3310101022321311-2102311011100020-0201320111112200-3331113213211131-2031222302101333-2133101233032223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201013120330002-0312231311233222-1320201012331201-3022002312000112-1312120201113202-0322201032230130-0010030232112203-2320013121031031"></a>

## proxy_config.https.tls_cert_params.no_mtls — no_mtls / 313133110022 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- proxy_config.https.tls_cert_params.no_mtls

<a id="canonical-3321313102220133-2311213030011233-0012110132133220-3203323230330000-1013313202132332-1232000110102031-0001132120123232-0210022212030103"></a>

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
no_mtls = {}
```

<a id="canonical-1003113102312020-0220123020211022-0111023302200223-0111222101101011-0102122113200200-0320332231113310-2221102032123001-2032002231233002"></a>

## Direct properties — no_mtls / 313133110022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1203233011030303-1201002000010322-3112111333322331-3223320321321233-3203302313220003-3003102211203132-2012303200113112-2212330213313013"></a>

## Next pages — no_mtls / 313133110022 / 4

- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222121010023320-0300130023013212-3312230202211121-0203020302222002-3213132002301112-3303102330303032-1020112122200110-1323013200322232"></a>

## proxy_config.https.tls_cert_params.tls_config — tls_config / 231221331132 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- proxy_config.https.tls_cert_params.tls_config

<a id="canonical-2133230230132223-0013202033222321-1032221331203013-1213030201201130-1012122111333111-0122020311333212-0030313001002123-3011211110212103"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100123033120021-0210330001212130-0200132133123200-3121112011321032-0200201332021031-1021321102313331-0123011221202201-3230223310003221"></a>

## Direct properties — tls_config / 231221331132 / 3

- [custom_security](resources--bigip_http_proxy--reference--group-003.md#canonical-0110222220210032-0331331320002032-0330112030312222-1132120120111312-3113131202131300-0331013230113132-2112213320303132-3220320330033310): complete subsection reference.

- [default_security](resources--bigip_http_proxy--reference--group-003.md#canonical-0022301131111233-1321032023222011-3000222310010121-0203121211232323-0031002300112021-1032033202203200-0222013023000101-3113313232330132): complete subsection reference.

- [low_security](resources--bigip_http_proxy--reference--group-003.md#canonical-1320200110221000-1031112223132333-3133121330331232-0223310122222011-3103231101200312-1213130230303300-2102201100233313-1121102332232022): complete subsection reference.

- [medium_security](resources--bigip_http_proxy--reference--group-003.md#canonical-3121031223203211-3002232101112320-0213210011001032-1232303220311202-0112030222012333-3003322220101030-3022330210023031-3232211130321003): complete subsection reference.

<a id="canonical-0003100203230312-3221303302312320-2100030213320312-3203010101300220-3101103012002100-0132111133231103-3232103322313333-3110030220011101"></a>

## Next pages — tls_config / 231221331132 / 4

- [proxy_config.https.tls_cert_params.tls_config.custom_security](resources--bigip_http_proxy--reference--group-003.md#canonical-0110222220210032-0331331320002032-0330112030312222-1132120120111312-3113131202131300-0331013230113132-2112213320303132-3220320330033310)
- [proxy_config.https.tls_cert_params.tls_config.default_security](resources--bigip_http_proxy--reference--group-003.md#canonical-0022301131111233-1321032023222011-3000222310010121-0203121211232323-0031002300112021-1032033202203200-0222013023000101-3113313232330132)
- [proxy_config.https.tls_cert_params.tls_config.low_security](resources--bigip_http_proxy--reference--group-003.md#canonical-1320200110221000-1031112223132333-3133121330331232-0223310122222011-3103231101200312-1213130230303300-2102201100233313-1121102332232022)
- [proxy_config.https.tls_cert_params.tls_config.medium_security](resources--bigip_http_proxy--reference--group-003.md#canonical-3121031223203211-3002232101112320-0213210011001032-1232303220311202-0112030222012333-3003322220101030-3022330210023031-3232211130321003)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0110222220210032-0331331320002032-0330112030312222-1132120120111312-3113131202131300-0331013230113132-2112213320303132-3220320330033310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130120202011330-1111032013232013-0000213203333101-0312320220003112-2303331313110110-2130322012313213-2203113033203120-3223313321213212"></a>

## proxy_config.https.tls_cert_params.tls_config.custom_security — custom_security / 103110323303 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320)
- proxy_config.https.tls_cert_params.tls_config.custom_security

<a id="canonical-0031303213213230-2003331031200220-2032201120012222-1300113230130211-3321122000021131-1133222112012330-2331302010011132-3122231122103321"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231200201321311-3121223232011303-2112320323130123-1113303302333112-0101333331320133-3233032012210033-1303330222222023-1001333111110331"></a>

## Direct properties — custom_security / 103110323303 / 3

<a id="canonical-2100012232300311-1022312321330200-0320113031110201-2331220210122300-1332223301233032-3220331312020101-3010122131101022-2001023212020100"></a>

<a id="canonical-2113102122333012-1212121122003001-1321000011210303-0332010100123313-1233000213131220-1302011203000321-1120232231023212-2011302012321233"></a>

## cipher_suites property — custom_security / 103110323303 / 4

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2131110222323031-1310312332103231-0003203223032310-0211103112012003-0230231030321332-3330232003300312-3200302001020032-0032222000110130"></a>

<a id="canonical-0021332112123221-3311022302023023-2310120110030311-3203231221220123-0101022002101110-3133002100122031-0212213332223120-0233120301301301"></a>

## max_version property — custom_security / 103110323303 / 5

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

<a id="canonical-2120223232223003-2030201231223003-0131120133022011-3333131302213011-0322333323333033-0320321203031003-1311202231002232-1212222322102011"></a>

<a id="canonical-0222023030223221-2133322020132100-1100333021210132-3300032331031200-1323110331211202-0020103331211031-1321210011323202-0222220132130313"></a>

## min_version property — custom_security / 103110323303 / 6

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

<a id="canonical-2100311122300233-0231010230321330-3123333011223131-0123100112213020-1113302222013321-0331002212120112-2102210223203023-0333113023120210"></a>

## Next pages — custom_security / 103110323303 / 7

- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0022301131111233-1321032023222011-3000222310010121-0203121211232323-0031002300112021-1032033202203200-0222013023000101-3113313232330132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311213012210030-0213120323001110-2033102100011233-0130102102021101-1313033220333010-0021010213030103-0210213110023133-1123202231012112"></a>

## proxy_config.https.tls_cert_params.tls_config.default_security — default_security / 301020123011 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320)
- proxy_config.https.tls_cert_params.tls_config.default_security

<a id="canonical-3013330031011133-0031101232332000-3320013320321210-2022112121323010-2200333123112331-1312301032101133-0212012003303311-0031300122010323"></a>

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
default_security = {}
```

<a id="canonical-0333122333323030-2211131011002000-3320031230010220-3101230222022223-1121313312203021-2322120220002210-3011112203131323-0300001132112000"></a>

## Direct properties — default_security / 301020123011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010303030201112-3022101210220031-2333333131313313-2121300113102000-1330213130030213-0202212332010233-0231023013233132-1202023101103032"></a>

## Next pages — default_security / 301020123011 / 4

- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1320200110221000-1031112223132333-3133121330331232-0223310122222011-3103231101200312-1213130230303300-2102201100233313-1121102332232022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023330033030030-0023021231310030-2001220023230210-3333201301021011-1220131330132001-1211031223220133-3222023200332010-3101221201302131"></a>

## proxy_config.https.tls_cert_params.tls_config.low_security — low_security / 221001001010 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320)
- proxy_config.https.tls_cert_params.tls_config.low_security

<a id="canonical-2010121312320302-0101102001333221-0103032311000230-0221230332312300-0033111103133232-2123120310213121-2002231312021323-3220230230033103"></a>

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
low_security = {}
```

<a id="canonical-3100312100312230-0303121130000200-0033220133330123-3233020132210033-1133030332233001-2010033303322030-1122131213001102-3022210313201310"></a>

## Direct properties — low_security / 221001001010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023030212121031-1302031131020220-3312302221132332-2223303131201210-2031102220021001-0010312120221123-2133100032131210-3123323101033110"></a>

## Next pages — low_security / 221001001010 / 4

- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3121031223203211-3002232101112320-0213210011001032-1232303220311202-0112030222012333-3003322220101030-3022330210023031-3232211130321003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332113220013330-0223232330223023-3002012033012322-3321021232122202-0031022033213231-1023031311330022-1111003102303021-3203210120121033"></a>

## proxy_config.https.tls_cert_params.tls_config.medium_security — medium_security / 001033230200 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320)
- proxy_config.https.tls_cert_params.tls_config.medium_security

<a id="canonical-2322331100001201-2230223332010032-2110202221222210-3220010010231032-0201023110212123-3231021120303112-1132012210021202-1231003202000130"></a>

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
medium_security = {}
```

<a id="canonical-2320021203222121-0100320020230300-2100112303133310-0001133333010202-0322112300033011-3103321102211201-2131310133123033-0212123123322130"></a>

## Direct properties — medium_security / 001033230200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323023312201130-1120130101020330-0220233230132302-1202203212121213-2102122022002030-2302202300321322-0220202033313311-0302121021021111"></a>

## Next pages — medium_security / 001033230200 / 4

- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321330011113121-2111132000122133-1220013123303230-3023112133021320-3212102310103110-0002113111212113-0201213003223021-3222002011112003"></a>

## proxy_config.https.tls_cert_params.use_mtls — use_mtls / 223113301110 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- proxy_config.https.tls_cert_params.use_mtls

<a id="canonical-2002303301031111-2132312201122033-3220010121322013-0333031200101322-0203103201332033-3030000210201212-3320203123132031-1220221212230001"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133011000132210-0112022122311133-1021233300023111-0332100013122212-1101302300303222-3222012203033121-0330031220121031-0320220302003301"></a>

## Direct properties — use_mtls / 223113301110 / 3

<a id="canonical-3013302213222033-3023202323332200-2003131122120220-3031201022201333-1231332233300133-1310320201223312-0312031001032213-1303303333013133"></a>

<a id="canonical-0103222021103003-1123313333110001-0233100101132000-0313210201113010-2301102232102232-1033003030013303-3330100210222120-2233302201022321"></a>

## client_certificate_optional property — use_mtls / 223113301110 / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](resources--bigip_http_proxy--reference--group-003.md#canonical-2001103200102102-3123312201111013-0233221112320031-2110311020310222-3103322321302002-1300320230202232-3003301332010033-3032110010131020): complete subsection reference.

- [no_crl](resources--bigip_http_proxy--reference--group-003.md#canonical-2303332100000202-0220310301012200-0120330311213201-3112033222320313-2000121202023110-1030303133332302-3222301013103021-0102232123300300): complete subsection reference.

- [trusted_ca](resources--bigip_http_proxy--reference--group-003.md#canonical-2302133000120033-0002333202301303-0103122120311021-0223300032001222-3010031112020220-1001001000233201-2321230203001113-3133020333310113): complete subsection reference.

<a id="canonical-0003220231231303-3130033332303332-3322310321001232-3020013321023310-0203000303200202-3302312200201110-2210302301030312-0122303302220230"></a>

<a id="canonical-0303322110000101-2002320011223113-0212311003212001-1210310213333202-2202021320030203-1120011232200310-3203020310133011-3123123110101323"></a>

## trusted_ca_url property — use_mtls / 223113301110 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](resources--bigip_http_proxy--reference--group-003.md#canonical-3211103101231003-3022300000021000-0232311030103131-0132120231222232-2022310023101321-1131323011110101-2110320321212130-1213311133112230): complete subsection reference.

- [xfcc_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1120102033101201-0123012302213200-3120300323230003-2111123203122310-0333003320221320-3121320101322321-0123002232323132-0211310201113221): complete subsection reference.

<a id="canonical-0000110013133033-3101332320120300-3033330202231121-2031031203213001-2220113033023301-3210030222032010-3100303332110302-2131001122202020"></a>

## Next pages — use_mtls / 223113301110 / 6

- [proxy_config.https.tls_cert_params.use_mtls.crl](resources--bigip_http_proxy--reference--group-003.md#canonical-2001103200102102-3123312201111013-0233221112320031-2110311020310222-3103322321302002-1300320230202232-3003301332010033-3032110010131020)
- [proxy_config.https.tls_cert_params.use_mtls.no_crl](resources--bigip_http_proxy--reference--group-003.md#canonical-2303332100000202-0220310301012200-0120330311213201-3112033222320313-2000121202023110-1030303133332302-3222301013103021-0102232123300300)
- [proxy_config.https.tls_cert_params.use_mtls.trusted_ca](resources--bigip_http_proxy--reference--group-003.md#canonical-2302133000120033-0002333202301303-0103122120311021-0223300032001222-3010031112020220-1001001000233201-2321230203001113-3133020333310113)
- [proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled](resources--bigip_http_proxy--reference--group-003.md#canonical-3211103101231003-3022300000021000-0232311030103131-0132120231222232-2022310023101321-1131323011110101-2110320321212130-1213311133112230)
- [proxy_config.https.tls_cert_params.use_mtls.xfcc_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1120102033101201-0123012302213200-3120300323230003-2111123203122310-0333003320221320-3121320101322321-0123002232323132-0211310201113221)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2001103200102102-3123312201111013-0233221112320031-2110311020310222-3103322321302002-1300320230202232-3003301332010033-3032110010131020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332032033012021-0011300113012000-3010133023011322-2010313213302122-0120130013111121-2331022211000130-1320213303123201-2303202102311030"></a>

## proxy_config.https.tls_cert_params.use_mtls.crl — crl / 321003112112 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- proxy_config.https.tls_cert_params.use_mtls.crl

<a id="canonical-1100022302103132-2021302201320123-0122000012202212-3221132210021322-1202201031101030-1302010230302232-0323200313222132-3200011223123233"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203310010003120-0023203103001332-3110311330302233-2112023233032031-0230113301213001-0331101123312202-0332212321200010-3002023121111130"></a>

## Direct properties — crl / 321003112112 / 3

<a id="canonical-3133102110030302-1230311303021030-0313313000311230-3013300130031221-0000223202221333-1310012302032123-3332202133031212-3022312213100312"></a>

<a id="canonical-0202023030300001-1312011132021002-0131022221232103-3100210230112131-1102331100301011-3223030113322022-2222231111110133-0133202213102323"></a>

## name property — crl / 321003112112 / 4

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

<a id="canonical-0321032210002011-1013222001011202-1020121230310020-2311331030000210-3332213103312212-1120303300333012-2213031203030102-2312232330311130"></a>

<a id="canonical-0112223011113120-1122100311302110-2220330122001023-0222320001010323-3022211002130113-2300322012222203-3112232003003332-0103202132333310"></a>

## namespace property — crl / 321003112112 / 5

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

<a id="canonical-2313032301310211-1030022122233213-3220202310322302-1300012333123332-3321222110230200-2212211102222123-1011013002020001-2333112131210302"></a>

<a id="canonical-0231213000103303-1111301331330021-0031103300233322-2322020322102121-0331103132020302-3121122213010122-1233330111223223-1211232312030202"></a>

## tenant property — crl / 321003112112 / 6

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

<a id="canonical-3023332313001111-0031230013311130-1213010301201200-1030003012202111-3222232301022100-2033133312321022-0110112102030002-1300001220203132"></a>

## Next pages — crl / 321003112112 / 7

- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2303332100000202-0220310301012200-0120330311213201-3112033222320313-2000121202023110-1030303133332302-3222301013103021-0102232123300300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023031202132103-0303033301002012-1303201100021122-1233221010110220-3103221331121320-1130000013132130-2233121112220333-1222101101132311"></a>

## proxy_config.https.tls_cert_params.use_mtls.no_crl — no_crl / 101212323301 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- proxy_config.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-3020011320103021-1330123313300030-0012020133130110-0332032322023001-0332300202022133-3201101230021303-0211202030211323-1231220220130221"></a>

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
no_crl = {}
```

<a id="canonical-0213031022101201-0002012131231033-2200133020311113-0121123121201023-3031121110311211-3203102122212020-0222310220000003-2320112200311012"></a>

## Direct properties — no_crl / 101212323301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133021232220011-2002321122230233-1310333332033013-1300301230313220-0332210322213303-2113000011002013-2032202320101123-1221111333312320"></a>

## Next pages — no_crl / 101212323301 / 4

- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2302133000120033-0002333202301303-0103122120311021-0223300032001222-3010031112020220-1001001000233201-2321230203001113-3133020333310113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010323121213301-2101122211030313-2213013121220221-1101033303333220-1300012323221211-3123222031012313-0301131001103220-3200113022223113"></a>

## proxy_config.https.tls_cert_params.use_mtls.trusted_ca — trusted_ca / 123330310103 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- proxy_config.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-2000232132011120-2011013201322020-1031122031302333-1302021203011212-2020221032312302-1331000203010313-1000222011111201-0103233332312232"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322101232101103-2201101233222320-3300000202220210-0212100232102123-2121023202212223-0103303313213021-1111003233232120-1312203302233203"></a>

## Direct properties — trusted_ca / 123330310103 / 3

<a id="canonical-1003202213002313-0031002130223112-1013121001110312-3200311020323310-3102020112003321-2102022112310031-2132332310111321-0201233031302321"></a>

<a id="canonical-3221232112013123-3303120221220100-0332203120210321-1111030030313313-3300210203032212-1303001203013031-1103102113233011-0020311311322232"></a>

## name property — trusted_ca / 123330310103 / 4

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

<a id="canonical-1320010130103133-0020223220303031-3231323210000230-3023233331131231-0010223232111330-2122110013223310-2223032320210023-2302223020010120"></a>

<a id="canonical-2033220332113021-3201121030032321-0232233131223033-1302210213012002-2002021222032011-1220011212220323-0301223123130032-3332321002132223"></a>

## namespace property — trusted_ca / 123330310103 / 5

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

<a id="canonical-2122330320130303-0011132332121021-0203120133010320-1101300202120102-1102330010212331-2012033213121100-2032031012321023-3132131000303301"></a>

<a id="canonical-3221311232300200-0213312113301230-1031133220020201-0120232003123202-2303303111110322-1301022113120200-0301132023110222-0213311121202023"></a>

## tenant property — trusted_ca / 123330310103 / 6

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

<a id="canonical-2020021300022200-1330130021321113-0303031110330032-0133222112312111-1332013230320000-0231313000000031-3320133230213313-3322110012310033"></a>

## Next pages — trusted_ca / 123330310103 / 7

- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3211103101231003-3022300000021000-0232311030103131-0132120231222232-2022310023101321-1131323011110101-2110320321212130-1213311133112230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122230331211231-1313022321120210-2213031231222303-0323222013012031-3300132112013331-3312102330110112-3231202323023000-0100322023311000"></a>

## proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled — xfcc_disabled / 023013203103 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-0111033302023221-3310003303303131-1310332121313111-1013031100131132-0130010033111122-1113222120323210-3200300013123322-3203323313213131"></a>

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
xfcc_disabled = {}
```

<a id="canonical-0022203302010103-0102100103212031-1121121010333232-1220103220221203-3010100110121323-2322303203322321-2122120323230001-0110213310210303"></a>

## Direct properties — xfcc_disabled / 023013203103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213200112130013-3312212220302111-3103013320221013-2330032223330022-3201021222003130-0321222133323000-2330013121100300-2202020010100023"></a>

## Next pages — xfcc_disabled / 023013203103 / 4

- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1120102033101201-0123012302213200-3120300323230003-2111123203122310-0333003320221320-3121320101322321-0123002232323132-0211310201113221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031221300210302-3320300311312002-2000232000103122-1031021101123002-1102133200211321-3221330001013030-2000202003311330-0211230310110202"></a>

## proxy_config.https.tls_cert_params.use_mtls.xfcc_options — xfcc_options / 233001020223 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- proxy_config.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-0332212230300113-2113001211110101-1300133321321132-3203023313200121-3101231200331111-2013121011120303-2311113123200033-1003222322321120"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110133022102210-0003303230020013-1000120321200110-1221210210023013-1333303212321103-1211032120111001-0333302230210222-2223222211003032"></a>

## Direct properties — xfcc_options / 233001020223 / 3

<a id="canonical-2212233001102010-0302012322301022-1331112010300111-0102001013012231-2301313011311012-3012201000313230-0123301213020113-2001130100331302"></a>

<a id="canonical-0311110212311333-2032113121011202-1103232002223122-2303202000023303-0232122303222120-0301333230300122-2100322303322031-1213001032323311"></a>

## xfcc_header_elements property — xfcc_options / 233001020223 / 4

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-1003103223200331-0313103303330313-1112000100102131-3023121101200310-1312330320303111-0221230203310330-1123310013300302-0111332301010221"></a>

## Next pages — xfcc_options / 233001020223 / 5

- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230032303200213-1000103312100211-3302031122311020-0000120222310223-1010230210213222-3232123300121003-0200131223211133-1123212010201221"></a>

## proxy_config.https.tls_parameters — tls_parameters / 230232121332 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.tls_parameters

<a id="canonical-0120102120100020-1100011221033102-0221220220332001-1333311021222120-2023123132110312-0002031111321112-1022231022201102-3131022230333133"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-1213321121302300-2003230330022320-2021300110012030-1301100331323211-0020122133100301-1132111333000322-0230000133113212-1313310020201330"></a>

## Direct properties — tls_parameters / 230232121332 / 3

- [no_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-2230003003322321-2030203023030130-0311103033230111-2312121003013022-0122112230330223-3233101232232332-3200133323000330-0231021223022121): complete subsection reference.

- [tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101): complete subsection reference.

- [tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200): complete subsection reference.

- [use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300): complete subsection reference.

<a id="canonical-2313112033303123-2222322200313311-0213220113230212-3232021222213012-2300333102203321-3300201220003131-0203010220212023-2213101300102002"></a>

## Next pages — tls_parameters / 230232121332 / 4

- [proxy_config.https.tls_parameters.no_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-2230003003322321-2030203023030130-0311103033230111-2312121003013022-0122112230330223-3233101232232332-3200133323000330-0231021223022121)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200)
- [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1330121210122133-3031122320310320-0011233103332023-0123132311203120-2203332113021122-2323200001230323-3203300122100330-2003232220200300)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2230003003322321-2030203023030130-0311103033230111-2312121003013022-0122112230330223-3233101232232332-3200133323000330-0231021223022121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011311011022132-0131120001301220-0302231333213022-0223230113011021-2103220011300011-1132032311110331-3211021002023112-1032301021201130"></a>

## proxy_config.https.tls_parameters.no_mtls — no_mtls / 111121110110 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- proxy_config.https.tls_parameters.no_mtls

<a id="canonical-3310030003230002-1220330232102210-3310322322222322-3120220231131103-2013100022001031-0232220231320012-3222200121113102-2231011021001022"></a>

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
no_mtls = {}
```

<a id="canonical-2013122230231131-2111202213112031-0211111231230020-1333033030330221-2000210322300131-0303220302002033-3233103011120232-1313303033331021"></a>

## Direct properties — no_mtls / 111121110110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003000320311111-0223131321001002-2133222331103001-1303331031120033-2233002120301231-3120333302332020-2210332213231231-0220032103022332"></a>

## Next pages — no_mtls / 111121110110 / 4

- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213310222221312-2223232133310030-0212013001010220-1230322300322223-2333032121100131-2310112100011201-1203100100003211-3322223132031210"></a>

## proxy_config.https.tls_parameters.tls_certificates — tls_certificates / 002233103320 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- proxy_config.https.tls_parameters.tls_certificates

<a id="canonical-0220301032330333-1323013100331100-0012033233323112-2320332311233123-3212210220220222-0112021100023210-3201100302313332-0122102311013310"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
  "maxItems": 16,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022032132200233-1122220202011211-0002011021113113-2113200222003311-2131331021101020-0323300200113033-2023211232032321-2031000113020330"></a>

## Direct properties — tls_certificates / 002233103320 / 3

<a id="canonical-0032213100120331-0231021203102030-3023322201023331-1112010001322213-3132232010200100-0320222200322132-3211320213133323-0013323130333301"></a>

<a id="canonical-2303002103013011-1333020333010222-0121201021030220-3121122010000121-3313000013311030-2020120213303230-1003313120223302-2332030003132123"></a>

## certificate_url property — tls_certificates / 002233103320 / 4

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

- [custom_hash_algorithms](resources--bigip_http_proxy--reference--group-003.md#canonical-2003101000123301-3003321201003100-0032122001302231-1022322320210120-2213101123103231-2003102111322102-3223300103032332-2111300000213013): complete subsection reference.

<a id="canonical-3133310331011323-0312013011321130-3133332322303112-2232201032230200-0132020301023012-0013213132203202-2232223120002310-3221330123213032"></a>

<a id="canonical-1122113332000113-2301122310202220-2001003010002331-1000110320113032-2001130003213233-3300230021011023-0033333332211323-0303022023023221"></a>

## description_spec property — tls_certificates / 002233103320 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--bigip_http_proxy--reference--group-003.md#canonical-3110232021220003-3003233101313123-2203211130303311-2032120100312000-1320013021303302-0203201310100032-0221200113301310-1100013033233030): complete subsection reference.

- [private_key](resources--bigip_http_proxy--reference--group-003.md#canonical-1211012313130312-0102200123311121-2030202110320000-0311221021302313-1123022320102220-3230011321013201-3233032100022220-3102330232022302): complete subsection reference.

- [use_system_defaults](resources--bigip_http_proxy--reference--group-003.md#canonical-1032003223022220-2130323332313130-3221323000322303-3110310230102211-0210333013303123-0310301121230130-3112223130103113-0332222220001330): complete subsection reference.

<a id="canonical-1012102331220112-2203220200313220-3300301310022033-2201130332133011-0330132303311332-0121333112002122-3212002331100212-2122132303231021"></a>

## Next pages — tls_certificates / 002233103320 / 6

- [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms](resources--bigip_http_proxy--reference--group-003.md#canonical-2003101000123301-3003321201003100-0032122001302231-1022322320210120-2213101123103231-2003102111322102-3223300103032332-2111300000213013)
- [proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--bigip_http_proxy--reference--group-003.md#canonical-3110232021220003-3003233101313123-2203211130303311-2032120100312000-1320013021303302-0203201310100032-0221200113301310-1100013033233030)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--reference--group-003.md#canonical-1211012313130312-0102200123311121-2030202110320000-0311221021302313-1123022320102220-3230011321013201-3233032100022220-3102330232022302)
- [proxy_config.https.tls_parameters.tls_certificates.use_system_defaults](resources--bigip_http_proxy--reference--group-003.md#canonical-1032003223022220-2130323332313130-3221323000322303-3110310230102211-0210333013303123-0310301121230130-3112223130103113-0332222220001330)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2003101000123301-3003321201003100-0032122001302231-1022322320210120-2213101123103231-2003102111322102-3223300103032332-2111300000213013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213111322030333-0311003223220322-0023323212031200-0330303313223031-1003330130211311-2223200310012311-2100333022302300-2110121312211210"></a>

## proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 022110222221 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-2333232103133300-0103332332102100-3233300330020313-0131011021330123-1023112210330310-3321023123023210-2132112123212203-3313221233102030"></a>

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

<a id="canonical-0130331103310132-3033213330311000-0322231021321103-2232321013331230-1020310332131323-3013003121000203-1011110012201100-1013013302213020"></a>

## Direct properties — custom_hash_algorithms / 022110222221 / 3

<a id="canonical-1100222003230221-2033331123031030-3220021020133123-3012122121031213-3002223122312312-1011012231100202-2131011223020320-2331303222311231"></a>

<a id="canonical-3202110303111221-3302332112320013-1330032122233222-1312322330032221-3121313131013312-1320010331031022-0320113031222012-0200330120003231"></a>

## hash_algorithms property — custom_hash_algorithms / 022110222221 / 4

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

<a id="canonical-3301331323303112-1130200221113122-2010321120222313-2033023120221332-3020323133220133-2313120131222132-3210221000300301-3310200022221303"></a>

## Next pages — custom_hash_algorithms / 022110222221 / 5

- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-3110232021220003-3003233101313123-2203211130303311-2032120100312000-1320013021303302-0203201310100032-0221200113301310-1100013033233030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020222222222330-3130000311113333-2323022302021031-0203111212112130-2222030330312111-0212023012231002-3101133122000031-1322123312321020"></a>

## proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 102002033032 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-3233023212023113-3300203231011311-2232232201223103-2311313202200030-1313313102330213-3112331101121320-0330033210230300-0312131322102203"></a>

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

<a id="canonical-2210012103011120-1310232211300200-0220022313331131-1101310213313333-0011213010112031-3121221311120312-0101222110132230-3010012103022131"></a>

## Direct properties — disable_ocsp_stapling / 102002033032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203103223000032-0123132332231221-2002330220310101-0210032110030020-1030203020101230-2322131320201111-2123322131110311-1222313021233001"></a>

## Next pages — disable_ocsp_stapling / 102002033032 / 4

- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1211012313130312-0102200123311121-2030202110320000-0311221021302313-1123022320102220-3230011321013201-3233032100022220-3102330232022302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121330001322003-0111231120333322-0201130230322330-1201010321111110-3233333232110133-3000111011320023-3131200222230021-1032020000230323"></a>

## proxy_config.https.tls_parameters.tls_certificates.private_key — private_key / 330121220100 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- proxy_config.https.tls_parameters.tls_certificates.private_key

<a id="canonical-3013133203032233-3112011013132000-0201232133010111-3321300130121111-0211113323130333-0223111102210233-1321012113010101-2100001101012221"></a>

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

<a id="canonical-3331303311302112-2013130311133230-3033110001212121-1021230313232023-1301012112332112-0333011231101010-3122221203311111-0211000302013012"></a>

## Direct properties — private_key / 330121220100 / 3

- [blindfold_secret_info](resources--bigip_http_proxy--reference--group-003.md#canonical-0222132222003310-0221302202120120-1130032201222333-2022311310103320-1131322012101010-0120331202202020-2221321101320131-2022111211000000): complete subsection reference.

- [clear_secret_info](resources--bigip_http_proxy--reference--group-003.md#canonical-0003112113212000-3122200022201033-3113113103003111-1230203021201222-1323302110011213-0230320311200211-0301200022221103-0010221222311120): complete subsection reference.

<a id="canonical-3213132202112321-1222012032320021-3033032330212333-3130122213232131-1130313022023012-3002113031220230-2012113210013233-1232311332300323"></a>

## Next pages — private_key / 330121220100 / 4

- [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--bigip_http_proxy--reference--group-003.md#canonical-0222132222003310-0221302202120120-1130032201222333-2022311310103320-1131322012101010-0120331202202020-2221321101320131-2022111211000000)
- [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--bigip_http_proxy--reference--group-003.md#canonical-0003112113212000-3122200022201033-3113113103003111-1230203021201222-1323302110011213-0230320311200211-0301200022221103-0010221222311120)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0222132222003310-0221302202120120-1130032201222333-2022311310103320-1131322012101010-0120331202202020-2221321101320131-2022111211000000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212031320112230-1021300130001110-3020002030120211-3111232331030010-1033321112301312-2311311321003313-1223210031332312-2010332011132233"></a>

## proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 023213301010 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--reference--group-003.md#canonical-1211012313130312-0102200123311121-2030202110320000-0311221021302313-1123022320102220-3230011321013201-3233032100022220-3102330232022302)
- proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3002230331121022-2011220313320330-2221223233211031-3212133202023210-0003333232030203-1221312322100203-2110002312132312-1303232011130120"></a>

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

<a id="canonical-2030212313321100-3210130113313332-3220310113102003-2211132332030010-2332011313232223-1002213021201002-1312210123223200-1332323123120130"></a>

## Direct properties — blindfold_secret_info / 023213301010 / 3

<a id="canonical-0212203223200200-2113132031320133-2132003302301211-2032130301011120-3210003130222203-0312202003201033-3223002230220022-0232200033113102"></a>

<a id="canonical-2201202300213011-0222210030032311-1110130130031011-0123220222303320-0113312212013321-3331021210132010-2111222032001322-0200031312112323"></a>

## decryption_provider property — blindfold_secret_info / 023213301010 / 4

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

<a id="canonical-0121123033113033-1022122331121000-2300103130203112-0213310122131231-2231020003332220-3003231111321113-3232112332320231-0211321310022301"></a>

<a id="canonical-2212230032103100-0001102301301113-0321031210323310-1010133212021312-2032223322013013-1310022002003310-3300133303333112-0023000321130033"></a>

## location property — blindfold_secret_info / 023213301010 / 5

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

<a id="canonical-0211002203112010-1101011032112213-0113213212112323-0013311030333312-2233033312212201-0211312122001100-2302320332300003-2100233231323110"></a>

<a id="canonical-3000001320100101-0122210212121203-1223113201301332-2001132202201010-2223131022201313-3032220022231330-3122031213003130-2100212311022100"></a>

## store_provider property — blindfold_secret_info / 023213301010 / 6

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

<a id="canonical-0300103303301230-1203110021033231-3212200323330110-3001120320230200-1303312003001113-0230022111011011-3210312323320313-2103303220212022"></a>

## Next pages — blindfold_secret_info / 023213301010 / 7

- [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--reference--group-003.md#canonical-1211012313130312-0102200123311121-2030202110320000-0311221021302313-1123022320102220-3230011321013201-3233032100022220-3102330232022302)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-0003112113212000-3122200022201033-3113113103003111-1230203021201222-1323302110011213-0230320311200211-0301200022221103-0010221222311120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212020033303202-1031111322220303-3131132003230223-1303022000021300-0022233222123012-0002333231331230-2222101321020023-2312023222011230"></a>

## proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info — clear_secret_info / 033322031231 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--reference--group-003.md#canonical-1211012313130312-0102200123311121-2030202110320000-0311221021302313-1123022320102220-3230011321013201-3233032100022220-3102330232022302)
- proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-3101321102031322-2010221022033033-3110011233010102-1132312122322000-1121203302023131-3301130212231331-0313000010132021-1200121313012321"></a>

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

<a id="canonical-0003111033323303-0103322012303131-3220302112002231-3101300332320013-1201330302201031-3220001211333220-0031012101130132-3300202333010200"></a>

## Direct properties — clear_secret_info / 033322031231 / 3

<a id="canonical-0110203200231022-2023013001221300-2203303231020311-2032020312220302-1203232200100002-2030101223302013-3011021003132223-2201332211032312"></a>

<a id="canonical-2100332222033003-0033100033311332-3112220000303230-1203120002231022-0202133312203213-0221213210121022-2010322233300331-1212110011011200"></a>

## provider_ref property — clear_secret_info / 033322031231 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2023202121220333-3310223010223223-3323320130200213-2022103011011220-3022133301303232-2130320133302202-0031110032320232-3230323032122111"></a>

<a id="canonical-0120002003021113-3300013320101120-0233112133110011-1121302213333220-2001101100100331-3113032322133202-1100011210210122-3212023321302122"></a>

## URL property — clear_secret_info / 033322031231 / 5

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

<a id="canonical-0202211102212300-3310330221132311-2203030103033120-2121013102123222-0310000212100011-2002030331102030-3100231323002013-1002132110013311"></a>

## Next pages — clear_secret_info / 033322031231 / 6

- [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--reference--group-003.md#canonical-1211012313130312-0102200123311121-2030202110320000-0311221021302313-1123022320102220-3230011321013201-3233032100022220-3102330232022302)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-1032003223022220-2130323332313130-3221323000322303-3110310230102211-0210333013303123-0310301121230130-3112223130103113-0332222220001330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213210231012133-1003203210321031-2112130132002201-2330202003023311-3203321323302011-3012011233131221-0033320333322020-2210222110300203"></a>

## proxy_config.https.tls_parameters.tls_certificates.use_system_defaults — use_system_defaults / 121110323233 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013)
- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- proxy_config.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-3300203102333201-1311112211011302-0031000102110232-2312210220211233-1002211132111022-1232202302023033-2232132100310132-0111012201021031"></a>

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

<a id="canonical-2120200002221332-2113222032032110-1212220313320200-0312131123333203-2003312131130212-1002213322033310-2022020323200131-0321130003013333"></a>

## Direct properties — use_system_defaults / 121110323233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303113000213013-0203010313323302-1113320223300000-3011120011033132-0300311310022010-0121220200233100-0020221133330120-0210100213311322"></a>

## Next pages — use_system_defaults / 121110323233 / 4

- [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-0113320203210020-3311313102130022-1113310033003111-0101223223130021-2103033312111323-1320323200131202-1212122013221010-2321312000121101)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)

<a id="canonical-2031033101211112-0033321320211013-1322122101220213-0120102020101130-1231000130120333-3103223203002122-0100003022021331-2323132322020200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
