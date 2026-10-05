---
page_title: "xcsh_service_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy reference."
---

# xcsh_service_policy reference

<a id="canonical-3100333221232232-1102311111100003-1013332101302210-1203113013100103-2213003113112021-0131110331310322-2321002101313210-0023303332330211"></a>

## rule_list.rules.spec.query_params.check_present — check_present / 220321323123 / 2

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
check_present = {}
```

<a id="canonical-2123232123000201-1310111012022030-3130211130133301-2012021032101130-2313102230221030-2002200120123122-3212013230103132-2230003323301310"></a>

## Direct properties — check_present / 220321323123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132103201220002-3033320310313311-1230233311333200-2332023233033000-1122201303320313-3232133230111202-2330132210112230-1001213211133323"></a>

## Next pages — check_present / 220321323123 / 4

- [rule_list.rules.spec.query_params](resources--service_policy--reference--group-002.md#canonical-2020223332102111-0033130221100121-2321221112100331-3032233101002012-0013223201032330-1010113011301300-2103120333221013-1121231031201231)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-0311031230301231-0322130312023310-0020223230031100-0320313212223123-2310203303201331-0030332103031301-2101000020012022-2330302301021003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300200011011100-3203213133211130-3133013123132000-3321113330321303-3011120132122300-0132321333003220-3022011120133013-0111232020100020"></a>

## rule_list.rules.spec.query_params.item — item / 113233200133 / 2

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

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-0111320201132330-2010300333222230-1220203321220102-0210333131003311-0321311330323012-0130302032132123-2000330310332230-2200133012301222"></a>

## Direct properties — item / 113233200133 / 3

<a id="canonical-2102110330320230-1312321310120202-1232201112111230-3220031122121120-3323010001110100-2223120220023210-3312130020112121-2030033320002211"></a>

<a id="canonical-1232333212121203-3033120100231013-3320022233201333-3130210210300001-3130313012202230-0103001332001222-0102033121113011-2221330213222300"></a>

## exact_values property — item / 113233200133 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2123103122113200-3332101232223301-2320031010020233-2112031320301232-1312132332003130-1332121301120221-1120003203133103-0222230213120331"></a>

## regex_values property — item / 113233200133 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0200223100102111-2220100103312302-1301323220332332-3331230132113031-2020101111033020-2121133231103302-0302011003201213-2220211233133022"></a>

## transformers property — item / 113233200133 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2231201013031011-1030113110211100-2003033323002011-0002321311111033-1202322020313103-2222112313121112-2222100130110133-3200120012301333"></a>

## Next pages — item / 113233200133 / 7

- [rule_list.rules.spec.query_params](resources--service_policy--reference--group-002.md#canonical-2020223332102111-0033130221100121-2321221112100331-3032233101002012-0013223201032330-1010113011301300-2103120333221013-1121231031201231)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200001232023233-2330323330320301-2002333102111112-1120000311211333-0130200111003020-1110210012303310-0200203023132022-0111131033133133"></a>

## rule_list.rules.spec.request_constraints — request_constraints / 011200002102 / 2

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

<a id="canonical-2210323101021103-0002020010300311-0313313330121021-0121202233000003-2313022103031232-1122300212202113-2002003123011323-0231302312011001"></a>

## Direct properties — request_constraints / 011200002102 / 3

<a id="canonical-3203200013110120-3210122032320032-1023211320310003-3201203132002321-0302323111131223-2320111221303221-2221132020310313-2312002030002020"></a>

<a id="canonical-2303230010322130-2001032222011311-3212323300313321-1120331310011321-2020132232200203-0111010003322331-0200133312330011-3112120022203132"></a>

## max_cookie_count_exceeds property — request_constraints / 011200002102 / 4

Type: `"number"`. Optional.

Match on the Count for all Cookies that exceed this value. Exclusive with
\[max\_cookie\_count\_none\]

Upstream description:

Exclusive with \[max\_cookie\_count\_none\]

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [max_cookie_count_none](resources--service_policy--reference--group-003.md#canonical-0030222301212332-0012300010222202-0111102300021200-2223301132322123-1133320011022321-0333102223220110-1332121101012231-0323022211010233): complete subsection reference.

<a id="canonical-0333323330233211-2131021103331213-3001003020031221-2031010111122201-2313223133223103-1332011330000301-2223131111330000-0102332013303223"></a>

<a id="canonical-3101232320322022-0303011213221303-2020302322102301-1133103130233300-2011122033301223-3302213211110212-2201000330001010-1013013213321312"></a>

## max_cookie_key_size_exceeds property — request_constraints / 011200002102 / 5

Type: `"number"`. Optional.

Exclusive with \[max\_cookie\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_key\_size\_none\]

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [max_cookie_key_size_none](resources--service_policy--reference--group-003.md#canonical-1113321011231222-2012220023330033-0012333012221100-3322133231023203-0201000310120310-2301110200320120-3022232232130212-3033303123222012): complete subsection reference.

<a id="canonical-0130201012313000-0021311132313200-0111023020000112-1332313212211123-3331001003310231-0031333113000223-2221012302000012-2100200110113201"></a>

<a id="canonical-1213131000031233-0131201031010020-3222001202102330-1000313222103233-1210301221132003-2023120131222133-0132301002013123-3121323330131002"></a>

## max_cookie_value_size_exceeds property — request_constraints / 011200002102 / 6

Type: `"number"`. Optional.

Exclusive with \[max\_cookie\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_value\_size\_none\]

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [max_cookie_value_size_none](resources--service_policy--reference--group-003.md#canonical-2003211303022323-3332223010332321-2301222030020233-1021330312213310-0011101303003300-0102320320312332-1102220311121300-0113023322332102): complete subsection reference.

<a id="canonical-1112210022221221-2313233002012221-0102022213231231-2203123232012013-0213021221201311-0132133013310021-0033023123323021-2303100300203331"></a>

<a id="canonical-3203230312110102-3022103012333323-1110221021022012-3330012211112321-1112102103222102-1001010011033232-0232202213110310-1022303301300011"></a>

## max_header_count_exceeds property — request_constraints / 011200002102 / 7

Type: `"number"`. Optional.

Match on the Count for all Headers that exceed this value. Exclusive with
\[max\_header\_count\_none\]

Upstream description:

Exclusive with \[max\_header\_count\_none\]

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [max_header_count_none](resources--service_policy--reference--group-003.md#canonical-1012312000301123-0220212010011032-0323103313020203-1113220000101222-2121211311220110-0302232302310222-2000233321101131-0130211130000100): complete subsection reference.

<a id="canonical-0113100021100122-3032110310322311-3233101200221210-1221111302321200-3033011113120202-2102011103133213-2322321202320202-3311302303122313"></a>

<a id="canonical-0310130333030002-3302103201121231-3313102211132021-0111022030233011-0331003000032211-0302310300031111-2011121023320210-3203120002023121"></a>

## max_header_key_size_exceeds property — request_constraints / 011200002102 / 8

Type: `"number"`. Optional.

Exclusive with \[max\_header\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_key\_size\_none\]

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [max_header_key_size_none](resources--service_policy--reference--group-003.md#canonical-3002331231230012-1002010202101030-1222010232332013-2202012002201021-1112100102110312-2032002121103332-2211110131031022-0302321222202111): complete subsection reference.

<a id="canonical-1111023220010330-2302131021033122-0212100322130101-0220303212301003-2220211011220330-1110013010010122-1111220311103113-3331211221113100"></a>

<a id="canonical-2112300210332002-3301003200132103-0120231210012002-0202203221113211-3101321232313130-2022011021312301-3210221330112211-2301301311333332"></a>

## max_header_value_size_exceeds property — request_constraints / 011200002102 / 9

Type: `"number"`. Optional.

Exclusive with \[max\_header\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_value\_size\_none\]

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [max_header_value_size_none](resources--service_policy--reference--group-003.md#canonical-0211210213001101-2133201323222011-2320223111203302-3021002322100332-1211023310210001-2133120232030322-0003301221033113-3133303200132203): complete subsection reference.

<a id="canonical-3213101212031103-1033001312322000-2111321122111110-0032202123332323-0103032121212313-0323112200122003-0133302222232213-1222201312220130"></a>

<a id="canonical-3310310021010212-3331330120231013-1332203131312103-0310313101000031-3233203333301122-3312012010102030-3210101030232120-3031120012201311"></a>

## max_parameter_count_exceeds property — request_constraints / 011200002102 / 10

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_count\_none\].

Upstream description:

Exclusive with \[max\_parameter\_count\_none\]

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [max_parameter_count_none](resources--service_policy--reference--group-003.md#canonical-3101302230010022-0013031312110133-0321121220133022-2311312302112322-3331023130332301-1133320301212222-3221320321311301-2031031023133113): complete subsection reference.

<a id="canonical-3301222111120221-3303003013202123-3021012001303030-0231023311211002-1022033323021001-3030002100123320-3030120130332111-3030111122111102"></a>

<a id="canonical-3313132022300120-2120031023031320-1122210213223101-2332203233330311-1031231202023231-2322333332131301-3320122031303031-3003230311121321"></a>

## max_parameter_name_size_exceeds property — request_constraints / 011200002102 / 11

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_name\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_name\_size\_none\]

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [max_parameter_name_size_none](resources--service_policy--reference--group-003.md#canonical-0002301100011231-2122311113330121-0311221202132100-1302333221302203-2230111002112032-1020033103213333-3132010001232303-0302120113120032): complete subsection reference.

<a id="canonical-1332233100201321-3321120120321310-1010000310013023-1132033123201110-1323232202003300-0000103302212012-1012111203010120-1310233330332123"></a>

<a id="canonical-0032112122332112-3011321101332031-3230220332030202-0231233223213330-1212201211133002-2211002131330313-3220133231210023-3300323020010123"></a>

## max_parameter_value_size_exceeds property — request_constraints / 011200002102 / 12

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_value\_size\_none\]

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [max_parameter_value_size_none](resources--service_policy--reference--group-003.md#canonical-3200323212233321-0033011302233121-2203302022132212-1120122210313321-2333213120030122-2020000123022320-1220121001203310-3202300320101201): complete subsection reference.

<a id="canonical-2313203203310333-0021121100033120-2210320220010032-0231103221333313-3302023232323011-0111120231033211-2203202023222332-3210022022020113"></a>

<a id="canonical-0100130131311001-1020110001300311-3002223113123330-2311310003000113-3320113303103303-2333110013013113-2313132321220121-2012001023231123"></a>

## max_query_size_exceeds property — request_constraints / 011200002102 / 13

Type: `"number"`. Optional.

Match on the URL Query Size that exceed this value. Exclusive with \[max\_query\_size\_none\]

Upstream description:

Exclusive with \[max\_query\_size\_none\]

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [max_query_size_none](resources--service_policy--reference--group-003.md#canonical-1013020222111032-2123302100312203-3232302320201033-1112132112021122-3311120131020020-1002132011310023-2100113111330221-0312301301200102): complete subsection reference.

<a id="canonical-1111333020123223-2003023013001313-1303133323210310-2213012202322203-1023032101301031-0222220020330311-0302102121012331-1311320300300331"></a>

<a id="canonical-0220221032220322-2002231301011000-0310010122132210-1021331323301223-2111110301011200-0302211012223213-2222032013220200-1131333323121233"></a>

## max_request_line_size_exceeds property — request_constraints / 011200002102 / 14

Type: `"number"`. Optional.

Exclusive with \[max\_request\_line\_size\_none\].

Upstream description:

Exclusive with \[max\_request\_line\_size\_none\]

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [max_request_line_size_none](resources--service_policy--reference--group-003.md#canonical-1312203332130002-0010103331023220-3230131321200133-3302332030033232-2130320000212332-2310302032332113-1122211012123030-3031233122111230): complete subsection reference.

<a id="canonical-3033133010232202-0232321002220220-0112102333333230-0121110000112013-1100330221323020-0231312323002010-0013000002310332-3303001211110031"></a>

<a id="canonical-3003220312011030-3131001122120013-3212001033102213-2321020112301033-3020101303312231-3233321223331332-1313313210301003-2132120223221111"></a>

## max_request_size_exceeds property — request_constraints / 011200002102 / 15

Type: `"number"`. Optional.

Match on the Request Size that exceed this value. Exclusive with \[max\_request\_size\_none\]

Upstream description:

Exclusive with \[max\_request\_size\_none\]

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [max_request_size_none](resources--service_policy--reference--group-003.md#canonical-2211003200031203-2030000102310111-3010311010013203-1310200022020121-2331102101212302-0231323030121333-1013000131130122-1120012210021030): complete subsection reference.

<a id="canonical-3112020102101133-0131033313220321-2110030302222023-0012111102232002-0222232132322312-1022020002330202-2223032032030221-0100130332013213"></a>

<a id="canonical-2020312211222231-2311100210112013-0000003131031002-0321201311112112-3233133130332122-1002233020012002-2311113203001003-0103212321032013"></a>

## max_url_size_exceeds property — request_constraints / 011200002102 / 16

Type: `"number"`. Optional.

Match on the URL Size that exceed this value. Exclusive with \[max\_url\_size\_none\]

Upstream description:

Exclusive with \[max\_url\_size\_none\]

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [max_url_size_none](resources--service_policy--reference--group-003.md#canonical-0001112012202010-3112323211321020-1010031231321302-2032201031020002-0032313122112232-3132020222101132-0103010031123132-3110203212320130): complete subsection reference.

<a id="canonical-3030003200311330-2311311223011220-1310033200322030-0130132202300001-2011110223201321-3110203023132020-1323122102120332-2303221221130232"></a>

## Next pages — request_constraints / 011200002102 / 17

- [rule_list.rules.spec.request_constraints.max_cookie_count_none](resources--service_policy--reference--group-003.md#canonical-0030222301212332-0012300010222202-0111102300021200-2223301132322123-1133320011022321-0333102223220110-1332121101012231-0323022211010233)
- [rule_list.rules.spec.request_constraints.max_cookie_key_size_none](resources--service_policy--reference--group-003.md#canonical-1113321011231222-2012220023330033-0012333012221100-3322133231023203-0201000310120310-2301110200320120-3022232232130212-3033303123222012)
- [rule_list.rules.spec.request_constraints.max_cookie_value_size_none](resources--service_policy--reference--group-003.md#canonical-2003211303022323-3332223010332321-2301222030020233-1021330312213310-0011101303003300-0102320320312332-1102220311121300-0113023322332102)
- [rule_list.rules.spec.request_constraints.max_header_count_none](resources--service_policy--reference--group-003.md#canonical-1012312000301123-0220212010011032-0323103313020203-1113220000101222-2121211311220110-0302232302310222-2000233321101131-0130211130000100)
- [rule_list.rules.spec.request_constraints.max_header_key_size_none](resources--service_policy--reference--group-003.md#canonical-3002331231230012-1002010202101030-1222010232332013-2202012002201021-1112100102110312-2032002121103332-2211110131031022-0302321222202111)
- [rule_list.rules.spec.request_constraints.max_header_value_size_none](resources--service_policy--reference--group-003.md#canonical-0211210213001101-2133201323222011-2320223111203302-3021002322100332-1211023310210001-2133120232030322-0003301221033113-3133303200132203)
- [rule_list.rules.spec.request_constraints.max_parameter_count_none](resources--service_policy--reference--group-003.md#canonical-3101302230010022-0013031312110133-0321121220133022-2311312302112322-3331023130332301-1133320301212222-3221320321311301-2031031023133113)
- [rule_list.rules.spec.request_constraints.max_parameter_name_size_none](resources--service_policy--reference--group-003.md#canonical-0002301100011231-2122311113330121-0311221202132100-1302333221302203-2230111002112032-1020033103213333-3132010001232303-0302120113120032)
- [rule_list.rules.spec.request_constraints.max_parameter_value_size_none](resources--service_policy--reference--group-003.md#canonical-3200323212233321-0033011302233121-2203302022132212-1120122210313321-2333213120030122-2020000123022320-1220121001203310-3202300320101201)
- [rule_list.rules.spec.request_constraints.max_query_size_none](resources--service_policy--reference--group-003.md#canonical-1013020222111032-2123302100312203-3232302320201033-1112132112021122-3311120131020020-1002132011310023-2100113111330221-0312301301200102)
- [rule_list.rules.spec.request_constraints.max_request_line_size_none](resources--service_policy--reference--group-003.md#canonical-1312203332130002-0010103331023220-3230131321200133-3302332030033232-2130320000212332-2310302032332113-1122211012123030-3031233122111230)
- [rule_list.rules.spec.request_constraints.max_request_size_none](resources--service_policy--reference--group-003.md#canonical-2211003200031203-2030000102310111-3010311010013203-1310200022020121-2331102101212302-0231323030121333-1013000131130122-1120012210021030)
- [rule_list.rules.spec.request_constraints.max_url_size_none](resources--service_policy--reference--group-003.md#canonical-0001112012202010-3112323211321020-1010031231321302-2032201031020002-0032313122112232-3132020222101132-0103010031123132-3110203212320130)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-0030222301212332-0012300010222202-0111102300021200-2223301132322123-1133320011022321-0333102223220110-1332121101012231-0323022211010233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132222332102131-1003330221012320-1220020120003002-1321330220103200-3220212020212201-3123310310311132-3320220202001032-1210210113011313"></a>

## rule_list.rules.spec.request_constraints.max_cookie_count_none — max_cookie_count_none / 202100332121 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_cookie_count_none

<a id="canonical-2033331210013202-1301200033023110-3033023210233200-0302110203221311-1220113021303013-0220112200201323-1202133232320022-3201301133031123"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie count none.

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
max_cookie_count_none = {}
```

<a id="canonical-0220333333302320-0120213300213020-2211021113012232-1131210231013002-3200032003120130-0213002310212113-3232210233102331-1200030020203230"></a>

## Direct properties — max_cookie_count_none / 202100332121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100230111210333-3012310332123230-0030200031023321-3331323022022120-1112322103101033-0132011333200031-3000210231300003-2202113013112131"></a>

## Next pages — max_cookie_count_none / 202100332121 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-1113321011231222-2012220023330033-0012333012221100-3322133231023203-0201000310120310-2301110200320120-3022232232130212-3033303123222012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302321303212211-0012323311302021-3320302323131002-3111210330203000-3011002112230310-1320302301222301-0233332223001200-1231203021012113"></a>

## rule_list.rules.spec.request_constraints.max_cookie_key_size_none — max_cookie_key_size_none / 012102221000 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_cookie_key_size_none

<a id="canonical-1322302200033123-1131111212300303-2032233013003123-0321101002232202-2111302012322112-3322311310310010-0322012121223303-3211120220000103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie key size none.

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
max_cookie_key_size_none = {}
```

<a id="canonical-1113122001212101-3332300313313202-2000322030313133-1320300333330221-3221003322101211-0333330102001033-2022031030223223-3223233332210210"></a>

## Direct properties — max_cookie_key_size_none / 012102221000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0121303112012130-3232000201201321-1310201322220100-1220200003232321-3033202002132002-0330121331303220-1023000120202310-2133322233121122"></a>

## Next pages — max_cookie_key_size_none / 012102221000 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-2003211303022323-3332223010332321-2301222030020233-1021330312213310-0011101303003300-0102320320312332-1102220311121300-0113023322332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112231203221000-1203122112213221-0011321233302020-2330023313322000-2321212132231333-1331010122013011-3302033013220110-1010321010020102"></a>

## rule_list.rules.spec.request_constraints.max_cookie_value_size_none — max_cookie_value_size_none / 201011010110 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_cookie_value_size_none

<a id="canonical-3100112101300020-0201221301112022-2302330012220132-2211112023333231-2133310233130032-0031201202130210-2012202300310311-0210003230001023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie value size none.

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
max_cookie_value_size_none = {}
```

<a id="canonical-0012201113002310-1311022322202322-2113302000321121-0030331233010313-3311100000132301-0332333303112123-1003011210310123-3133313130001211"></a>

## Direct properties — max_cookie_value_size_none / 201011010110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320200332013210-1201100331230012-3122112103322022-2223011333100132-1032231102322222-2003102321300020-1111331031331001-1130113003223031"></a>

## Next pages — max_cookie_value_size_none / 201011010110 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-1012312000301123-0220212010011032-0323103313020203-1113220000101222-2121211311220110-0302232302310222-2000233321101131-0130211130000100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132312110122022-2112110322021021-3321223022012012-3223021131201301-3113230233233100-2313131212131122-3312223013232122-0001221112333320"></a>

## rule_list.rules.spec.request_constraints.max_header_count_none — max_header_count_none / 030320321130 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_header_count_none

<a id="canonical-3231131322130322-2322120113213002-3323331321003102-3233000130312032-3302111222211022-0000322033033110-3133033133030003-0330320331322213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header count none.

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
max_header_count_none = {}
```

<a id="canonical-1312310001333133-0103231100023123-1321300121321102-0113321103003012-3023100113033033-3232332121120323-0312313223120332-0301202330021223"></a>

## Direct properties — max_header_count_none / 030320321130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111112012120022-3020321130223013-1210032312310331-0111132003003012-2002200130002201-0111000221233000-0311032012131002-2120300003031233"></a>

## Next pages — max_header_count_none / 030320321130 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-3002331231230012-1002010202101030-1222010232332013-2202012002201021-1112100102110312-2032002121103332-2211110131031022-0302321222202111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303013111133313-0103011101023231-0201201200222111-2233133323132031-0003101332102321-0200223320100001-2113030232322110-0220103110320133"></a>

## rule_list.rules.spec.request_constraints.max_header_key_size_none — max_header_key_size_none / 132200101013 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_header_key_size_none

<a id="canonical-2311321222221211-2202013013113133-1210132321200222-2333313033032131-1212032303301333-1203320020310031-0003112033312031-1002332221110221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header key size none.

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
max_header_key_size_none = {}
```

<a id="canonical-0203331033330220-3200333233321002-3323120100101132-0310323123331131-2113112310123103-2033200033122313-2100122300203332-3103002300310102"></a>

## Direct properties — max_header_key_size_none / 132200101013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033003303311032-3022212033230121-0022130122120122-0110103312230200-1220110201101310-1012323131033331-1301012122333302-3221021122231110"></a>

## Next pages — max_header_key_size_none / 132200101013 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-0211210213001101-2133201323222011-2320223111203302-3021002322100332-1211023310210001-2133120232030322-0003301221033113-3133303200132203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212113101121020-1020231122230002-3231301203233000-1232213223120222-1210003130310102-0211112220202032-2013030023332133-2003123332131113"></a>

## rule_list.rules.spec.request_constraints.max_header_value_size_none — max_header_value_size_none / 301003130000 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_header_value_size_none

<a id="canonical-0200202312023131-3221023120000003-1022133301231332-1032131001023223-0300111321312032-3102330222211120-0111230313322321-0320000122002211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header value size none.

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
max_header_value_size_none = {}
```

<a id="canonical-1312021120303303-3323212031000130-3123010102321301-3003020103200233-0331233022033030-1322012203101132-2213033022020300-0322110102313122"></a>

## Direct properties — max_header_value_size_none / 301003130000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232321030333102-1322301122101023-1000131201031120-2010333102333213-3022320321122312-0130123022231300-0332030332311233-2012323200001211"></a>

## Next pages — max_header_value_size_none / 301003130000 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-3101302230010022-0013031312110133-0321121220133022-2311312302112322-3331023130332301-1133320301212222-3221320321311301-2031031023133113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222333020131222-2110223102100230-2022012220301102-0032100223202302-2130003232130020-0223330130010030-0010210331111130-2132102101301100"></a>

## rule_list.rules.spec.request_constraints.max_parameter_count_none — max_parameter_count_none / 203311012300 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_parameter_count_none

<a id="canonical-1011003033220003-2231200213222112-1301323310323030-2201010130331310-3202122333003000-2000311032021302-2213203011033310-2332210032333022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max parameter count none.

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
max_parameter_count_none = {}
```

<a id="canonical-1023020202331121-2322302311232220-3120210001323320-2000110333111030-1211220323202232-2201300330201032-2200112332222302-0202323330103122"></a>

## Direct properties — max_parameter_count_none / 203311012300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202212021311322-0323210000132121-0230333111101322-2310022222122122-3332201311012001-0213021103001330-2113021101101310-0101102302132030"></a>

## Next pages — max_parameter_count_none / 203311012300 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-0002301100011231-2122311113330121-0311221202132100-1302333221302203-2230111002112032-1020033103213333-3132010001232303-0302120113120032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311121131203332-2122000103123030-2113330230111031-2212221220023002-2202230121112333-0113133033112030-1002121023101132-2200231230000311"></a>

## rule_list.rules.spec.request_constraints.max_parameter_name_size_none — max_parameter_name_size_none / 322030003332 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_parameter_name_size_none

<a id="canonical-1002200331203103-1221101103003231-3231221331020301-1310323010120031-3331230220333111-0203001303130220-3310300221310330-0100310333320332"></a>

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
max_parameter_name_size_none = {}
```

<a id="canonical-2032332321213203-1003221330030111-0133310201001300-2030132220113332-2011022330123333-2023003230002330-3033130230003123-3021021320110101"></a>

## Direct properties — max_parameter_name_size_none / 322030003332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121031301312013-1100122212311211-1123131100323101-1322032121122223-2220121313312223-1311333332022302-3200331031100003-2113012121320323"></a>

## Next pages — max_parameter_name_size_none / 322030003332 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-3200323212233321-0033011302233121-2203302022132212-1120122210313321-2333213120030122-2020000123022320-1220121001203310-3202300320101201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101132300000303-0001310311103311-0031013021133023-2200311123112121-2210221222123313-2322212001103202-1223003302120331-2130120123222011"></a>

## rule_list.rules.spec.request_constraints.max_parameter_value_size_none — max_parameter_value_size_none / 311222310130 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_parameter_value_size_none

<a id="canonical-1232133101130103-2132120111110022-1130120331112013-2303013230120131-3111332322200001-3002103222212012-1320033011210231-3221001232333330"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max parameter value size none.

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
max_parameter_value_size_none = {}
```

<a id="canonical-3103132103231231-1221221313212013-0222111202321120-3123121222323333-2301102301221302-2310300223232303-3212012330223230-3233323012200223"></a>

## Direct properties — max_parameter_value_size_none / 311222310130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100333130101320-2130023111300200-0233222313302112-1200303212123012-3311233232200000-2322113210003011-3310130022103103-0333100021130320"></a>

## Next pages — max_parameter_value_size_none / 311222310130 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-1013020222111032-2123302100312203-3232302320201033-1112132112021122-3311120131020020-1002132011310023-2100113111330221-0312301301200102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003300020333323-0101020103100103-3302013223321332-1320311301222213-0031101230302021-0313332201230132-2312031121022121-0023310012311230"></a>

## rule_list.rules.spec.request_constraints.max_query_size_none — max_query_size_none / 302221322233 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_query_size_none

<a id="canonical-2220223202000320-2313110303213033-1011313210202330-0032032201011313-0033131310303033-2321033102112233-1132323233001102-1103103023332223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max query size none.

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
max_query_size_none = {}
```

<a id="canonical-3203001003203101-2200031121003211-1220110312003321-1300320030201100-2132300013332310-1222102330202000-2223001101323230-1321302022200000"></a>

## Direct properties — max_query_size_none / 302221322233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130111333130023-3202032211331101-2111122232201213-1320131033311221-3032003031322111-1221002322032121-3322120020231301-0012103013320101"></a>

## Next pages — max_query_size_none / 302221322233 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-1312203332130002-0010103331023220-3230131321200133-3302332030033232-2130320000212332-2310302032332113-1122211012123030-3031233122111230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022133032313223-3123110032230231-3211212330203300-3202012333220000-3332222132000333-2323012112312122-2101123312003301-2121031103200233"></a>

## rule_list.rules.spec.request_constraints.max_request_line_size_none — max_request_line_size_none / 210130002300 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_request_line_size_none

<a id="canonical-1321210222121100-2232102302000320-1201232200012221-1022201300033120-0212223323111031-1012300122330001-3122002210110322-2300103212232203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max request line size none.

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
max_request_line_size_none = {}
```

<a id="canonical-3012312133101201-0012100021230203-2310321210311212-0100301211210133-3203220210021203-3111112022223320-2012210313302221-1132223100132202"></a>

## Direct properties — max_request_line_size_none / 210130002300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323100320203100-2231200023111201-3030011103131001-3002211211223223-2301003333111221-2023221023130210-2211211203112233-3000212130311113"></a>

## Next pages — max_request_line_size_none / 210130002300 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-2211003200031203-2030000102310111-3010311010013203-1310200022020121-2331102101212302-0231323030121333-1013000131130122-1120012210021030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020223312110020-0110322301011201-3323101213022111-3211313331133110-3132033201211032-1322220302102023-3312120111330211-1331003201021322"></a>

## rule_list.rules.spec.request_constraints.max_request_size_none — max_request_size_none / 011001301010 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_request_size_none

<a id="canonical-1203310032203211-3200210000223202-1202030320200112-1133330220010203-1310011202111322-3320130300131030-1020200313321202-1223211320331332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max request size none.

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
max_request_size_none = {}
```

<a id="canonical-0013300211000000-3301002303312012-1032002033313120-0113232300030210-3011111113230031-2102202012100303-1002031032311001-1123230322001121"></a>

## Direct properties — max_request_size_none / 011001301010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133023131021003-2231131331332332-2223031302213000-2330102103021210-0223313332102013-2230300222203020-2210212323131332-3223100132223123"></a>

## Next pages — max_request_size_none / 011001301010 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-0001112012202010-3112323211321020-1010031231321302-2032201031020002-0032313122112232-3132020222101132-0103010031123132-3110203212320130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212332330000033-2133310121101123-3200022232101233-1013223333223022-3332031222232122-2301101320202002-3303031012110003-1231203222303102"></a>

## rule_list.rules.spec.request_constraints.max_url_size_none — max_url_size_none / 021012102022 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- rule_list.rules.spec.request_constraints.max_url_size_none

<a id="canonical-3331101010322200-1032332023221003-1013000330313113-2302113330213310-2333323333113030-0222003132132002-1122031332000212-3101320130031302"></a>

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
max_url_size_none = {}
```

<a id="canonical-0020311202201313-0011212200021000-3110222203102020-0313003213210020-2202101233323323-1122103011202332-0302233001222332-3320312221032033"></a>

## Direct properties — max_url_size_none / 021012102022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322232000202312-3031022102200100-3113130301330313-0232011323111013-2320020131230310-1121100102322021-0020102130302011-0210221231232022"></a>

## Next pages — max_url_size_none / 021012102022 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-003.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031130201113220-3120312123313120-0003312323202121-0202111120221221-0301231023102202-2110021132103313-1002003001101332-2001033020332130"></a>

## rule_list.rules.spec.segment_policy — segment_policy / 123023321033 / 2

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

<a id="canonical-3310000020213002-2003001331123320-0303202231302031-1230311020231120-2232222110203212-3012033200211120-3312221100220311-2300010211020301"></a>

## Direct properties — segment_policy / 123023321033 / 3

- [dst_any](resources--service_policy--reference--group-003.md#canonical-3002010211101110-3221323322013110-0000111130011312-2121102330232021-1223010230013233-0010033023203121-3203130003222121-1320110231313030): complete subsection reference.

- [dst_segments](resources--service_policy--reference--group-003.md#canonical-2023102300100100-0010302223303330-1300202232022033-0103023130302032-1023012301210211-3320130333031013-2033320130101133-1131333303332132): complete subsection reference.

- [intra_segment](resources--service_policy--reference--group-003.md#canonical-3201323020022131-3013300300300130-0212302313301301-2300301302210303-2021000303132231-0132322023231112-0122112233202333-1000013300223122): complete subsection reference.

- [src_any](resources--service_policy--reference--group-003.md#canonical-0013302203311202-1010220003121332-3210230121232312-3002023203002320-1200113032330123-3222112032312123-2032333002210113-3312101313022300): complete subsection reference.

- [src_segments](resources--service_policy--reference--group-003.md#canonical-1321111033030303-3021331113201220-2102331103121003-1222300031003013-3100131131110121-1301223332213100-1032130232201220-1032222211221031): complete subsection reference.

<a id="canonical-0212312021321203-3321300111230121-2212103001110212-3332002321111312-1322100213313322-2302010333312321-2330110222133331-0233121032201213"></a>

## Next pages — segment_policy / 123023321033 / 4

- [rule_list.rules.spec.segment_policy.dst_any](resources--service_policy--reference--group-003.md#canonical-3002010211101110-3221323322013110-0000111130011312-2121102330232021-1223010230013233-0010033023203121-3203130003222121-1320110231313030)
- [rule_list.rules.spec.segment_policy.dst_segments](resources--service_policy--reference--group-003.md#canonical-2023102300100100-0010302223303330-1300202232022033-0103023130302032-1023012301210211-3320130333031013-2033320130101133-1131333303332132)
- [rule_list.rules.spec.segment_policy.intra_segment](resources--service_policy--reference--group-003.md#canonical-3201323020022131-3013300300300130-0212302313301301-2300301302210303-2021000303132231-0132322023231112-0122112233202333-1000013300223122)
- [rule_list.rules.spec.segment_policy.src_any](resources--service_policy--reference--group-003.md#canonical-0013302203311202-1010220003121332-3210230121232312-3002023203002320-1200113032330123-3222112032312123-2032333002210113-3312101313022300)
- [rule_list.rules.spec.segment_policy.src_segments](resources--service_policy--reference--group-003.md#canonical-1321111033030303-3021331113201220-2102331103121003-1222300031003013-3100131131110121-1301223332213100-1032130232201220-1032222211221031)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-3002010211101110-3221323322013110-0000111130011312-2121102330232021-1223010230013233-0010033023203121-3203130003222121-1320110231313030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301301311233000-0220231232230310-3123111123023002-0120130103320021-3333311323032220-0002010233321300-3030202200333313-0003332031103223"></a>

## rule_list.rules.spec.segment_policy.dst_any — dst_any / 031311030113 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- rule_list.rules.spec.segment_policy.dst_any

<a id="canonical-1012233013201321-0113011332231330-0302110123312323-3023220002210313-2113120201231301-1120132132033200-2332013221313031-1323113322120203"></a>

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
dst_any = {}
```

<a id="canonical-0123032210122303-0210323223000323-1322233202220201-0221331012101010-3120103011201313-2023023222031113-0222303223030200-1120313311302012"></a>

## Direct properties — dst_any / 031311030113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122231211032120-2131001212212022-3033011313031032-3013021131111310-1032330102011110-1100110030111223-1133000130232013-3011113032301101"></a>

## Next pages — dst_any / 031311030113 / 4

- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-2023102300100100-0010302223303330-1300202232022033-0103023130302032-1023012301210211-3320130333031013-2033320130101133-1131333303332132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330132333330010-0110122311313220-1022333303122132-0020133313202310-3323212112103201-2020230310201033-3323003320201122-2303310011200113"></a>

## rule_list.rules.spec.segment_policy.dst_segments — dst_segments / 033101213333 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- rule_list.rules.spec.segment_policy.dst_segments

<a id="canonical-0132200022132130-3300330010322312-1032121303001121-2130111312331302-3102000010312323-0333100201100331-0310212123312312-3110131320102302"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dst segments.

Upstream description:

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

<a id="canonical-0131221320203222-1102012030122101-0001123013013312-2303331103001330-2020121030231323-3311301202333231-2301211121311030-1212312133311211"></a>

## Direct properties — dst_segments / 033101213333 / 3

- [segments](resources--service_policy--reference--group-003.md#canonical-0131100222231033-1122330332010133-3131222232012322-3313011122311303-1221010121213112-1311310110321100-0331212101032032-3220003021113223): complete subsection reference.

<a id="canonical-0210132213210102-1200021302201021-0312323300300320-2332110330213112-1003113102130313-0310321100203112-3121003213133233-1231220000230002"></a>

## Next pages — dst_segments / 033101213333 / 4

- [rule_list.rules.spec.segment_policy.dst_segments.segments](resources--service_policy--reference--group-003.md#canonical-0131100222231033-1122330332010133-3131222232012322-3313011122311303-1221010121213112-1311310110321100-0331212101032032-3220003021113223)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-0131100222231033-1122330332010133-3131222232012322-3313011122311303-1221010121213112-1311310110321100-0331212101032032-3220003021113223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200222232002200-3312122010223311-0010200130310023-3113213111131301-0101130102312133-2312113101210031-2112331122003232-2013310122333112"></a>

## rule_list.rules.spec.segment_policy.dst_segments.segments — segments / 103023032310 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- [rule_list.rules.spec.segment_policy.dst_segments](resources--service_policy--reference--group-003.md#canonical-2023102300100100-0010302223303330-1300202232022033-0103023130302032-1023012301210211-3320130333031013-2033320130101133-1131333303332132)
- rule_list.rules.spec.segment_policy.dst_segments.segments

<a id="canonical-2332110131120102-2220301302310211-0031031223031213-2301301133122223-3320123101012320-2323001302332220-1030120300230221-2321030013100123"></a>

Type: `"object"`. list nested block, Optional.

Segments. Select list of segments.

Upstream description:

Select list of segments.

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

<a id="canonical-3032013133122012-3213211231131320-3212112221212220-1122020320210311-3333001110110121-2112110323212130-0211323233301120-2223100202232031"></a>

## Direct properties — segments / 103023032310 / 3

<a id="canonical-0020203332031121-1111112312110101-0031013003001013-1311213202121032-0333310200300331-1121113103202301-1222220021311031-3303233113312001"></a>

<a id="canonical-0300233130103300-3223131200000103-1302210203030132-2321200323302333-0112222103011303-0030131230031310-0102031033130332-3111113233230013"></a>

## name property — segments / 103023032310 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3101113001230221-1113030331002101-3121302131311002-3333120331002110-2221313211220102-3211223100320101-1002030210200320-3030201200032202"></a>

## namespace property — segments / 103023032310 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0312333230103110-2120023000110000-0010130123212003-0201132302130123-3002020231201003-0233033212220003-2003222220123132-3323012211033323"></a>

## tenant property — segments / 103023032310 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2203132203110212-0100101032031110-2330312100201213-3113100000023103-2011000222100230-2203230212301311-1323302202330020-1332220300222323"></a>

## Next pages — segments / 103023032310 / 7

- [rule_list.rules.spec.segment_policy.dst_segments](resources--service_policy--reference--group-003.md#canonical-2023102300100100-0010302223303330-1300202232022033-0103023130302032-1023012301210211-3320130333031013-2033320130101133-1131333303332132)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-3201323020022131-3013300300300130-0212302313301301-2300301302210303-2021000303132231-0132322023231112-0122112233202333-1000013300223122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220112013010220-0103233111323330-1230022030131013-1121022222111222-3032121201231211-0013300122121203-1113131332003312-1122102320223220"></a>

## rule_list.rules.spec.segment_policy.intra_segment — intra_segment / 202200112321 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- rule_list.rules.spec.segment_policy.intra_segment

<a id="canonical-1210033002102202-1012321302012001-1221021120312231-0210110320120122-1301110110203120-1220333211010130-1313010111312021-0223323120003321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for intra segment.

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
intra_segment = {}
```

<a id="canonical-3333003222311203-2030020230020313-0101311010120202-2211112113231011-1302130231232012-2333301303022133-1112213302300302-3301203221332230"></a>

## Direct properties — intra_segment / 202200112321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002320323202201-1223232032231301-1030313123232001-2112013010223200-2021212330310302-2200330113130232-2021022212321111-3132010110220122"></a>

## Next pages — intra_segment / 202200112321 / 4

- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-0013302203311202-1010220003121332-3210230121232312-3002023203002320-1200113032330123-3222112032312123-2032333002210113-3312101313022300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223032011333320-1012203022210130-1211020310123210-3000030023102331-1120210002113021-0232200111330302-0002320000010302-1300113111221300"></a>

## rule_list.rules.spec.segment_policy.src_any — src_any / 013230103230 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- rule_list.rules.spec.segment_policy.src_any

<a id="canonical-1021110201210013-2301132113303322-1212121222233110-1330302030322332-0131130002000303-1022331220210313-3023002332330031-2330102232022201"></a>

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
src_any = {}
```

<a id="canonical-2303232333023100-1110223002203323-2113112330213121-3001203233311010-0322300213001222-2313330301012113-1101222312202233-0031301221022122"></a>

## Direct properties — src_any / 013230103230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202030220300313-0122320233133011-1302103313212002-2211003332313023-1301010101300233-3201010122203231-0303100302231233-0001313013121330"></a>

## Next pages — src_any / 013230103230 / 4

- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-1321111033030303-3021331113201220-2102331103121003-1222300031003013-3100131131110121-1301223332213100-1032130232201220-1032222211221031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233113012030020-0001020133010232-0100311110300132-2223003310232103-2212101101120030-1122132301010213-1001313122120220-1303013323131301"></a>

## rule_list.rules.spec.segment_policy.src_segments — src_segments / 333310001332 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- rule_list.rules.spec.segment_policy.src_segments

<a id="canonical-3332231220130101-1011320302021122-1303123303132300-0311302332120211-2132013011112123-2210133031301012-2210112211011003-0032013310233112"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for src segments.

Upstream description:

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

<a id="canonical-0130331313301323-1323202110322212-1101332210110212-2100120311112333-0133223100313210-3013021331230221-0210221103221112-1233032333330231"></a>

## Direct properties — src_segments / 333310001332 / 3

- [segments](resources--service_policy--reference--group-003.md#canonical-2122210222323033-3030011311312001-3323232201220221-1201303302003330-0001201301100210-3122330110301321-2012032221200013-1303001132201130): complete subsection reference.

<a id="canonical-2121200133100321-3222211323332030-3332232333023331-3121122233303231-3311320232102200-3203010231032130-1211332313100320-0031231302020110"></a>

## Next pages — src_segments / 333310001332 / 4

- [rule_list.rules.spec.segment_policy.src_segments.segments](resources--service_policy--reference--group-003.md#canonical-2122210222323033-3030011311312001-3323232201220221-1201303302003330-0001201301100210-3122330110301321-2012032221200013-1303001132201130)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-2122210222323033-3030011311312001-3323232201220221-1201303302003330-0001201301100210-3122330110301321-2012032221200013-1303001132201130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222312301011220-3022010103010313-0321211123010330-3312023031313123-0201032223222310-2112213332012121-3233133001332112-3201000110300133"></a>

## rule_list.rules.spec.segment_policy.src_segments.segments — segments / 130011033202 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223)
- [rule_list.rules.spec.segment_policy.src_segments](resources--service_policy--reference--group-003.md#canonical-1321111033030303-3021331113201220-2102331103121003-1222300031003013-3100131131110121-1301223332213100-1032130232201220-1032222211221031)
- rule_list.rules.spec.segment_policy.src_segments.segments

<a id="canonical-0102023111220023-0000302012323012-1121022023130210-1303022321020333-2130213231100210-1301213210011131-2201020001200002-3003302202213232"></a>

Type: `"object"`. list nested block, Optional.

Segments. Select list of segments.

Upstream description:

Select list of segments.

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

<a id="canonical-1102321110131121-2321203101322100-3031312330133022-0333301301212320-3121032332332030-0113012333122123-1132320031312011-2121030103033003"></a>

## Direct properties — segments / 130011033202 / 3

<a id="canonical-0332131131300213-0301113323313102-2320303301300302-3311002230202122-0120031101330203-2121121022333231-2001120002132013-2220122002122233"></a>

<a id="canonical-0321121111113130-2332203020031102-1133310033333233-0301002131331033-0012320201233231-1322330131232002-3031132231032222-0312302111220201"></a>

## name property — segments / 130011033202 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3320032303100311-2012312223333311-1233033022213113-3030120113013020-0102323333301002-1310000002213330-1221320322031011-2020021312101311"></a>

## namespace property — segments / 130011033202 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1131222311000210-1210233112000033-3202100031211003-3203312011020333-3133303000230011-3202313021132020-0032312022310313-3222313012132113"></a>

## tenant property — segments / 130011033202 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3220031033323320-3021100121333220-0312211110123302-1200111332011331-1023123230103202-2230330020232332-0023031213110122-3220202230201333"></a>

## Next pages — segments / 130011033202 / 7

- [rule_list.rules.spec.segment_policy.src_segments](resources--service_policy--reference--group-003.md#canonical-1321111033030303-3021331113201220-2102331103121003-1222300031003013-3100131131110121-1301223332213100-1032130232201220-1032222211221031)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-1020320201300000-0002323231201110-2333023122032003-3301321123001032-2102310103222032-0133000133300130-0123200131201022-1013312312033202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310220233100323-2031002020323100-0301112030022121-2103133221201021-2302030231312033-1321321021110020-3132010313320320-2012312131222211"></a>

## rule_list.rules.spec.tls_fingerprint_matcher — tls_fingerprint_matcher / 200132121233 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-2211302223202211-0133312332123122-3001120111203221-2001331203120000-2122122033200113-0203222022113013-3310203232101331-2101130030200002"></a>

Type: `"object"`. single nested block, Optional.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

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

<a id="canonical-0120133112100232-1312213103230003-1110111213311331-1110032202330232-3131202301302003-0323021220032111-1323210312002203-1201311023010201"></a>

## Direct properties — tls_fingerprint_matcher / 200132121233 / 3

<a id="canonical-0331220121023231-0021112103101300-0031313222021121-3210131000332312-2331212321201033-0321332301010021-1201203310300313-0303030100323313"></a>

<a id="canonical-0333300311213123-0231113033220203-0100232311230030-0233122313012033-0212332003122123-1333000333131311-1212122321001100-0203132001000110"></a>

## classes property — tls_fingerprint_matcher / 200132121233 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2210331331110110-2330131021113122-1123111311303011-3112133101012112-2022213100312300-2201313133020312-2033301333003200-3021102023211211"></a>

## exact_values property — tls_fingerprint_matcher / 200132121233 / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2311333020000321-0220011231032112-2220020131100230-0031031333322132-1111122212101032-0023223330320320-2001213101210111-1103330330302001"></a>

## excluded_values property — tls_fingerprint_matcher / 200132121233 / 6

Type: `["list", "string"]`. Optional.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2102003330032222-0013320031222320-3211313132120013-2311013130012220-1322323130301130-2001221110232310-0202321211311330-2323332021330121"></a>

## Next pages — tls_fingerprint_matcher / 200132121233 / 7

- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-3001003323033212-0123331232232220-3131203230202012-2011331002320011-1212313031310311-0101011121113132-2102222322031110-3333021300010302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100233020113002-2320000002312123-0211303021030312-2302321010023012-0330222211113010-3213113030321320-3312211013301213-2302213023001023"></a>

## rule_list.rules.spec.user_identity_matcher — user_identity_matcher / 001201101121 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.user_identity_matcher

<a id="canonical-1232010221230100-0121021000010310-2213323203322133-0302300133131212-2033103123311120-0012331103220112-1112122302102113-0011130020320302"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-1130312330033203-3030310222321233-1131122231022020-2231013020301123-3032202312213120-0022023003121000-0222202321320012-3100021320330303"></a>

## Direct properties — user_identity_matcher / 001201101121 / 3

<a id="canonical-1321310003332203-0023133101222121-3312023221111313-3123033321102023-3112220231302223-1031201210222330-3003333102230021-0011001103133301"></a>

<a id="canonical-3300231010232000-2222031013210213-1223303321332000-0120120310031303-2000232112020022-2231110210332200-0010011200132210-3001232201211201"></a>

## exact_values property — user_identity_matcher / 001201101121 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1102310000011003-0102000000120112-0103303102111003-1211123333231003-3213221013320130-0202021303111113-0000310222302023-2231100322111020"></a>

## regex_values property — user_identity_matcher / 001201101121 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1230010303021323-2300320002303120-2013200202010211-1130230011232313-2310332222022222-3303300101022030-0020021320000120-1322011103223012"></a>

## Next pages — user_identity_matcher / 001201101121 / 6

- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-1330120330031000-3202113320223100-2303020333222232-1113322320113320-2131320031302300-3021303213200102-0321303213312133-3121220001012010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321123201100001-0002332311031111-1113212310021110-1021313200021130-3112332231211231-2030320012302110-3222100111320331-3031301212133313"></a>

## rule_list.rules.spec.waf_action — waf_action / 222131210323 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.waf_action

<a id="canonical-2323131031010332-3002023002320130-2322222211000312-3130213100113011-2320002112103301-2101033310021200-1130113033130203-0233221333232023"></a>

Type: `"object"`. single nested block, Optional.

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Upstream description:

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall_detection_control",
    "none"),
  validators.ConflictingObjectAttributes("app_firewall_detection_control",
    "waf_skip_processing"),
  validators.ConflictingObjectAttributes("none",
    "waf_skip_processing")}
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
  "x-ves-oneof-field-action_type": "[\"app_firewall_detection_control\",\"none\",\"waf_skip_processing\"]"
}
```

Terraform syntax:

```terraform
waf_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211301231011033-2202222300133310-1322312222203223-2323301123231010-3101211033000033-0331011000313230-0000003020130001-1323000331013110"></a>

## Direct properties — waf_action / 222131210323 / 3

- [app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-3022321321013201-0020033302122032-2311011301300323-1202110132110022-1233301221333112-0030120230100322-3103323111103001-1210013301233011): complete subsection reference.

- [none](resources--service_policy--reference--group-003.md#canonical-0203032203021300-2101223011122200-2011010020200031-1121003201202320-3121133102312123-0230221030232323-0331000320232021-3330311131131220): complete subsection reference.

- [waf_skip_processing](resources--service_policy--reference--group-003.md#canonical-0332300220323213-1323323030222322-1131312000330222-2232032001110301-0212213131220202-0321130113302233-2100102223233213-2032013001120230): complete subsection reference.

<a id="canonical-1210130321300132-2100110231301311-0231332023333311-2120203101010221-0101012020311303-3212131121001101-0210000202010001-0210201303220033"></a>

## Next pages — waf_action / 222131210323 / 4

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-3022321321013201-0020033302122032-2311011301300323-1202110132110022-1233301221333112-0030120230100322-3103323111103001-1210013301233011)
- [rule_list.rules.spec.waf_action.none](resources--service_policy--reference--group-003.md#canonical-0203032203021300-2101223011122200-2011010020200031-1121003201202320-3121133102312123-0230221030232323-0331000320232021-3330311131131220)
- [rule_list.rules.spec.waf_action.waf_skip_processing](resources--service_policy--reference--group-003.md#canonical-0332300220323213-1323323030222322-1131312000330222-2232032001110301-0212213131220202-0321130113302233-2100102223233213-2032013001120230)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-3022321321013201-0020033302122032-2311011301300323-1202110132110022-1233301221333112-0030120230100322-3103323111103001-1210013301233011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013010010223230-3210311003123021-1200310232231101-3302000333122000-3123022132031332-1312302221013132-1323120301020111-2023112201023122"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control — app_firewall_detection_control / 103111213232 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-1330120330031000-3202113320223100-2303020333222232-1113322320113320-2131320031302300-3021303213200102-0321303213312133-3121220001012010)
- rule_list.rules.spec.waf_action.app_firewall_detection_control

<a id="canonical-3230320101223013-1011031223033032-0101122123300100-3100203130020222-2332330322003133-3033201132303311-2212003122110223-2113023230100000"></a>

Type: `"object"`. single nested block, Optional.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

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
app_firewall_detection_control {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312300132031130-3320013132333031-1210223111231201-3111203212023322-0013322303310220-0311022003032120-3032210000012201-0333220101222322"></a>

## Direct properties — app_firewall_detection_control / 103111213232 / 3

- [exclude_attack_type_contexts](resources--service_policy--reference--group-003.md#canonical-1121200201210132-3112020211031112-3020333222132202-0302323013321113-3012230111333203-0033112022201112-3230012331302102-1212300013110211): complete subsection reference.

- [exclude_bot_name_contexts](resources--service_policy--reference--group-003.md#canonical-0111310013301302-1200320013221100-1221313300203131-0031233311320110-2311000301300122-1023321132320232-1212000220201110-3123101331210330): complete subsection reference.

- [exclude_signature_contexts](resources--service_policy--reference--group-003.md#canonical-2230232030112002-3133213202011322-2231200121302112-0332111221013003-0311220100113122-1300330031002113-3213232001133020-3232112322220211): complete subsection reference.

- [exclude_violation_contexts](resources--service_policy--reference--group-003.md#canonical-3131312213101323-2132132033200322-1112310220210201-2110202010012102-3232313130203301-2122221333313001-1201332103102311-2020230301021103): complete subsection reference.

<a id="canonical-3101031010100032-3102003323221231-1200233011000222-3200332110003311-2113221302013113-3203031323210202-2332320002212333-3032000222131312"></a>

## Next pages — app_firewall_detection_control / 103111213232 / 4

- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts](resources--service_policy--reference--group-003.md#canonical-1121200201210132-3112020211031112-3020333222132202-0302323013321113-3012230111333203-0033112022201112-3230012331302102-1212300013110211)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts](resources--service_policy--reference--group-003.md#canonical-0111310013301302-1200320013221100-1221313300203131-0031233311320110-2311000301300122-1023321132320232-1212000220201110-3123101331210330)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts](resources--service_policy--reference--group-003.md#canonical-2230232030112002-3133213202011322-2231200121302112-0332111221013003-0311220100113122-1300330031002113-3213232001133020-3232112322220211)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts](resources--service_policy--reference--group-003.md#canonical-3131312213101323-2132132033200322-1112310220210201-2110202010012102-3232313130203301-2122221333313001-1201332103102311-2020230301021103)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-1330120330031000-3202113320223100-2303020333222232-1113322320113320-2131320031302300-3021303213200102-0321303213312133-3121220001012010)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-1121200201210132-3112020211031112-3020333222132202-0302323013321113-3012230111333203-0033112022201112-3230012331302102-1212300013110211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223000321030010-2230033023112132-2111213003001222-0321033123231132-3111032002232300-2123002320333031-0113201331202202-0321221121002313"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts — exclude_attack_type_contexts / 121120200220 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-1330120330031000-3202113320223100-2303020333222232-1113322320113320-2131320031302300-3021303213200102-0321303213312133-3121220001012010)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-3022321321013201-0020033302122032-2311011301300323-1202110132110022-1233301221333112-0030120230100322-3103323111103001-1210013301233011)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-3203111012121003-3002323231231031-2230103200131210-1113000133301133-2333002220003102-1320230003323300-0213021033331300-0022112312120100"></a>

Type: `"object"`. list nested block, Optional.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_attack_type_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113330110202033-0332212110232211-0300020002121212-2023003101202131-2223101111031332-0111200210001310-0322212233213102-1113002023002223"></a>

## Direct properties — exclude_attack_type_contexts / 121120200220 / 3

<a id="canonical-2233211310100100-1030322201001210-0323223223213122-1022322332310120-0113000201323230-1320223132001333-0233202331000013-1203220332032300"></a>

<a id="canonical-2313300333323122-0331000121330120-3103113013332313-0002102303232103-0211331110232102-2322213102230013-0113230000210221-2001222133321033"></a>

## context property — exclude_attack_type_contexts / 121120200220 / 4

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CONTEXT_ANY","CONTEXT_BODY","CONTEXT_COOKIE","CONTEXT_HEADER","CONTEXT_PARAMETER","CONTEXT_REQUEST","CONTEXT_RESPONSE","CONTEXT_URI","CONTEXT_URL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1301331011100031-2132301110033000-3313120300231121-1100330301201021-0020030022203020-2023103330122233-2101322211332013-1023233320232200"></a>

<a id="canonical-0322020303111212-3021112301222313-3122112310312121-2033023010323321-0212323322200001-1323332312333301-1323230013313000-3331030323102300"></a>

## context_name property — exclude_attack_type_contexts / 121120200220 / 5

Type: `"string"`. Optional.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0211231312311010-0130013111211202-2012300323311011-1322213322330330-1210111222303312-3112033330133302-2022033203003321-2213310301213103"></a>

<a id="canonical-3020021110002321-2103111021232122-3010013132102013-1130023113313223-0111222102222223-1301212202113021-2121231332101300-0033200133100023"></a>

## exclude_attack_type property — exclude_attack_type_contexts / 121120200220 / 6

Type: `"string"`. Optional.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Upstream description:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY","ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS","ATTACK_TYPE_BUFFER_OVERFLOW","ATTACK_TYPE_COMMAND_EXECUTION","ATTACK_TYPE_CROSS_SITE_SCRIPTING","ATTACK_TYPE_DENIAL_OF_SERVICE","ATTACK_TYPE_DETECTION_EVASION","ATTACK_TYPE_DIRECTORY_INDEXING","ATTACK_TYPE_FORCEFUL_BROWSING","ATTACK_TYPE_GRAPHQL_PARSER_ATTACK","ATTACK_TYPE_HTTP_PARSER_ATTACK","ATTACK_TYPE_HTTP_RESPONSE_SPLITTING","ATTACK_TYPE_INFORMATION_LEAKAGE","ATTACK_TYPE_LDAP_INJECTION","ATTACK_TYPE_MALICIOUS_FILE_UPLOAD","ATTACK_TYPE_NONE","ATTACK_TYPE_NON_BROWSER_CLIENT","ATTACK_TYPE_OTHER_APPLICATION_ATTACKS","ATTACK_TYPE_PATH_TRAVERSAL","ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION","ATTACK_TYPE_REMOTE_FILE_INCLUDE","ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION","ATTACK_TYPE_SESSION_HIJACKING","ATTACK_TYPE_SQL_INJECTION","ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE","ATTACK_TYPE_VULNERABILITY_SCAN","ATTACK_TYPE_XPATH_INJECTION"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3031113223022300-3133211122300021-1220201200231201-3210233313301300-0221030331103130-1322133131303302-3101002122121131-1003012022332313"></a>

## Next pages — exclude_attack_type_contexts / 121120200220 / 7

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-3022321321013201-0020033302122032-2311011301300323-1202110132110022-1233301221333112-0030120230100322-3103323111103001-1210013301233011)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-0111310013301302-1200320013221100-1221313300203131-0031233311320110-2311000301300122-1023321132320232-1212000220201110-3123101331210330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200300301230200-2312123122130003-1001000323330023-0121231310110030-2131330320221200-3311010213321002-0201120212010133-1211130221323212"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts — exclude_bot_name_contexts / 031032322001 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-1330120330031000-3202113320223100-2303020333222232-1113322320113320-2131320031302300-3021303213200102-0321303213312133-3121220001012010)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-3022321321013201-0020033302122032-2311011301300323-1202110132110022-1233301221333112-0030120230100322-3103323111103001-1210013301233011)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-3300111112133121-0310032000322201-3013331110020120-2302211202200011-1320321012202131-1123103101113102-2331200000102300-2102322321020130"></a>

Type: `"object"`. list nested block, Optional.

Bot Names to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("bot_name")}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_bot_name_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030021211213231-0131303323033011-1310223130021020-1201020212201302-2032221101022120-2030121201131021-0233300202022122-2203021211221112"></a>

## Direct properties — exclude_bot_name_contexts / 031032322001 / 3

<a id="canonical-2000131203123232-0320021223212232-3130133213210003-3313113003003033-1120103000001011-0123300020231310-0221230101002223-2312002313111003"></a>

<a id="canonical-0101003320220032-3230230001132131-0230312312302330-1312022321223212-3131133103022031-2011220012303331-2323322020220100-3231133210131330"></a>

## bot_name property — exclude_bot_name_contexts / 031032322001 / 4

Type: `"string"`. Optional.

Bot Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0322103000132131-0122032312200220-3020110322130130-3212320013020122-3000223110322012-1302310013102103-0232130012130130-0121320321302312"></a>

## Next pages — exclude_bot_name_contexts / 031032322001 / 5

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-3022321321013201-0020033302122032-2311011301300323-1202110132110022-1233301221333112-0030120230100322-3103323111103001-1210013301233011)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-2230232030112002-3133213202011322-2231200121302112-0332111221013003-0311220100113122-1300330031002113-3213232001133020-3232112322220211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011022301113032-3212120210032111-0120001030320122-3311130221012213-2020212320201313-2021321331232031-1201113231001223-2010111113303122"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts — exclude_signature_contexts / 210100123130 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-1330120330031000-3202113320223100-2303020333222232-1113322320113320-2131320031302300-3021303213200102-0321303213312133-3121220001012010)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-3022321321013201-0020033302122032-2311011301300323-1202110132110022-1233301221333112-0030120230100322-3103323111103001-1210013301233011)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-1013000012311210-3230012300322003-2031223220130102-0133321333310221-2101203023211023-2002202011021323-2203312011002113-3003112002113233"></a>

Type: `"object"`. list nested block, Optional.

Signature IDs to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("signature_id")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_signature_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301122200220300-1212333120230311-0222310111331322-1123003102220013-3310222100203123-1203213300010002-1212312120223313-0013333330123023"></a>

## Direct properties — exclude_signature_contexts / 210100123130 / 3

<a id="canonical-1000001223231030-1132101302013320-1311210102013211-3120113202210103-2330332231000211-2233230301232011-1303030332013231-0110321130010113"></a>

<a id="canonical-2033113330323303-0323310103231123-2230333111133220-3013323311011311-2201310312031020-1110322103311300-2023201232110301-0102211133130021"></a>

## context property — exclude_signature_contexts / 210100123130 / 4

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CONTEXT_ANY","CONTEXT_BODY","CONTEXT_COOKIE","CONTEXT_HEADER","CONTEXT_PARAMETER","CONTEXT_REQUEST","CONTEXT_RESPONSE","CONTEXT_URI","CONTEXT_URL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3203212010030220-2230123130120200-3110101022210220-3211100030330013-0130133121031221-1121312013031212-2232333013103302-3320023102223303"></a>

<a id="canonical-3112223012330301-2231111203333330-2120032212330001-3310003001013310-2201200201202121-2121232113221220-0120132310120111-1003322001133123"></a>

## context_name property — exclude_signature_contexts / 210100123130 / 5

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1330222030012132-1103030312020202-2033310031321100-0301013131001133-3302320311322332-1130311021003322-2111100323133302-2321322111302021"></a>

<a id="canonical-3112230021131232-0220102032222013-1303022130320320-2313232110211113-2322022333310123-2102310222200220-0221130213322130-1200330001023212"></a>

## signature_id property — exclude_signature_contexts / 210100123130 / 6

Type: `"number"`. Optional.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 299999999),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

<a id="canonical-1002132330301222-3121123120023112-0102233103232303-1312330313103213-1232000103212222-2100321002131302-1312112211200223-1303230323001011"></a>

## Next pages — exclude_signature_contexts / 210100123130 / 7

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-3022321321013201-0020033302122032-2311011301300323-1202110132110022-1233301221333112-0030120230100322-3103323111103001-1210013301233011)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-3131312213101323-2132132033200322-1112310220210201-2110202010012102-3232313130203301-2122221333313001-1201332103102311-2020230301021103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302000000230221-3110213231223002-0201321232031132-1001233001001012-0013132233211232-1002311330210120-3323231123023011-3202002333303130"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts — exclude_violation_contexts / 330003302022 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-1330120330031000-3202113320223100-2303020333222232-1113322320113320-2131320031302300-3021303213200102-0321303213312133-3121220001012010)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-3022321321013201-0020033302122032-2311011301300323-1202110132110022-1233301221333112-0030120230100322-3103323111103001-1210013301233011)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-1100221322321231-2212130303000130-2203333230031211-3212022103122302-1320212320000101-3120322232302131-3220111232313130-2213211331131032"></a>

Type: `"object"`. list nested block, Optional.

Violations to be excluded for the defined match criteria.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_violation_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133001122031231-2212021031022001-1002012330301123-1230302332133301-2221110001303312-2312101011312213-2230212302332203-3302100221013333"></a>

## Direct properties — exclude_violation_contexts / 330003302022 / 3

<a id="canonical-1301232333003013-0312112312303321-1112110003213320-0330231101001300-1333102030210100-3131121302101202-2001121313123323-3212101311103000"></a>

<a id="canonical-1322330300330002-1111222100200203-2202132113020123-1333122312003131-3311111311300332-3023302331132203-0320232002212000-0011213323111303"></a>

## context property — exclude_violation_contexts / 330003302022 / 4

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CONTEXT_ANY","CONTEXT_BODY","CONTEXT_COOKIE","CONTEXT_HEADER","CONTEXT_PARAMETER","CONTEXT_REQUEST","CONTEXT_RESPONSE","CONTEXT_URI","CONTEXT_URL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1303003230231210-0012101312223201-2013011200200300-0000113223012300-3333013231110312-2301120211313211-1111003011233203-2331012000321100"></a>

<a id="canonical-3222201203002112-3001231120230000-2033310022131000-0002213012230220-1323101320312003-3112202211220002-1223310031010233-2333212221013121"></a>

## context_name property — exclude_violation_contexts / 330003302022 / 5

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3313311302333121-3233332223103332-0200323223120033-3223303330111031-0333332212110131-3021032223223300-3322131203301003-3112021223010323"></a>

<a id="canonical-1223032112122110-0120131001321010-0221312120202302-2333221002020223-0222033202020020-0312323320332111-1201212213230011-0000330122201222"></a>

## exclude_violation property — exclude_violation_contexts / 330003302022 / 6

Type: `"string"`. Optional.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
List of all supported Violation Types VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER
VIOL\_HTTP\_RESPONSE\_STATUS VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD
VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED
VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS.. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Upstream description:

List of all supported Violation Types

VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER VIOL\_HTTP\_RESPONSE\_STATUS
VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED
VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS
VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT
VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION
VIOL\_HTTP\_PROTOCOL\_CRLF\_CHARACTERS\_BEFORE\_REQUEST\_START
VIOL\_HTTP\_PROTOCOL\_NO\_HOST\_HEADER\_IN\_HTTP\_1\_1\_REQUEST
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_PARAMETERS\_PARSING
VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS
VIOL\_HTTP\_PROTOCOL\_CONTENT\_LENGTH\_SHOULD\_BE\_A\_POSITIVE\_NUMBER
VIOL\_EVASION\_DIRECTORY\_TRAVERSALS VIOL\_MALFORMED\_REQUEST VIOL\_EVASION\_MULTIPLE\_DECODING
VIOL\_DATA\_GUARD VIOL\_EVASION\_APACHE\_WHITESPACE VIOL\_COOKIE\_MODIFIED
VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS VIOL\_EVASION\_IIS\_BACKSLASHES
VIOL\_EVASION\_PERCENT\_U\_DECODING VIOL\_EVASION\_BARE\_BYTE\_DECODING VIOL\_EVASION\_BAD\_UNESCAPE
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_FORMDATA\_REQUEST\_PARSING
VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST
VIOL\_HTTP\_PROTOCOL\_HIGH\_ASCII\_CHARACTERS\_IN\_HEADERS VIOL\_ENCODING VIOL\_COOKIE\_MALFORMED
VIOL\_GRAPHQL\_FORMAT VIOL\_GRAPHQL\_MALFORMED VIOL\_GRAPHQL\_INTROSPECTION\_QUERY.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIOL_ASM_COOKIE_MODIFIED","VIOL_COOKIE_MALFORMED","VIOL_COOKIE_MODIFIED","VIOL_DATA_GUARD","VIOL_ENCODING","VIOL_EVASION_APACHE_WHITESPACE","VIOL_EVASION_BAD_UNESCAPE","VIOL_EVASION_BARE_BYTE_DECODING","VIOL_EVASION_DIRECTORY_TRAVERSALS","VIOL_EVASION_IIS_BACKSLASHES","VIOL_EVASION_IIS_UNICODE_CODEPOINTS","VIOL_EVASION_MULTIPLE_DECODING","VIOL_EVASION_PERCENT_U_DECODING","VIOL_FILETYPE","VIOL_FILE_UPLOAD","VIOL_FILE_UPLOAD_IN_BODY","VIOL_GRAPHQL_FORMAT","VIOL_GRAPHQL_INTROSPECTION_QUERY","VIOL_GRAPHQL_MALFORMED","VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE","VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION","VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST","VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS","VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST","VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS","VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT","VIOL_HTTP_RESPONSE_STATUS","VIOL_JSON_MALFORMED","VIOL_MALFORMED_REQUEST","VIOL_MANDATORY_HEADER","VIOL_METHOD","VIOL_NONE","VIOL_REQUEST_MAX_LENGTH","VIOL_XML_MALFORMED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIOL_NONE",
  "enum": [
    "VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START",
    "VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1322323233303313-1022122132113100-3302320000102220-1313010033133023-2003330031113321-0332031030233311-0221030231221233-3001331022232313"></a>

## Next pages — exclude_violation_contexts / 330003302022 / 7

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-3022321321013201-0020033302122032-2311011301300323-1202110132110022-1233301221333112-0030120230100322-3103323111103001-1210013301233011)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-0203032203021300-2101223011122200-2011010020200031-1121003201202320-3121133102312123-0230221030232323-0331000320232021-3330311131131220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300323012322210-1300232111003212-1113102222020113-1301010212222212-2122112010131033-0010032222122323-2330011332222123-0121031003013300"></a>

## rule_list.rules.spec.waf_action.none — none / 110301200200 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-1330120330031000-3202113320223100-2303020333222232-1113322320113320-2131320031302300-3021303213200102-0321303213312133-3121220001012010)
- rule_list.rules.spec.waf_action.none

<a id="canonical-1130211132010033-0012023333330123-3122312222333202-2303000132033133-2030012230211010-0123300221231223-1013103130011323-3332101101000120"></a>

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
none = {}
```

<a id="canonical-1102111112012011-3121120011001230-2311123230021002-0313232210021000-0230021300012030-1333021230002102-1220311220021102-3322132313321113"></a>

## Direct properties — none / 110301200200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322032202331303-3332103131332331-0133002012021212-2323201111300023-0233203231311202-1122320333130301-2313200323230333-1132321003030030"></a>

## Next pages — none / 110301200200 / 4

- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-1330120330031000-3202113320223100-2303020333222232-1113322320113320-2131320031302300-3021303213200102-0321303213312133-3121220001012010)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-0332300220323213-1323323030222322-1131312000330222-2232032001110301-0212213131220202-0321130113302233-2100102223233213-2032013001120230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313103221002323-3113320013132202-0322001031303232-2101222130203232-3322130132212213-2111020233312233-0112011121031103-0201112332213011"></a>

## rule_list.rules.spec.waf_action.waf_skip_processing — waf_skip_processing / 202131223013 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-1330120330031000-3202113320223100-2303020333222232-1113322320113320-2131320031302300-3021303213200102-0321303213312133-3121220001012010)
- rule_list.rules.spec.waf_action.waf_skip_processing

<a id="canonical-0332332021232010-3301000202212332-3313331233100032-0132212001212200-3331300332220011-2032001311031303-2120131033213332-1222312203332022"></a>

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
waf_skip_processing = {}
```

<a id="canonical-1121000321212012-1010103013322220-0330233232213121-3101122200330131-3302200031201122-2110000130132222-3220113211200313-0313002322231322"></a>

## Direct properties — waf_skip_processing / 202131223013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003003203223013-2001101223112103-3313312111011123-2010233133132132-1201121201331202-3003020120100300-1132131321012301-0221210323123220"></a>

## Next pages — waf_skip_processing / 202131223013 / 4

- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-1330120330031000-3202113320223100-2303020333222232-1113322320113320-2131320031302300-3021303213200102-0321303213312133-3121220001012010)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-3132031220110012-1111121333101001-0122230010332012-3320330110332133-3010320201321321-1211023113112212-1300122212023203-2020020300213013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111213033320111-2123321322132023-2220022020200300-2200322102012201-0211310030003310-3020300330130103-0002032022232012-0313033120112120"></a>

## server_name_matcher — server_name_matcher / 100031320101 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- server_name_matcher

<a id="canonical-2000320332113331-2121031221100100-1221133200122002-3322012030312113-3103022233103020-1100222220211102-1020331122111223-0203320320001212"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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
server_name_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330223022202120-3003010102103013-1201303310313010-2302113121321311-3331030213301033-1211303020111311-0103000223132322-3303101301323231"></a>

## Direct properties — server_name_matcher / 100031320101 / 3

<a id="canonical-0332001021322311-2131220233010003-1000102232330210-3023121000321032-1122011030103022-3221023131331303-0222331231030333-2200112232233231"></a>

<a id="canonical-2323310033200132-3220121021020110-0223100220032331-1321301013030000-0030132203023211-1230021231002133-3203222022101033-0233001131321313"></a>

## exact_values property — server_name_matcher / 100031320101 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1130030021202301-1031122212120132-2131301011322220-3301123133131033-1331010012013311-2320113302032322-1231120123232022-1213102020330010"></a>

<a id="canonical-2233312231211222-2031313201023011-0221213021112101-0220213033321321-0000303223301020-1023032000131213-0301020133101302-1131032203002312"></a>

## regex_values property — server_name_matcher / 100031320101 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2302130212000311-2130210222221022-2233012102221020-2320110333303202-2111110112231002-3110110003220122-3210131021200102-2020220232232301"></a>

## Next pages — server_name_matcher / 100031320101 / 6

- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-3032002123132023-0202122031333133-0202113320332213-2000203232033130-0232320232222201-3022133310223233-0213130313331202-2000111311332223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132221021210110-1131231221012123-0010010320223013-1113132011120102-0302312121032220-0021023110212321-0020203132200300-3321022032200213"></a>

## server_selector — server_selector / 201210002030 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- server_selector

<a id="canonical-2100103232022233-0113101000311301-3031331021032201-3032212030030022-3321011201111130-3223030011011331-2010300313013200-1101302322210021"></a>

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
server_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221002112131120-3220111313132001-0100211200210111-2011120202132110-3031330000303000-3001321211000312-0003120020220330-1121223332033222"></a>

## Direct properties — server_selector / 201210002030 / 3

<a id="canonical-0223003301030203-2231011011012312-3230302101132103-0203220212030132-0133300202003202-1033110312221230-0011131113030210-0212001131033220"></a>

<a id="canonical-1110202200311311-1320210123103232-0120323113201322-3020001231031023-1012211332203010-1011211203203302-0111111130030202-3203110113231122"></a>

## expressions property — server_selector / 201210002030 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0102302010000213-0023222100220220-3320100131130310-1310030112322133-1330133012223113-2100221013000333-0012230311232133-0030030011111010"></a>

## Next pages — server_selector / 201210002030 / 5

- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)

<a id="canonical-1030022313211210-1123201233300031-2301231122113220-0310203223011103-2122310011221001-2023101232323101-1031133331010301-1220301100022212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311302313332033-2011003323210330-1320321303032311-0231010001022211-2233011203100220-2033213211003001-2231310120133020-2031103320000313"></a>

## timeouts — timeouts / 301312131330 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- timeouts

<a id="canonical-2332330300313210-3032220230231002-2210001000213311-1120030113311201-3323010032130212-0301322332310212-1223300320310131-1002210230210203"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311010220010000-2013003111212122-3302311233333122-3001132211231120-1103133230312202-1223003310212121-2301120333333102-3323020310212103"></a>

## Direct properties — timeouts / 301312131330 / 3

<a id="canonical-0130223333312220-1300211333313231-3021121330032221-2302232003113011-3203233212213213-0200122002320223-2213001103032131-2310210302000030"></a>

<a id="canonical-3113232112320323-1320332021210131-0133313000333201-1012311212002012-1322132203010022-1131321030230331-1101310010201000-3210110210033301"></a>

## create property — timeouts / 301312131330 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1031302123101121-1101013020003133-0032312221131121-1100303013031110-1322133010222220-3001330301302222-1303300033011013-3211003211212222"></a>

<a id="canonical-3322101023311030-2213103003111003-1023012212310211-0011123323200120-2221310213011302-0112221122131021-1303203113313323-1113220123123031"></a>

## delete property — timeouts / 301312131330 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3222332311310031-1300301003212221-1121213321301231-2303133033030033-0131120223312201-0130113013130102-2223121101102113-3010302222133212"></a>

<a id="canonical-3301211323311013-0113211113120222-2211130020001010-1120013332222100-3033312232231231-0001233120001032-3321231203211002-0311031200221010"></a>

## read property — timeouts / 301312131330 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2223033213323213-0023113311030323-2230232220032313-0212011012110310-2023321002030132-0232001232202000-0121200122002113-2300001313013322"></a>

<a id="canonical-2300311323331131-3000202032202101-2013133120022033-2330200133321200-0300302022000002-1323112021210131-0302021100220123-0013021100301010"></a>

## update property — timeouts / 301312131330 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3223111030210022-1022012223223230-0300221210312330-0210311230323102-3223100313001311-3212201033323302-3332111311303000-2322120312011311"></a>

## Next pages — timeouts / 301312131330 / 8

- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
