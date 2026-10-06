---
page_title: "xcsh_network_interface reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface reference."
---

# xcsh_network_interface reference

<a id="canonical-2213321332300020-1020311031120301-1321111203210032-2013120101333021-2110012000121100-3331203023010013-2013000231120231-2130103113123333"></a>

### Direct properties for `tunnel_interface.static_ip.cluster_static_ip`

<a id="canonical-3021112021230211-2011132303113103-1132300102111101-2113300222131230-2232221113031211-2322211022021220-3220331230030101-2121332323103210"></a>

#### `tunnel_interface.static_ip.cluster_static_ip.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"128\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "128"
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

<a id="canonical-1330112132002121-2232301100230130-1120113013030222-3130123110103313-0322132122123220-1101120020011221-2031102121101002-3123012113133331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tunnel_interface.static_ip.node_static_ip` properties

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [tunnel_interface](resources--network_interface--reference--group-001.md#canonical-0031231223100101-0221202300220330-0013311301301203-3301232310130302-2330031103323032-2221301010223301-3302220320030311-3202002232312200)
- [tunnel_interface.static_ip](resources--network_interface--reference--group-001.md#canonical-3000301202120000-2120302233320331-1032223120223111-2120030310022133-0301211132223101-1120232202101331-0120120333222230-2033213332133001)
- tunnel_interface.static_ip.node_static_ip

<a id="canonical-3013331022112231-2323212002202010-3130033020132332-2321321200021333-2023202301230232-0233303311333101-0231300033023230-3313103103220122"></a>

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

<a id="canonical-0321023303002033-1310011221311100-2110210113332330-1311331300111330-1112230000102320-0330102120022221-0230112300022103-1110102221212332"></a>

### Direct properties for `tunnel_interface.static_ip.node_static_ip`

<a id="canonical-1121033311100333-2110301111120302-3323210303221231-1333231220012123-0302222102100010-3230211010302110-1022323132100012-1123233111020313"></a>

#### `tunnel_interface.static_ip.node_static_ip.default_gw` property

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

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

<a id="canonical-2111132203200303-0013330110230223-2120031331021202-2132332131002123-2122220232021013-2310213122320220-3332102110331110-2111020231333203"></a>

<a id="canonical-1302202010211220-3311122003211311-3020000320003002-1133300211300203-1023200223203032-3233332231122331-1222301031123203-3130332221333133"></a>

#### `tunnel_interface.static_ip.node_static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-1323032322302331-1321121231032011-3322211203011033-1121112001010120-2013220322233223-3220323233100031-0132223212311112-2311213310033302"></a>

<a id="canonical-2202310002111010-0011330020210131-3001212333031030-1100221300102302-3030303110230011-0332002332022231-1300132000302222-2332300331332220"></a>

#### `tunnel_interface.static_ip.node_static_ip.ip_address` property

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

<a id="canonical-0203320221220111-3202022023030223-1220012000322020-2312313100112201-3033021300110121-0303002220110300-2033222100133313-1312333112003123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tunnel_interface.tunnel` properties

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [tunnel_interface](resources--network_interface--reference--group-001.md#canonical-0031231223100101-0221202300220330-0013311301301203-3301232310130302-2330031103323032-2221301010223301-3302220320030311-3202002232312200)
- tunnel_interface.tunnel

<a id="canonical-3212213203211123-2000113103201332-0001201121300110-1233233100320031-0200021100332330-0130022202223103-1023222331130212-0002330320010200"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-0220000132201013-1212112110100222-1023130023333122-0300311213313033-1230213011332220-1100102022201022-3032223333310112-2102210211222212"></a>

### Direct properties for `tunnel_interface.tunnel`

<a id="canonical-3212013302223121-3223222331233320-0333211023011121-1201100122200301-2321321130323023-3310231130210320-2200012333232101-2322222231301303"></a>

#### `tunnel_interface.tunnel.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1200323130131130-3112002002101213-2032300101230022-3132210323111102-3002013002002312-2202333120300300-3013213331100321-2122101313122302"></a>

<a id="canonical-0311131022220201-0110110122132102-2313301223003310-2022103011010311-0132030312231331-2001202012101300-2133032000030330-2223101300013000"></a>

#### `tunnel_interface.tunnel.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-1321330323202020-0023302321230213-0330223031000031-0323333302100233-0122012310131130-2021222121122333-1021102201231313-2321012102111101"></a>

<a id="canonical-0300123033112330-0232211002300320-3130032302301301-2113101312330110-2303211200120002-1101032101212130-2220201231213200-0302222112013201"></a>

#### `tunnel_interface.tunnel.tenant` property

Type: `"string"`. Computed.

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
