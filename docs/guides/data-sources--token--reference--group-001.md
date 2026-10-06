---
page_title: "xcsh_token reference"
subcategory: "Identity"
description: "Complete grouped canonical reference for xcsh_token reference."
---

# xcsh_token reference

<a id="canonical-1302021003003321-1033031123003102-3221300113210300-3012332223200031-1300210333011012-3120203201330301-0001020001012101-1203222133001122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_token](../data-sources/token.md#canonical-2000130020133110-1113012332320202-2223121002312130-2001210311231111-3323332201023323-2311120322030210-1131312201010133-0322033200102202)
- Property reference

<a id="canonical-0223331101113313-0001121021301002-0032322110333331-0103023211133212-2301312131323303-1221220003332220-2010133003033211-3031123101233003"></a>

### Direct properties for `xcsh_token`

<a id="canonical-0301311303102120-3030331330302310-0203122311203210-1131122202030223-1003131200203023-3303333323203231-0030212012032232-1110003210023301"></a>

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

<a id="canonical-2021113313123121-1300223311020310-1020213321113021-3001130332033230-0130113230300013-2200022112312003-3010331312112022-3331320112302112"></a>

<a id="canonical-1022022010020110-0232122312313032-1221023200103301-0331111300121002-2323110130301132-0221301032331022-2100231112233333-1001131312120230"></a>

#### `content` property

Type: `"string"`. Computed, Sensitive.

Server-issued JWT registration credential.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-sensitive": true
}
```

<a id="canonical-1332113303200003-3111120120012121-0213012303033233-0011112002030222-0200031200011231-2022102020123232-1111322323120000-1111130322311021"></a>

<a id="canonical-2213232332010210-1112221000130311-1013200120031121-3021301022333132-2012201301232003-1220233132020103-3311300220223033-0012332011301213"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Token.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2033213222302330-2013000033131120-3212213200322233-3123120102203221-1032121221133220-1131330013303222-1131022313333210-1010200220101103"></a>

<a id="canonical-0211223021333303-2000021333032312-1223300123012221-1000100130200030-3003132232031220-0303210330002033-1133310020331331-2313331313233210"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3103122121020122-1000101012121012-2030120121111012-0330021121133001-0323331321003302-1302331220113221-3110332303101330-2210331320230320"></a>

<a id="canonical-0212103211100333-1121102333213200-2132131023220310-0030201121031121-2201103220232031-2222302111001030-3101112033103102-1103132312303121"></a>

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

<a id="canonical-2311332320332211-1333030323312333-1221300030203132-3332002230300113-2302033332110032-0110011023012232-1203322002301031-2212303211330203"></a>

<a id="canonical-0030033301123220-2311332022323122-3100233312121232-3122100302311123-2303121330103111-3032012300203103-3113021211103020-0032233131313232"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Token.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3112002213000011-1023322003221010-1221232200202302-1233201031210323-3230332211310200-0122213033010313-0101302021002103-0332133221230310"></a>

<a id="canonical-2213120110312330-2302103101233212-1113233200030233-2332030012310212-2032122300021223-1312331313011201-2113221132303120-2102120301123333"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace where the Token exists.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3123213003013221-2221012122022121-2311122132023032-2212223121310333-3010232121021000-1232303320222120-2131102001303330-0023320121132313"></a>

<a id="canonical-0310011102311020-1133021330002303-3211100002220310-2333120203323133-2313103030302020-0010132210312302-1002301130230200-3210000323323103"></a>

#### `site_name` property

Type: `"string"`. Computed.

Secure Mesh Site v2 name bound into a JWT token.

<a id="canonical-3202312233002202-1223331013033021-1222311002101220-0030301003003011-0323300012100301-3021202130100312-0200123303133022-0232201333101212"></a>

<a id="canonical-1212003032303022-1310332301001121-0111123030130132-0320311232323021-2330112203023220-1313131303023302-2323111233200132-1020121223332003"></a>

#### `type` property

Type: `"number"`. Computed.

\[Enum: 0|1\] Token type, where 0 is NORMAL and 1 is JWT. Possible values are \`0\`, \`1\`.

Receipt-pinned upstream constraints:

```json
{
  "default": 0,
  "enum": [
    0,
    1
  ]
}
```

<a id="canonical-2101312311223303-2232020322230111-3300011120133101-1021310212301311-0210032110231130-0312320021011000-2013123331131331-1100232022233222"></a>

<a id="canonical-2210213032013102-0102000023211313-1120131230301323-2030031221203121-2123030013122331-2301022010223032-0200323221211322-3031020123111333"></a>

#### `uid` property

Type: `"string"`. Computed, Sensitive.

Effective sensitive CE registration credential. NORMAL tokens use \`system\_metadata.uid\`; JWT
tokens use \`spec.content\`. This value is stored in plain text in the Terraform state file; ensure
your state file is properly secured.

<a id="canonical-3001120301112023-1113112022121230-1311033333210110-3122302123102122-1332123012233310-0320022030300312-2022012232130301-2120003023212021"></a>

### All schema paths for `xcsh_token`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--token--reference--group-001.md#canonical-0301311303102120-3030331330302310-0203122311203210-1131122202030223-1003131200203023-3303333323203231-0030212012032232-1110003210023301) |
| `content` | [content](data-sources--token--reference--group-001.md#canonical-2021113313123121-1300223311020310-1020213321113021-3001130332033230-0130113230300013-2200022112312003-3010331312112022-3331320112302112) |
| `description` | [description](data-sources--token--reference--group-001.md#canonical-1332113303200003-3111120120012121-0213012303033233-0011112002030222-0200031200011231-2022102020123232-1111322323120000-1111130322311021) |
| `id` | [ID](data-sources--token--reference--group-001.md#canonical-2033213222302330-2013000033131120-3212213200322233-3123120102203221-1032121221133220-1131330013303222-1131022313333210-1010200220101103) |
| `labels` | [labels](data-sources--token--reference--group-001.md#canonical-3103122121020122-1000101012121012-2030120121111012-0330021121133001-0323331321003302-1302331220113221-3110332303101330-2210331320230320) |
| `name` | [name](data-sources--token--reference--group-001.md#canonical-2311332320332211-1333030323312333-1221300030203132-3332002230300113-2302033332110032-0110011023012232-1203322002301031-2212303211330203) |
| `namespace` | [namespace](data-sources--token--reference--group-001.md#canonical-3112002213000011-1023322003221010-1221232200202302-1233201031210323-3230332211310200-0122213033010313-0101302021002103-0332133221230310) |
| `site_name` | [site_name](data-sources--token--reference--group-001.md#canonical-3123213003013221-2221012122022121-2311122132023032-2212223121310333-3010232121021000-1232303320222120-2131102001303330-0023320121132313) |
| `type` | [type](data-sources--token--reference--group-001.md#canonical-3202312233002202-1223331013033021-1222311002101220-0030301003003011-0323300012100301-3021202130100312-0200123303133022-0232201333101212) |
| `uid` | [uid](data-sources--token--reference--group-001.md#canonical-2101312311223303-2232020322230111-3300011120133101-1021310212301311-0210032110231130-0312320021011000-2013123331131331-1100232022233222) |
