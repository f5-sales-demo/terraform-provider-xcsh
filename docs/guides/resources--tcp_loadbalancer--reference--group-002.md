---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-3120202130210103-3030312112113121-0333233002110322-3233011130210210-3321300030302100-0131303022121202-3011213011203112-3331333131312013"></a>

## name property — site / 232213120303 / 4

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

<a id="canonical-1232133211032003-1113323021002211-2332123230032210-2120232323022132-3323200311210220-3302102123233030-0030003333103222-1132020230133100"></a>

<a id="canonical-3211211023032100-0130020230023013-2032001113321223-2201233233323211-1233311203010310-3211001021320223-2112233101033210-3012222131003220"></a>

## namespace property — site / 232213120303 / 5

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

<a id="canonical-0000312230110203-3123331312133333-1022111132213330-2000001210002232-3120112202210003-0333003010100301-2012231031233212-3131310002000233"></a>

<a id="canonical-0332303320221000-2023220000202321-3201110211200000-3322203331013311-1003000201313211-1221021333133101-0320331013110331-1330320231223231"></a>

## tenant property — site / 232213120303 / 6

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

<a id="canonical-3221011112131000-0131310021332010-3321210330323222-3310002333300331-0130001103312111-1210101032110231-3110311012020322-3230333311311120"></a>

## Next pages — site / 232213120303 / 7

- [advertise_custom.advertise_where.vk8s_service](resources--tcp_loadbalancer--reference--group-001.md#canonical-2310032102311332-1320010320220300-3000131113132003-3031300330001303-3011312212131020-2233012333120212-1031220012331120-3102210330032022)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0102002101330311-3203311313223300-1030302122122311-2331332203030221-0201130111312302-1310112222312130-3133301001133013-2333130010233330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231002010331021-3010221222022102-2213010012131212-0102110310222213-2102123231100232-0102122230330022-3301321022003202-3221002013231012"></a>

## advertise_custom.advertise_where.vk8s_service.virtual_site — virtual_site / 020022021122 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- [advertise_custom.advertise_where.vk8s_service](resources--tcp_loadbalancer--reference--group-001.md#canonical-2310032102311332-1320010320220300-3000131113132003-3031300330001303-3011312212131020-2233012333120212-1031220012331120-3102210330032022)
- advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-2112123100031210-3120032221323121-2033302333321130-3323200332003223-0033120310200021-2310123321010101-0023120132332210-1312203221311012"></a>

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

<a id="canonical-3312131220112303-2123101313121013-3102210322313313-0133100320300203-3033221120132232-3010120311101011-3330032001303022-0113322222102222"></a>

## Direct properties — virtual_site / 020022021122 / 3

<a id="canonical-0111300312122122-1011232001322002-2333203111203012-0222021231213031-1312213121121211-1231012102130332-0020303021030331-2201020200031100"></a>

<a id="canonical-3000213222013312-2032103330333010-2011122332031011-0333012232112231-3210220231231303-3022213101210110-3011103221202232-3321021201122010"></a>

## name property — virtual_site / 020022021122 / 4

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

<a id="canonical-1112303020300010-1203231213133311-2033200023002001-3211030020102031-2333320332030320-3330201013133013-0003122023011321-3231201331220222"></a>

<a id="canonical-0301133022122320-2333103230321211-1122230331222300-1332031102230311-2000231123023222-1131322113223202-2103103032222312-0100110002001232"></a>

## namespace property — virtual_site / 020022021122 / 5

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

<a id="canonical-2200121221021111-1122220022210202-0222321303323220-1102231331110211-1301000020323311-3232300100011202-3301233232010120-2021130302312302"></a>

<a id="canonical-3013213332323202-3012023231231230-0100220213230103-3001011330011310-1110132323000122-1230220021023133-3303112200022200-3011111220120210"></a>

## tenant property — virtual_site / 020022021122 / 6

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

<a id="canonical-0233232232313003-2122300203211231-3231223220203123-3301011022211330-2212222101321232-1223121102301332-2321020021332303-2001033201103013"></a>

## Next pages — virtual_site / 020022021122 / 7

- [advertise_custom.advertise_where.vk8s_service](resources--tcp_loadbalancer--reference--group-001.md#canonical-2310032102311332-1320010320220300-3000131113132003-3031300330001303-3011312212131020-2233012333120212-1031220012331120-3102210330032022)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-1101101122210003-2330111331303010-2130211300213202-2120322121300012-2033321200023022-3103313203330222-2302000021033232-1221001021130221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210200100333300-3310022323203110-2211010032122103-3030312302112303-2120030331223232-0101131021112331-1322110011313002-0001000233311031"></a>

## advertise_on_public — advertise_on_public / 232233132300 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- advertise_on_public

<a id="canonical-0302313131313103-2320203031130213-3133301303203210-2220200131111230-0233011312332320-0230200102303112-2311223301221223-3032232213032211"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
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
advertise_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003202320333201-2103220322031312-2310133131200222-0011121313100232-1000223202333333-2113302330230320-2121020312013302-0313031003010122"></a>

## Direct properties — advertise_on_public / 232233132300 / 3

- [public_ip](resources--tcp_loadbalancer--reference--group-002.md#canonical-1202122111210021-3012203133200102-2103133031030132-1001321000010000-0321112231023232-3212123332122232-0230203020213311-2312120301320102): complete subsection reference.

<a id="canonical-0311332121210323-0310312020210002-3232320103122310-3230233230123322-2010022022210331-0101213121313112-1100201130130300-2231203210132022"></a>

## Next pages — advertise_on_public / 232233132300 / 4

- [advertise_on_public.public_ip](resources--tcp_loadbalancer--reference--group-002.md#canonical-1202122111210021-3012203133200102-2103133031030132-1001321000010000-0321112231023232-3212123332122232-0230203020213311-2312120301320102)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-1202122111210021-3012203133200102-2103133031030132-1001321000010000-0321112231023232-3212123332122232-0230203020213311-2312120301320102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203110313130210-2030112033131202-0001230233033003-2200133123010032-3120312103211100-3301121030123013-1100200123111012-0031212032112102"></a>

## advertise_on_public.public_ip — public_ip / 220110211100 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_on_public](resources--tcp_loadbalancer--reference--group-002.md#canonical-1101101122210003-2330111331303010-2130211300213202-2120322121300012-2033321200023022-3103313203330222-2302000021033232-1221001021130221)
- advertise_on_public.public_ip

<a id="canonical-2322221111012111-1133332103010101-0013232322003313-2002023331303131-3211221122003301-3301212102232332-3201101333132303-2211210301321233"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021232113311023-0310130301202313-1232330003030302-3221001312122221-3312301021010313-0202121200000231-1130002120113303-0313101000222101"></a>

## Direct properties — public_ip / 220110211100 / 3

<a id="canonical-0002200333201201-1121312313131333-3222011322012223-2011231033301000-0302211010211332-2232221113001132-1023211223120230-0300200300332131"></a>

<a id="canonical-2313110011132012-2322331012133112-0232020232023123-3102131111311332-3221310111332030-3211113222003310-2330202020121231-1012331331332323"></a>

## name property — public_ip / 220110211100 / 4

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

<a id="canonical-1312201323033122-3123210203001233-2320223320220233-2222201201301212-0230301203333032-3310212022000033-3000203121322123-2333333213231000"></a>

<a id="canonical-0200032002310000-3102002030323201-1330102102100121-0232113312031120-2032232303103232-3120032313110111-1200210320031101-1312033013131230"></a>

## namespace property — public_ip / 220110211100 / 5

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

<a id="canonical-1303111231321103-0130122311331131-0210132201310122-0310003201212112-1001120102120013-2222110330131000-2000310133201330-0131022101011333"></a>

<a id="canonical-1321211123023002-1211202020213323-2222122000032312-1212100031332211-0013032212022312-1301122003123203-3331220002303022-0322022212033102"></a>

## tenant property — public_ip / 220110211100 / 6

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

<a id="canonical-1330021002121200-0323131322230111-2320312131312213-2200031331210001-2211103201100131-1210133301002210-0023010030332120-3200231212303022"></a>

## Next pages — public_ip / 220110211100 / 7

- [advertise_on_public](resources--tcp_loadbalancer--reference--group-002.md#canonical-1101101122210003-2330111331303010-2130211300213202-2120322121300012-2033321200023022-3103313203330222-2302000021033232-1221001021130221)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3333100232210200-1000122300031311-3012302322223020-3032022132323133-3310012110333012-2333013330001120-1130222201132013-1120103002011310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033223331100033-3200210300200223-1301322012300023-1203212202201003-1223113032030312-0311231221301231-3211101300112030-2113201122100312"></a>

## advertise_on_public_default_vip — advertise_on_public_default_vip / 110300030331 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- advertise_on_public_default_vip

<a id="canonical-0321202223002312-2000122002021122-1000333311113001-0202202003021011-3030231120013213-3312123301110310-3323231213110121-0311213111231321"></a>

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
advertise_on_public_default_vip = {}
```

<a id="canonical-0301210033130030-3003112110131102-0032131133012001-0121303121103120-2131313233123010-3001312203312311-3313000221012123-3011212311001331"></a>

## Direct properties — advertise_on_public_default_vip / 110300030331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110323100323020-2232320001231022-0302132223113323-3231012110102130-3311132303022203-3101302311131013-0122303213013012-2222210313211123"></a>

## Next pages — advertise_on_public_default_vip / 110300030331 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-1233300113221313-3110203321303320-1132100212321211-3233111001220210-2202120330231303-1320322300203123-3132233222313123-1030123011100010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022302121312003-0001122320002132-1202112000201211-3031213200323312-1102010200113312-3202331333100000-0202001020202121-2221030322331230"></a>

## default_lb_with_sni — default_lb_with_sni / 133202213130 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- default_lb_with_sni

<a id="canonical-0313202033230201-1300210123230312-1001011322302321-0120211103330203-3120232100300130-1120303231202201-1102310202111313-2312111111000220"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_lb\_with\_sni, no\_sni, sni; Default: default\_lb\_with\_sni\] Configuration
parameter for default lb with sni.

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

- [default_lb_with_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-0313202033230201-1300210123230312-1001011322302321-0120211103330203-3120232100300130-1120303231202201-1102310202111313-2312111111000220)
- [no_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-3013302200020113-0111322303322230-0000110230220313-0232032310130333-2323102130310131-1010123111330300-2220231131123000-2211320231132021)
- [sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-2121331132131131-3033031100120302-0330132302022103-3000220102112000-3011311311311201-2200311121231223-3011221000201220-3201310112120322)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_lb_with_sni = {}
```

<a id="canonical-1113030301220122-1202111231201222-1312023312122312-0320322322202011-0031103323133201-2232310232130232-1123222133322303-2010111101111001"></a>

## Direct properties — default_lb_with_sni / 133202213130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122221332100311-1330121301122232-3120100201200221-3113233112131030-2321030333201202-1223011202100213-1232121100133132-1300010100123023"></a>

## Next pages — default_lb_with_sni / 133202213130 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2322213221112110-0231131200021100-2323031112011321-2113223331230331-1110202310201022-0323112230111011-3203121033311100-1130120122122120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003120001013300-3202113030003321-3223231333210230-1220023322210313-3303120101002132-1021000112221021-0320132010233123-3132211132103330"></a>

## do_not_advertise — do_not_advertise / 220130201301 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- do_not_advertise

<a id="canonical-3132133212102201-2021300201322133-0333101230320001-0201211013310101-2132321123323201-1233011231122201-0032202131001100-1033022212130331"></a>

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

<a id="canonical-3200323120310320-2331232233222331-2233103001111320-0212032131022023-0333302001231323-1303101003110211-1110323113021233-0001312331032111"></a>

## Direct properties — do_not_advertise / 220130201301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011203311001233-0311321022313233-3212110303112011-3101220202120020-1132331001133001-1200201033321000-0102111233030112-1333102032221110"></a>

## Next pages — do_not_advertise / 220130201301 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2322113213320330-0102033300000032-1130020131303320-1130033301222310-0100131123323030-1300033212333030-2033201321130131-2130003330320013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203033123033101-0202210133323210-1222212111110001-0312321213120211-0001221222310310-0203121222000323-1321323121333321-1333120002021031"></a>

## do_not_retract_cluster — do_not_retract_cluster / 121201023133 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- do_not_retract_cluster

<a id="canonical-2132121202202312-3303103110203131-0123121231302102-1022300131113230-2101310330031313-1031133323302233-2310000223033320-3130011202233233"></a>

Type: `["object", {}]`. Optional.

\[OneOf: do\_not\_retract\_cluster, retract\_cluster\] Enable this option

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

- [do_not_retract_cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-2132121202202312-3303103110203131-0123121231302102-1022300131113230-2101310330031313-1031133323302233-2310000223033320-3130011202233233)
- [retract_cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-0223212221221323-1003333011230323-2311103300002012-3022112031333333-0211120323003211-0323113322301032-0013033220133100-0321130210120033)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
do_not_retract_cluster = {}
```

<a id="canonical-0000312310322021-2330221233130121-1123011132030032-2213320100331122-3012323320332301-0210002100010010-2231002000223211-1302302211022103"></a>

## Direct properties — do_not_retract_cluster / 121201023133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312112103031111-2020212110221303-3112023032012213-1100113200332121-1210032311311013-2321230031312023-2220131313020111-2221211102002301"></a>

## Next pages — do_not_retract_cluster / 121201023133 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3231312101111010-1230322232102131-1310111121030233-1233021020330300-2012003120131022-3230203210022022-1331310221131200-1310111032002330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332233112033321-3102100320211223-2322203011210131-3130222121110212-1013120102130310-2313332010323123-1303312111130120-1220022003031310"></a>

## hash_policy_choice_least_active — hash_policy_choice_least_active / 011230110333 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- hash_policy_choice_least_active

<a id="canonical-2120022010132212-3200121332222303-3312131310001132-1220100002110001-3013320221312220-0031002113013130-1223201110102021-3102010323112103"></a>

Type: `["object", {}]`. Optional.

\[OneOf: hash\_policy\_choice\_least\_active, hash\_policy\_choice\_random,
hash\_policy\_choice\_round\_robin, hash\_policy\_choice\_source\_ip\_stickiness\] Enable this
option

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

- [hash_policy_choice_least_active](resources--tcp_loadbalancer--reference--group-002.md#canonical-2120022010132212-3200121332222303-3312131310001132-1220100002110001-3013320221312220-0031002113013130-1223201110102021-3102010323112103)
- [hash_policy_choice_random](resources--tcp_loadbalancer--reference--group-002.md#canonical-0320110201222002-0210223103211030-3332203031121102-3103000311221313-2002102321221001-2330032210323323-3301102223121201-1031013122301231)
- [hash_policy_choice_round_robin](resources--tcp_loadbalancer--reference--group-002.md#canonical-3322321013232230-1300300103220333-2220321012300001-3102030202321111-2033300233313210-0220020121331312-2032311221020001-1221330012002212)
- [hash_policy_choice_source_ip_stickiness](resources--tcp_loadbalancer--reference--group-002.md#canonical-3101023012331020-0102212011231001-1131323230001301-1323130122000230-3202123120312213-3200111331232121-1032310313320003-1122101121103222)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
hash_policy_choice_least_active = {}
```

<a id="canonical-3001032123301022-1330313020132121-2231120332131112-1223323311303013-0032112002113021-0323310122011030-0322300320101102-2120010210033023"></a>

## Direct properties — hash_policy_choice_least_active / 011230110333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203220303310133-3330333131100133-3313113123322333-1103110001212210-1113202212110111-2002010112000130-1012233033233233-0110232222101113"></a>

## Next pages — hash_policy_choice_least_active / 011230110333 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2133201212213303-3220313201223303-0031130110231110-1213131231101232-3321233020122003-3302010033002303-0202202110320320-1222200211330032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212323303200222-2213223210311132-2310322212120211-0312231130323103-0201112122200230-2231220303012100-2101101232333130-0123101320011310"></a>

## hash_policy_choice_random — hash_policy_choice_random / 212011130303 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- hash_policy_choice_random

<a id="canonical-0320110201222002-0210223103211030-3332203031121102-3103000311221313-2002102321221001-2330032210323323-3301102223121201-1031013122301231"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for hash policy choice random.

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
hash_policy_choice_random = {}
```

<a id="canonical-3322100230311030-3233112033002231-0203220332112021-0333232013120211-1132212101102332-1210011011333110-0213230130110300-1320112303213123"></a>

## Direct properties — hash_policy_choice_random / 212011130303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312203031112311-1330022311102223-1322321001212120-3010130330301303-2112201223020232-3330131220031200-2023101322311333-1102213333320202"></a>

## Next pages — hash_policy_choice_random / 212011130303 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2011303133011102-2021121122120220-1133330230131122-0002121330231010-2103102030013022-2020210201310012-0000100210012131-3213113301223032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031113311322133-3300200301300220-2233333230023111-3210221103112222-0121122121003113-3031313012100333-0030313301133131-3210233113101300"></a>

## hash_policy_choice_round_robin — hash_policy_choice_round_robin / 022331103231 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- hash_policy_choice_round_robin

<a id="canonical-3322321013232230-1300300103220333-2220321012300001-3102030202321111-2033300233313210-0220020121331312-2032311221020001-1221330012002212"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for hash policy choice round robin. Defaults to \`map\[\]\`. Server applies
default when omitted.

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
hash_policy_choice_round_robin = {}
```

<a id="canonical-1003311113331012-0013202100320021-1322102220312021-1013132312011112-1102312111233303-0332311312200202-0130301332331010-0322203033212333"></a>

## Direct properties — hash_policy_choice_round_robin / 022331103231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122020232220033-2220012021322333-3312232332103331-2203221102120021-0320131213300113-0232120111321213-1121021331011221-2323212112033022"></a>

## Next pages — hash_policy_choice_round_robin / 022331103231 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3123323133211320-2223021102110033-2023122331312031-1312102001331031-2023122213321233-0303012321200222-1233301221213131-3010122320223012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330323233313123-0201213202322021-3203301220231003-3213123021120313-1002311111221223-0212113211211232-2113121330312030-1202101131220001"></a>

## hash_policy_choice_source_ip_stickiness — hash_policy_choice_source_ip_stickiness / 331001231101 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- hash_policy_choice_source_ip_stickiness

<a id="canonical-3101023012331020-0102212011231001-1131323230001301-1323130122000230-3202123120312213-3200111331232121-1032310313320003-1122101121103222"></a>

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
hash_policy_choice_source_ip_stickiness = {}
```

<a id="canonical-2012033031332330-3031211112032220-3311031103222000-1323232110001033-3022310102121233-3213212030110201-1032101202001023-1222221003001221"></a>

## Direct properties — hash_policy_choice_source_ip_stickiness / 331001231101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312012311032321-2011100000310012-2330232221220202-1011223022320220-3110221013221100-3313021033023020-2112222102320110-3203221202120002"></a>

## Next pages — hash_policy_choice_source_ip_stickiness / 331001231101 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0211303031122133-3002101102330112-3303102231112031-0323212233100110-3023331222023022-0122331220121021-0201223330103012-0220122230220001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332233030232203-2003320302302310-1203202210123331-0311232332010223-1030022032022133-1033130220333203-1220200231110233-3302221000220213"></a>

## no_service_policies — no_service_policies / 302003112223 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- no_service_policies

<a id="canonical-0102302021033021-2301111022111023-2331021210033212-0131133021322033-3000003112132020-3310331033321310-3001103300120102-2322303233312111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no service policies.

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
no_service_policies = {}
```

<a id="canonical-0310022122132330-1323330012032222-2221300230132021-0130320323011212-1231101303230011-3000231223303203-2233212232021222-1231223030212102"></a>

## Direct properties — no_service_policies / 302003112223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100103003320001-0022130033330302-0132302200321132-0212210200312030-3203203322212010-2230120101313203-3223101302301320-3233101102301300"></a>

## Next pages — no_service_policies / 302003112223 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0332121320022012-0131110131002221-3211220033203220-3333221011312320-1303130303113032-2011002222333033-3001212201031002-1010101011201223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320120223112330-0032011111200313-0020003100202231-0112210213021010-2033233301322231-1330022220131012-3103000021212031-1302231323331230"></a>

## no_sni — no_sni / 133213032030 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- no_sni

<a id="canonical-3013302200020113-0111322303322230-0000110230220313-0232032310130333-2323102130310131-1010123111330300-2220231131123000-2211320231132021"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
no_sni = {}
```

<a id="canonical-1131313110203110-2201030330333221-1330102032301222-2112100320111122-1212300323231021-0331023212322200-3311120300001000-2133133211101331"></a>

## Direct properties — no_sni / 133213032030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200202302130230-1301312213331120-1112021032121023-0111013120000020-2123332311133312-2302122020132233-3222032012122123-0002301202010120"></a>

## Next pages — no_sni / 133213032030 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2100010003132033-1331210120313310-0212003223033302-1320122132000102-1302201010312200-3333221331300211-3332023322113011-0102012310233122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002000111000330-3030223113203102-2133333000321010-2311210223332131-2312210010332123-0233333023312203-2333203132220302-3233233330201000"></a>

## origin_pools_weights — origin_pools_weights / 231122201031 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- origin_pools_weights

<a id="canonical-3300102231110023-1122311101122321-3133123320300330-3102222002313323-1001120223303033-3223101033122101-0032121010131101-3233033211031232"></a>

Type: `"object"`. list nested block, Optional.

Origin pools and weights used for this load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cluster",
    "pool")}
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
origin_pools_weights {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223123022011200-0203003211033201-2023321311303332-2113330312211030-2011333111203033-0102012121202203-1022123100111120-2312023023032022"></a>

## Direct properties — origin_pools_weights / 231122201031 / 3

- [cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-1011120330133113-3221310102110212-3012110211103032-1212313330022332-2303330232010101-3212113330300201-2331100110002001-3330111113013333): complete subsection reference.

- [endpoint_subsets](resources--tcp_loadbalancer--reference--group-002.md#canonical-0233132302001113-1013122023210222-1222022303323200-3120101033031121-0232202032202111-0031100302133121-1303233330230201-1301123310033123): complete subsection reference.

- [pool](resources--tcp_loadbalancer--reference--group-002.md#canonical-0101012032110312-0012103013221021-2000123112202033-2211320112310201-0213110000103010-1321023222113211-3031331212323121-2320012220221031): complete subsection reference.

<a id="canonical-2023023300223021-1030013322302200-0323020103311210-3121211203123312-0132122310301111-2310333003000230-2320310032131322-3231122110302300"></a>

<a id="canonical-3001133123212230-3323222011321301-1001120202113303-3013211001201201-0020221203213232-1333331210323303-1301023012203011-0311110123302130"></a>

## priority property — origin_pools_weights / 231122201031 / 4

Type: `"number"`. Optional.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-1300023231131310-0000310231101120-1111103312212312-1102222100102213-1113123311303013-1121213322000130-2330000132310232-1100321300303130"></a>

<a id="canonical-2303103033330302-3003133323013032-3203000200001232-3011333131030200-1330122201222302-1311220133020120-3231311302230132-1110302221022132"></a>

## weight property — origin_pools_weights / 231122201031 / 5

Type: `"number"`. Optional.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2123031331123131-2013311303020131-0033220103101312-3203222122001222-1012232221030020-3002112001022220-1201301030100302-2333112321010122"></a>

## Next pages — origin_pools_weights / 231122201031 / 6

- [origin_pools_weights.cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-1011120330133113-3221310102110212-3012110211103032-1212313330022332-2303330232010101-3212113330300201-2331100110002001-3330111113013333)
- [origin_pools_weights.endpoint_subsets](resources--tcp_loadbalancer--reference--group-002.md#canonical-0233132302001113-1013122023210222-1222022303323200-3120101033031121-0232202032202111-0031100302133121-1303233330230201-1301123310033123)
- [origin_pools_weights.pool](resources--tcp_loadbalancer--reference--group-002.md#canonical-0101012032110312-0012103013221021-2000123112202033-2211320112310201-0213110000103010-1321023222113211-3031331212323121-2320012220221031)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-1011120330133113-3221310102110212-3012110211103032-1212313330022332-2303330232010101-3212113330300201-2331100110002001-3330111113013333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031120003132011-3032221320110312-3302213120202210-3011303211220322-3011220100013113-3122032212223122-1101231030031202-0223201101201111"></a>

## origin_pools_weights.cluster — cluster / 113121202110 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-2100010003132033-1331210120313310-0212003223033302-1320122132000102-1302201010312200-3333221331300211-3332023322113011-0102012310233122)
- origin_pools_weights.cluster

<a id="canonical-2110022331031320-2323210200030022-2201303121210111-2302002332211320-1202301111330323-3122312123013131-3222123001321121-1202200010031330"></a>

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
cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332013031122031-3032231223210321-3000003020130003-2303122133310231-1021222212031211-2330001321033000-2331001130110200-2201132230030201"></a>

## Direct properties — cluster / 113121202110 / 3

<a id="canonical-2221201022222220-1101232211132331-1211322311311012-0231322230031012-3120232331301301-2201122100233320-1022002012301012-1202111222101233"></a>

<a id="canonical-0231202100120323-1232303110333022-0221002000220121-2201231233333321-2132013333223232-2313333002323231-3000103201212012-1120100320312122"></a>

## name property — cluster / 113121202110 / 4

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

<a id="canonical-1312321133020223-2023020303320302-1113321133221211-3310033222323121-2220011331211020-1223021233121022-2002130332202303-1101110012302002"></a>

<a id="canonical-1300313032223300-0123331020013000-0120110313010100-0132030331223232-2003100100223000-3210030323123310-0311212303311012-0330122013011301"></a>

## namespace property — cluster / 113121202110 / 5

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

<a id="canonical-0303130010011030-3130300323203021-1212130313322132-2232300322020230-3330311013001331-3111111123120112-3333021020320230-0311201223201221"></a>

<a id="canonical-1011013232010110-2121101230132003-1220233221220031-0102020033023031-0212221323120032-1020131211121200-1221112123130012-1212113121233302"></a>

## tenant property — cluster / 113121202110 / 6

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

<a id="canonical-2230031231021020-3032231020123123-0322010210003312-3122300233010220-1213221003201102-0330233100132213-1013000023320000-0103212012323123"></a>

## Next pages — cluster / 113121202110 / 7

- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-2100010003132033-1331210120313310-0212003223033302-1320122132000102-1302201010312200-3333221331300211-3332023322113011-0102012310233122)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0233132302001113-1013122023210222-1222022303323200-3120101033031121-0232202032202111-0031100302133121-1303233330230201-1301123310033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300323203133300-0020330003232111-0200300120120130-3233023023113201-2323203021030300-2300021112021012-1321132203211113-1021131001033332"></a>

## origin_pools_weights.endpoint_subsets — endpoint_subsets / 013320302330 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-2100010003132033-1331210120313310-0212003223033302-1320122132000102-1302201010312200-3333221331300211-3332023322113011-0102012310233122)
- origin_pools_weights.endpoint_subsets

<a id="canonical-3120233110133312-0321210100022022-3000321113133310-3130101101021221-1223301332330133-1103223110230230-2230020003011021-0233130133022230"></a>

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

<a id="canonical-0100131100323033-1021302002003101-3013222231010203-3213011321230200-1320220221123110-0013230310022230-1003113103222231-2123001030311212"></a>

## Direct properties — endpoint_subsets / 013320302330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003301110300201-2131330110311320-2023030231322322-3002201121000130-0021203331233233-3311120320120221-0210312333320211-0120002202001201"></a>

## Next pages — endpoint_subsets / 013320302330 / 4

- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-2100010003132033-1331210120313310-0212003223033302-1320122132000102-1302201010312200-3333221331300211-3332023322113011-0102012310233122)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0101012032110312-0012103013221021-2000123112202033-2211320112310201-0213110000103010-1321023222113211-3031331212323121-2320012220221031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021131332322131-2221222301001023-3000301322233223-2133132013103211-0232232001232323-1330021113312320-0022313033230321-1100221302020202"></a>

## origin_pools_weights.pool — pool / 222213123233 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-2100010003132033-1331210120313310-0212003223033302-1320122132000102-1302201010312200-3333221331300211-3332023322113011-0102012310233122)
- origin_pools_weights.pool

<a id="canonical-2211110010101203-3132220223100012-1211011020310013-3133233312232003-2030302002310022-0201203300032312-2202322100111100-2111033302210000"></a>

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
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110021320201212-0212002011011230-2232322313032331-1010130330011031-3122222020202222-1301211001002100-2210033231001022-0330223132321031"></a>

## Direct properties — pool / 222213123233 / 3

<a id="canonical-3231133211120113-0210010220020230-3032001333300132-3330221200210310-2003211213221020-2022321330203001-3201322003333331-2230210321333301"></a>

<a id="canonical-1011333222013332-1221111032103123-0312012231121111-1013130113201223-1320112333000122-1030003001321332-1101200230300130-3011002110200003"></a>

## name property — pool / 222213123233 / 4

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

<a id="canonical-1311301302321321-1103201112111120-2303010031013133-3031313320321301-2213220113003101-2320121311011330-2210330332320103-3323320323200323"></a>

<a id="canonical-1313321322111201-0302103000231130-1033312011220020-0123200033121301-1130303221132202-3330303320101322-0330121223013111-0103201202222123"></a>

## namespace property — pool / 222213123233 / 5

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

<a id="canonical-0210031323123022-3033220210002201-2122322311203000-2131001200333233-3020321333003202-2100011023013333-1312210001212013-2132103311100202"></a>

<a id="canonical-3210030021322110-3212223120330202-3111201111000002-3123113002301100-0133033233110002-1023130212301013-3133103202313300-3201223322030231"></a>

## tenant property — pool / 222213123233 / 6

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

<a id="canonical-3333201311103002-0102323013020033-2222232110023013-2212303112202221-1101003111310332-3132322003312313-1311000103232331-1130202010230113"></a>

## Next pages — pool / 222213123233 / 7

- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-2100010003132033-1331210120313310-0212003223033302-1320122132000102-1302201010312200-3333221331300211-3332023322113011-0102012310233122)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2203330232231221-0231100222101101-2032200123103102-3022031012302100-0322123221223231-2330022031223301-3202022122302011-3102020113332333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102002132223311-2313320223001130-2110302211300213-0311222102312223-3121032113323033-2131211021310210-0313002020232323-2333212010230020"></a>

## retract_cluster — retract_cluster / 220133312300 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- retract_cluster

<a id="canonical-0223212221221323-1003333011230323-2311103300002012-3022112031333333-0211120323003211-0323113322301032-0013033220133100-0321130210120033"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
retract_cluster = {}
```

<a id="canonical-3021233321223331-0023011100020133-1121220013301010-0030321003021331-1013123321011113-3221120320123012-0323102222031232-3133111320013331"></a>

## Direct properties — retract_cluster / 220133312300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000121332311313-1103210211333223-1020032110023000-1000131001101322-1133102123322010-0030213102221331-2210021330223211-3102013200113302"></a>

## Next pages — retract_cluster / 220133312300 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-1302000233122001-0330320220303103-0201012112213303-1133102211110130-2333133012210201-3013203312103310-2212130213030102-1122113100303223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033212000003222-1111023013223213-0213002121201300-0023121123223031-1112211003031111-2121120111121302-0002333322303222-3312130111332320"></a>

## service_policies_from_namespace — service_policies_from_namespace / 030131302200 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- service_policies_from_namespace

<a id="canonical-3001223103333333-1020031000231122-0311203031231103-3133312332031032-1302211202001202-3211013112120332-1003020031013122-3233111310220230"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
service_policies_from_namespace = {}
```

<a id="canonical-3213130303112112-0120110323310221-3100301132022021-3030222133001302-1103222223212220-0001202021032013-3102220113003033-0121310110021221"></a>

## Direct properties — service_policies_from_namespace / 030131302200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020213132221233-0022021121003010-1212102230030021-2121122031011032-2222112032130223-0001323310132303-0312013100103130-0102220000112132"></a>

## Next pages — service_policies_from_namespace / 030131302200 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-1211312003133330-1223211300322220-2113300203322232-3010131033013022-2332011121303131-3303223200010132-3300002312210210-2000310102320322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012102020332203-3030102320000112-2113213013313232-2011302231132300-0023030320131301-0322302011212310-3312302210203123-3221301322220221"></a>

## sni — sni / 003313221221 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- sni

<a id="canonical-2121331132131131-3033031100120302-0330132302022103-3000220102112000-3011311311311201-2200311121231223-3011221000201220-3201310112120322"></a>

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
sni = {}
```

<a id="canonical-3100012121310032-2210121221112320-3010330023000120-2100201101100113-3122332011001301-1033230131310100-2033213211210121-2131300001222202"></a>

## Direct properties — sni / 003313221221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112101100101301-3033313123302320-1101110133220121-3302312103023023-0203121320210302-2333000132102000-2131301011022130-0033320310230103"></a>

## Next pages — sni / 003313221221 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2220232031020123-0333221323000013-0330311311100200-2313132103001132-1300312202210020-1103313303133002-0011300213213202-1132012101023211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130012313022220-0012110220113121-3221312213233033-3133221121102221-2211100202313113-2123300130230011-2330200031331313-3123002312232300"></a>

## tcp — tcp / 303330303322 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- tcp

<a id="canonical-3320220303000012-1212310221333003-2023232000001311-2003113122201032-0320213100330200-3200012310233331-1203312023032303-2033123000332201"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: tcp, tls\_tcp, tls\_tcp\_auto\_cert\] Enable this option. Defaults to \`map\[\]\`. Server
applies default when omitted.

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

- [tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3320220303000012-1212310221333003-2023232000001311-2003113122201032-0320213100330200-3200012310233331-1203312023032303-2033123000332201)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3311310101231101-0312312311123203-2112200103233023-0231013101001013-1232132002021323-0011020322221211-1311301030203010-2212103030203323)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1021002212100302-1311013322023113-0033102003303321-1102223200321033-3111102121313232-0020131233022011-3121111200002222-0320203122133203)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
tcp = {}
```

<a id="canonical-0023232113303131-0213322303212010-1020012212133212-1113113311211302-0122213112210223-3033233013033130-2333020130120332-0033213212023221"></a>

## Direct properties — tcp / 303330303322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112210113030201-1230111032101223-3330111022121002-1022121322120023-3313212232200121-1333101032013201-2103120221300222-3023120222120000"></a>

## Next pages — tcp / 303330303322 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-1223011333230122-2303021022201212-1030021321212322-1311000300020110-1130221333112233-3223223303001030-2231022021122211-0113200021221011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202331302322332-2231203121223220-2130301102301230-3012102112131032-0010201020122312-2133221103330322-2302313100102101-0232132211231012"></a>

## timeouts — timeouts / 231130022312 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- timeouts

<a id="canonical-0231313301332031-0221012333210323-2202220031030022-3300010221320220-0102332011321200-0321222233312130-1133231323310012-1333333313112222"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322101011313330-0322103332331101-3133303203032311-3313221031321102-3313303003011100-1101220003130313-1133222103133231-2133131120322312"></a>

## Direct properties — timeouts / 231130022312 / 3

<a id="canonical-1223233012111011-0320010120330323-2110100223101002-0021021131120322-3212003211111103-0020232031013220-3302212121020312-3332323131102132"></a>

<a id="canonical-3000221021220213-3002231100321313-3332331001031012-1010320322112133-1333011303321331-3302301122011001-2302012223330002-0333310303320121"></a>

## create property — timeouts / 231130022312 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2301320011230031-3113330321132003-2132312201210201-3031201321333030-0312012330232203-0201010223302230-2233330222312123-2132112230212103"></a>

<a id="canonical-0112010013020330-2023111200021333-1010231002233201-1310301232330101-0112320111201112-3211112022201002-3301001311121310-2030302120222333"></a>

## delete property — timeouts / 231130022312 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2312030213132213-0021202101111112-2102221101011331-0332303033022122-0303123120222001-1001221110032001-0211033113032320-1010122102000231"></a>

<a id="canonical-3312001300031322-2131233220122020-1330011021202132-3133203320011113-2230213231220203-3001332132122233-0300311100122201-2322120022012033"></a>

## read property — timeouts / 231130022312 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3323331302032012-3330110330212101-2303232133322130-0223003322013132-2302100202101301-2203312321331313-2011112230130313-1310032331303010"></a>

<a id="canonical-2203312032313333-1222212011112112-2232113310112201-0211003130102303-0113021000112212-1033002210130113-3332001122312133-3111213220002301"></a>

## update property — timeouts / 231130022312 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0221111212113033-2131100123002102-3121312203023013-3030000232323012-0330202233012022-2310311332223303-0330123203222132-3303123221113113"></a>

## Next pages — timeouts / 231130022312 / 8

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001303022321131-3300200220200203-3300320202311011-1211021011330123-0330130000301033-3101200223121100-2223231111132210-0001102221033023"></a>

## tls_tcp — tls_tcp / 013223112320 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- tls_tcp

<a id="canonical-3311310101231101-0312312311123203-2112200103233023-0231013101001013-1232132002021323-0011020322221211-1311301030203010-2212103030203323"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting TLS over TCP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("tls_cert_params",
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
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
tls_tcp {
  # Configure direct properties listed below.
}
```

<a id="canonical-0233311333313033-2133233031200203-3010223220112020-0002311130111312-3103330003133302-2132000223110030-2313011212123021-0020221200013123"></a>

## Direct properties — tls_tcp / 013223112320 / 3

- [tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301): complete subsection reference.

- [tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132): complete subsection reference.

<a id="canonical-0201213213301212-2211102120113202-0010003022100231-3202302301320321-2230132100333131-3010112013321332-0322132002331032-2220003333321021"></a>

## Next pages — tls_tcp / 013223112320 / 4

- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103133200032023-2301221133211102-2322230221131311-0122110202230130-0000102133012021-1312330001211121-0000010230131202-1331231001133123"></a>

## tls_tcp.tls_cert_params — tls_cert_params / 031100112200 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- tls_tcp.tls_cert_params

<a id="canonical-2003212231231230-2122303233200032-0123102210223221-2301333103012021-0022212032211012-2002001201000110-3021200222310310-2223313022132223"></a>

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

<a id="canonical-2311233332221210-1032030030123100-2223213220133131-3133130113302233-1301213020123331-1211333333211312-0312132011232133-2003011210111202"></a>

## Direct properties — tls_cert_params / 031100112200 / 3

- [certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-3213110003102033-1303112111202003-2023331320110320-3133112022030112-2203320300332012-0231011200113210-0123111101030111-0201203323200003): complete subsection reference.

- [no_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-0011323331201223-2233033222323031-0123021332131023-2101130020120033-1322010103013022-0000022023200202-2333212121323232-1031132032111203): complete subsection reference.

- [tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130): complete subsection reference.

- [use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313): complete subsection reference.

<a id="canonical-0102312330110203-0300000133103321-0332011003302320-2230310200120132-1332300001310232-3311223212113101-2310112011232320-2101133010302300"></a>

## Next pages — tls_cert_params / 031100112200 / 4

- [tls_tcp.tls_cert_params.certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-3213110003102033-1303112111202003-2023331320110320-3133112022030112-2203320300332012-0231011200113210-0123111101030111-0201203323200003)
- [tls_tcp.tls_cert_params.no_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-0011323331201223-2233033222323031-0123021332131023-2101130020120033-1322010103013022-0000022023200202-2333212121323232-1031132032111203)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3213110003102033-1303112111202003-2023331320110320-3133112022030112-2203320300332012-0231011200113210-0123111101030111-0201203323200003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331330213020103-2100111223300331-1333023221332313-2213021030122121-1321003110112322-0221001131221223-0323030301021031-2201023212232213"></a>

## tls_tcp.tls_cert_params.certificates — certificates / 032313310001 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- tls_tcp.tls_cert_params.certificates

<a id="canonical-1023331330101033-1020020103122030-2123332131230111-2230101121013212-1023301223011112-2011303033120331-2021112200211310-0113223113132333"></a>

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

<a id="canonical-2101032030213231-2122012220211001-1311132131211322-0220120302133033-0230122002212103-0333321123332101-3032003232002330-1313301332103330"></a>

## Direct properties — certificates / 032313310001 / 3

<a id="canonical-0012113110101030-3122010132010113-1313312023003013-1023002333303203-2100120033120011-1230300102221222-0021013032200010-2031321030102233"></a>

<a id="canonical-0102122131133230-1332003033332212-1112213122312213-2223210312003233-3332320003100323-3012223232103311-3322132321300122-2323210101111303"></a>

## name property — certificates / 032313310001 / 4

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

<a id="canonical-3201012123201210-1112131301001222-1133201300123010-0213203332033303-3303030130330123-0320301212130210-1233331300132103-1313122220201313"></a>

<a id="canonical-0222133313310220-0222313201332032-2100133232222232-2131013012332222-0230322030113020-2003211320330323-1032321131301020-0312002003231113"></a>

## namespace property — certificates / 032313310001 / 5

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

<a id="canonical-3221111031212330-0122002102310230-2301001123303210-3333231002223120-3010123300203330-2320102201202030-2200331320021202-0331020202113010"></a>

<a id="canonical-2231113132133101-0103331321313011-0220110013202100-3213032212022310-0033321103333100-1010223020223102-3301310321132101-0131331020212211"></a>

## tenant property — certificates / 032313310001 / 6

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

<a id="canonical-0123011032030023-3303121101131301-3301133322223301-1220310312311103-2233221031312223-2123203202220310-2033133010122110-2033222233322002"></a>

## Next pages — certificates / 032313310001 / 7

- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0011323331201223-2233033222323031-0123021332131023-2101130020120033-1322010103013022-0000022023200202-2333212121323232-1031132032111203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302300101020301-1220311002033223-0221130231331113-0030033122320302-2020133131312231-2323031113133011-2301102330033012-0222300311232310"></a>

## tls_tcp.tls_cert_params.no_mtls — no_mtls / 221210133212 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- tls_tcp.tls_cert_params.no_mtls

<a id="canonical-0231032331233303-3221202323213112-2122302002313202-0322312221003231-0230200212111233-3222231300120331-0211303122102312-2001133332112332"></a>

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

<a id="canonical-3012232123313203-3313332303031013-2310022111211231-1331320312120130-2001321012231210-0203112131231220-2302030010230022-0120013333223112"></a>

## Direct properties — no_mtls / 221210133212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122220331100230-0132300103313303-1121320030130103-3022133010220020-2211300330300330-0130212120020230-1013031013331212-3302030111233033"></a>

## Next pages — no_mtls / 221210133212 / 4

- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113220113210132-0302310001112120-1300211000103121-0032020030313201-3131231212013112-2033232122111022-3312310021321330-0103113111131013"></a>

## tls_tcp.tls_cert_params.tls_config — tls_config / 211010332120 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- tls_tcp.tls_cert_params.tls_config

<a id="canonical-0300232220332032-1221030130223021-3111102102101230-3303020222132332-0221110032000221-1320031310322313-2001122212000022-3332230202330021"></a>

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

<a id="canonical-1100122230031111-0200233220321333-0221331020031032-0030112010023010-2020332120110203-1013011212133300-2313200232323203-0102111020233220"></a>

## Direct properties — tls_config / 211010332120 / 3

- [custom_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-3320220203210203-3010313213202302-0013120331113012-3331210220313000-0303222000322212-3000121230313001-2332320121112310-3112221030310331): complete subsection reference.

- [default_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-2211123331233312-2131002001130330-1033033000202002-1113120302302011-2303321232232213-0030213231213022-0021123310020203-0212203202100301): complete subsection reference.

- [low_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-3211203330313011-1301102230023230-1221213021020013-2103100111020231-1231100223233220-3102232113203303-0222110120112032-2133031112020113): complete subsection reference.

- [medium_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-3331020220302000-0002023130121130-3302231310301101-0330001232103201-2220022300103220-2221032222003310-3222021022322002-2312232130120111): complete subsection reference.

<a id="canonical-1121332302032021-3120302130202232-3131311332021112-3101000221100131-1210330321312113-2201301102013030-0021202121030030-0110331230231233"></a>

## Next pages — tls_config / 211010332120 / 4

- [tls_tcp.tls_cert_params.tls_config.custom_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-3320220203210203-3010313213202302-0013120331113012-3331210220313000-0303222000322212-3000121230313001-2332320121112310-3112221030310331)
- [tls_tcp.tls_cert_params.tls_config.default_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-2211123331233312-2131002001130330-1033033000202002-1113120302302011-2303321232232213-0030213231213022-0021123310020203-0212203202100301)
- [tls_tcp.tls_cert_params.tls_config.low_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-3211203330313011-1301102230023230-1221213021020013-2103100111020231-1231100223233220-3102232113203303-0222110120112032-2133031112020113)
- [tls_tcp.tls_cert_params.tls_config.medium_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-3331020220302000-0002023130121130-3302231310301101-0330001232103201-2220022300103220-2221032222003310-3222021022322002-2312232130120111)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3320220203210203-3010313213202302-0013120331113012-3331210220313000-0303222000322212-3000121230313001-2332320121112310-3112221030310331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031133022013120-2110202131030002-2321222310223123-0012013202101000-0011022233001120-3312310021101000-3313213013012233-0203001312212010"></a>

## tls_tcp.tls_cert_params.tls_config.custom_security — custom_security / 300333120102 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130)
- tls_tcp.tls_cert_params.tls_config.custom_security

<a id="canonical-2113323331002013-1311022023212033-3312022212313032-0122213021312121-1303211030101120-1330002303320230-0032333120113002-2323312212101221"></a>

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

<a id="canonical-2132221132233111-1312203223330201-1133132001113320-0032112303222123-3120321133131213-1322212121102012-1100103302123312-3000123013002211"></a>

## Direct properties — custom_security / 300333120102 / 3

<a id="canonical-3121112032303131-2113333330102113-2302002222011123-1321012321213122-2301100112203022-1302000331302101-0021302310222301-1023212320320132"></a>

<a id="canonical-2122030302022121-1233200200222221-0212312103222233-0200313210201201-1232101022130223-3203131213100303-1112121332213302-1231210111033101"></a>

## cipher_suites property — custom_security / 300333120102 / 4

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

<a id="canonical-2213033221303032-2133232233212220-0010120110122031-1131332230001231-1032003012030000-1201202020100303-2102311131322232-1312010102200300"></a>

<a id="canonical-0312211201133121-2003202210232002-1321020012210022-3201101100331003-0231033101233020-1010201111333232-2001000121133311-1002130120101222"></a>

## max_version property — custom_security / 300333120102 / 5

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

<a id="canonical-1121223320123033-2012032320221312-1303211110000233-3332230021321030-2202030111211230-0313330010310102-0221133303120323-0102111222030332"></a>

<a id="canonical-2300030311133201-3333000033031123-3133321223130033-0322102201112231-0201332210021030-1021022232302320-1023110133222220-0222000322030222"></a>

## min_version property — custom_security / 300333120102 / 6

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

<a id="canonical-0233101302112103-3113122331301030-2300113123202030-1323011111102020-1210231202201233-2012232012112311-1020332010210013-2023100101232103"></a>

## Next pages — custom_security / 300333120102 / 7

- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2211123331233312-2131002001130330-1033033000202002-1113120302302011-2303321232232213-0030213231213022-0021123310020203-0212203202100301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130212223200203-3333232110201031-2332233301203111-0321323123322100-1002110031020002-0001012301012111-0202033002032233-1120333011213322"></a>

## tls_tcp.tls_cert_params.tls_config.default_security — default_security / 033122231023 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130)
- tls_tcp.tls_cert_params.tls_config.default_security

<a id="canonical-1312101011311300-2121123211331302-0322000312232030-1321002111111312-2211110032231032-0030213020110012-2022321120002130-3313030002300012"></a>

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

<a id="canonical-1122121211202210-1313231211211231-2003012132203133-1301302303131212-1211101223213213-1011233230030102-0133101001123301-0021130230130110"></a>

## Direct properties — default_security / 033122231023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310123011123130-3110000002313332-3012311311103101-1323021330311011-3333101320211320-1310322232100230-0211331013033333-2013030331212010"></a>

## Next pages — default_security / 033122231023 / 4

- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3211203330313011-1301102230023230-1221213021020013-2103100111020231-1231100223233220-3102232113203303-0222110120112032-2133031112020113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032333023020133-1022323213020312-1220031010011321-2020300030303030-3120322232312110-0003321133111121-1032003201213310-1110013320022320"></a>

## tls_tcp.tls_cert_params.tls_config.low_security — low_security / 123131001011 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130)
- tls_tcp.tls_cert_params.tls_config.low_security

<a id="canonical-1330110313322011-0310203320203031-3121310203122013-0101231031333102-3201221010321121-2311210210011003-2000223310310022-3131021220323320"></a>

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

<a id="canonical-0210211222112003-1222033322033333-2323222200110312-3231022112223333-2003323222003302-1101211122013212-2010203222131201-3022202331220122"></a>

## Direct properties — low_security / 123131001011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130020033202233-0200301023300203-2312311321100033-0100321210202312-2211223030111201-1322200013222301-3210313120001032-3013221021011212"></a>

## Next pages — low_security / 123131001011 / 4

- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3331020220302000-0002023130121130-3302231310301101-0330001232103201-2220022300103220-2221032222003310-3222021022322002-2312232130120111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120330002230320-3310120022330121-2231212332332111-3313013003010203-1213023322200202-0131203312303323-1222230003232213-0020101231023031"></a>

## tls_tcp.tls_cert_params.tls_config.medium_security — medium_security / 203122033102 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130)
- tls_tcp.tls_cert_params.tls_config.medium_security

<a id="canonical-2212022333012322-1311013322202201-0332230112313031-0020113110133021-3200103330203022-1031311301012301-2010030212133332-3100012033212000"></a>

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

<a id="canonical-1021203301003333-0102321000132023-0030211300033132-0202223211012022-3133132332231023-0112110003123333-0220023103023010-1031132113221123"></a>

## Direct properties — medium_security / 203122033102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320001001201201-0112312122033003-1130210030313233-2011133323130030-2313323331313113-2200312203131023-0001133103211203-2313210121322101"></a>

## Next pages — medium_security / 203122033102 / 4

- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-0322111112121223-3032310231220110-0301012010113032-3030001001013032-3222020211121322-0310302230223200-2310031131231120-2022130022302130)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222203122203030-0001202123331220-1021310031023301-2312013001132302-0313021113031211-0233010130200100-2202231112200131-1131133303323112"></a>

## tls_tcp.tls_cert_params.use_mtls — use_mtls / 010201133111 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- tls_tcp.tls_cert_params.use_mtls

<a id="canonical-3111121020110031-3020010003330322-1122023201122301-3322201021112202-1223113202201010-1103102120111200-3022123010301001-1213101302223032"></a>

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

<a id="canonical-0012300033110121-3323302102332111-2323310133022122-1223321332103231-1330132331131323-0112020221211021-3231210130001033-1303100333021001"></a>

## Direct properties — use_mtls / 010201133111 / 3

<a id="canonical-3013320211222232-1113312010311002-2013331322303032-3122101112232213-2121333130330020-2313232131331221-1231201031202110-2000330200313310"></a>

<a id="canonical-3123023220321020-0023301101331333-0111123200310002-3300120233332330-3223033101230021-1002000302110323-2331101112211321-2102023203011322"></a>

## client_certificate_optional property — use_mtls / 010201133111 / 4

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

- [crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-0013203232221122-3131111332200030-1203322032032201-3020301221113322-2031221333332022-1223321321032033-1230301223121200-3020112202333022): complete subsection reference.

- [no_crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-3112230111023311-3101102123323312-0012002232001030-2310033001231110-3230312221301102-2001013123020332-1113120033111310-0200211220101232): complete subsection reference.

- [trusted_ca](resources--tcp_loadbalancer--reference--group-002.md#canonical-0200003100030200-2110123332302230-3021233032202021-2113003301313222-1010010331203122-1331032301012223-3233003001322033-3201021231230232): complete subsection reference.

<a id="canonical-1331110031120010-0133112021222233-3310123203101103-1120030233300330-2102123221100213-2032103221311323-3021201201200321-3002003101100213"></a>

<a id="canonical-1100120321012330-2202303120112100-3202212031123332-0220201203202333-0000312102011111-1103112200102011-3102330032311031-1100130313201001"></a>

## trusted_ca_url property — use_mtls / 010201133111 / 5

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

- [xfcc_disabled](resources--tcp_loadbalancer--reference--group-002.md#canonical-0232003220003202-2001302003310012-3101300112231123-2000131303013013-0230313130011112-2201230112233231-2311110123132210-0221213112031311): complete subsection reference.

- [xfcc_options](resources--tcp_loadbalancer--reference--group-002.md#canonical-2211233203111001-0112213331212320-2030100120211101-1123203223220031-0213311102100010-1322132321202230-1210010301011010-3231020331123320): complete subsection reference.

<a id="canonical-1102210231202311-1212223130311300-3220001222333312-1330003201030203-1233103312331122-2011300303203303-3223023130323113-0112002011101121"></a>

## Next pages — use_mtls / 010201133111 / 6

- [tls_tcp.tls_cert_params.use_mtls.crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-0013203232221122-3131111332200030-1203322032032201-3020301221113322-2031221333332022-1223321321032033-1230301223121200-3020112202333022)
- [tls_tcp.tls_cert_params.use_mtls.no_crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-3112230111023311-3101102123323312-0012002232001030-2310033001231110-3230312221301102-2001013123020332-1113120033111310-0200211220101232)
- [tls_tcp.tls_cert_params.use_mtls.trusted_ca](resources--tcp_loadbalancer--reference--group-002.md#canonical-0200003100030200-2110123332302230-3021233032202021-2113003301313222-1010010331203122-1331032301012223-3233003001322033-3201021231230232)
- [tls_tcp.tls_cert_params.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--reference--group-002.md#canonical-0232003220003202-2001302003310012-3101300112231123-2000131303013013-0230313130011112-2201230112233231-2311110123132210-0221213112031311)
- [tls_tcp.tls_cert_params.use_mtls.xfcc_options](resources--tcp_loadbalancer--reference--group-002.md#canonical-2211233203111001-0112213331212320-2030100120211101-1123203223220031-0213311102100010-1322132321202230-1210010301011010-3231020331123320)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0013203232221122-3131111332200030-1203322032032201-3020301221113322-2031221333332022-1223321321032033-1230301223121200-3020112202333022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010120302302122-2223322333122133-3101023312100320-1030032131122132-0023201011002111-1011132033012000-0023300032332303-0000020110210301"></a>

## tls_tcp.tls_cert_params.use_mtls.crl — crl / 230010232112 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- tls_tcp.tls_cert_params.use_mtls.crl

<a id="canonical-0201000232020212-2202311101120211-1303001110313122-1323113011333302-0023300000331133-2101223132022230-2001321002132032-1111132022101323"></a>

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

<a id="canonical-1111132213230213-0311231310131323-1312011221013133-3301002312302010-1202120312021300-3320121112000031-3102032021311331-1320333101010320"></a>

## Direct properties — crl / 230010232112 / 3

<a id="canonical-0122220321213333-3032112020213112-0000131113302311-0130313203103131-0031312022112301-2313202322130323-2111333000333032-2133223233320111"></a>

<a id="canonical-1311220130102012-1302121203021030-3320232230110322-1121300311210323-0001021000302230-3211122000120103-0303330321012201-0033222011013022"></a>

## name property — crl / 230010232112 / 4

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

<a id="canonical-3222211200132001-3031220310033032-1310310032033320-3300220203220201-1103321122010233-1023001202031002-2032003321031202-2223222313132122"></a>

<a id="canonical-1311311130322320-2203312320103100-2020103232121012-2001301311331201-0121210312310321-0120132200032010-3332332101122122-1331103032332003"></a>

## namespace property — crl / 230010232112 / 5

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

<a id="canonical-3202100131322103-0100202301112031-2233311310302301-3323223101332211-3020310300300033-2211113331313103-2311312300112033-3302303033300002"></a>

<a id="canonical-3012202100003330-3020002223003220-1011110101132102-3330000321302300-2222311012232131-0323333233302011-1003211012111112-3212030202110032"></a>

## tenant property — crl / 230010232112 / 6

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

<a id="canonical-3030302103130031-2112020031203203-2011000110331211-2203033233213213-3112101112200130-2200001021120213-3311030111202001-1011222213003023"></a>

## Next pages — crl / 230010232112 / 7

- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3112230111023311-3101102123323312-0012002232001030-2310033001231110-3230312221301102-2001013123020332-1113120033111310-0200211220101232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331303132312033-2033300322121233-1311113011022321-1013111120001221-0200212230003120-0223113033013321-1322001231123023-2220112320231201"></a>

## tls_tcp.tls_cert_params.use_mtls.no_crl — no_crl / 230333230022 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- tls_tcp.tls_cert_params.use_mtls.no_crl

<a id="canonical-3321012202302322-2300022233010223-0002002320103100-2201212300331212-2131232302233030-2311321311113112-1100223113313200-2203112310201333"></a>

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

<a id="canonical-2201023300200011-0320023113112312-3122131223011333-2031132022023013-3132001200012131-3032202331233312-2130113110033020-1233021032201030"></a>

## Direct properties — no_crl / 230333230022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212231120113010-3130222211301333-3121003103332333-1132102103013013-1310033321123300-0300112020030101-0330133113310120-1023212200333320"></a>

## Next pages — no_crl / 230333230022 / 4

- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0200003100030200-2110123332302230-3021233032202021-2113003301313222-1010010331203122-1331032301012223-3233003001322033-3201021231230232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331222022232120-3013210101003101-2211013213231222-3033311131113121-3031033121311113-3011023033120033-0203232103332100-2220221301033222"></a>

## tls_tcp.tls_cert_params.use_mtls.trusted_ca — trusted_ca / 123232131212 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- tls_tcp.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-2001121202320001-2202231211201110-2003012013201230-0220313302201133-2330001003222200-1012133301012132-2332233323330302-1302031300000212"></a>

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

<a id="canonical-0021133212211202-3010133232303302-2031000130220031-1123212301232131-3301023310231100-1121111013103011-3111032213023103-0020233203222313"></a>

## Direct properties — trusted_ca / 123232131212 / 3

<a id="canonical-0030003332002301-3033213103121332-2202100321303020-3032001133100103-1320313301321003-1103233113131113-2120211020010311-0331221112301033"></a>

<a id="canonical-3233232312301212-3213031302233020-2322232221330132-1132203112322023-2033113112131130-1030333333112123-2223313000011103-3331113222021003"></a>

## name property — trusted_ca / 123232131212 / 4

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

<a id="canonical-3033303232333313-3122210031010110-1102013100233333-3221003022213232-1133233010020222-1101122133102123-0320200000011021-1313122332030222"></a>

<a id="canonical-1033210033200220-2133031311012301-1220123112332101-1223133022200202-2120103020111220-3213330010003211-0110200213303321-1232122103132202"></a>

## namespace property — trusted_ca / 123232131212 / 5

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

<a id="canonical-0133131301311303-0200121022110032-1313302030123033-2000333320301323-2023030002320232-1003302101033332-2112302300111012-0000133021222213"></a>

<a id="canonical-1203211021122210-1220213323012233-0221221302233210-1223231202323221-2230112211310213-0013120033031101-0313322202102131-3133012132133103"></a>

## tenant property — trusted_ca / 123232131212 / 6

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

<a id="canonical-0132232031020133-2212120033310110-3111321020110032-2021031300132301-1203303232111121-3110030202300201-3123003110020332-2032310130000213"></a>

## Next pages — trusted_ca / 123232131212 / 7

- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0232003220003202-2001302003310012-3101300112231123-2000131303013013-0230313130011112-2201230112233231-2311110123132210-0221213112031311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201012100210200-1202301220332323-2021331013313013-2330012111203123-2213301012321101-3133012201230113-3210213313313333-1323303021222102"></a>

## tls_tcp.tls_cert_params.use_mtls.xfcc_disabled — xfcc_disabled / 231202132200 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- tls_tcp.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-1102003200120212-0001032231120020-3010012232313310-2102131301303203-2110213321321002-0103100013232032-1012012330132120-3012111012230230"></a>

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

<a id="canonical-1312101021303122-1133033002123002-1100033233220330-1230320021210333-2101003113133300-0321311021111302-0110301202301233-2103011012121120"></a>

## Direct properties — xfcc_disabled / 231202132200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003223231102322-2302312310021301-0021110011330123-2320033231221112-2112103101100220-2300202222221003-1213000330123011-0002211323111311"></a>

## Next pages — xfcc_disabled / 231202132200 / 4

- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2211233203111001-0112213331212320-2030100120211101-1123203223220031-0213311102100010-1322132321202230-1210010301011010-3231020331123320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101303300332021-1112133310330111-0022121333233231-1220320110330022-2320131232322321-2022121000011001-3223322311110012-3221201320001101"></a>

## tls_tcp.tls_cert_params.use_mtls.xfcc_options — xfcc_options / 233330333220 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102302012003330-3112123130012120-2323212311013211-2232213133230020-2101330212110310-0123002001102200-3023210032122023-2030132220321301)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- tls_tcp.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-2001133311013203-0331320222002031-1020230000001303-1230320001331322-0033131323332332-2110301023013303-0333023120112033-3333330002310022"></a>

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

<a id="canonical-3121032120123212-3312003201330031-2220100101210122-1130213130001003-1111310333230211-1111100211131313-0023111011333333-2120132213002320"></a>

## Direct properties — xfcc_options / 233330333220 / 3

<a id="canonical-3110021330311023-2330101103023010-2310200210001001-3330102302120121-1231201313101111-0212112322133232-3112023211112133-0223203312230311"></a>

<a id="canonical-0332230011221031-0130122133021210-0223101331330302-2303200002010233-0121100222231031-0212321013331213-2301233303310331-1123001312302122"></a>

## xfcc_header_elements property — xfcc_options / 233330333220 / 4

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

<a id="canonical-2130000211201103-3022313101311111-0023321032122222-0211012232001110-0120123233322210-1201022332313031-3320103302230020-2010013010223013"></a>

## Next pages — xfcc_options / 233330333220 / 5

- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-2003301203003021-0233021122121211-2233213111211112-2233232003100000-0310023133013131-0000031312323032-1332232311131110-1021031210223313)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130213001110132-2323010012310223-3030020120322112-2322202003123101-3301322303001133-2122003131322222-2003033221312111-3303311311003032"></a>

## tls_tcp.tls_parameters — tls_parameters / 100312303013 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- tls_tcp.tls_parameters

<a id="canonical-3110131221302012-0020201332212302-2232113113030133-1301112203103333-1123112003212201-2021231011313230-3030310112330201-3023113233130131"></a>

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

<a id="canonical-2221213202320301-0112330200202011-0333230120111000-3110100321121010-1230200032120322-1120221002331302-0313103031200232-1330212003100013"></a>

## Direct properties — tls_parameters / 100312303013 / 3

- [no_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-3312032102310132-3313333033231033-2033013201212031-0312003323222322-1333102130012332-3313300133122000-1013213111230231-1333131222103221): complete subsection reference.

- [tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120): complete subsection reference.

- [tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003): complete subsection reference.

- [use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000): complete subsection reference.

<a id="canonical-2232202101112102-1320113303322103-3010031310133030-1332123133313000-2120323130202210-2131323330323311-1302011213111203-2232101200310103"></a>

## Next pages — tls_parameters / 100312303013 / 4

- [tls_tcp.tls_parameters.no_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-3312032102310132-3313333033231033-2033013201212031-0312003323222322-1333102130012332-3313300133122000-1013213111230231-1333131222103221)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3312032102310132-3313333033231033-2033013201212031-0312003323222322-1333102130012332-3313300133122000-1013213111230231-1333131222103221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123220332211320-2333101003332003-2233303210011122-3311310203003211-2020021221033331-1123303330120113-2023311223130223-1111111321031210"></a>

## tls_tcp.tls_parameters.no_mtls — no_mtls / 230231133011 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- tls_tcp.tls_parameters.no_mtls

<a id="canonical-1101321203110110-2003011002113030-3132001310310233-1003003302213320-0301113232132202-1101023221222301-3222011302000021-2013123121323323"></a>

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

<a id="canonical-0233201013102223-1111231012321001-1030131320122330-2130233312111221-0002211032313012-2003012003102111-2110202231221100-1012100311231203"></a>

## Direct properties — no_mtls / 230231133011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221000230221313-3300303320201320-0103021002131022-1300222120003030-2321012023030203-1000223100330121-3223030112132010-0232123011121112"></a>

## Next pages — no_mtls / 230231133011 / 4

- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131132211231033-3321233100121200-1012010322331223-3320021132310323-1121320000100111-3231202201021203-1021300112311300-2302023133111130"></a>

## tls_tcp.tls_parameters.tls_certificates — tls_certificates / 321220010233 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- tls_tcp.tls_parameters.tls_certificates

<a id="canonical-1202231212120100-0230313232012210-0101220333230100-3133010011111033-3100212131003130-2032111121032021-3213120013320200-1223223221323012"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0100023133333322-2113003030103310-1120233210023100-3012231123010123-2100312213313321-2030302320201203-2332101110011303-1033131231302111"></a>

## Direct properties — tls_certificates / 321220010233 / 3

<a id="canonical-3323312300322303-0101303101133320-0113031230010102-0112102212033023-1100213323020213-0233230302230231-2031210032231230-3001200130112212"></a>

<a id="canonical-2221302132200103-1303000122031332-2223131333302131-1103120112033330-1030231131210011-3011211003331200-3103100301322230-0333330211210302"></a>

## certificate_url property — tls_certificates / 321220010233 / 4

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

- [custom_hash_algorithms](resources--tcp_loadbalancer--reference--group-002.md#canonical-3301303122202032-0100230000013222-3021011212311010-0230130121121012-1321031020230022-2221312323210312-2333010103101231-0313112121312113): complete subsection reference.

<a id="canonical-2213300101033130-3302203302213113-0322332331210012-0201321200032210-3020333310031000-0133011302021313-1233002210320112-3113323033122330"></a>

<a id="canonical-2300031221113320-3101321213202122-0200311032033201-0010223230311113-1331031133021122-0200233203000110-0100102032110101-0022031201230313"></a>

## description_spec property — tls_certificates / 321220010233 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--tcp_loadbalancer--reference--group-002.md#canonical-1130133232111030-1313232203303110-3131200030333223-0333033223000000-0222210123323020-0202130300123102-1122110121333233-1003022103200212): complete subsection reference.

- [private_key](resources--tcp_loadbalancer--reference--group-002.md#canonical-0100333102101300-3220200132001223-0030120230300122-3331201232232302-1232301113321120-0330221011030203-2110332331210010-2021112323003013): complete subsection reference.

- [use_system_defaults](resources--tcp_loadbalancer--reference--group-002.md#canonical-0102010120213331-3223223023110201-0122232110302212-0323110331103310-3031012201011103-1013011002200011-3202010123303321-1231222302131132): complete subsection reference.

<a id="canonical-2231230121020100-0201221013122302-0223123111101311-1012101110022301-0102103132333222-2313300103312323-2221300013121232-0130200322200220"></a>

## Next pages — tls_certificates / 321220010233 / 6

- [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms](resources--tcp_loadbalancer--reference--group-002.md#canonical-3301303122202032-0100230000013222-3021011212311010-0230130121121012-1321031020230022-2221312323210312-2333010103101231-0313112121312113)
- [tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--tcp_loadbalancer--reference--group-002.md#canonical-1130133232111030-1313232203303110-3131200030333223-0333033223000000-0222210123323020-0202130300123102-1122110121333233-1003022103200212)
- [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--reference--group-002.md#canonical-0100333102101300-3220200132001223-0030120230300122-3331201232232302-1232301113321120-0330221011030203-2110332331210010-2021112323003013)
- [tls_tcp.tls_parameters.tls_certificates.use_system_defaults](resources--tcp_loadbalancer--reference--group-002.md#canonical-0102010120213331-3223223023110201-0122232110302212-0323110331103310-3031012201011103-1013011002200011-3202010123303321-1231222302131132)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3301303122202032-0100230000013222-3021011212311010-0230130121121012-1321031020230022-2221312323210312-2333010103101231-0313112121312113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230231131031123-3232211302102332-3123130022010301-3210021102212112-3230133130133223-2310321331012201-2202033022023113-2212322223021201"></a>

## tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 031320222032 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-3023032010200202-0123321213302310-0200011320123020-2301033322202011-3221011013213010-3000231003323200-3101222231311010-0100003033232233"></a>

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

<a id="canonical-2230331310121031-2200321033300113-2133331200120313-0210321303100120-1100010011121310-1330230211011101-0211000123220230-2123122022232332"></a>

## Direct properties — custom_hash_algorithms / 031320222032 / 3

<a id="canonical-2212220120123321-2133031300330331-0111223332223332-3231033131323101-2102203211122021-1000202233110013-3331303113201222-0100103310000211"></a>

<a id="canonical-0321202301220103-2112302203313331-3102031333210021-2201300200232033-1110103020321023-2000001303100133-3123220201001033-2022333112233103"></a>

## hash_algorithms property — custom_hash_algorithms / 031320222032 / 4

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

<a id="canonical-3300312332030113-0211332330230123-1230210220320123-2330303033023132-3302130002223202-2101321333030112-0322001012101112-0301333120000311"></a>

## Next pages — custom_hash_algorithms / 031320222032 / 5

- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-1130133232111030-1313232203303110-3131200030333223-0333033223000000-0222210123323020-0202130300123102-1122110121333233-1003022103200212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000031200012200-2000121121232302-2320123313220212-2210011112331032-3001123012320112-1300233031012212-1231211110302122-2133122203223202"></a>

## tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 333230312223 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-3023311011102113-3021223201133112-1323211121000232-0303210323200303-2330130013033010-3002032023100001-0212323313231232-2121232130122333"></a>

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

<a id="canonical-3133211013212232-1000230311222003-0033002312020300-3302020133223230-3001033321132130-0232202203312303-1110323020003033-0022222011221132"></a>

## Direct properties — disable_ocsp_stapling / 333230312223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100033322111203-1320212311223022-2010200312323112-3200012001211003-1110111301323120-2213012130122010-0031212201221332-3313332112102330"></a>

## Next pages — disable_ocsp_stapling / 333230312223 / 4

- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0100333102101300-3220200132001223-0030120230300122-3331201232232302-1232301113321120-0330221011030203-2110332331210010-2021112323003013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122113230102120-0223231100210221-2130230031011022-3121001232021312-3002230132013033-2100333021321302-3203312001221301-2030110230111110"></a>

## tls_tcp.tls_parameters.tls_certificates.private_key — private_key / 213312022220 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- tls_tcp.tls_parameters.tls_certificates.private_key

<a id="canonical-1123011310203113-3021011221132332-1033103301311113-1000100031023220-1333112013232012-0000112221132213-3322033303103231-2223231110110032"></a>

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

<a id="canonical-0000211010000300-3000003210030123-1323121013130132-2103030220111320-3033033133332331-0311332231121133-1113030123033023-1020311113321211"></a>

## Direct properties — private_key / 213312022220 / 3

- [blindfold_secret_info](resources--tcp_loadbalancer--reference--group-002.md#canonical-0320313320210230-1002123322100231-1303211120310011-1001023000201010-0110021301010220-1120210310223301-3320002323332001-0323323232130210): complete subsection reference.

- [clear_secret_info](resources--tcp_loadbalancer--reference--group-002.md#canonical-0203120320001120-1303310322101331-1333232131022333-1101032230333221-2201100331122131-0311223122332330-0203003300131013-2010200203233322): complete subsection reference.

<a id="canonical-0012112123022311-0221113333020012-0013020300203201-1320330102001211-3200210221033000-3301230010201133-3022020011323032-0223320122011000"></a>

## Next pages — private_key / 213312022220 / 4

- [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--tcp_loadbalancer--reference--group-002.md#canonical-0320313320210230-1002123322100231-1303211120310011-1001023000201010-0110021301010220-1120210310223301-3320002323332001-0323323232130210)
- [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--tcp_loadbalancer--reference--group-002.md#canonical-0203120320001120-1303310322101331-1333232131022333-1101032230333221-2201100331122131-0311223122332330-0203003300131013-2010200203233322)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0320313320210230-1002123322100231-1303211120310011-1001023000201010-0110021301010220-1120210310223301-3320002323332001-0323323232130210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203233022131121-3213000111221331-2130010023032333-0332103331321231-0032000002210212-0222331330200330-3030121103200333-1033313230331300"></a>

## tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 000000312022 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--reference--group-002.md#canonical-0100333102101300-3220200132001223-0030120230300122-3331201232232302-1232301113321120-0330221011030203-2110332331210010-2021112323003013)
- tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3332301012313030-0020232102210020-3000311111333222-2133112010222120-3222012301032002-0310232101232002-3323230131122013-2331232301032132"></a>

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

<a id="canonical-3300012230213133-1302130123321202-3013122231110333-3330232313020111-3132212003222131-0002333101003000-3100302212233131-3120023003012000"></a>

## Direct properties — blindfold_secret_info / 000000312022 / 3

<a id="canonical-2012210230101302-0132203223022232-1220023022312111-2032112132230330-2301222033320111-0330020313110321-0200003003200313-0011213000113331"></a>

<a id="canonical-1330012120023103-1100302113200232-2211320332313320-3133321320332102-2030321211303223-3210223221321033-3003100332120122-0211111032332321"></a>

## decryption_provider property — blindfold_secret_info / 000000312022 / 4

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

<a id="canonical-3303301233102000-2312323022220232-0223230133231013-2101201011133311-3113222303012022-0013323103231121-3121230313121131-3011300303220132"></a>

<a id="canonical-0222000021233031-1231123201013011-1322103312033233-1211303312232312-3333211020332201-1012101010311322-3321220113223020-3302300302100013"></a>

## location property — blindfold_secret_info / 000000312022 / 5

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

<a id="canonical-1023331223321223-0021203233231033-3110103223203020-0303221210102323-1232203131220130-2113000202030023-1020013013113100-2031022321022321"></a>

<a id="canonical-3301101213330332-1022300132102230-1313130212031031-3001310012210023-3111321302010132-0100211121303013-3313101330103231-2110203231200200"></a>

## store_provider property — blindfold_secret_info / 000000312022 / 6

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

<a id="canonical-3310010201211300-2131000022210122-1112032010121012-1330333123030003-2230321003231011-0010333222212230-1201131312230202-2023200320131300"></a>

## Next pages — blindfold_secret_info / 000000312022 / 7

- [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--reference--group-002.md#canonical-0100333102101300-3220200132001223-0030120230300122-3331201232232302-1232301113321120-0330221011030203-2110332331210010-2021112323003013)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0203120320001120-1303310322101331-1333232131022333-1101032230333221-2201100331122131-0311223122332330-0203003300131013-2010200203233322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113222200321323-0012112222113022-2101111303012010-2231220100333211-2113030321021020-1320201031203022-2112311023332332-2303212100322233"></a>

## tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info — clear_secret_info / 230202311000 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--reference--group-002.md#canonical-0100333102101300-3220200132001223-0030120230300122-3331201232232302-1232301113321120-0330221011030203-2110332331210010-2021112323003013)
- tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-1121332121001303-1330031030003332-2332023200211220-1113111101212011-0113332002311202-2312203002010102-3000023000101023-1120111233020003"></a>

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

<a id="canonical-3310332011330120-1113301102011202-0020322132121330-3332102313213231-0031131011212121-2001320332332003-3333310103102330-1331203233232321"></a>

## Direct properties — clear_secret_info / 230202311000 / 3

<a id="canonical-0102311301232332-0233123210030330-2203000022130033-2032202210202330-0331102020232103-3313133111031201-0203030310332121-3210013321223232"></a>

<a id="canonical-3222233231123101-2022303210213322-3123200223311321-3210300132121030-3221003310133302-3211230210110202-0322231131232313-0223321301003102"></a>

## provider_ref property — clear_secret_info / 230202311000 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0130202101332222-3202202301103120-0212010100131101-1120020103112031-3033201300100321-2003300301101320-1102130200313011-3131223233102310"></a>

<a id="canonical-0313021333312123-3020332310330011-0333102132332102-1202120233130133-2323032120113002-0113201300221322-1202330303231022-0102130113332032"></a>

## URL property — clear_secret_info / 230202311000 / 5

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

<a id="canonical-0031022210231232-0221102030312311-3232113211113100-3020120222323300-1300233332201300-1330312121003302-0222332030100020-0033212202200310"></a>

## Next pages — clear_secret_info / 230202311000 / 6

- [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--reference--group-002.md#canonical-0100333102101300-3220200132001223-0030120230300122-3331201232232302-1232301113321120-0330221011030203-2110332331210010-2021112323003013)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0102010120213331-3223223023110201-0122232110302212-0323110331103310-3031012201011103-1013011002200011-3202010123303321-1231222302131132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112012111120032-2320120121220003-3003013011211122-0133220001323122-1031202330031201-1210202313312133-1020022231211220-1333120023111332"></a>

## tls_tcp.tls_parameters.tls_certificates.use_system_defaults — use_system_defaults / 230312020113 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- tls_tcp.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-3321332112202331-0332220020203310-0303033222211031-3111032001010021-3022330202330210-0230302202001231-1203010221201321-3320220322313121"></a>

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

<a id="canonical-0101033221223113-3202300332230231-3132021233021202-0111023013020303-1002320321120130-2113203123230123-0311210212002301-3200213130100100"></a>

## Direct properties — use_system_defaults / 230312020113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322301032333010-2310123010000132-3322103302103322-2131123131132200-0100010333102330-3101000101201320-2201332010111033-2331203110121012"></a>

## Next pages — use_system_defaults / 230312020113 / 4

- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001323033032222-2331303031031202-1010002210033302-0101030011033001-1102121331203212-0132020010132202-1210103112300002-3122221122033302"></a>

## tls_tcp.tls_parameters.tls_config — tls_config / 022202000012 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- tls_tcp.tls_parameters.tls_config

<a id="canonical-0212300030220000-1030102030120032-3230012030010030-0000021233232103-0220012112012103-0232213020123322-2232021021001302-0012002033001322"></a>

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

<a id="canonical-1013013033213210-2300230212333103-0101032213130131-1232031201033210-1002121113113020-2213013220223310-3222122200032331-2321133111201233"></a>

## Direct properties — tls_config / 022202000012 / 3

- [custom_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-0130120303200203-2030112320030203-0133122321202013-1233130133222211-2322200122103223-0120023212110033-0231012031200131-1233012301231213): complete subsection reference.

- [default_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-3011313102103121-3231233210201221-0022213002310313-1032322300233321-0200302013232133-3300313101210033-2330210102022303-2303013013130210): complete subsection reference.

- [low_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-2300003002320102-2022222323333210-2301101111110221-0202120131321310-1011003021132121-1012223013100001-0302333020233012-0111332122032303): complete subsection reference.

- [medium_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-0122101212100122-0123222332303022-3031313101311112-3132323113331202-1021023112312302-0003111030202031-1320222133300130-2222121221213212): complete subsection reference.

<a id="canonical-3020032120203133-1323231310321302-0122213003320330-1010013120311111-1033212213302032-1331020232101012-3010320011312313-2011012131030112"></a>

## Next pages — tls_config / 022202000012 / 4

- [tls_tcp.tls_parameters.tls_config.custom_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-0130120303200203-2030112320030203-0133122321202013-1233130133222211-2322200122103223-0120023212110033-0231012031200131-1233012301231213)
- [tls_tcp.tls_parameters.tls_config.default_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-3011313102103121-3231233210201221-0022213002310313-1032322300233321-0200302013232133-3300313101210033-2330210102022303-2303013013130210)
- [tls_tcp.tls_parameters.tls_config.low_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-2300003002320102-2022222323333210-2301101111110221-0202120131321310-1011003021132121-1012223013100001-0302333020233012-0111332122032303)
- [tls_tcp.tls_parameters.tls_config.medium_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-0122101212100122-0123222332303022-3031313101311112-3132323113331202-1021023112312302-0003111030202031-1320222133300130-2222121221213212)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0130120303200203-2030112320030203-0133122321202013-1233130133222211-2322200122103223-0120023212110033-0231012031200131-1233012301231213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210121110010032-3212123100020220-1310213330303222-3011213020013123-1221330302232031-0210001032022131-2022310131023022-0231332320221312"></a>

## tls_tcp.tls_parameters.tls_config.custom_security — custom_security / 330120133203 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003)
- tls_tcp.tls_parameters.tls_config.custom_security

<a id="canonical-2101203230103003-0022303032202231-0131223223111012-3233012312103213-1001023021033001-2100022031011221-1203102001323321-1301100001220121"></a>

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

<a id="canonical-2211100213231113-3300203112120113-3321300012302120-0013332123302100-1210210222210320-2103202033030233-1313312333103020-2212300003123103"></a>

## Direct properties — custom_security / 330120133203 / 3

<a id="canonical-2120111230032212-3301001021200020-1023133233032212-1110303201110011-2021022230221201-1131020200030133-2223232121220313-2103311220122113"></a>

<a id="canonical-2320112313111021-1221022220123302-1020213133120032-1311333321003101-0201232332133231-0320201113221233-2001002222320021-1002103211233230"></a>

## cipher_suites property — custom_security / 330120133203 / 4

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

<a id="canonical-1031333132300100-0133123202123002-0000101321001333-2031122131003303-2022221001220200-0233300233000321-2332323322132110-1313021210333320"></a>

<a id="canonical-1002212200110000-0113121010231131-0331012203010220-3030032303131011-3122010011003112-3312213110203330-1133100022003011-3011200233021021"></a>

## max_version property — custom_security / 330120133203 / 5

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

<a id="canonical-0332013010102312-0233000311211020-2331000222113201-0120313002200032-1213122300031232-1003222000210023-3122233313232021-0310303231010021"></a>

<a id="canonical-1010223220022302-1231203300311233-3222300121001120-2323031331210332-2032100031312320-1023231111131121-2030231331203100-2212313010203023"></a>

## min_version property — custom_security / 330120133203 / 6

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

<a id="canonical-3011122212033033-1130312003211122-2031231320300233-3001131020210130-2200300000321113-3020013323323202-1102131212032323-0032101022220133"></a>

## Next pages — custom_security / 330120133203 / 7

- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3011313102103121-3231233210201221-0022213002310313-1032322300233321-0200302013232133-3300313101210033-2330210102022303-2303013013130210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022112022100021-0303001230110131-3230212200010221-3321131213001323-1210021123331032-0022132032203133-1232330022102201-2130110012302002"></a>

## tls_tcp.tls_parameters.tls_config.default_security — default_security / 003101102313 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003)
- tls_tcp.tls_parameters.tls_config.default_security

<a id="canonical-2100100011331032-1320201311013202-1220313213333021-0132130000330331-3020102312322220-2201223031321111-3113031332123203-3321301231210332"></a>

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

<a id="canonical-0322230013021100-3302031020010023-1011310013311003-2103311300032213-2221221132200031-2322220310032203-3012232132133112-1200301212102200"></a>

## Direct properties — default_security / 003101102313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200112231131333-1311231121220310-3033133322013113-3133203300023221-3312120021212212-1202310220033212-0313210013123210-3211111002121300"></a>

## Next pages — default_security / 003101102313 / 4

- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2300003002320102-2022222323333210-2301101111110221-0202120131321310-1011003021132121-1012223013100001-0302333020233012-0111332122032303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122231101233303-3330233031332123-2013023202001022-0001010121223303-1213133223121112-3130000121201112-0203330331221001-1331002011120323"></a>

## tls_tcp.tls_parameters.tls_config.low_security — low_security / 110031002132 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003)
- tls_tcp.tls_parameters.tls_config.low_security

<a id="canonical-1221332221122311-3130011130300221-1202023201232230-3013132212022033-2022330030303302-2230311211331020-3013323230032002-2121232033311301"></a>

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

<a id="canonical-2123222123200330-0131300120230333-1032220010230001-1102332033301112-2323201010030033-1030332233232311-2032321221030132-0111330102202100"></a>

## Direct properties — low_security / 110031002132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021332121303223-2000210300002013-3302230220302322-0031130130331002-0123032210101113-3310330121233210-3023133301110123-0211131233322100"></a>

## Next pages — low_security / 110031002132 / 4

- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0122101212100122-0123222332303022-3031313101311112-3132323113331202-1021023112312302-0003111030202031-1320222133300130-2222121221213212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003222303200131-3032223013030131-1220210010002010-2311221023112213-1012331323131022-2000332111023031-3201123010332221-3223313233331210"></a>

## tls_tcp.tls_parameters.tls_config.medium_security — medium_security / 313230213313 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003)
- tls_tcp.tls_parameters.tls_config.medium_security

<a id="canonical-1020031022132103-0201111302022203-3133231013000123-1122132022031103-2230013201200003-1031232100303201-1001212120302120-0321332000232002"></a>

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

<a id="canonical-1001302201331002-0010000223313333-2121330103312302-0013223200331112-3321311323013133-1131120222301300-0030323303211231-1130213011321000"></a>

## Direct properties — medium_security / 313230213313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113322220022100-1222310021002220-0221212210222210-0101210230213111-1010020213213232-2023022101220010-3203133321210213-2222100232012203"></a>

## Next pages — medium_security / 313230213313 / 4

- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3312211030133232-3210300013231003-3101021322121230-0131120002223323-1231030021010300-2313111331303031-1312020033300201-1030323332322003)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320001103232103-1132122010310113-1311010303001311-3013002220132320-1221202031120131-1331011230322200-3000200010023002-0130022303312300"></a>

## tls_tcp.tls_parameters.use_mtls — use_mtls / 312300023313 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- tls_tcp.tls_parameters.use_mtls

<a id="canonical-1232221302231213-0031210310023113-0211110311201223-1120002312213120-2011321300110012-3200030013001222-2100103323222132-3302122312012030"></a>

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

<a id="canonical-0300203002021301-0130032201210120-2100121300201022-3100221032103212-1020200321310231-3013112032032102-0012002300102131-1122000121121303"></a>

## Direct properties — use_mtls / 312300023313 / 3

<a id="canonical-3221111013201111-3213023210020303-3131030321232001-3203102023003103-0102001010302210-2012103231122331-1203201220231100-3221211132120123"></a>

<a id="canonical-3322212113002221-0133232112021332-1103231231322213-1010331222130212-2330300303000023-3032213121332033-0023202223031301-2320210121213203"></a>

## client_certificate_optional property — use_mtls / 312300023313 / 4

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

- [crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-0221011232132102-3013223103032110-1212203323321123-1200223110001020-2321033103122302-2120320330031323-2303023100221123-2200033122213033): complete subsection reference.

- [no_crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-0123330020221011-0130332120131212-0123313133232112-3111230010033011-2221002101021100-0200001231202201-1212210333330011-0011131110130111): complete subsection reference.

- [trusted_ca](resources--tcp_loadbalancer--reference--group-002.md#canonical-2011233303322103-3011101331001312-3210102103312000-2200300102303003-0101022230330103-1013222332313000-2002201113122021-1203021300001330): complete subsection reference.

<a id="canonical-2003211222130021-0102013022031220-0311111202133301-1223000212300330-0120033121121310-2102321021313102-1300102312222321-2321210202130323"></a>

<a id="canonical-3200303032321312-3001200230231222-3331301002102203-1203111002321021-1212210022112111-2301203123031120-2223231233002012-3331222300322110"></a>

## trusted_ca_url property — use_mtls / 312300023313 / 5

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

- [xfcc_disabled](resources--tcp_loadbalancer--reference--group-003.md#canonical-3221120302133200-2013113331332203-0323332322332331-2113030121313132-3201101323233220-1002123231331100-3030233310332120-3033231201103012): complete subsection reference.

- [xfcc_options](resources--tcp_loadbalancer--reference--group-003.md#canonical-0000312003001302-1122232200100212-2311230011331001-2201013210031300-0230220320310321-3013203311333033-3322330201331111-2213231322023111): complete subsection reference.

<a id="canonical-3131332312222230-1303011221222312-1111302310213331-0331111121110032-0200023212332322-3000222122101120-1221311321023220-0111102223211102"></a>

## Next pages — use_mtls / 312300023313 / 6

- [tls_tcp.tls_parameters.use_mtls.crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-0221011232132102-3013223103032110-1212203323321123-1200223110001020-2321033103122302-2120320330031323-2303023100221123-2200033122213033)
- [tls_tcp.tls_parameters.use_mtls.no_crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-0123330020221011-0130332120131212-0123313133232112-3111230010033011-2221002101021100-0200001231202201-1212210333330011-0011131110130111)
- [tls_tcp.tls_parameters.use_mtls.trusted_ca](resources--tcp_loadbalancer--reference--group-002.md#canonical-2011233303322103-3011101331001312-3210102103312000-2200300102303003-0101022230330103-1013222332313000-2002201113122021-1203021300001330)
- [tls_tcp.tls_parameters.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--reference--group-003.md#canonical-3221120302133200-2013113331332203-0323332322332331-2113030121313132-3201101323233220-1002123231331100-3030233310332120-3033231201103012)
- [tls_tcp.tls_parameters.use_mtls.xfcc_options](resources--tcp_loadbalancer--reference--group-003.md#canonical-0000312003001302-1122232200100212-2311230011331001-2201013210031300-0230220320310321-3013203311333033-3322330201331111-2213231322023111)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0221011232132102-3013223103032110-1212203323321123-1200223110001020-2321033103122302-2120320330031323-2303023100221123-2200033122213033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031212221223211-3022121223331103-0212203303202220-3232013103202312-2022000123230122-0021023010332023-1113320201002322-0220033022121021"></a>

## tls_tcp.tls_parameters.use_mtls.crl — crl / 231200100012 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- tls_tcp.tls_parameters.use_mtls.crl

<a id="canonical-2231002300022321-0310320020013321-3013231001302330-0222022201020003-1333312021222020-0123002203132031-3221023102230103-2323210312230300"></a>

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

<a id="canonical-0101103223011323-2011231101203212-0003011320031323-0033201130313332-2131110102203133-3123030020013130-0133100102130032-3203130323133013"></a>

## Direct properties — crl / 231200100012 / 3

<a id="canonical-0230012313210303-0033010223133013-0300200112331132-0321122221223033-3132301012030211-3302203013101313-3231320322233230-3000032003213223"></a>

<a id="canonical-3003213021111132-1001321322101300-3101020123012231-2203202323131113-3321100222121231-2101102002000113-2322331132313321-1232012031130113"></a>

## name property — crl / 231200100012 / 4

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

<a id="canonical-0113221301022030-1221222321132301-0010123121103202-0300021231303212-0212320221200121-1223111003120310-3321311201112000-2313202122122223"></a>

<a id="canonical-0100020013313010-1110131102222212-2300321010313132-3221033332213013-2032110312023011-1230330102112302-1122330203012133-3010312023113030"></a>

## namespace property — crl / 231200100012 / 5

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

<a id="canonical-3321103023302322-1130302103302300-3232302121212231-3213320201031033-3232213100110303-2101303022312333-2103112303122232-1220320212122000"></a>

<a id="canonical-1100311020223312-2121031122233312-2033210302101113-2211211132013101-3300333231322230-2012320112120001-2100100022103312-2132001111103330"></a>

## tenant property — crl / 231200100012 / 6

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

<a id="canonical-1010210122232111-3302001032332303-1222203313231033-0031103301223323-1131321313301022-2312220102203020-1303322122310233-0033303132132020"></a>

## Next pages — crl / 231200100012 / 7

- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0123330020221011-0130332120131212-0123313133232112-3111230010033011-2221002101021100-0200001231202201-1212210333330011-0011131110130111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313113120011013-2302231333013031-3203322113012322-0321222223113233-1103013100231201-3332233023313112-1301310212002100-2320011010312332"></a>

## tls_tcp.tls_parameters.use_mtls.no_crl — no_crl / 302330330220 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- tls_tcp.tls_parameters.use_mtls.no_crl

<a id="canonical-3130121333221110-2303030212130233-0101213010223232-3212112032320032-3310033333332110-0330323323320023-1122000222002203-2230112223222001"></a>

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

<a id="canonical-2011322132302221-3330131302311110-0031013232022013-0010110030032310-3221033210110020-0120101302113203-3231003231003200-1332030332030102"></a>

## Direct properties — no_crl / 302330330220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030210320021000-1012312333223331-0302232313002010-0122131213220300-0131022201232232-0003332131323221-0320212200311013-1000111323232023"></a>

## Next pages — no_crl / 302330330220 / 4

- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2011233303322103-3011101331001312-3210102103312000-2200300102303003-0101022230330103-1013222332313000-2002201113122021-1203021300001330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212210231001311-0200231202130201-0130323031010200-0232033231003320-0230203302111310-3003322123321213-1300020003303020-0310231221131021"></a>

## tls_tcp.tls_parameters.use_mtls.trusted_ca — trusted_ca / 332322112100 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- tls_tcp.tls_parameters.use_mtls.trusted_ca

<a id="canonical-0223202131300212-0310312222130113-3132132013321133-1113010120111302-1112023223102000-3113203211000312-1300033232131000-2303101302210311"></a>

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
