---
page_title: "xcsh_forwarding_class reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_forwarding_class reference."
---

# xcsh_forwarding_class reference

<a id="canonical-1232301301333312-0233001311220232-0313221003110113-2220222123020101-3332223221102212-0221300232013330-0022330020133203-2230200222233213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003302013310103-3311021323131131-3012312100000011-0011032211023001-3000203331030111-1302333331102231-0313323223310230-0210130112213023"></a>

## Property reference — Property reference / 121112130033 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203)
- Property reference

<a id="canonical-2201102312101102-1002130002221131-2203100301333223-3311012332211210-3020121200130213-2331331110313203-2300020312200311-0011100021130003"></a>

## Direct properties — Property reference / 121112130033 / 3

<a id="canonical-0222132112201303-1132122230203100-3201323311320123-0020323132011102-2321002232131323-1302121012211323-0200320103233310-1020131113131130"></a>

<a id="canonical-2132131230111100-2220320330233310-2320203333320001-0021300013333130-0132201333112131-0010112232120111-2210330122120231-0322013133212333"></a>

## annotations property — Property reference / 121112130033 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

<a id="canonical-1033122133003123-0321230112113103-2112001223211300-1200212032203232-0313230301320300-3001110131232213-1311303322112112-0211303031221233"></a>

<a id="canonical-0310202130320233-0203211200210111-2203300200333230-3203332232021122-2122001310123111-1223203022222003-1223230213112132-3220032221211001"></a>

## description property — Property reference / 121112130033 / 5

Type: `"string"`. Computed.

Description of the ForwardingClass.

Upstream description:

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

- [dscp](data-sources--forwarding_class--reference--group-001.md#canonical-2311202013321033-0100101033323120-3111331212200200-3133301132032033-3122320200333113-2201133101021101-2300230120331132-1111122123032022): complete subsection reference.

- [dscp_based_queue](data-sources--forwarding_class--reference--group-001.md#canonical-3000331101023302-0301122121210323-3001002202033121-2023220023320201-0010023030032003-3030033010033210-1323020033133212-3313113321131301): complete subsection reference.

<a id="canonical-0212103112110122-0003223003322023-0322103322003101-3023131021002022-3021102230110302-3030220330011113-0221113132310131-1031231303123120"></a>

<a id="canonical-0310033121002120-2121032203021010-1012031021013013-2002100031023313-2223201113300012-1310320302031133-2112103111032021-1223211113101302"></a>

## ID property — Property reference / 121112130033 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3033320010121200-0122213030021302-2210000100231233-3103031212131120-2002210110111303-1033111303323121-3211013132303010-0332211031323132"></a>

<a id="canonical-1312132322133131-3321213311233133-1030200330122103-3230331312320332-2132210311311122-2213113312303333-0130323313103030-1302312200032131"></a>

## interface_group property — Property reference / 121112130033 / 7

Type: `"string"`. Computed.

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

<a id="canonical-2013101320112200-3332223203010223-3202213122222230-2013103232333231-0301113333231312-2200210332122303-3012132230302030-1233030321102001"></a>

<a id="canonical-2103130023021230-3133313031230123-0111132203121211-0012103231100221-1010322013321313-1311331110312213-0101333002033013-0200123033013232"></a>

## labels property — Property reference / 121112130033 / 8

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-0003302223022223-3310102203213012-1133101100321002-3223222032331200-3212212113110303-0030111123310032-0103003322122002-2121000311223233"></a>

<a id="canonical-2202110310133203-1323031010300330-3202231330202000-2000211122322100-3211302322032301-1023100133031233-0233013133233132-0002220323203201"></a>

## name property — Property reference / 121112130033 / 9

Type: `"string"`. Required.

Name of the ForwardingClass.

Upstream description:

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

<a id="canonical-2203012011032221-3013102330100020-0332123223312223-1322203331333233-0231021121113211-1103130301223122-2332102203011002-2031120231113232"></a>

<a id="canonical-2322221103220210-1231232233023210-3132303030131302-0322332202303212-1200230102332120-2122331001221232-1303202333122123-3210300310002233"></a>

## namespace property — Property reference / 121112130033 / 10

Type: `"string"`. Required.

Namespace where the ForwardingClass exists.

Upstream description:

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

- [no_marking](data-sources--forwarding_class--reference--group-001.md#canonical-2010202022030221-0311030100111131-3223133032133220-1013301000022230-1023310112012131-2111303233022313-1020123303120302-0010222321212332): complete subsection reference.

- [no_policer](data-sources--forwarding_class--reference--group-001.md#canonical-0132002012332231-3230221000012120-0132211221330131-0112303112310211-2132222113213031-0223021031202132-3200211212100232-0203202200211203): complete subsection reference.

- [policer](data-sources--forwarding_class--reference--group-001.md#canonical-1301022313310121-1121313222030230-2303013012221111-1101000332011123-1003331033330001-1230203012023120-0311102312122330-2220031210303013): complete subsection reference.

<a id="canonical-1322120213303313-2233012120002330-0330000012133001-3101321233001333-3220301321132112-3203123033311211-3120102111021201-3011232230202222"></a>

<a id="canonical-2111303012300103-0122133023001213-2302122113111311-2101331301122012-1202203212331312-1000301130221302-3323020220001233-2122331103331130"></a>

## queue_id_to_use property — Property reference / 121112130033 / 11

Type: `"string"`. Computed.

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

<a id="canonical-1201003210323311-0330202033011112-1102032113003221-0031011000031023-1302220320223012-3023300011311311-3321322001032300-1132122301300212"></a>

<a id="canonical-1003230203132323-2222220220222301-1213220032220301-0132010110011132-0110023322123222-0130032232000213-1030120130301332-0213233301220121"></a>

## tos_value property — Property reference / 121112130033 / 12

Type: `"number"`. Computed.

Exclusive with \[dscp no\_marking\] Decimal value of raw 8 bit TOS. In above example DSCP 10 =
Precedence Class 1 and drop precedence low.

Upstream description:

Exclusive with \[dscp no\_marking\] Decimal value of raw 8 bit TOS. In above example DSCP 10 =
Precedence Class 1 and drop precedence low.

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

<a id="canonical-3223013100233331-2033311033321320-1022200300312032-2033010301013311-3030332132013120-0131132312130221-1201210331202210-1123311220133301"></a>

## All schema paths — Property reference / 121112130033 / 13

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--forwarding_class--reference--group-001.md#canonical-0222132112201303-1132122230203100-3201323311320123-0020323132011102-2321002232131323-1302121012211323-0200320103233310-1020131113131130) |
| `description` | [description](data-sources--forwarding_class--reference--group-001.md#canonical-1033122133003123-0321230112113103-2112001223211300-1200212032203232-0313230301320300-3001110131232213-1311303322112112-0211303031221233) |
| `dscp` | [dscp](data-sources--forwarding_class--reference--group-001.md#canonical-0120121133231103-1121123012323230-2001032131012113-2132203311313323-0031231322200200-3202222130221030-1120310033322220-2200333222231003) |
| `dscp.drop_precedence` | [dscp.drop_precedence](data-sources--forwarding_class--reference--group-001.md#canonical-2331313231302113-2312133013011003-1332210303310020-1213022130011001-3012113011322131-2310330212301010-1312132302311023-3233220231300302) |
| `dscp.dscp_class` | [dscp.dscp_class](data-sources--forwarding_class--reference--group-001.md#canonical-1032200123130113-2233020211003020-2113211211332033-2333021312010213-1113000302122120-2022000120132032-3312301231210330-2332031130331231) |
| `dscp_based_queue` | [dscp_based_queue](data-sources--forwarding_class--reference--group-001.md#canonical-0102132033313302-3032002130132013-1000021201232011-2122330310302220-0111302033112220-2031021323320301-0222131102230300-0001010312112003) |
| `id` | [ID](data-sources--forwarding_class--reference--group-001.md#canonical-0212103112110122-0003223003322023-0322103322003101-3023131021002022-3021102230110302-3030220330011113-0221113132310131-1031231303123120) |
| `interface_group` | [interface_group](data-sources--forwarding_class--reference--group-001.md#canonical-3033320010121200-0122213030021302-2210000100231233-3103031212131120-2002210110111303-1033111303323121-3211013132303010-0332211031323132) |
| `labels` | [labels](data-sources--forwarding_class--reference--group-001.md#canonical-2013101320112200-3332223203010223-3202213122222230-2013103232333231-0301113333231312-2200210332122303-3012132230302030-1233030321102001) |
| `name` | [name](data-sources--forwarding_class--reference--group-001.md#canonical-0003302223022223-3310102203213012-1133101100321002-3223222032331200-3212212113110303-0030111123310032-0103003322122002-2121000311223233) |
| `namespace` | [namespace](data-sources--forwarding_class--reference--group-001.md#canonical-2203012011032221-3013102330100020-0332123223312223-1322203331333233-0231021121113211-1103130301223122-2332102203011002-2031120231113232) |
| `no_marking` | [no_marking](data-sources--forwarding_class--reference--group-001.md#canonical-0310311202331030-2321020001123203-1001023131010320-0332012320331330-1102333202233331-3020103103202230-0121220320230103-3230033230132033) |
| `no_policer` | [no_policer](data-sources--forwarding_class--reference--group-001.md#canonical-0011101103230310-0010030301001330-3022320113220123-1013102220312320-0130231100133322-3102023013003211-2330113100003100-1210213131202301) |
| `policer` | [policer](data-sources--forwarding_class--reference--group-001.md#canonical-2110121310020302-2121101221330030-3132330020113203-3113000113302102-2130031101122212-3013201100032020-2030010222132231-1010030003132023) |
| `policer.name` | [policer.name](data-sources--forwarding_class--reference--group-001.md#canonical-2230100113100113-3231311311030012-1201311011231321-0203121132313203-0210123123233112-3120010330300312-0132001231103203-3130321212213221) |
| `policer.namespace` | [policer.namespace](data-sources--forwarding_class--reference--group-001.md#canonical-2100011122203300-3132200123013322-1122132002203122-1032013003212210-3213201302122003-3330002120330113-3111121131320023-2200323313122203) |
| `policer.tenant` | [policer.tenant](data-sources--forwarding_class--reference--group-001.md#canonical-1321211222013231-0102020203110120-0303212031233212-0231313011220212-0320100310020113-3303032212201032-1330000003000200-0131003231222220) |
| `queue_id_to_use` | [queue_id_to_use](data-sources--forwarding_class--reference--group-001.md#canonical-1322120213303313-2233012120002330-0330000012133001-3101321233001333-3220301321132112-3203123033311211-3120102111021201-3011232230202222) |
| `tos_value` | [tos_value](data-sources--forwarding_class--reference--group-001.md#canonical-1201003210323311-0330202033011112-1102032113003221-0031011000031023-1302220320223012-3023300011311311-3321322001032300-1132122301300212) |

<a id="canonical-3212203023010232-0201222313002320-0233312121122230-0313002210313113-0221020113330022-1313212110211332-1213022310203213-2110222033210030"></a>

## Next pages — Property reference / 121112130033 / 14

- [dscp](data-sources--forwarding_class--reference--group-001.md#canonical-2311202013321033-0100101033323120-3111331212200200-3133301132032033-3122320200333113-2201133101021101-2300230120331132-1111122123032022)
- [dscp_based_queue](data-sources--forwarding_class--reference--group-001.md#canonical-3000331101023302-0301122121210323-3001002202033121-2023220023320201-0010023030032003-3030033010033210-1323020033133212-3313113321131301)
- [no_marking](data-sources--forwarding_class--reference--group-001.md#canonical-2010202022030221-0311030100111131-3223133032133220-1013301000022230-1023310112012131-2111303233022313-1020123303120302-0010222321212332)
- [no_policer](data-sources--forwarding_class--reference--group-001.md#canonical-0132002012332231-3230221000012120-0132211221330131-0112303112310211-2132222113213031-0223021031202132-3200211212100232-0203202200211203)
- [policer](data-sources--forwarding_class--reference--group-001.md#canonical-1301022313310121-1121313222030230-2303013012221111-1101000332011123-1003331033330001-1230203012023120-0311102312122330-2220031210303013)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203)

<a id="canonical-2311202013321033-0100101033323120-3111331212200200-3133301132032033-3122320200333113-2201133101021101-2300230120331132-1111122123032022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022323302231001-2303112132332232-1313012332203033-3103011330012231-3310212130201130-0023123213121311-0313220033302313-3303322032001210"></a>

## dscp — dscp / 201303323122 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203)
- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-1232301301333312-0233001311220232-0313221003110113-2220222123020101-3332223221102212-0221300232013330-0022330020133203-2230200222233213)
- dscp

<a id="canonical-0120121133231103-1121123012323230-2001032131012113-2132203311313323-0031231322200200-3202222130221030-1120310033322220-2200333222231003"></a>

Type: `"single"`. Computed.

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

- [dscp](data-sources--forwarding_class--reference--group-001.md#canonical-0120121133231103-1121123012323230-2001032131012113-2132203311313323-0031231322200200-3202222130221030-1120310033322220-2200333222231003)
- [no_marking](data-sources--forwarding_class--reference--group-001.md#canonical-0310311202331030-2321020001123203-1001023131010320-0332012320331330-1102333202233331-3020103103202230-0121220320230103-3230033230132033)
- [tos_value](data-sources--forwarding_class--reference--group-001.md#canonical-1201003210323311-0330202033011112-1102032113003221-0031011000031023-1302220320223012-3023300011311311-3321322001032300-1132122301300212)

Select alternatives according to the provider validators above.

<a id="canonical-0001311333003221-3010310123313301-0211111122113023-3223300220200300-1012332101130030-3231133100031133-3301331112301211-0020311221201301"></a>

## Direct properties — dscp / 201303323122 / 3

<a id="canonical-2331313231302113-2312133013011003-1332210303310020-1213022130011001-3012113011322131-2310330212301010-1312132302311023-3233220231300302"></a>

<a id="canonical-3302121002100333-2223312101032310-0330013330113210-2232131230320022-1232030113122212-3031110133131231-2303232313101332-3021320220313330"></a>

## drop_precedence property — dscp / 201303323122 / 4

Type: `"string"`. Computed.

\[Enum: DSCP\_AF\_LOW|DSCP\_AF\_MEDIUM|DSCP\_AF\_HIGH|DSCP\_AF\_POLICER\] DSCP Assured forwarding
drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP drop
precedence value is taken from output of policer. Possible values are \`DSCP\_AF\_LOW\`,
\`DSCP\_AF\_MEDIUM\`, \`DSCP\_AF\_HIGH\`, \`DSCP\_AF\_POLICER\`.

Upstream description:

DSCP Assured forwarding drop precedence

DSCP Low drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP drop precedence
value is taken from output of policer.

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

<a id="canonical-1032200123130113-2233020211003020-2113211211332033-2333021312010213-1113000302122120-2022000120132032-3312301231210330-2332031130331231"></a>

<a id="canonical-3332330101113201-2302311130320132-1201030311123201-0020101111330232-0321231030032011-2020130022103310-3110103223110303-1231023130222100"></a>

## dscp_class property — dscp / 201303323122 / 5

Type: `"string"`. Computed.

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

<a id="canonical-2023033003100113-2110003221100213-1331001311232103-1001013031232122-0001131302223121-1332013302101103-1100333020212202-2103301032330212"></a>

## Next pages — dscp / 201303323122 / 6

- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-1232301301333312-0233001311220232-0313221003110113-2220222123020101-3332223221102212-0221300232013330-0022330020133203-2230200222233213)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203)

<a id="canonical-3000331101023302-0301122121210323-3001002202033121-2023220023320201-0010023030032003-3030033010033210-1323020033133212-3313113321131301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330012301333300-2000013222323121-0222111221310121-3033312022110301-3220223013333213-0333310232011133-1303333033013223-2333231230303203"></a>

## dscp_based_queue — dscp_based_queue / 230102031203 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203)
- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-1232301301333312-0233001311220232-0313221003110113-2220222123020101-3332223221102212-0221300232013330-0022330020133203-2230200222233213)
- dscp_based_queue

<a id="canonical-0102132033313302-3032002130132013-1000021201232011-2122330310302220-0111302033112220-2031021323320301-0222131102230300-0001010312112003"></a>

Type: `["object", {}]`. Computed.

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

- [dscp_based_queue](data-sources--forwarding_class--reference--group-001.md#canonical-0102132033313302-3032002130132013-1000021201232011-2122330310302220-0111302033112220-2031021323320301-0222131102230300-0001010312112003)
- [queue_id_to_use](data-sources--forwarding_class--reference--group-001.md#canonical-1322120213303313-2233012120002330-0330000012133001-3101321233001333-3220301321132112-3203123033311211-3120102111021201-3011232230202222)

Select alternatives according to the provider validators above.

<a id="canonical-1200002331120222-2222130111233000-1310313013012322-0321011002311130-2221223212333203-3310103322231110-1030303002212112-2132111031011300"></a>

## Direct properties — dscp_based_queue / 230102031203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102331011020312-1103230031300300-1102010020333113-3101113113100320-0220223221322122-1230330000112311-1031202310230231-0330102133120030"></a>

## Next pages — dscp_based_queue / 230102031203 / 4

- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-1232301301333312-0233001311220232-0313221003110113-2220222123020101-3332223221102212-0221300232013330-0022330020133203-2230200222233213)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203)

<a id="canonical-2010202022030221-0311030100111131-3223133032133220-1013301000022230-1023310112012131-2111303233022313-1020123303120302-0010222321212332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023133332113231-2202332003301323-1302320012120302-2110310321132111-3303110222210303-0232121230120203-0000022333213321-1120330111033211"></a>

## no_marking — no_marking / 211000010232 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203)
- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-1232301301333312-0233001311220232-0313221003110113-2220222123020101-3332223221102212-0221300232013330-0022330020133203-2230200222233213)
- no_marking

<a id="canonical-0310311202331030-2321020001123203-1001023131010320-0332012320331330-1102333202233331-3020103103202230-0121220320230103-3230033230132033"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0131010311000202-1220222000031212-2013221213312232-0123032233020230-0031213210220032-1330232330212201-0122111113021001-1313231003303333"></a>

## Direct properties — no_marking / 211000010232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220032111020223-2121300022000110-3212031110321101-0201100113323212-3023120213332232-3312103230030211-3020001210000011-0313332112320232"></a>

## Next pages — no_marking / 211000010232 / 4

- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-1232301301333312-0233001311220232-0313221003110113-2220222123020101-3332223221102212-0221300232013330-0022330020133203-2230200222233213)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203)

<a id="canonical-0132002012332231-3230221000012120-0132211221330131-0112303112310211-2132222113213031-0223021031202132-3200211212100232-0203202200211203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300032210102213-3330330031202010-1132111300113213-2010300310202033-1303033133021213-3121322213021213-2112000102023301-1301233103203102"></a>

## no_policer — no_policer / 313101320101 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203)
- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-1232301301333312-0233001311220232-0313221003110113-2220222123020101-3332223221102212-0221300232013330-0022330020133203-2230200222233213)
- no_policer

<a id="canonical-0011101103230310-0010030301001330-3022320113220123-1013102220312320-0130231100133322-3102023013003211-2330113100003100-1210213131202301"></a>

Type: `["object", {}]`. Computed.

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

- [no_policer](data-sources--forwarding_class--reference--group-001.md#canonical-0011101103230310-0010030301001330-3022320113220123-1013102220312320-0130231100133322-3102023013003211-2330113100003100-1210213131202301)
- [policer](data-sources--forwarding_class--reference--group-001.md#canonical-2110121310020302-2121101221330030-3132330020113203-3113000113302102-2130031101122212-3013201100032020-2030010222132231-1010030003132023)

Select alternatives according to the provider validators above.

<a id="canonical-2101323333012122-1223303300331023-0103221011231020-0300322120111220-0123000212313230-1331212032032003-2223001103221310-3033210220130322"></a>

## Direct properties — no_policer / 313101320101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303032021111223-3232031111322312-3003023202030111-3231133020223003-0101001002023001-0200010031233330-3023213021303202-0301201200021022"></a>

## Next pages — no_policer / 313101320101 / 4

- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-1232301301333312-0233001311220232-0313221003110113-2220222123020101-3332223221102212-0221300232013330-0022330020133203-2230200222233213)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203)

<a id="canonical-1301022313310121-1121313222030230-2303013012221111-1101000332011123-1003331033330001-1230203012023120-0311102312122330-2220031210303013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112312020301230-3310111001312123-2330000323100110-3330011020300200-1132333310102231-1023001213201110-3100210202330310-1332012012220321"></a>

## policer — policer / 113100322212 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203)
- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-1232301301333312-0233001311220232-0313221003110113-2220222123020101-3332223221102212-0221300232013330-0022330020133203-2230200222233213)
- policer

<a id="canonical-2110121310020302-2121101221330030-3132330020113203-3113000113302102-2130031101122212-3013201100032020-2030010222132231-1010030003132023"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0303202003231310-3320300333010312-0020103302001201-1310331203013103-2110111232000120-3113310100221011-0212101033230213-3131101301230312"></a>

## Direct properties — policer / 113100322212 / 3

<a id="canonical-2230100113100113-3231311311030012-1201311011231321-0203121132313203-0210123123233112-3120010330300312-0132001231103203-3130321212213221"></a>

<a id="canonical-2220100000303033-3302330203211001-0310320231012330-1123213320312322-3202130232022013-1130111110120200-1011223003312333-3022001103100022"></a>

## name property — policer / 113100322212 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-2100011122203300-3132200123013322-1122132002203122-1032013003212210-3213201302122003-3330002120330113-3111121131320023-2200323313122203"></a>

<a id="canonical-1212230100003333-3002310100023331-3313332210322303-0210200033133230-3312110032303012-0202221032133311-3322223021333103-2001030213033011"></a>

## namespace property — policer / 113100322212 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-1321211222013231-0102020203110120-0303212031233212-0231313011220212-0320100310020113-3303032212201032-1330000003000200-0131003231222220"></a>

<a id="canonical-0220002302322310-2313220130200331-2200320100003102-1111232323301011-2222302302122132-1002321103213311-0211112103112232-2300010322021011"></a>

## tenant property — policer / 113100322212 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-1322020320221211-3301212300320222-2131102212321133-2203330020011030-3300110133100200-3322020310130311-1221002231303302-0002122332020000"></a>

## Next pages — policer / 113100322212 / 7

- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-1232301301333312-0233001311220232-0313221003110113-2220222123020101-3332223221102212-0221300232013330-0022330020133203-2230200222233213)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203)
