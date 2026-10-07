---
page_title: "xcsh_virtual_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_virtual_site reference."
---

# xcsh_virtual_site reference

<a id="canonical-0001012332120212-3321121210011001-1030302032121112-1211101313330022-1313122033210013-3011100213131302-3231210233123131-0230121213221310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_virtual_site](../data-sources/virtual_site.md#canonical-1031213012322000-0133010132320022-1031302313301212-0223023031301131-3333302202323221-1032021131132033-3323021133330011-1320213202212000)
- Property reference

<a id="canonical-0133102123111130-0110312132033332-2331121203232033-3013012103000012-0020002100233211-3320232120222233-3110231313203221-3220323123212212"></a>

### Direct properties for `xcsh_virtual_site`

<a id="canonical-0333120201133211-3300023123000213-3230203130030123-2013010230130132-2031312231003023-1122330121220113-0231201232131212-2003033011211232"></a>

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

<a id="canonical-0020111213123012-0130121031013212-2302203020332221-3331233131002312-1100332330111210-2112020010213302-1100231232031120-2122111132301001"></a>

<a id="canonical-0023133131230020-3312322210300131-2121213012323033-0103210310123002-0032113020012202-1111203321223010-2232310301213111-0132310330223202"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the VirtualSite.

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

<a id="canonical-0203132013121302-1232230112012201-3111002110210133-3202221312322310-2333020033013033-3201132230003331-2321021001101012-3330111202231311"></a>

<a id="canonical-0110123312303021-2023111212111110-0232321230133320-1112121120202231-3020301000211222-0331331322333211-0033212210222310-3022020021221301"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0202103222212133-2011220231331013-1123203133321022-1212222333203132-0233320223233120-0232023020320220-0333111111011003-1113123303120120"></a>

<a id="canonical-3213333002232100-2020030221222122-2322202011103133-2320112211112121-2112033311030332-0322000200001321-2311302111010211-0303301022010033"></a>

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

<a id="canonical-0201112213121101-1303013330222123-2011300301300121-2121213121301123-0213111033123112-3112231223202003-0022130321331203-3113231011122100"></a>

<a id="canonical-0031002303212201-0011122313200011-1122120233210102-1232033010211031-1330121320001030-1202200221212221-1300222021213002-1202023030220123"></a>

#### `name` property

Type: `"string"`. Required.

Name of the VirtualSite.

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

<a id="canonical-1220323210302002-3022100103002030-2303022302123332-1013200123312302-2221011131010233-3023323222021123-2100222021201101-3032210323003010"></a>

<a id="canonical-1120330232330222-3203103113232031-0132301020102121-0333330201022212-1002033221130010-1331130320322213-0213303321123333-1121222300313333"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the VirtualSite exists.

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

- [site_selector](data-sources--virtual_site--reference--group-001.md#canonical-2301103133123210-2313021021321321-1233320023113120-0032210032021100-1111330222102332-2223333000332112-0023322211321132-1200220130123003): complete subsection reference.

<a id="canonical-2133321121131010-3213320121023011-1302230113133110-2022221121231000-3201333322110310-1132121100211221-2212313212113221-1120120232121110"></a>

<a id="canonical-3202231022112313-3003233031222021-0103122020120203-1121213013232303-2102223322003313-2220133033222200-1113113330201231-2332013111110010"></a>

#### `site_type` property

Type: `"string"`. Computed.

\[Enum: INVALID|REGIONAL\_EDGE|CUSTOMER\_EDGE|NGINX\_ONE\] Site Type which can either RE or CE
Invalid type of site Regional Edge site Customer Edge site. Possible values are \`INVALID\`,
\`REGIONAL\_EDGE\`, \`CUSTOMER\_EDGE\`, \`NGINX\_ONE\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALID",
  "enum": [
    "INVALID",
    "REGIONAL_EDGE",
    "CUSTOMER_EDGE",
    "NGINX_ONE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1221023203331332-0311011320103000-3221020232311100-3231323003210032-3010320032322222-2301122032003003-2121310112301330-2231102112012332"></a>

### All schema paths for `xcsh_virtual_site`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--virtual_site--reference--group-001.md#canonical-0333120201133211-3300023123000213-3230203130030123-2013010230130132-2031312231003023-1122330121220113-0231201232131212-2003033011211232) |
| `description` | [description](data-sources--virtual_site--reference--group-001.md#canonical-0020111213123012-0130121031013212-2302203020332221-3331233131002312-1100332330111210-2112020010213302-1100231232031120-2122111132301001) |
| `id` | [ID](data-sources--virtual_site--reference--group-001.md#canonical-0203132013121302-1232230112012201-3111002110210133-3202221312322310-2333020033013033-3201132230003331-2321021001101012-3330111202231311) |
| `labels` | [labels](data-sources--virtual_site--reference--group-001.md#canonical-0202103222212133-2011220231331013-1123203133321022-1212222333203132-0233320223233120-0232023020320220-0333111111011003-1113123303120120) |
| `name` | [name](data-sources--virtual_site--reference--group-001.md#canonical-0201112213121101-1303013330222123-2011300301300121-2121213121301123-0213111033123112-3112231223202003-0022130321331203-3113231011122100) |
| `namespace` | [namespace](data-sources--virtual_site--reference--group-001.md#canonical-1220323210302002-3022100103002030-2303022302123332-1013200123312302-2221011131010233-3023323222021123-2100222021201101-3032210323003010) |
| `site_selector` | [site_selector](data-sources--virtual_site--reference--group-001.md#canonical-2301012130331220-2121320013322022-0223012123003200-2322001022331111-1133011333213002-1002331002122200-3203231222021312-3312131003321331) |
| `site_selector.expressions` | [site_selector.expressions](data-sources--virtual_site--reference--group-001.md#canonical-1111221032312121-1130010113110030-3312230132101223-2112310220101322-3030322320330022-1102232300323333-2122100222331231-1110131100112211) |
| `site_type` | [site_type](data-sources--virtual_site--reference--group-001.md#canonical-2133321121131010-3213320121023011-1302230113133110-2022221121231000-3201333322110310-1132121100211221-2212313212113221-1120120232121110) |

<a id="canonical-2301103133123210-2313021021321321-1233320023113120-0032210032021100-1111330222102332-2223333000332112-0023322211321132-1200220130123003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_selector` properties

Breadcrumbs:

- [xcsh_virtual_site](../data-sources/virtual_site.md#canonical-1031213012322000-0133010132320022-1031302313301212-0223023031301131-3333302202323221-1032021131132033-3323021133330011-1320213202212000)
- [Property reference](data-sources--virtual_site--reference--group-001.md#canonical-0001012332120212-3321121210011001-1030302032121112-1211101313330022-1313122033210013-3011100213131302-3231210233123131-0230121213221310)
- site_selector

<a id="canonical-2301012130331220-2121320013322022-0223012123003200-2322001022331111-1133011333213002-1002331002122200-3203231222021312-3312131003321331"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0330221323222322-2313023200113022-0021023022122122-3031021221100000-2322313231101331-2022002033032133-0200230003331100-0122212010310000"></a>

### Direct properties for `site_selector`

<a id="canonical-1111221032312121-1130010113110030-3312230132101223-2112310220101322-3030322320330022-1102232300323333-2122100222331231-1110131100112211"></a>

#### `site_selector.expressions` property

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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
