---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-2210220133203030-3332033032010023-1022131210020212-2223021021230232-3233103120301210-1210130121301102-1031121022223200-1231200320222213"></a>

## voltstack_cluster.az_nodes.local_subnet.subnet_param — subnet_param / 013221200223 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.az_nodes](resources--azure_vnet_site--reference--group-008.md#canonical-1110030203112331-1013331332022220-1131111111202200-1313201113200032-0100300023212210-1230202123102322-2322310110221331-2322301103133021)
- [voltstack_cluster.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-008.md#canonical-1132001220313023-0311122233001020-2302231132120203-0333203213323333-2123300302232030-3011010101320130-3203200033122300-1333101030103021)
- voltstack_cluster.az_nodes.local_subnet.subnet_param

<a id="canonical-0102002333320300-1021220131001103-0311300330012021-3130101122133022-3230332210231011-2223220232230132-1332133300330001-3322031000111232"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332032220320123-2003103301000210-2323020302201321-3202032132000010-0113222120301220-0321010331221210-1103202000111212-2311331300112231"></a>

## Direct properties — subnet_param / 013221200223 / 3

<a id="canonical-0121111331032312-3202302101031110-2202103100330120-2100311022102100-2130330112031200-3130311213102013-1113133101111302-3223003130031011"></a>

<a id="canonical-1311101112213313-3302212102120310-2312001332200001-0223322310022000-0001033021032012-3323122120122331-2211133013221330-3321123333113320"></a>

## IPv4 property — subnet_param / 013221200223 / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-2303123101003100-0301003020211330-0111202030000000-3222333211131331-3203223122110200-3301311110213223-2230132011213211-1301112001112330"></a>

## Next pages — subnet_param / 013221200223 / 5

- [voltstack_cluster.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-008.md#canonical-1132001220313023-0311122233001020-2302231132120203-0333203213323333-2123300302232030-3011010101320130-3203200033122300-1333101030103021)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2131023101022201-2303100013112312-1313313100100112-0110212020122023-0210320020111301-1210330123131001-2010103112113320-3030303103310022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012202000303120-2122232102322100-2000220010130122-3133322223233122-3002333030112203-3020202113201123-2223300013230000-1320000011203112"></a>

## voltstack_cluster.dc_cluster_group — dc_cluster_group / 221011131113 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.dc_cluster_group

<a id="canonical-1213023333113223-1011223310020022-2321220010223320-1023302013010130-1000201021203130-2230031332013203-1130230023113230-1101320031312301"></a>

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
dc_cluster_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-0131231212203032-0013030110333010-0002013102022120-2301321311201031-0321101230000301-1132320330311030-0301132310121233-1220212001313113"></a>

## Direct properties — dc_cluster_group / 221011131113 / 3

<a id="canonical-2120033103100120-1300000030201201-1201200020021212-1112322033133302-0100233203102021-1232212100101220-0310111031030330-2230332201020310"></a>

<a id="canonical-3120133101131003-1103102101233310-1310102011133133-2001001121033220-0132002123002202-2011001111133122-3222321211132231-3212102210233021"></a>

## name property — dc_cluster_group / 221011131113 / 4

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

<a id="canonical-1322112300122101-3320020211323110-2100121112333301-1331003131222221-1212223030210133-3023112231130121-3010322330220213-2331122210220002"></a>

<a id="canonical-2022313230233113-0113033123233133-1313133102123023-1001110000202310-3313130323223011-2131202202021231-3311212233120321-1331231320322223"></a>

## namespace property — dc_cluster_group / 221011131113 / 5

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

<a id="canonical-1131201212132001-0021000022303230-1132100313303020-2223120200333101-1110301201201300-3333230033123003-3033233221211320-1330112313101300"></a>

<a id="canonical-3000100311120020-0101101102020120-0101231110111120-0220132333103012-2230122133010321-2011101222000222-1000013230132003-0302221120320332"></a>

## tenant property — dc_cluster_group / 221011131113 / 6

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

<a id="canonical-0300301322021300-0302110200101303-3330200321322231-2311102120111132-1332123233120331-1220132300213212-2233202112223113-3323013110203311"></a>

## Next pages — dc_cluster_group / 221011131113 / 7

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3231030123031233-2321122232032332-0111011100221313-0121010311223223-3123111032023331-2301221310121230-2332103012002200-0300320110100211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313030211100313-0201001021123201-3033122121231323-3313202332011220-2202303310022213-3201322130303333-3323312020133011-2320312310303100"></a>

## voltstack_cluster.default_storage — default_storage / 220031031112 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.default_storage

<a id="canonical-1032100102012123-1230021022103033-2310222032232130-2133211120102000-3230031012101232-0021303101220030-0002032110021332-0102023303122221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default storage.

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
default_storage = {}
```

<a id="canonical-2131213103030300-2200123003313302-1230301101331022-1210121330311103-3033020223113013-1311023323212101-2210303301332232-3222231330010001"></a>

## Direct properties — default_storage / 220031031112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002131331202321-0231131233302323-2302232301120133-3000320123312232-3122200020130300-2311202133311113-3002222113303132-3111213033120132"></a>

## Next pages — default_storage / 220031031112 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3001113132222023-0130111230232332-1003013200230301-3132233032321131-0113133333000021-3222300233023303-3112022230123232-0312301003101033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003233222121111-1102203110011003-1300003202002000-0101030331003013-3202130020010001-1222313130303013-2201221310003031-3121330210220110"></a>

## voltstack_cluster.forward_proxy_allow_all — forward_proxy_allow_all / 302233013113 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.forward_proxy_allow_all

<a id="canonical-0212003030202222-0232132323231001-2001101030112231-3011333330302033-0310111201120001-3013301332331222-2222032333330001-3123020303210121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for forward proxy allow all.

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
forward_proxy_allow_all = {}
```

<a id="canonical-3020320310131301-1201023102320203-1313013313221130-1300213302001002-0123231132010313-1213322131222110-0021033323230010-0213302121122032"></a>

## Direct properties — forward_proxy_allow_all / 302233013113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012121001023211-3210202311303210-2012012102102311-3013310131033100-0323103032200011-0220303030130320-3330032030320022-0301310233223130"></a>

## Next pages — forward_proxy_allow_all / 302233013113 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2220210311032001-2331030022323023-3230301222203031-0003110121033101-3301121312120023-0223331112203311-1110012300221302-0031310321121233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321033300333300-1330023121321333-0131112010232111-2312223113123220-1212213200312330-3203012211233010-3203230321222022-1302132023322133"></a>

## voltstack_cluster.global_network_list — global_network_list / 023021320301 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.global_network_list

<a id="canonical-3300122013130122-1031320001110113-2131331201301011-0331313223022132-1130322133131132-3322130333323213-0113303002023111-1301302030023221"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322020012102120-1033122302011012-0302111120301101-0112011120320230-1310030201102133-3333130200300020-3220131121031221-1131212303310211"></a>

## Direct properties — global_network_list / 023021320301 / 3

- [global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-1033332000030130-3310210202231022-1121022301302203-2322031031012222-0032003012003102-1133032221223121-1233300203320120-1002102111231101): complete subsection reference.

<a id="canonical-0110220102223030-2201311023312032-0111112012231302-1133123023300212-2002120201312220-1132323000111002-2331013330221302-3311330133001001"></a>

## Next pages — global_network_list / 023021320301 / 4

- [voltstack_cluster.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-1033332000030130-3310210202231022-1121022301302203-2322031031012222-0032003012003102-1133032221223121-1233300203320120-1002102111231101)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1033332000030130-3310210202231022-1121022301302203-2322031031012222-0032003012003102-1133032221223121-1233300203320120-1002102111231101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011003120312103-0001332121130111-1012211030221223-1202231300211223-3321232122303001-2310232001032100-2221020111123102-3322001102100320"></a>

## voltstack_cluster.global_network_list.global_network_connections — global_network_connections / 323300320103 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-2220210311032001-2331030022323023-3230301222203031-0003110121033101-3301121312120023-0223331112203311-1110012300221302-0031310321121233)
- voltstack_cluster.global_network_list.global_network_connections

<a id="canonical-1102311301122123-3112203011002321-1023312021211222-2110201031030010-3012210122303112-1031311331112302-2210211020331131-0002022033121003"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
global_network_connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130011231202311-0023031303133031-0301033032310222-2313011100220112-1003103001030221-3223201120033110-1030013033211033-1120332200333021"></a>

## Direct properties — global_network_connections / 323300320103 / 3

- [sli_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-0223121010121213-0013212312223120-1232030011102101-2303130011211001-1000012233211201-2032331023012321-0203121213302123-0113013101111310): complete subsection reference.

- [slo_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-0323121202121301-0213113313230200-1101131332123201-2103133232111000-0010031133203200-0003011222212302-2131033330110012-3022003311203003): complete subsection reference.

<a id="canonical-1000202012002120-2013100001000222-2202133220300013-2311223133102121-0103201301121113-0311102202000232-1323200113010000-3221312013110323"></a>

## Next pages — global_network_connections / 323300320103 / 4

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-0223121010121213-0013212312223120-1232030011102101-2303130011211001-1000012233211201-2032331023012321-0203121213302123-0113013101111310)
- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-0323121202121301-0213113313230200-1101131332123201-2103133232111000-0010031133203200-0003011222212302-2131033330110012-3022003311203003)
- [voltstack_cluster.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-2220210311032001-2331030022323023-3230301222203031-0003110121033101-3301121312120023-0223331112203311-1110012300221302-0031310321121233)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0223121010121213-0013212312223120-1232030011102101-2303130011211001-1000012233211201-2032331023012321-0203121213302123-0113013101111310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120110312103030-3232230020130013-1200030010203220-2020010122220332-0112223330112130-3121023012230221-1010001021230001-2112220020310322"></a>

## voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr — sli_to_global_dr / 323110122200 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-2220210311032001-2331030022323023-3230301222203031-0003110121033101-3301121312120023-0223331112203311-1110012300221302-0031310321121233)
- [voltstack_cluster.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-1033332000030130-3310210202231022-1121022301302203-2322031031012222-0032003012003102-1133032221223121-1233300203320120-1002102111231101)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-3221210312322011-1210300012021020-0022023202203130-1013022312310212-0000302032130003-2030223033020302-1111100121210310-2032201122111231"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
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
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211322113030112-2332230011120012-3001302110221130-1033132033110032-3223011221123311-1120200211201032-0030101223021222-2010000103100113"></a>

## Direct properties — sli_to_global_dr / 323110122200 / 3

- [global_vn](resources--azure_vnet_site--reference--group-009.md#canonical-0030223033333312-3211300111012301-1000101202132113-0010202310303232-3210221000212023-2000131131221213-2220010002101212-0033010000320020): complete subsection reference.

<a id="canonical-3231133130121311-3022111131112223-1222110211333022-2121001310033213-0330120032330133-3221103012000010-1230330020111333-3220132123301201"></a>

## Next pages — sli_to_global_dr / 323110122200 / 4

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--azure_vnet_site--reference--group-009.md#canonical-0030223033333312-3211300111012301-1000101202132113-0010202310303232-3210221000212023-2000131131221213-2220010002101212-0033010000320020)
- [voltstack_cluster.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-1033332000030130-3310210202231022-1121022301302203-2322031031012222-0032003012003102-1133032221223121-1233300203320120-1002102111231101)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0030223033333312-3211300111012301-1000101202132113-0010202310303232-3210221000212023-2000131131221213-2220010002101212-0033010000320020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110123320332032-1301323312001233-1033301022212103-2131201212031122-0122303133031230-2220231133312333-1030323032311103-1030123102313033"></a>

## voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn — global_vn / 330133201010 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-2220210311032001-2331030022323023-3230301222203031-0003110121033101-3301121312120023-0223331112203311-1110012300221302-0031310321121233)
- [voltstack_cluster.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-1033332000030130-3310210202231022-1121022301302203-2322031031012222-0032003012003102-1133032221223121-1233300203320120-1002102111231101)
- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-0223121010121213-0013212312223120-1232030011102101-2303130011211001-1000012233211201-2032331023012321-0203121213302123-0113013101111310)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-2222012330133202-1002031303223120-0103023131002120-2011202203031221-3300032023320311-2031010001003012-0001031312023303-1332330303110110"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211301201110223-0301031330113121-2321210211022031-0232230320202120-2330220021111302-1120120223320322-0130302220221320-3220312021011303"></a>

## Direct properties — global_vn / 330133201010 / 3

<a id="canonical-1203031333131022-2120020003313001-1232130231111023-1010032001200011-0120220032001212-1323103003212320-1333331333223322-2210310122130332"></a>

<a id="canonical-1022301203332312-3121023130110101-1102013121233023-3100333212100110-2201003130222000-3121211201331101-1012110122320012-2010213300231030"></a>

## name property — global_vn / 330133201010 / 4

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

<a id="canonical-2330020201120202-0112323223010203-3011100311130123-2113003122002102-0221320323103020-2010031132333220-1230210331213123-0301303312220012"></a>

<a id="canonical-3020331303333230-3101310031001200-2131032213000120-2131132300112321-1202313111223211-1221033331311312-3301320030220330-1210200200111110"></a>

## namespace property — global_vn / 330133201010 / 5

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

<a id="canonical-1231200210012313-2131010012000003-3221321100101011-1313311002331320-1221322201010322-3122313322102223-3331203323332011-2313312031022313"></a>

<a id="canonical-0312003333311000-3110102321110322-3303331023003300-0121211310323010-3200300200212210-1110201133300120-1233130123111202-2001022113001221"></a>

## tenant property — global_vn / 330133201010 / 6

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

<a id="canonical-2011330333300120-2300310222113022-3313121310033220-3300122312301120-1030201121233031-2010301132202033-1110011212311021-3101333300120301"></a>

## Next pages — global_vn / 330133201010 / 7

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-0223121010121213-0013212312223120-1232030011102101-2303130011211001-1000012233211201-2032331023012321-0203121213302123-0113013101111310)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0323121202121301-0213113313230200-1101131332123201-2103133232111000-0010031133203200-0003011222212302-2131033330110012-3022003311203003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320323301031112-0133231233130203-2133003323220200-3222110302332211-3321221331330322-0111130201233010-0221112330331313-1231103013020302"></a>

## voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr — slo_to_global_dr / 301101001123 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-2220210311032001-2331030022323023-3230301222203031-0003110121033101-3301121312120023-0223331112203311-1110012300221302-0031310321121233)
- [voltstack_cluster.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-1033332000030130-3310210202231022-1121022301302203-2322031031012222-0032003012003102-1133032221223121-1233300203320120-1002102111231101)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-0320211113333020-2310300113113302-2203301312012303-3330113022132110-1000031022132011-0110211302102320-3011120320331303-3030031231330313"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
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
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132312101130331-0201201310120203-2203101102203010-3131010010110213-1033003031223211-1211133230123111-0222010102113121-0201312321023220"></a>

## Direct properties — slo_to_global_dr / 301101001123 / 3

- [global_vn](resources--azure_vnet_site--reference--group-009.md#canonical-3012020201303312-1200312231033120-0123030322201312-2212122213212130-2103110223021203-1131003022221223-0311220231110133-0223133132331322): complete subsection reference.

<a id="canonical-1003032310010132-0033221333102111-3230311033002300-3022231130330210-2000311201311012-2232120221121000-1031212031020001-1221211132111013"></a>

## Next pages — slo_to_global_dr / 301101001123 / 4

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--azure_vnet_site--reference--group-009.md#canonical-3012020201303312-1200312231033120-0123030322201312-2212122213212130-2103110223021203-1131003022221223-0311220231110133-0223133132331322)
- [voltstack_cluster.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-1033332000030130-3310210202231022-1121022301302203-2322031031012222-0032003012003102-1133032221223121-1233300203320120-1002102111231101)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3012020201303312-1200312231033120-0123030322201312-2212122213212130-2103110223021203-1131003022221223-0311220231110133-0223133132331322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313303203122131-1002232103201021-0103020131102033-0132112310001003-0221231330301002-1211220032123102-2321001333210323-0211003031012020"></a>

## voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn — global_vn / 221012322112 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-2220210311032001-2331030022323023-3230301222203031-0003110121033101-3301121312120023-0223331112203311-1110012300221302-0031310321121233)
- [voltstack_cluster.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-1033332000030130-3310210202231022-1121022301302203-2322031031012222-0032003012003102-1133032221223121-1233300203320120-1002102111231101)
- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-0323121202121301-0213113313230200-1101131332123201-2103133232111000-0010031133203200-0003011222212302-2131033330110012-3022003311203003)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-0302323112302012-3033100022102210-1320130121011322-0200131312032130-0102201011023123-3002220313133101-1230200131113233-3133000113301100"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212310322220311-1311010130033013-2200330110121322-1132002111323030-2023012220120323-2312333311011213-3211002220033110-1323001310101201"></a>

## Direct properties — global_vn / 221012322112 / 3

<a id="canonical-3011232333201310-2011000332322031-0311021021332201-1112100231232013-3320311012011123-3111213211331102-3021133323120203-0323200320112232"></a>

<a id="canonical-1011010111213100-1320033023331322-1131332303220132-0303001103032200-2220000223211022-1003202221310210-3323223120101301-0013001010010021"></a>

## name property — global_vn / 221012322112 / 4

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

<a id="canonical-2210230300103032-1101320312001320-1003100030021101-2200031233222130-1130223022202111-0220101232311203-2002222133231101-2320330112232100"></a>

<a id="canonical-2110121313211011-3000110301021331-0331300103100100-0032322333323013-1002321010330232-1123230011210212-2202103130031111-0220121022312310"></a>

## namespace property — global_vn / 221012322112 / 5

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

<a id="canonical-2121101022313000-0301101103131202-2312002002202032-1023202230312002-3300200033011312-1002320020112300-1232120212223101-3003000000301223"></a>

<a id="canonical-0103230201332222-1023223030301233-2210111033110120-2210001112012101-0000122301130303-2220202123011233-0013010020201223-2122003122232233"></a>

## tenant property — global_vn / 221012322112 / 6

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

<a id="canonical-2323212232300132-0213113103223120-3122232011202201-1210131032231210-1123330003223003-2322132032033011-0323323311203112-2332210010221221"></a>

## Next pages — global_vn / 221012322112 / 7

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-0323121202121301-0213113313230200-1101131332123201-2103133232111000-0010031133203200-0003011222212302-2131033330110012-3022003311203003)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3103003200122121-0220123331322220-1122001133011002-0102133010322110-0102300310331323-3330010211130200-1122223212011103-3133221130012103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301330121132031-0320102210213232-1301212002233031-3022311330211010-0203022020130212-2322310232211321-1022213330213303-1230112223311111"></a>

## voltstack_cluster.k8s_cluster — k8s_cluster / 231302223221 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.k8s_cluster

<a id="canonical-1102003211031000-1100232131233232-2312023032111321-2213033021121232-3010220013300303-1301221130223321-0121130121110133-3312001100013022"></a>

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
k8s_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032101321103111-0101321001031332-2333002131122221-3030321320232113-1320333231333203-3332200331022330-3110112113123331-0122130012233122"></a>

## Direct properties — k8s_cluster / 231302223221 / 3

<a id="canonical-2010312302231231-3333233223213103-1113001001303230-1103322232001220-1221301100112033-1133312113200002-2221023132332121-0110303103310030"></a>

<a id="canonical-3022223112302321-2022002202102123-0312310010012100-1121223033132222-1131310233031210-3223022000303301-0211133203022112-2122323303033301"></a>

## name property — k8s_cluster / 231302223221 / 4

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

<a id="canonical-0300102003011102-1021031323220201-1100300203321032-0211211123000110-0130120212021011-0011322110223311-0013122100003000-2121112320322200"></a>

<a id="canonical-2312121310101231-0202122113323002-1200023123313032-3223210213132020-3111312312322303-3233123132233330-3021212312311301-0310332322012032"></a>

## namespace property — k8s_cluster / 231302223221 / 5

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

<a id="canonical-2330201312300312-1031110303220120-3100233322232130-1120033333023303-1331311213213112-3302200111221221-1321123223232013-0130121013210223"></a>

<a id="canonical-2313232130312122-0132202212201233-1021200322131123-0220033212010013-1010122132123033-3211323022311332-0203322131312112-2222103300112222"></a>

## tenant property — k8s_cluster / 231302223221 / 6

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

<a id="canonical-1300203030011110-0121121301122110-0020323123203000-2222233131120103-1231303330003003-3221033312310033-2032010232130010-2211123000131033"></a>

## Next pages — k8s_cluster / 231302223221 / 7

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0220213211122001-2221303302321103-0210023130302010-3120103321133022-1321013323003003-0320032112022220-2322013121221303-2231222322202120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023222113221111-1311111201011220-2011001332000200-2302323031013012-2033223300233122-0031332200310310-0103132133020112-2033322131300021"></a>

## voltstack_cluster.no_dc_cluster_group — no_dc_cluster_group / 012212233033 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.no_dc_cluster_group

<a id="canonical-0001302120030202-1303121321201231-0213322110330300-2103330110210222-2000012101323212-0221110310331213-3301023313131133-0003113121333031"></a>

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
no_dc_cluster_group = {}
```

<a id="canonical-1301130030031300-0102332211131302-0200231210312303-3201201210333103-1013200013332222-3203211101320130-3130013131322311-0013330331222321"></a>

## Direct properties — no_dc_cluster_group / 012212233033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130330312333012-0132222101001123-2100032113300132-3330320023202301-0010212310220230-2311321023011021-2331001231332100-2213213310333131"></a>

## Next pages — no_dc_cluster_group / 012212233033 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3122002221331312-0123201123230102-2100132111122313-0330213132001330-0001033102220202-0122323120230320-3223022121321102-1333201202112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222032012010030-2222110023122212-0103201222333003-2330113323230121-3113331033032010-0302102111223023-2213033222231110-0001311030111013"></a>

## voltstack_cluster.no_forward_proxy — no_forward_proxy / 022100032133 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.no_forward_proxy

<a id="canonical-1033311300222123-3300021113011110-3311131202233032-1332222301232331-2321031122331023-1221003202102033-3020001122103201-3013320010210020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no forward proxy.

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
no_forward_proxy = {}
```

<a id="canonical-2230303001322011-1113013200100310-2321011313321110-2131022232003223-0022121011320211-3120221211021001-2113132120121101-2221130232223102"></a>

## Direct properties — no_forward_proxy / 022100032133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0300233221322100-2023013212223030-3101001201211330-3000202221220320-0010001330102131-2300032032103112-2320110012010133-2233313130231332"></a>

## Next pages — no_forward_proxy / 022100032133 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2220030233023310-2103121010213212-3122122202213103-0123211113223031-3200320010331321-3000213303021020-1123110232302223-3012132223022030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011300313130021-3223111223002102-0200331131022320-1113022331120312-1032032232313031-3211211121122021-1302331121020300-2011232312120321"></a>

## voltstack_cluster.no_global_network — no_global_network / 202301113210 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.no_global_network

<a id="canonical-0113120300321231-1002220333000103-3020001112212302-2012223030011032-1121222012031131-1330122011122313-2333200302301322-1223333110232223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no global network.

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
no_global_network = {}
```

<a id="canonical-0313302011022031-3333031201133113-2222022310021013-0232030021131113-2320000321300333-3121010302131210-3032210202013200-1101230233112222"></a>

## Direct properties — no_global_network / 202301113210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120310232132213-0001222223211032-1232002211110233-1212233333203100-0112131322101311-2232031211102013-3033311033020311-2012110332231200"></a>

## Next pages — no_global_network / 202301113210 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3232233010211212-1013020303022030-1112121330121331-0230002302323311-0133033122223313-2232121032230133-3113332123302103-2313233130131231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030003022132310-3023211100022132-3131011202031022-0131122101030122-1002312220031000-0032231223013023-0011031020013010-3030003002322100"></a>

## voltstack_cluster.no_k8s_cluster — no_k8s_cluster / 000133230021 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.no_k8s_cluster

<a id="canonical-1030010131002221-1022322200201210-3022232321010201-3023023233212233-3102103200110113-2031322130103332-0322113212200121-1032203120003310"></a>

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
no_k8s_cluster = {}
```

<a id="canonical-1112113232111033-2320031000102200-0010213003122031-3101020121102030-2303000203221310-0213000013011020-1203203322011123-3103213102331101"></a>

## Direct properties — no_k8s_cluster / 000133230021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022202332323020-3222301233003233-3231010112202013-1030330133130300-1031110123120231-3232311032310233-0021122130013120-1011131222311113"></a>

## Next pages — no_k8s_cluster / 000133230021 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0013010110333221-3201233330000210-2333131132212103-2012102221331033-1121101101212312-3200221103100000-2313232003323000-1132220230320121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320121111323131-2000200301013201-2012003103113113-3103020331220101-0001210320022133-1301321131120313-2202212232312331-2121102013030201"></a>

## voltstack_cluster.no_network_policy — no_network_policy / 113001310010 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.no_network_policy

<a id="canonical-0101223120121120-0133100312302202-2300030100223102-0020131101322012-1211312101003011-3211110200101202-0300303130221112-2210112010103311"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_network_policy = {}
```

<a id="canonical-0101020133332021-1303131301012111-2310012110033103-3113212302011323-2000131232321010-3131023210003303-2301202121202213-2032120202002013"></a>

## Direct properties — no_network_policy / 113001310010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123122212113121-2321020210233101-1121113222300322-0023101230211213-0312132211202220-3033121323121123-2330210010333103-3100312322210203"></a>

## Next pages — no_network_policy / 113001310010 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1011203111010233-2230200221322310-2000113233223010-2000001222101230-3222333020032210-2220200011200320-3202012101130301-0102130013202210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203300102002200-3010332101201101-3000202323313231-3210302201333323-3330300210233232-3303011323212021-0002121323323302-0203232023022213"></a>

## voltstack_cluster.no_outside_static_routes — no_outside_static_routes / 102211002300 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.no_outside_static_routes

<a id="canonical-2111320001230210-1112002203232013-1300110211103202-3110200303321020-0223310330211000-2011210212030201-1301331213112200-0203313103011032"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no outside static routes.

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
no_outside_static_routes = {}
```

<a id="canonical-2223201231013122-0300323331230033-3010222023030031-0132000033300221-1233231123111111-0110003131332310-3111001013113321-2213232320131310"></a>

## Direct properties — no_outside_static_routes / 102211002300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032110222122333-0202103201023010-3030301302212303-3122133302210111-1002303000113010-2231220302332313-3101200002123023-1232201132313133"></a>

## Next pages — no_outside_static_routes / 102211002300 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123121200112332-3002113201311100-2100320011132201-0313013032033233-1303030220310122-1313213120200113-2010311233033121-3021000101122132"></a>

## voltstack_cluster.outside_static_routes — outside_static_routes / 133121003030 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.outside_static_routes

<a id="canonical-0000112220223011-0211331031021311-2100220103221103-0102021230223312-2130323302323122-3031121311111220-2202210102302112-2300032031133333"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
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
outside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333102023310001-2203020020200132-2223231313101000-3210201212001210-1021120312130003-0033310121310002-3300303232300330-2112113001230200"></a>

## Direct properties — outside_static_routes / 133121003030 / 3

- [static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212): complete subsection reference.

<a id="canonical-1122123212012000-2320123002212010-2322313330123100-2101120310213132-1100022001223011-2001032321301220-2210031302311222-0200100311002101"></a>

## Next pages — outside_static_routes / 133121003030 / 4

- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222330303302220-2123003232221000-2302201333013211-3331322010213132-0221211011312023-1130030300000021-1131203212313032-0022230122033313"></a>

## voltstack_cluster.outside_static_routes.static_route_list — static_route_list / 202100122203 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- voltstack_cluster.outside_static_routes.static_route_list

<a id="canonical-2031233033323212-3313122021033100-2132223211033033-1231310100303311-3033330103233301-0210123110330011-0003301123003110-3233331330002002"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223302212203233-1231133321021133-2203022110201022-0110102203133132-3201210110013233-1000222301022202-3030330203321130-3333313203020112"></a>

## Direct properties — static_route_list / 202100122203 / 3

- [custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213): complete subsection reference.

<a id="canonical-1300002022021133-0023211120301212-2112111000300311-3320120232032003-2021130323020021-3322030133001021-0220301203311312-0220232010203222"></a>

<a id="canonical-2033300013011011-0322213300102203-3320002111221121-1033103221331200-2201320303032132-3130033030132301-2312313222333221-3212120122201020"></a>

## simple_static_route property — static_route_list / 202100122203 / 4

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-3002230332011232-1103120121002323-1323310100202112-0001302010212211-3230101113231123-1031021130023000-3121322000123323-0330321232312002"></a>

## Next pages — static_route_list / 202100122203 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333132121310010-2110010033102213-3011021012113022-2320021012212001-1303212032210021-2031313331311333-1320211012112033-0301032233201112"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route — custom_static_route / 332021311002 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-2123030121331123-0222232012301032-0331113013210320-1130331002321211-1031323012202001-2220320323330210-2302300012010231-1031303230222310"></a>

Type: `"object"`. single nested block, Optional.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnets")}
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
custom_static_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133033010021003-2110130331331222-3103010223233001-1300313100121013-2211103213011333-3330233322331010-0013203222221210-3000120232200123"></a>

## Direct properties — custom_static_route / 332021311002 / 3

<a id="canonical-2233223213031002-2020031120211120-3333102112211002-2210130201103310-3303320113120321-1210022131011201-2221131020221233-2111331112131220"></a>

<a id="canonical-0022200120322221-1132123331003030-3021012303211121-0233330233113122-1111133113012212-1020103211323202-1231132100212212-2332122011012200"></a>

## attrs property — custom_static_route / 332021311002 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](resources--azure_vnet_site--reference--group-009.md#canonical-3100000331200121-1223113310312211-2011023031011030-0212131013310012-2200021113032111-1323021003020013-2102200331003121-2010321333010021): complete subsection reference.

- [nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-3031013011323111-2122323303113302-1031202023131133-2313222233112331-2302112133323203-0023112330311322-0123003201311121-3113201331313100): complete subsection reference.

- [subnets](resources--azure_vnet_site--reference--group-009.md#canonical-2312223302201322-2110312130320113-2022022322110311-2000013313031011-0010011323200111-2230222201310013-2112111302103011-3200110310323222): complete subsection reference.

<a id="canonical-1122012211033223-0033002113300121-0133021302122221-0220132110101131-0310123021113003-2032223311022302-0302022233221120-2201311113122023"></a>

## Next pages — custom_static_route / 332021311002 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](resources--azure_vnet_site--reference--group-009.md#canonical-3100000331200121-1223113310312211-2011023031011030-0212131013310012-2200021113032111-1323021003020013-2102200331003121-2010321333010021)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-3031013011323111-2122323303113302-1031202023131133-2313222233112331-2302112133323203-0023112330311322-0123003201311121-3113201331313100)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-009.md#canonical-2312223302201322-2110312130320113-2022022322110311-2000013313031011-0010011323200111-2230222201310013-2112111302103011-3200110310323222)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3100000331200121-1223113310312211-2011023031011030-0212131013310012-2200021113032111-1323021003020013-2102200331003121-2010321333010021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122120303111230-0230130200013112-1100311201103013-0313323003313320-0311333301301233-2101033301033213-0222012220321223-1210312131133323"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels — labels / 012030331223 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-1013210323201311-0122023320321012-3033323111332212-0233111032200031-3322332021023212-2133033032302012-1123230301021331-3033010311332120"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Static Route, these labels can be used in network policy.

Receipt-pinned upstream constraints:

```json
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
labels {}
```

<a id="canonical-2231310231123300-1032300233320223-3112021112300232-3111223031000101-0323021123013331-1200032221232133-3033122100100223-1221203313333112"></a>

## Direct properties — labels / 012030331223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233321312310003-1220021221332303-2113302301102001-2300303223320000-1110002211323213-2211331102102202-1330321322220210-0233033030030002"></a>

## Next pages — labels / 012030331223 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3031013011323111-2122323303113302-1031202023131133-2313222233112331-2302112133323203-0023112330311322-0123003201311121-3113201331313100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322131231303220-2032200111112113-1213332123021222-0111302121233131-0303332031310311-2012332021221132-1102303012011223-3302322111130030"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 102123311200 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-1021133203011213-2220021303121010-2131321011313102-2201300132111122-1322232203311323-3213110002213021-3123211132100230-1130021122012200"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
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
nexthop {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320130231233210-1203231330203213-2211031311120012-2132031031032320-3032202302120331-0220121231133130-2113321232321221-1010023100301112"></a>

## Direct properties — nexthop / 102123311200 / 3

- [interface](resources--azure_vnet_site--reference--group-009.md#canonical-1031133210110232-3000201311223221-3012012002023013-0100010300320322-0021310322123323-2032220312332121-2000120231230212-0032021122303211): complete subsection reference.

- [nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-3110000001012121-1102233111212033-2203320120100213-1222302113223111-2102022220312212-2023100320120101-3221233223322313-0100231302011020): complete subsection reference.

<a id="canonical-0103302332322303-0331011211130303-0310102121110220-2122322211121001-3231321131031210-3200203300210323-1301022331131013-0232221303022212"></a>

<a id="canonical-0311200332022310-1321213323030220-1230223013133031-2000233301031331-0313011333311133-1233020112111002-3222000111000321-1121313201020331"></a>

## type property — nexthop / 102123311200 / 4

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1330310213023213-0122101310103020-3312303133301001-3213321332103022-3020233010312200-1130130102220130-2001100232021312-0312322332131120"></a>

## Next pages — nexthop / 102123311200 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--azure_vnet_site--reference--group-009.md#canonical-1031133210110232-3000201311223221-3012012002023013-0100010300320322-0021310322123323-2032220312332121-2000120231230212-0032021122303211)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-3110000001012121-1102233111212033-2203320120100213-1222302113223111-2102022220312212-2023100320120101-3221233223322313-0100231302011020)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1031133210110232-3000201311223221-3012012002023013-0100010300320322-0021310322123323-2032220312332121-2000120231230212-0032021122303211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330130223013322-1030013032020132-1013211012303223-3112002300111013-3113113013322001-0212022111030133-0120332103001230-3232210300302233"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 130232313010 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-3031013011323111-2122323303113302-1031202023131133-2313222233112331-2302112133323203-0023112330311322-0123003201311121-3113201331313100)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-2121000212212222-2132130103110202-0311331100330013-2223031112311223-0103012302103320-0001003211112032-3223133212010331-0221121312013001"></a>

Type: `"object"`. list nested block, Optional.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1221313202221001-2202010212311022-0102123021133130-0121213213033312-0011202133033230-3320123132312321-0321100213222101-0023102302012310"></a>

## Direct properties — interface / 130232313010 / 3

<a id="canonical-0031122121132320-3323210131032011-1031012023202312-2310012021031023-1022300210211033-1322211311022011-1323311021231302-2030202331233311"></a>

<a id="canonical-0012231300320312-3103233102213202-3311313022302010-2133211033011113-2221010302212020-0230013311221232-2002010303301022-0333002220312201"></a>

## kind property — interface / 130232313010 / 4

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

<a id="canonical-1220130310202123-0323331003321201-2031313203020130-1030132320203131-3132202231002200-2122313303202323-3323211103030321-3333302220200333"></a>

<a id="canonical-0202003210030000-2332220212230023-2021311230001012-2313031300330232-3001310031002103-0020113100010102-3202320231211332-3110331122111100"></a>

## name property — interface / 130232313010 / 5

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

<a id="canonical-2133330320210120-3101001100320223-2033130332020202-0221131320230332-1333231122301021-1011000330311033-1311311123310033-0032113131103032"></a>

<a id="canonical-0002313320331330-3001200313202132-1201013322111212-3300103011311222-3003332210123320-1321233100120201-1010300333123212-1212132220200213"></a>

## namespace property — interface / 130232313010 / 6

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
  }
}
```

<a id="canonical-2210031102120033-0202011121233230-0331212302031003-0300103201010203-2001022302200113-3333021130223311-1100330032120300-3323221112022111"></a>

<a id="canonical-3313100300013133-2011022303200103-2002231220012001-1230322000303000-0013333331332001-2003033010223132-3312210320131103-1210203110003300"></a>

## tenant property — interface / 130232313010 / 7

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

<a id="canonical-1301133323131210-3023131010221211-3102031112232321-1033011103113032-2211212321132022-2212233113220301-2113231310201021-1211212101032300"></a>

<a id="canonical-0312303010023232-0010201000201101-1322112332101022-2320312213232123-0113220133321321-1012323311220312-3122123113121212-1130003033032230"></a>

## uid property — interface / 130232313010 / 8

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

<a id="canonical-2033210210231302-0110102322313230-0103321101333312-3123322300311110-1123133203222322-0212202232302012-0100301030113320-3320213212201310"></a>

## Next pages — interface / 130232313010 / 9

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-3031013011323111-2122323303113302-1031202023131133-2313222233112331-2302112133323203-0023112330311322-0123003201311121-3113201331313100)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3110000001012121-1102233111212033-2203320120100213-1222302113223111-2102022220312212-2023100320120101-3221233223322313-0100231302011020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133202333203102-2310002233211101-1113321311102013-3200000321331313-2331113320302101-2000020200310232-2023031331120103-1020311130022302"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 120203003132 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-3031013011323111-2122323303113302-1031202023131133-2313222233112331-2302112133323203-0023112330311322-0123003201311121-3113201331313100)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-0012310102023110-2331133303122021-1312321231322333-1303113033312213-3123313300123323-0022131221031321-0033122200000030-2233111332301022"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023100130332313-0330303131303110-3320231131320220-3130000220012130-1331301103031011-1313032302212300-1122200100112002-1121220003232003"></a>

## Direct properties — nexthop_address / 120203003132 / 3

- [dual_stack](resources--azure_vnet_site--reference--group-009.md#canonical-1221332102302312-1033131321012122-0102110302223311-1100011110213212-2032301001001300-1222132012102120-1320013302330322-2301320202030312): complete subsection reference.

- [ipv4](resources--azure_vnet_site--reference--group-009.md#canonical-3201010321230230-2301333130030203-1111022211102010-2301221302003102-2301213232201222-0213023223120312-0130313310102211-1112310300020303): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-009.md#canonical-1020013033220121-3122013320032102-1112211022012101-1210111323211131-2302131023223201-1211221121330220-1310333013212301-3033123330211210): complete subsection reference.

<a id="canonical-1031101321113021-1223021332320301-2323112211301121-2323222130002030-0222022333311331-2102111030233201-0020321232303222-2333210330103101"></a>

## Next pages — nexthop_address / 120203003132 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-009.md#canonical-1221332102302312-1033131321012122-0102110302223311-1100011110213212-2032301001001300-1222132012102120-1320013302330322-2301320202030312)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--azure_vnet_site--reference--group-009.md#canonical-3201010321230230-2301333130030203-1111022211102010-2301221302003102-2301213232201222-0213023223120312-0130313310102211-1112310300020303)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--azure_vnet_site--reference--group-009.md#canonical-1020013033220121-3122013320032102-1112211022012101-1210111323211131-2302131023223201-1211221121330220-1310333013212301-3033123330211210)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-3031013011323111-2122323303113302-1031202023131133-2313222233112331-2302112133323203-0023112330311322-0123003201311121-3113201331313100)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1221332102302312-1033131321012122-0102110302223311-1100011110213212-2032301001001300-1222132012102120-1320013302330322-2301320202030312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311023022200223-1012303130111012-3220022322330131-3223021133211201-1203113220133103-2010312203122112-1003310312210102-0113231111322012"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 203313132113 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-3031013011323111-2122323303113302-1031202023131133-2313222233112331-2302112133323203-0023112330311322-0123003201311121-3113201331313100)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-3110000001012121-1102233111212033-2203320120100213-1222302113223111-2102022220312212-2023100320120101-3221233223322313-0100231302011020)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-3312121033032220-2132121121022201-1331232030023121-2321312033210022-0013312000013001-3001132323233102-2021111311120131-0202031331013131"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
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
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312301133123201-3120203231310322-2311021331121003-2201112322221132-1311211203022010-2113230100212201-2013321112211332-1002121021022211"></a>

## Direct properties — dual_stack / 203313132113 / 3

- [ipv4](resources--azure_vnet_site--reference--group-009.md#canonical-1313130220003022-3101313112122110-2130303032123312-3332233210333003-0323122030003301-1100103013303302-3102021101333020-0312311312211013): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-009.md#canonical-1230213221310110-3023113322233310-3303130033210100-2002132003020101-1130333021332110-0200020301233102-2231313111231310-1300213302320000): complete subsection reference.

<a id="canonical-1210333122212011-2022322101003200-3020023123101312-2031031231332021-1213030033112233-2020002313302030-1000231001031323-3310013202301313"></a>

## Next pages — dual_stack / 203313132113 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--azure_vnet_site--reference--group-009.md#canonical-1313130220003022-3101313112122110-2130303032123312-3332233210333003-0323122030003301-1100103013303302-3102021101333020-0312311312211013)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--azure_vnet_site--reference--group-009.md#canonical-1230213221310110-3023113322233310-3303130033210100-2002132003020101-1130333021332110-0200020301233102-2231313111231310-1300213302320000)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-3110000001012121-1102233111212033-2203320120100213-1222302113223111-2102022220312212-2023100320120101-3221233223322313-0100231302011020)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1313130220003022-3101313112122110-2130303032123312-3332233210333003-0323122030003301-1100103013303302-3102021101333020-0312311312211013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112111130120303-1320201121232210-1310302312111210-0100120101231223-1103322303122100-0030311100011032-1132303120111033-3212303120302002"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 120211233120 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-3031013011323111-2122323303113302-1031202023131133-2313222233112331-2302112133323203-0023112330311322-0123003201311121-3113201331313100)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-3110000001012121-1102233111212033-2203320120100213-1222302113223111-2102022220312212-2023100320120101-3221233223322313-0100231302011020)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-009.md#canonical-1221332102302312-1033131321012122-0102110302223311-1100011110213212-2032301001001300-1222132012102120-1320013302330322-2301320202030312)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-2321002212222110-1221002203022200-3201200112202331-1022321303202311-2121000000210323-3132202302213230-0013333132103020-0103231033232233"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-0101132213121132-2230231112012201-3203330312120312-3301213120020320-1200010013300323-1021021313232010-2310011231000123-2300013203010102"></a>

## Direct properties — IPv4 / 120211233120 / 3

<a id="canonical-1013003002333313-0130303210132201-3220200200330310-1031030331121133-1312023120330003-3231222032222221-2011113002310123-0210113231133311"></a>

<a id="canonical-0321020332322110-0022103122322130-2302230311202101-1322322321310001-1320301132033112-3032222113330131-3333210313230133-2230231000122233"></a>

## addr property — IPv4 / 120211233120 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2032230021022220-1312021323311202-1013233301130020-2132020211131101-3302222322122120-2202311121202112-3313301022223132-2330310012221131"></a>

## Next pages — IPv4 / 120211233120 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-009.md#canonical-1221332102302312-1033131321012122-0102110302223311-1100011110213212-2032301001001300-1222132012102120-1320013302330322-2301320202030312)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1230213221310110-3023113322233310-3303130033210100-2002132003020101-1130333021332110-0200020301233102-2231313111231310-1300213302320000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201100012233330-3231201121132330-2031110102132032-2132102110203233-2300330130303203-2110310002001013-1003302003301223-2313302120312001"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 123320320110 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-3031013011323111-2122323303113302-1031202023131133-2313222233112331-2302112133323203-0023112330311322-0123003201311121-3113201331313100)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-3110000001012121-1102233111212033-2203320120100213-1222302113223111-2102022220312212-2023100320120101-3221233223322313-0100231302011020)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-009.md#canonical-1221332102302312-1033131321012122-0102110302223311-1100011110213212-2032301001001300-1222132012102120-1320013302330322-2301320202030312)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-3332100121131120-3132001111010020-0212232010113211-0303131300211131-2122111002000033-2113302021110223-0331333313023330-2332223130333010"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212232133222211-3310110311223210-0231223122231211-2001033202211223-3211231231102320-0230032121133101-2011111111300001-2131123033030100"></a>

## Direct properties — IPv6 / 123320320110 / 3

<a id="canonical-2012001011112230-1211211133132120-1232120220213300-1230030113113210-3332022230003002-2013231303203102-2133333203231021-0113132332203101"></a>

<a id="canonical-1021013011013031-1130300302022120-3321021101011022-0101103011223120-2132121333020231-1201310133110302-3222013121302231-0121131203103001"></a>

## addr property — IPv6 / 123320320110 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-1020001131002333-1232220202020230-2211332033122320-2013202232211322-1112200223321102-3001112013133010-1020132113231201-1323130222110121"></a>

## Next pages — IPv6 / 123320320110 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-009.md#canonical-1221332102302312-1033131321012122-0102110302223311-1100011110213212-2032301001001300-1222132012102120-1320013302330322-2301320202030312)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3201010321230230-2301333130030203-1111022211102010-2301221302003102-2301213232201222-0213023223120312-0130313310102211-1112310300020303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122102131131301-0102031101320223-3022313113002210-3122101212123321-0323132021112110-2202000032000302-2200122332123211-2223013212231011"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 033031111330 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-3031013011323111-2122323303113302-1031202023131133-2313222233112331-2302112133323203-0023112330311322-0123003201311121-3113201331313100)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-3110000001012121-1102233111212033-2203320120100213-1222302113223111-2102022220312212-2023100320120101-3221233223322313-0100231302011020)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-1213133323103021-3023100001323221-3203201003120311-2120202012011113-3103330222121320-2233313030231322-3220221012330033-1213213013023323"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-1301131103123231-0303000312110100-2023113113303032-0131102331321002-2300113023110001-0300113102231131-2200132023222022-0000230223130231"></a>

## Direct properties — IPv4 / 033031111330 / 3

<a id="canonical-1102103121000310-2210111303133123-1311213323320331-3121110300003200-2123003121230022-3332222010311332-0300110101132211-0022311331113003"></a>

<a id="canonical-2210200221113132-2301132230131211-2333020121320010-0120320311121033-2123303301022102-2003120030033013-1210321110210000-0002120220211023"></a>

## addr property — IPv4 / 033031111330 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1021021200001330-3302121231123003-2200010331130310-1030320113131000-3333122310333032-3212100132200322-0101203032232211-2003020011231313"></a>

## Next pages — IPv4 / 033031111330 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-3110000001012121-1102233111212033-2203320120100213-1222302113223111-2102022220312212-2023100320120101-3221233223322313-0100231302011020)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1020013033220121-3122013320032102-1112211022012101-1210111323211131-2302131023223201-1211221121330220-1310333013212301-3033123330211210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300320101210313-3020211031311031-1133303011000012-3012010302033002-0133301001132111-1122302131300203-3331020323231320-1322210002311132"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 201013112332 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-3031013011323111-2122323303113302-1031202023131133-2313222233112331-2302112133323203-0023112330311322-0123003201311121-3113201331313100)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-3110000001012121-1102233111212033-2203320120100213-1222302113223111-2102022220312212-2023100320120101-3221233223322313-0100231302011020)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-2321311112210023-3221322233211211-3131300010012023-0330303113033222-0312013122213310-1222012300100333-1031113332011333-3122312002103330"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2023032301033131-3202132332312320-3101312012301013-3210103032322301-0003010103132031-1230033221031213-3100233321100002-3232133111033102"></a>

## Direct properties — IPv6 / 201013112332 / 3

<a id="canonical-1120321212301313-1220302311032032-1200303103121223-2103330031031020-0233203121331010-3032013101112320-3102313112033332-0132013211202313"></a>

<a id="canonical-1222122100333100-3031230120322012-2031110203313331-3102331311300211-3130030100100122-3320332200121331-2131133301013202-3033333130012211"></a>

## addr property — IPv6 / 201013112332 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-2000110303332002-2133100100000002-0031030110200122-3313331303201133-0001300102121232-1122330131222300-1103322333111202-3331002201220021"></a>

## Next pages — IPv6 / 201013112332 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-3110000001012121-1102233111212033-2203320120100213-1222302113223111-2102022220312212-2023100320120101-3221233223322313-0100231302011020)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2312223302201322-2110312130320113-2022022322110311-2000013313031011-0010011323200111-2230222201310013-2112111302103011-3200110310323222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333223032233333-1201122132300033-2301133112113000-3313330022311003-3212112210222220-1330322010211010-1101201203230232-1122113220003112"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets — subnets / 003202120223 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-3332100300110302-2333302031101112-0220230310113202-0231113220301323-1112020211312111-2123132031302321-1002301331200023-0001001321112301"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211030000230103-1302010232330310-1000113320120112-0231000301010003-3011301002132031-2003201003023121-0313000133301300-3323001231333300"></a>

## Direct properties — subnets / 003202120223 / 3

- [ipv4](resources--azure_vnet_site--reference--group-009.md#canonical-0231210321023033-2001023323011111-1001330220012302-1121131002131111-2323123200023101-1201232020021302-0303231233322201-2011021322013302): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-009.md#canonical-0020202333032120-3222223303331102-2032120001203033-3212322021221113-2230231331103112-2130330322202110-3300132232331211-0300200203023200): complete subsection reference.

<a id="canonical-2330302200202213-1102001222112300-1133313120011033-2000310302312123-3030113130121133-2311000023332200-2312312313020022-0222102000012221"></a>

## Next pages — subnets / 003202120223 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--azure_vnet_site--reference--group-009.md#canonical-0231210321023033-2001023323011111-1001330220012302-1121131002131111-2323123200023101-1201232020021302-0303231233322201-2011021322013302)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--azure_vnet_site--reference--group-009.md#canonical-0020202333032120-3222223303331102-2032120001203033-3212322021221113-2230231331103112-2130330322202110-3300132232331211-0300200203023200)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0231210321023033-2001023323011111-1001330220012302-1121131002131111-2323123200023101-1201232020021302-0303231233322201-2011021322013302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100022301123133-1331201202000202-0323012200023211-1120202021110121-1002102312213030-0130003302332312-3203102100312231-2013103122122212"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 301213302021 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-009.md#canonical-2312223302201322-2110312130320113-2022022322110311-2000013313031011-0010011323200111-2230222201310013-2112111302103011-3200110310323222)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-0131212013312020-3020311301212223-0222322102332211-0330132323201302-2130302330330131-1011323030320222-2111131321031313-0113102000310012"></a>

Type: `"object"`. single nested block, Optional.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2322002020133012-1113223331010210-2120210000030331-1232313000113322-2332012112103233-3212030203301200-0003012211301212-1223222013031122"></a>

## Direct properties — IPv4 / 301213302021 / 3

<a id="canonical-1313331310311320-0300012311212232-3022001123002220-1021120330212123-1302323033311130-2221231020203302-2012312233001310-3210130230332210"></a>

<a id="canonical-0120021000030102-0023023120213310-2220101320330122-0200301213202222-3123031121220302-0310302102221232-1313012011122232-3133001222022322"></a>

## plen property — IPv4 / 301213302021 / 4

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32),
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
    }
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

<a id="canonical-0213021001010233-0332032210130133-1331222303113312-2200023031222121-0003101221221231-2321111122103310-3121113312131330-3231210200223320"></a>

<a id="canonical-2231012102023302-3313132101312331-2300300031331311-0231111223311113-0330221231010330-3130210132002222-1200211010333223-0233311033321001"></a>

## prefix property — IPv4 / 301213302021 / 5

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1311222231303231-1132203200232121-2110121013132231-0121222123122213-0323031021302203-0013003123222120-3031020301201112-2332111221100333"></a>

## Next pages — IPv4 / 301213302021 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-009.md#canonical-2312223302201322-2110312130320113-2022022322110311-2000013313031011-0010011323200111-2230222201310013-2112111302103011-3200110310323222)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0020202333032120-3222223303331102-2032120001203033-3212322021221113-2230231331103112-2130330322202110-3300132232331211-0300200203023200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330103130212120-3310310001333001-3312320132331233-0131003200311221-3113113323223303-2202003112102100-1322101033322110-3021102031121202"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 232330331021 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-0100231323113011-3022021102111220-2021101030010220-2000022103132232-0310320321133320-2021103332300231-3132123131102013-0121113233022222)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-3301312032033021-3012112130321120-1201031210100321-3223103233230211-3111003133330133-2331322310000123-1003023001133313-3330001130320212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-3331010032220130-2123112313013211-0022220010303033-1223131111312203-3003333133133231-3120022223231120-3103131310321231-3203010223331213)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-009.md#canonical-2312223302201322-2110312130320113-2022022322110311-2000013313031011-0010011323200111-2230222201310013-2112111302103011-3200110310323222)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-1312022223320333-3112022203033203-2030002211001033-1022030022212101-1310133121233121-2030322210031211-2300021331200122-2230113232021300"></a>

Type: `"object"`. single nested block, Optional.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323003210322321-1331233002133312-3120133022013200-3212023331121201-3311010213312222-0331111320330203-0012023320103210-3103321321332222"></a>

## Direct properties — IPv6 / 232330331021 / 3

<a id="canonical-2101131003022010-1200123213130101-3020230113122130-3312011000123012-0323302213203330-0303101113230213-3031013200021300-1203330322211302"></a>

<a id="canonical-3023210102300221-0101300001012103-1001103211111023-1211231221323213-2010301221231332-3220130210131312-1001231223223002-1020213132111310"></a>

## plen property — IPv6 / 232330331021 / 4

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-0322233013322113-1233031210033323-3111333233220023-2021003211332113-1130220201323230-0201230210300033-1301323031331012-1230101332130001"></a>

<a id="canonical-0211321032330010-2203113320110332-0231231013013313-1202303302001101-0131223220330001-3021223223331320-0232113010302022-0222130122100031"></a>

## prefix property — IPv6 / 232330331021 / 5

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3100303021103211-3031120331132232-3012213320122003-3132122022230110-1022130033023201-1101020203012131-3202022123200112-2301033331120003"></a>

## Next pages — IPv6 / 232330331021 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-009.md#canonical-2312223302201322-2110312130320113-2022022322110311-2000013313031011-0010011323200111-2230222201310013-2112111302103011-3200110310323222)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2321133213021020-0212121322131221-2110132012310211-2131222311000202-1112033321010300-2002022321211122-2020021222122012-1001212101003111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332233233230330-3223320210131111-1200200021113321-1230310100212222-0201002120202111-1100023331030200-1310122210210322-0213101311211032"></a>

## voltstack_cluster.sm_connection_public_ip — sm_connection_public_ip / 021312030101 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.sm_connection_public_ip

<a id="canonical-0032121013010031-1211313102120130-2320020312032220-0322000012110133-3302233223312212-1123131301201322-0220321103112013-1133012301102003"></a>

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
sm_connection_public_ip = {}
```

<a id="canonical-2312103110320211-0300301120321020-3231200321222030-1321030022202210-0023132230023132-2131330223223233-3313032121030311-0122112033332311"></a>

## Direct properties — sm_connection_public_ip / 021312030101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101221212113020-2111123230123113-1220023033333123-3033121121000203-0133133120332203-0203210221001331-3122030333203112-1103110111201332"></a>

## Next pages — sm_connection_public_ip / 021312030101 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3013311330102123-3233032312200212-3200203223111212-2332321312031020-2101122223222222-1022200220220111-3123111020023213-2311323300213221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032332221023220-1113023212212203-0031031330132301-0320200033212121-0001321320301310-3102311011110220-3031202313133102-2202321103113303"></a>

## voltstack_cluster.sm_connection_pvt_ip — sm_connection_pvt_ip / 301121210022 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.sm_connection_pvt_ip

<a id="canonical-1321310333022132-2001321031312312-1113320133331111-2113122331122202-1302313332330323-2222113221311012-0230010220321012-3020210132303223"></a>

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
sm_connection_pvt_ip = {}
```

<a id="canonical-2130212230121131-0202020323100021-3302302303200102-2213330233230120-3202323121203322-0122122111231031-3121020121012332-1331100303330222"></a>

## Direct properties — sm_connection_pvt_ip / 301121210022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231220030203321-3021233103200112-3032300332321122-3330121221020121-3200103112023203-1131213123213220-3121312023230131-1131221231311113"></a>

## Next pages — sm_connection_pvt_ip / 301121210022 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3311322312000000-0231210321003323-2300131023112023-2003010121321010-0331001121323123-1103330000301123-0311230333000223-0302132310200030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211323322302202-1311230113213333-3101222223222032-0212210112002231-0132231002301331-1032230030311121-0202011003232021-1203023210301101"></a>

## voltstack_cluster.storage_class_list — storage_class_list / 201322003131 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- voltstack_cluster.storage_class_list

<a id="canonical-3223232322301302-2113313112213232-2120232120133023-1132200230012131-0012331312300323-2002001111002222-0123013232312101-3322021303131130"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this site.

Receipt-pinned upstream constraints:

```json
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
storage_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2100110232010113-0111123112001321-2030210011133032-2303221031011113-2301130330212311-0031311033322133-0132221202002320-2333201330032103"></a>

## Direct properties — storage_class_list / 201322003131 / 3

- [storage_classes](resources--azure_vnet_site--reference--group-009.md#canonical-0310003233021333-1233033012301212-3212331021002322-2300331102212320-1310113332033123-2010022000102222-1311331231220120-0033111013023030): complete subsection reference.

<a id="canonical-3020211011121313-3113001202103303-1023103100123302-2111020322120002-0022223330320011-2320121033023303-0322020212122202-3023321030121010"></a>

## Next pages — storage_class_list / 201322003131 / 4

- [voltstack_cluster.storage_class_list.storage_classes](resources--azure_vnet_site--reference--group-009.md#canonical-0310003233021333-1233033012301212-3212331021002322-2300331102212320-1310113332033123-2010022000102222-1311331231220120-0033111013023030)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0310003233021333-1233033012301212-3212331021002322-2300331102212320-1310113332033123-2010022000102222-1311331231220120-0033111013023030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012330320333022-2022033223202222-0332331203300210-1303201121130312-3011131011301322-0232312233232121-1112333201313033-3332013111111023"></a>

## voltstack_cluster.storage_class_list.storage_classes — storage_classes / 203032233331 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110)
- [voltstack_cluster.storage_class_list](resources--azure_vnet_site--reference--group-009.md#canonical-3311322312000000-0231210321003323-2300131023112023-2003010121321010-0331001121323123-1103330000301123-0311230333000223-0302132310200030)
- voltstack_cluster.storage_class_list.storage_classes

<a id="canonical-0113033330122001-0120111302101101-3333120212002130-0010223232333210-1031322002000312-2231321002110211-2133103303300320-1300233301102312"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_class_name")}
```

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_classes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011103011132031-3322022213332113-1222213220321322-1010011103200032-3111320002031130-2213033232313233-2011100002303023-3320201020110203"></a>

## Direct properties — storage_classes / 203032233331 / 3

<a id="canonical-1332300112121202-3010103133203010-3023312220112102-1212221221223121-0231202220230203-1213132313303112-1132131312133131-2123323031331120"></a>

<a id="canonical-2323101012232113-0011113311132032-1111332331121023-2032331311231032-1132101003310011-3023111010223031-1100212013312131-2322111210100223"></a>

## default_storage_class property — storage_classes / 203032233331 / 4

Type: `"bool"`. Optional.

Make this storage class default storage class for the K8s cluster.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0223121330020011-1300200310013101-2030301113213001-2003002233113211-0223032023113132-1303210332100013-3323323322100130-2310111112132132"></a>

<a id="canonical-0133303301121101-3022201021220331-2111111233300101-1232312121303300-2030021133322221-3132230021333330-1330201102032100-2020030132320130"></a>

## storage_class_name property — storage_classes / 203032233331 / 5

Type: `"string"`. Optional.

Name of the storage class as it will appear in K8s.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-1222002012011030-1022131123233303-2133101210202130-2310013133310000-2201201230020121-0132133103222001-3013210112321321-1111111100022230"></a>

## Next pages — storage_classes / 203032233331 / 6

- [voltstack_cluster.storage_class_list](resources--azure_vnet_site--reference--group-009.md#canonical-3311322312000000-0231210321003323-2300131023112023-2003010121321010-0331001121323123-1103330000301123-0311230333000223-0302132310200030)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221003021221103-2313023012310030-3132323210202031-3110000331212101-0223020103023321-2011303111103101-1111223102311201-1200332301310130"></a>

## voltstack_cluster_ar — voltstack_cluster_ar / 112002133033 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- voltstack_cluster_ar

<a id="canonical-1313311132320000-1301223112230310-0121230020333323-1030200331312123-2030031332122210-2113212122202120-2123022120331223-2222210121113022"></a>

Type: `"object"`. single nested block, Optional.

App Stack Cluster of single interface Azure nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("azure_certified_hw"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("default_storage",
    "storage_class_list"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("k8s_cluster",
    "no_k8s_cluster"),
  validators.ConflictingObjectAttributes("no_outside_static_routes",
    "outside_static_routes"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-k8s_cluster_choice": "[\"k8s_cluster\",\"no_k8s_cluster\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage\",\"storage_class_list\"]"
}
```

Terraform syntax:

```terraform
voltstack_cluster_ar {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121211312202132-1001332212330333-1123321232320013-0210321300321011-0231101223201313-3222300220222032-0001333311323223-1313022312022013"></a>

## Direct properties — voltstack_cluster_ar / 112002133033 / 3

- [accelerated_networking](resources--azure_vnet_site--reference--group-009.md#canonical-3002321320303033-1213323210311010-0303311132021330-2121303300120221-2221302120102332-1300102103231102-1133203211321202-2002031232133031): complete subsection reference.

- [active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-009.md#canonical-1301330232303323-2200110003013023-2121133213033011-0220001212130002-2332201223211310-1212130332121021-0223323003002111-1310131211131231): complete subsection reference.

- [active_forward_proxy_policies](resources--azure_vnet_site--reference--group-009.md#canonical-3323010331233020-2211330010321210-1303030013330122-1201200113113113-2223212323300031-1133103123120310-2132021301110321-0210321320030110): complete subsection reference.

- [active_network_policies](resources--azure_vnet_site--reference--group-009.md#canonical-1233133200231232-0111002100320203-3302220112013003-0121323211220001-2120131000333102-3000131221132322-3231003211301311-1000112203333023): complete subsection reference.

<a id="canonical-3001211031220131-2311302203120021-1300013222101132-1112122022211213-0110211323012003-0133303233110222-2230103233301200-2122303101110222"></a>

<a id="canonical-0312323122312201-3131201010120123-3102300232221230-2303213021113312-0103033031030112-0010122310002202-0301031200000310-1322201300020322"></a>

## azure_certified_hw property — voltstack_cluster_ar / 112002133033 / 4

Type: `"string"`. Optional.

\[Enum: Azure-byol-voltstack-combo\] Azure Certified Hardware. Name for Azure certified hardware.
The only possible value is \`azure-byol-voltstack-combo\`.

Upstream description:

Name for Azure certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("azure-byol-voltstack-combo"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-voltstack-combo"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [dc_cluster_group](resources--azure_vnet_site--reference--group-009.md#canonical-0203101011132113-2202222021100200-2213321231233321-1132112100122230-1021212020111311-2121113122110313-1300312001030233-1132330023111201): complete subsection reference.

- [default_storage](resources--azure_vnet_site--reference--group-009.md#canonical-0232003110213133-2001122003300233-1030011331010030-3023331010023221-0310202003112033-3012320001231022-0332301211101212-3300013331011033): complete subsection reference.

- [forward_proxy_allow_all](resources--azure_vnet_site--reference--group-009.md#canonical-0112221012310323-2302101231000122-2003230023300322-3333221012100022-1300311131222303-2021201030333111-0103301222032101-2003012331231223): complete subsection reference.

- [global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-0302133020210101-2221032211110211-1032111131302013-3223120131021120-3213310203101213-1110303313200321-2030023121010000-1210111333331312): complete subsection reference.

- [k8s_cluster](resources--azure_vnet_site--reference--group-010.md#canonical-0110300210201321-2233321233202030-0203210323012120-2311222321101002-0110210121020330-0320321103302133-0021332210302312-1100012010211331): complete subsection reference.

- [no_dc_cluster_group](resources--azure_vnet_site--reference--group-010.md#canonical-1312111013220222-1211300132311131-3210122100203322-3220302230110020-0303331211110022-3331201313233122-1030212312123333-2033112102210032): complete subsection reference.

- [no_forward_proxy](resources--azure_vnet_site--reference--group-010.md#canonical-3003200021130023-1233323111132230-1011120203303301-0230303301011312-0211003011133211-1000331102100331-3221202211111312-1222010233330313): complete subsection reference.

- [no_global_network](resources--azure_vnet_site--reference--group-010.md#canonical-3303132032223003-0010031311031112-3312103133303103-0222233230331001-3103303201211301-3313021100103310-0030313021132011-0120112203023332): complete subsection reference.

- [no_k8s_cluster](resources--azure_vnet_site--reference--group-010.md#canonical-3013312222012123-1001012001300003-1303113001030332-2212223031300022-0012022323110023-1102301123102303-0101303233100120-1213220011032303): complete subsection reference.

- [no_network_policy](resources--azure_vnet_site--reference--group-010.md#canonical-3210030201312131-2301022101213032-3120333110130331-3221302301101010-1233123102210000-3222133213211020-0222302332321320-0330301231133113): complete subsection reference.

- [no_outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-2113103013030123-3333113313303300-3313310320000031-1220022002113312-0332300321010030-2330010210211013-3310103010231303-3022211021103012): complete subsection reference.

- [node](resources--azure_vnet_site--reference--group-010.md#canonical-3323133232110233-1133112012201220-0312330113021030-1133010221211103-3133033331010223-3031023302233023-3010112130013103-0032221030002012): complete subsection reference.

- [outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211): complete subsection reference.

- [sm_connection_public_ip](resources--azure_vnet_site--reference--group-010.md#canonical-0201321201012211-3131100121131013-0010232103310121-3330001101211033-1231022213213233-2112001113230320-1301330313102111-1130121102200103): complete subsection reference.

- [sm_connection_pvt_ip](resources--azure_vnet_site--reference--group-010.md#canonical-1000100033323233-1201011002000310-1223202013111033-0000100231103002-2022132213031021-3230322203312213-2002023303301213-3213201333110102): complete subsection reference.

- [storage_class_list](resources--azure_vnet_site--reference--group-010.md#canonical-3333220321002101-1012013013201131-3112320212130103-2132312101213111-3011030103300022-0201231001332133-2223211331031221-1312311232113101): complete subsection reference.

<a id="canonical-2301031023313133-3221133131102103-2111331122000322-2123032013230113-3112200310030310-0300201203002003-1320122031110122-1333202112211031"></a>

## Next pages — voltstack_cluster_ar / 112002133033 / 5

- [voltstack_cluster_ar.accelerated_networking](resources--azure_vnet_site--reference--group-009.md#canonical-3002321320303033-1213323210311010-0303311132021330-2121303300120221-2221302120102332-1300102103231102-1133203211321202-2002031232133031)
- [voltstack_cluster_ar.active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-009.md#canonical-1301330232303323-2200110003013023-2121133213033011-0220001212130002-2332201223211310-1212130332121021-0223323003002111-1310131211131231)
- [voltstack_cluster_ar.active_forward_proxy_policies](resources--azure_vnet_site--reference--group-009.md#canonical-3323010331233020-2211330010321210-1303030013330122-1201200113113113-2223212323300031-1133103123120310-2132021301110321-0210321320030110)
- [voltstack_cluster_ar.active_network_policies](resources--azure_vnet_site--reference--group-009.md#canonical-1233133200231232-0111002100320203-3302220112013003-0121323211220001-2120131000333102-3000131221132322-3231003211301311-1000112203333023)
- [voltstack_cluster_ar.dc_cluster_group](resources--azure_vnet_site--reference--group-009.md#canonical-0203101011132113-2202222021100200-2213321231233321-1132112100122230-1021212020111311-2121113122110313-1300312001030233-1132330023111201)
- [voltstack_cluster_ar.default_storage](resources--azure_vnet_site--reference--group-009.md#canonical-0232003110213133-2001122003300233-1030011331010030-3023331010023221-0310202003112033-3012320001231022-0332301211101212-3300013331011033)
- [voltstack_cluster_ar.forward_proxy_allow_all](resources--azure_vnet_site--reference--group-009.md#canonical-0112221012310323-2302101231000122-2003230023300322-3333221012100022-1300311131222303-2021201030333111-0103301222032101-2003012331231223)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-0302133020210101-2221032211110211-1032111131302013-3223120131021120-3213310203101213-1110303313200321-2030023121010000-1210111333331312)
- [voltstack_cluster_ar.k8s_cluster](resources--azure_vnet_site--reference--group-010.md#canonical-0110300210201321-2233321233202030-0203210323012120-2311222321101002-0110210121020330-0320321103302133-0021332210302312-1100012010211331)
- [voltstack_cluster_ar.no_dc_cluster_group](resources--azure_vnet_site--reference--group-010.md#canonical-1312111013220222-1211300132311131-3210122100203322-3220302230110020-0303331211110022-3331201313233122-1030212312123333-2033112102210032)
- [voltstack_cluster_ar.no_forward_proxy](resources--azure_vnet_site--reference--group-010.md#canonical-3003200021130023-1233323111132230-1011120203303301-0230303301011312-0211003011133211-1000331102100331-3221202211111312-1222010233330313)
- [voltstack_cluster_ar.no_global_network](resources--azure_vnet_site--reference--group-010.md#canonical-3303132032223003-0010031311031112-3312103133303103-0222233230331001-3103303201211301-3313021100103310-0030313021132011-0120112203023332)
- [voltstack_cluster_ar.no_k8s_cluster](resources--azure_vnet_site--reference--group-010.md#canonical-3013312222012123-1001012001300003-1303113001030332-2212223031300022-0012022323110023-1102301123102303-0101303233100120-1213220011032303)
- [voltstack_cluster_ar.no_network_policy](resources--azure_vnet_site--reference--group-010.md#canonical-3210030201312131-2301022101213032-3120333110130331-3221302301101010-1233123102210000-3222133213211020-0222302332321320-0330301231133113)
- [voltstack_cluster_ar.no_outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-2113103013030123-3333113313303300-3313310320000031-1220022002113312-0332300321010030-2330010210211013-3310103010231303-3022211021103012)
- [voltstack_cluster_ar.node](resources--azure_vnet_site--reference--group-010.md#canonical-3323133232110233-1133112012201220-0312330113021030-1133010221211103-3133033331010223-3031023302233023-3010112130013103-0032221030002012)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [voltstack_cluster_ar.sm_connection_public_ip](resources--azure_vnet_site--reference--group-010.md#canonical-0201321201012211-3131100121131013-0010232103310121-3330001101211033-1231022213213233-2112001113230320-1301330313102111-1130121102200103)
- [voltstack_cluster_ar.sm_connection_pvt_ip](resources--azure_vnet_site--reference--group-010.md#canonical-1000100033323233-1201011002000310-1223202013111033-0000100231103002-2022132213031021-3230322203312213-2002023303301213-3213201333110102)
- [voltstack_cluster_ar.storage_class_list](resources--azure_vnet_site--reference--group-010.md#canonical-3333220321002101-1012013013201131-3112320212130103-2132312101213111-3011030103300022-0201231001332133-2223211331031221-1312311232113101)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3002321320303033-1213323210311010-0303311132021330-2121303300120221-2221302120102332-1300102103231102-1133203211321202-2002031232133031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100233300010333-0200102222021131-3030002200333100-3301302211113221-2012300221330331-0211112331230230-3002131300333131-2311113033031010"></a>

## voltstack_cluster_ar.accelerated_networking — accelerated_networking / 323133220133 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.accelerated_networking

<a id="canonical-0333323323223301-1032230033020102-1323100333320023-1223101130133023-0010003132232331-1332100323112113-1320002102123130-1231012233323301"></a>

Type: `"object"`. single nested block, Optional.

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Upstream description:

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-accelerated_networking": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
accelerated_networking {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112132210300001-0103233303123123-3000033230330313-1033210111110120-2032303203302032-2110001333003010-0102001320112312-0001332033203221"></a>

## Direct properties — accelerated_networking / 323133220133 / 3

- [disable_spec](resources--azure_vnet_site--reference--group-009.md#canonical-3332331303123331-1200112123121001-2202320122013033-1013303331132310-0203300032120203-2111013321030031-0113221012230302-0301100101130022): complete subsection reference.

- [enable](resources--azure_vnet_site--reference--group-009.md#canonical-1101201210320310-2023013232023130-0301231312101003-1332011131221113-3301112022211103-1231233223203122-1202321231111021-1030200313110023): complete subsection reference.

<a id="canonical-3300301203301201-3231133221211203-0202300030101112-2322331302233022-1000123132212302-0102301302102220-2302121231131230-2302103301211211"></a>

## Next pages — accelerated_networking / 323133220133 / 4

- [voltstack_cluster_ar.accelerated_networking.disable_spec](resources--azure_vnet_site--reference--group-009.md#canonical-3332331303123331-1200112123121001-2202320122013033-1013303331132310-0203300032120203-2111013321030031-0113221012230302-0301100101130022)
- [voltstack_cluster_ar.accelerated_networking.enable](resources--azure_vnet_site--reference--group-009.md#canonical-1101201210320310-2023013232023130-0301231312101003-1332011131221113-3301112022211103-1231233223203122-1202321231111021-1030200313110023)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3332331303123331-1200112123121001-2202320122013033-1013303331132310-0203300032120203-2111013321030031-0113221012230302-0301100101130022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201220202210321-2220032002133031-0101203223101101-2000100223312330-1303211300001102-2022021121102213-3103333330301110-2312333101102330"></a>

## voltstack_cluster_ar.accelerated_networking.disable_spec — disable_spec / 302130320300 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.accelerated_networking](resources--azure_vnet_site--reference--group-009.md#canonical-3002321320303033-1213323210311010-0303311132021330-2121303300120221-2221302120102332-1300102103231102-1133203211321202-2002031232133031)
- voltstack_cluster_ar.accelerated_networking.disable_spec

<a id="canonical-3232113232203100-3033031010311233-2222320033333311-1110131021221132-1021101113213233-0003032010103113-0123232330003103-1111120002030033"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-3322032013230310-3330301001011033-1313210321110332-1023323021013013-2122300222122012-1101032301012031-1213021000231123-1003302031320010"></a>

## Direct properties — disable_spec / 302130320300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210123133120202-3102111002212102-2100301112023033-3321030023203133-3200113223202223-3212333020311313-2012231110211230-3333221110233321"></a>

## Next pages — disable_spec / 302130320300 / 4

- [voltstack_cluster_ar.accelerated_networking](resources--azure_vnet_site--reference--group-009.md#canonical-3002321320303033-1213323210311010-0303311132021330-2121303300120221-2221302120102332-1300102103231102-1133203211321202-2002031232133031)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1101201210320310-2023013232023130-0301231312101003-1332011131221113-3301112022211103-1231233223203122-1202321231111021-1030200313110023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130022032122223-2112302100311230-1311230230202112-1220030020223220-2310200030023020-0111012032210000-2303322013312222-2131003130020332"></a>

## voltstack_cluster_ar.accelerated_networking.enable — enable / 202302203020 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.accelerated_networking](resources--azure_vnet_site--reference--group-009.md#canonical-3002321320303033-1213323210311010-0303311132021330-2121303300120221-2221302120102332-1300102103231102-1133203211321202-2002031232133031)
- voltstack_cluster_ar.accelerated_networking.enable

<a id="canonical-3131220133000012-1301001133302333-2132100320112022-3102033200230002-3320013331020122-0321322220212101-1213313222302012-2322121200213213"></a>

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
enable = {}
```

<a id="canonical-2200012022332010-2111002213302200-1313113221010012-1300001331130012-3202300023300130-3111310100210230-1332103231132112-1231231232322300"></a>

## Direct properties — enable / 202302203020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302011120223021-2122002130130320-0321120230021011-2232320133203203-1110302331323302-1220111323202122-1131313303032201-3020013121131211"></a>

## Next pages — enable / 202302203020 / 4

- [voltstack_cluster_ar.accelerated_networking](resources--azure_vnet_site--reference--group-009.md#canonical-3002321320303033-1213323210311010-0303311132021330-2121303300120221-2221302120102332-1300102103231102-1133203211321202-2002031232133031)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1301330232303323-2200110003013023-2121133213033011-0220001212130002-2332201223211310-1212130332121021-0223323003002111-1310131211131231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112001023020011-3120102032000221-3000033122322000-3222332002231320-3320320221332123-3112233032021331-0202223100310101-0101130033111110"></a>

## voltstack_cluster_ar.active_enhanced_firewall_policies — active_enhanced_firewall_policies / 110132112223 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.active_enhanced_firewall_policies

<a id="canonical-0013200333323000-2301301331032121-1220011023130001-0232130103000312-2000022320332232-3301201131022332-0030310301232001-3320203123133212"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011030003013121-2023301033130301-2012233231101122-0321332302320010-3220021332202231-2111130013102303-1033302331002323-0212011022102231"></a>

## Direct properties — active_enhanced_firewall_policies / 110132112223 / 3

- [enhanced_firewall_policies](resources--azure_vnet_site--reference--group-009.md#canonical-1333300110101300-2113011000032112-0210022313122112-2000121213213110-2033212130130230-1023110313112111-0233102101201000-0203030013010122): complete subsection reference.

<a id="canonical-1332310023031212-0120302220102233-2113032331331222-0122100130231022-0003122320303113-3130332320310002-2223211010322232-2103133221332333"></a>

## Next pages — active_enhanced_firewall_policies / 110132112223 / 4

- [voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--azure_vnet_site--reference--group-009.md#canonical-1333300110101300-2113011000032112-0210022313122112-2000121213213110-2033212130130230-1023110313112111-0233102101201000-0203030013010122)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1333300110101300-2113011000032112-0210022313122112-2000121213213110-2033212130130230-1023110313112111-0233102101201000-0203030013010122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032300033223203-3211113011202022-2212120210002220-3322201231303100-2211301022200113-1022203011112122-0102133133231123-3220311201110330"></a>

## voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 332032011221 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-009.md#canonical-1301330232303323-2200110003013023-2121133213033011-0220001212130002-2332201223211310-1212130332121021-0223323003002111-1310131211131231)
- voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-0132211300031312-0110332112210011-3212101000231211-1130030312011300-2231303012223210-0021323113033202-2320130101121231-3121233313121031"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020301100320100-2113020303210332-2322310103220301-3323201221022113-1101010131121002-2133100033030322-1031102130032120-2011121303220223"></a>

## Direct properties — enhanced_firewall_policies / 332032011221 / 3

<a id="canonical-1111112230031332-3131023311301212-3320132303200122-2313110123232312-3100323033200220-2211010232110033-2011013123020131-2303333311010320"></a>

<a id="canonical-2011023132303313-2000201220120003-2302112312133333-0320122003330322-1201233310122210-2300303200302322-2302211111122323-3011330003012323"></a>

## name property — enhanced_firewall_policies / 332032011221 / 4

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

<a id="canonical-1213030133202313-3330232213202021-2002213002001120-2332301122121121-0131300230131210-2332312021130212-0201313012012333-3331131001131233"></a>

<a id="canonical-3123010210220032-2122301301230321-3302202022312020-0120102012230201-0323001021201001-1030020210131102-3103202212000211-3011213021202112"></a>

## namespace property — enhanced_firewall_policies / 332032011221 / 5

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

<a id="canonical-0300103120202311-0221130012230323-0310003020001203-3333101322200102-1113233032312132-0230201002031213-3311130310222101-3121221231002103"></a>

<a id="canonical-0001323032220311-2210033132233013-0211201023221020-2113300231230301-3331213131331033-1323021212201023-0301201111032203-2222233221012220"></a>

## tenant property — enhanced_firewall_policies / 332032011221 / 6

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

<a id="canonical-3131302102302222-1000220102210122-0102103202302312-1003030231131031-1313222023222220-1123332210120003-2103320120013121-1312232303132201"></a>

## Next pages — enhanced_firewall_policies / 332032011221 / 7

- [voltstack_cluster_ar.active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-009.md#canonical-1301330232303323-2200110003013023-2121133213033011-0220001212130002-2332201223211310-1212130332121021-0223323003002111-1310131211131231)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3323010331233020-2211330010321210-1303030013330122-1201200113113113-2223212323300031-1133103123120310-2132021301110321-0210321320030110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022300211322213-3121200223302000-3111021032133121-1301022103031003-1232321030212023-0330310303313233-2023031120333232-0113123101022232"></a>

## voltstack_cluster_ar.active_forward_proxy_policies — active_forward_proxy_policies / 010333012133 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.active_forward_proxy_policies

<a id="canonical-0123022310332022-2003202101023223-3321113033223213-2322231012201320-0331211111320102-0011111122122302-1331011223002323-0202120013202330"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111133220001102-2132223002111023-2312013122203131-0302303011322203-1012302312210220-1303220103311013-0301101021333311-0113301010121303"></a>

## Direct properties — active_forward_proxy_policies / 010333012133 / 3

- [forward_proxy_policies](resources--azure_vnet_site--reference--group-009.md#canonical-1320021221120312-0100321221020313-3033212121000213-3002232310313303-1210013033110100-1112112301033130-3232210321023221-1202303303130023): complete subsection reference.

<a id="canonical-1103020320333010-1200311202020033-1023100200132100-2010123023303233-3322313231122321-2303131312121221-2001030130311122-0113002122312101"></a>

## Next pages — active_forward_proxy_policies / 010333012133 / 4

- [voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies](resources--azure_vnet_site--reference--group-009.md#canonical-1320021221120312-0100321221020313-3033212121000213-3002232310313303-1210013033110100-1112112301033130-3232210321023221-1202303303130023)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1320021221120312-0100321221020313-3033212121000213-3002232310313303-1210013033110100-1112112301033130-3232210321023221-1202303303130023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223310212223112-2332131213213211-3112301031200330-1210022233121222-2133022231331203-2331130032111231-1323132111232111-2131130312331100"></a>

## voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 133201010230 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.active_forward_proxy_policies](resources--azure_vnet_site--reference--group-009.md#canonical-3323010331233020-2211330010321210-1303030013330122-1201200113113113-2223212323300031-1133103123120310-2132021301110321-0210321320030110)
- voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-2132301123312212-2202101131112203-2002122320220232-0233133330001213-0203012010221200-3103333303103133-2102010312022111-1013132033223131"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111231012333131-2213011200131120-0232031333201013-2000313022322121-1120112102023323-0311331310023123-1230213032323003-1113210331313310"></a>

## Direct properties — forward_proxy_policies / 133201010230 / 3

<a id="canonical-0201120333211131-1021233222031302-1302133112302010-0220010223132123-2302301212031210-3020233103131121-2030330311133233-2312123321222211"></a>

<a id="canonical-2011111212003302-0313313122101321-1001131122003110-3211200022333100-1031112101113311-3332303223323332-2030310011301223-3221130021120033"></a>

## name property — forward_proxy_policies / 133201010230 / 4

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

<a id="canonical-1310331210122103-0221022011021301-3011123020133111-0322220300032013-3320323331310112-3231310321101313-0013300331301322-3223301121012202"></a>

<a id="canonical-0100313311103321-1101112001323023-0231131012213012-2321311022000231-0213130000033030-2323230232122122-2133020212110331-2202021203031323"></a>

## namespace property — forward_proxy_policies / 133201010230 / 5

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

<a id="canonical-1203120333201310-1312200200222130-2030033020030110-3222011112021230-3000312110303121-2302123300102331-1220232223303130-0111131202232122"></a>

<a id="canonical-2301032322332233-3312122101130121-3012220120221221-1300233023201312-1212231211202122-0321323303323203-3320020302323012-3002112313012223"></a>

## tenant property — forward_proxy_policies / 133201010230 / 6

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

<a id="canonical-2300001223122311-3320030320321022-1113122230030100-0231320113310323-3220103223011031-0210031002102312-3230122120101133-3101322330301231"></a>

## Next pages — forward_proxy_policies / 133201010230 / 7

- [voltstack_cluster_ar.active_forward_proxy_policies](resources--azure_vnet_site--reference--group-009.md#canonical-3323010331233020-2211330010321210-1303030013330122-1201200113113113-2223212323300031-1133103123120310-2132021301110321-0210321320030110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1233133200231232-0111002100320203-3302220112013003-0121323211220001-2120131000333102-3000131221132322-3231003211301311-1000112203333023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223210210133002-1322000031030031-0022311112302120-1000101011200220-1003101122201022-3222030013312332-1123303223103330-0323200021003122"></a>

## voltstack_cluster_ar.active_network_policies — active_network_policies / 101010020323 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.active_network_policies

<a id="canonical-0230223000112222-2313103230000011-3032303100001200-1302011221020233-3032202033132033-0203203013322332-2103200230221033-3112030002203331"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1112001103020132-3000211203220121-0111322030010101-1002223231202113-0312112131230132-2320020322322312-3223300030031012-1302332322202001"></a>

## Direct properties — active_network_policies / 101010020323 / 3

- [network_policies](resources--azure_vnet_site--reference--group-009.md#canonical-2013321030011033-0012102012023302-2330011010330103-1030200312123303-3130101220113103-3011010011103100-3120023100031300-0003030122012103): complete subsection reference.

<a id="canonical-2021103220021020-2233020112303203-2333112211120202-2220323022033113-0310130032330231-2310220113201212-1121102002123302-3331003300001032"></a>

## Next pages — active_network_policies / 101010020323 / 4

- [voltstack_cluster_ar.active_network_policies.network_policies](resources--azure_vnet_site--reference--group-009.md#canonical-2013321030011033-0012102012023302-2330011010330103-1030200312123303-3130101220113103-3011010011103100-3120023100031300-0003030122012103)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2013321030011033-0012102012023302-2330011010330103-1030200312123303-3130101220113103-3011010011103100-3120023100031300-0003030122012103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232211033031203-3301223301312122-2320231312303203-0120230110110300-2331220001113210-2113301031023011-1112021003222312-2112020031100211"></a>

## voltstack_cluster_ar.active_network_policies.network_policies — network_policies / 102110321203 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.active_network_policies](resources--azure_vnet_site--reference--group-009.md#canonical-1233133200231232-0111002100320203-3302220112013003-0121323211220001-2120131000333102-3000131221132322-3231003211301311-1000112203333023)
- voltstack_cluster_ar.active_network_policies.network_policies

<a id="canonical-1200220303002103-3211123100030121-1230122332123000-2200032030103310-0231222102010332-0232032022211130-0330023002012110-2003133210120331"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223103322330212-2003133333022233-2323333211223231-0131100110223012-2123100210300103-0012010233310133-2230001032230111-3302211203123201"></a>

## Direct properties — network_policies / 102110321203 / 3

<a id="canonical-3102032032100220-0101003201201131-1112133330233012-1302001333033233-3303000311330001-2030223223201121-2130232133222212-0121301303323332"></a>

<a id="canonical-2001002021213203-3320031230100103-2133232212013002-0012020320012231-0002212323333303-0210230302111103-1330301322322310-0201133323303001"></a>

## name property — network_policies / 102110321203 / 4

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

<a id="canonical-0310200101011220-0301013231330133-3022010021011302-0020302121003012-2232033033031013-2002231122322320-2132112231223130-0320103031101303"></a>

<a id="canonical-2112113213033233-0133302012313102-0003123300331111-0032330103032321-1130222323022123-3231100022033312-1322130230033200-3313123211000201"></a>

## namespace property — network_policies / 102110321203 / 5

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

<a id="canonical-2013301020023100-1001223023021112-2032202011000303-2221322001031203-1022231320001230-3302122032200203-0123033130301200-3131101301321302"></a>

<a id="canonical-3212331202122031-3332303013001122-0103010202133120-1301202001103103-0321033100333302-2003231212110211-0001321302022013-0132120211301221"></a>

## tenant property — network_policies / 102110321203 / 6

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

<a id="canonical-3320032131023030-3331102212330003-1231001030313230-1002010122001031-1132232302002231-0101312110120203-3022300232231313-0110213212330102"></a>

## Next pages — network_policies / 102110321203 / 7

- [voltstack_cluster_ar.active_network_policies](resources--azure_vnet_site--reference--group-009.md#canonical-1233133200231232-0111002100320203-3302220112013003-0121323211220001-2120131000333102-3000131221132322-3231003211301311-1000112203333023)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0203101011132113-2202222021100200-2213321231233321-1132112100122230-1021212020111311-2121113122110313-1300312001030233-1132330023111201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010032311230330-1111003122113331-1310131112231112-3102000002112221-2230220002311031-2020313220213222-2011013001220320-2221123130012132"></a>

## voltstack_cluster_ar.dc_cluster_group — dc_cluster_group / 121232201200 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.dc_cluster_group

<a id="canonical-1301212121303003-3200223210220203-1002103131210101-3133121201311221-2231323020000320-3101101122301021-3000221021213130-3013131020211223"></a>

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
dc_cluster_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332010301110130-0010123330320300-0223323031232121-2132222222322220-0030133233102323-1211320001321110-3212130203012033-1131232112022223"></a>

## Direct properties — dc_cluster_group / 121232201200 / 3

<a id="canonical-3132012100023013-3122211313031311-3312303223200120-2233122110121132-0331201212202111-3022311201113330-0002223332010312-2111330013022033"></a>

<a id="canonical-3102112102220332-0023231312213331-0331200110031012-0321110331123201-1231002232210220-0230223223130312-3300223030100022-1201011121200300"></a>

## name property — dc_cluster_group / 121232201200 / 4

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

<a id="canonical-1033121221132032-3111000023203132-1203120113030312-3301111222030032-3301132331200232-1231111210320133-0222312022213301-0220212033011013"></a>

<a id="canonical-0202030023331121-0213300333300003-0223221232003123-3032313320232211-3101321221301302-2133301233300012-3121123333233231-2120300112233033"></a>

## namespace property — dc_cluster_group / 121232201200 / 5

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

<a id="canonical-0230202230203021-2331332132032303-3001313110021302-3123102121011210-0213031331011301-0111332003213023-3022010233331132-2313012221123312"></a>

<a id="canonical-1232330113102121-1100203113213221-1320132230003130-0313131113223220-3301203333023033-1022321312001012-2130131133331313-0111220102201033"></a>

## tenant property — dc_cluster_group / 121232201200 / 6

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

<a id="canonical-0001332013113232-3222102033103003-1320212101021331-2011031201031122-1302203100230212-0203033112200021-0321021203222100-2020120200103301"></a>

## Next pages — dc_cluster_group / 121232201200 / 7

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0232003110213133-2001122003300233-1030011331010030-3023331010023221-0310202003112033-3012320001231022-0332301211101212-3300013331011033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011323032013001-3301200102000000-1133220100020000-3023132033200030-2310212231110221-1213321101311030-0300013210000323-0122203011020013"></a>

## voltstack_cluster_ar.default_storage — default_storage / 033012021110 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.default_storage

<a id="canonical-3213300220130213-0112023032300320-2301222213132330-0003312323323323-1003123213023203-1132001202221323-1223133012130211-3330230233123131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default storage.

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
default_storage = {}
```

<a id="canonical-1011312331132210-1031201021323101-3130030133032212-2120230221321002-1310212022202222-2031103333111011-1230101330311323-2333120333323210"></a>

## Direct properties — default_storage / 033012021110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122022221223221-3301221012310012-1231010221013303-0122132102020033-1321113101212213-2330123213222302-2030302133330003-1201001300003032"></a>

## Next pages — default_storage / 033012021110 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0112221012310323-2302101231000122-2003230023300322-3333221012100022-1300311131222303-2021201030333111-0103301222032101-2003012331231223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030002302023232-3130311120310232-3122023333311012-0311131131230132-1231202023102221-0132011322200203-1023301110231132-3313020312210321"></a>

## voltstack_cluster_ar.forward_proxy_allow_all — forward_proxy_allow_all / 321310121110 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.forward_proxy_allow_all

<a id="canonical-1033002200131301-2010112021332213-3333003010133230-2312212332030013-0131332310203231-2310320001003321-2132221003111330-3320123032333201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for forward proxy allow all.

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
forward_proxy_allow_all = {}
```

<a id="canonical-0112223133012002-2202211012020003-0323021333121202-3311212113231013-1121023321233311-2230011032222222-0322003233233002-3332320313201302"></a>

## Direct properties — forward_proxy_allow_all / 321310121110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332220323300221-0000120121011331-3211231320103102-1202133303111213-1200022200133132-0223333313320333-0112202233200330-1313113002102123"></a>

## Next pages — forward_proxy_allow_all / 321310121110 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0302133020210101-2221032211110211-1032111131302013-3223120131021120-3213310203101213-1110303313200321-2030023121010000-1210111333331312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133131320330023-2220212100010313-3101332012033322-0002321320133010-1322023232031023-3011231111201233-2222210023002102-0133110003211032"></a>

## voltstack_cluster_ar.global_network_list — global_network_list / 231303301210 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.global_network_list

<a id="canonical-0022122033130001-3200333011322112-0001021313310001-3102303310033232-0313220110113322-3312131021313212-3131020122023133-1101220012003003"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```
