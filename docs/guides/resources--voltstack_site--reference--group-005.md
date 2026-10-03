---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-3210213020011333-0233323330210212-3201123302012310-2030221002121311-0221011310303203-1302222303120101-3203320011221121-3323202021332202"></a>

## custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.cluster_static_ip — cluster_static_ip / 332130012232 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.tunnel_interface](resources--voltstack_site--reference--group-004.md#canonical-1012322230030023-3000301300302131-3203221332301322-2103012002021321-0111030132311312-3313233303332202-2323030121333310-3100101233000101)
- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](resources--voltstack_site--reference--group-004.md#canonical-0201332103101310-0131121012221230-2030111132322122-3320003020121002-0321232233301130-2013120022103301-0112121302232331-2330000322001012)
- custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.cluster_static_ip

<a id="canonical-0212120311220100-0222302031320020-3010003210212101-3331012321030332-1313231003222000-0231201232031200-0211202300233002-3023013031120121"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for cluster.

Receipt-pinned upstream constraints:

```json
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
cluster_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313320003021202-2122220303312020-1223203022200023-2330130213123111-3233230332211020-0210301332132331-1232332111000202-2202122203120331"></a>

## Direct properties — cluster_static_ip / 332130012232 / 3

<a id="canonical-3113232020002333-0101203000103201-0030012221031323-3223200213133123-2033121333011002-2010201020013301-2003233222101122-0031310301000111"></a>

<a id="canonical-0212301333122132-2203332100213301-0230312201233023-1002211200223211-3022203200001021-1131033330123032-0133023221011331-1121333311111303"></a>

## interface_ip_map property — cluster_static_ip / 332130012232 / 4

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 128,
    "metadata": {
      "confidence": 0.75,
      "source": "inferred",
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-3302002311320321-1111221202233232-1210020321120332-1232130312121333-3232333210033221-1123330013132300-1001010332220021-3022302213002133"></a>

## Next pages — cluster_static_ip / 332130012232 / 5

- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](resources--voltstack_site--reference--group-004.md#canonical-0201332103101310-0131121012221230-2030111132322122-3320003020121002-0321232233301130-2013120022103301-0112121302232331-2330000322001012)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1310022322011031-1033323233223213-3310333231023311-3000233230332023-3311301311013330-0323210113010020-2013000203230123-3202101120120102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123311231303133-0032331031003120-3020131322302300-3020220223120322-0130212213023322-3110131110223023-2221033322213230-0033011330002032"></a>

## custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.node_static_ip — node_static_ip / 211331202011 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.tunnel_interface](resources--voltstack_site--reference--group-004.md#canonical-1012322230030023-3000301300302131-3203221332301322-2103012002021321-0111030132311312-3313233303332202-2323030121333310-3100101233000101)
- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](resources--voltstack_site--reference--group-004.md#canonical-0201332103101310-0131121012221230-2030111132322122-3320003020121002-0321232233301130-2013120022103301-0112121302232331-2330000322001012)
- custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.node_static_ip

<a id="canonical-0321203320202120-3201132023311221-1230121110332103-1122123201131023-1231303320021302-0112123231121231-0320200131020203-2001120202320120"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
node_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033101303231210-0102221132133133-3333022001112312-1321223310223133-2110031102011123-3123123121120223-2030201200132003-0110201112221012"></a>

## Direct properties — node_static_ip / 211331202011 / 3

<a id="canonical-1220330103312112-1133112131023310-3103031323213201-3322003122333033-3302330113231300-1013233030202330-1021133201032332-3320331331131122"></a>

<a id="canonical-3322103131231221-0001133010233010-3313112132220331-2220020202002013-2221210021323131-2132122032011111-2001111213130223-3331012110330112"></a>

## default_gw property — node_static_ip / 211331202011 / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-0101303102220223-2102210322210233-3121202201123302-2121210020323102-0120031212331030-2100301201332231-1001220103322022-0333312033213312"></a>

<a id="canonical-2131210332303032-0023201203302111-2300002032303120-2103213111332302-3000333103030211-1300121220213031-0133010230302311-3223001002030012"></a>

## dns_server property — node_static_ip / 211331202011 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3110032312221023-1103103301332033-1030122002003310-3233302112123320-3202013021321333-1310222201233231-2210212012122313-2031201213011023"></a>

<a id="canonical-2232033331131122-1013020030123300-0303002133313112-2320331321321321-1333233322210010-1201103031221332-3321102021330231-2233113212201101"></a>

## ip_address property — node_static_ip / 211331202011 / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-1003010031103101-3322131003232011-3320302000122312-1333321213301102-3032010313323201-0123031331000001-3021120210231120-3300330223133313"></a>

## Next pages — node_static_ip / 211331202011 / 7

- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](resources--voltstack_site--reference--group-004.md#canonical-0201332103101310-0131121012221230-2030111132322122-3320003020121002-0321232233301130-2013120022103301-0112121302232331-2330000322001012)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0220322301001222-3022112332322211-2232333231020032-0201320212231012-3331231230023133-1032013112021312-1002130020322021-3220133323211300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332101000123221-0012030000033331-0111000023233133-3322303010010312-0032312332223320-3133201010202121-0221322331331320-2023033211111020"></a>

## custom_network_config.interface_list.interfaces.tunnel_interface.tunnel — tunnel / 311110020103 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.interface_list](resources--voltstack_site--reference--group-003.md#canonical-0100203223120222-0230103002031132-0332203202020213-0212233113013100-3330311013022230-0332030200210121-2020333030213323-0133002123021131)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--reference--group-003.md#canonical-3323232300012332-3233133221220120-1001103003112132-0332100111102230-2023131333013233-0131310021221101-0000121002122332-2202120030322211)
- [custom_network_config.interface_list.interfaces.tunnel_interface](resources--voltstack_site--reference--group-004.md#canonical-1012322230030023-3000301300302131-3203221332301322-2103012002021321-0111030132311312-3313233303332202-2323030121333310-3100101233000101)
- custom_network_config.interface_list.interfaces.tunnel_interface.tunnel

<a id="canonical-3103020031132120-1022020110101203-2321120223211320-3232031023132223-2133233300111230-0112012102323333-0121001030303220-3310023303232002"></a>

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
tunnel {
  # Configure direct properties listed below.
}
```

<a id="canonical-1100103201022110-0301122210210201-0013020020001011-3210012121120321-0212033001102132-0021103030011231-3213122032330201-3213112223312211"></a>

## Direct properties — tunnel / 311110020103 / 3

<a id="canonical-3230332201102221-3020133200313330-3101012121303130-3323021012003110-1002122323123131-0122311100331312-1302210230013113-2202320313321212"></a>

<a id="canonical-1321133033302010-3032100200332032-3203011313333103-0102112311331310-3132011121000001-2110310021121132-0323001012000113-0202303120101212"></a>

## name property — tunnel / 311110020103 / 4

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

<a id="canonical-0033023211012121-1013302232000132-1111133032233130-0223220203220330-3001213013211201-2023310211021102-3221333013113211-2110233102230102"></a>

<a id="canonical-1231233112320331-1302222333110302-1133011312221233-1132322310310221-3120311321210301-0330310333221031-3332012123210321-3300301020333131"></a>

## namespace property — tunnel / 311110020103 / 5

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

<a id="canonical-2331313233023101-2310233123230133-0020223320210301-1311300121212013-3122211312211331-1203212203111312-2211022033033220-3200020030231233"></a>

<a id="canonical-2112100032320210-1223011202311233-3011132322102210-3032211002110031-2010212033012233-0233332323003330-1231023313133202-3113231121031200"></a>

## tenant property — tunnel / 311110020103 / 6

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

<a id="canonical-2130322010332303-2223221032113023-2103201003331132-2311111033323101-1213122110310013-2333010302030302-3110333103233113-2230313212033103"></a>

## Next pages — tunnel / 311110020103 / 7

- [custom_network_config.interface_list.interfaces.tunnel_interface](resources--voltstack_site--reference--group-004.md#canonical-1012322230030023-3000301300302131-3203221332301322-2103012002021321-0111030132311312-3313233303332202-2323030121333310-3100101233000101)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2120233011220333-3312223321223103-2102033112021112-1020023232012202-1010202022212101-2113300132330313-0303120303110132-0311202003220300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032011003120031-0131323232131200-0123020210020122-3332120101133010-1003013110332310-3312120101320023-3102323211220103-1303011300110103"></a>

## custom_network_config.no_forward_proxy — no_forward_proxy / 122300033101 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- custom_network_config.no_forward_proxy

<a id="canonical-0203032003212021-2330122002130233-0111013332302202-2312131201313000-0303220322333020-1222100131110210-3102011331012203-3012232202301020"></a>

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

<a id="canonical-1302102010212000-2123222131223302-3331221321131310-3322010131010331-1110131330330133-2132221102212320-0223313021203202-2102301013100303"></a>

## Direct properties — no_forward_proxy / 122300033101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313233310030220-2011113320122320-1302123123020121-2011031031122112-2012220211033233-2233110031320220-2033122032330000-3223203222131103"></a>

## Next pages — no_forward_proxy / 122300033101 / 4

- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3110332012322003-3322011320311033-0000323333000210-0210323011312201-0330122331230323-3020001333023311-2101022000213102-3212132300031121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232033220000101-3303320321311030-1001030132130221-0231230111232333-2123310202211031-2310311332213230-2211011132200212-0113000122003032"></a>

## custom_network_config.no_global_network — no_global_network / 320230221233 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- custom_network_config.no_global_network

<a id="canonical-3210202133333133-0110220201233130-3003321323033323-2320121222332313-1011121021111313-2103330120133130-3231002031201330-1013002202302223"></a>

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

<a id="canonical-1130101330232332-1022123220112013-2230133133323321-3223210120031310-3021121121001100-0103302130103022-3022201330021111-3233220012301222"></a>

## Direct properties — no_global_network / 320230221233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320222132232021-3333202223331230-1030233103332000-3020231300112310-2332030033230032-3011212330222323-1330213203023002-3110201120031111"></a>

## Next pages — no_global_network / 320230221233 / 4

- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1031100003302212-3220302111000131-1223201301122201-0301203203102223-2121101210001333-3030023023320300-1231222003011220-2233331301320213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310132102203330-0212210222103112-3012010311233202-0223030333101213-1212300133033211-0021213330133232-0022330100132303-0230000003110032"></a>

## custom_network_config.no_network_policy — no_network_policy / 220122333303 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- custom_network_config.no_network_policy

<a id="canonical-1200033010133212-1110001011323111-1313030100211333-3222313220012121-0032021200111122-3111123230220121-1032120323132033-2013223011220023"></a>

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

<a id="canonical-3020310100101310-1223001031212300-0021233133123030-3221013203023003-0220312212031201-2223300323021001-2322221101313222-1233333212200100"></a>

## Direct properties — no_network_policy / 220122333303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133021131112230-1220333120322201-0313120111211313-1200203211033223-0111023213322123-0110313003210313-2113220220300100-1000220301213301"></a>

## Next pages — no_network_policy / 220122333303 / 4

- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012230220313313-0333303333000131-2320113231323100-1023200231012320-2231312330333202-1302300011310000-1232020133300002-0223110322331123"></a>

## custom_network_config.sli_config — sli_config / 032330320311 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- custom_network_config.sli_config

<a id="canonical-2210333310200032-1202201133303332-2210030133111020-3132012202000123-2233032331233103-1330102220202202-1231230021313322-3020002311021220"></a>

Type: `"object"`. single nested block, Optional.

Site local inside network configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_v6_static_routes",
    "static_v6_routes")}
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
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
sli_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0021111022300202-0211310230113320-0213312130230322-1200201213120223-2113032132021223-1022212311211120-1002100120113212-3102010223131033"></a>

## Direct properties — sli_config / 032330320311 / 3

- [no_static_routes](resources--voltstack_site--reference--group-005.md#canonical-2000212223331231-3112311021130330-0100203022233300-2021033331031210-0200120202213103-1031222200011332-3013030320312202-2313110321322211): complete subsection reference.

- [no_v6_static_routes](resources--voltstack_site--reference--group-005.md#canonical-2212310133133230-3233010331112221-2331312301003110-0121130322222133-1020133031303300-1233102002331000-2311213310311011-3130202000101230): complete subsection reference.

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-3010201321130121-1003131132022102-1321123300131211-1311222020210013-1210221231201101-3211022010320113-2202213020220333-3232132331233201): complete subsection reference.

- [static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-3303202101313311-2332103211112231-1223201102030123-1322100301333302-2333211100213230-2222122132030211-0311221100323131-2200112333120203): complete subsection reference.

<a id="canonical-0301001232012023-2123102312020132-2301002032012033-1313330233002102-0210013202020003-1010323132312122-1202203012030303-2021103312131231"></a>

## Next pages — sli_config / 032330320311 / 4

- [custom_network_config.sli_config.no_static_routes](resources--voltstack_site--reference--group-005.md#canonical-2000212223331231-3112311021130330-0100203022233300-2021033331031210-0200120202213103-1031222200011332-3013030320312202-2313110321322211)
- [custom_network_config.sli_config.no_v6_static_routes](resources--voltstack_site--reference--group-005.md#canonical-2212310133133230-3233010331112221-2331312301003110-0121130322222133-1020133031303300-1233102002331000-2311213310311011-3130202000101230)
- [custom_network_config.sli_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3010201321130121-1003131132022102-1321123300131211-1311222020210013-1210221231201101-3211022010320113-2202213020220333-3232132331233201)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-3303202101313311-2332103211112231-1223201102030123-1322100301333302-2333211100213230-2222122132030211-0311221100323131-2200112333120203)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2000212223331231-3112311021130330-0100203022233300-2021033331031210-0200120202213103-1031222200011332-3013030320312202-2313110321322211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221203010110110-3213003211113020-3113111333333100-1012113013222002-3010131333000203-2133332323021301-1102211101003100-1330012211120020"></a>

## custom_network_config.sli_config.no_static_routes — no_static_routes / 031223132232 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- custom_network_config.sli_config.no_static_routes

<a id="canonical-1202102213123131-1202130201033220-0203032003200313-3323103113111122-3212320313220012-1202223211212201-0201131030123133-3003023300103031"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static routes.

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
no_static_routes = {}
```

<a id="canonical-3122233121103120-3110011023230032-0211021312123223-1303012000203033-0020331312232121-0032020220300212-2131132023131003-0321003231232301"></a>

## Direct properties — no_static_routes / 031223132232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213313323221330-1023111003132012-0022003131312222-1030222230022022-1220021211110033-1110030133110030-2201221130322010-3330321201220310"></a>

## Next pages — no_static_routes / 031223132232 / 4

- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2212310133133230-3233010331112221-2331312301003110-0121130322222133-1020133031303300-1233102002331000-2311213310311011-3130202000101230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230301032200000-2001102033313210-1113312001103013-1323131121213123-2223323311201131-2031333221113320-3031232230012123-2323010211131122"></a>

## custom_network_config.sli_config.no_v6_static_routes — no_v6_static_routes / 221111220231 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- custom_network_config.sli_config.no_v6_static_routes

<a id="canonical-0200303033231322-3110201210023101-3211322233313131-0333010123310213-0112321320102121-1013003020233033-1031311301221102-3100313023212303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no v6 static routes.

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
no_v6_static_routes = {}
```

<a id="canonical-2133133033120213-2020111300121300-3232221323033201-1233120330030322-0200030210013123-0110313132110103-2003022320222111-3112303032003112"></a>

## Direct properties — no_v6_static_routes / 221111220231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112233330030120-3203100031100131-3101103100011010-2301312212223023-2011101111111101-2013222313331333-3201300132131100-1202212133311330"></a>

## Next pages — no_v6_static_routes / 221111220231 / 4

- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3010201321130121-1003131132022102-1321123300131211-1311222020210013-1210221231201101-3211022010320113-2202213020220333-3232132331233201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231133023123312-0032120002121333-2211212222032103-3320030000212102-2212003302213213-0012130312331110-2202301111103001-2013002120221221"></a>

## custom_network_config.sli_config.static_routes — static_routes / 130322230020 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- custom_network_config.sli_config.static_routes

<a id="canonical-3231202033110132-2021223020213032-1133303123011122-3333130111230011-3010130000100200-1213123012102132-0030332123331232-0110200012200031"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022211012133233-0030230121320123-1030011122211332-3132132031213102-0330311333311201-3123023113301222-0032022201321032-2120233023233221"></a>

## Direct properties — static_routes / 130322230020 / 3

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-3131302122323111-0022332121303122-1031021121211112-3113303020123222-2212011023100203-3010220013323131-0102320221020032-0112300221321101): complete subsection reference.

<a id="canonical-3323033232212011-3323213123003132-2111132231021202-3310231010132001-3001132132030132-1332010233002231-3031201113001102-0300302123003222"></a>

## Next pages — static_routes / 130322230020 / 4

- [custom_network_config.sli_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3131302122323111-0022332121303122-1031021121211112-3113303020123222-2212011023100203-3010220013323131-0102320221020032-0112300221321101)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3131302122323111-0022332121303122-1031021121211112-3113303020123222-2212011023100203-3010220013323131-0102320221020032-0112300221321101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011001212113311-1332122210133230-2131312310312211-2133210232223323-2232220301232011-1313121013023323-0230222230313103-1113213003122201"></a>

## custom_network_config.sli_config.static_routes.static_routes — static_routes / 330022103333 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- [custom_network_config.sli_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3010201321130121-1003131132022102-1321123300131211-1311222020210013-1210221231201101-3211022010320113-2202213020220333-3232132331233201)
- custom_network_config.sli_config.static_routes.static_routes

<a id="canonical-3121321102130231-2233001130122120-2233330013230210-1223212132312021-1302213122230012-1322312331332000-1333013030130303-1112033301022100"></a>

Type: `"object"`. list nested block, Optional.

Static Routes. List of static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111133001112001-0010311312332033-0021121013202213-0033121333303220-1133230013022023-0331102112011233-3001012131302133-0110200010130030"></a>

## Direct properties — static_routes / 330022103333 / 3

<a id="canonical-1302000331233322-0020000032111323-3203311220020321-0121133011301012-1211102133130020-0320311221232210-3223121032112231-2300321301033303"></a>

<a id="canonical-2010203130000101-3231222032213120-0320111132030103-2030102311031212-3300232210202310-0330013310202200-2302221333122032-1001002032030231"></a>

## attrs property — static_routes / 330022103333 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--voltstack_site--reference--group-005.md#canonical-0121133032222000-1131203022200023-2010213302300220-2102033103212112-3202102010123002-1003020300110022-3330000020031232-0031131012223302): complete subsection reference.

<a id="canonical-2200320222000030-2112010332313013-0000311330133232-1303323010311201-2302333031322333-0331032032223102-1310132232031113-1303022203310311"></a>

<a id="canonical-3102311200123110-3221323202331011-1321002123310332-0301203130200012-1223320122022210-0002002333131130-2213031332202130-0201101002102203"></a>

## ip_address property — static_routes / 330022103333 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-3320311313332331-2131102101022023-2032103112103133-0011031203120122-1213301102313021-1202103300131123-2300030200213013-2312320213032012"></a>

<a id="canonical-0110013210312332-3300101303232011-3230323323032102-3330211030122232-3022023100213231-3020222200202110-0231002233321300-2221032323303002"></a>

## ip_prefixes property — static_routes / 330022103333 / 6

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--voltstack_site--reference--group-005.md#canonical-2331132302012223-2300102101131321-0331011120011100-0333133222330022-3230313123210312-3223122220231033-1001232310132002-2311120222313022): complete subsection reference.

<a id="canonical-0321010233331311-2133130232102222-0203103012203023-3201213101311020-0102030113323002-1203103031003110-0122133110322323-3203013112130002"></a>

## Next pages — static_routes / 330022103333 / 7

- [custom_network_config.sli_config.static_routes.static_routes.default_gateway](resources--voltstack_site--reference--group-005.md#canonical-0121133032222000-1131203022200023-2010213302300220-2102033103212112-3202102010123002-1003020300110022-3330000020031232-0031131012223302)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-2331132302012223-2300102101131321-0331011120011100-0333133222330022-3230313123210312-3223122220231033-1001232310132002-2311120222313022)
- [custom_network_config.sli_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3010201321130121-1003131132022102-1321123300131211-1311222020210013-1210221231201101-3211022010320113-2202213020220333-3232132331233201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0121133032222000-1131203022200023-2010213302300220-2102033103212112-3202102010123002-1003020300110022-3330000020031232-0031131012223302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212022223120113-3312212002221212-2103201130121213-1021103033221120-0320030300321121-2031212000320331-0312033003321200-2123303310103132"></a>

## custom_network_config.sli_config.static_routes.static_routes.default_gateway — default_gateway / 303312021011 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- [custom_network_config.sli_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3010201321130121-1003131132022102-1321123300131211-1311222020210013-1210221231201101-3211022010320113-2202213020220333-3232132331233201)
- [custom_network_config.sli_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3131302122323111-0022332121303122-1031021121211112-3113303020123222-2212011023100203-3010220013323131-0102320221020032-0112300221321101)
- custom_network_config.sli_config.static_routes.static_routes.default_gateway

<a id="canonical-2021321211211012-1222011311123123-3213231222132003-2101333020301112-1310311222300102-3332132221111121-1212302302002300-3310101103330201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

<a id="canonical-2121332211032322-2202023100213322-3003031301301110-2220232313322131-1223102203202230-3030131301112202-3212222031300012-3103203230030132"></a>

## Direct properties — default_gateway / 303312021011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220032021003112-0030322302333031-1220332101202211-1322310332330131-3130033212013023-0330302123201201-3211203111323120-1220013303323131"></a>

## Next pages — default_gateway / 303312021011 / 4

- [custom_network_config.sli_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3131302122323111-0022332121303122-1031021121211112-3113303020123222-2212011023100203-3010220013323131-0102320221020032-0112300221321101)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2331132302012223-2300102101131321-0331011120011100-0333133222330022-3230313123210312-3223122220231033-1001232310132002-2311120222313022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000132310000303-1001200200322102-2103131020001020-1032221033111220-0110100311101332-1212000232010331-1110310212203021-0020331232213121"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface — node_interface / 131020200121 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- [custom_network_config.sli_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3010201321130121-1003131132022102-1321123300131211-1311222020210013-1210221231201101-3211022010320113-2202213020220333-3232132331233201)
- [custom_network_config.sli_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3131302122323111-0022332121303122-1031021121211112-3113303020123222-2212011023100203-3010220013323131-0102320221020032-0112300221321101)
- custom_network_config.sli_config.static_routes.static_routes.node_interface

<a id="canonical-1000013020322000-1112022200230230-1221103311101021-1333233113011112-0013222111022200-1210100233032000-3332030231202203-3021330313103032"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

Receipt-pinned upstream constraints:

```json
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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222020312211320-3103223301110100-3020003130320211-2021031132103230-3111212323330003-2022332200311212-3120030113010322-3300033232210030"></a>

## Direct properties — node_interface / 131020200121 / 3

- [list](resources--voltstack_site--reference--group-005.md#canonical-1232002103313210-1200101221200330-0330100300303111-2110010233111213-0313101220111222-1203032332311310-3221102021001021-3002023201111111): complete subsection reference.

<a id="canonical-2303131010003110-2303002122233321-2303230230210322-3331131003122232-1222303210021002-1311310202012011-0300113131321330-0213002303013221"></a>

## Next pages — node_interface / 131020200121 / 4

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-1232002103313210-1200101221200330-0330100300303111-2110010233111213-0313101220111222-1203032332311310-3221102021001021-3002023201111111)
- [custom_network_config.sli_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3131302122323111-0022332121303122-1031021121211112-3113303020123222-2212011023100203-3010220013323131-0102320221020032-0112300221321101)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1232002103313210-1200101221200330-0330100300303111-2110010233111213-0313101220111222-1203032332311310-3221102021001021-3002023201111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032222231002010-1103130330303222-0020323321030102-3220213013022021-2030323213131332-1321021323012200-2323111231310013-0330032202013223"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface.list — list / 320202313101 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- [custom_network_config.sli_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3010201321130121-1003131132022102-1321123300131211-1311222020210013-1210221231201101-3211022010320113-2202213020220333-3232132331233201)
- [custom_network_config.sli_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3131302122323111-0022332121303122-1031021121211112-3113303020123222-2212011023100203-3010220013323131-0102320221020032-0112300221321101)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-2331132302012223-2300102101131321-0331011120011100-0333133222330022-3230313123210312-3223122220231033-1001232310132002-2311120222313022)
- custom_network_config.sli_config.static_routes.static_routes.node_interface.list

<a id="canonical-0031101330133032-3320021101233313-2030011032012101-0201112300231222-0100030033033202-1313232312313002-0301123133312210-0231223021101232"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213122212230122-0123032232321131-1333031102320221-1331000210003113-3131100223223310-0300312130212112-3313013323332302-3313222233320013"></a>

## Direct properties — list / 320202313101 / 3

- [interface](resources--voltstack_site--reference--group-005.md#canonical-0310002211133323-3000320110330223-0222212201021200-1011201103222313-2031332333011010-3311220333201031-1023113013002120-2101311133002201): complete subsection reference.

<a id="canonical-0222102333211022-2330023031032022-0213020222011123-3321022301332233-1202322020302132-2201133001330211-2332120310110220-1011333112212022"></a>

<a id="canonical-1201303123031232-2221023120200013-2333113233120313-2211301331120111-2213200121201011-0100112131232210-3330031331013322-1202021222323030"></a>

## node property — list / 320202313101 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-3310101001330023-2000330210013030-0133223001223110-0113013122301023-2132022132310102-1203302110123113-0023333232020321-2000022112231233"></a>

## Next pages — list / 320202313101 / 5

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface](resources--voltstack_site--reference--group-005.md#canonical-0310002211133323-3000320110330223-0222212201021200-1011201103222313-2031332333011010-3311220333201031-1023113013002120-2101311133002201)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-2331132302012223-2300102101131321-0331011120011100-0333133222330022-3230313123210312-3223122220231033-1001232310132002-2311120222313022)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0310002211133323-3000320110330223-0222212201021200-1011201103222313-2031332333011010-3311220333201031-1023113013002120-2101311133002201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230100103200310-0210323310030230-1013320103022133-2212011032221330-1321300033302032-2002002201212300-2331320323302220-3210212133002331"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface — interface / 313220203023 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- [custom_network_config.sli_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3010201321130121-1003131132022102-1321123300131211-1311222020210013-1210221231201101-3211022010320113-2202213020220333-3232132331233201)
- [custom_network_config.sli_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3131302122323111-0022332121303122-1031021121211112-3113303020123222-2212011023100203-3010220013323131-0102320221020032-0112300221321101)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-2331132302012223-2300102101131321-0331011120011100-0333133222330022-3230313123210312-3223122220231033-1001232310132002-2311120222313022)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-1232002103313210-1200101221200330-0330100300303111-2110010233111213-0313101220111222-1203032332311310-3221102021001021-3002023201111111)
- custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-3223223311211132-1331111002100333-0020213213320023-2101010122201021-2021322311123121-1333102222233021-2223212103311300-0302321330102011"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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

<a id="canonical-1220220332323131-0130330332033302-2001233100323203-2122231123220101-0003201022200233-3202223023101103-2013111111133120-1013332002210211"></a>

## Direct properties — interface / 313220203023 / 3

<a id="canonical-0232003131211120-0001200012133323-0303201123320003-1301311123310301-2103312321310202-1112101312323203-2100220122232330-1022313123132002"></a>

<a id="canonical-1320221223201123-1103111313032220-3221003202322230-3133000200210120-0311000033313112-0020301101101111-3121001123002302-0031111021132103"></a>

## kind property — interface / 313220203023 / 4

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

<a id="canonical-2120130210202322-3100022320102330-1331312222023323-2033100101030100-1130211200302332-3210112020113331-0201010010311010-2233113233021010"></a>

<a id="canonical-1231021320321013-1213120320123320-2133023133021230-2022020103201112-2301123122303112-1322200211302312-2010020021013112-3122310131120313"></a>

## name property — interface / 313220203023 / 5

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

<a id="canonical-1331110123003021-0332311302113330-3111032132320210-0203122020002031-3132202010230010-1103112003312123-2212331303010012-3311001223103012"></a>

<a id="canonical-2322130122333313-2301302100103030-0332133220300103-1320112013002130-0100331022203323-2132123002310201-1211211010001331-0333310012332012"></a>

## namespace property — interface / 313220203023 / 6

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

<a id="canonical-0231013200013003-2231123010032330-1001003113113003-2131212131102200-2022031003002220-3221122030132100-3000322302011021-2001233101223221"></a>

<a id="canonical-1002133120200202-3112031302113231-0230023301101032-2201223212202123-1102312112130202-3313201032202220-2113213033301002-2122330111021132"></a>

## tenant property — interface / 313220203023 / 7

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

<a id="canonical-3123320111101213-3102133002031032-1202223212302303-2201003120320300-3021022010223313-2310210231122303-1122222133002233-1013223122202020"></a>

<a id="canonical-1002130102131023-3230012322133103-1202232033122323-0132203012022313-2232010212122233-0032232011012122-0223131103201331-0312310201131323"></a>

## uid property — interface / 313220203023 / 8

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

<a id="canonical-3030021133102023-1332323333301203-0333233030330133-1002202112011223-3202300011232220-2103130131233300-1213202323331331-1203001020113020"></a>

## Next pages — interface / 313220203023 / 9

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-1232002103313210-1200101221200330-0330100300303111-2110010233111213-0313101220111222-1203032332311310-3221102021001021-3002023201111111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3303202101313311-2332103211112231-1223201102030123-1322100301333302-2333211100213230-2222122132030211-0311221100323131-2200112333120203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323223203130131-1313133110022130-0133232010123103-3301010021230202-1333123120202010-3311220310300103-2122013103323211-3323112203203132"></a>

## custom_network_config.sli_config.static_v6_routes — static_v6_routes / 022203132232 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- custom_network_config.sli_config.static_v6_routes

<a id="canonical-1130220031030003-0332003013330221-1310211121201121-3221222110232110-2312230301002232-1233011002030320-0300033232220221-2232310230313120"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_v6_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2320203303221322-3222001320032303-1130323021211030-1110132103021003-0021101210232131-0012233232102223-3033213100300131-2331110000221231"></a>

## Direct properties — static_v6_routes / 022203132232 / 3

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-0020333010323222-3112120021020001-1301000233112123-1101331330133020-0332002001213110-1030123223001013-0033231233011031-2021311003300012): complete subsection reference.

<a id="canonical-1231001303202323-1101330000232020-2310332213220220-0020122201031303-2002313312100100-3003212001030123-2100201210221012-1131113112112012"></a>

## Next pages — static_v6_routes / 022203132232 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0020333010323222-3112120021020001-1301000233112123-1101331330133020-0332002001213110-1030123223001013-0033231233011031-2021311003300012)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0020333010323222-3112120021020001-1301000233112123-1101331330133020-0332002001213110-1030123223001013-0033231233011031-2021311003300012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103322101020323-3002322131233003-3000100231303011-2320220101230203-1031033110101013-0313020100232300-3021332102323210-0213022130033323"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes — static_routes / 010213213233 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-3303202101313311-2332103211112231-1223201102030123-1322100301333302-2333211100213230-2222122132030211-0311221100323131-2200112333120203)
- custom_network_config.sli_config.static_v6_routes.static_routes

<a id="canonical-1111220103212222-3033122211300310-2211321121311131-0011122033123202-2122210102031110-2123101111033302-3001321010211021-0221300322011333"></a>

Type: `"object"`. list nested block, Optional.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122133031121202-1100100102031332-0103210321033203-0312202300221303-2320123331221133-0103030232020301-1312131130122311-1101330132220321"></a>

## Direct properties — static_routes / 010213213233 / 3

<a id="canonical-0133301323222220-3032320112211010-1113210202330123-0332013330303100-0230210220133000-0010122123002233-1230032023332222-0120322023322102"></a>

<a id="canonical-1133020031202132-2101233001230000-2030312123130101-1112201113333321-0223011131302110-0021330200112311-3310112322110031-3022132223332200"></a>

## attrs property — static_routes / 010213213233 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--voltstack_site--reference--group-005.md#canonical-0020222221123002-0120322213020033-0000000201132210-1010201103012333-3112120231313001-0100132311021232-2131311212311230-2323120132223233): complete subsection reference.

<a id="canonical-3203320000332333-3220321313232121-3031030102233313-1200110110330012-1221223003123120-2222000203101313-2121321131303321-3223120002232300"></a>

<a id="canonical-0210102111131212-1031333010212300-0023201120031012-1032033030020212-3221111320110110-0330113233201200-3332023031321332-1023200211332233"></a>

## ip_address property — static_routes / 010213213233 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-2003330102201130-1332202213230103-2312222233111021-0203100202331022-1330230232023303-2113301223223012-0133103330021123-3311312101031302"></a>

<a id="canonical-3122323333330020-0021303111210002-1233112032000110-1222210222300313-1320200213131121-3332133313003320-3333010002132213-0103121231303302"></a>

## ip_prefixes property — static_routes / 010213213233 / 6

Type: `["list", "string"]`. Optional.

List of IPv6 route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--voltstack_site--reference--group-005.md#canonical-0021021030111233-3211302212333230-1211121002303111-2013133113220330-3321301030233110-0233020310033021-0311333313123232-3322300222131232): complete subsection reference.

<a id="canonical-0130213010112220-3012112232022231-3320212312101320-0311302103112331-3130223210310131-2021032101002123-1311301112323210-0102111133201132"></a>

## Next pages — static_routes / 010213213233 / 7

- [custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway](resources--voltstack_site--reference--group-005.md#canonical-0020222221123002-0120322213020033-0000000201132210-1010201103012333-3112120231313001-0100132311021232-2131311212311230-2323120132223233)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-0021021030111233-3211302212333230-1211121002303111-2013133113220330-3321301030233110-0233020310033021-0311333313123232-3322300222131232)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-3303202101313311-2332103211112231-1223201102030123-1322100301333302-2333211100213230-2222122132030211-0311221100323131-2200112333120203)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0020222221123002-0120322213020033-0000000201132210-1010201103012333-3112120231313001-0100132311021232-2131311212311230-2323120132223233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101301220032132-0210033213011023-2201223020020012-0021020003302133-0321203312010303-0233003000132231-2330321131111320-1231312002313202"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway — default_gateway / 001010202322 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-3303202101313311-2332103211112231-1223201102030123-1322100301333302-2333211100213230-2222122132030211-0311221100323131-2200112333120203)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0020333010323222-3112120021020001-1301000233112123-1101331330133020-0332002001213110-1030123223001013-0033231233011031-2021311003300012)
- custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-3312312100103122-3303302121112300-2202200021333131-3112113303011113-0130332202030023-2201232312121303-1113330310113013-3011013331133033"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

<a id="canonical-1122132021111331-1001103202322200-1022033030002302-0133222002310221-2131213000033300-2013200113003013-1031031311023312-2321030323002320"></a>

## Direct properties — default_gateway / 001010202322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222002303103112-3120233232210121-2122231003001212-1232301312003111-2210321133312201-0021302011301003-1302133002121330-3332123002023333"></a>

## Next pages — default_gateway / 001010202322 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0020333010323222-3112120021020001-1301000233112123-1101331330133020-0332002001213110-1030123223001013-0033231233011031-2021311003300012)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0021021030111233-3211302212333230-1211121002303111-2013133113220330-3321301030233110-0233020310033021-0311333313123232-3322300222131232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203023033010023-2230331322213021-0322113211212122-2323302211102230-0112011203111111-1201022102030232-0122233001003013-2022002303111010"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface — node_interface / 130320012220 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-3303202101313311-2332103211112231-1223201102030123-1322100301333302-2333211100213230-2222122132030211-0311221100323131-2200112333120203)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0020333010323222-3112120021020001-1301000233112123-1101331330133020-0332002001213110-1030123223001013-0033231233011031-2021311003300012)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface

<a id="canonical-0211033330201213-1203130333013031-1020033301331321-3312303323122101-2110001100011213-0200133211023012-2300120131122232-2013230220331202"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

Receipt-pinned upstream constraints:

```json
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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033021321112112-3112200211332103-2120002321232323-3200232021032102-0112230331303000-2021133030310130-0301110233311030-2111100103112013"></a>

## Direct properties — node_interface / 130320012220 / 3

- [list](resources--voltstack_site--reference--group-005.md#canonical-1201322100232212-2332211112120120-3031130021221010-2202321001302323-0213011211223110-0312021120002312-2112220221021322-2013320021210001): complete subsection reference.

<a id="canonical-3100012232030321-0002330321103302-2221313111201200-3202020033103133-1320202322031301-2131220113102123-1130221001233030-0002202231112112"></a>

## Next pages — node_interface / 130320012220 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-1201322100232212-2332211112120120-3031130021221010-2202321001302323-0213011211223110-0312021120002312-2112220221021322-2013320021210001)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0020333010323222-3112120021020001-1301000233112123-1101331330133020-0332002001213110-1030123223001013-0033231233011031-2021311003300012)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1201322100232212-2332211112120120-3031130021221010-2202321001302323-0213011211223110-0312021120002312-2112220221021322-2013320021210001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020012200230130-1120121332311231-2213212013323010-3312201321123111-2102201301321113-2313102011210021-2223013212011001-3332131211121010"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list — list / 110011023333 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-3303202101313311-2332103211112231-1223201102030123-1322100301333302-2333211100213230-2222122132030211-0311221100323131-2200112333120203)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0020333010323222-3112120021020001-1301000233112123-1101331330133020-0332002001213110-1030123223001013-0033231233011031-2021311003300012)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-0021021030111233-3211302212333230-1211121002303111-2013133113220330-3321301030233110-0233020310033021-0311333313123232-3322300222131232)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-2100022030020113-1301311302310200-1201222322001122-0202232001123200-0203211131220022-0112033120203023-3021123031123233-1122111032321323"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303013121302213-1321132012222011-2310113233012232-3200022210102112-3212123213313020-2101010132033311-1000002200110023-3030030000112002"></a>

## Direct properties — list / 110011023333 / 3

- [interface](resources--voltstack_site--reference--group-005.md#canonical-1021012303003120-0131320332232202-2220012002020201-1133103302200022-3130230012332211-2121111020332033-1120013103023232-2300021200011200): complete subsection reference.

<a id="canonical-2120233312011211-0110210333331222-1103310311232321-0110330311332211-0301230031112102-2000120022211330-2032300212030233-1332122312231333"></a>

<a id="canonical-3133121330321133-0222333123030120-2120222030103101-1312031210030023-2010132311033300-1202201303011312-3322303021032011-1232230320323212"></a>

## node property — list / 110011023333 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-3103312001222023-1122333123021013-3112203211211020-2110211132320311-3330102310022103-0003311002021201-2313001331232021-1221232021313213"></a>

## Next pages — list / 110011023333 / 5

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface](resources--voltstack_site--reference--group-005.md#canonical-1021012303003120-0131320332232202-2220012002020201-1133103302200022-3130230012332211-2121111020332033-1120013103023232-2300021200011200)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-0021021030111233-3211302212333230-1211121002303111-2013133113220330-3321301030233110-0233020310033021-0311333313123232-3322300222131232)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1021012303003120-0131320332232202-2220012002020201-1133103302200022-3130230012332211-2121111020332033-1120013103023232-2300021200011200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133230000030210-1012111331001301-0113313010010311-3030000022310021-3310131130201003-2033130102020010-1301000011210323-3123031100020023"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface — interface / 113100211333 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.sli_config](resources--voltstack_site--reference--group-005.md#canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-3303202101313311-2332103211112231-1223201102030123-1322100301333302-2333211100213230-2222122132030211-0311221100323131-2200112333120203)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0020333010323222-3112120021020001-1301000233112123-1101331330133020-0332002001213110-1030123223001013-0033231233011031-2021311003300012)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-0021021030111233-3211302212333230-1211121002303111-2013133113220330-3321301030233110-0233020310033021-0311333313123232-3322300222131232)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-1201322100232212-2332211112120120-3031130021221010-2202321001302323-0213011211223110-0312021120002312-2112220221021322-2013320021210001)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-1131211302230020-2020303303032321-1223121313131322-0103333032003133-1333122223212020-2132030321020310-1030003231203223-0332120102321100"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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

<a id="canonical-1201313223111301-2130200321003033-3001320023321210-0033222100110312-0323313011133203-0002020101022122-1012223112333333-1302311213302022"></a>

## Direct properties — interface / 113100211333 / 3

<a id="canonical-3011212020311123-1002020321332303-0001222012112332-1302000120130121-2113321220303123-2300131130032121-0311031333222321-3020022230010202"></a>

<a id="canonical-3100212302113223-2220130001301021-0100220023202001-2330110222212202-0123203133120210-3113120331012103-0323230103223310-0232020212213321"></a>

## kind property — interface / 113100211333 / 4

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

<a id="canonical-1223121200122310-2000201013110130-0121132301132011-1111011321210311-3200310303032023-2220100112101032-3301213130011303-1123122233121122"></a>

<a id="canonical-0333322322330222-2223310002010301-1333230323030030-2212212312210323-3132031222303000-1213200101032110-2301120113321113-0001302100031023"></a>

## name property — interface / 113100211333 / 5

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

<a id="canonical-1323023311212130-1113332212011211-1332233020231013-1212000111311130-3122232103201222-1121023310132333-3032311311010212-3033131121312233"></a>

<a id="canonical-0233212300213013-1003323201330101-1312303122212020-0102331210011303-0112220301012122-1133122222230320-3213033220001321-1021020021203202"></a>

## namespace property — interface / 113100211333 / 6

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

<a id="canonical-2000232121033333-3131133022222031-1320132210313321-0001302131122120-1111311132131210-3221211223122232-3100103123112210-1101103213032333"></a>

<a id="canonical-2320200101321201-2001101010023032-3232002303120123-0320220201311332-0311123231011221-0103312223200110-2121333330213203-0001112120033022"></a>

## tenant property — interface / 113100211333 / 7

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

<a id="canonical-0121122330130330-2030110201213211-2003320211121202-1202213101300123-3222212130330330-2330221312310222-3030123302300300-3303031102010022"></a>

<a id="canonical-1101301333001333-0123123020002203-0120112120322312-3221100110120020-0101012202202312-0112113220001220-1320001223333103-1112213121222300"></a>

## uid property — interface / 113100211333 / 8

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

<a id="canonical-1333102220132030-0120320111323001-1110131320123012-0110223232022203-1322322211133213-0223222321121333-0233002011220003-1203212110221003"></a>

## Next pages — interface / 113100211333 / 9

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-1201322100232212-2332211112120120-3031130021221010-2202321001302323-0213011211223110-0312021120002312-2112220221021322-2013320021210001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102001311110100-0300000031020131-2201310213312010-3111322100233033-0120322312300103-1102132010010323-3120221210220203-2302022310323311"></a>

## custom_network_config.slo_config — slo_config / 312332131220 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- custom_network_config.slo_config

<a id="canonical-0220330220202313-2332133003031313-0023111102132122-0201321312331220-0103212031330223-1120111012220013-3323123222202012-1320312122102100"></a>

Type: `"object"`. single nested block, Optional.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_static_v6_routes",
    "static_v6_routes")}
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
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_static_v6_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
slo_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203122000020212-3003231330332012-1011033231302220-1012032331213110-1320033133110221-2332133210321031-3011200300013033-1333022020332113"></a>

## Direct properties — slo_config / 312332131220 / 3

- [dc_cluster_group](resources--voltstack_site--reference--group-005.md#canonical-2131302313022303-2213121110223333-0203133211211100-2311322201121312-1112333030111120-0001301212322221-2020110311311121-3130020330010302): complete subsection reference.

- [labels](resources--voltstack_site--reference--group-005.md#canonical-3010111313021021-0123120213200031-0023023303322303-1110003213133110-0101210020320223-0113232221300222-2120330322001101-2313210232131321): complete subsection reference.

- [no_dc_cluster_group](resources--voltstack_site--reference--group-005.md#canonical-3130200111333330-0303110300230033-1013020200101122-0021232022031122-2100213210311332-0032210133302310-2302320101011022-0031333123023011): complete subsection reference.

- [no_static_routes](resources--voltstack_site--reference--group-005.md#canonical-0130213112111003-1133111311110122-1023120321103131-2322320003023021-2213102333223230-3013231321301331-3113213112021122-3302033000203012): complete subsection reference.

- [no_static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-3310000003111133-2310132332112313-3303330312233002-1212323003023131-2310203130212022-2132112202212211-0213213103300212-1013021130221321): complete subsection reference.

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-1222233330322332-1311022211003101-0230120133122311-3030023301210101-0021032123033010-0203312312123023-3130210113323311-3100333122023020): complete subsection reference.

- [static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-1003112130012011-2001132133020101-3313022022312333-1303310032020322-1022302212213212-2120201321002113-1002101031021031-0101120022113123): complete subsection reference.

<a id="canonical-3232232011322100-0130102220131210-3130111100213103-3331000133311110-3101210100230020-2032030002022102-3123303003012220-3232110000133320"></a>

## Next pages — slo_config / 312332131220 / 4

- [custom_network_config.slo_config.dc_cluster_group](resources--voltstack_site--reference--group-005.md#canonical-2131302313022303-2213121110223333-0203133211211100-2311322201121312-1112333030111120-0001301212322221-2020110311311121-3130020330010302)
- [custom_network_config.slo_config.labels](resources--voltstack_site--reference--group-005.md#canonical-3010111313021021-0123120213200031-0023023303322303-1110003213133110-0101210020320223-0113232221300222-2120330322001101-2313210232131321)
- [custom_network_config.slo_config.no_dc_cluster_group](resources--voltstack_site--reference--group-005.md#canonical-3130200111333330-0303110300230033-1013020200101122-0021232022031122-2100213210311332-0032210133302310-2302320101011022-0031333123023011)
- [custom_network_config.slo_config.no_static_routes](resources--voltstack_site--reference--group-005.md#canonical-0130213112111003-1133111311110122-1023120321103131-2322320003023021-2213102333223230-3013231321301331-3113213112021122-3302033000203012)
- [custom_network_config.slo_config.no_static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-3310000003111133-2310132332112313-3303330312233002-1212323003023131-2310203130212022-2132112202212211-0213213103300212-1013021130221321)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-1222233330322332-1311022211003101-0230120133122311-3030023301210101-0021032123033010-0203312312123023-3130210113323311-3100333122023020)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-1003112130012011-2001132133020101-3313022022312333-1303310032020322-1022302212213212-2120201321002113-1002101031021031-0101120022113123)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2131302313022303-2213121110223333-0203133211211100-2311322201121312-1112333030111120-0001301212322221-2020110311311121-3130020330010302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131132312111003-3313031010131233-3020112311323012-2111020321231023-1110332222002101-0331322031030030-2232012133310221-3321211133311322"></a>

## custom_network_config.slo_config.dc_cluster_group — dc_cluster_group / 033223103320 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- custom_network_config.slo_config.dc_cluster_group

<a id="canonical-1321223320322211-3210030233222231-2302130201230122-2001210223100302-3313311030131103-0330101000311032-2012010220002001-0023311120303311"></a>

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

<a id="canonical-0103200010330221-1312113032100022-1220213010230303-1210013201013002-1111020312122233-2003203332201231-2223301303331331-0233123030332223"></a>

## Direct properties — dc_cluster_group / 033223103320 / 3

<a id="canonical-3333020213123320-1212123033130112-2000123301303112-2131310310133101-1002201003131110-1020120101112123-3120012011303333-1330213331031331"></a>

<a id="canonical-2023310231100322-2211300110312311-3013300002202121-3133100003031022-3013133000332130-3033123300201220-2222030112211231-1101331023331123"></a>

## name property — dc_cluster_group / 033223103320 / 4

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

<a id="canonical-1200022002033100-2023202033003113-3302013301031201-3000221303130022-0323022123000003-0321101330301220-0202133021032210-1102210010032131"></a>

<a id="canonical-1132003313130013-1302001310213000-3310031022321203-1010132031230230-2021011121203112-1210013030202110-1112231133112010-0123103020123212"></a>

## namespace property — dc_cluster_group / 033223103320 / 5

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

<a id="canonical-0112022000020030-2003012321123322-2002113310221321-2100133232113333-2103120001303121-3301122032120322-0020332303123111-1212120321223002"></a>

<a id="canonical-2312131210113011-2232133232130022-2030232003132002-3302302223030312-3212313302233011-2203302120030131-1001103021110033-3003102013031322"></a>

## tenant property — dc_cluster_group / 033223103320 / 6

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

<a id="canonical-3320312321211002-3131230202012103-2223330112232112-2110111002233122-0110202332020010-2010210321233131-1322020212221111-0303201231311232"></a>

## Next pages — dc_cluster_group / 033223103320 / 7

- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3010111313021021-0123120213200031-0023023303322303-1110003213133110-0101210020320223-0113232221300222-2120330322001101-2313210232131321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203031212200202-0331030023311321-3113201330123100-3000002022232130-0100212301310131-1221201200130020-2302032030311333-3020201013020202"></a>

## custom_network_config.slo_config.labels — labels / 032203101100 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- custom_network_config.slo_config.labels

<a id="canonical-1200121131323201-1210132330213220-1000331323102220-3311222231131022-0130313201332023-0303123123032102-3102001130003023-0323312131003100"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this network, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-0121130223213313-1011200230221311-3221112130203113-2010012213103210-3002230112233020-0233003301132231-2111133212113331-0012020112201223"></a>

## Direct properties — labels / 032203101100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233332330123011-2322123222111203-1201113120210321-1102123020213310-2321133333320030-0301132013200101-2221213110231131-2333313120221323"></a>

## Next pages — labels / 032203101100 / 4

- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3130200111333330-0303110300230033-1013020200101122-0021232022031122-2100213210311332-0032210133302310-2302320101011022-0031333123023011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133210112320210-3312121031313222-3002322203332002-2233220102321030-1113300230320121-2310112200223130-0103123010231132-0221213121020322"></a>

## custom_network_config.slo_config.no_dc_cluster_group — no_dc_cluster_group / 002311210331 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- custom_network_config.slo_config.no_dc_cluster_group

<a id="canonical-3002212103113322-2002003101321122-1210331030112303-0003232002002223-1110021030331302-3201031202333001-0332221322120332-3210131223220332"></a>

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

<a id="canonical-2122121233123202-3312330133321001-3021111313030212-0202001221101123-1131132033223011-1202331020223122-3022102332332102-1210231320212013"></a>

## Direct properties — no_dc_cluster_group / 002311210331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013330123212103-1023312102311233-0212030313233031-0020132201112201-3311131303032300-2330223132011123-3132302212010211-2011123202212121"></a>

## Next pages — no_dc_cluster_group / 002311210331 / 4

- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0130213112111003-1133111311110122-1023120321103131-2322320003023021-2213102333223230-3013231321301331-3113213112021122-3302033000203012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100302032113311-3223303103323202-0013220301131330-0333232022320110-1132102223223132-3300003030023032-2220201233120302-1133211010121102"></a>

## custom_network_config.slo_config.no_static_routes — no_static_routes / 213123033313 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- custom_network_config.slo_config.no_static_routes

<a id="canonical-1230000100222033-2223102211113331-1332122210310232-1233322023211203-0011321131111222-3301131103102010-2032233022312111-1223102312302102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static routes.

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
no_static_routes = {}
```

<a id="canonical-2323001321210013-0232220203010322-3123331132222121-2313103321121232-1103033300311103-2032110210131332-3123012200001222-3231232202102300"></a>

## Direct properties — no_static_routes / 213123033313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110122033000101-3003031123311320-0102011223310102-0201232110231133-3330303020321322-2332001311200113-1200201331132220-3312302333213013"></a>

## Next pages — no_static_routes / 213123033313 / 4

- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3310000003111133-2310132332112313-3303330312233002-1212323003023131-2310203130212022-2132112202212211-0213213103300212-1013021130221321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220031331210213-1023031322110011-0200323311230132-2323231313301232-2010321113100022-0301303111002323-2220122333221013-1313001021113313"></a>

## custom_network_config.slo_config.no_static_v6_routes — no_static_v6_routes / 102001220030 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- custom_network_config.slo_config.no_static_v6_routes

<a id="canonical-3200230013221333-3033321220200210-1131233130310131-0011322232221203-0322302033312003-1222310032303122-2111202121301132-3213030220323221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static v6 routes.

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
no_static_v6_routes = {}
```

<a id="canonical-1021020000000001-1330010110221321-0301321300203323-1113323201000302-1020312330313212-3002032231113001-1031200110102031-0222122230013033"></a>

## Direct properties — no_static_v6_routes / 102001220030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112311033200120-2011021033221322-1031312333333132-2331232231210121-0121313232100102-0332221023002113-0012221121122133-1203000211132003"></a>

## Next pages — no_static_v6_routes / 102001220030 / 4

- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1222233330322332-1311022211003101-0230120133122311-3030023301210101-0021032123033010-0203312312123023-3130210113323311-3100333122023020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201323313132021-1102223310303033-0212310133002301-0232110211231221-1020330101030103-0111102013302000-2303223033001222-2223110220332232"></a>

## custom_network_config.slo_config.static_routes — static_routes / 212201103200 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- custom_network_config.slo_config.static_routes

<a id="canonical-0330231103203032-0032313201120303-0303002102312102-0300103131323031-3211310202223110-2131222220200221-3300122031321100-3313112013222223"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123102211121123-1100110020333333-1323232021231110-2002300120322202-0020001201101221-3013120130003012-3133021112200300-0100212332313120"></a>

## Direct properties — static_routes / 212201103200 / 3

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-2102220100231101-2233321202110130-3131001011222331-1222110220020312-3211102102212312-0130031110002102-3022120103111002-3100322213130101): complete subsection reference.

<a id="canonical-2200001231323200-2210113203310120-0322021221013110-0021202030110031-3230020103203332-3232031023321232-2111222312221112-3033301202010103"></a>

## Next pages — static_routes / 212201103200 / 4

- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-2102220100231101-2233321202110130-3131001011222331-1222110220020312-3211102102212312-0130031110002102-3022120103111002-3100322213130101)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2102220100231101-2233321202110130-3131001011222331-1222110220020312-3211102102212312-0130031110002102-3022120103111002-3100322213130101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022331320023332-1013313201231213-1033302012120231-3021203320301202-2021102020123100-0001221012322103-3301323002321213-0120023131230012"></a>

## custom_network_config.slo_config.static_routes.static_routes — static_routes / 302103212103 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-1222233330322332-1311022211003101-0230120133122311-3030023301210101-0021032123033010-0203312312123023-3130210113323311-3100333122023020)
- custom_network_config.slo_config.static_routes.static_routes

<a id="canonical-1113123002231301-3321211131012311-3302020112303330-3333000012131212-0332022210011113-2222031303233112-1300232332222033-2203303202313101"></a>

Type: `"object"`. list nested block, Optional.

Static Routes. List of static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333121233302332-0321211110201233-3330301210230132-0203103331303021-1012330332122303-0301132332103003-2112100312111233-3021002322031220"></a>

## Direct properties — static_routes / 302103212103 / 3

<a id="canonical-2223002313010233-3110232311202233-2200003333032101-0031001133322202-3121301233221331-0223311031121023-2300322121120111-3301230310212231"></a>

<a id="canonical-3020021210230000-0211313010111210-3022021011022032-0321211232222032-0121332232222103-0003333013113122-0301221133231200-2210322030211120"></a>

## attrs property — static_routes / 302103212103 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--voltstack_site--reference--group-005.md#canonical-3332102200203231-0130312010330110-1000212230002300-2033030022111312-2112023211310310-3101211022302110-3030123333133331-1011121132133233): complete subsection reference.

<a id="canonical-1231010303032332-3322223123330121-1032331230203132-0002233131120102-1230310223220112-2233233030232232-0103322132131121-0023333003312221"></a>

<a id="canonical-3303100031230130-1203030321302300-1101102010032122-0332112131320231-3133310120231211-2323130222101110-2131203030223221-1022303033311333"></a>

## ip_address property — static_routes / 302103212103 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-1331212322212022-1121013230132222-2201031100112230-2022011311120021-2022132301220201-0031002333221113-0201330123110301-3313303031012313"></a>

<a id="canonical-1211133130301302-2210333001312310-0102030332132132-0000103121030120-1301321032332302-3323011032010031-3101201003113213-2330222022212312"></a>

## ip_prefixes property — static_routes / 302103212103 / 6

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--voltstack_site--reference--group-005.md#canonical-2200221310302012-3322200200333010-3030232302020222-2302033101133001-0202223003010333-3223021332310122-3311330231012222-1133303221012201): complete subsection reference.

<a id="canonical-1123212320312322-3122332030320311-1021033320033011-1100001120013202-0111120212211310-1122221210132022-3123333302303011-1200331312112003"></a>

## Next pages — static_routes / 302103212103 / 7

- [custom_network_config.slo_config.static_routes.static_routes.default_gateway](resources--voltstack_site--reference--group-005.md#canonical-3332102200203231-0130312010330110-1000212230002300-2033030022111312-2112023211310310-3101211022302110-3030123333133331-1011121132133233)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-2200221310302012-3322200200333010-3030232302020222-2302033101133001-0202223003010333-3223021332310122-3311330231012222-1133303221012201)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-1222233330322332-1311022211003101-0230120133122311-3030023301210101-0021032123033010-0203312312123023-3130210113323311-3100333122023020)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3332102200203231-0130312010330110-1000212230002300-2033030022111312-2112023211310310-3101211022302110-3030123333133331-1011121132133233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020113123110220-1001103331202201-3131031333232213-0322202112001210-0132121011221031-0223121100023031-3130231222220022-3023200101212212"></a>

## custom_network_config.slo_config.static_routes.static_routes.default_gateway — default_gateway / 132233202112 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-1222233330322332-1311022211003101-0230120133122311-3030023301210101-0021032123033010-0203312312123023-3130210113323311-3100333122023020)
- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-2102220100231101-2233321202110130-3131001011222331-1222110220020312-3211102102212312-0130031110002102-3022120103111002-3100322213130101)
- custom_network_config.slo_config.static_routes.static_routes.default_gateway

<a id="canonical-0131013133302220-1021132220223012-1101210223210231-3220102222202100-0003203223133233-2030300120200133-3022132131201023-1201011132312003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

<a id="canonical-2110000023200211-0331030221023022-2230321031221031-2232331033320023-0210033130301331-0321230101013333-0313102222112211-2323121223213321"></a>

## Direct properties — default_gateway / 132233202112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332223231110330-1101133220302203-3331010001101331-2120001313110211-0112131012310100-1010020212033212-1230003111220110-1112321330230032"></a>

## Next pages — default_gateway / 132233202112 / 4

- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-2102220100231101-2233321202110130-3131001011222331-1222110220020312-3211102102212312-0130031110002102-3022120103111002-3100322213130101)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2200221310302012-3322200200333010-3030232302020222-2302033101133001-0202223003010333-3223021332310122-3311330231012222-1133303221012201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002121233100311-3313032001232201-2331131032210303-1120031023212232-2110013331230131-1000021311000031-2022333320330333-0100012012213033"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface — node_interface / 201021322002 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-1222233330322332-1311022211003101-0230120133122311-3030023301210101-0021032123033010-0203312312123023-3130210113323311-3100333122023020)
- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-2102220100231101-2233321202110130-3131001011222331-1222110220020312-3211102102212312-0130031110002102-3022120103111002-3100322213130101)
- custom_network_config.slo_config.static_routes.static_routes.node_interface

<a id="canonical-3031010211313031-0030012121010113-0130133231133312-3312033122212201-0112331312003000-2221203012330300-1020333332010103-2113031313121320"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

Receipt-pinned upstream constraints:

```json
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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321033220320132-2023121231113222-0312220200031231-3303132320232101-2231120230202121-2120023123311201-2222223232010331-1322221223130020"></a>

## Direct properties — node_interface / 201021322002 / 3

- [list](resources--voltstack_site--reference--group-005.md#canonical-1112122030113231-0102301311303111-1210302001301002-1303131321023120-0311122003000231-3130300203113123-2332321023320010-3020322001133233): complete subsection reference.

<a id="canonical-1210212130023233-3300323333120302-0131231113113112-0203331020113232-1121223222303210-3100001030102333-1230131132222000-1020210231312212"></a>

## Next pages — node_interface / 201021322002 / 4

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-1112122030113231-0102301311303111-1210302001301002-1303131321023120-0311122003000231-3130300203113123-2332321023320010-3020322001133233)
- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-2102220100231101-2233321202110130-3131001011222331-1222110220020312-3211102102212312-0130031110002102-3022120103111002-3100322213130101)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1112122030113231-0102301311303111-1210302001301002-1303131321023120-0311122003000231-3130300203113123-2332321023320010-3020322001133233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222232031111200-3112112032022332-0303202222230201-3302300133033311-1331331011000332-0122103111011010-3303322233210223-3001110331203322"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface.list — list / 223120012122 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-1222233330322332-1311022211003101-0230120133122311-3030023301210101-0021032123033010-0203312312123023-3130210113323311-3100333122023020)
- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-2102220100231101-2233321202110130-3131001011222331-1222110220020312-3211102102212312-0130031110002102-3022120103111002-3100322213130101)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-2200221310302012-3322200200333010-3030232302020222-2302033101133001-0202223003010333-3223021332310122-3311330231012222-1133303221012201)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list

<a id="canonical-3030313110221230-1131113020111332-1301213003333021-0002202030121113-2311123133220120-1011022010000321-3112012021013311-3303300031001301"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121030230220103-2233031320230302-0221322223330000-1012322220312031-3223230323301110-0231003120020111-2001330311032312-1223000330200100"></a>

## Direct properties — list / 223120012122 / 3

- [interface](resources--voltstack_site--reference--group-005.md#canonical-3010101313312112-0311302003200010-0022202303221100-3310003222231203-0312221001001323-3113021001110020-3202123310121031-0221120131103030): complete subsection reference.

<a id="canonical-2001223310000211-0133021222121023-0210300101101221-0221221123013321-2002313203022212-2122212100220302-0311231011023232-3332230313022030"></a>

<a id="canonical-2313203111121233-3103112300322013-0203321302132200-1030111120330111-2110232231022312-0202212211333322-0111032232223002-1112233001232103"></a>

## node property — list / 223120012122 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-3120010232301010-1021000323312331-0012023313001320-0033201033011201-3000210102210200-3322321113333231-3022301313113200-1212010101221302"></a>

## Next pages — list / 223120012122 / 5

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](resources--voltstack_site--reference--group-005.md#canonical-3010101313312112-0311302003200010-0022202303221100-3310003222231203-0312221001001323-3113021001110020-3202123310121031-0221120131103030)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-2200221310302012-3322200200333010-3030232302020222-2302033101133001-0202223003010333-3223021332310122-3311330231012222-1133303221012201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3010101313312112-0311302003200010-0022202303221100-3310003222231203-0312221001001323-3113021001110020-3202123310121031-0221120131103030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103222012310323-0012022200101212-1011202313121020-3311303023213100-3232313123023000-0100130222312123-2323030022123123-2110132023221111"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface — interface / 032010003131 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-1222233330322332-1311022211003101-0230120133122311-3030023301210101-0021032123033010-0203312312123023-3130210113323311-3100333122023020)
- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-2102220100231101-2233321202110130-3131001011222331-1222110220020312-3211102102212312-0130031110002102-3022120103111002-3100322213130101)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-2200221310302012-3322200200333010-3030232302020222-2302033101133001-0202223003010333-3223021332310122-3311330231012222-1133303221012201)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-1112122030113231-0102301311303111-1210302001301002-1303131321023120-0311122003000231-3130300203113123-2332321023320010-3020322001133233)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-0321302312033021-0230121130102021-3103333312221211-0000030033202213-3212212132103203-3213233012112222-0000311130130330-0300333133333133"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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

<a id="canonical-2303322033133031-2122233033331103-0001022231123332-1232312232000221-2300122123002201-2113000033100023-2333010031223230-2330332003302132"></a>

## Direct properties — interface / 032010003131 / 3

<a id="canonical-2221231132123301-3013101013112030-0131323012021301-0000000223112220-3131013232322203-3203212120200133-2123132110100203-1322032311231212"></a>

<a id="canonical-2313332101112011-0222322201222311-2122102332022321-0101202211032111-2210200110303100-1212320332232212-3212001301201313-1322133031202121"></a>

## kind property — interface / 032010003131 / 4

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

<a id="canonical-1011200320202200-0003002322312012-0022221301123112-3331303013013312-0113321233323002-1022211031022311-0213020023132231-2133313133123023"></a>

<a id="canonical-3220311000023231-1212331032111123-2222301100312320-1230133133033120-0322110120303213-0023102130222222-0202212201121311-0132213233103200"></a>

## name property — interface / 032010003131 / 5

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

<a id="canonical-1020011023233110-3103230012121321-2213132130231232-3323222331212101-0120212222211223-1123033113200232-0313033003302120-2013011222100031"></a>

<a id="canonical-0001311303202013-0002201312130133-1313310232300303-2320031023001013-0121320220323231-0330312023223023-2001110230012310-0021022300213102"></a>

## namespace property — interface / 032010003131 / 6

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

<a id="canonical-2210332223020031-1233213330313311-2100122113000313-1220120132231321-2223202221212012-0123110212130320-3301313203200221-2123200233211323"></a>

<a id="canonical-1003231310122112-2300111010200122-1210232023202023-2123000013011110-0032301323112313-1302210231323323-0112301111300300-2102103321111203"></a>

## tenant property — interface / 032010003131 / 7

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

<a id="canonical-0022133031112001-0212210103332030-2231201232131000-0322113321201330-2210211220233322-3023021022320122-1320332002313112-3310112132003302"></a>

<a id="canonical-1020301302220223-1103132311133031-1322320110310232-0030312321021131-3322102133033213-3221221103332111-2030100111133301-3211120110211121"></a>

## uid property — interface / 032010003131 / 8

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

<a id="canonical-2220112123030031-3232312302110103-1323120321220313-0012211232200103-1321030223313223-0032202311022013-1311233313103332-0312331212001233"></a>

## Next pages — interface / 032010003131 / 9

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-1112122030113231-0102301311303111-1210302001301002-1303131321023120-0311122003000231-3130300203113123-2332321023320010-3020322001133233)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1003112130012011-2001132133020101-3313022022312333-1303310032020322-1022302212213212-2120201321002113-1002101031021031-0101120022113123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031023313020312-3300131111232133-3333220231103020-0011233002121220-0301013230102100-2312130212331133-0222333203113332-0212221332222130"></a>

## custom_network_config.slo_config.static_v6_routes — static_v6_routes / 100003213230 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- custom_network_config.slo_config.static_v6_routes

<a id="canonical-3201302220203203-3223002023113332-0012101011120312-1221132133020332-2023232130301213-0023310203002301-3221030230033113-3321311110031002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_v6_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002313220223111-3111121023330210-2220221010301200-3003201122033211-2321131023232300-3012013213330002-0331003013212233-2100302331210220"></a>

## Direct properties — static_v6_routes / 100003213230 / 3

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-3203302113010101-2211213322220323-3233330323210333-1010111002322003-0321111120322022-3132122302203230-2133103120023310-3021022112231010): complete subsection reference.

<a id="canonical-1222020131113001-2011130032120231-3032222203310223-1023031133200133-2111001330112110-1302101331130310-3130301111022330-1200310211201233"></a>

## Next pages — static_v6_routes / 100003213230 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3203302113010101-2211213322220323-3233330323210333-1010111002322003-0321111120322022-3132122302203230-2133103120023310-3021022112231010)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3203302113010101-2211213322220323-3233330323210333-1010111002322003-0321111120322022-3132122302203230-2133103120023310-3021022112231010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112303333333230-2110033101110320-3133100123220011-3003320022222333-2301023100131330-2113333333301021-3003133303212120-1010331220223132"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes — static_routes / 322013302212 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-1003112130012011-2001132133020101-3313022022312333-1303310032020322-1022302212213212-2120201321002113-1002101031021031-0101120022113123)
- custom_network_config.slo_config.static_v6_routes.static_routes

<a id="canonical-2010123023113100-2201320322330223-0233101120302123-2011033132312110-1302210333030323-0010221321222312-0311200321233130-0123130310211031"></a>

Type: `"object"`. list nested block, Optional.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321220003201013-1322131123122332-0211311120111120-3123110101302111-2332000130123220-1100231321112301-0203020131123323-1312331112311111"></a>

## Direct properties — static_routes / 322013302212 / 3

<a id="canonical-3213002312310233-2310030113303021-1301211322202112-1232321011002010-1311212030222111-1101032003233303-0313022303021111-3201202002030010"></a>

<a id="canonical-2310111300200010-1323010000033100-2131001232332103-3022101300231232-0001312000300231-3130233120332030-0330302320012302-3010303000003132"></a>

## attrs property — static_routes / 322013302212 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--voltstack_site--reference--group-005.md#canonical-0302211301100303-1011113302110331-1012132031310311-0201103001123031-1321112312230331-3201313300020020-0032121132032222-0033122002130301): complete subsection reference.

<a id="canonical-2223011030122131-2102021220111332-2013013110020200-1002321032033010-3332002322111311-0033312013302201-1113201232103002-1110230210123202"></a>

<a id="canonical-1310112112322111-0330023122202301-1011102310332231-2220113133233001-3112110103130333-2022312100101200-2010001203230223-3203100200222302"></a>

## ip_address property — static_routes / 322013302212 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-1331102002033322-3210210210211001-0211012010202103-2223233030212012-2032100013113322-2301030003022010-2102331212202210-1003103022322133"></a>

<a id="canonical-2132303031313013-3332330301003002-3012223032101321-1322022213231123-2122330031020201-0322010111011221-0110010021100301-3121322102133231"></a>

## ip_prefixes property — static_routes / 322013302212 / 6

Type: `["list", "string"]`. Optional.

List of IPv6 route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--voltstack_site--reference--group-005.md#canonical-3200012230033122-2201130301301330-1013312323232321-2133221122321132-1310010100311101-0320112330332022-1331022222013123-1102220011222200): complete subsection reference.

<a id="canonical-3020223203023121-3123311111333010-3102001002322000-2033200100121011-3111111302201223-1230220202222203-1331123200031122-3131313120122001"></a>

## Next pages — static_routes / 322013302212 / 7

- [custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway](resources--voltstack_site--reference--group-005.md#canonical-0302211301100303-1011113302110331-1012132031310311-0201103001123031-1321112312230331-3201313300020020-0032121132032222-0033122002130301)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-3200012230033122-2201130301301330-1013312323232321-2133221122321132-1310010100311101-0320112330332022-1331022222013123-1102220011222200)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-1003112130012011-2001132133020101-3313022022312333-1303310032020322-1022302212213212-2120201321002113-1002101031021031-0101120022113123)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0302211301100303-1011113302110331-1012132031310311-0201103001123031-1321112312230331-3201313300020020-0032121132032222-0033122002130301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012013210023103-3203011212301123-1233123112120120-2223200121331002-1010013231300312-3102110002102202-0100123132233022-0310002111311101"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway — default_gateway / 301220201203 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-1003112130012011-2001132133020101-3313022022312333-1303310032020322-1022302212213212-2120201321002113-1002101031021031-0101120022113123)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3203302113010101-2211213322220323-3233330323210333-1010111002322003-0321111120322022-3132122302203230-2133103120023310-3021022112231010)
- custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-0221313333210300-3220302203303313-2313223202330312-0222131230320033-2133010330100102-0322002201230333-2202203122123103-2232131230010003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

<a id="canonical-1123022230003123-3323003131003331-2321233320030012-2300222100323320-2021022200230103-1323121231323132-0103031222300031-3232002022100222"></a>

## Direct properties — default_gateway / 301220201203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023110202022211-3001120312331211-0113201013103011-2003013321333332-2112201233123112-2231323032222301-1133220212301031-1102230231032020"></a>

## Next pages — default_gateway / 301220201203 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3203302113010101-2211213322220323-3233330323210333-1010111002322003-0321111120322022-3132122302203230-2133103120023310-3021022112231010)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3200012230033122-2201130301301330-1013312323232321-2133221122321132-1310010100311101-0320112330332022-1331022222013123-1102220011222200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310033330210231-3120333222001102-0023101322203013-3020323030222301-3222022000023231-3331221002001201-0113323312011211-3012023011310002"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface — node_interface / 111200023012 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-1003112130012011-2001132133020101-3313022022312333-1303310032020322-1022302212213212-2120201321002113-1002101031021031-0101120022113123)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3203302113010101-2211213322220323-3233330323210333-1010111002322003-0321111120322022-3132122302203230-2133103120023310-3021022112231010)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface

<a id="canonical-3022200130233323-1311021111102323-0131310020210212-0101113013311121-0102011002221013-0333310120132320-2232111012332031-1020330123313230"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

Receipt-pinned upstream constraints:

```json
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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200033210211320-1031013113221023-2301031333211312-1231122220111230-1103033112012131-0101103223030100-1113003033120321-3111230322320230"></a>

## Direct properties — node_interface / 111200023012 / 3

- [list](resources--voltstack_site--reference--group-005.md#canonical-1111311123110222-3032223330020021-3003110311010112-1032202023132121-1000111100021000-1202230203133000-1120313001300110-0112202031312201): complete subsection reference.

<a id="canonical-3220012313123201-2023232223213332-0220320001103013-3330323211202330-1122322001201012-1233322232100102-1222310320320033-1201301302311131"></a>

## Next pages — node_interface / 111200023012 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-1111311123110222-3032223330020021-3003110311010112-1032202023132121-1000111100021000-1202230203133000-1120313001300110-0112202031312201)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3203302113010101-2211213322220323-3233330323210333-1010111002322003-0321111120322022-3132122302203230-2133103120023310-3021022112231010)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1111311123110222-3032223330020021-3003110311010112-1032202023132121-1000111100021000-1202230203133000-1120313001300110-0112202031312201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322233310200301-2333322121320111-3332031031003330-2122330131320132-3020103020210323-0310203010131132-2301120220301123-2331011231202220"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list — list / 213222032102 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-1003112130012011-2001132133020101-3313022022312333-1303310032020322-1022302212213212-2120201321002113-1002101031021031-0101120022113123)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3203302113010101-2211213322220323-3233330323210333-1010111002322003-0321111120322022-3132122302203230-2133103120023310-3021022112231010)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-3200012230033122-2201130301301330-1013312323232321-2133221122321132-1310010100311101-0320112330332022-1331022222013123-1102220011222200)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-2111221312303021-0323111101330231-1311320030323323-2102332031221121-2210331132320031-1232100103320022-2312121103120030-0203103313202231"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113211300221320-2113110233333330-2130220121131030-1302302113121002-2323322130021201-0112331331021010-2010300321003311-1032332223321102"></a>

## Direct properties — list / 213222032102 / 3

- [interface](resources--voltstack_site--reference--group-005.md#canonical-2020120010131023-2332003110130002-1033112011302321-2013000202312113-0231101321210111-3030103010131111-0023120031212222-3123221331012232): complete subsection reference.

<a id="canonical-1000211212330102-1302211023203102-3022202231301111-2133323132220200-3033113120302300-0001020221123030-1323233222013123-1322021030102321"></a>

<a id="canonical-2200203220000302-3311232011101101-3000313301002112-2330322322220213-1010021210202213-3233023100030133-1001232201310030-3321030033311122"></a>

## node property — list / 213222032102 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-3002112212013022-1210201332331010-1223133001022100-1212311030320300-2322200020303131-2130132113122200-3330113002221332-3212110031121221"></a>

## Next pages — list / 213222032102 / 5

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](resources--voltstack_site--reference--group-005.md#canonical-2020120010131023-2332003110130002-1033112011302321-2013000202312113-0231101321210111-3030103010131111-0023120031212222-3123221331012232)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-3200012230033122-2201130301301330-1013312323232321-2133221122321132-1310010100311101-0320112330332022-1331022222013123-1102220011222200)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2020120010131023-2332003110130002-1033112011302321-2013000202312113-0231101321210111-3030103010131111-0023120031212222-3123221331012232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301110023110113-3220320311023001-1200100333122022-3102320220311102-3103323220312202-2222113122220201-0303333111032313-1221012212023333"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface — interface / 233231311311 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [custom_network_config.slo_config](resources--voltstack_site--reference--group-005.md#canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--reference--group-005.md#canonical-1003112130012011-2001132133020101-3313022022312333-1303310032020322-1022302212213212-2120201321002113-1002101031021031-0101120022113123)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-3203302113010101-2211213322220323-3233330323210333-1010111002322003-0321111120322022-3132122302203230-2133103120023310-3021022112231010)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--voltstack_site--reference--group-005.md#canonical-3200012230033122-2201130301301330-1013312323232321-2133221122321132-1310010100311101-0320112330332022-1331022222013123-1102220011222200)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-1111311123110222-3032223330020021-3003110311010112-1032202023132121-1000111100021000-1202230203133000-1120313001300110-0112202031312201)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-2212013032323223-0233101023232231-1000333202010331-2122232103223212-2333231012313020-1032322110222122-0101132123312003-2321313000213110"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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

<a id="canonical-3103102203001310-3221002203202132-0133102231213103-3023013011200033-1203330121321310-2011030201303101-3223201331123222-0203023112112110"></a>

## Direct properties — interface / 233231311311 / 3

<a id="canonical-1013301031223111-2120001103122221-3030331323233032-0211121220110000-3311322322210321-2122002102033002-0212302003311123-1030133302301111"></a>

<a id="canonical-1022121132001202-3120102322221322-0300200302120202-2103020021320113-0001031012030020-2331210022030110-0002003103311131-2220033023122011"></a>

## kind property — interface / 233231311311 / 4

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

<a id="canonical-0330030220220333-3002131211110103-0231113202211001-3001031032311321-2100121023220201-2320220023210011-0112112033232011-3301001001322222"></a>

<a id="canonical-1133333220011311-1021011003110333-2302020121300022-0122002010003122-1200003202102001-0223131020101112-0233322020221131-2110010311110221"></a>

## name property — interface / 233231311311 / 5

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

<a id="canonical-0221132323211103-1112013202223213-3230133320310230-3110331302120123-2013103113332112-2100333222311132-1231110303030323-2222223031110312"></a>

<a id="canonical-1020100002030112-0202303102221312-3202301200113210-0033001133113232-3103000222003123-1300300110032313-3011121201322001-3202100310232310"></a>

## namespace property — interface / 233231311311 / 6

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

<a id="canonical-1113032230222210-2131313220323012-3222133021330033-3333202313302230-3201310001211321-1330223303132121-3212202203332221-3303303011212213"></a>

<a id="canonical-1111330312333213-3033322100221122-2331132131322133-0100210000310333-0221233210123011-2000011231012030-3030033322331211-0123320200030111"></a>

## tenant property — interface / 233231311311 / 7

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

<a id="canonical-1030132013000000-1121031122113112-3002211220210332-3300212303001333-3233000100211333-2002021233231110-3213123310002213-3122033300222131"></a>

<a id="canonical-1100020202021131-3332130323131111-0113030212101223-2313233301032202-2223212132211102-1133002210201203-2130211320103003-3110012003000323"></a>

## uid property — interface / 233231311311 / 8

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

<a id="canonical-1333312313311030-3123202230233112-3021201022020131-3111331013220111-1230233132012133-1101210012332103-2011023003003112-1223001000120320"></a>

## Next pages — interface / 233231311311 / 9

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-005.md#canonical-1111311123110222-3032223330020021-3003110311010112-1032202023132121-1000111100021000-1202230203133000-1120313001300110-0112202031312201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0310133100133310-2222131222012001-1310311313333101-0213210223133331-1231312300220032-0103011301103021-3200303130212111-3121313102230010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021212031213112-1233320001022220-3310213232002231-2210331001302022-3321000100110112-0102320313303110-1332200121012032-0032230011022032"></a>

## custom_network_config.sm_connection_public_ip — sm_connection_public_ip / 023111320312 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- custom_network_config.sm_connection_public_ip

<a id="canonical-2233002111122202-0333033001033111-0203230003310333-1030102200303213-3212112123223133-2102102003100023-0330130100212010-2212123222022102"></a>

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

<a id="canonical-1010333001020111-3132111322032303-3111012030102213-2321010220223130-0121320003013212-2121111130233223-3002122330022012-0003322011010200"></a>

## Direct properties — sm_connection_public_ip / 023111320312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301031003302331-3133322330011113-1320231222220220-3222032210200032-0113231302000200-3210111100220301-3301333020031020-1032302301320110"></a>

## Next pages — sm_connection_public_ip / 023111320312 / 4

- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3121323010222232-2331132302030203-3122001100110001-1210223031133013-1120010000213111-0310001310101133-0200210132333123-1133313122320123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211201013011231-1123111213123001-0332221021120132-0213333302113002-2002311001201101-2301210130000033-2310032123002111-0202211133121322"></a>

## custom_network_config.sm_connection_pvt_ip — sm_connection_pvt_ip / 113313210011 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- custom_network_config.sm_connection_pvt_ip

<a id="canonical-3103102310111322-3310110212223002-0231100133021110-0211303221210213-1232203130132130-2332211000003020-3033331012330212-0312311003200313"></a>

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

<a id="canonical-1213223202302032-3132101101211130-3322102012102223-2221132100211330-2123001220120030-3313002023122000-2331111130031202-3230223202211022"></a>

## Direct properties — sm_connection_pvt_ip / 113313210011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002122012013120-3203202010211330-0330030103111332-3202303230332003-3011213213223033-0231122003100131-3330322332003102-2002131000023313"></a>

## Next pages — sm_connection_pvt_ip / 113313210011 / 4

- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110210311332221-0002231330021013-0330022211312201-3231210313110033-2112000230200120-1133113100111312-2000130301330022-2113233001300003"></a>

## custom_storage_config — custom_storage_config / 020312031033 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- custom_storage_config

<a id="canonical-2012001120122131-2020300012013020-2123011021211210-1120121102101231-0121120112203012-3000003230002333-1021202331312113-1120033013331233"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_storage\_config, default\_storage\_config; Default: default\_storage\_config\]
VssStorageConfiguration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_storage_class",
    "storage_class_list"),
  validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_storage_device",
    "storage_device_list"),
  validators.ConflictingObjectAttributes("no_storage_interfaces",
    "storage_interface_list")}
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
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage_class\",\"storage_class_list\"]",
  "x-ves-oneof-field-storage_device_choice": "[\"no_storage_device\",\"storage_device_list\"]",
  "x-ves-oneof-field-storage_interface_choice": "[\"no_storage_interfaces\",\"storage_interface_list\"]"
}
```

OneOf alternatives in this subsection:

- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-2012001120122131-2020300012013020-2123011021211210-1120121102101231-0121120112203012-3000003230002333-1021202331312113-1120033013331233)
- [default_storage_config](resources--voltstack_site--reference--group-009.md#canonical-3001000033122001-3121013231330231-0313132013202222-3302100123312300-1210133311213110-0302220212110122-1123031230103122-2130212033123121)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_storage_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202232333101222-3321201211113231-3030021010200123-1203322002011331-2201131302120313-2133023232212222-2102303013312002-0132131123131130"></a>

## Direct properties — custom_storage_config / 020312031033 / 3

- [default_storage_class](resources--voltstack_site--reference--group-005.md#canonical-1303121221130112-2121130210012120-0332002232001012-2300100300001001-0013130011132300-1203112101200132-3030310311123003-3203301230333233): complete subsection reference.

- [no_static_routes](resources--voltstack_site--reference--group-005.md#canonical-2310311311320232-3103122132012122-3200302311113123-2330020021021130-0200131030200020-1023333133300102-0131023322031222-2121023032203230): complete subsection reference.

- [no_storage_device](resources--voltstack_site--reference--group-005.md#canonical-0303321012032212-1200203113023231-3120321113322332-1133212210230230-1003032221011212-0231002223021233-2112103003200103-3120303210202210): complete subsection reference.

- [no_storage_interfaces](resources--voltstack_site--reference--group-005.md#canonical-2012023023231333-0111012103223220-1103212131313222-1023220003231003-2022102330211233-2020331111121013-1201012331310121-3120000223303321): complete subsection reference.

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-0010013332203001-0113132310200201-1011031031323020-3130002332300123-3110122210200111-1330032111302123-2100222030001223-3213031122112122): complete subsection reference.

- [storage_class_list](resources--voltstack_site--reference--group-006.md#canonical-3200223201220132-0311220310300003-2103003133003311-1323333023323123-3011330010121132-2203203012313212-2030201103311030-0213120200223032): complete subsection reference.

- [storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321): complete subsection reference.

- [storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333): complete subsection reference.

<a id="canonical-0011133320230231-2121210230022111-0333013312322001-0120023301030132-2102200000201011-1331020022300222-0223110103000133-3302120303330222"></a>

## Next pages — custom_storage_config / 020312031033 / 4

- [custom_storage_config.default_storage_class](resources--voltstack_site--reference--group-005.md#canonical-1303121221130112-2121130210012120-0332002232001012-2300100300001001-0013130011132300-1203112101200132-3030310311123003-3203301230333233)
- [custom_storage_config.no_static_routes](resources--voltstack_site--reference--group-005.md#canonical-2310311311320232-3103122132012122-3200302311113123-2330020021021130-0200131030200020-1023333133300102-0131023322031222-2121023032203230)
- [custom_storage_config.no_storage_device](resources--voltstack_site--reference--group-005.md#canonical-0303321012032212-1200203113023231-3120321113322332-1133212210230230-1003032221011212-0231002223021233-2112103003200103-3120303210202210)
- [custom_storage_config.no_storage_interfaces](resources--voltstack_site--reference--group-005.md#canonical-2012023023231333-0111012103223220-1103212131313222-1023220003231003-2022102330211233-2020331111121013-1201012331310121-3120000223303321)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0010013332203001-0113132310200201-1011031031323020-3130002332300123-3110122210200111-1330032111302123-2100222030001223-3213031122112122)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-006.md#canonical-3200223201220132-0311220310300003-2103003133003311-1323333023323123-3011330010121132-2203203012313212-2030201103311030-0213120200223032)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1303121221130112-2121130210012120-0332002232001012-2300100300001001-0013130011132300-1203112101200132-3030310311123003-3203301230333233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032330312333103-1013102021011210-0130331301333300-1313310232110120-0213203310331113-2323030133131032-3123012123303301-3023002003130313"></a>

## custom_storage_config.default_storage_class — default_storage_class / 201122322323 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- custom_storage_config.default_storage_class

<a id="canonical-1000302122312320-2211311113230222-1230001220333231-2300020110313012-1202203330000233-2122113012132202-2131231132102331-3123212023320121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default storage class.

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
default_storage_class = {}
```

<a id="canonical-1033303311202231-0123310001113100-3313200111202000-0022123122311233-3323012020120132-2203023130231110-2103333030112300-1021213331020210"></a>

## Direct properties — default_storage_class / 201122322323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210300213133303-2201100212003103-2213010031332211-1322310023132313-1312133310001021-3232222213200210-0110213210011200-2221121101311122"></a>

## Next pages — default_storage_class / 201122322323 / 4

- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2310311311320232-3103122132012122-3200302311113123-2330020021021130-0200131030200020-1023333133300102-0131023322031222-2121023032203230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021020333332332-1030013223301230-2112220300231331-2320103123123301-3323130333101132-3302121021323222-2220120301203203-2021031120011330"></a>

## custom_storage_config.no_static_routes — no_static_routes / 021010320220 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- custom_storage_config.no_static_routes

<a id="canonical-1323333022001123-1021120032133012-0012110231303100-3320331232300122-0121001311320213-2120102323111211-2201201120032203-2222310003332023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static routes.

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
no_static_routes = {}
```

<a id="canonical-2002100232301120-1320330020122033-0030033223001321-1030203001202031-1221101131321132-0110213200231333-3233210331002113-2320013202302322"></a>

## Direct properties — no_static_routes / 021010320220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2310131303123323-2221231031032030-3022010301122210-2212301000111020-1223321020013213-0113110132130020-2022123210033222-3030331322001202"></a>

## Next pages — no_static_routes / 021010320220 / 4

- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0303321012032212-1200203113023231-3120321113322332-1133212210230230-1003032221011212-0231002223021233-2112103003200103-3120303210202210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122021010332331-2222233230120311-0231123322000320-0021101223211213-0230013202232120-2010222333323222-2000211212312103-1130103230202310"></a>

## custom_storage_config.no_storage_device — no_storage_device / 133312002300 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- custom_storage_config.no_storage_device

<a id="canonical-1201230211113221-0320230320332333-3333201232303330-0323131321211122-2003033310222221-2213120332330111-0233211130133030-0030311120230133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no storage device.

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
no_storage_device = {}
```

<a id="canonical-3010030330121333-1021001321002332-0220133213213223-3022312230210333-3302321012013332-1010232220220300-1020032032013222-1112113000032003"></a>

## Direct properties — no_storage_device / 133312002300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301022230230233-1011231030113000-0231303001321102-3011030132100213-1012213013212133-0012033200222030-2000313122222313-0103121233232210"></a>

## Next pages — no_storage_device / 133312002300 / 4

- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2012023023231333-0111012103223220-1103212131313222-1023220003231003-2022102330211233-2020331111121013-1201012331310121-3120000223303321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100120203233032-2031110221231222-1311202110102000-3003212333131330-3131302031120010-3301111202000230-3320321012221330-1122203032213202"></a>

## custom_storage_config.no_storage_interfaces — no_storage_interfaces / 023023102320 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- custom_storage_config.no_storage_interfaces

<a id="canonical-2213113230200221-0223233220322111-2122100112320331-3312010101303111-2220112303003210-0202113302200021-1002313222223131-3231121332233101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no storage interfaces.

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
no_storage_interfaces = {}
```

<a id="canonical-3101302231011201-2303030133232031-3323323223032322-3201120233023133-1100233322020113-3223013322331031-3321122033000330-1331211220222212"></a>

## Direct properties — no_storage_interfaces / 023023102320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211203021121220-1110120302033102-1230223012133120-0131230332322331-0201122303221332-0303311010310222-1220000130323013-0111310310220000"></a>

## Next pages — no_storage_interfaces / 023023102320 / 4

- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0010013332203001-0113132310200201-1011031031323020-3130002332300123-3110122210200111-1330032111302123-2100222030001223-3213031122112122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032023301022031-0321103131222302-3201330131022031-1000313031303203-1021313010130310-3122013213311012-3120031233233210-0031023022203222"></a>

## custom_storage_config.static_routes — static_routes / 230103213231 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- custom_storage_config.static_routes

<a id="canonical-2200311230212213-0013200103310200-1312130231111311-1110201221130102-1133323020212023-3231011033220102-0223001202220110-1300013210111011"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303201322031032-1032133112313322-3213021101033133-1203200013111332-3031021203320223-0023111031233333-0000202202013220-1321003123320013"></a>

## Direct properties — static_routes / 230103213231 / 3

- [static_routes](resources--voltstack_site--reference--group-005.md#canonical-2300310110003322-0200012021220130-2123203021230230-2303022333110233-2313312311210211-0232003100312110-2230131203001132-2203002101122230): complete subsection reference.

<a id="canonical-0210323100310220-1332231012310133-1302022331010231-3312012310121030-2320010003112200-2121313231022031-2101003302203002-0131221122220000"></a>

## Next pages — static_routes / 230103213231 / 4

- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-2300310110003322-0200012021220130-2123203021230230-2303022333110233-2313312311210211-0232003100312110-2230131203001132-2203002101122230)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2300310110003322-0200012021220130-2123203021230230-2303022333110233-2313312311210211-0232003100312110-2230131203001132-2203002101122230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221113311100210-3300312203030121-1311112031003011-2303130033233220-3210031201323110-0301032203331101-0231311300002021-2200233122013331"></a>

## custom_storage_config.static_routes.static_routes — static_routes / 131122103332 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0010013332203001-0113132310200201-1011031031323020-3130002332300123-3110122210200111-1330032111302123-2100222030001223-3213031122112122)
- custom_storage_config.static_routes.static_routes

<a id="canonical-3120231323201122-0230012120212300-0020022110013133-0222311231313203-1133232302133222-3322212331203001-1020223103333312-1010012033200102"></a>

Type: `"object"`. list nested block, Optional.

Static Routes. List of static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130200021112202-2202110100230233-1220010103103311-3103022300300020-2110300321330130-3101012311332023-0311210231313311-3330023221101133"></a>

## Direct properties — static_routes / 131122103332 / 3

<a id="canonical-1120313311131110-1010320312130013-1233010022103110-0000103022312101-0323102302032033-3130112232322132-2222110013233202-3312203311313233"></a>

<a id="canonical-3320312213221001-3001013022020203-1001131213122023-1023031300313002-1332301202021110-1032200301013301-3132322203132122-2113303303003121"></a>

## attrs property — static_routes / 131122103332 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--voltstack_site--reference--group-005.md#canonical-0220110121113211-0123200320311213-2033231212131113-0101303012033002-0331200212333220-0100030201202310-0203300222031112-1212131332212032): complete subsection reference.

<a id="canonical-2000002110220110-2303202203021322-2132113332333201-0302002123321320-2010131223201032-0221022310011010-1303232031203111-2020213101333122"></a>

<a id="canonical-2222322010101231-0113132131003033-3200332202012230-0303322131132302-2330331121331232-3001011102023121-0212132232230033-2311030111310221"></a>

## ip_address property — static_routes / 131122103332 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-0120132003321120-3020112030203232-3233012312133203-0222200131223130-0100322002020331-2223011221030213-2110010231330332-3211322132032220"></a>

<a id="canonical-0320021332132320-3322322101302000-0113321202130312-2130121323111122-2320312312000133-1032112301232320-2330032132330201-3203103023233031"></a>

## ip_prefixes property — static_routes / 131122103332 / 6

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--voltstack_site--reference--group-006.md#canonical-2103010001332321-3000301122011331-2103133111021230-0301232203102112-3333122212301320-0010123301123323-1031033201333332-2031303213320033): complete subsection reference.

<a id="canonical-0203002210301012-3113101232011231-0200010330310100-1233031233321103-3120032331213210-2221033323020301-0230113313233101-1121133201221110"></a>

## Next pages — static_routes / 131122103332 / 7

- [custom_storage_config.static_routes.static_routes.default_gateway](resources--voltstack_site--reference--group-005.md#canonical-0220110121113211-0123200320311213-2033231212131113-0101303012033002-0331200212333220-0100030201202310-0203300222031112-1212131332212032)
- [custom_storage_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-006.md#canonical-2103010001332321-3000301122011331-2103133111021230-0301232203102112-3333122212301320-0010123301123323-1031033201333332-2031303213320033)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0010013332203001-0113132310200201-1011031031323020-3130002332300123-3110122210200111-1330032111302123-2100222030001223-3213031122112122)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0220110121113211-0123200320311213-2033231212131113-0101303012033002-0331200212333220-0100030201202310-0203300222031112-1212131332212032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
