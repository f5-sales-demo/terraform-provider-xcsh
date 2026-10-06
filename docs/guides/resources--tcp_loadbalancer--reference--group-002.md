---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-0021210012032100-2132021033222132-1011100102011210-0332000131011220-2322330320300000-2221003302023231-1313021331032231-0032033010030303"></a>

## `advertise_custom.advertise_where.virtual_network.specific_v6_vip` property

Type: `"string"`. Optional.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0222332231021212-3213220030121313-1300022001031030-0212122201001200-0301322230020123-1112103233131312-1120020332322011-2133131021022333"></a>

<a id="canonical-0010213003013131-3033233223000220-0121001321223212-0311223123123003-2333201120020303-1033230133003101-1230000002210120-0221102120233023"></a>

## `advertise_custom.advertise_where.virtual_network.specific_vip` property

Type: `"string"`. Optional.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](resources--tcp_loadbalancer--reference--group-002.md#canonical-0320000010022301-3300003102221232-3203200210312200-0013230101012112-0012033102333001-2000100301113312-3220103030121213-2102232111101311): complete subsection reference.

<a id="canonical-3022220111330023-2033301113121011-0103101011332303-2132012011023202-2013032210001032-3030011120200301-1033300322022213-2033331112303313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network.default_v6_vip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-1003112313011332-2020131311011001-0030223322103102-1202300132320101-1111022212122021-1030301120012112-0331021223030333-0330020330031020)
- advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-2300230330233301-3220031330202032-0311013320002003-1311323021331311-1020131330131030-0132021021132223-3032203123132010-2312313323010301"></a>

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
default_v6_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032130230213230-3303113002020202-3030213303311100-1301120021221311-2201300123303031-2332211032000302-1000323111133012-2212221120012222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network.default_vip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-1003112313011332-2020131311011001-0030223322103102-1202300132320101-1111022212122021-1030301120012112-0331021223030333-0330020330031020)
- advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-2213333200110023-3303310232003132-0012011013120201-3110100303102200-1031012011220110-1121102000032333-1232133120312133-2211203032231001"></a>

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
default_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320000010022301-3300003102221232-3203200210312200-0013230101012112-0012033102333001-2000100301113312-3220103030121213-2102232111101311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network.virtual_network` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-1003112313011332-2020131311011001-0030223322103102-1202300132320101-1111022212122021-1030301120012112-0331021223030333-0330020330031020)
- advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-2013303220021031-3010201222213310-3201121113013223-0220120113123300-3330010232230213-3022010032100010-1111302213333313-2200300113123021"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320310202003112-3322202221010101-1133010103100200-3112331333300332-1322331120101332-1303311323323312-1013220320003201-3313200012110323"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_network.virtual_network`

<a id="canonical-2111332202231002-2133301010003332-0220213322013133-2113011322121202-1031113200311232-0212130101220311-1320110122122020-0213111313222221"></a>

#### `advertise_custom.advertise_where.virtual_network.virtual_network.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2210312203132311-0302010313032130-0210310312223100-0000203122103221-1032322202020311-0231112112201013-2112311330322220-2212131223210002"></a>

<a id="canonical-0003123303001122-3302102000120222-0002023120101302-0000123030111220-2302333223110122-2311021022111110-1021330221122311-3111200330132021"></a>

#### `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3300312221232222-1021131103013010-1323211203211030-3110322221320310-2011220332222010-3223333312322313-3321211330202223-2111332103230301"></a>

<a id="canonical-2001012231020223-0112122213101223-1330131210223023-0032133111120232-2322300331320012-1232220303012130-0213201012131310-0122322213112001"></a>

#### `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1122231030213111-1302221130230031-0331311330131011-0113323220002101-1200331300203123-1110330313212100-3233313332101212-2310321002101013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- advertise_custom.advertise_where.virtual_site

<a id="canonical-2101030020321030-0020230030332110-0302121201011233-0232202000223022-0333133322220100-2311021122323033-0333031032200232-1120303320320100"></a>

Type: `"object"`. single nested block, Optional.

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-2133101220001312-0030221023201023-3210311032120302-2213313101302031-3003312030013323-0200203000310110-1021030133202123-1301122211301123"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site`

<a id="canonical-2301303022310123-2132120022313212-0201313013313312-2201213323302000-1313200131130131-1221312100003112-2131122320220230-0021302313232211"></a>

#### `advertise_custom.advertise_where.virtual_site.network` property

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_NETWORK_INSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_IP_FABRIC","SITE_NETWORK_OUTSIDE","SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_SERVICE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](resources--tcp_loadbalancer--reference--group-002.md#canonical-0023122310132301-1032230311102222-0133211121311212-1302103110110131-0202200031103103-1312221201211012-0012101231330202-1301002220111023): complete subsection reference.

<a id="canonical-0023122310132301-1032230311102222-0133211121311212-1302103110110131-0202200031103103-1312221201211012-0012101231330202-1301002220111023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site.virtual_site` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- [advertise_custom.advertise_where.virtual_site](resources--tcp_loadbalancer--reference--group-002.md#canonical-1122231030213111-1302221130230031-0331311330131011-0113323220002101-1200331300203123-1110330313212100-3233313332101212-2310321002101013)
- advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-2211322122230300-3333302211031120-3102332010002203-3121213331101203-0030103030113021-0133303121013211-1013332322203013-2130123113201220"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3101302231110101-0232110112302210-2322133121132213-1202111232210022-0001232323210010-0010031010130223-0022130123331120-1221032023221123"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site.virtual_site`

<a id="canonical-3311113020331212-3021133312100032-0231203321312322-0320022012011303-0213300123223311-1033211010201131-1312222322203331-0322330230121022"></a>

#### `advertise_custom.advertise_where.virtual_site.virtual_site.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3201212132120202-3233100321103132-0212301123212310-2233230200031313-2233032313031033-3031221310033103-3023013200221313-1220130231121022"></a>

<a id="canonical-2111031132222200-1130200111313223-2312100211130121-1122131100021321-0322313002310211-3111132322323103-2320111023313331-3222231321323100"></a>

#### `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0020020132032100-3013000112130012-3213222301110023-0003110112110013-2131312112220222-1011001123310230-2321001211323311-2121102223230230"></a>

<a id="canonical-1330131212021222-3032012201300312-0232120002223202-3133003012013203-0000201220332003-0101002011303221-2203231203121220-3330220302323221"></a>

#### `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1121201333003001-3211103001030232-1120020013233002-1201212112111322-2313032201232300-2313132333210321-3020122003000202-2203220100301301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site_with_vip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-0331011202312231-2001120230211033-2322202110213221-3101121130032223-3101133203330023-3113302012121122-1320122101033030-0210332013213123"></a>

Type: `"object"`. single nested block, Optional.

This defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

Receipt-pinned upstream constraints:

```json
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
virtual_site_with_vip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323022310121321-0130322321231122-2201011111011233-1001320133332300-2021233012330030-3100123321002312-0021330222003233-2332131313221332"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site_with_vip`

<a id="canonical-1311000001220100-0032223213001313-1020333321030333-3010333111302320-1130002231032222-3103102000111023-1311301330201112-3023311101100022"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.ip` property

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1200202002213230-3320332101322113-2003220320330220-0031313021101033-1202031012210013-2232332121120131-1320103011220031-0130000333232001"></a>

<a id="canonical-3311313010313202-2023213332233320-0023021031211323-2033022020123012-3213310203002110-3232233322311321-0122122011303012-0332232332012331"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.network` property

Type: `"string"`. Optional.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on virtual-site with specified VIP

All outside networks.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_NETWORK_SPECIFIED_VIP_INSIDE","SITE_NETWORK_SPECIFIED_VIP_OUTSIDE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
  "enum": [
    "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](resources--tcp_loadbalancer--reference--group-002.md#canonical-2011102020302121-2213133103233212-3030222111013333-1023313331013001-1123031331021123-0203213100310001-2101011002203333-2033111322013123): complete subsection reference.

<a id="canonical-2011102020302121-2213133103233212-3030222111013333-1023313331013001-1123031331021123-0203213100310001-2101011002203333-2033111322013123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- [advertise_custom.advertise_where.virtual_site_with_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-1121201333003001-3211103001030232-1120020013233002-1201212112111322-2313032201232300-2313132333210321-3020122003000202-2203220100301301)
- advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-1001202121302300-2231113032010323-3222121313021223-1131311012013301-3103201030112201-1010313110102030-1133101003210011-2311300302223002"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2003212332123031-1302130320031332-1311301020032230-3002301022202030-1020032112000111-2202023032300331-0112002222132011-3201310211211203"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site`

<a id="canonical-3300122301013231-2022130100012202-1102220032121033-1313123233122100-0322113020111120-2022212030002213-1233121100013313-3133002220012230"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0003003330211323-1131132201203213-0311320021031021-2100200001023202-1122311021320213-2010223310111032-2113030022130310-3101213313010230"></a>

<a id="canonical-2230001323030113-2110321000033332-1123220311301113-3111132202012211-3212323202000212-1020102233012110-3122312231310223-2233313121010012"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2311000230231031-3102332121030203-0210002112013300-1312301320301332-0002133333213102-0100330301130030-2020120010303223-2213222313133332"></a>

<a id="canonical-2310233100330313-2021020013103131-1000031113031201-3331230011001011-3011333021200303-2002323111012003-1202203112122300-3233220220033320"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2310032102311332-1320010320220300-3000131113132003-3031300330001303-3011312212131020-2233012333120212-1031220012331120-3102210330032022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.vk8s_service` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- advertise_custom.advertise_where.vk8s_service

<a id="canonical-2032303212230003-0200223230112002-3132232213103110-2001011022332302-0332200200001332-2033111311013003-3012002113232130-3102101201313131"></a>

Type: `"object"`. single nested block, Optional.

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2122323113321312-0332012122122032-0112110323313323-2100311221032133-0003120331022200-3301212010311233-1022212131102312-0302102130312223"></a>

### Direct properties for `advertise_custom.advertise_where.vk8s_service`

- [site](resources--tcp_loadbalancer--reference--group-002.md#canonical-0011122310030103-3322230232311222-2232133023303210-2301302032000120-3001000330230130-0232311010133030-1021101232130312-2031310113021012): complete subsection reference.

- [virtual_site](resources--tcp_loadbalancer--reference--group-002.md#canonical-0102002101330311-3203311313223300-1030302122122311-2331332203030221-0201130111312302-1310112222312130-3133301001133013-2333130010233330): complete subsection reference.

<a id="canonical-0011122310030103-3322230232311222-2232133023303210-2301302032000120-3001000330230130-0232311010133030-1021101232130312-2031310113021012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.vk8s_service.site` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- [advertise_custom.advertise_where.vk8s_service](resources--tcp_loadbalancer--reference--group-002.md#canonical-2310032102311332-1320010320220300-3000131113132003-3031300330001303-3011312212131020-2233012333120212-1031220012331120-3102210330032022)
- advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-3033230022332321-0231212233003132-1132003120132200-2211031231220011-1231101313123321-2310010200120010-1113023220233311-0213023330031301"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3133133211030331-2123302113031033-1302112310231333-2321100133112221-1130313302002213-3302111303101112-3012123312232213-0112110231310310"></a>

### Direct properties for `advertise_custom.advertise_where.vk8s_service.site`

<a id="canonical-3130032322213030-3032022102012313-3313210030131300-3211013020303030-1301122220320022-2201103230332313-1201132101202320-3100221303012111"></a>

#### `advertise_custom.advertise_where.vk8s_service.site.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1232133211032003-1113323021002211-2332123230032210-2120232323022132-3323200311210220-3302102123233030-0030003333103222-1132020230133100"></a>

<a id="canonical-2013212303101033-2102013201013210-3312303120103221-2331103032333221-1121030320101102-3323023312232113-1230322201322033-3012130000012302"></a>

#### `advertise_custom.advertise_where.vk8s_service.site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0000312230110203-3123331312133333-1022111132213330-2000001210002232-3120112202210003-0333003010100301-2012231031233212-3131310002000233"></a>

<a id="canonical-3120202130210103-3030312112113121-0333233002110322-3233011130210210-3321300030302100-0131303022121202-3011213011203112-3331333131312013"></a>

#### `advertise_custom.advertise_where.vk8s_service.site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0102002101330311-3203311313223300-1030302122122311-2331332203030221-0201130111312302-1310112222312130-3133301001133013-2333130010233330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.vk8s_service.virtual_site` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-0101013130323102-0111020001013232-1223003210233021-3112202320110020-1310022023311310-1003000113120300-2312202022202013-0222311212101130)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332)
- [advertise_custom.advertise_where.vk8s_service](resources--tcp_loadbalancer--reference--group-002.md#canonical-2310032102311332-1320010320220300-3000131113132003-3031300330001303-3011312212131020-2233012333120212-1031220012331120-3102210330032022)
- advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-2112123100031210-3120032221323121-2033302333321130-3323200332003223-0033120310200021-2310123321010101-0023120132332210-1312203221311012"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0231002010331021-3010221222022102-2213010012131212-0102110310222213-2102123231100232-0102122230330022-3301321022003202-3221002013231012"></a>

### Direct properties for `advertise_custom.advertise_where.vk8s_service.virtual_site`

<a id="canonical-0111300312122122-1011232001322002-2333203111203012-0222021231213031-1312213121121211-1231012102130332-0020303021030331-2201020200031100"></a>

#### `advertise_custom.advertise_where.vk8s_service.virtual_site.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1112303020300010-1203231213133311-2033200023002001-3211030020102031-2333320332030320-3330201013133013-0003122023011321-3231201331220222"></a>

<a id="canonical-3312131220112303-2123101313121013-3102210322313313-0133100320300203-3033221120132232-3010120311101011-3330032001303022-0113322222102222"></a>

#### `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2200121221021111-1122220022210202-0222321303323220-1102231331110211-1301000020323311-3232300100011202-3301233232010120-2021130302312302"></a>

<a id="canonical-3000213222013312-2032103330333010-2011122332031011-0333012232112231-3210220231231303-3022213101210110-3011103221202232-3321021201122010"></a>

#### `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1101101122210003-2330111331303010-2130211300213202-2120322121300012-2033321200023022-3103313203330222-2302000021033232-1221001021130221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_on_public` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- advertise_on_public

<a id="canonical-0302313131313103-2320203031130213-3133301303203210-2220200131111230-0233011312332320-0230200102303112-2311223301221223-3032232213032211"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2210200100333300-3310022323203110-2211010032122103-3030312302112303-2120030331223232-0101131021112331-1322110011313002-0001000233311031"></a>

### Direct properties for `advertise_on_public`

- [public_ip](resources--tcp_loadbalancer--reference--group-002.md#canonical-1202122111210021-3012203133200102-2103133031030132-1001321000010000-0321112231023232-3212123332122232-0230203020213311-2312120301320102): complete subsection reference.

<a id="canonical-1202122111210021-3012203133200102-2103133031030132-1001321000010000-0321112231023232-3212123332122232-0230203020213311-2312120301320102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [advertise_on_public](resources--tcp_loadbalancer--reference--group-002.md#canonical-1101101122210003-2330111331303010-2130211300213202-2120322121300012-2033321200023022-3103313203330222-2302000021033232-1221001021130221)
- advertise_on_public.public_ip

<a id="canonical-2322221111012111-1133332103010101-0013232322003313-2002023331303131-3211221122003301-3301212102232332-3201101333132303-2211210301321233"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3203110313130210-2030112033131202-0001230233033003-2200133123010032-3120312103211100-3301121030123013-1100200123111012-0031212032112102"></a>

### Direct properties for `advertise_on_public.public_ip`

<a id="canonical-0002200333201201-1121312313131333-3222011322012223-2011231033301000-0302211010211332-2232221113001132-1023211223120230-0300200300332131"></a>

#### `advertise_on_public.public_ip.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1312201323033122-3123210203001233-2320223320220233-2222201201301212-0230301203333032-3310212022000033-3000203121322123-2333333213231000"></a>

<a id="canonical-3021232113311023-0310130301202313-1232330003030302-3221001312122221-3312301021010313-0202121200000231-1130002120113303-0313101000222101"></a>

#### `advertise_on_public.public_ip.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1303111231321103-0130122311331131-0210132201310122-0310003201212112-1001120102120013-2222110330131000-2000310133201330-0131022101011333"></a>

<a id="canonical-2313110011132012-2322331012133112-0232020232023123-3102131111311332-3221310111332030-3211113222003310-2330202020121231-1012331331332323"></a>

#### `advertise_on_public.public_ip.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3333100232210200-1000122300031311-3012302322223020-3032022132323133-3310012110333012-2333013330001120-1130222201132013-1120103002011310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_on_public_default_vip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- advertise_on_public_default_vip

<a id="canonical-0321202223002312-2000122002021122-1000333311113001-0202202003021011-3030231120013213-3312123301110310-3323231213110121-0311213111231321"></a>

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
advertise_on_public_default_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233300113221313-3110203321303320-1132100212321211-3233111001220210-2202120330231303-1320322300203123-3132233222313123-1030123011100010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_lb_with_sni` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- default_lb_with_sni

<a id="canonical-0313202033230201-1300210123230312-1001011322302321-0120211103330203-3120232100300130-1120303231202201-1102310202111313-2312111111000220"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_lb\_with\_sni, no\_sni, sni; Default: default\_lb\_with\_sni\] Configuration
parameter for default lb with sni.

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

OneOf alternatives in this subsection:

- [default_lb_with_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-0313202033230201-1300210123230312-1001011322302321-0120211103330203-3120232100300130-1120303231202201-1102310202111313-2312111111000220)
- [no_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-3013302200020113-0111322303322230-0000110230220313-0232032310130333-2323102130310131-1010123111330300-2220231131123000-2211320231132021)
- [sni](resources--tcp_loadbalancer--reference--group-003.md#canonical-2121331132131131-3033031100120302-0330132302022103-3000220102112000-3011311311311201-2200311121231223-3011221000201220-3201310112120322)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_lb_with_sni = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322213221112110-0231131200021100-2323031112011321-2113223331230331-1110202310201022-0323112230111011-3203121033311100-1130120122122120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `do_not_advertise` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- do_not_advertise

<a id="canonical-3132133212102201-2021300201322133-0333101230320001-0201211013310101-2132321123323201-1233011231122201-0032202131001100-1033022212130331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise.

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
do_not_advertise = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322113213320330-0102033300000032-1130020131303320-1130033301222310-0100131123323030-1300033212333030-2033201321130131-2130003330320013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `do_not_retract_cluster` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- do_not_retract_cluster

<a id="canonical-2132121202202312-3303103110203131-0123121231302102-1022300131113230-2101310330031313-1031133323302233-2310000223033320-3130011202233233"></a>

Type: `["object", {}]`. Optional.

\[OneOf: do\_not\_retract\_cluster, retract\_cluster\] Enable this option

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

OneOf alternatives in this subsection:

- [do_not_retract_cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-2132121202202312-3303103110203131-0123121231302102-1022300131113230-2101310330031313-1031133323302233-2310000223033320-3130011202233233)
- [retract_cluster](resources--tcp_loadbalancer--reference--group-003.md#canonical-0223212221221323-1003333011230323-2311103300002012-3022112031333333-0211120323003211-0323113322301032-0013033220133100-0321130210120033)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
do_not_retract_cluster = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231312101111010-1230322232102131-1310111121030233-1233021020330300-2012003120131022-3230203210022022-1331310221131200-1310111032002330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hash_policy_choice_least_active` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- hash_policy_choice_least_active

<a id="canonical-2120022010132212-3200121332222303-3312131310001132-1220100002110001-3013320221312220-0031002113013130-1223201110102021-3102010323112103"></a>

Type: `["object", {}]`. Optional.

\[OneOf: hash\_policy\_choice\_least\_active, hash\_policy\_choice\_random,
hash\_policy\_choice\_round\_robin, hash\_policy\_choice\_source\_ip\_stickiness\] Enable this
option

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133201212213303-3220313201223303-0031130110231110-1213131231101232-3321233020122003-3302010033002303-0202202110320320-1222200211330032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hash_policy_choice_random` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- hash_policy_choice_random

<a id="canonical-0320110201222002-0210223103211030-3332203031121102-3103000311221313-2002102321221001-2330032210323323-3301102223121201-1031013122301231"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for hash policy choice random.

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
hash_policy_choice_random = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011303133011102-2021121122120220-1133330230131122-0002121330231010-2103102030013022-2020210201310012-0000100210012131-3213113301223032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hash_policy_choice_round_robin` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- hash_policy_choice_round_robin

<a id="canonical-3322321013232230-1300300103220333-2220321012300001-3102030202321111-2033300233313210-0220020121331312-2032311221020001-1221330012002212"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for hash policy choice round robin. Defaults to \`map\[\]\`. Server applies
default when omitted.

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
hash_policy_choice_round_robin = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123323133211320-2223021102110033-2023122331312031-1312102001331031-2023122213321233-0303012321200222-1233301221213131-3010122320223012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hash_policy_choice_source_ip_stickiness` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- hash_policy_choice_source_ip_stickiness

<a id="canonical-3101023012331020-0102212011231001-1131323230001301-1323130122000230-3202123120312213-3200111331232121-1032310313320003-1122101121103222"></a>

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
hash_policy_choice_source_ip_stickiness = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211303031122133-3002101102330112-3303102231112031-0323212233100110-3023331222023022-0122331220121021-0201223330103012-0220122230220001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_service_policies` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- no_service_policies

<a id="canonical-0102302021033021-2301111022111023-2331021210033212-0131133021322033-3000003112132020-3310331033321310-3001103300120102-2322303233312111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no service policies.

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
no_service_policies = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332121320022012-0131110131002221-3211220033203220-3333221011312320-1303130303113032-2011002222333033-3001212201031002-1010101011201223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_sni` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- no_sni

<a id="canonical-3013302200020113-0111322303322230-0000110230220313-0232032310130333-2323102130310131-1010123111330300-2220231131123000-2211320231132021"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
no_sni = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100010003132033-1331210120313310-0212003223033302-1320122132000102-1302201010312200-3333221331300211-3332023322113011-0102012310233122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools_weights` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- origin_pools_weights

<a id="canonical-3300102231110023-1122311101122321-3133123320300330-3102222002313323-1001120223303033-3223101033122101-0032121010131101-3233033211031232"></a>

Type: `"object"`. list nested block, Optional.

Origin pools and weights used for this load balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1002000111000330-3030223113203102-2133333000321010-2311210223332131-2312210010332123-0233333023312203-2333203132220302-3233233330201000"></a>

### Direct properties for `origin_pools_weights`

- [cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-1011120330133113-3221310102110212-3012110211103032-1212313330022332-2303330232010101-3212113330300201-2331100110002001-3330111113013333): complete subsection reference.

- [endpoint_subsets](resources--tcp_loadbalancer--reference--group-003.md#canonical-0233132302001113-1013122023210222-1222022303323200-3120101033031121-0232202032202111-0031100302133121-1303233330230201-1301123310033123): complete subsection reference.

- [pool](resources--tcp_loadbalancer--reference--group-003.md#canonical-0101012032110312-0012103013221021-2000123112202033-2211320112310201-0213110000103010-1321023222113211-3031331212323121-2320012220221031): complete subsection reference.

<a id="canonical-2023023300223021-1030013322302200-0323020103311210-3121211203123312-0132122310301111-2310333003000230-2320310032131322-3231122110302300"></a>

<a id="canonical-2223123022011200-0203003211033201-2023321311303332-2113330312211030-2011333111203033-0102012121202203-1022123100111120-2312023023032022"></a>

#### `origin_pools_weights.priority` property

Type: `"number"`. Optional.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3001133123212230-3323222011321301-1001120202113303-3013211001201201-0020221203213232-1333331210323303-1301023012203011-0311110123302130"></a>

#### `origin_pools_weights.weight` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1011120330133113-3221310102110212-3012110211103032-1212313330022332-2303330232010101-3212113330300201-2331100110002001-3330111113013333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools_weights.cluster` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-2100010003132033-1331210120313310-0212003223033302-1320122132000102-1302201010312200-3333221331300211-3332023322113011-0102012310233122)
- origin_pools_weights.cluster

<a id="canonical-2110022331031320-2323210200030022-2201303121210111-2302002332211320-1202301111330323-3122312123013131-3222123001321121-1202200010031330"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1031120003132011-3032221320110312-3302213120202210-3011303211220322-3011220100013113-3122032212223122-1101231030031202-0223201101201111"></a>

### Direct properties for `origin_pools_weights.cluster`

<a id="canonical-2221201022222220-1101232211132331-1211322311311012-0231322230031012-3120232331301301-2201122100233320-1022002012301012-1202111222101233"></a>

#### `origin_pools_weights.cluster.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1312321133020223-2023020303320302-1113321133221211-3310033222323121-2220011331211020-1223021233121022-2002130332202303-1101110012302002"></a>

<a id="canonical-3332013031122031-3032231223210321-3000003020130003-2303122133310231-1021222212031211-2330001321033000-2331001130110200-2201132230030201"></a>

#### `origin_pools_weights.cluster.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0303130010011030-3130300323203021-1212130313322132-2232300322020230-3330311013001331-3111111123120112-3333021020320230-0311201223201221"></a>

<a id="canonical-0231202100120323-1232303110333022-0221002000220121-2201231233333321-2132013333223232-2313333002323231-3000103201212012-1120100320312122"></a>

#### `origin_pools_weights.cluster.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
