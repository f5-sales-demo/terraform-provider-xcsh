---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-3102121211210100-2321031112221031-3332121031020212-0221332022333322-0121222111232033-3002133121101100-3313113102121223-2133232111221313"></a>

## bot_defense.policy.protected_app_endpoints.headers.item — item / 321212321122 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.headers](resources--cdn_loadbalancer--reference--group-008.md#canonical-1100110031020322-1021330213010003-0313300320311220-0101003231320232-2021332111310001-1330021302111232-3020223001131222-0310111230322302)
- bot_defense.policy.protected_app_endpoints.headers.item

<a id="canonical-0000022102332312-3312113303122120-2322112321300132-0000011222101333-3213222013333123-1112001013031012-3033032101330233-0320233303122003"></a>

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

<a id="canonical-2003322003112031-1030331123113332-3203230031010102-0012232000103001-1013023312103322-0333222212100321-1220132002311001-2303212300223100"></a>

## Direct properties — item / 321212321122 / 3

<a id="canonical-2030113323010210-1210210220002032-1033323003113123-3332110301312031-0122310122230331-0021200112302021-1021223000313131-2122111030232002"></a>

<a id="canonical-0230010003321223-3202102021010212-2132302031313123-0013331301031303-2123303303331311-0112213322332130-3231303321102133-1133233332012100"></a>

## exact_values property — item / 321212321122 / 4

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

<a id="canonical-3321221000101103-3232020001030213-3132010201020330-1101001200211110-0121121121221122-3231213321231332-2312013210101022-3313002022211210"></a>

<a id="canonical-2022103000203111-2121111200203301-2123012323132320-3130210102111133-0012120331303301-3011333332230323-0310021223331033-2203021021202211"></a>

## regex_values property — item / 321212321122 / 5

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

<a id="canonical-2120310000310101-1201331203320111-0313302111310030-0030232103000212-0233312121210002-1213010212323233-1100213132123102-2130112011121322"></a>

<a id="canonical-3021231001122031-0233222330222330-3020000322012231-3203223331231102-1333022300011121-0331130123232201-0111001011320010-2312030230120030"></a>

## transformers property — item / 321212321122 / 6

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

<a id="canonical-1221113313322022-0320101102123033-0202212222333133-3311301112102232-3102000303211322-0120311101131210-1212201230303103-2012300123023313"></a>

## Next pages — item / 321212321122 / 7

- [bot_defense.policy.protected_app_endpoints.headers](resources--cdn_loadbalancer--reference--group-008.md#canonical-1100110031020322-1021330213010003-0313300320311220-0101003231320232-2021332111310001-1330021302111232-3020223001131222-0310111230322302)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3310130131321203-3331122133211030-1121020112102101-0213323211311213-0102323212323210-3022231033021330-1000312320331032-2330022020312100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333100313221303-2303121221232331-3022232101222331-3100003211000032-1203103212101302-0110121221310301-1122031120201022-1330130212010003"></a>

## bot_defense.policy.protected_app_endpoints.metadata — metadata / 000303233202 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.metadata

<a id="canonical-3323121012011302-1002221011223033-3303130323033013-1133101300101123-0213000133111313-2211323303120232-0200103221320103-1313233013110022"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223232323120223-0202103320122003-0012221003312013-1013001003031101-2023212320102310-0123331301111031-3100001303211200-3323121031100301"></a>

## Direct properties — metadata / 000303233202 / 3

<a id="canonical-3110013123302312-0112311333133020-2011033021113002-1212103103011030-2300010310022002-3103010123123202-0220322210112103-1010131300013100"></a>

<a id="canonical-1300022200200030-1023003000112313-1210113303022231-3213201222000331-0223231100201100-3113112202331033-1113311301010211-0000121312122333"></a>

## description_spec property — metadata / 000303233202 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2303310300330213-0123030323120032-3232113330321300-2003221011303023-3321021102132021-0121221132211221-1310320211102311-2220102220221223"></a>

<a id="canonical-1223220221233012-3233223122010213-2103313012210123-1020012210022233-2003302323030030-0011021130330132-2032113130210202-3333131123312200"></a>

## name property — metadata / 000303233202 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0031031233231112-1011100121110121-2031333302111331-2202013132232223-1310032031111331-3320003203111211-2011003323102101-3202112310223122"></a>

## Next pages — metadata / 000303233202 / 6

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0311023021301223-0213131330202333-3031023111101230-0230200120220300-0032012201030313-3202221121302331-3031112300021232-2220312200203022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120030131033102-3302003201221013-3301110002321220-3323113101021002-0231113011200200-1313323011310021-2023331212312121-0213333121331122"></a>

## bot_defense.policy.protected_app_endpoints.mitigate_good_bots — mitigate_good_bots / 311303333323 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.mitigate_good_bots

<a id="canonical-1013302330330003-0123313002013210-2123031030203131-1111313123121222-3001032303013320-2030213111103033-2202212010102000-3113130200132231"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for mitigate good bots.

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
mitigate_good_bots = {}
```

<a id="canonical-3313332230131302-1123122001000012-1203132233201021-3230320221120301-1301103232222003-3012321101032230-3322031033121323-2103111031100231"></a>

## Direct properties — mitigate_good_bots / 311303333323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003213010101102-0110320102013332-3222310121322122-3011202110011131-2011001013311010-0123221321201021-3031302220202003-2212332311110312"></a>

## Next pages — mitigate_good_bots / 311303333323 / 4

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330313233230120-0033003213020022-1221333003320022-3113300221030200-2213301020121011-2331002300231002-1210223130313303-0333320211102220"></a>

## bot_defense.policy.protected_app_endpoints.mitigation — mitigation / 220231022213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.mitigation

<a id="canonical-3120013201021020-1230100333010221-1220101222133200-3333230330012101-3210033001301312-3302332202232313-1023201220212033-0032332231131321"></a>

Type: `"object"`. single nested block, Optional.

Modify Bot Defense behavior for a matching request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "flag"),
  validators.ConflictingObjectAttributes("block",
    "redirect"),
  validators.ConflictingObjectAttributes("flag",
    "redirect")}
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
  "x-ves-oneof-field-action_type": "[\"block\",\"flag\",\"redirect\"]"
}
```

Terraform syntax:

```terraform
mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-1221003022332203-1101321132203213-0300112110330020-2131013221311132-0110011300200130-0132103121310333-0212313323233001-3030303221000211"></a>

## Direct properties — mitigation / 220231022213 / 3

- [block](resources--cdn_loadbalancer--reference--group-009.md#canonical-3302302022001231-3010031103003012-3133003223203201-1222003000111013-3020303020323011-2101021023230301-3212130200232123-2221312212222312): complete subsection reference.

- [flag](resources--cdn_loadbalancer--reference--group-009.md#canonical-3000200131002020-3221332213012332-2210213230212022-3133311113302123-3310211110332232-2230130301130221-0020030110210203-3233023122232110): complete subsection reference.

- [redirect](resources--cdn_loadbalancer--reference--group-009.md#canonical-1313313330213113-1313101230021132-1120203311013321-3211102113010322-1100103012220022-0323132213332203-2132220310032211-1113203122030233): complete subsection reference.

<a id="canonical-2102002201013210-0013013010320333-3203332300321202-0120001121332101-0302100223001121-3102312110033303-3103201133311203-2323023131002113"></a>

## Next pages — mitigation / 220231022213 / 4

- [bot_defense.policy.protected_app_endpoints.mitigation.block](resources--cdn_loadbalancer--reference--group-009.md#canonical-3302302022001231-3010031103003012-3133003223203201-1222003000111013-3020303020323011-2101021023230301-3212130200232123-2221312212222312)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--cdn_loadbalancer--reference--group-009.md#canonical-3000200131002020-3221332213012332-2210213230212022-3133311113302123-3310211110332232-2230130301130221-0020030110210203-3233023122232110)
- [bot_defense.policy.protected_app_endpoints.mitigation.redirect](resources--cdn_loadbalancer--reference--group-009.md#canonical-1313313330213113-1313101230021132-1120203311013321-3211102113010322-1100103012220022-0323132213332203-2132220310032211-1113203122030233)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3302302022001231-3010031103003012-3133003223203201-1222003000111013-3020303020323011-2101021023230301-3212130200232123-2221312212222312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311103331310310-2221030330112321-1013222320220331-3122012012302113-1023111301321210-3330332132121322-3132210032111213-2220311302030330"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.block — block / 113221331220 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-009.md#canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101)
- bot_defense.policy.protected_app_endpoints.mitigation.block

<a id="canonical-0323030001213121-2133101120032000-1321222331302110-1000111003221032-0310202233221012-3330110103122123-2330320013023013-3312200100310132"></a>

Type: `"object"`. single nested block, Optional.

Block request and respond with custom content.

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
block {
  # Configure direct properties listed below.
}
```

<a id="canonical-1123011323322222-0300021332011303-1200301122323032-1233220112122020-1201002010233201-1312131030103313-2112123331232221-1120022011001321"></a>

## Direct properties — block / 113221331220 / 3

<a id="canonical-0110112102103012-0131113133333331-3001012011213303-3202233313013100-3200131333322302-0021300230033111-1203011123000231-1121030000023310"></a>

<a id="canonical-2033000222031021-2123013102121200-0303130203223222-1131000131330101-0003023202011231-2113000123331210-1203310300013331-0333020231202210"></a>

## body property — block / 113221331220 / 4

Type: `"string"`. Optional.

Custom body message is of type URI\_ref. Currently supported URL schemes is string:///. For
string:/// scheme, message needs to be encoded in base64 format.

Upstream description:

Custom body message is of type URI\_ref. Currently supported URL schemes is string:///. For
string:/// scheme, message needs to be encoded in base64 format. You can specify this message as
base64 encoded plain text message e.g. "Your request was blocked" or it can be HTML paragraph or a
body string encoded as base64 string E.g. "&lt;p&gt; Your request was blocked &lt;/p&gt;". base64
encoded string for this HTML is "LzxwPiBZb3VyIHJlcXVlc3Qgd2FzIGJsb2NrZWQgPC9wPg=="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2302312332023201-2122100230012212-3221102201013001-0111013132213111-3231312112201122-2021202231302332-1013320022331230-2333330211302222"></a>

<a id="canonical-0110323233022301-0330213212303101-1301311233120123-3210132110121201-1010311132113220-2230322130113300-0121123223110330-3322123323333031"></a>

## status property — block / 113221331220 / 5

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3233100233201132-1131210331001002-0111020201011123-0312033021310021-0312132313130123-2110103122131311-0021322221111020-2031301010312021"></a>

## Next pages — block / 113221331220 / 6

- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-009.md#canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3000200131002020-3221332213012332-2210213230212022-3133311113302123-3310211110332232-2230130301130221-0020030110210203-3233023122232110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223302031323203-3003211310113312-2330001101002301-0023110011101312-0033323310003111-0233223130222121-2222123012322333-0212033333031222"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.flag — flag / 220313011111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-009.md#canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101)
- bot_defense.policy.protected_app_endpoints.mitigation.flag

<a id="canonical-0122001120010023-1001223032122011-2220320331210223-1033303123010310-2333133112332031-3202131330233311-2232330232321202-2200013112012002"></a>

Type: `"object"`. single nested block, Optional.

Select Flag Bot Mitigation Action. Flag mitigation action.

Upstream description:

Flag mitigation action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_headers",
    "no_headers")}
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
  "x-ves-oneof-field-send_headers_choice": "[\"append_headers\",\"no_headers\"]"
}
```

Terraform syntax:

```terraform
flag {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011223211112020-2312202302110232-0321013210212333-3120100101103323-3330213312301100-1321220000002110-3112013003022032-2222022312101201"></a>

## Direct properties — flag / 220313011111 / 3

- [append_headers](resources--cdn_loadbalancer--reference--group-009.md#canonical-2021023122003030-3032103120101121-2110000210020320-0120132233233113-3103022001101000-3033222000102221-3320231003123303-2001111023022230): complete subsection reference.

- [no_headers](resources--cdn_loadbalancer--reference--group-009.md#canonical-2232022031111012-2123120231300022-0230012213330022-0013201210111203-0320113121001121-1221212111201023-1102322003200313-2102003100103220): complete subsection reference.

<a id="canonical-3010023211112211-1111221211033221-1223010020333023-0111021332211232-1313302133332032-2111120310021201-1233313120333322-3301001210010302"></a>

## Next pages — flag / 220313011111 / 4

- [bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers](resources--cdn_loadbalancer--reference--group-009.md#canonical-2021023122003030-3032103120101121-2110000210020320-0120132233233113-3103022001101000-3033222000102221-3320231003123303-2001111023022230)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers](resources--cdn_loadbalancer--reference--group-009.md#canonical-2232022031111012-2123120231300022-0230012213330022-0013201210111203-0320113121001121-1221212111201023-1102322003200313-2102003100103220)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-009.md#canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2021023122003030-3032103120101121-2110000210020320-0120132233233113-3103022001101000-3033222000102221-3320231003123303-2001111023022230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203233203312223-3223022021000210-2012211122021201-2132032300330123-2332230301121131-0321011131103302-1003223123203023-3002103133103010"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers — append_headers / 331013102030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-009.md#canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--cdn_loadbalancer--reference--group-009.md#canonical-3000200131002020-3221332213012332-2210213230212022-3133311113302123-3310211110332232-2230130301130221-0020030110210203-3233023122232110)
- bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers

<a id="canonical-3222321320000133-0211121331230002-1100003132013313-2233230123013023-1100120111212232-0133112312001303-3332322332030020-2233100302322100"></a>

Type: `"object"`. single nested block, Optional.

Append flag mitigation headers to forwarded request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("auto_type_header_name",
    "inference_header_name")}
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
append_headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303300300302223-0022323021013212-1021212201102113-0011331120030332-3321130231113132-3030302001203313-3201113101302132-3100321123232023"></a>

## Direct properties — append_headers / 331013102030 / 3

<a id="canonical-1311312212020101-2123001310301013-1133112000033320-2212223222330012-2003210022100202-3311320233301332-1202111002223201-3321033110321312"></a>

<a id="canonical-0202231300102220-1031312112230133-3011223222013103-0003031001310113-3033311313210112-3311131122103303-0321101211132111-2330013010032301"></a>

## auto_type_header_name property — append_headers / 331013102030 / 4

Type: `"string"`. Optional.

Automation Type Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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

<a id="canonical-2233303033211200-3100031102033320-0131212001310210-1313302300320123-2233110103203010-2011313332101323-3321220120133321-2000220212131330"></a>

<a id="canonical-3133301103310203-0113223122221210-2221232123131313-0201033200203302-2133321003002320-1320301130221331-3030003002320310-3331311113003111"></a>

## inference_header_name property — append_headers / 331013102030 / 5

Type: `"string"`. Optional.

Inference Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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

<a id="canonical-1110002023213302-1031223223202220-2122311222133303-3331330221022200-0322211032203321-2031123312323320-1032011331122331-0111211221012232"></a>

## Next pages — append_headers / 331013102030 / 6

- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--cdn_loadbalancer--reference--group-009.md#canonical-3000200131002020-3221332213012332-2210213230212022-3133311113302123-3310211110332232-2230130301130221-0020030110210203-3233023122232110)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2232022031111012-2123120231300022-0230012213330022-0013201210111203-0320113121001121-1221212111201023-1102322003200313-2102003100103220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200200003322202-2032031213031010-0131001011310121-1200022313033131-2202131311223310-3203010301100201-3101231130322113-3202302302230121"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers — no_headers / 331101031302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-009.md#canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--cdn_loadbalancer--reference--group-009.md#canonical-3000200131002020-3221332213012332-2210213230212022-3133311113302123-3310211110332232-2230130301130221-0020030110210203-3233023122232110)
- bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers

<a id="canonical-0330230120000312-1033111321333311-3103332013330210-1020121331013122-3113100111013011-0113310230320033-2011102023120130-0333112201322221"></a>

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
no_headers = {}
```

<a id="canonical-3021300210233112-1310012232002021-3031130320103333-1300211210113021-2310012112223113-1101331032323230-3112113210313303-3102201020321130"></a>

## Direct properties — no_headers / 331101031302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303220020321001-3120200303302212-2022111211102223-2013121122231113-3102023330310231-2202130210132332-1231331231203100-2200212110021032"></a>

## Next pages — no_headers / 331101031302 / 4

- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--cdn_loadbalancer--reference--group-009.md#canonical-3000200131002020-3221332213012332-2210213230212022-3133311113302123-3310211110332232-2230130301130221-0020030110210203-3233023122232110)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1313313330213113-1313101230021132-1120203311013321-3211102113010322-1100103012220022-0323132213332203-2132220310032211-1113203122030233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002221312131033-3333201313333120-1210301010000200-0133100121221021-1233111130103113-0122023110323200-0201320123320101-0300330121303010"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.redirect — redirect / 313121202222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-009.md#canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101)
- bot_defense.policy.protected_app_endpoints.mitigation.redirect

<a id="canonical-2202121020212130-0031330113001033-2102111001103232-0203002112020112-2333030123320032-0300313132030111-0200332120220211-1200022100232033"></a>

Type: `"object"`. single nested block, Optional.

Redirect bot mitigation. Redirect request to a custom URI.

Upstream description:

Redirect request to a custom URI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("uri")}
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
redirect {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003233322300223-1001203213223033-3200321232220210-3001222013022020-1322110212321223-0213323003022123-2222120200023010-2230301301230231"></a>

## Direct properties — redirect / 313121202222 / 3

<a id="canonical-2201302121231000-0221002320320203-3213321312020231-3210101200130113-3222013022102233-3101012310220310-2312233112311003-2021133220202222"></a>

<a id="canonical-3321021221213110-0011133302012221-1212223303101133-2301133232212110-2111121223020211-1312231001213321-0021110332211111-1020323222203113"></a>

## URI property — redirect / 313121202222 / 4

Type: `"string"`. Optional.

URI location for redirect may be relative or absolute.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  }
}
```

<a id="canonical-1032233022101002-1133120130002313-2013201022212323-2200233010110310-3022003033303030-3313220120331310-2121033031020020-2113310211010133"></a>

## Next pages — redirect / 313121202222 / 5

- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-009.md#canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2320231113201121-2020300211032022-1102102103223120-0111130001113120-3301232221202322-3003200222210111-3331112021232103-2000212121221231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303000202333131-3203111110311110-1030200133022012-3120301201200101-1022230012000312-2211303201320110-0100110032210033-1023313312212331"></a>

## bot_defense.policy.protected_app_endpoints.mobile — mobile / 111303130211 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.mobile

<a id="canonical-1231131102032000-2023133131033223-0001210110013110-3023311211312033-1312101111111220-2300311321200122-3013130131331221-1211221330111101"></a>

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
mobile = {}
```

<a id="canonical-1112300312111110-2233302200021222-1031022010220233-1203233302311012-3002130212121232-3230231122311213-2332023222230113-1301202321231331"></a>

## Direct properties — mobile / 111303130211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312321221312211-1320101020111233-3003021123121032-1233021023101212-1222202211210110-0232023112030211-0100122232102133-2112331031303303"></a>

## Next pages — mobile / 111303130211 / 4

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2223313223123202-2233212130010311-3312012021323023-3002302323032331-3131221133330101-2202011230010131-3013113021220322-1223300321232230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030312002203003-3230003110122303-1313232021130122-0130303103103020-3003200330202230-3130032031332013-0022322331213233-0313101221313102"></a>

## bot_defense.policy.protected_app_endpoints.path — path / 032002113302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.path

<a id="canonical-2122100031203313-3232303022000332-2121223211032220-3331310221311232-3122231312211332-3212012122110122-1010003222310232-2001101033010013"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202033031023013-0110023011120302-3023201303223102-1221013112123301-2130233300131102-1321131300111231-3310113000223210-3332103213202013"></a>

## Direct properties — path / 032002113302 / 3

<a id="canonical-0113123112200103-3003330212200101-0023310322303303-1220331313003230-0312322111232202-1032200301002202-2020121131202102-2200022323303001"></a>

<a id="canonical-2133321001132033-3312030311113022-0221220321113102-1330002023130202-1211113031101231-1001213020202201-2222031201111023-0322032120230012"></a>

## path property — path / 032002113302 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1233212010212312-2022003132303231-1322322132311102-2012312202122311-0310123032200021-1102233201002312-0010311001031111-2212020320302100"></a>

<a id="canonical-0002133112031101-2110222212003131-1100312131312200-1220032300013230-1313311223010220-1013113223012210-0012201031102213-2101131000132313"></a>

## prefix property — path / 032002113302 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0011111310101223-3110232230120200-1301021322321031-2310301121111023-1220110133032032-0101123321320231-1312020001220321-0213030233222132"></a>

<a id="canonical-0232321302002311-1221031112033220-0310232313321100-2201210012302312-3300323301102330-3302103122112010-3213323110203032-1022131030310232"></a>

## regular expression property — path / 032002113302 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2212313210130103-1013123320330313-2023120020221130-3211200322313011-0033320321210111-1020330122031311-0111101012033130-0210010133312330"></a>

## Next pages — path / 032002113302 / 7

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2222021120300201-2020223023102321-3221032331210030-3302030312102232-3312302330311213-2130033013233102-3100313111202202-0211201310111123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303111211232203-0203223120331000-2003002123321003-2201320303310010-1032212321301230-1222003311030300-1320233013113003-1013100100020202"></a>

## bot_defense.policy.protected_app_endpoints.query_params — query_params / 323010121301 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.query_params

<a id="canonical-2031331220312022-2012111002021311-3201132002110103-3131111122312201-3222101102211000-1011010232030121-1310031023033311-3232111010102231"></a>

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

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-0311321320121210-3001101121011210-3203121001013121-0133000111320200-2120203131221111-0331112322223332-0033213010303011-0012110303310022"></a>

## Direct properties — query_params / 323010121301 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-009.md#canonical-3031013010131032-2002031330002310-2120213101302321-1120103322331222-3330232000112310-1220103013313211-2133022212030332-1231232102310222): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-009.md#canonical-3022002210001313-3002300223003212-0020330310222223-2202022123021200-2312031310222110-0231012020102021-1100100113000002-1312001201303103): complete subsection reference.

<a id="canonical-3332331232011003-1310023210311112-3121300033103301-3320120233330133-0230331333331003-3331000021321311-0122111022203320-0222321010112000"></a>

<a id="canonical-3231103120222133-3113231100021013-0300210020122131-1012231313202120-0112002111120333-2212321320302233-0202013203030230-0012122022202122"></a>

## invert_matcher property — query_params / 323010121301 / 4

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

- [item](resources--cdn_loadbalancer--reference--group-009.md#canonical-3211223223120223-2320210230221112-0301311210130120-1303000122132012-0021120313233110-3020001121123120-2021110232322001-0123113103320013): complete subsection reference.

<a id="canonical-3231333230023131-2113103220203301-2221232311000301-1021310023101011-2000031120030013-1000032300303333-3123332001311102-3020011330311003"></a>

<a id="canonical-3302100100120230-0103122110013330-3312210232133022-2301122231131123-1103031012230123-0123101031101302-2013112221012020-1200000322223100"></a>

## key property — query_params / 323010121301 / 5

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

<a id="canonical-3233332331320200-3310133210300333-3130310010212101-0012222221112211-1120021110033132-2132233333230310-1003132302201123-2002021221323131"></a>

## Next pages — query_params / 323010121301 / 6

- [bot_defense.policy.protected_app_endpoints.query_params.check_not_present](resources--cdn_loadbalancer--reference--group-009.md#canonical-3031013010131032-2002031330002310-2120213101302321-1120103322331222-3330232000112310-1220103013313211-2133022212030332-1231232102310222)
- [bot_defense.policy.protected_app_endpoints.query_params.check_present](resources--cdn_loadbalancer--reference--group-009.md#canonical-3022002210001313-3002300223003212-0020330310222223-2202022123021200-2312031310222110-0231012020102021-1100100113000002-1312001201303103)
- [bot_defense.policy.protected_app_endpoints.query_params.item](resources--cdn_loadbalancer--reference--group-009.md#canonical-3211223223120223-2320210230221112-0301311210130120-1303000122132012-0021120313233110-3020001121123120-2021110232322001-0123113103320013)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3031013010131032-2002031330002310-2120213101302321-1120103322331222-3330232000112310-1220103013313211-2133022212030332-1231232102310222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031112313323011-1320110133310002-3322201110230221-3010133322200103-2013120132113210-0121001013103003-2330011023021201-3123120211313210"></a>

## bot_defense.policy.protected_app_endpoints.query_params.check_not_present — check_not_present / 230321031111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-009.md#canonical-2222021120300201-2020223023102321-3221032331210030-3302030312102232-3312302330311213-2130033013233102-3100313111202202-0211201310111123)
- bot_defense.policy.protected_app_endpoints.query_params.check_not_present

<a id="canonical-0001030012111310-2332332113031230-2231223101120332-1000032223322132-0300112100122210-1310322210132021-0021200031202012-1100030133310021"></a>

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

<a id="canonical-2012201320133122-0210021112102133-3022220133130030-3221010132122211-1001312232232023-3010012120002301-3230323220330101-1220100330003121"></a>

## Direct properties — check_not_present / 230321031111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111302010033233-2310301001110220-1112231232320331-0233033031331200-3032033012012200-3123333301210101-1003110001001331-0133302331100330"></a>

## Next pages — check_not_present / 230321031111 / 4

- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-009.md#canonical-2222021120300201-2020223023102321-3221032331210030-3302030312102232-3312302330311213-2130033013233102-3100313111202202-0211201310111123)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3022002210001313-3002300223003212-0020330310222223-2202022123021200-2312031310222110-0231012020102021-1100100113000002-1312001201303103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231012320332323-3022202021332010-2233332323010320-2323231211212023-0210003021132223-0310010133233101-3120301010101211-3220020332131323"></a>

## bot_defense.policy.protected_app_endpoints.query_params.check_present — check_present / 230302231111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-009.md#canonical-2222021120300201-2020223023102321-3221032331210030-3302030312102232-3312302330311213-2130033013233102-3100313111202202-0211201310111123)
- bot_defense.policy.protected_app_endpoints.query_params.check_present

<a id="canonical-0032230330000320-1233212100213131-3101031003333333-1002000122131130-3313133323202133-1321322201121111-1002010322003032-3231100311300020"></a>

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

<a id="canonical-2001120111310332-1313213012311202-1323032032202330-0112032202012321-3302220102122132-3100333213130001-3010001020132330-2331203003021233"></a>

## Direct properties — check_present / 230302231111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210013213313333-1201013120030330-2312010322011020-1133231220320120-2210330323312213-1230312001321320-2312111101021122-0332111001231110"></a>

## Next pages — check_present / 230302231111 / 4

- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-009.md#canonical-2222021120300201-2020223023102321-3221032331210030-3302030312102232-3312302330311213-2130033013233102-3100313111202202-0211201310111123)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3211223223120223-2320210230221112-0301311210130120-1303000122132012-0021120313233110-3020001121123120-2021110232322001-0123113103320013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133113001011312-3211002102212310-1301031101201300-1000333020233232-1333131312101232-0203331301222232-3102201133121113-3300111230232103"></a>

## bot_defense.policy.protected_app_endpoints.query_params.item — item / 010320333032 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-009.md#canonical-2222021120300201-2020223023102321-3221032331210030-3302030312102232-3312302330311213-2130033013233102-3100313111202202-0211201310111123)
- bot_defense.policy.protected_app_endpoints.query_params.item

<a id="canonical-0130331303033033-0112111030002201-2133311023230313-3312300231223032-3201332302030012-0002200112011133-0230211312311000-3222331222003123"></a>

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

<a id="canonical-1302211301300221-1303333323011333-3330020201222222-1011113211311200-0322230310232010-1310301023302000-2103022200030331-0121310202101123"></a>

## Direct properties — item / 010320333032 / 3

<a id="canonical-3101310023133231-2101110022022331-0030221213230103-3021011032022022-3012131211232230-3000001330200301-0212120123210001-1023003111103113"></a>

<a id="canonical-2000212011103001-0132000301110130-2222112023131011-3103333113101233-0012322201202200-0011203331132003-0020312201121021-0303001212322231"></a>

## exact_values property — item / 010320333032 / 4

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

<a id="canonical-2022322023011001-0123033101203302-3220011222123332-0231313311302020-1123300000222001-3213101310320010-0113222103123123-0103102131202231"></a>

<a id="canonical-3323001212002120-1311300310221111-0033000230032010-3003032100210300-3310010223030333-0002010303010232-3313202233310232-1210320333322201"></a>

## regex_values property — item / 010320333032 / 5

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

<a id="canonical-2310212223302230-0102322010133003-3000313312030021-3221330230222310-3332210310313121-3123122330310000-3322213223120012-1100331000330230"></a>

<a id="canonical-0003222202133130-1002122303211101-1302003322122012-1011031003310013-2133201120322323-0222322212231120-2121232020023000-0300130303222323"></a>

## transformers property — item / 010320333032 / 6

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

<a id="canonical-0121133012111211-2011302103222310-1033111032322321-3120103313223232-3201212320331312-2221201033002312-3220233113322330-3201332313130103"></a>

## Next pages — item / 010320333032 / 7

- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-009.md#canonical-2222021120300201-2020223023102321-3221032331210030-3302030312102232-3312302330311213-2130033013233102-3100313111202202-0211201310111123)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1103122121102201-2310232002301010-2300221230313032-0023230013011103-2102023331202303-1133033313011223-2120000031313003-1320100210123030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221100000310032-0203032022203211-1321310020123310-2322012132233131-1211210231322103-0222011013012313-2103111310313110-1003130113312300"></a>

## bot_defense.policy.protected_app_endpoints.undefined_flow_label — undefined_flow_label / 323221112110 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.undefined_flow_label

<a id="canonical-2332300202221301-0312322022201032-2130002330010201-0020222130213231-0313101133223010-0230310022321032-0102233010310210-2322212120103102"></a>

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
undefined_flow_label = {}
```

<a id="canonical-3333220320302021-2100332310002000-3331013001021201-0020312200221130-1101131200111321-0023210330120333-2311003322300200-0330233323103220"></a>

## Direct properties — undefined_flow_label / 323221112110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330130212230330-3112110032300022-0200323010213121-3010321330232223-0331003102113322-1012323020110333-2310202200321312-0101120301103202"></a>

## Next pages — undefined_flow_label / 323221112110 / 4

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0113200330202102-0131023200210002-0120011001302123-3003313020231311-0232312022111032-0210130032310120-0320301113013123-0130001112002010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001021302212121-3001132333000130-2322213111222233-0001311130000303-0232220210310021-3231221031300010-2101113010310202-3111230222111113"></a>

## bot_defense.policy.protected_app_endpoints.web — web / 201333303000 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.web

<a id="canonical-1333322122333021-1132211132002112-0203230121320312-2311030103221131-0132323312230223-1323330323113030-0211113003322130-2331212302130323"></a>

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
web = {}
```

<a id="canonical-2021112220203321-1301132021111000-2032032302100332-2100102121111002-2103020033320001-1210100323322021-1023030303222230-1310112320010333"></a>

## Direct properties — web / 201333303000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020021101231111-1031320130010313-2202122033332313-1222112020232221-1000113332210310-3200011022101032-3030332013033302-3123002132032312"></a>

## Next pages — web / 201333303000 / 4

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1033303310333201-2322202210330231-1303323202223200-0121131311331201-2331312332001013-3100331211333112-3213330112110230-3230321011112100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031322110131132-0233121000120321-1332022132320213-1300001232021133-1320022100011011-0221133211003332-1322311033122223-0002233032210301"></a>

## bot_defense.policy.protected_app_endpoints.web_mobile — web_mobile / 032312333111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.web_mobile

<a id="canonical-3223133323220033-0003211312031211-0033100222000010-0122312112201212-3331020021020123-2131023310002231-0232312130133322-1123131020132301"></a>

Type: `"object"`. single nested block, Optional.

Web and Mobile traffic type. Web and Mobile traffic type.

Upstream description:

Web and Mobile traffic type.

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
web_mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232032213113212-1312120300232020-3233300320003220-1221021302120322-2321120122220102-3011110033030231-0101221103232221-2223130133321320"></a>

## Direct properties — web_mobile / 032312333111 / 3

<a id="canonical-2310021023310113-1132331323311031-3030220130000032-1101212220021232-2223323132320213-3233202103132333-0003332300111020-1000330322222031"></a>

<a id="canonical-2200301020313011-2223021321331122-1220313131122212-1131221022103132-1101302300100013-3331122302211320-1312330330222100-3133121232330210"></a>

## mobile_identifier property — web_mobile / 032312333111 / 4

Type: `"string"`. Optional.

\[Enum: HEADERS\] Mobile identifier type - HEADERS: Headers Headers. The only possible value is
\`HEADERS\`. Defaults to \`HEADERS\`.

Upstream description:

Mobile identifier type

&#8203;- HEADERS: Headers

Headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("HEADERS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "HEADERS",
  "enum": [
    "HEADERS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3121220320010013-1222132022133101-0333202333231132-2201012313000211-3021133321221022-2000030332032113-3102120133001032-2030210132311113"></a>

## Next pages — web_mobile / 032312333111 / 5

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0130021111111203-2032310323002002-0122300313003322-0110301130313122-2232111201113100-3201003313113133-0220131013102102-1223311220131210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312111213120210-1231013222130110-1110323223220200-2323323001113202-2020000301120332-3010301003331201-1302201211323133-1333033332021012"></a>

## captcha_challenge — captcha_challenge / 022332323102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- captcha_challenge

<a id="canonical-2300021221221121-3320002311031123-1112212300010133-0223113230131030-2120000032200201-2210002122002320-1101331010022310-0302201032122202"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: captcha\_challenge, enable\_challenge, js\_challenge, no\_challenge,
policy\_based\_challenge; Default: no\_challenge\] Enables loadbalancer to perform captcha challenge
Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that
pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is
configured to do Captcha Challenge, it will redirect..

Upstream description:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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

- [captcha_challenge](resources--cdn_loadbalancer--reference--group-009.md#canonical-2300021221221121-3320002311031123-1112212300010133-0223113230131030-2120000032200201-2210002122002320-1101331010022310-0302201032122202)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-1201312312020123-3230111211203203-0323103120212310-2002022312313232-2200320302313111-0002122232302200-2233200202311231-0212233210001213)
- [js_challenge](resources--cdn_loadbalancer--reference--group-011.md#canonical-1102301031312322-1122311201230301-2021213310112100-3011201130322112-3031202232133333-1230111121313021-2312100212122201-2010022231031120)
- [no_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-3122322231130001-0110033223312111-0211313102301220-1122203023011032-2300330122000333-2313023000331020-1121211003020210-1230333230203311)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-3110233220333200-3212201333323201-2020113221230221-3322221202220102-2132220310202212-2231113002233113-3013031220202301-3203221133222020)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
captcha_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303332213011011-2212312212232323-1102103321102011-3220203310011030-2013130220202102-2000113201300333-1121210302011300-0102120121011231"></a>

## Direct properties — captcha_challenge / 022332323102 / 3

<a id="canonical-2030232213300110-2222201010123011-0211022030231200-1213120220320031-3131301123122013-3312000012011302-3100203230120323-0212013233110010"></a>

<a id="canonical-1130302030301322-1020322013112321-1020221321320121-2323123330011102-2313200022211332-3110000010210302-1310130222010022-2032122110121112"></a>

## cookie_expiry property — captcha_challenge / 022332323102 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-2123032030231021-1112001101300133-2333311122312112-3320010130020110-0112101132313020-1030122222300213-1110332100322002-3202031012333330"></a>

<a id="canonical-1222312011203113-3312000110122300-2231030133301211-3222112221120302-2200222313320012-0232101202120312-0302013032113311-0103321221311002"></a>

## custom_page property — captcha_challenge / 022332323102 / 5

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format.

Upstream description:

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1021330120000312-3013003320111131-2132103032323100-0002110112022210-0302322310330032-3321323112231121-1330033320301211-2212223203322132"></a>

## Next pages — captcha_challenge / 022332323102 / 6

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302023020100313-3013011022202202-3120222013103102-1011310201002122-0300203222133321-3311230112010033-3020222311202112-0102322031320303"></a>

## client_side_defense — client_side_defense / 201220111123 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- client_side_defense

<a id="canonical-2101210301333033-0320000320311100-3201232120000003-0222313110211113-2323001021102312-3310231102100213-0131220011312311-3102233221201121"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: client\_side\_defense, disable\_client\_side\_defense; Default:
disable\_client\_side\_defense\] Defines various configuration OPTIONS for Client-Side Defense
Policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense Policy.

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

- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2101210301333033-0320000320311100-3201232120000003-0222313110211113-2323001021102312-3310231102100213-0131220011312311-3102233221201121)
- [disable_client_side_defense](resources--cdn_loadbalancer--reference--group-010.md#canonical-2323022203312322-3012211311100012-3031300132031113-3310223011233211-0120103103100233-1101032030002322-1221023301030013-2002330013332333)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
client_side_defense {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301311311200101-3010310021023333-2212322321201120-2203320103122222-0102311321010200-2301130232010102-1022002312200100-3211201323102020"></a>

## Direct properties — client_side_defense / 201220111123 / 3

- [policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011): complete subsection reference.

<a id="canonical-1030002022333311-3313102033322010-3333302211102103-2221001020100030-3320223023130011-3311300303002000-1003030010211323-2323110233232021"></a>

## Next pages — client_side_defense / 201220111123 / 4

- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212131022302101-1031123320012321-1113011121003112-2000130313311312-1200113201000101-3313212001311133-3231331100103230-0331211100311322"></a>

## client_side_defense.policy — policy / 010131332233 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- client_side_defense.policy

<a id="canonical-1133131220002112-3130113012120313-1302213332101112-0003131211102100-2111120030001303-0210030023313103-2020302210021301-3122230122021002"></a>

Type: `"object"`. single nested block, Optional.

Defines various configuration OPTIONS for Client-Side Defense policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

Terraform syntax:

```terraform
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110122012203120-1330213323120330-3010301201020313-2303223031211113-2123322013230102-1200312202113310-3120333131200323-3001221120210301"></a>

## Direct properties — policy / 010131332233 / 3

- [disable_js_insert](resources--cdn_loadbalancer--reference--group-009.md#canonical-2130300320210021-3202220131032010-1230100320310002-2323321113200213-1110220232032020-0311120013201303-1100011220313032-1002330301022113): complete subsection reference.

- [js_insert_all_pages](resources--cdn_loadbalancer--reference--group-009.md#canonical-2201311330313320-1232113110111220-2223231323030033-3320030202232020-2202301102020001-0011312003220012-2111022000112000-0111112132212221): complete subsection reference.

- [js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110): complete subsection reference.

- [js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303): complete subsection reference.

<a id="canonical-1300003232312201-2001032101013112-2132030320002300-1102333223233033-1202000133002201-3020330120333121-0113123112231001-1112100221333221"></a>

## Next pages — policy / 010131332233 / 4

- [client_side_defense.policy.disable_js_insert](resources--cdn_loadbalancer--reference--group-009.md#canonical-2130300320210021-3202220131032010-1230100320310002-2323321113200213-1110220232032020-0311120013201303-1100011220313032-1002330301022113)
- [client_side_defense.policy.js_insert_all_pages](resources--cdn_loadbalancer--reference--group-009.md#canonical-2201311330313320-1232113110111220-2223231323030033-3320030202232020-2202301102020001-0011312003220012-2111022000112000-0111112132212221)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2130300320210021-3202220131032010-1230100320310002-2323321113200213-1110220232032020-0311120013201303-1100011220313032-1002330301022113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230013212113232-2033132210123013-2123131230230032-0311100033032122-2022220332323222-0120202232213011-0331311132001201-1003023320022301"></a>

## client_side_defense.policy.disable_js_insert — disable_js_insert / 112112202200 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- client_side_defense.policy.disable_js_insert

<a id="canonical-0213112322220203-0310233202232221-2230313003232302-0210220121333122-2021302033003221-1320212013012002-2101210132200120-0312202223202032"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

<a id="canonical-0000123200110320-2223211211303311-0001301321332313-2203311132013013-2232231010203102-1111012323010120-1322313211201122-2032222000200102"></a>

## Direct properties — disable_js_insert / 112112202200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231322023020312-2121033031031131-1312333101231121-1313011223230320-2323313211211321-1110310122313223-0003321002310033-2111221130101003"></a>

## Next pages — disable_js_insert / 112112202200 / 4

- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2201311330313320-1232113110111220-2223231323030033-3320030202232020-2202301102020001-0011312003220012-2111022000112000-0111112132212221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211012321010133-0132231010133220-0303032313230013-3132022121132002-1202310222322220-3021030233213201-0301233202321201-3212200130022001"></a>

## client_side_defense.policy.js_insert_all_pages — js_insert_all_pages / 002131010030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- client_side_defense.policy.js_insert_all_pages

<a id="canonical-2121220033211002-0233333332013100-2322121200020033-0301232000001020-0222221001331131-0311021122100101-1021202002121202-3230322301211020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for js insert all pages.

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
js_insert_all_pages = {}
```

<a id="canonical-2002010303103230-3002210230021032-3332200121312221-1020103133013011-2232000011221202-0330032012212010-3023302332331211-3210232331313111"></a>

## Direct properties — js_insert_all_pages / 002131010030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011310101210112-2230123020100312-1010110013221200-3310112000223301-3301311031212123-3103311203211121-3333202020233132-2331302030221313"></a>

## Next pages — js_insert_all_pages / 002131010030 / 4

- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033200300102030-3123123333220130-3222101212101201-3130110232321210-0331132311133020-0012320212031331-0210230111210300-2222023221132230"></a>

## client_side_defense.policy.js_insert_all_pages_except — js_insert_all_pages_except / 223221230023 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- client_side_defense.policy.js_insert_all_pages_except

<a id="canonical-3022312310012312-3220213003200033-3023021110223323-3103300333103312-2212032213232033-0213313331020332-2110003200300032-1031112220030203"></a>

Type: `"object"`. single nested block, Optional.

Insert Client-Side Defense JavaScript in all pages with the exceptions.

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
js_insert_all_pages_except {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323303301113311-3123320213023202-2312321012002012-0011302131113201-3112013312033012-1320330213212033-2312102100000020-2331100011110213"></a>

## Direct properties — js_insert_all_pages_except / 223221230023 / 3

- [exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020): complete subsection reference.

<a id="canonical-1020220101311300-0002232331033130-0132000101202000-1222112213210210-1311203220020300-2102220103022223-1302133101331030-0300331210300231"></a>

## Next pages — js_insert_all_pages_except / 223221230023 / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101230201131123-0223323002100233-2310331123323123-1122312201231232-3331113313230200-0132331310001222-3311323313233320-3033021321232120"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list — exclude_list / 031210111121 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-2003312100300120-0031320302101323-1000212101123100-2303312101311020-2111332112301013-1012133113331203-3031333321011311-0132310233233201"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2320320111322111-3023103130302111-0121110213212212-2202310111023322-3212312022201012-3122333020323221-3302200210323202-0000002133322123"></a>

## Direct properties — exclude_list / 031210111121 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-0233222230112133-1021222222312133-0100031110310123-1131002012220222-1111213330011213-1121030313011130-2313231233320303-0331131310212310): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-1311320022132201-2220101011200313-0300021221032302-1230121331022101-3022330003311200-0313221232201302-2121010112231133-0000201220321011): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-009.md#canonical-2221033023222032-0121123212121332-1301303032103003-0202021211132110-1001022112100333-2301200022020012-3133013333312020-3213003301012111): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-009.md#canonical-2210033010113310-2031231013213022-3023300301231331-1322202021312311-3331311033121332-2130111022203022-3203102121120233-1021113311213012): complete subsection reference.

<a id="canonical-2311211322301200-2120101313103120-3232002321100313-2103033003100003-1000301023232312-0012223322023210-1003303320023131-1011212131013022"></a>

## Next pages — exclude_list / 031210111121 / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-0233222230112133-1021222222312133-0100031110310123-1131002012220222-1111213330011213-1121030313011130-2313231233320303-0331131310212310)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-1311320022132201-2220101011200313-0300021221032302-1230121331022101-3022330003311200-0313221232201302-2121010112231133-0000201220321011)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata](resources--cdn_loadbalancer--reference--group-009.md#canonical-2221033023222032-0121123212121332-1301303032103003-0202021211132110-1001022112100333-2301200022020012-3133013333312020-3213003301012111)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.path](resources--cdn_loadbalancer--reference--group-009.md#canonical-2210033010113310-2031231013213022-3023300301231331-1322202021312311-3331311033121332-2130111022203022-3203102121120233-1021113311213012)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0233222230112133-1021222222312133-0100031110310123-1131002012220222-1111213330011213-1121030313011130-2313231233320303-0331131310212310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323312223013322-3030102020303002-3203021031220010-2212220232131200-1133201332030301-2313110032130222-0201302201021001-2002102303111123"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain — any_domain / 011320202032 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-1130300003131313-0103122113331002-2213333311203221-3303012111011220-3302231130212120-0301333303222121-2010103211020110-0032103121111202"></a>

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
any_domain = {}
```

<a id="canonical-1213323121302200-0003201322201231-0333010223331312-2020122232201213-1033011330211321-2303231322030100-1222233201121032-2300110121211232"></a>

## Direct properties — any_domain / 011320202032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201320222211210-3111212033130213-3123013121111131-2233001232123103-2302033211232032-2012201111301101-3112321120330303-1121020000223123"></a>

## Next pages — any_domain / 011320202032 / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1311320022132201-2220101011200313-0300021221032302-1230121331022101-3022330003311200-0313221232201302-2121010112231133-0000201220321011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203030020333233-0232313330030203-1010212202301033-2111033201233333-3212201320022102-1021323330303121-1101001221333203-2112233232122210"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain — domain / 302203031132 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-3313022330020111-2133321333201211-0132202031003023-1000110030102213-0003222213232033-1013201112023021-1003012110231201-1320121312103110"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013223230013103-1101201220212221-1131101211130230-3000111103011302-2111320200320122-1233222100013210-0113031010322313-0300223311020000"></a>

## Direct properties — domain / 302203031132 / 3

<a id="canonical-3031232201300203-2213222003112000-0330323300122211-1020321223133111-3002112103002321-1313121021102302-3322201322221212-0321000021122133"></a>

<a id="canonical-2212330223310021-1303123201012221-3301313121032102-2300312221203230-1110331311013313-3330333300201011-1320122021300311-1003222131013013"></a>

## exact_value property — domain / 302203031132 / 4

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2222012312200333-1123101102132003-1103020130112032-0012133032130032-1323320201230230-3022220023313332-2011332303002232-3203320132300213"></a>

<a id="canonical-3332230030213010-1100203100102222-1130002301131111-2010310212021110-2031211111311110-1231111213122010-2210330332211122-3001311132013132"></a>

## regex_value property — domain / 302203031132 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3323212322200030-1210202321130001-3310120223012203-1311013121230220-3111030123012322-1013112213031310-0100100010032130-2330131210002022"></a>

<a id="canonical-3301112331233202-3032230003000232-0321323230003230-1230020312031310-0332223330133120-1121211320303101-2223213323102131-0322303321002011"></a>

## suffix_value property — domain / 302203031132 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0011031033320113-1031332001322231-1302113103020303-1011000233302002-0002330202332021-2012123301003122-1130232111131003-3111122012012300"></a>

## Next pages — domain / 302203031132 / 7

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2221033023222032-0121123212121332-1301303032103003-0202021211132110-1001022112100333-2301200022020012-3133013333312020-3213003301012111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113003100123221-1020320302331013-3012123233031212-1221013330013200-2020313213122322-2112212202213103-0010122032112300-1220101020133101"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata — metadata / 323202030300 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-1333312112123212-1230011203202302-1300030112311020-0103021120002232-1322213202002011-0121321230123112-3020303220312201-1313112323231110"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133302223322211-3223030313312022-3222003122200212-0133003231111223-2000300302213022-1312303222320030-1231211132311132-2003300200033210"></a>

## Direct properties — metadata / 323202030300 / 3

<a id="canonical-0231003011131203-2120020330021030-1300031011201033-1311030110200131-3213003130132302-3120301003311102-3100103020213020-1312320221131312"></a>

<a id="canonical-3131020010122203-1223320121300101-2303031303001202-2300022003231012-3210302313132332-1003010010123232-1021311033012111-0023113231311113"></a>

## description_spec property — metadata / 323202030300 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3211112033010001-1320031313211302-3123302330203201-1021200033202300-2320303231320311-3321030301201103-3230111332333332-2120111300012331"></a>

<a id="canonical-3333312001031001-1102201032123231-2322230011322101-2322220201002310-3032202233112132-3000010021301110-0001110001103311-0202002101111132"></a>

## name property — metadata / 323202030300 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0313232201333302-2223010003301211-1312003211022021-2333113131330112-2133212111222320-2110233233130231-0203002110322032-1000233313102033"></a>

## Next pages — metadata / 323202030300 / 6

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2210033010113310-2031231013213022-3023300301231331-1322202021312311-3331311033121332-2130111022203022-3203102121120233-1021113311213012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022132112003321-3013200112303313-2101120331201331-2230123313023120-3231133103022201-2310221232233332-3112231012010303-1031323100320323"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.path — path / 212120212301 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-1213131310333101-2322120110001030-0011132222013320-1210002122100011-2000033032201222-1022132323122210-3303110322101231-3321133123112011"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323323323123231-0011333301211313-2303033112313120-2213303220333133-1011311311100100-1220232110122111-1021331232300313-3203311110002112"></a>

## Direct properties — path / 212120212301 / 3

<a id="canonical-0311300300323311-2001233312101312-3222232222320030-3032123120121220-3233223303323203-0212122332301011-0101000223103100-2010223132021321"></a>

<a id="canonical-1210121220223002-0131231112212100-3202211323132302-0110112303233130-2202311020122112-3100301001033333-1123112100221001-2132233220121201"></a>

## path property — path / 212120212301 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3203213220001312-0022131212132032-2110002132100202-3103123233200203-2210202323233313-2213002213001013-1330122011210133-1011120321001201"></a>

<a id="canonical-3200223023312033-0102300311310031-1003133013122310-3102323220120120-3211122300200023-3010100112223113-3221111130310013-2023131020233322"></a>

## prefix property — path / 212120212301 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0112122122200121-3213100113323323-1220332211320223-0310322303113321-1201001313230131-1122022002221312-2231331312323001-3023233013013333"></a>

<a id="canonical-1101130232103203-3320013211101033-1200223321213130-0312230033220131-3010003023332022-2030121003230031-3020300323023333-3121322311033302"></a>

## regular expression property — path / 212120212301 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2000321310312113-1121123211203321-1013333022003103-2333300111320330-3021330113230130-1333032322033012-0230220102020230-2131111020211332"></a>

## Next pages — path / 212120212301 / 7

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000100322120130-1312331200233121-3320022002232300-1203101021112021-2322122223010122-0110323233000213-0323131320022120-0003103212101322"></a>

## client_side_defense.policy.js_insertion_rules — js_insertion_rules / 030010323312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- client_side_defense.policy.js_insertion_rules

<a id="canonical-0032123211201322-1203331031033202-0330112130003233-3131003101312202-3122010122211332-1011012231212020-3131233000000301-3121023303313023"></a>

Type: `"object"`. single nested block, Optional.

Defines custom JavaScript insertion rules for Client-Side Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Client-Side Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3031101111303321-1320133300123230-0220331000202201-1220220001221131-0333200221202122-2002332203100023-3200131332200003-1032030033013222"></a>

## Direct properties — js_insertion_rules / 030010323312 / 3

- [exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231): complete subsection reference.

- [rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331): complete subsection reference.

<a id="canonical-1132332203200233-3032223301223022-2010221303210132-3331033211023302-2033000102332332-1132311232313120-0213122200033100-0102231312120100"></a>

## Next pages — js_insertion_rules / 030010323312 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020231332002210-3213220322012332-1332020103013313-0211332011213302-0333233303301331-1312323023021002-2001133133332132-3202230002000031"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list — exclude_list / 220203033211 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- client_side_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-0033021021103030-3110320120300010-1122011303020221-0021330230023023-0322032222322102-0022031122332200-2113211301221130-3311233000001233"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002021200312230-0010211030211303-1103210321011122-1110302220331321-2203001300233321-1112011101223101-0200313310330220-2200121331202101"></a>

## Direct properties — exclude_list / 220203033211 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-2112322300003322-1231031202111323-2123300201312323-1221203333221323-0323111022202232-1302100100013210-0331231111021000-0003213221322332): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-2322123032311121-1222301003302223-1201000321312313-2021323231210223-0302012213331013-3302232021032223-0203300101322100-1202012031110121): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-009.md#canonical-0331132301132103-0021110300121031-2221103332210322-3200133303110302-3323201121112011-0023322210033132-0132213000113303-0013332122001320): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-009.md#canonical-0331030011303300-1311003002101020-1030002303213032-1331121301120132-2233212033130012-2300200020230001-0323010022110202-0002222120010302): complete subsection reference.

<a id="canonical-0222311230311023-3011010102101211-1031022311203030-0023120003200120-3333221202312232-1033302311011212-2213023222102100-3010330121302031"></a>

## Next pages — exclude_list / 220203033211 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list.any_domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-2112322300003322-1231031202111323-2123300201312323-1221203333221323-0323111022202232-1302100100013210-0331231111021000-0003213221322332)
- [client_side_defense.policy.js_insertion_rules.exclude_list.domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-2322123032311121-1222301003302223-1201000321312313-2021323231210223-0302012213331013-3302232021032223-0203300101322100-1202012031110121)
- [client_side_defense.policy.js_insertion_rules.exclude_list.metadata](resources--cdn_loadbalancer--reference--group-009.md#canonical-0331132301132103-0021110300121031-2221103332210322-3200133303110302-3323201121112011-0023322210033132-0132213000113303-0013332122001320)
- [client_side_defense.policy.js_insertion_rules.exclude_list.path](resources--cdn_loadbalancer--reference--group-009.md#canonical-0331030011303300-1311003002101020-1030002303213032-1331121301120132-2233212033130012-2300200020230001-0323010022110202-0002222120010302)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2112322300003322-1231031202111323-2123300201312323-1221203333221323-0323111022202232-1302100100013210-0331231111021000-0003213221322332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220030223202303-3033001021000110-2331330233202111-1101110330013132-0022232132002101-3100321123032123-3310110132201122-3230331212223330"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.any_domain — any_domain / 210100101220 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231)
- client_side_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-0122333333321202-1001111023302220-0220002310310012-2301212111233330-0230200112032231-1010331133013013-1031213113332110-2001202023322121"></a>

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
any_domain = {}
```

<a id="canonical-0323313210011211-1113012321132012-2221332103122313-2002213330113322-2201331311220201-0231222122221330-0022220313232102-3103311233120301"></a>

## Direct properties — any_domain / 210100101220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0300003030330032-0302233210101300-3013113332222233-2002130210030313-0112030330311110-3213001030001013-2002011023221022-2321201113322231"></a>

## Next pages — any_domain / 210100101220 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2322123032311121-1222301003302223-1201000321312313-2021323231210223-0302012213331013-3302232021032223-0203300101322100-1202012031110121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302002303310210-1103032213123233-1130301021031321-2220333101112332-1101123220203102-1120031200131120-0311201210302020-0133323013202012"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.domain — domain / 113021113033 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231)
- client_side_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-2302222233200123-2310123201220120-3322003123221331-0000110223010012-3332231200012221-1102331033031023-2022300330020331-1333320003113121"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2030103313012331-1320233210301300-3323130131233102-2223232333131131-0010310020110013-0323323313300012-0023301332221002-1112312113233320"></a>

## Direct properties — domain / 113021113033 / 3

<a id="canonical-2203132331100031-0232203032010003-0122203133231020-3212332012111201-0110322312311003-3033131012122100-2121230302330111-0202320120032012"></a>

<a id="canonical-2301001300300130-1212002331311103-2323313201313020-0133220033113230-2323310111203123-3130331321103103-2222312311222212-1121023032201023"></a>

## exact_value property — domain / 113021113033 / 4

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3003222111223032-0213020211121220-2213030201333032-0213121030310110-1021333130220020-2111132103030323-3001033203200222-1303001231233032"></a>

<a id="canonical-0223012100132112-3110010131113030-3330112120103102-0230012303200101-2323311010133103-1332203132302310-2111112123113110-2020022002122103"></a>

## regex_value property — domain / 113021113033 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3300302002310002-2321231110120132-1133332333231311-0322023123011033-1120302022332032-1033120023203013-3000223030231011-3221021300110103"></a>

<a id="canonical-2201123330122330-1200130233320112-2222322111010120-1210210130313200-3230302101232321-1030301201000301-1321112210332032-3220221001312233"></a>

## suffix_value property — domain / 113021113033 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0121303032133232-2032303333031010-1032021013321321-1212313030313120-3212100213011012-3320123110132020-3120333023311102-2323211022300103"></a>

## Next pages — domain / 113021113033 / 7

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0331132301132103-0021110300121031-2221103332210322-3200133303110302-3323201121112011-0023322210033132-0132213000113303-0013332122001320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012312303221202-0233013330130021-3233222002113122-0201301313210222-0032130103110330-2202323201213202-2003320102212103-2112131132021331"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.metadata — metadata / 102320221323 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231)
- client_side_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-2332121222032311-3323130233322121-0332222320222031-3032332302010130-0030110223021123-2230220222211030-3010113310301300-2333232211201212"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331303221020321-0110310033303113-3203321311201233-3010221122223233-0102121312031223-3120101220133112-0211102121323323-3000300113111221"></a>

## Direct properties — metadata / 102320221323 / 3

<a id="canonical-2033222001310103-2110021031213101-3022000320113220-1213322112013321-3232330212301131-1233023133230133-1100211230311233-1223023333100230"></a>

<a id="canonical-3211033222201200-2013013222333312-1301112021021333-3322113301302023-2120102211012232-3212323313232121-2131221122033302-2221330101321030"></a>

## description_spec property — metadata / 102320221323 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1130230332310103-3011303331030303-1001000002031021-0300232301222211-3223332102223212-3333031320123032-3102202230002030-2230133012313101"></a>

<a id="canonical-1201303123320012-2202302201220201-3313021333303210-3233231001011100-3312100322110323-3030020131230000-0231002320332300-0123123213110202"></a>

## name property — metadata / 102320221323 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-3203123132233033-3110231012222011-2233230113312303-0201130232131002-0222030000233133-0121321222033221-0111130221220302-2202233321103332"></a>

## Next pages — metadata / 102320221323 / 6

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0331030011303300-1311003002101020-1030002303213032-1331121301120132-2233212033130012-2300200020230001-0323010022110202-0002222120010302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201131033221011-3132102103000003-2321331301130310-2110013213033101-1300131220231230-1013333013332110-3032102013220213-0110203321113211"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.path — path / 112021023111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231)
- client_side_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-1010303203000310-3301211222303301-3230012010303132-3112021013023021-1312322301102113-3300103001100311-2200030311111211-2231313221333032"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011302030131203-3120133123303211-1200100203302130-2330322121233101-3122220203221103-0020333003333300-2213321010121121-2311133312110213"></a>

## Direct properties — path / 112021023111 / 3

<a id="canonical-0203023003023120-2130003023021122-3310032220022312-3222213330211301-0002010130232110-3123131112203001-1331010302001221-1331211220202031"></a>

<a id="canonical-2002123331112101-2231303022233230-1121122120201112-0132031222301303-0212130132100202-3101002200113210-2222210320002300-1013311113302333"></a>

## path property — path / 112021023111 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3213310002001320-3312302331123122-2201112010333330-2303233113320002-0330021321333001-1230002223233130-2133300002220131-3030220000011130"></a>

<a id="canonical-3333003333201123-0112312320301223-1333000232221113-1003111221012100-1312020013333300-3212221230113010-3131301313012210-3330210202113330"></a>

## prefix property — path / 112021023111 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2201010323011010-0001002002301220-2311313323201031-0221020032301132-0120301231323221-0221131303122301-0321031231022211-2112010001123122"></a>

<a id="canonical-2112202202330331-2202321131112023-1333320202011211-3233232200033030-1303211110331020-3100020103030031-2032000110211220-3133130103223231"></a>

## regular expression property — path / 112021023111 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3030030211103033-3030101220221032-1102222232001202-0300100221313112-3200221330111310-1203233322331302-2321212001332300-1200121200210222"></a>

## Next pages — path / 112021023111 / 7

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001112322322232-3003020031033302-3130120011113003-3221100000211100-2023320003013331-3002112122301122-0100232333230201-3333030103310012"></a>

## client_side_defense.policy.js_insertion_rules.rules — rules / 212211030311 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- client_side_defense.policy.js_insertion_rules.rules

<a id="canonical-2113231100130120-2203032320230203-3232002223231003-2200321232311320-0013032321302030-0201220321103001-3031003300230001-0211321332131013"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Client-Side Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122311321103022-3130313031230032-1201002202333022-0133213111310133-2002121301101101-1021033133303311-3013231101301101-2311022331100302"></a>

## Direct properties — rules / 212211030311 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-2110102120002011-3132103223032223-1223230211332301-0333000212113300-3111011223123300-3131320001210310-3311132323121032-0023310010322003): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-0022003231313222-1230033102113003-1301032020211022-1130332002032233-2022100333320303-1200201202213233-0120131200011211-0321023332212020): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-009.md#canonical-3222013222323100-0313302221020131-1021202113310121-2121123020112312-3130121112230031-0123000322121111-2221112013302112-0012130222222213): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-009.md#canonical-2003333022203031-1121112303022232-2331322031223212-1123103020213010-1210100002023102-2211202302302300-1200131301023033-1311100233200000): complete subsection reference.

<a id="canonical-2032320302101101-0121211313101332-1020303300320013-2011232132210331-0200212232022113-3020312133021321-0211000321030001-2123002203331000"></a>

## Next pages — rules / 212211030311 / 4

- [client_side_defense.policy.js_insertion_rules.rules.any_domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-2110102120002011-3132103223032223-1223230211332301-0333000212113300-3111011223123300-3131320001210310-3311132323121032-0023310010322003)
- [client_side_defense.policy.js_insertion_rules.rules.domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-0022003231313222-1230033102113003-1301032020211022-1130332002032233-2022100333320303-1200201202213233-0120131200011211-0321023332212020)
- [client_side_defense.policy.js_insertion_rules.rules.metadata](resources--cdn_loadbalancer--reference--group-009.md#canonical-3222013222323100-0313302221020131-1021202113310121-2121123020112312-3130121112230031-0123000322121111-2221112013302112-0012130222222213)
- [client_side_defense.policy.js_insertion_rules.rules.path](resources--cdn_loadbalancer--reference--group-009.md#canonical-2003333022203031-1121112303022232-2331322031223212-1123103020213010-1210100002023102-2211202302302300-1200131301023033-1311100233200000)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2110102120002011-3132103223032223-1223230211332301-0333000212113300-3111011223123300-3131320001210310-3311132323121032-0023310010322003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021330123113133-0200032131033220-1032320133023002-0100123220301210-3020320103012321-0100103112322012-0201122312111003-0223221210300011"></a>

## client_side_defense.policy.js_insertion_rules.rules.any_domain — any_domain / 113322202210 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331)
- client_side_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-3003200331231231-2033231203133300-0103201222320033-3221212323302222-3023203202001131-2221112210103023-3313011032300332-0031313023312031"></a>

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
any_domain = {}
```

<a id="canonical-2023001310321213-3102011021012133-0331032112320020-0111103321121112-3300323113012221-3300021203113210-3120112212131331-0113223110311013"></a>

## Direct properties — any_domain / 113322202210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211030031333202-2302230013013201-1000212333111101-1131021000020120-3202331311022322-2123233000121303-2021322021220003-3123121203301232"></a>

## Next pages — any_domain / 113322202210 / 4

- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0022003231313222-1230033102113003-1301032020211022-1130332002032233-2022100333320303-1200201202213233-0120131200011211-0321023332212020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122022303001211-2213120033000303-0331211321232002-2230331222002233-3032113333322120-2232011021003231-0223332121111211-0233132032002302"></a>

## client_side_defense.policy.js_insertion_rules.rules.domain — domain / 323013302103 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331)
- client_side_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-3112032003220220-0132023321312013-2332222302033200-0010303110332231-2231300102232203-1021202231001000-2231322310032021-2112222012200133"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033031120321032-3031020133310121-3202302232320103-1102031230313303-3311213113111220-0223231131031120-0020101223033022-1211120332121020"></a>

## Direct properties — domain / 323013302103 / 3

<a id="canonical-0211201333031123-0121202332131233-2121302130301001-1130323013212230-0102020133230221-2023331300110100-0031331001101210-1033200210113230"></a>

<a id="canonical-3321303313220200-2031010232132012-0032120001312213-3321213111202313-1301222101332012-2002123312220303-1232133223221000-2132100331101330"></a>

## exact_value property — domain / 323013302103 / 4

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2232121011202312-3031232013232020-3211002133032231-2103222322002321-1320201330333333-2213122101010112-2211202202301130-2322203101200112"></a>

<a id="canonical-1333122003111003-0000011311112201-3003020213332311-0331211300122100-0122230313013000-1313213301130021-2102331020312023-1121002112200333"></a>

## regex_value property — domain / 323013302103 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2202213302301200-2110113313201033-3111213303323333-0132200310113022-1001323102221130-3112222001030202-1221131132031102-0320130222113032"></a>

<a id="canonical-0321323131231102-0221232112313302-2011020130102002-2103022132130222-2013210000130300-0013000113332221-1323012230120120-3113332223020331"></a>

## suffix_value property — domain / 323013302103 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0231331000001330-1332321033203010-3120122310201020-3211123303332120-0130333212310223-0112000111033030-3333112102021013-2312213132313102"></a>

## Next pages — domain / 323013302103 / 7

- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3222013222323100-0313302221020131-1021202113310121-2121123020112312-3130121112230031-0123000322121111-2221112013302112-0012130222222213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100110233022010-2302310332022130-0223332030033201-2202121223003223-1223021001030202-0013300123012021-0331330113133222-3010011130133111"></a>

## client_side_defense.policy.js_insertion_rules.rules.metadata — metadata / 230132301303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331)
- client_side_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-3112211113001233-3122232313020030-1120300233221102-3021303000203333-3120013120103001-3030021123330200-1302100331132201-2330120022301010"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230221003123130-0103121112202300-2200303322010230-0222001231212330-0220300103113101-3323112331331211-0020232333102203-3130033111033302"></a>

## Direct properties — metadata / 230132301303 / 3

<a id="canonical-1033200021301103-3003320100331232-1300302232221120-3311120032212030-1120033332130132-3322100111202022-2200021220120101-2200023222223231"></a>

<a id="canonical-3310323301113212-1221322222132012-0110023131210331-3133122103331030-1100033322102110-0331320303201301-0301213030232230-0300320313023101"></a>

## description_spec property — metadata / 230132301303 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-0103011023303321-1133010032021202-1102310211010123-3123032112230002-1002022310200131-2202213120333200-0313120000132221-3003022201312303"></a>

<a id="canonical-0033122023211211-2022010000200003-1131303120111033-1002311223210101-3101211311223110-2123113333230122-2323312302003032-0002002002021013"></a>

## name property — metadata / 230132301303 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-1333032120311132-0222130331333001-1231231211203020-0100230122110112-1021011232311002-0233103022230222-3111300112323310-2033223020122202"></a>

## Next pages — metadata / 230132301303 / 6

- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2003333022203031-1121112303022232-2331322031223212-1123103020213010-1210100002023102-2211202302302300-1200131301023033-1311100233200000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010001131112332-1301233031132310-3232120003001121-1132311011300130-1130203130130313-0311333112110003-0302022310102021-3030233203211022"></a>

## client_side_defense.policy.js_insertion_rules.rules.path — path / 232021013211 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331)
- client_side_defense.policy.js_insertion_rules.rules.path

<a id="canonical-1022320212001002-1331132122202203-3211012120310110-2210101300113101-1311131032210103-0301000203301330-3002310211010010-2033001313022330"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132300120002010-0222320331102221-3012211022103322-3313310320112103-2221021011331112-0121311022103131-0310322330022323-3033211133013130"></a>

## Direct properties — path / 232021013211 / 3

<a id="canonical-1132131031330031-1101201031101132-2212101133031113-0232032010300313-3131103011200111-1020101230030033-1303313323203000-0231202311103310"></a>

<a id="canonical-0022333100022110-2110002032030022-1310032100302230-3330222020312103-0031111323030110-2033321013100102-3002123203323211-2021012011231300"></a>

## path property — path / 232021013211 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3212132311111110-2122301310332323-1012000123113003-3121121201311130-1221210231210222-2220031010112121-3001032230230201-0321232221031033"></a>

<a id="canonical-2302202203101020-0131221030333100-1103103231113000-0022301322222013-3313012203303120-3330112101303312-1232111231021312-0302131320003030"></a>

## prefix property — path / 232021013211 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0223120020230003-3130032322033301-2200012333033200-1203213001000013-1233122022020310-1303010321312330-0301332130331212-0031232022200021"></a>

<a id="canonical-0021212300100033-0200123210031120-1033223220222212-0010220103030331-2221032010330332-1303132220301132-2101332031310120-0233201010331203"></a>

## regular expression property — path / 232021013211 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2301333222202303-1021123132231001-3003300212002112-0320133331310113-1133233311232231-3102202202110113-1211120300302022-3021023100322200"></a>

## Next pages — path / 232021013211 / 7

- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1201021310323311-1300002333101322-1212100020212023-2033122312102030-0331223230312000-1222113322200232-2203221130113100-0233323002331212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003333131010311-1110110301102303-1220010102121121-2123232031001200-0111030302112302-3112123100232322-3300133102113001-0033030232010122"></a>

## cors_policy — cors_policy / 112231302233 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- cors_policy

<a id="canonical-0132123231310220-2230102001020102-2012331000011001-0221223011133020-3322112130010013-1113030322330222-1103030313010023-0301201112010201"></a>

Type: `"object"`. single nested block, Optional.

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence. An example of an Cross origin HTTP request GET
/resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel macOS
X 10.5..

Upstream description:

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence.

An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel macOS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130
Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8
Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset:
ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive Referrer:
http&#58;//foo.example/examples/access-control/simplexsinvocation.HTML Origin:
http&#58;//foo.example

HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61
Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive
Transfer-Encoding: chunked Content-Type: application/XML

An example for cross origin HTTP OPTIONS request with Access-Control-Request-\* header

OPTIONS /resources/POST-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel
macOS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept:
text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8 Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive
Origin: http&#58;//foo.example Access-Control-Request-Method: POST Access-Control-Request-Headers:
X-PINGOTHER, Content-Type

HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http&#58;//foo.example Access-Control-Allow-Methods: POST, GET, OPTIONS
Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary:
Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive.

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
cors_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2022033230232033-0100323330202132-2031221103210300-0130232100031330-0330233020201111-3110130023003121-3220103331321110-2231301330200011"></a>

## Direct properties — cors_policy / 112231302233 / 3

<a id="canonical-0022023213000031-2001030301003112-2130233302030020-0113130132110221-3321122023230220-1310033330112322-1312002301220010-0002221201320333"></a>

<a id="canonical-2133321012033130-3213000213211302-2133330300022011-0011201321303221-1131220323211021-3212023332131120-2210030302331132-2021000332231110"></a>

## allow_credentials property — cors_policy / 112231302233 / 4

Type: `"bool"`. Optional.

Specifies whether the resource allows credentials.

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

<a id="canonical-2332033032033223-1323101102213103-3120211022323110-0220300122211003-1111302321102232-2322012301303110-0033023112100012-1202310000320211"></a>

<a id="canonical-3311311210101232-2310202113032202-3200321301313201-0310022311210012-1102101002233301-3032020330023131-3320302013023300-3030010010113232"></a>

## allow_headers property — cors_policy / 112231302233 / 5

Type: `"string"`. Optional.

Specifies the content for the access-control-allow-headers header.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0100203323233313-1232012133331113-0331122302332102-2030333100201213-1101311020000211-3132103300010103-1220132321300131-0320311121133331"></a>

<a id="canonical-2103021312032302-3201003112333000-1021222101203020-3011321311033302-3223331212300321-0212220331201331-3121323032311300-1311021231303131"></a>

## allow_methods property — cors_policy / 112231302233 / 6

Type: `"string"`. Optional.

Specifies the content for the access-control-allow-methods header.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-2100000121033132-3122202020033131-0131102200110313-2201230222120231-1300321113333233-2031202222122130-0333213131303003-3003121011110330"></a>

<a id="canonical-2012100103132010-3232130213202200-3302232303031101-0113302200310231-0203230111101020-3132201020312231-2012232213330311-1221030313231222"></a>

## allow_origin property — cors_policy / 112231302233 / 7

Type: `["list", "string"]`. Optional.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Upstream description:

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2331202031132032-2012022231031320-2211131022123012-3301200333231031-1321220203011213-2203101011302121-0123330112330211-2031113022023112"></a>

<a id="canonical-3103031230132212-1002212200102220-1021011212203203-1000003311330211-3012023111033300-0221133221100220-2000122330113020-1311003003011221"></a>

## allow_origin_regex property — cors_policy / 112231302233 / 8

Type: `["list", "string"]`. Optional.

Specifies regular expression patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Upstream description:

Specifies regular expression patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1021230203102223-2210120122110211-0331332022123131-3102202030233330-2113012110122203-2011200333011011-2000200332231322-1210020210121312"></a>

<a id="canonical-2200132113311201-0103132101201221-1030131202221012-2320021323112230-1031331022321302-1213321331001330-1231033100122022-3321220232331312"></a>

## disabled property — cors_policy / 112231302233 / 9

Type: `"bool"`. Optional.

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="canonical-1010213102320233-3030321001023322-1200331130322103-2001333132002201-1021223332230120-3110230302330123-0331000230112130-1031033201322222"></a>

<a id="canonical-0202030200100201-3323002300323212-1103033223203232-3030030100033211-1223323120111200-3322001223331213-2132201312330031-3013313222210031"></a>

## expose_headers property — cors_policy / 112231302233 / 10

Type: `"string"`. Optional.

Specifies the content for the access-control-expose-headers header.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1301000120302231-1011121300120321-1312331230210133-0111301320221222-2112321303102113-2333110133202100-3111200223333223-1332323210020011"></a>

<a id="canonical-0302303132112211-2303030001313011-0020301313123301-0303301210132030-2103203333220303-2122300303203113-0210232000002330-1121222030100131"></a>

## maximum_age property — cors_policy / 112231302233 / 11

Type: `"number"`. Optional.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Upstream description:

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(-1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": -1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  }
}
```

<a id="canonical-3030301010011130-1231231323212301-3111130220313300-1010031133303232-3211122100311221-2321101012131001-3222111113020303-3002211330322310"></a>

## Next pages — cors_policy / 112231302233 / 12

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1000130300033331-3002211233300130-2313130100130200-1302332133032013-0013323333213201-3223012120111001-2100331133200101-2301200003031301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000120320012211-1231301111102010-1213201220313030-3322203321332331-2102333300301231-0222132101233220-1013310232231111-2101121111103013"></a>

## csrf_policy — csrf_policy / 132203110331 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- csrf_policy

<a id="canonical-1231222033032220-2322020333233302-3212230323212330-2001210011133312-1302002032221112-1111212100012202-3030031103121313-0101331103211120"></a>

Type: `"object"`. single nested block, Optional.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host. 1. The origin that caused the user
agent..

Upstream description:

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host.

&#8203;1. The origin that caused the user agent to issue the request (source origin). &#8203;2. The
origin that the request is going to (target origin). When the policy evaluating a request, it
ensures both pieces of information are present and compare their values. If the source origin is
missing or origins do not match the request is rejected. The exception to this being if the
source-origin has been added to they policy as valid. Because CSRF attacks specifically target
state-changing requests, the policy only acts on the HTTP requests that have state-changing method
(PUT,POST, etc.).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "custom_domain_list"),
  validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "disabled"),
  validators.ConflictingObjectAttributes("custom_domain_list",
    "disabled")}
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
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

Terraform syntax:

```terraform
csrf_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032233303113220-0001103333303212-0002301333132310-0302222102122123-0002332323130022-2110220232100330-1313232001331303-1110120020233101"></a>

## Direct properties — csrf_policy / 132203110331 / 3

- [all_load_balancer_domains](resources--cdn_loadbalancer--reference--group-009.md#canonical-3311132011232321-3011321303100313-0121230222211323-3302032233300132-2212023303132123-0030201203101312-1103033012121111-3013232011311211): complete subsection reference.

- [custom_domain_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-0213202201233021-0103110200200122-2030230232330203-3212322123021201-3121030303212001-2331012333101230-2210223032333301-1102301030102112): complete subsection reference.

- [disabled](resources--cdn_loadbalancer--reference--group-009.md#canonical-3133310213001213-0202133032123301-3120123112332301-2202202021033120-2101030113200103-0322303110010301-1131320013123210-2110113132201013): complete subsection reference.

<a id="canonical-0130221112211202-1132330113102211-0202213303020210-0022212210032001-2220223233332002-3032033202033201-2203021323232121-1201323303303303"></a>

## Next pages — csrf_policy / 132203110331 / 4

- [csrf_policy.all_load_balancer_domains](resources--cdn_loadbalancer--reference--group-009.md#canonical-3311132011232321-3011321303100313-0121230222211323-3302032233300132-2212023303132123-0030201203101312-1103033012121111-3013232011311211)
- [csrf_policy.custom_domain_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-0213202201233021-0103110200200122-2030230232330203-3212322123021201-3121030303212001-2331012333101230-2210223032333301-1102301030102112)
- [csrf_policy.disabled](resources--cdn_loadbalancer--reference--group-009.md#canonical-3133310213001213-0202133032123301-3120123112332301-2202202021033120-2101030113200103-0322303110010301-1131320013123210-2110113132201013)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3311132011232321-3011321303100313-0121230222211323-3302032233300132-2212023303132123-0030201203101312-1103033012121111-3013232011311211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011120330202030-2120113130323201-0333122300021223-0230121213103302-3000013211202122-3322302300331031-3311212200112320-3201030312303212"></a>

## csrf_policy.all_load_balancer_domains — all_load_balancer_domains / 121122230201 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [csrf_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-1000130300033331-3002211233300130-2313130100130200-1302332133032013-0013323333213201-3223012120111001-2100331133200101-2301200003031301)
- csrf_policy.all_load_balancer_domains

<a id="canonical-0323213320011002-0021333221121313-3031011113033020-3132302130023100-3331031122021100-0311023132233330-2311333200033032-3303300113130131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all load balancer domains.

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
all_load_balancer_domains = {}
```

<a id="canonical-3001232021202332-2112233232332311-3121133111202102-1010232202221322-1332202300213111-0223111112030220-1031010110201230-0023000233321002"></a>

## Direct properties — all_load_balancer_domains / 121122230201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220033301132213-2302120202112203-2310102303300013-0232232231111113-0001202320221313-0230023022230022-2102223020031300-2210230013323313"></a>

## Next pages — all_load_balancer_domains / 121122230201 / 4

- [csrf_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-1000130300033331-3002211233300130-2313130100130200-1302332133032013-0013323333213201-3223012120111001-2100331133200101-2301200003031301)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0213202201233021-0103110200200122-2030230232330203-3212322123021201-3121030303212001-2331012333101230-2210223032333301-1102301030102112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203233323220221-2320222221013100-2013320021222000-0103130230322130-1320332200012311-3021320112233022-3101002221331233-1213133212233123"></a>

## csrf_policy.custom_domain_list — custom_domain_list / 302331203100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [csrf_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-1000130300033331-3002211233300130-2313130100130200-1302332133032013-0013323333213201-3223012120111001-2100331133200101-2301200003031301)
- csrf_policy.custom_domain_list

<a id="canonical-2130002011211003-1331130323232231-3320112001031123-2213302013232313-3031110130102112-1233023113112130-0000110300202133-3120110223122300"></a>

Type: `"object"`. single nested block, Optional.

List of domain names used for Host header matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
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
custom_domain_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010313023320313-0323323120110031-0121030310013012-3203322030032002-1000332032211033-1132323110223110-3010222033103312-2230322010323333"></a>

## Direct properties — custom_domain_list / 302331203100 / 3

<a id="canonical-3110312123011331-0112122333032210-0211011011000021-3023312032212011-3203011221013032-2302033100300301-0201203310213003-2202213023020213"></a>

<a id="canonical-3111022132221030-1001213132313312-1200232321303233-2110311322112022-2312231323203313-2023201101102212-2330120310232220-2222100333033233"></a>

## domains property — custom_domain_list / 302331203100 / 4

Type: `["list", "string"]`. Optional.

List of domain names that will be matched to loadbalancer. These domains are not used for SNI match.
Wildcard names are supported in the suffix or prefix form.

Upstream description:

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2123311022123110-2031013212103223-0012100102321013-1131012131013203-2133220331202302-2012113323320211-0223323210311333-0020333201232211"></a>

## Next pages — custom_domain_list / 302331203100 / 5

- [csrf_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-1000130300033331-3002211233300130-2313130100130200-1302332133032013-0013323333213201-3223012120111001-2100331133200101-2301200003031301)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3133310213001213-0202133032123301-3120123112332301-2202202021033120-2101030113200103-0322303110010301-1131320013123210-2110113132201013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110332112321201-2312223202321123-1333311022230331-3012101303320012-3013233311312313-0202030223001103-1223032322112311-2233220332211123"></a>

## csrf_policy.disabled — disabled / 103201001131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [csrf_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-1000130300033331-3002211233300130-2313130100130200-1302332133032013-0013323333213201-3223012120111001-2100331133200101-2301200003031301)
- csrf_policy.disabled

<a id="canonical-3223333013113013-3012312102113001-2011211223333030-2101120321031200-3033231332301112-2011133001003122-1010032223000012-2200032203320120"></a>

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
disabled = {}
```

<a id="canonical-0102322322203020-0112000013321311-0233030030120111-3132121000301013-2121202030223003-3003121311020011-2300113220003330-2023331323122322"></a>

## Direct properties — disabled / 103201001131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233102133022110-0102211133003211-0203110122320220-0303302202223111-3011213233133013-1330201103221102-0111030033001130-3231012310101100"></a>

## Next pages — disabled / 103201001131 / 4

- [csrf_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-1000130300033331-3002211233300130-2313130100130200-1302332133032013-0013323333213201-3223012120111001-2100331133200101-2301200003031301)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2210332032110200-0210201322313110-1131031023311132-0213211002011103-0233222021320322-2330011210220101-1033313321031003-0323003302010310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231201012033022-2021032311223013-0330302021303331-2310101223102321-1110333023111110-1123112230001312-2320222001132102-0001032111221103"></a>

## custom_cache_rule — custom_cache_rule / 030131010020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- custom_cache_rule

<a id="canonical-1002233320231332-3203021211001213-1021010122022123-3100121021131012-2033012002133102-2000110310333000-0020013303232003-0313121100332232"></a>

Type: `"object"`. single nested block, Optional.

Custom Cache Rules. Caching policies for CDN.

Upstream description:

Caching policies for CDN.

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
custom_cache_rule {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320330111121033-0310002032001221-3122312123002332-1300000100021011-3111120212102323-2111020032000221-3032301213010132-0020232320211311"></a>

## Direct properties — custom_cache_rule / 030131010020 / 3

- [cdn_cache_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-3233131223322012-0200133201103211-2200131122332302-2332122101113013-3011230210012221-2231010010110130-2122212100210230-1311120330120300): complete subsection reference.

<a id="canonical-1010031200100000-0112311222311030-0323130011102313-2232132111221321-3213101232231331-2021001131100303-0323313331230131-2132003130001210"></a>

## Next pages — custom_cache_rule / 030131010020 / 4

- [custom_cache_rule.cdn_cache_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-3233131223322012-0200133201103211-2200131122332302-2332122101113013-3011230210012221-2231010010110130-2122212100210230-1311120330120300)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3233131223322012-0200133201103211-2200131122332302-2332122101113013-3011230210012221-2231010010110130-2122212100210230-1311120330120300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013132323010132-1121012223130100-3002111232121133-3311100320302322-2331110032030233-1331132030301333-1011101102323200-1233300101320010"></a>

## custom_cache_rule.cdn_cache_rules — cdn_cache_rules / 022030112212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [custom_cache_rule](resources--cdn_loadbalancer--reference--group-009.md#canonical-2210332032110200-0210201322313110-1131031023311132-0213211002011103-0233222021320322-2330011210220101-1033313321031003-0323003302010310)
- custom_cache_rule.cdn_cache_rules

<a id="canonical-3123311213201213-3123301233333232-1122033112101303-0020321130001230-0322223100132232-1202231032222131-2202100303333203-1310120022220120"></a>

Type: `"object"`. list nested block, Optional.

Reference to CDN Cache Rule configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
cdn_cache_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001121310032000-0223323232030102-0003301232203121-0222313132133100-1213022322323222-1100031022003111-3333130330031321-1203233331113311"></a>

## Direct properties — cdn_cache_rules / 022030112212 / 3

<a id="canonical-3003301112211301-3120203301303103-0102203223111012-3002021330113311-0230232031133212-3111212101020103-1011131232221030-0220033100000031"></a>

<a id="canonical-3332300112300222-3312123031223212-2310302223313301-3220202311121201-2300213122313121-1212030101301331-2312301323023113-3302312302322011"></a>

## name property — cdn_cache_rules / 022030112212 / 4

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

<a id="canonical-0132210303121312-1303203323112201-1111203011110111-3303103033022203-0022233202112022-0003033032102011-3121211033213000-0220210013300003"></a>

<a id="canonical-3000213020010232-1211202131203101-2212103000012103-1330330003133031-3221213020303320-2131030113322311-3030300312000202-2332121113031021"></a>

## namespace property — cdn_cache_rules / 022030112212 / 5

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

<a id="canonical-1332133323121021-0212123220223302-3033112330232020-2030011101322131-3120101102202202-2230202232333302-3220002210000301-0100333311232033"></a>

<a id="canonical-2032131301310113-0303101301032010-1022023301000213-1001013212131300-1102221012021230-0003120233110213-0013323003311032-3321031200320311"></a>

## tenant property — cdn_cache_rules / 022030112212 / 6

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
