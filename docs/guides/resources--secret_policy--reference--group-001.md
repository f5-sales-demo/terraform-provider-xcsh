---
page_title: "xcsh_secret_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_secret_policy reference."
---

# xcsh_secret_policy reference

<a id="canonical-1213313330302120-1100122222322313-2323333222220332-2211122300300103-3232201121133200-0113220013210033-2322130013023011-1132310023231200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_secret_policy](../resources/secret_policy.md#canonical-1312021301020310-0131200100331301-3002003211003233-3010322032311123-0022212322233110-2113210022031113-0031130332112033-3113012311332111)
- Property reference

<a id="canonical-3101130330132312-1233212322100120-3011210220303013-3022100110002313-1303033031233023-0000011101011333-3211332332211111-2000120123300102"></a>

### Direct properties for `xcsh_secret_policy`

<a id="canonical-1300310323301312-2211032002222023-1023123132201302-1201333113010032-3102110203001212-0200131302303102-2301121231120033-2230323301332020"></a>

#### `allow_f5xc` property

Type: `"bool"`. Optional, Computed.

Additional upstream details:

If allow\_f5xc is set to true, it allows relevant F5XC infrastructure services to decrypt the secret
encrypted using this policy.

Provider validators and defaults (from schema source):

```go
Default: booldefault.StaticBool(false)
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

<a id="canonical-2210233301111102-2221012112213313-1031031122313011-3011023112001220-0333031123023002-0130133333221313-2101010133133110-2123303221111033"></a>

<a id="canonical-2010333202020232-1131021212333121-2010212023033002-2210110131210013-2201110122022302-2300311003333203-0320302201102213-2331120310330013"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

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

<a id="canonical-1101113121000313-3232113000332022-1132233320330321-3203200012110121-0203120303010001-0122223031020010-1033203021330321-2030133121033133"></a>

<a id="canonical-0023000020030003-2322030301101320-3001202302023301-2113231220030300-2211311002302223-2212333112101100-1200121320110333-0103300222213210"></a>

#### `decrypt_cache_timeout` property

Type: `"string"`. Optional, Computed.

Additional upstream details:

Decrypt\_cache\_timeout contains the amount of time a decrypted secret is cached in wingman. Value
for this parameter is a string ending in the suffix "s" (indicating seconds), suffix "m" (indicating
minutes) or suffix "h" (indicating hours)

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("30s")
EnumExtractionComplete: true
Validators: []validator.String{stringvalidator.RegexMatches(secretPolicyDuration,
    "Use a positive duration ending in s, m, or h.")}
```

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
      "validatedAt": "2026-10-16T10:31:02+00:00"
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

<a id="canonical-0203332013321131-3030003300022011-1323020111213132-1113210321233021-3123232312021231-0123000311113100-0202120110220213-1011022013101202"></a>

<a id="canonical-0230301210022111-3110222222000020-3100221331111210-1213231300201212-1320211321000033-1310330203132120-3022300132030321-0213000012013312"></a>

#### `deletion_time` property

Type: `"string"`. Computed.

<a id="canonical-1300211223201123-3110232002232123-0033312221132110-1003222330023331-2001201032022023-0301210212211010-0012032233330213-3303001223320132"></a>

<a id="canonical-2333133211222332-1210112003121321-1312310330201111-2013300200320101-3112110312123212-2220001311233331-0112122303103302-2113212220311333"></a>

#### `description` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-16T10:31:02+00:00"
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

<a id="canonical-0220332201332030-1002010212231032-0330000010200311-2233230211311131-2021011033231012-2011213202020330-0310201323230211-2010331221322320"></a>

<a id="canonical-1231110211311212-2213211033102032-3331113100031303-1223303333130203-3203112330202010-2102310200321202-2012313333231000-0201120313120202"></a>

#### `id` property

Type: `"string"`. Computed.

<a id="canonical-0132201201220222-2020222310221223-3001221231131123-3321012022311010-3212320311223311-1133111003131130-3111110030310013-2101112132301000"></a>

<a id="canonical-2013313013102213-0330320303113213-0211331101221001-2012013032200231-3100021120331301-3023011322211201-1110113011211121-3101201033001210"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

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

<a id="canonical-3330102003322001-1200320222031303-2100020130121201-0121100332332123-0031020112211003-2031212033011132-0123333310220202-1203011312210321"></a>

<a id="canonical-0303221132033121-0123021132223231-2021221010001212-3210022131111213-3132203131030201-2033232022311301-1010111211311203-1003213133233320"></a>

#### `marked_for_delete` property

Type: `"bool"`. Computed.

<a id="canonical-1013310232311101-2213011332322002-2303031111020211-0222232102230131-1223302122111203-1013310030311002-2311033121001003-3011320023331222"></a>

<a id="canonical-1321200302320313-0221010230333231-1322022000101033-1322322121000220-0103321032021112-1312200310021101-2030100200112133-1211033031022003"></a>

#### `name` property

Type: `"string"`. Required.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{validators.NameValidator()}
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
      "validatedAt": "2026-10-16T10:31:02+00:00"
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

<a id="canonical-3013303030313211-0112333003002133-1110030201310122-1011010020312303-0222123323033101-2320002113223221-0232203313120003-0110303311021121"></a>

<a id="canonical-3020122002232032-3222100132112113-3230211121232333-3312022000123112-2313132120132112-1101013101012203-3010113031330323-3033011100112012"></a>

#### `namespace` property

Type: `"string"`. Required.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{validators.NamespaceValidator()}
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
      "validatedAt": "2026-10-16T10:31:02+00:00"
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

<a id="canonical-1021110202020311-3122212110002211-3223203331313303-1331322321332100-1000223130200212-3202100202003133-3133202223202110-0122310300022202"></a>

<a id="canonical-0332230312322203-1102222302013021-1020122102220223-1012221221011000-2023032023302113-1001332000310220-3312302121330021-3203203020233300"></a>

#### `rule_combining_algorithm` property

Type: `"string"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("FIRST_MATCH")
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["FIRST_MATCH"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{stringvalidator.OneOf("FIRST_MATCH")}
```

- [rules](resources--secret_policy--reference--group-001.md#canonical-1111232303130011-1030301132230333-3012212320310123-1213022030003030-3202032333121311-2232110321210330-3200001233322211-0333333311023332): complete subsection reference.

<a id="canonical-2022212033013021-0303212312000103-3120333323130323-1013333002133302-0303320302211113-3303210020233310-3323201110013233-1010010331011331"></a>

<a id="canonical-0101223302201013-2303111110300013-1300110030323123-1203033132013123-2120330301123133-2130330131220000-1112303021332320-2210023122023000"></a>

#### `uid` property

Type: `"string"`. Computed.

<a id="canonical-2030022330023103-1023121231031320-3031213000220201-1123103133123101-2220012013312202-1110120202220111-0302310010100002-2110320321231220"></a>

### All schema paths for `xcsh_secret_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_f5xc` | [allow_f5xc](resources--secret_policy--reference--group-001.md#canonical-1300310323301312-2211032002222023-1023123132201302-1201333113010032-3102110203001212-0200131302303102-2301121231120033-2230323301332020) |
| `annotations` | [annotations](resources--secret_policy--reference--group-001.md#canonical-2210233301111102-2221012112213313-1031031122313011-3011023112001220-0333031123023002-0130133333221313-2101010133133110-2123303221111033) |
| `decrypt_cache_timeout` | [decrypt_cache_timeout](resources--secret_policy--reference--group-001.md#canonical-1101113121000313-3232113000332022-1132233320330321-3203200012110121-0203120303010001-0122223031020010-1033203021330321-2030133121033133) |
| `deletion_time` | [deletion_time](resources--secret_policy--reference--group-001.md#canonical-0203332013321131-3030003300022011-1323020111213132-1113210321233021-3123232312021231-0123000311113100-0202120110220213-1011022013101202) |
| `description` | [description](resources--secret_policy--reference--group-001.md#canonical-1300211223201123-3110232002232123-0033312221132110-1003222330023331-2001201032022023-0301210212211010-0012032233330213-3303001223320132) |
| `id` | [ID](resources--secret_policy--reference--group-001.md#canonical-0220332201332030-1002010212231032-0330000010200311-2233230211311131-2021011033231012-2011213202020330-0310201323230211-2010331221322320) |
| `labels` | [labels](resources--secret_policy--reference--group-001.md#canonical-0132201201220222-2020222310221223-3001221231131123-3321012022311010-3212320311223311-1133111003131130-3111110030310013-2101112132301000) |
| `marked_for_delete` | [marked_for_delete](resources--secret_policy--reference--group-001.md#canonical-3330102003322001-1200320222031303-2100020130121201-0121100332332123-0031020112211003-2031212033011132-0123333310220202-1203011312210321) |
| `name` | [name](resources--secret_policy--reference--group-001.md#canonical-1013310232311101-2213011332322002-2303031111020211-0222232102230131-1223302122111203-1013310030311002-2311033121001003-3011320023331222) |
| `namespace` | [namespace](resources--secret_policy--reference--group-001.md#canonical-3013303030313211-0112333003002133-1110030201310122-1011010020312303-0222123323033101-2320002113223221-0232203313120003-0110303311021121) |
| `rule_combining_algorithm` | [rule_combining_algorithm](resources--secret_policy--reference--group-001.md#canonical-1021110202020311-3122212110002211-3223203331313303-1331322321332100-1000223130200212-3202100202003133-3133202223202110-0122310300022202) |
| `rules` | [rules](resources--secret_policy--reference--group-001.md#canonical-2312033100022131-0220012211202103-1200010210132330-2032003030312321-3323303303233013-1011111121112023-3112332332333200-0203322313322030) |
| `rules.action` | [rules.action](resources--secret_policy--reference--group-001.md#canonical-2113303201111100-0233020203311123-1110021232013112-0013001331133112-3021301012322303-1313002301303031-1131312232331221-3013101103033020) |
| `rules.client_selector` | [rules.client_selector](resources--secret_policy--reference--group-001.md#canonical-3303001022233322-2203230021332200-2210213210223012-1120020133013213-0033101022321101-1031013102103331-3221103232002020-0101003010013003) |
| `rules.description` | [rules.description](resources--secret_policy--reference--group-001.md#canonical-0003102312123301-1222212212011102-0033003330131213-1201123230323101-0203233301210102-2303003003111320-0301133020323323-2322221120333213) |
| `rules.name` | [rules.name](resources--secret_policy--reference--group-001.md#canonical-2131102032133103-1213220330102231-3313013232131332-0003022302223133-3212100321301001-3211231303030020-0123322321131213-3311302100220331) |
| `uid` | [uid](resources--secret_policy--reference--group-001.md#canonical-2022212033013021-0303212312000103-3120333323130323-1013333002133302-0303320302211113-3303210020233310-3323201110013233-1010010331011331) |

<a id="canonical-1111232303130011-1030301132230333-3012212320310123-1213022030003030-3202032333121311-2232110321210330-3200001233322211-0333333311023332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules` properties

Breadcrumbs:

- [xcsh_secret_policy](../resources/secret_policy.md#canonical-1312021301020310-0131200100331301-3002003211003233-3010322032311123-0022212322233110-2113210022031113-0031130332112033-3113012311332111)
- [Property reference](resources--secret_policy--reference--group-001.md#canonical-1213313330302120-1100122222322313-2323333222220332-2211122300300103-3232201121133200-0113220013210033-2322130013023011-1132310023231200)
- rules

<a id="canonical-2312033100022131-0220012211202103-1200010210132330-2032003030312321-3323303303233013-1011111121112023-3112332332333200-0203322313322030"></a>

Type: `"list"`. Required.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{listvalidator.SizeBetween(1, 256)}
```

<a id="canonical-2021313200011330-0301311330220102-3323032131312210-1001012121130231-2010013221130133-1012332000033203-2211303100231002-3201103200201301"></a>

### Direct properties for `rules`

<a id="canonical-2113303201111100-0233020203311123-1110021232013112-0013001331133112-3021301012322303-1313002301303031-1131312232331221-3013101103033020"></a>

#### `rules.action` property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALLOW","DENY"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{stringvalidator.OneOf("ALLOW",
    "DENY")}
```

<a id="canonical-3303001022233322-2203230021332200-2210213210223012-1120020133013213-0033101022321101-1031013102103331-3221103232002020-0101003010013003"></a>

<a id="canonical-3212232122210113-0103003003102231-1222022221332302-1332223231120221-1223130103031332-2003101132000001-1213311302220123-2313120332211123"></a>

#### `rules.client_selector` property

Type: `["list", "string"]`. Optional.

Kubernetes selector expression. Omitted selector matches every client; use only for a final DENY
rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{listvalidator.SizeAtMost(1)}
```

<a id="canonical-0003102312123301-1222212212011102-0033003330131213-1201123230323101-0203233301210102-2303003003111320-0301133020323323-2322221120333213"></a>

<a id="canonical-1012302033221023-3310012003333313-0332031201320132-0301202332013000-0031321030100322-2231312232033103-0301201300213311-1013220322120032"></a>

#### `rules.description` property

Type: `"string"`. Optional.

<a id="canonical-2131102032133103-1213220330102231-3313013232131332-0003022302223133-3212100321301001-3211231303030020-0123322321131213-3311302100220331"></a>

<a id="canonical-2122132103011031-2333220030223312-1003023120013001-3002011100322032-3003230201222211-1212212012331312-3321210202230033-0132331311113200"></a>

#### `rules.name` property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{validators.NameValidator()}
```
