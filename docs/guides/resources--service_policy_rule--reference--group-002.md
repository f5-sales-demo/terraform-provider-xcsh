---
page_title: "xcsh_service_policy_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule reference."
---

# xcsh_service_policy_rule reference

<a id="canonical-0102022231201013-1330312103021123-2211102010033323-2331222132131330-1111020313030011-3321233012202120-0020232303113033-3030322101200032"></a>

## name property — jwt_claims / 013131122032 / 5

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0210212132233321-2313323022030202-2101330101132322-0032120220230122-0212032031210122-3220122101323310-0211331030102321-1113121002331002"></a>

## Next pages — jwt_claims / 013131122032 / 6

- [jwt_claims.check_not_present](resources--service_policy_rule--reference--group-002.md#canonical-2303320103130132-2101102022031220-3233210122010200-3132323231322222-1202332312102313-0031113302231103-1221230211023220-1321231232313002)
- [jwt_claims.check_present](resources--service_policy_rule--reference--group-002.md#canonical-2031030212122331-1231122323313101-1302300202103032-0110222223333122-0001021313011023-2200310200000100-3330131313202003-3320120231032210)
- [jwt_claims.item](resources--service_policy_rule--reference--group-002.md#canonical-2311232010330201-3102323300101200-2132333132303330-0332202102312221-0220131212103330-0210002331320300-1023120010330210-2111203310303201)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2303320103130132-2101102022031220-3233210122010200-3132323231322222-1202332312102313-0031113302231103-1221230211023220-1321231232313002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302121323300223-1321031003031000-3132212213022200-3311103123222221-3223312121110120-3020313033120012-2133010302132320-1003132111031333"></a>

## jwt_claims.check_not_present — check_not_present / 321001031022 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-3121210103300010-0232302121023210-3221102033333021-3313032210213000-2323222213320112-2302302332002332-3203103000231332-2012033133032113)
- jwt_claims.check_not_present

<a id="canonical-3211003220132220-2111200033001312-0211232230111330-3000011132030100-2110300023103031-0203130031210320-3200222002201133-3221130312102132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-3023230213330200-0030132213331130-1301033202212310-2212033213210213-0021223331012132-2112212212301330-2311030133111333-1122001010302103"></a>

## Direct properties — check_not_present / 321001031022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223313113302013-0122000113132202-3232210223232111-1012301202110330-1013331300113130-1023103201233320-1112000311010132-1013332132113203"></a>

## Next pages — check_not_present / 321001031022 / 4

- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-3121210103300010-0232302121023210-3221102033333021-3313032210213000-2323222213320112-2302302332002332-3203103000231332-2012033133032113)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2031030212122331-1231122323313101-1302300202103032-0110222223333122-0001021313011023-2200310200000100-3330131313202003-3320120231032210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021023012303101-1200000111101120-2321312322212221-0300032233310200-0322223012023101-1233121020331023-0113330110212003-3123020103101302"></a>

## jwt_claims.check_present — check_present / 022110313303 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-3121210103300010-0232302121023210-3221102033333021-3313032210213000-2323222213320112-2302302332002332-3203103000231332-2012033133032113)
- jwt_claims.check_present

<a id="canonical-3010233313020311-0333203110111132-3110113212022133-1020022110221232-3232011133320320-0003022122322222-0123212200202213-1032013222333231"></a>

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

<a id="canonical-0331333310002223-0011033021121220-0023322101000111-0001320133231013-1020021132100210-2003110221232322-3222233100302213-0312232003202101"></a>

## Direct properties — check_present / 022110313303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330301013220302-3013302320223032-1312223111230202-3122222020111232-3110020113113221-0220122202331211-2221033130013202-1113031233213303"></a>

## Next pages — check_present / 022110313303 / 4

- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-3121210103300010-0232302121023210-3221102033333021-3313032210213000-2323222213320112-2302302332002332-3203103000231332-2012033133032113)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2311232010330201-3102323300101200-2132333132303330-0332202102312221-0220131212103330-0210002331320300-1023120010330210-2111203310303201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021222330100113-3301003320111303-2311303300202321-2303032030200022-1223231301133011-1121012132100100-2322103031020022-3103220220303110"></a>

## jwt_claims.item — item / 132100323302 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-3121210103300010-0232302121023210-3221102033333021-3313032210213000-2323222213320112-2302302332002332-3203103000231332-2012033133032113)
- jwt_claims.item

<a id="canonical-3122203122311031-0320223010303311-2303023021112000-1200201201000221-2321323301132011-0002230313021110-1312102132223122-3221301301011011"></a>

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

<a id="canonical-2201201221201131-1221222320003230-0112100301330321-2120111111213313-0023111300110130-3231302323313202-2222310312021203-3130123200131322"></a>

## Direct properties — item / 132100323302 / 3

<a id="canonical-1131203020211002-1132013133221130-2330323330200033-3320120000301132-2021301102313331-2333232232133302-2001322133101011-1003030101331331"></a>

<a id="canonical-0311131213032130-3212233221030123-1312211210333311-3230230302130022-2320010333103301-1000121230223023-2200132231100310-3211000110030022"></a>

## exact_values property — item / 132100323302 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0333211231120231-2221101333230013-2332220003030120-0320011313320131-3322220312003013-0201102031323122-3111010002110233-0121101120223200"></a>

<a id="canonical-2313330001133231-1103333123020231-2210033000011232-2101332213000213-1212110300122301-1120121022132322-0030003132310131-0133100023232310"></a>

## regex_values property — item / 132100323302 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2201100002003013-0013312020120131-3123210331020232-3233123332033123-3321302130020302-0013031102222131-3132331313323100-2203111311220022"></a>

<a id="canonical-2213332033221030-1030003332201123-2220121023133032-0210103203030231-1132201000000033-2331022103132103-0303232130113131-0132110220113223"></a>

## transformers property — item / 132100323302 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3300333202213010-2220323112121111-2113013112033202-0030131310000301-2010323030133303-0033322023232110-1101333303232120-2003031003110033"></a>

## Next pages — item / 132100323302 / 7

- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-3121210103300010-0232302121023210-3221102033333021-3313032210213000-2323222213320112-2302302332002332-3203103000231332-2012033133032113)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2020111211100111-2033231320322121-0121111211102333-3113322310323132-1300301223000233-0231032111303233-0333010313331210-0100231321203000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123110323220300-0113311120120012-1110011111232311-3011302013222013-1311202220110311-1030303332223020-0313231332131313-0203132011330123"></a>

## label_matcher — label_matcher / 121102032232 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- label_matcher

<a id="canonical-1312003200110013-1123113121200330-0303232330132123-3130321201330321-2333032023303210-3110021030303000-3113023202210211-2302312332233300"></a>

Type: `"object"`. single nested block, Optional.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

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

<a id="canonical-2300331100330103-2301211212001302-2320011321122301-2000300313131000-3213322003102202-1033233020302121-0322213000022110-3022121331121212"></a>

## Direct properties — label_matcher / 121102032232 / 3

<a id="canonical-2303011232220331-2131220312311033-1100103323302001-1132322033300202-0223223333113312-1103211202000323-0011012300003323-0312120022330012"></a>

<a id="canonical-3102000231232222-1330202223112231-0120122231233100-0212223213223310-0022123313332012-3312013331332321-2100021021223001-3101222000203230"></a>

## keys property — label_matcher / 121102032232 / 4

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3122233001313110-2232302033223330-1201003222101033-3002331030131301-3010031210032211-2321221231320333-2331021101321320-3232133321112032"></a>

## Next pages — label_matcher / 121102032232 / 5

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2001130322132301-0223123123302223-3201113011220210-2021122302201100-0100103003130323-3131100121201120-2120001312003313-2010030020332131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220201102001210-0021123312200103-1001201203032311-3023013020310220-2022203111122001-3221121000230013-3221223220132012-1311000103132021"></a>

## mum_action — mum_action / 101023223300 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- mum_action

<a id="canonical-2221121310120320-0122232013020111-3111313002230103-0012100300022122-2233030220321313-1231102130223133-0323221002032002-3032000311011100"></a>

Type: `"object"`. single nested block, Optional.

Modify behavior for a matching request. The modification could be to entirely skip processing.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0332012031333121-0312202232121302-0220310203320313-0313303321103200-0000323332011220-1113223131131122-1312313120000022-1031313010023013"></a>

## Direct properties — mum_action / 101023223300 / 3

- [default](resources--service_policy_rule--reference--group-002.md#canonical-3211102001223032-0311103320033133-3130122202310202-0331113011003213-3201120320303100-2321003200213302-1200120123321323-3112121031201102): complete subsection reference.

- [skip_processing](resources--service_policy_rule--reference--group-002.md#canonical-0122323332132112-1310201121311332-0000021312313130-2303132030300200-1103301333032313-0023213311211230-3300103103220222-3232213011100333): complete subsection reference.

<a id="canonical-3023110211021331-0012113201320130-3300033321103203-1213331311011101-2013300130201103-0321300302101311-0320313102033000-2003212303232321"></a>

## Next pages — mum_action / 101023223300 / 4

- [mum_action.default](resources--service_policy_rule--reference--group-002.md#canonical-3211102001223032-0311103320033133-3130122202310202-0331113011003213-3201120320303100-2321003200213302-1200120123321323-3112121031201102)
- [mum_action.skip_processing](resources--service_policy_rule--reference--group-002.md#canonical-0122323332132112-1310201121311332-0000021312313130-2303132030300200-1103301333032313-0023213311211230-3300103103220222-3232213011100333)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-3211102001223032-0311103320033133-3130122202310202-0331113011003213-3201120320303100-2321003200213302-1200120123321323-3112121031201102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033220121012033-0301303020013012-0003300213332000-3013222100331032-3231003233033201-3321212322011311-3131002321233320-2033200030111022"></a>

## mum_action.default — default / 212111221112 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [mum_action](resources--service_policy_rule--reference--group-002.md#canonical-2001130322132301-0223123123302223-3201113011220210-2021122302201100-0100103003130323-3131100121201120-2120001312003313-2010030020332131)
- mum_action.default

<a id="canonical-2102113111213301-3030032320301122-0101232002213212-0322211221310200-1101210110003000-2131012033221002-1022221121113013-1031100210203302"></a>

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
default = {}
```

<a id="canonical-0033302120210221-1133313110223110-2131011122321302-2220330111023030-1110232101023100-1023333221021301-1132231121023100-2031200233113030"></a>

## Direct properties — default / 212111221112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310202123100212-2302233303023212-2120111102313120-0102220312212230-2032303123203201-0211331032330332-0333213001233130-2313023312023003"></a>

## Next pages — default / 212111221112 / 4

- [mum_action](resources--service_policy_rule--reference--group-002.md#canonical-2001130322132301-0223123123302223-3201113011220210-2021122302201100-0100103003130323-3131100121201120-2120001312003313-2010030020332131)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-0122323332132112-1310201121311332-0000021312313130-2303132030300200-1103301333032313-0023213311211230-3300103103220222-3232213011100333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100320110111300-3320310222203310-1203200200202003-0002320131002303-3211022231322033-1230210322232303-2321031022032131-2131212332332231"></a>

## mum_action.skip_processing — skip_processing / 310032220203 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [mum_action](resources--service_policy_rule--reference--group-002.md#canonical-2001130322132301-0223123123302223-3201113011220210-2021122302201100-0100103003130323-3131100121201120-2120001312003313-2010030020332131)
- mum_action.skip_processing

<a id="canonical-0312310211231330-1203021111031102-3103233002012120-1210130133030212-1322020300113131-3311120303211301-3330102301313021-0213020111213311"></a>

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
skip_processing = {}
```

<a id="canonical-3122000131332323-1113113002203111-1211000112033010-3210132330122133-0120230031212311-2201311311110201-2031130023103221-3230002022221331"></a>

## Direct properties — skip_processing / 310032220203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331332330020101-2102232213100011-0210130300113113-0130231113300103-1020122211232002-2130033223013000-1322022003032030-1112131320231103"></a>

## Next pages — skip_processing / 310032220203 / 4

- [mum_action](resources--service_policy_rule--reference--group-002.md#canonical-2001130322132301-0223123123302223-3201113011220210-2021122302201100-0100103003130323-3131100121201120-2120001312003313-2010030020332131)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-3010110322330023-1103101211122201-1133323113223332-1313232321223110-0101110303212213-1030200111110001-0302033113300012-0303132213131330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030213002123113-1100011232010032-1333302032330312-0122202210023012-2302023020220130-3101002112120131-2201322011210322-1231012300310233"></a>

## path — path / 122203120302 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- path

<a id="canonical-1222332103311013-2322202031102010-2310333211132301-2133303132123131-3101032220033221-2332323101120012-0023222202301233-0100221110233300"></a>

Type: `"object"`. single nested block, Optional.

Path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

Upstream description:

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

<a id="canonical-3232113221313231-2201233211310220-3203302110222111-2000010103322220-2331100303102222-2123310030022131-3120211312102221-2102220032213021"></a>

## Direct properties — path / 122203120302 / 3

<a id="canonical-1203321001120102-3230123100011102-3310122110021233-1010032220322303-1321202311111101-2101232203313223-3031113123221003-0123331113312113"></a>

<a id="canonical-3102120230201203-0121323013002121-3133210112231203-3003210320020332-0031000130300303-2131121202132230-1321120132202013-1022321320211312"></a>

## encoded_path_matcher property — path / 122203120302 / 4

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

<a id="canonical-3300211023023011-0100330002020032-0221232220001003-2023130313203113-1110113023021102-0203110332303122-0301132030231013-0232211211002022"></a>

<a id="canonical-3133110010320230-1322122213310131-2032031023031003-1102021120023220-3131330003233132-3310033033002113-3003220221020022-3111330302130323"></a>

## exact_values property — path / 122203120302 / 5

Type: `["list", "string"]`. Optional.

List of exact path values to match the input HTTP path against.

Upstream description:

A list of exact path values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0122202103323232-3332313000012033-0031003032201200-1303213100000221-1122210032020210-3221322012000233-0230203021001300-2003012113330110"></a>

<a id="canonical-3320212203220302-3002023322330112-2111301302122100-3121110100120013-1320122311220331-0100200101223101-2010302020302032-0221211001200222"></a>

## invert_matcher property — path / 122203120302 / 6

Type: `"bool"`. Optional.

Invert Path Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-0033131231323313-2133303031222313-3330221011331312-0132022203002230-3233300231311030-1232003330022330-1003213132011232-2033012301112130"></a>

<a id="canonical-0331321321312231-2301111111330300-2113300312322310-2002112313202333-2030033311113010-1222200031223121-1211113313013020-0122312113032000"></a>

## prefix_values property — path / 122203120302 / 7

Type: `["list", "string"]`. Optional.

List of path prefix values to match the input HTTP path against.

Upstream description:

A list of path prefix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3231203123331011-1010021300103133-3100010310231333-2112002200321102-2213301011312013-2121110313020232-3301210032010221-2203323010311321"></a>

<a id="canonical-1121102022212021-0301212113202233-1020220211131331-1012113032201213-3310122112012130-1000233211201110-2212303220300322-0001312112200231"></a>

## regex_values property — path / 122203120302 / 8

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input HTTP path against.

Upstream description:

A list of regular expressions to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0012332022230031-0201113331012212-0221033003102023-2030201101331300-2322202311021232-1102202303102103-0022332222100002-3123213102000001"></a>

<a id="canonical-0321112200333201-0122301002231013-3201130013101021-1231310101203330-1212302311303002-0231012000200031-0331210323212102-2010322312323010"></a>

## suffix_values property — path / 122203120302 / 9

Type: `["list", "string"]`. Optional.

List of path suffix values to match the input HTTP path against.

Upstream description:

A list of path suffix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2332221310210123-1303210232003020-1011212330201122-0320023322033321-0332301020311103-1223013320302002-0213010132303030-1212211131110230"></a>

<a id="canonical-3022330303023102-2312202333002300-3023100100300110-0033210202000121-0113200122320220-3112112003013211-1323013233023333-2000212022001211"></a>

## transformers property — path / 122203120302 / 10

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3230023012030103-3123203033202023-1122302332012312-3231230331333100-1233103302120023-3312002222002221-0103103203213011-1210313312332233"></a>

## Next pages — path / 122203120302 / 11

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2203222212003231-2202300100330200-1310211032320202-1313000110311333-0331313210023112-2112333302313030-1110023000021210-2033311231020133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032130301303201-3310332220312231-2110203123201232-2330310112331221-2012131233111210-0011001330332132-0110220232022122-1113011011030131"></a>

## port_matcher — port_matcher / 331200322310 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- port_matcher

<a id="canonical-2322112120320222-3113003012110013-3110231221003322-0221120100321303-2132110222311133-2030211130312102-0231232102133322-2023023231200211"></a>

Type: `"object"`. single nested block, Optional.

Port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true. Server applies default when omitted.

Upstream description:

A port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0232303133221212-3332131223200113-2121331003110032-3131033211230031-1201202200123312-2131122022232322-1103300110112000-1131113331022031"></a>

## Direct properties — port_matcher / 331200322310 / 3

<a id="canonical-0120221311003100-1000230031222100-1322323000312032-1031130332211231-1210111013010030-3210012031021302-3103210102000132-3032321110230021"></a>

<a id="canonical-0230113100110202-0121311000101310-1320122131110002-0013202101013321-1310200002031033-3213110003333113-2001113130121012-2312113033100121"></a>

## invert_matcher property — port_matcher / 331200322310 / 4

Type: `"bool"`. Optional.

Invert Port Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-2223320321103203-0023122212323001-3002223020320021-3213202020100230-1231130222132002-0132303232200030-3131023001013101-2010121223331303"></a>

<a id="canonical-1033311001012303-0210223221211301-2330112333220133-0031231330322210-0133120100203121-1301112103103231-1120330110330130-2000211300332202"></a>

## ports property — port_matcher / 331200322310 / 5

Type: `["list", "string"]`. Optional.

List of strings, each of which is a single port value or a tuple of start and end port values
separated by '-'. The start and end values are considered to be part of the range.

Upstream description:

A list of strings, each of which is a single port value or a tuple of start and end port values
separated by "-". The start and end values are considered to be part of the range.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3030110101211101-1333123210312221-2321303002111213-1203102222231103-3222001021033220-2333120323012012-1100303113000113-2010123323301123"></a>

## Next pages — port_matcher / 331200322310 / 6

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-1012032130001102-2011001331322032-1103021200202322-0322031221221110-0133202010012100-1121312101103032-2022003321212212-0312123221010101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301200300003021-2303331200211221-2123130200012202-1201200111103323-1113302221122103-2112000032311330-1033133302222322-0301010333112101"></a>

## query_params — query_params / 332312032030 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- query_params

<a id="canonical-1132121303300022-3222130331310311-3012221303320222-0300303322210021-2012131020223121-0203302231323312-3113331121030333-0011112032010322"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1211110211301031-0103111210102021-2201033020001011-3103130010320120-3301002313013200-0322010123020301-2202101203100230-2231303100212332"></a>

## Direct properties — query_params / 332312032030 / 3

- [check_not_present](resources--service_policy_rule--reference--group-002.md#canonical-2233100010311321-0013300020103330-2232123000232130-2313333300323232-2021121122103012-3321120200110032-2222013322013111-3111323320011102): complete subsection reference.

- [check_present](resources--service_policy_rule--reference--group-002.md#canonical-2213203220132122-1123311113310120-0200132121322222-1320110200011231-3213213300100303-2100122210033231-1120321110122000-1001021132121000): complete subsection reference.

<a id="canonical-2033101023010311-0310112222301102-2100232312001011-1100313301133203-1131022123000031-0102213230111001-1303220330120300-0212300011311202"></a>

<a id="canonical-0321302201211212-2332030000302032-0001133000113303-1120123103130103-1021010003321133-1310113312132210-1130111210132302-2102223001323221"></a>

## invert_matcher property — query_params / 332312032030 / 4

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--service_policy_rule--reference--group-002.md#canonical-1311331210221220-1223101220322011-0301012322132023-0333111323012133-0011313130123200-3102231330313002-2322220322101021-1202323313133012): complete subsection reference.

<a id="canonical-2100301210133310-2333223312301201-3131000303310131-1333223330212101-2121213331212121-2031222321033333-0202130301313312-3231000113313312"></a>

<a id="canonical-1300031132211323-1132123113111133-2123133100212133-1131031013033312-3013013023330023-3103221223330102-0002220013210003-3211221322011301"></a>

## key property — query_params / 332312032030 / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3310203301003031-2203211222323002-2102122130231023-0020130320132301-2323320203333213-1122220211131321-1032002210212220-3002133203322000"></a>

## Next pages — query_params / 332312032030 / 6

- [query_params.check_not_present](resources--service_policy_rule--reference--group-002.md#canonical-2233100010311321-0013300020103330-2232123000232130-2313333300323232-2021121122103012-3321120200110032-2222013322013111-3111323320011102)
- [query_params.check_present](resources--service_policy_rule--reference--group-002.md#canonical-2213203220132122-1123311113310120-0200132121322222-1320110200011231-3213213300100303-2100122210033231-1120321110122000-1001021132121000)
- [query_params.item](resources--service_policy_rule--reference--group-002.md#canonical-1311331210221220-1223101220322011-0301012322132023-0333111323012133-0011313130123200-3102231330313002-2322220322101021-1202323313133012)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2233100010311321-0013300020103330-2232123000232130-2313333300323232-2021121122103012-3321120200110032-2222013322013111-3111323320011102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130021120101312-3221122020232100-3323201203130320-3123233321033220-2031010222331321-0121010210000201-1223102233230133-0020120310032222"></a>

## query_params.check_not_present — check_not_present / 102333001201 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [query_params](resources--service_policy_rule--reference--group-002.md#canonical-1012032130001102-2011001331322032-1103021200202322-0322031221221110-0133202010012100-1121312101103032-2022003321212212-0312123221010101)
- query_params.check_not_present

<a id="canonical-3202330001122331-1130330220001301-2201120212330221-2000333031220132-0331013333332133-0030011200120031-0232302232121212-1200221333202002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-1332102321223323-0321332131010031-3100132120313031-3032020330200112-3203320212301232-1022023301321212-0133032103123131-2120301020203221"></a>

## Direct properties — check_not_present / 102333001201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121231331020323-1232210323132113-3202111111113223-3133202230231321-1232332112121310-0333130101112300-0120000213312211-1303122012330120"></a>

## Next pages — check_not_present / 102333001201 / 4

- [query_params](resources--service_policy_rule--reference--group-002.md#canonical-1012032130001102-2011001331322032-1103021200202322-0322031221221110-0133202010012100-1121312101103032-2022003321212212-0312123221010101)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2213203220132122-1123311113310120-0200132121322222-1320110200011231-3213213300100303-2100122210033231-1120321110122000-1001021132121000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203023230310101-0102013111210131-1332020211131121-0102331332032013-2332130131301211-0313121212101123-0323230320003113-2222221011213313"></a>

## query_params.check_present — check_present / 010030000312 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [query_params](resources--service_policy_rule--reference--group-002.md#canonical-1012032130001102-2011001331322032-1103021200202322-0322031221221110-0133202010012100-1121312101103032-2022003321212212-0312123221010101)
- query_params.check_present

<a id="canonical-0203031220002032-2321320122323321-0330112110213330-0210112333032311-2300003111301101-1133223023031032-1103032000122312-3123133310113121"></a>

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

<a id="canonical-0333133023223300-3333100333220323-0230131002201202-0022312011113123-1001301313211120-0213013202000232-2013010301030112-0322112312123331"></a>

## Direct properties — check_present / 010030000312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300331003003103-2301113303033032-1110111312332223-2231003322201002-0311020212223213-1133003203122111-1100303102100013-2103011203020302"></a>

## Next pages — check_present / 010030000312 / 4

- [query_params](resources--service_policy_rule--reference--group-002.md#canonical-1012032130001102-2011001331322032-1103021200202322-0322031221221110-0133202010012100-1121312101103032-2022003321212212-0312123221010101)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-1311331210221220-1223101220322011-0301012322132023-0333111323012133-0011313130123200-3102231330313002-2322220322101021-1202323313133012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012110112031103-0020130230300002-1103023101012131-0223013000331203-0113220000000012-0330103031000232-1313332132202000-2320310332221133"></a>

## query_params.item — item / 003021232132 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [query_params](resources--service_policy_rule--reference--group-002.md#canonical-1012032130001102-2011001331322032-1103021200202322-0322031221221110-0133202010012100-1121312101103032-2022003321212212-0312123221010101)
- query_params.item

<a id="canonical-2233103232121213-2131110023313322-0313133320302120-1121000310110033-2111303130113330-0132132000302301-2101021032003213-3132031301203312"></a>

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

<a id="canonical-3201022212200211-3222202301022310-3020023331110231-3012123201122001-3103121010033012-3120210102302122-1122000020300130-2201030112122010"></a>

## Direct properties — item / 003021232132 / 3

<a id="canonical-3323111210031331-1133123113121110-3131322121022313-1122302110012211-0011213013231301-2301000132230213-3332312100301303-2113311203232112"></a>

<a id="canonical-2303101212301033-3103313231103000-0133330213200011-2112312002300320-3333221102010333-2300323013303120-1222002232123123-0320020012001133"></a>

## exact_values property — item / 003021232132 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1003310032001103-3332232022122202-1031211322230031-3033231223011030-3201321033301331-3002013130130221-3000231122302010-2123131313211021"></a>

<a id="canonical-3212302310323122-3013211232220221-3102103100220110-2200010333030012-0232123100223120-2030331030112033-3122220022021131-3021103003113213"></a>

## regex_values property — item / 003021232132 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0201023212001002-2020130011313311-2230133011222110-2032132231300101-1303230311113220-3320322221201132-1332030201323221-2230300212220301"></a>

<a id="canonical-0112220113322220-0131221122011002-0111013223103230-1222121331122032-3121231203120311-1210312033212110-1232011202133300-0011102333303112"></a>

## transformers property — item / 003021232132 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2332222313212031-3320000320110313-1303002320221011-1111302120130220-0113002232033031-0103203010310331-3333030233300120-0310002211101200"></a>

## Next pages — item / 003021232132 / 7

- [query_params](resources--service_policy_rule--reference--group-002.md#canonical-1012032130001102-2011001331322032-1103021200202322-0322031221221110-0133202010012100-1121312101103032-2022003321212212-0312123221010101)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010313321232010-1200202213331203-2130000310110112-2303030200003332-3110133011013202-1111201031323112-0033031012122320-3010030330311323"></a>

## request_constraints — request_constraints / 332123312223 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- request_constraints

<a id="canonical-3321300223001100-2120321321003300-3311003010320003-2302220203322302-3333221002312102-0223300030120210-0203232233110201-3321030211133300"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request constraints.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2132312311232302-1133303013120203-2312321213222231-1123223111003300-1012221031131120-0033112132033121-3030303223110113-1300123130012223"></a>

## Direct properties — request_constraints / 332123312223 / 3

<a id="canonical-2313200310211300-0211203202113123-0110012003032132-2111022330333210-2011232022312322-0302102323031210-2233021313131230-2331300203213300"></a>

<a id="canonical-3100102323203333-1221332333203330-1100023312210003-3033232022101202-0303211133222322-2303021002100332-3100111131302000-2130133300121013"></a>

## max_cookie_count_exceeds property — request_constraints / 332123312223 / 4

Type: `"number"`. Optional.

Match on the Count for all Cookies that exceed this value. Exclusive with
\[max\_cookie\_count\_none\]

Upstream description:

Exclusive with \[max\_cookie\_count\_none\]

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_cookie_count_none](resources--service_policy_rule--reference--group-002.md#canonical-1122022011000132-2310232120103311-1302322200300230-2231133020012123-0101113203221130-3103331101010130-0200032210220122-2333200330022310): complete subsection reference.

<a id="canonical-2020212101221102-2131103121230032-3211101120300013-1030032123131233-2001211302331313-0103012001113311-2113222133001310-3123100200033100"></a>

<a id="canonical-2311303021203232-0132001003313010-2013333110323331-2000130331120023-2130010023132310-3020130113023310-0302301223031112-2222020022000033"></a>

## max_cookie_key_size_exceeds property — request_constraints / 332123312223 / 5

Type: `"number"`. Optional.

Exclusive with \[max\_cookie\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_key\_size\_none\]

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_cookie_key_size_none](resources--service_policy_rule--reference--group-002.md#canonical-3210311330313210-1030032003013232-2310012332113123-0200213031011110-3032031210102331-2131213220101131-3321201211111213-3111123300123122): complete subsection reference.

<a id="canonical-1021031032330313-1110311310220223-1130223231212220-1123010110233113-0020332200121200-1120100002222312-2321333002213310-0131331132002221"></a>

<a id="canonical-3331002233210112-3302321030230122-3322223201031102-0011212113332322-1220021231132332-1303213101320332-3330212001232000-3213232211313332"></a>

## max_cookie_value_size_exceeds property — request_constraints / 332123312223 / 6

Type: `"number"`. Optional.

Exclusive with \[max\_cookie\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_value\_size\_none\]

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_cookie_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-2210332330211102-3323302021220120-0333020011333130-3231233221330020-0311010101023031-1213120210302130-3032020321123303-3220013033000320): complete subsection reference.

<a id="canonical-3223011123303132-0200013211133320-1333131213002323-3011012101112031-0231010330311111-1120113021100003-1212310200231232-0102032323220303"></a>

<a id="canonical-1132031131300201-1120302210233032-3002311311002011-1302200221013230-1100110333132002-0100133311322013-3131203030000022-2122110001012031"></a>

## max_header_count_exceeds property — request_constraints / 332123312223 / 7

Type: `"number"`. Optional.

Match on the Count for all Headers that exceed this value. Exclusive with
\[max\_header\_count\_none\]

Upstream description:

Exclusive with \[max\_header\_count\_none\]

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_header_count_none](resources--service_policy_rule--reference--group-002.md#canonical-2031333020012231-3100132320333220-1321002111211212-1121322112320113-1033112301122111-3221032333001212-3332000110133331-1031013122313000): complete subsection reference.

<a id="canonical-0100220101123010-3112011323022110-1113131203013001-3322030123200030-2111302021101313-2133201313112312-0303122231321113-1113123000211112"></a>

<a id="canonical-2222310223030030-3202100200200003-1021110310233331-1122212320100120-0233132131212220-3323013121020100-0131113121033021-3301112103333111"></a>

## max_header_key_size_exceeds property — request_constraints / 332123312223 / 8

Type: `"number"`. Optional.

Exclusive with \[max\_header\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_key\_size\_none\]

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_header_key_size_none](resources--service_policy_rule--reference--group-002.md#canonical-1132101021313102-3332033311133003-2210133001000322-0010232331203032-2322133101011321-1202032320332312-3132033012032313-3332110002331110): complete subsection reference.

<a id="canonical-2022303122012113-0223332122112030-2223312303331122-1120322031002123-0322301012001233-1331210232313113-1002202203022222-0300111231011013"></a>

<a id="canonical-2110132002230123-2133210233202002-0310103021133012-1033200101002110-0120232233303030-1330312012011331-1230012103030112-3020001110133302"></a>

## max_header_value_size_exceeds property — request_constraints / 332123312223 / 9

Type: `"number"`. Optional.

Exclusive with \[max\_header\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_value\_size\_none\]

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_header_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-2102120230210231-2100101103110120-1123203011201001-3031333211322330-1210231111011330-2310102033001102-1320213011013311-0021120132333103): complete subsection reference.

<a id="canonical-1233030203201113-1233220312102010-3013333203201322-1220312010103230-3222021021222031-2102103331312232-0212133332321003-0021011200132000"></a>

<a id="canonical-1303120311332031-2112221311120100-2131203221102302-2212232021033332-1303033030220112-2033220010303000-0033102331021133-1133312121100201"></a>

## max_parameter_count_exceeds property — request_constraints / 332123312223 / 10

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_count\_none\].

Upstream description:

Exclusive with \[max\_parameter\_count\_none\]

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_parameter_count_none](resources--service_policy_rule--reference--group-002.md#canonical-3221300301123333-0233121001221021-1311302003000011-2010203103033011-3100131213301022-0011232323122220-0300231113313310-1313310111333232): complete subsection reference.

<a id="canonical-0020332332023322-2123001321300233-0110333230222311-0032031322110002-1102101112103112-1033320301221033-1330032202122220-1302203122201011"></a>

<a id="canonical-1132331031231100-0031203203302003-1011221311120002-3233330122311320-0003203031220202-3132201113030200-3202103232131203-1223031332120231"></a>

## max_parameter_name_size_exceeds property — request_constraints / 332123312223 / 11

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_name\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_name\_size\_none\]

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_parameter_name_size_none](resources--service_policy_rule--reference--group-002.md#canonical-2230202133131132-2331002310102321-2203332230123312-3120102201131031-2000222113110033-2122102213110001-3312010333210032-0032213011111222): complete subsection reference.

<a id="canonical-0030320130001130-3333231301332021-2320213102111133-2332312131010011-2320030020100132-2333021032223221-3201132230222313-1302330101111230"></a>

<a id="canonical-0012001101012313-2101111330322311-2001020120021103-2102102331311203-2222313122301232-2123002201333331-1010330323223020-1032023303003112"></a>

## max_parameter_value_size_exceeds property — request_constraints / 332123312223 / 12

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_value\_size\_none\]

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_parameter_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-1023121102232230-3110132110130012-2003011223012323-0330131103322022-1233311030100300-0113123310232012-3233100301000331-3200002330112200): complete subsection reference.

<a id="canonical-3223021010002003-3011033230333121-3132233230301132-0133001310023221-3331100321113030-1233121111210133-2002313321310202-3330332331001332"></a>

<a id="canonical-2233111203233002-3113012332313333-1203012113132300-3321313223322101-3330203201123233-2200123022100301-1133230223211020-3222111220133222"></a>

## max_query_size_exceeds property — request_constraints / 332123312223 / 13

Type: `"number"`. Optional.

Match on the URL Query Size that exceed this value. Exclusive with \[max\_query\_size\_none\]

Upstream description:

Exclusive with \[max\_query\_size\_none\]

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_query_size_none](resources--service_policy_rule--reference--group-002.md#canonical-0100103003131210-2220030203003130-1213013113101012-2123313132020002-3211122013110230-0202020232331211-3111300202321122-1301331102130200): complete subsection reference.

<a id="canonical-2111022021010231-3030203322312131-0133222011013033-2230101220213302-2323112200100313-2122120333133031-3020333331103222-3133012301010332"></a>

<a id="canonical-0102031031003102-1311332313300303-2312322332212321-3031133102313113-3322320131010310-2020001220123223-3022130201030113-0000123320012120"></a>

## max_request_line_size_exceeds property — request_constraints / 332123312223 / 14

Type: `"number"`. Optional.

Exclusive with \[max\_request\_line\_size\_none\].

Upstream description:

Exclusive with \[max\_request\_line\_size\_none\]

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_request_line_size_none](resources--service_policy_rule--reference--group-002.md#canonical-0000013323113031-2033013133320032-2010233210331233-2202002001102333-3222323012130033-2132022123121011-0231320022032213-0133013002000130): complete subsection reference.

<a id="canonical-3011003132011022-0130310200023122-2020210310200003-3111113130013010-1000222210303013-2302222331212011-1323123120232121-2303101010230003"></a>

<a id="canonical-2121022102100213-2122112210023032-1011102311123031-2210222031232113-3010302021310311-3313103320230012-2120313101323320-1200313230220132"></a>

## max_request_size_exceeds property — request_constraints / 332123312223 / 15

Type: `"number"`. Optional.

Match on the Request Size that exceed this value. Exclusive with \[max\_request\_size\_none\]

Upstream description:

Exclusive with \[max\_request\_size\_none\]

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_request_size_none](resources--service_policy_rule--reference--group-002.md#canonical-0211021211211232-0231130103220023-2211300110200302-1330123120002032-1233331131120002-0331131010010032-1001003322020101-2232031202032200): complete subsection reference.

<a id="canonical-3002332102130132-1023102213020233-3301321313011303-3130311311302021-1211031221320201-3312013130003120-3001020230310222-0111101030121022"></a>

<a id="canonical-1032113121233110-0223101212012321-3003322230322003-2010321001101013-3302320222213232-1100323131013232-1010003121003130-2131031013123120"></a>

## max_url_size_exceeds property — request_constraints / 332123312223 / 16

Type: `"number"`. Optional.

Match on the URL Size that exceed this value. Exclusive with \[max\_url\_size\_none\]

Upstream description:

Exclusive with \[max\_url\_size\_none\]

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_url_size_none](resources--service_policy_rule--reference--group-002.md#canonical-2330331322221000-2213212201103011-1120230100211311-3201011320111220-1123011221100112-1002130230011321-0213111023121131-3303130013023331): complete subsection reference.

<a id="canonical-2000133132021033-3101230123303023-1320331113233212-1203012301330101-3233221230010211-2133102233100201-0202312222132203-1331100102133012"></a>

## Next pages — request_constraints / 332123312223 / 17

- [request_constraints.max_cookie_count_none](resources--service_policy_rule--reference--group-002.md#canonical-1122022011000132-2310232120103311-1302322200300230-2231133020012123-0101113203221130-3103331101010130-0200032210220122-2333200330022310)
- [request_constraints.max_cookie_key_size_none](resources--service_policy_rule--reference--group-002.md#canonical-3210311330313210-1030032003013232-2310012332113123-0200213031011110-3032031210102331-2131213220101131-3321201211111213-3111123300123122)
- [request_constraints.max_cookie_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-2210332330211102-3323302021220120-0333020011333130-3231233221330020-0311010101023031-1213120210302130-3032020321123303-3220013033000320)
- [request_constraints.max_header_count_none](resources--service_policy_rule--reference--group-002.md#canonical-2031333020012231-3100132320333220-1321002111211212-1121322112320113-1033112301122111-3221032333001212-3332000110133331-1031013122313000)
- [request_constraints.max_header_key_size_none](resources--service_policy_rule--reference--group-002.md#canonical-1132101021313102-3332033311133003-2210133001000322-0010232331203032-2322133101011321-1202032320332312-3132033012032313-3332110002331110)
- [request_constraints.max_header_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-2102120230210231-2100101103110120-1123203011201001-3031333211322330-1210231111011330-2310102033001102-1320213011013311-0021120132333103)
- [request_constraints.max_parameter_count_none](resources--service_policy_rule--reference--group-002.md#canonical-3221300301123333-0233121001221021-1311302003000011-2010203103033011-3100131213301022-0011232323122220-0300231113313310-1313310111333232)
- [request_constraints.max_parameter_name_size_none](resources--service_policy_rule--reference--group-002.md#canonical-2230202133131132-2331002310102321-2203332230123312-3120102201131031-2000222113110033-2122102213110001-3312010333210032-0032213011111222)
- [request_constraints.max_parameter_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-1023121102232230-3110132110130012-2003011223012323-0330131103322022-1233311030100300-0113123310232012-3233100301000331-3200002330112200)
- [request_constraints.max_query_size_none](resources--service_policy_rule--reference--group-002.md#canonical-0100103003131210-2220030203003130-1213013113101012-2123313132020002-3211122013110230-0202020232331211-3111300202321122-1301331102130200)
- [request_constraints.max_request_line_size_none](resources--service_policy_rule--reference--group-002.md#canonical-0000013323113031-2033013133320032-2010233210331233-2202002001102333-3222323012130033-2132022123121011-0231320022032213-0133013002000130)
- [request_constraints.max_request_size_none](resources--service_policy_rule--reference--group-002.md#canonical-0211021211211232-0231130103220023-2211300110200302-1330123120002032-1233331131120002-0331131010010032-1001003322020101-2232031202032200)
- [request_constraints.max_url_size_none](resources--service_policy_rule--reference--group-002.md#canonical-2330331322221000-2213212201103011-1120230100211311-3201011320111220-1123011221100112-1002130230011321-0213111023121131-3303130013023331)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-1122022011000132-2310232120103311-1302322200300230-2231133020012123-0101113203221130-3103331101010130-0200032210220122-2333200330022310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212100113020003-3201310212330103-0013010223010223-2120311203300300-2203312322332201-1310033133200331-1323332310010020-3310200233303123"></a>

## request_constraints.max_cookie_count_none — max_cookie_count_none / 332123222201 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_cookie_count_none

<a id="canonical-3330003322002303-3112011103313310-2310133300001032-0301112113002100-2231301330310022-1113310301123111-1211300233030201-0231330331021111"></a>

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

<a id="canonical-3103202003203133-3111322132012031-3221210303122230-1020101012030030-2202130302121321-3210320223310320-2233003021100331-0011213310113021"></a>

## Direct properties — max_cookie_count_none / 332123222201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003301322203102-3020302302013023-3133010102303132-2121131211131132-0212032022302303-2100223110132132-2213111200112203-0023300312011012"></a>

## Next pages — max_cookie_count_none / 332123222201 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-3210311330313210-1030032003013232-2310012332113123-0200213031011110-3032031210102331-2131213220101131-3321201211111213-3111123300123122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312010232110232-1123013223102220-1312322032121311-0131321122322121-1200233112120030-0321231010022033-1110221023201303-1202302110012200"></a>

## request_constraints.max_cookie_key_size_none — max_cookie_key_size_none / 322321300320 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_cookie_key_size_none

<a id="canonical-0122132010120230-1020232110103112-3221012123023201-1301101202303202-0301310003310113-3210303313321301-0313331323231233-1323000201213320"></a>

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

<a id="canonical-2031012302222310-0110300322030120-1333311311220222-3020123221312331-2313030010031312-3302023322210013-0012003013010311-3223322111322230"></a>

## Direct properties — max_cookie_key_size_none / 322321300320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030131032031130-3103301202000212-1311132212102011-3100213131002102-2013302332112302-1331012230212310-2230001120323011-3123212011010222"></a>

## Next pages — max_cookie_key_size_none / 322321300320 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2210332330211102-3323302021220120-0333020011333130-3231233221330020-0311010101023031-1213120210302130-3032020321123303-3220013033000320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012112221003322-2021303130330013-2100200311032310-1011321000033013-1011212100132201-0322132303003231-3130100023323123-2011020223322231"></a>

## request_constraints.max_cookie_value_size_none — max_cookie_value_size_none / 333112120201 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_cookie_value_size_none

<a id="canonical-0130202311032031-3021022220311213-2202111002300213-3112133122112130-3131100320201132-1230023003103202-2022301131222121-1130033202332102"></a>

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

<a id="canonical-3313100102132002-2210201221010310-2010112313333113-1222322220233121-0333102321031321-1300012301301131-1111031130032210-3201220112001113"></a>

## Direct properties — max_cookie_value_size_none / 333112120201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302103011023203-0111020200331120-0303203011103322-2111310233103103-1311330312112103-3211102301102322-2112112102013022-0310010321201222"></a>

## Next pages — max_cookie_value_size_none / 333112120201 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2031333020012231-3100132320333220-1321002111211212-1121322112320113-1033112301122111-3221032333001212-3332000110133331-1031013122313000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101313032003131-2310112202230103-2211222211233003-1220132032323130-3313301212112021-3101223132101213-2032212221022111-0333211300131000"></a>

## request_constraints.max_header_count_none — max_header_count_none / 211032220123 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_header_count_none

<a id="canonical-0231113303012011-0211220222302012-1200101001132312-3010010221132210-3233332300122222-0322021112300010-0300203121103023-3012210210021313"></a>

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

<a id="canonical-0302332201123201-1232102031010221-0331332020013103-2131202323110103-1022133232003232-0321010131131302-2210331113001132-0203130011232211"></a>

## Direct properties — max_header_count_none / 211032220123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221310210132031-2021320212022233-3032230003021303-2122231302230320-2201231130132330-2333322001133322-0200321001321031-2032310030323200"></a>

## Next pages — max_header_count_none / 211032220123 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-1132101021313102-3332033311133003-2210133001000322-0010232331203032-2322133101011321-1202032320332312-3132033012032313-3332110002331110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330020232232232-3003002231133320-1012002322322122-2311331323331133-1320122002110032-3220131013201310-3332210000232332-0012320132210233"></a>

## request_constraints.max_header_key_size_none — max_header_key_size_none / 322302003230 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_header_key_size_none

<a id="canonical-0203311333010103-1213002103233110-3332002301220100-0103103332020023-0001232020213231-0132101001323123-0021202133022020-3002032121223032"></a>

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

<a id="canonical-3330130201120211-2301322201212211-2231033200323201-2303223210200333-3021203020221111-3222212031023311-3312323101130130-3133032320011112"></a>

## Direct properties — max_header_key_size_none / 322302003230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303001312022203-3302302202211333-3002232230322101-0201211310321223-2122223300023302-0110012231012302-0103230011033223-1300032320222011"></a>

## Next pages — max_header_key_size_none / 322302003230 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2102120230210231-2100101103110120-1123203011201001-3031333211322330-1210231111011330-2310102033001102-1320213011013311-0021120132333103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012312200313223-0203113021002332-1110200311213133-0000020132002000-3013332330123123-1221212120112321-2033031011203132-3312001021322210"></a>

## request_constraints.max_header_value_size_none — max_header_value_size_none / 303202333200 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_header_value_size_none

<a id="canonical-3332010032003122-2002310031220112-0200323132113003-0232210311322332-1102020313301231-0100201222233130-1113113121320301-3123210320132232"></a>

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

<a id="canonical-2230012033200121-1112001333233121-1032103301200022-0123311023110201-2120003330002122-0123313002311333-3203303220221012-2301220110301210"></a>

## Direct properties — max_header_value_size_none / 303202333200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003000220121031-1020103132013121-1022120213031001-0213211302031133-3131001120121021-2021330001032003-2331110201012102-3031100120323111"></a>

## Next pages — max_header_value_size_none / 303202333200 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-3221300301123333-0233121001221021-1311302003000011-2010203103033011-3100131213301022-0011232323122220-0300231113313310-1313310111333232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220010222002012-2301333011021011-0130113113201021-2103112210021320-0223210002022231-2030130231302101-1200033121210031-0000302321021212"></a>

## request_constraints.max_parameter_count_none — max_parameter_count_none / 322133120202 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_parameter_count_none

<a id="canonical-1110333230123311-0312322321312122-1330211131123201-0010320132111230-0111232322111312-0021010103323333-1112221012202133-1303112320333201"></a>

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

<a id="canonical-1030001332010213-2320200010023303-3122210132222330-0111222021013030-1121331223330110-1300323302211320-1101331132313023-3012023232102321"></a>

## Direct properties — max_parameter_count_none / 322133120202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130133122221233-2303300033230113-1013112021200320-1332333130001101-1301323102200230-0212311011133330-1233312023313010-3201013303230333"></a>

## Next pages — max_parameter_count_none / 322133120202 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2230202133131132-2331002310102321-2203332230123312-3120102201131031-2000222113110033-2122102213110001-3312010333210032-0032213011111222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301100111010222-1033130122322312-0003212230121323-0120303301101030-2323001320322031-0330213103233213-2111300221321133-1131012231310001"></a>

## request_constraints.max_parameter_name_size_none — max_parameter_name_size_none / 231222321132 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_parameter_name_size_none

<a id="canonical-1223322223300233-0130113021331330-2323201312220132-1322022321310110-0202232103120301-1012123212220203-1011023221322021-3122033220301211"></a>

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

<a id="canonical-1120000311031002-0220113133303002-0033221010002033-1001333313123330-0210021032013233-3301223300133312-0020103303333013-0113233332333233"></a>

## Direct properties — max_parameter_name_size_none / 231222321132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131131103310231-3110032123313020-0311031210223221-3220320012101112-3012233102120031-3300021302122232-0013012033000310-1030330022222011"></a>

## Next pages — max_parameter_name_size_none / 231222321132 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-1023121102232230-3110132110130012-2003011223012323-0330131103322022-1233311030100300-0113123310232012-3233100301000331-3200002330112200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023000030022021-3222300232313212-2112230212223323-2313311030220133-3113212110033011-3301131312210221-0313020102312332-3321100232313320"></a>

## request_constraints.max_parameter_value_size_none — max_parameter_value_size_none / 133110001110 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_parameter_value_size_none

<a id="canonical-1323313230133220-0222102003010130-2130000311203101-2221202332011303-2303321031302101-3311022303220021-0120321203333320-1221213122130020"></a>

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

<a id="canonical-0332022302132212-1321212312333031-0001030132123002-0010313323221220-2331032031123330-1221302332123023-2202231021010320-1021122011333232"></a>

## Direct properties — max_parameter_value_size_none / 133110001110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102203130101102-3320221101330100-2102033323213233-3030312032112113-2200132313030312-3202321220100123-2130211101333111-1321313023233000"></a>

## Next pages — max_parameter_value_size_none / 133110001110 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-0100103003131210-2220030203003130-1213013113101012-2123313132020002-3211122013110230-0202020232331211-3111300202321122-1301331102130200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033231123300032-1300000332121300-0110010121021131-1313132231010132-0231123010222030-0111102001113313-0312103233230211-2033333103020030"></a>

## request_constraints.max_query_size_none — max_query_size_none / 312300213333 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_query_size_none

<a id="canonical-1203232220211220-0211100200300303-0022133022033311-1200031312211111-2322013231313110-0230232003011333-1100221300223112-3323222232010110"></a>

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

<a id="canonical-0013301011211233-0213111101032303-0121203021302131-3301020203100103-2030222122022323-0012002231201231-1321012113132101-1221101310121122"></a>

## Direct properties — max_query_size_none / 312300213333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203201323230312-3120033322123313-3320101122210223-1200201013130013-3033330230323311-3302131202100321-3332112223001003-0232322130210313"></a>

## Next pages — max_query_size_none / 312300213333 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-0000013323113031-2033013133320032-2010233210331233-2202002001102333-3222323012130033-2132022123121011-0231320022032213-0133013002000130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313132022312003-2121322022230101-0021121031131113-3211011111003330-1212120200300220-3133111012012132-3112011011320203-0333102021113312"></a>

## request_constraints.max_request_line_size_none — max_request_line_size_none / 303022311010 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_request_line_size_none

<a id="canonical-3033011001110323-3122020112332330-0230210321000010-1213233000333111-0210010303213123-2203220130031212-0021101202312223-1002021313203221"></a>

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

<a id="canonical-1320020303030120-3202333000203212-0301223311032022-0210312213102001-2330300101212200-0203320212010102-2320202121020220-2032332323231121"></a>

## Direct properties — max_request_line_size_none / 303022311010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310110332203020-2103131000330121-1130320013231300-2322302322300222-3131202222003320-1121112330101030-1222110102220003-3331332312330212"></a>

## Next pages — max_request_line_size_none / 303022311010 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-0211021211211232-0231130103220023-2211300110200302-1330123120002032-1233331131120002-0331131010010032-1001003322020101-2232031202032200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221212130113102-3223123310022123-2301122101303022-0133210320022020-1211121030121312-2023232302231221-0233032010333110-2203300012011110"></a>

## request_constraints.max_request_size_none — max_request_size_none / 120020321210 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_request_size_none

<a id="canonical-1110021232200100-3021311022230001-3331231112100033-1332110300222123-0222022201021032-2032300302303102-3200320111320133-3030022013230313"></a>

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

<a id="canonical-0321313031130112-0330111232012301-3323021221213122-3333202330000002-1200310201300120-3333231302112313-2132210122333132-3020130223033130"></a>

## Direct properties — max_request_size_none / 120020321210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230132202232101-3001111220313020-1331121222122200-3311121031323031-1303323011222223-3313223332310201-3012321120030231-3131132331111123"></a>

## Next pages — max_request_size_none / 120020321210 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2330331322221000-2213212201103011-1120230100211311-3201011320111220-1123011221100112-1002130230011321-0213111023121131-3303130013023331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220133310013313-2303033122032130-0330023222203220-2210323023313023-0333221012221213-0323002013312332-3031130232000302-1323232202322103"></a>

## request_constraints.max_url_size_none — max_url_size_none / 033230231200 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_url_size_none

<a id="canonical-3212210122123013-0021120011333112-2333023110102022-1033202113011330-1332300213030323-1122132301332232-2333002000233000-2323333110322112"></a>

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

<a id="canonical-0222111311030112-1030201011131311-1002132312121211-3311200112230310-1022021210110223-1003023011231330-0112310320103131-3132322103303022"></a>

## Direct properties — max_url_size_none / 033230231200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332311123013003-2332322002232320-2300232202312301-2011130031032022-2302113100000310-2000203232232213-0133223220000321-2102111201233232"></a>

## Next pages — max_url_size_none / 033230231200 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102212311203021-2320303011122133-2302233322132200-1220230233322203-3321121233032001-2023101331322102-1121121231331030-3213012032111231"></a>

## segment_policy — segment_policy / 033110021103 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- segment_policy

<a id="canonical-0000310300000020-3030312333212300-1033312300110000-2012001210311021-3333123312322130-3100112131120320-3322132313111200-1332233120101131"></a>

Type: `"object"`. single nested block, Optional.

Configure source and destination segment for policy.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2130331302230331-1122131223331231-2000202320010321-3102001011131200-3131001233123331-0033233322032011-2313302022202131-0022120000132301"></a>

## Direct properties — segment_policy / 033110021103 / 3

- [dst_any](resources--service_policy_rule--reference--group-002.md#canonical-2113201130221021-1032210031323120-3302033001311130-3010031030233213-2223213113211003-0220220301230130-0113013133110302-0033310222310123): complete subsection reference.

- [dst_segments](resources--service_policy_rule--reference--group-002.md#canonical-2022332223212133-1222113102221310-0111310321311112-0032220010012112-0311310031133322-1221210033230122-3132203110031011-1312332032002013): complete subsection reference.

- [intra_segment](resources--service_policy_rule--reference--group-002.md#canonical-2323132121321022-3120112231121021-1221330022210220-2023333321111201-0123312221120313-2323021022122033-0103131123103200-2201010312333323): complete subsection reference.

- [src_any](resources--service_policy_rule--reference--group-002.md#canonical-1021033100031322-0210332232032110-0310012210003102-3310330313201232-1013122333100201-1010021111102100-0221022030003001-1011013331111111): complete subsection reference.

- [src_segments](resources--service_policy_rule--reference--group-002.md#canonical-0101332320113122-2200203113113133-1130203002313322-0202221012003110-2321113003133011-3130122110120132-1322213110010211-0010122011230310): complete subsection reference.

<a id="canonical-0323203211030222-2300211312120131-0233031123013303-0323300013121233-1023220123013233-0210021213011102-0012113031113132-2132130122323201"></a>

## Next pages — segment_policy / 033110021103 / 4

- [segment_policy.dst_any](resources--service_policy_rule--reference--group-002.md#canonical-2113201130221021-1032210031323120-3302033001311130-3010031030233213-2223213113211003-0220220301230130-0113013133110302-0033310222310123)
- [segment_policy.dst_segments](resources--service_policy_rule--reference--group-002.md#canonical-2022332223212133-1222113102221310-0111310321311112-0032220010012112-0311310031133322-1221210033230122-3132203110031011-1312332032002013)
- [segment_policy.intra_segment](resources--service_policy_rule--reference--group-002.md#canonical-2323132121321022-3120112231121021-1221330022210220-2023333321111201-0123312221120313-2323021022122033-0103131123103200-2201010312333323)
- [segment_policy.src_any](resources--service_policy_rule--reference--group-002.md#canonical-1021033100031322-0210332232032110-0310012210003102-3310330313201232-1013122333100201-1010021111102100-0221022030003001-1011013331111111)
- [segment_policy.src_segments](resources--service_policy_rule--reference--group-002.md#canonical-0101332320113122-2200203113113133-1130203002313322-0202221012003110-2321113003133011-3130122110120132-1322213110010211-0010122011230310)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2113201130221021-1032210031323120-3302033001311130-3010031030233213-2223213113211003-0220220301230130-0113013133110302-0033310222310123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002000023301031-2123023201230301-3021021101013203-0331311010123122-0101231311120003-1001101030313113-2222103132121013-1320312033300313"></a>

## segment_policy.dst_any — dst_any / 122200011011 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- segment_policy.dst_any

<a id="canonical-2332321232122233-1303323222111121-1220010021332231-2031211320212330-2001012101311130-0310102211102210-0322223312220030-2221230311322313"></a>

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

<a id="canonical-0001203313120103-1111320122321212-2013032312311211-1332030001030321-0010221331331002-2022103022301232-1211310011003121-1023000301100310"></a>

## Direct properties — dst_any / 122200011011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003010022022233-1011300331203021-0211330020332132-3120212100031232-1330221022203331-1110021010010011-1231220022330300-3230103102230102"></a>

## Next pages — dst_any / 122200011011 / 4

- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2022332223212133-1222113102221310-0111310321311112-0032220010012112-0311310031133322-1221210033230122-3132203110031011-1312332032002013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111301333333310-2323130002200130-1002122133222231-2013033112233200-2303322123200022-0021321322021301-2323322333330123-1102311303233213"></a>

## segment_policy.dst_segments — dst_segments / 322100110113 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- segment_policy.dst_segments

<a id="canonical-3113011300132303-2321300122002120-2020323213211011-1031022001212230-2122011332330022-3111201320100133-3002320222223022-0201022330113323"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dst segments.

Upstream description:

List of references to Segments.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0300021220120201-3331001122133230-3323230031323321-2133233202113120-0201323020312302-1201331311330312-2220033211102101-0011322032031323"></a>

## Direct properties — dst_segments / 322100110113 / 3

- [segments](resources--service_policy_rule--reference--group-002.md#canonical-0102102011212003-2000031103020230-2021320331231313-1313122001333300-3002001311202123-3202320310031110-1202101220012321-0033220113231320): complete subsection reference.

<a id="canonical-2311213333030103-2233121131030100-2100011223302011-3023112132101002-1221311220221033-1022222310301230-3030231313231321-3032102310032012"></a>

## Next pages — dst_segments / 322100110113 / 4

- [segment_policy.dst_segments.segments](resources--service_policy_rule--reference--group-002.md#canonical-0102102011212003-2000031103020230-2021320331231313-1313122001333300-3002001311202123-3202320310031110-1202101220012321-0033220113231320)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-0102102011212003-2000031103020230-2021320331231313-1313122001333300-3002001311202123-3202320310031110-1202101220012321-0033220113231320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130003331032303-0310232303023220-3021111131120022-0312202300222330-1013100221233112-2023231232030331-2323213212313003-1020322330331331"></a>

## segment_policy.dst_segments.segments — segments / 021022211022 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- [segment_policy.dst_segments](resources--service_policy_rule--reference--group-002.md#canonical-2022332223212133-1222113102221310-0111310321311112-0032220010012112-0311310031133322-1221210033230122-3132203110031011-1312332032002013)
- segment_policy.dst_segments.segments

<a id="canonical-2202013210133013-0011231011230330-3111001023132023-1000201001113100-1231320031223030-1101010200302311-2021021111110302-0130312101001000"></a>

Type: `"object"`. list nested block, Optional.

Segments. Select list of segments.

Upstream description:

Select list of segments.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3103321030302111-0231100110321332-0201333332211212-2223333133201200-0110031323132033-2331102103000112-3323202020300110-3320013230112110"></a>

## Direct properties — segments / 021022211022 / 3

<a id="canonical-0031001033010112-3333100222122113-2333202111213332-1332030103113201-1113013303003120-0300211113000120-3223302103102220-0313301033231133"></a>

<a id="canonical-3313120030101312-1330012330113231-2112333301223103-0120030001101310-1210100313103122-1322302301000002-2303002302313330-2303310223000132"></a>

## name property — segments / 021022211022 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0213322130102120-3323021203320203-1220022200203100-1011323023311213-0332212223030110-0131231122113313-0110220312132221-3301321231031301"></a>

<a id="canonical-1121020221023203-2203133130333113-2021311323110313-3123112131310010-3131233333113103-0202112310212002-3133223000313020-0320022223222013"></a>

## namespace property — segments / 021022211022 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2221200313302300-2313130133003030-0001223103030203-0101131201221012-0011032323103103-1132112021123020-2231333133322031-2112333223333000"></a>

<a id="canonical-2230313121233203-2330030203220132-1312200130030010-0130111010133003-0312100002332223-0231030030310022-0003332112021211-2231121003031031"></a>

## tenant property — segments / 021022211022 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3121033223222321-1313320011020131-2032133110302333-2320023133123020-3201101200313220-3021331003303323-1130122220221202-0313221330203000"></a>

## Next pages — segments / 021022211022 / 7

- [segment_policy.dst_segments](resources--service_policy_rule--reference--group-002.md#canonical-2022332223212133-1222113102221310-0111310321311112-0032220010012112-0311310031133322-1221210033230122-3132203110031011-1312332032002013)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2323132121321022-3120112231121021-1221330022210220-2023333321111201-0123312221120313-2323021022122033-0103131123103200-2201010312333323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230231000321111-2033120331202223-2110301031133112-0033133333120100-3122012121232200-0320031203210313-3321112112123323-0220010103313033"></a>

## segment_policy.intra_segment — intra_segment / 302120322202 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- segment_policy.intra_segment

<a id="canonical-0230132313131123-0123032033333211-1203232110211110-3133123130013103-0333333132312203-3313310313311102-0133021021001210-3220000033031211"></a>

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

<a id="canonical-2130133123021313-0022211230211003-0133112130012321-0003102230231111-3200021111032302-3100212000020100-1030132203113031-1333201233133222"></a>

## Direct properties — intra_segment / 302120322202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231321033310121-2123001110230220-0103333330300303-3222010013210333-2120000011300031-1112200230021203-3203230003121132-1111010123112033"></a>

## Next pages — intra_segment / 302120322202 / 4

- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-1021033100031322-0210332232032110-0310012210003102-3310330313201232-1013122333100201-1010021111102100-0221022030003001-1011013331111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020322312322223-1030320021100332-3330031332000133-3013120211223233-3330320100022223-3022211323332000-0003212133313322-3202233330111302"></a>

## segment_policy.src_any — src_any / 123011120201 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- segment_policy.src_any

<a id="canonical-2133300232030230-3302313230113313-2230112133321120-1332211003133132-3132203311230011-0321232132203331-2303133222030121-3223131223223100"></a>

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

<a id="canonical-1103222333223022-1322213020233203-1000203103032232-3303011312001330-2303233220320331-1110021000213011-0332222200012010-2032122033333003"></a>

## Direct properties — src_any / 123011120201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130223112201333-0330002003121011-1302201033203111-3213303221123202-0311012312133003-1230301331032131-0210122120312312-0203213211221203"></a>

## Next pages — src_any / 123011120201 / 4

- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-0101332320113122-2200203113113133-1130203002313322-0202221012003110-2321113003133011-3130122110120132-1322213110010211-0010122011230310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120210303010322-2212022123012031-1022311020030101-2003023320312012-0111333230132013-1231331203221323-0123302220221000-1223220221011233"></a>

## segment_policy.src_segments — src_segments / 001213300003 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- segment_policy.src_segments

<a id="canonical-1010311123002311-3011223222303031-0010001002300303-1012200110130003-3312203003320212-0132220003222002-3010330210130202-2023310332233331"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for src segments.

Upstream description:

List of references to Segments.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2312102111010201-2201002331131102-0303031003213101-3120130012312102-2320212030211231-2211001331112011-0131233010003131-1102000323303323"></a>

## Direct properties — src_segments / 001213300003 / 3

- [segments](resources--service_policy_rule--reference--group-002.md#canonical-0222002030202103-3002001110223120-1111302113013331-2122323332211302-0131202200010210-1103103302211023-3232001301123333-3132310231302323): complete subsection reference.

<a id="canonical-2010220233303123-3331110113300210-1032110013301202-2223321021321231-3120303330302111-1210110122013202-2301223321333222-0201122311122212"></a>

## Next pages — src_segments / 001213300003 / 4

- [segment_policy.src_segments.segments](resources--service_policy_rule--reference--group-002.md#canonical-0222002030202103-3002001110223120-1111302113013331-2122323332211302-0131202200010210-1103103302211023-3232001301123333-3132310231302323)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-0222002030202103-3002001110223120-1111302113013331-2122323332211302-0131202200010210-1103103302211023-3232001301123333-3132310231302323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311121331223301-1303232210013222-3122221100311312-1301120133031230-0030333001103111-3333210313102222-3013233021012111-3232110122223233"></a>

## segment_policy.src_segments.segments — segments / 020121021012 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- [segment_policy.src_segments](resources--service_policy_rule--reference--group-002.md#canonical-0101332320113122-2200203113113133-1130203002313322-0202221012003110-2321113003133011-3130122110120132-1322213110010211-0010122011230310)
- segment_policy.src_segments.segments

<a id="canonical-3033220131033301-0221220332120201-0230121221221110-1300033010331320-3231120233003320-1233310312232321-1213032210033323-0332010221113213"></a>

Type: `"object"`. list nested block, Optional.

Segments. Select list of segments.

Upstream description:

Select list of segments.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0312013220031000-2311122001020212-1112322202110121-3300310320110102-3231323203233213-0213333001210332-2100033322130113-3132303020232113"></a>

## Direct properties — segments / 020121021012 / 3

<a id="canonical-1021300030021222-1320103122013202-2022123003021031-0220021310110031-1312020031222001-0013203312301302-0332031320013132-0031333323032213"></a>

<a id="canonical-0133220133230331-2200003230212211-1022032311210110-3031003230232122-0102321131303212-3301103020122101-0232133320221200-3300121100211212"></a>

## name property — segments / 020121021012 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2313332313212032-1010011002013002-0003111333002113-0330032321023103-1321120322322300-0033020202201312-2331230000232031-1203010320231331"></a>

<a id="canonical-3031122303303202-1330222100113000-2011230320003012-1321320102300223-1120332113031123-2310032232313310-2120212313123130-3332333302220112"></a>

## namespace property — segments / 020121021012 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0301231033320131-0303133320301200-0021113210022300-0303123011102101-3022032301321332-1221021321211323-1330103010102103-1111212220302321"></a>

<a id="canonical-1210210202131012-0330222211101330-3111100013103111-3103032130120012-3201333302331332-0131003223312200-3233230003001232-0112230001031011"></a>

## tenant property — segments / 020121021012 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3301330131320312-2100133022131330-1230111121302230-1323013333132212-1031031202021203-0312321323213200-1030133012100103-3211322212312221"></a>

## Next pages — segments / 020121021012 / 7

- [segment_policy.src_segments](resources--service_policy_rule--reference--group-002.md#canonical-0101332320113122-2200203113113133-1130203002313322-0202221012003110-2321113003133011-3130122110120132-1322213110010211-0010122011230310)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-1021310131303231-3111121233212022-1000112321100121-1023301303021202-3313011013101213-2203330203232131-3231103133122222-2233231203321311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210022313331330-3322323331201030-1303303103012300-2233022210312011-0330302132321011-2332331133030311-0300333030031013-1101003111211021"></a>

## timeouts — timeouts / 133201100210 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- timeouts

<a id="canonical-3220311022103123-2030031323233001-0321211311322230-1333222313202203-0233320303323213-3133330130021203-3233123132133020-0320211203311000"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323312321110010-1010022101001021-0120232123030010-0322100102202331-0013313110213320-1100011031113313-0130002101101131-1010030001301303"></a>

## Direct properties — timeouts / 133201100210 / 3

<a id="canonical-0222333130210210-2122032222313301-0321310123301310-3231302222101003-1030221121110102-2131020211020112-3112300131113212-2030203130111300"></a>

<a id="canonical-1102302231122322-3331033133301031-1230313200031220-2221033200333103-1211130003032233-2131303331210203-0210021131131313-3021310033333103"></a>

## create property — timeouts / 133201100210 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3310021130020030-3212212120032333-1100120210100100-3313113223313122-0222013210030010-0211003202300320-1211010013102103-1222023100022233"></a>

<a id="canonical-2010122231022000-3010103101033213-0032310030133001-2332323111010221-1202000301012100-1302001111223212-1131010110322003-1131321301300313"></a>

## delete property — timeouts / 133201100210 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3013131233001012-2230120010033112-3331333322011332-3031030002030231-1222232132220233-0023310310001220-0212211202200112-2023303010113101"></a>

<a id="canonical-2300323020112212-3131102230033320-3313033213332221-1213310000303201-3120223101333302-0121121021321300-2112013120213101-3210320323013303"></a>

## read property — timeouts / 133201100210 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1120022113001331-2130320113031203-2010200322130313-2302311201332123-2312301110123333-1123322112031032-1122000320011102-2330131200103321"></a>

<a id="canonical-1201332320010210-1122311120103323-2100302133211000-2112120022221031-1212131001311031-1302331030020001-1002231032121012-2222300321323200"></a>

## update property — timeouts / 133201100210 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2132120222303301-3303021102103123-3101003210020210-0210023201101130-1221202233131222-2002032020112103-3001210130113313-1011211030121103"></a>

## Next pages — timeouts / 133201100210 / 8

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2200000321123222-1121231032320123-3112320132222010-3101030202331210-2103131131000030-1333130312003211-1330202233211000-0302103210011223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013200113333233-1232302213301102-2300303020121311-2202020222010210-2321120333013131-2131120223222012-2303131103322120-0320133203020112"></a>

## tls_fingerprint_matcher — tls_fingerprint_matcher / 323223132332 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- tls_fingerprint_matcher

<a id="canonical-2300023332310032-3012332003021312-3311001110322002-3132320121320002-3212132123303111-2220301013203202-0001020200302013-3212122333321223"></a>

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

<a id="canonical-3310200013222202-0003221022223113-1101021230003111-1032333303023333-2100300300023013-3100322201120100-2331323311330201-0031030022200310"></a>

## Direct properties — tls_fingerprint_matcher / 323223132332 / 3

<a id="canonical-2213113001321332-2230010132132201-3131200121012322-3110112223120201-1022203210311023-1112312302231023-1032132201320101-1032321103021011"></a>

<a id="canonical-0031311102201111-0030331032133003-2231111121202310-2002020312213110-2033102312100013-0233333313230030-3003323103233233-0221220313120031"></a>

## classes property — tls_fingerprint_matcher / 323223132332 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0311230130200231-1033100312020100-1322132333033120-0201003033310310-1023101322130221-0331223123301332-1332011031310202-1311100322001031"></a>

<a id="canonical-3220202310320013-0122333203101032-1020121332201230-1231102323300020-2032101331101312-1331211103233100-0210310021111233-0330320231132322"></a>

## exact_values property — tls_fingerprint_matcher / 323223132332 / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3120011321131203-0012123131322323-1101203311102121-3213221221210010-0302310333001031-1102100220102211-0000021211231302-1033300200332323"></a>

<a id="canonical-2001112113132223-2331030231111031-1332313200312031-2132211232002102-1010102231131011-2331002332200031-2331112033320101-0232023132002111"></a>

## excluded_values property — tls_fingerprint_matcher / 323223132332 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3220202110331221-2020033333332123-0210002231330211-3202013122113232-3311321210010211-2232203313020022-1110212321210003-2313132031002002"></a>

## Next pages — tls_fingerprint_matcher / 323223132332 / 7

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203011233200203-1310000123233023-1322133103312130-0011013101200312-2321121332112013-3113223200000003-2311102213211333-0230221212312300"></a>

## waf_action — waf_action / 003002323113 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- waf_action

<a id="canonical-0112112222221031-1331310203113201-2000213313323030-0130211021321221-2010210010030031-3223302012121302-1312233021121013-0000331203202102"></a>

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

<a id="canonical-1230333003131301-2121330313020212-0201303113302130-0102313121011223-2032332211333100-1131200021113210-0130003132213103-0301332011220231"></a>

## Direct properties — waf_action / 003002323113 / 3

- [app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100): complete subsection reference.

- [none](resources--service_policy_rule--reference--group-002.md#canonical-0122121012021132-1000022311102111-3131301323221132-0220303102210311-3223101332203121-2201303111022201-2320030331311003-3211100200333212): complete subsection reference.

- [waf_skip_processing](resources--service_policy_rule--reference--group-002.md#canonical-2300310201333330-3133013011222002-2102212230202020-3113323220303320-0031002131200313-1213322102302103-0221203123001033-2303330302222312): complete subsection reference.

<a id="canonical-1302110111323101-2320203122201023-0030321322300232-1200210001023232-2113132103232031-0230001011122012-2111221330102113-1111301310220032"></a>

## Next pages — waf_action / 003002323113 / 4

- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100)
- [waf_action.none](resources--service_policy_rule--reference--group-002.md#canonical-0122121012021132-1000022311102111-3131301323221132-0220303102210311-3223101332203121-2201303111022201-2320030331311003-3211100200333212)
- [waf_action.waf_skip_processing](resources--service_policy_rule--reference--group-002.md#canonical-2300310201333330-3133013011222002-2102212230202020-3113323220303320-0031002131200313-1213322102302103-0221203123001033-2303330302222312)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223330200331303-0332201102300301-3322222123303010-3233030313032211-2302113220301111-1111213012211323-0300001213032321-3133313033222020"></a>

## waf_action.app_firewall_detection_control — app_firewall_detection_control / 233332211211 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- waf_action.app_firewall_detection_control

<a id="canonical-2302020223330013-0320000111031132-0032321313011101-1220221302020001-1023310022311120-3220322131302302-2112320233110232-1201003221033023"></a>

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

<a id="canonical-2030331300202001-3313230300220221-1011301203020033-1223223332120003-2103320230310021-0010033011321223-0120232301212022-0002310220033211"></a>

## Direct properties — app_firewall_detection_control / 233332211211 / 3

- [exclude_attack_type_contexts](resources--service_policy_rule--reference--group-002.md#canonical-0221033123313311-2131133301001003-2221110233020201-2110000023011300-0302133300313312-2320022333022111-1133323323131212-3132220121231210): complete subsection reference.

- [exclude_bot_name_contexts](resources--service_policy_rule--reference--group-002.md#canonical-0321130232020121-0010023111030323-3023100322011320-3222001110203201-1212221212323003-0013202010231232-0131203313132331-2023221130323302): complete subsection reference.

- [exclude_signature_contexts](resources--service_policy_rule--reference--group-002.md#canonical-2023301201222231-0311313312020102-3232203322212121-1323122133301123-3321220330323000-3202033033333020-2121210022301213-3021102030120201): complete subsection reference.

- [exclude_violation_contexts](resources--service_policy_rule--reference--group-002.md#canonical-2323023101030330-0311033202012133-3120010032230210-2311100301223311-0221003311113000-2301222003331101-0133202202221320-0222121311320033): complete subsection reference.

<a id="canonical-0011132110310012-0113021131013121-2231230221303132-3030221112123100-2233302033123322-3321132301003322-3220032131321131-2013131011210122"></a>

## Next pages — app_firewall_detection_control / 233332211211 / 4

- [waf_action.app_firewall_detection_control.exclude_attack_type_contexts](resources--service_policy_rule--reference--group-002.md#canonical-0221033123313311-2131133301001003-2221110233020201-2110000023011300-0302133300313312-2320022333022111-1133323323131212-3132220121231210)
- [waf_action.app_firewall_detection_control.exclude_bot_name_contexts](resources--service_policy_rule--reference--group-002.md#canonical-0321130232020121-0010023111030323-3023100322011320-3222001110203201-1212221212323003-0013202010231232-0131203313132331-2023221130323302)
- [waf_action.app_firewall_detection_control.exclude_signature_contexts](resources--service_policy_rule--reference--group-002.md#canonical-2023301201222231-0311313312020102-3232203322212121-1323122133301123-3321220330323000-3202033033333020-2121210022301213-3021102030120201)
- [waf_action.app_firewall_detection_control.exclude_violation_contexts](resources--service_policy_rule--reference--group-002.md#canonical-2323023101030330-0311033202012133-3120010032230210-2311100301223311-0221003311113000-2301222003331101-0133202202221320-0222121311320033)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-0221033123313311-2131133301001003-2221110233020201-2110000023011300-0302133300313312-2320022333022111-1133323323131212-3132220121231210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003002230330111-2012220333233130-0100211032221223-2030033233131201-3301320212012001-3030332333122223-1302110110023122-3231332233003303"></a>

## waf_action.app_firewall_detection_control.exclude_attack_type_contexts — exclude_attack_type_contexts / 201301320330 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100)
- waf_action.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-0001100132222033-3233203111303132-1302101100302133-0001120212102001-0013332011013322-3332101122303200-2221332012000110-1332301022301122"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0232003221121303-2203223132322311-1212200232132020-2022133222002001-2311313122131111-1030332101331113-1312031233113211-3112020221303230"></a>

## Direct properties — exclude_attack_type_contexts / 201301320330 / 3

<a id="canonical-1112112313312322-3303313323212300-1000023203011031-2113002332231213-3101113000020001-2033311222112200-1300131112311013-3331021112301001"></a>

<a id="canonical-1002313022100210-0010020102002021-2301122200330102-0333101132203002-0113213222311303-0132131331300131-2003332030320021-3031301331312203"></a>

## context property — exclude_attack_type_contexts / 201301320330 / 4

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

<a id="canonical-1002102131033021-1002323121230322-3221012131011021-2310003221323200-2122221201033021-0021231103230130-0330311130120232-2110300102211033"></a>

<a id="canonical-3133010333200120-2323030113310102-2110222211022303-0230020102133020-1030232001211211-2212302321330011-1111201123330222-0002232313102301"></a>

## context_name property — exclude_attack_type_contexts / 201301320330 / 5

Type: `"string"`. Optional.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1013333333330111-3301332333313001-0332010122001102-2333102303321222-2120201302231111-1302302312201221-2232022023200030-1001113202120131"></a>

<a id="canonical-3031321211223312-3123203002010103-0033303010002001-2100000332331103-1323020003011022-3020200310131332-3320332031131201-2201101201303232"></a>

## exclude_attack_type property — exclude_attack_type_contexts / 201301320330 / 6

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

<a id="canonical-0002122232322331-0023301202132232-3311002213330011-2211323332013100-3022311211300321-1111222122330203-1330332023020310-2201310310313023"></a>

## Next pages — exclude_attack_type_contexts / 201301320330 / 7

- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-0321130232020121-0010023111030323-3023100322011320-3222001110203201-1212221212323003-0013202010231232-0131203313132331-2023221130323302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233322021121021-1020101030332310-1202103223100132-2131012001322332-1232122023333000-0210103031110231-0331231210011003-0222301201212322"></a>

## waf_action.app_firewall_detection_control.exclude_bot_name_contexts — exclude_bot_name_contexts / 031001131102 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100)
- waf_action.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-2231121312223131-1320132030312333-2213010200222123-3220300221010031-1332021211103133-3322232020002310-1330100201031010-3023130323130302"></a>

Type: `"object"`. list nested block, Optional.

Bot Names to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3102333131221231-2113102332021333-0010000323210313-2132100121200122-1200101003331300-2203331330111322-3100231103313012-3013023123223102"></a>

## Direct properties — exclude_bot_name_contexts / 031001131102 / 3

<a id="canonical-1321301102011012-0310203102012333-2113223321021332-3113032022330013-3211001120101012-2333013030133200-1330031210330003-1010332331213121"></a>

<a id="canonical-1222123110012233-3211013100033323-3131131212103320-0022310132120032-2221310333230020-0232310232222001-1302012323003300-1101211121102301"></a>

## bot_name property — exclude_bot_name_contexts / 031001131102 / 4

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1101113311100032-3130321130322313-2213223202322112-2220121231322000-0110221230313112-2012332102301020-2032232012023100-1012013312330002"></a>

## Next pages — exclude_bot_name_contexts / 031001131102 / 5

- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2023301201222231-0311313312020102-3232203322212121-1323122133301123-3321220330323000-3202033033333020-2121210022301213-3021102030120201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023232013102103-0310202330011213-1112122123033131-3200101330233312-3221032000032222-0113113200222103-2223021033100013-2120312122132222"></a>

## waf_action.app_firewall_detection_control.exclude_signature_contexts — exclude_signature_contexts / 201332032331 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100)
- waf_action.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-2002232121300211-3232321313232120-2123203121030122-1110010100222330-3220123313310231-1203132200202311-2202013020330233-0330220012112210"></a>

Type: `"object"`. list nested block, Optional.

Signature IDs to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1221032131303311-1232313133202112-1201000313332332-3123210213011311-1211130103102320-2220100332212103-0131002033002020-2330123302022220"></a>

## Direct properties — exclude_signature_contexts / 201332032331 / 3

<a id="canonical-0302331300021002-2202300131013122-3312102321333212-1331330313212033-0203031212000012-3112230120120031-3303100213201232-1000032133303030"></a>

<a id="canonical-3320210010102320-0303012330033322-1022122023020100-3122121022123100-0112022211330031-2112320122110331-3211021200331313-3200223122121003"></a>

## context property — exclude_signature_contexts / 201332032331 / 4

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

<a id="canonical-0211210033323210-1303120212112233-0220101030201301-1001122232200022-0330130023310100-2323332112012113-3021323233203322-3101001013023222"></a>

<a id="canonical-0233230002233111-1110020102130022-0011122320332000-3232123112322131-2320201332323133-3130111322122230-0021122213300122-0332310212031100"></a>

## context_name property — exclude_signature_contexts / 201332032331 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1303233330233001-2100200201011001-3133220013222203-0131100220332012-3113211013321200-2031322310321033-3212003110011311-0113123121223023"></a>

<a id="canonical-1232100021102011-3212111133003203-2031120021111323-0222131201100212-1002010203110132-3323200312322321-0220213011030123-3332323200330231"></a>

## signature_id property — exclude_signature_contexts / 201332032331 / 6

Type: `"number"`. Optional.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1202231233003202-0212010000010102-2020220311202302-3210211132302321-1120113311200010-3022330121210100-3013032322003212-3102320212001112"></a>

## Next pages — exclude_signature_contexts / 201332032331 / 7

- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2323023101030330-0311033202012133-3120010032230210-2311100301223311-0221003311113000-2301222003331101-0133202202221320-0222121311320033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130131102022212-3223010121321022-2101123101233002-2012000001220333-3332132120320220-2132132233033301-2220311313123310-3310313010202303"></a>

## waf_action.app_firewall_detection_control.exclude_violation_contexts — exclude_violation_contexts / 332030012000 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100)
- waf_action.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-2002213110202130-2303030331103001-2313312323000001-1321231022313032-2011322002310132-2110121132211122-2230111031233322-3100013033001321"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1322232212201130-1320022300032032-3013000031222223-2122200312201203-2230011131120330-2122123120131121-1101132112110100-0301202233000000"></a>

## Direct properties — exclude_violation_contexts / 332030012000 / 3

<a id="canonical-2213132312321330-2111101320303132-2310111300031003-3311110023111200-2013133303030133-0000211013301221-1113132131211113-0121201011120111"></a>

<a id="canonical-3233302200232200-0213032023132331-0323130321322201-2021300021230210-0033301132320312-3122033320331022-1012132311110021-3111012021323113"></a>

## context property — exclude_violation_contexts / 332030012000 / 4

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

<a id="canonical-0312113122213230-1031201020012332-2011002031322133-2322321323001012-3331200211033320-0011331021302131-0233301011003231-3213212101100231"></a>

<a id="canonical-1300033132131320-2102022312312200-1323312331323103-2322110022132230-0011002333323112-3002322313033333-0203231120113333-3230220110303311"></a>

## context_name property — exclude_violation_contexts / 332030012000 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0333110202323322-1120213031231030-3332332221131011-0302102232020300-2012033000201110-1130100133333112-1210023010302103-3201000102230023"></a>

<a id="canonical-3013033321123202-2112210333222000-2321030222211120-3323033031312130-3333001111001023-0000301010211113-1310303212221231-3333012010002231"></a>

## exclude_violation property — exclude_violation_contexts / 332030012000 / 6

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

<a id="canonical-0303212102113311-0110310203231003-2032300001103123-1123312112230030-2001311030112103-1331120233231021-0201212231102021-1032131331303131"></a>

## Next pages — exclude_violation_contexts / 332030012000 / 7

- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-0122121012021132-1000022311102111-3131301323221132-0220303102210311-3223101332203121-2201303111022201-2320030331311003-3211100200333212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313021212000103-0321213212331011-2313113200011302-0320231313130121-2012111132010002-3013100112121103-2212022323210210-3302123200110023"></a>

## waf_action.none — none / 230301211330 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- waf_action.none

<a id="canonical-0203313012102132-1232220131100013-2102313200231111-2202023011200331-0321000311321331-2112212300301323-2303313202121100-0203313013320102"></a>

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

<a id="canonical-0120202211310231-2113223233023310-1020231110121230-2233323323223333-1332230301113131-2222223232023132-1331002010021303-1333110012210203"></a>

## Direct properties — none / 230301211330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030303332013302-3012212313302122-2230111011013122-2003302233213312-0232220101011301-0202312110300220-3130211220223302-0301022133233230"></a>

## Next pages — none / 230301211330 / 4

- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)

<a id="canonical-2300310201333330-3133013011222002-2102212230202020-3113323220303320-0031002131200313-1213322102302103-0221203123001033-2303330302222312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100323131200012-1013122113233310-2120132220123022-3130203302031221-0010312033321300-0232001233203003-3002010022033322-3203031132100013"></a>

## waf_action.waf_skip_processing — waf_skip_processing / 203213300003 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- waf_action.waf_skip_processing

<a id="canonical-3313002212212011-1012223132001111-2033223121201032-0333003111331031-1323210230303010-3113320221320021-2300200010122121-0103203202012311"></a>

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

<a id="canonical-2112103223112112-1233121212011222-1000021001323231-3210103220203302-3112231121222022-2133012322321133-0333003323233010-1230321021230233"></a>

## Direct properties — waf_skip_processing / 203213300003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000301032020131-3310303131102330-3122101231232203-3112323202022130-0303020310310122-3123130230132112-3332200121200233-0032031322203300"></a>

## Next pages — waf_skip_processing / 203213300003 / 4

- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
