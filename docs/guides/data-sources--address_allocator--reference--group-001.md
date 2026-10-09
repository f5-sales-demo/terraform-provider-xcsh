---
page_title: "xcsh_address_allocator reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_address_allocator reference."
---

# xcsh_address_allocator reference

<a id="canonical-0012311232330222-1032233333322133-2322023003332231-0123033020123120-0101033200130133-1131010113203203-2210031202200112-2212010003230021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_address_allocator](../data-sources/address_allocator.md#canonical-2303133311012021-3322130130031103-2120032332031012-3013200021213232-0100310121031123-3230030130121113-3213033330111232-3203303021311023)
- Property reference

<a id="canonical-2033011303322310-3300330303331313-3331333112001232-2102303330311200-0202020213223011-3013113132003313-0113223311113332-2111010000023221"></a>

### Direct properties for `xcsh_address_allocator`

- [address_allocation_scheme](data-sources--address_allocator--reference--group-001.md#canonical-2111321100330323-3002321212213203-0330313012110130-1321112130102233-1320122010220022-0000130220031132-2330210313130332-1223112100032321): complete subsection reference.

<a id="canonical-0011323130330331-2033223030001313-0110202322100103-3020320002223123-3003310332130313-0233012330020200-2100130132011103-0310321203323333"></a>

<a id="canonical-0130123202110103-3311203023130320-3222100003212030-0233230202102230-2111121203212323-2303203103012121-2320023121122323-0301130200210020"></a>

#### `address_pool` property

Type: `["list", "string"]`. Computed.

Address pool from which the allocator carves out subnets or addresses to its clients.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2102332112022003-2022323032321200-0201310023210012-2111120000133303-0031003333323021-0223000123120222-1233101301223030-3102032222221233"></a>

<a id="canonical-2110311300130332-2013300130213333-2213131333102120-2232133013020102-1121331221301310-0031031223123002-2111333033231330-3223102021320330"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-3231303311322210-3333203121033120-3011013203002101-3213321000110220-0213330030022031-0002023100013230-1211130122302130-0313003303322023"></a>

<a id="canonical-0302300203120300-0102111312333203-2101022212233330-1222201323312103-3330223311311233-0333333310310222-1032112101133213-2203133102023022"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the AddressAllocator.

Additional upstream details:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-3012111303233010-1331320232313012-2303131200323000-0203101311313232-3120221020200321-2121212203230130-3131001110230101-2103212331003120"></a>

<a id="canonical-2121300332333212-3211131213033223-1223121130201100-1213132011032232-1101210332213330-2223323311300330-0221300333111022-3301230320222021"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3120302321203302-3232022311100000-1212023212002102-1012002200223332-1320210312113122-3110121020313123-3010332110231020-3331333223312312"></a>

<a id="canonical-3332102312212200-1120220200022110-2032023323310331-3201331003120201-3322320133113123-0021120312331101-0001231210002312-3030323030220000"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1232211331211222-3100030323103121-2201012022332200-3132020333230221-3010120030012333-2121221302312201-2013210311010032-0130322133121020"></a>

<a id="canonical-3003233011020121-1120102122120211-0200223303030223-1013303320021330-1302103123303011-1110213020212112-0000031022101332-2001320312121130"></a>

#### `mode` property

Type: `"string"`. Computed.

\[Enum: LOCAL|GLOBAL\_PER\_SITE\_NODE\] Mode of the address allocator Address allocator is for VERs
within the local cluster or site Allocation is per site and then per node. Possible values are
\`LOCAL\`, \`GLOBAL\_PER\_SITE\_NODE\`. Defaults to \`LOCAL\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "LOCAL",
  "enum": [
    "LOCAL",
    "GLOBAL_PER_SITE_NODE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1111233312003031-2002323320303031-1321331321112231-3002301212232221-1001000321110122-1010310201112200-3003233322203301-1131211100320103"></a>

<a id="canonical-2132220013103233-2001033231330223-1022213201212202-0121220011210302-1021111321100030-1113313020022132-2123322322131220-2203013031113013"></a>

#### `name` property

Type: `"string"`. Required.

Name of the AddressAllocator.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3233011032100110-3303213103230101-0023113022012112-2233013003023020-3113332233203233-2200131012100323-3022231202013030-0212011220222022"></a>

<a id="canonical-2320201230130312-1230213230201222-3210231203320321-2023211012033011-2312221321221223-3002232032010203-3111231100020201-3110302101233330"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the AddressAllocator exists.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

<a id="canonical-2102330130332101-1031123221230230-0133022003203111-1213102002023133-3102212100023210-0303220001132321-0032110102312023-3223111320330313"></a>

### All schema paths for `xcsh_address_allocator`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address_allocation_scheme` | [address_allocation_scheme](data-sources--address_allocator--reference--group-001.md#canonical-3133222322313300-1122332023312120-0102023102202112-0210013130130320-2323321032122012-3101111101331033-1232203230321322-1202200120201222) |
| `address_allocation_scheme.allocation_unit` | [address_allocation_scheme.allocation_unit](data-sources--address_allocator--reference--group-001.md#canonical-3120011222231321-1321112023333213-2130011202003233-1231110101323013-0210221102200213-0313013333020101-0031322220203123-0033010201231232) |
| `address_allocation_scheme.local_interface_address_offset` | [address_allocation_scheme.local_interface_address_offset](data-sources--address_allocator--reference--group-001.md#canonical-0120032320313311-1320011330023231-2203002230330113-3012311310020012-3203312000003001-3323302010130003-0021021120211010-3131120221133223) |
| `address_allocation_scheme.local_interface_address_type` | [address_allocation_scheme.local_interface_address_type](data-sources--address_allocator--reference--group-001.md#canonical-0211100122333300-1321023221032333-3032131320312103-3302212212010310-0003310111233022-3002303132231020-2113312332210111-1231200302023323) |
| `address_pool` | [address_pool](data-sources--address_allocator--reference--group-001.md#canonical-0011323130330331-2033223030001313-0110202322100103-3020320002223123-3003310332130313-0233012330020200-2100130132011103-0310321203323333) |
| `annotations` | [annotations](data-sources--address_allocator--reference--group-001.md#canonical-2102332112022003-2022323032321200-0201310023210012-2111120000133303-0031003333323021-0223000123120222-1233101301223030-3102032222221233) |
| `description` | [description](data-sources--address_allocator--reference--group-001.md#canonical-3231303311322210-3333203121033120-3011013203002101-3213321000110220-0213330030022031-0002023100013230-1211130122302130-0313003303322023) |
| `id` | [ID](data-sources--address_allocator--reference--group-001.md#canonical-3012111303233010-1331320232313012-2303131200323000-0203101311313232-3120221020200321-2121212203230130-3131001110230101-2103212331003120) |
| `labels` | [labels](data-sources--address_allocator--reference--group-001.md#canonical-3120302321203302-3232022311100000-1212023212002102-1012002200223332-1320210312113122-3110121020313123-3010332110231020-3331333223312312) |
| `mode` | [mode](data-sources--address_allocator--reference--group-001.md#canonical-1232211331211222-3100030323103121-2201012022332200-3132020333230221-3010120030012333-2121221302312201-2013210311010032-0130322133121020) |
| `name` | [name](data-sources--address_allocator--reference--group-001.md#canonical-1111233312003031-2002323320303031-1321331321112231-3002301212232221-1001000321110122-1010310201112200-3003233322203301-1131211100320103) |
| `namespace` | [namespace](data-sources--address_allocator--reference--group-001.md#canonical-3233011032100110-3303213103230101-0023113022012112-2233013003023020-3113332233203233-2200131012100323-3022231202013030-0212011220222022) |

<a id="canonical-2111321100330323-3002321212213203-0330313012110130-1321112130102233-1320122010220022-0000130220031132-2330210313130332-1223112100032321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `address_allocation_scheme` properties

Breadcrumbs:

- [xcsh_address_allocator](../data-sources/address_allocator.md#canonical-2303133311012021-3322130130031103-2120032332031012-3013200021213232-0100310121031123-3230030130121113-3213033330111232-3203303021311023)
- [Property reference](data-sources--address_allocator--reference--group-001.md#canonical-0012311232330222-1032233333322133-2322023003332231-0123033020123120-0101033200130133-1131010113203203-2210031202200112-2212010003230021)
- address_allocation_scheme

<a id="canonical-3133222322313300-1122332023312120-0102023102202112-0210013130130320-2323321032122012-3101111101331033-1232203230321322-1202200120201222"></a>

Type: `"single"`. Computed.

Decides the scheme to be used to allocate addresses from the configured address pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2103302220201202-2001332300202021-2101223330232113-2112313320201011-2103031000300032-3333232232232013-2231101222031101-2212012303231221"></a>

### Direct properties for `address_allocation_scheme`

<a id="canonical-3120011222231321-1321112023333213-2130011202003233-1231110101323013-0210221102200213-0313013333020101-0031322220203123-0033010201231232"></a>

#### `address_allocation_scheme.allocation_unit` property

Type: `"number"`. Computed.

Prefix length indicating the size of each allocated subnet. For example, if this is specified as 30,
subnets of /30 will be allocated from the given address pool.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-0120032320313311-1320011330023231-2203002230330113-3012311310020012-3203312000003001-3323302010130003-0021021120211010-3131120221133223"></a>

<a id="canonical-2321201133223103-0311230202200233-0012023321313100-0213131300033010-1131210201300102-2023003013232020-3132331201310332-0111013233220101"></a>

#### `address_allocation_scheme.local_interface_address_offset` property

Type: `"number"`. Computed.

Used to derive address for the local interface from the allocated subnet. If Local Interface Address
Type is set to 'Offset from beginning of Subnet', this offset value is added to the allocated subnet
and used as the local interface address. For example, if the allocated subnet is 192.0.2.0/24..

Additional upstream details:

This is used to derive address for the local interface from the allocated subnet. If Local Interface
Address Type is set to "Offset from beginning of Subnet", this offset value is added to the
allocated subnet and used as the local interface address. For example, if the allocated subnet is
192.0.2.0/24 and offset is set to 2 with Local Interface Address Type set to "Offset from beginning
of Subnet", local interface address of 192.0.2.204 is used. If Local Interface Address Type is set
to "Offset from end of Subnet", this offset value is subtracted from the end of the allocated subnet
and used as the local interface address. For example, if the allocated subnet is 192.0.2.0/24 and
offset is set to 1 with Local Interface Address Type set to "Offset from end of Subnet", local
interface address of 192.0.2.204 is used.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-0211100122333300-1321023221032333-3032131320312103-3302212212010310-0003310111233022-3002303132231020-2113312332210111-1231200302023323"></a>

<a id="canonical-2121203121103213-0013201321323032-0023330122033001-2320321130130030-2033313210021001-3301313301123230-2021031302013022-3223031010220202"></a>

#### `address_allocation_scheme.local_interface_address_type` property

Type: `"string"`. Computed.

\[Enum:
LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN|LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_END|LOCAL\_INTERFACE\_ADDRESS\_FROM\_PREFIX\]
Dictates how local interface address is derived from the allocated subnet Use Nth address of the
allocated subnet as the local interface address, N being the Local Interface Address Offset. For
example, if the allocated subnet is 192.0.2.0/24, Local Interface Address Offset is set to 2 and
Local.. Possible values are \`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN\`,
\`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_END\`,
\`LOCAL\_INTERFACE\_ADDRESS\_FROM\_PREFIX\`. Defaults to
\`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN\`.

Additional upstream details:

Dictates how local interface address is derived from the allocated subnet

Use Nth address of the allocated subnet as the local interface address, N being the Local Interface
Address Offset. For example, if the allocated subnet is 192.0.2.0/24, Local Interface Address Offset
is set to 2 and Local Interface Address Type is set to "Offset from beginning of Subnet", local
address of 192.0.2.204 is used. Use Nth last address of the allocated subnet as the local interface
address, N being the Local Interface Address Offset. For example, if the allocated subnet is
192.0.2.0/24, Local Interface Address Offset is set to 1 and Local Interface Address Type is set to
"Offset from end of Subnet", local address of 192.0.2.204 is used. This case is used for
external\_connector.

Receipt-pinned upstream constraints:

```json
{
  "default": "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN",
  "enum": [
    "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN",
    "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_END",
    "LOCAL_INTERFACE_ADDRESS_FROM_PREFIX"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
