---
page_title: "xcsh_secret_management_access reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access reference."
---

# xcsh_secret_management_access reference

<a id="canonical-1313312212212220-1131131232320221-3202033013012233-2322101003330200-0201230310303321-0023002110113211-1221120012012311-1220000103222212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-001.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- where.virtual_site

<a id="canonical-2001033311032022-3000322323110031-3221011302233023-1300031320330201-0112132010323310-2221301210010212-2301013033120303-2032321220111202"></a>

Type: `"single"`. Computed.

Virtual Site. A reference to virtual\_site object.

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

<a id="canonical-1321001113133033-1112102222030233-1032202013301230-3320013220212001-3000312130312213-1210100012103002-1201221320220100-1332003201132222"></a>

### Direct properties for `where.virtual_site`

- [disable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-2202023202011210-2132033110202311-0001122031213032-3323000030010202-3213121322211221-0000213302332000-0021313210223312-2032331032311300): complete subsection reference.

- [enable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-0231020332010130-0212221121331013-2113132202133330-3032233032130112-3322131230211113-0330102332002200-1230031220300210-3233103212332300): complete subsection reference.

<a id="canonical-1333032023122311-3110002120230000-0200021213113010-3121030313210303-2002002201001113-2301010232013211-2101322221203110-1323303303200021"></a>

<a id="canonical-2233323121123332-3122111211203003-1112130033103302-3213030321120311-1311021033310001-3022230013311122-3133000111120131-3231332201330033"></a>

#### `where.virtual_site.network_type` property

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

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

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

- [ref](data-sources--secret_management_access--reference--group-002.md#canonical-1113100123233122-3321100023011022-3203000032000302-1320110201021112-2130013230100010-0213121312010030-0102200202033232-3333202000022121): complete subsection reference.

<a id="canonical-2202023202011210-2132033110202311-0001122031213032-3323000030010202-3213121322211221-0000213302332000-0021313210223312-2032331032311300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-001.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-1313312212212220-1131131232320221-3202033013012233-2322101003330200-0201230310303321-0023002110113211-1221120012012311-1220000103222212)
- where.virtual_site.disable_internet_vip

<a id="canonical-1132313222220030-3302030131103023-0000322313002232-0221120003000322-2011313122022312-0333303011102330-0132121301002002-0223321212101203"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231020332010130-0212221121331013-2113132202133330-3032233032130112-3322131230211113-0330102332002200-1230031220300210-3233103212332300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-001.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-1313312212212220-1131131232320221-3202033013012233-2322101003330200-0201230310303321-0023002110113211-1221120012012311-1220000103222212)
- where.virtual_site.enable_internet_vip

<a id="canonical-1120002302323123-0012311313203013-0200032113300311-2030232003001321-3312010021210023-0201002020323002-0113303130112323-0301221133302001"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113100123233122-3321100023011022-3203000032000302-1320110201021112-2130013230100010-0213121312010030-0102200202033232-3333202000022121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.ref` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-001.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-1313312212212220-1131131232320221-3202033013012233-2322101003330200-0201230310303321-0023002110113211-1221120012012311-1220000103222212)
- where.virtual_site.ref

<a id="canonical-1232031000220131-0320210023132201-0233211020113323-3032000231123321-0020213122020321-2221002211321031-0031213203220013-3000022201231213"></a>

Type: `"list"`. Computed.

Reference. A virtual\_site direct reference.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2203111110032231-2023333113003221-1112103123020222-2311301011010302-0013310110031231-3330010223320222-0011120100333320-0330130230131000"></a>

### Direct properties for `where.virtual_site.ref`

<a id="canonical-2202231031301311-3112211023032113-2302032230320220-0232123232133030-3313130222220100-3313301121023100-2113122112232221-0211030101330320"></a>

#### `where.virtual_site.ref.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3232232030013323-0203121131202312-2322212303200312-3131010003100030-3213023132001000-1003301212231222-1211231122330020-2012012113233330"></a>

<a id="canonical-2212020133200212-3031200122100122-0313110311213000-3202123220310313-1123013033120211-2011213013102321-3002220022202130-0132033100213202"></a>

#### `where.virtual_site.ref.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1021230013230332-2210123030202320-0233211200001322-3301113201322201-0130231201123001-1211120132102222-2231003301111031-2313310131000020"></a>

<a id="canonical-3020012133020322-0310210213021022-2213313133000123-3320022200231133-3010310010332322-3220332323333031-1321302211020101-3213131203210001"></a>

#### `where.virtual_site.ref.namespace` property

Type: `"string"`. Computed.

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
  }
}
```

<a id="canonical-0220112000000022-1113232310030133-0333012222320322-3020331312311010-0123103320331032-0223231323123300-3322232031311202-1030330001001103"></a>

<a id="canonical-1022103310220132-1033021023110202-1020222220232112-0123333231210330-2003001231233200-0020221211211122-2320000132233001-1301102221321300"></a>

#### `where.virtual_site.ref.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3311022013110010-2133332013330030-2020231331130110-2233232200330123-3131303133133111-0033013113332232-0302321021000013-1023330330013103"></a>

<a id="canonical-2321033230330132-3221230213211331-2223213212010000-0120101122311322-3112102301100011-1110032313102311-3021001301102321-3010300231021220"></a>

#### `where.virtual_site.ref.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
