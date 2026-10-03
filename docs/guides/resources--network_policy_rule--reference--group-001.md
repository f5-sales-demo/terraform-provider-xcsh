---
page_title: "xcsh_network_policy_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_rule reference."
---

# xcsh_network_policy_rule reference

<a id="canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013210120322222-3112221333100023-1011200131133233-1232020000322233-3000120220322010-2212002210011103-2230221221333012-0132001221122213"></a>

## Property reference — Property reference / 011303130013 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)
- Property reference

<a id="canonical-0313210310023223-0021221120032323-2010133201022120-1312210223110230-2112321010310110-0100322101203011-2101012320100201-3111233020330312"></a>

## Direct properties — Property reference / 011303130013 / 3

<a id="canonical-0211133210103031-2301000101101221-1233130031301230-0001103213102201-1323103223222220-1310200321210210-0100023312230320-3313102222313122"></a>

<a id="canonical-3122231222120120-0133023221301012-3323323023033123-2130133122230232-1003330000303130-0000311202012221-2121230031133212-0012013011223130"></a>

## action property — Property reference / 011303130013 / 4

Type: `"string"`. Optional, Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [advanced_action](resources--network_policy_rule--reference--group-001.md#canonical-2201003201203002-1001021032313331-0000232203010002-0000231003020102-3001203210231122-3222031020113232-2022101230331123-3232310100313020): complete subsection reference.

<a id="canonical-1313303013030122-2121201321111113-1213323022002322-0202113111211310-1233332221302232-0321123121321111-2310221233301203-3323330200001231"></a>

<a id="canonical-1002310223331102-2130323002211302-3023120121101031-1332103103332122-1001321133030313-0032323300033302-3033323031333312-3131231332030022"></a>

## annotations property — Property reference / 011303130013 / 5

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

<a id="canonical-2232131011010311-3211110212301112-1310121103203311-0300012121102222-2011200021300222-1123030100102221-2120011203231030-2101301130220231"></a>

<a id="canonical-3230110321113333-3030302220132232-2233103102030333-1330212020123022-3333222220122021-0211023223233223-2333322113132013-1200133322122200"></a>

## description property — Property reference / 011303130013 / 6

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

<a id="canonical-3011303221123322-2012020110130020-3211233330113133-3023321022020021-0203331301313200-2120320231021331-1130220200232012-3222002333123331"></a>

<a id="canonical-0302001330102332-0000033110330101-3303210111132023-1131021331203123-2201103303202001-2030231232030132-0303132332100023-0221013200200130"></a>

## disable property — Property reference / 011303130013 / 7

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

<a id="canonical-0100000103213012-2002102033230302-1300032233213033-0321213133002332-1013220003111012-3213230332122003-0020002103031031-2330302231112130"></a>

<a id="canonical-2202210023020332-2211001001312130-0020001322213211-1030020202013303-2022211213132203-1302033203003200-0110332003313230-0102022231133202"></a>

## ID property — Property reference / 011303130013 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_prefix_set](resources--network_policy_rule--reference--group-001.md#canonical-1330023003330311-0120320212023000-3011113312022212-3321212113003322-1322030310230311-0020213122332002-2032202202130303-0101011123223132): complete subsection reference.

- [label_matcher](resources--network_policy_rule--reference--group-001.md#canonical-3222211312301230-0213003220131133-1103113301100322-1312010120330201-0032020133120322-2230013122010001-3003031332303220-0010012123030031): complete subsection reference.

<a id="canonical-2000320203132220-2320033123320232-2331223120022331-0233032101333030-2213023121013310-0221012213310120-3030122200223332-0031133110202032"></a>

<a id="canonical-0311000331331311-3032303311300010-2120312322210000-2022313101032113-1023312220022123-3102101031030022-2231032203311111-0310121112102220"></a>

## labels property — Property reference / 011303130013 / 9

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

<a id="canonical-1123120232232212-1322100213021201-2003300200210203-0310111331132022-1333210121000202-0213013121031121-0332010300010101-0300023222001313"></a>

<a id="canonical-1332001031302303-0123022321222031-3201212113111221-0333111011022300-0113230121002102-0231300312000131-0210212333230120-1232313232022012"></a>

## name property — Property reference / 011303130013 / 10

Type: `"string"`. Required.

Name of the Network Policy Rule. Must be unique within the namespace.

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

<a id="canonical-1130112013211212-2131031210130212-3000210331201013-3023320003012112-2202111002211110-2103231130321000-1121121222310000-3030223010002223"></a>

<a id="canonical-3100113323212013-1211321331132203-0003301201100130-2031320103220131-1112311131320020-3233030121320320-3211003110323111-2021210003110022"></a>

## namespace property — Property reference / 011303130013 / 11

Type: `"string"`. Required.

Namespace where the Network Policy Rule is created.

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

<a id="canonical-0213030301112230-2312222231301233-0132020022221322-2010102111302111-2122020001212230-0313313000303120-2323323113313030-0220230222202311"></a>

<a id="canonical-0301010211120311-1203003331012032-3320321012311302-2202202012132200-2222103321000302-3203300201103012-1021130202230030-1032031032233001"></a>

## ports property — Property reference / 011303130013 / 12

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [prefix](resources--network_policy_rule--reference--group-001.md#canonical-1322332132013232-0332313113100120-1122031211122021-3202302023323010-0032030221021030-1300222000022322-1223313012310311-1320233203031322): complete subsection reference.

- [prefix_selector](resources--network_policy_rule--reference--group-001.md#canonical-0123133033332220-1332302021201021-1311031103013223-0132030013110331-1012010100011001-3211212222100321-3201102203021112-3332122012022222): complete subsection reference.

<a id="canonical-1213120110121332-1131213321000203-1311220311203101-0312133100300301-2301231113203330-0332220102212333-2231330022321131-2220011103200132"></a>

<a id="canonical-2021130321112001-2303302323310231-1112221231003211-0203133220021302-1023133101012212-0312101311030011-0231312303312203-2001203313310000"></a>

## protocol property — Property reference / 011303130013 / 13

Type: `"string"`. Optional, Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ALL",
    "TCP",
    "UDP",
    "ICMP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

- [timeouts](resources--network_policy_rule--reference--group-001.md#canonical-1300032020203311-0001033032221233-3321113100111211-3032113200300222-2300201033230220-2033011300322212-3322021131132110-2301210013002332): complete subsection reference.

<a id="canonical-0003031033101222-2310212210321302-1000112000122303-1301333203031221-1210132300311301-0003333222311002-0231330131223010-2002031202311321"></a>

## All schema paths — Property reference / 011303130013 / 14

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](resources--network_policy_rule--reference--group-001.md#canonical-0211133210103031-2301000101101221-1233130031301230-0001103213102201-1323103223222220-1310200321210210-0100023312230320-3313102222313122) |
| `advanced_action` | [advanced_action](resources--network_policy_rule--reference--group-001.md#canonical-1213112312010113-1233020111010301-3030313013013320-2102001033331320-2132233013302122-0221131132132130-2021011030321322-3323220303222011) |
| `advanced_action.action` | [advanced_action.action](resources--network_policy_rule--reference--group-001.md#canonical-2030220030232120-1031320233031200-1312100301231313-3031113123122331-1230212213211132-3130303233103033-1012011021023122-1003131130010211) |
| `annotations` | [annotations](resources--network_policy_rule--reference--group-001.md#canonical-1313303013030122-2121201321111113-1213323022002322-0202113111211310-1233332221302232-0321123121321111-2310221233301203-3323330200001231) |
| `description` | [description](resources--network_policy_rule--reference--group-001.md#canonical-2232131011010311-3211110212301112-1310121103203311-0300012121102222-2011200021300222-1123030100102221-2120011203231030-2101301130220231) |
| `disable` | [disable](resources--network_policy_rule--reference--group-001.md#canonical-3011303221123322-2012020110130020-3211233330113133-3023321022020021-0203331301313200-2120320231021331-1130220200232012-3222002333123331) |
| `id` | [ID](resources--network_policy_rule--reference--group-001.md#canonical-0100000103213012-2002102033230302-1300032233213033-0321213133002332-1013220003111012-3213230332122003-0020002103031031-2330302231112130) |
| `ip_prefix_set` | [ip_prefix_set](resources--network_policy_rule--reference--group-001.md#canonical-1220331221100333-2122122320021210-2310121030231032-1213033300131213-2312033301220001-3001332021210202-0321001120111132-1031121322003120) |
| `ip_prefix_set.ref` | [ip_prefix_set.ref](resources--network_policy_rule--reference--group-001.md#canonical-1110113200012112-3132132313130022-1000223111212130-3202222030113120-1310201122222121-2323100310331103-1221000212320210-0010322102023301) |
| `ip_prefix_set.ref.kind` | [ip_prefix_set.ref.kind](resources--network_policy_rule--reference--group-001.md#canonical-3212322113100221-1031332201223113-2210132001220101-3311313031030122-0101002130221122-2301110323300013-1113332122032221-2012123330111231) |
| `ip_prefix_set.ref.name` | [ip_prefix_set.ref.name](resources--network_policy_rule--reference--group-001.md#canonical-3131132123223120-3131131310332002-0221112311133232-0313222230220130-2311222033132220-3200222303133302-3132023223321222-2232122032133002) |
| `ip_prefix_set.ref.namespace` | [ip_prefix_set.ref.namespace](resources--network_policy_rule--reference--group-001.md#canonical-0031211213102321-0020112222031211-1200231323010321-0021012031322233-0131112023020223-0020311132000032-3300121201022032-3003203001032111) |
| `ip_prefix_set.ref.tenant` | [ip_prefix_set.ref.tenant](resources--network_policy_rule--reference--group-001.md#canonical-2103330322111012-0122033031303221-0312101133002231-1333220131033331-1223303233330102-2320013033101003-0133222001000323-1020020330023111) |
| `ip_prefix_set.ref.uid` | [ip_prefix_set.ref.uid](resources--network_policy_rule--reference--group-001.md#canonical-2023201223223123-2133202211133133-2133122030312101-1001022022332310-2232121101131313-3102113020122103-2113001013310023-2132120033202113) |
| `label_matcher` | [label_matcher](resources--network_policy_rule--reference--group-001.md#canonical-3333112010110111-3201300112222012-0303003021312311-1232100102310300-3233022003130133-2320131131021322-0211233221200101-3210210300031302) |
| `label_matcher.keys` | [label_matcher.keys](resources--network_policy_rule--reference--group-001.md#canonical-0112310131100032-2223020213302201-2210023111023100-3202032232322003-1110022212233222-0130031010201122-2101102023332200-2023221232200132) |
| `labels` | [labels](resources--network_policy_rule--reference--group-001.md#canonical-2000320203132220-2320033123320232-2331223120022331-0233032101333030-2213023121013310-0221012213310120-3030122200223332-0031133110202032) |
| `name` | [name](resources--network_policy_rule--reference--group-001.md#canonical-1123120232232212-1322100213021201-2003300200210203-0310111331132022-1333210121000202-0213013121031121-0332010300010101-0300023222001313) |
| `namespace` | [namespace](resources--network_policy_rule--reference--group-001.md#canonical-1130112013211212-2131031210130212-3000210331201013-3023320003012112-2202111002211110-2103231130321000-1121121222310000-3030223010002223) |
| `ports` | [ports](resources--network_policy_rule--reference--group-001.md#canonical-0213030301112230-2312222231301233-0132020022221322-2010102111302111-2122020001212230-0313313000303120-2323323113313030-0220230222202311) |
| `prefix` | [prefix](resources--network_policy_rule--reference--group-001.md#canonical-0130030202032220-0213030001022100-2121322330220121-2030211210232323-3132101023031003-2310210111201331-3330121133020213-1333133032101103) |
| `prefix.prefix` | [prefix.prefix](resources--network_policy_rule--reference--group-001.md#canonical-0022310031113230-3321130203212013-0021320300110031-0232301101231332-0330301122130213-3313003100320101-0102020133131222-0003333112100301) |
| `prefix_selector` | [prefix_selector](resources--network_policy_rule--reference--group-001.md#canonical-3102211033013012-2230100222330123-0200112112031212-1033010111122221-2103231202000323-3012130232330310-0021002230022233-0032202300130333) |
| `prefix_selector.expressions` | [prefix_selector.expressions](resources--network_policy_rule--reference--group-001.md#canonical-1303233210032333-2310011001220333-0303201113203001-1201113021113313-2331302311321202-0330322030212112-3202331320013132-3330223303012103) |
| `protocol` | [protocol](resources--network_policy_rule--reference--group-001.md#canonical-1213120110121332-1131213321000203-1311220311203101-0312133100300301-2301231113203330-0332220102212333-2231330022321131-2220011103200132) |
| `timeouts` | [timeouts](resources--network_policy_rule--reference--group-001.md#canonical-2032212123200333-1311002210201033-3013013303211011-2232131123330331-1302212120123000-2121230001333300-3322030000223032-0232121012110020) |
| `timeouts.create` | [timeouts.create](resources--network_policy_rule--reference--group-001.md#canonical-0222333301321122-1302122332312200-0110023322030322-1322212320322200-2213130002220330-1230303012102330-1233122122323123-2311300002201302) |
| `timeouts.delete` | [timeouts.delete](resources--network_policy_rule--reference--group-001.md#canonical-2010033113300000-3231200133010220-2201102313312002-1212313111201011-3313111030222131-1313011222331203-2332330231212220-3120011202230032) |
| `timeouts.read` | [timeouts.read](resources--network_policy_rule--reference--group-001.md#canonical-1321211100101102-1323331110213301-0132023101201002-0001010132101321-2200232310131320-2333331022322111-3333001332001121-3203220232010010) |
| `timeouts.update` | [timeouts.update](resources--network_policy_rule--reference--group-001.md#canonical-1211323320133201-3031321113023101-2130233330312322-3311200000301222-2121211000203010-1022020110012231-3323200313132230-2021031321121321) |

<a id="canonical-0322211231111000-3322101310132213-0031221321212213-1232310230311313-1012201103111212-3210323221302312-0030030030211330-3112213310112303"></a>

## Next pages — Property reference / 011303130013 / 15

- [advanced_action](resources--network_policy_rule--reference--group-001.md#canonical-2201003201203002-1001021032313331-0000232203010002-0000231003020102-3001203210231122-3222031020113232-2022101230331123-3232310100313020)
- [ip_prefix_set](resources--network_policy_rule--reference--group-001.md#canonical-1330023003330311-0120320212023000-3011113312022212-3321212113003322-1322030310230311-0020213122332002-2032202202130303-0101011123223132)
- [label_matcher](resources--network_policy_rule--reference--group-001.md#canonical-3222211312301230-0213003220131133-1103113301100322-1312010120330201-0032020133120322-2230013122010001-3003031332303220-0010012123030031)
- [prefix](resources--network_policy_rule--reference--group-001.md#canonical-1322332132013232-0332313113100120-1122031211122021-3202302023323010-0032030221021030-1300222000022322-1223313012310311-1320233203031322)
- [prefix_selector](resources--network_policy_rule--reference--group-001.md#canonical-0123133033332220-1332302021201021-1311031103013223-0132030013110331-1012010100011001-3211212222100321-3201102203021112-3332122012022222)
- [timeouts](resources--network_policy_rule--reference--group-001.md#canonical-1300032020203311-0001033032221233-3321113100111211-3032113200300222-2300201033230220-2033011300322212-3322021131132110-2301210013002332)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)

<a id="canonical-2201003201203002-1001021032313331-0000232203010002-0000231003020102-3001203210231122-3222031020113232-2022101230331123-3232310100313020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331120323233302-1032013032002322-2133123203103203-0210211022100332-0222332123333211-1022110101231300-2121203011230303-1302123232022303"></a>

## advanced_action — advanced_action / 201031021312 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300)
- advanced_action

<a id="canonical-1213112312010113-1233020111010301-3030313013013320-2102001033331320-2132233013302122-0221131132132130-2021011030321322-3323220303222011"></a>

Type: `"object"`. single nested block, Optional.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

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
advanced_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202300013210333-2303212213023210-1320000112123333-1230203101333003-0203302202133233-1110000032311330-1003012322201201-1031323113230020"></a>

## Direct properties — advanced_action / 201031021312 / 3

<a id="canonical-2030220030232120-1031320233031200-1312100301231313-3031113123122331-1230212213211132-3130303233103033-1012011021023122-1003131130010211"></a>

<a id="canonical-2103112201133321-0000122113113302-1131113020222223-2323103013330303-0011203321332213-0102302102331120-2301032013212110-2211322111232100"></a>

## action property — advanced_action / 201031021312 / 4

Type: `"string"`. Optional.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Upstream description:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NOLOG",
    "LOG"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1012331300030022-3212122322000322-3102320123123022-1330332313123030-2123302001101301-2103011121031111-3023122032022201-3313021010121303"></a>

## Next pages — advanced_action / 201031021312 / 5

- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)

<a id="canonical-1330023003330311-0120320212023000-3011113312022212-3321212113003322-1322030310230311-0020213122332002-2032202202130303-0101011123223132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331113121312100-1102001302220200-0203011300003022-2103132122021133-2110001300332210-1312021132112021-1120300002121230-0002313310222103"></a>

## ip_prefix_set — ip_prefix_set / 203331230333 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300)
- ip_prefix_set

<a id="canonical-1220331221100333-2122122320021210-2310121030231032-1213033300131213-2312033301220001-3001332021210202-0321001120111132-1031121322003120"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ip\_prefix\_set, prefix, prefix\_selector\] List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

- [ip_prefix_set](resources--network_policy_rule--reference--group-001.md#canonical-1220331221100333-2122122320021210-2310121030231032-1213033300131213-2312033301220001-3001332021210202-0321001120111132-1031121322003120)
- [prefix](resources--network_policy_rule--reference--group-001.md#canonical-0130030202032220-0213030001022100-2121322330220121-2030211210232323-3132101023031003-2310210111201331-3330121133020213-1333133032101103)
- [prefix_selector](resources--network_policy_rule--reference--group-001.md#canonical-3102211033013012-2230100222330123-0200112112031212-1033010111122221-2103231202000323-3012130232330310-0021002230022233-0032202300130333)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213222223131312-1210031220322013-0321332100030123-3030112211033023-1023103022320231-2200002131202101-1320101303311111-1322220020211000"></a>

## Direct properties — ip_prefix_set / 203331230333 / 3

- [ref](resources--network_policy_rule--reference--group-001.md#canonical-1100211033101221-0000111111010132-3031132012203313-1310320132010201-0100312011123123-2313123211332030-3322302322311302-3212311002030223): complete subsection reference.

<a id="canonical-2013112000300211-2103001121022003-3203333113323020-0021132132132332-1331322222210030-0003131202311232-3010123221200223-0022130003112012"></a>

## Next pages — ip_prefix_set / 203331230333 / 4

- [ip_prefix_set.ref](resources--network_policy_rule--reference--group-001.md#canonical-1100211033101221-0000111111010132-3031132012203313-1310320132010201-0100312011123123-2313123211332030-3322302322311302-3212311002030223)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)

<a id="canonical-1100211033101221-0000111111010132-3031132012203313-1310320132010201-0100312011123123-2313123211332030-3322302322311302-3212311002030223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111233013021133-3220110100101032-2333031103312303-1000200133210101-1202223210211033-0112212232333331-2102230211310203-2300313331231100"></a>

## ip_prefix_set.ref — ref / 013001012321 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300)
- [ip_prefix_set](resources--network_policy_rule--reference--group-001.md#canonical-1330023003330311-0120320212023000-3011113312022212-3321212113003322-1322030310230311-0020213122332002-2032202202130303-0101011123223132)
- ip_prefix_set.ref

<a id="canonical-1110113200012112-3132132313130022-1000223111212130-3202222030113120-1310201122222121-2323100310331103-1221000212320210-0010322102023301"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202322132120210-0330022002133233-3013123111211331-1210022122013223-3111331122300103-0012223100133231-0003121210122132-1320210110203112"></a>

## Direct properties — ref / 013001012321 / 3

<a id="canonical-3212322113100221-1031332201223113-2210132001220101-3311313031030122-0101002130221122-2301110323300013-1113332122032221-2012123330111231"></a>

<a id="canonical-0131023223121010-3300331113020201-1233303002021120-3201220230033021-2220321313200301-2213022130222103-0133122230221201-3230101332011130"></a>

## kind property — ref / 013001012321 / 4

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

<a id="canonical-3131132123223120-3131131310332002-0221112311133232-0313222230220130-2311222033132220-3200222303133302-3132023223321222-2232122032133002"></a>

<a id="canonical-2033100330231202-1233210131201132-0020032202022030-3223123032322312-0000332202313113-3021002121211010-3233233023003331-0022323101101123"></a>

## name property — ref / 013001012321 / 5

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

<a id="canonical-0031211213102321-0020112222031211-1200231323010321-0021012031322233-0131112023020223-0020311132000032-3300121201022032-3003203001032111"></a>

<a id="canonical-1013022021312332-1321013322301102-2303011102331013-0010201123021210-2131130021120323-3231311102310313-3320310312130232-3011223320322121"></a>

## namespace property — ref / 013001012321 / 6

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

<a id="canonical-2103330322111012-0122033031303221-0312101133002231-1333220131033331-1223303233330102-2320013033101003-0133222001000323-1020020330023111"></a>

<a id="canonical-0032301130312031-1331033113230100-2220033000133123-3220321220323033-1320211000030100-1312030221300030-3302313131020200-0210112103301000"></a>

## tenant property — ref / 013001012321 / 7

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

<a id="canonical-2023201223223123-2133202211133133-2133122030312101-1001022022332310-2232121101131313-3102113020122103-2113001013310023-2132120033202113"></a>

<a id="canonical-0102212201231002-2131122301002231-3013112232230203-1131031210233031-0022211331202102-2022310013113000-0320111332031321-0102133301231201"></a>

## uid property — ref / 013001012321 / 8

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

<a id="canonical-2023332130113100-1311311302313331-1333310312301000-1311023000031032-3121331220331213-2332120233213001-3302310003030123-0032210300213101"></a>

## Next pages — ref / 013001012321 / 9

- [ip_prefix_set](resources--network_policy_rule--reference--group-001.md#canonical-1330023003330311-0120320212023000-3011113312022212-3321212113003322-1322030310230311-0020213122332002-2032202202130303-0101011123223132)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)

<a id="canonical-3222211312301230-0213003220131133-1103113301100322-1312010120330201-0032020133120322-2230013122010001-3003031332303220-0010012123030031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100221210333201-3011120300000131-2030122230033122-1331110323102203-2200220020232230-3030230003313031-3332112202112320-1300320200313311"></a>

## label_matcher — label_matcher / 232131110302 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300)
- label_matcher

<a id="canonical-3333112010110111-3201300112222012-0303003021312311-1232100102310300-3233022003130133-2320131131021322-0211233221200101-3210210300031302"></a>

Type: `"object"`. single nested block, Optional.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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
label_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221311233103300-2033323101032210-2131202120233030-3321111101323033-0002300232102323-2122102310223203-2300222131300302-3120010120021132"></a>

## Direct properties — label_matcher / 232131110302 / 3

<a id="canonical-0112310131100032-2223020213302201-2210023111023100-3202032232322003-1110022212233222-0130031010201122-2101102023332200-2023221232200132"></a>

<a id="canonical-0033213211113023-0020303221211111-3313011330233201-2203120031003201-0111323030110110-1333010002202212-3231203302212303-0003220130030132"></a>

## keys property — label_matcher / 232131110302 / 4

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0300003333101102-2331131112302032-3110121032211000-2310111103132112-0202120212232321-0212013021312011-3213112120022211-3303300010001222"></a>

## Next pages — label_matcher / 232131110302 / 5

- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)

<a id="canonical-1322332132013232-0332313113100120-1122031211122021-3202302023323010-0032030221021030-1300222000022322-1223313012310311-1320233203031322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103113022003023-2212031312310312-2121320231031331-3230000110011213-1323330210221201-2311020123212032-3220132011211223-3233310202313222"></a>

## prefix — prefix / 002330102203 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300)
- prefix

<a id="canonical-0130030202032220-0213030001022100-2121322330220121-2030211210232323-3132101023031003-2310210111201331-3330121133020213-1333133032101103"></a>

Type: `"object"`. single nested block, Optional.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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
prefix {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000220230132011-2022102110202012-1002230122110213-2303300213213310-2322023202112003-0213331122212230-1021230131112202-2212103330333213"></a>

## Direct properties — prefix / 002330102203 / 3

<a id="canonical-0022310031113230-3321130203212013-0021320300110031-0232301101231332-0330301122130213-3313003100320101-0102020133131222-0003333112100301"></a>

<a id="canonical-0021220022311310-1223021322311223-3003210003222213-3313312020122123-3231320132103320-0321131303012223-0131031020022211-1332031113201200"></a>

## prefix property — prefix / 002330102203 / 4

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-1023132131101203-1113300331203310-1130322301202103-0231231300210320-2000212100313222-0213313200003222-3312130111013012-3033120001030130"></a>

## Next pages — prefix / 002330102203 / 5

- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)

<a id="canonical-0123133033332220-1332302021201021-1311031103013223-0132030013110331-1012010100011001-3211212222100321-3201102203021112-3332122012022222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112032203120222-3321032130332100-3203112111020002-1031303231310202-3100012100102010-0201021201103113-1220221313033222-3212333021133220"></a>

## prefix_selector — prefix_selector / 010311310031 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300)
- prefix_selector

<a id="canonical-3102211033013012-2230100222330123-0200112112031212-1033010111122221-2103231202000323-3012130232330310-0021002230022233-0032202300130333"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
prefix_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303023212102132-0103000100221031-3100033202302110-2233132101112213-3200220322113033-3310202221220300-2031122121023213-0132130103122010"></a>

## Direct properties — prefix_selector / 010311310031 / 3

<a id="canonical-1303233210032333-2310011001220333-0303201113203001-1201113021113313-2331302311321202-0330322030212112-3202331320013132-3330223303012103"></a>

<a id="canonical-2222112303013110-1200110112220322-1332111010301111-0332320312333010-1133010331133321-2230320100220311-0032102312211102-1102313303030131"></a>

## expressions property — prefix_selector / 010311310031 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0310223032000022-0112203222103211-2230012011021023-0003221300013033-0100120333232230-2130123320220211-2001030133002321-0013131111002132"></a>

## Next pages — prefix_selector / 010311310031 / 5

- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)

<a id="canonical-1300032020203311-0001033032221233-3321113100111211-3032113200300222-2300201033230220-2033011300322212-3322021131132110-2301210013002332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203013032330203-1301330231320032-3310031102302312-1112102110232001-2021222211332311-0121321130102333-3313031131112320-1013110220323123"></a>

## timeouts — timeouts / 021200233223 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300)
- timeouts

<a id="canonical-2032212123200333-1311002210201033-3013013303211011-2232131123330331-1302212120123000-2121230001333300-3322030000223032-0232121012110020"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012031302231133-3212333230102230-3213012001033200-3120003321303123-0032302310130310-2322203022203033-2003013130201113-2023020101201132"></a>

## Direct properties — timeouts / 021200233223 / 3

<a id="canonical-0222333301321122-1302122332312200-0110023322030322-1322212320322200-2213130002220330-1230303012102330-1233122122323123-2311300002201302"></a>

<a id="canonical-0332232110323300-2312120132100302-2111130111203000-2122333130300132-2320021200000220-1202203131210221-2121303200011212-1103011022031012"></a>

## create property — timeouts / 021200233223 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2010033113300000-3231200133010220-2201102313312002-1212313111201011-3313111030222131-1313011222331203-2332330231212220-3120011202230032"></a>

<a id="canonical-2331313310032123-0301332211013230-1213303222310130-1133311131000311-1020132223202112-3222332202223213-2201202303320322-1312000320213022"></a>

## delete property — timeouts / 021200233223 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1321211100101102-1323331110213301-0132023101201002-0001010132101321-2200232310131320-2333331022322111-3333001332001121-3203220232010010"></a>

<a id="canonical-1203201132202332-1011222312231103-2121320333220032-0302320302001230-2330223203200131-1032033233230223-1013012031021112-2133333330321130"></a>

## read property — timeouts / 021200233223 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1211323320133201-3031321113023101-2130233330312322-3311200000301222-2121211000203010-1022020110012231-3323200313132230-2021031321121321"></a>

<a id="canonical-3313132323013202-3001302003012322-0113302113221311-3101123023000032-2303233313220123-3122310220332312-0213330031300121-0020012321320101"></a>

## update property — timeouts / 021200233223 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1021013020002312-3100230230313322-2232313210200232-0322100322112133-0002223222210001-1301110131301321-1011220313112332-1203303222223201"></a>

## Next pages — timeouts / 021200233223 / 8

- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)
