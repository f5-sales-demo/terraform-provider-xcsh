---
page_title: "xcsh_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery reference."
---

# xcsh_discovery reference

<a id="canonical-0330103230210002-2023131122013200-3000313131202303-3332020130030231-3212123120231201-3022003312202221-2202130203022010-3120213322311130"></a>

## name property — ref / 013312022032 / 5

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

<a id="canonical-2121022210310020-2010032232022003-3003333101133323-1002203312000231-3232312022221021-1020333301101231-3110112210120212-0003210012112232"></a>

<a id="canonical-3001021000223021-1121113220222231-1300113323123110-1321211202120312-3120213012130021-2111002001131322-2102333203333221-0112131030130300"></a>

## namespace property — ref / 013312022032 / 6

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

<a id="canonical-0122032130001322-1231202123102221-3121021003121013-1010113003330231-0231211122130013-1223130121002202-2201101230210001-3201023333311010"></a>

<a id="canonical-3330130003021122-2313322233022002-3000303031002320-2023203131212320-1221102332011011-3113030102201310-3202311221010202-0130321130121320"></a>

## tenant property — ref / 013312022032 / 7

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

<a id="canonical-2303322321221301-3110012132223100-1333011120031230-1031222221000231-3212331333012313-0331033003122111-1200313121011202-0232212310312122"></a>

<a id="canonical-0311030321103010-3231321211131101-0132000032030230-3110023100323101-1122311233103231-1002201133102000-3113133121021102-0312001311303100"></a>

## uid property — ref / 013312022032 / 8

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

<a id="canonical-0302231033102111-0111220332303113-3113332002003332-3120113122121233-1001303010031132-1111210031202203-1230302331021023-0102003222311201"></a>

## Next pages — ref / 013312022032 / 9

- [where.virtual_network](data-sources--discovery--reference--group-001.md#canonical-2000100203122133-0301232002323310-3233310132013200-2022221330110313-1210122133103010-3121123023112123-0033000321111322-1210332230303323)
- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)

<a id="canonical-2133223112022020-1203131112101120-1123221132012123-1012101130010321-2213300002121320-2233022032323113-3111330010231222-0322222111031122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233032212031121-2312310231023322-3133221322310012-3021132001032230-0021030133022310-3230012002121203-0213012222100223-0112103303011213"></a>

## where.virtual_site — virtual_site / 201120211112 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [where](data-sources--discovery--reference--group-001.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- where.virtual_site

<a id="canonical-3121122332022100-2322022213221212-0030221202020033-0013010231331211-3232021320122012-3112321021330201-0311211221311133-0333021120321313"></a>

Type: `"single"`. Computed.

Virtual Site. A reference to virtual\_site object.

Upstream description:

A reference to virtual\_site object.

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

<a id="canonical-2122313100301200-0130300323013330-0102203110012001-1231012023200023-0120003210331312-1201210111002321-0012121230231333-3103020123000310"></a>

## Direct properties — virtual_site / 201120211112 / 3

- [disable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-2310001332100202-0111030221230112-1322023213221100-2012333111000330-1232321320031301-0131300201321203-2321211122012232-1333223332300122): complete subsection reference.

- [enable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-2020131222103111-2223331232232312-3212101302212223-1200010333003320-3310010232102320-3303103331031332-1000011103202330-2011022321023222): complete subsection reference.

<a id="canonical-3320020330211301-1102122020300322-2013331030233300-3113002012313023-0313002001233231-3100203330121111-1331013312002022-0000232011033021"></a>

<a id="canonical-0302033111213203-2331230202220113-0013213123010031-1033210311321310-0323110201100130-0000310203333330-3303110112222003-2321331221301030"></a>

## network_type property — virtual_site / 201120211112 / 4

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

- [ref](data-sources--discovery--reference--group-002.md#canonical-2112101231113022-3113100010022100-2233130012103200-0011131222221220-1320202012310013-3232220012222010-1222300020330213-3022121113131011): complete subsection reference.

<a id="canonical-0010320222031222-2222022220120230-0232221333210221-2312323032131312-1022122132212303-3330121011220122-1013122323021300-3100011133010010"></a>

## Next pages — virtual_site / 201120211112 / 5

- [where.virtual_site.disable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-2310001332100202-0111030221230112-1322023213221100-2012333111000330-1232321320031301-0131300201321203-2321211122012232-1333223332300122)
- [where.virtual_site.enable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-2020131222103111-2223331232232312-3212101302212223-1200010333003320-3310010232102320-3303103331031332-1000011103202330-2011022321023222)
- [where.virtual_site.ref](data-sources--discovery--reference--group-002.md#canonical-2112101231113022-3113100010022100-2233130012103200-0011131222221220-1320202012310013-3232220012222010-1222300020330213-3022121113131011)
- [where](data-sources--discovery--reference--group-001.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)

<a id="canonical-2310001332100202-0111030221230112-1322023213221100-2012333111000330-1232321320031301-0131300201321203-2321211122012232-1333223332300122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333000333311003-2221001311011032-2322133213321321-1222031002320123-0031100302233021-2023233030230133-1302312022002110-2132233013310113"></a>

## where.virtual_site.disable_internet_vip — disable_internet_vip / 312032033033 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [where](data-sources--discovery--reference--group-001.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- [where.virtual_site](data-sources--discovery--reference--group-002.md#canonical-2133223112022020-1203131112101120-1123221132012123-1012101130010321-2213300002121320-2233022032323113-3111330010231222-0322222111031122)
- where.virtual_site.disable_internet_vip

<a id="canonical-3121010010312310-2012012330223202-0133303311112213-1300211123033101-3112301131231222-1200110102021221-3123330030130012-0323213030231012"></a>

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

<a id="canonical-3112313013330323-1112213021020133-2230201031332032-2121030103110313-2213201012301320-1313101001231230-3100102021311223-1021122022312010"></a>

## Direct properties — disable_internet_vip / 312032033033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231322013201132-3210002313120033-2132110002300330-3000121210221203-2123333310222031-0112230103301012-0322003310121233-1302203301331320"></a>

## Next pages — disable_internet_vip / 312032033033 / 4

- [where.virtual_site](data-sources--discovery--reference--group-002.md#canonical-2133223112022020-1203131112101120-1123221132012123-1012101130010321-2213300002121320-2233022032323113-3111330010231222-0322222111031122)
- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)

<a id="canonical-2020131222103111-2223331232232312-3212101302212223-1200010333003320-3310010232102320-3303103331031332-1000011103202330-2011022321023222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202203200213021-1202100331023033-1230302133110321-2132122131120202-1111030311311031-2002202132222213-1023032103130211-3223300221000012"></a>

## where.virtual_site.enable_internet_vip — enable_internet_vip / 111022130233 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [where](data-sources--discovery--reference--group-001.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- [where.virtual_site](data-sources--discovery--reference--group-002.md#canonical-2133223112022020-1203131112101120-1123221132012123-1012101130010321-2213300002121320-2233022032323113-3111330010231222-0322222111031122)
- where.virtual_site.enable_internet_vip

<a id="canonical-3011122123133320-1303211013132313-1112123011003222-0003232113101030-3110122102333102-1322203212322021-1021003202331110-3013333001212003"></a>

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

<a id="canonical-2323233002323100-3330013122221003-3133200130303302-1323333021112310-2010010121322222-3301313231310122-0301021020130001-2122233120210101"></a>

## Direct properties — enable_internet_vip / 111022130233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231102230300331-3311210331020302-0231022200332211-0211021321002020-2112221310300322-1332033030231213-3231111300333002-3310231320301023"></a>

## Next pages — enable_internet_vip / 111022130233 / 4

- [where.virtual_site](data-sources--discovery--reference--group-002.md#canonical-2133223112022020-1203131112101120-1123221132012123-1012101130010321-2213300002121320-2233022032323113-3111330010231222-0322222111031122)
- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)

<a id="canonical-2112101231113022-3113100010022100-2233130012103200-0011131222221220-1320202012310013-3232220012222010-1222300020330213-3022121113131011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001330203223203-3001022021322213-1013232321311303-1200220100010130-1221233323111001-3002001231020222-2102200211202010-3313320221202301"></a>

## where.virtual_site.ref — ref / 131002202133 / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [where](data-sources--discovery--reference--group-001.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- [where.virtual_site](data-sources--discovery--reference--group-002.md#canonical-2133223112022020-1203131112101120-1123221132012123-1012101130010321-2213300002121320-2233022032323113-3111330010231222-0322222111031122)
- where.virtual_site.ref

<a id="canonical-0003002230031332-3213000322300231-0312212202233210-1002020331313320-2002103231103312-2201201210233320-0032021202013310-3001311232330013"></a>

Type: `"list"`. Computed.

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

<a id="canonical-2003131210022320-0103031120202002-3113112123202003-1312320021331003-3213102311102033-2021133223323032-0203333022101113-3202233332002211"></a>

## Direct properties — ref / 131002202133 / 3

<a id="canonical-1230000133212313-1302201130321013-0300211302313300-3103213201302210-2101020230310013-0223031133213113-3320202002211122-3131330010330301"></a>

<a id="canonical-0013031212112310-0030133023131203-1232002333222021-3013213031123022-0101131310221331-3220311100020202-1200130133201213-0313133120020200"></a>

## kind property — ref / 131002202133 / 4

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

<a id="canonical-2121102110133313-3103303223012200-3012220100312133-1311300020001322-3113312213223103-1301113333120331-3303113330100031-1311031232233200"></a>

<a id="canonical-3131103003010030-2211322213103301-0233322231213102-0001013112101321-3130213211012203-2031100303303322-0001130230202200-3120231331320133"></a>

## name property — ref / 131002202133 / 5

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

<a id="canonical-0312223301202212-2200221213013001-0230121133112010-0313300000303022-1332232122113310-1102020233010220-2111030231100022-2211212130213320"></a>

<a id="canonical-1333113210230112-0130212231331303-0032330220321132-0200102013321232-3113200311303213-1212113110023200-3133022321331100-2211332302012211"></a>

## namespace property — ref / 131002202133 / 6

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

<a id="canonical-2320113201012322-3011103012123211-1121033302223312-0032222302222130-3012132013310330-0020323320333332-2010230311000323-2303301212020011"></a>

<a id="canonical-0001101030103212-1233111102320111-2200322023110011-2211233100332001-1130033320010030-0330130201130032-0122112003121201-3112103223303102"></a>

## tenant property — ref / 131002202133 / 7

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

<a id="canonical-3212130122302310-0120202023312212-3301200101013110-3102303122203200-3003002201111120-2001020013020220-3032223301302123-3101011122011122"></a>

<a id="canonical-0331103302220030-1203013110311313-0310233132132131-0321232013311131-2300033303312123-2310113101323112-1333220223113021-1110032132222021"></a>

## uid property — ref / 131002202133 / 8

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

<a id="canonical-0120232231120123-0030212323030032-1322303322022112-1311023222120301-0231003321122320-0223232300133233-2321011010222230-2231311332113320"></a>

## Next pages — ref / 131002202133 / 9

- [where.virtual_site](data-sources--discovery--reference--group-002.md#canonical-2133223112022020-1203131112101120-1123221132012123-1012101130010321-2213300002121320-2233022032323113-3111330010231222-0322222111031122)
- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
