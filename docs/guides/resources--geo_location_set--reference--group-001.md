---
page_title: "xcsh_geo_location_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_geo_location_set reference."
---

# xcsh_geo_location_set reference

<a id="canonical-1123112132012331-2320021312120003-1120132113110332-0331323012021123-1320320333300013-2001331312211030-0311313333232210-2132322132112012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233311230331003-0321100300131001-1212223200303013-2131103111103021-1300323301201001-1120001221212033-1303033310322102-2133021111111333"></a>

## Property reference — Property reference / 012003013212 / 2

Breadcrumbs:

- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-3103011222330300-2233200322023210-0200023031210213-3231332213130232-1013021130033033-2013330101030012-3301030030300101-3321313130311100)
- Property reference

<a id="canonical-3223322102101220-2200000311321333-0131312010312310-2220311332020311-2101331310000132-0233122130133223-0013321033012102-0032213311213120"></a>

## Direct properties — Property reference / 012003013212 / 3

<a id="canonical-3302102032112332-2030112203311223-1122322300301032-1213313011113003-2300003020222321-3103320320003212-3310113013121003-2321130123323013"></a>

<a id="canonical-1231022320312232-3022323013032303-1220222031201232-2313300330210333-2221212321033010-2311100122201333-0120110332311313-0201333221312123"></a>

## annotations property — Property reference / 012003013212 / 4

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

- [custom_geo_location_selector](resources--geo_location_set--reference--group-001.md#canonical-1312002213210212-0132131021201320-0011002030203311-3110323132311123-0312132032233211-2300021120300311-1223210211220101-2202203322211021): complete subsection reference.

<a id="canonical-0213303120013302-0320103033312223-0131002022213020-2113220230033023-0022203013301120-2203233312132223-2331122200002102-1022211232230333"></a>

<a id="canonical-1130333330021002-3001311121223110-1221003330022322-0332201331111123-3310123130011111-0331113210311322-0220013100110211-1033302110101232"></a>

## description property — Property reference / 012003013212 / 5

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

<a id="canonical-0121211001002022-1333322213332233-0032020121111112-3110001030000303-3123212330011212-3001310332220100-2100102121031132-2211200231221332"></a>

<a id="canonical-2012300231323233-3122120133323221-3101331331223023-0111101220311321-2212232223132020-2033001330223100-0210221022200012-1103102133133120"></a>

## disable property — Property reference / 012003013212 / 6

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

- [global](resources--geo_location_set--reference--group-001.md#canonical-3321203011103202-0103022333122011-1322313223010213-1110200030333021-1103232001011002-0001332310101033-3132221322222331-2122032032321002): complete subsection reference.

<a id="canonical-1022203303202320-2232033110203113-2203100333302011-0320200023220002-1103203000110033-2002213032323321-0322333123023331-1033023111133122"></a>

<a id="canonical-3110033333003122-0331221312011303-1332201320023101-0033101100120112-1100130313211332-3103300201221330-0320120001331202-0102012220233232"></a>

## ID property — Property reference / 012003013212 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3100022002110201-1012330133221032-0100312011023123-0300322331333103-1022003300211122-2322200110210112-1331120310031302-0213110112302130"></a>

<a id="canonical-3333203200110202-2202012323210333-3120301033013222-0211002232301311-0331332103012013-0321002010213203-3233230023113021-3313213020211323"></a>

## labels property — Property reference / 012003013212 / 8

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

<a id="canonical-0111332110200000-2331232000201103-0132113220302200-0112022211121311-1122233203223313-3122313310301333-2331213312302023-3313333113122323"></a>

<a id="canonical-3320102212330001-2233100020011002-3002102232301200-2131130010211300-2320323120330021-1230013111121200-1322000321112102-0311012331110331"></a>

## name property — Property reference / 012003013212 / 9

Type: `"string"`. Required.

Name of the Geo Location Set. Must be unique within the namespace.

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

<a id="canonical-1023030013032312-1320200220201202-2120221222310303-1020330131031121-1321330331230201-0102233222202302-1102301301231132-1001302013312230"></a>

<a id="canonical-2121202333111113-0223212231100120-2300033101001222-3310001002031320-0031123031312022-2102010221121301-2320230023022212-3113223120021330"></a>

## namespace property — Property reference / 012003013212 / 10

Type: `"string"`. Optional, Computed.

Namespace for the Geo Location Set. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
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

- [timeouts](resources--geo_location_set--reference--group-001.md#canonical-3131320102322030-1221310213222021-0222131000323103-2033130021100111-0010212112013320-2302231122113231-2033321211333003-3013032012213122): complete subsection reference.

<a id="canonical-3030113100011233-0220230210213303-3222122110302001-3021303031321102-3003213213021103-3000300101012333-2201300102020221-1113313002331223"></a>

## All schema paths — Property reference / 012003013212 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--geo_location_set--reference--group-001.md#canonical-3302102032112332-2030112203311223-1122322300301032-1213313011113003-2300003020222321-3103320320003212-3310113013121003-2321130123323013) |
| `custom_geo_location_selector` | [custom_geo_location_selector](resources--geo_location_set--reference--group-001.md#canonical-1213322210102021-0132012110310302-0213003101030033-3230020233313030-0333132033211320-3130003301211132-0010033011330001-0002211220313122) |
| `custom_geo_location_selector.expressions` | [custom_geo_location_selector.expressions](resources--geo_location_set--reference--group-001.md#canonical-3311320010103221-3020220222121302-3323320011033331-1101301323200133-3102232223111310-3323200203333033-0130231130103101-2301310322210102) |
| `description` | [description](resources--geo_location_set--reference--group-001.md#canonical-0213303120013302-0320103033312223-0131002022213020-2113220230033023-0022203013301120-2203233312132223-2331122200002102-1022211232230333) |
| `disable` | [disable](resources--geo_location_set--reference--group-001.md#canonical-0121211001002022-1333322213332233-0032020121111112-3110001030000303-3123212330011212-3001310332220100-2100102121031132-2211200231221332) |
| `global` | [global](resources--geo_location_set--reference--group-001.md#canonical-3113223103231012-0231223131322103-3101010122312103-3301230223330230-1330120033312121-0113003020023333-3132131230213000-0332232221112213) |
| `id` | [id](resources--geo_location_set--reference--group-001.md#canonical-1022203303202320-2232033110203113-2203100333302011-0320200023220002-1103203000110033-2002213032323321-0322333123023331-1033023111133122) |
| `labels` | [labels](resources--geo_location_set--reference--group-001.md#canonical-3100022002110201-1012330133221032-0100312011023123-0300322331333103-1022003300211122-2322200110210112-1331120310031302-0213110112302130) |
| `name` | [name](resources--geo_location_set--reference--group-001.md#canonical-0111332110200000-2331232000201103-0132113220302200-0112022211121311-1122233203223313-3122313310301333-2331213312302023-3313333113122323) |
| `namespace` | [namespace](resources--geo_location_set--reference--group-001.md#canonical-1023030013032312-1320200220201202-2120221222310303-1020330131031121-1321330331230201-0102233222202302-1102301301231132-1001302013312230) |
| `timeouts` | [timeouts](resources--geo_location_set--reference--group-001.md#canonical-0131002021300030-2131111220002030-1000221322123320-1100323232023200-1122113312302011-2030032101301110-0232133323210220-3213032232313333) |
| `timeouts.create` | [timeouts.create](resources--geo_location_set--reference--group-001.md#canonical-2220212121301212-0212212021223122-0112322300003103-1113010210232022-1231101331003200-2031101121032202-0110122121100310-2310313312211022) |
| `timeouts.delete` | [timeouts.delete](resources--geo_location_set--reference--group-001.md#canonical-2310002033200201-3032233232011112-2223013312211103-0300023220132130-1210232101120020-2321132021312220-3212303200103123-2333231113310120) |
| `timeouts.read` | [timeouts.read](resources--geo_location_set--reference--group-001.md#canonical-1131300023110321-3211111333100012-3120200213013110-0312301320030112-1122022333123000-0301132221023313-1112332130123121-0322131033121303) |
| `timeouts.update` | [timeouts.update](resources--geo_location_set--reference--group-001.md#canonical-0032030033211010-2123211231311110-2102002101133132-0320003110110230-1202302002220213-3301121323022103-1012023223121230-0310030331133233) |

<a id="canonical-2121022332133211-2022211131121202-3311133123332222-1322223200310302-2310213213220202-0110021200301001-3122012302333100-0132330330003231"></a>

## Next pages — Property reference / 012003013212 / 12

- [custom_geo_location_selector](resources--geo_location_set--reference--group-001.md#canonical-1312002213210212-0132131021201320-0011002030203311-3110323132311123-0312132032233211-2300021120300311-1223210211220101-2202203322211021)
- [global](resources--geo_location_set--reference--group-001.md#canonical-3321203011103202-0103022333122011-1322313223010213-1110200030333021-1103232001011002-0001332310101033-3132221322222331-2122032032321002)
- [timeouts](resources--geo_location_set--reference--group-001.md#canonical-3131320102322030-1221310213222021-0222131000323103-2033130021100111-0010212112013320-2302231122113231-2033321211333003-3013032012213122)
- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-3103011222330300-2233200322023210-0200023031210213-3231332213130232-1013021130033033-2013330101030012-3301030030300101-3321313130311100)

<a id="canonical-1312002213210212-0132131021201320-0011002030203311-3110323132311123-0312132032233211-2300021120300311-1223210211220101-2202203322211021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300112303111003-1003131021230000-1010301300232312-2210230223001203-1310323013033110-2001013230332222-2001110033121313-3033133203033213"></a>

## custom_geo_location_selector — custom_geo_location_selector / 110321122332 / 2

Breadcrumbs:

- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-3103011222330300-2233200322023210-0200023031210213-3231332213130232-1013021130033033-2013330101030012-3301030030300101-3321313130311100)
- [Property reference](resources--geo_location_set--reference--group-001.md#canonical-1123112132012331-2320021312120003-1120132113110332-0331323012021123-1320320333300013-2001331312211030-0311313333232210-2132322132112012)
- custom_geo_location_selector

<a id="canonical-1213322210102021-0132012110310302-0213003101030033-3230020233313030-0333132033211320-3130003301211132-0010033011330001-0002211220313122"></a>

Type: `"object"`. single nested block, Optional.

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

OneOf alternatives in this subsection:

- [custom_geo_location_selector](resources--geo_location_set--reference--group-001.md#canonical-1213322210102021-0132012110310302-0213003101030033-3230020233313030-0333132033211320-3130003301211132-0010033011330001-0002211220313122)
- [global](resources--geo_location_set--reference--group-001.md#canonical-3113223103231012-0231223131322103-3101010122312103-3301230223330230-1330120033312121-0113003020023333-3132131230213000-0332232221112213)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_geo_location_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032010232322230-0013330203030030-1320133333001010-2202332211102310-1111111323102223-2123101001310210-0121020202312023-0230020221232013"></a>

## Direct properties — custom_geo_location_selector / 110321122332 / 3

<a id="canonical-3311320010103221-3020220222121302-3323320011033331-1101301323200133-3102232223111310-3323200203333033-0130231130103101-2301310322210102"></a>

<a id="canonical-3220223212111302-0213231010012222-1223023221002232-0321301211321201-1013221313233230-0021333003332332-0120322220300010-2011221013321123"></a>

## expressions property — custom_geo_location_selector / 110321122332 / 4

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

<a id="canonical-1203113002023230-3130221100110332-0231212230302303-0122333000312330-0020333222131022-3001030102302233-0203211130213010-2230111120331120"></a>

## Next pages — custom_geo_location_selector / 110321122332 / 5

- [Property reference](resources--geo_location_set--reference--group-001.md#canonical-1123112132012331-2320021312120003-1120132113110332-0331323012021123-1320320333300013-2001331312211030-0311313333232210-2132322132112012)
- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-3103011222330300-2233200322023210-0200023031210213-3231332213130232-1013021130033033-2013330101030012-3301030030300101-3321313130311100)

<a id="canonical-3321203011103202-0103022333122011-1322313223010213-1110200030333021-1103232001011002-0001332310101033-3132221322222331-2122032032321002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022033222223330-2131310313332321-1303330201032131-1123110110102233-3102123020000123-3202322320132011-3011332300321012-1232030002332113"></a>

## global — global / 333331033301 / 2

Breadcrumbs:

- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-3103011222330300-2233200322023210-0200023031210213-3231332213130232-1013021130033033-2013330101030012-3301030030300101-3321313130311100)
- [Property reference](resources--geo_location_set--reference--group-001.md#canonical-1123112132012331-2320021312120003-1120132113110332-0331323012021123-1320320333300013-2001331312211030-0311313333232210-2132322132112012)
- global

<a id="canonical-3113223103231012-0231223131322103-3101010122312103-3301230223330230-1330120033312121-0113003020023333-3132131230213000-0332232221112213"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
global = {}
```

<a id="canonical-0133213322333112-2110231333221302-0000131022110222-3300002210202323-1301303120021103-3230033301033213-3121212103320303-3211110122200322"></a>

## Direct properties — global / 333331033301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202313303201012-3202120223312033-0210333213210100-0120303102102013-2103230201213110-0200303011210003-1320220113013002-3110133231012113"></a>

## Next pages — global / 333331033301 / 4

- [Property reference](resources--geo_location_set--reference--group-001.md#canonical-1123112132012331-2320021312120003-1120132113110332-0331323012021123-1320320333300013-2001331312211030-0311313333232210-2132322132112012)
- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-3103011222330300-2233200322023210-0200023031210213-3231332213130232-1013021130033033-2013330101030012-3301030030300101-3321313130311100)

<a id="canonical-3131320102322030-1221310213222021-0222131000323103-2033130021100111-0010212112013320-2302231122113231-2033321211333003-3013032012213122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023333032222330-0332201232322011-2302102021023100-3000202023023333-3131310322102021-3311031322230201-2122200012102220-1122300120323023"></a>

## timeouts — timeouts / 232110020012 / 2

Breadcrumbs:

- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-3103011222330300-2233200322023210-0200023031210213-3231332213130232-1013021130033033-2013330101030012-3301030030300101-3321313130311100)
- [Property reference](resources--geo_location_set--reference--group-001.md#canonical-1123112132012331-2320021312120003-1120132113110332-0331323012021123-1320320333300013-2001331312211030-0311313333232210-2132322132112012)
- timeouts

<a id="canonical-0131002021300030-2131111220002030-1000221322123320-1100323232023200-1122113312302011-2030032101301110-0232133323210220-3213032232313333"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223032210133022-2113133220330213-0122301332222121-0022200202211221-1111200313232320-1112100331330011-3310231101213120-1222002213232013"></a>

## Direct properties — timeouts / 232110020012 / 3

<a id="canonical-2220212121301212-0212212021223122-0112322300003103-1113010210232022-1231101331003200-2031101121032202-0110122121100310-2310313312211022"></a>

<a id="canonical-2213301203302002-3321030321033132-2201212132103312-2212323000021313-0102310122003331-0222110220032231-0211132103202120-0012122122033120"></a>

## create property — timeouts / 232110020012 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2310002033200201-3032233232011112-2223013312211103-0300023220132130-1210232101120020-2321132021312220-3212303200103123-2333231113310120"></a>

<a id="canonical-2321320331222102-1012212031101121-0230200120231302-1122101220010333-2123232311302033-2301113320310303-1210132033213122-0001210223331202"></a>

## delete property — timeouts / 232110020012 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1131300023110321-3211111333100012-3120200213013110-0312301320030112-1122022333123000-0301132221023313-1112332130123121-0322131033121303"></a>

<a id="canonical-1133231331133013-0003023213313333-0013302333103233-3123222130310111-0003013010232120-3332130110030301-3012202032123231-1130131221003310"></a>

## read property — timeouts / 232110020012 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0032030033211010-2123211231311110-2102002101133132-0320003110110230-1202302002220213-3301121323022103-1012023223121230-0310030331133233"></a>

<a id="canonical-3322113031202303-3221201231102001-0133230322032031-1333330303003120-0200313012011011-2012330210203122-1102020213302203-0320321032213001"></a>

## update property — timeouts / 232110020012 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2101220111201003-3303202331221202-1331313101221232-2333332301130032-0102120010130302-2112200031100110-0302112300121223-2212123010012233"></a>

## Next pages — timeouts / 232110020012 / 8

- [Property reference](resources--geo_location_set--reference--group-001.md#canonical-1123112132012331-2320021312120003-1120132113110332-0331323012021123-1320320333300013-2001331312211030-0311313333232210-2132322132112012)
- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-3103011222330300-2233200322023210-0200023031210213-3231332213130232-1013021130033033-2013330101030012-3301030030300101-3321313130311100)
