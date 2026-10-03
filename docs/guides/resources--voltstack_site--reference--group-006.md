---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

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

- [custom_storage_config.static_routes.static_routes.default_gateway](resources--voltstack_site--reference--group-006.md#canonical-0220110121113211-0123200320311213-2033231212131113-0101303012033002-0331200212333220-0100030201202310-0203300222031112-1212131332212032)
- [custom_storage_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-006.md#canonical-2103010001332321-3000301122011331-2103133111021230-0301232203102112-3333122212301320-0010123301123323-1031033201333332-2031303213320033)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0010013332203001-0113132310200201-1011031031323020-3130002332300123-3110122210200111-1330032111302123-2100222030001223-3213031122112122)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0220110121113211-0123200320311213-2033231212131113-0101303012033002-0331200212333220-0100030201202310-0203300222031112-1212131332212032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302300023002311-3023312222021023-0013021232000331-0302103223333233-0332220010300202-3303131011222013-0103302022200133-0021212233311202"></a>

## custom_storage_config.static_routes.static_routes.default_gateway — default_gateway / 122020333100 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0010013332203001-0113132310200201-1011031031323020-3130002332300123-3110122210200111-1330032111302123-2100222030001223-3213031122112122)
- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-2300310110003322-0200012021220130-2123203021230230-2303022333110233-2313312311210211-0232003100312110-2230131203001132-2203002101122230)
- custom_storage_config.static_routes.static_routes.default_gateway

<a id="canonical-3003110103233320-1302113330112201-3333213012123203-2113013221332323-2120221012011323-0223213003200200-3023323230003211-3031031210021322"></a>

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

<a id="canonical-2303231301023232-3330120301011211-3332020003310203-0330101233003321-3231132312320101-0221213321101023-0330210103110333-0210201232211132"></a>

## Direct properties — default_gateway / 122020333100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211210101301120-2031120103310112-2033302001032123-2101302212131223-3133221323003222-3101032033103330-2220211213110302-0001203330323230"></a>

## Next pages — default_gateway / 122020333100 / 4

- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-2300310110003322-0200012021220130-2123203021230230-2303022333110233-2313312311210211-0232003100312110-2230131203001132-2203002101122230)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2103010001332321-3000301122011331-2103133111021230-0301232203102112-3333122212301320-0010123301123323-1031033201333332-2031303213320033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103212203321211-2011121310120231-2003233032330322-2311010213111322-0002102203301333-0131133230121032-3011013001302012-2033220312113303"></a>

## custom_storage_config.static_routes.static_routes.node_interface — node_interface / 300302012310 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0010013332203001-0113132310200201-1011031031323020-3130002332300123-3110122210200111-1330032111302123-2100222030001223-3213031122112122)
- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-2300310110003322-0200012021220130-2123203021230230-2303022333110233-2313312311210211-0232003100312110-2230131203001132-2203002101122230)
- custom_storage_config.static_routes.static_routes.node_interface

<a id="canonical-3012120100101202-2100132100223002-1131332222311112-0030000012222321-1211202033031020-0211012303300322-3023121301312022-0123122330130302"></a>

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

<a id="canonical-3212201211300211-3001332230011323-3000110202000300-1003303121230113-1003001022232121-1000321011203022-0132211002013100-2022103202022010"></a>

## Direct properties — node_interface / 300302012310 / 3

- [list](resources--voltstack_site--reference--group-006.md#canonical-2332202310310112-0123021022031011-2112210001013103-3202111132133020-1323012322132231-3020123130302100-3030311202333000-3001011013300321): complete subsection reference.

<a id="canonical-3331120202103201-1121021303100122-3310003233103211-3121132013301011-2011120003013032-2323011311303110-3313220123112011-2202000102302211"></a>

## Next pages — node_interface / 300302012310 / 4

- [custom_storage_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-006.md#canonical-2332202310310112-0123021022031011-2112210001013103-3202111132133020-1323012322132231-3020123130302100-3030311202333000-3001011013300321)
- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-2300310110003322-0200012021220130-2123203021230230-2303022333110233-2313312311210211-0232003100312110-2230131203001132-2203002101122230)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2332202310310112-0123021022031011-2112210001013103-3202111132133020-1323012322132231-3020123130302100-3030311202333000-3001011013300321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231013320232120-3131322021010021-3123311023222021-2023123003120223-3220332211310031-1101122102121333-2320301001323110-3322213023332101"></a>

## custom_storage_config.static_routes.static_routes.node_interface.list — list / 203112221213 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0010013332203001-0113132310200201-1011031031323020-3130002332300123-3110122210200111-1330032111302123-2100222030001223-3213031122112122)
- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-2300310110003322-0200012021220130-2123203021230230-2303022333110233-2313312311210211-0232003100312110-2230131203001132-2203002101122230)
- [custom_storage_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-006.md#canonical-2103010001332321-3000301122011331-2103133111021230-0301232203102112-3333122212301320-0010123301123323-1031033201333332-2031303213320033)
- custom_storage_config.static_routes.static_routes.node_interface.list

<a id="canonical-2033310333000312-3323202102033312-1330232313030011-3133101021332103-3023101201331102-1101112223121323-2202313310000331-3321101213112332"></a>

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

<a id="canonical-2103011000001303-1123221100221100-0231331110310013-2202122303211210-1131313000230221-0021023031132022-2331001211013000-3122201122131131"></a>

## Direct properties — list / 203112221213 / 3

- [interface](resources--voltstack_site--reference--group-006.md#canonical-1123213133320103-2122120332332123-0222021133100210-3120032131003220-2320002310331013-3120101311300001-3233301110333130-1313003232333231): complete subsection reference.

<a id="canonical-1303100232010002-2231112031311310-2032232013213113-0221013311112011-1020213330312033-1320132200311322-1320111231312232-3121302021031232"></a>

<a id="canonical-2123131031223303-2210120023331011-2202031120312230-3332001102122312-0133303130032130-0022223130013003-3102130122113020-2200120013000123"></a>

## node property — list / 203112221213 / 4

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

<a id="canonical-1111131022133103-3122110221230130-2131103231023233-1100232330323311-3002213301321033-0023212322102100-2233130030212212-2100310103013001"></a>

## Next pages — list / 203112221213 / 5

- [custom_storage_config.static_routes.static_routes.node_interface.list.interface](resources--voltstack_site--reference--group-006.md#canonical-1123213133320103-2122120332332123-0222021133100210-3120032131003220-2320002310331013-3120101311300001-3233301110333130-1313003232333231)
- [custom_storage_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-006.md#canonical-2103010001332321-3000301122011331-2103133111021230-0301232203102112-3333122212301320-0010123301123323-1031033201333332-2031303213320033)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1123213133320103-2122120332332123-0222021133100210-3120032131003220-2320002310331013-3120101311300001-3233301110333130-1313003232333231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102133022303321-2210001133322321-3310033230311212-0110012322002212-2120112231123222-1203011200021322-2313323033133121-3321120101133310"></a>

## custom_storage_config.static_routes.static_routes.node_interface.list.interface — interface / 330203011332 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.static_routes](resources--voltstack_site--reference--group-005.md#canonical-0010013332203001-0113132310200201-1011031031323020-3130002332300123-3110122210200111-1330032111302123-2100222030001223-3213031122112122)
- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--reference--group-005.md#canonical-2300310110003322-0200012021220130-2123203021230230-2303022333110233-2313312311210211-0232003100312110-2230131203001132-2203002101122230)
- [custom_storage_config.static_routes.static_routes.node_interface](resources--voltstack_site--reference--group-006.md#canonical-2103010001332321-3000301122011331-2103133111021230-0301232203102112-3333122212301320-0010123301123323-1031033201333332-2031303213320033)
- [custom_storage_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-006.md#canonical-2332202310310112-0123021022031011-2112210001013103-3202111132133020-1323012322132231-3020123130302100-3030311202333000-3001011013300321)
- custom_storage_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-2301221032322321-3230033121031223-0130313320332102-2303200113313231-2330320132020120-0131203031221331-2111111332330303-1320333221313321"></a>

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

<a id="canonical-3332030100210133-0212223002312301-3003120131123000-2220223132023010-0012122100123332-1031032232322231-2313302111302002-1100131210121312"></a>

## Direct properties — interface / 330203011332 / 3

<a id="canonical-2023113011233333-3202113221230231-2331221313113130-0222130002300201-1302203023123020-2202002301100211-3221133101221012-3302223320001032"></a>

<a id="canonical-2003323121210200-0132130323021111-1021311033300130-3010220002202123-1132331332001200-0032012331321031-0230312002103033-3010110030003212"></a>

## kind property — interface / 330203011332 / 4

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

<a id="canonical-0010132021310020-0101230222301101-0000323010013200-0013232123023332-1121212110301332-3120102223000300-2312232230210212-2022100033323102"></a>

<a id="canonical-3312110133031320-0300220230132321-1222120013031003-0000222300221023-0130032031033033-1010113021300220-1311222020102331-3010032330130220"></a>

## name property — interface / 330203011332 / 5

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

<a id="canonical-0013001212321313-2231332031312111-2323122212200200-0300323110133201-3230333230130121-1131132221303130-0300213210130120-0001020202232332"></a>

<a id="canonical-0010211020212301-2323202331311023-1131021023020330-1121002332120232-0123202011331200-0012211330222002-2220231021123112-3230320322100232"></a>

## namespace property — interface / 330203011332 / 6

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

<a id="canonical-1232110201101210-2230303223323233-1120102131300032-3301133213122012-2101033021330200-3002310003222012-2031321001133333-2333230002312011"></a>

<a id="canonical-0110110111020131-0120330002221310-1221332222223133-3311013231111020-0113111210113022-0102310020023112-0331301232332322-1113023003323212"></a>

## tenant property — interface / 330203011332 / 7

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

<a id="canonical-2112020133121203-0211033111101323-2203230311010311-1332210023233033-2201021002332112-2010111110323103-3003122312130200-3111013121232120"></a>

<a id="canonical-2320122102211312-1311231120001010-3302101212203302-1131321002310031-0023003033310130-2103103210232002-2001232331120120-3310331023222012"></a>

## uid property — interface / 330203011332 / 8

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

<a id="canonical-2232130221102001-1032231121131311-0310130212323201-0201033011110310-2033021301103230-1001112033023010-1102233220312210-3133110020021232"></a>

## Next pages — interface / 330203011332 / 9

- [custom_storage_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--reference--group-006.md#canonical-2332202310310112-0123021022031011-2112210001013103-3202111132133020-1323012322132231-3020123130302100-3030311202333000-3001011013300321)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3200223201220132-0311220310300003-2103003133003311-1323333023323123-3011330010121132-2203203012313212-2030201103311030-0213120200223032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300321301323202-0030331132030110-3210232231332112-0132103232011222-0110212333010323-1121303302113312-3210303131311211-1300333310103112"></a>

## custom_storage_config.storage_class_list — storage_class_list / 133220303120 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- custom_storage_config.storage_class_list

<a id="canonical-2113011302032212-0032101122112311-1002113330210132-1002113011003201-1223200201123123-3310030211030230-1200222003331030-1331301002201203"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this fleet.

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

<a id="canonical-1023011130010020-2003201312001221-0122103221322330-1321303000023023-1131100221102131-1121012130331231-3211030101000331-0001032101100133"></a>

## Direct properties — storage_class_list / 133220303120 / 3

- [storage_classes](resources--voltstack_site--reference--group-006.md#canonical-3300313120223201-3231110231130201-3211021103013321-1022230203211310-3133011130131233-0203022112023110-1211220112313103-1133323202301233): complete subsection reference.

<a id="canonical-2232202022130031-0133232021013031-3102121000030113-3010011231210303-2002321122023213-0013020303030112-1030002032112222-2221310320112223"></a>

## Next pages — storage_class_list / 133220303120 / 4

- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-006.md#canonical-3300313120223201-3231110231130201-3211021103013321-1022230203211310-3133011130131233-0203022112023110-1211220112313103-1133323202301233)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3300313120223201-3231110231130201-3211021103013321-1022230203211310-3133011130131233-0203022112023110-1211220112313103-1133323202301233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003232112031330-1123230112310113-1122120312030013-2033301002332112-1030032002330331-2313133111122122-1333031123231301-2222113332012010"></a>

## custom_storage_config.storage_class_list.storage_classes — storage_classes / 311101120300 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-006.md#canonical-3200223201220132-0311220310300003-2103003133003311-1323333023323123-3011330010121132-2203203012313212-2030201103311030-0213120200223032)
- custom_storage_config.storage_class_list.storage_classes

<a id="canonical-3102120000313213-0121302031023001-3221020033132311-0111131322303030-1023213311333221-3031211130110101-3321211233120032-2010121223103232"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_class_name",
    "storage_device"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "hpe_storage"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("netapp_trident",
    "pure_service_orchestrator")}
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

Terraform syntax:

```terraform
storage_classes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323333112311132-0011100131002000-1022021200312303-0322222022232223-3110322321330130-1320003112001130-2303212020313033-0303303311211131"></a>

## Direct properties — storage_classes / 311101120300 / 3

<a id="canonical-1230033032322320-1331003213320200-3302313310220223-2010213022111221-2322220311101112-1311313122021000-0103101212111002-0101223321100312"></a>

<a id="canonical-3010031320111010-0233031312201031-3313223102231121-3223002212202223-1322232231200210-1200201033100223-2102033011101203-1321202001012300"></a>

## advanced_storage_parameters property — storage_classes / 311101120300 / 4

Type: `["map", "string"]`. Optional.

Advanced Parameters. Map of parameter name and string value.

Upstream description:

Map of parameter name and string value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
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
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0021102132022130-2211033232122113-3311013322331201-0322032202131322-1002311322112321-1212330231300123-0100320310320120-1020231021200011"></a>

<a id="canonical-1220221221130223-2221022131332213-3301111231012120-2211110231301311-3223131331200000-3033213002022302-0320132131003211-0212223101131211"></a>

## allow_volume_expansion property — storage_classes / 311101120300 / 5

Type: `"bool"`. Optional.

Allow Volume Expansion. Allow volume expansion.

Upstream description:

Allow volume expansion.

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

- [custom_storage](resources--voltstack_site--reference--group-006.md#canonical-1021332220311002-1101203313100120-1121201211311303-3310101003111012-1330232303030220-1303113102013032-2322100121010323-3213230321310120): complete subsection reference.

<a id="canonical-1330112103031133-2002213010222110-0131020030330123-1323123100313203-0202100213313100-3100211001230301-2021132220123202-3211222201230001"></a>

<a id="canonical-1232232020012131-0031322000233012-1323213022200032-2323130232311012-3111202323113302-3011300023100202-0310301222322100-1001311323233020"></a>

## default_storage_class property — storage_classes / 311101120300 / 6

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

<a id="canonical-3313022213100222-1033001021331231-0021000333012230-1103210120112110-3120000303002023-0211321300133012-0122223013332203-3123231103321031"></a>

<a id="canonical-0120312131332303-1202332012333332-2310311232312333-0110320121032302-3232213331231013-3232332130000221-3033201010131021-3131021122110110"></a>

## description_spec property — storage_classes / 311101120300 / 7

Type: `"string"`. Optional.

Storage Class Description. Description for this storage class.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-2032132303130012-3231033313311200-2003222212102222-1002100021103100-3121212221021223-0212102131230331-3212332031233032-3320321032323102): complete subsection reference.

- [netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-3010213222022331-2313320003203232-3130302120220302-0130222212302333-2103232323303223-2210112231333120-1311222210032110-3323312220122300): complete subsection reference.

- [pure_service_orchestrator](resources--voltstack_site--reference--group-006.md#canonical-2013033101131001-2020332001123121-0301320113203232-1310233032132303-0031122121121110-0300100302023122-3200301322122133-2320202010323200): complete subsection reference.

<a id="canonical-2330000310121132-1133112121331103-3020200211103022-0203112001301032-0323130301131002-1213323010311212-2211131200301232-0122323200310212"></a>

<a id="canonical-2120032202310130-0110120101300323-3310312000232022-3002302220332321-1022320003303323-1303223233022013-2132103301131011-1023013330321000"></a>

## reclaim_policy property — storage_classes / 311101120300 / 8

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Reclaim Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 16,
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
    "ves.io.schema.rules.string.max_len": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16"
  }
}
```

<a id="canonical-2000103120232320-3001100003113232-3032311301023110-2021300331211101-3113331200101131-0121220032210202-1033022100110311-2011001122023230"></a>

<a id="canonical-0021321212233023-3200100022021302-3320212132331011-2312300102110130-0203310311200201-2331132112103132-3201300010211130-1200310023303231"></a>

## storage_class_name property — storage_classes / 311101120300 / 9

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0210220312032120-0210230122000112-2323222013121200-0320300220232012-2230120212231310-3132332102131010-1300122013131322-1233101231230022"></a>

<a id="canonical-2020123311202030-2123213121122220-0020003231110011-1201222112021332-2003303110301223-1010112333123000-3131113031211112-3211203122202202"></a>

## storage_device property — storage_classes / 311101120300 / 10

Type: `"string"`. Optional.

Storage device that this class will use. The Device name defined at previous step.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3310200020330011-2100221011000133-0100302123110001-0200031232102120-3023320200002310-2033122010103221-3021113333220233-2012023232112013"></a>

## Next pages — storage_classes / 311101120300 / 11

- [custom_storage_config.storage_class_list.storage_classes.custom_storage](resources--voltstack_site--reference--group-006.md#canonical-1021332220311002-1101203313100120-1121201211311303-3310101003111012-1330232303030220-1303113102013032-2322100121010323-3213230321310120)
- [custom_storage_config.storage_class_list.storage_classes.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-2032132303130012-3231033313311200-2003222212102222-1002100021103100-3121212221021223-0212102131230331-3212332031233032-3320321032323102)
- [custom_storage_config.storage_class_list.storage_classes.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-3010213222022331-2313320003203232-3130302120220302-0130222212302333-2103232323303223-2210112231333120-1311222210032110-3323312220122300)
- [custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator](resources--voltstack_site--reference--group-006.md#canonical-2013033101131001-2020332001123121-0301320113203232-1310233032132303-0031122121121110-0300100302023122-3200301322122133-2320202010323200)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-006.md#canonical-3200223201220132-0311220310300003-2103003133003311-1323333023323123-3011330010121132-2203203012313212-2030201103311030-0213120200223032)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1021332220311002-1101203313100120-1121201211311303-3310101003111012-1330232303030220-1303113102013032-2322100121010323-3213230321310120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331013201013331-3310021010302313-3101010212111130-3121222002110032-1201223102330112-1000133232031312-1332112022322331-0330111311132332"></a>

## custom_storage_config.storage_class_list.storage_classes.custom_storage — custom_storage / 221133231111 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-006.md#canonical-3200223201220132-0311220310300003-2103003133003311-1323333023323123-3011330010121132-2203203012313212-2030201103311030-0213120200223032)
- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-006.md#canonical-3300313120223201-3231110231130201-3211021103013321-1022230203211310-3133011130131233-0203022112023110-1211220112313103-1133323202301233)
- custom_storage_config.storage_class_list.storage_classes.custom_storage

<a id="canonical-2233032201232023-3203212102011003-3213100323012011-1331303121113310-2122002120133322-0223102032303201-2011003312102031-2212130330232121"></a>

Type: `"object"`. single nested block, Optional.

Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into
given site.

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
custom_storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202123032230030-2312221310131313-1221000121302220-3300031003332122-1222303002222113-1212320232320133-0200232312203101-3313102333130122"></a>

## Direct properties — custom_storage / 221133231111 / 3

<a id="canonical-2122210312312233-0113132130313213-2210210320323300-0021313213332202-1122221111231032-2302021011230131-3312212203013220-0122200213031312"></a>

<a id="canonical-1130012111113010-2111331000003110-0031230233321012-1032123220010111-0212013103031103-0330000222230221-0200201121022002-3020223233210322"></a>

## YAML property — custom_storage / 221133231111 / 4

Type: `"string"`. Optional.

Storage Class YAML. K8s YAML for StorageClass.

Upstream description:

K8s YAML for StorageClass.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "Valid parseable YAML",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "validation": {
      "customRule": "Must be valid YAML"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0321322002122021-2013031022031023-2032032220230001-2321303330201010-0113212011010333-2312030333113120-3210322231220222-0003331002203301"></a>

## Next pages — custom_storage / 221133231111 / 5

- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-006.md#canonical-3300313120223201-3231110231130201-3211021103013321-1022230203211310-3133011130131233-0203022112023110-1211220112313103-1133323202301233)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2032132303130012-3231033313311200-2003222212102222-1002100021103100-3121212221021223-0212102131230331-3212332031233032-3320321032323102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211322031230101-0300012121011101-1200021122111111-0012310020322333-2232312202321232-3103223313320032-2103202302330100-3300112220102023"></a>

## custom_storage_config.storage_class_list.storage_classes.hpe_storage — hpe_storage / 333132320230 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-006.md#canonical-3200223201220132-0311220310300003-2103003133003311-1323333023323123-3011330010121132-2203203012313212-2030201103311030-0213120200223032)
- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-006.md#canonical-3300313120223201-3231110231130201-3211021103013321-1022230203211310-3133011130131233-0203022112023110-1211220112313103-1133323202301233)
- custom_storage_config.storage_class_list.storage_classes.hpe_storage

<a id="canonical-1023202320032331-2121311121110120-0313203133211012-1102303132220123-3302002112021213-0010101000232210-1311120030332113-2022331010121330"></a>

Type: `"object"`. single nested block, Optional.

Storage class Device configuration for HPE Storage.

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
hpe_storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013103031011310-1013120312333000-3212332113112220-3123133122000120-0331110313200330-2320023101002111-3030030323331031-1133332232310202"></a>

## Direct properties — hpe_storage / 333132320230 / 3

<a id="canonical-1213211103231320-2111221013032301-2213131202022130-1101221320311111-2310332311333011-2203332333122111-2322320313332003-0231323323123003"></a>

<a id="canonical-3122122322121103-3012221302211230-3130200133313220-1121230103003022-1120230303113323-3320110111300311-3031200221300300-2220031223231310"></a>

## allow_mutations property — hpe_storage / 333132320230 / 4

Type: `"string"`. Optional.

Mutation can override specified parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1101332100212033-0013003222003103-2232302332101133-3010223023110331-3312221333302210-0103222203121212-0133303210011303-3333203031301323"></a>

<a id="canonical-2012321331103333-2313022332022230-3132323220000212-3111032121112313-3003302321133321-3323111320132110-0212022321120023-2012322223023102"></a>

## allow_overrides property — hpe_storage / 333132320230 / 5

Type: `"string"`. Optional.

AllowOverrides. PVC can override specified parameters.

Upstream description:

PVC can override specified parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1122000333120003-2232230110031330-2102102311102210-3101201310021003-2230113012301302-0101002100302030-1111033322112032-1313033212131031"></a>

<a id="canonical-0102232013300002-1021102233032121-0310003000221320-3021113020011300-3130102311131231-0321032110112313-0332223321031002-0001013222133100"></a>

## dedupe_enabled property — hpe_storage / 333132320230 / 6

Type: `"bool"`. Optional.

Indicates that the volume should enable deduplication.

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

<a id="canonical-1020030300232023-1113022331323033-2233312211130033-1312121121111123-0110231213021023-3032012301212002-1012210210310220-2302313133032213"></a>

<a id="canonical-0121112210010303-3221121120210210-3303221312221221-0201013332113020-1311303101013321-0023212333230001-0100033013233210-1331032201011311"></a>

## description_spec property — hpe_storage / 333132320230 / 7

Type: `"string"`. Optional.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

<a id="canonical-3112032002310033-3033010321220302-3301211110031003-3031032312231212-0031210302022013-1221130211200232-3020000110120021-0011310132011111"></a>

<a id="canonical-1031322102121133-2303331022200303-2200333221232002-2333312213210112-1322131203332231-1020021102123230-2111101202200303-0203120023332113"></a>

## destroy_on_delete property — hpe_storage / 333132320230 / 8

Type: `"bool"`. Optional.

Indicates the backing Nimble volume (including snapshots) should be destroyed when the PVC is
deleted.

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

<a id="canonical-3113102321000312-1322031311333203-0011221322022022-0230131312301122-0123001311012112-1011223322103332-0322232302131323-2303221223130013"></a>

<a id="canonical-1302033322032113-3022202010300333-0120131020120301-2230321111302303-2222032303121302-3122103031233210-3021303122002012-1030323222133232"></a>

## encrypted property — hpe_storage / 333132320230 / 9

Type: `"bool"`. Optional.

Indicates that the volume should be encrypted.

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

<a id="canonical-1120103233122031-0323010011202132-3212023302112310-1322322103323202-0332331221110011-2011210021312333-2312001101012121-3300322300110232"></a>

<a id="canonical-3110011032331321-3320113231332032-1123232020300001-0132132011312333-2101102000310211-2200313211330013-2013033030320202-3010220001312132"></a>

## folder property — hpe_storage / 333132320230 / 10

Type: `"string"`. Optional.

The name of the folder in which to place the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0133021233323313-2210010113013201-1123221123133332-2003300212030100-2010032030332131-3301331213122030-0000331333331203-0312320113233313"></a>

<a id="canonical-3221312203302121-3011021033230330-2030113010221211-0232212101103331-3221333121132001-0221302000121310-0323313022331021-3000301131332031"></a>

## limit_iops property — hpe_storage / 333132320230 / 11

Type: `"string"`. Optional.

LimitIops. The IOPS limit of the volume.

Upstream description:

The IOPS limit of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "int64",
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

<a id="canonical-1222102211312022-3230022311332121-3322211110202211-1323202233301231-3120311210013120-3003013221002011-3013111030112202-1111332132130103"></a>

<a id="canonical-3031320203302012-0300230101321103-2003322323303013-1301230121010021-1311331310122101-3023213323010131-2301111213230233-1203013231230321"></a>

## limit_mbps property — hpe_storage / 333132320230 / 12

Type: `"string"`. Optional.

LimitMbps. The IOPS limit of the volume.

Upstream description:

The IOPS limit of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "int64",
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

<a id="canonical-2333202001022031-0212222211223131-1201232022321002-2311303032001021-0321311302331200-2110122331203333-1022030131112003-3102331210011331"></a>

<a id="canonical-0001221220110131-3033301101113311-3021320001100230-0300102323011133-3301002232121222-0301130301111202-3001211333010202-3220210011223120"></a>

## performance_policy property — hpe_storage / 333132320230 / 13

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

The name of the performance policy to assign to the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0303300011013301-0333013222000203-1213023030302200-3211131202100210-2022333123331023-2203110301003020-3200321200330120-2301321232022111"></a>

<a id="canonical-2320031201122102-3101220122100030-1311223233021233-2022010211130010-1113020201122023-1002033111330203-1320000000220222-2120230122312221"></a>

## pool property — hpe_storage / 333132320230 / 14

Type: `"string"`. Optional.

The name of the pool in which to place the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2310121313031313-3322201310230120-0320020213023021-3231031200223102-1001001130101122-1201103310331302-3331220101020001-2312001301301030"></a>

<a id="canonical-3132211113031203-0310111013203220-3003212333123200-3000001103330333-3222203202000212-3001333230202031-0131301013220201-3101033220300110"></a>

## protection_template property — hpe_storage / 333132320230 / 15

Type: `"string"`. Optional.

The name of the performance policy to assign to the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3011013332000113-0331323233033222-3133013030020233-1203310210312121-3310003121203123-2103301111000222-1321001033122311-2003302333030100"></a>

<a id="canonical-0031210033213013-0202123002023200-3312010132210130-1231303222310233-2211210002032100-1001330100030231-1330120121201012-1231222012133111"></a>

## secret_name property — hpe_storage / 333132320230 / 16

Type: `"string"`. Optional.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2013030321321111-0030031313321313-2012310033212302-1100100313022233-1122212032231332-2120200032332033-1320323233303301-2012121202131222"></a>

<a id="canonical-2132223002231202-3133222010311302-1123310133201111-1222131023111330-1202001121013301-3201002232202230-0331110113231122-1012332132100310"></a>

## secret_namespace property — hpe_storage / 333132320230 / 17

Type: `"string"`. Optional.

The SecretNamespace parameter is used to identify name of namespace where secret resides.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1323210303202333-2312031022130102-1100111121212232-3001332020333131-1300013333010221-0302032130011211-2112321320303133-3123222333323131"></a>

<a id="canonical-3132010230213232-2230311210023212-0121012033320222-2003212022031103-3313033023300000-2210033022011111-0001212332000120-2311311211211002"></a>

## sync_on_detach property — hpe_storage / 333132320230 / 18

Type: `"bool"`. Optional.

Indicates that a snapshot of the volume should be synced to the replication partner each time it is
detached from a node.

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

<a id="canonical-0330020111012223-0313031220100321-1330031132330012-1231022331201003-3122333102313010-1130020003031103-0131110222033320-1202112200032022"></a>

<a id="canonical-3103021102101303-1120000123320302-2022311332311121-1010103103133002-2233203323200132-2131012033303231-1032133222330012-3221222112022032"></a>

## thick property — hpe_storage / 333132320230 / 19

Type: `"bool"`. Optional.

Indicates that the volume should be thick provisioned.

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

<a id="canonical-0331021112102033-2122013120012231-0103121131320012-1101002231032300-0120322002202331-1133000203113110-2231212021302311-1303223333121212"></a>

## Next pages — hpe_storage / 333132320230 / 20

- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-006.md#canonical-3300313120223201-3231110231130201-3211021103013321-1022230203211310-3133011130131233-0203022112023110-1211220112313103-1133323202301233)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3010213222022331-2313320003203232-3130302120220302-0130222212302333-2103232323303223-2210112231333120-1311222210032110-3323312220122300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010300322000133-0221033221020032-0203033130301230-0321223220020103-1012100201212212-3303131102020132-1221030313011003-2211121211001021"></a>

## custom_storage_config.storage_class_list.storage_classes.netapp_trident — netapp_trident / 322021212313 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-006.md#canonical-3200223201220132-0311220310300003-2103003133003311-1323333023323123-3011330010121132-2203203012313212-2030201103311030-0213120200223032)
- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-006.md#canonical-3300313120223201-3231110231130201-3211021103013321-1022230203211310-3133011130131233-0203022112023110-1211220112313103-1133323202301233)
- custom_storage_config.storage_class_list.storage_classes.netapp_trident

<a id="canonical-3322133201120333-0330130332122133-0222331220111122-2300303323202113-2301311222202310-1113110102103021-1100131232030020-1032321301201002"></a>

Type: `"object"`. single nested block, Optional.

Storage class Device configuration for NetApp Trident.

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
netapp_trident {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013313203332221-0232333223101222-0213020032221321-1213303133130230-0221332130001321-2030111320110000-0121332323010323-3031312321320303"></a>

## Direct properties — netapp_trident / 322021212313 / 3

- [selector](resources--voltstack_site--reference--group-006.md#canonical-1123220132112313-2013212322310210-1002130031032333-1222332033103322-0000122030223111-2201133302211310-2010320032313333-1003233032311000): complete subsection reference.

<a id="canonical-2112301302122310-0003331321130202-2200003310310300-3301202022123133-3132321123133332-1210000122002031-1322222311011101-2233331322103330"></a>

<a id="canonical-0232110110001301-2131322300233330-2312310222203321-3002023210233112-0322201221023203-1013133101212332-2131321220212110-1000112003321000"></a>

## storage_pools property — netapp_trident / 322021212313 / 4

Type: `"string"`. Optional.

The storagePools parameter is used to further restrict the set of pools that match any specified
attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-0131212103001032-3102322112033121-0011200311131031-0211012231313323-3110113000311223-2021010103033303-2220121312013312-3032003123332001"></a>

## Next pages — netapp_trident / 322021212313 / 5

- [custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector](resources--voltstack_site--reference--group-006.md#canonical-1123220132112313-2013212322310210-1002130031032333-1222332033103322-0000122030223111-2201133302211310-2010320032313333-1003233032311000)
- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-006.md#canonical-3300313120223201-3231110231130201-3211021103013321-1022230203211310-3133011130131233-0203022112023110-1211220112313103-1133323202301233)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1123220132112313-2013212322310210-1002130031032333-1222332033103322-0000122030223111-2201133302211310-2010320032313333-1003233032311000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323100001021100-0220032322133222-0220210231311231-2220112013323003-0212330230011131-1300302311033021-3121001323220202-3003303133212122"></a>

## custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector — selector / 220303213232 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-006.md#canonical-3200223201220132-0311220310300003-2103003133003311-1323333023323123-3011330010121132-2203203012313212-2030201103311030-0213120200223032)
- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-006.md#canonical-3300313120223201-3231110231130201-3211021103013321-1022230203211310-3133011130131233-0203022112023110-1211220112313103-1133323202301233)
- [custom_storage_config.storage_class_list.storage_classes.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-3010213222022331-2313320003203232-3130302120220302-0130222212302333-2103232323303223-2210112231333120-1311222210032110-3323312220122300)
- custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector

<a id="canonical-0113013313033122-2012203232110223-1113230221131011-1011302222020131-2301302302031020-0131110113023112-2201130222033313-0131100100203310"></a>

Type: `"object"`. single nested block, Optional.

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

Upstream description:

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

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
selector {}
```

<a id="canonical-0003223320322010-1121000130222103-0232011003011200-2010213031023212-2303001120133323-3110230330203210-0323303222002133-0120030313000312"></a>

## Direct properties — selector / 220303213232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101110223102233-0302203133001100-0332303202013210-3302110331233302-2023121320003333-3111010211122022-2123222331110020-0233230103032100"></a>

## Next pages — selector / 220303213232 / 4

- [custom_storage_config.storage_class_list.storage_classes.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-3010213222022331-2313320003203232-3130302120220302-0130222212302333-2103232323303223-2210112231333120-1311222210032110-3323312220122300)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2013033101131001-2020332001123121-0301320113203232-1310233032132303-0031122121121110-0300100302023122-3200301322122133-2320202010323200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132333033010321-0302020331022311-0222113000311001-2313201001303130-1230133031010101-0021013102131031-1210030322010010-2300123311211132"></a>

## custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator — pure_service_orchestrator / 210120203230 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-006.md#canonical-3200223201220132-0311220310300003-2103003133003311-1323333023323123-3011330010121132-2203203012313212-2030201103311030-0213120200223032)
- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-006.md#canonical-3300313120223201-3231110231130201-3211021103013321-1022230203211310-3133011130131233-0203022112023110-1211220112313103-1133323202301233)
- custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator

<a id="canonical-3203233230320012-2200300123031331-2301111323123303-1122011323332120-2223312121030201-2310201202210120-0121100321321321-3111101330200322"></a>

Type: `"object"`. single nested block, Optional.

Storage class Device configuration for Pure Service Orchestrator.

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
pure_service_orchestrator {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122000030031321-0022113121002012-0222330312313312-0322132230103330-2133020100311110-0023231303203211-3321222320301221-2323310333113002"></a>

## Direct properties — pure_service_orchestrator / 210120203230 / 3

<a id="canonical-2000112130001320-3222201301332102-0313302220230302-3130102220222222-0231230010033010-3131120110301200-1001013032130300-3311113112311231"></a>

<a id="canonical-3231131110130031-3032331033022001-1322311333300211-0112000001033310-3000010003123130-0102210231020222-3202203201010031-0113112110310113"></a>

## backend property — pure_service_orchestrator / 210120203230 / 4

Type: `"string"`. Optional.

\[Enum: block|file\] Defines type of Pure storage backend block or file. The volume will have the
aspects defined in the chosen virtual pool. Possible values are \`block\`, \`file\`.

Upstream description:

Defines type of Pure storage backend block or file. The volume will have the aspects defined in the
chosen virtual pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("block",
    "file"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "block",
    "file"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  }
}
```

<a id="canonical-3212231031112002-2000301011100130-2013303020310311-0012103203313320-3021322323001213-3030213221332131-1031320000101232-2012330311222131"></a>

<a id="canonical-1133212123121220-1231310130033313-3102131332130332-2000132131133330-3222022323001023-3322302123033220-3230300123032223-0121213232013113"></a>

## bandwidth_limit property — pure_service_orchestrator / 210120203230 / 5

Type: `"string"`. Optional.

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Upstream description:

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(12),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 12,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 12,
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
    "ves.io.schema.rules.string.max_len": "12"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "12"
  }
}
```

<a id="canonical-0133201010332123-3120111310023231-2321233122131110-2210203130331333-3032113330213122-3120012203022130-0123132100302321-0021332023033022"></a>

<a id="canonical-1230202301312020-0213232223102111-3223231003201302-3101213300333321-2330101023023103-0102012300133132-3023020020013010-3331002001321021"></a>

## iops_limit property — pure_service_orchestrator / 210120203230 / 6

Type: `"number"`. Optional.

Enable IOPS limitation. It must be between 100 and 100 million. If value is 0, IOPS limit is not
defined.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 100, Maximum: 100000000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100000000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  }
}
```

<a id="canonical-2320010321112333-2200220113130123-0030213230213012-1010220330121102-0312112100312101-0033111222313021-2020021020323221-0230210100022300"></a>

## Next pages — pure_service_orchestrator / 210120203230 / 7

- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-006.md#canonical-3300313120223201-3231110231130201-3211021103013321-1022230203211310-3133011130131233-0203022112023110-1211220112313103-1133323202301233)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013012301322130-3230310232011010-3232220001002313-0011330122031302-1222303113323012-3130001010320010-0000012301120200-3310223310323002"></a>

## custom_storage_config.storage_device_list — storage_device_list / 010121012011 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- custom_storage_config.storage_device_list

<a id="canonical-0132210012211322-2122221202302203-1023130202202023-3113302130123110-1031220230011133-1121112233333132-3001022121131301-1233310020110321"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this fleet.

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
storage_device_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213020322300110-1000330212133232-0123301301012233-2333013211021020-3013103032221212-3333123311231232-1011123100100131-3032133220012012"></a>

## Direct properties — storage_device_list / 010121012011 / 3

- [storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323): complete subsection reference.

<a id="canonical-3200000022001202-3102322110001232-1210203020000322-3012011111031023-2111233233103031-1102121111020210-3221211113111033-3300310201123021"></a>

## Next pages — storage_device_list / 010121012011 / 4

- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323012330101213-1100110020011331-2010311311302212-1312121301012031-1022120000010131-3210013222122322-0100132302231220-3201103103123003"></a>

## custom_storage_config.storage_device_list.storage_devices — storage_devices / 132232100331 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- custom_storage_config.storage_device_list.storage_devices

<a id="canonical-2003122312300220-3130320313123323-0131023122012011-3320312313203022-2201302322301232-3233110333230133-0130001023013310-2012133310300310"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Devices. List of custom storage devices.

Upstream description:

List of custom storage devices.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_device"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "hpe_storage"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("netapp_trident",
    "pure_service_orchestrator")}
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

Terraform syntax:

```terraform
storage_devices {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221003021223313-1112231100223321-0200132110002312-2000001332220331-0023122022033102-2131001000203000-1222103300122220-0233211103021211"></a>

## Direct properties — storage_devices / 132232100331 / 3

<a id="canonical-1230313023100131-0121021111031320-2120023333220332-1022310012031030-1133213123131002-0201013230322030-2021132011233303-1313131002020113"></a>

<a id="canonical-3201011133010203-2203001313010011-3111131032021213-3303321012120232-2021030323120130-0323112331023100-1303010201313111-3101033003330331"></a>

## advanced_advanced_parameters property — storage_devices / 132232100331 / 4

Type: `["map", "string"]`. Optional.

Advanced Parameters. Map of parameter name and string value.

Upstream description:

Map of parameter name and string value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
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
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [custom_storage](resources--voltstack_site--reference--group-006.md#canonical-3322001302302223-0010031102103130-3333322002323232-3010223122302132-3221301320321103-1030210211023312-1031031021111003-3031012123032301): complete subsection reference.

- [hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-1000202310213110-2303000320310002-0303220122011131-0111301010332022-0211221202010230-1320100011122033-1120330012123212-2112010102133201): complete subsection reference.

- [netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123): complete subsection reference.

- [pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100): complete subsection reference.

<a id="canonical-3223323023313123-3230331322001022-0011032332331202-3313003301113121-0202121112211311-2011132221202201-1310130330313111-1113233132101222"></a>

<a id="canonical-1323020133130031-0133013020223223-2111232213011133-3033213301000001-3003011230132233-2112200201123111-3122130132001133-3101210330030320"></a>

## storage_device property — storage_devices / 132232100331 / 5

Type: `"string"`. Optional.

Storage Device. Storage device and device unit.

Upstream description:

Storage device and device unit.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2111202222320231-2113102110301010-1110200333010021-0302221103232203-1220323111010311-0113220013331001-0121000112021200-1301031312021112"></a>

## Next pages — storage_devices / 132232100331 / 6

- [custom_storage_config.storage_device_list.storage_devices.custom_storage](resources--voltstack_site--reference--group-006.md#canonical-3322001302302223-0010031102103130-3333322002323232-3010223122302132-3221301320321103-1030210211023312-1031031021111003-3031012123032301)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-1000202310213110-2303000320310002-0303220122011131-0111301010332022-0211221202010230-1320100011122033-1120330012123212-2112010102133201)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3322001302302223-0010031102103130-3333322002323232-3010223122302132-3221301320321103-1030210211023312-1031031021111003-3031012123032301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100201312311313-3020111112203023-3223331033021122-0132301213313132-1210020230313122-3213202202111023-2121332221323220-1300033013300031"></a>

## custom_storage_config.storage_device_list.storage_devices.custom_storage — custom_storage / 133002333213 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- custom_storage_config.storage_device_list.storage_devices.custom_storage

<a id="canonical-1300132012332302-1333320213021330-2011003130333033-0322213012333213-2021111023320121-3011031133321121-0233123103232320-1210100212000011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for custom storage.

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
custom_storage = {}
```

<a id="canonical-0320032021330030-3133010303303330-3113110212210220-1321123231020020-2122222222011220-3100112220003012-0332331223103312-1122003203101233"></a>

## Direct properties — custom_storage / 133002333213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211301030221010-1323202013002102-1110003001012030-2210203202231223-1300020213322130-3330013122333200-3321301202020120-1303323120112120"></a>

## Next pages — custom_storage / 133002333213 / 4

- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1000202310213110-2303000320310002-0303220122011131-0111301010332022-0211221202010230-1320100011122033-1120330012123212-2112010102133201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021333321033010-2220300331102313-1332201203110313-2011232233123020-1010000101012031-3232222132011123-1322332333130131-0113322022312233"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage — hpe_storage / 031102220233 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage

<a id="canonical-0301030011121032-1020010212302211-3311321021112001-1121032030110130-0121000223331012-3331132201311311-1320110230112301-0002010311013233"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for hpe storage.

Upstream description:

Device configuration for HPE Storage.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_server_port",
    "username")}
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
hpe_storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333202321211110-2020022033221013-2031201012103320-3300303303020122-3333100020113211-0220333222122123-1301232031223200-1100330100130113"></a>

## Direct properties — hpe_storage / 031102220233 / 3

<a id="canonical-3223332133033110-2312123323233023-3033220301021330-2201232310023213-0130031132232013-3103320220300231-2000323310132323-1102001013103023"></a>

<a id="canonical-3012201223202310-3311203122211011-1321132230202212-3032020232332221-0121213330020110-0320312030220311-1102233013130322-3001322201030300"></a>

## api_server_port property — hpe_storage / 031102220233 / 4

Type: `"number"`. Optional.

Storage server Port. Enter Storage Server Port.

Upstream description:

Enter Storage Server Port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [iscsi_chap_password](resources--voltstack_site--reference--group-006.md#canonical-3002023122213001-1201323231321233-1331331032312113-0112303002001310-0211221330312123-1033033230333000-3330013010023110-2320220313320220): complete subsection reference.

<a id="canonical-2213102123203130-1120210102332020-3232030311022120-3033133220132030-1100122003011020-3032322120211010-2031100202022231-1100301233322213"></a>

<a id="canonical-1312030223002012-0213203233212012-3131223000111310-2310001032101232-0011001122133323-1330123031102211-2022322322322002-2020030300301002"></a>

## iscsi_chap_user property — hpe_storage / 031102220233 / 5

Type: `"string"`. Optional.

Chap Username to connect to the HPE storage.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [password](resources--voltstack_site--reference--group-006.md#canonical-0032202121033001-2113321233323312-0312110113211323-0312120322221233-1113200322022310-3131223113133322-0013202213113101-2020203230200001): complete subsection reference.

<a id="canonical-0300113010300330-2213033232013310-0302201232121033-0332303200032231-1331032232120232-2313233202003221-0321121321320013-0133311012023022"></a>

<a id="canonical-3121233023010021-2331212331313210-1302022012102122-0221122300332022-2121221303012133-1113021032111011-3000102103030300-0100121313102320"></a>

## storage_server_ip_address property — hpe_storage / 031102220233 / 6

Type: `"string"`. Optional.

Storage Server IP address. Enter storage server IP address.

Upstream description:

Enter storage server IP address.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2223101033323200-0210121330332002-1033001322311121-0210330032003311-3320233030023033-1201013233301321-0311311133131311-2033231130320202"></a>

<a id="canonical-2303103312321232-0032133213323103-1120120131210330-3000302201322223-2321202002202112-0310013313301223-0203303133123210-2313113231033121"></a>

## storage_server_name property — hpe_storage / 031102220233 / 7

Type: `"string"`. Optional.

Storage Server Name. Enter storage server Name.

Upstream description:

Enter storage server Name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-1010011230111031-3211233210133101-1122102213003212-2132123001232110-3200102311122111-0201320103021100-1113221231213011-2100210011011233"></a>

<a id="canonical-2232220312333332-2311121230132331-2131113032211311-1220013101123131-2101331200000120-1122220012021032-0033100033312102-0200132330310133"></a>

## username property — hpe_storage / 031102220233 / 8

Type: `"string"`. Optional.

Username to connect to the HPE storage management IP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0301302223032033-3233232020102131-0121100103123323-0021222312201021-1111003132203002-0123312112332232-3303332200122113-3321313002313220"></a>

## Next pages — hpe_storage / 031102220233 / 9

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--voltstack_site--reference--group-006.md#canonical-3002023122213001-1201323231321233-1331331032312113-0112303002001310-0211221330312123-1033033230333000-3330013010023110-2320220313320220)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](resources--voltstack_site--reference--group-006.md#canonical-0032202121033001-2113321233323312-0312110113211323-0312120322221233-1113200322022310-3131223113133322-0013202213113101-2020203230200001)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3002023122213001-1201323231321233-1331331032312113-0112303002001310-0211221330312123-1033033230333000-3330013010023110-2320220313320220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022112202000031-1232021103231311-2200100210103301-0211201022320233-0020022233130111-3220321130000231-0230110223120121-3113220213321103"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password — iscsi_chap_password / 212332121122 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-1000202310213110-2303000320310002-0303220122011131-0111301010332022-0211221202010230-1320100011122033-1120330012123212-2112010102133201)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password

<a id="canonical-1000331330303013-2021230220333101-0111223203022121-1332302110102020-0212201312013132-1110320112122310-1132031011101032-1331102123130020"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
iscsi_chap_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130122021313120-3133230232232220-2013202332021310-1201011202130000-1112100120321103-0321111211212101-3311303002021021-0032001002133213"></a>

## Direct properties — iscsi_chap_password / 212332121122 / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-2013001021231130-1212201230111202-1231321322030011-2131220033320333-2132110101203021-0100113032303313-2230233130322101-0022210123212001): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-0333110310330231-2320210110012310-3000311010301311-1100030320330222-2111100110103122-1013223320031232-1203100211103112-3033203211122022): complete subsection reference.

<a id="canonical-2203212011232123-3020003033133031-0232032202123110-1133210320120112-2122031000313303-3212213022211100-3031220330220311-1100130232032013"></a>

## Next pages — iscsi_chap_password / 212332121122 / 4

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-2013001021231130-1212201230111202-1231321322030011-2131220033320333-2132110101203021-0100113032303313-2230233130322101-0022210123212001)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-0333110310330231-2320210110012310-3000311010301311-1100030320330222-2111100110103122-1013223320031232-1203100211103112-3033203211122022)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-1000202310213110-2303000320310002-0303220122011131-0111301010332022-0211221202010230-1320100011122033-1120330012123212-2112010102133201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2013001021231130-1212201230111202-1231321322030011-2131220033320333-2132110101203021-0100113032303313-2230233130322101-0022210123212001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122222031021300-3322013023121023-2101133332313110-1003320133020023-3120001102220220-1200331223103013-0200311230001313-3333110332111021"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info — blindfold_secret_info / 001300330201 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-1000202310213110-2303000320310002-0303220122011131-0111301010332022-0211221202010230-1320100011122033-1120330012123212-2112010102133201)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--voltstack_site--reference--group-006.md#canonical-3002023122213001-1201323231321233-1331331032312113-0112303002001310-0211221330312123-1033033230333000-3330013010023110-2320220313320220)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info

<a id="canonical-3012113222130313-1031010211233213-2333310230122320-3020002201301310-1103023322221303-2202220313010230-2132011310030120-2210212112133201"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2201102212221301-3132131222020133-1323230032021132-2333322312020303-3331223303101101-1010213213003220-0132002232021321-3023020023001321"></a>

## Direct properties — blindfold_secret_info / 001300330201 / 3

<a id="canonical-3213113300003211-1300000122331233-0110331013102211-2010033021312100-0021212002002302-0222002021323101-1120111013113031-1100212321031203"></a>

<a id="canonical-2321202022110333-2011100010303333-3130210110330310-1013302133212003-1020301010202001-2311323202030121-2212310332203231-3203222121002321"></a>

## decryption_provider property — blindfold_secret_info / 001300330201 / 4

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

<a id="canonical-1131322032112313-1001330030003302-1331330300003202-3212232030002212-3311300322122220-3203000221022022-3233300123003231-2122121312001021"></a>

<a id="canonical-1312322010311301-3323330221023002-1101030333221331-0301333120333213-1111320102131220-2332313010033333-1233012313000302-2120303323120332"></a>

## location property — blindfold_secret_info / 001300330201 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2211202121333331-1202111212112222-3112033233012111-1101231131033302-2220310230103303-2302120202030120-3110312221210201-3321133233210000"></a>

<a id="canonical-1030203210020203-3002120021100223-0300222232303211-2322313321013303-0312032133230322-0022031310332022-3012300203030020-1123320210302113"></a>

## store_provider property — blindfold_secret_info / 001300330201 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-2100332301231110-3010200120201023-1110333010021211-0101023003103021-2332131303012100-1221203111200012-3101003001011212-3200111331101220"></a>

## Next pages — blindfold_secret_info / 001300330201 / 7

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--voltstack_site--reference--group-006.md#canonical-3002023122213001-1201323231321233-1331331032312113-0112303002001310-0211221330312123-1033033230333000-3330013010023110-2320220313320220)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0333110310330231-2320210110012310-3000311010301311-1100030320330222-2111100110103122-1013223320031232-1203100211103112-3033203211122022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220100320311300-0223322110131022-2103121131311020-1120000201110312-1312032102120010-3122113131012033-1000133320302123-2031211001132220"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info — clear_secret_info / 303300332112 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-1000202310213110-2303000320310002-0303220122011131-0111301010332022-0211221202010230-1320100011122033-1120330012123212-2112010102133201)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--voltstack_site--reference--group-006.md#canonical-3002023122213001-1201323231321233-1331331032312113-0112303002001310-0211221330312123-1033033230333000-3330013010023110-2320220313320220)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info

<a id="canonical-3202213221202111-0000230231333321-2022332010330322-1122121020120211-0111110013230000-2203332120203010-2033000023111310-2002033211331223"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020012330233213-0210322120220202-3303212332133203-1303021132121120-2320323332001231-2102230202110202-2122003222013030-1212132103111132"></a>

## Direct properties — clear_secret_info / 303300332112 / 3

<a id="canonical-3010122320300020-2211220323210113-3203123232323123-1102200130122100-2002103130331331-2212211300223332-0003313222022103-1220330223012013"></a>

<a id="canonical-3122102213033331-1033213101002223-0133223000320122-0001033311032321-0211131302231130-3210021122102221-0332321031333232-1110111330021220"></a>

## provider_ref property — clear_secret_info / 303300332112 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1323121023000010-1113321030112303-2002221222230030-1301122220030233-1303220002020330-3112002132123122-3331130333031310-1203300002200301"></a>

<a id="canonical-3210110001201102-0332112301302121-1132323322333101-2132001003332213-2121213300010212-1133101203231001-2333113312030010-2231030313001330"></a>

## URL property — clear_secret_info / 303300332112 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0220210300302302-0220032132310301-2303220030112210-2332320132103002-0002301231020123-1101312313021320-0100132001031023-0100323312200202"></a>

## Next pages — clear_secret_info / 303300332112 / 6

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--voltstack_site--reference--group-006.md#canonical-3002023122213001-1201323231321233-1331331032312113-0112303002001310-0211221330312123-1033033230333000-3330013010023110-2320220313320220)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0032202121033001-2113321233323312-0312110113211323-0312120322221233-1113200322022310-3131223113133322-0013202213113101-2020203230200001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111220220312323-2001322303213013-2303000321002131-3020033320332322-3103302200301313-2210121112021321-3012211112220310-0212022303013211"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.password — password / 332001131212 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-1000202310213110-2303000320310002-0303220122011131-0111301010332022-0211221202010230-1320100011122033-1120330012123212-2112010102133201)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.password

<a id="canonical-1211323323030013-1102220303312123-1131311002231233-0031202310212221-1020220213103223-0013110202003123-3121211331212123-3111100023220003"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233333231111211-2210020013131122-3200311020033003-1033131223002221-1221131133213313-1323230230230200-0030111312000220-3020201323200322"></a>

## Direct properties — password / 332001131212 / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-3210030020211213-0200112332032203-1102220101103212-2103231001330200-3033320033002032-0302122333013330-2220022213130212-2302022022113130): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-0301021303003300-0200133310012000-1000110131031202-3102112202110003-3011310233232300-2221110102113110-2210020311321301-0122213101311233): complete subsection reference.

<a id="canonical-1121222203133002-3101132313221112-0101320322332312-0013120311122012-0020223000202332-3031112212332020-2232300311311132-0311323302323100"></a>

## Next pages — password / 332001131212 / 4

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-3210030020211213-0200112332032203-1102220101103212-2103231001330200-3033320033002032-0302122333013330-2220022213130212-2302022022113130)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-0301021303003300-0200133310012000-1000110131031202-3102112202110003-3011310233232300-2221110102113110-2210020311321301-0122213101311233)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-1000202310213110-2303000320310002-0303220122011131-0111301010332022-0211221202010230-1320100011122033-1120330012123212-2112010102133201)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3210030020211213-0200112332032203-1102220101103212-2103231001330200-3033320033002032-0302122333013330-2220022213130212-2302022022113130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112002000222033-3333332013303020-1222200120133231-2322300020100213-3100033130131333-2110210011321021-0100123322233032-1133201230133022"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info — blindfold_secret_info / 011231230103 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-1000202310213110-2303000320310002-0303220122011131-0111301010332022-0211221202010230-1320100011122033-1120330012123212-2112010102133201)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](resources--voltstack_site--reference--group-006.md#canonical-0032202121033001-2113321233323312-0312110113211323-0312120322221233-1113200322022310-3131223113133322-0013202213113101-2020203230200001)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info

<a id="canonical-1103231300110012-3311303023103130-2211300210202013-0132320203022313-2310201320000210-2110002113033233-2003221133322123-2121111202000121"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132323100101233-3323111233332200-3322230323320110-0330302211130032-1320211303000003-0210303110001103-2101133232112223-1113222232332130"></a>

## Direct properties — blindfold_secret_info / 011231230103 / 3

<a id="canonical-2322202220011311-1201021130232130-0302031302102112-0122131200202220-3113313332031313-2223121010320112-1022233020332220-2223223232310132"></a>

<a id="canonical-1221211320202230-2132202022221111-3220012022230021-2313332023311121-0101222232000212-0330133212112030-3333331111320313-3322023122301002"></a>

## decryption_provider property — blindfold_secret_info / 011231230103 / 4

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

<a id="canonical-3011302020003121-2230110123120121-3030230313000211-3322110130303110-0303131321103001-3033033022220121-2320322110213201-2320101000303333"></a>

<a id="canonical-1032201122330221-0121020332121113-2001121331233312-0200331021200021-0133311133132300-2032200101000212-3111232031031303-3202211333100313"></a>

## location property — blindfold_secret_info / 011231230103 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2202121103023210-1003123232122003-1101302303120110-2233301333230002-0223201011323023-2030123203011210-2121022300013000-3021012121311233"></a>

<a id="canonical-1011121003120133-0031202000120010-1021211003010222-0203332221300210-0203330111222010-1220021202123232-3002102210230003-1311300133330202"></a>

## store_provider property — blindfold_secret_info / 011231230103 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-2320333002202211-2132101303333330-3111003231000210-3221213122012232-1102000200011021-1120131001233322-2112033122031201-0022023233112013"></a>

## Next pages — blindfold_secret_info / 011231230103 / 7

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](resources--voltstack_site--reference--group-006.md#canonical-0032202121033001-2113321233323312-0312110113211323-0312120322221233-1113200322022310-3131223113133322-0013202213113101-2020203230200001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0301021303003300-0200133310012000-1000110131031202-3102112202110003-3011310233232300-2221110102113110-2210020311321301-0122213101311233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330231132222301-3102322313210302-3111030132030233-0311132303200113-1030100032123033-1300101112311313-2012113113022031-2312323202030130"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.clear_secret_info — clear_secret_info / 120113100321 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-1000202310213110-2303000320310002-0303220122011131-0111301010332022-0211221202010230-1320100011122033-1120330012123212-2112010102133201)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](resources--voltstack_site--reference--group-006.md#canonical-0032202121033001-2113321233323312-0312110113211323-0312120322221233-1113200322022310-3131223113133322-0013202213113101-2020203230200001)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.clear_secret_info

<a id="canonical-3101020032022321-3303220300220321-1231020112101031-0101223221112333-0110213121331103-0002100201132230-1002301102133111-0310100032233231"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0300132300113011-3330033020322133-0302231013223310-1030323033231302-2112320102030220-3212303333233101-1111210322332002-1300233212330032"></a>

## Direct properties — clear_secret_info / 120113100321 / 3

<a id="canonical-1301101131112120-1311020331121002-3213121121231322-1023131323013320-2032200232132332-2112210223000303-1300323203131330-0332222130201223"></a>

<a id="canonical-1023122011231212-3330133332003312-2222231333222213-3302222013330013-1000110113012201-3110000200210311-0233020123323201-1221130121223232"></a>

## provider_ref property — clear_secret_info / 120113100321 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2312210331131232-3113332220131033-0111002303213201-1012320332121330-2323320100301121-1212303103332320-3130130300220131-2003121203330202"></a>

<a id="canonical-1303312231212123-3203223221013201-3321200123020031-0032320102220211-0002232030331120-2222131201122132-2123331213212112-1111000213300010"></a>

## URL property — clear_secret_info / 120113100321 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0013012123110133-0213110223033331-1032011122101010-2033311002310122-1000321022132013-3220310213123030-2313322330221021-2232021001023310"></a>

## Next pages — clear_secret_info / 120113100321 / 6

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](resources--voltstack_site--reference--group-006.md#canonical-0032202121033001-2113321233323312-0312110113211323-0312120322221233-1113200322022310-3131223113133322-0013202213113101-2020203230200001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013120331101010-1311030333220012-2021121311020131-2100021120302123-0013021232033312-2012003210331112-0133023332130323-0312323103302122"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident — netapp_trident / 013212231021 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident

<a id="canonical-0201120230201133-0202220222212232-1102320232223212-1213332100213033-1200113121313311-3330202032220011-0001013032131332-0330322321302333"></a>

Type: `"object"`. single nested block, Optional.

Device configuration for NetApp Trident Storage.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("netapp_backend_ontap_nas",
    "netapp_backend_ontap_san")}
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
  "x-ves-oneof-field-backend_choice": "[\"netapp_backend_ontap_nas\",\"netapp_backend_ontap_san\"]"
}
```

Terraform syntax:

```terraform
netapp_trident {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311032021130112-0210331230223210-0212322321330222-0012203033201312-1203203130110023-2232331213323222-1113020302233102-3020110332012313"></a>

## Direct properties — netapp_trident / 013212231021 / 3

- [netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213): complete subsection reference.

- [netapp_backend_ontap_san](resources--voltstack_site--reference--group-007.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010): complete subsection reference.

<a id="canonical-1113033003033011-1213212300030002-1033121203102103-3210332222231032-1132210100330133-0100312223012001-1022123031033223-3231031232203100"></a>

## Next pages — netapp_trident / 013212231021 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-007.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022333322332333-3120101202033311-2232220113032310-0312313230130023-1321032302323110-0232033103210023-1312230022131023-0111301233210322"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas — netapp_backend_ontap_nas / 211312330321 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas

<a id="canonical-3001223200020321-3032212111300001-2110300323212113-1312103233212300-2323133233213231-0332100303222220-0233112313313200-1223032132230012"></a>

Type: `"object"`. single nested block, Optional.

Configuration of storage backend for NetApp ONTAP NAS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_driver_name",
    "username"),
  validators.ConflictingObjectAttributes("data_lif_dns_name",
    "data_lif_ip"),
  validators.ConflictingObjectAttributes("management_lif_dns_name",
    "management_lif_ip")}
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
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

Terraform syntax:

```terraform
netapp_backend_ontap_nas {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103323223100330-0212010321002213-0121132012032002-2001011122213012-0002112113203233-2022030313012221-0230103323310030-1300113133231101"></a>

## Direct properties — netapp_backend_ontap_nas / 211312330321 / 3

- [auto_export_cidrs](resources--voltstack_site--reference--group-006.md#canonical-0031002113320100-3311111012000012-2321203302213320-2331001211100030-2331301112021303-0310202213332130-3012311010322211-1332221233102301): complete subsection reference.

<a id="canonical-1102002310033330-0210320001113312-1202012230211101-1033200300023101-1212233221311303-3102010013332220-3132311311330002-2210233220121120"></a>

<a id="canonical-1330230220310313-1302303332330332-0100122203120023-0012030332101013-3000101013001333-1210030200122020-1301000022110011-2130203103101022"></a>

## auto_export_policy property — netapp_backend_ontap_nas / 211312330321 / 4

Type: `"bool"`. Optional.

Policy configuration for this feature.

Upstream description:

Enable automatic export policy creation and updating.

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

<a id="canonical-3111101213233202-0331232021111110-1323122213213301-0111100100032121-0133030112030121-2221302321312213-0113132203021113-1032223030313201"></a>

<a id="canonical-2102213200022322-3220333120203131-0010120001231013-1331300333030131-1312313211333012-0012101010030123-3031131111303313-2002222230213012"></a>

## backend_name property — netapp_backend_ontap_nas / 211312330321 / 5

Type: `"string"`. Optional.

Configuration of Backend Name. Driver is name + '\_' + dataLIF.

Upstream description:

Configuration of Backend Name. Driver is name + "\_" + dataLIF.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 50,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1032110113112331-2202002213110000-1121202200112132-1010230310032201-2021011110323011-1032003202222123-3111232100002231-3202330133010033"></a>

<a id="canonical-0020301032023330-3101221021132101-2013312200210210-0022032011221200-3302312032321032-3120011233333000-0231222021333002-3323102101320222"></a>

## client_certificate property — netapp_backend_ontap_nas / 211312330321 / 6

Type: `"string"`. Optional.

Please Enter base64-encoded value of client certificate. Used for certificate-based auth.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](resources--voltstack_site--reference--group-006.md#canonical-3323310221101132-3311012010300323-0011210100201033-3313230212023313-3331001103110002-1130233312101201-1313330201032132-3020112210110130): complete subsection reference.

<a id="canonical-0101031212312013-0310220003123122-0000313332223321-0230231031120332-3120022323101321-3231003113130031-3003102023322300-2302211101132333"></a>

<a id="canonical-2022303222002331-3100131113303011-3112132331021130-3112122221113110-1012232233020211-3111111321300120-1223021003013010-0333031013332203"></a>

## data_lif_dns_name property — netapp_backend_ontap_nas / 211312330321 / 7

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3222220322220030-0201103322302012-3121232001302011-1013323031111111-2333123112233101-1113023033222313-0122233320131001-1203101211301323"></a>

<a id="canonical-3103113133133000-2012123211032201-1110312103320030-2130021221111133-1222201213231301-2331221323203031-0202310002300002-3012030333122120"></a>

## data_lif_ip property — netapp_backend_ontap_nas / 211312330321 / 8

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

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

<a id="canonical-3230201001203022-0020332001332230-2222333132132012-0221301203100212-2010313311101132-1031221023123231-1003331331212120-2231232333330213"></a>

<a id="canonical-1010023102012000-1232100311332032-3313321131120331-3223230011312321-0022112011120011-2222113332301311-3230222022110301-2031110002101130"></a>

## labels property — netapp_backend_ontap_nas / 211312330321 / 9

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":20},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"20\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
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
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-3132320030010211-1332033031333230-2220011223102131-2203000021102313-3102100203032303-0121333110300022-3233120023202302-1021113330112200"></a>

<a id="canonical-2000130312111203-2233132301031032-0031213101210121-2001330230212102-2332310030023300-2122222321130223-3031020020301223-3333012113030032"></a>

## limit_aggregate_usage property — netapp_backend_ontap_nas / 211312330321 / 10

Type: `"string"`. Optional.

Fail provisioning if usage is above this percentage. Not enforced by default.

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

<a id="canonical-0323103232323103-0132122111210010-2223132231032221-3012220233101210-2322333122122233-3120203322331300-0011013232013002-2211031323333300"></a>

<a id="canonical-0023113310110011-0030103223312322-0132011113100112-1312002012222002-2100313331312313-2103222332101102-3031213013111310-3011221130120132"></a>

## limit_volume_size property — netapp_backend_ontap_nas / 211312330321 / 11

Type: `"string"`. Optional.

Fail provisioning if requested volume size is above this value. Not enforced by default.

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

<a id="canonical-0100311133021302-0001321121200302-1203210133031233-0323133210113333-2123231123320102-2333032321001233-0232111322313333-2221322000111033"></a>

<a id="canonical-0321213013002203-3101013013333200-1231021210231032-3330332312333221-1022021112200020-0103220320233013-3333203102202322-2103331000233223"></a>

## management_lif_dns_name property — netapp_backend_ontap_nas / 211312330321 / 12

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1123331100003200-2331200110001222-2320323330310132-3333223022020332-1203333030222232-3100002113123132-0130110103113303-1320122102231311"></a>

<a id="canonical-2323301132213022-1010202110110120-2000200333330102-0022200323100332-1200110231321220-0310223133132110-3332310223030111-0332031322021113"></a>

## management_lif_ip property — netapp_backend_ontap_nas / 211312330321 / 13

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

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

<a id="canonical-2232332110323123-3330200221331202-1020003210101220-1002102020101021-0313223220212001-2332230133133023-3033020223130223-2012001331113010"></a>

<a id="canonical-1223333001130111-1021231021100202-1013132013113003-0132020012322310-3000022123302133-1011202331002230-3021331320222120-3231011310223202"></a>

## nfs_mount_options property — netapp_backend_ontap_nas / 211312330321 / 14

Type: `"string"`. Optional.

Comma-separated list of NFS mount OPTIONS. Not enforced by default.

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

- [password](resources--voltstack_site--reference--group-006.md#canonical-0113330003101111-2120321331121121-2132013220212333-0013110203013132-2103112030213122-3032333122221222-1131332203001223-3010221331303233): complete subsection reference.

<a id="canonical-1020000223321011-1030201333012201-1312223331313013-3020332200010320-3331020013311310-1020303320330202-3330131003322220-3233300021310032"></a>

<a id="canonical-2022201302203120-3111111101002013-2023210212222330-1222012123301023-1310023320100101-3301312131320003-0022013230330110-2321311103011320"></a>

## region property — netapp_backend_ontap_nas / 211312330321 / 15

Type: `"string"`. Optional.

Backend Region. Virtual Pool Region.

Upstream description:

Virtual Pool Region.

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

- [storage](resources--voltstack_site--reference--group-006.md#canonical-0003212002302320-0230230222133000-1230132033022210-1303013121300202-1221031233022031-2131010312122002-0032102203210130-1123120030221322): complete subsection reference.

<a id="canonical-2311323100313210-2101133131110323-1132023031203100-0011220301220301-2112203013113302-0331320121023301-0312122202013133-3032221231221003"></a>

<a id="canonical-0122232311222330-2020103023110222-3032000122031222-2303132313003221-0213212313020321-2200030321023320-0201003012200132-3223233002201111"></a>

## storage_driver_name property — netapp_backend_ontap_nas / 211312330321 / 16

Type: `"string"`. Optional.

\[Enum: ontap-nas|ontap-nas-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-nas\`, \`ontap-nas-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"
  ],
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-1332310133032011-3031232321332221-3133110013033000-3203311021203021-0301023313310313-2101230203322003-0222000201221100-1021020103131033"></a>

<a id="canonical-1033221112212111-3303023130303132-2100231230021332-3022102122321213-3332221112003232-2333010011103201-3332133123331000-1212301132022022"></a>

## storage_prefix property — netapp_backend_ontap_nas / 211312330321 / 17

Type: `"string"`. Optional.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

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

<a id="canonical-3322033110301230-3000010330101030-0103333131303010-0011011001120303-3210112320310103-1323032220312011-2220131333123001-2212233220213023"></a>

<a id="canonical-3312113000130210-1221332123232012-1112111023322010-2022031033333023-3231103101001323-1311330302312122-0221102030120101-0021212111132322"></a>

## svm property — netapp_backend_ontap_nas / 211312330321 / 18

Type: `"string"`. Optional.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0220222300333323-0021032030222030-0022210120120211-0210120222112003-3332020010002302-1201320032032122-3101121103032313-3122111022020303"></a>

<a id="canonical-0110311213313230-1210230330300110-1220112111133301-1112020303003032-1010213012012220-3212223010130000-1313321330331210-0001012122113120"></a>

## trusted_ca_certificate property — netapp_backend_ontap_nas / 211312330321 / 19

Type: `"string"`. Optional.

Please Enter base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

<a id="canonical-0001302013312301-3331301302332231-2132112113211203-0023220202131210-1122002311133203-1011333023111112-2112313032310311-3331000230333000"></a>

<a id="canonical-2300321232202211-2222203000031021-1033001021101102-0011121021300031-0131310132300132-0333132212123011-3132113321120032-1210020213231302"></a>

## username property — netapp_backend_ontap_nas / 211312330321 / 20

Type: `"string"`. Optional.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--voltstack_site--reference--group-006.md#canonical-2113210123032330-2220213133030132-0302222012203211-0333131032213120-0232121001321012-1121201012031013-3033112013220010-2332201222201032): complete subsection reference.

<a id="canonical-0133003312213001-1012103102122211-1131122221223203-1023003331122033-3122330000002213-3122312210123330-1011123302232010-0001033001121123"></a>

## Next pages — netapp_backend_ontap_nas / 211312330321 / 21

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](resources--voltstack_site--reference--group-006.md#canonical-0031002113320100-3311111012000012-2321203302213320-2331001211100030-2331301112021303-0310202213332130-3012311010322211-1332221233102301)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-3323310221101132-3311012010300323-0011210100201033-3313230212023313-3331001103110002-1130233312101201-1313330201032132-3020112210110130)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--voltstack_site--reference--group-006.md#canonical-0113330003101111-2120321331121121-2132013220212333-0013110203013132-2103112030213122-3032333122221222-1131332203001223-3010221331303233)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--voltstack_site--reference--group-006.md#canonical-0003212002302320-0230230222133000-1230132033022210-1303013121300202-1221031233022031-2131010312122002-0032102203210130-1123120030221322)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](resources--voltstack_site--reference--group-006.md#canonical-2113210123032330-2220213133030132-0302222012203211-0333131032213120-0232121001321012-1121201012031013-3033112013220010-2332201222201032)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0031002113320100-3311111012000012-2321203302213320-2331001211100030-2331301112021303-0310202213332130-3012311010322211-1332221233102301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223133200223132-3333311011030111-2020302201300132-2201222001030312-0120120010211123-0232231222032031-3130213310133212-1030322131332330"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs — auto_export_cidrs / 112221231322 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs

<a id="canonical-1033302221312333-3130020302221310-3201211002322222-0130332213303201-2322223111310030-1003321231332332-3131323130310203-0100011231222310"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
auto_export_cidrs {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302131323121332-1013030320212223-3000012133301123-1122132223301310-3021313023230303-0000132221200000-1223120323030123-3213330213303311"></a>

## Direct properties — auto_export_cidrs / 112221231322 / 3

<a id="canonical-0031233331213212-3312200313020030-2223200200202212-3322020001020303-3012010002331012-1123330001212130-2202210232003210-1200232023110222"></a>

<a id="canonical-0322002213021011-1033220232331130-2103332033011122-2020313323303002-1021303021220130-3130230033331100-3102200120302011-3232122232133032"></a>

## prefixes property — auto_export_cidrs / 112221231322 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2230303111330101-0330213113302033-3121222233310311-0202131303320101-3222131020120013-0110212331200220-1122332001202321-2313103333012032"></a>

## Next pages — auto_export_cidrs / 112221231322 / 5

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3323310221101132-3311012010300323-0011210100201033-3313230212023313-3331001103110002-1130233312101201-1313330201032132-3020112210110130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123022122302221-0032321132111331-2103110131200120-0101131331110233-3013132110001122-2113313132021113-0211013011020333-3112210102312112"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key — client_private_key / 000110301001 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key

<a id="canonical-2120231213112001-0220121031321133-3122130120022203-3323121200100122-1210323121031130-2210220001001230-0010330132220030-2013021020131330"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
client_private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121002230302333-2233030031220011-0300233020123122-0303120212231031-1001020322212200-3333012323022031-2202213200011121-2213210222032313"></a>

## Direct properties — client_private_key / 000110301001 / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-0002121003322223-2010221013002112-3030213031221123-1320131133220322-3131321213322300-0231220202311110-3120210102033323-2332113100202012): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-1013033200330220-2011321030102311-3303302203200032-3203303201100331-2221212331213221-3011031010201123-2203220322023133-1201033221302000): complete subsection reference.

<a id="canonical-2130330030201123-3133303311310032-3222121133203230-2001220233031100-0230033010002003-1332111133323003-0132023322212131-2103322312203012"></a>

## Next pages — client_private_key / 000110301001 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-0002121003322223-2010221013002112-3030213031221123-1320131133220322-3131321213322300-0231220202311110-3120210102033323-2332113100202012)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-1013033200330220-2011321030102311-3303302203200032-3203303201100331-2221212331213221-3011031010201123-2203220322023133-1201033221302000)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0002121003322223-2010221013002112-3030213031221123-1320131133220322-3131321213322300-0231220202311110-3120210102033323-2332113100202012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101333313013313-1230330030023323-3002323001032322-0301002002000332-2032033110330131-1322302310330331-0310200321213121-0310121231221001"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info — blindfold_secret_info / 032032330112 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-3323310221101132-3311012010300323-0011210100201033-3313230212023313-3331001103110002-1130233312101201-1313330201032132-3020112210110130)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info

<a id="canonical-0320120003333133-2003121230123231-1201300223122131-1210313023231233-0112213002110100-0101221011313121-1233112010031033-0322012132103220"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102012123102103-0302123123220033-0203202003001302-3113103030311322-1013211132201031-2033332001133000-3111210230130310-0130212031002002"></a>

## Direct properties — blindfold_secret_info / 032032330112 / 3

<a id="canonical-3330123033210023-2132002003103222-2333110033323131-3211210210312121-3023010111012232-3223203000021023-0203030232031012-1320011133130210"></a>

<a id="canonical-1303213333323132-1020031230002201-2111322112032123-0311311333113213-0102231021333232-2101023321210310-2323033312122221-2022331022003333"></a>

## decryption_provider property — blindfold_secret_info / 032032330112 / 4

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

<a id="canonical-0033000112330132-3313213010332102-2221200002331131-2012003120200123-2322110132231222-1212330112313311-2303333230133212-3020021231221202"></a>

<a id="canonical-2102233120030133-3001300103002301-0111322332013020-2201122322131223-0222311012000223-2003231133130222-1010101020303030-2311010320121202"></a>

## location property — blindfold_secret_info / 032032330112 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1230133002021000-1321302033111320-1131202030103200-3211003102022131-3200313312223320-2011332100001222-2020031003203032-2113010210103212"></a>

<a id="canonical-3020023030313303-1033211033312012-2221103113302233-1231301322211330-2321001223303221-3233313000101010-3003012300232301-0013100321300021"></a>

## store_provider property — blindfold_secret_info / 032032330112 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-1332230112321302-0000113321110310-2310133333102310-3133310030033032-0321311110021003-0012203002131303-0011001322210021-1131111300011203"></a>

## Next pages — blindfold_secret_info / 032032330112 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-3323310221101132-3311012010300323-0011210100201033-3313230212023313-3331001103110002-1130233312101201-1313330201032132-3020112210110130)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1013033200330220-2011321030102311-3303302203200032-3203303201100331-2221212331213221-3011031010201123-2203220322023133-1201033221302000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133230302312312-0111110012202300-2200000203133332-2223332202010333-2330230003010001-0203011333111303-0203320323113121-0313233220313232"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info — clear_secret_info / 003120301203 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-3323310221101132-3311012010300323-0011210100201033-3313230212023313-3331001103110002-1130233312101201-1313330201032132-3020112210110130)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info

<a id="canonical-3030313030221132-1213013001030301-2313031012131232-2211123323120320-3112112311123003-2101132312320030-1202102130302321-1112002210300120"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200121220210322-2202100122102201-3121221233332331-0033301111233123-3130210123122212-0213000121121131-0122101331012333-3313012231101230"></a>

## Direct properties — clear_secret_info / 003120301203 / 3

<a id="canonical-2123300101220200-1233121013213312-0330201201300100-2312303033210321-1102023103320121-3300123200322103-0221010311331112-2220121013103310"></a>

<a id="canonical-3001321033021232-2211231210120220-3030213203000221-0133023201012320-0133220030010013-3311013032213101-1230131031012203-0122221100012332"></a>

## provider_ref property — clear_secret_info / 003120301203 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0030211213102220-2133221000223301-1122323323003200-2002323210033123-1311223112022022-2302323212212213-3021212310301323-1233130211103000"></a>

<a id="canonical-3033310232212331-3323010330002033-2033112210301300-1132231023233223-1123000130322113-1202333001120012-3333222232112210-1111303332001011"></a>

## URL property — clear_secret_info / 003120301203 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2021321022133222-1212030012212210-1000211112121223-3331131002123302-0121112300303030-3030130300103300-3020032211212301-1023120002221201"></a>

## Next pages — clear_secret_info / 003120301203 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-3323310221101132-3311012010300323-0011210100201033-3313230212023313-3331001103110002-1130233312101201-1313330201032132-3020112210110130)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0113330003101111-2120321331121121-2132013220212333-0013110203013132-2103112030213122-3032333122221222-1131332203001223-3010221331303233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110222203101132-1123102230313022-1113320330311203-2330220121221130-0021102312100122-1232302313232021-1130231002121223-2201110123111222"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password — password / 231203103010 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password

<a id="canonical-2122131322323222-1221323010331210-3320302101210222-3100131011122301-3112323122222331-0221121202233030-3222123312122211-2022012320021001"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103111111210221-0321023021020301-2023330321023320-2021103321330113-0023032313010321-0330130222111220-0131033012111303-2011333021221312"></a>

## Direct properties — password / 231203103010 / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-3103232323013000-0333102303300130-1002331022123132-2230303023113000-3130002003121311-1130111330031112-3122030032033122-3211212011200001): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-1213231011320213-1032230221203231-3000211123313320-1110013101102201-0330302212000300-3002220110030123-2203210200121030-3021302212322123): complete subsection reference.

<a id="canonical-1202303003211322-0130201212123112-3330002130313110-2022332220010130-3331212100322032-3023201001102223-0312112131131023-0232013010233010"></a>

## Next pages — password / 231203103010 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-3103232323013000-0333102303300130-1002331022123132-2230303023113000-3130002003121311-1130111330031112-3122030032033122-3211212011200001)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-1213231011320213-1032230221203231-3000211123313320-1110013101102201-0330302212000300-3002220110030123-2203210200121030-3021302212322123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3103232323013000-0333102303300130-1002331022123132-2230303023113000-3130002003121311-1130111330031112-3122030032033122-3211212011200001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302111012023011-3210103102320022-0331100330032101-1013311003313110-3000233103230322-1122332221233023-1111201210003213-2313103130230230"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info — blindfold_secret_info / 310101102000 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--voltstack_site--reference--group-006.md#canonical-0113330003101111-2120321331121121-2132013220212333-0013110203013132-2103112030213122-3032333122221222-1131332203001223-3010221331303233)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info

<a id="canonical-0222011231023101-1200023310320030-3002310222131031-0300102313313220-2232230221331000-1032231111101102-2333300001221112-3320211233113201"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012131010310032-1000131311320100-0121123023222102-3200322122112221-2313221223322032-2103002022320203-3022002010233021-0313032032223100"></a>

## Direct properties — blindfold_secret_info / 310101102000 / 3

<a id="canonical-3233032032122000-0130201311020301-0302111033123021-3121013030320032-2133132133121121-2002212310013031-2102310211330313-0031112023103310"></a>

<a id="canonical-1021121200031333-0333320233210000-3103112111133030-2012112331221221-3322211110023111-0200213031210023-2000012132331101-2121333003300313"></a>

## decryption_provider property — blindfold_secret_info / 310101102000 / 4

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

<a id="canonical-3332321202312023-3313231013132303-3313320312021232-3020223003200001-3313300233200130-1303302122122103-1213131211013030-0121101232203202"></a>

<a id="canonical-3300211223323131-0013231031303113-3203231000301032-2011131331123232-0121203200230030-3023112010021002-2032102320020320-2100120110303100"></a>

## location property — blindfold_secret_info / 310101102000 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1120213123320113-3200223033201013-1300331103002232-3222232320121121-0112110333230221-1103330101002301-0233211011032013-2100011303111110"></a>

<a id="canonical-0012311212100002-0012302100203313-2201130301312102-1013303321112220-0032302023231201-0311130331113113-0022300101111012-1033123320012310"></a>

## store_provider property — blindfold_secret_info / 310101102000 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-0121221220301231-1032311003300033-3133333102023102-1220232201211321-0210033311121113-0010213212220123-1313023033323000-3101010101312131"></a>

## Next pages — blindfold_secret_info / 310101102000 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--voltstack_site--reference--group-006.md#canonical-0113330003101111-2120321331121121-2132013220212333-0013110203013132-2103112030213122-3032333122221222-1131332203001223-3010221331303233)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1213231011320213-1032230221203231-3000211123313320-1110013101102201-0330302212000300-3002220110030123-2203210200121030-3021302212322123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312011223001303-2222123012033000-1030323102311302-1012331020101022-2123321003300231-0333323013321313-1322221131313113-3110223201030220"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info — clear_secret_info / 033031201111 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--voltstack_site--reference--group-006.md#canonical-0113330003101111-2120321331121121-2132013220212333-0013110203013132-2103112030213122-3032333122221222-1131332203001223-3010221331303233)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info

<a id="canonical-0211201030011303-0323231302033130-1233310313201021-2101013213310111-3021203132121203-0033120301330001-1030223231123033-1101200220233020"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233112130012111-3230002032033232-0230320321213212-2100003211310323-1321300303102122-2032012310201131-1101101232301302-2200023112101231"></a>

## Direct properties — clear_secret_info / 033031201111 / 3

<a id="canonical-0122020311223122-1302322013123330-1320210032102202-2210313131310211-1011123032203320-1310301333002133-1022032203330223-2300230301133301"></a>

<a id="canonical-2212020012121132-2033333111013110-1003210221111320-2221330202212212-0010302111322220-3123120210011102-3021323220112133-0001313031000322"></a>

## provider_ref property — clear_secret_info / 033031201111 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3111120001321331-0331011002210201-0130333233010023-2232312321032033-3332302302312102-2200201322323101-3031002233123232-0232013120133120"></a>

<a id="canonical-0131021330320133-3232210122201031-1013002333312021-3130012130100301-1200132333000213-3323201013331120-3122032312012301-1221011131020012"></a>

## URL property — clear_secret_info / 033031201111 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1121212230323222-3120110312022011-0233112322323302-0301111032120211-1222030231222030-0111303123203332-0100223001222331-0002203130002012"></a>

## Next pages — clear_secret_info / 033031201111 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--voltstack_site--reference--group-006.md#canonical-0113330003101111-2120321331121121-2132013220212333-0013110203013132-2103112030213122-3032333122221222-1131332203001223-3010221331303233)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0003212002302320-0230230222133000-1230132033022210-1303013121300202-1221031233022031-2131010312122002-0032102203210130-1123120030221322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111323211210030-1213203311101113-1232003233103303-0031021333312200-2132223022113220-3013111202122120-3303033301323130-0222231302223011"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage — storage / 010221312201 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage

<a id="canonical-0232230230121013-0310120311020220-3232312111202022-0300031111002131-0202122033211100-1112101320200310-0201233300303222-2112201030012031"></a>

Type: `"object"`. list nested block, Optional.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131300320023313-2001201323101030-0231002230132122-0110332203023011-3223102111311323-3201232311120112-0103312111102110-1233220032030313"></a>

## Direct properties — storage / 010221312201 / 3

<a id="canonical-3013213031331122-1123030131000311-1312231021231033-0121131023030203-1301110212003022-1212100010001332-1213231131130220-2222231312311323"></a>

<a id="canonical-1201300033023321-1001222121010033-3111213331202211-2232123323001023-1202302001302020-3300222001221121-1022233301111101-0001020000011011"></a>

## labels property — storage / 010221312201 / 4

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":20},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"20\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
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
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--voltstack_site--reference--group-006.md#canonical-2120111121320133-0022321320313331-2132202303320221-3101330310100212-0030302300221333-0210323203203332-3113022013211001-1030313221311311): complete subsection reference.

<a id="canonical-0333300003111120-0303012302120232-2312233321323133-3102022103212022-3031301102331113-1112223000122121-2022001100000133-2300212333122330"></a>

<a id="canonical-0221020322121120-0231321312301222-1202110112130323-0103332112203113-3121110211112213-3103100133203130-0122323032331033-2013233003010333"></a>

## zone property — storage / 010221312201 / 5

Type: `"string"`. Optional.

Virtual Pool Zone. Virtual Storage Pool zone definition.

Upstream description:

Virtual Storage Pool zone definition.

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

<a id="canonical-0333331023123303-1220032331323330-2233121220331121-3003230312001331-0222022121301031-1120231122103010-1301033020331121-1013123012331031"></a>

## Next pages — storage / 010221312201 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--voltstack_site--reference--group-006.md#canonical-2120111121320133-0022321320313331-2132202303320221-3101330310100212-0030302300221333-0210323203203332-3113022013211001-1030313221311311)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2120111121320133-0022321320313331-2132202303320221-3101330310100212-0030302300221333-0210323203203332-3113022013211001-1030313221311311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022023332030132-0120210203103231-2220202333332312-1311103302000200-3221011011212101-0312110303232031-1212201000210210-0211120221131011"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults — volume_defaults / 103012211211 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--voltstack_site--reference--group-006.md#canonical-0003212002302320-0230230222133000-1230132033022210-1303013121300202-1221031233022031-2131010312122002-0032102203210130-1123120030221322)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults

<a id="canonical-1211232321130010-2013120120203213-0203303311323232-1311131111203320-1012000031031322-3000123232233302-3130301310301111-2210112012003311"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
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
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223203301011200-1332032103203031-2320102013333303-3233132100133003-2200013330210210-2030111000013131-3302111132313101-2011321122023222"></a>

## Direct properties — volume_defaults / 103012211211 / 3

<a id="canonical-0300110003233013-0012332333213121-0300033101213333-2032303131311021-2011330030021333-1010122021230113-2101332303302000-1010003200220002"></a>

<a id="canonical-3031312003001012-0013120211021101-1210103032010222-1320103123203101-3002202232113332-2132120011211000-3303221000332112-3110132312211331"></a>

## adaptive_qos_policy property — volume_defaults / 103012211211 / 4

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3003132021321131-0212003121022323-1203012133112323-0311321200200330-1321202103303322-1001210132302203-2301011022303220-1310003313212212"></a>

<a id="canonical-3122120331210021-1101203121103022-3031102111221320-0123322323213300-3010003221033003-0103221321031212-2300001023220101-3323011313010221"></a>

## encryption property — volume_defaults / 103012211211 / 5

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

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

<a id="canonical-2110322103133000-1122232201203212-2211001330022122-1321220211000310-0011133130113120-3121200002210310-3013321333321130-1032321322103103"></a>

<a id="canonical-1001100013212310-2310231120000302-0032322323331100-1300320323330133-3310322123220103-0122013332330010-2222312022333231-3112122133132311"></a>

## export_policy property — volume_defaults / 103012211211 / 6

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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

- [no_qos](resources--voltstack_site--reference--group-006.md#canonical-3011333011112312-3012222103031110-1220130001220321-3003123211002012-3330211003000321-0200122323000301-2111322222320221-0022033133203001): complete subsection reference.

<a id="canonical-3122211333232120-2302033322032310-3032232232020133-1210022320110321-3031023313233220-3122122123010303-2002332013330333-2313132320121213"></a>

<a id="canonical-0112011221002002-1022131113301133-2202323020323123-3003302110332331-3103111233011201-3003102130012103-2032232202021331-1213010330210022"></a>

## qos_policy property — volume_defaults / 103012211211 / 7

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0222021313201131-2002112022030130-2030231131210133-0330123113223130-0220332200320313-3002021131231101-1101322133322001-0323203031222300"></a>

<a id="canonical-3131123200323213-2003023212302200-1111111021031100-3122203130313123-2313002331100132-1330222221102022-2001030011233320-1102202303000323"></a>

## security_style property — volume_defaults / 103012211211 / 8

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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

<a id="canonical-2110022333213332-0330132002112233-3000330130011112-1330201311001223-2333203131231330-3223221110211312-2000202332311301-0013231131203323"></a>

<a id="canonical-3212120323031201-0010302213330303-3230232303122231-1333001013121333-0100203211312300-2201002012102321-1001023213213333-2232102022032300"></a>

## snapshot_dir property — volume_defaults / 103012211211 / 9

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

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

<a id="canonical-2331302020103022-0220222013001312-2021033333130321-0120111301213110-3233031100331300-3232213102322221-1100132033132012-1010032302131232"></a>

<a id="canonical-0211222102130220-3002222013210031-1003131133222100-1211023213113210-0202310321002202-0232000231113331-2331123002331113-3203331122203200"></a>

## snapshot_policy property — volume_defaults / 103012211211 / 10

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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

<a id="canonical-3013313203331121-3231132020121220-0333203210000313-1010021133222100-3020211223222323-2122132232003232-3003330001231111-1110212032101331"></a>

<a id="canonical-2122210202300220-1013122101300001-1011110022132220-1122102211330112-3330101231212013-3221320200313213-2301022233201200-3232113212331001"></a>

## snapshot_reserve property — volume_defaults / 103012211211 / 11

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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

<a id="canonical-3330320011102313-1011131023031203-0310031030232222-3300303133111001-1222033032011113-0220322030323110-2012101131130120-3003212300111123"></a>

<a id="canonical-1333100032120322-2320113213333202-2230333120013010-2132133102012010-0303012133020002-2110303211031202-0302103213302231-2023013122332100"></a>

## space_reserve property — volume_defaults / 103012211211 / 12

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-1201000010201323-0131121300311320-2001322312023102-2321030023322330-1201210313123132-3210310311123323-2222211112322301-1211203222103113"></a>

<a id="canonical-1123303113111200-1101011033100112-2323310111133331-3313330320313121-2211313303312012-3301210201212312-2313022010330112-2303011302111310"></a>

## split_on_clone property — volume_defaults / 103012211211 / 13

Type: `"bool"`. Optional.

Split a clone from its parent upon creation.

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

<a id="canonical-2122002321133221-3231023012011211-3102231211021013-0132331230113210-3131330312212010-2010012023223100-1123313000203213-3021020312030200"></a>

<a id="canonical-0121112033031023-2031313033131213-3203302210221102-0002010101310111-3002101233231031-1030311033111012-0323001030123031-0111110323003101"></a>

## tiering_policy property — volume_defaults / 103012211211 / 14

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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

<a id="canonical-1030203001000121-3001231132233110-1322303312113012-2311223210213330-2003300012301023-0202213023201301-3220212320021323-0000121302020130"></a>

<a id="canonical-2220313011120120-1321121111133001-2013132023233132-0311230103203113-3332221023001111-2002313221112330-2000022212332213-2121222311100100"></a>

## unix_permissions property — volume_defaults / 103012211211 / 15

Type: `"number"`. Optional.

Unix permission mode for new volumes. All allowed 777.

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

<a id="canonical-0102301111101123-3111022013222023-3032303330220100-0223330213013233-2010021310222303-2220230211012122-0333230222323113-2100012332322232"></a>

## Next pages — volume_defaults / 103012211211 / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos](resources--voltstack_site--reference--group-006.md#canonical-3011333011112312-3012222103031110-1220130001220321-3003123211002012-3330211003000321-0200122323000301-2111322222320221-0022033133203001)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--voltstack_site--reference--group-006.md#canonical-0003212002302320-0230230222133000-1230132033022210-1303013121300202-1221031233022031-2131010312122002-0032102203210130-1123120030221322)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3011333011112312-3012222103031110-1220130001220321-3003123211002012-3330211003000321-0200122323000301-2111322222320221-0022033133203001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132332212112211-3122231010003311-1301013103000003-3211301213313233-1233320013302203-1320030200111303-1233012300130001-2233201321311003"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos — no_qos / 111123320101 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--voltstack_site--reference--group-006.md#canonical-0003212002302320-0230230222133000-1230132033022210-1303013121300202-1221031233022031-2131010312122002-0032102203210130-1123120030221322)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--voltstack_site--reference--group-006.md#canonical-2120111121320133-0022321320313331-2132202303320221-3101330310100212-0030302300221333-0210323203203332-3113022013211001-1030313221311311)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos

<a id="canonical-1330121230223221-1313302221330323-2230001203230113-3013111312023021-0211313101323300-2130303011310320-1100003101033101-3031222121302201"></a>

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
no_qos = {}
```

<a id="canonical-1010210221132323-1223030011121111-3210310121233201-1103311133311311-1030131201020312-3222220312110030-3210213222301110-3312202021212220"></a>

## Direct properties — no_qos / 111123320101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102012101311020-3133030120211233-0332002222301010-0030112331302011-3311132130032103-0131110210300312-0113323031200133-0331320002211030"></a>

## Next pages — no_qos / 111123320101 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--voltstack_site--reference--group-006.md#canonical-2120111121320133-0022321320313331-2132202303320221-3101330310100212-0030302300221333-0210323203203332-3113022013211001-1030313221311311)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2113210123032330-2220213133030132-0302222012203211-0333131032213120-0232121001321012-1121201012031013-3033112013220010-2332201222201032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033211301110231-0323302230330033-1232121303011113-1201300310122131-2013132032310230-3033123003033132-1032210312112111-2212112113012223"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults — volume_defaults / 331001322121 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults

<a id="canonical-0300103213011023-1211310010122333-0323102202313311-3322210233310322-1223321301021013-3201021213330223-3011221310203330-3123100331020010"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
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
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330130331012011-2300321212201331-3302121313312033-3223013212211313-3022130213123310-0012313101331112-1330320213010013-0001133320111033"></a>

## Direct properties — volume_defaults / 331001322121 / 3

<a id="canonical-2313130120012312-2230122010303213-2323332001203002-2223032121302030-3332133230212022-2033302221332222-1033100222012212-3221132301102123"></a>

<a id="canonical-3132230203133303-3021303031303320-2221010133032000-3313100122302103-0111133310012121-1121213010103122-0101220030231112-3331231333110101"></a>

## adaptive_qos_policy property — volume_defaults / 331001322121 / 4

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3331231332303101-1310033133233020-1110200320313212-2030210321021220-3013302111132321-2013221202112033-0002302222221303-2001102302122003"></a>

<a id="canonical-0123323323201302-1100211123230113-2322300313110131-3011000301110031-3321013312102223-2301213021003211-1010023021203323-2332001211302013"></a>

## encryption property — volume_defaults / 331001322121 / 5

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

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

<a id="canonical-0101221330323120-1230222120101320-1003202100022203-3001310221021302-2213333130313310-3301113213311132-0221231101013122-2232013111130212"></a>

<a id="canonical-2021332302332100-1202110222032223-0222011021021032-3012223110001133-2032300120023112-1320002313110133-3130103223122100-2233013323010032"></a>

## export_policy property — volume_defaults / 331001322121 / 6

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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

- [no_qos](resources--voltstack_site--reference--group-006.md#canonical-2202130202333322-3322111120011303-3112331121233130-3202010230200103-3212003322333103-1023300103312023-1213130313210320-3321111101320112): complete subsection reference.

<a id="canonical-1133013203132201-2210011133320130-1131020332201311-3101210331111012-0100011103122233-0032133310113101-2211021210212102-2001221230011012"></a>

<a id="canonical-1332311331103132-1221003102332122-2330203122112110-0310123122202322-1021310230131323-1221211231022233-1211131202003011-2132133113212100"></a>

## qos_policy property — volume_defaults / 331001322121 / 7

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1011123311130001-0302013203311320-2120013313133301-0033202012333030-3322301230310002-2333332202121013-0111222312031031-0333310012322123"></a>

<a id="canonical-0321100131100010-1112032133311030-2232031312110122-3221133221130213-3133202312302003-1023113330120130-0120310022203202-2013231133210031"></a>

## security_style property — volume_defaults / 331001322121 / 8

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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

<a id="canonical-0301100301313312-0222012030132320-2113302222013021-2112101110013300-0301321332233202-1000333011330313-0233131213213101-1221110320111002"></a>

<a id="canonical-0331133310200233-2030220312122131-1220323211110012-0111101231310203-1301120112230300-2122031320011321-0002332030222200-0031020220021212"></a>

## snapshot_dir property — volume_defaults / 331001322121 / 9

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

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

<a id="canonical-1332222221113213-0312113231210120-0031113002011031-0203101210000032-3030012033221013-2321311021213003-3200310122120113-3200133313031201"></a>

<a id="canonical-1023320302311032-2002132130201332-1210130310320233-1223211031130321-2121222302202201-2111132100123210-0000321020131001-2111030302022302"></a>

## snapshot_policy property — volume_defaults / 331001322121 / 10

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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

<a id="canonical-3020333303110013-0011132210332110-2311203012020322-1101120330002233-2200031131023131-0302123010113012-1011121033303010-3300111323321311"></a>

<a id="canonical-1131233202302303-2033001311012301-0230031231130330-3112221032310103-2121312132222321-1120333103310123-3031102233321130-1102112301200320"></a>

## snapshot_reserve property — volume_defaults / 331001322121 / 11

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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

<a id="canonical-2022203330331123-2203022021013132-0022333200000211-3330133130311202-1103212301210212-2022000031021110-2200200300311323-3233013310021131"></a>

<a id="canonical-0001202131223122-2223023001011320-0111022101110333-3211101222011300-0220030303200331-0031002030233330-2001033012310012-2123030021000232"></a>

## space_reserve property — volume_defaults / 331001322121 / 12

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-0313311222021013-1120301032203233-3322020021123231-0003232222332213-0133320303033133-2202311121310103-1202020020200011-3213213100301210"></a>

<a id="canonical-3020320113201301-2022012302012122-0220223023321200-0312013120303113-2200003303133032-2020023120020023-3310311313332322-1121002032330310"></a>

## split_on_clone property — volume_defaults / 331001322121 / 13

Type: `"bool"`. Optional.

Split a clone from its parent upon creation.

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

<a id="canonical-2031220222210033-0332031312313213-1323132223031032-2333311333312211-3232130312023132-3131323011023202-1202112123333021-1023201320213223"></a>

<a id="canonical-1300120302100201-2301320032133211-2232210210003032-1131210203013330-0310101223233301-3332200100133113-0233110310221023-3220331122210301"></a>

## tiering_policy property — volume_defaults / 331001322121 / 14

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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

<a id="canonical-1231131122232333-1130101012100201-1232031032033232-1101120203311331-0133220123010311-2230320332030003-1013120111233301-3110302020100223"></a>

<a id="canonical-1112013321112133-1021322102311000-0222102202112102-0301132222120311-3332132012120303-2212011123111013-0120333220330000-2303033200313310"></a>

## unix_permissions property — volume_defaults / 331001322121 / 15

Type: `"number"`. Optional.

Unix permission mode for new volumes. All allowed 777.

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

<a id="canonical-3123301201000010-2003010310010231-0030203212111010-1011023300123133-2200332132001123-1330010233103211-1322212031211000-2030102032222221"></a>

## Next pages — volume_defaults / 331001322121 / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos](resources--voltstack_site--reference--group-006.md#canonical-2202130202333322-3322111120011303-3112331121233130-3202010230200103-3212003322333103-1023300103312023-1213130313210320-3321111101320112)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-3033222101303330-1011123222310233-1330220110033121-3331113222323101-3201110213321220-1120133323300233-1122122130112111-0131320210221213)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2202130202333322-3322111120011303-3112331121233130-3202010230200103-3212003322333103-1023300103312023-1213130313210320-3321111101320112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
