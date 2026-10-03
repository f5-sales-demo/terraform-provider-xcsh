---
page_title: "xcsh_service_policy_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule reference."
---

# xcsh_service_policy_rule reference

<a id="canonical-3101203321210221-0110210200130002-0133300232120301-1132333333302213-0023222223113112-3122300010121033-3120102131012231-2221232213013212"></a>

## jwt_claims.item — item / 201032231311 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-3233303020320223-0201322230301223-1111013223211322-2311021222122102-0313111032112232-3202102031123001-0123113023331220-1103011102320312)
- jwt_claims.item

<a id="canonical-1211020202230012-3220003301020233-3203312001201133-3001032323312311-1110031110122001-3233023031130232-2231133233133311-3102023302121230"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0301002301222012-3230101302310330-0012030321013132-2112333231313221-2210203131311320-2331130322201020-1301010013101223-1120022312030333"></a>

## Direct properties — item / 201032231311 / 3

<a id="canonical-2133301320131231-0020011102112213-2011122110332202-2333103013020020-3301311300322012-3103032201213131-1233011223002230-2122020100001123"></a>

<a id="canonical-0002020211331233-2223330323101223-0322302330121032-0222030303233302-1332312022303031-2103212023003222-3332023010023123-1211313011001203"></a>

## exact_values property — item / 201032231311 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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

<a id="canonical-1332033123211330-0022220030300003-1121030333122222-0132221332332020-2102032112320103-3303031310013210-3323030010203022-3133103202032312"></a>

<a id="canonical-2103101010233002-3222330130220011-2233322113130201-0013031103113223-3010312121120133-1220323320003332-0202002311121303-0302002022213302"></a>

## regex_values property — item / 201032231311 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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

<a id="canonical-1101001201330102-0200132301102310-0123011303130023-2212101210003122-2321220202121011-1001031322130232-3123301302321232-0303313012123023"></a>

<a id="canonical-2120023131102302-0323333001101003-2221020221102010-2122303330020110-2223020222003102-0320001222103101-3012211131132102-0001022310001203"></a>

## transformers property — item / 201032231311 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

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

<a id="canonical-0220321002002320-0210302331302301-3232232103032200-1322201233300000-0322310000323001-0003130022302303-0032201033311302-0212331202021313"></a>

## Next pages — item / 201032231311 / 7

- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-3233303020320223-0201322230301223-1111013223211322-2311021222122102-0313111032112232-3202102031123001-0123113023331220-1103011102320312)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3210311000320320-2030222121122100-0220110100321312-1121102131023230-3300133221311211-2200300323311103-1132001130323212-2010010230320023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201100101223030-1222210022200130-3110103013323310-3231302113001013-0000120020213213-2110323322031023-2122323031023133-3132033321112122"></a>

## label_matcher — label_matcher / 013002023003 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- label_matcher

<a id="canonical-0200330310302023-3300303133133011-2123020230313101-1331223303030333-1303021022210130-2310222013032010-2320211001011312-1130301211201100"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0222203310013100-2032302202011301-2300012333131122-1233021111021320-0110331032013220-1102001213010113-2322102211312332-1131302303311110"></a>

## Direct properties — label_matcher / 013002023003 / 3

<a id="canonical-1001130000233222-3122310003013300-1303221203220123-1323100333203231-2300311003330313-3220301303212000-0022220212232010-3022000123332132"></a>

<a id="canonical-3210332120333101-1223301202331101-0031101202123100-3001202032202013-3210021121222231-1213100100201312-0320022211000323-3131232113321222"></a>

## keys property — label_matcher / 013002023003 / 4

Type: `["list", "string"]`. Computed.

The list of label key names that have to match.

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

<a id="canonical-3330310120332223-2212320301210201-0311122222213322-0323232320122021-1231331112020033-2312120211113212-0230003232121111-1202031300131110"></a>

## Next pages — label_matcher / 013002023003 / 5

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0103133211113112-2110333301000123-0132303232311023-0033020113003322-3132123120231322-1200002211210031-0222200323230132-3323113100322303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012020010000323-0302330223323331-1302001030301220-2320212130320120-2233103331302022-3010010223203020-0200311220010001-3323322311202331"></a>

## mum_action — mum_action / 311123333120 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- mum_action

<a id="canonical-1232230311312032-2121331211020302-0213033312300333-0012232230033211-3200033003020130-2312132120000303-0211133100312120-1112302303011322"></a>

Type: `"single"`. Computed.

Modify behavior for a matching request. The modification could be to entirely skip processing.

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

<a id="canonical-3301321133011331-2232311100001113-3222120020230030-0022110102120313-2323001232321010-3123032311332102-1100210033333023-0300103100003202"></a>

## Direct properties — mum_action / 311123333120 / 3

- [default](data-sources--service_policy_rule--reference--group-002.md#canonical-2003020203321223-0123233121100233-1132113230022032-2232011221333200-1312120330233100-3131101132302300-1010003020000323-0231332001211221): complete subsection reference.

- [skip_processing](data-sources--service_policy_rule--reference--group-002.md#canonical-3212231102301113-1020002133121323-3321230232033030-2130133321112213-0102232131102032-1131133320123332-1322221320320130-1203133200212132): complete subsection reference.

<a id="canonical-3221012132203331-2100212013221121-2311021103020021-0132303121312131-0313200002132312-3303000300223331-3133301131310122-3121023121001133"></a>

## Next pages — mum_action / 311123333120 / 4

- [mum_action.default](data-sources--service_policy_rule--reference--group-002.md#canonical-2003020203321223-0123233121100233-1132113230022032-2232011221333200-1312120330233100-3131101132302300-1010003020000323-0231332001211221)
- [mum_action.skip_processing](data-sources--service_policy_rule--reference--group-002.md#canonical-3212231102301113-1020002133121323-3321230232033030-2130133321112213-0102232131102032-1131133320123332-1322221320320130-1203133200212132)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2003020203321223-0123233121100233-1132113230022032-2232011221333200-1312120330233100-3131101132302300-1010003020000323-0231332001211221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312030202110221-0033121213233120-1321010022122212-2213032200231301-3030032302133202-0122220030310220-1132331333102101-3033311110202201"></a>

## mum_action.default — default / 211333202312 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [mum_action](data-sources--service_policy_rule--reference--group-002.md#canonical-0103133211113112-2110333301000123-0132303232311023-0033020113003322-3132123120231322-1200002211210031-0222200323230132-3323113100322303)
- mum_action.default

<a id="canonical-3033220123200322-2031112332001123-2020131332123131-2132100310012032-0330122211030322-1112132322010133-0123033030113010-1001102023302202"></a>

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

<a id="canonical-3233110120023130-1212231311303331-2313220133311113-2213000001132113-3122330000013303-1222332103121022-1010121121301202-2313222122133311"></a>

## Direct properties — default / 211333202312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112002120231330-2212121333102102-0302110110021330-0232102323003310-3221100111133110-0331321212223132-3131231201222011-0032122130210121"></a>

## Next pages — default / 211333202312 / 4

- [mum_action](data-sources--service_policy_rule--reference--group-002.md#canonical-0103133211113112-2110333301000123-0132303232311023-0033020113003322-3132123120231322-1200002211210031-0222200323230132-3323113100322303)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3212231102301113-1020002133121323-3321230232033030-2130133321112213-0102232131102032-1131133320123332-1322221320320130-1203133200212132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023121012222300-3110112012332132-1202133330223122-0323110331032313-0312231303100123-2210000212032301-3312031211010331-2123303210310210"></a>

## mum_action.skip_processing — skip_processing / 233330221020 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [mum_action](data-sources--service_policy_rule--reference--group-002.md#canonical-0103133211113112-2110333301000123-0132303232311023-0033020113003322-3132123120231322-1200002211210031-0222200323230132-3323113100322303)
- mum_action.skip_processing

<a id="canonical-0100312101210231-0302220100200300-3002131133112312-0220212012323312-0100210000133313-0100021301232022-0211222230120232-0023333023103300"></a>

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

<a id="canonical-1013012000012013-2331102113020222-3123312003211020-0012022300112112-0220303212230231-2021301331220022-1110210023320011-1301023200102320"></a>

## Direct properties — skip_processing / 233330221020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322111122203033-0033120212311221-0013122003330230-0000313300200130-1131010203213233-3231320001012113-3310213012133132-3221132212100211"></a>

## Next pages — skip_processing / 233330221020 / 4

- [mum_action](data-sources--service_policy_rule--reference--group-002.md#canonical-0103133211113112-2110333301000123-0132303232311023-0033020113003322-3132123120231322-1200002211210031-0222200323230132-3323113100322303)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1202323031323101-3010322010211020-2201013123220300-0120101231322302-1330322232232230-0002133000323121-2021123103113212-2023022232033013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201230013200313-3302213210113000-0011102120121221-2030301100211110-0311000231012032-1010100111222302-0132320020023203-1032032332220223"></a>

## path — path / 203320102011 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- path

<a id="canonical-0031213201320113-0231310121110121-0121200331113332-0010202322232200-0011332112301101-0322223221223202-2103011210010201-0213023020212223"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1200201332313331-1003321122013311-3130023301210213-0120030200132321-1103303201002300-1222302101102121-2230110313211331-2123022100123111"></a>

## Direct properties — path / 203320102011 / 3

<a id="canonical-2233111103220210-0022110311303002-1021121233222031-1032011321031322-3011332103003213-0010111122202122-0203220132302331-0132212011220102"></a>

<a id="canonical-2103013312320223-0313323320023202-2111123121003032-1011333112033023-1233030321323302-0203311233113323-1121333213022201-3010203032003311"></a>

## encoded_path_matcher property — path / 203320102011 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-1021112133022313-1113200020120301-3112323111032321-1212111001101223-1223331200132322-2122031320002121-0031311213003312-3102303312330211"></a>

<a id="canonical-2130023232020331-3020020030333203-1112220333003302-3133300310031233-2112111331211111-1332011002300331-2231123033322301-3312101002002303"></a>

## exact_values property — path / 203320102011 / 5

Type: `["list", "string"]`. Computed.

List of exact path values to match the input HTTP path against.

Upstream description:

A list of exact path values to match the input HTTP path against.

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

<a id="canonical-0122002222120103-2300311231001111-1311011013331230-2212113133301030-2200212210200301-3200133220230132-0032221313013111-2002222333212313"></a>

<a id="canonical-2113321122310130-0311002131233103-3222012111310111-1223321320133001-1001322112032012-0123310231302200-2333221032312012-1330001310002201"></a>

## invert_matcher property — path / 203320102011 / 6

Type: `"bool"`. Computed.

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

<a id="canonical-3323220100003020-2333331113121203-3310122032320133-0022011132123112-0122122331313132-2301220000322102-1220220211212231-3031330301200212"></a>

<a id="canonical-0033000313300200-3313033232222211-2332101132233122-2033322120200003-3222320122012202-0101323132103222-3332210033012320-1200313130102211"></a>

## prefix_values property — path / 203320102011 / 7

Type: `["list", "string"]`. Computed.

List of path prefix values to match the input HTTP path against.

Upstream description:

A list of path prefix values to match the input HTTP path against.

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

<a id="canonical-2311313133011303-3332310211130130-0201203330202201-1300002230232211-0002001110301003-2323210233001032-3201002301122202-0121200000031030"></a>

<a id="canonical-2020022111320030-0010311220203003-3332203233302312-0320311132112201-1101021230012100-1210333110212232-2333232120211303-3133313301133001"></a>

## regex_values property — path / 203320102011 / 8

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input HTTP path against.

Upstream description:

A list of regular expressions to match the input HTTP path against.

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

<a id="canonical-1222230001213311-0322221330330021-2013011220303322-0002110130221311-0221001011101022-2230231000332100-1302022031220301-2022211030222111"></a>

<a id="canonical-0301102332310330-3112122131211213-2021102322133202-1313030100013002-1023122201220221-3211230121023133-1303103110122132-3220122010102132"></a>

## suffix_values property — path / 203320102011 / 9

Type: `["list", "string"]`. Computed.

List of path suffix values to match the input HTTP path against.

Upstream description:

A list of path suffix values to match the input HTTP path against.

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

<a id="canonical-0123102202300202-3231310213302101-3230011013211232-0303202020230313-1310020313203131-0130000132320312-3201220320000200-3231200022221202"></a>

<a id="canonical-0011223132022301-0033222221322220-1322122230333202-2302220311211331-1313320001010133-2013022330110133-2211111213203132-2000011333010211"></a>

## transformers property — path / 203320102011 / 10

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

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

<a id="canonical-3113112320203120-2311333010020232-2133222132130122-3030310320122130-1013302222023020-2002212003131231-3011013313302331-0330102120220123"></a>

## Next pages — path / 203320102011 / 11

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2101210020323103-3303202202131321-1133232111123110-1302223220031001-1031030010123130-1010013200001221-1302232103001023-1320212133031120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110312220032211-1120012011113000-3310133033333021-0130202310211223-3100333232103102-0102120310111113-1100232303310132-2121103200032112"></a>

## port_matcher — port_matcher / 220010331200 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- port_matcher

<a id="canonical-0131213012220100-3233021322333010-0232201212000122-1202123103003121-0133002011221000-3020313202313003-1303103303310102-0011031113301112"></a>

Type: `"single"`. Computed.

Port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true. Server applies default when omitted.

Upstream description:

A port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true.

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

<a id="canonical-1322320233123301-3203203331010312-2321110312200212-1130222300003321-0231201331210221-2331211302121333-3123333033303312-3333022000022131"></a>

## Direct properties — port_matcher / 220010331200 / 3

<a id="canonical-3330332300030201-2203330233023021-0002032132001020-2012201330321203-1020331230110211-1222103030122130-2321312211030033-3131103112310113"></a>

<a id="canonical-3313020033212011-0102000012020002-3232130220103003-2112322233003333-3223333133133122-3323000022031330-1110320303121120-3223223210120133"></a>

## invert_matcher property — port_matcher / 220010331200 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-3101130102333312-2110110312330220-3020123011133333-0233110000213322-1222311330211123-0101303232310120-0122121013221113-3020002312112110"></a>

<a id="canonical-0122200223001131-1023003330103113-0322023211333000-1312331223312301-2030211201010312-3113203320101001-3121233031020102-0301133130202310"></a>

## ports property — port_matcher / 220010331200 / 5

Type: `["list", "string"]`. Computed.

List of strings, each of which is a single port value or a tuple of start and end port values
separated by '-'. The start and end values are considered to be part of the range.

Upstream description:

A list of strings, each of which is a single port value or a tuple of start and end port values
separated by "-". The start and end values are considered to be part of the range.

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

<a id="canonical-3110223300213321-3301333001112022-3230323231001310-0122310101322132-3003300121300121-1102200301022132-1003232310110320-0022103021033300"></a>

## Next pages — port_matcher / 220010331200 / 6

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0010003202113012-3010221330311030-1033010003322312-0003011202030230-0310211010032310-1130310201331200-0130002311302322-3313032310023203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212331311311303-0321322311230213-1201232202012013-2213132130222221-2132032230320210-1030013032233203-2202113100332202-0312123133111322"></a>

## query_params — query_params / 112033232320 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- query_params

<a id="canonical-0110022223012231-1011022222022021-1231213210021012-2020023231312213-0120003203031103-1110212020200121-1320200132312313-1310013021311300"></a>

Type: `"list"`. Computed.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

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

<a id="canonical-2100333011300110-0021200200230121-0312201333333101-3120032300320331-0000320113322332-1202010221221010-2011233323031310-2030312001223132"></a>

## Direct properties — query_params / 112033232320 / 3

- [check_not_present](data-sources--service_policy_rule--reference--group-002.md#canonical-3220033332213100-2033312120023122-2231132332010002-2033031303113023-2132212313323222-1001212132301000-1210033320000312-2101012002123022): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-002.md#canonical-2330013133300202-1312212010221301-2102030201103121-3303231100230202-0000303311230222-0002320123000121-2011023000302222-1332232020022211): complete subsection reference.

<a id="canonical-1103033200210122-3201213211221303-2230013333312012-1222031203122022-2022101301223122-2330130211320000-2021230222233110-2301132123331121"></a>

<a id="canonical-3313110021012103-2233212133313000-1031111133032031-1013203202012122-0032103112011203-0123120032101032-1320322300230333-3111301111323010"></a>

## invert_matcher property — query_params / 112033232320 / 4

Type: `"bool"`. Computed.

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

- [item](data-sources--service_policy_rule--reference--group-002.md#canonical-3133001330323313-2012212130033223-1311101230031121-0231002310113211-3111102301030122-2300120031010000-0310221131132102-0302320330303002): complete subsection reference.

<a id="canonical-0312121020113322-3203333312133313-1203103311320333-3121102221233032-3201130110310000-0000311232031333-3002323232211121-3023333131220320"></a>

<a id="canonical-0302312011131023-2003313122301313-1100201020212110-2103312321303203-3330300131132202-0130302302313221-1313212023332320-1311322300132231"></a>

## key property — query_params / 112033232320 / 5

Type: `"string"`. Computed.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3102323313013011-2030132333212312-1101212331200201-2112213223120331-2333332101010003-0030113123123303-2120011222113302-0302303203013010"></a>

## Next pages — query_params / 112033232320 / 6

- [query_params.check_not_present](data-sources--service_policy_rule--reference--group-002.md#canonical-3220033332213100-2033312120023122-2231132332010002-2033031303113023-2132212313323222-1001212132301000-1210033320000312-2101012002123022)
- [query_params.check_present](data-sources--service_policy_rule--reference--group-002.md#canonical-2330013133300202-1312212010221301-2102030201103121-3303231100230202-0000303311230222-0002320123000121-2011023000302222-1332232020022211)
- [query_params.item](data-sources--service_policy_rule--reference--group-002.md#canonical-3133001330323313-2012212130033223-1311101230031121-0231002310113211-3111102301030122-2300120031010000-0310221131132102-0302320330303002)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3220033332213100-2033312120023122-2231132332010002-2033031303113023-2132212313323222-1001212132301000-1210033320000312-2101012002123022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321200321121313-3300202221200130-0021121210002031-0210320332112211-3102312012222020-0203321331102311-2212333212020231-0330212320123001"></a>

## query_params.check_not_present — check_not_present / 200113211033 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-0010003202113012-3010221330311030-1033010003322312-0003011202030230-0310211010032310-1130310201331200-0130002311302322-3313032310023203)
- query_params.check_not_present

<a id="canonical-3011101203123210-1300300210101020-3321130320311201-3310230230131332-2213302110000323-0122201221030113-2011323332311220-3000300112000112"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2330323302002300-0123223101203201-1120022320031211-1333000201131132-0231201221100103-1030300101233210-1301101131032233-0032202220333112"></a>

## Direct properties — check_not_present / 200113211033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321113233212333-1213100130023021-1112331121013131-2212212203200003-0113301133021032-0210123231212302-2132101110121221-2220332332012022"></a>

## Next pages — check_not_present / 200113211033 / 4

- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-0010003202113012-3010221330311030-1033010003322312-0003011202030230-0310211010032310-1130310201331200-0130002311302322-3313032310023203)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2330013133300202-1312212010221301-2102030201103121-3303231100230202-0000303311230222-0002320123000121-2011023000302222-1332232020022211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121233012312123-1020203202320231-3323001230102313-0110213301320111-0220001032313010-2113230020022131-0230300210211001-0211020220312223"></a>

## query_params.check_present — check_present / 132300102123 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-0010003202113012-3010221330311030-1033010003322312-0003011202030230-0310211010032310-1130310201331200-0130002311302322-3313032310023203)
- query_params.check_present

<a id="canonical-1002312133133323-0223313003230001-3331200333031002-1332121010001221-2012123200022211-2211201002322331-3223232333131320-1303131103120003"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3223321312301202-1011302113200013-3323012110322031-2221103123001333-3122130031020012-0020222130011220-1212021320203020-1020210313003020"></a>

## Direct properties — check_present / 132300102123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023313311132202-1220021133203331-1233201101103130-1310312112320010-1311030001113013-0103203311212123-0020222111332213-2133312023013132"></a>

## Next pages — check_present / 132300102123 / 4

- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-0010003202113012-3010221330311030-1033010003322312-0003011202030230-0310211010032310-1130310201331200-0130002311302322-3313032310023203)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3133001330323313-2012212130033223-1311101230031121-0231002310113211-3111102301030122-2300120031010000-0310221131132102-0302320330303002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001313021121202-2103033310122110-1203221120001320-3010303002331202-3311022221310203-3232323112100213-2102032003220330-0002031123113230"></a>

## query_params.item — item / 323121313012 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-0010003202113012-3010221330311030-1033010003322312-0003011202030230-0310211010032310-1130310201331200-0130002311302322-3313032310023203)
- query_params.item

<a id="canonical-0330101112120322-3333133122212133-2332122021103233-1322100002311122-3333313123111132-1133023013130202-0221213301303033-2102321000311213"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1220201202012003-0030211021332300-0333020322220003-3133303231022221-2003323331321122-2210331201121132-1122111301311202-3313032033113202"></a>

## Direct properties — item / 323121313012 / 3

<a id="canonical-0233120011232212-2331113232020021-3322333002132232-1133021001021230-3210112012102013-2212101320301333-1120222012313232-2232020233211120"></a>

<a id="canonical-0323011033120131-0020203101231133-1102110000000130-3302110133223103-3122020113120011-1002013210221301-3032000021023132-3210031310210203"></a>

## exact_values property — item / 323121313012 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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

<a id="canonical-1131323033011300-2123101133231321-3221130011313103-3132232220303320-3331131121200131-0002313032300221-2033320002333012-1331313001110000"></a>

<a id="canonical-0110221213232012-0003113030230101-0102112133300032-3322000303210320-3132020332322001-0110300022033133-1213323320130211-0013001023101032"></a>

## regex_values property — item / 323121313012 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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

<a id="canonical-0300320221022220-3000131220203013-1103213222132210-2120301022320023-2123003333033112-1000032112123101-0000111222300023-1001002002122313"></a>

<a id="canonical-0001130233013212-3230122033110322-3120310111000010-2103233030313022-3213201220231303-0021012320222001-0331001203122111-2223311020102112"></a>

## transformers property — item / 323121313012 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

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

<a id="canonical-0313230230002001-0133002120322102-1302330112233322-3121301201022121-1133201212003032-0110230213101031-1221302010131103-0003121023031102"></a>

## Next pages — item / 323121313012 / 7

- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-0010003202113012-3010221330311030-1033010003322312-0003011202030230-0310211010032310-1130310201331200-0130002311302322-3313032310023203)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111302313300233-0210121300112313-3001003322312011-3313123233023032-3311101002121303-2213330312121133-1233300012221322-0332122030220320"></a>

## request_constraints — request_constraints / 312021113323 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- request_constraints

<a id="canonical-0121303130321000-2010323021321020-0130112232322010-2222221123133001-3032121201323213-1210123132003033-2233210030211101-2031120112210020"></a>

Type: `"single"`. Computed.

Configuration parameter for request constraints.

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

<a id="canonical-1033010023301200-2220000033330102-2110120023232100-0111102303011111-2013321212223122-2210211012012132-2210333212220320-0211031023302021"></a>

## Direct properties — request_constraints / 312021113323 / 3

<a id="canonical-2133230320223220-1113123020232100-0233011100311311-2220223033232133-3111131323323231-3232322210332223-2332213312020201-3232132112311022"></a>

<a id="canonical-3101030313130130-1232301313213212-0112210010120313-1222103203003133-1003202112002321-0301303032203122-3210002332211212-2112323011211211"></a>

## max_cookie_count_exceeds property — request_constraints / 312021113323 / 4

Type: `"number"`. Computed.

Match on the Count for all Cookies that exceed this value. Exclusive with
\[max\_cookie\_count\_none\]

Upstream description:

Exclusive with \[max\_cookie\_count\_none\]

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

- [max_cookie_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3322130202102222-2002033301313123-0322013111331003-3201133202022132-2301323122123113-2220101110303020-1322001332112212-0012003221013332): complete subsection reference.

<a id="canonical-1303231131001120-1001120031110232-1103010020023232-3331030223330223-1331132022322023-0322322123113202-0030300232030133-3003313021120013"></a>

<a id="canonical-1010322201331311-1003003003122212-1112311310032302-3032322120112301-2010301030311312-0001310020311230-2330000333101101-0131222121021111"></a>

## max_cookie_key_size_exceeds property — request_constraints / 312021113323 / 5

Type: `"number"`. Computed.

Exclusive with \[max\_cookie\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_key\_size\_none\]

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

- [max_cookie_key_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-0000033220333232-3211001012013030-0300033323331123-3022331331122131-1222202312230032-2133033211233301-2210203013012211-1330313023103330): complete subsection reference.

<a id="canonical-0103100211320112-3303303103023200-0221102113130302-2233122010100221-0132302032002032-0132232101102221-0211311320202031-3102102113001113"></a>

<a id="canonical-0210100010303033-1321332101200102-1200332311202021-1102203221202013-3331123101023213-0212123203112013-0300011130131002-3300010231323332"></a>

## max_cookie_value_size_exceeds property — request_constraints / 312021113323 / 6

Type: `"number"`. Computed.

Exclusive with \[max\_cookie\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_value\_size\_none\]

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

- [max_cookie_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3021022221302101-1312003030231221-0002211013021200-2000330000130120-3310201001130112-0333223310022132-3321210010212212-2011313233033202): complete subsection reference.

<a id="canonical-1011320330133023-1012201121103201-0220020110013000-1103113123322021-0303331033111012-0222221110223300-0121300123013110-3012233211103301"></a>

<a id="canonical-2201113231102010-3031010113122010-0211312112120223-3100212310002001-1333300012122002-0013020012213102-0113012002100110-0023121301320120"></a>

## max_header_count_exceeds property — request_constraints / 312021113323 / 7

Type: `"number"`. Computed.

Match on the Count for all Headers that exceed this value. Exclusive with
\[max\_header\_count\_none\]

Upstream description:

Exclusive with \[max\_header\_count\_none\]

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

- [max_header_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1123213300220232-2101212233311012-1023031233101111-1133011231331301-3023033020130231-2302310003221210-2031231031003312-2122012002210303): complete subsection reference.

<a id="canonical-0021211210121302-2200300210020031-1011203031203000-1102310111013122-2210300122303131-1321322330121202-1101123303012123-1010010202232310"></a>

<a id="canonical-1221323123220202-1221300102311323-3333020120021320-3301313302300100-0033120320213101-2221022122233312-1211130130320003-0231330112023201"></a>

## max_header_key_size_exceeds property — request_constraints / 312021113323 / 8

Type: `"number"`. Computed.

Exclusive with \[max\_header\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_key\_size\_none\]

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

- [max_header_key_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3313303021111131-1302003331000213-3231032112323223-0300030232230302-2211032133132212-1330210212312310-3111332021210102-1021012223230013): complete subsection reference.

<a id="canonical-3020010003200211-0221120102310101-0200113112132022-3220032210112210-1300312232100223-0312220230303033-3110111223313202-0233000323022122"></a>

<a id="canonical-1331310130201313-3213333221113222-1033321223011331-0323022020312321-1132231121300210-1133312203102200-0201012023012001-0220330030213320"></a>

## max_header_value_size_exceeds property — request_constraints / 312021113323 / 9

Type: `"number"`. Computed.

Exclusive with \[max\_header\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_value\_size\_none\]

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

- [max_header_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-0020232011313002-2133122103132211-3000120033100032-3220320232021001-1032321100113000-3311323230103210-1311020111232100-3033222222030303): complete subsection reference.

<a id="canonical-2230210011311131-3220103311102300-3112202003332123-2313200201113222-3011111012102202-2332010230231223-1312300231011212-2302110110222232"></a>

<a id="canonical-3100031222020103-1222231223010013-1100032133001213-2131103312032300-1200223001030101-1322021013122130-0023301112103231-1011022220122001"></a>

## max_parameter_count_exceeds property — request_constraints / 312021113323 / 10

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_count\_none\].

Upstream description:

Exclusive with \[max\_parameter\_count\_none\]

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

- [max_parameter_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3033220222300300-0312132120213313-3011220332122230-1113333033131312-0302330300123213-0300300322310022-3310003202121003-0233013221122221): complete subsection reference.

<a id="canonical-2011121232102213-3111131112321201-0122200021020202-0232203201103211-1211301033333121-1201003213032123-1201033231132220-3230213123111223"></a>

<a id="canonical-1033231013032123-1012203223010301-1030122211221313-1213103222200232-0113002122013312-2030003222000023-2213023130322312-0111002301201120"></a>

## max_parameter_name_size_exceeds property — request_constraints / 312021113323 / 11

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_name\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_name\_size\_none\]

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

- [max_parameter_name_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-0011311103233113-0020332013313232-3021002032111022-0203011021103123-1221333330033203-0121320221312221-2311221221201123-1000231123111131): complete subsection reference.

<a id="canonical-1202300203103111-1022032020110102-0000101131000020-1323231311011110-0201130121012113-3001321201123102-3332012303122003-1231120333210300"></a>

<a id="canonical-1100222203121231-2033132120221211-1113320212312203-2311001032032033-0213202002303311-0131122332103002-2311021011003012-3012002223121030"></a>

## max_parameter_value_size_exceeds property — request_constraints / 312021113323 / 12

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_value\_size\_none\]

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

- [max_parameter_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1312213213333021-1112322013103230-2013321221110002-2031101322120113-2102200330010300-3230203301312001-1100102031023300-3310310200121322): complete subsection reference.

<a id="canonical-2111112211231200-3121212100122210-2203032122130232-1331022232111100-2303220210233302-1320120203011133-3332222213313110-2120333320201121"></a>

<a id="canonical-0020012301321333-2123102030103023-3303221013321023-2203101201213111-3012300001032201-1003321133200221-1203130233223211-3003131021202333"></a>

## max_query_size_exceeds property — request_constraints / 312021113323 / 13

Type: `"number"`. Computed.

Match on the URL Query Size that exceed this value. Exclusive with \[max\_query\_size\_none\]

Upstream description:

Exclusive with \[max\_query\_size\_none\]

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

- [max_query_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1221101332103232-3222233020213003-2223321331211011-0211013303022331-0032003131003020-3223331120000131-2313110313120021-2030012203112000): complete subsection reference.

<a id="canonical-1102112313031123-1113223221133223-0021013202022231-0232200231230232-2121000203210332-0030122300330213-0313312231313112-1213332012132020"></a>

<a id="canonical-2200322120132320-1221331310131331-1233331102231132-1323222323111311-1202233123023101-3120033221322022-3231201201221332-1000313130332323"></a>

## max_request_line_size_exceeds property — request_constraints / 312021113323 / 14

Type: `"number"`. Computed.

Exclusive with \[max\_request\_line\_size\_none\].

Upstream description:

Exclusive with \[max\_request\_line\_size\_none\]

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

- [max_request_line_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1003333222232313-0033202110033020-1231100320211000-0303120101200113-1111233102320332-1322231222013131-1102322123013121-0022031230311130): complete subsection reference.

<a id="canonical-0223002231203300-3122020102321311-1002001001122023-1133210130323002-1103311212233233-3113200020230020-2233021302033331-3200001133030313"></a>

<a id="canonical-3303232311110330-1301233203331322-0002223330313200-1332011023032221-1100003233022123-0113122221303222-0000203130323201-0111011230302000"></a>

## max_request_size_exceeds property — request_constraints / 312021113323 / 15

Type: `"number"`. Computed.

Match on the Request Size that exceed this value. Exclusive with \[max\_request\_size\_none\]

Upstream description:

Exclusive with \[max\_request\_size\_none\]

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

- [max_request_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3103320223032223-3030132100322323-3232212132012203-3001100112123333-3131111113110120-3330312202013022-1332020030033333-0122320311212323): complete subsection reference.

<a id="canonical-2113220021310133-3312302133003221-2120222232020321-0201023231321213-2332101031203212-2330121103303231-2203230022033021-2222030023031102"></a>

<a id="canonical-0031122311311013-1020002200110223-0000302330320302-3300103332220301-2130211232222303-2113020001230223-2312211332202023-2310331211230223"></a>

## max_url_size_exceeds property — request_constraints / 312021113323 / 16

Type: `"number"`. Computed.

Match on the URL Size that exceed this value. Exclusive with \[max\_url\_size\_none\]

Upstream description:

Exclusive with \[max\_url\_size\_none\]

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

- [max_url_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-2210112122120113-1023110233212102-2203103322200020-0322210012030020-3332223301032223-2020120330212011-1010021311233031-1011330313322303): complete subsection reference.

<a id="canonical-1132323032301001-2310231020303133-1223012132033330-0122112030030233-3231101210323033-2202132110323210-3031333132123112-3011012312130232"></a>

## Next pages — request_constraints / 312021113323 / 17

- [request_constraints.max_cookie_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3322130202102222-2002033301313123-0322013111331003-3201133202022132-2301323122123113-2220101110303020-1322001332112212-0012003221013332)
- [request_constraints.max_cookie_key_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-0000033220333232-3211001012013030-0300033323331123-3022331331122131-1222202312230032-2133033211233301-2210203013012211-1330313023103330)
- [request_constraints.max_cookie_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3021022221302101-1312003030231221-0002211013021200-2000330000130120-3310201001130112-0333223310022132-3321210010212212-2011313233033202)
- [request_constraints.max_header_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1123213300220232-2101212233311012-1023031233101111-1133011231331301-3023033020130231-2302310003221210-2031231031003312-2122012002210303)
- [request_constraints.max_header_key_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3313303021111131-1302003331000213-3231032112323223-0300030232230302-2211032133132212-1330210212312310-3111332021210102-1021012223230013)
- [request_constraints.max_header_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-0020232011313002-2133122103132211-3000120033100032-3220320232021001-1032321100113000-3311323230103210-1311020111232100-3033222222030303)
- [request_constraints.max_parameter_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3033220222300300-0312132120213313-3011220332122230-1113333033131312-0302330300123213-0300300322310022-3310003202121003-0233013221122221)
- [request_constraints.max_parameter_name_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-0011311103233113-0020332013313232-3021002032111022-0203011021103123-1221333330033203-0121320221312221-2311221221201123-1000231123111131)
- [request_constraints.max_parameter_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1312213213333021-1112322013103230-2013321221110002-2031101322120113-2102200330010300-3230203301312001-1100102031023300-3310310200121322)
- [request_constraints.max_query_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1221101332103232-3222233020213003-2223321331211011-0211013303022331-0032003131003020-3223331120000131-2313110313120021-2030012203112000)
- [request_constraints.max_request_line_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1003333222232313-0033202110033020-1231100320211000-0303120101200113-1111233102320332-1322231222013131-1102322123013121-0022031230311130)
- [request_constraints.max_request_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3103320223032223-3030132100322323-3232212132012203-3001100112123333-3131111113110120-3330312202013022-1332020030033333-0122320311212323)
- [request_constraints.max_url_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-2210112122120113-1023110233212102-2203103322200020-0322210012030020-3332223301032223-2020120330212011-1010021311233031-1011330313322303)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3322130202102222-2002033301313123-0322013111331003-3201133202022132-2301323122123113-2220101110303020-1322001332112212-0012003221013332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120000020010303-3320212210232020-1232111230312133-1333230221020013-2103311112112330-1320020012300020-0000033013210012-0001023222322030"></a>

## request_constraints.max_cookie_count_none — max_cookie_count_none / 110020032133 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_cookie_count_none

<a id="canonical-1030113020022100-3110132033201201-3012322102131031-2003100123311002-1013100132030132-0013103030302130-0022123120223321-2203020213100101"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2122233022211221-2313032201301221-2330103223202302-3232220012001032-3313201232111110-1001103023211100-2313123110112230-3010120023130330"></a>

## Direct properties — max_cookie_count_none / 110020032133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311233022023311-3213210020313110-1332010311201131-3322000110301221-1311022200302023-3200103323230210-0110033332000210-3100202213102212"></a>

## Next pages — max_cookie_count_none / 110020032133 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0000033220333232-3211001012013030-0300033323331123-3022331331122131-1222202312230032-2133033211233301-2210203013012211-1330313023103330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202332013220130-1000311201033020-1013012121231333-1023213231333200-3022223223103113-0312323203130031-1110122113002213-0322210200020001"></a>

## request_constraints.max_cookie_key_size_none — max_cookie_key_size_none / 221030311031 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_cookie_key_size_none

<a id="canonical-1033333332111222-3003302032233333-0211010132033231-3310003311301003-3212033101301303-1032122303212112-2211301202111332-1031120033122313"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2231223121332312-2011312121231201-3101123103122322-2030203120331303-1333213211332030-2101022113121302-3020013031000012-2030101013131030"></a>

## Direct properties — max_cookie_key_size_none / 221030311031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220212012011302-0302133310131321-0330313312132212-2130112303013321-0211333122322231-2001030303230103-1010322203130211-0301230003030113"></a>

## Next pages — max_cookie_key_size_none / 221030311031 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3021022221302101-1312003030231221-0002211013021200-2000330000130120-3310201001130112-0333223310022132-3321210010212212-2011313233033202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321021013113330-2213310322010120-2101103132211012-1121001131301321-1310013022311300-2231112331312131-1323132330130120-2212131321320131"></a>

## request_constraints.max_cookie_value_size_none — max_cookie_value_size_none / 213310333002 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_cookie_value_size_none

<a id="canonical-0210323111300301-1111012003030002-1322022032031323-3021122231020033-0221120021330323-3121231110103200-1201231031302223-2021232101312332"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2021300131301012-2221301001312132-0021102201332330-1033003102001301-1212031333131103-2210022032200110-0100323200212131-1203300313332033"></a>

## Direct properties — max_cookie_value_size_none / 213310333002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031133201020201-3201202213022301-3223100330220111-3221200023020322-3101312100331222-3032132333333101-0233133002010222-1103111122230210"></a>

## Next pages — max_cookie_value_size_none / 213310333002 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1123213300220232-2101212233311012-1023031233101111-1133011231331301-3023033020130231-2302310003221210-2031231031003312-2122012002210303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112211223130232-0002011111123213-3301201023012321-0222312130211112-3201022010132233-2213031020302103-1021320332000333-3130230232200323"></a>

## request_constraints.max_header_count_none — max_header_count_none / 220321212121 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_header_count_none

<a id="canonical-1011200232203120-1120023222123130-3012332100333113-0203030232001031-2203132022312323-2223131332313013-2310122210301230-3211321030023300"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1100200212303301-1111221003122213-2110331201123220-1232122003313210-3113311023012223-2310101033012132-0122210012132312-2001011333233202"></a>

## Direct properties — max_header_count_none / 220321212121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031232221131222-3133313232021030-0201232320002313-1331332211102202-1332131012030320-2332330202001131-0003331200330231-1311232032322113"></a>

## Next pages — max_header_count_none / 220321212121 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3313303021111131-1302003331000213-3231032112323223-0300030232230302-2211032133132212-1330210212312310-3111332021210102-1021012223230013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221122113013202-2000321011032001-1032331103100313-2310122231213012-1311002330010012-0121022121313123-0201330120332032-2002113321013102"></a>

## request_constraints.max_header_key_size_none — max_header_key_size_none / 223211012211 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_header_key_size_none

<a id="canonical-0101223120103230-1001211233301323-1112113320113113-1203303322022001-2011130311001020-3132003002210131-2331311332120212-1130111111331002"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1110012300011232-0132111001023201-2011023121000123-1333312322233131-3113003211230301-1303213303310313-1122110011013110-2200031002300120"></a>

## Direct properties — max_header_key_size_none / 223211012211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002323233102102-2132022103132012-2202101320023330-2212000112213311-3031213201310210-1130200320220321-3230113111031103-3313302020120103"></a>

## Next pages — max_header_key_size_none / 223211012211 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0020232011313002-2133122103132211-3000120033100032-3220320232021001-1032321100113000-3311323230103210-1311020111232100-3033222222030303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210023120203012-0230231113320210-1033030233011030-3200321203300203-3103123320132133-1313110121002002-0023033312130322-0031123031211031"></a>

## request_constraints.max_header_value_size_none — max_header_value_size_none / 213331222313 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_header_value_size_none

<a id="canonical-3220222232320100-3012231320230101-0123211102311011-1301202121302323-0013311022300300-1220113302230332-0330210112021120-1301310000221100"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3212200233031310-3320000203303322-1113032333012100-1130122210120203-2123230132113321-2013201311300331-1211320002021211-3203132212031033"></a>

## Direct properties — max_header_value_size_none / 213331222313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231120221011103-2311320223333123-2223123330201202-2303033131302013-3013331003023210-2030033333002210-3203211203003013-3320010103031111"></a>

## Next pages — max_header_value_size_none / 213331222313 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3033220222300300-0312132120213313-3011220332122230-1113333033131312-0302330300123213-0300300322310022-3310003202121003-0233013221122221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230231021312221-0023311232220300-1202010131331223-0233002100331333-3332223022310132-0302310012102332-0102113210223011-1221100221111022"></a>

## request_constraints.max_parameter_count_none — max_parameter_count_none / 301223013333 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_parameter_count_none

<a id="canonical-1300321310102013-2312333223021233-0221302313001122-0300013002001321-3023232221132022-0200210310231210-1323002132331321-2022020122222212"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2301121202331300-3130323032233303-0121011011302231-1321330222203322-0102111301120033-1123202201310130-2001300302311233-3133222020000033"></a>

## Direct properties — max_parameter_count_none / 301223013333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102302212210131-0221023213112113-3220303011022020-0110001222112130-1303032101313333-1132100113303101-1011213202011223-3103311100100133"></a>

## Next pages — max_parameter_count_none / 301223013333 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0011311103233113-0020332013313232-3021002032111022-0203011021103123-1221333330033203-0121320221312221-2311221221201123-1000231123111131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321022210011030-0000232202223132-2131132333233132-3133221302311120-2112203021332030-1023103030221333-0212211110112133-3123020202213120"></a>

## request_constraints.max_parameter_name_size_none — max_parameter_name_size_none / 121023121210 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_parameter_name_size_none

<a id="canonical-1312322021102121-2103201031100211-3230031321110122-3102130021333311-0100330121032110-2132301222032332-2213021332010103-2101111300003003"></a>

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

<a id="canonical-2020213310001112-3333220031130130-0013100332313333-0111011120330303-0131232333010002-3333312233103112-1030031222322333-1032333222111102"></a>

## Direct properties — max_parameter_name_size_none / 121023121210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010222000300000-2220103030123113-2331330012022010-0212120102012022-0220310120302320-1010023130201132-1113102302032331-3333230300021032"></a>

## Next pages — max_parameter_name_size_none / 121023121210 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1312213213333021-1112322013103230-2013321221110002-2031101322120113-2102200330010300-3230203301312001-1100102031023300-3310310200121322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132200222202011-1002302023321301-0110302223012012-0001031011123102-3102201022020003-2021012331031331-3310122012120301-3213020011010213"></a>

## request_constraints.max_parameter_value_size_none — max_parameter_value_size_none / 330332322122 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_parameter_value_size_none

<a id="canonical-2002332310321030-2000302113032022-3321111201101313-3332023001302200-0202033321102113-3011330102300031-0330322121001002-0303032201223003"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1120203010001311-0313113320211223-2333130322010313-1213001311232230-3233113111023303-3031313233331201-3133133311121203-0201301013221123"></a>

## Direct properties — max_parameter_value_size_none / 330332322122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133013020111232-2020311021113333-1011202011000201-0022033023001023-3310002301023303-0233230100022311-2333130131003020-3303023232102220"></a>

## Next pages — max_parameter_value_size_none / 330332322122 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1221101332103232-3222233020213003-2223321331211011-0211013303022331-0032003131003020-3223331120000131-2313110313120021-2030012203112000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021032031110203-0200012211302120-0130210121010102-1210100330323302-3310201121123020-1010232321323032-1002333132323102-1223031201120110"></a>

## request_constraints.max_query_size_none — max_query_size_none / 101331201302 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_query_size_none

<a id="canonical-1312001012111122-3011222231323112-0313303221200022-3123133210220220-2302201211112122-2021012313222222-3221332020130101-2231113210010302"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1310001201021113-1332230220111233-1003112103210222-2210210030002111-1333303202211012-1030102000302330-1020220311211101-3120110123212202"></a>

## Direct properties — max_query_size_none / 101331201302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332010212113133-1212001200301310-2013310320101301-3012030030222022-2113331011120022-2311232023002132-0121330101233212-0211122001013010"></a>

## Next pages — max_query_size_none / 101331201302 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1003333222232313-0033202110033020-1231100320211000-0303120101200113-1111233102320332-1322231222013131-1102322123013121-0022031230311130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333032111322011-1312310300002033-2032202011021311-0113211301221122-3113112110312213-2333223310203023-3320102120110012-3102201310111212"></a>

## request_constraints.max_request_line_size_none — max_request_line_size_none / 200210120223 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_request_line_size_none

<a id="canonical-0323333201031003-0221202303322213-1211000011202220-2103103033331130-0230033213322231-3022022322031102-2202101323201231-1310221322311022"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3132100003133122-0032210103210322-2313003103313203-0000103331101330-2133120231201233-0120331230130221-1131110213300131-2113301200130001"></a>

## Direct properties — max_request_line_size_none / 200210120223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133220113333233-3122013020032222-1322320231202310-0331222003003331-0003103321312111-0120221200000003-1222021211121333-3021232321123221"></a>

## Next pages — max_request_line_size_none / 200210120223 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3103320223032223-3030132100322323-3232212132012203-3001100112123333-3131111113110120-3330312202013022-1332020030033333-0122320311212323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333311331010323-1122212302022330-2312110330010110-0221112002133113-1121001013311303-2003031021030031-2021122322221032-3012133310101033"></a>

## request_constraints.max_request_size_none — max_request_size_none / 211012231131 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_request_size_none

<a id="canonical-3133031123331012-0233222031021311-2223123201320023-1333300122021021-3303210120101113-3223310112302002-0003022331123312-0323113122002112"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3111111002311322-3333210202303331-2012331010321030-2131033103321113-0003231222220333-0011122022012013-0302323220331320-0122233331333120"></a>

## Direct properties — max_request_size_none / 211012231131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020101131103202-1311221012020210-1031020320131123-3230230103222302-2230302320022001-2121211333033301-0011220232331111-3221302311013112"></a>

## Next pages — max_request_size_none / 211012231131 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2210112122120113-1023110233212102-2203103322200020-0322210012030020-3332223301032223-2020120330212011-1010021311233031-1011330313322303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112333002331103-2130321322203121-3213230232013031-2132030111012133-2131133130333110-1033123103223313-1201103102030102-1130230312202100"></a>

## request_constraints.max_url_size_none — max_url_size_none / 110021121323 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_url_size_none

<a id="canonical-2231332203033130-2130030012322122-0213330212320022-0232123320130031-0333022121213313-2130200132132020-3022211301330302-3230321033321012"></a>

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

<a id="canonical-0100231111021023-2223232320211301-0102320320012103-2003230121020111-3100213311222033-2130201132112220-0111000303221323-0303003101320032"></a>

## Direct properties — max_url_size_none / 110021121323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303133320031102-3220232020111113-2222200212003123-1003023123003220-3123033321323302-3100013333310030-2311231231303031-3232013303132122"></a>

## Next pages — max_url_size_none / 110021121323 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231221123130011-1101301311032123-3011232230310312-2131330021310031-2321223201003302-3322132102102020-1130132103123231-2013122002201021"></a>

## segment_policy — segment_policy / 321102002132 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- segment_policy

<a id="canonical-2031020012133013-2110330033230313-0220002102100223-3000000223311301-2113321321312110-0312203203302331-1220302301011132-2020123231020323"></a>

Type: `"single"`. Computed.

Configure source and destination segment for policy.

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

<a id="canonical-0213132001120033-1011311011200112-3133003323112020-2321030031223033-2033122223231133-3200222211023333-1310222310210121-0111312021012211"></a>

## Direct properties — segment_policy / 321102002132 / 3

- [dst_any](data-sources--service_policy_rule--reference--group-002.md#canonical-3312001003211322-3111102023210032-2333122021111331-2322113233302331-0023031013010223-0010002123012101-0213223222301232-2221202200330233): complete subsection reference.

- [dst_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-0203122001021002-2032112320100211-0031113222131100-2021121120122222-2132022301122321-1130020103021330-2012101000013121-0023322202200011): complete subsection reference.

- [intra_segment](data-sources--service_policy_rule--reference--group-002.md#canonical-3110120333223021-2310003013131111-1333011202323121-3132323011332011-0002321033221011-3223031123122031-3203222000210122-1310110132001330): complete subsection reference.

- [src_any](data-sources--service_policy_rule--reference--group-002.md#canonical-2221222002213300-0230001012121200-0331310302322332-3020302301210200-0232020133031333-3222110033023212-1120032322233301-2303122233320211): complete subsection reference.

- [src_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-1220001122222130-0303323031120210-0203112221301021-0121332332221201-3321313332210022-1023022113320133-1123213021331003-3121302033103123): complete subsection reference.

<a id="canonical-3123010112101130-0213120312033113-0100313131121110-2000232330101121-0232102323231111-1230103310200013-3300223123200101-1320203203033321"></a>

## Next pages — segment_policy / 321102002132 / 4

- [segment_policy.dst_any](data-sources--service_policy_rule--reference--group-002.md#canonical-3312001003211322-3111102023210032-2333122021111331-2322113233302331-0023031013010223-0010002123012101-0213223222301232-2221202200330233)
- [segment_policy.dst_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-0203122001021002-2032112320100211-0031113222131100-2021121120122222-2132022301122321-1130020103021330-2012101000013121-0023322202200011)
- [segment_policy.intra_segment](data-sources--service_policy_rule--reference--group-002.md#canonical-3110120333223021-2310003013131111-1333011202323121-3132323011332011-0002321033221011-3223031123122031-3203222000210122-1310110132001330)
- [segment_policy.src_any](data-sources--service_policy_rule--reference--group-002.md#canonical-2221222002213300-0230001012121200-0331310302322332-3020302301210200-0232020133031333-3222110033023212-1120032322233301-2303122233320211)
- [segment_policy.src_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-1220001122222130-0303323031120210-0203112221301021-0121332332221201-3321313332210022-1023022113320133-1123213021331003-3121302033103123)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3312001003211322-3111102023210032-2333122021111331-2322113233302331-0023031013010223-0010002123012101-0213223222301232-2221202200330233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032201301130131-2230111323310223-2301311201003220-0323021012103113-0333030033010000-3300021011132332-3101330210332121-0210310003300010"></a>

## segment_policy.dst_any — dst_any / 002102100313 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100)
- segment_policy.dst_any

<a id="canonical-3313032031020120-1013330202221033-2110130331221320-2111001310103112-1122012202100103-0233120001322132-3213000321300112-3223302212323022"></a>

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

<a id="canonical-1223300321310001-1031132003002331-1232330011222210-2111102102033130-0201330330221220-2013120110101330-0230200321001022-2013020303322101"></a>

## Direct properties — dst_any / 002102100313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002322133103123-3311210320120033-3321010123320232-1210310212231101-3130020010110200-1030330111302102-2212130103212303-3030103212131113"></a>

## Next pages — dst_any / 002102100313 / 4

- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0203122001021002-2032112320100211-0031113222131100-2021121120122222-2132022301122321-1130020103021330-2012101000013121-0023322202200011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303001012122121-1133223302120231-2223103312312211-2011222301123021-2100002222320001-3222310332201222-0200331122103312-2230033022002212"></a>

## segment_policy.dst_segments — dst_segments / 221132232211 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100)
- segment_policy.dst_segments

<a id="canonical-0302121223330221-2323103000132013-1012113101320232-2330302231221131-1321322203212230-1220201211000102-2112131033200213-0230231031221233"></a>

Type: `"single"`. Computed.

Configuration parameter for dst segments.

Upstream description:

List of references to Segments.

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

<a id="canonical-0132210131233001-3311013120211331-2232320323332011-1020223001203130-1201011200011111-2103102322322302-1120101223003312-1001303223200323"></a>

## Direct properties — dst_segments / 221132232211 / 3

- [segments](data-sources--service_policy_rule--reference--group-002.md#canonical-0332111202220313-1211021011203133-2101011123002113-1012203023101210-2311232233022002-3120220111032022-3020310330010322-1222030313301322): complete subsection reference.

<a id="canonical-2223013321332200-2233323012103313-1121103032020220-1311022033303103-0103023201130102-1323011301323030-2312331220023213-0311201220110213"></a>

## Next pages — dst_segments / 221132232211 / 4

- [segment_policy.dst_segments.segments](data-sources--service_policy_rule--reference--group-002.md#canonical-0332111202220313-1211021011203133-2101011123002113-1012203023101210-2311232233022002-3120220111032022-3020310330010322-1222030313301322)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0332111202220313-1211021011203133-2101011123002113-1012203023101210-2311232233022002-3120220111032022-3020310330010322-1222030313301322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000303213102022-2322201122011132-0013103011121030-3331022232220220-1222101310221301-0222330210111032-0311132123120331-1113303321201022"></a>

## segment_policy.dst_segments.segments — segments / 130211231332 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100)
- [segment_policy.dst_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-0203122001021002-2032112320100211-0031113222131100-2021121120122222-2132022301122321-1130020103021330-2012101000013121-0023322202200011)
- segment_policy.dst_segments.segments

<a id="canonical-3232232211120100-3120230303110001-1102300322102301-3022023303310101-2020200322022331-2313200300212333-3103002113013110-0132133233132312"></a>

Type: `"list"`. Computed.

Segments. Select list of segments.

Upstream description:

Select list of segments.

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

<a id="canonical-3132322133001231-0122100130010113-0012232111201120-0131220313221122-0223110021322313-0110110223230231-1233101302031130-3023322312312130"></a>

## Direct properties — segments / 130211231332 / 3

<a id="canonical-3000332302212201-0223022223100003-3221203230123302-0211132331113323-3121331020011130-2302132300122321-2112322002213313-2302120011033133"></a>

<a id="canonical-3130201021223100-1330001300230332-1101331302331312-0311120230113312-1223223332112213-2221222013320311-2122021232101110-0031011001100130"></a>

## name property — segments / 130211231332 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0200330213223012-1002321113132022-2112311013013021-2213220213133301-2313231010012113-2021133230311122-1202320030200222-1303123223133302"></a>

<a id="canonical-0130121110200121-0030211233312310-0201122130231021-2302313131132202-1120322203012333-2223223033130131-1000312021212230-2203110323020132"></a>

## namespace property — segments / 130211231332 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2111113012033321-2301213202120313-0301120211112311-0202222112212212-3031332031110122-1133311113133011-2112032012230300-0111123122332120"></a>

<a id="canonical-3021100212102223-2310313320312333-0323022111130333-2033322000003322-2031200013213103-1012002200202133-0211312233010303-1130012022310003"></a>

## tenant property — segments / 130211231332 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0323103101033213-1333232213012202-3012301233022012-2302033221213323-1110213222112203-1223202201102321-3021122221320231-0212000122120011"></a>

## Next pages — segments / 130211231332 / 7

- [segment_policy.dst_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-0203122001021002-2032112320100211-0031113222131100-2021121120122222-2132022301122321-1130020103021330-2012101000013121-0023322202200011)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3110120333223021-2310003013131111-1333011202323121-3132323011332011-0002321033221011-3223031123122031-3203222000210122-1310110132001330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202320130131321-2223331102123021-3332130030332300-3100110232302223-0003123131220001-0112123030311311-3213331020232301-0333321210020330"></a>

## segment_policy.intra_segment — intra_segment / 212010313132 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100)
- segment_policy.intra_segment

<a id="canonical-3112210131131120-1002021133031301-2030012123301030-0322332021312201-2020230200203122-2103312230122302-2111301122321032-2031310133310120"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0333331013220031-3102313212110033-0121321122100023-1132320212222213-3103000322230302-3230230012113210-1103102123202211-1301201321022020"></a>

## Direct properties — intra_segment / 212010313132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213313010230312-1131302220002001-0330222220221221-1223110213223121-2300230213213013-1012102111133101-0202312102101222-3323320202323133"></a>

## Next pages — intra_segment / 212010313132 / 4

- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2221222002213300-0230001012121200-0331310302322332-3020302301210200-0232020133031333-3222110033023212-1120032322233301-2303122233320211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001333100112020-0012202031130110-1123313021230223-1013321110203021-0233110021011322-3103012331232020-0301112113021111-0021200030222333"></a>

## segment_policy.src_any — src_any / 233212211030 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100)
- segment_policy.src_any

<a id="canonical-0122132221001123-3333313131030020-0322321233112230-2221301320210123-1012231022310201-3112130223001331-2101111320212020-1320131003333212"></a>

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

<a id="canonical-2021200321213000-2321210021320222-1003313333330103-0130101032101123-2313001322230123-2131303012022302-0122100030012221-3031031103122312"></a>

## Direct properties — src_any / 233212211030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332011313132012-2203133312230200-1302322010211012-2300102130121013-2210121201210211-2312201033303331-3023031320133132-3300110133103110"></a>

## Next pages — src_any / 233212211030 / 4

- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1220001122222130-0303323031120210-0203112221301021-0121332332221201-3321313332210022-1023022113320133-1123213021331003-3121302033103123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213310121210132-2300212112131022-2300210330001132-2303313223011220-3122122200020221-2000322302220001-1302130032210022-2223103022212332"></a>

## segment_policy.src_segments — src_segments / 111232331323 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100)
- segment_policy.src_segments

<a id="canonical-1331220222003103-0321323122232233-1220010130200003-0211232323121110-3113211321121002-0203320300311103-3022011211201023-2120312032120000"></a>

Type: `"single"`. Computed.

Configuration parameter for src segments.

Upstream description:

List of references to Segments.

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

<a id="canonical-0210211001103303-2313302223000202-1303220201113210-0300222002222301-1100003221132322-0331220313201200-2013221022003222-1322303200312313"></a>

## Direct properties — src_segments / 111232331323 / 3

- [segments](data-sources--service_policy_rule--reference--group-002.md#canonical-1111103113133102-1212131032120233-2013302111013020-0313133233220232-3000002302212102-3220233300003203-1331310312301000-0132213321030222): complete subsection reference.

<a id="canonical-3310001013133311-0323200231213322-1321333101032111-1111321302312023-1331322131121033-2003132311032223-2020121132333031-3321030002101010"></a>

## Next pages — src_segments / 111232331323 / 4

- [segment_policy.src_segments.segments](data-sources--service_policy_rule--reference--group-002.md#canonical-1111103113133102-1212131032120233-2013302111013020-0313133233220232-3000002302212102-3220233300003203-1331310312301000-0132213321030222)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1111103113133102-1212131032120233-2013302111013020-0313133233220232-3000002302212102-3220233300003203-1331310312301000-0132213321030222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303233313312211-2331320323301201-0311310323031112-1201313133213231-0032203311100212-1031013331323123-3113121233211230-2030200210021101"></a>

## segment_policy.src_segments.segments — segments / 230223200230 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100)
- [segment_policy.src_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-1220001122222130-0303323031120210-0203112221301021-0121332332221201-3321313332210022-1023022113320133-1123213021331003-3121302033103123)
- segment_policy.src_segments.segments

<a id="canonical-2122202332330112-1003103211130200-3211002200113032-3332311003220322-3111333032313300-1212002102110302-2101112101020131-2021100111101023"></a>

Type: `"list"`. Computed.

Segments. Select list of segments.

Upstream description:

Select list of segments.

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

<a id="canonical-3313211213302233-0132102312033022-2221000020103311-0302231032102000-0300032330012033-2213212212032010-0120212110220021-1310203101101232"></a>

## Direct properties — segments / 230223200230 / 3

<a id="canonical-0303203332220210-2100313121200330-0333010311222322-3130223110002112-3302120000111133-1033010200322013-0010330213312301-0333110132102000"></a>

<a id="canonical-1310111012113213-3030211122222231-2311300301013231-3031303202322233-0200201210321200-3102030030102312-2000122333332121-1210010201032103"></a>

## name property — segments / 230223200230 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3220122331000311-1012132222132112-3100333013233133-1130222033231310-3021002332003020-0031331121202332-0013131101302211-0120200120221231"></a>

<a id="canonical-2032302100333312-2001330231030202-1131231221001313-2131011330023312-2102221010222102-0000011310022032-2023322221000330-1230130002311132"></a>

## namespace property — segments / 230223200230 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2100203211202333-0002303211210220-3303302120230023-2122221102032310-3322320030122033-3031300020132213-1031310210312030-0113223203220302"></a>

<a id="canonical-0011210202101211-2010123022001222-0211312021130332-2033331221000203-0330011333030213-1031330021212100-2300002210130122-1222112200301220"></a>

## tenant property — segments / 230223200230 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2330310320101131-3111113102313231-3211232120101210-1233223212311013-2302300120131320-1102122131233031-0321300330320110-1003302111313030"></a>

## Next pages — segments / 230223200230 / 7

- [segment_policy.src_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-1220001122222130-0303323031120210-0203112221301021-0121332332221201-3321313332210022-1023022113320133-1123213021331003-3121302033103123)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0133301223123320-2232020220130121-1213202001010012-3222021111322322-0203313233311020-3313132002210111-1222220233313311-1202000232133022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123232230011023-2310223201232123-2011203233331123-3030223302013320-0333211110330113-1301022032323132-2013231322001303-2301203311120110"></a>

## tls_fingerprint_matcher — tls_fingerprint_matcher / 103032013101 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- tls_fingerprint_matcher

<a id="canonical-3210102313131132-1011210000200333-2001031333102312-0013313211310201-0323202000203330-1101132130313031-1020332221301212-0331111311133010"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1033013111213031-0132202220203032-1121010120200131-1120231001113031-1101322101231230-1201202211001022-3233113133110120-2220300312033110"></a>

## Direct properties — tls_fingerprint_matcher / 103032013101 / 3

<a id="canonical-1201310202132333-2011313030300331-1111021022231213-0300320033123200-1001313130111100-0202010121032222-0302120232113203-0211132203223320"></a>

<a id="canonical-1002113112200310-2232112222331100-3002020300101020-0103301130223303-2330330001113302-2000231122333300-0100331213220202-1002101111212110"></a>

## classes property — tls_fingerprint_matcher / 103032013101 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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

<a id="canonical-0111300310230322-0313031200211310-2203021030102020-2220212211033201-2121323001320022-1313300002012111-2202223323213133-2323023032321023"></a>

<a id="canonical-2100323312103311-1020010110303302-0031222001120220-2002231221331201-0300322222200220-0202102103012333-2321212333303221-1203101300320002"></a>

## exact_values property — tls_fingerprint_matcher / 103032013101 / 5

Type: `["list", "string"]`. Computed.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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

<a id="canonical-2311311013303022-2012303321220002-1112212003323102-3023102302300221-1311031022303202-0211211312323221-2323132003000201-0132232332130003"></a>

<a id="canonical-3112302300211123-3023122120213123-0022013001221322-0300033320020223-1222123133210123-2330110211010231-2013233011222302-1313121213010132"></a>

## excluded_values property — tls_fingerprint_matcher / 103032013101 / 6

Type: `["list", "string"]`. Computed.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

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

<a id="canonical-3220313221323003-0210201320023211-1002330233311122-1132312201303333-3332112333222003-1323333030211101-2122101033101322-2003310022022223"></a>

## Next pages — tls_fingerprint_matcher / 103032013101 / 7

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311320232333123-0201011102112110-0210321131122222-1302330222122313-2230221120302302-0111132211002020-3003002113101011-3002003211011222"></a>

## waf_action — waf_action / 323123130300 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- waf_action

<a id="canonical-1220313201012031-1031100221113212-2331033213200133-0012122003101020-1232003033001223-0233201223112301-1301311031211221-1212332121012202"></a>

Type: `"single"`. Computed.

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Upstream description:

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

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

<a id="canonical-3233001232033222-3130201000323323-1102300222200313-3233031031001320-0321003023030010-2300111131312221-2001220113311001-2232230201111330"></a>

## Direct properties — waf_action / 323123130300 / 3

- [app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-1100223002122311-3102112101011001-3323030310110133-2001211223312132-0323301301012311-0321211203211032-2320021103030123-3323102223202123): complete subsection reference.

- [none](data-sources--service_policy_rule--reference--group-002.md#canonical-3121120301100101-1331232132103033-1110101130132000-1130213310122112-1002120313101300-2212300123223332-0322130032212330-1101320023313000): complete subsection reference.

- [waf_skip_processing](data-sources--service_policy_rule--reference--group-002.md#canonical-2311333232203120-3100201011120211-1202102120222132-0330112333122010-1332320311113211-0323311120133301-2023110211212033-0211231032301230): complete subsection reference.

<a id="canonical-3223200012013011-1101133310303101-3201102320012230-0310211112312220-2223213303233213-2332223032110332-0012003013321010-1130001111011221"></a>

## Next pages — waf_action / 323123130300 / 4

- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-1100223002122311-3102112101011001-3323030310110133-2001211223312132-0323301301012311-0321211203211032-2320021103030123-3323102223202123)
- [waf_action.none](data-sources--service_policy_rule--reference--group-002.md#canonical-3121120301100101-1331232132103033-1110101130132000-1130213310122112-1002120313101300-2212300123223332-0322130032212330-1101320023313000)
- [waf_action.waf_skip_processing](data-sources--service_policy_rule--reference--group-002.md#canonical-2311333232203120-3100201011120211-1202102120222132-0330112333122010-1332320311113211-0323311120133301-2023110211212033-0211231032301230)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1100223002122311-3102112101011001-3323030310110133-2001211223312132-0323301301012311-0321211203211032-2320021103030123-3323102223202123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233303211031232-1330012331002200-2022013331020022-1133132123311303-1313002102223300-1103101000203231-0112000133202013-1300333111200212"></a>

## waf_action.app_firewall_detection_control — app_firewall_detection_control / 222022021012 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322)
- waf_action.app_firewall_detection_control

<a id="canonical-2331311321100102-3032303201210011-1330122001130320-0220011133303031-2032020232100221-2212011011122111-1201010021202110-1112003013022322"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3130033023112001-2000030113133133-0121323223322103-1331212202321003-3213300312302010-2200030212121132-1012023003123232-2031100320221013"></a>

## Direct properties — app_firewall_detection_control / 222022021012 / 3

- [exclude_attack_type_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-2332000210223121-0000311311302203-3122301132322312-2120112201212100-1031011232221031-1002202103121002-2213010122131020-0313223220130310): complete subsection reference.

- [exclude_bot_name_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-3302033000300011-3013302321111020-2110010100000121-2010323033233131-0300311211201222-1303320102013123-2201212112123000-2230102300331302): complete subsection reference.

- [exclude_signature_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-1211112033312221-0103030130013020-1222232010230220-3023110230102122-2301201220123121-2032201233020130-1113213233333231-1222030232023033): complete subsection reference.

- [exclude_violation_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-2221110203023000-3300303310313220-3323122130001232-1311211120003200-0203220113030132-3321231002310331-0200013222222033-3011321232223312): complete subsection reference.

<a id="canonical-1131111210010302-0231003332102300-1123020312313232-3113221312120333-2212130311233223-1133013032211023-3133322121313313-3200011231330032"></a>

## Next pages — app_firewall_detection_control / 222022021012 / 4

- [waf_action.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-2332000210223121-0000311311302203-3122301132322312-2120112201212100-1031011232221031-1002202103121002-2213010122131020-0313223220130310)
- [waf_action.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-3302033000300011-3013302321111020-2110010100000121-2010323033233131-0300311211201222-1303320102013123-2201212112123000-2230102300331302)
- [waf_action.app_firewall_detection_control.exclude_signature_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-1211112033312221-0103030130013020-1222232010230220-3023110230102122-2301201220123121-2032201233020130-1113213233333231-1222030232023033)
- [waf_action.app_firewall_detection_control.exclude_violation_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-2221110203023000-3300303310313220-3323122130001232-1311211120003200-0203220113030132-3321231002310331-0200013222222033-3011321232223312)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2332000210223121-0000311311302203-3122301132322312-2120112201212100-1031011232221031-1002202103121002-2213010122131020-0313223220130310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223121333132121-2310223202001201-0202211203303020-3312110321232101-3203021021012233-2131330231223212-2203110202123310-1131110321211131"></a>

## waf_action.app_firewall_detection_control.exclude_attack_type_contexts — exclude_attack_type_contexts / 331222023221 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322)
- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-1100223002122311-3102112101011001-3323030310110133-2001211223312132-0323301301012311-0321211203211032-2320021103030123-3323102223202123)
- waf_action.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-2110333001322033-3212330132312023-2330020120202231-0121221113313303-0101321111020230-1022323002031220-3022010131310031-1331113203131100"></a>

Type: `"list"`. Computed.

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

<a id="canonical-0311101030212300-3220223212322230-3011030113233101-0012110330233123-2332223223013131-1323101123333311-1232100100021211-2131310032100103"></a>

## Direct properties — exclude_attack_type_contexts / 331222023221 / 3

<a id="canonical-2113310221211320-2213003332212012-1010311030023122-3131030133333303-2023021211310101-1333201322312013-0123003330103331-0123001101320300"></a>

<a id="canonical-2130300210230310-2013323100233030-0221132213311233-1233300301330323-3010201331100311-3032111231232031-1102322123322131-3013220203230111"></a>

## context property — exclude_attack_type_contexts / 331222023221 / 4

Type: `"string"`. Computed.

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

<a id="canonical-3300002011323003-0100210313232001-1110331121321031-2231223130233123-3231311103331202-0102113320311201-0121023102100021-1213311223213221"></a>

<a id="canonical-0223011301031311-3220301323331130-1333210213102203-3022310332031003-2201310112020222-0231233113320023-1021330321013021-1312202223232231"></a>

## context_name property — exclude_attack_type_contexts / 331222023221 / 5

Type: `"string"`. Computed.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

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

<a id="canonical-1211033121233210-2001330110022110-1212012101323131-1231002323300101-1332301302330003-1030013122233100-2001002000032232-1212302103320212"></a>

<a id="canonical-3200032121233123-3121231033322233-1121203231103013-0100300023333120-1130001311101331-3300131131330110-1301223332320312-1021333200001232"></a>

## exclude_attack_type property — exclude_attack_type_contexts / 331222023221 / 6

Type: `"string"`. Computed.

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

<a id="canonical-1322333211003332-1010013020110010-3331202233233300-2332323010033331-2030123130132131-0101002122331322-1012212012103110-2131300020223331"></a>

## Next pages — exclude_attack_type_contexts / 331222023221 / 7

- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-1100223002122311-3102112101011001-3323030310110133-2001211223312132-0323301301012311-0321211203211032-2320021103030123-3323102223202123)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3302033000300011-3013302321111020-2110010100000121-2010323033233131-0300311211201222-1303320102013123-2201212112123000-2230102300331302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032023233202003-2302120033232000-2200111101012032-1000101300211031-3301230111033010-0133133131210230-2213001332322330-3330332333200330"></a>

## waf_action.app_firewall_detection_control.exclude_bot_name_contexts — exclude_bot_name_contexts / 320002120121 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322)
- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-1100223002122311-3102112101011001-3323030310110133-2001211223312132-0323301301012311-0321211203211032-2320021103030123-3323102223202123)
- waf_action.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-1323012021120132-1031132233323102-3211013020103321-2112000300232212-2023312121323103-0231322012312103-3211313001323113-0220330100211112"></a>

Type: `"list"`. Computed.

Bot Names to be excluded for the defined match criteria.

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

<a id="canonical-2232322032201232-1012230202131022-1021112303203031-1322000000011222-2300331312111213-2112000130011011-0320003132022033-1131301022322332"></a>

## Direct properties — exclude_bot_name_contexts / 320002120121 / 3

<a id="canonical-1232210032121031-2102112302203332-0311331022231000-1203203211311212-0003220223012213-2222223113122222-0000203231330200-0133122123020101"></a>

<a id="canonical-0322230303111002-3211333022132302-1030023233303301-3302333331031332-3100130033213213-0013031111301313-0111332313123112-0321223201010232"></a>

## bot_name property — exclude_bot_name_contexts / 320002120121 / 4

Type: `"string"`. Computed.

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

<a id="canonical-3111012300332301-3121120221111012-0033222033302221-3313222031330213-0220112211001303-3010303101130003-0300331320010231-2031331202120002"></a>

## Next pages — exclude_bot_name_contexts / 320002120121 / 5

- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-1100223002122311-3102112101011001-3323030310110133-2001211223312132-0323301301012311-0321211203211032-2320021103030123-3323102223202123)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1211112033312221-0103030130013020-1222232010230220-3023110230102122-2301201220123121-2032201233020130-1113213233333231-1222030232023033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022302312131200-1110302000132002-2123132103201113-3133001202312322-2103211332302211-0131322333221310-0023303023201310-1312130233300201"></a>

## waf_action.app_firewall_detection_control.exclude_signature_contexts — exclude_signature_contexts / 033330223220 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322)
- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-1100223002122311-3102112101011001-3323030310110133-2001211223312132-0323301301012311-0321211203211032-2320021103030123-3323102223202123)
- waf_action.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-0322300030322013-2013113332113001-3200111021300020-1302012212311003-1202002210220120-1010211321110331-2002123332003200-3003012101302230"></a>

Type: `"list"`. Computed.

Signature IDs to be excluded for the defined match criteria.

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

<a id="canonical-2230123213131110-0100223011003122-1220330232000121-1110332022231001-2122330121310022-2023232212223103-0332020213221311-0101312113300022"></a>

## Direct properties — exclude_signature_contexts / 033330223220 / 3

<a id="canonical-1231302122311322-0001222220331223-1102032101332033-0330312010310000-1033203001012233-2313310113220301-0221312031013023-0121203013312203"></a>

<a id="canonical-1303010330022213-2231113131230122-2121020213001323-0032200000321133-2211033023201322-1012200032130000-1323102221333213-1030212210132311"></a>

## context property — exclude_signature_contexts / 033330223220 / 4

Type: `"string"`. Computed.

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

<a id="canonical-0203021032133203-3300110231231312-0022101322201303-2231112010212010-2032033320333023-0133332111302112-2103033013300202-3210003332301013"></a>

<a id="canonical-2233130000202212-1232012233133301-3102112021213100-3121233102103210-3113013112010113-1202112321210322-1211210003021101-1210121012202032"></a>

## context_name property — exclude_signature_contexts / 033330223220 / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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

<a id="canonical-0232033020320133-3020132120123303-2103000223030333-2312313233022222-2023031302120310-3231213311020303-0232301220322330-2000012010230320"></a>

<a id="canonical-3211020221300123-2030021022100210-0330330110023013-1110023122203212-3103011103301010-2021311132311303-3331010312123033-0331300321211030"></a>

## signature_id property — exclude_signature_contexts / 033330223220 / 6

Type: `"number"`. Computed.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

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

<a id="canonical-0011032030321001-0032221211000013-1321223223131231-0013130221312120-1123110311131221-0022202001021010-3112132310033021-0113032123302111"></a>

## Next pages — exclude_signature_contexts / 033330223220 / 7

- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-1100223002122311-3102112101011001-3323030310110133-2001211223312132-0323301301012311-0321211203211032-2320021103030123-3323102223202123)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2221110203023000-3300303310313220-3323122130001232-1311211120003200-0203220113030132-3321231002310331-0200013222222033-3011321232223312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210111130121311-3220030303103100-0122330012310323-0213131213032003-0333323120203013-0320013121001021-0322130230301002-0020000021123313"></a>

## waf_action.app_firewall_detection_control.exclude_violation_contexts — exclude_violation_contexts / 001313222221 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322)
- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-1100223002122311-3102112101011001-3323030310110133-2001211223312132-0323301301012311-0321211203211032-2320021103030123-3323102223202123)
- waf_action.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-3003011100221213-3333131100032013-2321120031100020-3001030001013121-3003223202131031-1323203230200311-1233330110300100-2130031301233321"></a>

Type: `"list"`. Computed.

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

<a id="canonical-1330223020233003-1221031320011323-1013310131132213-1332011013010010-2121231231201221-3021030322322220-3211301110332012-0312133232200103"></a>

## Direct properties — exclude_violation_contexts / 001313222221 / 3

<a id="canonical-2123203313000031-1200103131002010-1030132003131312-2133021111313213-1011233331103331-0131003113102120-0320203331032220-3202131210111012"></a>

<a id="canonical-3010110321301122-3323002322220000-1332312300323200-2001022032332132-0000033022203110-2022212220213313-1320000203031111-0123022333300000"></a>

## context property — exclude_violation_contexts / 001313222221 / 4

Type: `"string"`. Computed.

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

<a id="canonical-1123323020223102-0222030001322223-2000011333213332-2223032301212022-2300311102133032-1100121133013232-1331020112200312-1001310101011233"></a>

<a id="canonical-3120121300100131-3221100121312010-1021012033022222-2330131100111301-0310300131200121-1310223000133331-1012000122133123-3230102221032210"></a>

## context_name property — exclude_violation_contexts / 001313222221 / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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

<a id="canonical-2321313031323311-1212331103003003-2010021010303023-2211310001132312-0303210332002232-3011223201030122-0223103203003320-0010220212112320"></a>

<a id="canonical-2133133101233332-3312211331001020-2233212301110001-2323222221023001-3230131103112103-2232201113110301-2112123313222332-3323203230333300"></a>

## exclude_violation property — exclude_violation_contexts / 001313222221 / 6

Type: `"string"`. Computed.

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

<a id="canonical-1313100033132111-2201213212223120-1300111203121312-0112210012020100-0222112222011020-1211202330322000-0003131113030223-3201033002102022"></a>

## Next pages — exclude_violation_contexts / 001313222221 / 7

- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-1100223002122311-3102112101011001-3323030310110133-2001211223312132-0323301301012311-0321211203211032-2320021103030123-3323102223202123)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3121120301100101-1331232132103033-1110101130132000-1130213310122112-1002120313101300-2212300123223332-0322130032212330-1101320023313000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032131203231031-2333101211212121-1210110113333030-3013130130002031-3032302112102021-0311302010121332-0132001331211312-0032101130231312"></a>

## waf_action.none — none / 312003012311 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322)
- waf_action.none

<a id="canonical-1312012132001212-2311312302222231-1013103310113003-2220023033202101-2211301333032312-3302222131002000-2321012332232203-0303132310333322"></a>

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

<a id="canonical-0233201231011330-2130302302011122-3020310021132011-2030202332100021-3313323101132210-0200321323030201-0123002002100130-0303030030212031"></a>

## Direct properties — none / 312003012311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301100222030331-3211132121313213-3223231103031203-2200003003200011-3232230100301301-0232100300223001-3311102331301311-1322012123103300"></a>

## Next pages — none / 312003012311 / 4

- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2311333232203120-3100201011120211-1202102120222132-0330112333122010-1332320311113211-0323311120133301-2023110211212033-0211231032301230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212101232100031-3332020023020033-0223012101033132-1321211300230131-2222313331122313-0331131233221010-0321000323301010-2001010030000112"></a>

## waf_action.waf_skip_processing — waf_skip_processing / 201230101232 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322)
- waf_action.waf_skip_processing

<a id="canonical-1222213123001231-0232113320132011-3130011302332122-3301210313321330-3323320310132212-3122133011230113-2020312100111231-1022031213132011"></a>

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

<a id="canonical-3100032203200211-2232300310100030-2010232023201032-0110302331023212-3103320121212220-1012321022011001-0121000313233301-0222100210333303"></a>

## Direct properties — waf_skip_processing / 201230101232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331012312222101-0223221311213212-2233313202333030-1003232012223320-3012222022111212-0100111221233211-3102020112311312-2313002321132033"></a>

## Next pages — waf_skip_processing / 201230101232 / 4

- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
