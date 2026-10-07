---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-0000021310102303-2201201300103032-3133321003021010-0301210023122001-2003320011023003-0102210320130331-3020230030232112-0211223310330021"></a>

## `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [virtual_network](resources--bigip_http_proxy--reference--group-003.md#canonical-3222021333132003-2333132213331001-3133203331320301-0111213323031103-2331130213031223-3120230321300133-1220111302102031-3002001333122330): complete subsection reference.

<a id="canonical-2232132011202230-0011020032122231-2322101110210313-2213210330131212-3213102111002233-0023310113202322-0112000111002302-2120310221102111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-0303023203101011-3310022012313303-0310200232211333-2320120210030203-3001331213320020-1120321132301203-0320311000020123-1101213123321031)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-3031102321023301-0003233200231212-2100221311200220-1021002103112311-2121132033022033-1101123331113010-3202012121110012-3213323011331103"></a>

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

<a id="canonical-3020010230021021-3130200302202230-0313320010003112-3211122021202113-3033220110320231-3023202202311100-3303230112201010-2123122113332211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-0303023203101011-3310022012313303-0310200232211333-2320120210030203-3001331213320020-1120321132301203-0320311000020123-1101213123321031)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-2333123311103322-0333002203111322-2010313023322001-0331122223103000-1222300102001212-0033200321233203-3133131011121032-1330311333132333"></a>

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

<a id="canonical-3222021333132003-2333132213331001-3133203331320301-0111213323031103-2331130213031223-3120230321300133-1220111302102031-3002001333122330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-0303023203101011-3310022012313303-0310200232211333-2320120210030203-3001331213320020-1120321132301203-0320311000020123-1101213123321031)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-1022031103002101-3231301032320231-1213212331322033-1122203102323223-1001122032030333-3110033020100112-1321122200333333-0202022222121102"></a>

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

<a id="canonical-1231022200131103-2010303001033232-3133332221113330-3321331130121021-0112223030323023-1211012002231122-3120011130233310-1000013203022223"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network`

<a id="canonical-0320002002100211-0320000022000020-3223111021211312-1031232232221131-1231232232101020-2130313110100132-1131030330011012-0013020300001302"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3123332312122310-0331020312300103-0210332230130013-1120310123001102-0202301002200322-1003233131321123-0123222100201111-1123030300112021"></a>

<a id="canonical-1101012013003310-2212023000322021-3310303302303312-3131231213020010-1201031020111322-2200311021001213-3201003101030113-1110313302221101"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2321120232220111-0110200332222033-0123103222102023-0131312031323322-0022212331301001-1313333002300202-3002212312210311-2130303331100103"></a>

<a id="canonical-1333023221313332-2121133122203202-3232223232232230-3020212002110330-1311322220023313-0200020022032102-0030123031211123-3333121011000310"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2333212000313303-2000232320323310-1132310130123202-1021333123010110-2300203221100121-1202232211130120-3321213112310102-1011112111330230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site

<a id="canonical-3123331013130022-3233130113001230-2231022020010003-1301212100030103-2112231313201100-2323230330003103-2223313213113322-2200123202013330"></a>

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

<a id="canonical-0223202011110113-3230312030210323-2021323320121110-2222001301203313-2323211111032100-1210122320112311-2302003223323133-1033222033110030"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site`

<a id="canonical-3132133303020301-2123313103311312-1221330302311101-0210211230201022-3210011232011213-1231302112130112-0220330002311000-2021012203023123"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` property

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

- [virtual_site](resources--bigip_http_proxy--reference--group-003.md#canonical-1222223210211213-1022003231022203-1230103210102110-2011323003212021-3232303322313301-0012230303202332-0302221323220213-2132300222033023): complete subsection reference.

<a id="canonical-1222223210211213-1022003231022203-1230103210102110-2011323003212021-3232303322313301-0012230303202332-0302221323220213-2132300222033023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--bigip_http_proxy--reference--group-003.md#canonical-2333212000313303-2000232320323310-1132310130123202-1021333123010110-2300203221100121-1202232211130120-3321213112310102-1011112111330230)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-3212032320331233-0110203001111003-1213113311213331-0030213020210013-3201211003302002-0133330120022102-3203131031312311-1031003122010022"></a>

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

<a id="canonical-0320312112300223-2121222133203032-3003001031013320-0111330211322331-0031101111302132-0310100110231010-0132012320112003-1200211221011220"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site`

<a id="canonical-2330020322302033-0133011231010320-2101021131302231-0111301120221212-3321012233032031-3132132221203220-0023130202211203-0331133003322230"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2020122000310211-2023032133333312-1001001102013313-0332101210303201-1010200131233233-3003323232312320-2032311230332221-1102011231223012"></a>

<a id="canonical-3222120101231320-0032222210121013-3133330321021002-2333002330123211-2202333120122011-0023023213222210-2023303020203132-3102201013130133"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0323332222331013-0120012112012230-2121323123232122-1221100022330002-2303211203023020-3321133120020230-0210211001311212-2211311303023322"></a>

<a id="canonical-0222232100320032-1321220123230232-1230033333210112-3331123122203230-1003333020021012-1233111323102201-1310003313232202-2121002300321112"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2210222003023302-3322232233330320-1133232011210322-2302301032121031-3332022221001312-0223333301221201-1212020213132130-1210122103301303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-0220300231313210-3332311301321130-3030320113320313-3300310102103031-3003312001331100-2123113302222023-2000113330300331-1312213033101121"></a>

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

<a id="canonical-2210121011321131-2103222200323312-1212130331230221-2131033031202130-2313101031030203-3103101100222310-0222102010123111-3321232230221000"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip`

<a id="canonical-3023233110110223-0112012031033223-3122313322000200-3122102330023132-2330112123322313-0012021000302000-0231330203303003-0310121300110223"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2011101021100212-2333221011113000-2233212112220011-0302312103123032-1220002131021311-2003010303300233-0231033033200223-0211103031323101"></a>

<a id="canonical-3203220301133103-2023230231212301-3313300100102301-2102020303210031-2313202101131221-0213031220333131-2213303020021101-2030230330332121"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` property

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

- [virtual_site](resources--bigip_http_proxy--reference--group-003.md#canonical-3221002021223232-1133123232100002-0310220220322123-0322230130211012-2213130323310111-2022230232312333-3021102101131312-1120313032231133): complete subsection reference.

<a id="canonical-3221002021223232-1133123232100002-0310220220322123-0322230130211012-2213130323310111-2022230232312333-3021102101131312-1120313032231133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--bigip_http_proxy--reference--group-003.md#canonical-2210222003023302-3322232233330320-1133232011210322-2302301032121031-3332022221001312-0223333301221201-1212020213132130-1210122103301303)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-1002001211210232-2213230213230221-2112321102310310-1203230222110311-3133303313213330-2000222130200222-2330233033132232-2002002311111123"></a>

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

<a id="canonical-2131023031330231-0222331131013312-1220323030121102-3332130111221103-3131303033212132-1020120333000030-0132332332231231-1031101301211123"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site`

<a id="canonical-3322203021211320-0032023300311101-2101202110112303-1223302233203113-1031331302200130-3311330303220332-1302111321000000-0303100123112311"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0133212222131103-1101200233122000-0221322320102311-1303012110121020-2230103303020101-0110133322321321-1002100332301023-1112301213312012"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3020111323231331-0022202122030313-0313230112033301-0300021003303233-1131302111303033-2230231121321332-2010023131020102-0321123202212221"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3201012210311222-0013322302303000-0033333010112020-1110200303031321-2310032110213210-2223222313100032-1222203010020321-1021221212131032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service

<a id="canonical-3301003031203220-1330003030032321-0233023302121333-1101120101330111-3003011300032322-3032233230133012-2022332121321220-0020132330021030"></a>

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

<a id="canonical-0023112010011332-3302211322231133-3331033221323130-2131211103311202-2222332010033021-0203223331221211-2330122332010132-2212301302311212"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.vk8s_service`

- [site](resources--bigip_http_proxy--reference--group-003.md#canonical-3032002303300103-3232203103200230-3331001123323011-3131201300332120-2220320123203201-1301232230000001-2113021333321222-2010000031311032): complete subsection reference.

- [virtual_site](resources--bigip_http_proxy--reference--group-003.md#canonical-2201030210312032-3310031013032100-2321330132010000-2312323032101330-3202100123013002-1200320032230213-3121002102010103-1032302201021202): complete subsection reference.

<a id="canonical-3032002303300103-3232203103200230-3331001123323011-3131201300332120-2220320123203201-1301232230000001-2113021333321222-2010000031311032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` properties

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

<a id="canonical-3230001231301113-3301203221010223-3022023212213330-3331320333020133-3000110030220202-2011331220320120-3121002120032010-3213101121013111"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site`

<a id="canonical-3332121322001001-1102133321203012-1332213032333031-2102201023102331-3123101321021122-0211003010021032-1021011023200320-1301112130122321"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2230220022032022-1013320331120302-3222312133232113-0133032320212212-1212301123111102-0032132320121223-2123123023130122-1333020123103221"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0023013212333102-2333110011100213-1132220111331202-0123222222012013-0112210221012333-3022013300220220-1220231113113303-3031123311311221"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2201030210312032-3310031013032100-2321330132010000-2312323032101330-3202100123013002-1200320032230213-3121002102010103-1032302201021202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` properties

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

<a id="canonical-0202230231233123-1012001330233102-2131030021020232-1302100000121020-3020121323301332-2222111022202000-3330220201313001-1030102020002120"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site`

<a id="canonical-2023311033103001-1233202311221312-1310221033223232-0013213213010321-2312321102101120-0121201102100012-0031121210113011-0302200210230322"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0133203300013300-1332103311123031-2131222032212222-2310323201223223-1333212302233220-0122210330030131-3120100103203102-1102220321223311"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2202002221331002-0011020301211332-2213330121212113-0311220320123200-2202303321023202-3203311210121100-2101320002303120-2023122021210211"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1332033121310100-2011203031301030-3100012202102111-0231031031001013-0100201312332130-2311031012211232-0303100202100230-3112212033310122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.do_not_advertise` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- proxy_advertisement.do_not_advertise

<a id="canonical-1020302001020012-1111200022023222-0013230011313120-3011031001012211-3220221112212112-3313110121011001-3330011201311013-3230123323013210"></a>

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

<a id="canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- proxy_config

<a id="canonical-0321032001132330-3220030230232212-1130321320202323-1011112321033102-0312012203223033-1331130030003002-0021310112303213-3332210012102121"></a>

Type: `"object"`. single nested block, Optional.

HTTP/HTTPS Load Balancer. HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3010212002320300-0002323133310113-2121011332201010-3113102030332203-0110221010021223-2211230110111200-0130112030221013-1133223021202030"></a>

### Direct properties for `proxy_config`

<a id="canonical-0323031323301101-1123202220310110-0213023010203111-1211132003303011-3333123002012113-0013302202301133-1033002132222030-1301222000000021"></a>

#### `proxy_config.domains` property

Type: `["list", "string"]`. Optional.

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
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3120122013321203-0002100130200113-1320200332112232-3110210102320231-1333133322131232-0010022221131321-2221331010133313-3113033222022112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.http` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- proxy_config.http

<a id="canonical-2033312303200203-3102110023031002-2221300310210132-2211131003033322-0130333102110000-3122333132321131-3203232331122012-3212113300332313"></a>

Type: `"object"`. single nested block, Optional.

HTTP Choice. Choice for selecting HTTP proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3003210200013012-3013033232303122-2021132213020310-2131133130201201-0033202130022130-1232221000333331-2202312012011002-2032330300212222"></a>

### Direct properties for `proxy_config.http`

<a id="canonical-1130111333031312-1312121322022202-1303200330122021-0020232133222311-2221010102201131-1232333002230100-1030213112103200-1103232223132113"></a>

#### `proxy_config.http.dns_volterra_managed` property

Type: `"bool"`. Optional.

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

<a id="canonical-1232311213031101-0131213230112033-3332302023110231-3022022210223333-1122030020312201-3021000132111312-0012031130133022-3030102131002100"></a>

#### `proxy_config.http.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3303112110233102-3201100333302121-3133233103101320-1303321101031120-0322101130100120-3302130220033113-2330122002323303-2200300210010320"></a>

#### `proxy_config.http.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https` properties

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
EnumExtractionComplete: false
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

<a id="canonical-2233130103303123-1001313130120220-1302000222100033-0322333330123001-0311023033120102-3311221003302232-0202121013223300-0123213023033111"></a>

### Direct properties for `proxy_config.https`

<a id="canonical-0011233003023213-2013302103113212-3200223313232323-1112231002120013-0011123120133300-0230333100230231-1130131101222100-2003103121000020"></a>

#### `proxy_config.https.add_hsts` property

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

<a id="canonical-1203112122111013-3301123301333000-3321020232323110-2000022023101020-0030202201232131-1303213223201223-3300012321331113-3201001101230322"></a>

#### `proxy_config.https.append_server_name` property

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2220101221203011-1023012323130003-1321030213011010-3230030121131112-1002302211031121-2112220313012210-3132102332210222-0002130021222111"></a>

#### `proxy_config.https.connection_idle_timeout` property

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2112103012313300-1323103101330011-3221233112230201-0031320303321303-0212032100211301-0132002112002232-3113202230122222-2023201033212323"></a>

#### `proxy_config.https.http_redirect` property

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-0320201322322010-1200000331203132-3320120132230020-1311212030232102-2302021333031130-2010130300313112-2301200333210032-1222021001012311"></a>

#### `proxy_config.https.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2302230020233202-2100103120301332-2230111231112311-1230323323321030-2013131212332103-3321131332233223-1223100001101012-1111221002012202"></a>

#### `proxy_config.https.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1301213320222020-2002132333230033-2111001130321313-0223303313323020-3320212011101133-0332312323022122-3311110311331010-2011130011212331"></a>

#### `proxy_config.https.server_name` property

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0122023333332003-0230210110313232-2302012011112302-2312312013000100-0332122101232130-1301100301123232-3320133330333302-0122001022012013): complete subsection reference.

<a id="canonical-0220013003110000-1011123121320123-0211100123130311-3133200103222013-0102323300313331-3332301003131101-2003012100230103-3101321311200333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.coalescing_options` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.coalescing_options

<a id="canonical-1123120103122012-3002213212320332-1213303222100133-2313330212013122-2130201213310200-2023213132130031-3212033231211013-3303232011310133"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1021110203310021-1332231331012133-2233131233310002-2333303301022033-3232230031123212-2331222012321301-0212301010302013-0202021131200223"></a>

### Direct properties for `proxy_config.https.coalescing_options`

- [default_coalescing](resources--bigip_http_proxy--reference--group-003.md#canonical-1312103010121233-2211320312212203-3322211300123333-0300311130132213-0332021111032223-0022202002033213-0323030203230320-3023032113212232): complete subsection reference.

- [strict_coalescing](resources--bigip_http_proxy--reference--group-003.md#canonical-2121023322022333-0122312212302220-0132301323100121-0233323321200233-0123003322211030-2203311233131231-1330133230121010-1131133012213112): complete subsection reference.

<a id="canonical-1312103010121233-2211320312212203-3322211300123333-0300311130132213-0332021111032223-0022202002033213-0323030203230320-3023032113212232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.coalescing_options.default_coalescing` properties

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

<a id="canonical-2121023322022333-0122312212302220-0132301323100121-0233323321200233-0123003322211030-2203311233131231-1330133230121010-1131133012213112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.coalescing_options.strict_coalescing` properties

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

<a id="canonical-2113133303201003-2332010203301203-0101132133223020-2103011222121331-3331032000330032-1122031221022100-1320003323032023-2032210102100232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.default_header` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.default_header

<a id="canonical-2333033300201021-1330001010030110-3013123320212003-2231113231011030-3020112302212202-1123103312231003-1203013012330222-1320301312330110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default header.

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
default_header = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100001333201312-1023331023100332-0212202323021102-1101331122232322-0231020133013033-3313332200103113-0010022200133210-3030203301323022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.default_loadbalancer

<a id="canonical-0311203212033031-3312023023011121-0320122331211221-1111313321230321-0000211133211113-2102031332230220-1011301333220202-1233110311011322"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default loadbalancer.

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
default_loadbalancer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021220202312122-1322031221010112-0102301300133130-2021113023103020-3112011122321213-0132002221323021-0120121133211323-0132013313102131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.disable_path_normalize

<a id="canonical-3223230200123312-3320312110313232-3113113123100232-2330300313111023-0223001001200313-2013300220100200-1330312303221131-0033232210120210"></a>

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

<a id="canonical-1202213120122303-3301230233233302-1112111001003131-3220033321332232-1303003103222301-3223033131003333-1231200112123320-1300123301103120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.enable_path_normalize

<a id="canonical-3002223111210111-3303032202311333-2003003230033302-0300003121103233-1321000030203010-1022233010011000-1101211102211102-1223322100200222"></a>

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

<a id="canonical-1200103230101333-2113021032322120-3323321233020213-3302012133221032-2222302001101103-1013231202302013-3031233132323311-3300212010030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options` properties

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
EnumExtractionComplete: false
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

<a id="canonical-3203312220323030-2102113331323223-2123322220302103-0122112011100110-1122000203301032-0030230203002220-2032213002010222-1020331300102211"></a>

### Direct properties for `proxy_config.https.http_protocol_options`

- [http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-0313013020312011-2221103131030112-0220030323120120-0110023302321311-1222312200331132-1201132212330331-2232221200202231-1223233021132232): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--bigip_http_proxy--reference--group-003.md#canonical-0210030223123132-2223301031131002-0012120003033233-1001211032112301-2302300030200311-0010200230211001-2320333023331210-0313312310021121): complete subsection reference.

- [http_protocol_enable_v2_only](resources--bigip_http_proxy--reference--group-003.md#canonical-3111003033233120-2303101110200112-0001132101323022-2310111232301102-0210131222023233-1012002332220302-0010103003323133-1320201033200231): complete subsection reference.

<a id="canonical-0313013020312011-2221103131030112-0220030323120120-0110023302321311-1222312200331132-1201132212330331-2232221200202231-1223233021132232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only` properties

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

<a id="canonical-3322011332313113-3022301222110301-2301322223010123-3320132321120200-2101001002030121-0033210101311011-1312022323002231-3110110312101202"></a>

### Direct properties for `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-2232331113101101-1303111213011013-2133032203200333-2113230322210220-2102030302011300-3312003131232323-2010132203213013-3302123331122130): complete subsection reference.

<a id="canonical-2232331113101101-1303111213011013-2133032203200333-2113230322210220-2102030302011300-3312003131232323-2010132203213013-3302123331122130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

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
EnumExtractionComplete: false
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

<a id="canonical-2310001103331331-2010110111123022-0213120111030011-0022312302303013-1313032322010321-1012123002201112-1303201323312212-0113120011030130"></a>

### Direct properties for `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-3030022221012231-1301333130221002-2033210310110333-1110133302020231-0311233323201002-1022302033132330-1212223111033022-0113221120230231): complete subsection reference.

- [preserve_case_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-1113112123230132-0030023231003131-1002312111103200-1201310020331322-3331320023222121-1023233100000022-3303112021101212-0001131011330003): complete subsection reference.

- [proper_case_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-1223120320313001-3313200202233210-2022130230120103-3123133211022313-3212303101131131-1033023101212103-1113323033202101-2031121100101303): complete subsection reference.

<a id="canonical-3030022221012231-1301333130221002-2033210310110333-1110133302020231-0311233323201002-1022302033132330-1212223111033022-0113221120230231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113112123230132-0030023231003131-1002312111103200-1201310020331322-3331320023222121-1023233100000022-3303112021101212-0001131011330003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223120320313001-3313200202233210-2022130230120103-3123133211022313-3212303101131131-1033023101212103-1113323033202101-2031121100101303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210030223123132-2223301031131002-0012120003033233-1001211032112301-2302300030200311-0010200230211001-2320333023331210-0313312310021121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2` properties

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
http_protocol_enable_v1_v2 = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111003033233120-2303101110200112-0001132101323022-2310111232301102-0210131222023233-1012002332220302-0010103003323133-1320201033200231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options.http_protocol_enable_v2_only` properties

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
http_protocol_enable_v2_only = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122020120213001-1002322011131121-1002302020013313-1333031232020323-3221231022002002-2330001003133021-3333221020001033-3220323333010003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.non_default_loadbalancer

<a id="canonical-3121011321013220-0322322003203110-0113320333110021-3212330003310130-0203122002230033-0021210220203030-0121110333330103-3300020120032122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for non default loadbalancer.

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
non_default_loadbalancer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2310222311302101-2121110111210313-3033122233131333-3320311202312031-1102212112100311-2120323112302121-3313233100202102-3132232232313210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.pass_through` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.pass_through

<a id="canonical-2100303113221032-3211313031330102-1010121010103101-3302131111021313-1302300103212302-2101103000311001-0223310330002100-0301201121003131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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
pass_through = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123311201031311-3201123311021011-2113321033022200-1320122102112130-3203011231331232-0100103211121032-1321012122113211-2213331312220102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033)
- [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111)
- proxy_config.https.tls_cert_params

<a id="canonical-0233332000132102-1313202110002220-1210132103232310-0011130031230122-2100211121001120-0330212132020020-2331202011301010-0110310320220330"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Additional upstream details:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0101333103320331-3011133101031300-3221301201323210-0003010012313102-1322320202133023-0300031122310012-2133031030203212-3203133132132102"></a>

### Direct properties for `proxy_config.https.tls_cert_params`

- [certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-1013202012320033-1110130002111103-2003230201332002-3003000103011220-3103220201133133-2310231333132013-1330012122312111-0032201130223000): complete subsection reference.

- [no_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-3023033212320010-1200320302302111-3310101022321311-2102311011100020-0201320111112200-3331113213211131-2031222302101333-2133101233032223): complete subsection reference.

- [tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-3312301111313113-0202213300013222-2003020101033232-3302331231131213-3311303312020033-3211313102120301-1033030023232130-2331210133302320): complete subsection reference.

- [use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310): complete subsection reference.

<a id="canonical-1013202012320033-1110130002111103-2003230201332002-3003000103011220-3103220201133133-2310231333132013-1330012122312111-0032201130223000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.certificates` properties

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
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1211310121000321-0200133030230213-0203301203232021-1231303102201311-2102102002203223-0101310211102303-1223303100310212-2331113323213113"></a>

### Direct properties for `proxy_config.https.tls_cert_params.certificates`

<a id="canonical-3020130213211231-0110022032110102-0103110201032332-3130011011102112-0000032222200003-0123132112301022-0321313310003111-2101002230132021"></a>

#### `proxy_config.https.tls_cert_params.certificates.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
