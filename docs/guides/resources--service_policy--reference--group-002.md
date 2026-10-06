---
page_title: "xcsh_service_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy reference."
---

# xcsh_service_policy reference

<a id="canonical-3013111121202120-1033033012323330-1020021320132301-3202002003210211-0333123103122320-1100213133023033-2021101000133111-2002101210222223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.asn_matcher](resources--service_policy--reference--group-001.md#canonical-0232211031132121-3002321133223230-3303210013302032-0323302103003332-0131313120223121-3031312302333221-3012100033212200-2323100313231113)
- rule_list.rules.spec.asn_matcher.asn_sets

<a id="canonical-3002022331323320-1223010221333030-1232321121220102-1211213203301123-1221001020311203-3030131120033132-1000220012130230-3230210113102001"></a>

Type: `"object"`. list nested block, Optional.

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111121131133321-2030320112122331-3302031102310222-2101022332102021-0332230232133030-0232233023322203-0210321002110022-2322111101220230"></a>

### Direct properties for `rule_list.rules.spec.asn_matcher.asn_sets`

<a id="canonical-3013121001110201-3121330022333310-3032312310303200-0012331200121213-3133221220110312-1203313123201022-2202200002121101-3301312321100100"></a>

#### `rule_list.rules.spec.asn_matcher.asn_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2200311223230301-3211210233332022-1332110300031323-3211113022023130-3020303210120133-3030022213011200-1121100111330023-1330221113022130"></a>

<a id="canonical-2211211120130322-3321132303002232-1232112300120322-1200310213221012-3002233221231232-0210222100133022-1131203110101323-1202132011313232"></a>

#### `rule_list.rules.spec.asn_matcher.asn_sets.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2100013200133312-1110123113210012-2121213111021011-2232212322320030-2121323102011120-1301101003122301-2020232301020110-3310300120223220"></a>

<a id="canonical-1130213021110333-1310000020331333-2310322132000312-3101211211301130-1100003103100212-1202133030102020-1113030133033220-2331230222113303"></a>

#### `rule_list.rules.spec.asn_matcher.asn_sets.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-0013301213123311-3032032323110202-3101321112022300-3003312132130330-1221001000130020-1203012033102022-1330233111322330-0331110030212132"></a>

<a id="canonical-3021022201311300-3332130011122120-1201132031103210-2332232021322303-0230100302022310-0111030310031301-1113122121320032-0122300322000230"></a>

#### `rule_list.rules.spec.asn_matcher.asn_sets.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2231001300032230-3220021301232210-3002222031133020-2113030200022222-0103100022213302-1033133311103012-0032233201233312-1220303210321122"></a>

<a id="canonical-2133012021323331-2310200302331212-2211102110303011-1130011110012322-2323113200301123-2112033233330020-2231310301200211-1300223303213303"></a>

#### `rule_list.rules.spec.asn_matcher.asn_sets.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0030321002132020-0123311113120033-3011301123233201-0303332212102210-0332111031220132-2100030113212110-1130111021302232-1020213123301102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.body_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.body_matcher

<a id="canonical-1132300332130030-2212330302311231-3031330013112203-1300311030020032-1210003020030112-0232210320320300-1002111312131231-2212011203311320"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
body_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222101330332232-1110322330001023-3011313002311130-0001013311001021-0323122022011101-2021112221112300-1101210011001022-3331321100130001"></a>

### Direct properties for `rule_list.rules.spec.body_matcher`

<a id="canonical-0320100001331130-1330310130320122-2102032302221013-3013103111030331-3013033112332111-3010202220222202-2120321210131101-2322032310122201"></a>

#### `rule_list.rules.spec.body_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0323120321021113-0323311220323013-0112312110100202-3312332112103333-0033013213101021-3300202013000110-1032300002001333-0211220323230131"></a>

<a id="canonical-1232320101201221-2322110002210102-1232210301302022-2011310121030312-0133012320312030-1203311111323330-0133223103111022-3021221310103322"></a>

#### `rule_list.rules.spec.body_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0003222100230331-2202133133320111-3233030222221100-1331032111201332-3010220213030320-2323133101300130-1313130233322020-0233002100223010"></a>

<a id="canonical-2101332330212113-3300331210212221-1322303333221223-3300320132101002-3002031001113312-2312220030101302-3202211223003010-0100102201211333"></a>

#### `rule_list.rules.spec.body_matcher.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0302200302211013-0331210113301001-2003331301230333-2332131010012033-3111300123223102-2333001133220102-0110113210223131-3201213110101121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.bot_action` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.bot_action

<a id="canonical-0033121133202131-2133212303020203-1332123113310000-2021312303302022-3213001013120202-3102223120100231-3133123201011300-2210030303112131"></a>

Type: `"object"`. single nested block, Optional.

Modify Bot protection behavior for a matching request. The modification could be to entirely skip
Bot processing.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("bot_skip_processing",
    "none")}
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
  "x-ves-oneof-field-action_type": "[\"bot_skip_processing\",\"none\"]"
}
```

Terraform syntax:

```terraform
bot_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303321231011003-3131332103231221-0233232212032210-3231000322031031-2100221122300210-0212020101213011-3122013211101132-2330311211303020"></a>

### Direct properties for `rule_list.rules.spec.bot_action`

- [bot_skip_processing](resources--service_policy--reference--group-002.md#canonical-3232001302313223-1320003300222130-2223200232110131-0023132132122032-1031323101022022-0012231023123131-0001232121213210-0301322332310032): complete subsection reference.

- [none](resources--service_policy--reference--group-002.md#canonical-3320201122303030-2231320201121033-0120021203101012-3202003122023221-1001333123311221-0201220033131321-0232321221033130-0031330123012212): complete subsection reference.

<a id="canonical-3232001302313223-1320003300222130-2223200232110131-0023132132122032-1031323101022022-0012231023123131-0001232121213210-0301322332310032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.bot_action.bot_skip_processing` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.bot_action](resources--service_policy--reference--group-002.md#canonical-0302200302211013-0331210113301001-2003331301230333-2332131010012033-3111300123223102-2333001133220102-0110113210223131-3201213110101121)
- rule_list.rules.spec.bot_action.bot_skip_processing

<a id="canonical-2312132211322012-3021201303020020-0031003231301221-1100002310333210-0000033012303022-3212122033101021-0221131330330032-3022033213132302"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
bot_skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320201122303030-2231320201121033-0120021203101012-3202003122023221-1001333123311221-0201220033131321-0232321221033130-0031330123012212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.bot_action.none` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.bot_action](resources--service_policy--reference--group-002.md#canonical-0302200302211013-0331210113301001-2003331301230333-2332131010012033-3111300123223102-2333001133220102-0110113210223131-3201213110101121)
- rule_list.rules.spec.bot_action.none

<a id="canonical-0130312213301320-2303112020231313-2031231323301301-3320330131120203-1011130332030030-3113013321133122-2033223110113023-2230021331111330"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013211000212013-2212232120132201-3311130211322132-3320030201030231-3310233022020230-1023302200222133-1112120002313123-1031200213012010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.client_name_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.client_name_matcher

<a id="canonical-1122221102210123-2032002011013332-0232112002321231-2130223110233202-0212232233323101-2230300132312311-3232023001011201-0201231012221233"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
client_name_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012323331333013-1330021032320213-3210301003101112-0232022222132203-1100311301010031-1101213322203221-3300300333312101-1301322111221222"></a>

### Direct properties for `rule_list.rules.spec.client_name_matcher`

<a id="canonical-0311313132010320-3200301001320000-2210112133000211-2230023302301130-0022132003301233-3101330031313130-3133010102113213-1301033023321103"></a>

#### `rule_list.rules.spec.client_name_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2100013303002233-3020112023211021-2313103033111200-0012132012031301-3331130020310001-2013210032012302-2002112321220200-1120030123031011"></a>

<a id="canonical-2331211211021331-1010012201221200-3313302033302200-2221310123103131-0311023020002231-2023003032131102-3312121222033023-0231300302200303"></a>

#### `rule_list.rules.spec.client_name_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0320300110001300-2133011003200221-0202030200323133-1210302233302020-2330320021322130-3312002201203120-0220012120133111-3023331012012020"></a>

<a id="canonical-3211122100323212-1102222301233332-2031312121023132-0323321321012012-1220203330303031-3002113232200202-2110212121203203-3103001113333030"></a>

#### `rule_list.rules.spec.client_name_matcher.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0211201021232311-0030301003210011-2310000123031310-0130033112310121-3032233333123100-0131113122312203-3011123320102131-0203220020213233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.client_selector` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.client_selector

<a id="canonical-2003111002023221-1103220223332023-3013333202303122-2231213020002130-3001012030303012-0233321330221322-0300321113203130-2002033233332211"></a>

Type: `"object"`. single nested block, Optional.

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
EnumExtractionComplete: false
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130003321132130-1120313023201011-2001300100200203-1213011320103222-1202120003302212-3120012022203322-2200001130300120-1110113332033203"></a>

### Direct properties for `rule_list.rules.spec.client_selector`

<a id="canonical-1332221020302120-2032210223023031-0130011330132321-0313130301312022-0020233102232332-0003222310303112-0212110011302122-1302002100102120"></a>

#### `rule_list.rules.spec.client_selector.expressions` property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3121212120020200-2301203330120221-0103311213020123-2031130122012111-2001303101221200-3022101002012312-1221331031322310-2110303302010321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.cookie_matchers` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.cookie_matchers

<a id="canonical-2112011002113230-2211013332221222-2300302112132333-2010300033210300-3231300000303200-2212222112010320-0332213213332311-1230023222023331"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321023022003313-3001201231212112-0331230312322313-1220302323103122-0123130221022322-3312233323333100-3220312102302201-2012312200322202"></a>

### Direct properties for `rule_list.rules.spec.cookie_matchers`

- [check_not_present](resources--service_policy--reference--group-002.md#canonical-1210221003132022-1120023320332100-1220203213011133-2202131330310022-2311223303132311-2230000333010201-2203221133102000-3220313010310132): complete subsection reference.

- [check_present](resources--service_policy--reference--group-002.md#canonical-1003120310230002-1230102100023000-2103333033320033-0301221212102112-3102210221303121-1023033210030333-1301112332101222-3303120332303022): complete subsection reference.

<a id="canonical-3213331130021230-0233233123032033-0332121002120023-3120323302002330-1231030211322023-3211301330203001-1232303033303132-0103133022103122"></a>

<a id="canonical-1201012330300010-3312210020122210-0021010110311320-2000123000321213-3100312230021230-3333232032202313-3300223001222322-2021022233122032"></a>

#### `rule_list.rules.spec.cookie_matchers.invert_matcher` property

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

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

- [item](resources--service_policy--reference--group-002.md#canonical-1003122112323313-2200130212313233-1002302002001323-3110023221302013-3023323021013013-1103200302030111-1132311030311232-2233023113001020): complete subsection reference.

<a id="canonical-1311021021022212-1020220132233122-3333023203332220-2312210312123113-2011211110230313-2232311321203310-3121210010200133-2332120313123100"></a>

<a id="canonical-0322030132032231-1112320231113323-3231120202011333-0320023132310000-3223211021100230-0123333133301322-0221033311303033-1302010121123112"></a>

#### `rule_list.rules.spec.cookie_matchers.name` property

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1210221003132022-1120023320332100-1220203213011133-2202131330310022-2311223303132311-2230000333010201-2203221133102000-3220313010310132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.cookie_matchers](resources--service_policy--reference--group-002.md#canonical-3121212120020200-2301203330120221-0103311213020123-2031130122012111-2001303101221200-3022101002012312-1221331031322310-2110303302010321)
- rule_list.rules.spec.cookie_matchers.check_not_present

<a id="canonical-2100331203021313-1020101120103220-3230123331101200-1123000203222120-0303120230100112-0302022130001311-3320312132213033-2022032213003132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Additional upstream details:

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
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003120310230002-1230102100023000-2103333033320033-0301221212102112-3102210221303121-1023033210030333-1301112332101222-3303120332303022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.cookie_matchers](resources--service_policy--reference--group-002.md#canonical-3121212120020200-2301203330120221-0103311213020123-2031130122012111-2001303101221200-3022101002012312-1221331031322310-2110303302010321)
- rule_list.rules.spec.cookie_matchers.check_present

<a id="canonical-3032100303112312-2220033032301231-3210201221223331-1022100012113020-3133023112100331-0000022003010301-0211000322202023-0103122200213112"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Additional upstream details:

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
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003122112323313-2200130212313233-1002302002001323-3110023221302013-3023323021013013-1103200302030111-1132311030311232-2233023113001020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.cookie_matchers](resources--service_policy--reference--group-002.md#canonical-3121212120020200-2301203330120221-0103311213020123-2031130122012111-2001303101221200-3022101002012312-1221331031322310-2110303302010321)
- rule_list.rules.spec.cookie_matchers.item

<a id="canonical-3231132302133321-0200311113212233-3301231031112302-1133202000230333-3030302012032311-3121012010202020-3202322100300212-3302132301001302"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230223010011230-1111001211011211-2110313223033132-1331002300012232-1222103310323020-2201012031133011-3313231123011201-1002303110302123"></a>

### Direct properties for `rule_list.rules.spec.cookie_matchers.item`

<a id="canonical-0022331112020300-3012021202230030-1131121223123110-3203233321233110-3020302003313110-2202331300312023-2300131302132200-0120201000223113"></a>

#### `rule_list.rules.spec.cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2310122021011012-3122221312330131-1231233222131212-1131300221002023-2232202002123311-0120011011010030-3212231030001302-0311002211301303"></a>

<a id="canonical-1123302002011222-1110222222301302-3213012213132001-1012233232331103-1001101010313003-1101230233203023-3312310030102011-1331132033212011"></a>

#### `rule_list.rules.spec.cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3111123013322003-0031333203013332-0032213333211330-3312202011012102-1003100121233100-2312001000310201-3120223013023132-2300031311233223"></a>

<a id="canonical-0230030111023001-3300313301313303-3233023303333003-0203331031020131-1012001022113322-3123100103311233-3103322331001110-2322133330303133"></a>

#### `rule_list.rules.spec.cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0002120030013310-3002023311102221-2221120010233003-0112023021102232-2110300301110330-3222223323210112-1002023003211130-1212030032102211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.domain_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.domain_matcher

<a id="canonical-3013030122221221-0321233330320223-3223222031223102-2010233030133031-3010100103013111-3021230133303300-2203123232133122-3131020003001022"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
domain_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303211231102022-3002033303123300-0201311022322131-1032223030232323-2301221011132201-2220200000320232-1010303033201102-1332231300113223"></a>

### Direct properties for `rule_list.rules.spec.domain_matcher`

<a id="canonical-3202131312112222-1323211210231002-0111002131012332-0203010203303122-3102310030202222-2320300310300132-0201022100220013-3120011321113133"></a>

#### `rule_list.rules.spec.domain_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2213010313213022-0210311331222022-1323013113310031-3200011223132010-0203322313022031-3022301101210032-1302231030210132-2321230332202102"></a>

<a id="canonical-1003012100121013-3330220110232012-0020210032001203-1121012021210313-2300322200003201-1023221331223332-2330121320323233-0031232311031122"></a>

#### `rule_list.rules.spec.domain_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3301301003231322-3300012022231202-0203302333223321-2013030110222013-3021213232310311-3103102131033032-1300212331312220-1013110010320323"></a>

<a id="canonical-2023201212222230-2332132103110230-1012000122233110-0000212003311222-3130301000230213-2300021212003332-0022133213100203-3323221003133112"></a>

#### `rule_list.rules.spec.domain_matcher.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3303211112230231-3131220013211212-3131022320033101-1130301232233213-0301112303322203-0023113002110131-0203122003122231-0230111120022301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.headers` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.headers

<a id="canonical-0213120231022333-1331110113330110-3302230212201103-0212013302220011-3200303113333210-0202311002002032-0113132010112303-1022031230212220"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1312230113221112-0230320231130121-2002112313323001-2130301101101223-0231332133120011-0021222033301300-1330203210320123-0231303031333023"></a>

### Direct properties for `rule_list.rules.spec.headers`

- [check_not_present](resources--service_policy--reference--group-002.md#canonical-1113301033011203-1102303031023211-2123003110101102-3103300010103110-1012001221103020-2002301313333033-0112230320303331-0123220333133121): complete subsection reference.

- [check_present](resources--service_policy--reference--group-002.md#canonical-3320221012111201-3311303210100010-1121223312103301-1000001033122300-3230020201100321-1331022321312120-2032100210201033-1001113331020121): complete subsection reference.

<a id="canonical-0220313230022012-3103100122130010-1022001130120230-2112001330102211-0330213101233313-0313213301211200-1110001300003012-3000230233032322"></a>

<a id="canonical-1031022130330322-2011202222011032-1110013123012313-2331313310331203-0112332103222222-3230113212101212-2210300031113130-0030202222133132"></a>

#### `rule_list.rules.spec.headers.invert_matcher` property

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

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

- [item](resources--service_policy--reference--group-002.md#canonical-3011120033300333-2202312110320331-2333310303331223-3131132202120321-3130203232320011-0021123023032301-2223033131210122-3033231112211201): complete subsection reference.

<a id="canonical-1131322331000101-2232300020230331-3102210220030021-2201111033321100-1211023321111013-1202022121330103-1011302012131202-1120012312032201"></a>

<a id="canonical-0131111020320333-0230102122203301-3201010011330002-0031101203312220-1131012302132301-1032211211023202-2031332113322031-0312012030322023"></a>

#### `rule_list.rules.spec.headers.name` property

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1113301033011203-1102303031023211-2123003110101102-3103300010103110-1012001221103020-2002301313333033-0112230320303331-0123220333133121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.headers](resources--service_policy--reference--group-002.md#canonical-3303211112230231-3131220013211212-3131022320033101-1130301232233213-0301112303322203-0023113002110131-0203122003122231-0230111120022301)
- rule_list.rules.spec.headers.check_not_present

<a id="canonical-3222332233120023-0031033121231313-1002303013323121-0213131030313002-3332130123220130-1023010232300312-3230321321102030-3020012302110000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Additional upstream details:

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
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320221012111201-3311303210100010-1121223312103301-1000001033122300-3230020201100321-1331022321312120-2032100210201033-1001113331020121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.headers.check_present` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.headers](resources--service_policy--reference--group-002.md#canonical-3303211112230231-3131220013211212-3131022320033101-1130301232233213-0301112303322203-0023113002110131-0203122003122231-0230111120022301)
- rule_list.rules.spec.headers.check_present

<a id="canonical-1333113230013120-2200102130022102-1031331222030131-2211231023120313-3103001032231210-3223321012102121-1020320111211232-3212212130300100"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Additional upstream details:

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
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011120033300333-2202312110320331-2333310303331223-3131132202120321-3130203232320011-0021123023032301-2223033131210122-3033231112211201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.headers.item` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.headers](resources--service_policy--reference--group-002.md#canonical-3303211112230231-3131220013211212-3131022320033101-1130301232233213-0301112303322203-0023113002110131-0203122003122231-0230111120022301)
- rule_list.rules.spec.headers.item

<a id="canonical-1312310200301223-3210102303211310-3223031201102033-0210102021001002-3321012000031100-0133310132133223-1013201001110201-0330311223023101"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232110303020113-0023303320110222-2120300200323231-3201130300021301-1313100313003331-1223322330031303-2311310220111321-1003010123112001"></a>

### Direct properties for `rule_list.rules.spec.headers.item`

<a id="canonical-3123333112013232-3100030201003232-0230110130123223-2003332302032213-1221231111110210-1131200223312001-0333100200333213-2231220331232230"></a>

#### `rule_list.rules.spec.headers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3302120023100222-1113133221031002-3320330201233103-0222302302320212-0220320031020100-3200133232320011-1222001321333302-3312033300203013"></a>

<a id="canonical-1123122202100320-3331313301321231-2221202102310111-1233101021332032-3013300203022220-0202211130302301-0313021112302011-3132201232322011"></a>

#### `rule_list.rules.spec.headers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2223200003130022-2033321033031100-1220133221033333-0002312210313303-1200232300323300-3330110100322222-3122201003010133-3200323000322130"></a>

<a id="canonical-3230101100133032-0112120012100100-2212022222002102-3032130320131323-2031000301022320-0231302302030332-0332303111133223-2330131232303202"></a>

#### `rule_list.rules.spec.headers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3200303221303321-0332123202022111-0032113310031110-3333200123232130-3321331111121012-1030202033031002-1030130232323200-1112120101333313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.http_method` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.http_method

<a id="canonical-1021002231113020-0212311202231001-0023222012220312-0121323112302230-2012230031310122-1013021011230333-3002300311231320-2033033231313113"></a>

Type: `"object"`. single nested block, Optional.

An HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

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
http_method {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011010112330200-2223331113133310-1223120121122001-2132030201313131-1310102223020112-2022321332202103-2011113320232122-1301133223010222"></a>

### Direct properties for `rule_list.rules.spec.http_method`

<a id="canonical-3000121122320203-1013130121101033-1102012310033200-2222300232000333-0213002111102211-3132232230221211-0123133323020330-2133231313120211"></a>

#### `rule_list.rules.spec.http_method.invert_matcher` property

Type: `"bool"`. Optional.

Invert Method Matcher. Invert the match result.

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

<a id="canonical-0331131330322111-0232323132211210-0320012133202000-1020121000233122-1011320022212313-0300000310001231-2121123210200213-1121000201330322"></a>

<a id="canonical-3102330202020112-3211102132011013-0131323001110303-1132012101121213-2303302133102310-2131022313000333-1122302210203201-0213323202230313"></a>

#### `rule_list.rules.spec.http_method.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1030111002320303-0232312311303320-3110022123122122-1021111312200023-0113312211223231-0101120320103223-1232331122313210-1312313222312210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.ip_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.ip_matcher

<a id="canonical-2300301111202111-2100311132323011-1112011132022031-2110302023001212-2101210122110321-2230231001030221-0100221012300200-3322220122330323"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001313032110230-3122033123012213-1020310012321210-0121301230331312-3123022201132301-0132332221310120-3001001210201303-1111023023312102"></a>

### Direct properties for `rule_list.rules.spec.ip_matcher`

<a id="canonical-2130232212313311-1322122202121310-1332122210303230-0331203300330123-0212220331232331-0213103322030222-2122111200020002-0213130223000020"></a>

#### `rule_list.rules.spec.ip_matcher.invert_matcher` property

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

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

- [prefix_sets](resources--service_policy--reference--group-002.md#canonical-1223311302203303-0322113302231131-3020111222231300-0130330102003013-3203133003233111-0301011313232003-0012000111300303-2213011103122101): complete subsection reference.

<a id="canonical-1223311302203303-0322113302231131-3020111222231300-0130330102003013-3203133003233111-0301011313232003-0012000111300303-2213011103122101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.ip_matcher](resources--service_policy--reference--group-002.md#canonical-1030111002320303-0232312311303320-3110022123122122-1021111312200023-0113312211223231-0101120320103223-1232331122313210-1312313222312210)
- rule_list.rules.spec.ip_matcher.prefix_sets

<a id="canonical-3013331101120231-3330123011110330-1320232201033322-0201100213312011-3130110303121120-0322130232213030-1132001221222210-3120323203331031"></a>

Type: `"object"`. list nested block, Optional.

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130303110320231-1313200233232133-3133011003122323-0323012133010000-1010110231322220-2201213033122023-2133222003332213-0231333101303212"></a>

### Direct properties for `rule_list.rules.spec.ip_matcher.prefix_sets`

<a id="canonical-0022322313323221-0102131330313210-0011312323030003-1130022230010120-0111202222201122-1031121233223021-2332003332312021-0332332222033202"></a>

#### `rule_list.rules.spec.ip_matcher.prefix_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3321030231013221-2211002013033121-2230231223203300-2022321010100332-1133231220122122-3131330312033313-2223100030233102-2020210103111021"></a>

<a id="canonical-1011112212030123-2131103302031001-3323333303013331-0030333102123222-3121212000213233-1112330032221333-0310303323202312-1103210220032112"></a>

#### `rule_list.rules.spec.ip_matcher.prefix_sets.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0321112330133013-2311230200322112-0122002323111122-1121132033030223-3302321122203202-3122013203100013-0322332212111030-1103220130003030"></a>

<a id="canonical-1220022033330110-2313322321231333-0323211203011031-2302313033310231-0121112032003111-2000230331222023-1313303002322221-2001130301113100"></a>

#### `rule_list.rules.spec.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-1210031232301203-2111013202213111-3010233033313203-3233102312210113-3122221322321210-3333111200123233-0103003030221233-3210031122131000"></a>

<a id="canonical-3232233111320321-3203033313033303-3333213331103102-0300223033023030-2023012132100210-0133112001303322-3000013323322233-0111111111021010"></a>

#### `rule_list.rules.spec.ip_matcher.prefix_sets.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0112120110031132-1220100221221130-3202122120123313-1231330213103130-0303103133121012-0212301032011021-0102232222032200-2322202200221300"></a>

<a id="canonical-3332303033010322-1032131133022322-3303202230300123-0322100223023003-0032200002231113-0122110330222312-0302330332101212-0313123012113022"></a>

#### `rule_list.rules.spec.ip_matcher.prefix_sets.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1320110103211332-0010132210330323-2221011311200322-1022311233301132-3210220223000122-3203222200021011-1130101332123213-1211103331231223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.ip_prefix_list

<a id="canonical-2312122221023221-3203212000322101-0313001101033121-0310111103332002-0030010013303333-0331300011022321-1311311213320200-2030123213321312"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

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
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221330210001302-0012310232211311-2231110101033312-1321230313012003-1302331200321212-1023103110302211-2320213210103332-1320002303130321"></a>

### Direct properties for `rule_list.rules.spec.ip_prefix_list`

<a id="canonical-2203220011330210-0231203013231030-0013312020001002-2113113203301221-0100322313210312-1022113202001123-1310120213231220-0122322230302033"></a>

#### `rule_list.rules.spec.ip_prefix_list.invert_match` property

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

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

<a id="canonical-0023233302301232-3333103103323230-2113200133113000-2122332230133231-0010022013223011-2011023002311111-2100333331210112-0121223013020001"></a>

<a id="canonical-3230213313330333-0020323031132200-0030212223312033-1031102332332120-2121112122223220-3221033111010123-3021032312210203-0132023201003331"></a>

#### `rule_list.rules.spec.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3333211202130202-0133211023331201-3301313020033000-0233020133130232-2320303320330222-2010201313212003-1110123210010000-1101102021012033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.ip_threat_category_list

<a id="canonical-1312030330312331-3321100332010323-2022233012332322-1000100030230000-3101010320232101-3012122020033300-1021022213303332-0010322032221300"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102312323233031-0232230000213011-3320223000230233-0232122213012310-2030333221012022-0113020000120120-1122300322231333-0012100331002322"></a>

### Direct properties for `rule_list.rules.spec.ip_threat_category_list`

<a id="canonical-3300001222002020-2101003021130010-2212031201222311-2203211023222302-3113331330303330-0123121112032203-2001323113003113-0300110211312133"></a>

#### `rule_list.rules.spec.ip_threat_category_list.ip_threat_categories` property

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3030030233300323-2201301300112210-1201003030313223-2320311230330003-2013232332000022-0131000212013130-1121021010010202-1011100231001020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.ja4_tls_fingerprint` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.ja4_tls_fingerprint

<a id="canonical-2200331211100333-2110313001203310-0011223222311212-2333230223322132-2302012110023130-1103003310303213-3012210111302003-3303132020211123"></a>

Type: `"object"`. single nested block, Optional.

An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

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
ja4_tls_fingerprint {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301131320112103-3031111323233221-0201033320213010-0200120230131133-2103320102312323-3122200221331112-3200312233200202-1123300100300220"></a>

### Direct properties for `rule_list.rules.spec.ja4_tls_fingerprint`

<a id="canonical-0100320003130232-2231313322333333-0032032221301002-0122022312300202-2032001001031022-3330011203233003-0203210320112111-2133203202333131"></a>

#### `rule_list.rules.spec.ja4_tls_fingerprint.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0232333201113100-3110011212311231-2210020330303010-1321112033033300-2312211113223103-1101003302101330-3313322102011022-2013130320332331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.jwt_claims` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.jwt_claims

<a id="canonical-1000201331321112-2103012033023310-1122030000010121-1000300302013312-1101332122310000-3011223330020211-2021133023212232-0032231221311110"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131000223021233-2021012130223321-2213002331022220-3003110011003312-3212302333212221-2300332112101202-0333213221031110-1232333331201231"></a>

### Direct properties for `rule_list.rules.spec.jwt_claims`

- [check_not_present](resources--service_policy--reference--group-002.md#canonical-0021122330202312-0020302110013030-2121213320033001-0110023313120123-2003200022313221-2113331020222131-0103011113122321-2212311301030023): complete subsection reference.

- [check_present](resources--service_policy--reference--group-002.md#canonical-2020113331113322-3302201102313012-3031211032131112-2011030321312230-1103313112300220-2113203202100033-0320310332131332-0221012211031101): complete subsection reference.

<a id="canonical-0031133331031031-1321010301330100-1033020011302112-1222022013131011-1131010201320211-0303212012003320-3122332313200113-2022233012011331"></a>

<a id="canonical-2323203332012132-1010211333223100-1122003132123332-1010201023121021-0203131311312130-1000111031120101-2020103233203220-0233311201233323"></a>

#### `rule_list.rules.spec.jwt_claims.invert_matcher` property

Type: `"bool"`. Optional.

Invert Matcher. Invert the match result.

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

- [item](resources--service_policy--reference--group-002.md#canonical-2122100032203300-1321230301312201-2010032311323100-2302332323301221-1200102121002132-3322330301300300-1303033111002002-1321033212312303): complete subsection reference.

<a id="canonical-3223111323113330-2112123222131122-1303322201210103-2310033012221230-1131111012101103-0110302310223221-1220320200033112-0013131232301120"></a>

<a id="canonical-2122203221310310-3013303300031331-2013313023301120-1032203031000010-0313103310302033-3303103102020102-0300001211013133-2203200203310322"></a>

#### `rule_list.rules.spec.jwt_claims.name` property

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0021122330202312-0020302110013030-2121213320033001-0110023313120123-2003200022313221-2113331020222131-0103011113122321-2212311301030023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.jwt_claims](resources--service_policy--reference--group-002.md#canonical-0232333201113100-3110011212311231-2210020330303010-1321112033033300-2312211113223103-1101003302101330-3313322102011022-2013130320332331)
- rule_list.rules.spec.jwt_claims.check_not_present

<a id="canonical-1000010230030210-0322122332332233-3101302030202331-3013001000010232-1203032210033133-2231202001333102-1003311130101020-1001233001201322"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Additional upstream details:

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
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020113331113322-3302201102313012-3031211032131112-2011030321312230-1103313112300220-2113203202100033-0320310332131332-0221012211031101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.jwt_claims](resources--service_policy--reference--group-002.md#canonical-0232333201113100-3110011212311231-2210020330303010-1321112033033300-2312211113223103-1101003302101330-3313322102011022-2013130320332331)
- rule_list.rules.spec.jwt_claims.check_present

<a id="canonical-1333301220120303-2000323100113200-2003313311313303-3231302232010213-2231120310133101-1331200031000031-0122032003132103-0123113230223321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Additional upstream details:

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
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122100032203300-1321230301312201-2010032311323100-2302332323301221-1200102121002132-3322330301300300-1303033111002002-1321033212312303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.jwt_claims](resources--service_policy--reference--group-002.md#canonical-0232333201113100-3110011212311231-2210020330303010-1321112033033300-2312211113223103-1101003302101330-3313322102011022-2013130320332331)
- rule_list.rules.spec.jwt_claims.item

<a id="canonical-1210213111000001-1123321320023333-2123323131200330-1312013303103312-0100303013001310-0200321010222023-0333103010311331-2032220233001202"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201133301012213-1022303023101130-2310110213103331-0210101333103103-1121312201110000-0311001012021200-0133000031303311-0013220011002021"></a>

### Direct properties for `rule_list.rules.spec.jwt_claims.item`

<a id="canonical-1213220120020323-0033203133311133-0302233311123220-0202110233102231-1222132012112023-2111121033302131-1302313213301113-0011113101003112"></a>

#### `rule_list.rules.spec.jwt_claims.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3332201012122120-1223101313313001-1233310312022310-3331032030220313-1030201333301231-3231321322121010-2120112132030103-1020103021130222"></a>

<a id="canonical-2332130231113122-3323032110201102-1332331300332211-2030101132202011-3201310112000223-0023133232013320-1301313113101012-0210231233300301"></a>

#### `rule_list.rules.spec.jwt_claims.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2013011113222133-3201301120112011-1223202113311111-3321003003133203-2230322133123313-3110030033320103-0313003202210230-1132331012323222"></a>

<a id="canonical-1123233331113033-2112000002032201-1131203332111213-0102213311311211-1131322311103112-1101200233202213-3212013321223223-0012131233000320"></a>

#### `rule_list.rules.spec.jwt_claims.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3232100012122002-1210310013101032-0310322320030203-3203213320013100-3120100320020220-0100032322102131-1013203100213003-2110320203300311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.label_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.label_matcher

<a id="canonical-2231321110123331-2012213321010303-2303111210130201-1332001113000330-3100302101222213-0301023003010033-1201110222203211-2203003013121223"></a>

Type: `"object"`. single nested block, Optional.

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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
label_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233112010222213-3222130200310312-0033110323023122-1001122003321031-0001012120211030-0310123132233133-1030211303303211-0222033320022101"></a>

### Direct properties for `rule_list.rules.spec.label_matcher`

<a id="canonical-2322031213102022-0121220000021030-3322103300103101-3200223230103012-3230320110330131-0020130331230330-0012221311211031-0323311112121300"></a>

#### `rule_list.rules.spec.label_matcher.keys` property

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1031203103211101-1320130203110223-1203222000022323-0200122233332223-3330121000123133-2233010210331321-1120111112011131-2031122333013131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.mum_action` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.mum_action

<a id="canonical-3130111322110332-1003313330213002-0202120002122032-1332322232212331-1130231322331111-0210221103222231-0222131231130223-1022023311230020"></a>

Type: `"object"`. single nested block, Optional.

Modify behavior for a matching request. The modification could be to entirely skip processing.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default",
    "skip_processing")}
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
  "x-ves-oneof-field-action_type": "[\"default\",\"skip_processing\"]"
}
```

Terraform syntax:

```terraform
mum_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230000213012200-0020003302212333-1203000220003213-0030133221033220-3312013030301222-0123212120013212-0232220013010132-2111311003212020"></a>

### Direct properties for `rule_list.rules.spec.mum_action`

- [default](resources--service_policy--reference--group-002.md#canonical-3303232231013311-0311112021113232-0213101120023021-2003110022121232-1302332203231212-1212111030100200-3003223133331313-3322211310320010): complete subsection reference.

- [skip_processing](resources--service_policy--reference--group-002.md#canonical-3110033032111322-1210310112302101-2322300321312232-3323220030301122-3130001023123331-1033230213022032-3112320113210013-3301212333233012): complete subsection reference.

<a id="canonical-3303232231013311-0311112021113232-0213101120023021-2003110022121232-1302332203231212-1212111030100200-3003223133331313-3322211310320010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.mum_action.default` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.mum_action](resources--service_policy--reference--group-002.md#canonical-1031203103211101-1320130203110223-1203222000022323-0200122233332223-3330121000123133-2233010210331321-1120111112011131-2031122333013131)
- rule_list.rules.spec.mum_action.default

<a id="canonical-1203210221011221-1231322030122330-3011213013320322-1023213231300112-1113310122001202-2323212221133313-1033121310302010-0321210112030102"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3110033032111322-1210310112302101-2322300321312232-3323220030301122-3130001023123331-1033230213022032-3112320113210013-3301212333233012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.mum_action.skip_processing` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.mum_action](resources--service_policy--reference--group-002.md#canonical-1031203103211101-1320130203110223-1203222000022323-0200122233332223-3330121000123133-2233010210331321-1120111112011131-2031122333013131)
- rule_list.rules.spec.mum_action.skip_processing

<a id="canonical-1220022123030333-3220201200011133-2213221212020322-2032213023303312-3031203203210220-2303231103121320-0011231230023221-2033211322012222"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100030012203210-1100201310311210-1001120230313212-2000131113131300-3023102000302131-3003123100333233-0213221032313101-0303033021301201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.path` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.path

<a id="canonical-3103230311331210-0300023330003333-3212322321221123-2330313301220312-2230332311233331-3120022111220210-2201010211121101-2013010022000210"></a>

Type: `"object"`. single nested block, Optional.

A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

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
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003101203012300-2300131221200003-2221103122213220-1120130113220201-1112102232321223-2113232323312310-2020203012311331-2120132322211111"></a>

### Direct properties for `rule_list.rules.spec.path`

<a id="canonical-3222322321200301-1332010330122220-3021301301320111-0312300000331213-1023212310013120-0320203012331023-1012121200223120-3123313333013101"></a>

#### `rule_list.rules.spec.path.encoded_path_matcher` property

Type: `"bool"`. Optional.

Match against the encoded, escaped path.

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

<a id="canonical-3210333102012033-3131323100311030-3001002321213112-2201333211232032-3301321110102133-3020122200011001-3022321002333323-2230302032021222"></a>

<a id="canonical-2011333333322202-2202221013103310-1001300202132123-1021020321222321-3232000302022202-2233302311122232-3212103031203300-1103000233003233"></a>

#### `rule_list.rules.spec.path.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact path values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1003023133300002-1120000233122310-1330220211201110-1030331331010323-2102222332231122-1323222323101020-2101031011223333-1222202321301221"></a>

<a id="canonical-1310300312033301-1212331311112300-3123012132333210-3321001212233112-2023031301023033-2101302222210302-3203332000232130-3301212301110233"></a>

#### `rule_list.rules.spec.path.invert_matcher` property

Type: `"bool"`. Optional.

Invert Path Matcher. Invert the match result.

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

<a id="canonical-2331002311030320-2230313213123023-1013222322211130-3312031331210313-3113130000023231-2333120323220101-2332320333311213-3013221130100331"></a>

<a id="canonical-1330302132003222-3122030310310231-0130010002301300-2011223011110203-1120331000232301-3321033030023021-0010202302010010-3201000232130132"></a>

#### `rule_list.rules.spec.path.prefix_values` property

Type: `["list", "string"]`. Optional.

A list of path prefix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1123122013100312-0311103102303323-1300000223011211-3202333313123230-3103231110302321-0220112133321301-2023002102011032-2121231032221013"></a>

<a id="canonical-2310000030131012-1223111230020232-3103100022232111-1100310333132210-3102120321121113-3311121222202231-0111323002211122-2230011321301023"></a>

#### `rule_list.rules.spec.path.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0100023222031213-3010233320022000-2001302010201112-3032010123310202-0023300220222220-1303230021021010-1213122032002022-1122300330213003"></a>

<a id="canonical-2111233002223203-1120223013101320-2121012223322121-1123310021201320-2100102111212231-3023033232020110-2023103031130231-2203032000032310"></a>

#### `rule_list.rules.spec.path.suffix_values` property

Type: `["list", "string"]`. Optional.

A list of path suffix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1230231102330211-2311300033000311-0322220020110023-0110303301233101-1222020333212303-2123300022021211-0322001210320332-2230013113320212"></a>

<a id="canonical-1231231323020032-1202131100132231-1222131130223223-3201011131110032-3300331220213011-1003102322100003-3032023000223200-1020100123131002"></a>

#### `rule_list.rules.spec.path.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3032102120213313-0312100211232033-3112223313112201-0031231231213331-3302120020022021-0000132323100020-3012123121223223-2302111000133001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.port_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.port_matcher

<a id="canonical-0322300223131313-2202330032010311-1302231231102213-2323120213102321-3210322230000132-0221322233331332-1201020132320013-1220332013322130"></a>

Type: `"object"`. single nested block, Optional.

Port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true. Server applies default when omitted.

Additional upstream details:

A port matcher specifies a list of port ranges as match criteria.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
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
port_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030103211211201-2022002303031002-0113302110030212-1231130222122021-3013122033112221-2110211100332100-1331102300302033-2212212331202321"></a>

### Direct properties for `rule_list.rules.spec.port_matcher`

<a id="canonical-1102222233230023-2003132031021323-1331003211012101-0110330011132130-2321311033100330-2113121032312211-0133000132023030-3203000001230311"></a>

#### `rule_list.rules.spec.port_matcher.invert_matcher` property

Type: `"bool"`. Optional.

Invert Port Matcher. Invert the match result.

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

<a id="canonical-2120332231000300-3331130013300021-3313323011221211-1320223000113102-1331220121321220-0000100103311022-1311001313033331-0121123032101302"></a>

<a id="canonical-2230111001133311-3030031221312212-2332300113113101-1123011001121113-0010100120131321-3013233001032223-2020202232133001-3231030002200201"></a>

#### `rule_list.rules.spec.port_matcher.ports` property

Type: `["list", "string"]`. Optional.

List of strings, each of which is a single port value or a tuple of start and end port values
separated by '-'. The start and end values are considered to be part of the range.

Additional upstream details:

A list of strings, each of which is a single port value or a tuple of start and end port values
separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2020223332102111-0033130221100121-2321221112100331-3032233101002012-0013223201032330-1010113011301300-2103120333221013-1121231031201231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.query_params` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.query_params

<a id="canonical-3313303312311300-0033303303001212-3310323232203103-0021112110221120-0121003322120210-0231003023330323-2330221000013103-1220103001003120"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113232210203101-0332232102131013-2120020322030111-2033231232212010-1232211311322310-2213330311310330-3021203103021033-1230311210231300"></a>

### Direct properties for `rule_list.rules.spec.query_params`

- [check_not_present](resources--service_policy--reference--group-002.md#canonical-3201111130122200-1332321030100132-2310213320212111-3322223023312103-0230322302011201-3132021323311123-0121123030322210-0100201303003022): complete subsection reference.

- [check_present](resources--service_policy--reference--group-002.md#canonical-2010302320222233-0330002332120132-1323100010012213-3003031011211100-0220000000131232-2221203330301212-2301020010322321-2013201102031230): complete subsection reference.

<a id="canonical-2000211101313102-1233220222331033-2003322310110202-3111112301010102-2012023233330220-3021122310023100-2021321111021233-1232033030210312"></a>

<a id="canonical-3222133002200301-3020000030323302-2230000111331030-2122012120030033-3320011112020220-2013321310010112-1303213213012111-1200333133200003"></a>

#### `rule_list.rules.spec.query_params.invert_matcher` property

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](resources--service_policy--reference--group-002.md#canonical-0311031230301231-0322130312023310-0020223230031100-0320313212223123-2310203303201331-0030332103031301-2101000020012022-2330302301021003): complete subsection reference.

<a id="canonical-3022020032001113-1221011220231301-0212001222333231-1111010230033102-2103310321133333-0103323013312032-0033303032320332-0021120233131111"></a>

<a id="canonical-2131221133020120-3113330120322120-3003133322202123-2123311023200010-3122232021033201-1323101130213000-3032300003223322-3133333100301233"></a>

#### `rule_list.rules.spec.query_params.key` property

Type: `"string"`. Optional.

A case-sensitive HTTP query parameter name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3201111130122200-1332321030100132-2310213320212111-3322223023312103-0230322302011201-3132021323311123-0121123030322210-0100201303003022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.query_params](resources--service_policy--reference--group-002.md#canonical-2020223332102111-0033130221100121-2321221112100331-3032233101002012-0013223201032330-1010113011301300-2103120333221013-1121231031201231)
- rule_list.rules.spec.query_params.check_not_present

<a id="canonical-2213311101101232-1311032112333231-1032121203300013-1302132020312221-1203331211003003-1221331121130103-1300110311112211-3202110200122122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Additional upstream details:

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
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010302320222233-0330002332120132-1323100010012213-3003031011211100-0220000000131232-2221203330301212-2301020010322321-2013201102031230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.query_params.check_present` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.query_params](resources--service_policy--reference--group-002.md#canonical-2020223332102111-0033130221100121-2321221112100331-3032233101002012-0013223201032330-1010113011301300-2103120333221013-1121231031201231)
- rule_list.rules.spec.query_params.check_present

<a id="canonical-3320122212110131-3133133000022033-0012111013212211-2330221333212002-3302320300303030-3033112132321121-0332211011122123-3320112223302312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Additional upstream details:

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
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311031230301231-0322130312023310-0020223230031100-0320313212223123-2310203303201331-0030332103031301-2101000020012022-2330302301021003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.query_params.item` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.query_params](resources--service_policy--reference--group-002.md#canonical-2020223332102111-0033130221100121-2321221112100331-3032233101002012-0013223201032330-1010113011301300-2103120333221013-1121231031201231)
- rule_list.rules.spec.query_params.item

<a id="canonical-1202132022201203-0312121113221100-3102011200212320-2210112330323031-3101111001102213-2233332203130333-2211103201021122-2103032021111002"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300200011011100-3203213133211130-3133013123132000-3321113330321303-3011120132122300-0132321333003220-3022011120133013-0111232020100020"></a>

### Direct properties for `rule_list.rules.spec.query_params.item`

<a id="canonical-2102110330320230-1312321310120202-1232201112111230-3220031122121120-3323010001110100-2223120220023210-3312130020112121-2030033320002211"></a>

#### `rule_list.rules.spec.query_params.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0022222102111200-2031211300222203-1032201022322230-1233203213131222-0020020022103322-0121311000311001-1131011203011131-0111103322033111"></a>

<a id="canonical-0111320201132330-2010300333222230-1220203321220102-0210333131003311-0321311330323012-0130302032132123-2000330310332230-2200133012301222"></a>

#### `rule_list.rules.spec.query_params.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3033131211220301-3103011222321111-2202301031012003-3310110130302321-0221311210203002-1012320022321132-1300010230002033-0101210201120220"></a>

<a id="canonical-1232333212121203-3033120100231013-3320022233201333-3130210210300001-3130313012202230-0103001332001222-0102033121113011-2221330213222300"></a>

#### `rule_list.rules.spec.query_params.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.request_constraints

<a id="canonical-3333302322101323-0201111313220103-0020221122010103-1212101111101001-2023210312120022-1132203131330102-3000302301311000-1110312003233023"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request constraints.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_cookie_count_exceeds",
    "max_cookie_count_none"),
  validators.ConflictingObjectAttributes("max_cookie_key_size_exceeds",
    "max_cookie_key_size_none"),
  validators.ConflictingObjectAttributes("max_cookie_value_size_exceeds",
    "max_cookie_value_size_none"),
  validators.ConflictingObjectAttributes("max_header_count_exceeds",
    "max_header_count_none"),
  validators.ConflictingObjectAttributes("max_header_key_size_exceeds",
    "max_header_key_size_none"),
  validators.ConflictingObjectAttributes("max_header_value_size_exceeds",
    "max_header_value_size_none"),
  validators.ConflictingObjectAttributes("max_parameter_count_exceeds",
    "max_parameter_count_none"),
  validators.ConflictingObjectAttributes("max_parameter_name_size_exceeds",
    "max_parameter_name_size_none"),
  validators.ConflictingObjectAttributes("max_parameter_value_size_exceeds",
    "max_parameter_value_size_none"),
  validators.ConflictingObjectAttributes("max_query_size_exceeds",
    "max_query_size_none"),
  validators.ConflictingObjectAttributes("max_request_line_size_exceeds",
    "max_request_line_size_none"),
  validators.ConflictingObjectAttributes("max_request_size_exceeds",
    "max_request_size_none"),
  validators.ConflictingObjectAttributes("max_url_size_exceeds",
    "max_url_size_none")}
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
  "x-ves-oneof-field-max_cookie_count_choice": "[\"max_cookie_count_exceeds\",\"max_cookie_count_none\"]",
  "x-ves-oneof-field-max_cookie_key_size_choice": "[\"max_cookie_key_size_exceeds\",\"max_cookie_key_size_none\"]",
  "x-ves-oneof-field-max_cookie_value_size_choice": "[\"max_cookie_value_size_exceeds\",\"max_cookie_value_size_none\"]",
  "x-ves-oneof-field-max_header_count_choice": "[\"max_header_count_exceeds\",\"max_header_count_none\"]",
  "x-ves-oneof-field-max_header_key_size_choice": "[\"max_header_key_size_exceeds\",\"max_header_key_size_none\"]",
  "x-ves-oneof-field-max_header_value_size_choice": "[\"max_header_value_size_exceeds\",\"max_header_value_size_none\"]",
  "x-ves-oneof-field-max_parameter_count_choice": "[\"max_parameter_count_exceeds\",\"max_parameter_count_none\"]",
  "x-ves-oneof-field-max_parameter_name_size_choice": "[\"max_parameter_name_size_exceeds\",\"max_parameter_name_size_none\"]",
  "x-ves-oneof-field-max_parameter_value_size_choice": "[\"max_parameter_value_size_exceeds\",\"max_parameter_value_size_none\"]",
  "x-ves-oneof-field-max_query_size_choice": "[\"max_query_size_exceeds\",\"max_query_size_none\"]",
  "x-ves-oneof-field-max_request_line_size_choice": "[\"max_request_line_size_exceeds\",\"max_request_line_size_none\"]",
  "x-ves-oneof-field-max_request_size_choice": "[\"max_request_size_exceeds\",\"max_request_size_none\"]",
  "x-ves-oneof-field-max_url_size_choice": "[\"max_url_size_exceeds\",\"max_url_size_none\"]"
}
```

Terraform syntax:

```terraform
request_constraints {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200001232023233-2330323330320301-2002333102111112-1120000311211333-0130200111003020-1110210012303310-0200203023132022-0111131033133133"></a>

### Direct properties for `rule_list.rules.spec.request_constraints`

<a id="canonical-3203200013110120-3210122032320032-1023211320310003-3201203132002321-0302323111131223-2320111221303221-2221132020310313-2312002030002020"></a>

#### `rule_list.rules.spec.request_constraints.max_cookie_count_exceeds` property

Type: `"number"`. Optional.

Match on the Count for all Cookies that exceed this value. Exclusive with
\[max\_cookie\_count\_none\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_cookie_count_none](resources--service_policy--reference--group-002.md#canonical-0030222301212332-0012300010222202-0111102300021200-2223301132322123-1133320011022321-0333102223220110-1332121101012231-0323022211010233): complete subsection reference.

<a id="canonical-0333323330233211-2131021103331213-3001003020031221-2031010111122201-2313223133223103-1332011330000301-2223131111330000-0102332013303223"></a>

<a id="canonical-2210323101021103-0002020010300311-0313313330121021-0121202233000003-2313022103031232-1122300212202113-2002003123011323-0231302312011001"></a>

#### `rule_list.rules.spec.request_constraints.max_cookie_key_size_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_cookie\_key\_size\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_cookie_key_size_none](resources--service_policy--reference--group-002.md#canonical-1113321011231222-2012220023330033-0012333012221100-3322133231023203-0201000310120310-2301110200320120-3022232232130212-3033303123222012): complete subsection reference.

<a id="canonical-0130201012313000-0021311132313200-0111023020000112-1332313212211123-3331001003310231-0031333113000223-2221012302000012-2100200110113201"></a>

<a id="canonical-2303230010322130-2001032222011311-3212323300313321-1120331310011321-2020132232200203-0111010003322331-0200133312330011-3112120022203132"></a>

#### `rule_list.rules.spec.request_constraints.max_cookie_value_size_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_cookie\_value\_size\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

- [max_cookie_value_size_none](resources--service_policy--reference--group-002.md#canonical-2003211303022323-3332223010332321-2301222030020233-1021330312213310-0011101303003300-0102320320312332-1102220311121300-0113023322332102): complete subsection reference.

<a id="canonical-1112210022221221-2313233002012221-0102022213231231-2203123232012013-0213021221201311-0132133013310021-0033023123323021-2303100300203331"></a>

<a id="canonical-3101232320322022-0303011213221303-2020302322102301-1133103130233300-2011122033301223-3302213211110212-2201000330001010-1013013213321312"></a>

#### `rule_list.rules.spec.request_constraints.max_header_count_exceeds` property

Type: `"number"`. Optional.

Match on the Count for all Headers that exceed this value. Exclusive with
\[max\_header\_count\_none\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "40"
  }
}
```

- [max_header_count_none](resources--service_policy--reference--group-002.md#canonical-1012312000301123-0220212010011032-0323103313020203-1113220000101222-2121211311220110-0302232302310222-2000233321101131-0130211130000100): complete subsection reference.

<a id="canonical-0113100021100122-3032110310322311-3233101200221210-1221111302321200-3033011113120202-2102011103133213-2322321202320202-3311302303122313"></a>

<a id="canonical-1213131000031233-0131201031010020-3222001202102330-1000313222103233-1210301221132003-2023120131222133-0132301002013123-3121323330131002"></a>

#### `rule_list.rules.spec.request_constraints.max_header_key_size_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_header\_key\_size\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_header_key_size_none](resources--service_policy--reference--group-002.md#canonical-3002331231230012-1002010202101030-1222010232332013-2202012002201021-1112100102110312-2032002121103332-2211110131031022-0302321222202111): complete subsection reference.

<a id="canonical-1111023220010330-2302131021033122-0212100322130101-0220303212301003-2220211011220330-1110013010010122-1111220311103113-3331211221113100"></a>

<a id="canonical-3203230312110102-3022103012333323-1110221021022012-3330012211112321-1112102103222102-1001010011033232-0232202213110310-1022303301300011"></a>

#### `rule_list.rules.spec.request_constraints.max_header_value_size_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_header\_value\_size\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 64000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [max_header_value_size_none](resources--service_policy--reference--group-002.md#canonical-0211210213001101-2133201323222011-2320223111203302-3021002322100332-1211023310210001-2133120232030322-0003301221033113-3133303200132203): complete subsection reference.

<a id="canonical-3213101212031103-1033001312322000-2111321122111110-0032202123332323-0103032121212313-0323112200122003-0133302222232213-1222201312220130"></a>

<a id="canonical-0310130333030002-3302103201121231-3313102211132021-0111022030233011-0331003000032211-0302310300031111-2011121023320210-3203120002023121"></a>

#### `rule_list.rules.spec.request_constraints.max_parameter_count_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_count\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_parameter_count_none](resources--service_policy--reference--group-002.md#canonical-3101302230010022-0013031312110133-0321121220133022-2311312302112322-3331023130332301-1133320301212222-3221320321311301-2031031023133113): complete subsection reference.

<a id="canonical-3301222111120221-3303003013202123-3021012001303030-0231023311211002-1022033323021001-3030002100123320-3030120130332111-3030111122111102"></a>

<a id="canonical-2112300210332002-3301003200132103-0120231210012002-0202203221113211-3101321232313130-2022011021312301-3210221330112211-2301301311333332"></a>

#### `rule_list.rules.spec.request_constraints.max_parameter_name_size_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_name\_size\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_parameter_name_size_none](resources--service_policy--reference--group-002.md#canonical-0002301100011231-2122311113330121-0311221202132100-1302333221302203-2230111002112032-1020033103213333-3132010001232303-0302120113120032): complete subsection reference.

<a id="canonical-1332233100201321-3321120120321310-1010000310013023-1132033123201110-1323232202003300-0000103302212012-1012111203010120-1310233330332123"></a>

<a id="canonical-3310310021010212-3331330120231013-1332203131312103-0310313101000031-3233203333301122-3312012010102030-3210101030232120-3031120012201311"></a>

#### `rule_list.rules.spec.request_constraints.max_parameter_value_size_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_value\_size\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 1073741824),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1073741824,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1073741824"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1073741824"
  }
}
```

- [max_parameter_value_size_none](resources--service_policy--reference--group-002.md#canonical-3200323212233321-0033011302233121-2203302022132212-1120122210313321-2333213120030122-2020000123022320-1220121001203310-3202300320101201): complete subsection reference.

<a id="canonical-2313203203310333-0021121100033120-2210320220010032-0231103221333313-3302023232323011-0111120231033211-2203202023222332-3210022022020113"></a>

<a id="canonical-3313132022300120-2120031023031320-1122210213223101-2332203233330311-1031231202023231-2322333332131301-3320122031303031-3003230311121321"></a>

#### `rule_list.rules.spec.request_constraints.max_query_size_exceeds` property

Type: `"number"`. Optional.

Match on the URL Query Size that exceed this value. Exclusive with \[max\_query\_size\_none\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

- [max_query_size_none](resources--service_policy--reference--group-002.md#canonical-1013020222111032-2123302100312203-3232302320201033-1112132112021122-3311120131020020-1002132011310023-2100113111330221-0312301301200102): complete subsection reference.

<a id="canonical-1111333020123223-2003023013001313-1303133323210310-2213012202322203-1023032101301031-0222220020330311-0302102121012331-1311320300300331"></a>

<a id="canonical-0032112122332112-3011321101332031-3230220332030202-0231233223213330-1212201211133002-2211002131330313-3220133231210023-3300323020010123"></a>

#### `rule_list.rules.spec.request_constraints.max_request_line_size_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_request\_line\_size\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_line_size_none](resources--service_policy--reference--group-002.md#canonical-1312203332130002-0010103331023220-3230131321200133-3302332030033232-2130320000212332-2310302032332113-1122211012123030-3031233122111230): complete subsection reference.

<a id="canonical-3033133010232202-0232321002220220-0112102333333230-0121110000112013-1100330221323020-0231312323002010-0013000002310332-3303001211110031"></a>

<a id="canonical-0100130131311001-1020110001300311-3002223113123330-2311310003000113-3320113303103303-2333110013013113-2313132321220121-2012001023231123"></a>

#### `rule_list.rules.spec.request_constraints.max_request_size_exceeds` property

Type: `"number"`. Optional.

Match on the Request Size that exceed this value. Exclusive with \[max\_request\_size\_none\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_size_none](resources--service_policy--reference--group-002.md#canonical-2211003200031203-2030000102310111-3010311010013203-1310200022020121-2331102101212302-0231323030121333-1013000131130122-1120012210021030): complete subsection reference.

<a id="canonical-3112020102101133-0131033313220321-2110030302222023-0012111102232002-0222232132322312-1022020002330202-2223032032030221-0100130332013213"></a>

<a id="canonical-0220221032220322-2002231301011000-0310010122132210-1021331323301223-2111110301011200-0302211012223213-2222032013220200-1131333323121233"></a>

#### `rule_list.rules.spec.request_constraints.max_url_size_exceeds` property

Type: `"number"`. Optional.

Match on the URL Size that exceed this value. Exclusive with \[max\_url\_size\_none\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 128000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128000"
  }
}
```

- [max_url_size_none](resources--service_policy--reference--group-002.md#canonical-0001112012202010-3112323211321020-1010031231321302-2032201031020002-0032313122112232-3132020222101132-0103010031123132-3110203212320130): complete subsection reference.

<a id="canonical-0030222301212332-0012300010222202-0111102300021200-2223301132322123-1133320011022321-0333102223220110-1332121101012231-0323022211010233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_cookie_count_none` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_cookie_count_none

<a id="canonical-2033331210013202-1301200033023110-3033023210233200-0302110203221311-1220113021303013-0220112200201323-1202133232320022-3201301133031123"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie count none.

Additional upstream details:

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
max_cookie_count_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113321011231222-2012220023330033-0012333012221100-3322133231023203-0201000310120310-2301110200320120-3022232232130212-3033303123222012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_cookie_key_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_cookie_key_size_none

<a id="canonical-1322302200033123-1131111212300303-2032233013003123-0321101002232202-2111302012322112-3322311310310010-0322012121223303-3211120220000103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie key size none.

Additional upstream details:

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
max_cookie_key_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003211303022323-3332223010332321-2301222030020233-1021330312213310-0011101303003300-0102320320312332-1102220311121300-0113023322332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_cookie_value_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_cookie_value_size_none

<a id="canonical-3100112101300020-0201221301112022-2302330012220132-2211112023333231-2133310233130032-0031201202130210-2012202300310311-0210003230001023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie value size none.

Additional upstream details:

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
max_cookie_value_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012312000301123-0220212010011032-0323103313020203-1113220000101222-2121211311220110-0302232302310222-2000233321101131-0130211130000100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_header_count_none` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_header_count_none

<a id="canonical-3231131322130322-2322120113213002-3323331321003102-3233000130312032-3302111222211022-0000322033033110-3133033133030003-0330320331322213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header count none.

Additional upstream details:

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
max_header_count_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002331231230012-1002010202101030-1222010232332013-2202012002201021-1112100102110312-2032002121103332-2211110131031022-0302321222202111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_header_key_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_header_key_size_none

<a id="canonical-2311321222221211-2202013013113133-1210132321200222-2333313033032131-1212032303301333-1203320020310031-0003112033312031-1002332221110221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header key size none.

Additional upstream details:

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
max_header_key_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211210213001101-2133201323222011-2320223111203302-3021002322100332-1211023310210001-2133120232030322-0003301221033113-3133303200132203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_header_value_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_header_value_size_none

<a id="canonical-0200202312023131-3221023120000003-1022133301231332-1032131001023223-0300111321312032-3102330222211120-0111230313322321-0320000122002211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header value size none.

Additional upstream details:

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
max_header_value_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101302230010022-0013031312110133-0321121220133022-2311312302112322-3331023130332301-1133320301212222-3221320321311301-2031031023133113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_parameter_count_none` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_parameter_count_none

<a id="canonical-1011003033220003-2231200213222112-1301323310323030-2201010130331310-3202122333003000-2000311032021302-2213203011033310-2332210032333022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max parameter count none.

Additional upstream details:

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
max_parameter_count_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002301100011231-2122311113330121-0311221202132100-1302333221302203-2230111002112032-1020033103213333-3132010001232303-0302120113120032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_parameter_name_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_parameter_name_size_none

<a id="canonical-1002200331203103-1221101103003231-3231221331020301-1310323010120031-3331230220333111-0203001303130220-3310300221310330-0100310333320332"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
max_parameter_name_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200323212233321-0033011302233121-2203302022132212-1120122210313321-2333213120030122-2020000123022320-1220121001203310-3202300320101201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_parameter_value_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_parameter_value_size_none

<a id="canonical-1232133101130103-2132120111110022-1130120331112013-2303013230120131-3111332322200001-3002103222212012-1320033011210231-3221001232333330"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max parameter value size none.

Additional upstream details:

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
max_parameter_value_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013020222111032-2123302100312203-3232302320201033-1112132112021122-3311120131020020-1002132011310023-2100113111330221-0312301301200102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_query_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_query_size_none

<a id="canonical-2220223202000320-2313110303213033-1011313210202330-0032032201011313-0033131310303033-2321033102112233-1132323233001102-1103103023332223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max query size none.

Additional upstream details:

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
max_query_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312203332130002-0010103331023220-3230131321200133-3302332030033232-2130320000212332-2310302032332113-1122211012123030-3031233122111230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_request_line_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_request_line_size_none

<a id="canonical-1321210222121100-2232102302000320-1201232200012221-1022201300033120-0212223323111031-1012300122330001-3122002210110322-2300103212232203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max request line size none.

Additional upstream details:

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
max_request_line_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211003200031203-2030000102310111-3010311010013203-1310200022020121-2331102101212302-0231323030121333-1013000131130122-1120012210021030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_request_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_request_size_none

<a id="canonical-1203310032203211-3200210000223202-1202030320200112-1133330220010203-1310011202111322-3320130300131030-1020200313321202-1223211320331332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max request size none.

Additional upstream details:

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
max_request_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001112012202010-3112323211321020-1010031231321302-2032201031020002-0032313122112232-3132020222101132-0103010031123132-3110203212320130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_url_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_url_size_none

<a id="canonical-3331101010322200-1032332023221003-1013000330313113-2302113330213310-2333323333113030-0222003132132002-1122031332000212-3101320130031302"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
max_url_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.segment_policy

<a id="canonical-3123012000130222-0021031033231122-1021210223003202-0321331233310102-0310003132123300-2220102021201131-2300100313011321-0220311300302121"></a>

Type: `"object"`. single nested block, Optional.

Configure source and destination segment for policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("dst_any",
    "dst_segments"),
  validators.ConflictingObjectAttributes("dst_any",
    "intra_segment"),
  validators.ConflictingObjectAttributes("dst_segments",
    "intra_segment"),
  validators.ConflictingObjectAttributes("src_any",
    "src_segments")}
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
  "x-ves-oneof-field-dst_segment_choice": "[\"dst_any\",\"dst_segments\",\"intra_segment\"]",
  "x-ves-oneof-field-src_segment_choice": "[\"src_any\",\"src_segments\"]"
}
```

Terraform syntax:

```terraform
segment_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031130201113220-3120312123313120-0003312323202121-0202111120221221-0301231023102202-2110021132103313-1002003001101332-2001033020332130"></a>

### Direct properties for `rule_list.rules.spec.segment_policy`

- [dst_any](resources--service_policy--reference--group-002.md#canonical-3002010211101110-3221323322013110-0000111130011312-2121102330232021-1223010230013233-0010033023203121-3203130003222121-1320110231313030): complete subsection reference.

- [dst_segments](resources--service_policy--reference--group-002.md#canonical-2023102300100100-0010302223303330-1300202232022033-0103023130302032-1023012301210211-3320130333031013-2033320130101133-1131333303332132): complete subsection reference.

- [intra_segment](resources--service_policy--reference--group-002.md#canonical-3201323020022131-3013300300300130-0212302313301301-2300301302210303-2021000303132231-0132322023231112-0122112233202333-1000013300223122): complete subsection reference.

- [src_any](resources--service_policy--reference--group-002.md#canonical-0013302203311202-1010220003121332-3210230121232312-3002023203002320-1200113032330123-3222112032312123-2032333002210113-3312101313022300): complete subsection reference.

- [src_segments](resources--service_policy--reference--group-002.md#canonical-1321111033030303-3021331113201220-2102331103121003-1222300031003013-3100131131110121-1301223332213100-1032130232201220-1032222211221031): complete subsection reference.

<a id="canonical-3002010211101110-3221323322013110-0000111130011312-2121102330232021-1223010230013233-0010033023203121-3203130003222121-1320110231313030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy.dst_any` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-002.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- rule_list.rules.spec.segment_policy.dst_any

<a id="canonical-1012233013201321-0113011332231330-0302110123312323-3023220002210313-2113120201231301-1120132132033200-2332013221313031-1323113322120203"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
dst_any = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023102300100100-0010302223303330-1300202232022033-0103023130302032-1023012301210211-3320130333031013-2033320130101133-1131333303332132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy.dst_segments` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-002.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- rule_list.rules.spec.segment_policy.dst_segments

<a id="canonical-0132200022132130-3300330010322312-1032121303001121-2130111312331302-3102000010312323-0333100201100331-0310212123312312-3110131320102302"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dst segments.

Additional upstream details:

List of references to Segments.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("segments")}
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
dst_segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330132333330010-0110122311313220-1022333303122132-0020133313202310-3323212112103201-2020230310201033-3323003320201122-2303310011200113"></a>

### Direct properties for `rule_list.rules.spec.segment_policy.dst_segments`

- [segments](resources--service_policy--reference--group-002.md#canonical-0131100222231033-1122330332010133-3131222232012322-3313011122311303-1221010121213112-1311310110321100-0331212101032032-3220003021113223): complete subsection reference.

<a id="canonical-0131100222231033-1122330332010133-3131222232012322-3313011122311303-1221010121213112-1311310110321100-0331212101032032-3220003021113223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy.dst_segments.segments` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-002.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- [rule_list.rules.spec.segment_policy.dst_segments](resources--service_policy--reference--group-002.md#canonical-2023102300100100-0010302223303330-1300202232022033-0103023130302032-1023012301210211-3320130333031013-2033320130101133-1131333303332132)
- rule_list.rules.spec.segment_policy.dst_segments.segments

<a id="canonical-2332110131120102-2220301302310211-0031031223031213-2301301133122223-3320123101012320-2323001302332220-1030120300230221-2321030013100123"></a>

Type: `"object"`. list nested block, Optional.

Segments. Select list of segments.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
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

Terraform syntax:

```terraform
segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200222232002200-3312122010223311-0010200130310023-3113213111131301-0101130102312133-2312113101210031-2112331122003232-2013310122333112"></a>

### Direct properties for `rule_list.rules.spec.segment_policy.dst_segments.segments`

<a id="canonical-0020203332031121-1111112312110101-0031013003001013-1311213202121032-0333310200300331-1121113103202301-1222220021311031-3303233113312001"></a>

#### `rule_list.rules.spec.segment_policy.dst_segments.segments.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0123001233200023-3120321003310010-2022330322100231-0233220300022322-2111201332131331-0210300013121022-1222021130130331-2101123330331311"></a>

<a id="canonical-3032013133122012-3213211231131320-3212112221212220-1122020320210311-3333001110110121-2112110323212130-0211323233301120-2223100202232031"></a>

#### `rule_list.rules.spec.segment_policy.dst_segments.segments.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2200112020130223-3310233030031100-3100232310033212-2101310320220112-3211333332312110-1120003311113321-2233321320310103-2022111302131002"></a>

<a id="canonical-0300233130103300-3223131200000103-1302210203030132-2321200323302333-0112222103011303-0030131230031310-0102031033130332-3111113233230013"></a>

#### `rule_list.rules.spec.segment_policy.dst_segments.segments.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3201323020022131-3013300300300130-0212302313301301-2300301302210303-2021000303132231-0132322023231112-0122112233202333-1000013300223122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy.intra_segment` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-002.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- rule_list.rules.spec.segment_policy.intra_segment

<a id="canonical-1210033002102202-1012321302012001-1221021120312231-0210110320120122-1301110110203120-1220333211010130-1313010111312021-0223323120003321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for intra segment.

Additional upstream details:

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
intra_segment = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013302203311202-1010220003121332-3210230121232312-3002023203002320-1200113032330123-3222112032312123-2032333002210113-3312101313022300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy.src_any` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-002.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- rule_list.rules.spec.segment_policy.src_any

<a id="canonical-1021110201210013-2301132113303322-1212121222233110-1330302030322332-0131130002000303-1022331220210313-3023002332330031-2330102232022201"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
src_any = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321111033030303-3021331113201220-2102331103121003-1222300031003013-3100131131110121-1301223332213100-1032130232201220-1032222211221031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy.src_segments` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-002.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- rule_list.rules.spec.segment_policy.src_segments

<a id="canonical-3332231220130101-1011320302021122-1303123303132300-0311302332120211-2132013011112123-2210133031301012-2210112211011003-0032013310233112"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for src segments.

Additional upstream details:

List of references to Segments.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("segments")}
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
src_segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233113012030020-0001020133010232-0100311110300132-2223003310232103-2212101101120030-1122132301010213-1001313122120220-1303013323131301"></a>

### Direct properties for `rule_list.rules.spec.segment_policy.src_segments`

- [segments](resources--service_policy--reference--group-002.md#canonical-2122210222323033-3030011311312001-3323232201220221-1201303302003330-0001201301100210-3122330110301321-2012032221200013-1303001132201130): complete subsection reference.

<a id="canonical-2122210222323033-3030011311312001-3323232201220221-1201303302003330-0001201301100210-3122330110301321-2012032221200013-1303001132201130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy.src_segments.segments` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-002.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- [rule_list.rules.spec.segment_policy.src_segments](resources--service_policy--reference--group-002.md#canonical-1321111033030303-3021331113201220-2102331103121003-1222300031003013-3100131131110121-1301223332213100-1032130232201220-1032222211221031)
- rule_list.rules.spec.segment_policy.src_segments.segments

<a id="canonical-0102023111220023-0000302012323012-1121022023130210-1303022321020333-2130213231100210-1301213210011131-2201020001200002-3003302202213232"></a>

Type: `"object"`. list nested block, Optional.

Segments. Select list of segments.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
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

Terraform syntax:

```terraform
segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222312301011220-3022010103010313-0321211123010330-3312023031313123-0201032223222310-2112213332012121-3233133001332112-3201000110300133"></a>

### Direct properties for `rule_list.rules.spec.segment_policy.src_segments.segments`

<a id="canonical-0332131131300213-0301113323313102-2320303301300302-3311002230202122-0120031101330203-2121121022333231-2001120002132013-2220122002122233"></a>

#### `rule_list.rules.spec.segment_policy.src_segments.segments.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0220020133112222-1003310100122112-0010320333012211-0212320110000111-3203133101030311-2011303103301130-3333100202203011-1233013013210321"></a>

<a id="canonical-1102321110131121-2321203101322100-3031312330133022-0333301301212320-3121032332332030-0113012333122123-1132320031312011-2121030103033003"></a>

#### `rule_list.rules.spec.segment_policy.src_segments.segments.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1112111121033002-0220100101230130-3312331303110222-3302102001312012-2222300323302333-1300203332303132-2112021123121200-0212101311120103"></a>

<a id="canonical-0321121111113130-2332203020031102-1133310033333233-0301002131331033-0012320201233231-1322330131232002-3031132231032222-0312302111220201"></a>

#### `rule_list.rules.spec.segment_policy.src_segments.segments.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1020320201300000-0002323231201110-2333023122032003-3301321123001032-2102310103222032-0133000133300130-0123200131201022-1013312312033202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-2211302223202211-0133312332123122-3001120111203221-2001331203120000-2122122033200113-0203222022113013-3310203232101331-2101130030200002"></a>

Type: `"object"`. single nested block, Optional.

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310220233100323-2031002020323100-0301112030022121-2103133221201021-2302030231312033-1321321021110020-3132010313320320-2012312131222211"></a>

### Direct properties for `rule_list.rules.spec.tls_fingerprint_matcher`

<a id="canonical-0331220121023231-0021112103101300-0031313222021121-3210131000332312-2331212321201033-0321332301010021-1201203310300313-0303030100323313"></a>

#### `rule_list.rules.spec.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1320002201100011-3203312220113103-2312130030332322-0302001320032322-3311013320011030-1330311112132300-1323120212100113-0001120200032202"></a>

<a id="canonical-0120133112100232-1312213103230003-1110111213311331-1110032202330232-3131202301302003-0323021220032111-1323210312002203-1201311023010201"></a>

#### `rule_list.rules.spec.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0131312323300001-1110011110132013-3101210201212221-0212001023323310-3011303221022330-1220301200110033-2003201232202112-0010001321232023"></a>

<a id="canonical-0333300311213123-0231113033220203-0100232311230030-0233122313012033-0212332003122123-1333000333131311-1212122321001100-0203132001000110"></a>

#### `rule_list.rules.spec.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Optional.

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3001003323033212-0123331232232220-3131203230202012-2011331002320011-1212313031310311-0101011121113132-2102222322031110-3333021300010302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.user_identity_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.user_identity_matcher

<a id="canonical-1232010221230100-0121021000010310-2213323203322133-0302300133131212-2033103123311120-0012331103220112-1112122302102113-0011130020320302"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
user_identity_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2100233020113002-2320000002312123-0211303021030312-2302321010023012-0330222211113010-3213113030321320-3312211013301213-2302213023001023"></a>

### Direct properties for `rule_list.rules.spec.user_identity_matcher`

<a id="canonical-1321310003332203-0023133101222121-3312023221111313-3123033321102023-3112220231302223-1031201210222330-3003333102230021-0011001103133301"></a>

#### `rule_list.rules.spec.user_identity_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2033313221121303-2030322301120303-2030310003323020-0210012222022101-3002231102231012-2323020331211031-2031302201102303-0221121113033230"></a>

<a id="canonical-1130312330033203-3030310222321233-1131122231022020-2231013020301123-3032202312213120-0022023003121000-0222202321320012-3100021320330303"></a>

#### `rule_list.rules.spec.user_identity_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
