---
page_title: "xcsh_forwarding_class reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_forwarding_class reference."
---

# xcsh_forwarding_class reference

<a id="canonical-0102111112020232-1131131013230311-1223320113131321-1013330221212303-3022123312213323-3323312223011303-2202303002230110-1331212022130030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331111211013312-1222311331320310-0003213331322102-3120111121111331-1313000312112302-2110013320010113-1202333223131330-1320312233001101"></a>

## Property reference — Property reference / 031112300221 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)
- Property reference

<a id="canonical-0210330120230111-1300323111302000-0131001123330003-0322101213032022-1310330133302121-2323223110003032-0333303133122231-1113011300303002"></a>

## Direct properties — Property reference / 031112300221 / 3

<a id="canonical-1331002201302121-0233230202221002-1030011220010112-1331131333123013-3130322223301102-0012223311111321-1100102123122310-2103220010111120"></a>

<a id="canonical-2302213231121111-0133300211110121-1113001332233011-0122320211033223-2331012211010131-1130330313032122-2011202133120333-0000012010330303"></a>

## annotations property — Property reference / 031112300221 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-3322101132221210-3321301133010120-0231333113133212-0313210100023330-0111132020000032-0121302100203022-2110022120113312-3301011012223122"></a>

<a id="canonical-2132230210312113-1102111203131130-1301322223033211-3122323102023003-0100130200130113-1022233223203231-1313030103000221-3120010102001310"></a>

## description property — Property reference / 031112300221 / 5

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2120121201203020-1220313300322203-0020110322002310-3213010021301033-1123130322213120-2121232013003010-2113033332330132-1332111021101100"></a>

<a id="canonical-2022213003230112-2211200122200210-3301001130012102-3031011110222322-0223221123010100-1111201200101032-2110211102330021-1010210102123220"></a>

## disable property — Property reference / 031112300221 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [dscp](resources--forwarding_class--reference--group-001.md#canonical-1211022211101312-3321220233030131-0310113331012110-3002002203030223-0201230200010121-3001101111002300-3333213313003122-3033122032222001): complete subsection reference.

- [dscp_based_queue](resources--forwarding_class--reference--group-001.md#canonical-3302113323222301-0133012130232111-0221312101202123-3223011101220123-0303002133022202-0022010122103221-0232201320101121-3001131002330320): complete subsection reference.

<a id="canonical-3320021021302332-1333233030013212-2312231011023133-0103132130102213-3102233222220331-3301113332211303-0113301221012013-0103130103102012"></a>

<a id="canonical-0020130022020301-3300131221100033-1023330331212222-0301211322031312-3321113021200312-3033020012121003-3113120132021002-3110033300030122"></a>

## ID property — Property reference / 031112300221 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2120111112202230-0012311210121032-2231301010022033-3032203220022133-1100102010222230-3022123321000121-1120223113321223-2212002002231332"></a>

<a id="canonical-0133333310101101-0311101102031002-1033213132023330-0332001230332200-0302112313300032-0220223301330100-2000020000212130-1123012233123210"></a>

## interface_group property — Property reference / 031112300221 / 8

Type: `"string"`. Optional, Computed.

\[Enum: ANY\_AVAILABLE\_INTERFACE|INTERFACE\_GROUP1|INTERFACE\_GROUP2|INTERFACE\_GROUP3\] Interface
group, group membership by adding group label to interface Choose any of the available interfaces
Choose all interfaces with label group1 Choose all interfaces with label group2 Choose all
interfaces with label group3. Possible values are \`ANY\_AVAILABLE\_INTERFACE\`,
\`INTERFACE\_GROUP1\`, \`INTERFACE\_GROUP2\`, \`INTERFACE\_GROUP3\`. Defaults to
\`ANY\_AVAILABLE\_INTERFACE\`.

Upstream description:

Interface group, group membership by adding group label to interface

Choose any of the available interfaces Choose all interfaces with label group1 Choose all interfaces
with label group2 Choose all interfaces with label group3.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY_AVAILABLE_INTERFACE",
    "INTERFACE_GROUP1",
    "INTERFACE_GROUP2",
    "INTERFACE_GROUP3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY_AVAILABLE_INTERFACE",
  "enum": [
    "ANY_AVAILABLE_INTERFACE",
    "INTERFACE_GROUP1",
    "INTERFACE_GROUP2",
    "INTERFACE_GROUP3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1102002013223221-3203232311330012-2301211331131002-1323222120033201-3031301202011102-0231133312211013-0311232233201012-0321113001020122"></a>

<a id="canonical-3212122123031102-3000223322113023-1011220012002023-2222031000110233-0031102303011333-0333311321220332-1231211230323310-2013333221122200"></a>

## labels property — Property reference / 031112300221 / 9

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Upstream description:

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

<a id="canonical-1322132101222221-1231231032001010-0322020313333032-3231021202302032-2121000131000213-3013302223232011-2220101110323010-3011313131100301"></a>

<a id="canonical-2132031331213302-0130310131321122-3202222202311321-1331010022223133-1122033232133110-1022102111310120-3332121213010120-3201101133332020"></a>

## name property — Property reference / 031112300221 / 10

Type: `"string"`. Required.

Name of the Forwarding Class. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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

<a id="canonical-3212132011101310-1231013330313120-1233213332211031-1033032330013102-0112120011220213-2310333201212113-1201013202333031-1313111132112121"></a>

<a id="canonical-0130000101133000-2232221101031203-1313221001210213-1012030300130032-2203303101300313-0002321001310113-1022200100301230-0120121123012122"></a>

## namespace property — Property reference / 031112300221 / 11

Type: `"string"`. Required.

Namespace where the Forwarding Class is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [no_marking](resources--forwarding_class--reference--group-001.md#canonical-1021103130221111-1123033322313003-2032001203200330-0001333002012023-3323021332212211-3113320333001230-1101113133200332-1310201221313033): complete subsection reference.

- [no_policer](resources--forwarding_class--reference--group-001.md#canonical-1102211202003033-1303000013132220-0201013320301323-2022132212011332-3111131211020213-2203020212003232-3311010001321232-2132301130323132): complete subsection reference.

- [policer](resources--forwarding_class--reference--group-001.md#canonical-1110212130302302-3120310323312010-0221323020122301-1332103131311011-2232123320231231-3301310021300320-2102320131210102-3112321331101001): complete subsection reference.

<a id="canonical-1322111122031103-3103323233203003-3112230313100031-0100033221020110-2203020000120200-0332200321213112-3323022021320010-3000022023212223"></a>

<a id="canonical-3030130133102211-0123133102023221-0231313220001121-2231212213233010-1112203322221100-3222012323100212-1330200231232132-2323300311000122"></a>

## queue_id_to_use property — Property reference / 031112300221 / 12

Type: `"string"`. Optional, Computed.

\[Enum:
DSCP\_BEST\_EFFORT|DSCP\_CLASS1|DSCP\_CLASS2|DSCP\_CLASS3|DSCP\_CLASS4|DSCP\_EXPRESS\_FORWARDING|DSCP\_CONTROL\_L3|DSCP\_CONTROL\_L2\]
DSCP Precedence Level Values Best Effort service will GET any available bandwidth DSCP Class 1
service DSCP Class 2 service DSCP Class 3 service DSCP Class 4 service Express Forwarding is used
for low latency traffic Control is used for routing traffic, not recommended Link Layer traffic
like.. Possible values are \`DSCP\_BEST\_EFFORT\`, \`DSCP\_CLASS1\`, \`DSCP\_CLASS2\`,
\`DSCP\_CLASS3\`, \`DSCP\_CLASS4\`, \`DSCP\_EXPRESS\_FORWARDING\`, \`DSCP\_CONTROL\_L3\`,
\`DSCP\_CONTROL\_L2\`. Defaults to \`DSCP\_BEST\_EFFORT\`.

Upstream description:

DSCP Precedence Level Values

Best Effort service will GET any available bandwidth DSCP Class 1 service DSCP Class 2 service DSCP
Class 3 service DSCP Class 4 service Express Forwarding is used for low latency traffic Control is
used for routing traffic, not recommended Link Layer traffic like LACP or keepalive, not
recommended.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DSCP_BEST_EFFORT",
    "DSCP_CLASS1",
    "DSCP_CLASS2",
    "DSCP_CLASS3",
    "DSCP_CLASS4",
    "DSCP_EXPRESS_FORWARDING",
    "DSCP_CONTROL_L3",
    "DSCP_CONTROL_L2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DSCP_BEST_EFFORT",
  "enum": [
    "DSCP_BEST_EFFORT",
    "DSCP_CLASS1",
    "DSCP_CLASS2",
    "DSCP_CLASS3",
    "DSCP_CLASS4",
    "DSCP_EXPRESS_FORWARDING",
    "DSCP_CONTROL_L3",
    "DSCP_CONTROL_L2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [timeouts](resources--forwarding_class--reference--group-001.md#canonical-1232131200303203-0031203210012133-2012313021120110-2230300330031020-2123313132213221-3101202203102203-2301230332323320-3233212010001321): complete subsection reference.

<a id="canonical-1120331003030132-2233201101101201-3123212313120330-2033111210020030-0122113023022301-2213113113323300-3002300310221322-2330122230132223"></a>

<a id="canonical-0200231212321300-2112100321210311-3110213133030021-3203233212010021-1100031001310211-3103033132202210-1231001132232332-3231113231001102"></a>

## tos_value property — Property reference / 031112300221 / 13

Type: `"number"`. Optional, Computed.

Exclusive with \[dscp no\_marking\] Decimal value of raw 8 bit TOS. In above example DSCP 10 =
Precedence Class 1 and drop precedence low.

Upstream description:

Exclusive with \[dscp no\_marking\] Decimal value of raw 8 bit TOS. In above example DSCP 10 =
Precedence Class 1 and drop precedence low.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
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
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-0210003122102300-1002220131022030-0012213231212213-3233311301010111-1230020120322003-0212123312132300-3031213032320131-1203221033230020"></a>

## All schema paths — Property reference / 031112300221 / 14

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--forwarding_class--reference--group-001.md#canonical-1331002201302121-0233230202221002-1030011220010112-1331131333123013-3130322223301102-0012223311111321-1100102123122310-2103220010111120) |
| `description` | [description](resources--forwarding_class--reference--group-001.md#canonical-3322101132221210-3321301133010120-0231333113133212-0313210100023330-0111132020000032-0121302100203022-2110022120113312-3301011012223122) |
| `disable` | [disable](resources--forwarding_class--reference--group-001.md#canonical-2120121201203020-1220313300322203-0020110322002310-3213010021301033-1123130322213120-2121232013003010-2113033332330132-1332111021101100) |
| `dscp` | [dscp](resources--forwarding_class--reference--group-001.md#canonical-3120011111020103-1332133221102231-1122323212300030-0332313310210002-1122300212231102-2032200033301330-2011132120311123-2120212220131221) |
| `dscp.drop_precedence` | [dscp.drop_precedence](resources--forwarding_class--reference--group-001.md#canonical-1201211012312201-0011100321031232-0123221331132022-2132222130122323-0321332213021201-3012330022313321-0320113112210020-3013313032032013) |
| `dscp.dscp_class` | [dscp.dscp_class](resources--forwarding_class--reference--group-001.md#canonical-3120210213132320-2220201011201233-3112200223100103-2203233033121130-0110032302022322-1200210320321133-0221230130321001-2213113331102111) |
| `dscp_based_queue` | [dscp_based_queue](resources--forwarding_class--reference--group-001.md#canonical-2003303032021301-1232211011322121-3213200122000002-3300223200211003-2203003201223312-1032010130230010-2321013131010223-2302220220222212) |
| `id` | [ID](resources--forwarding_class--reference--group-001.md#canonical-3320021021302332-1333233030013212-2312231011023133-0103132130102213-3102233222220331-3301113332211303-0113301221012013-0103130103102012) |
| `interface_group` | [interface_group](resources--forwarding_class--reference--group-001.md#canonical-2120111112202230-0012311210121032-2231301010022033-3032203220022133-1100102010222230-3022123321000121-1120223113321223-2212002002231332) |
| `labels` | [labels](resources--forwarding_class--reference--group-001.md#canonical-1102002013223221-3203232311330012-2301211331131002-1323222120033201-3031301202011102-0231133312211013-0311232233201012-0321113001020122) |
| `name` | [name](resources--forwarding_class--reference--group-001.md#canonical-1322132101222221-1231231032001010-0322020313333032-3231021202302032-2121000131000213-3013302223232011-2220101110323010-3011313131100301) |
| `namespace` | [namespace](resources--forwarding_class--reference--group-001.md#canonical-3212132011101310-1231013330313120-1233213332211031-1033032330013102-0112120011220213-2310333201212113-1201013202333031-1313111132112121) |
| `no_marking` | [no_marking](resources--forwarding_class--reference--group-001.md#canonical-1021021232202130-0331303132133101-0203200111202303-2133300311022332-3113320321332201-1013122030310303-0020001112113302-3011130210330310) |
| `no_policer` | [no_policer](resources--forwarding_class--reference--group-001.md#canonical-3022012312002022-1132231322001333-0102031010313100-2013202021201000-0320303323033030-2301131023231121-1012311132123023-1003211331301231) |
| `policer` | [policer](resources--forwarding_class--reference--group-001.md#canonical-2323331002012330-2003001101221233-0000222212132011-3031011202312322-3020001322030022-3211012120223213-0131111202331112-1102113322213301) |
| `policer.name` | [policer.name](resources--forwarding_class--reference--group-001.md#canonical-1311101010330000-0101121122120002-2200111120233303-3023012210123230-3321230112131220-2213012113010313-1030003202021213-2120102103132133) |
| `policer.namespace` | [policer.namespace](resources--forwarding_class--reference--group-001.md#canonical-2003202000101000-0230021031122220-1032311010203131-1333121103210200-1232223101303020-0000202132212023-1231031100111012-0120000010320032) |
| `policer.tenant` | [policer.tenant](resources--forwarding_class--reference--group-001.md#canonical-2030230001310000-0100330021032033-1130023301223220-2301200303030001-3032100211211122-0323203110121212-3031213330201220-3132131200220003) |
| `queue_id_to_use` | [queue_id_to_use](resources--forwarding_class--reference--group-001.md#canonical-1322111122031103-3103323233203003-3112230313100031-0100033221020110-2203020000120200-0332200321213112-3323022021320010-3000022023212223) |
| `timeouts` | [timeouts](resources--forwarding_class--reference--group-001.md#canonical-3033312332113202-1200133232200323-2202103220110330-1210002110203302-2102230220210330-2311331200310302-3231233200130110-2313322103302001) |
| `timeouts.create` | [timeouts.create](resources--forwarding_class--reference--group-001.md#canonical-1201000011031201-3203220211013220-1113102112300313-0220213032323032-1122113320003132-3011231223301033-2222033020033023-1101132013202130) |
| `timeouts.delete` | [timeouts.delete](resources--forwarding_class--reference--group-001.md#canonical-3012332323201130-2001213310032021-0113332100021022-3310312322111122-0320011103313323-0232010033210202-2321132013223210-1321313302333111) |
| `timeouts.read` | [timeouts.read](resources--forwarding_class--reference--group-001.md#canonical-2123322301122111-2112203220302212-0231220133201110-0123030023303000-2002231230013023-2101030310021010-0001110122132130-3130102013122003) |
| `timeouts.update` | [timeouts.update](resources--forwarding_class--reference--group-001.md#canonical-1230011013121103-0201212310202113-3111033332103133-1122301222120223-3013231221230102-2010013032002210-0300001100111113-2021333012231300) |
| `tos_value` | [tos_value](resources--forwarding_class--reference--group-001.md#canonical-1120331003030132-2233201101101201-3123212313120330-2033111210020030-0122113023022301-2213113113323300-3002300310221322-2330122230132223) |

<a id="canonical-2220321302020021-0221103313100333-0012110131003312-1212313202100022-0312223221220300-2200101023112110-0020030230032032-2012321111200013"></a>

## Next pages — Property reference / 031112300221 / 15

- [dscp](resources--forwarding_class--reference--group-001.md#canonical-1211022211101312-3321220233030131-0310113331012110-3002002203030223-0201230200010121-3001101111002300-3333213313003122-3033122032222001)
- [dscp_based_queue](resources--forwarding_class--reference--group-001.md#canonical-3302113323222301-0133012130232111-0221312101202123-3223011101220123-0303002133022202-0022010122103221-0232201320101121-3001131002330320)
- [no_marking](resources--forwarding_class--reference--group-001.md#canonical-1021103130221111-1123033322313003-2032001203200330-0001333002012023-3323021332212211-3113320333001230-1101113133200332-1310201221313033)
- [no_policer](resources--forwarding_class--reference--group-001.md#canonical-1102211202003033-1303000013132220-0201013320301323-2022132212011332-3111131211020213-2203020212003232-3311010001321232-2132301130323132)
- [policer](resources--forwarding_class--reference--group-001.md#canonical-1110212130302302-3120310323312010-0221323020122301-1332103131311011-2232123320231231-3301310021300320-2102320131210102-3112321331101001)
- [timeouts](resources--forwarding_class--reference--group-001.md#canonical-1232131200303203-0031203210012133-2012313021120110-2230300330031020-2123313132213221-3101202203102203-2301230332323320-3233212010001321)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)

<a id="canonical-1211022211101312-3321220233030131-0310113331012110-3002002203030223-0201230200010121-3001101111002300-3333213313003122-3033122032222001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222333100212211-1303200121003130-3322131013202031-1302233303220123-3011323202200322-3110200303302200-3332133230310020-1102031011201132"></a>

## dscp — dscp / 320012321330 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)
- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-0102111112020232-1131131013230311-1223320113131321-1013330221212303-3022123312213323-3323312223011303-2202303002230110-1331212022130030)
- dscp

<a id="canonical-3120011111020103-1332133221102231-1122323212300030-0332313310210002-1122300212231102-2032200033301330-2011132120311123-2120212220131221"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dscp, no\_marking, tos\_value; Default: no\_marking\] DSCP Marking setting. DSCP marking
setting as per RFC 2475.

Upstream description:

DSCP marking setting as per RFC 2475.

Receipt-pinned upstream constraints:

```json
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

- [dscp](resources--forwarding_class--reference--group-001.md#canonical-3120011111020103-1332133221102231-1122323212300030-0332313310210002-1122300212231102-2032200033301330-2011132120311123-2120212220131221)
- [no_marking](resources--forwarding_class--reference--group-001.md#canonical-1021021232202130-0331303132133101-0203200111202303-2133300311022332-3113320321332201-1013122030310303-0020001112113302-3011130210330310)
- [tos_value](resources--forwarding_class--reference--group-001.md#canonical-1120331003030132-2233201101101201-3123212313120330-2033111210020030-0122113023022301-2213113113323300-3002300310221322-2330122230132223)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dscp {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312112200103221-2320321202002112-0023332133032213-3200120311312231-3120130030011220-2003120011302030-2333130031320113-3200102000203311"></a>

## Direct properties — dscp / 320012321330 / 3

<a id="canonical-1201211012312201-0011100321031232-0123221331132022-2132222130122323-0321332213021201-3012330022313321-0320113112210020-3013313032032013"></a>

<a id="canonical-2231333001302131-3131231112213332-2321223302211210-1202020311120112-3333032112113001-0021232033132310-1323102112333300-3120032212133312"></a>

## drop_precedence property — dscp / 320012321330 / 4

Type: `"string"`. Optional.

\[Enum: DSCP\_AF\_LOW|DSCP\_AF\_MEDIUM|DSCP\_AF\_HIGH|DSCP\_AF\_POLICER\] DSCP Assured forwarding
drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP drop
precedence value is taken from output of policer. Possible values are \`DSCP\_AF\_LOW\`,
\`DSCP\_AF\_MEDIUM\`, \`DSCP\_AF\_HIGH\`, \`DSCP\_AF\_POLICER\`.

Upstream description:

DSCP Assured forwarding drop precedence

DSCP Low drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP drop precedence
value is taken from output of policer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DSCP_AF_LOW",
    "DSCP_AF_MEDIUM",
    "DSCP_AF_HIGH",
    "DSCP_AF_POLICER"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "DSCP_AF_LOW",
    "DSCP_AF_MEDIUM",
    "DSCP_AF_HIGH",
    "DSCP_AF_POLICER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3120210213132320-2220201011201233-3112200223100103-2203233033121130-0110032302022322-1200210320321133-0221230130321001-2213113331102111"></a>

<a id="canonical-3230211311333331-0232210312220132-1201313021030332-1102223100323303-3221231203003222-0310331012323122-2301013232133100-3120200313133012"></a>

## dscp_class property — dscp / 320012321330 / 5

Type: `"string"`. Optional.

\[Enum:
DSCP\_BEST\_EFFORT|DSCP\_CLASS1|DSCP\_CLASS2|DSCP\_CLASS3|DSCP\_CLASS4|DSCP\_EXPRESS\_FORWARDING|DSCP\_CONTROL\_L3|DSCP\_CONTROL\_L2\]
DSCP Precedence Level Values Best Effort service will GET any available bandwidth DSCP Class 1
service DSCP Class 2 service DSCP Class 3 service DSCP Class 4 service Express Forwarding is used
for low latency traffic Control is used for routing traffic, not recommended Link Layer traffic
like.. Possible values are \`DSCP\_BEST\_EFFORT\`, \`DSCP\_CLASS1\`, \`DSCP\_CLASS2\`,
\`DSCP\_CLASS3\`, \`DSCP\_CLASS4\`, \`DSCP\_EXPRESS\_FORWARDING\`, \`DSCP\_CONTROL\_L3\`,
\`DSCP\_CONTROL\_L2\`. Defaults to \`DSCP\_BEST\_EFFORT\`.

Upstream description:

DSCP Precedence Level Values

Best Effort service will GET any available bandwidth DSCP Class 1 service DSCP Class 2 service DSCP
Class 3 service DSCP Class 4 service Express Forwarding is used for low latency traffic Control is
used for routing traffic, not recommended Link Layer traffic like LACP or keepalive, not
recommended.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DSCP_BEST_EFFORT",
    "DSCP_CLASS1",
    "DSCP_CLASS2",
    "DSCP_CLASS3",
    "DSCP_CLASS4",
    "DSCP_EXPRESS_FORWARDING",
    "DSCP_CONTROL_L3",
    "DSCP_CONTROL_L2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DSCP_BEST_EFFORT",
  "enum": [
    "DSCP_BEST_EFFORT",
    "DSCP_CLASS1",
    "DSCP_CLASS2",
    "DSCP_CLASS3",
    "DSCP_CLASS4",
    "DSCP_EXPRESS_FORWARDING",
    "DSCP_CONTROL_L3",
    "DSCP_CONTROL_L2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3023312131311012-3010222333322313-3032321033323020-0100322323200020-3301302300223210-3221311310333022-3122122321031011-1210123321031310"></a>

## Next pages — dscp / 320012321330 / 6

- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-0102111112020232-1131131013230311-1223320113131321-1013330221212303-3022123312213323-3323312223011303-2202303002230110-1331212022130030)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)

<a id="canonical-3302113323222301-0133012130232111-0221312101202123-3223011101220123-0303002133022202-0022010122103221-0232201320101121-3001131002330320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233333103320130-0302310020012313-0110220313122213-3111211202002120-2313032010220200-3312033010113003-2110300211331101-3122303211123302"></a>

## dscp_based_queue — dscp_based_queue / 131002223230 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)
- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-0102111112020232-1131131013230311-1223320113131321-1013330221212303-3022123312213323-3323312223011303-2202303002230110-1331212022130030)
- dscp_based_queue

<a id="canonical-2003303032021301-1232211011322121-3213200122000002-3300223200211003-2203003201223312-1032010130230010-2321013131010223-2302220220222212"></a>

Type: `["object", {}]`. Optional.

\[OneOf: dscp\_based\_queue, queue\_id\_to\_use\] Configuration parameter for dscp based queue.

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

OneOf alternatives in this subsection:

- [dscp_based_queue](resources--forwarding_class--reference--group-001.md#canonical-2003303032021301-1232211011322121-3213200122000002-3300223200211003-2203003201223312-1032010130230010-2321013131010223-2302220220222212)
- [queue_id_to_use](resources--forwarding_class--reference--group-001.md#canonical-1322111122031103-3103323233203003-3112230313100031-0100033221020110-2203020000120200-0332200321213112-3323022021320010-3000022023212223)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dscp_based_queue = {}
```

<a id="canonical-3130020120302102-3113003111320333-3121220213031030-1201120331222021-1013033030233012-2211313302110032-0003223102121302-2100321032112212"></a>

## Direct properties — dscp_based_queue / 131002223230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212122211230010-1332122320200003-0101003021300201-2303320200002031-1211033231202321-2300103002233311-0231033121110030-1023101203300201"></a>

## Next pages — dscp_based_queue / 131002223230 / 4

- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-0102111112020232-1131131013230311-1223320113131321-1013330221212303-3022123312213323-3323312223011303-2202303002230110-1331212022130030)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)

<a id="canonical-1021103130221111-1123033322313003-2032001203200330-0001333002012023-3323021332212211-3113320333001230-1101113133200332-1310201221313033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323130323120110-1121021022030223-3013001322200102-2020211323021202-0023011333110322-2103023231303100-3320101313233331-0122203233110223"></a>

## no_marking — no_marking / 020200122002 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)
- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-0102111112020232-1131131013230311-1223320113131321-1013330221212303-3022123312213323-3323312223011303-2202303002230110-1331212022130030)
- no_marking

<a id="canonical-1021021232202130-0331303132133101-0203200111202303-2133300311022332-3113320321332201-1013122030310303-0020001112113302-3011130210330310"></a>

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
no_marking = {}
```

<a id="canonical-2000321300102021-1003321230011111-2200132202031001-3303210001130202-2322120232320113-0330002132021313-1322222123203113-0023230001302312"></a>

## Direct properties — no_marking / 020200122002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213330211320213-1302020132111201-1321200311121001-2000210211222020-0112112010201210-1232330102132012-3122220021200132-0303201223000111"></a>

## Next pages — no_marking / 020200122002 / 4

- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-0102111112020232-1131131013230311-1223320113131321-1013330221212303-3022123312213323-3323312223011303-2202303002230110-1331212022130030)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)

<a id="canonical-1102211202003033-1303000013132220-0201013320301323-2022132212011332-3111131211020213-2203020212003232-3311010001321232-2132301130323132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321031220130110-3111000020103321-1220331312001211-0220203113211012-0002223202113010-2112233201020031-0230023011223302-0330001303010322"></a>

## no_policer — no_policer / 033033001012 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)
- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-0102111112020232-1131131013230311-1223320113131321-1013330221212303-3022123312213323-3323312223011303-2202303002230110-1331212022130030)
- no_policer

<a id="canonical-3022012312002022-1132231322001333-0102031010313100-2013202021201000-0320303323033030-2301131023231121-1012311132123023-1003211331301231"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_policer, policer; Default: no\_policer\] Enable this option

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

OneOf alternatives in this subsection:

- [no_policer](resources--forwarding_class--reference--group-001.md#canonical-3022012312002022-1132231322001333-0102031010313100-2013202021201000-0320303323033030-2301131023231121-1012311132123023-1003211331301231)
- [policer](resources--forwarding_class--reference--group-001.md#canonical-2323331002012330-2003001101221233-0000222212132011-3031011202312322-3020001322030022-3211012120223213-0131111202331112-1102113322213301)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_policer = {}
```

<a id="canonical-2110002110211212-2332201121030220-3330131301133101-1100312032331320-1201021322121021-1312103030101100-1213033213331103-0321103313202222"></a>

## Direct properties — no_policer / 033033001012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001111002313113-2112333110303003-1213232033110123-2321102233332011-2321300223102333-0223312230031123-1101210333320203-1233113303231302"></a>

## Next pages — no_policer / 033033001012 / 4

- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-0102111112020232-1131131013230311-1223320113131321-1013330221212303-3022123312213323-3323312223011303-2202303002230110-1331212022130030)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)

<a id="canonical-1110212130302302-3120310323312010-0221323020122301-1332103131311011-2232123320231231-3301310021300320-2102320131210102-3112321331101001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112330102013000-0312033002222220-1331213021220113-2002010020133011-3001020102002201-0201303211322013-0010111123013130-2202113111232013"></a>

## policer — policer / 123222020131 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)
- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-0102111112020232-1131131013230311-1223320113131321-1013330221212303-3022123312213323-3323312223011303-2202303002230110-1331212022130030)
- policer

<a id="canonical-2323331002012330-2003001101221233-0000222212132011-3031011202312322-3020001322030022-3211012120223213-0131111202331112-1102113322213301"></a>

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
policer {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102003001131102-1001221301303021-0311300310210212-0033321203313002-3211031012331230-0232103230032101-1031131200333100-2102002321312103"></a>

## Direct properties — policer / 123222020131 / 3

<a id="canonical-1311101010330000-0101121122120002-2200111120233303-3023012210123230-3321230112131220-2213012113010313-1030003202021213-2120102103132133"></a>

<a id="canonical-1011013033130123-3311333233112013-2023310021033020-1100212031203232-0123311032033301-2131011021202321-3000013331303213-1012021021023133"></a>

## name property — policer / 123222020131 / 4

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

<a id="canonical-2003202000101000-0230021031122220-1032311010203131-1333121103210200-1232223101303020-0000202132212023-1231031100111012-0120000010320032"></a>

<a id="canonical-3332233213123030-0211100003102212-0130010301113310-3003320221023222-2212012003021110-2303313323022332-3102210021330103-2223102011310121"></a>

## namespace property — policer / 123222020131 / 5

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

<a id="canonical-2030230001310000-0100330021032033-1130023301223220-2301200303030001-3032100211211122-0323203110121212-3031213330201220-3132131200220003"></a>

<a id="canonical-2130320222113203-0002110200200232-2230131200110223-2302103003220213-3323313201103001-2233012233002001-1302233333021133-0133013113332110"></a>

## tenant property — policer / 123222020131 / 6

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

<a id="canonical-1330112022200013-2020330102110120-0203130221033001-2103032301130101-1133111021302210-0022321133210203-2032023320120022-0211210231001122"></a>

## Next pages — policer / 123222020131 / 7

- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-0102111112020232-1131131013230311-1223320113131321-1013330221212303-3022123312213323-3323312223011303-2202303002230110-1331212022130030)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)

<a id="canonical-1232131200303203-0031203210012133-2012313021120110-2230300330031020-2123313132213221-3101202203102203-2301230332323320-3233212010001321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103022302010133-1323331123311030-0110210121103303-1031221033312121-0011213203132003-3331030221300203-0300010011010033-3002301002131311"></a>

## timeouts — timeouts / 333123310033 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)
- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-0102111112020232-1131131013230311-1223320113131321-1013330221212303-3022123312213323-3323312223011303-2202303002230110-1331212022130030)
- timeouts

<a id="canonical-3033312332113202-1200133232200323-2202103220110330-1210002110203302-2102230220210330-2311331200310302-3231233200130110-2313322103302001"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302213323120021-0212132203231233-3213030032303331-2310223331323311-1223002220113103-0220233131100302-0021222200221023-3223302013323212"></a>

## Direct properties — timeouts / 333123310033 / 3

<a id="canonical-1201000011031201-3203220211013220-1113102112300313-0220213032323032-1122113320003132-3011231223301033-2222033020033023-1101132013202130"></a>

<a id="canonical-1032033013303101-1233321233033230-3333303002202101-1221113330010130-2121021213221131-1121310111000203-3111333332020201-0221213101203022"></a>

## create property — timeouts / 333123310033 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3012332323201130-2001213310032021-0113332100021022-3310312322111122-0320011103313323-0232010033210202-2321132013223210-1321313302333111"></a>

<a id="canonical-3011311021013002-2232021211233030-1201321002132331-2333223133223130-1032210202311320-1002322221222000-1210101321311102-0233320202011223"></a>

## delete property — timeouts / 333123310033 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2123322301122111-2112203220302212-0231220133201110-0123030023303000-2002231230013023-2101030310021010-0001110122132130-3130102013122003"></a>

<a id="canonical-0120103112210110-3113333201311333-2312131203203200-1023031321113310-0111100011021231-0221110310013313-1030132221210110-0033312122313202"></a>

## read property — timeouts / 333123310033 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1230011013121103-0201212310202113-3111033332103133-1122301222120223-3013231221230102-2010013032002210-0300001100111113-2021333012231300"></a>

<a id="canonical-2212210220101213-3013133213301001-3322210311312312-1121220122120321-0123100301332111-3323300113320222-0302313121031121-1033133231001003"></a>

## update property — timeouts / 333123310033 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2131301320102322-3113112310232331-1023112102020020-1112330212132112-2112203322231003-2330333300100322-1112312110002222-3112002320322223"></a>

## Next pages — timeouts / 333123310033 / 8

- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-0102111112020232-1131131013230311-1223320113131321-1013330221212303-3022123312213323-3323312223011303-2202303002230110-1331212022130030)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)
