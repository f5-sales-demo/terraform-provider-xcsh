---
page_title: "xcsh_certificate_chain reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate_chain reference."
---

# xcsh_certificate_chain reference

<a id="canonical-3121100312232202-0231200300013012-2200330002333301-1222103120032112-3203332013220113-1331112100302301-1010131310310303-2011022230112320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_certificate_chain](../data-sources/certificate_chain.md#canonical-3032033102223330-1001120313322113-3332030202222110-0322003202122313-3333310012310100-1102220222021332-3222113133323101-1133210103103321)
- Property reference

<a id="canonical-2103301130303133-3033101321301110-1123001230202001-2023113220300213-1302320030231333-2103021213233222-3022301020030113-1200131001302201"></a>

### Direct properties for `xcsh_certificate_chain`

<a id="canonical-3221030233332320-3223111233223020-0311213113010110-0110101012230223-3111013223132113-1301021000310120-3311031222210102-1301110321032210"></a>

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

<a id="canonical-3023102130001130-2200000013201321-2302302130303100-3121012000320322-2332002032001223-1030010213202301-2210313103031110-2331203302030130"></a>

<a id="canonical-1022200310330311-0310310102232120-3121301310022113-0001123320021123-1201013332330312-1103333002203122-1323103030333032-2010203310222101"></a>

#### `certificate_url` property

Type: `"string"`. Computed.

Certificate chain is the list of intermediate certificates in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.intermediate_certificate_chain_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.intermediate_certificate_chain_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3230313022023313-2131201013032210-3211001301103033-0331002300012231-3202321201323330-2220002211103031-0020000201033200-1320313223022322"></a>

<a id="canonical-1300101100222321-2222012303331113-2100212002200202-3120022231113333-2100130220313001-2312102133112230-3321232202330221-2211002333103122"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the CertificateChain.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0103013132121113-3103232113200303-3001203221021333-2320011033033222-1021321020012302-2211321321110331-2313103203301302-1221202223110332"></a>

<a id="canonical-2220320230301023-3010010100001102-1201003203310133-2130203213011320-3233133001301321-1022231113231212-2121132021313121-2303003231100213"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2120130332010202-2031100101102231-0133123303313320-0210203202113111-2032113231310303-0003220121302000-2102021113331211-3032111011022003"></a>

<a id="canonical-3212332330212023-2203132100230130-3133320201023212-3012031320323321-2332010302120221-2302331230203200-2313312000312312-3202311200310133"></a>

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

<a id="canonical-1102120310202110-2122103210320031-3110221002001331-2100230221201210-0023023312310111-3131021322003123-0012013122203012-1323013130330211"></a>

<a id="canonical-0211112100210331-2211021311211313-0202312102110001-0220100122220200-3211110021102000-1313121001132131-1130322301110200-2331130102102322"></a>

#### `name` property

Type: `"string"`. Required.

Name of the CertificateChain.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0032321121222320-2323232222102223-3101001110220200-2303231023330213-2102203333130202-2012003131112321-0131301230201203-1112020123332102"></a>

<a id="canonical-1231113300232030-1200322122033202-3221012332331133-2010010023020200-1130022211023200-1221222313133011-1302033013131100-2011231031102302"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the CertificateChain exists.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0332011302301203-1103112011011010-2100220212220103-2311300211100133-3230123312301032-3232203012102303-1132112121202232-1302303132110021"></a>

### All schema paths for `xcsh_certificate_chain`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--certificate_chain--reference--group-001.md#canonical-3221030233332320-3223111233223020-0311213113010110-0110101012230223-3111013223132113-1301021000310120-3311031222210102-1301110321032210) |
| `certificate_url` | [certificate_url](data-sources--certificate_chain--reference--group-001.md#canonical-3023102130001130-2200000013201321-2302302130303100-3121012000320322-2332002032001223-1030010213202301-2210313103031110-2331203302030130) |
| `description` | [description](data-sources--certificate_chain--reference--group-001.md#canonical-3230313022023313-2131201013032210-3211001301103033-0331002300012231-3202321201323330-2220002211103031-0020000201033200-1320313223022322) |
| `id` | [ID](data-sources--certificate_chain--reference--group-001.md#canonical-0103013132121113-3103232113200303-3001203221021333-2320011033033222-1021321020012302-2211321321110331-2313103203301302-1221202223110332) |
| `labels` | [labels](data-sources--certificate_chain--reference--group-001.md#canonical-2120130332010202-2031100101102231-0133123303313320-0210203202113111-2032113231310303-0003220121302000-2102021113331211-3032111011022003) |
| `name` | [name](data-sources--certificate_chain--reference--group-001.md#canonical-1102120310202110-2122103210320031-3110221002001331-2100230221201210-0023023312310111-3131021322003123-0012013122203012-1323013130330211) |
| `namespace` | [namespace](data-sources--certificate_chain--reference--group-001.md#canonical-0032321121222320-2323232222102223-3101001110220200-2303231023330213-2102203333130202-2012003131112321-0131301230201203-1112020123332102) |
