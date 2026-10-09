---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-0032333010121002-0332331033300212-2002033020321220-2332202013303113-2313133231100331-0003222333003100-3133200232230203-0121213102202020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- csrf_policy

<a id="canonical-3000331031023322-3131110331302003-2320301022032003-1010232010300203-1103201100200333-0321313303002220-3023320100320230-1120102030310133"></a>

Type: `"object"`. single nested block, Optional.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host.

&#8203;1. The origin that caused the user agent to issue the request (source origin). &#8203;2. The
origin that the request is going to (target origin). When the policy evaluating a request, it
ensures both pieces of information are present and compare their values. If the source origin is
missing or origins do not match the request is rejected. The exception to this being if the
source-origin has been added to they policy as valid. Because CSRF attacks specifically target
state-changing requests, the policy only acts on the HTTP requests that have state-changing method
(PUT,POST, etc.).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

Terraform syntax:

```terraform
csrf_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223112302012100-1030233033310030-3101103101130100-1011020032312122-3321121302021133-1322010211011211-0321133100233111-3130031131331013"></a>

### Direct properties for `csrf_policy`

- [all_load_balancer_domains](resources--virtual_host--reference--group-002.md#canonical-0110122323300011-3202101110312103-1212210232132200-3123012203220303-0012003220313200-1100012030223133-3011220003003221-2131002333120213): complete subsection reference.

- [custom_domain_list](resources--virtual_host--reference--group-002.md#canonical-0210302321013100-1300220003111020-1330033301211103-3021331002113120-2311332011003021-2100020010022101-3320121011300333-3030122322323001): complete subsection reference.

- [disabled](resources--virtual_host--reference--group-002.md#canonical-1002103013312222-1311000201003223-0013030330132112-0012122333221101-2120131022311012-0030223010122223-0301322123230103-2210301232001223): complete subsection reference.

<a id="canonical-0110122323300011-3202101110312103-1212210232132200-3123012203220303-0012003220313200-1100012030223133-3011220003003221-2131002333120213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy.all_load_balancer_domains` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [csrf_policy](resources--virtual_host--reference--group-002.md#canonical-0032333010121002-0332331033300212-2002033020321220-2332202013303113-2313133231100331-0003222333003100-3133200232230203-0121213102202020)
- csrf_policy.all_load_balancer_domains

<a id="canonical-3133102211130012-2123322312022330-0023111120020320-1231211020211130-3320222020302212-3221122332320201-0312001003320213-0121333032231313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all load balancer domains.

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
all_load_balancer_domains = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210302321013100-1300220003111020-1330033301211103-3021331002113120-2311332011003021-2100020010022101-3320121011300333-3030122322323001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy.custom_domain_list` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [csrf_policy](resources--virtual_host--reference--group-002.md#canonical-0032333010121002-0332331033300212-2002033020321220-2332202013303113-2313133231100331-0003222333003100-3133200232230203-0121213102202020)
- csrf_policy.custom_domain_list

<a id="canonical-3022331111332211-0031131112222202-3131030223330302-0110001000321200-3232202013212100-1320302022121232-3001233213210200-2200323031112011"></a>

Type: `"object"`. single nested block, Optional.

List of domain names used for Host header matching.

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
custom_domain_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221320123123210-0202221320010332-2012002221132103-3230301231003330-2312102231210111-1301201202032122-3011212001103201-2231000332101310"></a>

### Direct properties for `csrf_policy.custom_domain_list`

<a id="canonical-2101300102020301-2101200313113331-0220321031323312-2230003102122130-3301001111312030-3033323303300033-0031233233321032-0321210100033133"></a>

#### `csrf_policy.custom_domain_list.domains` property

Type: `["list", "string"]`. Optional.

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1002103013312222-1311000201003223-0013030330132112-0012122333221101-2120131022311012-0030223010122223-0301322123230103-2210301232001223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy.disabled` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [csrf_policy](resources--virtual_host--reference--group-002.md#canonical-0032333010121002-0332331033300212-2002033020321220-2332202013303113-2313133231100331-0003222333003100-3133200232230203-0121213102202020)
- csrf_policy.disabled

<a id="canonical-3121103112021120-0011030112000311-2302123123110223-3223313113231001-2131230123131131-0202132333333211-1022120220221321-1323331310320122"></a>

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
disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332133000210111-0223310312103001-1010233010211033-1102013020202323-2130113222000300-0232202222221101-2023022020231220-2012203012032313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_header` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- default_header

<a id="canonical-3103110113233101-2311030300030002-3133031023300010-1132231000122223-1300222101112311-3130112321310012-0333210202330121-0301221023322011"></a>

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

<a id="canonical-3121022200033310-1220301130221312-2021030331023200-1312300232120133-1231101321130312-2222300001330221-1303311113210322-3100130111300312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_loadbalancer` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- default_loadbalancer

<a id="canonical-3000212102323033-3010302213033131-1311032222332011-3102311311323303-3231233111020000-3032233222331012-3132021303003223-3203122122312100"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_loadbalancer, non\_default\_loadbalancer; Default: default\_loadbalancer\]
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

OneOf alternatives in this subsection:

- [default_loadbalancer](resources--virtual_host--reference--group-002.md#canonical-3000212102323033-3010302213033131-1311032222332011-3102311311323303-3231233111020000-3032233222331012-3132021303003223-3203122122312100)
- [non_default_loadbalancer](resources--virtual_host--reference--group-002.md#canonical-0203302010201112-0331211102220230-0300132232103103-1210332110032101-1332230330310132-1011232020203331-2233222103311321-3203201030110132)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_loadbalancer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002023231210102-0232311230110330-2231021201011211-3000222000033100-3121230313301113-1303121121223033-1322331301232132-0220030301200000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_path_normalize` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- disable_path_normalize

<a id="canonical-2301113032002302-0100100301320003-2232331201100022-0101013130330333-0310113103302030-1201021302033333-3100320010022221-2210331012103001"></a>

Type: `["object", {}]`. Optional.

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

- [disable_path_normalize](resources--virtual_host--reference--group-002.md#canonical-2301113032002302-0100100301320003-2232331201100022-0101013130330333-0310113103302030-1201021302033333-3100320010022221-2210331012103001)
- [enable_path_normalize](resources--virtual_host--reference--group-002.md#canonical-3220103012323132-3123012003332321-1010200201202130-0001120230103122-2231023300220233-0201200010231213-0212213100033130-1203010112203200)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133031213301110-2320211003030001-3300321001112212-0020212013122303-0311232110311233-2131100131212002-2313301132131310-2201010033313212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_reverse_proxy` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- dynamic_reverse_proxy

<a id="canonical-1313301200300110-2211133100220020-0323222013011332-0203113310011312-0030120202123032-1323133201003113-1312022023331003-1322222111213110"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
dynamic_reverse_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011002020020030-0212221301313021-3013122103011212-3022322103111233-2023002033201320-3233220231303001-3221133203322332-2311303011320113"></a>

### Direct properties for `dynamic_reverse_proxy`

<a id="canonical-0223112013223232-2011132122110100-2330212121311333-3121002322123130-0322113120121013-3302013032003110-2032032123132223-2101203222023123"></a>

#### `dynamic_reverse_proxy.connection_timeout` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [resolution_network](resources--virtual_host--reference--group-002.md#canonical-1130103032111333-2121033223012323-3121001312103121-2320022022021301-2211303233101222-0120030200131130-2310203320231103-0030110032003223): complete subsection reference.

<a id="canonical-2000023100203002-2312211232023010-0113221312023022-3002021230113223-3203303333112002-0222023102001002-1310331300111101-0030332222312033"></a>

<a id="canonical-2231130011332302-2122221230022012-1323301031031232-1313202303202212-1302111123323210-2113202122022100-3111131201312303-1132120003130331"></a>

#### `dynamic_reverse_proxy.resolution_network_type` property

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

<a id="canonical-0121231000330131-1303300020303311-0302201331312312-3002032021323221-2220113112220200-3031100122322131-1333323003310103-1222020232110200"></a>

<a id="canonical-3322002300022212-1010000303323232-1021231332030002-1311113322113300-3022223123220211-1031110003233011-1123321231012122-2023212013100200"></a>

#### `dynamic_reverse_proxy.resolve_endpoint_dynamically` property

Type: `"bool"`. Optional.

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

<a id="canonical-1130103032111333-2121033223012323-3121001312103121-2320022022021301-2211303233101222-0120030200131130-2310203320231103-0030110032003223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_reverse_proxy.resolution_network` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [dynamic_reverse_proxy](resources--virtual_host--reference--group-002.md#canonical-2133031213301110-2320211003030001-3300321001112212-0020212013122303-0311232110311233-2131100131212002-2313301132131310-2201010033313212)
- dynamic_reverse_proxy.resolution_network

<a id="canonical-3102311130233210-3021002322331210-2111323113033013-1100321012211322-0312100303210020-3000301012221103-0123130030223233-0233312230220133"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
resolution_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223331312313210-1203001223032021-3023033330331031-1232332303232112-2202302121213101-0120302111132302-2031022103302200-2303103101200312"></a>

### Direct properties for `dynamic_reverse_proxy.resolution_network`

<a id="canonical-0232123323033222-3031320122332211-2103111102110222-1032122111301213-0211132120111012-1020103311022210-1220303022120122-0131031330002013"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2002233033212032-3120023223122312-3310023130223322-0232202111123311-2320232110300220-1132220213123200-1220303103003112-0133033021121131"></a>

<a id="canonical-1133002210032121-0311000121313101-3103222132231003-2322123333301311-2201211333302332-1033000313022001-1331323211132300-3213010120321220"></a>

#### `dynamic_reverse_proxy.resolution_network.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0031013311021113-2201033233210000-1132220231323112-3213101233000310-2121131123122230-2321103230333332-2230301213131221-1320101332330333"></a>

<a id="canonical-0031320332321231-0121023220010113-3213000311011332-2000330203033113-0113213030001332-1302231332030010-2022100212020131-2022200311103011"></a>

#### `dynamic_reverse_proxy.resolution_network.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0130313230221120-1322312032310331-1032310333133310-0133200231233123-0311212132232230-0220200001130201-0020012000313113-3012312212002231"></a>

<a id="canonical-3123131031031023-0100101000130311-1220102303303130-2231032312023300-1320333010021112-0101312021133030-3103331101013210-1132313311213201"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0011332222022203-1321313000323223-2001001020220102-2103303333333333-1213213201131200-3122223123123202-1233112131013321-1002302010312332"></a>

<a id="canonical-0012022330130010-2032202233032322-1320333201111301-3000001312020301-0330002030012010-1221122012032221-2232201033101102-0122030221012213"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1222113320121333-2000121121321100-2112033320330111-1023323321010321-0122103013112310-0102102320310031-1303320233230330-2311203122001130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_path_normalize` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- enable_path_normalize

<a id="canonical-3220103012323132-3123012003332321-1010200201202130-0001120230103122-2231023300220233-0201200010231213-0212213100033130-1203010112203200"></a>

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

<a id="canonical-0212202230011113-2203132220233101-1321131330330010-2030212021032322-3133023002222200-2133232330210133-3312230332103333-3123111320010133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- http_protocol_options

<a id="canonical-0121200200323330-3000031332102221-0220123010203231-2022233221110203-0012032001020321-0320130130030032-1312231312302031-0233031323201121"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020031013132202-3013203002322100-2122231211020222-0131223203300003-0331311211121222-2020131130300000-3213320213301203-3221313010010203"></a>

### Direct properties for `http_protocol_options`

- [http_protocol_enable_v1_only](resources--virtual_host--reference--group-002.md#canonical-3210022102331030-1102301122203233-1032122010103223-3301302200330220-3133302113323110-0303211131002113-0323210330020030-2332000211221030): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--virtual_host--reference--group-002.md#canonical-0031201323133231-2323233103310111-2003113221311133-3301000132022012-2210101221110112-2132232133020121-2220313323310121-1230130323032203): complete subsection reference.

- [http_protocol_enable_v2_only](resources--virtual_host--reference--group-002.md#canonical-3132300100302013-0130131332313101-3020312012102220-0221323111203101-1031323231301023-3213120230010110-0100003010002212-0223122111013021): complete subsection reference.

<a id="canonical-3210022102331030-1102301122203233-1032122010103223-3301302200330220-3133302113323110-0303211131002113-0323210330020030-2332000211221030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-0212202230011113-2203132220233101-1321131330330010-2030212021032322-3133023002222200-2133232330210133-3312230332103333-3123111320010133)
- http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-3330312023321112-3330313332212130-0120133011211001-3103312021201100-3000332101020231-3033021210323010-0131023230331013-1322332220302110"></a>

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

<a id="canonical-2210110232010213-3110033000101120-0332110210321303-2300002223112001-3100100201113100-0320222311310202-1323011122210002-2223302310111310"></a>

### Direct properties for `http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](resources--virtual_host--reference--group-002.md#canonical-3311022031030333-1010333211021130-1120130210200120-2002132001202131-2103010223302112-3211120210301023-2210220131121001-1131020001111110): complete subsection reference.

<a id="canonical-3311022031030333-1010333211021130-1120130210200120-2002132001202131-2103010223302112-3211120210301023-2210220131121001-1131020001111110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-0212202230011113-2203132220233101-1321131330330010-2030212021032322-3133023002222200-2133232330210133-3312230332103333-3123111320010133)
- [http_protocol_options.http_protocol_enable_v1_only](resources--virtual_host--reference--group-002.md#canonical-3210022102331030-1102301122203233-1032122010103223-3301302200330220-3133302113323110-0303211131002113-0323210330020030-2332000211221030)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-2223100312013213-0332203100322230-1123022310201220-1312022210333331-3121310103120000-1122130003323012-1112222301332130-3123123011311123"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103313202202200-2310212303112031-2102133010212000-3113031112011112-1010332120331223-0201020000011032-0331331321121002-2233232102213211"></a>

### Direct properties for `http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](resources--virtual_host--reference--group-002.md#canonical-3021332131210123-3203200332330321-2021022202222020-3212310220323210-3211011132121111-2130313200201020-0103012312310322-3332312331012220): complete subsection reference.

- [preserve_case_header_transformation](resources--virtual_host--reference--group-002.md#canonical-1103321023013312-0322220013332302-2311321011032331-1223322102031303-3123213113230031-0311331333022222-1311032212300232-0003200002333323): complete subsection reference.

- [proper_case_header_transformation](resources--virtual_host--reference--group-002.md#canonical-1300330310211123-1002221103021231-1332220203301323-0322211132303210-0030232003133212-3220201020132300-0232120213331220-1023032013100330): complete subsection reference.

<a id="canonical-3021332131210123-3203200332330321-2021022202222020-3212310220323210-3211011132121111-2130313200201020-0103012312310322-3332312331012220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-0212202230011113-2203132220233101-1321131330330010-2030212021032322-3133023002222200-2133232330210133-3312230332103333-3123111320010133)
- [http_protocol_options.http_protocol_enable_v1_only](resources--virtual_host--reference--group-002.md#canonical-3210022102331030-1102301122203233-1032122010103223-3301302200330220-3133302113323110-0303211131002113-0323210330020030-2332000211221030)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--reference--group-002.md#canonical-3311022031030333-1010333211021130-1120130210200120-2002132001202131-2103010223302112-3211120210301023-2210220131121001-1131020001111110)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1211112003010222-1132210302020312-3200213211203023-3223132223223323-3131312011333031-1323012200202331-3020233322032303-0321212032132302"></a>

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

<a id="canonical-1103321023013312-0322220013332302-2311321011032331-1223322102031303-3123213113230031-0311331333022222-1311032212300232-0003200002333323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-0212202230011113-2203132220233101-1321131330330010-2030212021032322-3133023002222200-2133232330210133-3312230332103333-3123111320010133)
- [http_protocol_options.http_protocol_enable_v1_only](resources--virtual_host--reference--group-002.md#canonical-3210022102331030-1102301122203233-1032122010103223-3301302200330220-3133302113323110-0303211131002113-0323210330020030-2332000211221030)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--reference--group-002.md#canonical-3311022031030333-1010333211021130-1120130210200120-2002132001202131-2103010223302112-3211120210301023-2210220131121001-1131020001111110)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-0112122310122312-0130121212231300-0303120201203002-0203203023010322-2302320323320123-0312003322102021-1030102032322312-3000331302121301"></a>

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

<a id="canonical-1300330310211123-1002221103021231-1332220203301323-0322211132303210-0030232003133212-3220201020132300-0232120213331220-1023032013100330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-0212202230011113-2203132220233101-1321131330330010-2030212021032322-3133023002222200-2133232330210133-3312230332103333-3123111320010133)
- [http_protocol_options.http_protocol_enable_v1_only](resources--virtual_host--reference--group-002.md#canonical-3210022102331030-1102301122203233-1032122010103223-3301302200330220-3133302113323110-0303211131002113-0323210330020030-2332000211221030)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--reference--group-002.md#canonical-3311022031030333-1010333211021130-1120130210200120-2002132001202131-2103010223302112-3211120210301023-2210220131121001-1131020001111110)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-3333100110223012-0101032322023230-0302033322210012-3200111120013222-1030121321300203-2223323133303102-1203012331311232-3122113001220111"></a>

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

<a id="canonical-0031201323133231-2323233103310111-2003113221311133-3301000132022012-2210101221110112-2132232133020121-2220313323310121-1230130323032203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-0212202230011113-2203132220233101-1321131330330010-2030212021032322-3133023002222200-2133232330210133-3312230332103333-3123111320010133)
- http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-0003033122202032-3012112010322112-3130100200211113-0222303121320312-1132233323223312-2202230312233131-0033223332200231-2223233132201301"></a>

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

<a id="canonical-3132300100302013-0130131332313101-3020312012102220-0221323111203101-1031323231301023-3213120230010110-0100003010002212-0223122111013021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-0212202230011113-2203132220233101-1321131330330010-2030212021032322-3133023002222200-2133232330210133-3312230332103333-3123111320010133)
- http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-2300022212233023-0322210221310110-0301000331312331-3001203310203123-2133232023321231-2001123210011021-1101212313022103-1123233110110001"></a>

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

<a id="canonical-0101231302233001-2011223111023233-0001132321323232-0221122312210111-3200231102333020-3311221202021012-2021103121020012-0121132032203123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `js_challenge` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- js_challenge

<a id="canonical-2301110332123321-2201100001322312-3310232321120302-2100202013231333-1032213203023203-3311131310311320-2313310010120313-0030132221110033"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011232221223320-0233202013221320-3321303212122121-2313002111132112-2010320211212200-1312010303123313-0220301102003132-3202213302022002"></a>

### Direct properties for `js_challenge`

<a id="canonical-2200112100333330-2303233000113213-3212132031201320-0212303202022102-2120120101323201-2003211321300313-3332323023023231-1100212112031211"></a>

#### `js_challenge.cookie_expiry` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2030230111000310-0320101310031231-0321202033312222-0333120230011323-2102311023321321-1012110011313222-3113231302011121-3120110323033332"></a>

<a id="canonical-2001312320231331-3010101013002121-0132022211133210-2022212213331023-0312031130302012-2101200332100331-1322112011103003-1022221020032111"></a>

#### `js_challenge.custom_page` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0123013001000222-2133130113221231-3300212132013200-3020131221210312-2323031211333202-3113331323203212-3002023122101303-2300231001110222"></a>

<a id="canonical-2021021032320012-1030200130320111-0113200202231313-2231201032132202-0003231221133111-2113110232232212-0000300310133123-0210111222032122"></a>

#### `js_challenge.js_script_delay` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3013320031330300-3111311312212033-2213131023230113-1110310112213020-3030011311320203-2011102330101110-1110113223010031-3121032131222123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_authentication` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- no_authentication

<a id="canonical-0103231132201100-0003112003220201-1232112230200001-0013023021010100-2023323122300101-0222212120020221-3323111231232231-2003131212032321"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_authentication = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032230233200032-3113230021023123-1002001030322103-2001220130000333-3010022102100323-0113113322202322-2132212011203032-2132231103032230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_challenge` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- no_challenge

<a id="canonical-1023032113232030-0020323313122121-0322301020011111-0223320231112323-1001102230211202-1023032220132002-0312223223230123-1312013321231030"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_challenge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201021232211320-1113031321300211-0230223223322233-2013130232110130-1300111221111201-1102111321313123-0013222103000031-1210301000110230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_request_limit_per_connection` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- no_request_limit_per_connection

<a id="canonical-1220233000321320-2201210110230233-3022023021301313-3131222033332030-3021323101221112-2320200132210013-0112101301202311-0300202201122113"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_request_limit_per_connection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103112131222102-3010030322132022-3211222333113012-3301232022132313-0320122322331301-1020131023201302-2322122130131023-1210231202212021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- non_default_loadbalancer

<a id="canonical-0203302010201112-0331211102220230-0300132232103103-1210332110032101-1332230330310132-1011232020203331-2233222103311321-3203201030110132"></a>

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

<a id="canonical-0323002221202331-0013320123000331-2220112010231031-3313320111013331-3201130132130201-2323102121111212-0000301210022202-0212030313323103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `pass_through` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- pass_through

<a id="canonical-0132112102332232-0123332032222330-0301020323101110-3123022022131303-1232332032232110-0303301323030131-0113120332013210-1113310021333202"></a>

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

<a id="canonical-1012012232023120-0032122201003333-2110302001323221-0232313320231102-2031333000031122-0111010022332321-0120313223113103-2321021201002203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limiter_allowed_prefixes` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- rate_limiter_allowed_prefixes

<a id="canonical-3301131211121202-3332111111230013-2113212301003120-2021213232103210-0002020320300330-0100112121211312-3010223232330201-2112231233011130"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
rate_limiter_allowed_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0131220021013120-3330310101033102-3312100133230202-1133220112000032-3022010111031110-0333011031330001-0203303133030232-0032233213223211"></a>

### Direct properties for `rate_limiter_allowed_prefixes`

<a id="canonical-1112011231300102-3300100010132003-1031313112202220-2020211133203122-2003031213203312-3303012123313231-3202003000300233-0332112203130000"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0132122021311320-3023233222230102-1001101200123310-3002333121301330-2231132101133302-0021131122121120-0213022303302030-1022231032333103"></a>

<a id="canonical-2003022222232311-0123312230112332-1100013131112011-0002113001011010-3233201222102220-3021103013212231-3121320321211301-2201001033112111"></a>

#### `rate_limiter_allowed_prefixes.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2133312030102020-0200220310130122-2232012120132030-0303233321113312-0032022203331210-2212203323100320-3020313131330011-1132200300002223"></a>

<a id="canonical-1102300021333033-0131211221002030-2223300311321300-3122333022101203-3132201021221010-3111220032020001-0121211221310320-1313221022212203"></a>

#### `rate_limiter_allowed_prefixes.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3110002303210333-2033011203312313-0310330021233230-3120311200230003-3102103203211023-0013200012012103-3201030312132213-3332213321003223"></a>

<a id="canonical-1322132300301312-3033321020332330-0212022130202000-2322131120021033-2011303210010032-1020332222002001-1220110333333002-2223330322100002"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1213110300200123-3100001102333012-1202001002031132-3231313223311102-1231032121032110-0110301321212033-1323110220122300-0110121002301231"></a>

<a id="canonical-1133203111010321-0121302232122020-1223331313210333-2301113312122021-0112132203311230-0211112330033012-0203221102320103-3311111300322022"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0111131022031311-0102023300100223-2120311331022200-0220230223131102-3320022000312001-1021212300331101-1201203011120311-1323012103110312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_cookies_to_add` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- request_cookies_to_add

<a id="canonical-1132200202122012-0120231300202311-1120202301232030-0133023100121320-0313102032300112-1303032313321032-3233123332123220-2123322212001013"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
request_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221302110100123-3101301210311221-1100202031332323-0010323112111320-1113111100321100-3312100231133330-0000203323020302-0130101312333011"></a>

### Direct properties for `request_cookies_to_add`

<a id="canonical-3110333310310322-2332330123133021-2321112123103312-1321100111332112-2102130131001310-0312330021023323-1013220111001332-1122210101331122"></a>

#### `request_cookies_to_add.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0332100333321211-3331133232201230-3112220021022313-3331213102300221-2220310001201232-1231223310231213-3211103223022210-3013002000333003"></a>

<a id="canonical-1320103010030310-1212333013221330-3303202202112101-1111200003021223-2331312332113002-0122011131013031-1113222320012123-2302132112210022"></a>

#### `request_cookies_to_add.overwrite` property

Type: `"bool"`. Optional.

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

- [secret_value](resources--virtual_host--reference--group-002.md#canonical-1002131030330032-0110121232021222-0103211301003113-2320020231120122-1000102033211111-3130322311000021-0203302200212301-0100121330320103): complete subsection reference.

<a id="canonical-1312300232101312-3321320301021013-2303102103100002-0010110101010002-3210223332310320-1010222010031032-1330200230130322-2003000230110203"></a>

<a id="canonical-3220133322012031-3333110012122110-1121213111320112-1113213223210032-2130001201101023-2131321113000030-1311102002001102-2123122021211211"></a>

#### `request_cookies_to_add.value` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1002131030330032-0110121232021222-0103211301003113-2320020231120122-1000102033211111-3130322311000021-0203302200212301-0100121330320103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [request_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-0111131022031311-0102023300100223-2120311331022200-0220230223131102-3320022000312001-1021212300331101-1201203011120311-1323012103110312)
- request_cookies_to_add.secret_value

<a id="canonical-3000303221003030-1131111221011200-0032123200103013-2232133132011313-1333231331130232-1030300222313003-3023233002321200-0213112001202101"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330003201231233-0233102031213031-3123010322300123-2022030303321101-2201000223023010-3112132032202002-3312233020200203-2211321130133323"></a>

### Direct properties for `request_cookies_to_add.secret_value`

- [blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-3132013211313321-0102210013031220-3203331002031103-3020303011333103-0102122013121110-3012003020311023-2221323130212223-2031102203321100): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-2000122111133200-3110310130311333-3111200023022323-1000220321213023-1013302111333013-1101011321101100-2330200200323022-3202112032301202): complete subsection reference.

<a id="canonical-3132013211313321-0102210013031220-3203331002031103-3020303011333103-0102122013121110-3012003020311023-2221323130212223-2031102203321100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [request_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-0111131022031311-0102023300100223-2120311331022200-0220230223131102-3320022000312001-1021212300331101-1201203011120311-1323012103110312)
- [request_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-1002131030330032-0110121232021222-0103211301003113-2320020231120122-1000102033211111-3130322311000021-0203302200212301-0100121330320103)
- request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-1132212111133203-0110230032011110-3230012121010213-1131230013030212-3213011322011102-2001131102120123-2302302323212231-1110002113213111"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223013113331111-3223232230012033-1133030303313332-0133310220130102-3222313001020303-0232130130132333-0332300311110133-1230300100032132"></a>

### Direct properties for `request_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1323002121230133-3211001201313111-1200023313022123-0221110121332201-1033230220331201-3022333002032100-3101100031012130-1223210312012030"></a>

#### `request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3222233101030000-2321122212102221-2222320323313020-0231102101102102-1330331310210210-1300111222021132-2200023033212333-0120032002313231"></a>

<a id="canonical-2110230222133023-1203203130102223-3112011322003300-2303312300030120-1102332123100223-1221031222302020-3321313121130012-1123313133312010"></a>

#### `request_cookies_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1121322202213333-1003103022211131-1212033133310101-1031220003213123-3113023000332220-0213230132202311-2000131003133121-2020310112000112"></a>

<a id="canonical-1302202120301113-0111310103002123-0211201131032113-2300221201113300-2030110303100320-2022302013130321-2023312003212312-0331112230020210"></a>

#### `request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2000122111133200-3110310130311333-3111200023022323-1000220321213023-1013302111333013-1101011321101100-2330200200323022-3202112032301202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [request_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-0111131022031311-0102023300100223-2120311331022200-0220230223131102-3320022000312001-1021212300331101-1201203011120311-1323012103110312)
- [request_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-1002131030330032-0110121232021222-0103211301003113-2320020231120122-1000102033211111-3130322311000021-0203302200212301-0100121330320103)
- request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-3111133101200302-3112022120333133-0220331201122321-0021122331010131-0130001312201013-2111031300222023-3201000233321301-0010021301321323"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113233201123223-0103003231312311-3333022031333010-1021333013322010-1111131300000230-2033323332333003-2233300232020000-1002003321012330"></a>

### Direct properties for `request_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-0333121232013311-2020101201033302-1201032132322101-1100022322333120-1013110233310021-1030232131233112-3111201000110130-1300210102121211"></a>

#### `request_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2203223313011111-0301002302020021-3320331302111210-0210020020130003-2021100300132200-0333321211301200-0210331002221233-2303301120212313"></a>

<a id="canonical-1021301103333001-2212201020220131-0002102121223231-1210312002320212-2033211331133033-0130030001133102-3132103110233332-2022003001112211"></a>

#### `request_cookies_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3023100112011132-2302030100320013-2030100013010220-0112333330230033-0110121200123102-2012130203132311-0221233210230001-2212303220211221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_headers_to_add` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- request_headers_to_add

<a id="canonical-3012313223321220-3121120103002220-3002111110010303-2131002021322310-1130100033030200-0333222300022220-2303003031131332-1233011000012013"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231020212303030-1020212103231031-3300322102002223-0002023020302031-2131103131211131-0230212312030033-2130322023033221-2133033313312231"></a>

### Direct properties for `request_headers_to_add`

<a id="canonical-2010323003001201-0231101202111211-2120113301310303-3312133100101001-2201303020221012-2332033001021213-3302203123233222-1012220033201003"></a>

#### `request_headers_to_add.append` property

Type: `"bool"`. Optional.

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

<a id="canonical-1203233002121312-3211321130122001-3200301223231123-0110230220213120-3013022120013201-2222213300031321-3120301131121223-1322013112130100"></a>

<a id="canonical-0331130003100321-3320132123033302-0100011112222230-1233223311123232-0003121213011300-0331123101223303-1121020010120310-3233231032102103"></a>

#### `request_headers_to_add.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [secret_value](resources--virtual_host--reference--group-002.md#canonical-2301022310030122-2011302321130001-1212132333110003-0212200001032330-1112033303013302-3100023023022010-2102113011233123-1021302003310321): complete subsection reference.

<a id="canonical-0013233321213220-2212201230202132-3232211100113332-0322133131011032-2233300322320021-1321333312010100-2333100103200222-1021322122132023"></a>

<a id="canonical-2201320101121331-1331030011122203-1302120110031102-3011122223212023-0311103102011231-0321320332021133-1221310032332012-0022121102300231"></a>

#### `request_headers_to_add.value` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2301022310030122-2011302321130001-1212132333110003-0212200001032330-1112033303013302-3100023023022010-2102113011233123-1021302003310321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [request_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-3023100112011132-2302030100320013-2030100013010220-0112333330230033-0110121200123102-2012130203132311-0221233210230001-2212303220211221)
- request_headers_to_add.secret_value

<a id="canonical-3232230202330303-2130001230123102-1132113011220232-2023301202100301-2323121322221323-0330032213301110-0300212303011010-2201232102001011"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123323000320111-2013221211030233-3112203231112232-1302001332010223-0131321101230311-3000310202002331-3121232320222232-2313323320121313"></a>

### Direct properties for `request_headers_to_add.secret_value`

- [blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-2331323120322022-2020202122332003-0323130120310030-2122122201113312-2321131330233113-0332130223123013-1302101103201332-0333320023100203): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-1001131332231313-0303023303001121-0321131021120312-2320130103323311-0102333202033233-1301122220021021-3312231020311013-0120232011300331): complete subsection reference.

<a id="canonical-2331323120322022-2020202122332003-0323130120310030-2122122201113312-2321131330233113-0332130223123013-1302101103201332-0333320023100203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [request_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-3023100112011132-2302030100320013-2030100013010220-0112333330230033-0110121200123102-2012130203132311-0221233210230001-2212303220211221)
- [request_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-2301022310030122-2011302321130001-1212132333110003-0212200001032330-1112033303013302-3100023023022010-2102113011233123-1021302003310321)
- request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-2313103131023113-1011002230021132-2213101022023113-0021213311030232-3113002303232212-0220101233120032-1032311110333022-3030202211221233"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1301232313033323-1032030022213031-3312021320232302-3012330222112201-2232210133123103-3223332213113030-1000232222302231-2231101121113312"></a>

### Direct properties for `request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1233302001023212-2322221231211312-2010222210202121-3220033013323333-1321313201012031-2112023201132333-0302101121303103-0233102201333131"></a>

#### `request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3310302302211221-0020102100103113-3303122133103013-3230221202313123-1120112010232310-3321333321233303-0233133320122202-2333312310331112"></a>

<a id="canonical-0202102010331132-1322031122110210-0020020002032020-3032120100010111-1330103230333011-2120223033221322-2033303130302112-1300002000002102"></a>

#### `request_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3102211033132301-0100133322312313-2023311210012200-0010000120013332-1211032020311010-0212303322100201-2003030003113300-2122302300012130"></a>

<a id="canonical-2323200201223322-3211221011021200-3131103120222221-0023302121001132-0103121320203000-3130332110022023-3312103210012100-2120202231132222"></a>

#### `request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1001131332231313-0303023303001121-0321131021120312-2320130103323311-0102333202033233-1301122220021021-3312231020311013-0120232011300331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [request_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-3023100112011132-2302030100320013-2030100013010220-0112333330230033-0110121200123102-2012130203132311-0221233210230001-2212303220211221)
- [request_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-2301022310030122-2011302321130001-1212132333110003-0212200001032330-1112033303013302-3100023023022010-2102113011233123-1021302003310321)
- request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-1221212101301212-0112121012232333-2033300101312312-0022201332120231-0112223312102120-0233302322133312-1223312131120002-3332313032111230"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102332032302213-0303310003113302-0103113020012300-0103010200012021-3223333322021122-0233300113303213-2023202221022102-1232321033233000"></a>

### Direct properties for `request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-2112002133311301-0323013311233013-3133110323300322-0212131031201223-1233022211001013-2200122210030320-3122001012311202-1023321311312030"></a>

#### `request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2323211222111331-2300102300130302-2011102332122120-0211211221221322-0132001101030031-3220313130030123-0201132001203312-0130012210133203"></a>

<a id="canonical-1312332130320111-1332020110122103-2020331132203122-0302100222233012-3302011102113302-2110130202133032-1332203320212030-0120323231101320"></a>

#### `request_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- response_cookies_to_add

<a id="canonical-2222003011301030-3202222031130133-1300311132321213-0002212031300032-1232213100222100-1111332011333020-0302312231231200-1310300303220332"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
response_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-1213102331112001-2213030232213201-0120213232102032-3012023131021030-3022032020331110-0302022211103231-2303313030112132-0130021302103032"></a>

### Direct properties for `response_cookies_to_add`

<a id="canonical-2021011121031211-3210023310302303-2020203012111211-0221120301100103-3212101023203021-2130102121211231-2322331312031022-3030102330312111"></a>

#### `response_cookies_to_add.add_domain` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1332223231213232-1112200333232311-3210301020332000-3220212222132230-2010311223203223-0120130320131032-2233120110111203-1210320033102300"></a>

<a id="canonical-1322003211111211-3131230221322230-0323332221301212-1110321003221110-1112202032230011-0321323012132332-3212010323302121-1013030313030231"></a>

#### `response_cookies_to_add.add_expiry` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [add_httponly](resources--virtual_host--reference--group-002.md#canonical-0121330302100113-3120013113022120-2011102111212232-1131231201303301-2301002012032320-3300120320100213-3232001203321130-3123230202331331): complete subsection reference.

- [add_partitioned](resources--virtual_host--reference--group-002.md#canonical-2023002300033302-3331013320000001-2103101332023300-3213112223101300-0233003132230103-1011131101303231-0331122311331202-0033221103003321): complete subsection reference.

<a id="canonical-1100032020320211-1220111000033233-2222030022020113-1120131203023020-1021320121010212-0230031013010310-1213312003212021-1212013200113203"></a>

<a id="canonical-1133101013133012-2302012002110200-3233330033011032-3032102123300220-3003103012302322-1112300002103123-1313212201132010-2133231322220121"></a>

#### `response_cookies_to_add.add_path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [add_secure](resources--virtual_host--reference--group-002.md#canonical-2220302231103210-3020332130223302-3323012232313223-3230210330103011-3210030311210322-2010001121322323-0213021023113212-2212230323102121): complete subsection reference.

- [ignore_domain](resources--virtual_host--reference--group-002.md#canonical-3201100313033331-3021302030332320-1031022021211101-0112002311330312-1102000110001113-2323002121322130-3301322102003301-1102202322300212): complete subsection reference.

- [ignore_expiry](resources--virtual_host--reference--group-002.md#canonical-0210213113321323-0022010203302111-1210020132111333-0323322023030212-2132102103013200-0322031100230111-2102033202133320-1132101020302200): complete subsection reference.

- [ignore_httponly](resources--virtual_host--reference--group-002.md#canonical-3021021000020303-3132010100310213-3232233323323001-1232110121210132-0222003122132021-3131332232133003-1321321302032200-1221300030321202): complete subsection reference.

- [ignore_max_age](resources--virtual_host--reference--group-002.md#canonical-3030101211213013-2012310330131231-1112110210223013-3321312001103003-3030130000301021-1331131112233323-3011032332020100-0130002032120031): complete subsection reference.

- [ignore_partitioned](resources--virtual_host--reference--group-002.md#canonical-0323032210213111-1233102002320133-2223031120200013-0002312313333133-1213210120100200-0131121211231200-2103303010321301-1032311303023200): complete subsection reference.

- [ignore_path](resources--virtual_host--reference--group-002.md#canonical-0131011203101031-1232030321310130-3200331201303231-2321123333213320-0312233330221313-1232232322301323-0300202220320103-0313231313313320): complete subsection reference.

- [ignore_samesite](resources--virtual_host--reference--group-002.md#canonical-3310012300203022-2122332302223001-1332032010010102-2303203320323131-0320100032301111-3320103233133222-1311233112333032-2130330323101110): complete subsection reference.

- [ignore_secure](resources--virtual_host--reference--group-002.md#canonical-3311322330133013-0002222123021220-0300211113100121-0113102331201102-0312013333010222-1011321131132331-0012103210231222-1110222120213302): complete subsection reference.

- [ignore_value](resources--virtual_host--reference--group-002.md#canonical-3020132301301100-2332100203221311-0211103012132321-2121022123021002-1122311223201331-2230332333201130-1201230002323301-0033311230201031): complete subsection reference.

<a id="canonical-3012301010312230-1222133313103133-0112000321231213-3013031310000203-3113002232133110-1301020220111203-0232113120001010-2332011101102002"></a>

<a id="canonical-1023303000323332-2030230033020203-3211212113222221-3130111000121003-3132222132213332-3333100102002033-2222032330012132-0213203133020022"></a>

#### `response_cookies_to_add.max_age_value` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3221033320123302-3333110200003230-2130020301100001-1323103313233310-2200203320223331-3113333032202330-2112201333001131-2001113300303032"></a>

<a id="canonical-3013230110332112-1113210133010112-1202212201302121-0130212033213030-0123002333330112-2130232233202213-2332210322200311-1230323010302303"></a>

#### `response_cookies_to_add.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3013021010000112-2133312130123322-1002313233033200-0013222210323001-3221111203332023-3221100221313022-3303231030010021-0311011101233013"></a>

<a id="canonical-1110031000201111-1211220121133330-3221221313110122-1223133331111133-0210022002310120-1300113211110301-0020313302133111-2003302122010202"></a>

#### `response_cookies_to_add.overwrite` property

Type: `"bool"`. Optional.

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

- [samesite_lax](resources--virtual_host--reference--group-002.md#canonical-3003211130301030-2123220322312113-2320032202102130-0131323200322130-1021203001021112-1130122231121233-3210112001112110-0233221303311022): complete subsection reference.

- [samesite_none](resources--virtual_host--reference--group-002.md#canonical-0132300020202331-0331200133003121-0113000221102102-1333231020013220-2321020013221321-1220120222022302-3211230202002331-3311012133132320): complete subsection reference.

- [samesite_strict](resources--virtual_host--reference--group-002.md#canonical-3333212232123132-1213103200121230-0310032101323202-2101213202131030-3020031200002021-2211301100003033-2312103102113310-0032303122313020): complete subsection reference.

- [secret_value](resources--virtual_host--reference--group-002.md#canonical-1010303131122313-3322223020312220-2031002021323312-2321220101102002-2222302113021322-2120000333313022-0320332221303010-0110313302022130): complete subsection reference.

<a id="canonical-1330120322003022-3120321110121302-2111132010331030-1130313121303303-3010330130001001-2233020120121030-3113023322023311-2312322200303101"></a>

<a id="canonical-0000313310102000-1331102022012220-3211100202132200-3232132302011210-1123210213030133-2312200133320332-1002030302310233-3110020001022320"></a>

#### `response_cookies_to_add.value` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0121330302100113-3120013113022120-2011102111212232-1131231201303301-2301002012032320-3300120320100213-3232001203321130-3123230202331331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.add_httponly` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.add_httponly

<a id="canonical-3120210223030310-3030233311200121-3021100300230021-3132333020111033-0003022231312221-1012310033020022-1133301002130011-0120202210231212"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
add_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023002300033302-3331013320000001-2103101332023300-3213112223101300-0233003132230103-1011131101303231-0331122311331202-0033221103003321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.add_partitioned` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.add_partitioned

<a id="canonical-3002222321133222-1211333333013321-2111122313120123-3332021132331033-3002132122211310-2311123013230003-2113321000322202-0302201010332331"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
add_partitioned = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220302231103210-3020332130223302-3323012232313223-3230210330103011-3210030311210322-2010001121322323-0213021023113212-2212230323102121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.add_secure` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.add_secure

<a id="canonical-0221210230121010-1030110221213031-1212220021223203-3021301321222013-2303320130322233-2311033332322112-0313020111313022-1202131101300122"></a>

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
add_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201100313033331-3021302030332320-1031022021211101-0112002311330312-1102000110001113-2323002121322130-3301322102003301-1102202322300212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_domain` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.ignore_domain

<a id="canonical-3231320210312130-0220031111132303-0330101001001033-2001223300313031-1302330022010033-2022210301031031-2212103221031231-1220030203203223"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210213113321323-0022010203302111-1210020132111333-0323322023030212-2132102103013200-0322031100230111-2102033202133320-1132101020302200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_expiry` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.ignore_expiry

<a id="canonical-1222320221003333-2312002000233333-3222002211220222-2313012020122000-0111312231201212-3321220110201123-3231110220122321-1232123003133010"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_expiry = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021021000020303-3132010100310213-3232233323323001-1232110121210132-0222003122132021-3131332232133003-1321321302032200-1221300030321202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_httponly` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.ignore_httponly

<a id="canonical-2033020133203330-2012130130232022-2021120112011101-1332122031131330-1332112230133333-3111001210302103-1232303021030320-3110121002021100"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030101211213013-2012310330131231-1112110210223013-3321312001103003-3030130000301021-1331131112233323-3011032332020100-0130002032120031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_max_age` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.ignore_max_age

<a id="canonical-2332113113003020-1322101231212120-0132020311110113-2123000110131322-3022313133210310-1203232330112103-0210110133103031-2212133323322313"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_max_age = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323032210213111-1233102002320133-2223031120200013-0002312313333133-1213210120100200-0131121211231200-2103303010321301-1032311303023200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_partitioned` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.ignore_partitioned

<a id="canonical-0011210320120210-2232220310231120-0000311010223303-0330111213202302-2321101301222001-0320231101333332-1120100000130120-1120310301102203"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_partitioned = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131011203101031-1232030321310130-3200331201303231-2321123333213320-0312233330221313-1232232322301323-0300202220320103-0313231313313320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_path` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.ignore_path

<a id="canonical-3210022222033021-2202133320331330-1003301111223131-2010103012122113-2200222222002232-3133210131103312-0230321211002201-2000231230213131"></a>

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
ignore_path = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310012300203022-2122332302223001-1332032010010102-2303203320323131-0320100032301111-3320103233133222-1311233112333032-2130330323101110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_samesite` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.ignore_samesite

<a id="canonical-2011020332310022-0201003300002320-1211123211021232-3313203000233302-2033210131020030-2131230221222332-2103102131103313-2321012230230303"></a>

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
ignore_samesite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311322330133013-0002222123021220-0300211113100121-0113102331201102-0312013333010222-1011321131132331-0012103210231222-1110222120213302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_secure` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.ignore_secure

<a id="canonical-2101322000003222-1012302202233233-1001013031032132-2000212121201330-2011220311133030-3202003220100023-1122132213000220-2023130133102313"></a>

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
ignore_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020132301301100-2332100203221311-0211103012132321-2121022123021002-1122311223201331-2230332333201130-1201230002323301-0033311230201031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.ignore_value` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.ignore_value

<a id="canonical-0213002033003332-0111222102033220-1311130033221331-3132230322210022-0130101102330101-2120002032320202-3321011113130031-3032111330112323"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_value = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003211130301030-2123220322312113-2320032202102130-0131323200322130-1021203001021112-1130122231121233-3210112001112110-0233221303311022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.samesite_lax` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.samesite_lax

<a id="canonical-0132233221213211-0232200202330301-1022100132323301-0331221033100233-1311231131203122-0000300022002000-2320201011123101-3023233201132130"></a>

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
samesite_lax = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132300020202331-0331200133003121-0113000221102102-1333231020013220-2321020013221321-1220120222022302-3211230202002331-3311012133132320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.samesite_none` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.samesite_none

<a id="canonical-3122300331321103-2100331310330333-1201132103301031-1311222322233233-0322303010101011-3200310021101203-0033222003020020-1222301120320110"></a>

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
samesite_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333212232123132-1213103200121230-0310032101323202-2101213202131030-3020031200002021-2211301100003033-2312103102113310-0032303122313020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.samesite_strict` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.samesite_strict

<a id="canonical-0300322121023121-3203023200020310-2302121113011232-3232231320202203-3231003330201110-1132233331313202-0013100301202300-0100022110313211"></a>

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
samesite_strict = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010303131122313-3322223020312220-2031002021323312-2321220101102002-2222302113021322-2120000333313022-0320332221303010-0110313302022130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- response_cookies_to_add.secret_value

<a id="canonical-0210333030013002-2101002202331002-2011133222322133-0210132303330022-0212321333231213-1030321310103123-1000302123000020-1102231123032101"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310000133130200-0320002013103020-0020223100120131-1002303213003103-3230332202213233-1003122100320023-3120312103231312-3310203301022020"></a>

### Direct properties for `response_cookies_to_add.secret_value`

- [blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-3001001033223233-0203321100100022-3131222020121222-0012320320333121-1131302223311010-1313310112331202-3020030110033320-2320121330320123): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-1222030333223110-3200212101022212-2332333023201200-0032120320223203-3201132013000233-0011121110113022-1310213121022010-2302312001020332): complete subsection reference.

<a id="canonical-3001001033223233-0203321100100022-3131222020121222-0012320320333121-1131302223311010-1313310112331202-3020030110033320-2320121330320123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- [response_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-1010303131122313-3322223020312220-2031002021323312-2321220101102002-2222302113021322-2120000333313022-0320332221303010-0110313302022130)
- response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-3201302200000310-0311222231330013-1231313303010133-2301001233012010-1303321020231113-3022120110230123-0211031323212100-3102220201320122"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230122113331203-1111322220323200-0121111020111000-0033221021131001-0312201020132222-3310211212021310-1333312102031312-2320013233131213"></a>

### Direct properties for `response_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3310132220002331-3030312200210132-1132100302302032-0313012103112312-0320113033300302-0020030022123311-1331333021021310-2011210310030210"></a>

#### `response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1002013121110101-0211100012202320-2112301112231220-0222333310122230-3200010313011113-3310213213213013-0220011122131010-1220300210322031"></a>

<a id="canonical-2302320230102311-0030020220112001-3131210310133331-1123002120112312-3300203203131311-2020111311113122-1220102211303033-1020031322320013"></a>

#### `response_cookies_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0032332111130133-3321030330002122-2032023012133030-0221122002222210-3201101021300313-1031020232323333-0320221302332311-1112003320121300"></a>

<a id="canonical-1212300203323132-1202222022222103-1222113121233310-0100131300323231-2021022232212301-1030031223123123-2302013313012132-1022011002133010"></a>

#### `response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1222030333223110-3200212101022212-2332333023201200-0032120320223203-3201132013000233-0011121110113022-1310213121022010-2302312001020332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- [response_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-1010303131122313-3322223020312220-2031002021323312-2321220101102002-2222302113021322-2120000333313022-0320332221303010-0110313302022130)
- response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2201120200032020-3231311213302112-1301111132002220-0202333213133220-0021121202211221-3030301001202113-1022222200202312-3003001322031221"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333230131223232-0200103011030200-3131113332320211-0030232211301021-2322030132223221-2201221201221312-3212123000230213-3003300323033303"></a>

### Direct properties for `response_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-0320023110210030-0132120000023010-2003233200032232-1033033333102032-2302232311011110-2201231232310331-0220033101211311-1023031221330303"></a>

#### `response_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0122102121020011-2032120130332333-2200123132330033-0220030130232103-0130313312010111-3220231110302223-3331313222023220-1130131333330003"></a>

<a id="canonical-1012033032130033-2000020030000202-2030210212223012-3021121303213131-0021023102133233-3310003033321011-1122320132112022-2032130202313323"></a>

#### `response_cookies_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2031023100033202-2330100210303322-3113323303313133-2030111112021212-1302210011133231-0221213311321333-0210033213203311-0133030211030022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_headers_to_add` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- response_headers_to_add

<a id="canonical-1332201100013111-2211122001031301-3021122213021002-3103112322011321-1203121332113223-1233013010332233-2313001300333001-1313003003202120"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
response_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201220201322331-0103302213323313-1211221122220231-1223323022130130-2200202010331133-2221303032300110-1201230101032033-3021212030120223"></a>

### Direct properties for `response_headers_to_add`

<a id="canonical-1002223013013103-2102120223002030-1132001311122000-0310111231212011-0030132021203103-1200302310031220-3320213123333230-2020120032130001"></a>

#### `response_headers_to_add.append` property

Type: `"bool"`. Optional.

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

<a id="canonical-1110020330011320-0301111023102030-0003130000010223-2001003023103133-0000032033331303-1132211333120122-2310221231211001-3013002233202111"></a>

<a id="canonical-0313223110303201-0200323132031021-1013013132113023-3032201121210210-0032201222003032-2210312213031231-3033322330320112-0310023010200000"></a>

#### `response_headers_to_add.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [secret_value](resources--virtual_host--reference--group-002.md#canonical-2111212032020003-3033322032323303-2321321231321000-1310211200131101-1030333200010020-3101110203003230-1013011320132113-3312210222133312): complete subsection reference.

<a id="canonical-1020202333133332-0012231300003001-1303210322300100-2030022011210310-2101011223100332-0321110122110112-0100302200222313-2100011130210121"></a>

<a id="canonical-0221311332133221-3302031302010312-0121232211213313-2301330312332213-3201012303131330-1211220313103200-2332103031131131-3220310101033013"></a>

#### `response_headers_to_add.value` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2111212032020003-3033322032323303-2321321231321000-1310211200131101-1030333200010020-3101110203003230-1013011320132113-3312210222133312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-2031023100033202-2330100210303322-3113323303313133-2030111112021212-1302210011133231-0221213311321333-0210033213203311-0133030211030022)
- response_headers_to_add.secret_value

<a id="canonical-1022221223111302-1022022232321022-1120002030211230-1101301311000311-3111213102210232-2132233222132130-1231321212320312-3133213103111301"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-3002223211213130-1321032302330003-1200220310001300-1003002313110333-1233230303331020-1100211112230022-3122003311113221-1022221110312231"></a>

### Direct properties for `response_headers_to_add.secret_value`

- [blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-1233013201222331-3030133300002211-3201230110320221-1000211222112222-0232333220101033-1322020211330332-0210031213033123-1022332322120233): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-1210303002202023-1322031021321021-3331232302132303-3201130332331023-0303111303030100-0213221021012223-1222230233321230-2123233200312131): complete subsection reference.

<a id="canonical-1233013201222331-3030133300002211-3201230110320221-1000211222112222-0232333220101033-1322020211330332-0210031213033123-1022332322120233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-2031023100033202-2330100210303322-3113323303313133-2030111112021212-1302210011133231-0221213311321333-0210033213203311-0133030211030022)
- [response_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-2111212032020003-3033322032323303-2321321231321000-1310211200131101-1030333200010020-3101110203003230-1013011320132113-3312210222133312)
- response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0130000101120133-3330110310003133-1030200231211102-0332201203201230-3102013302233322-1212300113012200-2213203321130132-2232031310101020"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011131223000110-1122333311031200-2100123322332323-0112301312310312-0130122023122033-1123002102221102-1022012303321313-1101332123022212"></a>

### Direct properties for `response_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1111111033210122-1101122203100302-0221001130032300-3102022022333022-2330121333133333-2232210230332203-2113123220013313-3313222123232310"></a>

#### `response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1111130000031302-0133021123222020-2320332132300230-3313201201030202-2221201333312132-2323012303313133-3133323121221232-1232032332310330"></a>

<a id="canonical-0021211222122020-3233002021320231-1222200113331301-2101102032322102-2300120323111323-3323120311201111-3101130321112321-3302230101122001"></a>

#### `response_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0022001000331301-1333032031222220-3320311321313032-0030313223332031-3003300301330300-1103101031131120-1132111233220002-0112201201312321"></a>

<a id="canonical-1213031000221201-2111122211222003-2001122020023302-0200323130131320-1223302231100132-0331112203211303-0130013202310000-2203132113220032"></a>

#### `response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1210303002202023-1322031021321021-3331232302132303-3201130332331023-0303111303030100-0213221021012223-1222230233321230-2123233200312131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-2031023100033202-2330100210303322-3113323303313133-2030111112021212-1302210011133231-0221213311321333-0210033213203311-0133030211030022)
- [response_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-2111212032020003-3033322032323303-2321321231321000-1310211200131101-1030333200010020-3101110203003230-1013011320132113-3312210222133312)
- response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-3201300033031132-3102213312230110-3200121221130002-1023200120331300-3230013111301132-2030330103012213-0301003133123000-0132311302103310"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2100102000321121-1201231312112221-0111211333303210-1100131310132233-3302130002022101-2010112133030213-2212003120122113-3011122010332101"></a>

### Direct properties for `response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-0301030202012021-1031112013213332-0020010013030001-0132201123200003-0213013101010120-3011122021023131-0011023032220332-3011103001102211"></a>

#### `response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1320102133121300-3203103220333200-3010331312220321-1133333132303110-1111330002332321-2220023213320031-3003310101013301-0130323201220222"></a>

<a id="canonical-3330122233110102-3233233333223132-1301231120200100-1002223032332122-1312121302103103-1121112221102200-2110232030332130-0332300113330202"></a>

#### `response_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3120210313011012-3220120012303030-2320221131003020-2003111301323011-0131032003311130-3311130033201323-1332312222032333-0100213232033200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `retry_policy` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- retry_policy

<a id="canonical-2030200310111133-3030032003003120-3002230020001330-3301223321122321-1312122333313220-0210331300030232-3323011033030330-1311130010323222"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
retry_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-1103120320021100-3132322131301320-3303002101020223-2203313120301131-2133232032010222-2110101122130112-0122301022110322-1210210230333331"></a>

### Direct properties for `retry_policy`

- [back_off](resources--virtual_host--reference--group-002.md#canonical-3232203011321132-1200220222313300-0222313113000111-2210230121110011-1033231311222001-3213020130210323-3330003232013232-2301313032100120): complete subsection reference.

<a id="canonical-3321101220111111-0210303233310200-3322021320321323-3330223120110123-1012032303123103-2222330310020030-1321031311333220-2031130212212310"></a>

<a id="canonical-2210212300231223-0202130200102012-2022001032200020-2302031330312231-0101213010112220-0010230323331203-1012120022031110-0313113103122121"></a>

#### `retry_policy.num_retries` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0330120133021002-2200030001111231-1010010302211330-1110133103332012-2212201003221211-0120331032312010-1021120021131332-3013112231103000"></a>

<a id="canonical-3231103000111123-1120120120302200-3031112331210010-2311311332220023-0300103301203221-0331212022032101-0002111213113230-2333313301310202"></a>

#### `retry_policy.per_try_timeout` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2103233032202121-2131010003130032-2113013210000222-1111110310120113-0210013031311000-3323331321011201-1030332120032200-3003202201132113"></a>

<a id="canonical-3002033122131031-1313102202113110-2203102223110111-0321001323000312-0203123003010213-0301002123010332-3030012311110011-2133322022101013"></a>

#### `retry_policy.retriable_status_codes` property

Type: `["list", "number"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2313310213313002-3233220321023030-0133101030030101-0310300223013100-1123210020031310-0133131313012233-3330200322311330-2011031110102212"></a>

<a id="canonical-2032030330333200-3313031103131330-2223320111033012-2211232100200202-2320130101323011-0003231132001202-3112010031112001-3201011211132200"></a>

#### `retry_policy.retry_condition` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3232203011321132-1200220222313300-0222313113000111-2210230121110011-1033231311222001-3213020130210323-3330003232013232-2301313032100120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `retry_policy.back_off` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [retry_policy](resources--virtual_host--reference--group-002.md#canonical-3120210313011012-3220120012303030-2320221131003020-2003111301323011-0131032003311130-3311130033201323-1332312222032333-0100213232033200)
- retry_policy.back_off

<a id="canonical-2022202132020122-0313220232010302-3323220003131000-2031030233333230-3131220223033201-0011222012332112-0112320230011102-2023131003303230"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
back_off {
  # Configure direct properties listed below.
}
```

<a id="canonical-1330003022232110-0030132122220001-2302130310332011-0123130131120133-1233112213203332-0030031213203032-1111130000322302-2301102000312102"></a>

### Direct properties for `retry_policy.back_off`

<a id="canonical-0000230233313210-2020102122303230-3321200120031222-1303022020310110-1013313323231311-1333003213221020-3001311030132301-3011222031210332"></a>

#### `retry_policy.back_off.base_interval` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3231131112331230-2022002123220210-3113011203301202-2302330310302201-1333332112023330-2310323300012111-1332023211003102-3032211120133331"></a>

<a id="canonical-1132332312100003-0111032002223301-1020233210123300-3113302230123030-1022121021220013-0003022232123010-3013300231230120-3323201001131130"></a>

#### `retry_policy.back_off.max_interval` property

Type: `"number"`. Optional.

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

<a id="canonical-1230120312310211-1032010120013201-1323220112101321-0101321330220132-1310202220130210-3031233111323031-2123332330332300-2030220320303223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- routes

<a id="canonical-1001231111002131-3010222122130310-3000212310212132-2332312133330201-3310103133323231-3122230203211313-0231120223222120-1013010013030320"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203332011202103-3312112230113330-3022230011132010-3131100222111221-2100031222002021-0121221001223333-2023213131132010-0210103111000213"></a>

### Direct properties for `routes`

<a id="canonical-1332302203002203-1323131112012302-3212130321233131-2222020231232012-0112202202120210-0003100001213322-2223322311313002-2012200003301002"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3233200101223303-0331333031002211-3222223003320121-1322231321313012-0121030223213223-0322132220323222-3021200203131133-1120233203031200"></a>

<a id="canonical-0321301102012312-0130312231201111-1211001001321310-0011123122012321-3013110133100300-1303130201220321-0111220010121101-3013031202030200"></a>

#### `routes.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2333233100103023-3300021033112020-3031120032113020-0122110223123013-0002132123121023-3033013230301313-3211201223013131-0220133132133222"></a>

<a id="canonical-0133332003131113-1221032333330000-1111303301311232-1100322231332030-1230133100012312-3233013011232123-1102123023131233-3323212032101031"></a>

#### `routes.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0103130202322010-0213301300021223-3112001333123233-0103211230330111-1321103321202003-0301023010232321-1303032013202001-2032030210101102"></a>

<a id="canonical-1323033112303330-3121112113230311-0330312330122233-3001330012202103-0132123333011222-0022130333210112-3130111003322200-1332201100210321"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2102332203002001-0323322233101200-1331003332313110-3223012300133232-1213122100320301-2020213310212303-2200001301312132-1311101333021310"></a>

<a id="canonical-0201303320023131-2232010101210301-2110303012122330-0003130032300121-2333313202223122-1213201131100121-0210303011231211-1332310011021213"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1113203023200020-3223013230231323-2212021220030131-0310020233201101-0121323110323203-3122333302133331-0211123332303330-1130022122020210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sensitive_data_policy` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- sensitive_data_policy

<a id="canonical-0222123113222031-2013231001121032-1101203330320312-2232211013122133-3202312112000010-0012121230102013-2101110323013332-3013221200021231"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
sensitive_data_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301111313230320-3012201202020311-2230300320120003-0331113130331233-3001332303020033-0023330323321032-2020130012013102-1002011231231223"></a>

### Direct properties for `sensitive_data_policy`

<a id="canonical-0131000121302011-2211210111203202-3202112101330022-3313330121012122-2021031221310010-3111111301310100-0110013310300230-3222333202322002"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3233322232232023-0203210130330331-0303200101211000-2132112013233123-3320231200012123-2210220313133102-2232031332322121-3100231132323232"></a>

<a id="canonical-2213312222221233-0210023001231011-2122232120123323-3111200233230032-0020332333132311-3123132200301313-3201021131300222-2220321212010023"></a>

#### `sensitive_data_policy.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3110012223312303-1313130030112020-0331100232120323-3321001203301003-0200233222200113-2323133010211300-1320001100121123-2233313303101100"></a>

<a id="canonical-3030332012323121-1001203320330301-2332302332020222-2031102200203200-1310000303002103-0030233012301021-2232002001202233-0210031231202000"></a>

#### `sensitive_data_policy.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0113201221033133-3130332221120333-1032123330203122-2122101012020032-3103202022233021-1210012302020032-0202132120312133-2231220202012031"></a>

<a id="canonical-1321130203331002-0203222233130200-2031120132301123-2323330031211232-1113330111123230-1112131023100303-0000223000031003-3100201222233102"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2110331321310100-1220302312211220-2313213200133332-1303221103013133-1201323220013221-1202021020312132-2200123111300130-0103030110123330"></a>

<a id="canonical-3130331231112132-1130130310120210-1323230223323003-0232102331231312-1111032221222103-1203223300032112-0300300311332020-2202201003010300"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2333312030122121-1100330320331032-1120232213211221-3131120330330231-2221003131100222-1230211011000221-1033321202131202-1012333211023301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slow_ddos_mitigation` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- slow_ddos_mitigation

<a id="canonical-0213112013313021-0210021331121222-0031003332330302-2013322310222003-2302001202130331-0023232133221022-2330101022202311-2132022201020312"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
slow_ddos_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023200001202302-0103003132331233-1310013232202123-1121131312032122-2011031003022333-0311102200300101-0313122023302130-3131032330331230"></a>

### Direct properties for `slow_ddos_mitigation`

- [disable_request_timeout](resources--virtual_host--reference--group-002.md#canonical-0302112002310101-2021323133201110-2302121112300222-2022200302212321-1020202132131320-1021002113020201-2320232221231103-3212312020311302): complete subsection reference.

<a id="canonical-0310121203030120-2202111030132111-3230013223020011-0232201133232112-3210322310111002-1020201320121213-2211132302123003-0003003120012112"></a>

<a id="canonical-1320203012320333-1100320111121122-0323210332002103-3010223211003010-3330310113213320-0130333301122213-2010221310003211-0332112022000221"></a>

#### `slow_ddos_mitigation.request_headers_timeout` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1010331231022331-1132221033201011-3230133101222221-2013121321230003-3031003203021123-1310310002323232-2012301111230332-3011023231213112"></a>

<a id="canonical-0102221122133221-1022020211010311-1212310032003123-3302131220313113-2120020013222122-1331120213132312-0230131233100031-0231231023300121"></a>

#### `slow_ddos_mitigation.request_timeout` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0302112002310101-2021323133201110-2302121112300222-2022200302212321-1020202132131320-1021002113020201-2320232221231103-3212312020311302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slow_ddos_mitigation.disable_request_timeout` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [slow_ddos_mitigation](resources--virtual_host--reference--group-002.md#canonical-2333312030122121-1100330320331032-1120232213211221-3131120330330231-2221003131100222-1230211011000221-1033321202131202-1012333211023301)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-0032120312201033-3200032031011122-0011133112303101-1033331011003302-3232321223332332-3022232232310020-1232233310013321-0201211221133333"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_request_timeout = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133311002322212-0233231012302121-3111302012013321-0301321213223223-0023120022022300-3003323023033302-0131213332220332-1033313111212310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- timeouts

<a id="canonical-1223303000101110-3011111313112031-1002101201030211-1001300332202033-0331102311120200-0233023330322221-2003333130122211-3310101221230210"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120210332031313-0232213211230300-2222320022033110-2113332303321012-2113202310302101-1000100320023030-1230312332333312-3221111310102233"></a>

### Direct properties for `timeouts`

<a id="canonical-1113332021000102-1022012011032030-2313220010330013-3333310312000023-1122113130313220-1013111221012123-0032002330101330-3012130212102132"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3030233233112021-1210231122222201-0022210333223313-0013233323323122-0011321101322012-1120210233132011-3303000120023020-2313103233221213"></a>

<a id="canonical-0333023101313303-2203310133113000-2210021123222322-2011003133223320-0101203312213210-2110032011101210-0010210020301233-3220020030001231"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1302101300113101-2232013120221032-2130001111123003-1011023213020302-0311120111212333-1223211000133333-0210003110220230-3230132211103030"></a>

<a id="canonical-3020022312213010-3302300230032012-0123210003222313-0033113231003021-0220102303120220-3210313102101133-3223331300311123-0022302232333202"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1310103030131222-1332220101011001-1012020331030322-1323110131211203-1231000312020102-1331323032132213-1133012111210022-3001023220230031"></a>

<a id="canonical-1211203233201033-3301333112033333-3313100001102231-1210210021132120-0002122012210303-3313310323201331-0202223111331001-0212310120222101"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- tls_cert_params

<a id="canonical-2123330110100202-3132023320131302-1212121211011102-3110211230311121-1121003310302011-2120012132033323-2102333031220302-2103213231223331"></a>

Type: `"object"`. single nested block, Optional.

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

- [tls_cert_params](resources--virtual_host--reference--group-002.md#canonical-2123330110100202-3132023320131302-1212121211011102-3110211230311121-1121003310302011-2120012132033323-2102333031220302-2103213231223331)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0221010032112232-0002311102011103-1002323023221012-1223320011202010-1100310130023232-0301233120212133-0021101133201301-2201231323313333)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223313210310003-3130320133220110-0020002300110120-1321100103200302-2020323332123231-0102330022300012-0202210332203312-2022210222212321"></a>

### Direct properties for `tls_cert_params`

- [certificates](resources--virtual_host--reference--group-002.md#canonical-1333132230002120-2220113212031031-1100313100101320-2112323110202120-2103200122222302-2101113102020123-3020020213230333-1303213032212222): complete subsection reference.

<a id="canonical-3332232212210020-0203212331130212-1300121313102121-1221223322011312-3333022100321011-2130003130102020-3332020313022023-0321302322211331"></a>

<a id="canonical-1302331222310233-0002032133320203-2021113202031331-2330211212102221-2301000110130200-2103331301333223-1321303201020101-3311011013301332"></a>

#### `tls_cert_params.cipher_suites` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [client_certificate_optional](resources--virtual_host--reference--group-002.md#canonical-2220311223310000-1311303333221131-0123212110113222-3113013210012112-3331322002321331-0123212233121233-2003331313111101-0301010133230130): complete subsection reference.

- [client_certificate_required](resources--virtual_host--reference--group-002.md#canonical-1003023201132001-2031223012013002-0033031311323313-1000012312300110-3220100130333033-2010311200120011-0223020021323130-0212233213122231): complete subsection reference.

<a id="canonical-3202210313202032-2103210032220213-2333313310011023-3320000022101131-0032313032232223-0101023100203130-2320333202112301-0003222302323300"></a>

<a id="canonical-2023023111200213-1310102301203203-2122223230321101-2103230020333130-1111020121002010-2121010300303132-0110320022210221-3302032032233331"></a>

#### `tls_cert_params.maximum_protocol_version` property

Type: `"string"`. Optional.

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

<a id="canonical-3113032321331133-1323121012031212-3213020313001013-3302110032002301-2213313313322121-0230022201023101-2332323023210222-2013122300311010"></a>

<a id="canonical-3102200032010101-3010201303313322-3211330211113300-2333221303110222-0200320020031110-3002212130301300-3321222001232131-0322323223300320"></a>

#### `tls_cert_params.minimum_protocol_version` property

Type: `"string"`. Optional.

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

- [no_client_certificate](resources--virtual_host--reference--group-002.md#canonical-2101132133101113-1022120223131223-3310112111001032-1310100230310130-3020001203133321-0211202000130300-3012200301033002-3323223231312321): complete subsection reference.

- [validation_params](resources--virtual_host--reference--group-002.md#canonical-2111212302120220-0101203001130310-1133231301002100-3121101200220032-0321213313100212-2200021303113021-0023002112012032-3022331122002112): complete subsection reference.

<a id="canonical-3200233311230203-0032131020020012-0003320011300330-2100210331110120-2023023332113213-2013120223101120-0313131322110120-3022030102321311"></a>

<a id="canonical-0322313210000210-0311320131211300-3201213003303113-2102201320210032-3202023020011131-2123031321202023-0130311113001301-3113211221321202"></a>

#### `tls_cert_params.xfcc_header_elements` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1333132230002120-2220113212031031-1100313100101320-2112323110202120-2103200122222302-2101113102020123-3020020213230333-1303213032212222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_cert_params](resources--virtual_host--reference--group-002.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- tls_cert_params.certificates

<a id="canonical-1002330313231012-3132011231022303-2330322112320020-3333223010230220-2230020323030221-3301101101021312-0001221121123213-3222202330100022"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032121102300303-3131122002332201-0333000231202230-0230013103102331-1220321201030021-1022330220010001-0010200331122201-1320301132100233"></a>

### Direct properties for `tls_cert_params.certificates`

<a id="canonical-2030112331132002-3113310132003032-2331232100230000-0200230011323213-0022103130310033-1103130211221313-2003301233203132-0300301222330133"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2000100313113101-1301313210100011-2200312233122122-2221320303323322-0102032301030311-0030302132331230-2332202022231223-0112023233333321"></a>

<a id="canonical-1110311310020331-1103321012101210-3330323100002013-2130313112002110-3121213101330233-1113123033022220-0312111331002010-2000002003000132"></a>

#### `tls_cert_params.certificates.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2320202202302103-3203100131231231-2202033310030023-3323131020333232-2202213202302212-3210222330301000-2320300010333030-1013121122130111"></a>

<a id="canonical-2330203322303221-1231012133201213-1120213031223302-2130012330110133-1032120030233113-3332211203112130-3032020123111103-3032231111330333"></a>

#### `tls_cert_params.certificates.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0210103321000312-3101022020303002-0222032131002212-0301330303330320-0110201223120300-2323233303233003-1221103312002120-1200112220101233"></a>

<a id="canonical-0311023203332223-0030230110321211-3323121000200023-0023103122030123-1231022312033010-2112332102320330-2102233233300000-2121000000212212"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0033103333210122-3131033102132023-2103000002320311-2201223212010302-2003002203313230-2311333021012310-3010020111030332-3212120020301103"></a>

<a id="canonical-2321313231330003-1322301112013313-3113313222001123-1000221333220313-0203010030330302-0122011121122220-1223230100021333-3203332101131201"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2220311223310000-1311303333221131-0123212110113222-3113013210012112-3331322002321331-0123212233121233-2003331313111101-0301010133230130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params.client_certificate_optional` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_cert_params](resources--virtual_host--reference--group-002.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- tls_cert_params.client_certificate_optional

<a id="canonical-1223302233121202-2213311121111220-0322222120322002-2333330312333021-0301101302320000-1331020233022201-0301232202203023-1322033023121103"></a>

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
client_certificate_optional = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003023201132001-2031223012013002-0033031311323313-1000012312300110-3220100130333033-2010311200120011-0223020021323130-0212233213122231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params.client_certificate_required` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_cert_params](resources--virtual_host--reference--group-002.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- tls_cert_params.client_certificate_required

<a id="canonical-2102010012200310-3201101313122222-3023032020132310-0123202022323332-0220210320300102-0230223003110030-3100321002312200-0300000220221331"></a>

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
client_certificate_required = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101132133101113-1022120223131223-3310112111001032-1310100230310130-3020001203133321-0211202000130300-3012200301033002-3323223231312321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params.no_client_certificate` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_cert_params](resources--virtual_host--reference--group-002.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- tls_cert_params.no_client_certificate

<a id="canonical-3310213000123010-3132121002122102-3103230132120112-3301100200202131-3310211001131130-2332131103111121-0010012310330323-1132323202112103"></a>

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
no_client_certificate = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111212302120220-0101203001130310-1133231301002100-3121101200220032-0321213313100212-2200021303113021-0023002112012032-3022331122002112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params.validation_params` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_cert_params](resources--virtual_host--reference--group-002.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- tls_cert_params.validation_params

<a id="canonical-1032202302300211-3001333330201033-2112023021131021-2121022110300031-2233123033220131-3032133101222000-3022121132022232-2331211111102331"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032013002112111-1202102003233113-1222033130312011-0310303110313100-2233301333213120-3010232113200033-2120313012230031-0110103100102221"></a>

### Direct properties for `tls_cert_params.validation_params`

<a id="canonical-0110332132111203-2133212021330330-3212100333320203-0022213121203020-0233000003322130-3211110100121032-0312120230130232-1323222112100233"></a>

#### `tls_cert_params.validation_params.skip_hostname_verification` property

Type: `"bool"`. Optional.

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

- [trusted_ca](resources--virtual_host--reference--group-002.md#canonical-3102322203102032-1010001013023210-2012213032232022-3012220300112201-1212033212213030-3021001301210011-1330200222200322-0000203010231131): complete subsection reference.

<a id="canonical-0102232033030023-0113332012201000-1312102323333021-2031133211031031-3313020313320333-1300232331030302-3302013232323323-1011120210323103"></a>

<a id="canonical-0201301313202103-2000001323032032-2000302212302022-2213202201022132-1211313033101301-1130111232013223-2300120123201112-3031102303210222"></a>

#### `tls_cert_params.validation_params.trusted_ca_url` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1330331320133131-3230022131231000-1320110033311002-0013002210032220-3201200210123011-3210101321101131-3020300122020130-3002301301100201"></a>

<a id="canonical-2102110233011301-2303131202200200-2123332333210323-2230122201233231-3122033020002130-1212033023003031-1201233110011011-1202200003102312"></a>

#### `tls_cert_params.validation_params.verify_subject_alt_names` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-3102322203102032-1010001013023210-2012213032232022-3012220300112201-1212033212213030-3021001301210011-1330200222200322-0000203010231131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params.validation_params.trusted_ca` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_cert_params](resources--virtual_host--reference--group-002.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- [tls_cert_params.validation_params](resources--virtual_host--reference--group-002.md#canonical-2111212302120220-0101203001130310-1133231301002100-3121101200220032-0321213313100212-2200021303113021-0023002112012032-3022331122002112)
- tls_cert_params.validation_params.trusted_ca

<a id="canonical-1002033131101220-0211131331011102-3222212131332030-3030112200103030-2121330011020101-1333203122110120-0322323002313131-1131020202321102"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123033302102100-1220322312103133-1213210230300221-2120200222003132-1333132110230012-0213032120221210-1133132122322131-1210010020311201"></a>

### Direct properties for `tls_cert_params.validation_params.trusted_ca`

- [trusted_ca_list](resources--virtual_host--reference--group-002.md#canonical-0231232332310003-0020000023320100-2002120131220321-2130121301021230-0201332013033300-1130032113023210-1112333030122301-3201313113210030): complete subsection reference.

<a id="canonical-0231232332310003-0020000023320100-2002120131220321-2130121301021230-0201332013033300-1130032113023210-1112333030122301-3201313113210030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_cert_params.validation_params.trusted_ca.trusted_ca_list` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_cert_params](resources--virtual_host--reference--group-002.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- [tls_cert_params.validation_params](resources--virtual_host--reference--group-002.md#canonical-2111212302120220-0101203001130310-1133231301002100-3121101200220032-0321213313100212-2200021303113021-0023002112012032-3022331122002112)
- [tls_cert_params.validation_params.trusted_ca](resources--virtual_host--reference--group-002.md#canonical-3102322203102032-1010001013023210-2012213032232022-3012220300112201-1212033212213030-3021001301210011-1330200222200322-0000203010231131)
- tls_cert_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-1111101121102001-0321110223122321-2230131322101220-0303201223120133-3123230230020330-3232201232300022-0001002321302333-1002213333312003"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1021201233310002-1331113033002010-1132320021013012-1303222032202111-3032310100030011-3322100011202312-0230113021010122-1032031021001122"></a>

### Direct properties for `tls_cert_params.validation_params.trusted_ca.trusted_ca_list`

<a id="canonical-1333112023322221-3313100101333133-1002320321022103-0220322010211232-3200303231013330-3201333310012122-0331001331123100-0232021113032202"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3323131031113030-3113210102313220-3103232022202130-3303030002230221-3323133211013021-3130223000030013-3011223210330211-0123211231100323"></a>

<a id="canonical-2010131103023031-2122223112000033-1131203112013112-0002100230012322-2003100302230102-2310213203030131-0012330303022311-0211133011201330"></a>

#### `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3121030130012330-2210330211031333-1132310012230213-1103113012210102-2213032002020302-0210123121221030-2110311331000333-1002331332001011"></a>

<a id="canonical-1311221031121233-0331300210000112-0100122231332012-3231203332102203-0111130201231313-3013301032202230-3202200300102332-0000011322131302"></a>

#### `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1100201012230233-0213301233120301-3033323121112133-1023222100113222-0002201103130021-1100313113220113-1212201220111220-1203033102002311"></a>

<a id="canonical-1222031132233302-2312010221202012-1021201023330010-0323112211100330-3131023203233011-2023031000221101-1322022012011220-2222322223103310"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0233133023123231-3003132311323220-2110120101023033-3323303310130133-1120030212121311-1232200123221003-1121130203203202-0010210011200322"></a>

<a id="canonical-3321210211003110-1001212022213032-1333003112021013-2213120233310312-0302323213220000-1010323013133030-1102010002213330-0120110101200120"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- tls_parameters

<a id="canonical-0221010032112232-0002311102011103-1002323023221012-1223320011202010-1100310130023232-0301233120212133-0021101133201301-2201231323313333"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202300301300021-0201230301102010-3133213031032021-1310112212332023-2030020331322110-1001021320131223-2233102230201022-1032130013313230"></a>

### Direct properties for `tls_parameters`

- [client_certificate_optional](resources--virtual_host--reference--group-002.md#canonical-2322133313323222-3132132032131311-2200332013110013-2311323311021120-2130302130010230-0122102013022233-3013312030021222-2121002101111211): complete subsection reference.

- [client_certificate_required](resources--virtual_host--reference--group-002.md#canonical-3202130002100313-3301122123011011-0330322120202303-0333011230321202-1223210103310323-2303032110201303-3012000030223131-3231311133101122): complete subsection reference.

- [common_params](resources--virtual_host--reference--group-002.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221): complete subsection reference.

- [no_client_certificate](resources--virtual_host--reference--group-003.md#canonical-2213212201113300-0223221321112131-2300013131310300-2330130313033001-3003001223303232-3000310231222000-2101120210302122-2300231020001031): complete subsection reference.

<a id="canonical-0213200310122130-0101300331002010-3102233222333003-2122132330021030-3201110012210001-3122300102120002-0031012231012200-2300100230221000"></a>

<a id="canonical-2300231130210032-0202022323200113-3023213011202030-1102212000231313-2102020002303332-2222231331213230-1212003120013103-0123223030131030"></a>

#### `tls_parameters.xfcc_header_elements` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2322133313323222-3132132032131311-2200332013110013-2311323311021120-2130302130010230-0122102013022233-3013312030021222-2121002101111211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.client_certificate_optional` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- tls_parameters.client_certificate_optional

<a id="canonical-0333121111220133-0323020310031210-3333011101131031-3201121301201312-3211211020311333-0123102320230100-3032302023333002-3010000221101202"></a>

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
client_certificate_optional = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202130002100313-3301122123011011-0330322120202303-0333011230321202-1223210103310323-2303032110201303-3012000030223131-3231311133101122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.client_certificate_required` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- tls_parameters.client_certificate_required

<a id="canonical-3130232032303203-0232011333223321-0030113131322311-0313010120132021-1001000100003301-2121210321312131-2013311200100131-1313302322110000"></a>

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
client_certificate_required = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- tls_parameters.common_params

<a id="canonical-0303120133133233-0132220103133322-1022302231330312-1302033300200132-0301031130022023-3013200203121330-3031202120031130-0010023113130121"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-1031103330110211-0101210231203021-2223103233122030-3122201003331331-2313123211223010-3233121023000023-0302021323121200-0311100020323201"></a>

### Direct properties for `tls_parameters.common_params`

<a id="canonical-2310321300133003-0230000103032103-1231003120202213-3113000213210122-1131113331001110-3222122220323311-1212222013202333-0213312132100201"></a>

#### `tls_parameters.common_params.cipher_suites` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3133332330120023-2322121212320313-2131311120313023-2222221001230020-3221112322200303-3231220021213000-0022100122003132-2333333332323110"></a>

<a id="canonical-1301300101003101-3322132033112203-2322213321031301-0023103220133113-3023000333131002-2010002202210121-0031013223132312-1032021111003110"></a>

#### `tls_parameters.common_params.maximum_protocol_version` property

Type: `"string"`. Optional.

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

<a id="canonical-0213203310111320-3203301022001113-0203122220203000-3200302123110222-3231321112312320-1303333303122323-3232311230032312-0323311013031321"></a>

<a id="canonical-0222203000002123-2031022220122100-3200233223320113-3101323213310133-0222112202221013-1323311100222311-0312221202331111-3212312131312300"></a>

#### `tls_parameters.common_params.minimum_protocol_version` property

Type: `"string"`. Optional.

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

- [tls_certificates](resources--virtual_host--reference--group-002.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203): complete subsection reference.

- [validation_params](resources--virtual_host--reference--group-003.md#canonical-1100132303312233-2302331223202323-2232231202112111-2233122222201102-3320121233323003-0022333231032202-1210011013012132-1022023001212123): complete subsection reference.

<a id="canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-002.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- tls_parameters.common_params.tls_certificates

<a id="canonical-1011323231323101-2001121311000010-1021211103001120-2203300311211232-0231231033120330-0102330323332022-1001313320010300-1222312131021200"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222312033203012-2310212010220201-2311031210321032-3230113200332301-3301112312323100-0223301230002223-2033311311110022-1222331300301323"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates`

- [blindfold](resources--virtual_host--reference--group-002.md#canonical-3122321303023200-3322022033201103-2132001313022101-1230300222021130-1011230033131322-2323310332013032-3233200100132233-0331330031002123): complete subsection reference.

<a id="canonical-3302211131233331-1203311302023312-0302121312102232-2033201220322003-0033302011231110-2211111233000312-0111112211001122-2012210123233002"></a>

<a id="canonical-0011330000300003-0232002132220312-0323031231210121-1101213331133212-3113010300231333-2022033233121130-3210033132021220-2232232210033310"></a>

#### `tls_parameters.common_params.tls_certificates.certificate_url` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [custom_hash_algorithms](resources--virtual_host--reference--group-002.md#canonical-1000011133010022-3111113122203210-0111223313002300-0133022021203110-0330111101023322-1230023110021010-1223202212030000-1320021132303212): complete subsection reference.

<a id="canonical-1011010332212323-1313312303231200-0033031122132330-3033330310233021-3030002010001233-2123323233323022-2120221301330301-2322220212223200"></a>

<a id="canonical-0213003303301332-0023101021001330-3132021101223131-1132323033210231-0220003201311312-0323321111000131-1332023001203002-1200212033202101"></a>

#### `tls_parameters.common_params.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--virtual_host--reference--group-002.md#canonical-3230303332000320-0100030312022332-1212333031013131-1303111211023103-2201222130212130-0230030321331231-0003312300322110-0003032221332333): complete subsection reference.

- [private_key](resources--virtual_host--reference--group-002.md#canonical-0100101001331011-1232022323322231-2223232000102312-0213110033331032-0321030213022011-3300202030223321-1111110020102320-1230230111233113): complete subsection reference.

- [use_system_defaults](resources--virtual_host--reference--group-003.md#canonical-1321321311010002-0333130122331222-0233010210103222-1020300033133113-0220033020222102-2031100330302111-3223213023203230-1020203123301203): complete subsection reference.

<a id="canonical-3122321303023200-3322022033201103-2132001313022101-1230300222021130-1011230033131322-2323310332013032-3233200100132233-0331330031002123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.blindfold` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-002.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-002.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- tls_parameters.common_params.tls_certificates.blindfold

<a id="canonical-2131013333202302-3322331311023130-0330332231222302-2032102231113211-0030130202021103-1201220223102011-1223002113201110-3200023113321023"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-2230301201102102-1322200202212033-0233031101121303-2021212111031203-3211121313031230-2020111233001323-2302102203021312-2012032003301013"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.blindfold`

<a id="canonical-2111030032211111-1222300110123223-0023303311201212-3231110203033023-1220223233232232-3221133323100121-1330030111303013-1230323320223021"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-2113122000033331-0122330312232220-1303020331021011-3313313203012131-3102323310120231-1101331032203030-0210311202330220-2213112302001331"></a>

<a id="canonical-2022003020031102-1230331220322131-1013133310003222-3011300001102231-3111321021033302-2100221031033223-3331020113113312-1022223030232333"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-1030320302211010-3122001322023102-2300033012103112-1100220222311233-1101032133312103-0111210033231312-2321031201220032-3322003102223013"></a>

<a id="canonical-1203230221120330-3010212313002211-0003021202013320-1002200213103303-1332101122331333-0333223212011231-2313001310213312-1132301200330310"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-3110112030112010-0032121303123322-0310022021333001-3032233100213113-2331011010121133-2021022200111223-2232013321000000-0311201333330311"></a>

<a id="canonical-3023011033210312-1101112022001302-2010022301130302-0221101120333123-3221301020110012-2103012112013123-1032132100131310-0130310231201213"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-0321331303223322-2011023100101102-3333213131020231-2200203323132112-1121230123222320-3110200113023033-1132313133230323-3031302230032302"></a>

<a id="canonical-3220031231231223-2331221132221020-0112010210330331-0000211230321331-2113103200320230-0123203302210101-1000302311332122-3120133300023323"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-1032232100121130-2103013330130102-3120203030102213-0200202220320000-1310313223202000-2122231220013123-3023312301032230-1130301300101301"></a>

<a id="canonical-3132233003303103-2002222100001210-1300023230103031-0030313103120301-2111300303330332-0231133333223332-3212110302031130-0120330133131233"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-2111233133033323-3133333310320200-2231210303201311-2032331120020002-2010033011002221-1213112103000330-0123030233220133-0213303302112331"></a>

<a id="canonical-0111113303313120-2130312001223311-3203123010123030-1312022221100122-0220321320312201-1332232013010033-3302313023013321-0112311032213310"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-3112330313312111-0210021000011312-2020011203333211-3223220232223211-3301120110230212-3223230111322320-1030101022230033-3311221203303100"></a>

<a id="canonical-0103220311110303-1102322010232103-1103313130033212-1130003113132133-3132032013011201-3230032233302132-3103321313232220-0111122223233200"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-0232312301212302-2122031033033002-2023202120133220-0201312002203230-1013102102210200-1111123033331102-2222132030101111-0332110300012322"></a>

<a id="canonical-2020021020330101-0201201220220101-3130310032113333-3101310313203120-0011122202113322-0132002231032303-2322210001331210-1032003301323213"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-3210112011002302-1201213023003100-0123001010101201-0203201311112131-1200333233231031-3310221110030231-0113333333111301-3330111222021001"></a>

<a id="canonical-3331312312012101-3212233232022110-1000031002222333-3032102212302131-3010303201100002-3200333332231111-3123301122232031-0110312032032003"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-3113121133002313-0222030222022221-2001302011100310-3000000323011310-3200003310330130-0101030021012320-3323331323103033-3121001020303231"></a>

<a id="canonical-0133030230310003-2111032303010212-3002330131101132-2022320103103032-2310003102103120-0123211300221223-0110301023220001-0100100330012201"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-0211323101122220-1323332330122102-3320003301201211-3110122123210133-0331002002110302-1220200312120300-3012303212323230-0313023100233121"></a>

<a id="canonical-1010030101130201-1303331113311030-2102130321120133-1333223222322202-1310301203011312-2033230310332110-2132130012001032-3202333022303230"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1313211020223001-1023332330122210-3112122321032301-0001213033233210-2233310210030132-2332330103221000-1210332203331311-3101101321303002"></a>

<a id="canonical-3011120112222310-3101102332121112-1233023012313231-1102030331312233-1330011213022330-2311003220233101-3231120132120232-3233101111002202"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-2230312111130110-1333110130101210-2002323122310031-2201332333312123-1103033120201023-2312320300030102-3332222311022113-1123113121003302"></a>

<a id="canonical-1030023121222123-0010212013233031-1322130311000232-3113232122213102-2223332002220333-1011300300031031-3012221132033210-0331123101120130"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-3231222230122000-3312032303110203-0323320323131020-1022122333103111-2202211311111323-3111123320220010-1233323233231220-1333331013011022"></a>

<a id="canonical-0130323012302002-3113012213330132-2223300203023020-1222130321302330-2020010211332232-2110032320200212-1002201002203011-3203113122130100"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-1231300322121223-1010211020103120-0301023122122121-3001232210230301-0303021333003001-3300333312201213-2130001322110303-3231120201212203"></a>

<a id="canonical-1102000122121233-3300300110033203-0130310330313311-3022300213320202-2100121333200000-0013313033311033-0300113121023122-3230211221122200"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-3231010013322311-0133101023333222-3010301330000130-0102301013121212-1321302220301022-1021312013331030-0020201130120032-3012023200032001"></a>

<a id="canonical-1102032020210310-1333132233031101-2011121013133123-1231311302223111-2221223013301221-0011333033110233-1003031113333222-0332222122001100"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-2110101130300201-2332110320221020-2301200230013001-3011020032123323-3102113030123312-3310303333300210-1002011121011112-3130203211330200"></a>

<a id="canonical-3133303300102223-0233322111022113-1320322122222101-0202011220222222-1232331202113230-1321132212111033-0230032323203022-0210332100221121"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1330111212230013-0002023320030130-2113002230100303-1020302300002300-3113221330223213-0203212320321020-3131021300101211-1020310320230223"></a>

<a id="canonical-1112311220002113-2031130111111113-2213232313310232-0230323111313132-3111122301202222-3000201100122200-3121101223103021-1301013011333123"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.spki_identity` property

Type: `"string"`. Computed.

<a id="canonical-1000011133010022-3111113122203210-0111223313002300-0133022021203110-0330111101023322-1230023110021010-1223202212030000-1320021132303212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-002.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-002.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-1301331000212220-0123310303202130-1202122222101313-0021222302101313-3330133233100022-3013103103331121-2002100321130302-1031320201012133"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323110232232330-3301102323123221-0012210030233031-0312220201333220-1030331202200110-1313221022001022-1113333301110111-3002133012310210"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.custom_hash_algorithms`

<a id="canonical-2101210120001021-3312301120323023-1100101002233120-3322221303231330-3320000311132031-1330223130103001-3302010013211200-3331031002333311"></a>

#### `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3230303332000320-0100030312022332-1212333031013131-1303111211023103-2201222130212130-0230030321331231-0003312300322110-0003032221332333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-002.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-002.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-0101302010000031-3023011131112013-2233103023132011-2112211322201222-0333030013211102-1112301112212301-1323133003022312-1222332133003033"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100101001331011-1232022323322231-2223232000102312-0213110033331032-0321030213022011-3300202030223321-1111110020102320-1230230111233113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-002.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-002.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-2302212021121312-0000312312213033-0102223312230203-1231320211020201-1302122302102123-3130030311202313-1210120133033102-3012110331121213"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201002011011313-2121231322333130-0021212221313000-0001210220203230-3021111102330320-3021320031212323-0310201100131010-3232001212201110"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key`

- [blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-2303223212323122-2111120020132110-2303311132222211-3122020120123201-1301200220013030-0121301312210220-3202033203132021-2010113331023333): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-2321012223121233-0320222032102230-2330222220013231-1100112312301103-2312121002131223-3302012231330100-2130300301011211-3120012312133212): complete subsection reference.

<a id="canonical-2303223212323122-2111120020132110-2303311132222211-3122020120123201-1301200220013030-0121301312210220-3202033203132021-2010113331023333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-002.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-002.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--reference--group-002.md#canonical-0100101001331011-1232022323322231-2223232000102312-0213110033331032-0321030213022011-3300202030223321-1111110020102320-1230230111233113)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3231101123122320-1201100232102023-2010212231312012-1001113320211332-3213000223222222-2013031102003301-2131313102300223-3023321110213033"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3002001010000323-0221021330011113-3001300321111120-1020122020201010-2300232030223203-3103030130133030-1123002032011313-2220312320220220"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-3030230232232312-3303112101112320-3212030010310032-1010013100021230-3031122122133033-2220333111231020-2003122101201021-0213213332311220"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3100331112122311-3220232233111123-0321301322130101-2002001331113332-2202302003200313-3322032002010310-2322303201033033-2113030033012023"></a>

<a id="canonical-0203003230200123-0030312100320110-2103333010100311-2101112223102203-2120010201130232-0212021300223132-0022202323111130-0333201032001333"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0303011010303232-3221003323310310-3110111121213110-0123330223321133-2323313211311122-1233331330121300-0032333223022030-0232030103311203"></a>

<a id="canonical-2233300002303213-1033313331331201-2233331233331213-0222300132110113-0020102112021103-1301033213112220-2123312111330131-0011023303212023"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2321012223121233-0320222032102230-2330222220013231-1100112312301103-2312121002131223-3302012231330100-2130300301011211-3120012312133212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-002.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-002.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--reference--group-002.md#canonical-0100101001331011-1232022323322231-2223232000102312-0213110033331032-0321030213022011-3300202030223321-1111110020102320-1230230111233113)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-0301020201000130-1303100322320331-2223120202001311-0022210010221201-3120021103322203-3002120231130230-2002132332210033-0030221312122011"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211031122221310-3131033323020012-1233030000323012-2202300013120112-2212330300220310-1322302223223320-2100012310203112-1212030100130030"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info`

<a id="canonical-3313010330101223-3000011030300100-1313321023120311-1230220310221010-2011111133113232-1330302100232100-0131211033333122-0102000321130332"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1022201333211033-3102322311210000-3323103323031202-1213302020101211-2330103303103001-0202000111131313-3223221132332232-1031323102102212"></a>
