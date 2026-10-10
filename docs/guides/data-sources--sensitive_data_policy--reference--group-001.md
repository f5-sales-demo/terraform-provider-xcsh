---
page_title: "xcsh_sensitive_data_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_sensitive_data_policy reference."
---

# xcsh_sensitive_data_policy reference

<a id="canonical-2012220221230301-2110311100333312-1310211131030013-1101130223100303-2002300211120322-3233031313000120-3123212231312123-0033113022332033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-2321213032331000-2132213131031233-2100233222312121-0113013100321030-1011013210020203-2003033231120211-3313210223013012-0130113122220200)
- Property reference

<a id="canonical-3302020300010332-0230102230133311-3020332203203212-2321302311121212-3123213023133221-0331003210133212-1003320313233022-3203212132012201"></a>

### Direct properties for `xcsh_sensitive_data_policy`

<a id="canonical-0321030101222302-1200102023001112-3322010011221323-1010230300330010-0102332322310203-2023332211120302-3231013023221022-1133322330220101"></a>

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

<a id="canonical-1001331123233302-1220121112302123-1211013122223102-0223300321032021-3010232333102330-0210133010123002-1111021310333020-3221130001111011"></a>

<a id="canonical-3302221300221201-0330320302322320-3313211230233122-2322030132001110-0121320131000012-0132000131212021-0030131101002220-2312120022021313"></a>

#### `compliances` property

Type: `["list", "string"]`. Computed.

\[Enum:
GDPR|CCPA|PIPEDA|LGPD|DPA\_UK|PDPA\_SG|APPI|HIPAA|CPRA\_2023|CPA\_CO|SOC2|PCI\_DSS|ISO\_IEC\_27001|ISO\_IEC\_27701|EPRIVACY\_DIRECTIVE|GLBA|SOX\]
Select relevant compliance frameworks, such as GDPR, HIPAA, or PCI-DSS, to ensure monitoring under
your sensitive data discovery. Defaults to \`\[\]\`. Server applies default when omitted. Possible
values are \`GDPR\`, \`CCPA\`, \`PIPEDA\`, \`LGPD\`, \`DPA\_UK\`, \`PDPA\_SG\`, \`APPI\`, \`HIPAA\`,
\`CPRA\_2023\`, \`CPA\_CO\`, \`SOC2\`, \`PCI\_DSS\`, \`ISO\_IEC\_27001\`, \`ISO\_IEC\_27701\`,
\`EPRIVACY\_DIRECTIVE\`, \`GLBA\`, \`SOX\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 17,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 17,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "17",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "17",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [custom_data_types](data-sources--sensitive_data_policy--reference--group-001.md#canonical-2213331322133102-0313231331030120-0210321330120000-3203032221020333-2021113001131100-2132023212032102-1123203103322320-1230012010303112): complete subsection reference.

<a id="canonical-2011113322331121-1001131202212212-1012113121103322-1301021011123320-1132020303203031-3010200022002203-0031211201120002-1012312302210310"></a>

<a id="canonical-0313230033103230-1102213331013132-2331033320321131-2222003032032132-2313122302323231-2302330020312213-0000310020030010-0322112302112223"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the SensitiveDataPolicy.

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

<a id="canonical-1322312122300222-3102212232111112-1300301212222333-3333302202301110-3310201012201102-0301211002333110-1312223200031300-0220200331300332"></a>

<a id="canonical-1122220022203330-1320102211222022-1103302111000121-3310330230032030-0231123103020122-3222231310011000-0221123123003011-3003111031330332"></a>

#### `disabled_predefined_data_types` property

Type: `["list", "string"]`. Computed.

Select which pre-configured data types to disable, disabled data types will not be shown as
sensitive in the API discovery. Defaults to \`\[\]\`. Server applies default when omitted.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3301110310313000-0103332210021213-2013120313102321-2300310330313103-3330332312003120-0200031230222233-2001032211011312-3112102211101032"></a>

<a id="canonical-0212322200132212-3312030331110310-1232212300320220-0331231023210312-1102111003300323-1103230012303131-3333101212132112-2023020332300323"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0102322110120312-2321032013013002-3122212013120303-0031200223210202-3220023113123033-2132301211003333-2302220322333122-3202202300201003"></a>

<a id="canonical-2002233311302330-2210112232012323-1211200132322003-3203112223233200-2011300222120313-0102233211313121-1222000003000223-1312213302313112"></a>

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

<a id="canonical-3033031120232101-0121121113003321-3022021230323212-3200232213130220-3130033311301011-0031032111213231-1300203032032123-1302223223222212"></a>

<a id="canonical-0332100031001232-3332001302023323-2013003103203303-0123030000222030-2201131120230010-2001131212310323-2111113001203321-2132323322121130"></a>

#### `name` property

Type: `"string"`. Required.

Name of the SensitiveDataPolicy.

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

<a id="canonical-0233310023133220-2320101230231110-0210300201303212-3021220013322322-1323021033331112-2311101012323033-1103030223211023-3321003203123232"></a>

<a id="canonical-1330231122031010-1202123313122221-3322212032130322-2011212212303300-0213211300102033-2303232323320320-3203111110230333-1323123133233230"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the SensitiveDataPolicy exists.

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

<a id="canonical-3321111332011222-2010320100000121-0030121231203331-1121331302111332-0312032121122020-2102002030022233-2121002010112102-0022200201003101"></a>

### All schema paths for `xcsh_sensitive_data_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--sensitive_data_policy--reference--group-001.md#canonical-0321030101222302-1200102023001112-3322010011221323-1010230300330010-0102332322310203-2023332211120302-3231013023221022-1133322330220101) |
| `compliances` | [compliances](data-sources--sensitive_data_policy--reference--group-001.md#canonical-1001331123233302-1220121112302123-1211013122223102-0223300321032021-3010232333102330-0210133010123002-1111021310333020-3221130001111011) |
| `custom_data_types` | [custom_data_types](data-sources--sensitive_data_policy--reference--group-001.md#canonical-3003231100003123-3131223022302102-3131322021332202-0201312203312113-1331323202321320-3112233013321130-2012111303010202-3033110132233333) |
| `custom_data_types.custom_data_type_ref` | [custom_data_types.custom_data_type_ref](data-sources--sensitive_data_policy--reference--group-001.md#canonical-0223331020100332-0302010001130333-0102023031211202-1022013313121212-0312230203021113-0103230130202100-3320223333013123-2231301120112320) |
| `custom_data_types.custom_data_type_ref.name` | [custom_data_types.custom_data_type_ref.name](data-sources--sensitive_data_policy--reference--group-001.md#canonical-1200103003112013-0330221221021012-3321200132021330-3030301202032331-2313223202200010-3000031221130203-3232123300001200-3123331333103112) |
| `custom_data_types.custom_data_type_ref.namespace` | [custom_data_types.custom_data_type_ref.namespace](data-sources--sensitive_data_policy--reference--group-001.md#canonical-2101211133012132-2223311133323230-3212123011010230-2032101233212322-3233023210100111-0021213302301013-2123321102203223-3110321331220102) |
| `custom_data_types.custom_data_type_ref.tenant` | [custom_data_types.custom_data_type_ref.tenant](data-sources--sensitive_data_policy--reference--group-001.md#canonical-0302002113012212-3310301211132311-2112130332101031-1203022210333031-0011311203312221-1332101031012213-0112333031311012-1113310022231322) |
| `description` | [description](data-sources--sensitive_data_policy--reference--group-001.md#canonical-2011113322331121-1001131202212212-1012113121103322-1301021011123320-1132020303203031-3010200022002203-0031211201120002-1012312302210310) |
| `disabled_predefined_data_types` | [disabled_predefined_data_types](data-sources--sensitive_data_policy--reference--group-001.md#canonical-1322312122300222-3102212232111112-1300301212222333-3333302202301110-3310201012201102-0301211002333110-1312223200031300-0220200331300332) |
| `id` | [ID](data-sources--sensitive_data_policy--reference--group-001.md#canonical-3301110310313000-0103332210021213-2013120313102321-2300310330313103-3330332312003120-0200031230222233-2001032211011312-3112102211101032) |
| `labels` | [labels](data-sources--sensitive_data_policy--reference--group-001.md#canonical-0102322110120312-2321032013013002-3122212013120303-0031200223210202-3220023113123033-2132301211003333-2302220322333122-3202202300201003) |
| `name` | [name](data-sources--sensitive_data_policy--reference--group-001.md#canonical-3033031120232101-0121121113003321-3022021230323212-3200232213130220-3130033311301011-0031032111213231-1300203032032123-1302223223222212) |
| `namespace` | [namespace](data-sources--sensitive_data_policy--reference--group-001.md#canonical-0233310023133220-2320101230231110-0210300201303212-3021220013322322-1323021033331112-2311101012323033-1103030223211023-3321003203123232) |

<a id="canonical-2213331322133102-0313231331030120-0210321330120000-3203032221020333-2021113001131100-2132023212032102-1123203103322320-1230012010303112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_data_types` properties

Breadcrumbs:

- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-2321213032331000-2132213131031233-2100233222312121-0113013100321030-1011013210020203-2003033231120211-3313210223013012-0130113122220200)
- [Property reference](data-sources--sensitive_data_policy--reference--group-001.md#canonical-2012220221230301-2110311100333312-1310211131030013-1101130223100303-2002300211120322-3233031313000120-3123212231312123-0033113022332033)
- custom_data_types

<a id="canonical-3003231100003123-3131223022302102-3131322021332202-0201312203312113-1331323202321320-3112233013321130-2012111303010202-3033110132233333"></a>

Type: `"list"`. Computed.

Select your custom data types to be monitored in the API discovery. Defaults to \`\[\]\`. Server
applies default when omitted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0131023103111100-0313013130001112-1110121102202301-2301303311020012-0002032203333201-0101232322210102-2002101012133132-3323302120020113"></a>

### Direct properties for `custom_data_types`

- [custom_data_type_ref](data-sources--sensitive_data_policy--reference--group-001.md#canonical-1312000322012203-0023222002023230-3102330000112032-0111231230300202-0130001330011121-1213010001233203-0031232100120211-3302111011220100): complete subsection reference.

<a id="canonical-1312000322012203-0023222002023230-3102330000112032-0111231230300202-0130001330011121-1213010001233203-0031232100120211-3302111011220100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_data_types.custom_data_type_ref` properties

Breadcrumbs:

- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-2321213032331000-2132213131031233-2100233222312121-0113013100321030-1011013210020203-2003033231120211-3313210223013012-0130113122220200)
- [Property reference](data-sources--sensitive_data_policy--reference--group-001.md#canonical-2012220221230301-2110311100333312-1310211131030013-1101130223100303-2002300211120322-3233031313000120-3123212231312123-0033113022332033)
- [custom_data_types](data-sources--sensitive_data_policy--reference--group-001.md#canonical-2213331322133102-0313231331030120-0210321330120000-3203032221020333-2021113001131100-2132023212032102-1123203103322320-1230012010303112)
- custom_data_types.custom_data_type_ref

<a id="canonical-0223331020100332-0302010001130333-0102023031211202-1022013313121212-0312230203021113-0103230130202100-3320223333013123-2231301120112320"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3233110300211313-3211013033012113-0310321033201102-2312220323220212-1330230210221002-1312020330130002-2331221103031312-0033131311203013"></a>

### Direct properties for `custom_data_types.custom_data_type_ref`

<a id="canonical-1200103003112013-0330221221021012-3321200132021330-3030301202032331-2313223202200010-3000031221130203-3232123300001200-3123331333103112"></a>

#### `custom_data_types.custom_data_type_ref.name` property

Type: `"string"`. Computed.

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

<a id="canonical-2101211133012132-2223311133323230-3212123011010230-2032101233212322-3233023210100111-0021213302301013-2123321102203223-3110321331220102"></a>

<a id="canonical-2021221033332221-1032202021013103-0121313311223122-2101213300222010-0002101332020331-1113231201220300-1121230100101033-2200313002223230"></a>

#### `custom_data_types.custom_data_type_ref.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-0302002113012212-3310301211132311-2112130332101031-1203022210333031-0011311203312221-1332101031012213-0112333031311012-1113310022231322"></a>

<a id="canonical-2123203310202113-0323311321222333-1121221232333210-1302211131312201-1223102303321030-3003311130121302-2000222032013122-3311233021111131"></a>

#### `custom_data_types.custom_data_type_ref.tenant` property

Type: `"string"`. Computed.

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
