---
page_title: "xcsh_container_registry reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_container_registry reference."
---

# xcsh_container_registry reference

<a id="canonical-1032120112320001-2222011121203202-3330011001232203-0333133013110231-2202321322121232-1121012101311030-0100022220223123-1211111333222203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md#canonical-3121031110330001-1021233032301331-1200313112012333-2030023322330002-2021231201222110-1023032102203110-1103221200322013-1102022332013332)
- Property reference

<a id="canonical-0303303223233032-0012322313033010-1122233210200033-0312202332222103-1211310300233321-2102133213220230-3202120220110113-0222021330132012"></a>

### Direct properties for `xcsh_container_registry`

<a id="canonical-3300011230110313-1231123233100320-2301322032321310-1302023200120022-3330301302033102-1110302211323333-1320020123111023-0212102201302311"></a>

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

<a id="canonical-1001013330222220-2112112332221030-1132112212023013-2133122212323120-3302230202001220-3013030202013213-3030231020132113-0310332123222131"></a>

<a id="canonical-2310222020300310-1212230212100031-0233111030012003-3132332232222031-1201333132132011-1032202233201102-3000020303333131-2230133121320303"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1121131023112302-2032122211313220-0110101223110121-2122101231001222-0103030110231020-0111033222313223-2202200323112033-1032322312230300"></a>

<a id="canonical-0333000031313313-1231223330332220-3203002003030033-2220201123022232-0123223020111233-2313322320322231-2122113110003201-1010032120120302"></a>

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

<a id="canonical-0121311320022323-2103303000213310-3302120311203201-0130030103201031-0331000203030312-2022321211310332-1321311121111031-2303032000112020"></a>

<a id="canonical-1101212003112133-2002120012330321-0333023313310111-3202113010022223-2123111302332000-0203222330111102-2120000232221000-0000210120231333"></a>

#### `email` property

Type: `"string"`. Optional, Computed.

Email. Email used for the registry.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "email",
    "formatDescription": "RFC 5322 email address, max 254 characters",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
    "validation": {
      "rfc": "RFC 5322"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  }
}
```

<a id="canonical-3231013003001012-0223000232030123-2231211230100120-2020230300033111-3000201011122110-3203112013110013-0200313232331202-3122311122213010"></a>

<a id="canonical-2020120220130322-2013331332101113-0223231131220021-2300332202102220-3301132330132001-0003303023032011-2113231111310332-1011121213122032"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1012013112222320-3213011222233123-1230321112032213-2111211002131001-1121213003232300-1213312002120000-2012201130122111-0120301012010022"></a>

<a id="canonical-0311021130020001-0100003313310232-3030233011101100-2310112131002330-3233101021011213-2122033301213203-2313011032131132-0230111010000330"></a>

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

<a id="canonical-0210220221010113-2332133023033221-1301110323311212-1012213202130222-2303132203100332-0110122121322100-3232333121033231-1023312110331331"></a>

<a id="canonical-2223123133330123-3211103221332330-1323310110100211-2122103121303022-0222000302013121-3213221312233211-1330212213133020-1012113320202133"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Container Registry. Must be unique within the namespace.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3102331233222000-0130330130203031-0031113211003102-3201311212013000-3322003311031330-3322131023001323-1211033331102231-2123100233230030"></a>

<a id="canonical-3033100013023130-2120012333100211-0312111120133033-1101331020011130-1102331212100022-0121012302022330-0301000302130223-0310032020202100"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Container Registry is created.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [password](resources--container_registry--reference--group-001.md#canonical-2333301020100202-1213333012031101-2112222013010032-1133212101000011-3131130330211132-2113313300202133-3232233203032013-0020230201003030): complete subsection reference.

<a id="canonical-3123222222313213-2210222112130222-0221123112121002-2130213112312323-0002312210122210-1300211001230102-2331032012001102-3312122011000112"></a>

<a id="canonical-2023233033312032-3003022120210020-3330112122103110-3002233231001023-3233230133030313-2233100132100301-2223120203303201-1102203231010330"></a>

#### `registry` property

Type: `"string"`. Required.

Fully qualified name of the registry login server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [timeouts](resources--container_registry--reference--group-001.md#canonical-2133111212322010-1310302120021302-3113221011233131-2210220003322330-1330202222220030-2022121210313202-0112122112231220-1030210111021010): complete subsection reference.

<a id="canonical-0133203131221113-0012033003200203-3320003112100213-2010020302200000-1011311030031310-2333301222222312-3331111113020331-3010113131302130"></a>

<a id="canonical-3202222201200200-1232001113203200-0201122200023023-1132211210210312-3033300000220201-3030033121311302-3103022223113100-1302130122013010"></a>

#### `user_name` property

Type: `"string"`. Required.

username. Username used to access the registry.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1200233101313320-0003323021102233-0023221202220023-3303100010021301-1211232211002313-3101133012333123-3213033133213030-2310122313012110"></a>

### All schema paths for `xcsh_container_registry`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--container_registry--reference--group-001.md#canonical-3300011230110313-1231123233100320-2301322032321310-1302023200120022-3330301302033102-1110302211323333-1320020123111023-0212102201302311) |
| `description` | [description](resources--container_registry--reference--group-001.md#canonical-1001013330222220-2112112332221030-1132112212023013-2133122212323120-3302230202001220-3013030202013213-3030231020132113-0310332123222131) |
| `disable` | [disable](resources--container_registry--reference--group-001.md#canonical-1121131023112302-2032122211313220-0110101223110121-2122101231001222-0103030110231020-0111033222313223-2202200323112033-1032322312230300) |
| `email` | [email](resources--container_registry--reference--group-001.md#canonical-0121311320022323-2103303000213310-3302120311203201-0130030103201031-0331000203030312-2022321211310332-1321311121111031-2303032000112020) |
| `id` | [ID](resources--container_registry--reference--group-001.md#canonical-3231013003001012-0223000232030123-2231211230100120-2020230300033111-3000201011122110-3203112013110013-0200313232331202-3122311122213010) |
| `labels` | [labels](resources--container_registry--reference--group-001.md#canonical-1012013112222320-3213011222233123-1230321112032213-2111211002131001-1121213003232300-1213312002120000-2012201130122111-0120301012010022) |
| `name` | [name](resources--container_registry--reference--group-001.md#canonical-0210220221010113-2332133023033221-1301110323311212-1012213202130222-2303132203100332-0110122121322100-3232333121033231-1023312110331331) |
| `namespace` | [namespace](resources--container_registry--reference--group-001.md#canonical-3102331233222000-0130330130203031-0031113211003102-3201311212013000-3322003311031330-3322131023001323-1211033331102231-2123100233230030) |
| `password` | [password](resources--container_registry--reference--group-001.md#canonical-1230021313131110-1003111003002230-3011032330332211-3000311320023003-0332201003320232-1021230301003030-0100031320320003-3011113110322222) |
| `password.blindfold_secret_info` | [password.blindfold_secret_info](resources--container_registry--reference--group-001.md#canonical-3003221031223321-0231223322222100-2112222220111003-3111331213323123-1132300233320202-0131221102301330-0220332021132102-3102212020202333) |
| `password.blindfold_secret_info.decryption_provider` | [password.blindfold_secret_info.decryption_provider](resources--container_registry--reference--group-001.md#canonical-0021213120223121-2200111033331022-0212220210100221-0302200023202133-3012100322010300-1310130213233313-0103321003023332-2112033003022130) |
| `password.blindfold_secret_info.location` | [password.blindfold_secret_info.location](resources--container_registry--reference--group-001.md#canonical-3210210300302300-1032020030002101-3011312310101132-3103200303023130-2303103121113320-2123120010323302-1202221001330021-0202001323001101) |
| `password.blindfold_secret_info.store_provider` | [password.blindfold_secret_info.store_provider](resources--container_registry--reference--group-001.md#canonical-2312200120132313-3121103101130201-0303313230032230-1300101300220223-3202001133222110-0212302113133200-3301311301111102-1300313210333322) |
| `password.clear_secret_info` | [password.clear_secret_info](resources--container_registry--reference--group-001.md#canonical-2223301100033101-0331111020302121-2332210110120311-1133100000111201-1133221310120210-3331233210023122-3313121232101300-3211132301022031) |
| `password.clear_secret_info.provider_ref` | [password.clear_secret_info.provider_ref](resources--container_registry--reference--group-001.md#canonical-0322101123101132-1323300020012002-2011223333320103-2300202223133320-2002113033122013-2012223322330023-2331301112313001-3312313002032321) |
| `password.clear_secret_info.url` | [password.clear_secret_info.url](resources--container_registry--reference--group-001.md#canonical-2120301332003223-0220230232222310-0010132133213030-2133302211130300-1201103113200032-3000010112033131-1211221332132332-0011300330322202) |
| `registry` | [registry](resources--container_registry--reference--group-001.md#canonical-3123222222313213-2210222112130222-0221123112121002-2130213112312323-0002312210122210-1300211001230102-2331032012001102-3312122011000112) |
| `timeouts` | [timeouts](resources--container_registry--reference--group-001.md#canonical-2010130023223322-1201022031021113-1300033112130100-1003110200201303-0000211301133023-1332021031320122-3311223333122002-3301231103210201) |
| `timeouts.create` | [timeouts.create](resources--container_registry--reference--group-001.md#canonical-3000223000121323-2101222111030233-0200232331320113-2022333003232201-1020122212123201-1010020130200203-2311121333301113-3200021210213000) |
| `timeouts.delete` | [timeouts.delete](resources--container_registry--reference--group-001.md#canonical-3112033113202220-2021213101122310-3211222110302312-0302302232032200-0220123113020310-2110202010312221-0213020133001011-0220131302202001) |
| `timeouts.read` | [timeouts.read](resources--container_registry--reference--group-001.md#canonical-1300222122003301-1100003232232003-2003113121230032-2332031223333100-0322232030033000-2110213112323130-2301312123130121-2232111230102310) |
| `timeouts.update` | [timeouts.update](resources--container_registry--reference--group-001.md#canonical-0211320213130233-1110212233133001-2323100221222320-1303210020010230-3330023301020311-0221030213313333-0022012201000131-2333022111332113) |
| `user_name` | [user_name](resources--container_registry--reference--group-001.md#canonical-0133203131221113-0012033003200203-3320003112100213-2010020302200000-1011311030031310-2333301222222312-3331111113020331-3010113131302130) |

<a id="canonical-2333301020100202-1213333012031101-2112222013010032-1133212101000011-3131130330211132-2113313300202133-3232233203032013-0020230201003030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `password` properties

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md#canonical-3121031110330001-1021233032301331-1200313112012333-2030023322330002-2021231201222110-1023032102203110-1103221200322013-1102022332013332)
- [Property reference](resources--container_registry--reference--group-001.md#canonical-1032120112320001-2222011121203202-3330011001232203-0333133013110231-2202321322121232-1121012101311030-0100022220223123-1211111333222203)
- password

<a id="canonical-1230021313131110-1003111003002230-3011032330332211-3000311320023003-0332201003320232-1021230301003030-0100031320320003-3011113110322222"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3320321310221120-1301133022300202-2201220023233103-2300200011100220-0303322120003022-0003130011133200-3302010310020311-0212213130323232"></a>

### Direct properties for `password`

- [blindfold_secret_info](resources--container_registry--reference--group-001.md#canonical-2101311200200132-0322131031022112-2323212022001212-1211312112131020-3021321122032030-1310232000302011-2032203210123233-2031101031012201): complete subsection reference.

- [clear_secret_info](resources--container_registry--reference--group-001.md#canonical-0023123110203112-3230220012112200-0101120212003331-1313320210323233-2123130203120002-1232031210203120-0222302332022101-3103123121223003): complete subsection reference.

<a id="canonical-2101311200200132-0322131031022112-2323212022001212-1211312112131020-3021321122032030-1310232000302011-2032203210123233-2031101031012201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md#canonical-3121031110330001-1021233032301331-1200313112012333-2030023322330002-2021231201222110-1023032102203110-1103221200322013-1102022332013332)
- [Property reference](resources--container_registry--reference--group-001.md#canonical-1032120112320001-2222011121203202-3330011001232203-0333133013110231-2202321322121232-1121012101311030-0100022220223123-1211111333222203)
- [password](resources--container_registry--reference--group-001.md#canonical-2333301020100202-1213333012031101-2112222013010032-1133212101000011-3131130330211132-2113313300202133-3232233203032013-0020230201003030)
- password.blindfold_secret_info

<a id="canonical-3003221031223321-0231223322222100-2112222220111003-3111331213323123-1132300233320202-0131221102301330-0220332021132102-3102212020202333"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0202001121001323-3303221033321023-1111033200331101-1013001130033021-2320031223203031-0133323022102232-2010222121120022-3131013221201122"></a>

### Direct properties for `password.blindfold_secret_info`

<a id="canonical-0021213120223121-2200111033331022-0212220210100221-0302200023202133-3012100322010300-1310130213233313-0103321003023332-2112033003022130"></a>

#### `password.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3210210300302300-1032020030002101-3011312310101132-3103200303023130-2303103121113320-2123120010323302-1202221001330021-0202001323001101"></a>

<a id="canonical-0320303312331233-2302311120212233-1330322011133031-1310023333323113-2011130322030131-0032332303132220-0202221103102313-0210130233113311"></a>

#### `password.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2312200120132313-3121103101130201-0303313230032230-1300101300220223-3202001133222110-0212302113133200-3301311301111102-1300313210333322"></a>

<a id="canonical-3301013322000111-0123332323221313-1221111103010021-2102213111000300-3231302010311003-0021110223132331-0332010310023122-3223013032332013"></a>

#### `password.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0023123110203112-3230220012112200-0101120212003331-1313320210323233-2123130203120002-1232031210203120-0222302332022101-3103123121223003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md#canonical-3121031110330001-1021233032301331-1200313112012333-2030023322330002-2021231201222110-1023032102203110-1103221200322013-1102022332013332)
- [Property reference](resources--container_registry--reference--group-001.md#canonical-1032120112320001-2222011121203202-3330011001232203-0333133013110231-2202321322121232-1121012101311030-0100022220223123-1211111333222203)
- [password](resources--container_registry--reference--group-001.md#canonical-2333301020100202-1213333012031101-2112222013010032-1133212101000011-3131130330211132-2113313300202133-3232233203032013-0020230201003030)
- password.clear_secret_info

<a id="canonical-2223301100033101-0331111020302121-2332210110120311-1133100000111201-1133221310120210-3331233210023122-3313121232101300-3211132301022031"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3010332311100330-2020320003031311-1232222002112223-3232022102302230-2111211330122332-1000211331023012-3100313323002033-0123030121302022"></a>

### Direct properties for `password.clear_secret_info`

<a id="canonical-0322101123101132-1323300020012002-2011223333320103-2300202223133320-2002113033122013-2012223322330023-2331301112313001-3312313002032321"></a>

#### `password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2120301332003223-0220230232222310-0010132133213030-2133302211130300-1201103113200032-3000010112033131-1211221332132332-0011300330322202"></a>

<a id="canonical-2022201032102023-1131121210310200-3322001133232200-0013331210233102-0131122021023203-2001201033033232-0121232300232323-1232213133302230"></a>

#### `password.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2133111212322010-1310302120021302-3113221011233131-2210220003322330-1330202222220030-2022121210313202-0112122112231220-1030210111021010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md#canonical-3121031110330001-1021233032301331-1200313112012333-2030023322330002-2021231201222110-1023032102203110-1103221200322013-1102022332013332)
- [Property reference](resources--container_registry--reference--group-001.md#canonical-1032120112320001-2222011121203202-3330011001232203-0333133013110231-2202321322121232-1121012101311030-0100022220223123-1211111333222203)
- timeouts

<a id="canonical-2010130023223322-1201022031021113-1300033112130100-1003110200201303-0000211301133023-1332021031320122-3311223333122002-3301231103210201"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203233032013331-0033022223122323-0130203023133320-3321213000232222-0010030202022331-1220330131103110-1131231123010222-1201332200121320"></a>

### Direct properties for `timeouts`

<a id="canonical-3000223000121323-2101222111030233-0200232331320113-2022333003232201-1020122212123201-1010020130200203-2311121333301113-3200021210213000"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3112033113202220-2021213101122310-3211222110302312-0302302232032200-0220123113020310-2110202010312221-0213020133001011-0220131302202001"></a>

<a id="canonical-2301103012200133-0013010302121230-1222320222223233-1302022112220032-0100313233220232-0132103113020120-2101301321013113-2321110020100113"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1300222122003301-1100003232232003-2003113121230032-2332031223333100-0322232030033000-2110213112323130-2301312123130121-2232111230102310"></a>

<a id="canonical-1002211120001222-1031031100022113-2011311021130121-2110230220113203-3102121200002312-0231210012012332-0200233133123121-0213112333200202"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0211320213130233-1110212233133001-2323100221222320-1303210020010230-3330023301020311-0221030213313333-0022012201000131-2333022111332113"></a>

<a id="canonical-0300033011010010-1023220031032031-0113001221110000-2223022023032132-1101203310133301-1203310033312133-3231222113310032-0220200011033021"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
