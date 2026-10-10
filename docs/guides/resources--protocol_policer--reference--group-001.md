---
page_title: "xcsh_protocol_policer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_policer reference."
---

# xcsh_protocol_policer reference

<a id="canonical-2131200202001010-0000130302120213-0011120213023122-1123322102232000-1023213310323311-0232321203023223-1310032211102313-0221101102322221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-3330333012213023-3331002310331112-1210012223120003-2212223210101200-3103330300201100-0230232202320033-0111223122201102-2133311232122222)
- Property reference

<a id="canonical-3220333032231322-2023212332132323-2013212223002131-2122202303321011-3022033131003312-3321023113322202-3301311301110302-2113123002301203"></a>

### Direct properties for `xcsh_protocol_policer`

<a id="canonical-2222300132201122-0121231333320332-0122023330100223-0033222300213322-3312302333223012-3231333112202303-2101213103000221-2002303112111321"></a>

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

<a id="canonical-1332013332201211-1022322300300213-2300320033312203-3211112012233103-0121313020010120-1222212112203320-0130300211001231-2230312113023112"></a>

<a id="canonical-2121320233220100-1131033033222220-3022212013301133-3012013311211210-1313203212022121-3033323310201220-2303101023022212-1231203301201102"></a>

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

<a id="canonical-1002330112112120-0201132000322102-1020112112030202-1332033023010130-2333102211031220-3232023200120013-3230033021302232-3000002013000330"></a>

<a id="canonical-0113220212111012-2101230221233133-2233300312133022-2112330333222230-2332100023221002-3131033300123201-1211003112220212-2220333200123012"></a>

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

<a id="canonical-2200322302130202-2111300022331103-0210002312323311-1213331313032101-1011232011130331-0311022020023022-3311022221302100-1313333001013330"></a>

<a id="canonical-2221033201310003-1212200302313300-3202133233122303-1212233122132213-3220020210003212-1101230020323032-0021202213220331-1023231220113320"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1210011223031333-2111203001122011-3032302110213330-3012132301111010-3001310330320332-3300321322130113-2213201213102103-2221121200032003"></a>

<a id="canonical-0300322002122301-3100311011101201-3311300130131300-3312010031211033-0011130121233323-0310230323300003-2223023132332233-3101110211333201"></a>

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

<a id="canonical-2221212210001322-1121200212323333-0311101030013322-0013223003113131-3102002333023023-2103030103000300-3213130131201010-1223013103101101"></a>

<a id="canonical-3110320032233020-3312210310330103-3223020212323003-3321211330213103-3103112231200010-3311133232103222-3020301111102100-3303333133130033"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Protocol Policer. Must be unique within the namespace.

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

<a id="canonical-0233202320331310-0122121032312011-2221103203021223-1310103133223323-0113312011211030-2201201022302302-1313311112021213-0012222020001003"></a>

<a id="canonical-2102020303231111-2222100210000310-2120113100011301-1302010000003101-0022110312103132-1221233301000111-1222202220122002-2011021332130133"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace for the Protocol Policer. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
EnumExtractionComplete: false
EnumValidators: [{"version":1,"validator":"OneOf","values":["system"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-1103203330131213-3232300131010200-1312133231130010-2111331121132203-3131111112032203-2132000223302321-0222101101022113-1011231101032132): complete subsection reference.

- [timeouts](resources--protocol_policer--reference--group-001.md#canonical-3102310322120033-2332320330102313-2130000322100222-1223003320213122-0112111112001303-2111110102123120-2223111202002132-1030301021111212): complete subsection reference.

<a id="canonical-0212303101220320-0131203012100012-3023202233032201-0022213331300230-1333120022122032-3320213101231121-3030200331030203-1202312303023032"></a>

### All schema paths for `xcsh_protocol_policer`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--protocol_policer--reference--group-001.md#canonical-2222300132201122-0121231333320332-0122023330100223-0033222300213322-3312302333223012-3231333112202303-2101213103000221-2002303112111321) |
| `description` | [description](resources--protocol_policer--reference--group-001.md#canonical-1332013332201211-1022322300300213-2300320033312203-3211112012233103-0121313020010120-1222212112203320-0130300211001231-2230312113023112) |
| `disable` | [disable](resources--protocol_policer--reference--group-001.md#canonical-1002330112112120-0201132000322102-1020112112030202-1332033023010130-2333102211031220-3232023200120013-3230033021302232-3000002013000330) |
| `id` | [ID](resources--protocol_policer--reference--group-001.md#canonical-2200322302130202-2111300022331103-0210002312323311-1213331313032101-1011232011130331-0311022020023022-3311022221302100-1313333001013330) |
| `labels` | [labels](resources--protocol_policer--reference--group-001.md#canonical-1210011223031333-2111203001122011-3032302110213330-3012132301111010-3001310330320332-3300321322130113-2213201213102103-2221121200032003) |
| `name` | [name](resources--protocol_policer--reference--group-001.md#canonical-2221212210001322-1121200212323333-0311101030013322-0013223003113131-3102002333023023-2103030103000300-3213130131201010-1223013103101101) |
| `namespace` | [namespace](resources--protocol_policer--reference--group-001.md#canonical-0233202320331310-0122121032312011-2221103203021223-1310103133223323-0113312011211030-2201201022302302-1313311112021213-0012222020001003) |
| `protocol_policer` | [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-3320310323201310-0333000312300301-2311131220212110-1001023230120222-3222213222023223-2303131030220030-2300220332010312-3101302203120130) |
| `protocol_policer.policer` | [protocol_policer.policer](resources--protocol_policer--reference--group-001.md#canonical-1203313002011222-3301000223121001-1222020212020102-3320311202131101-2103003132202310-2022001322100333-2202202133303033-3032002001133131) |
| `protocol_policer.policer.kind` | [protocol_policer.policer.kind](resources--protocol_policer--reference--group-001.md#canonical-3113330332110231-0212001310313123-1301110303032310-0132333223133112-3212202111132312-2130231003211312-3230333320233311-2131000102013212) |
| `protocol_policer.policer.name` | [protocol_policer.policer.name](resources--protocol_policer--reference--group-001.md#canonical-1111303003031203-0001232033301032-2303310120002112-2230130302033100-1230133310201312-2020013223102203-0313213031323320-1031112111331233) |
| `protocol_policer.policer.namespace` | [protocol_policer.policer.namespace](resources--protocol_policer--reference--group-001.md#canonical-0122300221133011-2202110013220323-2312323101030302-2000022001120020-3222233212012212-0130013312301310-2131033120001023-2001120331132012) |
| `protocol_policer.policer.tenant` | [protocol_policer.policer.tenant](resources--protocol_policer--reference--group-001.md#canonical-2332312332102112-2303301211020223-3322133331113220-1210330003232123-1110113100011102-1202013001131011-1131131133232232-1332033033311312) |
| `protocol_policer.policer.uid` | [protocol_policer.policer.uid](resources--protocol_policer--reference--group-001.md#canonical-0202330113121000-0020222003122302-0003320230032022-2230311000003011-3203113223031012-3022121113110030-2013221031110000-1313330230021032) |
| `protocol_policer.protocol` | [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-0021312311131021-0300202131021323-0203311300130320-0230111301333233-1121213303300332-2211303320323030-1203221022233131-2320311210032213) |
| `protocol_policer.protocol.dns` | [protocol_policer.protocol.dns](resources--protocol_policer--reference--group-001.md#canonical-0020020222010110-3003331033122202-0202210210130122-1202311232002003-1333030210002111-1223332233001111-3300122223133313-1211122133333023) |
| `protocol_policer.protocol.icmp` | [protocol_policer.protocol.icmp](resources--protocol_policer--reference--group-001.md#canonical-1331011212201303-2103213331120130-2000013331203022-1133310131212031-1122012201201221-0310221112001102-3201003003222133-0112220023210201) |
| `protocol_policer.protocol.icmp.type` | [protocol_policer.protocol.icmp.type](resources--protocol_policer--reference--group-001.md#canonical-2231131131332012-0220032010120031-1233001021002131-0310303132012032-2133300130013120-1130302101200201-2023132310201102-1111022130230232) |
| `protocol_policer.protocol.tcp` | [protocol_policer.protocol.tcp](resources--protocol_policer--reference--group-001.md#canonical-2132012323301030-2030203111302223-0110313230120112-2202120321101201-2301121333013013-3021012020211133-1120233102012133-3203131221210321) |
| `protocol_policer.protocol.tcp.flags` | [protocol_policer.protocol.tcp.flags](resources--protocol_policer--reference--group-001.md#canonical-3021031333103230-0132330010010010-0112233112301311-0222320130333223-1201100213311210-0131101302133232-2311032133120031-0220223220300101) |
| `protocol_policer.protocol.udp` | [protocol_policer.protocol.udp](resources--protocol_policer--reference--group-001.md#canonical-2023000222010123-2213323332021021-2021022300011310-2233210300021003-0233101101011130-3200012333031301-3100123232113002-2231210000121013) |
| `timeouts` | [timeouts](resources--protocol_policer--reference--group-001.md#canonical-1013333211210012-0002332112102032-3330020212132222-1131211320310010-1121312312030102-1022021110113000-2210132030132232-2121013101012101) |
| `timeouts.create` | [timeouts.create](resources--protocol_policer--reference--group-001.md#canonical-1101110001331320-2300103321311303-0031310101233332-1211212230022322-2322110222102001-0111113311201031-0102230323321032-0132333322020003) |
| `timeouts.delete` | [timeouts.delete](resources--protocol_policer--reference--group-001.md#canonical-1330132110220120-1203303231022012-3032102210311201-2231312102302130-3120233223112303-3101310323300232-0131213123112033-3013011331232211) |
| `timeouts.read` | [timeouts.read](resources--protocol_policer--reference--group-001.md#canonical-3112013300222203-0233220001330002-0202203201230022-2313321212200322-1011201032031122-0203112220331112-1212111100111110-1203133001322333) |
| `timeouts.update` | [timeouts.update](resources--protocol_policer--reference--group-001.md#canonical-1220221313203131-2300220211003120-0213310200321010-1301100213223310-0211102132212302-0121111212221322-1021031330311221-0113102020012302) |

<a id="canonical-1103203330131213-3232300131010200-1312133231130010-2111331121132203-3131111112032203-2132000223302321-0222101101022113-1011231101032132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_policer` properties

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-3330333012213023-3331002310331112-1210012223120003-2212223210101200-3103330300201100-0230232202320033-0111223122201102-2133311232122222)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-2131200202001010-0000130302120213-0011120213023122-1123322102232000-1023213310323311-0232321203023223-1310032211102313-0221101102322221)
- protocol_policer

<a id="canonical-3320310323201310-0333000312300301-2311131220212110-1001023230120222-3222213222023223-2303131030220030-2300220332010312-3101302203120130"></a>

Type: `"object"`. list nested block, Optional.

List of L4 protocol match condition and associated traffic rate limits.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("policer")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
protocol_policer {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202210133220220-2010032010013320-0332001331102113-1112023120103302-2213001113302230-0231122233213300-2220232020310120-2011132021323312"></a>

### Direct properties for `protocol_policer`

- [policer](resources--protocol_policer--reference--group-001.md#canonical-0211121213320101-3012122102131103-0330313002133011-0032111032201300-1223300233311020-0203301322031012-1312221122233321-1131302131301032): complete subsection reference.

- [protocol](resources--protocol_policer--reference--group-001.md#canonical-1100220312202001-1132123022333322-1131101032220112-2213013000211201-2032331131032232-0230102201211101-0223303131111000-0332211321020103): complete subsection reference.

<a id="canonical-0211121213320101-3012122102131103-0330313002133011-0032111032201300-1223300233311020-0203301322031012-1312221122233321-1131302131301032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_policer.policer` properties

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-3330333012213023-3331002310331112-1210012223120003-2212223210101200-3103330300201100-0230232202320033-0111223122201102-2133311232122222)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-2131200202001010-0000130302120213-0011120213023122-1123322102232000-1023213310323311-0232321203023223-1310032211102313-0221101102322221)
- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-1103203330131213-3232300131010200-1312133231130010-2111331121132203-3131111112032203-2132000223302321-0222101101022113-1011231101032132)
- protocol_policer.policer

<a id="canonical-1203313002011222-3301000223121001-1222020212020102-3320311202131101-2103003132202310-2022001322100333-2202202133303033-3032002001133131"></a>

Type: `"object"`. list nested block, Optional.

Reference to policer object to apply traffic rate limits.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
policer {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220102020022302-0111103332033110-3313323023301033-1131330210000222-2123112223113200-2101333113100201-3031100230101230-0223133212032101"></a>

### Direct properties for `protocol_policer.policer`

<a id="canonical-3113330332110231-0212001310313123-1301110303032310-0132333223133112-3212202111132312-2130231003211312-3230333320233311-2131000102013212"></a>

#### `protocol_policer.policer.kind` property

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

<a id="canonical-1111303003031203-0001232033301032-2303310120002112-2230130302033100-1230133310201312-2020013223102203-0313213031323320-1031112111331233"></a>

<a id="canonical-3002121211203222-0233203210013102-2113203333021322-2123211310321201-2323120322201230-2230032210332202-2320010131121212-1123103232303010"></a>

#### `protocol_policer.policer.name` property

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

<a id="canonical-0122300221133011-2202110013220323-2312323101030302-2000022001120020-3222233212012212-0130013312301310-2131033120001023-2001120331132012"></a>

<a id="canonical-3122301210010133-1203320200232320-1220100111311012-3333231200200221-2323101023302302-1202133021230301-2112020113330221-3202003321210133"></a>

#### `protocol_policer.policer.namespace` property

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

<a id="canonical-2332312332102112-2303301211020223-3322133331113220-1210330003232123-1110113100011102-1202013001131011-1131131133232232-1332033033311312"></a>

<a id="canonical-2203000312131320-3021303330003232-3033132223230033-3201323030222111-1112312303300301-0023013100113121-0101002012311300-2323320333323202"></a>

#### `protocol_policer.policer.tenant` property

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

<a id="canonical-0202330113121000-0020222003122302-0003320230032022-2230311000003011-3203113223031012-3022121113110030-2013221031110000-1313330230021032"></a>

<a id="canonical-3300130022212310-1330000312010130-0213312032301312-1300210001313023-3300013031130010-1033233211110112-2030212323121032-0020123301213323"></a>

#### `protocol_policer.policer.uid` property

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

<a id="canonical-1100220312202001-1132123022333322-1131101032220112-2213013000211201-2032331131032232-0230102201211101-0223303131111000-0332211321020103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_policer.protocol` properties

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-3330333012213023-3331002310331112-1210012223120003-2212223210101200-3103330300201100-0230232202320033-0111223122201102-2133311232122222)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-2131200202001010-0000130302120213-0011120213023122-1123322102232000-1023213310323311-0232321203023223-1310032211102313-0221101102322221)
- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-1103203330131213-3232300131010200-1312133231130010-2111331121132203-3131111112032203-2132000223302321-0222101101022113-1011231101032132)
- protocol_policer.protocol

<a id="canonical-0021312311131021-0300202131021323-0203311300130320-0230111301333233-1121213303300332-2211303320323030-1203221022233131-2320311210032213"></a>

Type: `"object"`. single nested block, Optional.

Protocol and protocol specific flags to be matched in packet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("dns",
    "icmp"),
  validators.ConflictingObjectAttributes("dns",
    "tcp"),
  validators.ConflictingObjectAttributes("dns",
    "udp"),
  validators.ConflictingObjectAttributes("icmp",
    "tcp"),
  validators.ConflictingObjectAttributes("icmp",
    "udp"),
  validators.ConflictingObjectAttributes("tcp",
    "udp")}
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
  "x-ves-oneof-field-type": "[\"dns\",\"icmp\",\"tcp\",\"udp\"]"
}
```

Terraform syntax:

```terraform
protocol {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232032333321023-2331210202322133-1121311033311221-2233002122310233-3103013131121113-1120233002300121-2001321212233011-2033100103130333"></a>

### Direct properties for `protocol_policer.protocol`

- [DNS](resources--protocol_policer--reference--group-001.md#canonical-1023101020133320-0101113221211233-2202301130212333-1321110111123120-1230023031330212-2321230012213102-3332320112230210-3300112231323233): complete subsection reference.

- [icmp](resources--protocol_policer--reference--group-001.md#canonical-3012120311331130-1110220122213203-2133122032110311-0133210023300302-0002032320212302-0122030321120221-1313031210131120-0101333111113001): complete subsection reference.

- [tcp](resources--protocol_policer--reference--group-001.md#canonical-0122020023100021-0020003133211223-1133032331211030-2223201003111100-0330213031012130-2223022230231131-1122312030003211-1123003022032011): complete subsection reference.

- [udp](resources--protocol_policer--reference--group-001.md#canonical-3232001013002031-3120311332330222-3021000303020022-2020032310230232-1030303131021111-2312112011011111-3212032322233022-3321311010112131): complete subsection reference.

<a id="canonical-1023101020133320-0101113221211233-2202301130212333-1321110111123120-1230023031330212-2321230012213102-3332320112230210-3300112231323233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_policer.protocol.dns` properties

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-3330333012213023-3331002310331112-1210012223120003-2212223210101200-3103330300201100-0230232202320033-0111223122201102-2133311232122222)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-2131200202001010-0000130302120213-0011120213023122-1123322102232000-1023213310323311-0232321203023223-1310032211102313-0221101102322221)
- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-1103203330131213-3232300131010200-1312133231130010-2111331121132203-3131111112032203-2132000223302321-0222101101022113-1011231101032132)
- [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-1100220312202001-1132123022333322-1131101032220112-2213013000211201-2032331131032232-0230102201211101-0223303131111000-0332211321020103)
- protocol_policer.protocol.DNS

<a id="canonical-0020020222010110-3003331033122202-0202210210130122-1202311232002003-1333030210002111-1223332233001111-3300122223133313-1211122133333023"></a>

Type: `["object", {}]`. Optional.

Match all DNS packets including UDP and TCP.

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
dns = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012120311331130-1110220122213203-2133122032110311-0133210023300302-0002032320212302-0122030321120221-1313031210131120-0101333111113001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_policer.protocol.icmp` properties

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-3330333012213023-3331002310331112-1210012223120003-2212223210101200-3103330300201100-0230232202320033-0111223122201102-2133311232122222)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-2131200202001010-0000130302120213-0011120213023122-1123322102232000-1023213310323311-0232321203023223-1310032211102313-0221101102322221)
- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-1103203330131213-3232300131010200-1312133231130010-2111331121132203-3131111112032203-2132000223302321-0222101101022113-1011231101032132)
- [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-1100220312202001-1132123022333322-1131101032220112-2213013000211201-2032331131032232-0230102201211101-0223303131111000-0332211321020103)
- protocol_policer.protocol.icmp

<a id="canonical-1331011212201303-2103213331120130-2000013331203022-1133310131212031-1122012201201221-0310221112001102-3201003003222133-0112220023210201"></a>

Type: `"object"`. single nested block, Optional.

ICMP Packet Type. ICMP message type to match in packet.

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
icmp {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230001301332123-2202211013031332-1321121020010102-0032023332022022-1333220012313201-1230203302010333-1312311313300220-0101122101212232"></a>

### Direct properties for `protocol_policer.protocol.icmp`

<a id="canonical-2231131131332012-0220032010120031-1233001021002131-0310303132012032-2133300130013120-1130302101200201-2023132310201102-1111022130230232"></a>

#### `protocol_policer.protocol.icmp.type` property

Type: `["list", "string"]`. Optional.

\[Enum: ECHO\_REPLY|ECHO\_REQUEST|ALL\_ICMP\_MSG\] ICMP message type to be matched in packet.
Possible values are \`ECHO\_REPLY\`, \`ECHO\_REQUEST\`, \`ALL\_ICMP\_MSG\`. Defaults to
\`ECHO\_REPLY\`.

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

<a id="canonical-0122020023100021-0020003133211223-1133032331211030-2223201003111100-0330213031012130-2223022230231131-1122312030003211-1123003022032011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_policer.protocol.tcp` properties

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-3330333012213023-3331002310331112-1210012223120003-2212223210101200-3103330300201100-0230232202320033-0111223122201102-2133311232122222)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-2131200202001010-0000130302120213-0011120213023122-1123322102232000-1023213310323311-0232321203023223-1310032211102313-0221101102322221)
- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-1103203330131213-3232300131010200-1312133231130010-2111331121132203-3131111112032203-2132000223302321-0222101101022113-1011231101032132)
- [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-1100220312202001-1132123022333322-1131101032220112-2213013000211201-2032331131032232-0230102201211101-0223303131111000-0332211321020103)
- protocol_policer.protocol.tcp

<a id="canonical-2132012323301030-2030203111302223-0110313230120112-2202120321101201-2301121333013013-3021012020211133-1120233102012133-3203131221210321"></a>

Type: `"object"`. single nested block, Optional.

Specification of TCP flag to be matched in a TCP packet.

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
tcp {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323223313103023-0112232310023011-0020221131133200-2301312323202323-1201302111101131-0032230132001223-3022131330132300-1201330102031322"></a>

### Direct properties for `protocol_policer.protocol.tcp`

<a id="canonical-3021031333103230-0132330010010010-0112233112301311-0222320130333223-1201100213311210-0131101302133232-2311032133120031-0220223220300101"></a>

#### `protocol_policer.protocol.tcp.flags` property

Type: `["list", "string"]`. Optional.

\[Enum: FIN|SYN|RST|PSH|ACK|URG|ALL\_TCP\_FLAGS|KEEPALIVE\] TCP flags. TCP flag to be matched in a
TCP packet. Possible values are \`FIN\`, \`SYN\`, \`RST\`, \`PSH\`, \`ACK\`, \`URG\`,
\`ALL\_TCP\_FLAGS\`, \`KEEPALIVE\`. Defaults to \`FIN\`.

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

<a id="canonical-3232001013002031-3120311332330222-3021000303020022-2020032310230232-1030303131021111-2312112011011111-3212032322233022-3321311010112131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_policer.protocol.udp` properties

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-3330333012213023-3331002310331112-1210012223120003-2212223210101200-3103330300201100-0230232202320033-0111223122201102-2133311232122222)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-2131200202001010-0000130302120213-0011120213023122-1123322102232000-1023213310323311-0232321203023223-1310032211102313-0221101102322221)
- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-1103203330131213-3232300131010200-1312133231130010-2111331121132203-3131111112032203-2132000223302321-0222101101022113-1011231101032132)
- [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-1100220312202001-1132123022333322-1131101032220112-2213013000211201-2032331131032232-0230102201211101-0223303131111000-0332211321020103)
- protocol_policer.protocol.udp

<a id="canonical-2023000222010123-2213323332021021-2021022300011310-2233210300021003-0233101101011130-3200012333031301-3100123232113002-2231210000121013"></a>

Type: `["object", {}]`. Optional.

UDP Packets. Match all UDP packets.

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
udp = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102310322120033-2332320330102313-2130000322100222-1223003320213122-0112111112001303-2111110102123120-2223111202002132-1030301021111212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-3330333012213023-3331002310331112-1210012223120003-2212223210101200-3103330300201100-0230232202320033-0111223122201102-2133311232122222)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-2131200202001010-0000130302120213-0011120213023122-1123322102232000-1023213310323311-0232321203023223-1310032211102313-0221101102322221)
- timeouts

<a id="canonical-1013333211210012-0002332112102032-3330020212132222-1131211320310010-1121312312030102-1022021110113000-2210132030132232-2121013101012101"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310011213103312-3333200023311233-3101133102133100-3023120010312003-2203132001333121-0303301210130112-0012031031211123-3221310000301031"></a>

### Direct properties for `timeouts`

<a id="canonical-1101110001331320-2300103321311303-0031310101233332-1211212230022322-2322110222102001-0111113311201031-0102230323321032-0132333322020003"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1330132110220120-1203303231022012-3032102210311201-2231312102302130-3120233223112303-3101310323300232-0131213123112033-3013011331232211"></a>

<a id="canonical-3111030011311302-3302121010012013-1221013222102213-1103232333313233-0232320010102013-3110330110302133-1033123211020120-0132330301202333"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3112013300222203-0233220001330002-0202203201230022-2313321212200322-1011201032031122-0203112220331112-1212111100111110-1203133001322333"></a>

<a id="canonical-2132220111202323-0302332030313101-2012321303013000-3033232031033231-3223103320030313-1310131233122102-1122312001033201-3303233110103113"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1220221313203131-2300220211003120-0213310200321010-1301100213223310-0211102132212302-0121111212221322-1021031330311221-0113102020012302"></a>

<a id="canonical-3133121203320222-1133130023122132-1000021302002331-3212221123003320-0311131120303221-3000120232023111-2320200003333132-2113003110331212"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
