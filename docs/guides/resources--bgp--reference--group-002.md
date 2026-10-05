---
page_title: "xcsh_bgp reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp reference."
---

# xcsh_bgp reference

<a id="canonical-1322010310120220-2322300123212001-2130121222320210-1002120132021002-3013333202222003-0010222133110202-0331110030020311-1121022033223112"></a>

## network_type property — virtual_site / 331330023202 / 4

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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIRTUAL_NETWORK_GLOBAL","VIRTUAL_NETWORK_IP_AUTO","VIRTUAL_NETWORK_IP_FABRIC","VIRTUAL_NETWORK_MANAGEMENT","VIRTUAL_NETWORK_PER_SITE","VIRTUAL_NETWORK_PUBLIC","VIRTUAL_NETWORK_SEGMENT","VIRTUAL_NETWORK_SITE_LOCAL","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE","VIRTUAL_NETWORK_SITE_SERVICE","VIRTUAL_NETWORK_SRV6_NETWORK","VIRTUAL_NETWORK_VER_INTERNAL","VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [ref](resources--bgp--reference--group-002.md#canonical-2110210021323002-1032333330101203-0210010203103202-0311133000022303-0300003012201032-0202230323011002-1322313223002123-2101112131103311): complete subsection reference.

<a id="canonical-0011003332311203-2030012332020301-2031302021332303-2021212322313201-3110010210013333-1300301133003300-2310101332330203-3020332200021032"></a>

## Next pages — virtual_site / 331330023202 / 5

- [where.virtual_site.disable_internet_vip](resources--bgp--reference--group-002.md#canonical-2001312332332021-1211103132101230-1211213112201110-0333023221322020-3331321010220030-2211022013211130-0220112223203322-0320003031312021)
- [where.virtual_site.enable_internet_vip](resources--bgp--reference--group-002.md#canonical-3112200331130322-3302030033020110-3102222103210303-2023131011102130-3110310002210231-1300212211231023-0330331201321103-2110110102133123)
- [where.virtual_site.ref](resources--bgp--reference--group-002.md#canonical-2110210021323002-1032333330101203-0210010203103202-0311133000022303-0300003012201032-0202230323011002-1322313223002123-2101112131103311)
- [where](resources--bgp--reference--group-001.md#canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231)
- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)

<a id="canonical-2001312332332021-1211103132101230-1211213112201110-0333023221322020-3331321010220030-2211022013211130-0220112223203322-0320003031312021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012233300131311-3333133032013112-2122313112101111-3121320333321210-0302120202223012-2201213233011202-2121310300331023-3202032123212213"></a>

## where.virtual_site.disable_internet_vip — disable_internet_vip / 002211211221 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [where](resources--bgp--reference--group-001.md#canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231)
- [where.virtual_site](resources--bgp--reference--group-001.md#canonical-0231022200202313-2033330230223132-0302221102233021-2022123222113212-2201231331231000-0221210311023333-0121201210301213-2231101310213200)
- where.virtual_site.disable_internet_vip

<a id="canonical-0300131333311032-0222313321100222-3221230201010202-1003022200310020-0022010312020023-0331303101133331-3123213210333210-3201130230330330"></a>

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

<a id="canonical-1112311113310123-1300103200301110-3133112320131121-2032222100200221-0313132300310001-2012122330201332-1212233122213002-1222311103221311"></a>

## Direct properties — disable_internet_vip / 002211211221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332021232130012-0220201033203211-3221120210013101-3332122333211202-3233310011320003-2000130331102132-2201000102012132-1302220112020100"></a>

## Next pages — disable_internet_vip / 002211211221 / 4

- [where.virtual_site](resources--bgp--reference--group-001.md#canonical-0231022200202313-2033330230223132-0302221102233021-2022123222113212-2201231331231000-0221210311023333-0121201210301213-2231101310213200)
- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)

<a id="canonical-3112200331130322-3302030033020110-3102222103210303-2023131011102130-3110310002210231-1300212211231023-0330331201321103-2110110102133123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311010302120222-2303301223002210-1012303202020013-0120232130010123-2220122210112213-2200031002011122-0231123220100000-1212202312113211"></a>

## where.virtual_site.enable_internet_vip — enable_internet_vip / 233303231111 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [where](resources--bgp--reference--group-001.md#canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231)
- [where.virtual_site](resources--bgp--reference--group-001.md#canonical-0231022200202313-2033330230223132-0302221102233021-2022123222113212-2201231331231000-0221210311023333-0121201210301213-2231101310213200)
- where.virtual_site.enable_internet_vip

<a id="canonical-3003323211221211-0232122113322013-2220301222332123-1121311313220012-1033303201311333-1330130101023011-3122132312022100-2220231213322111"></a>

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

<a id="canonical-1023323031030113-3333202103023210-1322121000031233-3221110000101223-1331232002311022-3022210000022113-2230132013132111-0032030232210231"></a>

## Direct properties — enable_internet_vip / 233303231111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131032021213021-2301123203010230-2203030320123031-3310313113121322-0112332232321211-0211213231302020-3122023020123200-2120113222302333"></a>

## Next pages — enable_internet_vip / 233303231111 / 4

- [where.virtual_site](resources--bgp--reference--group-001.md#canonical-0231022200202313-2033330230223132-0302221102233021-2022123222113212-2201231331231000-0221210311023333-0121201210301213-2231101310213200)
- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)

<a id="canonical-2110210021323002-1032333330101203-0210010203103202-0311133000022303-0300003012201032-0202230323011002-1322313223002123-2101112131103311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331232101223313-2011201220011201-2012030000300101-0010031233002331-1232303212300313-1032111012330111-0022223301301021-3102113003011313"></a>

## where.virtual_site.ref — ref / 202233123201 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [where](resources--bgp--reference--group-001.md#canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231)
- [where.virtual_site](resources--bgp--reference--group-001.md#canonical-0231022200202313-2033330230223132-0302221102233021-2022123222113212-2201231331231000-0221210311023333-0121201210301213-2231101310213200)
- where.virtual_site.ref

<a id="canonical-0102011002200311-0011100111101313-2313132101001233-0331122211313201-2132120201311322-3102033101132231-2132102022023013-1223102030322302"></a>

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

<a id="canonical-0220213313130130-0032011323200000-3123232012220033-1303322221133033-2130222122202001-3031202111030201-3002112020311233-0231230021201012"></a>

## Direct properties — ref / 202233123201 / 3

<a id="canonical-0032033333233030-2131302231320222-2003130333112330-1203210132201301-2011133212311313-1023111001111032-1020012120220001-2022103030000230"></a>

<a id="canonical-2210331111121222-0111122210023232-0122110102202001-2310210200130322-1303231030322012-3303331310113110-3312322231131332-2201121033201212"></a>

## kind property — ref / 202233123201 / 4

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

<a id="canonical-2110132321123303-1002131313200122-0200200333011121-1330001022021312-2110221012111332-3001101223012032-3210030012111203-0323031033120030"></a>

<a id="canonical-2031112323301321-1021222223133322-3023233232211312-0301300003212112-3030313320121301-0011313202232003-0101300011123313-2203223220311023"></a>

## name property — ref / 202233123201 / 5

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

<a id="canonical-3032301222230330-1203133332231311-2332313130213303-3210102002210120-1313013221312100-0021001203323312-3322221230021222-2111102101313010"></a>

<a id="canonical-0112130002112122-3313221212213001-2132310303010312-2000231023232011-2010231103011232-1230033230111003-0300101222232331-0223232012132031"></a>

## namespace property — ref / 202233123201 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0133120002032030-2133313313320012-3312233100321231-0102201113022123-2222223322221030-2203100021113101-2300222123323111-2031003200200200"></a>

<a id="canonical-1123223002030231-1302322122320101-3111113013330330-1113112220230103-2113021332203122-3201221312030120-1110113333233023-0231003112013121"></a>

## tenant property — ref / 202233123201 / 7

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

<a id="canonical-3101100121323210-3021302301200202-1122032321131321-3222132031031123-1131102123030131-2033013230032131-1021110103302222-3103322230113331"></a>

<a id="canonical-0212312130320112-0111201220320010-1010220323312211-3233213113113102-2113233201112133-0020102220320300-2220203132211110-0203011002303321"></a>

## uid property — ref / 202233123201 / 8

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

<a id="canonical-2312120221330310-2333112023113021-1101310322020311-0012001132030230-3223301212312323-0211102303313211-2320331211103312-2230012112013102"></a>

## Next pages — ref / 202233123201 / 9

- [where.virtual_site](resources--bgp--reference--group-001.md#canonical-0231022200202313-2033330230223132-0302221102233021-2022123222113212-2201231331231000-0221210311023333-0121201210301213-2231101310213200)
- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
