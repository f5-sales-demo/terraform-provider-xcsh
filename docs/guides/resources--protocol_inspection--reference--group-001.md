---
page_title: "xcsh_protocol_inspection reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_inspection reference."
---

# xcsh_protocol_inspection reference

<a id="canonical-3210031210322130-3133210121332000-1300320030213333-1220303203113101-3030330132202312-1103302221030102-1202010201212331-3030222101313220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-3030212022213332-3331201033021331-2130011232020201-3310322223233130-3212232100333301-3201232230312101-1003032323133112-1133013223230102)
- Property reference

<a id="canonical-3113120113101101-1223101211021333-3200133132300030-0103331022212333-2023013333033303-1120120031113010-3000020303002210-3233100001130222"></a>

### Direct properties for `xcsh_protocol_inspection`

<a id="canonical-3123320111023202-2100202111033233-2111310200223222-3212000212012020-0131332212132320-2021303110131101-1022122010021021-2113022210330131"></a>

#### `action` property

Type: `"string"`. Optional, Computed.

\[Enum: ALLOW|DENY|DROP\] Action after inspection - ALLOW: Allow Allow traffic - DENY: Deny Throw
RST error for TCP and ICMP error for UDP - DROP: DROP Silently drop traffic. Possible values are
\`ALLOW\`, \`DENY\`, \`DROP\`. Defaults to \`ALLOW\`. Server applies default when omitted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALLOW","DENY","DROP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ALLOW",
    "DENY",
    "DROP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ALLOW",
  "enum": [
    "ALLOW",
    "DENY",
    "DROP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2333203021223222-2223101303213231-0110122200001211-3301102022221113-1223221322001303-1110333313312202-2200311330230102-2111111332121012"></a>

<a id="canonical-2220231111333331-1222212302033032-2321232200000111-3022021033320230-1320303133210321-0013121100132211-2112311132022112-0003332022033202"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

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

<a id="canonical-2210002220031322-3330301022021001-3220303011020013-0312031013221112-2300002330330203-2231113331300101-2120203002232030-2330212001013310"></a>

<a id="canonical-3112302002332020-0313231000221111-1111311011230123-2233112221211002-2323000020003023-2101302323303013-1101232311133103-2220333130110113"></a>

#### `description` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3203013211032323-1021021013033022-3330210031220222-3323010221102320-1310010013123233-3303023233202201-1221330013301210-3303033030000322"></a>

<a id="canonical-0331332000202133-1122333123333302-3232020222003011-3332302123300303-0232302312130211-2031001331223331-2000003202201033-3103330121100103"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

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

- [enable_disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-3032303120120322-2021100000002111-2310220332010231-1233101110001211-1130323103302010-2213213331133112-1031321210203310-1002133331133202): complete subsection reference.

- [enable_disable_signatures](resources--protocol_inspection--reference--group-001.md#canonical-1103202120331001-3021223223230333-3311111113320300-0113233330323001-1310220331310130-3111230130303221-3033121302233302-0220122231222233): complete subsection reference.

<a id="canonical-0122331211333331-1003332201320300-2112313230223331-2031300211210200-3221321333303003-2030220112312011-2013330300132311-2000323030310232"></a>

<a id="canonical-2331112203000002-1131103123301001-2100020332101320-3313212132203010-3020003322120211-3133212300301331-3333231121112011-2003202003200202"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2232010210021032-1210102110013222-2011212210323002-3221100012213020-3102120223200012-1201012033000331-3233133000321132-2011303202130222"></a>

<a id="canonical-2032132031030322-3302000213211320-1220223022003213-1112212232213123-2220203103003303-3003301122301113-2003212312300220-0002321000212130"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-2220223312200230-1012113210313112-1233322021201303-3212213212120322-1310323201332101-1121130133311123-0112022013020121-1100302130003333"></a>

<a id="canonical-2031012311000322-0212303232231003-2023321330232333-2033000130021300-2101333211112201-2201002332200210-0300210131121023-0203223133011233"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Protocol Inspection. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0123212311310022-2312110233322010-2110331330000222-0301002221301312-3030203002212322-1123320321132012-3000320122330010-0020033100012123"></a>

<a id="canonical-0310321032210131-2031213302203033-1220130113310113-2023130322113113-3200332113231320-1023101222111033-3300002102132300-3031020320211113"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Protocol Inspection is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [timeouts](resources--protocol_inspection--reference--group-001.md#canonical-2210000312331211-1211032233233031-2003133032032133-2100302021230303-2103131332223001-0230103133220001-1031222033020033-0232310130220112): complete subsection reference.

<a id="canonical-2020311020332222-2323010000331131-2333221013222000-1201321203201012-2330102023223011-3101321011022002-0122133200100302-3132313010333231"></a>

### All schema paths for `xcsh_protocol_inspection`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](resources--protocol_inspection--reference--group-001.md#canonical-3123320111023202-2100202111033233-2111310200223222-3212000212012020-0131332212132320-2021303110131101-1022122010021021-2113022210330131) |
| `annotations` | [annotations](resources--protocol_inspection--reference--group-001.md#canonical-2333203021223222-2223101303213231-0110122200001211-3301102022221113-1223221322001303-1110333313312202-2200311330230102-2111111332121012) |
| `description` | [description](resources--protocol_inspection--reference--group-001.md#canonical-2210002220031322-3330301022021001-3220303011020013-0312031013221112-2300002330330203-2231113331300101-2120203002232030-2330212001013310) |
| `disable` | [disable](resources--protocol_inspection--reference--group-001.md#canonical-3203013211032323-1021021013033022-3330210031220222-3323010221102320-1310010013123233-3303023233202201-1221330013301210-3303033030000322) |
| `enable_disable_compliance_checks` | [enable_disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-2103000103231101-2333100031300000-1103312002103120-0102111003212231-1110102133100203-0112100101220110-0233111201313223-1333313112202230) |
| `enable_disable_compliance_checks.disable_compliance_checks` | [enable_disable_compliance_checks.disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-3110010022110230-1332323303021332-1022001300132323-3300011002133033-1013320333103023-1223003011231232-3133003111211131-0303333331231332) |
| `enable_disable_compliance_checks.enable_compliance_checks` | [enable_disable_compliance_checks.enable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-3012121000231002-2111332211212200-0312103031112210-2223121230333030-2022130333313332-1023332111101012-3021220021002031-3311210113031210) |
| `enable_disable_compliance_checks.enable_compliance_checks.name` | [enable_disable_compliance_checks.enable_compliance_checks.name](resources--protocol_inspection--reference--group-001.md#canonical-3333022202000001-0121023201223310-1121301123033311-3320310003131030-3302300313331113-3202223212202010-1100230333022103-0210122130131210) |
| `enable_disable_compliance_checks.enable_compliance_checks.namespace` | [enable_disable_compliance_checks.enable_compliance_checks.namespace](resources--protocol_inspection--reference--group-001.md#canonical-2132311313302000-0020232221121122-3332011101210311-1230203130230332-3010302232332312-1133012200220032-0312220200032012-1213310233032003) |
| `enable_disable_compliance_checks.enable_compliance_checks.tenant` | [enable_disable_compliance_checks.enable_compliance_checks.tenant](resources--protocol_inspection--reference--group-001.md#canonical-1223131121032332-3032303333330213-2133311311311122-0122110132002113-0133132310023003-1302332133330223-0213100233203330-2013103011100123) |
| `enable_disable_signatures` | [enable_disable_signatures](resources--protocol_inspection--reference--group-001.md#canonical-2012230130120310-3102121330002321-2303300303013003-0212302113013333-3331201020321311-0200001331323223-1330210023330022-1310201103321230) |
| `enable_disable_signatures.disable_signature` | [enable_disable_signatures.disable_signature](resources--protocol_inspection--reference--group-001.md#canonical-1103130201100200-1032230112000013-0020312032320120-0232010033203011-3222131033333223-1331130321122200-1222120331133331-0120123201231011) |
| `enable_disable_signatures.enable_signature` | [enable_disable_signatures.enable_signature](resources--protocol_inspection--reference--group-001.md#canonical-0122310232330322-0220333130323230-2030110331123001-0213302203111302-1110100301310112-2312021311310010-0132101020121232-2230200101000123) |
| `id` | [ID](resources--protocol_inspection--reference--group-001.md#canonical-0122331211333331-1003332201320300-2112313230223331-2031300211210200-3221321333303003-2030220112312011-2013330300132311-2000323030310232) |
| `labels` | [labels](resources--protocol_inspection--reference--group-001.md#canonical-2232010210021032-1210102110013222-2011212210323002-3221100012213020-3102120223200012-1201012033000331-3233133000321132-2011303202130222) |
| `name` | [name](resources--protocol_inspection--reference--group-001.md#canonical-2220223312200230-1012113210313112-1233322021201303-3212213212120322-1310323201332101-1121130133311123-0112022013020121-1100302130003333) |
| `namespace` | [namespace](resources--protocol_inspection--reference--group-001.md#canonical-0123212311310022-2312110233322010-2110331330000222-0301002221301312-3030203002212322-1123320321132012-3000320122330010-0020033100012123) |
| `timeouts` | [timeouts](resources--protocol_inspection--reference--group-001.md#canonical-2013211310211331-3100000313213132-2203132300021121-1100202221303333-2003321103330210-1303321131111031-2322001013110213-1012323000110010) |
| `timeouts.create` | [timeouts.create](resources--protocol_inspection--reference--group-001.md#canonical-3222312203200031-2012323020100000-1320120323321003-0003132003233013-2013101100112102-0331110201133233-3020003002332222-3133213303222022) |
| `timeouts.delete` | [timeouts.delete](resources--protocol_inspection--reference--group-001.md#canonical-0311030201323003-1333031012233022-0313313300001223-2300113023332102-1233210000003112-1201223221210133-0232203211122322-2213121132233302) |
| `timeouts.read` | [timeouts.read](resources--protocol_inspection--reference--group-001.md#canonical-1031313130233223-2330011330000313-1333223020010322-1001212311032101-0020033110202211-3122223002030301-0210333323332002-1220303213222122) |
| `timeouts.update` | [timeouts.update](resources--protocol_inspection--reference--group-001.md#canonical-1123111333003033-1222102320003110-2312032122322021-1330313102110003-3212122120213111-1331101123121333-3022210302032010-1302201203201033) |

<a id="canonical-3032303120120322-2021100000002111-2310220332010231-1233101110001211-1130323103302010-2213213331133112-1031321210203310-1002133331133202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_disable_compliance_checks` properties

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-3030212022213332-3331201033021331-2130011232020201-3310322223233130-3212232100333301-3201232230312101-1003032323133112-1133013223230102)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-3210031210322130-3133210121332000-1300320030213333-1220303203113101-3030330132202312-1103302221030102-1202010201212331-3030222101313220)
- enable_disable_compliance_checks

<a id="canonical-2103000103231101-2333100031300000-1103312002103120-0102111003212231-1110102133100203-0112100101220110-0233111201313223-1333313112202230"></a>

Type: `"object"`. single nested block, Optional.

Enable Disable Compliance Checks Choice.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_compliance_checks",
    "enable_compliance_checks")}
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
  "x-ves-oneof-field-compliance_check_choice": "[\"disable_compliance_checks\",\"enable_compliance_checks\"]"
}
```

Terraform syntax:

```terraform
enable_disable_compliance_checks {
  # Configure direct properties listed below.
}
```

<a id="canonical-1012103030322001-1112203231223111-3113032222111010-3113021303220332-2331311113320211-3022333133103010-2213012132221200-0311231203330023"></a>

### Direct properties for `enable_disable_compliance_checks`

- [disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-2002213210222313-1301203032330232-1130002101020131-1031220220200202-0211001113021203-2132333320232031-1300203131011000-1120103321023000): complete subsection reference.

- [enable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-1011012310000131-0213123130311333-2022231221201032-1133330213323231-0303220213023201-2320323302110303-1000202032201301-2002110303231311): complete subsection reference.

<a id="canonical-2002213210222313-1301203032330232-1130002101020131-1031220220200202-0211001113021203-2132333320232031-1300203131011000-1120103321023000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_disable_compliance_checks.disable_compliance_checks` properties

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-3030212022213332-3331201033021331-2130011232020201-3310322223233130-3212232100333301-3201232230312101-1003032323133112-1133013223230102)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-3210031210322130-3133210121332000-1300320030213333-1220303203113101-3030330132202312-1103302221030102-1202010201212331-3030222101313220)
- [enable_disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-3032303120120322-2021100000002111-2310220332010231-1233101110001211-1130323103302010-2213213331133112-1031321210203310-1002133331133202)
- enable_disable_compliance_checks.disable_compliance_checks

<a id="canonical-3110010022110230-1332323303021332-1022001300132323-3300011002133033-1013320333103023-1223003011231232-3133003111211131-0303333331231332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable compliance checks.

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
disable_compliance_checks = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011012310000131-0213123130311333-2022231221201032-1133330213323231-0303220213023201-2320323302110303-1000202032201301-2002110303231311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_disable_compliance_checks.enable_compliance_checks` properties

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-3030212022213332-3331201033021331-2130011232020201-3310322223233130-3212232100333301-3201232230312101-1003032323133112-1133013223230102)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-3210031210322130-3133210121332000-1300320030213333-1220303203113101-3030330132202312-1103302221030102-1202010201212331-3030222101313220)
- [enable_disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-3032303120120322-2021100000002111-2310220332010231-1233101110001211-1130323103302010-2213213331133112-1031321210203310-1002133331133202)
- enable_disable_compliance_checks.enable_compliance_checks

<a id="canonical-3012121000231002-2111332211212200-0312103031112210-2223121230333030-2022130333313332-1023332111101012-3021220021002031-3311210113031210"></a>

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
enable_compliance_checks {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102213033321330-0122333130010132-2112010323331020-1112320332222112-2231213233001230-2300200131033203-1100302221233133-0011323203220133"></a>

### Direct properties for `enable_disable_compliance_checks.enable_compliance_checks`

<a id="canonical-3333022202000001-0121023201223310-1121301123033311-3320310003131030-3302300313331113-3202223212202010-1100230333022103-0210122130131210"></a>

#### `enable_disable_compliance_checks.enable_compliance_checks.name` property

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

<a id="canonical-2132311313302000-0020232221121122-3332011101210311-1230203130230332-3010302232332312-1133012200220032-0312220200032012-1213310233032003"></a>

<a id="canonical-3111232133021133-2020013000101301-1001223331303111-1312023320103021-0032110222322021-0023223301112211-2130300030313100-1110032320320223"></a>

#### `enable_disable_compliance_checks.enable_compliance_checks.namespace` property

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

<a id="canonical-1223131121032332-3032303333330213-2133311311311122-0122110132002113-0133132310023003-1302332133330223-0213100233203330-2013103011100123"></a>

<a id="canonical-1011231221001000-0212031021031012-2000230223033032-3321110332010111-0132310320023211-2302310001310212-1231312103031330-1133112320000210"></a>

#### `enable_disable_compliance_checks.enable_compliance_checks.tenant` property

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

<a id="canonical-1103202120331001-3021223223230333-3311111113320300-0113233330323001-1310220331310130-3111230130303221-3033121302233302-0220122231222233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_disable_signatures` properties

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-3030212022213332-3331201033021331-2130011232020201-3310322223233130-3212232100333301-3201232230312101-1003032323133112-1133013223230102)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-3210031210322130-3133210121332000-1300320030213333-1220303203113101-3030330132202312-1103302221030102-1202010201212331-3030222101313220)
- enable_disable_signatures

<a id="canonical-2012230130120310-3102121330002321-2303300303013003-0212302113013333-3331201020321311-0200001331323223-1330210023330022-1310201103321230"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enable disable signatures.

Additional upstream details:

Enable Disable Signature Choice.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_signature",
    "enable_signature")}
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
  "x-ves-oneof-field-signature_choice": "[\"disable_signature\",\"enable_signature\"]"
}
```

Terraform syntax:

```terraform
enable_disable_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133323002210303-2203300230332333-1023122032312232-1120131310233001-1222202301322311-3233102211230102-0033001331123233-2233213313100110"></a>

### Direct properties for `enable_disable_signatures`

- [disable_signature](resources--protocol_inspection--reference--group-001.md#canonical-3330133210221230-3222233212003302-1012112101212302-2203233110220020-0211212000001121-3002003303032113-1122010212123113-0113301221233212): complete subsection reference.

- [enable_signature](resources--protocol_inspection--reference--group-001.md#canonical-1221131301013012-2020011133113023-3210333120032231-3220033201303003-3200323310232103-1332113111022201-1310003310020232-3220300012113303): complete subsection reference.

<a id="canonical-3330133210221230-3222233212003302-1012112101212302-2203233110220020-0211212000001121-3002003303032113-1122010212123113-0113301221233212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_disable_signatures.disable_signature` properties

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-3030212022213332-3331201033021331-2130011232020201-3310322223233130-3212232100333301-3201232230312101-1003032323133112-1133013223230102)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-3210031210322130-3133210121332000-1300320030213333-1220303203113101-3030330132202312-1103302221030102-1202010201212331-3030222101313220)
- [enable_disable_signatures](resources--protocol_inspection--reference--group-001.md#canonical-1103202120331001-3021223223230333-3311111113320300-0113233330323001-1310220331310130-3111230130303221-3033121302233302-0220122231222233)
- enable_disable_signatures.disable_signature

<a id="canonical-1103130201100200-1032230112000013-0020312032320120-0232010033203011-3222131033333223-1331130321122200-1222120331133331-0120123201231011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable signature.

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
disable_signature = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221131301013012-2020011133113023-3210333120032231-3220033201303003-3200323310232103-1332113111022201-1310003310020232-3220300012113303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_disable_signatures.enable_signature` properties

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-3030212022213332-3331201033021331-2130011232020201-3310322223233130-3212232100333301-3201232230312101-1003032323133112-1133013223230102)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-3210031210322130-3133210121332000-1300320030213333-1220303203113101-3030330132202312-1103302221030102-1202010201212331-3030222101313220)
- [enable_disable_signatures](resources--protocol_inspection--reference--group-001.md#canonical-1103202120331001-3021223223230333-3311111113320300-0113233330323001-1310220331310130-3111230130303221-3033121302233302-0220122231222233)
- enable_disable_signatures.enable_signature

<a id="canonical-0122310232330322-0220333130323230-2030110331123001-0213302203111302-1110100301310112-2312021311310010-0132101020121232-2230200101000123"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable signature.

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
enable_signature = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210000312331211-1211032233233031-2003133032032133-2100302021230303-2103131332223001-0230103133220001-1031222033020033-0232310130220112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-3030212022213332-3331201033021331-2130011232020201-3310322223233130-3212232100333301-3201232230312101-1003032323133112-1133013223230102)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-3210031210322130-3133210121332000-1300320030213333-1220303203113101-3030330132202312-1103302221030102-1202010201212331-3030222101313220)
- timeouts

<a id="canonical-2013211310211331-3100000313213132-2203132300021121-1100202221303333-2003321103330210-1303321131111031-2322001013110213-1012323000110010"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302113021210202-1112111203110331-2031203220200033-1013330111311303-3021131222012102-0002010110231113-3031323000102223-3221232013030220"></a>

### Direct properties for `timeouts`

<a id="canonical-3222312203200031-2012323020100000-1320120323321003-0003132003233013-2013101100112102-0331110201133233-3020003002332222-3133213303222022"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0311030201323003-1333031012233022-0313313300001223-2300113023332102-1233210000003112-1201223221210133-0232203211122322-2213121132233302"></a>

<a id="canonical-0002120233310312-3313201332020310-3201212311013333-1020320033200123-1001302032320323-3233212313311122-1131100020222302-3112133312110123"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1031313130233223-2330011330000313-1333223020010322-1001212311032101-0020033110202211-3122223002030301-0210333323332002-1220303213222122"></a>

<a id="canonical-0033220233010033-1111111122332303-3023310233103223-2113021311113113-3231213310223230-2122110112103021-3200333212000131-2132120312203333"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1123111333003033-1222102320003110-2312032122322021-1330313102110003-3212122120213111-1331101123121333-3022210302032010-1302201203201033"></a>

<a id="canonical-0300223310003333-1100310000130123-0001120213311122-0102301331232110-2330203301303023-1001122220221033-3020232330012010-1120023221033123"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
