---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- storage_static_routes

<a id="canonical-3300220221311210-2223123102101211-3220330103110130-2013011230213113-0302003112110032-0211100233331131-0000213033313103-2010020102112211"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for storage static routes.

Additional upstream details:

List of storage static routes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_routes")}
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
storage_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313032300003323-1200212133312312-0212122021211232-1303300232332111-2112103212022002-0103033330303011-2310121212100231-3130322003330220"></a>

### Direct properties for `storage_static_routes`

- [storage_routes](resources--fleet--reference--group-004.md#canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231): complete subsection reference.

<a id="canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330)
- storage_static_routes.storage_routes

<a id="canonical-1113103203120230-0103213211302232-2101121302323232-1122111212313231-2311322311200031-1121301130131233-1130133302322133-3311302231122210"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of storage static routes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("subnets")}
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220311232321311-2101323223330230-1103013003202300-2110122210120013-0010021330313201-3023010223031310-3102321301330120-0123212122123231"></a>

### Direct properties for `storage_static_routes.storage_routes`

<a id="canonical-2033130301313321-1232220332313121-3311112212031300-2211323031330223-1200021003201131-3310133300223310-2120201103231320-3022031112310003"></a>

#### `storage_static_routes.storage_routes.attrs` property

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [labels](resources--fleet--reference--group-004.md#canonical-3122130211122322-1301231021112223-1012121131110203-0212120302301103-0321330301322030-1322303112321312-2232110100220131-3212330121123110): complete subsection reference.

- [nexthop](resources--fleet--reference--group-004.md#canonical-2323102010131331-0222022123303131-0001320120102230-3002231130201221-0320012123110310-2111211303213332-2101222302031312-1321301311201213): complete subsection reference.

- [subnets](resources--fleet--reference--group-004.md#canonical-0011031323102333-2133013032021320-3023030012331100-2101022003022112-3302133213002111-2211232221112303-0333322213031131-1131101123022222): complete subsection reference.

<a id="canonical-3122130211122322-1301231021112223-1012121131110203-0212120302301103-0321330301322030-1322303112321312-2232110100220131-3212330121123110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.labels` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231)
- storage_static_routes.storage_routes.labels

<a id="canonical-3200200330113233-0203010303222300-1203213313200323-0002201003200323-2102302130320131-3200001233330020-2303323033120202-2032013003021232"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323102010131331-0222022123303131-0001320120102230-3002231130201221-0320012123110310-2111211303213332-2101222302031312-1321301311201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231)
- storage_static_routes.storage_routes.nexthop

<a id="canonical-1133203310222332-0201030230322312-3010300302332211-3030020321111100-1000213331231233-3213113130130232-2213202002102121-1300021112121130"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

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

<a id="canonical-3013300001010022-2221021022223122-0003323333130212-1331311302101001-1120031333112222-3000012210301023-3221000232201121-1030013220301321"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop`

- [interface](resources--fleet--reference--group-004.md#canonical-0023120323300020-0130223202211212-1030021320000101-0212220230300313-2130112320331220-2112230223221312-3031001230002133-1321032102310320): complete subsection reference.

- [nexthop_address](resources--fleet--reference--group-004.md#canonical-0113300103011330-2021221000112230-2120203033320002-0000031323332030-3330320102133210-3331212021003013-3132112030012121-0110211233022031): complete subsection reference.

<a id="canonical-3312013220222102-2000130020023100-1220022211110211-0000003103133000-3321032200002000-2122321300210323-2112021023213313-0312302203312332"></a>

<a id="canonical-1303200030013303-3100032131000311-3020200212321133-2111230302031023-3323030023031112-0111212230011331-2013010211200103-1322313320220133"></a>

#### `storage_static_routes.storage_routes.nexthop.type` property

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Additional upstream details:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Use the specified address as
nexthop Use the network interface as nexthop Discard nexthop, used when attr type is Advertise Used
in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["NEXT_HOP_DEFAULT_GATEWAY","NEXT_HOP_NETWORK_INTERFACE","NEXT_HOP_USE_CONFIGURED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-0023120323300020-0130223202211212-1030021320000101-0212220230300313-2130112320331220-2112230223221312-3031001230002133-1321032102310320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop.interface` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-2323102010131331-0222022123303131-0001320120102230-3002231130201221-0320012123110310-2111211303213332-2101222302031312-1321301311201213)
- storage_static_routes.storage_routes.nexthop.interface

<a id="canonical-2101122030320222-1011022121032223-2221103102123222-0320321300122312-0211111133123331-0202123303102012-3133231312131131-0210333030032200"></a>

Type: `"object"`. list nested block, Optional.

Nexthop is network interface when type is 'Network-Interface'.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2010002311331120-3020302021332222-0111231333102100-3012031030203010-3123203032033232-1102100010302030-3000120301133103-2310033332011121"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop.interface`

<a id="canonical-1113331231132223-2330021030122031-0131101120123303-1230000221030001-0001103122332133-0020303311021121-3020111331322212-0012121130302303"></a>

#### `storage_static_routes.storage_routes.nexthop.interface.kind` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1333312131013211-0313203020122000-3301102320123301-0132032201020333-0223211322330100-3002023032030030-2200123211033213-1321023132213223"></a>

<a id="canonical-0223210333310231-3300310211333211-2101121330201211-3010301233123131-3311312312321121-1222301231212312-3013322020120232-3031010000110223"></a>

#### `storage_static_routes.storage_routes.nexthop.interface.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1120121103003330-1121213013223210-1320032132100022-0103131021332331-0022022002211021-2222033102200122-2030120310302221-1322112023232323"></a>

<a id="canonical-0211203130031033-0113333031213323-2132301313021233-0322312301232113-1112011300203022-0020130000012102-3300312113031130-2202113210300211"></a>

#### `storage_static_routes.storage_routes.nexthop.interface.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1010221012231303-3111210300330220-3230120302321021-3331011112320213-0331201113330212-2113310212112301-1213323211103302-1310121302020020"></a>

<a id="canonical-3223112023001122-3010302103112232-3130310302001212-0323131213011033-0131102102100322-0301200302221112-3320311121201202-2232101003313332"></a>

#### `storage_static_routes.storage_routes.nexthop.interface.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1013311001101102-1220313102102201-1111300333221302-0210130310131123-2231201133011131-1102101232131003-0221211032032212-0020023000312301"></a>

<a id="canonical-2000010303231003-2121200232012033-3322321023102100-0202133322000222-3003111120320212-2210312303032023-2310130332100223-3030301133203011"></a>

#### `storage_static_routes.storage_routes.nexthop.interface.uid` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0113300103011330-2021221000112230-2120203033320002-0000031323332030-3330320102133210-3331212021003013-3132112030012121-0110211233022031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop.nexthop_address` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-2323102010131331-0222022123303131-0001320120102230-3002231130201221-0320012123110310-2111211303213332-2101222302031312-1321301311201213)
- storage_static_routes.storage_routes.nexthop.nexthop_address

<a id="canonical-1020220223202220-0330131020121100-0332133222131201-2113330101031223-0200301212310132-1130023310213131-1112002021300203-1100001320223111"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0313222021011230-0130122203123220-3131000031130211-3123132313112231-1222212322003312-3132002201110121-1313122020220300-2321131222230110"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop.nexthop_address`

- [dual_stack](resources--fleet--reference--group-004.md#canonical-3102121312220313-0113030102112222-1312112312222122-2222211021110113-3333023030230230-0211333200121331-3222310113021231-1001121303013130): complete subsection reference.

- [IPv4](resources--fleet--reference--group-004.md#canonical-1111103023232130-3012302323132032-1020130021030110-0132010331113111-2021231131202122-3102000212232011-1203002323232112-2333311000032211): complete subsection reference.

- [IPv6](resources--fleet--reference--group-004.md#canonical-3122032102121102-3002123231001213-3101301122202200-2133021320123223-3000333013000223-3020223000310003-2331311202011003-2200100110120102): complete subsection reference.

<a id="canonical-3102121312220313-0113030102112222-1312112312222122-2222211021110113-3333023030230230-0211333200121331-3222310113021231-1001121303013130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-2323102010131331-0222022123303131-0001320120102230-3002231130201221-0320012123110310-2111211303213332-2101222302031312-1321301311201213)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-0113300103011330-2021221000112230-2120203033320002-0000031323332030-3330320102133210-3331212021003013-3132112030012121-0110211233022031)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack

<a id="canonical-3312122000233102-1302032131331000-1132303333011232-2330202011012201-3313213233332222-2120210220032131-0012030130320302-0013133123100333"></a>

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

<a id="canonical-2001030330332102-1121132231310203-1030002003033303-2132122220231301-2123222000023133-3021132113322011-0111100220223313-1013123013322013"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack`

- [IPv4](resources--fleet--reference--group-004.md#canonical-3330230013310221-2103102113302010-0020322212232321-3322133312210301-3031001321232233-2012133120031011-1301031203022012-0233333323101031): complete subsection reference.

- [IPv6](resources--fleet--reference--group-004.md#canonical-2321102211320030-0223002220011332-1230213310102021-1131031000000012-3001331312132000-2000103313112030-3020012022321232-1101133321231212): complete subsection reference.

<a id="canonical-3330230013310221-2103102113302010-0020322212232321-3322133312210301-3031001321232233-2012133120031011-1301031203022012-0233333323101031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-2323102010131331-0222022123303131-0001320120102230-3002231130201221-0320012123110310-2111211303213332-2101222302031312-1321301311201213)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-0113300103011330-2021221000112230-2120203033320002-0000031323332030-3330320102133210-3331212021003013-3132112030012121-0110211233022031)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](resources--fleet--reference--group-004.md#canonical-3102121312220313-0113030102112222-1312112312222122-2222211021110113-3333023030230230-0211333200121331-3222310113021231-1001121303013130)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-3220103102133331-0223020201122211-0231322223121112-2232333303201211-2123010331323103-1110323203132033-1111101031313110-1112020223121121"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Additional upstream details:

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

<a id="canonical-1210122202322302-1203220322102013-1111302112021330-2120210032012222-1212301031321311-0022033003311323-0123020103230012-1200302101001133"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4`

<a id="canonical-2131013222133111-0000103202032110-1312012023102122-1011211001113302-0013001121333133-3121331020312102-3003333302301111-3320001222000010"></a>

#### `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4.addr` property

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2321102211320030-0223002220011332-1230213310102021-1131031000000012-3001331312132000-2000103313112030-3020012022321232-1101133321231212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-2323102010131331-0222022123303131-0001320120102230-3002231130201221-0320012123110310-2111211303213332-2101222302031312-1321301311201213)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-0113300103011330-2021221000112230-2120203033320002-0000031323332030-3330320102133210-3331212021003013-3132112030012121-0110211233022031)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](resources--fleet--reference--group-004.md#canonical-3102121312220313-0113030102112222-1312112312222122-2222211021110113-3333023030230230-0211333200121331-3222310113021231-1001121303013130)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-2311322200122000-1103223111300201-1130202332130112-0300020321023100-1202333113230012-3211223300002222-3322000313113211-1021011232331312"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

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

<a id="canonical-1122103030123212-2023120323203123-0032333022312302-0103031130202330-2103221002331222-2010020202322002-2003122223222203-3232000101020123"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6`

<a id="canonical-1300111333301232-3323201023033111-3320320023102323-3203023211110120-3000103022133210-0112013322303233-2312332000112102-1123131301302111"></a>

#### `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6.addr` property

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1111103023232130-3012302323132032-1020130021030110-0132010331113111-2021231131202122-3102000212232011-1203002323232112-2333311000032211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-2323102010131331-0222022123303131-0001320120102230-3002231130201221-0320012123110310-2111211303213332-2101222302031312-1321301311201213)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-0113300103011330-2021221000112230-2120203033320002-0000031323332030-3330320102133210-3331212021003013-3132112030012121-0110211233022031)
- storage_static_routes.storage_routes.nexthop.nexthop_address.IPv4

<a id="canonical-2122033231100233-1101100013211121-2200223211023020-1032033202321101-0313020102023121-2231113130011021-1203332002021222-0210303221103002"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Additional upstream details:

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

<a id="canonical-2010330031103210-0201201100330320-1000033233032113-2133020122231222-1011231123321222-2213223202313131-1310222100023320-0030330232312203"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4`

<a id="canonical-1012303303212231-1001212213102130-2113203301232132-3020102201203313-3212300023111111-1203123003211130-2221033102032311-0130101311311131"></a>

#### `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4.addr` property

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3122032102121102-3002123231001213-3101301122202200-2133021320123223-3000333013000223-3020223000310003-2331311202011003-2200100110120102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-2323102010131331-0222022123303131-0001320120102230-3002231130201221-0320012123110310-2111211303213332-2101222302031312-1321301311201213)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-0113300103011330-2021221000112230-2120203033320002-0000031323332030-3330320102133210-3331212021003013-3132112030012121-0110211233022031)
- storage_static_routes.storage_routes.nexthop.nexthop_address.IPv6

<a id="canonical-2320332312310002-2033230031020213-0121233021330300-3220300010333032-1112130000210320-1221313212201232-0320203111030323-1121012333210233"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

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

<a id="canonical-3331103330203321-2122322010121131-1110032033230132-2033202120313221-2122331331212300-3231223033130312-1321103012001130-2312223102002231"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6`

<a id="canonical-0331213132010301-3133333202102211-3003021030202120-0210030330332000-2233302201031323-1200101200023132-0133322010100212-0113122030001032"></a>

#### `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6.addr` property

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0011031323102333-2133013032021320-3023030012331100-2101022003022112-3302133213002111-2211232221112303-0333322213031131-1131101123022222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.subnets` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231)
- storage_static_routes.storage_routes.subnets

<a id="canonical-2110130201003120-3313131233010200-0331032132113333-0013023022000333-3132030321222020-2121121003312322-3023020031111031-1032000302312012"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3112020102310331-1021000212202120-3021033013110330-3100322133230030-0220300211332313-1203132031330020-1022133323120111-2332300222123031"></a>

### Direct properties for `storage_static_routes.storage_routes.subnets`

- [IPv4](resources--fleet--reference--group-004.md#canonical-2132212211202230-1200000233221133-3112130322200231-1321002020330023-3103032102321001-0110133031303221-3331211002321102-0022313200020201): complete subsection reference.

- [IPv6](resources--fleet--reference--group-004.md#canonical-0213122222022133-2200222013101300-3232201130331202-1332011213313032-2132031113010332-0131223200213322-1133133302202300-1132332200010302): complete subsection reference.

<a id="canonical-2132212211202230-1200000233221133-3112130322200231-1321002020330023-3103032102321001-0110133031303221-3331211002321102-0022313200020201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.subnets.ipv4` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231)
- [storage_static_routes.storage_routes.subnets](resources--fleet--reference--group-004.md#canonical-0011031323102333-2133013032021320-3023030012331100-2101022003022112-3302133213002111-2211232221112303-0333322213031131-1131101123022222)
- storage_static_routes.storage_routes.subnets.IPv4

<a id="canonical-3011302323032013-1312301320133210-0232233100310021-1111110203110220-2322203233013123-2201310020221320-2013130031121130-2233201123003333"></a>

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

<a id="canonical-1130332013300022-1231312003022301-2002110330012000-3121212031101023-2300112222021123-2000310021001013-1002001331322232-0033202030021303"></a>

### Direct properties for `storage_static_routes.storage_routes.subnets.ipv4`

<a id="canonical-1030302012123132-0113032112203100-2220012323302113-1211213122023013-0332202332313301-2133003021302302-0100013123323313-1001013011013011"></a>

#### `storage_static_routes.storage_routes.subnets.ipv4.plen` property

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3213201223213112-3221103030011301-2012022031303001-0211302013212103-1112332000102312-0331100323100221-2322012330011223-1110103113023130"></a>

<a id="canonical-3113332231013231-1302102230000232-3211213301202221-1033101200122231-1223332222013003-2200330320333131-0020313303303133-0203101313033211"></a>

#### `storage_static_routes.storage_routes.subnets.ipv4.prefix` property

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0213122222022133-2200222013101300-3232201130331202-1332011213313032-2132031113010332-0131223200213322-1133133302202300-1132332200010302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.subnets.ipv6` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231)
- [storage_static_routes.storage_routes.subnets](resources--fleet--reference--group-004.md#canonical-0011031323102333-2133013032021320-3023030012331100-2101022003022112-3302133213002111-2211232221112303-0333322213031131-1131101123022222)
- storage_static_routes.storage_routes.subnets.IPv6

<a id="canonical-2222033102211111-0100010122212323-2331000212000221-2011103211211320-1021123033102030-0030020002003223-1100203030003323-2201001330030331"></a>

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

<a id="canonical-1010202012010200-2332330230301013-1212012023000001-2011203310102011-1020100132312010-2011020202130111-3321210211203102-1233321110132210"></a>

### Direct properties for `storage_static_routes.storage_routes.subnets.ipv6`

<a id="canonical-1202321020130003-1012231131223213-3032033000222222-1022102012331101-1012203103330302-0213013323331032-1331220011311100-3111203113332221"></a>

#### `storage_static_routes.storage_routes.subnets.ipv6.plen` property

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3103212331223031-3200111123011211-3213213100303030-2002212031131330-3303220100022330-3310310303122132-1331121331323033-1111030113031013"></a>

<a id="canonical-2303221331213022-2002122021021210-0222320032112123-1302300002000131-3131301310232200-2210000213221003-2102103120223332-0331112201230130"></a>

#### `storage_static_routes.storage_routes.subnets.ipv6.prefix` property

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Additional upstream details:

IPv6 address must be specified as hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0"
The address can be compacted by suppressing zeros e.g. "2001:db8::2::"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3013312232300101-1023222303301232-0210002302033021-2230300221300103-1230132310201312-1210112221101000-0010122001311202-1320000331121210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- timeouts

<a id="canonical-0101333303030102-2023323213103220-3233013322301023-3332321203031320-3233333023230103-0230313320020131-0332230220222213-3301221323222031"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1012330202012211-3022013301103013-0110331231101323-3330033112332013-0322001301122123-1312231311300203-2202210302120002-2030111200002331"></a>

### Direct properties for `timeouts`

<a id="canonical-3322312311221323-0230233310031231-1031033013211122-0023211200003131-0132101030112131-2202022233032230-3333022320022223-0002011110130122"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3203133000103123-1222100322131033-1000111213021132-2121011133302120-2032112013132113-1020310210302103-2320112001100022-2323011312311032"></a>

<a id="canonical-2210211200233312-0322210012100233-1230200323320320-1232223220121110-3230103012121123-2013032212330322-3301013210200313-2233333010021231"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1020230301203111-0003103030020313-1023322320302003-2003303113021220-0301112330213002-1303012123012113-0203331001212110-1111302002202302"></a>

<a id="canonical-0300120021300131-3212020012102131-1311103010211013-0133231220031330-3322100112221222-1310330220031031-2013103300102020-1230111221223220"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2211201022321033-2330012313032120-0302200302012020-0133102300233231-1113112031323030-2221111031032122-3101323100013101-0200330001112030"></a>

<a id="canonical-3333102200032202-2012100011100231-3103020221222132-3110301113200131-3300311021123202-3202003133332123-3312131122021020-2310103000123232"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1313133033021303-0322312200331123-0101000033230330-1300123032023012-1022323323332131-3002113220222232-0302220132222213-1022313213313302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `usb_policy` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- usb_policy

<a id="canonical-1122110212330113-3101131220333203-2013203030011212-2310131001313330-1321230331331320-0133123320203112-1102233320013230-2122202121023032"></a>

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
usb_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011313030230232-2021332310120103-3222000030330132-2013220010301320-1011300102021023-2233333103103101-0002230000103023-0130010233213232"></a>

### Direct properties for `usb_policy`

<a id="canonical-0210031313123110-1013201323233220-3321022023222200-2131332103232323-2203032111130102-3313032333332030-0000332002311233-0320333133001220"></a>

#### `usb_policy.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2000203111211212-1122201033323001-3310033331330103-3320330022113132-0300023133131111-0131201313122323-2113011103200123-1002331020122212"></a>

<a id="canonical-0011300300110001-1300113313023110-2020013031213110-0333302033320131-2131320201320113-0213223331210331-2301302132120223-3133033013300130"></a>

#### `usb_policy.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1322013111303300-3212302013200320-3303113031311132-2220211312101300-3130101033211130-2010232231023103-3230111312312322-2223111210320202"></a>

<a id="canonical-2321002121312331-0100023330212123-2301033021003333-2020203003022013-2102022320133001-0120332212033002-1222332203300300-1310013300022332"></a>

#### `usb_policy.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
