---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-3323313313333101-1022230003023322-1121321233003231-3321002002120031-3231313020123102-0332221122301230-2301200331132022-2132211330222112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_path_normalize` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- disable_path_normalize

<a id="canonical-3301102111312121-0031222233000230-3002223322002213-3111133201103210-3210213120101312-2112223132033332-1123130311310222-0210010110330230"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_path\_normalize, enable\_path\_normalize; Default: disable\_path\_normalize\]
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

OneOf alternatives in this subsection:

- [disable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-3301102111312121-0031222233000230-3002223322002213-3111133201103210-3210213120101312-2112223132033332-1123130311310222-0210010110330230)
- [enable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-0201022300021320-1303303003222003-0131032130300103-2101300200200320-1102022211110023-3002001212313220-0232120011002133-0313322031310312)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320322030133330-2313030201332123-3133200313313203-0300032102221333-1001120222131112-1330331333122030-2312013001023021-0010002323330332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_reverse_proxy` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- dynamic_reverse_proxy

<a id="canonical-3312001003220201-1021111332030323-2320100230313010-0231102023313230-3231131000121133-2230033123002221-1131102320231311-0211223000310232"></a>

Type: `"single"`. Computed.

In this mode of proxy, virtual host will resolve the destination endpoint dynamically.

The dynamic resolution is done using a predefined field in the request. This predefined field
depends on the ProxyType configured on the Virtual Host.

For HTTP traffic, i.e. With ProxyType as HTTP\_PROXY or HTTPS\_PROXY, virtual host will use the
"HOST" HTTP header from the request and perform DNS resolution to select destination endpoint.

For TCP traffic with SNI, (If the ProxyType is TCP\_PROXY\_WITH\_SNI), virtual host will perform DNS
resolution using the SNI.

The DNS resolution is performed in the virtual network specified in outside\_network\_type or
outside\_network

In both modes of operation(either using Host header or SNI), the DNS resolution could return
multiple addresses. First IPv4 address from such returned list is used as endpoint for the request.
The DNS response is cached for 60s by default.

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

<a id="canonical-1321312220100011-2310312333000123-2030121003101222-0303213020033222-3011223310311222-1030030301110330-2032101100233113-1030021122121200"></a>

### Direct properties for `dynamic_reverse_proxy`

<a id="canonical-3211220132132303-3112133120120133-3202311011213200-1002000110020003-2320230021021132-1310002121231023-1132132011202020-0331230332012130"></a>

#### `dynamic_reverse_proxy.connection_timeout` property

Type: `"number"`. Computed.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Additional upstream details:

The default value is 2000 (2 seconds)

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

- [resolution_network](data-sources--virtual_host--reference--group-002.md#canonical-3300300230003011-1333312102330113-0310322301112003-1300032321101023-2110132310122301-1033000103331012-2033333123221333-2230311031113123): complete subsection reference.

<a id="canonical-3010130003111302-2130210213012133-2133122313331220-2021130112120331-0021231020013123-3121011111201010-2203011113302120-0333322103130113"></a>

<a id="canonical-3023031220313310-0032232222202113-0200301231333032-1130330110220331-2123131012022012-0100331200302221-0022111201100133-1013213333202101"></a>

#### `dynamic_reverse_proxy.resolution_network_type` property

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

<a id="canonical-2111003331333012-0102002000130221-1313311222331131-3203333123223321-1212303201102123-2220132000103222-3130113303221032-1200202002000221"></a>

<a id="canonical-1230020031330311-0103103110123132-1331131312113312-1202333312023322-3202220313130112-3233213111203222-3102000023122023-1330230003322100"></a>

#### `dynamic_reverse_proxy.resolve_endpoint_dynamically` property

Type: `"bool"`. Computed.

X-example : true In this mode of proxy, virtual host will resolve the destination endpoint
dynamically.

The dynamic resolution is done using a predefined field in the request. This predefined field
depends on the ProxyType configured on the Virtual Host.

For HTTP traffic, i.e. With ProxyType as HTTP\_PROXY or HTTPS\_PROXY, virtual host will use the
"HOST" HTTP header from the request and perform DNS resolution to select destination endpoint.

For TCP traffic with SNI, (If the ProxyType is TCP\_PROXY\_WITH\_SNI), virtual host will perform DNS
resolution using the SNI.

The DNS resolution is performed in the virtual network specified in outside\_network\_type or
outside\_network

In both modes of operation(either using Host header or SNI), the DNS resolution could return
multiple addresses. First IPv4 address from such returned list is used as endpoint for the request.
The DNS response is cached for 60s by default.

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

<a id="canonical-3300300230003011-1333312102330113-0310322301112003-1300032321101023-2110132310122301-1033000103331012-2033333123221333-2230311031113123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_reverse_proxy.resolution_network` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [dynamic_reverse_proxy](data-sources--virtual_host--reference--group-002.md#canonical-2320322030133330-2313030201332123-3133200313313203-0300032102221333-1001120222131112-1330331333122030-2312013001023021-0010002323330332)
- dynamic_reverse_proxy.resolution_network

<a id="canonical-3033013100002213-2212012332232022-0101032322103101-3203023313232302-2003311011303113-2011101113132023-1103120111033031-0211210301122103"></a>

Type: `"list"`. Computed.

Reference to virtual network where the endpoint is resolved. Reference is valid only when the
network type is VIRTUAL\_NETWORK\_PER\_SITE or VIRTUAL\_NETWORK\_GLOBAL. It is ignored for all other
network types.

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

<a id="canonical-2222101212101033-3320211102323202-0311110123120312-2121020232022031-1313100323021202-3312122132010331-3103202001101330-3221303130232303"></a>

### Direct properties for `dynamic_reverse_proxy.resolution_network`

<a id="canonical-3300130110213322-2033023121303131-0322022120321010-2031100303233130-0033131101331112-3100331302112101-0123032100001131-1103221033132010"></a>

#### `dynamic_reverse_proxy.resolution_network.kind` property

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

<a id="canonical-0332300122302032-2201103333100232-3201320201321221-3033020123112300-0102213011320201-2020331130323130-2331001000332332-3311230233032221"></a>

<a id="canonical-3300203103323013-0021213302321000-2110003313230132-3132133111331233-3331222221300220-0230131230331222-1000121330310031-0111033013101121"></a>

#### `dynamic_reverse_proxy.resolution_network.name` property

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

<a id="canonical-3030031013223202-3132003021232202-3032020020210201-1213002033221110-2222101222221300-1122002222231100-2030001102222002-3102113221212030"></a>

<a id="canonical-0113302112103030-1133012232232102-1010000332322122-1201023031323021-3330000210300212-2200102111101012-3321110232120130-0111332333332132"></a>

#### `dynamic_reverse_proxy.resolution_network.namespace` property

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

<a id="canonical-2221121323012333-2030311132131000-0121232223200221-2111301111231320-1202323202320031-3011112301121100-1321301331102200-2303313311123303"></a>

<a id="canonical-1201223212022131-1313012313132030-3033322001303333-1200200132303110-0301010232311013-3021202012210201-0102331103331132-1110300001101131"></a>

#### `dynamic_reverse_proxy.resolution_network.tenant` property

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

<a id="canonical-1132023131100102-0101122333332123-2320101021133230-2230123020222330-1022222131313332-3002001110221222-3111330122331220-3122100003321120"></a>

<a id="canonical-0330302331322231-0111133332110133-3012313031222313-1311100113012020-0320111030230103-1023202200122320-2312320223032212-1222033122233022"></a>

#### `dynamic_reverse_proxy.resolution_network.uid` property

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

<a id="canonical-0101130103132132-0220121022213231-2333332112030130-0121011300022021-3031032103223222-2102320203302331-2033233303102321-3213203032133202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_path_normalize` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- enable_path_normalize

<a id="canonical-0201022300021320-1303303003222003-0131032130300103-2101300200200320-1102022211110023-3002001212313220-0232120011002133-0313322031310312"></a>

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

<a id="canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- http_protocol_options

<a id="canonical-2310113310130010-3202002120331033-0311323332301131-0002201322002231-2330330123123022-1120101032001030-2313321211221113-1133330133022113"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

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

<a id="canonical-3322002001320020-3232021032121201-2203133033110103-1303032210212200-3013001001013010-1121232112020002-1112233002132012-3101030103032312"></a>

### Direct properties for `http_protocol_options`

- [http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--virtual_host--reference--group-002.md#canonical-3210120111312310-2121102031133313-3130022301322021-2012120212332123-0113011331121132-3312213132131213-2113212202323121-3011331111311333): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--virtual_host--reference--group-002.md#canonical-1032310021331000-0032112130231321-1331130221313133-1120003323332031-0032200222220100-0213231100001120-1230112322102122-0321132010132111): complete subsection reference.

<a id="canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-2120030322212322-2013100321103031-2132112213110203-2310201230132310-0312211110131133-3322120012323020-3003321110133113-0310310012222201"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3332203000222033-0111322002210233-0000101312222213-0213030133132033-3211310202212023-2112131233013231-1230323110221323-3300010223100333"></a>

### Direct properties for `http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130323112003123-0132103001001033-3311111001110313-3311020202100020-1302112201113113-1000033210213210-0130111202301032-1221110113013200): complete subsection reference.

<a id="canonical-0130323112003123-0132103001001033-3311111001110313-3311020202100020-1302112201113113-1000033210213210-0130111202301032-1221110113013200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0120332231310121-0113131120312310-1211221010210013-0031020333330210-1333013003123023-1113211132212222-0002003313032023-2310310221333231"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

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

<a id="canonical-2110031133303212-2321203112001102-2322212300131320-0310032102220201-0122310210123222-0312132001333321-1130002333013111-3131332032330213"></a>

### Direct properties for `http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1001302320233011-3002221201303200-0303303333210232-1101222111030322-0032202021201323-3231122113203201-0110201220101310-3011113120321300): complete subsection reference.

- [preserve_case_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130003111302320-1202220031111320-1200302232213332-0132330021113312-1000322121123001-0232210021233200-0021031133201322-3233010012213020): complete subsection reference.

- [proper_case_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1200133213310011-2013210300031221-0230122333221032-2223133010020231-1323212213101322-3010133032131223-2311201102100320-2013310023320012): complete subsection reference.

<a id="canonical-1001302320233011-3002221201303200-0303303333210232-1101222111030322-0032202021201323-3231122113203201-0110201220101310-3011113120321300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130323112003123-0132103001001033-3311111001110313-3311020202100020-1302112201113113-1000033210213210-0130111202301032-1221110113013200)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-0030121220133201-0213311121113220-2100110103320320-1221333110111101-0213031001333223-3202303333332302-3010213113200331-0211003311131303"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130003111302320-1202220031111320-1200302232213332-0132330021113312-1000322121123001-0232210021233200-0021031133201322-3233010012213020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130323112003123-0132103001001033-3311111001110313-3311020202100020-1302112201113113-1000033210213210-0130111202301032-1221110113013200)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2233331223213101-3023030222201313-0211012203010211-2131230000231202-1031111220030233-1211332231333121-1010123001311312-1222312101210023"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200133213310011-2013210300031221-0230122333221032-2223133010020231-1323212213101322-3010133032131223-2311201102100320-2013310023320012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-2130212003102213-1012210320213222-0323313123201010-1312132120032323-2120333313331231-1121001311221223-0320310003200300-3022021023121203)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0130323112003123-0132103001001033-3311111001110313-3311020202100020-1302112201113113-1000033210213210-0130111202301032-1221110113013200)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1230200223331302-2002002213010322-3013103231233030-3120303222000121-2023110213010322-1303313210001222-0101332311313131-3302301200120123"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210120111312310-2121102031133313-3130022301322021-2012120212332123-0113011331121132-3312213132131213-2113212202323121-3011331111311333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-3303201210330112-1103310130233322-1233213311213131-3233203100332101-3013230011313001-0311010121312000-1212302030323210-1111201323321210"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032310021331000-0032112130231321-1331130221313133-1120003323332031-0032200222220100-0213231100001120-1230112322102122-0321132010132111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-1022222323100221-0023222323033312-1203313231012101-0133220122330221-3323322030003302-2322233000132032-3030200101310212-2313110133001022)
- http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-2211100030121313-1121323330100333-3211112023122100-0030201230331100-3103030202000302-1111213101120220-1001021233020211-2022321222210223"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311311032210221-1213332233330222-2103321233000202-2020120320331211-3131003020011333-1211120300112103-3021112313203310-2031301312202110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `js_challenge` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- js_challenge

<a id="canonical-1233322231003110-0102032003010110-2120003213121113-1012123200113212-2002113320322202-0133131333003033-3333000122301333-0101331330331000"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

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

<a id="canonical-2002000030212230-1213310133211201-1221111002331021-3233300320021310-0020021323023212-2313112113002330-0120310102100321-0011302132020331"></a>

### Direct properties for `js_challenge`

<a id="canonical-2313113333022330-1020033223233003-2130201233133003-1002002123131033-0213020231321222-2033002121320310-1133123211230323-1011301023100201"></a>

#### `js_challenge.cookie_expiry` property

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-1302230100332332-2202313211121220-0112313320203313-1211131320323102-3310132321221203-1310313010201213-2113100130022103-0010311132323211"></a>

<a id="canonical-3012012323020111-0301012231123103-3301333132230312-1002011100211010-1023221230111110-2313223023002332-2232111201003020-0001032203102231"></a>

#### `js_challenge.custom_page` property

Type: `"string"`. Computed.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0103120220231031-1132002321101122-0112010200210321-3303312133213213-1302000012302123-0010032133330111-1122032303302303-1112323322232230"></a>

<a id="canonical-0303330020201110-2020123113202303-1222231313321013-3212221003310020-1131223203203302-0221233312133013-3233220130202132-0003132203021332"></a>

#### `js_challenge.js_script_delay` property

Type: `"number"`. Computed.

Delay introduced by JavaScript, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-1233312122313321-2011202230001010-3102133312211200-1130003122002233-2223112232011102-0131310120111220-3120020032023010-0313002202102123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_authentication` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- no_authentication

<a id="canonical-1102100302002300-1313133201331100-3113202203002110-2131030023211011-2113223023332022-0203202123003020-0213012011103301-3232331311333110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no authentication.

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

<a id="canonical-3110101102113030-0110100111222000-3312320302213100-0201010223313300-1213232220022000-2133222310110000-2331222120121232-2212130222311233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_challenge` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- no_challenge

<a id="canonical-0333002331231113-1021121120232222-1021202203033121-0100301010000210-0311202222011232-1313320122120233-3202220223102123-3332002100110330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no challenge.

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

<a id="canonical-1110112230220221-2303010133322201-1101300033010303-0103123220222322-3010312001321310-0123120221200321-1022302003222231-0020320101133303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_request_limit_per_connection` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- no_request_limit_per_connection

<a id="canonical-2222231010030030-3130021202122031-3100103020130033-3121222000100032-1102001222302230-2111131103111212-2133112303230032-3021011203213203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no request limit per connection.

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

<a id="canonical-2023112320112123-3300121231232030-1100122000331201-1003213223120203-1222330212012300-3302102223112001-3102332220100221-1121233221011031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- non_default_loadbalancer

<a id="canonical-1030023303220330-1233033230000310-3231300000123020-3023333313100103-3300223002022021-2032203120322203-1313212202312012-3321111231213003"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202312331000133-2122223121103310-2100301313303013-2310311201301110-2210333202010230-0211311013211222-0101120011203120-3303021112101311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `pass_through` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- pass_through

<a id="canonical-2011230100101232-2231130122231011-0101023133113300-2212113001232301-2201101303031211-2222001202121303-1333000311103303-2303303020030321"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133310002220332-1103221000011023-3032223311312122-2321113131121222-2110333100201023-1011110232221212-1300111103011301-1303120213023022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limiter_allowed_prefixes` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- rate_limiter_allowed_prefixes

<a id="canonical-0331330100322200-1002100120300333-2120120320132113-3002220323202030-1000321101320011-1133201200033030-1322111110211323-3011210232003110"></a>

Type: `"list"`. Computed.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-3330013312312021-0212021232033212-2211312220233022-2320211031333212-2113120300103331-2322323300201232-1030030122023123-1213310321312130"></a>

### Direct properties for `rate_limiter_allowed_prefixes`

<a id="canonical-3003022201012312-1020122303201010-1221030322123200-3032321102213223-2011202001121100-2210003123231310-3001133022303023-3030233101233031"></a>

#### `rate_limiter_allowed_prefixes.kind` property

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

<a id="canonical-2300030130022113-2002103122312233-0332110112313112-3202313123320002-1230311322013002-1031321111000223-2100220012333200-3003023301123303"></a>

<a id="canonical-0223031322321011-3001230200123303-0200120003302311-1331201120311203-0102223001020100-1031031130223311-2113000303323013-1000033103320202"></a>

#### `rate_limiter_allowed_prefixes.name` property

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

<a id="canonical-0123332313122333-2233322111222220-2232000033013202-0202221023100030-0320330123210302-0121323021303203-1211002111011002-3010202021201023"></a>

<a id="canonical-3002131221132102-3102013001222303-3020133111222120-2112210003313102-3131111323100130-1001230202200130-3021312233223213-2103112300230010"></a>

#### `rate_limiter_allowed_prefixes.namespace` property

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

<a id="canonical-0332132321331221-0010120222322311-0232113213223032-0302323220221311-3203021021121303-2301030303122210-0201302331321202-2311133033020220"></a>

<a id="canonical-1030300203111022-2302121210033102-0211030033023203-2323022300233323-0131302233331310-0111132223203012-3130201322223021-0131300130330301"></a>

#### `rate_limiter_allowed_prefixes.tenant` property

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

<a id="canonical-2003130203331213-2112231202131123-0011312113200000-3123033332122320-0200100301022012-0103211130201011-0332102312300101-0322031302120323"></a>

<a id="canonical-0023320222000003-0100312310221003-0221021031113120-2300201110122003-3101211130313333-2000210332212321-1313011013233223-2010323212013221"></a>

#### `rate_limiter_allowed_prefixes.uid` property

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

<a id="canonical-3012001313211032-0121110012120332-1211212103012331-1030201012110112-3012311230231023-0312310313121023-3122223302301012-1011130302002022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_cookies_to_add` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- request_cookies_to_add

<a id="canonical-1313120312033223-3223302312021001-2210111031213102-0032030313222120-1111023302121322-0321022003120100-1002030012222312-3203122220311202"></a>

Type: `"list"`. Computed.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

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

<a id="canonical-2232203312321211-0001013312201130-0210233022031300-2112120130110230-3302012223333132-0213033301111300-2202030101212032-0030201223113012"></a>

### Direct properties for `request_cookies_to_add`

<a id="canonical-3210111320021210-0032132032121020-3302313300112212-0032120213000131-3303212001021130-0210100012323310-2201200303232120-3330020121311200"></a>

#### `request_cookies_to_add.name` property

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3123011230320133-2030230232131002-2132301112010200-3000221223231033-2313112322222022-1020210231223000-2300113110033321-3213101011220211"></a>

<a id="canonical-3201101321100102-1320232233031022-1321001023332200-2220133000033001-0201112312211113-3232131231120131-0001302133130322-2010321003320100"></a>

#### `request_cookies_to_add.overwrite` property

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Additional upstream details:

If true, the value is overwritten to existing values. Default value is do not overwrite.

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

- [secret_value](data-sources--virtual_host--reference--group-002.md#canonical-2221012012102111-2312110102322101-0103322330302302-1323220122123302-3032133202210302-3121200131100310-1322210101030202-2123321301200213): complete subsection reference.

<a id="canonical-3022111331213332-0111133010110023-1121102122000323-3311001313022000-3321032110313012-0123130321032020-2320111232303303-0231121300231003"></a>

<a id="canonical-2012121113321111-1311213120323030-1320312231221323-1302100200021023-3121112203323102-3113010112100101-2332023330330203-1333132103123032"></a>

#### `request_cookies_to_add.value` property

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the Cookie header.

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

<a id="canonical-2221012012102111-2312110102322101-0103322330302302-1323220122123302-3032133202210302-3121200131100310-1322210101030202-2123321301200213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3012001313211032-0121110012120332-1211212103012331-1030201012110112-3012311230231023-0312310313121023-3122223302301012-1011130302002022)
- request_cookies_to_add.secret_value

<a id="canonical-0322303201132233-3331133033102320-3102210333112212-2121003031010100-2022302021021031-0203120133110221-0322311031233013-0123233230031222"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-2100332313033011-0021311232331002-0333212203102130-2231203132101232-0212112001003300-3201122211201210-0323011020121202-1122323303102211"></a>

### Direct properties for `request_cookies_to_add.secret_value`

- [blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-3322003220000330-2001322312113202-3022120321102120-3112112220322111-3303332001032122-2231023301212002-2120202311103032-0101303213310030): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-1021000123131222-3033303221301233-2130203001030221-2132332321310010-1013300122330020-0310222323300021-1012200331010301-1200011230231132): complete subsection reference.

<a id="canonical-3322003220000330-2001322312113202-3022120321102120-3112112220322111-3303332001032122-2231023301212002-2120202311103032-0101303213310030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3012001313211032-0121110012120332-1211212103012331-1030201012110112-3012311230231023-0312310313121023-3122223302301012-1011130302002022)
- [request_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-2221012012102111-2312110102322101-0103322330302302-1323220122123302-3032133202210302-3121200131100310-1322210101030202-2123321301200213)
- request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-1102332111131111-2312210102230132-2102301032003113-1222233223112201-0103032111122220-3020133233303122-0123020330232333-1122130322203001"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-3330103002321031-1032323002310331-1123113123102133-2333333130003033-2302101030001133-0300100203210322-1321122103033300-0013022000103020"></a>

### Direct properties for `request_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-0011112001231020-0202022121232302-1313002012223121-3332000123100131-1013300013130302-1030223302012111-3100330301110020-3321032102301203"></a>

#### `request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-3203200133231320-3031120000321011-2221130111333310-0200203230030211-2231210200203130-0000201320023200-2012003002032200-1212202200213233"></a>

<a id="canonical-2200213323131213-0230330220113120-2132230232023010-1222023332031330-1013023211312132-1031111330313210-3221100011032032-1220012013001333"></a>

#### `request_cookies_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0331331030323300-0132232222123003-2033300320120322-1032320321132331-1000223130031122-1221113010222033-3133002212133132-0200230301210200"></a>

<a id="canonical-1122303231333222-1002002133221030-0233222023130130-0021121320203120-2002132312210300-1122130033222303-0131300233201313-3013001212010222"></a>

#### `request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-1021000123131222-3033303221301233-2130203001030221-2132332321310010-1013300122330020-0310222323300021-1012200331010301-1200011230231132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3012001313211032-0121110012120332-1211212103012331-1030201012110112-3012311230231023-0312310313121023-3122223302301012-1011130302002022)
- [request_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-2221012012102111-2312110102322101-0103322330302302-1323220122123302-3032133202210302-3121200131100310-1322210101030202-2123321301200213)
- request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-3020312203131200-0020022130002033-0321323213310022-2313111212203031-2030211221112111-0213103303011130-1323002013112332-3320202133102220"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-2000100111031102-0221330232002030-0322310111132001-3002233200002000-3232311302311010-2012321000130000-3323333023112122-2213232232300321"></a>

### Direct properties for `request_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-3232213000032311-3301210102010031-3222232013311122-0233112321032130-2120322312202132-0330000323231203-2202231213101220-1003102131001330"></a>

#### `request_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0300321231110031-2331012232231000-0123331022213032-3030302211320333-0113230030130112-1201122031033032-0211323200112230-3233131303313211"></a>

<a id="canonical-0130221101230002-1001033220313220-1221020200032210-2202033233212003-1330013222301211-0101001203021310-0213233221103230-3321212111000300"></a>

#### `request_cookies_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3303213013223123-0333021210322333-3000030200231312-1002303202333220-2113133012232131-0030321011113311-3012010300230300-3131003321220113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_headers_to_add` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- request_headers_to_add

<a id="canonical-0212331111220023-1110003001131213-2230212300133010-0032203310311221-3132120120212033-2213331220010100-0122232301011223-1113212132012332"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

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

<a id="canonical-3233320103310223-2103332003102333-0210313011033032-0313122303230002-3211332202322133-3221110201022210-1120313223033313-0300320132203010"></a>

### Direct properties for `request_headers_to_add`

<a id="canonical-0111003322111311-0210311322203231-0033231212032013-0120201220320101-3322301111220313-0130312231102133-1010101013000110-0323013230030323"></a>

#### `request_headers_to_add.append` property

Type: `"bool"`. Computed.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Additional upstream details:

If true, the value is appended to existing values. Default value is do not append.

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

<a id="canonical-2013010322133123-3312231203130101-0001201323033122-0313203013300222-2002002231011321-0320010312120320-1311220121011202-3021332101220200"></a>

<a id="canonical-2222112303000101-1210102322011201-2131211033303122-1112113202230320-1301103330302122-2020111100233231-1203233300320130-1130222033120133"></a>

#### `request_headers_to_add.name` property

Type: `"string"`. Computed.

Name. Name of the HTTP header.

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

- [secret_value](data-sources--virtual_host--reference--group-002.md#canonical-1330032123203232-2013202012321330-0033330131300212-1213132002032112-0122223201331312-1123132323101310-2132002113033311-1331011333020313): complete subsection reference.

<a id="canonical-0002221320121320-1033122001021012-2021121302103102-0303312323133110-2231122220223301-3120003322302203-0211231223130313-2203011201111113"></a>

<a id="canonical-0130313333222233-3211333100101302-0123212020331222-2001132013313111-3133012131133300-0121032211210200-0001010302330310-1201233130332232"></a>

#### `request_headers_to_add.value` property

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

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

<a id="canonical-1330032123203232-2013202012321330-0033330131300212-1213132002032112-0122223201331312-1123132323101310-2132002113033311-1331011333020313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3303213013223123-0333021210322333-3000030200231312-1002303202333220-2113133012232131-0030321011113311-3012010300230300-3131003321220113)
- request_headers_to_add.secret_value

<a id="canonical-1113100003100331-1012223330010101-2311201233313310-2202132323023113-1002120103022112-2322022213110322-1002212202302302-0022023101201032"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-2320213120022321-3321110102332230-2122323223332121-2223110313223022-0231310103122030-1031212000311111-1022202221110130-2213332212000320"></a>

### Direct properties for `request_headers_to_add.secret_value`

- [blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-2033131003100211-0303122221321010-0333312111121321-0012031121300212-1333011022011210-3200220232000331-3020131301233330-2223303200330022): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-3011313210203032-1222222310212103-3111133313132122-1222112203233121-0322012013002212-2101111133232210-0101031020303223-0023103231210223): complete subsection reference.

<a id="canonical-2033131003100211-0303122221321010-0333312111121321-0012031121300212-1333011022011210-3200220232000331-3020131301233330-2223303200330022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3303213013223123-0333021210322333-3000030200231312-1002303202333220-2113133012232131-0030321011113311-3012010300230300-3131003321220113)
- [request_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-1330032123203232-2013202012321330-0033330131300212-1213132002032112-0122223201331312-1123132323101310-2132002113033311-1331011333020313)
- request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-1110020211333301-0201002333112132-1123323201210102-1122032113112330-0033210303230021-3311212110111200-2322201312102122-0103010302332210"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-2303033311110331-3101002100233020-2333232030110011-2012300023311301-0011332231303323-1133232003213210-1202032213320312-3212310333003211"></a>

### Direct properties for `request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3221300001313201-2202013112112012-3210001201320011-0211221022210112-1022212232103113-1310001122321233-3210230123221120-3323312303010110"></a>

#### `request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-1130010110321230-3102120121003322-1010221321200311-1132112033100303-0001301202313013-3130212230131221-1310230033220310-3200222213021312"></a>

<a id="canonical-0101302220203013-0301130213030323-0022210000130300-2310100202023122-3231131103320022-2302002212213233-1210100030221102-1022001331323303"></a>

#### `request_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3303102130111003-1333003000312221-0023001320311002-2133322032102330-1230230222003021-0100332210120222-2321030033213221-0130122233033021"></a>

<a id="canonical-2201322220232320-3221030212333111-3333221121120003-3230132133130331-3313030031302031-3330320323331032-3331321323003303-2333011331123210"></a>

#### `request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-3011313210203032-1222222310212103-3111133313132122-1222112203233121-0322012013002212-2101111133232210-0101031020303223-0023103231210223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-3303213013223123-0333021210322333-3000030200231312-1002303202333220-2113133012232131-0030321011113311-3012010300230300-3131003321220113)
- [request_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-1330032123203232-2013202012321330-0033330131300212-1213132002032112-0122223201331312-1123132323101310-2132002113033311-1331011333020313)
- request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-3121111213002323-1231101112112231-3101333032202332-0200013011002332-2223020023231333-1133033202321323-2031003023020022-1103101121113301"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-3031023132332012-1033112131021213-1313333322113232-3032312232131100-0002113221022012-1312222123033130-1213221123013020-3021213331121102"></a>

### Direct properties for `request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-1200322010133001-2100232100010330-1133302301320120-3201131331201221-0221220222211133-0120320101303330-3211133301010023-2230331232012120"></a>

#### `request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2330021311131111-1132011222001031-0010222130001213-1112132200302132-0322123330122212-1220113330112203-1210310130002200-2211211201220201"></a>

<a id="canonical-1233223330222022-2320013133221132-0003101220231331-2120030320202001-1211021111321331-2103302131311130-3330132102211031-2021012213303211"></a>

#### `request_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- response_cookies_to_add

<a id="canonical-0113212001232011-3133322010021302-3102230131212330-2020322230330311-0130213133032111-3122133022112220-3200232202122232-3312300203123230"></a>

Type: `"list"`. Computed.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

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

<a id="canonical-3003023123311110-1320212332103301-1212120102323032-1212323323322320-2011101120013013-2222003030301331-2232110331003233-0202113023001011"></a>

### Direct properties for `response_cookies_to_add`

<a id="canonical-0001011020132012-1232111312113132-0330300223301231-1011002313230031-3232220300331003-2300110310311011-3121013011113200-0013021222011120"></a>

#### `response_cookies_to_add.add_domain` property

Type: `"string"`. Computed.

Exclusive with \[ignore\_domain\] Add domain attribute.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0323331213312223-3021003123002211-1202213220023132-3010013212200011-1022332133021023-3323301323122213-0303012032122120-2201212310221002"></a>

<a id="canonical-3011222323203103-1210102001333311-3132020122120220-0202023100112033-1101331211121122-2330100111032113-3020102010312101-2223230000303010"></a>

#### `response_cookies_to_add.add_expiry` property

Type: `"string"`. Computed.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_httponly](data-sources--virtual_host--reference--group-002.md#canonical-3133121010112233-3322313331011132-2132333100021313-1103100203002133-0322310213022112-3130232120302020-1130301033103101-2131030113031102): complete subsection reference.

- [add_partitioned](data-sources--virtual_host--reference--group-002.md#canonical-1300012000233302-2313330232300323-2330201210102101-2032131201102111-1012033012111001-0303101333222210-3103312110002222-3110030211130311): complete subsection reference.

<a id="canonical-0211031222022221-1210232233232213-0123111010113301-1210201123011222-3321203311111332-0020311311131231-2211233232100232-0022110201032030"></a>

<a id="canonical-0103113301313211-2200022000302323-2223113201303330-1001330230012012-3132210230330002-1210112302033323-1300112221210032-1131003120111010"></a>

#### `response_cookies_to_add.add_path` property

Type: `"string"`. Computed.

Exclusive with \[ignore\_path\] Add path attribute.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_secure](data-sources--virtual_host--reference--group-002.md#canonical-1131111231223133-2310003201221100-1031203233212002-1112222321112123-1322013301002233-2203202100133322-0013112323023003-2211003331110111): complete subsection reference.

- [ignore_domain](data-sources--virtual_host--reference--group-002.md#canonical-3130310133231002-1131000313113312-2030022200200233-2011302002303013-3303011201121220-3202203133303302-2011203131131013-2332211220103303): complete subsection reference.

- [ignore_expiry](data-sources--virtual_host--reference--group-002.md#canonical-2030201030022210-2032320011321111-0112121020213023-0333200011123312-3333331012331211-1023111003112233-3010101100331300-0211322231032200): complete subsection reference.

- [ignore_httponly](data-sources--virtual_host--reference--group-002.md#canonical-2113311312232221-0013310131132212-3010132032113032-1201123320023131-0122300301100323-1020233312020223-2231033013201101-2032022103010000): complete subsection reference.

- [ignore_max_age](data-sources--virtual_host--reference--group-002.md#canonical-1213131003133130-1132100030000101-1312020211100300-1023102133112033-3323332101222100-0231120312322020-1322000302113023-3023223103021013): complete subsection reference.

- [ignore_partitioned](data-sources--virtual_host--reference--group-002.md#canonical-3013010331312000-2120221322022202-0012313203013322-1023230122220300-1011300220110232-2010213110003132-3110210010212011-2133323103202010): complete subsection reference.

- [ignore_path](data-sources--virtual_host--reference--group-002.md#canonical-1102211312201202-3111210321011033-3221202120032131-1310122310121202-1332103031102001-2210203322320113-3110230213230232-2212122010333220): complete subsection reference.

- [ignore_samesite](data-sources--virtual_host--reference--group-002.md#canonical-0330232211331123-0102120023332301-3231132200322201-3312221110102011-1310012301122323-1000113222013221-3330022122030133-0312230303222303): complete subsection reference.

- [ignore_secure](data-sources--virtual_host--reference--group-002.md#canonical-1203333011321131-3020123131103020-0121032013213000-1323211132232130-3202000322303300-1330102133201011-1212021303032023-0100122223001102): complete subsection reference.

- [ignore_value](data-sources--virtual_host--reference--group-002.md#canonical-0103130323203132-3111033320232102-1202231000313232-0102103013333002-2332301022003020-2020200133102202-3103033030130030-3331001202111002): complete subsection reference.

<a id="canonical-1223332230031002-3021221300023122-1202320001120111-1301132121100020-0021200302320200-1031203103230123-1212100133020002-1023133010220030"></a>

<a id="canonical-0030012303333322-0130233313322120-3332222230311321-2310203120023331-0131031123033123-1032220132122012-3201023301222123-0110320232033033"></a>

#### `response_cookies_to_add.max_age_value` property

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-0033210331200122-0220202000132121-0133233333200332-0202313322332222-3132303020321000-1030002033013000-2100222230022023-1011132021323013"></a>

<a id="canonical-1320232132100333-2302301130200021-0132102203023023-3001322002320130-3031122222313113-3130000210210221-0022110303103110-1020221201003133"></a>

#### `response_cookies_to_add.name` property

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0131110323003203-3300311130312322-0010021032303102-3033001202211300-3133130232132202-3203131033323323-2310212232320232-0322100132320110"></a>

<a id="canonical-2301003110113133-0223202100221100-3031003300232132-0103333121303323-2112203302200000-1130211310010020-2002100002022000-2013313331130313"></a>

#### `response_cookies_to_add.overwrite` property

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Additional upstream details:

If true, the value is overwritten to existing values. Default value is do not overwrite.

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

- [samesite_lax](data-sources--virtual_host--reference--group-002.md#canonical-1032311232333023-0123210200322102-2332221303132113-0311311123312313-0130113003223133-1222002122120001-1301232232200131-2113321032033121): complete subsection reference.

- [samesite_none](data-sources--virtual_host--reference--group-002.md#canonical-2233112002222232-2223030333212333-3201132101110203-0211330111331112-0221230222022203-0010121333033210-0122232002122312-1310233330303223): complete subsection reference.

- [samesite_strict](data-sources--virtual_host--reference--group-002.md#canonical-1322201213133222-1023323300033031-0231030302203233-0133010022003131-3302220101130320-3210200311111031-3001003010311323-1132010111102000): complete subsection reference.

- [secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3313221231300020-1321102321213000-1322132323323113-1300023011223200-2300033321130201-1012102200310331-3101122201123203-1330201223331132): complete subsection reference.

<a id="canonical-0132212221332200-0112030120132231-0303203332110213-0221133212132012-3112222223320130-3311011121331322-2300303301330321-0232012032110113"></a>

<a id="canonical-3202032203301331-2323110330323131-2321331212332111-1001012303123012-2200122023012321-0223232211231232-0332132310212001-2212322013001101"></a>

#### `response_cookies_to_add.value` property

Type: `"string"`. Computed.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

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

<a id="canonical-3133121010112233-3322313331011132-2132333100021313-1103100203002133-0322310213022112-3130232120302020-1130301033103101-2131030113031102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.add_httponly` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.add_httponly

<a id="canonical-3331200311131232-2111312122313330-3311022112022321-0210222232203233-2322120132222221-0330300102033233-0002221112121012-2231003222320123"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add httponly.

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

<a id="canonical-1300012000233302-2313330232300323-2330201210102101-2032131201102111-1012033012111001-0303101333222210-3103312110002222-3110030211130311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.add_partitioned` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.add_partitioned

<a id="canonical-2030311302202101-1021020302223000-2330003013102031-1031233032103232-3210302012231012-3212130200300211-1000222233330032-0202323320031200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add partitioned.

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

<a id="canonical-1131111231223133-2310003201221100-1031203233212002-1112222321112123-1322013301002233-2203202100133322-0013112323023003-2211003331110111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.add_secure` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.add_secure

<a id="canonical-3003210300322131-1210230110002132-0111332020330200-2212311130233302-0330011012323322-1132211132132020-3220101313231112-1023130121100010"></a>

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

<a id="canonical-3130310133231002-1131000313113312-2030022200200233-2011302002303013-3303011201121220-3202203133303302-2011203131131013-2332211220103303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_domain` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_domain

<a id="canonical-2300121232212121-0121210302102312-2113030112230210-3213202010320220-3000113302302223-2020031013213103-3000003131023032-3302121313313112"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore domain.

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

<a id="canonical-2030201030022210-2032320011321111-0112121020213023-0333200011123312-3333331012331211-1023111003112233-3010101100331300-0211322231032200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_expiry` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_expiry

<a id="canonical-0131302033331002-2302213311003033-0320130131231012-2312100333202102-0023100123222031-1213230122321313-3121023211212321-3101233013232101"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore expiry.

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

<a id="canonical-2113311312232221-0013310131132212-3010132032113032-1201123320023131-0122300301100323-1020233312020223-2231033013201101-2032022103010000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_httponly` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_httponly

<a id="canonical-0301030232121022-0112212033023121-2032223012003223-2230102331131002-3330102031321322-3333222213312032-3120310303312112-0330111332132312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore httponly.

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

<a id="canonical-1213131003133130-1132100030000101-1312020211100300-1023102133112033-3323332101222100-0231120312322020-1322000302113023-3023223103021013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_max_age` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_max_age

<a id="canonical-1312333330322012-1230201311201103-2333211023331033-2201220330121222-1321331313012310-1310302302331010-0333121011001310-1121031003120230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore max age.

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

<a id="canonical-3013010331312000-2120221322022202-0012313203013322-1023230122220300-1011300220110232-2010213110003132-3110210010212011-2133323103202010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_partitioned` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_partitioned

<a id="canonical-3110203131101201-2120323010230203-0202113001231330-3013301310123112-2002133221302221-0101033122100131-1331130012322133-2301222312322232"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore partitioned.

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

<a id="canonical-1102211312201202-3111210321011033-3221202120032131-1310122310121202-1332103031102001-2210203322320113-3110230213230232-2212122010333220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_path` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_path

<a id="canonical-2130320121310303-1332032310213100-2222113113013213-2303023222113201-1312310230031200-0122023033133131-1230233332231203-3230131012023332"></a>

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

<a id="canonical-0330232211331123-0102120023332301-3231132200322201-3312221110102011-1310012301122323-1000113222013221-3330022122030133-0312230303222303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_samesite` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_samesite

<a id="canonical-3021320121212200-3213213331333010-3110311312301132-0231220113203222-2131102100213123-1320001210231333-3131200320223223-1112322321131001"></a>

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

<a id="canonical-1203333011321131-3020123131103020-0121032013213000-1323211132232130-3202000322303300-1330102133201011-1212021303032023-0100122223001102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_secure` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_secure

<a id="canonical-2302001212203030-1130010112322122-2033200121331232-3200023020330333-1102221111111232-0221131320122112-2011203322210322-1030202221211032"></a>

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

<a id="canonical-0103130323203132-3111033320232102-1202231000313232-0102103013333002-2332301022003020-2020200133102202-3103033030130030-3331001202111002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_value` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.ignore_value

<a id="canonical-3222313001302022-0111130312012223-0231010310010311-1112112223103321-2201303303000332-3322220033221033-1301013100210230-0020213102300331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore value.

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

<a id="canonical-1032311232333023-0123210200322102-2332221303132113-0311311123312313-0130113003223133-1222002122120001-1301232232200131-2113321032033121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.samesite_lax` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.samesite_lax

<a id="canonical-2033121212113103-0210022033031232-2333211101213203-2320033333020211-0031301210202211-1103111012210301-0213003303311011-0321333020001302"></a>

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

<a id="canonical-2233112002222232-2223030333212333-3201132101110203-0211330111331112-0221230222022203-0010121333033210-0122232002122312-1310233330303223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.samesite_none` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.samesite_none

<a id="canonical-3220320102032323-0102112031021131-2101130101011021-3133023100130231-0301332213303331-0232020023003022-0200212031300200-3202111010023000"></a>

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

<a id="canonical-1322201213133222-1023323300033031-0231030302203233-0133010022003131-3302220101130320-3210200311111031-3001003010311323-1132010111102000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.samesite_strict` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.samesite_strict

<a id="canonical-0020132200132200-2312023030212103-0101121332113211-0300223301032223-3131310123120132-1030202201211311-3033223330103300-1001010231133231"></a>

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

<a id="canonical-3313221231300020-1321102321213000-1322132323323113-1300023011223200-2300033321130201-1012102200310331-3101122201123203-1330201223331132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- response_cookies_to_add.secret_value

<a id="canonical-0230220011133210-1121201122001221-2000212031123303-2011300012231322-0222223121102113-1020232220032232-2213211003211031-3233111203002212"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-0213032330233322-0221111301230031-3120000031213100-1220002330321002-2232121001120102-0230233013001300-0020232002220123-2202321100032132"></a>

### Direct properties for `response_cookies_to_add.secret_value`

- [blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-3300132111031211-1300013102022033-3332200320011200-2012102022101122-2001121123201220-1320332302132110-3321311012311203-1330022030301103): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-0102213233221021-3120100212310301-0312202123311301-3333333110333131-3030200311231130-1130102230300212-3133223031012311-2112320110222021): complete subsection reference.

<a id="canonical-3300132111031211-1300013102022033-3332200320011200-2012102022101122-2001121123201220-1320332302132110-3321311012311203-1330022030301103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [response_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3313221231300020-1321102321213000-1322132323323113-1300023011223200-2300033321130201-1012102200310331-3101122201123203-1330201223331132)
- response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-0121002103233003-2010103032030132-2022032223201103-3322111330020103-2222021232310220-1223233302111221-0230220220201122-1022002331131202"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-0111203033332113-0031131121233122-3210203323012320-3100020022123300-1210332200202012-0130313320020232-0122012331111210-0000120033232021"></a>

### Direct properties for `response_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-0012113123020121-1200013002032230-3330322213112113-1020322132300102-0110011223303332-2113013103003112-2300312331121330-3233223133103002"></a>

#### `response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-0302122213101131-1201112003300311-0100202202232231-3311121103123012-2323231121121112-2100023230202130-0333031121231000-1332322330021013"></a>

<a id="canonical-1221330212203313-0120312121103303-1020200330111213-0211222210130013-0121211211022012-3230323212030131-0023212233300301-0310302100031003"></a>

#### `response_cookies_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1020312200013313-2020300101100203-0232321123331320-3121113012310200-1231310002031023-0332120303101003-0112111203122001-0123002323122201"></a>

<a id="canonical-2322312100210300-3122321311010321-1212331320111003-1201302210131101-0131033101322133-0220111210313112-1030321220310332-3123211023100011"></a>

#### `response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-0102213233221021-3120100212310301-0312202123311301-3333333110333131-3030200311231130-1130102230300212-3133223031012311-2112320110222021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-2302322223131111-2321031223122023-1002301303112033-2332212120000303-1110013000232313-0223200003133213-1210023012222332-3121023023233022)
- [response_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3313221231300020-1321102321213000-1322132323323113-1300023011223200-2300033321130201-1012102200310331-3101122201123203-1330201223331132)
- response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-3313221201232231-0032112231301020-0000102331131233-2230002130021321-1300331221033122-2331012012100222-3333311331212011-1213221010013302"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-3122121313322003-3123000032222123-2313330212221313-2311211313302133-1011010022333202-0311003200133212-1201011203122200-3312221220232223"></a>

### Direct properties for `response_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-2130100100231230-0111333213301122-2223013210030123-0210223112111132-1332013333011322-0330223010022203-0101320132230022-3033120023232331"></a>

#### `response_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2322300220130300-2033203000011122-2111200122010030-2331010010200213-2223120132233322-3103001213000023-3221102222232333-0212231210102202"></a>

<a id="canonical-2123033101123213-3023313301231120-0023022130100200-3303210022131210-3210230230201033-3320310222313102-2131121232200012-1103332330103122"></a>

#### `response_cookies_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1122301030300332-0033230331023011-1102233120003323-3100320132003322-0030000111220313-3302203220201033-0222213312322000-2123201002112113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_headers_to_add` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- response_headers_to_add

<a id="canonical-0101202001030202-2332212330311113-0131001010330313-3313032330010131-2020200230233013-3311312012021031-2120012311323110-1223013022101202"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

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

<a id="canonical-3231000331231111-0013100102201301-2033030010110133-3331100332323110-0232233111030101-2232212202323320-0023030002331100-1112210233033102"></a>

### Direct properties for `response_headers_to_add`

<a id="canonical-3303003312131233-3030200320213123-3100123110201120-3131321300002023-2233310232001113-2010231121122122-3320330023303211-3030031322132203"></a>

#### `response_headers_to_add.append` property

Type: `"bool"`. Computed.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Additional upstream details:

If true, the value is appended to existing values. Default value is do not append.

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

<a id="canonical-2003113000313011-3210313331203312-2212323303332232-3131111332022133-1231002120203303-2020000303233020-2210103001332230-3332312130233310"></a>

<a id="canonical-3112322012231303-3032101212300301-0312220332001112-1200100310301111-2031003022310311-2031231023123211-3302021101033303-0011103003310110"></a>

#### `response_headers_to_add.name` property

Type: `"string"`. Computed.

Name. Name of the HTTP header.

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

- [secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3101230321222032-3323212212032112-0000210100132300-2200202302113313-3202033223320201-3113222013230022-1120310312103300-3330002320302120): complete subsection reference.

<a id="canonical-3322312333311133-0031013222000232-2233131202221213-1100201210300331-3301011233011300-3011131010300113-1200103303323130-0210113003213303"></a>

<a id="canonical-3000222000311012-2221012332310321-1020211213331120-2122222011332313-2322233123320301-2023312203301111-1023211022101013-1121200333211010"></a>

#### `response_headers_to_add.value` property

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

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

<a id="canonical-3101230321222032-3323212212032112-0000210100132300-2200202302113313-3202033223320201-3113222013230022-1120310312103300-3330002320302120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-1122301030300332-0033230331023011-1102233120003323-3100320132003322-0030000111220313-3302203220201033-0222213312322000-2123201002112113)
- response_headers_to_add.secret_value

<a id="canonical-3100032201022001-2001120211311323-2303132030311023-1223333122211033-2111333132122321-0032110321210032-1303101313331301-0220323001111102"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-3122010230223010-1311203312310322-3030130112002223-0020231303113331-0013102013323300-2023033311211001-0230011331233312-2301300310023122"></a>

### Direct properties for `response_headers_to_add.secret_value`

- [blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-2012122311123233-1323213122212312-1301311230331021-0011110120223113-2303130021033010-1120320122031200-2120330332002223-3103211030113022): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-0310123303110031-0303203333000323-2223131103332331-0200031312112020-1120223122110221-2312231113313233-0010021002321202-0100213233303303): complete subsection reference.

<a id="canonical-2012122311123233-1323213122212312-1301311230331021-0011110120223113-2303130021033010-1120320122031200-2120330332002223-3103211030113022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-1122301030300332-0033230331023011-1102233120003323-3100320132003322-0030000111220313-3302203220201033-0222213312322000-2123201002112113)
- [response_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3101230321222032-3323212212032112-0000210100132300-2200202302113313-3202033223320201-3113222013230022-1120310312103300-3330002320302120)
- response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0320122313112313-3331203313022000-0020013010012332-0133303123120012-0301223323113003-1323202202311113-0103112102302020-3111312210311233"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-3133112222113302-2012103301212233-3222221223213201-0302310032220003-2101312031210123-1211203033303101-2001011330330210-0203130022223313"></a>

### Direct properties for `response_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-2321220203333331-1210133320202232-0003321312003130-3020101102233113-2330202230100231-2103013212232033-2220000302030330-1233011223003331"></a>

#### `response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-3000332200020302-1200112232231112-2111231332323331-0213223100200310-0203001300120302-2032202130030311-1201102110132112-1032123023322310"></a>

<a id="canonical-1003100132033010-3200322100000231-0221332030330211-1311031102003130-3233000222033303-3113030100331020-3002233101013203-2020230120301112"></a>

#### `response_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1102310020100012-3102100223203311-3110012320302010-0123100022202212-0033032111133122-0103313212300302-3202100320123303-2101331100023100"></a>

<a id="canonical-0302032300303031-1232202202320323-2023323023012213-2010031113121203-1000021013303232-0303220201330000-1223302003231221-2001221221303000"></a>

#### `response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-0310123303110031-0303203333000323-2223131103332331-0200031312112020-1120223122110221-2312231113313233-0010021002321202-0100213233303303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-1122301030300332-0033230331023011-1102233120003323-3100320132003322-0030000111220313-3302203220201033-0222213312322000-2123201002112113)
- [response_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3101230321222032-3323212212032112-0000210100132300-2200202302113313-3202033223320201-3113222013230022-1120310312103300-3330002320302120)
- response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-3130022021012020-2111013303010033-1333320203310213-1113300030300221-3010301303313032-0131000023311021-1311011231100022-0033220130331233"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-2030232320213122-3023111310003102-1131202220220003-2203301320031312-2221310213132110-0232030303023133-3101131002230231-0020213111133030"></a>

### Direct properties for `response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-3130311230122321-1302312201303300-2201333323000103-2221203213132300-1221320233200232-0202033103310201-1130033232123211-1020312201022013"></a>

#### `response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3021103000001322-1332201012301002-2323332233021010-3223020132221322-3321333002212312-0002112103120230-0002122102112030-2000112103213322"></a>

<a id="canonical-2021030021230321-3130013031300223-0112021201221331-0220011221032000-3032120322201103-2001122112012212-1032030132102101-3030021321112302"></a>

#### `response_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1021031030330102-2100323023122311-3003021113211121-2312223333303303-2002222022211201-0131100013001013-2031022313222112-2203022023311021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `retry_policy` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- retry_policy

<a id="canonical-3121300022323212-3032100030021200-3220123111011103-0022232031022323-2212331012001311-2303233320201201-1331103322001123-2010123121001101"></a>

Type: `"single"`. Computed.

Retry policy configuration for route destination.

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

<a id="canonical-0310211322221223-3233311103233230-0132203200100030-3210121121113222-0312211323201121-1211031030203100-1033001331001123-3113303302011023"></a>

### Direct properties for `retry_policy`

- [back_off](data-sources--virtual_host--reference--group-002.md#canonical-3213001230013032-0020221021322333-1332330301110002-0020303010112211-3321113230333322-3203333212321332-0203333021032210-3112130033232321): complete subsection reference.

<a id="canonical-1220312222312301-1032103101122112-0131120022103202-3313321111013110-0222113311201112-0130211121333200-1110212201330031-1102103301033132"></a>

<a id="canonical-3032221232321322-1122001213223321-0311123201322113-0030112003300333-1123122000203312-2010230121200323-0003311030323222-1312021220230301"></a>

#### `retry_policy.num_retries` property

Type: `"number"`. Computed.

Specifies the allowed number of retries. Retries can be done any number of times. An exponential
back-off algorithm is used between each retry. Defaults to \`1\`.

Additional upstream details:

Defaults to 1.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8,
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
    "ves.io.schema.rules.uint32.lte": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  }
}
```

<a id="canonical-0320211020202131-1110221323012302-0333100001031313-0322023302330120-0201222331000302-2030031122110001-1330003133012202-3121210000302323"></a>

<a id="canonical-2020303221110022-1112213013330221-0200330130112102-3211121313101031-3013020202132032-0100030023321020-0003112302210300-0121021311021213"></a>

#### `retry_policy.per_try_timeout` property

Type: `"number"`. Computed.

Specifies a non-zero timeout per retry attempt. In milliseconds.

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

<a id="canonical-2121022023022001-3000132120321230-2023030231120210-3023100021111221-1021123323210330-2002210323013301-3110301013223023-0133022312311100"></a>

<a id="canonical-3221231302313311-1320131002022202-0221212233020322-1023231103131333-2023223223003032-0011031231030132-3311022120111301-3200213230222323"></a>

#### `retry_policy.retriable_status_codes` property

Type: `["list", "number"]`. Computed.

HTTP status codes that should trigger a retry in addition to those specified by retry\_on.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1330131133121001-0213001123313202-0333010330032300-2200023131011201-3032233303321102-0322311222232033-3011202221002323-1111222331132312"></a>

<a id="canonical-0223322130220012-2202310323330121-3203302320001302-0111200020310102-3231113203333213-2013331031031103-3101111003032323-0232110210200220"></a>

#### `retry_policy.retry_condition` property

Type: `["list", "string"]`. Computed.

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if
the..

Additional upstream details:

For example, network failure, all 5xx response codes, idempotent 4xx response codes, etc

The possible values are

"5xx" : Retry will be done if the upstream server responds with any 5xx response code, or does not
respond at all (disconnect/reset/read timeout).

"gateway-error" : Retry will be done only if the upstream server responds with 502, 503 or 504
responses (Included in 5xx)

"connect-failure" : Retry will be done if the request fails because of a connection failure to the
upstream server (connect timeout, etc.). (Included in 5xx)

"refused-stream" : Retry is done if the upstream server resets the stream with a REFUSED\_STREAM
error code (Included in 5xx)

"retriable-4xx" : Retry is done if the upstream server responds with a retriable 4xx response code.
The only response code in this category is HTTP CONFLICT (409)

"retriable-status-codes" : Retry is done if the upstream server responds with any response code
matching one defined in retriable\_status\_codes field

"reset" : Retry is done if the upstream server does not respond at all (disconnect/reset/read
timeout.)

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 7,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 7,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3213001230013032-0020221021322333-1332330301110002-0020303010112211-3321113230333322-3203333212321332-0203333021032210-3112130033232321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `retry_policy.back_off` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [retry_policy](data-sources--virtual_host--reference--group-002.md#canonical-1021031030330102-2100323023122311-3003021113211121-2312223333303303-2002222022211201-0131100013001013-2031022313222112-2203022023311021)
- retry_policy.back_off

<a id="canonical-2100010313230133-0131312130011120-3021201113200123-0111123210210102-2313130302102213-3321231130312130-2130210232001211-0132213311231233"></a>

Type: `"single"`. Computed.

Specifies parameters that control retry back off.

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

<a id="canonical-2000110230123330-1330333310310233-3312013201322233-2002030231320211-1123312310300003-1023320323223102-0303032322202000-3030311033223201"></a>

### Direct properties for `retry_policy.back_off`

<a id="canonical-2203132023322013-0121130211022203-3232303101100320-3232300022223013-3121331123220320-2021300332201203-3201002020331213-1130003023311201"></a>

#### `retry_policy.back_off.base_interval` property

Type: `"number"`. Computed.

Specifies the base interval between retries in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

<a id="canonical-2230122033122210-3301212221320032-2101022132012321-0313311303320300-2030013311022231-1011313330031023-0122310203312220-2101300302001102"></a>

<a id="canonical-1301320000002222-1232201300111020-3322113010030122-3320312110330131-1201011112233322-3002312003213322-2312302101212033-2311010100121120"></a>

#### `retry_policy.back_off.max_interval` property

Type: `"number"`. Computed.

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The times the base\_interval. Defaults to
\`10\`.

Additional upstream details:

The default is 10 times the base\_interval.

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

<a id="canonical-3211011200123230-0330211200101113-1112022302023320-3120120232213300-3223122013100230-0003102201231003-1303322122132321-0212222131122110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- routes

<a id="canonical-2002112323120121-1202312220122120-3322103020012232-0302122212212332-0222332110320333-1213010203011123-0000002223313331-0233310011312332"></a>

Type: `"list"`. Computed.

HTTP routing rules that match incoming requests based on path, headers, or query parameters and
forward them to appropriate backend origin pools.

Additional upstream details:

The list of routes that will be matched, in order, for incoming requests. The first route that
matches will be used. Currently route object is redundant in case of TCP proxy but required. For
TCP\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY VirtualHosts, the route object only specifies the
cluster/weighted-cluster as route destination without any match condition. In other words, match
condition in route object is ignored for TCP\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY VirtualHosts.
Routes used for TCP\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY VirtualHosts cannot have DirectResponse
or Redirect as actions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-3230102101023212-3022333312221101-0030300303202311-2020110231221212-3113313303121221-1033112321330123-0133002210121201-2030131231330023"></a>

### Direct properties for `routes`

<a id="canonical-1223020302332210-2023312112223323-0221101131313303-3023200210310011-2031033021232212-1012021110032002-0130032212230020-3222220230002003"></a>

#### `routes.kind` property

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

<a id="canonical-2202300320101110-0113011022002232-2110123233001222-1131030303001313-3232302022013011-3223233110113210-2012021133120233-1111023032031320"></a>

<a id="canonical-3203322102132213-1111322321310301-1100311202212200-1002331001023332-3100222113330012-1223221103301122-1022131022223002-2222130230202232"></a>

#### `routes.name` property

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

<a id="canonical-0210112210122000-0322130213000103-2333231131123030-2030311131121213-3031222313311233-3123322202222310-2211113012232110-1123030213303122"></a>

<a id="canonical-3322201321303032-1311131202310111-3322323222213033-0311210320303203-1211123332131203-1002020231320033-3031313232223331-2003200310131023"></a>

#### `routes.namespace` property

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

<a id="canonical-3320303210103110-2113232323122122-2230323000030202-1112313303333200-1000223103110013-1213132102221031-2133102330130322-1031112033033212"></a>

<a id="canonical-1303313322012321-1013332220121211-0331310220223333-2333203303123000-2012031200222000-1121120013123030-1003101330312133-3121210013110300"></a>

#### `routes.tenant` property

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

<a id="canonical-3123331032100112-0230032013120120-0232012012302122-3120122102110232-3332303313212120-2022020010212033-1322322303230002-2123032322023220"></a>

<a id="canonical-3030320121131210-3312333323200112-3321322001233210-2203302323001123-2223011031013022-2013211310033033-2322123001111122-3030103312012213"></a>

#### `routes.uid` property

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

<a id="canonical-1111130000112110-1102321322100210-0013310223010100-3213130303222101-3131013201101301-1000021000133131-0023212331022302-0120103003321211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sensitive_data_policy` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- sensitive_data_policy

<a id="canonical-2113012213211213-0001300300111232-0000020003302202-0323031131213222-2010031320310213-2023110002103000-0233002333102213-0010303110201000"></a>

Type: `"list"`. Computed.

Policy configuration for this feature.

Additional upstream details:

References to sensitive\_data\_policy objects.

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

<a id="canonical-3222233231020111-2022011320010212-1012001323201000-1132330011301331-3220031211121323-0220100011220012-1233320031013023-1123123312201310"></a>

### Direct properties for `sensitive_data_policy`

<a id="canonical-1131233313023311-2110020033320000-0020311223322131-0211222333323230-0030120022330022-1210321330303210-2322023311132123-1021330022311000"></a>

#### `sensitive_data_policy.kind` property

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

<a id="canonical-3323132320003002-0020103200100221-1223211013120003-0123101131101131-0102111300011232-1020000320333220-0322022111123232-0233331300323131"></a>

<a id="canonical-3221212123013022-2100211010103012-3013312231030201-2331233022013101-3020100323032100-3102102001212230-1321202101333121-3322311202030303"></a>

#### `sensitive_data_policy.name` property

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

<a id="canonical-1230101132333321-0231110301231201-3311323123113322-0231331010203132-3310020311101021-2013203010312033-0300220112033322-2000021123301300"></a>

<a id="canonical-1233020121002023-0221332232312312-2022010333031333-3331113332133113-1101233021323020-0332021320201223-3333111103133120-3120221112101122"></a>

#### `sensitive_data_policy.namespace` property

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

<a id="canonical-1023212213333333-2233313120032322-2311231322131322-2113102110330011-1320113113102000-3031303012020202-2232220233102131-1131030333123211"></a>

<a id="canonical-2233032231023330-2110200032210123-3022301100230030-0000211321121332-0002211013203330-0322223203111201-3023211333000102-3032032323230011"></a>

#### `sensitive_data_policy.tenant` property

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

<a id="canonical-2212202010100222-2331323201130111-2313301113223320-2323302021000202-1120203311123113-3330231210332120-1023131010203020-0330312323130012"></a>

<a id="canonical-3222200200111320-3022111311220023-3333333333233111-3210001203101130-0020213323203133-1330323211301320-2013231132111132-0000103032110021"></a>

#### `sensitive_data_policy.uid` property

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

<a id="canonical-2020302310011113-0110101030032202-1233233331203200-0321332220211010-1220330200213312-3321332320232003-2113021203322331-0330021133123211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slow_ddos_mitigation` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- slow_ddos_mitigation

<a id="canonical-1220330111031211-1120011223033003-2123202222131321-1201302331203110-1333233222300203-1200012010311001-1323233031232331-0100332123023100"></a>

Type: `"single"`. Computed.

'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Additional upstream details:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

<a id="canonical-2301331131323101-2332321331302031-2001303103323101-3120032002302023-2013020302222030-0131230023022323-3013033300112120-3232023202313133"></a>

### Direct properties for `slow_ddos_mitigation`

- [disable_request_timeout](data-sources--virtual_host--reference--group-002.md#canonical-2332311122002120-0102020211112123-3101022133022100-2132210101231300-2123013103112213-2123200322222131-3202202312030120-1203310031330002): complete subsection reference.

<a id="canonical-2201022000222211-2233111121033301-2331123120010132-0323013203231112-2102022322113110-2313110030333310-0011100210121223-0210320030232231"></a>

<a id="canonical-3120221223013210-1331033002013333-3232013103233331-1020220310303210-3012203322023131-1010203332130230-1032322313002000-1000220110213323"></a>

#### `slow_ddos_mitigation.request_headers_timeout` property

Type: `"number"`. Computed.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Additional upstream details:

The default value is 10000 milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-3230302201331122-3001013330302310-3223123011023012-3101202230323030-3121033310132003-3001132133233333-3011121011033001-0311321300131312"></a>

<a id="canonical-0020301313200002-0121023030333301-1310032111003031-2122310033000320-2310123001233200-0030022302112232-2100230131103201-1323122230312202"></a>

#### `slow_ddos_mitigation.request_timeout` property

Type: `"number"`. Computed.

Exclusive with \[disable\_request\_timeout\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

<a id="canonical-2332311122002120-0102020211112123-3101022133022100-2132210101231300-2123013103112213-2123200322222131-3202202312030120-1203310031330002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slow_ddos_mitigation.disable_request_timeout` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [slow_ddos_mitigation](data-sources--virtual_host--reference--group-002.md#canonical-2020302310011113-0110101030032202-1233233331203200-0321332220211010-1220330200213312-3321332320232003-2113021203322331-0330021133123211)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-3202112331010000-2313202003002100-0032311010100012-3120320312113322-2121103100232313-1010200233110313-3312323002331313-2323013233203122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable request timeout.

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

<a id="canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- tls_cert_params

<a id="canonical-3032031001222133-3201022301111333-2323000130331230-3210112110333013-1013322112101202-3313113011233020-3302022013222012-2000301331222210"></a>

Type: `"single"`. Computed.

\[OneOf: tls\_cert\_params, tls\_parameters\] Certificate Parameters for authentication, TLS
ciphers, and trust store.

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

OneOf alternatives in this subsection:

- [tls_cert_params](data-sources--virtual_host--reference--group-002.md#canonical-3032031001222133-3201022301111333-2323000130331230-3210112110333013-1013322112101202-3313113011233020-3302022013222012-2000301331222210)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-0013030212110331-2322313130021202-1331130021032333-0333100022023122-2301331212332300-0032103210001031-1030010233201111-2231131010212101)

Select alternatives according to the provider validators above.

<a id="canonical-3032203020232210-3103120322031031-1323023031121331-3011331010033330-1203122030332300-3032333111032233-2301330211232002-2121310110301003"></a>

### Direct properties for `tls_cert_params`

- [certificates](data-sources--virtual_host--reference--group-002.md#canonical-2330330032220120-0113103120120230-2302311102231101-2311310130322120-3110012121032233-3201221331111220-0301123210033033-2101230003011212): complete subsection reference.

<a id="canonical-3300013223123001-2110202103320120-1200121330023013-3332001132021311-0220123222132132-2203303231001013-0331100132323300-1331320322002033"></a>

<a id="canonical-3010320001323333-0200020303021300-1301223211302232-2302203323103111-3110131011120003-1222111012321010-3100212222100232-2003320021303012"></a>

#### `tls_cert_params.cipher_suites` property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [client_certificate_optional](data-sources--virtual_host--reference--group-002.md#canonical-2122001113001222-2122230031211023-0223212321103110-3000231312023013-2003222030013130-2013010331200302-1212033102132002-2020300113132312): complete subsection reference.

- [client_certificate_required](data-sources--virtual_host--reference--group-002.md#canonical-0001303310223111-1131201302232133-0110320020000022-1301012001003001-0200232012010012-1203002002323312-0012100122102031-3220031111002110): complete subsection reference.

<a id="canonical-1131110330012331-2221211112212321-3110313301221231-2330203202132300-3323300020312333-2122033010223003-3112221320102130-2323031130220211"></a>

<a id="canonical-0133302133233210-0023111101131330-3121000231210323-1213130221010101-3011022033323312-0122001003022131-1301120012103201-0023003031130101"></a>

#### `tls_cert_params.maximum_protocol_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-0101030003010310-1113103230302113-3230113031002121-1100010200200232-2003020012321302-0223211012123202-1001212101020032-3201222102003001"></a>

<a id="canonical-0102002032302330-2300331321230202-0221222121310030-2012321323101233-3033330133010000-2102301030231013-2030233020302011-1303212302210033"></a>

#### `tls_cert_params.minimum_protocol_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

- [no_client_certificate](data-sources--virtual_host--reference--group-002.md#canonical-0320000122121123-2012030032231200-3120023321321110-2033031203310332-3133231211310321-2221301133132232-1231103323331112-2020300331231203): complete subsection reference.

- [validation_params](data-sources--virtual_host--reference--group-002.md#canonical-3213003322230222-1311331232300133-0222202123333301-3000203322302130-2210010031322323-3023210333303300-3333200301120023-3212003022203233): complete subsection reference.

<a id="canonical-3310302001020000-3312322010233310-2000321003231101-3322201300010222-2300321120032212-3012303220211021-0313022203333323-0001021210021231"></a>

<a id="canonical-1213321021131230-0311032021310310-3000131303023202-1121113221201110-0212233320213103-2003012023300032-3033010103033100-0320023222120310"></a>

#### `tls_cert_params.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added. Possible values are \`XFCC\_NONE\`, \`XFCC\_CERT\`,
\`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to \`XFCC\_NONE\`.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2330330032220120-0113103120120230-2302311102231101-2311310130322120-3110012121032233-3201221331111220-0301123210033033-2101230003011212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_cert_params](data-sources--virtual_host--reference--group-002.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- tls_cert_params.certificates

<a id="canonical-3221233032310032-1212121233210100-0122010102321131-0033111303303211-1221232112012301-0301101013322103-0011110230221111-0213323331330132"></a>

Type: `"list"`. Computed.

Certificates. Set of certificates.

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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2132222001330203-0023103131102110-3331313113302123-2330231312230203-0233020101000123-0122120123011312-0020123323302213-1233222121020111"></a>

### Direct properties for `tls_cert_params.certificates`

<a id="canonical-2331312322233001-3121000113203011-2233010221331311-2322002221201221-2110010232321003-2233300323021321-3302121113321032-1333213011131300"></a>

#### `tls_cert_params.certificates.kind` property

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

<a id="canonical-0011133212323013-2121300212212320-1130201002031120-2301132023313210-1112312113321130-1132330032332313-1013222223323111-3123320231112223"></a>

<a id="canonical-3110300203303001-0111323332110230-1110333312210031-3301113132111203-2120001122203000-2303222322301331-1330222220012323-3202131112232203"></a>

#### `tls_cert_params.certificates.name` property

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

<a id="canonical-2010313321322010-2113320011031333-0202230332321300-0312330122133123-1003202103330303-3020311211200213-1030130312102012-3030113303203323"></a>

<a id="canonical-1021231123003122-3333003321332212-1111101121032230-3121130101102022-3320121230221202-3000323200213012-0131330102220121-0201232013001123"></a>

#### `tls_cert_params.certificates.namespace` property

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

<a id="canonical-2302323213233122-0002011332320311-3000323113031301-3132302330131222-2311233210310120-1333323030112103-3113001012102333-3003003202101332"></a>

<a id="canonical-0210201112101232-1013022012320001-0100311000202202-2113231130120022-0032303203202330-0323030102300221-2213231210210100-1332022131110210"></a>

#### `tls_cert_params.certificates.tenant` property

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

<a id="canonical-3311121320100133-1131023023331100-1312203313001132-3101201202113022-3123230123323222-3011311301200213-1203022211211210-3003213132012223"></a>

<a id="canonical-3120323110202011-2133021023121201-0120023310312333-1020031303310301-3200223221011220-2232120203330332-1001013230000113-1002133331310303"></a>

#### `tls_cert_params.certificates.uid` property

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

<a id="canonical-2122001113001222-2122230031211023-0223212321103110-3000231312023013-2003222030013130-2013010331200302-1212033102132002-2020300113132312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params.client_certificate_optional` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_cert_params](data-sources--virtual_host--reference--group-002.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- tls_cert_params.client_certificate_optional

<a id="canonical-0021020200323331-1101002130221300-3220121322010201-3023233303300022-0013223200023131-0221100021322011-1033001331022321-2131111100333022"></a>

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

<a id="canonical-0001303310223111-1131201302232133-0110320020000022-1301012001003001-0200232012010012-1203002002323312-0012100122102031-3220031111002110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params.client_certificate_required` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_cert_params](data-sources--virtual_host--reference--group-002.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- tls_cert_params.client_certificate_required

<a id="canonical-1310331110032112-3002212333220332-0011213300131032-1200013023013220-3212021223313301-0302122231011101-1321020103130121-2013232113023233"></a>

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

<a id="canonical-0320000122121123-2012030032231200-3120023321321110-2033031203310332-3133231211310321-2221301133132232-1231103323331112-2020300331231203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params.no_client_certificate` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_cert_params](data-sources--virtual_host--reference--group-002.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- tls_cert_params.no_client_certificate

<a id="canonical-0210111112322023-2302233322300033-2211201100101112-3302123300112212-0022310311221102-1112120223020133-1232222230213123-2012112021333133"></a>

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

<a id="canonical-3213003322230222-1311331232300133-0222202123333301-3000203322302130-2210010031322323-3023210333303300-3333200301120023-3212003022203233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params.validation_params` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_cert_params](data-sources--virtual_host--reference--group-002.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- tls_cert_params.validation_params

<a id="canonical-2300302222301320-1013220320220122-2011112013302323-0003012133101110-2211222202003322-3003231101333301-1321221232232330-1330130311101123"></a>

Type: `"single"`. Computed.

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

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

<a id="canonical-2111212331113332-0233210310121303-2332232113233223-1221022331031110-2313101331021210-3121231122222122-2232131002202311-1311301322230203"></a>

### Direct properties for `tls_cert_params.validation_params`

<a id="canonical-3320123120020312-2330230202230021-3222031121223112-0033231310023012-0121101333011131-3223211331322313-3003333221300120-1220202221101312"></a>

#### `tls_cert_params.validation_params.skip_hostname_verification` property

Type: `"bool"`. Computed.

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

- [trusted_ca](data-sources--virtual_host--reference--group-002.md#canonical-1023311023232233-3220101100112021-1002203003331330-3333121303303220-2010202202032031-2331011122033221-1232000233230233-2221230013312223): complete subsection reference.

<a id="canonical-3121111203320320-3330323231303011-3031120310021332-2112002002112231-0030031211132121-3312023201203212-0233001313311131-0011311310103012"></a>

<a id="canonical-0210310022322000-1022103300003321-3102212221221221-3110221033131010-0112113011133031-2023200131302322-1112112013100003-0023231123010002"></a>

#### `tls_cert_params.validation_params.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-0202011000311301-2131110123021013-3132112300332303-2121210000202202-1221313221210323-3223123000000313-2122012321133201-3030321130313212"></a>

<a id="canonical-2101022131300022-3020032123310010-0332132331102321-1302003112011233-2302200221102303-2011221322233101-1023013331323220-3023321113013313"></a>

#### `tls_cert_params.validation_params.verify_subject_alt_names` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-1023311023232233-3220101100112021-1002203003331330-3333121303303220-2010202202032031-2331011122033221-1232000233230233-2221230013312223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params.validation_params.trusted_ca` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_cert_params](data-sources--virtual_host--reference--group-002.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- [tls_cert_params.validation_params](data-sources--virtual_host--reference--group-002.md#canonical-3213003322230222-1311331232300133-0222202123333301-3000203322302130-2210010031322323-3023210333303300-3333200301120023-3212003022203233)
- tls_cert_params.validation_params.trusted_ca

<a id="canonical-2120312133211320-0213010030203123-2212023213213332-3301130020302332-2300100101203032-2002001113213200-1120233211121010-0122112323012112"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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

<a id="canonical-1311023033203213-0323331232233220-2320020112320222-0321132201323110-2032111212132021-3030000101012100-2223202012002012-1112133223131033"></a>

### Direct properties for `tls_cert_params.validation_params.trusted_ca`

- [trusted_ca_list](data-sources--virtual_host--reference--group-002.md#canonical-3123033001322311-3130310302032023-3021020200222003-3012011023123120-3210013130231123-2200312311130121-3212320230301331-3110222012020102): complete subsection reference.

<a id="canonical-3123033001322311-3130310302032023-3021020200222003-3012011023123120-3210013130231123-2200312311130121-3212320230301331-3110222012020102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params.validation_params.trusted_ca.trusted_ca_list` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_cert_params](data-sources--virtual_host--reference--group-002.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- [tls_cert_params.validation_params](data-sources--virtual_host--reference--group-002.md#canonical-3213003322230222-1311331232300133-0222202123333301-3000203322302130-2210010031322323-3023210333303300-3333200301120023-3212003022203233)
- [tls_cert_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-002.md#canonical-1023311023232233-3220101100112021-1002203003331330-3333121303303220-2010202202032031-2331011122033221-1232000233230233-2221230013312223)
- tls_cert_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-2322201223331011-1022202131013311-0012313030331113-1201110100110101-0000101131113030-2333112220310021-3010212122021220-0000300221301301"></a>

Type: `"list"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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

<a id="canonical-3010231220021020-1222323132123132-2000232110231222-1232102331221001-0031100130310022-0000012112111120-1123031332330333-1331122122321111"></a>

### Direct properties for `tls_cert_params.validation_params.trusted_ca.trusted_ca_list`

<a id="canonical-2322133201223202-0211300020200333-3112010321130301-1023303210332031-3130203131022211-3133123200323202-1301003231003211-2103110113320300"></a>

#### `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.kind` property

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

<a id="canonical-0203111133022121-0012200212331321-3333021230121123-2012003010301332-1033013301210213-0202330201300022-0231302103310221-1220120102102220"></a>

<a id="canonical-3220122022302233-1131210313211220-1003313320002020-2321002122322231-2202021003200110-1031311112130231-1111001131033221-1213220132021000"></a>

#### `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.name` property

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

<a id="canonical-0020100332200221-2113321020230013-0120130211110000-2022223031332200-3103132231131033-2301330123211133-1010301332010300-1020313311200332"></a>

<a id="canonical-2000122221333003-2130033021130200-0213333112312000-0231131230323112-0122201211310103-3212333010031001-0302311300230303-3001113301233000"></a>

#### `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.namespace` property

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

<a id="canonical-3210031221311021-0313012001202012-1220110122213101-2321102223232113-0223032000021133-1002301012003020-3311232103100303-0022010230332100"></a>

<a id="canonical-2110310002131022-2320120132201333-2021001330213010-1013222133021232-3312133022132200-2032132203132201-1331023100202113-1120230130312112"></a>

#### `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.tenant` property

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

<a id="canonical-3301020111002020-3320012021321331-0131013011312100-1002220220322213-0212131123100033-0320001330223211-1230002230201103-1200233010300100"></a>

<a id="canonical-1301110310023213-1301213123110300-1301130331200331-2021122021320020-2212200103023111-3123200213322032-0233132210010233-0100221123221000"></a>

#### `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.uid` property

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

<a id="canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- tls_parameters

<a id="canonical-0013030212110331-2322313130021202-1331130021032333-0333100022023122-2301331212332300-0032103210001031-1030010233201111-2231131010212101"></a>

Type: `"single"`. Computed.

TLS configuration for downstream connections.

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

<a id="canonical-0322032232033230-3102300022321212-1003102323301111-3100012323013120-3020232132220132-2232232332300101-0030101210232133-2123112000030200"></a>

### Direct properties for `tls_parameters`

- [client_certificate_optional](data-sources--virtual_host--reference--group-002.md#canonical-0311213131212331-3033101301300130-0312230321201212-1030121302231320-2031103222211101-2130230222311110-1331212232232032-3232101311001120): complete subsection reference.

- [client_certificate_required](data-sources--virtual_host--reference--group-002.md#canonical-3100311320103231-0130313003300212-2111010023132233-1002011031210200-1123131203311210-3013330113032213-1210102122213333-3201332021102312): complete subsection reference.

- [common_params](data-sources--virtual_host--reference--group-002.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111): complete subsection reference.

- [no_client_certificate](data-sources--virtual_host--reference--group-002.md#canonical-0003223203201211-1212131120130013-0331213103230031-2310001021211300-2022120121120211-0331000312213003-2223103133010011-2001303022303313): complete subsection reference.

<a id="canonical-3200332110312223-1011311221201220-1310300302322102-3213122331332330-1330232002303333-0010023311221110-3221003021203101-1030330303010122"></a>

<a id="canonical-2333102111322322-0322111203333221-1202102333330122-1213031021332203-3130001012023013-1320111020331032-2023110000030300-0132012133333111"></a>

#### `tls_parameters.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added. Possible values are \`XFCC\_NONE\`, \`XFCC\_CERT\`,
\`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to \`XFCC\_NONE\`.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0311213131212331-3033101301300130-0312230321201212-1030121302231320-2031103222211101-2130230222311110-1331212232232032-3232101311001120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.client_certificate_optional` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- tls_parameters.client_certificate_optional

<a id="canonical-3221212102110032-2212333010322130-1110301031211112-0233300112310032-1220113321013000-3223321211103002-1021230021203211-3100311011212110"></a>

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

<a id="canonical-3100311320103231-0130313003300212-2111010023132233-1002011031210200-1123131203311210-3013330113032213-1210102122213333-3201332021102312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.client_certificate_required` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- tls_parameters.client_certificate_required

<a id="canonical-1322121022021001-1211013131003133-2233122232330113-2203121213300003-3132332233300311-2211003133210313-0312020003120012-2020002303032033"></a>

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

<a id="canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- tls_parameters.common_params

<a id="canonical-1122233013102130-2021122230220223-3002031003230030-3020321300000222-1322010202110003-2122231302323200-2203220302313131-2323311001200032"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2210302220330121-3112232323121230-0220021100300021-3001103311033112-0221103221332023-1020333100232310-3320011121320333-3223321121031012"></a>

### Direct properties for `tls_parameters.common_params`

<a id="canonical-0210103132203322-1000023122230330-1203201210212103-2301100121230113-0221302200221113-2212212100332303-0021112331001121-3310301321303211"></a>

#### `tls_parameters.common_params.cipher_suites` property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3311002010103122-1013132310120330-0311130211112101-1233333323023213-2023300300221103-2123210233300320-1331211100323023-0123321210110212"></a>

<a id="canonical-0122231313122002-1313333311300302-1223101123210031-1011110033201333-1332001032221220-0203031302311312-3300020100302022-3312311131120101"></a>

#### `tls_parameters.common_params.maximum_protocol_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-2010113220100203-3122102002113033-1302302001303030-1203203331122233-1122311133000133-1033311000100001-0302011121113000-0322333003012230"></a>

<a id="canonical-0232132103112131-3011320300023121-2202130120121330-0000323231031131-1320020030002211-3212003212203033-0320333300230133-2231233132331303"></a>

#### `tls_parameters.common_params.minimum_protocol_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

- [tls_certificates](data-sources--virtual_host--reference--group-002.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312): complete subsection reference.

- [validation_params](data-sources--virtual_host--reference--group-002.md#canonical-1011200130001331-1320301220122001-3310301011012333-1301113202111223-0010221100002031-1203321002203131-1110303301010000-2102112302131230): complete subsection reference.

<a id="canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-002.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- tls_parameters.common_params.tls_certificates

<a id="canonical-3122211203131331-1110033333310310-1013132323031321-2023210123002111-2321322320000113-3321322313230001-2322021021212302-0321000021133001"></a>

Type: `"list"`. Computed.

TLS Certificates. Set of TLS certificates.

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

<a id="canonical-3121222333010020-3123123020332321-3303133031312130-0033212232231303-3320232200332303-2231233333023111-0222031222323302-1211323021332203"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates`

<a id="canonical-0130131200233131-3102001332233101-3111311001312113-0021333333233323-3201133301032220-3131301230200012-1213301010211021-3300123211231210"></a>

#### `tls_parameters.common_params.tls_certificates.certificate_url` property

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](data-sources--virtual_host--reference--group-002.md#canonical-2013221330301301-0200121011220020-0131100030133132-3111220100302013-0323301013010002-0032023213122022-1113030220331103-3321222302033332): complete subsection reference.

<a id="canonical-1112213221331013-1033133021011102-3230011200010333-1130220130001201-1201303333121131-0110032021011112-2132112133001020-0221123210030120"></a>

<a id="canonical-1202211223012312-0200202033201113-0300331030031213-1001320200311032-2031020232223210-2102333131111310-3332303320301020-1030010123013130"></a>

#### `tls_parameters.common_params.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--virtual_host--reference--group-002.md#canonical-1210223002300101-1110011013032331-1312303120101220-2322030232030110-3011330011000030-3022000300311013-3330122120300020-0310201103231212): complete subsection reference.

- [private_key](data-sources--virtual_host--reference--group-002.md#canonical-0032311131222331-3223213221002220-3130303232312133-0313321011032200-2101111120023012-2130010110101200-0311300221011321-3103120121101222): complete subsection reference.

- [use_system_defaults](data-sources--virtual_host--reference--group-002.md#canonical-2032000021233011-0333311310031021-0332130220312011-3322222321000123-1322322122132013-1131310310222201-0323023113221212-3333200213120130): complete subsection reference.

<a id="canonical-2013221330301301-0200121011220020-0131100030133132-3111220100302013-0323301013010002-0032023213122022-1113030220331103-3321222302033332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-002.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-002.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-1013100311311122-3133332231121321-0011003022010013-0022130200031322-1013213110331100-1302203032133203-2132003020022003-2203222120223211"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

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

<a id="canonical-1011313203221101-2221300220022132-0322011002023203-2310110311103232-2030323202110033-3312223101332312-2321010013010031-2331002031132020"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.custom_hash_algorithms`

<a id="canonical-0333323320231131-3112231210321031-1101203220011032-1030113210033133-0032022113023211-3302110133201303-1031332300002001-1221300003002300"></a>

#### `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

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

<a id="canonical-1210223002300101-1110011013032331-1312303120101220-2322030232030110-3011330011000030-3022000300311013-3330122120300020-0310201103231212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-002.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-002.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-1032000020301032-1220002203030322-3011131031012231-2021333010023032-1333201221002002-2120120103230222-3022023201322203-2202221221002310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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

<a id="canonical-0032311131222331-3223213221002220-3130303232312133-0313321011032200-2101111120023012-2130010110101200-0311300221011321-3103120121101222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-002.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-002.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-0132213332013300-3330132103202111-1123122112012003-2220302132012101-2032311122122211-2212023200322032-2310023213313333-3320300322330310"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-2132311033130203-0113311301301332-2132110131310201-0203322132023203-0203110122302122-1001212200023030-1211320211112003-2113203211300310"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-0212231131312301-3010010002131122-3333002010103203-0011203120313012-0133100230302122-3033100201321221-1331232010020232-2001310111231130): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-1030232231122130-3312120101201033-2312010133121330-3313001200230320-2020131003022200-1220101032132000-0313312030132203-3101322303003323): complete subsection reference.

<a id="canonical-0212231131312301-3010010002131122-3333002010103203-0011203120313012-0133100230302122-3033100201321221-1331232010020232-2001310111231130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-002.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-002.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--virtual_host--reference--group-002.md#canonical-0032311131222331-3223213221002220-3130303232312133-0313321011032200-2101111120023012-2130010110101200-0311300221011321-3103120121101222)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-1122221313200001-1233111222132323-2102220222033021-3301220113113232-2223011220201030-0020102100120322-0331001311223303-3303202102232133"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-1300233323212200-1333322023302220-0322203211220230-3130013023110002-3211222203121121-3302011020132220-0030322100221111-0330232230210212"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-2330012103000311-3301320210020113-1030130102020231-1300322021030020-1003010232212003-1002333231011310-2311231013322011-1322010331122310"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-0123303221223230-0110023130232133-1122313022202321-1321230303231133-2131013033210032-1100301301221121-1113300011322213-1131023312022123"></a>

<a id="canonical-2320223233303203-0000231323022021-3300033301013312-0312330131321222-2022023122213032-3123023200033332-2301022223212322-1233210321002033"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2231210311332313-1112022311211300-0200211232013110-1313032332000221-1211331230110022-3312311202120110-3022202310100322-3313201302021230"></a>

<a id="canonical-2130130321023102-0010010331022232-1130013323331001-2313232301030033-3232022323322201-1213221211032103-1133201311322020-3130330131322132"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-1030232231122130-3312120101201033-2312010133121330-3313001200230320-2020131003022200-1220101032132000-0313312030132203-3101322303003323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-002.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-002.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--virtual_host--reference--group-002.md#canonical-0032311131222331-3223213221002220-3130303232312133-0313321011032200-2101111120023012-2130010110101200-0311300221011321-3103120121101222)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-0102230201001020-3303302322012223-3222113230100120-0223231012231200-0303203121101011-1131310323021032-3030033333011100-3333000032000002"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-1200102131112001-0222030130123232-2103122232012230-1200202220011311-2003010001322301-3230022320333303-3100113323031001-1303222300221002"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info`

<a id="canonical-3332122332302233-2300202203032313-0113122321220012-2122120032220312-2103212212222201-2001123100233031-1232222020320132-0331300103212101"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0323302110030333-1013121220021123-1301121113031233-3131020213232021-3102311233211130-2311211201200132-1100031320002011-3221332202201013"></a>

<a id="canonical-1202201003000003-2112233121100220-2103212210121312-3332000331323312-2011121221000012-0313011010221302-1332300131212111-0232000222112223"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2032000021233011-0333311310031021-0332130220312011-3322222321000123-1322322122132013-1131310310222201-0323023113221212-3333200213120130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-002.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-002.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-2110303302121311-1330010201013331-1133321110131010-2033121213220212-3012002021330332-0133311012231022-2233202131232210-3203230332212322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

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

<a id="canonical-1011200130001331-1320301220122001-3310301011012333-1301113202111223-0010221100002031-1203321002203131-1110303301010000-2102112302131230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-002.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- tls_parameters.common_params.validation_params

<a id="canonical-2212012212101320-3112001032323123-2033201010100210-2133103003303313-3033022003211212-0213131112032203-2233331211102110-2032033120003310"></a>

Type: `"single"`. Computed.

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

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

<a id="canonical-1030133202311123-2110303220030213-2202021113323220-0110202123233120-3331222130121123-3010331022332022-2321332301333213-2100021112022231"></a>

### Direct properties for `tls_parameters.common_params.validation_params`

<a id="canonical-2030113320233213-0220133300211011-3122221331030231-1012202013002210-1200223021303102-3103213121121100-0323123213331301-0202123200321110"></a>

#### `tls_parameters.common_params.validation_params.skip_hostname_verification` property

Type: `"bool"`. Computed.

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

- [trusted_ca](data-sources--virtual_host--reference--group-002.md#canonical-3333331333311130-0131012003032211-0001030100130121-3211202103121133-2033233301023120-2201133131321102-3012333212020221-1322102322122012): complete subsection reference.

<a id="canonical-2221202123013033-0301222212222230-0002103322012211-1223210100330122-3332211023131310-0300223232202130-1310311203231300-2301201132223210"></a>

<a id="canonical-2013210020012333-1022113233222122-1331322312032132-1220311010001001-3223023332201023-2202203310322300-1300023333223011-1333312332132302"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-2030332013012000-2133110033200303-2303032112332220-0031023223132301-2031331013200132-3121311113000130-3113133002323002-3230013220202200"></a>

<a id="canonical-1222021223112312-0323330320132321-2023210202110003-3311231032002220-3231100011210322-3311120032002102-0212202323111003-0003301323000131"></a>

#### `tls_parameters.common_params.validation_params.verify_subject_alt_names` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-3333331333311130-0131012003032211-0001030100130121-3211202103121133-2033233301023120-2201133131321102-3012333212020221-1322102322122012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params.trusted_ca` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-002.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.validation_params](data-sources--virtual_host--reference--group-002.md#canonical-1011200130001331-1320301220122001-3310301011012333-1301113202111223-0010221100002031-1203321002203131-1110303301010000-2102112302131230)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-0110222322332222-3321211030330203-0222013221030130-2231232312323102-1203032313103202-0303220312110330-0330122230200302-0230033112222122"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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

<a id="canonical-1021303302112113-0321213131122131-0331122302322310-1102213331122302-1232300000103100-3330001302032103-3221232022332120-3320203323012220"></a>

### Direct properties for `tls_parameters.common_params.validation_params.trusted_ca`

- [trusted_ca_list](data-sources--virtual_host--reference--group-002.md#canonical-2212123223300122-3011100310011320-1221112011132033-1101133222300003-0211321302002122-1121130121210322-0330000210323101-3321132332322301): complete subsection reference.

<a id="canonical-2212123223300122-3011100310011320-1221112011132033-1101133222300003-0211321302002122-1121130121210322-0330000210323101-3321132332322301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-002.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.validation_params](data-sources--virtual_host--reference--group-002.md#canonical-1011200130001331-1320301220122001-3310301011012333-1301113202111223-0010221100002031-1203321002203131-1110303301010000-2102112302131230)
- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-002.md#canonical-3333331333311130-0131012003032211-0001030100130121-3211202103121133-2033233301023120-2201133131321102-3012333212020221-1322102322122012)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-1033000020332202-0023201321102003-0021322013210030-0121310110201111-0200220132013113-3211000033320120-1110013120222300-1321100223231013"></a>

Type: `"list"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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

<a id="canonical-2331333022120333-0223313300210322-1230122222022012-3321310030012233-3131230110233013-2133233232103102-2102311121001211-0221220122013313"></a>

### Direct properties for `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list`

<a id="canonical-3213102212203012-1011323021323312-2133323212032212-1110122023202100-3202131112120123-3220230301330321-3033312132230303-3133002212001320"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` property

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

<a id="canonical-3321030310313130-0123031000211330-3011131020210111-0112032230020130-0113212300003233-1203020021031122-1101202303111221-3022122322311321"></a>

<a id="canonical-0222201130321302-1211002023112132-0032131110011121-0310112302321031-0111323202011212-3232201013121111-3103310033330013-1132202332002122"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` property

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

<a id="canonical-3331123303122121-2031001311331303-2232012031321123-1012101132310002-3201200202011233-1201111103301222-0021112030303200-1333320310013212"></a>

<a id="canonical-3211201030111121-1333331233222232-3232002001002110-1220301201313130-1123333200232312-3331020233312221-0223003112030132-0001321030312003"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` property

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

<a id="canonical-2120122322022312-2330012223232210-0321010113112312-1203331221102222-0221032020003301-0132030133322001-1232000023320122-2202312001130002"></a>

<a id="canonical-1010200120311032-1113131212312230-1133331222322200-3032030200030320-2023233310220210-1333101031111221-1230023212012201-2301022221100123"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` property

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

<a id="canonical-1322130113222233-3221102020013002-2311301131232222-2231011212310112-2133022103213020-0022223130120321-1020223010310320-0213033333200123"></a>

<a id="canonical-2231130003222310-0111303133202102-1120110131210233-1321123311003023-0102213201302120-2002133221131113-1101003201313102-2333213131131213"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` property

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

<a id="canonical-0003223203201211-1212131120130013-0331213103230031-2310001021211300-2022120121120211-0331000312213003-2223103133010011-2001303022303313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.no_client_certificate` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-002.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- tls_parameters.no_client_certificate

<a id="canonical-2003121333303301-2111312111203022-1320320130233311-3033000220012333-0201022301310033-1312123300120330-3201231231111210-3213012001011200"></a>

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

<a id="canonical-1002302233122000-1323022011111121-3132010100020022-1310010213200321-0312233032021223-1020313001300323-0020033031030302-1102203102302301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_identification` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- user_identification

<a id="canonical-3332310012322303-2301321002210033-3200311030002213-3100332131002300-3021122202111201-1333223133120302-0111201333302303-2100033020001110"></a>

Type: `"list"`. Computed.

A reference to user\_identification object. The rules in the user\_identification object are
evaluated to determine the user identifier to be rate limited.

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

<a id="canonical-0231111102103000-2203112333000221-0322033111120212-1201221013221332-2310303022031212-2011131120120313-1132220212230030-0113002120011122"></a>

### Direct properties for `user_identification`

<a id="canonical-1211032333222130-1113032311112331-1322112023200233-2221210102302032-0111132202230233-0313311101113121-2231112101110013-2110120213032300"></a>

#### `user_identification.kind` property

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

<a id="canonical-1302023000122200-1233200213212012-0310002020132131-3212003200033311-1103203002212221-0012230220321313-0101102200232223-2111101323101303"></a>

<a id="canonical-2201330131112031-3031033130000203-3101022211033212-0012131011000132-1232033333230133-1322211110213122-1233321212220030-0032130230322321"></a>

#### `user_identification.name` property

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

<a id="canonical-1222313323133111-2012203111100120-2111302113221303-3130221231213012-0033000012131020-0301120010010010-1321330030230133-0220333320021122"></a>

<a id="canonical-1023233111003003-2333132331232121-0102000032233000-1332033333210033-1022233221021311-3113002002300300-1323023100231201-0120303212032001"></a>

#### `user_identification.namespace` property

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

<a id="canonical-1131031121132022-3200223321130120-3101323011322120-1220023033333233-3030131103211311-0001230112030013-0203121130311021-2311300302211020"></a>

<a id="canonical-3111332323111230-1232020330230002-2123330121110130-2000032302311331-1313221020000222-0211022202303021-0133321231021023-3211010021031000"></a>

#### `user_identification.tenant` property

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

<a id="canonical-0233310231303031-2221213101031123-0033311303301102-1023330030323230-2012313013130211-2322013030310211-2332122002033020-2220002133001131"></a>

<a id="canonical-0001230220010301-1310212332130303-1123133323021110-3003323022023101-0231331312220233-1230320033103311-1223032132011001-0012203130323010"></a>

#### `user_identification.uid` property

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

<a id="canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_type` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- waf_type

<a id="canonical-0212131312112320-3010013102333133-1003210023231021-0022011032133223-0310111333100300-0332102210001333-3113231121103003-0013001311203303"></a>

Type: `"single"`. Computed.

WAF instance will be pointing to an app\_firewall object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_type": "[\"app_firewall\",\"disable_waf\",\"inherit_waf\"]"
}
```

<a id="canonical-2331220111013021-1000033011133212-0000002130333012-2200101323102230-1212233033222220-3212101021123230-2011313022033200-0232111101122331"></a>

### Direct properties for `waf_type`

- [app_firewall](data-sources--virtual_host--reference--group-002.md#canonical-1333300322323113-1232131101000031-2233210011203232-3322023300230033-1233133023310221-1233001022120331-2003112011233330-2313020200000013): complete subsection reference.

- [disable_waf](data-sources--virtual_host--reference--group-003.md#canonical-2230010330302310-1312132323300100-3023320210221023-2030200010011000-1310013232033111-0133103323302222-1021312002012220-2103031000003013): complete subsection reference.

- [inherit_waf](data-sources--virtual_host--reference--group-003.md#canonical-3111320320023322-3033330033011220-2031221222111000-3211122202020221-2303103121031102-0213033333303222-0321322313130302-2022131333210103): complete subsection reference.

<a id="canonical-1333300322323113-1232131101000031-2233210011203232-3322023300230033-1233133023310221-1233001022120331-2003112011233330-2313020200000013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_type.app_firewall` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [waf_type](data-sources--virtual_host--reference--group-002.md#canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001)
- waf_type.app_firewall

<a id="canonical-3000222221111110-0300100323231110-3213130110203022-3102100032020310-1302332210112312-3031031231222333-0210011003021121-0131313212222203"></a>

Type: `"single"`. Computed.

A list of references to the app\_firewall configuration objects.

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

<a id="canonical-1222022012102102-3310130230221003-1103130230102223-0203210130303230-3231301030101110-1122023023213323-1130311020203203-1103120221133101"></a>

### Direct properties for `waf_type.app_firewall`

- [app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-3300211200203031-0102030220211032-2101332013321022-2020233013011313-3111012122302233-3321322111201122-3301232102311203-1002222023222302): complete subsection reference.
