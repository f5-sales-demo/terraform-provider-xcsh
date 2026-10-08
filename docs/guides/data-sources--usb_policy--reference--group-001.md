---
page_title: "xcsh_usb_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_usb_policy reference."
---

# xcsh_usb_policy reference

<a id="canonical-1000132301302331-1210321310212233-2332001022201121-3133211123011103-2123333120333133-1000202023301002-2333222023302121-1001000003033000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_usb_policy](../data-sources/usb_policy.md#canonical-3133321203222121-0333232330020202-2311221201233121-1211132233302312-0010122132101103-2031031312101012-3000023221320113-0202220320111000)
- Property reference

<a id="canonical-1313033231100122-0331202010330211-2221031313221213-0312302201210130-2102101203221232-0111332222032011-3031133332032313-2100211120213032"></a>

### Direct properties for `xcsh_usb_policy`

- [allowed_devices](data-sources--usb_policy--reference--group-001.md#canonical-1222313110031021-0033310100131201-3003112111300200-1012320303311210-3122230300102333-2330103233013011-0311111233223203-0223210002011200): complete subsection reference.

<a id="canonical-1212213332131130-1002301221031130-0203130121231011-3101112230220022-3221203021030221-3031310232333222-3013131010120133-2121121000301032"></a>

<a id="canonical-0331122122310002-0103201210313031-0300132231102031-2112212111133231-3310211001010300-2301110112221200-2011223302012000-0211001333012233"></a>

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

<a id="canonical-2130312212020301-1212112111110200-2320003032012322-0203001202101000-0101220013322223-2333020103312203-2021123311110323-1230230133201103"></a>

<a id="canonical-2021221032233221-0323302122212331-0300122300033303-3201203023323321-1331213032333020-1103101231203210-3113201130001021-2321132212330101"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the UsbPolicy.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3201321103321113-3002020201320012-0233223213000000-1102311100112231-0132320010223103-2332302033212330-3132111103300310-2301122011000323"></a>

<a id="canonical-2120133003232213-3200021133301223-2310333112011132-1133211320232100-3201203230330030-3320321031122031-1121011133231103-2132210132223001"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3103210211303100-1100332220221332-2303303031312110-0300031231301221-3131131031211222-3211220121333201-3121132001113321-3330223113322312"></a>

<a id="canonical-1231220033103200-2003232230311012-0030321030001122-1200012000113211-0002022213011232-0331122112331013-1213023130031323-1121003200330032"></a>

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

<a id="canonical-1220320022001210-0311233200131210-2333133310302310-0110103132133231-0001130212010000-3222221102132030-0211233320012322-2331200313132001"></a>

<a id="canonical-3123202113123202-2130320022113333-0223202122020212-0233221010312120-1033012001310003-0021002233011231-2223021100103101-2311023312320322"></a>

#### `name` property

Type: `"string"`. Required.

Name of the UsbPolicy.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2211113121131023-3010212330232230-2232311210220312-1013323201212333-3223131120022111-3101223213322202-2230000032103323-1320113212132231"></a>

<a id="canonical-1033033302130113-1213312223131213-2010120312232332-3210000301001132-1311323112300233-2322323311030000-1122130301231332-0222022320011331"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the UsbPolicy exists.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3222122123113211-3000202300211100-0030331223130301-1132002303031333-1302003210032323-3031311303331031-3232101232202133-3311300333231333"></a>

### All schema paths for `xcsh_usb_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allowed_devices` | [allowed_devices](data-sources--usb_policy--reference--group-001.md#canonical-3122013323320001-2333111023333130-0332020111032002-0223013122100303-2123023003332110-0312032331300212-0000021302132023-1003121321332002) |
| `allowed_devices.b_device_class` | [allowed_devices.b_device_class](data-sources--usb_policy--reference--group-001.md#canonical-2022330311221301-2220231032200132-1230030030201003-0103021002332220-0202203322331032-2112000301030032-1232101001220231-1313123331011221) |
| `allowed_devices.b_device_protocol` | [allowed_devices.b_device_protocol](data-sources--usb_policy--reference--group-001.md#canonical-1133020020201122-1010333230320100-1033000301122030-0131211322113311-0323121132232311-1113012230222102-2112312213320110-0003112001330210) |
| `allowed_devices.b_device_sub_class` | [allowed_devices.b_device_sub_class](data-sources--usb_policy--reference--group-001.md#canonical-2333222201121130-2231310211023233-1200322021012203-3200323011300302-1311001233023331-3033020320012232-1210332322310202-2023121030332231) |
| `allowed_devices.i_serial` | [allowed_devices.i_serial](data-sources--usb_policy--reference--group-001.md#canonical-0133013312230000-1213310021020321-1333111031201221-0313233230032233-3031223122202300-1202332303013313-3130102102120030-0111330022133131) |
| `allowed_devices.id_product` | [allowed_devices.id_product](data-sources--usb_policy--reference--group-001.md#canonical-0111323133121030-1311232003321002-1203233323122112-2320130211012302-0013133333201102-2001211211122010-3220221302133220-3221002022023033) |
| `allowed_devices.id_vendor` | [allowed_devices.id_vendor](data-sources--usb_policy--reference--group-001.md#canonical-1301311301110130-2130131003331103-2321212111313021-3223310002221003-3030211213212223-3022200311320311-3200002001031112-3211322123012033) |
| `annotations` | [annotations](data-sources--usb_policy--reference--group-001.md#canonical-1212213332131130-1002301221031130-0203130121231011-3101112230220022-3221203021030221-3031310232333222-3013131010120133-2121121000301032) |
| `description` | [description](data-sources--usb_policy--reference--group-001.md#canonical-2130312212020301-1212112111110200-2320003032012322-0203001202101000-0101220013322223-2333020103312203-2021123311110323-1230230133201103) |
| `id` | [ID](data-sources--usb_policy--reference--group-001.md#canonical-3201321103321113-3002020201320012-0233223213000000-1102311100112231-0132320010223103-2332302033212330-3132111103300310-2301122011000323) |
| `labels` | [labels](data-sources--usb_policy--reference--group-001.md#canonical-3103210211303100-1100332220221332-2303303031312110-0300031231301221-3131131031211222-3211220121333201-3121132001113321-3330223113322312) |
| `name` | [name](data-sources--usb_policy--reference--group-001.md#canonical-1220320022001210-0311233200131210-2333133310302310-0110103132133231-0001130212010000-3222221102132030-0211233320012322-2331200313132001) |
| `namespace` | [namespace](data-sources--usb_policy--reference--group-001.md#canonical-2211113121131023-3010212330232230-2232311210220312-1013323201212333-3223131120022111-3101223213322202-2230000032103323-1320113212132231) |

<a id="canonical-1222313110031021-0033310100131201-3003112111300200-1012320303311210-3122230300102333-2330103233013011-0311111233223203-0223210002011200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allowed_devices` properties

Breadcrumbs:

- [xcsh_usb_policy](../data-sources/usb_policy.md#canonical-3133321203222121-0333232330020202-2311221201233121-1211132233302312-0010122132101103-2031031312101012-3000023221320113-0202220320111000)
- [Property reference](data-sources--usb_policy--reference--group-001.md#canonical-1000132301302331-1210321310212233-2332001022201121-3133211123011103-2123333120333133-1000202023301002-2333222023302121-1001000003033000)
- allowed_devices

<a id="canonical-3122013323320001-2333111023333130-0332020111032002-0223013122100303-2123023003332110-0312032331300212-0000021302132023-1003121321332002"></a>

Type: `"list"`. Computed.

Allowed USB devices. List of allowed USB devices.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.message.required_one_nonzero_field": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.message.required_one_nonzero_field": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2202011303222103-3000012230003131-3231230310330010-3213222203010230-1202303131001121-1201231210103100-2100231220102020-3231130020203030"></a>

### Direct properties for `allowed_devices`

<a id="canonical-2022330311221301-2220231032200132-1230030030201003-0103021002332220-0202203322331032-2112000301030032-1232101001220231-1313123331011221"></a>

#### `allowed_devices.b_device_class` property

Type: `"string"`. Computed.

Class. The class of this device.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1133020020201122-1010333230320100-1033000301122030-0131211322113311-0323121132232311-1113012230222102-2112312213320110-0003112001330210"></a>

<a id="canonical-3131222030333123-3331223031001333-3021133311113313-3330210322310132-2231200013113230-2322113110120302-0113023120032310-0200302123113320"></a>

#### `allowed_devices.b_device_protocol` property

Type: `"string"`. Computed.

The protocol (within the subclass) of this device.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2333222201121130-2231310211023233-1200322021012203-3200323011300302-1311001233023331-3033020320012232-1210332322310202-2023121030332231"></a>

<a id="canonical-0010221200003010-3310310020300020-3232202012012101-2333230233133001-3323213031112230-0230001122032211-3310020203331333-1012200031212121"></a>

#### `allowed_devices.b_device_sub_class` property

Type: `"string"`. Computed.

The subclass (within the class) of this device.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0133013312230000-1213310021020321-1333111031201221-0313233230032233-3031223122202300-1202332303013313-3130102102120030-0111330022133131"></a>

<a id="canonical-3201323020010313-2301230001001212-1122322133103322-1313200120331021-2130103020320132-1013333322333131-2012323212102113-0221110102023010"></a>

#### `allowed_devices.i_serial` property

Type: `"string"`. Computed.

Index of Serial Number String Descriptor.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0111323133121030-1311232003321002-1203233323122112-2320130211012302-0013133333201102-2001211211122010-3220221302133220-3221002022023033"></a>

<a id="canonical-1203232313311110-1230200312123100-0103022101321322-3331230210331013-0122020212120110-1320222212020113-1322233330331211-2303322330102330"></a>

#### `allowed_devices.id_product` property

Type: `"string"`. Computed.

Product ID (Assigned by Manufacturer) in hex.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1301311301110130-2130131003331103-2321212111313021-3223310002221003-3030211213212223-3022200311320311-3200002001031112-3211322123012033"></a>

<a id="canonical-2202301132123310-2330211011001030-0021221220121200-3320320330230013-1003102220101303-1320213020031003-1210033220322223-3232200111000213"></a>

#### `allowed_devices.id_vendor` property

Type: `"string"`. Computed.

Vendor ID. Vendor ID (Assigned by USB Org) in hex.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
