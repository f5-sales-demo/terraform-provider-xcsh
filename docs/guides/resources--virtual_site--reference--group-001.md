---
page_title: "xcsh_virtual_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_virtual_site reference."
---

# xcsh_virtual_site reference

<a id="canonical-3012320130112032-2032311131023201-2230001222123031-0023020210201003-3230331331010002-2310130201122230-2211332320033010-2101122120231312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220312100100103-2202211333303212-1013102321102112-0111000331021332-2312132212300023-3112102001230031-3002112310123211-2010310201300103"></a>

## Property reference — Property reference / 012101021130 / 2

Breadcrumbs:

- [xcsh_virtual_site](../resources/virtual_site.md#canonical-0021200012330233-0110101013310110-3130112023213001-0301030013233201-0302311302331132-3122030310203321-2022313333210003-2333313130301011)
- Property reference

<a id="canonical-2232321333230331-3201303023213120-2320212010302102-1022002102031222-0311232332220121-2330130030200223-0333220312001201-3221231001202301"></a>

## Direct properties — Property reference / 012101021130 / 3

<a id="canonical-0103330203203133-3012130111010003-2110003002202113-3111131213031101-0003323323131102-3311121122033323-1331202102210223-3110221303201100"></a>

<a id="canonical-2333000023122301-3013113303122003-2301002102003100-0223200033130222-3321003301113300-3302002302221331-2333113103111020-3203202003332202"></a>

## annotations property — Property reference / 012101021130 / 4

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

<a id="canonical-0320102013123232-0110202210210331-2312111312002332-1132122131303023-2013030102320112-0330002323233322-2123310120103213-1313032232012322"></a>

<a id="canonical-0231020100033031-0332203231110311-3013102111022320-2013322030312311-1103032100221211-3201310231132222-3231112331203300-1303022033222323"></a>

## description property — Property reference / 012101021130 / 5

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

<a id="canonical-3323212320231232-1230112020331212-0110123032210010-2102001332011003-3010030330003011-3211000211030103-0200111232021211-0123131100111030"></a>

<a id="canonical-3112012233330213-1121213312021221-1023202003031303-2231332132133001-1000010011322120-2030212132122001-1013023211013120-1333212233120111"></a>

## disable property — Property reference / 012101021130 / 6

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

<a id="canonical-1012103211133122-3313200301122330-2322203132110300-3203110003031201-0113033123102033-3032212202100201-2132111112101231-3012321000201120"></a>

<a id="canonical-1112112003102320-2032113203031302-1002031322311203-2200202022301103-2022310122022033-1312220032332003-1103032330101132-0302013000012023"></a>

## ID property — Property reference / 012101021130 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0120000032103103-2102013323103011-2332001320213213-2303303321020322-3030002100220321-0102233211203213-3322033232022133-3303210202232211"></a>

<a id="canonical-1113300232201210-1331232033013301-2211211122121023-3312100211033131-0011130213213332-1223110120031223-0113030200001023-3132033021010013"></a>

## labels property — Property reference / 012101021130 / 8

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

<a id="canonical-0233212133211320-2330012033130303-2131033302100020-2103103312211323-1331303100312030-0330123220322223-2012300121131313-0320311230022123"></a>

<a id="canonical-2300202030202102-0101101031120332-3312302132200200-1030000303020320-3210223321232101-2022020031013330-2110231110333230-0113002123311130"></a>

## name property — Property reference / 012101021130 / 9

Type: `"string"`. Required.

Name of the Virtual Site. Must be unique within the namespace.

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

<a id="canonical-1213111112203232-3333330230323012-1131022222302312-1213331131023222-3130002213003033-2302201103102132-0131010322120012-3130311022100232"></a>

<a id="canonical-1232212220322111-3020123311133131-2211100011112302-1122222201313023-0303301103011021-3203112211101020-1300013212301032-0020103220300032"></a>

## namespace property — Property reference / 012101021130 / 10

Type: `"string"`. Required.

Namespace where the Virtual Site is created.

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

- [site_selector](resources--virtual_site--reference--group-001.md#canonical-2223123003130212-2302221302031110-1102023220300113-3300102323202122-3202123310120200-0111201111022123-2220031220320212-3221303320302201): complete subsection reference.

<a id="canonical-2123100331112003-3321321202111000-1331132000323130-1220311212113210-3020203102211000-2120123320100113-2332203121112312-1232313101331111"></a>

<a id="canonical-2202220320301100-2313233221012012-1020103121131133-2112133112113002-2203330330301032-0103113203320203-3223303023133223-3130330133233313"></a>

## site_type property — Property reference / 012101021130 / 11

Type: `"string"`. Optional, Computed.

\[Enum: INVALID|REGIONAL\_EDGE|CUSTOMER\_EDGE|NGINX\_ONE\] Site Type which can either RE or CE
Invalid type of site Regional Edge site Customer Edge site. Possible values are \`INVALID\`,
\`REGIONAL\_EDGE\`, \`CUSTOMER\_EDGE\`, \`NGINX\_ONE\`.

Upstream description:

Site Type which can either RE or CE

Invalid type of site Regional Edge site Customer Edge site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INVALID",
    "REGIONAL_EDGE",
    "CUSTOMER_EDGE",
    "NGINX_ONE"),
}
```

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

- [timeouts](resources--virtual_site--reference--group-001.md#canonical-2101121322312000-0133033232020332-0110221120212021-1023330001130013-1121320133203331-2312313310222222-3012223221231110-2010333212103333): complete subsection reference.

<a id="canonical-1003332122233332-2303101212010111-1201111200322003-3111132231320231-0130221103300223-0320303212101211-2232231111100201-0121120203022331"></a>

## All schema paths — Property reference / 012101021130 / 12

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--virtual_site--reference--group-001.md#canonical-0103330203203133-3012130111010003-2110003002202113-3111131213031101-0003323323131102-3311121122033323-1331202102210223-3110221303201100) |
| `description` | [description](resources--virtual_site--reference--group-001.md#canonical-0320102013123232-0110202210210331-2312111312002332-1132122131303023-2013030102320112-0330002323233322-2123310120103213-1313032232012322) |
| `disable` | [disable](resources--virtual_site--reference--group-001.md#canonical-3323212320231232-1230112020331212-0110123032210010-2102001332011003-3010030330003011-3211000211030103-0200111232021211-0123131100111030) |
| `id` | [id](resources--virtual_site--reference--group-001.md#canonical-1012103211133122-3313200301122330-2322203132110300-3203110003031201-0113033123102033-3032212202100201-2132111112101231-3012321000201120) |
| `labels` | [labels](resources--virtual_site--reference--group-001.md#canonical-0120000032103103-2102013323103011-2332001320213213-2303303321020322-3030002100220321-0102233211203213-3322033232022133-3303210202232211) |
| `name` | [name](resources--virtual_site--reference--group-001.md#canonical-0233212133211320-2330012033130303-2131033302100020-2103103312211323-1331303100312030-0330123220322223-2012300121131313-0320311230022123) |
| `namespace` | [namespace](resources--virtual_site--reference--group-001.md#canonical-1213111112203232-3333330230323012-1131022222302312-1213331131023222-3130002213003033-2302201103102132-0131010322120012-3130311022100232) |
| `site_selector` | [site_selector](resources--virtual_site--reference--group-001.md#canonical-2011211311200331-1332122312233301-0201331322213322-1030000130103213-2333032312200033-3022222012111002-1212212113331123-2103230210321302) |
| `site_selector.expressions` | [site_selector.expressions](resources--virtual_site--reference--group-001.md#canonical-3312033231333303-2321230022101112-1100102223310210-1123102233320000-1223021111231120-1130330112031121-0000002130311233-3131312312100312) |
| `site_type` | [site_type](resources--virtual_site--reference--group-001.md#canonical-2123100331112003-3321321202111000-1331132000323130-1220311212113210-3020203102211000-2120123320100113-2332203121112312-1232313101331111) |
| `timeouts` | [timeouts](resources--virtual_site--reference--group-001.md#canonical-2210022230013002-0233010221222223-3100312003030131-2132301332003000-3223021212130020-1312102022120321-2002322102121311-2221231133020103) |
| `timeouts.create` | [timeouts.create](resources--virtual_site--reference--group-001.md#canonical-3020030133010132-1211113203002120-1231013221223301-3213031223222323-3231221303233232-3121312123032023-0130321203333030-3320330320312011) |
| `timeouts.delete` | [timeouts.delete](resources--virtual_site--reference--group-001.md#canonical-1232102122023333-2000113222220231-0133132231200213-3130301112013200-0003202202011211-1012002033212320-2202211031222023-2313210200320132) |
| `timeouts.read` | [timeouts.read](resources--virtual_site--reference--group-001.md#canonical-0221312210332330-3032310030332302-2103320301301201-0122203233003232-2230211312213011-1200013010300202-1332202021203121-1111112121000113) |
| `timeouts.update` | [timeouts.update](resources--virtual_site--reference--group-001.md#canonical-2321220102103120-2133031332200333-1201103123230231-0332313220003011-1012330100321222-1232231012323113-3002013233022200-0211032211100031) |

<a id="canonical-0030010130001021-2022331200111113-1120111010212200-2021333310301032-0110233132113000-2202022001301303-0231033021131132-1321010310103112"></a>

## Next pages — Property reference / 012101021130 / 13

- [site_selector](resources--virtual_site--reference--group-001.md#canonical-2223123003130212-2302221302031110-1102023220300113-3300102323202122-3202123310120200-0111201111022123-2220031220320212-3221303320302201)
- [timeouts](resources--virtual_site--reference--group-001.md#canonical-2101121322312000-0133033232020332-0110221120212021-1023330001130013-1121320133203331-2312313310222222-3012223221231110-2010333212103333)
- [xcsh_virtual_site](../resources/virtual_site.md#canonical-0021200012330233-0110101013310110-3130112023213001-0301030013233201-0302311302331132-3122030310203321-2022313333210003-2333313130301011)

<a id="canonical-2223123003130212-2302221302031110-1102023220300113-3300102323202122-3202123310120200-0111201111022123-2220031220320212-3221303320302201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111031321311121-3302003022001133-3313233121020111-3032213013122101-0321232322310131-1321113033222213-1021120103212321-2111332302200031"></a>

## site_selector — site_selector / 321010131321 / 2

Breadcrumbs:

- [xcsh_virtual_site](../resources/virtual_site.md#canonical-0021200012330233-0110101013310110-3130112023213001-0301030013233201-0302311302331132-3122030310203321-2022313333210003-2333313130301011)
- [Property reference](resources--virtual_site--reference--group-001.md#canonical-3012320130112032-2032311131023201-2230001222123031-0023020210201003-3230331331010002-2310130201122230-2211332320033010-2101122120231312)
- site_selector

<a id="canonical-2011211311200331-1332122312233301-0201331322213322-1030000130103213-2333032312200033-3022222012111002-1212212113331123-2103230210321302"></a>

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
site_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-1123203123331222-2201321020023003-0302021130132122-2201031202130332-3201332202222123-3103222313322012-2231312023011210-1011213033201130"></a>

## Direct properties — site_selector / 321010131321 / 3

<a id="canonical-3312033231333303-2321230022101112-1100102223310210-1123102233320000-1223021111231120-1130330112031121-0000002130311233-3131312312100312"></a>

<a id="canonical-1101232211210201-1320301320133213-2000311010003011-2103210123322320-0131213332101212-1002122122231112-3323031230033233-2230200130101232"></a>

## expressions property — site_selector / 321010131321 / 4

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

<a id="canonical-0211211112031110-3010031210111002-3021023100102120-3122220111121103-1211330333000032-3031200111112102-3202223132103103-1200232332100022"></a>

## Next pages — site_selector / 321010131321 / 5

- [Property reference](resources--virtual_site--reference--group-001.md#canonical-3012320130112032-2032311131023201-2230001222123031-0023020210201003-3230331331010002-2310130201122230-2211332320033010-2101122120231312)
- [xcsh_virtual_site](../resources/virtual_site.md#canonical-0021200012330233-0110101013310110-3130112023213001-0301030013233201-0302311302331132-3122030310203321-2022313333210003-2333313130301011)

<a id="canonical-2101121322312000-0133033232020332-0110221120212021-1023330001130013-1121320133203331-2312313310222222-3012223221231110-2010333212103333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223212022130220-3003033322120311-3311003210023320-1032103102132103-0312232122221331-1233032002211012-1320313131332123-3111120222131133"></a>

## timeouts — timeouts / 212033231331 / 2

Breadcrumbs:

- [xcsh_virtual_site](../resources/virtual_site.md#canonical-0021200012330233-0110101013310110-3130112023213001-0301030013233201-0302311302331132-3122030310203321-2022313333210003-2333313130301011)
- [Property reference](resources--virtual_site--reference--group-001.md#canonical-3012320130112032-2032311131023201-2230001222123031-0023020210201003-3230331331010002-2310130201122230-2211332320033010-2101122120231312)
- timeouts

<a id="canonical-2210022230013002-0233010221222223-3100312003030131-2132301332003000-3223021212130020-1312102022120321-2002322102121311-2221231133020103"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310321032300333-1323012211322220-3211032121003331-0100012330212133-0222320232331212-1132132331302130-2222221331012322-2031332120302202"></a>

## Direct properties — timeouts / 212033231331 / 3

<a id="canonical-3020030133010132-1211113203002120-1231013221223301-3213031223222323-3231221303233232-3121312123032023-0130321203333030-3320330320312011"></a>

<a id="canonical-0122302323303130-1233020311131313-1312302102323212-3113113100032213-1131230000300132-1201033010013122-1230132312101323-1313131021330123"></a>

## create property — timeouts / 212033231331 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1232102122023333-2000113222220231-0133132231200213-3130301112013200-0003202202011211-1012002033212320-2202211031222023-2313210200320132"></a>

<a id="canonical-3113322332010030-3033130313230231-2120001133110312-2021230021021010-1032032311221030-2332131131311031-2333323021032103-2100122123133200"></a>

## delete property — timeouts / 212033231331 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0221312210332330-3032310030332302-2103320301301201-0122203233003232-2230211312213011-1200013010300202-1332202021203121-1111112121000113"></a>

<a id="canonical-2122320022221100-0232333111223130-1032220123112231-3022002103120332-3020213112333132-3111220320012300-0132123022333003-2022111003331112"></a>

## read property — timeouts / 212033231331 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2321220102103120-2133031332200333-1201103123230231-0332313220003011-1012330100321222-1232231012323113-3002013233022200-0211032211100031"></a>

<a id="canonical-0011222220021232-1002031201010023-1300322023023112-2102021300031202-0022220133322102-3311102103121102-2323303121232023-3300220320311010"></a>

## update property — timeouts / 212033231331 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0001233302311313-1321313113231011-2233122331330213-2210100202013330-0331233022331311-2311220212202223-1220321300213333-2012023303020001"></a>

## Next pages — timeouts / 212033231331 / 8

- [Property reference](resources--virtual_site--reference--group-001.md#canonical-3012320130112032-2032311131023201-2230001222123031-0023020210201003-3230331331010002-2310130201122230-2211332320033010-2101122120231312)
- [xcsh_virtual_site](../resources/virtual_site.md#canonical-0021200012330233-0110101013310110-3130112023213001-0301030013233201-0302311302331132-3122030310203321-2022313333210003-2333313130301011)
