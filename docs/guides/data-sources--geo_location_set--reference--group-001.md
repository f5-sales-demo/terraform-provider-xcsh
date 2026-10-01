---
page_title: "xcsh_geo_location_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_geo_location_set reference."
---

# xcsh_geo_location_set reference

<a id="canonical-1123331121000122-0300310122302103-1323002103311113-2222120103331101-2033133122011030-3102030113320122-1333103020013031-1112201210133330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332123302022121-1223230333211010-3112210001131222-3031210121311121-3213323230031032-3033021300331120-1012111232320120-1023130120332233"></a>

## Property reference — Property reference / 010201211012 / 2

Breadcrumbs:

- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-1001100000212311-3233013232003113-2232222223002331-1133222232220002-1101232133012211-1121021022203120-3302330011301312-1023310223132122)
- Property reference

<a id="canonical-3300310033110132-3020212130122303-0330102112220100-0300301130010022-1201200332013200-3031202030020203-1010022132101021-2323033302002203"></a>

## Direct properties — Property reference / 010201211012 / 3

<a id="canonical-1223120001001203-1213203220232112-0033122233222031-0303002012223101-1012213333133011-2212322120131102-3221323032200320-0021200030111110"></a>

<a id="canonical-1331011211023200-2211201323000023-1030223103003012-2202323330020103-0121112331013323-3001221032322103-0212000303321023-0112233120332312"></a>

## annotations property — Property reference / 010201211012 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
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

- [custom_geo_location_selector](data-sources--geo_location_set--reference--group-001.md#canonical-2020332100211112-3033212010031001-3133022212203121-1203000333110321-3032301101133301-1302321210330322-0211210121122321-0302331300332311): complete subsection reference.

<a id="canonical-3100310323233001-1023302310121313-1212321101021222-1231301131021103-2223100233303212-2002003221300330-3130030101313032-3300123010301110"></a>

<a id="canonical-2133011330320322-2121300101133030-3203011301023101-1010132113230302-2313323202031210-0113320232330120-3233202201132331-2122312220100323"></a>

## description property — Property reference / 010201211012 / 5

Type: `"string"`. Computed.

Description of the GeoLocationSet.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [global](data-sources--geo_location_set--reference--group-001.md#canonical-0323200211320033-2101113230310122-2202233031030132-0023121023113003-0033113133130230-3222222203132322-1103113222012121-2123312112303323): complete subsection reference.

<a id="canonical-2132222112020020-2202101222120301-2303020113212303-2202002100333220-3321031121303210-3232203313210213-0001311000100123-1123230333210023"></a>

<a id="canonical-2221120303123311-2200222113011120-1303123022313030-3210302301113122-2333322332123031-3233021320323020-1310210323121222-2010132301133022"></a>

## ID property — Property reference / 010201211012 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1203020030320301-2030121310333021-3023111021322331-1011013100332221-1031120213000012-3101301323201331-2003322022002323-2110120111110100"></a>

<a id="canonical-1203021220033012-2230012212222031-0222212203110303-3110022223303212-2101322112231111-3112211232020301-0211100122133302-1313122330133312"></a>

## labels property — Property reference / 010201211012 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-1210030310301313-2322031212312332-1022031112231022-3010222131213031-1010301112111220-3331031032223213-2201301203010123-0203003112210001"></a>

<a id="canonical-2003133033212131-2221301313110112-3211333332230120-2030020121130220-0233103332231222-0033101011220201-1301203112033303-3023211101202300"></a>

## name property — Property reference / 010201211012 / 8

Type: `"string"`. Required.

Name of the GeoLocationSet.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2032133131301113-1132213023011021-0323222111030332-1031313010220111-0112120230233132-3310131212233221-0211030203212021-1222323321112321"></a>

<a id="canonical-2322220211332112-0313200133112322-0002011110113003-2113102031323331-2210233321033222-2212122110211021-3312123003221111-3102002101313211"></a>

## namespace property — Property reference / 010201211012 / 9

Type: `"string"`. Optional, Computed.

Namespace where the GeoLocationSet exists.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2133201021301332-1311301323111110-1202323011323200-0212331130201321-3010232230222100-1113200230110001-3133310200220312-3232323120023030"></a>

## All schema paths — Property reference / 010201211012 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--geo_location_set--reference--group-001.md#canonical-1223120001001203-1213203220232112-0033122233222031-0303002012223101-1012213333133011-2212322120131102-3221323032200320-0021200030111110) |
| `custom_geo_location_selector` | [custom_geo_location_selector](data-sources--geo_location_set--reference--group-001.md#canonical-2322103221130012-3033032313132323-2023220112202113-0010011222101122-1310313122230330-3211131100011133-2202230023131113-1031033213233213) |
| `custom_geo_location_selector.expressions` | [custom_geo_location_selector.expressions](data-sources--geo_location_set--reference--group-001.md#canonical-3301031021013323-0230213120210132-0222323121103200-2103121130021313-1300032331221202-0123332013131102-3112122303203023-2001311230123223) |
| `description` | [description](data-sources--geo_location_set--reference--group-001.md#canonical-3100310323233001-1023302310121313-1212321101021222-1231301131021103-2223100233303212-2002003221300330-3130030101313032-3300123010301110) |
| `global` | [global](data-sources--geo_location_set--reference--group-001.md#canonical-0033202232001233-0011021231101202-2202331022331010-2001133302312130-3012333321311222-1030112312010303-2232303333110031-2211300132331101) |
| `id` | [id](data-sources--geo_location_set--reference--group-001.md#canonical-2132222112020020-2202101222120301-2303020113212303-2202002100333220-3321031121303210-3232203313210213-0001311000100123-1123230333210023) |
| `labels` | [labels](data-sources--geo_location_set--reference--group-001.md#canonical-1203020030320301-2030121310333021-3023111021322331-1011013100332221-1031120213000012-3101301323201331-2003322022002323-2110120111110100) |
| `name` | [name](data-sources--geo_location_set--reference--group-001.md#canonical-1210030310301313-2322031212312332-1022031112231022-3010222131213031-1010301112111220-3331031032223213-2201301203010123-0203003112210001) |
| `namespace` | [namespace](data-sources--geo_location_set--reference--group-001.md#canonical-2032133131301113-1132213023011021-0323222111030332-1031313010220111-0112120230233132-3310131212233221-0211030203212021-1222323321112321) |

<a id="canonical-2213311113000011-1023123030221231-1200233311212100-0112222303113000-3333023033031010-0121221021300221-3223101123033310-2331020221002302"></a>

## Next pages — Property reference / 010201211012 / 11

- [custom_geo_location_selector](data-sources--geo_location_set--reference--group-001.md#canonical-2020332100211112-3033212010031001-3133022212203121-1203000333110321-3032301101133301-1302321210330322-0211210121122321-0302331300332311)
- [global](data-sources--geo_location_set--reference--group-001.md#canonical-0323200211320033-2101113230310122-2202233031030132-0023121023113003-0033113133130230-3222222203132322-1103113222012121-2123312112303323)
- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-1001100000212311-3233013232003113-2232222223002331-1133222232220002-1101232133012211-1121021022203120-3302330011301312-1023310223132122)

<a id="canonical-2020332100211112-3033212010031001-3133022212203121-1203000333110321-3032301101133301-1302321210330322-0211210121122321-0302331300332311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020320303121313-0223031223000131-2333121221100232-3200200231221202-1130102202103302-1111130022032230-2032310302313231-3203132221331231"></a>

## custom_geo_location_selector — custom_geo_location_selector / 133312023130 / 2

Breadcrumbs:

- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-1001100000212311-3233013232003113-2232222223002331-1133222232220002-1101232133012211-1121021022203120-3302330011301312-1023310223132122)
- [Property reference](data-sources--geo_location_set--reference--group-001.md#canonical-1123331121000122-0300310122302103-1323002103311113-2222120103331101-2033133122011030-3102030113320122-1333103020013031-1112201210133330)
- custom_geo_location_selector

<a id="canonical-2322103221130012-3033032313132323-2023220112202113-0010011222101122-1310313122230330-3211131100011133-2202230023131113-1031033213233213"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_geo\_location\_selector, global\] Type can be used to establish a 'selector
reference' from one object(called selector) to a set of other objects(called selectees) based on the
value of expressions. A label selector is a label query over a set of resources. An empty label
selector matches all objects.

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

- [custom_geo_location_selector](data-sources--geo_location_set--reference--group-001.md#canonical-2322103221130012-3033032313132323-2023220112202113-0010011222101122-1310313122230330-3211131100011133-2202230023131113-1031033213233213)
- [global](data-sources--geo_location_set--reference--group-001.md#canonical-0033202232001233-0011021231101202-2202331022331010-2001133302312130-3012333321311222-1030112312010303-2232303333110031-2211300132331101)

Select alternatives according to the provider validators above.

<a id="canonical-2033131112322210-1020312112330010-2000130220001213-3003033122021011-3031300210322312-3333133131012212-1310023131100233-3220113320112010"></a>

## Direct properties — custom_geo_location_selector / 133312023130 / 3

<a id="canonical-3301031021013323-0230213120210132-0222323121103200-2103121130021313-1300032331221202-0123332013131102-3112122303203023-2001311230123223"></a>

<a id="canonical-3020311231003020-3200321232122002-2013133220011112-2322321311110032-3213011122021121-2231211112203201-3101002122232312-0213221003233222"></a>

## expressions property — custom_geo_location_selector / 133312023130 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2120020301011302-3201300310011220-1120331003232010-3121312111201031-2103030031130033-3302100111022331-2202020031113020-2010303032302123"></a>

## Next pages — custom_geo_location_selector / 133312023130 / 5

- [Property reference](data-sources--geo_location_set--reference--group-001.md#canonical-1123331121000122-0300310122302103-1323002103311113-2222120103331101-2033133122011030-3102030113320122-1333103020013031-1112201210133330)
- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-1001100000212311-3233013232003113-2232222223002331-1133222232220002-1101232133012211-1121021022203120-3302330011301312-1023310223132122)

<a id="canonical-0323200211320033-2101113230310122-2202233031030132-0023121023113003-0033113133130230-3222222203132322-1103113222012121-2123312112303323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033021312011021-2102113223323310-1120123301211101-3220300000020103-0111033123232302-1120211233233311-3210120331222032-3230311133013322"></a>

## global — global / 313213301323 / 2

Breadcrumbs:

- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-1001100000212311-3233013232003113-2232222223002331-1133222232220002-1101232133012211-1121021022203120-3302330011301312-1023310223132122)
- [Property reference](data-sources--geo_location_set--reference--group-001.md#canonical-1123331121000122-0300310122302103-1323002103311113-2222120103331101-2033133122011030-3102030113320122-1333103020013031-1112201210133330)
- global

<a id="canonical-0033202232001233-0011021231101202-2202331022331010-2001133302312130-3012333321311222-1030112312010303-2232303333110031-2211300132331101"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-3122111111213123-2012130230310133-0122003111031333-3320100111031120-0222101122201200-1020111013123132-1233132203113210-2100033110133202"></a>

## Direct properties — global / 313213301323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210120011331201-3032110222133332-2212200000033220-0211311113313233-2322012121133321-2330010033111201-3123133103210330-2331322120232312"></a>

## Next pages — global / 313213301323 / 4

- [Property reference](data-sources--geo_location_set--reference--group-001.md#canonical-1123331121000122-0300310122302103-1323002103311113-2222120103331101-2033133122011030-3102030113320122-1333103020013031-1112201210133330)
- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-1001100000212311-3233013232003113-2232222223002331-1133222232220002-1101232133012211-1121021022203120-3302330011301312-1023310223132122)
