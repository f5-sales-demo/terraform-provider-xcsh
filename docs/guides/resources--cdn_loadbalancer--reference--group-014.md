---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-0222010203333013-2300122132031123-1103133030020130-1233213103011120-1333102102133011-1010130011313221-3131322120012332-1232011202232131"></a>

## transformers property — item / 313103312103 / 6

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

<a id="canonical-3002222201012013-1022233233002322-2000232001002233-1332313011032132-2131321103312003-2103202221221211-0223012102211221-1331012002013121"></a>

## Next pages — item / 313103312103 / 7

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-3313200200332201-1200210121121211-3220003300020012-2002110001123113-3233122133221310-0212323322101310-1333012331212122-2002300120333130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0032332002211123-1202012303111223-1300130212323212-3131233302210131-2031312311130100-1303223302131323-2102102310301121-2010003203213221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332210113112200-0023023021110221-2230322030301112-0123233001033012-2330131103023100-2130111230021023-2222113213110302-1031013212130021"></a>

## policy_based_challenge.rule_list.rules.spec.disable_challenge — disable_challenge / 012030313030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.disable_challenge

<a id="canonical-0303000332313321-3320212331332232-0201122020300120-3211123012103021-2230320023003122-3012011010311201-1202213231130123-1111230210010233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable challenge.

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
disable_challenge = {}
```

<a id="canonical-0313233030233320-0222110133021030-0130001121121200-1000231030101131-2103231320233122-3103332200122222-2003211321132220-0201200223212100"></a>

## Direct properties — disable_challenge / 012030313030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123102331331220-1211021000133322-1021200320012222-0011030110222101-0112200031122011-1001223012203000-0132011231211001-3321223011100110"></a>

## Next pages — disable_challenge / 012030313030 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1031330331223002-1133112231311120-0010033010221011-3200232312203013-1321321213331110-3211120301031200-0232002103001122-1233011032121333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023000330232230-0201313302211200-2212013010232033-1111101223323223-2132320000121222-1331010030211202-3010313322100131-3223023103121220"></a>

## policy_based_challenge.rule_list.rules.spec.domain_matcher — domain_matcher / 221030300302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.domain_matcher

<a id="canonical-1101223320002201-0121100120130032-0022121223222023-0023113323133231-0003012013230310-0000000111121102-1003101103121320-3331122120213203"></a>

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
domain_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1002310121332002-2023222002232123-3001322232333001-1033010120033213-0312121003003033-2222320201222020-0212211131111113-1121110111202023"></a>

## Direct properties — domain_matcher / 221030300302 / 3

<a id="canonical-1111210100301113-3220313022122313-2120223223203132-3112330100312320-3330311203302110-0022332033121111-1030033331223332-0320202103132300"></a>

<a id="canonical-3301102100120203-2333332011130311-0023130021010131-3311111133033202-1203021313230313-2313321212303231-0131001120122211-0311203201320003"></a>

## exact_values property — domain_matcher / 221030300302 / 4

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

<a id="canonical-0323030030212232-3222300020331111-1301133310102021-2210121113202013-3031213322013100-2302011332210303-3121301103102231-3121213032311321"></a>

<a id="canonical-1003102010111133-1332320311331233-3333321332302310-2120203001013030-3300331120323032-0230101123033222-0201223220313211-3111211110112201"></a>

## regex_values property — domain_matcher / 221030300302 / 5

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

<a id="canonical-1213311013303013-3313113302321211-3103003223301003-1200203313021303-0031302002201031-3232312110320323-0133123113222010-3332103112031032"></a>

## Next pages — domain_matcher / 221030300302 / 6

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3103311303303301-3120122203132131-2113220010310230-3210210112101313-1321321213233223-2130133230221223-3223033330123030-1202010323003320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112321230223002-2032022333300322-1231010321120110-1210311210021223-1133130002200232-0130331301312213-3100320310220213-3101332020102030"></a>

## policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge — enable_captcha_challenge / 303001033221 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge

<a id="canonical-3232011221223112-1311003011100021-3023311130120120-2213322030301332-1212113021121301-2013230322112210-0010111131321010-2023112203220010"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable captcha challenge.

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
enable_captcha_challenge = {}
```

<a id="canonical-3333012232232232-3120120121031303-2230020200212220-3331132103320213-3000021312102231-3110133203033233-0202130031203311-0232210223330131"></a>

## Direct properties — enable_captcha_challenge / 303001033221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122023231333322-3200131200303020-2032333321221311-0002103130013200-2310302312102021-0303303303323110-1131301101133013-1231222010133022"></a>

## Next pages — enable_captcha_challenge / 303001033221 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1101321300022020-2012121100121213-1300112332120133-3301000323233111-0133300213322003-1130203021233202-1031230021202010-0002003123331010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132301202232122-2020132131133330-0020311302002013-0213302210322020-3322000201230300-2020031212020131-3313203320300020-0131113013101202"></a>

## policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge — enable_javascript_challenge / 300100202213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge

<a id="canonical-3130131113200201-3012320302303122-2200113101002110-3313132120231333-0021232020331303-1001222001331212-1132212103130213-0123220301123033"></a>

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
enable_javascript_challenge = {}
```

<a id="canonical-2121333110113030-3322221023103221-2111302123120122-0211310113220220-3213200301223020-0210300131000201-3231003213222132-0322313001210313"></a>

## Direct properties — enable_javascript_challenge / 300100202213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0121322221202113-2120033300301013-0333312211003011-0131112031310111-1211300313011232-3013310031123123-1202321002113330-2133233123210133"></a>

## Next pages — enable_javascript_challenge / 300100202213 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023030213100012-0033330020330010-1021112131001313-1220330303002220-1103322331300131-3123313323231133-3211303311201032-0211302300320011"></a>

## policy_based_challenge.rule_list.rules.spec.headers — headers / 032232021022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.headers

<a id="canonical-0303223210112332-0020033010231131-2023130000030012-2131023021323100-3210100130123301-3101212200200310-1111032020212301-1000313320033001"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1201302020312300-1131100003313122-3222103201130203-2021223030103130-0012221121323023-0210220110011230-2231223230201202-3322232011330232"></a>

## Direct properties — headers / 032232021022 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-2133232230322132-2013003023230313-2032312011333330-0323031212212313-1202223133002302-0130302300200130-1131302022302301-3222132301213310): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-2323233012323120-1101203011222330-0330221010331120-2003030011001211-2233323300223022-3110201200320233-0313322133232210-3320222232201200): complete subsection reference.

<a id="canonical-2123213032011220-3220320133331303-1310101330201121-1132310231132121-2003133133123112-3033013130021133-2002213233300002-1021332113001210"></a>

<a id="canonical-1012013332022101-3102113023112032-0301331333233331-2203100012002203-0232111111100110-0131212311311011-1002000113232332-3311331132022022"></a>

## invert_matcher property — headers / 032232021022 / 4

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

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

- [item](resources--cdn_loadbalancer--reference--group-014.md#canonical-0020010213313231-2230213012100103-3201102112031033-0120000120332332-2021003020323031-1222322303113232-1122212310101013-1031120222103022): complete subsection reference.

<a id="canonical-0121003112331220-1112103201302231-3010332210213111-3213100130103033-3230202230103332-2322113131302331-0330122313012030-3031132003310131"></a>

<a id="canonical-3313223330232310-1030112202122221-1000000032222300-3023200210333101-2132020221132313-0032223310113321-3033300101231131-2331331220122211"></a>

## name property — headers / 032232021022 / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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

<a id="canonical-2200030230313102-1020101131233032-2003021201023133-2230201211320310-1100132003312003-1202033322331111-3323320031310133-1203320322030313"></a>

## Next pages — headers / 032232021022 / 6

- [policy_based_challenge.rule_list.rules.spec.headers.check_not_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-2133232230322132-2013003023230313-2032312011333330-0323031212212313-1202223133002302-0130302300200130-1131302022302301-3222132301213310)
- [policy_based_challenge.rule_list.rules.spec.headers.check_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-2323233012323120-1101203011222330-0330221010331120-2003030011001211-2233323300223022-3110201200320233-0313322133232210-3320222232201200)
- [policy_based_challenge.rule_list.rules.spec.headers.item](resources--cdn_loadbalancer--reference--group-014.md#canonical-0020010213313231-2230213012100103-3201102112031033-0120000120332332-2021003020323031-1222322303113232-1122212310101013-1031120222103022)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2133232230322132-2013003023230313-2032312011333330-0323031212212313-1202223133002302-0130302300200130-1131302022302301-3222132301213310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301230030022003-0300123023013001-1023221022103113-3312132113003230-0101233310030213-3032312022330023-2020132331311200-1022220003211221"></a>

## policy_based_challenge.rule_list.rules.spec.headers.check_not_present — check_not_present / 120013231333 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020)
- policy_based_challenge.rule_list.rules.spec.headers.check_not_present

<a id="canonical-1203002130210320-0213000302220210-2010331232221233-1302222202222010-1021312133111030-3010123023110301-0001113203111132-0322032123213320"></a>

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

<a id="canonical-2123321122021300-3223312101121212-3132331021322203-2332312123122032-2131300212212232-0030100221320003-1212022002332211-3112301223011130"></a>

## Direct properties — check_not_present / 120013231333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313212103213000-3230021132030113-3132012323220033-1202223333310133-3031000301030100-2221232333220311-3023201122110131-1323330323200103"></a>

## Next pages — check_not_present / 120013231333 / 4

- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2323233012323120-1101203011222330-0330221010331120-2003030011001211-2233323300223022-3110201200320233-0313322133232210-3320222232201200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300033131120103-2222030111312032-3100031013322220-2112232021020321-1310301213321110-0111013311102003-1100233101113033-2233022110213011"></a>

## policy_based_challenge.rule_list.rules.spec.headers.check_present — check_present / 301122221303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020)
- policy_based_challenge.rule_list.rules.spec.headers.check_present

<a id="canonical-3220201211002330-0011220302023230-2223222222122032-3112101223201233-3213323321232000-1113132030031200-0333332010002120-1312031011003230"></a>

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

<a id="canonical-0133020332132003-3132001310020131-3320332113201122-1111012103131323-2003032231133201-1203010303101301-1223231013202223-2211231113132002"></a>

## Direct properties — check_present / 301122221303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231030313223233-1321332031233200-3310210230301113-1212332201033001-3203322102201303-2010321212323010-3131000230200010-2301303321302332"></a>

## Next pages — check_present / 301122221303 / 4

- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0020010213313231-2230213012100103-3201102112031033-0120000120332332-2021003020323031-1222322303113232-1122212310101013-1031120222103022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313001122132311-1313221030102131-1331021130312111-1000130003201000-2112220112113101-3020133322102000-3312321202102003-1221022131010321"></a>

## policy_based_challenge.rule_list.rules.spec.headers.item — item / 223332230102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020)
- policy_based_challenge.rule_list.rules.spec.headers.item

<a id="canonical-1001121230332002-2000120030103030-3112012121222110-2123030320333113-1010121333133320-0030303121301013-3312322022232023-1113320001011110"></a>

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

<a id="canonical-2132011333011110-0010113130331022-0321220021130010-1000100011010211-2203310312210132-1223230123320213-1123102013212323-0011323133323232"></a>

## Direct properties — item / 223332230102 / 3

<a id="canonical-1130213221202233-0303321311221210-1010001122233012-1130121100320201-2203230323311320-0000203012311003-3222101133233020-3120013313012002"></a>

<a id="canonical-2120212130002311-2131031022023112-0211311223311201-1122301101101102-1030000220332320-2120112201332221-2131130332303100-0112120332133210"></a>

## exact_values property — item / 223332230102 / 4

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

<a id="canonical-3000133123131333-0311133232321223-1322223102031322-2100103302030011-3030320010210132-2022032211212330-2121110322000112-1101210031301011"></a>

<a id="canonical-3201302330110323-3220210300011130-2322220221220103-2210020123212012-1033210131203111-3121201133130031-0223130303330001-0101300212213231"></a>

## regex_values property — item / 223332230102 / 5

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

<a id="canonical-0230101033312311-2210010000103120-1100201110021122-3033123232303112-3210120111021200-1303210202220213-1312001002023330-0010203130100300"></a>

<a id="canonical-3000100010022213-0112321302121032-2233120233313012-1003120303200100-1102032212323122-1103311112121331-1110212230133200-3223320312033231"></a>

## transformers property — item / 223332230102 / 6

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

<a id="canonical-1300011123003111-2301300322000223-2202002203100121-2020302330331331-1002232130200212-2313020022101022-2111000100222213-3001133013323312"></a>

## Next pages — item / 223332230102 / 7

- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2312131312221300-0230011100003021-0021131212211331-0123333322230033-1210310032303322-1320033320200030-2221332222001320-2330310311021003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000011311220310-3010010100122200-2213332310122103-3010031312312011-2100131103221231-1332113130102301-2020332223002021-0111321302112230"></a>

## policy_based_challenge.rule_list.rules.spec.http_method — http_method / 032230030010 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.http_method

<a id="canonical-3320212121102313-2230031103220110-2211201003313212-1320210223011323-0001012002020232-2200111211312211-0132203132213100-0323223001122232"></a>

Type: `"object"`. single nested block, Optional.

HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Upstream description:

A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
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

<a id="canonical-2102313213231110-2313103231123113-3013011013323301-2321011032012130-3210300012101202-0221322112212332-3113102010021201-0221103322113300"></a>

## Direct properties — http_method / 032230030010 / 3

<a id="canonical-2203232232111313-1221212000123010-3323200113023331-3200321032220321-2201110100201303-2222301123120121-0322202000303033-1213031011010112"></a>

<a id="canonical-3001133310101203-0013132132003232-1211223101003332-1200031233012020-3023103333121101-3212013213222210-1121323222122030-3200330110123212"></a>

## invert_matcher property — http_method / 032230030010 / 4

Type: `"bool"`. Optional.

Invert Method Matcher. Invert the match result.

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

<a id="canonical-0010102333021022-1312312211110233-2301231000333301-2033312002002101-2013303130302133-0113220022301202-1011222320210021-1031210211321100"></a>

<a id="canonical-3201222210113220-2330222123312101-3312013003213020-2131102300131011-0113312210022033-1331332310001103-3130131123321001-2330230211210311"></a>

## methods property — http_method / 032230030010 / 5

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of methods values to match against.

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

<a id="canonical-1132230013213112-0023110123223023-1011010122301113-1231133020211122-2313320013313111-0113211130212203-1211213313303131-0322312130211031"></a>

## Next pages — http_method / 032230030010 / 6

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1221203300113022-2132333200200013-1003100021200202-1120013111223113-1232003120333102-0311332302111030-2000323030332330-3022032101132020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021102121333021-3330120312111330-2130202313113013-3311032032313032-0300231230332212-0211103313202313-2122101232301131-3201323100200121"></a>

## policy_based_challenge.rule_list.rules.spec.ip_matcher — ip_matcher / 001213212132 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.ip_matcher

<a id="canonical-2322333132221020-1232232220213220-1232221023030310-1211311021103233-0330010333012110-3211320002300323-0332002303231310-2132123021123110"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1100003003132303-2100032023102020-2001320000230310-1101231233310230-2213030110100130-0331232231012123-0302231311333323-0113022013212013"></a>

## Direct properties — ip_matcher / 001213212132 / 3

<a id="canonical-2010112032010113-1330030033132030-1300312212033120-3232102120323323-1030232131012030-3232133302002001-3223122332103033-1021123300301313"></a>

<a id="canonical-2120031110001020-3021113033201312-3001230102032330-3110120122100130-2112022101113133-0120111311101230-3100230302300123-0220333223101002"></a>

## invert_matcher property — ip_matcher / 001213212132 / 4

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

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

- [prefix_sets](resources--cdn_loadbalancer--reference--group-014.md#canonical-2311012020023332-0201330100311030-3233112032312331-1222302001110020-1110031220112131-3112030003122113-1012032302112032-0312221130221132): complete subsection reference.

<a id="canonical-1312320013131202-3221311303102221-0132131222301301-1113012223302202-2103223232301301-1000230111231033-2112223011130001-0002330032313121"></a>

## Next pages — ip_matcher / 001213212132 / 5

- [policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets](resources--cdn_loadbalancer--reference--group-014.md#canonical-2311012020023332-0201330100311030-3233112032312331-1222302001110020-1110031220112131-3112030003122113-1012032302112032-0312221130221132)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2311012020023332-0201330100311030-3233112032312331-1222302001110020-1110031220112131-3112030003122113-1012032302112032-0312221130221132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231213013100021-3330023332021211-1203210100321020-2301231221223102-1300222021131000-0021032223202103-2222100233120220-2233021201321132"></a>

## policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets — prefix_sets / 113103102301 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-1221203300113022-2132333200200013-1003100021200202-1120013111223113-1232003120333102-0311332302111030-2000323030332330-3022032101132020)
- policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets

<a id="canonical-2123201001123223-1201000120120331-1030021130323111-1100313200303112-3232210032301101-2300133322301003-0121222303211323-3110330220122200"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

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

<a id="canonical-3030233132133021-3021023220220203-3011232211222010-3323200132210133-2131322012101122-1131032023200101-2202020311232130-1111321133331001"></a>

## Direct properties — prefix_sets / 113103102301 / 3

<a id="canonical-0232320111222030-1302203003330311-2331123202101230-3112132310130210-3232231202010102-1230122210230300-0333210333010111-2020101213212102"></a>

<a id="canonical-1010333022100323-2112310011033310-0310003013021231-2232001110001023-2022113213122011-3201203230003321-3220023213131300-2323312312220022"></a>

## kind property — prefix_sets / 113103102301 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3121132013100330-1220122031222003-2232213001333031-3213320213302331-3201203200212003-0102301321202333-3103310322133321-1101233112000301"></a>

<a id="canonical-1311200113013122-1023220122111011-2320322101211313-0230113102301313-3002202330223003-0031223222003210-0223232010101312-1101231213030212"></a>

## name property — prefix_sets / 113103102301 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0011012202233302-3220331121000213-0120233003300210-1321200212322232-3011023221031233-1113012010233021-3101321232222030-2300331213123313"></a>

<a id="canonical-3011210021020201-3102221203112112-2301022223132323-2001130002011302-3302203331211011-0322131230100110-1021100133222121-3110321221201210"></a>

## namespace property — prefix_sets / 113103102301 / 6

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

<a id="canonical-0312000323213231-0111012032202000-3033211033210133-3013230131121200-3213010021120312-1222022132301313-3013202302313023-1002131033012000"></a>

<a id="canonical-1212112232230102-3303222331313110-3313111131132302-1131331023022103-3313102211322203-0311303232032133-0200313030012301-0311133033022012"></a>

## tenant property — prefix_sets / 113103102301 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2300122200321331-3300303111213333-1112202332201133-3110022331232021-1203220200202301-1021203313300132-3102321003101113-3310002112332100"></a>

<a id="canonical-1103303332211033-3333123022013311-2112331111303311-0103201133133001-3120220330303021-1231000303032323-3332132212300030-2013212110011321"></a>

## uid property — prefix_sets / 113103102301 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0201122022130012-0010013020230103-2211013001223122-3032203322310220-3213302331202213-0023001030200201-2002200232220010-3320233103133032"></a>

## Next pages — prefix_sets / 113103102301 / 9

- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-1221203300113022-2132333200200013-1003100021200202-1120013111223113-1232003120333102-0311332302111030-2000323030332330-3022032101132020)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1012033320010022-2222231020323103-2321000021201310-3031023311331022-0012002023131321-2230011300323100-3130312013222320-2213121223200210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022213012310232-3210021222221233-0330332132332321-2223100031113101-2223113110113123-3013102120021333-2111000233221222-3202033131021100"></a>

## policy_based_challenge.rule_list.rules.spec.ip_prefix_list — ip_prefix_list / 003202312303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.ip_prefix_list

<a id="canonical-1123000020003201-1122230222010213-0301000031013301-3320012112200222-1230232102003203-2110011211121312-3232200333231210-0113303301210303"></a>

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

<a id="canonical-1233301313303023-0201011012032333-3021111331110333-2300020203011030-0201132032031113-2000001312002123-2310212121330213-3221020003323201"></a>

## Direct properties — ip_prefix_list / 003202312303 / 3

<a id="canonical-2312232003033113-2302032002221002-3032200012102023-3023122133002223-3011002031300010-2100101111132311-2333320033023322-1002323022212002"></a>

<a id="canonical-0012100212202131-2130221021221220-1312131100322202-1310032303202022-0201122211223030-3222131123210130-3023221313330003-0312201013033320"></a>

## invert_match property — ip_prefix_list / 003202312303 / 4

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

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

<a id="canonical-0330231110001002-0332222322223120-2100121122023210-3320332310302010-3233120113221123-0232010201020213-3210122311223300-0000110031131311"></a>

<a id="canonical-1033232300123210-1123311030313320-3330101223230320-0011332321200022-2000110131211323-1022330201313231-1030133313021020-2103132223202123"></a>

## ip_prefixes property — ip_prefix_list / 003202312303 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

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

<a id="canonical-1113320222203312-0231011202321310-0002030302111130-2323211203310230-2220023010111313-2132003311312113-2133332132100212-2231120123001323"></a>

## Next pages — ip_prefix_list / 003202312303 / 6

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1323031133100220-3003131300131310-1203112022000311-1230021132323020-2132121211131031-2202311300331212-3111230320020110-3011011131332111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113131132203211-1220321131033020-3230103032003001-3123031311130220-0101000203001023-3320310322131300-3111112233310233-3130132222222311"></a>

## policy_based_challenge.rule_list.rules.spec.path — path / 102111100000 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.path

<a id="canonical-1000020111011030-3203310221322031-0320302111200221-3200133321222322-2311112012030313-1301310222031232-1100211231322230-3031222002130313"></a>

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

<a id="canonical-2232233212103003-3233310321133303-3013301322221223-3002001113010110-1211332330201121-3113310021101132-0013003133120322-1321310020132102"></a>

## Direct properties — path / 102111100000 / 3

<a id="canonical-3310030332322101-3102210022203201-1220320102220233-0013001232011113-1202232320121100-3203202211331303-1333133203321303-2001003203011000"></a>

<a id="canonical-0331210032231020-1200123320321223-0100310221032200-0322230031302103-1320311210002222-3223232202231232-2032001220123033-1002121100110323"></a>

## encoded_path_matcher property — path / 102111100000 / 4

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

<a id="canonical-1122223033203001-3000311122102201-1112321202310120-0311211011310301-2203113031002213-2011131220320102-3100013221113313-3000223103212201"></a>

<a id="canonical-2011101301212203-2011030010130321-1000032021130213-3000212120121301-3331230001332121-3212021311323031-3001133130202303-0233300101031313"></a>

## exact_values property — path / 102111100000 / 5

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

<a id="canonical-1023110331303112-0102023321222113-0213003120013100-1301323300301220-1112023310321022-0030113121201213-0132123123211033-1231010312223202"></a>

<a id="canonical-3113302200330222-3121110211310313-1031313330300131-3030132023110201-2022331130333023-2101331030333201-2221333130231332-2003003123010212"></a>

## invert_matcher property — path / 102111100000 / 6

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

<a id="canonical-0311320231223032-2120202323130222-2203102302001202-3220123301312230-0223002331212213-1232312201330221-3231103223100010-2131210333231203"></a>

<a id="canonical-3023103333133333-1203022200002101-2032301121001101-0113112123131200-1003010200033111-2313021300003321-0211211231332113-0231102120013133"></a>

## prefix_values property — path / 102111100000 / 7

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

<a id="canonical-1333131200321220-0210021133020003-1310122003103122-3312320211030213-1130101133203003-3001321322211311-0033330023013123-0031220001122032"></a>

<a id="canonical-2022221022310113-0011221221300231-3222121003330122-3313323033322012-1320321030032132-2200023022322213-3101313011123013-1021333200022213"></a>

## regex_values property — path / 102111100000 / 8

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

<a id="canonical-0302133323201220-0320231030111112-1220331012311220-2121311010302022-1212033000301113-3020112222212302-3200123011013011-2020200121131103"></a>

<a id="canonical-0023322320222021-3133310331311022-0213101021230113-3321223230013220-0201021302020320-3002321322131223-0102000232301113-3000012112133303"></a>

## suffix_values property — path / 102111100000 / 9

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

<a id="canonical-3101323201032123-3130111021201301-2132330011201310-1001121301111111-3030112133003021-0200010300321033-0202010311201222-3220213200013312"></a>

<a id="canonical-3300002133111103-3321323312221330-0202311033220100-3200102132000312-1331113223001221-1333332123311013-1223223220111300-2110233222000112"></a>

## transformers property — path / 102111100000 / 10

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

<a id="canonical-2020310200213223-0120223110100010-3311130010212112-3203333321303033-2221321213333031-3332011013332303-2222113322202323-0230012220203011"></a>

## Next pages — path / 102111100000 / 11

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132210231130311-1002030311322103-3230021211120112-3321020321331033-3132303223231231-0322330302023210-0333113213203310-0303033201012100"></a>

## policy_based_challenge.rule_list.rules.spec.query_params — query_params / 211323210301 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.query_params

<a id="canonical-0330013112223030-0223233103011231-2112031210313113-2332310022120312-1232033200100301-2011020202113021-0013031301001100-2302032002200200"></a>

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

<a id="canonical-2222012103022311-0030322220330112-1102120332311130-0021023322323202-3002021122122121-0332302132022332-2303133301303121-2331030231331102"></a>

## Direct properties — query_params / 211323210301 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-1312300330202202-0230330032101303-2331213220023113-3312230312301330-3331100020223113-1303131000020002-2031223202300310-1031130011210022): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-0011200213331102-3303001022301002-0122112221113011-1101032132033112-2322010100320111-1203021310120331-0020211133132211-2133230123112201): complete subsection reference.

<a id="canonical-3102010112201132-3222113311210212-1102022201021013-0212100133323203-2302133113021031-2021311301103220-0223311303200101-0312321030212100"></a>

<a id="canonical-0133030032331312-1233120133000302-3213300120100313-2123023110210031-1332300102200210-3123021122200312-2331021102302123-0300313120210132"></a>

## invert_matcher property — query_params / 211323210301 / 4

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

- [item](resources--cdn_loadbalancer--reference--group-014.md#canonical-0300132120303011-2222013321322121-1103100003033103-0130000011023130-3102203201303332-3231303212012130-0201320300021112-0303323231120233): complete subsection reference.

<a id="canonical-0121103323123020-0322032200300210-1010000123320232-2213220011002320-0201100133022202-2313102110311320-1011212122101331-3220223033023222"></a>

<a id="canonical-1203200132132021-3110200320021122-2123223100030231-2120122001122320-1020302020332033-1312310310023111-3102331012211100-1211031100133323"></a>

## key property — query_params / 211323210301 / 5

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

<a id="canonical-3331312100131033-3100002220312322-0120023320232201-0131120322022013-0232223233232112-0130212331301232-1023133002200033-0033003320213230"></a>

## Next pages — query_params / 211323210301 / 6

- [policy_based_challenge.rule_list.rules.spec.query_params.check_not_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-1312300330202202-0230330032101303-2331213220023113-3312230312301330-3331100020223113-1303131000020002-2031223202300310-1031130011210022)
- [policy_based_challenge.rule_list.rules.spec.query_params.check_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-0011200213331102-3303001022301002-0122112221113011-1101032132033112-2322010100320111-1203021310120331-0020211133132211-2133230123112201)
- [policy_based_challenge.rule_list.rules.spec.query_params.item](resources--cdn_loadbalancer--reference--group-014.md#canonical-0300132120303011-2222013321322121-1103100003033103-0130000011023130-3102203201303332-3231303212012130-0201320300021112-0303323231120233)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1312300330202202-0230330032101303-2331213220023113-3312230312301330-3331100020223113-1303131000020002-2031223202300310-1031130011210022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002133120321013-3012301303121012-3322223120013100-2232213311120013-0122230321211020-3031213103130303-3121122230213023-1001300233033222"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.check_not_present — check_not_present / 203203211212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100)
- policy_based_challenge.rule_list.rules.spec.query_params.check_not_present

<a id="canonical-3323303321102130-1313123201200100-3000011111221332-0333223210323220-3211233200332202-3033233301200002-1313312001233111-1320232330220113"></a>

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

<a id="canonical-3110033230202320-1333032212000213-1201122120021232-0212030023303230-2112221322111100-0000131020300331-1331211001331302-0133023030322022"></a>

## Direct properties — check_not_present / 203203211212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310211000122030-2022230211320332-0213321011033231-1103202221122011-3212213102101022-3210020300300112-3111231311010212-2002020021233001"></a>

## Next pages — check_not_present / 203203211212 / 4

- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0011200213331102-3303001022301002-0122112221113011-1101032132033112-2322010100320111-1203021310120331-0020211133132211-2133230123112201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102131020211033-2000113220323031-2300233131332100-3202322123200001-3222312023132313-3030231232311220-2332032003312232-2022021313331022"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.check_present — check_present / 131310213212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100)
- policy_based_challenge.rule_list.rules.spec.query_params.check_present

<a id="canonical-1220320221120133-0033321113211020-1203312212003131-1313001232113100-3233202022203202-2313321203030320-3133133103002010-3333103002021312"></a>

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

<a id="canonical-0233211303300333-3010311202100213-1333320212003013-3220330210210101-2213303111200330-2220211332223020-3013220101333112-1230201020232113"></a>

## Direct properties — check_present / 131310213212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213332232002121-0300101313032313-1320201032013111-2102023203302003-1333230101030011-2132200330122112-3102133011200021-0002201310231220"></a>

## Next pages — check_present / 131310213212 / 4

- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0300132120303011-2222013321322121-1103100003033103-0130000011023130-3102203201303332-3231303212012130-0201320300021112-0303323231120233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320302312110211-2213212022021231-0002001211221020-1213133231032100-3113200102120021-2112132033230013-1110222310221103-0210313323303010"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.item — item / 232103313033 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100)
- policy_based_challenge.rule_list.rules.spec.query_params.item

<a id="canonical-0102033000211301-3202211131220331-2000333320311112-1232201230231123-3110102200032212-1133321033112230-3202312110233200-1323220303212130"></a>

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

<a id="canonical-2230113320132311-3112331203320201-2202013010133122-1331313012300010-1100330111030201-0113230033203030-3003320213102132-1132100133001000"></a>

## Direct properties — item / 232103313033 / 3

<a id="canonical-2302023302020011-3110113331121322-2323311121120322-3002132110300213-0301221333230233-3202122303131011-2031213033220032-2023112201313332"></a>

<a id="canonical-3020111202023211-1201032333023302-3312132103033210-1000332130003010-1111210023110000-1213010123330302-2120222123221121-1122203303232122"></a>

## exact_values property — item / 232103313033 / 4

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

<a id="canonical-0303000332223232-1232123101211223-1031132013231223-1313113202103030-2031230231323030-0201313012301232-0011021013103222-0022313031313211"></a>

<a id="canonical-0311221232300320-3003120000101020-0301232233003221-1212033002010321-1331331131313330-0312021113120010-3033000320033302-0012223323011120"></a>

## regex_values property — item / 232103313033 / 5

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

<a id="canonical-1303101121133330-0033031313313021-2302331021320122-3200200211103111-0112311321133021-3332111210220123-2120112213000020-0032123103120022"></a>

<a id="canonical-0200111323022122-3200023332201132-2223202310010313-2212002203131132-1300220020203100-2311320002221103-2333021012231122-0033100302222231"></a>

## transformers property — item / 232103313033 / 6

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

<a id="canonical-3321032121312220-2313113201012033-1321331021213110-1211033023211122-3123213033030000-1123111110222213-0101302110300101-2030133313320201"></a>

## Next pages — item / 232103313033 / 7

- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3232022003302213-0032131233110110-3231221110233032-1010300113202110-1012221322120230-1322201202221221-2233333330333200-3030322000303331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223320300213331-1200031102303232-0230131000301003-1010132103310221-2303123111012030-3203032231221000-2313232221223113-1332031033323203"></a>

## policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher — tls_fingerprint_matcher / 212011302322 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-2123022203130322-0022112310210121-3033311120121203-1120020002120211-3320221302133201-3012120322302120-1222232201031222-1313313331203301"></a>

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

<a id="canonical-2021320223331302-1012031110000302-1230311310202221-0032321113110212-2013021311003312-2010200121103002-0100131222101020-2201313210333011"></a>

## Direct properties — tls_fingerprint_matcher / 212011302322 / 3

<a id="canonical-0332102033203002-1300033000302223-1320111233030302-2011330223122011-0202033001013312-3103112203002020-1113210220023011-1310023101331132"></a>

<a id="canonical-0311221132103132-3311012203132000-1231022131021130-2121201313312331-1021120010023320-3213121120101101-1222021223000233-0232311230020123"></a>

## classes property — tls_fingerprint_matcher / 212011302322 / 4

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

<a id="canonical-0232122031310203-1130322311131121-1133103013313321-3121332003023112-1013321200030331-2202211230101021-0232311203201112-2203200311220213"></a>

<a id="canonical-3320211131120023-3313201033320001-1000211220100203-0301100012101111-3022331212001300-0212331003212221-0102122001311132-1001111112020221"></a>

## exact_values property — tls_fingerprint_matcher / 212011302322 / 5

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

<a id="canonical-2101002303102000-2012230021120100-3120111331223100-1023203333233010-3323322133212220-1101133231221321-1113220103111033-3000333121121203"></a>

<a id="canonical-0230332103221312-0303303002103123-1000022310320122-1100033221203331-3323130102301032-0201002210333110-3130232003103103-2223011211133212"></a>

## excluded_values property — tls_fingerprint_matcher / 212011302322 / 6

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

<a id="canonical-1003223003313013-1212303201120222-1012321301300021-0032000301210033-2021122010020222-3030013213222311-0122220130323310-0210132022313132"></a>

## Next pages — tls_fingerprint_matcher / 212011302322 / 7

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2020113203001122-0131003133101010-3112033112103122-1230132210230022-1110101121230003-3011322320122320-2001301222311323-1030230313310012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110003331013331-3130133013301231-3010203210213123-2222111310330220-1203123301030033-3221112233210000-2113132010310313-0201133201120020"></a>

## policy_based_challenge.temporary_user_blocking — temporary_user_blocking / 313131322111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.temporary_user_blocking

<a id="canonical-1201223313121033-1103213030333221-1331232332133120-3021210133032030-0032312022010311-2131030221021131-2322001320123223-3002233120103201"></a>

Type: `"object"`. single nested block, Optional.

Specifies configuration for temporary user blocking resulting from user behavior analysis. When
Malicious User Mitigation is enabled from service policy rules, users' accessing the application
will be analyzed for malicious activity and the configured mitigation actions will be taken on..

Upstream description:

Specifies configuration for temporary user blocking resulting from user behavior analysis.

When Malicious User Mitigation is enabled from service policy rules, users' accessing the
application will be analyzed for malicious activity and the configured mitigation actions will be
taken on identified malicious users. These mitigation actions include setting up temporary blocking
on that user. This configuration specifies settings on how that blocking should be done by the
loadbalancer.

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
temporary_user_blocking {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203323001221123-2201231232233133-2013311110311300-2222031023220302-1311232333322232-1321332230223330-0103202320320313-2232322110332220"></a>

## Direct properties — temporary_user_blocking / 313131322111 / 3

<a id="canonical-0010103303001310-0213230232201003-3123212123321232-2211122222002222-1330020101323220-1302331232121123-2113123000330012-2233321033223010"></a>

<a id="canonical-3110121210132233-2222202010111232-3212110031213231-2331302123321232-3333120232232230-2113210202212211-1012313232130303-2103332202302231"></a>

## custom_page property — temporary_user_blocking / 313131322111 / 4

Type: `"string"`. Optional.

Custom message is of type . Currently supported URL schemes is . For scheme, message needs to be
encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g.
'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '&lt;p&gt;
Blocked..

Upstream description:

Custom message is of type \`uri\_ref\`. Currently supported URL schemes is \`string:///\`. For
\`string:///\` scheme, message needs to be encoded in Base64 format. You can specify this message as
base64 encoded plain text message e.g. "Blocked.." or it can be HTML paragraph or a body string
encoded as base64 string E.g. "&lt;p&gt; Blocked &lt;/p&gt;". Base64 encoded string for this HTML is
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1300033232101210-2211002303231201-3123200003100110-1002111020211031-1010212013000222-1232031333002002-2210322002000301-1102133211022210"></a>

## Next pages — temporary_user_blocking / 313131322111 / 5

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313333230210103-0130123231011131-3012330332330013-3123011131033220-3133022002233021-3113131000213013-1221201010201311-1330033030030030"></a>

## protected_cookies — protected_cookies / 201113323322 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- protected_cookies

<a id="canonical-2223231201320222-0020301102201100-2031133233311002-1000230332031133-1303122030112310-0231202302013302-3122033213202302-1312113133230002"></a>

Type: `"object"`. list nested block, Optional.

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite.

Upstream description:

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF
(monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note:
We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("disable_tampering_protection",
    "enable_tampering_protection"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict")}
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

Terraform syntax:

```terraform
protected_cookies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200331313333211-3310003000030103-2030101213003232-2112111211232101-2131233333210222-1210300311013333-3120320230031331-2300320111311301"></a>

## Direct properties — protected_cookies / 201113323322 / 3

- [add_httponly](resources--cdn_loadbalancer--reference--group-014.md#canonical-1222133013311211-0333131120023113-3121113303130123-0312213332010220-1233033023120021-0221013100020132-1321230210000300-0210201010300313): complete subsection reference.

- [add_secure](resources--cdn_loadbalancer--reference--group-014.md#canonical-1132000333120232-3302322021300311-2313311210013112-0322232321221301-0212300032233331-1312322121103103-2330002320301013-1320232002031322): complete subsection reference.

- [disable_tampering_protection](resources--cdn_loadbalancer--reference--group-014.md#canonical-3300333222232023-0313031213010202-2102000001032100-2112313322020220-1233220033132310-2111330213212301-3323312202103000-1011031123202323): complete subsection reference.

- [enable_tampering_protection](resources--cdn_loadbalancer--reference--group-014.md#canonical-3022000331032231-1230000032102312-1211131333230321-3212123301010213-3131211011121332-2222331321231220-3023220111001123-1133210011210102): complete subsection reference.

- [ignore_httponly](resources--cdn_loadbalancer--reference--group-014.md#canonical-0322013100232300-3320012232320223-0302321121132122-2022111333203301-1030032201101211-0022020002113133-2123320001003022-0233133312121322): complete subsection reference.

- [ignore_max_age](resources--cdn_loadbalancer--reference--group-014.md#canonical-3331121030110230-3113323312022330-0300013331330211-3211201022230320-3002302301000022-3022011032202013-1012021332121122-1231123101031131): complete subsection reference.

- [ignore_samesite](resources--cdn_loadbalancer--reference--group-014.md#canonical-1001312120203320-0332332333210221-0021331233302310-2213332032130322-0332010223111130-2310201113022012-3033320130302030-0322220300311012): complete subsection reference.

- [ignore_secure](resources--cdn_loadbalancer--reference--group-014.md#canonical-2331301203101122-2122302001110110-2001122133213231-0103100311122302-1323130110210320-3002010000212312-2000232110301312-2130323323010010): complete subsection reference.

<a id="canonical-2012210210032003-3211210200031031-3202023210102201-2031322310333323-3031303112132202-3311331102301002-0121333302031110-1011001302130301"></a>

<a id="canonical-1200312033123300-3301000312201220-2021303012101001-0123102110203202-3220000312303003-0011010122121102-1113312300013010-0132200230021130"></a>

## max_age_value property — protected_cookies / 201113323322 / 4

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-0301330121223030-0212231032220331-0330232212023230-0210120333013003-0232030123123222-2323313322023331-1132022332021300-0011333122033301"></a>

<a id="canonical-2023230122201211-3113200021202111-1001201112111032-2300101222022003-1232300321330021-0233203200313021-3101012022123120-0312211132301133"></a>

## name property — protected_cookies / 201113323322 / 5

Type: `"string"`. Optional.

Cookie Name. Name of the Cookie.

Upstream description:

Name of the Cookie.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [samesite_lax](resources--cdn_loadbalancer--reference--group-014.md#canonical-2101022332031001-0323332310312332-2113121333322000-1313013133120330-0132010202201213-1330301200203231-1021011323232122-0010222230103230): complete subsection reference.

- [samesite_none](resources--cdn_loadbalancer--reference--group-014.md#canonical-1232203201132302-1331211313003233-3213032100202223-1223313032122220-2211332102102021-3133010113022030-3321023133303011-2023010011020020): complete subsection reference.

- [samesite_strict](resources--cdn_loadbalancer--reference--group-014.md#canonical-0003230130333321-0023020200021211-1100010112302123-3132331221102033-1313232222231120-2233011102220102-0211111032322201-1301311002131023): complete subsection reference.

<a id="canonical-3301003233320311-1323202230232202-3300121002113010-0322001103112302-2300101311203313-1012101110002331-0120312320131031-0032130131130313"></a>

## Next pages — protected_cookies / 201113323322 / 6

- [protected_cookies.add_httponly](resources--cdn_loadbalancer--reference--group-014.md#canonical-1222133013311211-0333131120023113-3121113303130123-0312213332010220-1233033023120021-0221013100020132-1321230210000300-0210201010300313)
- [protected_cookies.add_secure](resources--cdn_loadbalancer--reference--group-014.md#canonical-1132000333120232-3302322021300311-2313311210013112-0322232321221301-0212300032233331-1312322121103103-2330002320301013-1320232002031322)
- [protected_cookies.disable_tampering_protection](resources--cdn_loadbalancer--reference--group-014.md#canonical-3300333222232023-0313031213010202-2102000001032100-2112313322020220-1233220033132310-2111330213212301-3323312202103000-1011031123202323)
- [protected_cookies.enable_tampering_protection](resources--cdn_loadbalancer--reference--group-014.md#canonical-3022000331032231-1230000032102312-1211131333230321-3212123301010213-3131211011121332-2222331321231220-3023220111001123-1133210011210102)
- [protected_cookies.ignore_httponly](resources--cdn_loadbalancer--reference--group-014.md#canonical-0322013100232300-3320012232320223-0302321121132122-2022111333203301-1030032201101211-0022020002113133-2123320001003022-0233133312121322)
- [protected_cookies.ignore_max_age](resources--cdn_loadbalancer--reference--group-014.md#canonical-3331121030110230-3113323312022330-0300013331330211-3211201022230320-3002302301000022-3022011032202013-1012021332121122-1231123101031131)
- [protected_cookies.ignore_samesite](resources--cdn_loadbalancer--reference--group-014.md#canonical-1001312120203320-0332332333210221-0021331233302310-2213332032130322-0332010223111130-2310201113022012-3033320130302030-0322220300311012)
- [protected_cookies.ignore_secure](resources--cdn_loadbalancer--reference--group-014.md#canonical-2331301203101122-2122302001110110-2001122133213231-0103100311122302-1323130110210320-3002010000212312-2000232110301312-2130323323010010)
- [protected_cookies.samesite_lax](resources--cdn_loadbalancer--reference--group-014.md#canonical-2101022332031001-0323332310312332-2113121333322000-1313013133120330-0132010202201213-1330301200203231-1021011323232122-0010222230103230)
- [protected_cookies.samesite_none](resources--cdn_loadbalancer--reference--group-014.md#canonical-1232203201132302-1331211313003233-3213032100202223-1223313032122220-2211332102102021-3133010113022030-3321023133303011-2023010011020020)
- [protected_cookies.samesite_strict](resources--cdn_loadbalancer--reference--group-014.md#canonical-0003230130333321-0023020200021211-1100010112302123-3132331221102033-1313232222231120-2233011102220102-0211111032322201-1301311002131023)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1222133013311211-0333131120023113-3121113303130123-0312213332010220-1233033023120021-0221013100020132-1321230210000300-0210201010300313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300213001221221-2331012220310230-3300300321131311-3020020012332233-1120030320020311-1213103310121130-0322000210020321-0021313103002131"></a>

## protected_cookies.add_httponly — add_httponly / 003103232213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.add_httponly

<a id="canonical-0110023213031302-2102002121002103-2222332030011222-0022122301131021-2132231032111231-1101003022301212-3131302210022012-3130110113212001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

<a id="canonical-1022330333103211-0220132220101221-2231203022212021-2210022311323123-2111111102122223-3001013033321320-0233222001032320-0221013033303232"></a>

## Direct properties — add_httponly / 003103232213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300203220313310-3303233200112331-2120021112020332-3220001311231300-1120001102032033-1101300313213220-0022320000210221-3322313302302322"></a>

## Next pages — add_httponly / 003103232213 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1132000333120232-3302322021300311-2313311210013112-0322232321221301-0212300032233331-1312322121103103-2330002320301013-1320232002031322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002310023100100-0120011321003231-1000211120111003-2301303001221023-0113202232120222-2303102302310322-0123012032233033-1302303211333110"></a>

## protected_cookies.add_secure — add_secure / 330312220303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.add_secure

<a id="canonical-3210033102032123-3311002120120202-3030312001111322-0022332111131033-1231323210133321-2132020133212221-1112121010312212-3301233312110122"></a>

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
add_secure = {}
```

<a id="canonical-2223103031130010-2033311121311103-2332011003131212-3302313321123022-3230033112230121-2233001120331230-1000203203300310-1110021320002323"></a>

## Direct properties — add_secure / 330312220303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322232023311012-1013012333132113-0131231111222120-0103101333133110-2121032102213120-1313231301231113-2010331312000332-1133302123031122"></a>

## Next pages — add_secure / 330312220303 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3300333222232023-0313031213010202-2102000001032100-2112313322020220-1233220033132310-2111330213212301-3323312202103000-1011031123202323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303101322212302-2130133113103132-0011013021131033-3202221131323023-1032130220121330-1122001102112313-1130000120201133-0320300111311132"></a>

## protected_cookies.disable_tampering_protection — disable_tampering_protection / 000313211022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.disable_tampering_protection

<a id="canonical-3033132112112320-1011212120130211-0323321020112001-3032300020033031-3102233200102132-1200000303031213-3312212020301311-2132301230101232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable tampering protection.

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
disable_tampering_protection = {}
```

<a id="canonical-3112200200212301-3000200000013131-2300202331032332-1011022131303033-2300222102222221-3022123320231302-0313101020301212-0030102331230311"></a>

## Direct properties — disable_tampering_protection / 000313211022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003321333211311-1230003332232310-1220311032302033-0011220131033310-0321330201010232-2311010330030020-2111300013012201-0203333321213113"></a>

## Next pages — disable_tampering_protection / 000313211022 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3022000331032231-1230000032102312-1211131333230321-3212123301010213-3131211011121332-2222331321231220-3023220111001123-1133210011210102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030313132233113-0320000311023103-3212100131220231-3313221201302231-0000003031232213-0122101312011020-1032323300130023-1012231200330323"></a>

## protected_cookies.enable_tampering_protection — enable_tampering_protection / 300021020011 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.enable_tampering_protection

<a id="canonical-1011002213133222-1322310130030031-2112321030300220-3002113113003330-3031212122003131-2321031133323203-0111313130332032-0221121313212303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable tampering protection.

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
enable_tampering_protection = {}
```

<a id="canonical-1201020313312203-1020103012002202-3312221000003212-0110032302122100-3012210010212200-3011132000011133-1022010222233330-2111021021330322"></a>

## Direct properties — enable_tampering_protection / 300021020011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212003022011321-3023213303010212-2112233302003110-2311011003223023-2212102013030030-1010320022212112-2123202200130102-3121332322021021"></a>

## Next pages — enable_tampering_protection / 300021020011 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0322013100232300-3320012232320223-0302321121132122-2022111333203301-1030032201101211-0022020002113133-2123320001003022-0233133312121322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132130002120021-2020230001311000-1222201100200022-2020120303223333-1113132223121302-0133123022321331-3021213223023011-3022210010321333"></a>

## protected_cookies.ignore_httponly — ignore_httponly / 020211102233 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.ignore_httponly

<a id="canonical-0312331323131101-2112322330310133-3031100211220113-3232233321112201-3312013113023300-0110203202223303-2322222230010200-0020231133201202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

<a id="canonical-3301222223000121-1221113123013332-2300310213323020-0223031003123202-2333203123322303-2133310003010103-3120030122320101-0302331230231131"></a>

## Direct properties — ignore_httponly / 020211102233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2312002033130113-0122231012001201-0300221131002211-0110301001331322-1022021230233102-3221332210032123-2011112120202011-3212231223333031"></a>

## Next pages — ignore_httponly / 020211102233 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3331121030110230-3113323312022330-0300013331330211-3211201022230320-3002302301000022-3022011032202013-1012021332121122-1231123101031131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113220323010320-0311031332022233-0303302212130110-1202321123212330-3223213322020130-0010122111312210-1220003100010212-2030210232331132"></a>

## protected_cookies.ignore_max_age — ignore_max_age / 330211320032 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.ignore_max_age

<a id="canonical-3203231010202120-3020013211132010-0323322130221002-2131111011212111-2332101200120112-3201311233223303-3202232220310232-3113032122312121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore max age.

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
ignore_max_age = {}
```

<a id="canonical-1001010003301032-1002132022301031-1013133223032011-0131113131333303-0222102321132330-1033013032033031-2303123023131230-1032000330311332"></a>

## Direct properties — ignore_max_age / 330211320032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001021312021222-0231113023111211-2312103021030022-3202211030312003-1231111231301030-0320313123003002-2322000110111331-2231320330103010"></a>

## Next pages — ignore_max_age / 330211320032 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1001312120203320-0332332333210221-0021331233302310-2213332032130322-0332010223111130-2310201113022012-3033320130302030-0322220300311012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301120100032220-1233030003213311-1311111200013210-1221223320123033-3233001102310220-3230231233022302-0202110333333300-1333023213213220"></a>

## protected_cookies.ignore_samesite — ignore_samesite / 000022301111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.ignore_samesite

<a id="canonical-0200201121101123-2013330201300120-3033311302212030-1323332021020032-3302221332131130-0232200123110213-3033011232222021-1312301232200120"></a>

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
ignore_samesite = {}
```

<a id="canonical-0003303231213322-3120212002131033-1112001111331332-1330301330033130-0321310013323002-3323313233231233-1223023330022032-2001013113103200"></a>

## Direct properties — ignore_samesite / 000022301111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112212013321303-3010120331312322-0321210103331223-0121323203203333-1031223211332000-1113033112312311-2123323321213320-0023311032121311"></a>

## Next pages — ignore_samesite / 000022301111 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2331301203101122-2122302001110110-2001122133213231-0103100311122302-1323130110210320-3002010000212312-2000232110301312-2130323323010010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300210021003011-2123302330013332-0110122201321222-3203202100202320-0033012210321011-1200110300120303-0320020203023231-3222231213210100"></a>

## protected_cookies.ignore_secure — ignore_secure / 132323302022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.ignore_secure

<a id="canonical-1301203101223311-1132110222022301-3311320003230132-1311132323112002-0230330320323222-3120123003122011-0330110120021321-2322311220311232"></a>

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
ignore_secure = {}
```

<a id="canonical-2300330231023200-0001000311123232-0203300220310333-2011211233112001-2230033322032003-3302022300312313-1023101001120112-3311001102212211"></a>

## Direct properties — ignore_secure / 132323302022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230123321322213-0131300131202131-3332100203201133-0230201033022101-3133023323303020-1002323111100031-1231130011033102-3320133232011110"></a>

## Next pages — ignore_secure / 132323302022 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2101022332031001-0323332310312332-2113121333322000-1313013133120330-0132010202201213-1330301200203231-1021011323232122-0010222230103230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222033113330301-3133100121302123-1201223333323310-2013231100233222-3122302320101332-1113100231113113-2021130330323302-0101022232130201"></a>

## protected_cookies.samesite_lax — samesite_lax / 101102310103 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.samesite_lax

<a id="canonical-0112113002012121-3120032321130213-2322222022122032-2122330020000202-2202311210022201-2110313210021313-2233103002322110-1230102332323233"></a>

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
samesite_lax = {}
```

<a id="canonical-3120331111102121-0122222022213100-1131131003313300-1323101103101120-0322333113301232-1311100213020130-3003102103032102-1322221103333033"></a>

## Direct properties — samesite_lax / 101102310103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100103330302230-3231000320330212-1131212301120310-1322010321213023-1001312120002132-1213320110213113-0012222202323010-0223033303122302"></a>

## Next pages — samesite_lax / 101102310103 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1232203201132302-1331211313003233-3213032100202223-1223313032122220-2211332102102021-3133010113022030-3321023133303011-2023010011020020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321113120113120-0000113323132323-1221003132022132-3312032313102023-3032120133032021-2223021032320310-2001323222133013-0212032100032031"></a>

## protected_cookies.samesite_none — samesite_none / 221310321201 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.samesite_none

<a id="canonical-1102110013020121-2333000002012003-0112033002010123-1200122231033312-2223103103112103-1200203123110111-2211020221303321-0323220202211120"></a>

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
samesite_none = {}
```

<a id="canonical-0200123120102021-0210100133232033-1311033313012300-2233013031112111-3133002130131032-3333331320302310-2020102220221300-1103013331010003"></a>

## Direct properties — samesite_none / 221310321201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322021313033310-1023232011330323-2101030111112233-1021301001131001-3122202222030013-1300333110233311-2321302322312311-3233111212213213"></a>

## Next pages — samesite_none / 221310321201 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0003230130333321-0023020200021211-1100010112302123-3132331221102033-1313232222231120-2233011102220102-0211111032322201-1301311002131023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121103001213221-0133201133322333-3110311120333300-0211213222133330-0020121001021130-0332031003022112-1133210131230101-2232011301200331"></a>

## protected_cookies.samesite_strict — samesite_strict / 303320203000 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.samesite_strict

<a id="canonical-3132232222231122-1202121321110033-0012220121023132-0300320310313331-3020032031003020-3223222100030233-0012220221202212-2102233220113311"></a>

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
samesite_strict = {}
```

<a id="canonical-1030133232112213-0133030303323130-3331013032003113-2201300131312010-0300022301011123-3221222022131110-0122210111130000-1223330122111033"></a>

## Direct properties — samesite_strict / 303320203000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112330101300001-1012130122111112-0010002311020331-2333330310202013-2101100031222023-2113101221302302-1123211000223013-0102022301202023"></a>

## Next pages — samesite_strict / 303320203000 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023130321310200-1311313112001322-1023223002200032-2201302330030101-3120303121312333-2200121112033212-3121212013120311-1120122230112120"></a>

## rate_limit — rate_limit / 122112233122 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- rate_limit

<a id="canonical-3323121002313201-3031120322332203-2113131320323133-3022023331331231-1100333301321203-0121201122211101-2223013112203301-0210101213220100"></a>

Type: `"object"`. single nested block, Optional.

RateLimitConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("no_policies",
    "policies")}
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
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]",
  "x-ves-oneof-field-policy_choice": "[\"no_policies\",\"policies\"]"
}
```

Terraform syntax:

```terraform
rate_limit {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313020203200100-0002013221122011-3113010222120200-1112112103320113-0222123122120201-0333221012200131-2123322322023100-3101203331222310"></a>

## Direct properties — rate_limit / 122112233122 / 3

- [custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-3211202232232112-1332230121300022-3332022212230331-2133332330303223-2213132312211233-3323123331302103-2333032000301112-3332003332201331): complete subsection reference.

- [ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-0131111121111212-3101110311102221-2133123013002303-2210110213111330-1133030303203122-0003130031330133-1021013031213230-3221000033233130): complete subsection reference.

- [no_ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-2213010213120023-0100103320013021-3122212310121112-1322112200122212-2010212132123210-0012012333133011-2020101110020110-2132200212132230): complete subsection reference.

- [no_policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0132002221322313-3000222211302221-0201132202300003-2010311200322230-2303312220332001-0221100132130231-0122131031023220-0011032210230110): complete subsection reference.

- [policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0332012013133002-0110210320323201-1003122022020300-0120220113202102-3111330131211212-2231211322332003-0010201011000213-2030331032023213): complete subsection reference.

- [rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211): complete subsection reference.

<a id="canonical-1011303133003200-1302001131132123-1333201001110032-1003022003010303-0301301001200221-2101001103101331-0100001011101023-2022321231102013"></a>

## Next pages — rate_limit / 122112233122 / 4

- [rate_limit.custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-3211202232232112-1332230121300022-3332022212230331-2133332330303223-2213132312211233-3323123331302103-2333032000301112-3332003332201331)
- [rate_limit.ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-0131111121111212-3101110311102221-2133123013002303-2210110213111330-1133030303203122-0003130031330133-1021013031213230-3221000033233130)
- [rate_limit.no_ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-2213010213120023-0100103320013021-3122212310121112-1322112200122212-2010212132123210-0012012333133011-2020101110020110-2132200212132230)
- [rate_limit.no_policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0132002221322313-3000222211302221-0201132202300003-2010311200322230-2303312220332001-0221100132130231-0122131031023220-0011032210230110)
- [rate_limit.policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0332012013133002-0110210320323201-1003122022020300-0120220113202102-3111330131211212-2231211322332003-0010201011000213-2030331032023213)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3211202232232112-1332230121300022-3332022212230331-2133332330303223-2213132312211233-3323123331302103-2333032000301112-3332003332201331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203223300133100-2022022313033223-0002102310010310-3012203100120131-0213023012002232-3231231113132213-2133021001131233-1221223231232113"></a>

## rate_limit.custom_ip_allowed_list — custom_ip_allowed_list / 023022211203 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- rate_limit.custom_ip_allowed_list

<a id="canonical-3132311130311310-0313000312201223-2122001110133321-3232332123020213-3312000301320322-3212122030112232-3010021123303331-1030022223011001"></a>

Type: `"object"`. single nested block, Optional.

IP Allowed list using existing ip\_prefix\_set objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate_limiter_allowed_prefixes")}
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
custom_ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1112000233022101-3330332122133331-1221322321111331-3200311103333001-3202200103103333-0213212012010031-3103101203223212-2032212122330021"></a>

## Direct properties — custom_ip_allowed_list / 023022211203 / 3

- [rate_limiter_allowed_prefixes](resources--cdn_loadbalancer--reference--group-014.md#canonical-1133001332311010-3112323220120313-1220011323101130-1332132211101130-2312133303011233-0221102210200211-1302201101231200-2023100031230022): complete subsection reference.

<a id="canonical-2332231333103110-0300313213331330-1203322032021012-1222020002223020-0313031011022211-0032211313121133-0110000200101222-2013120211110301"></a>

## Next pages — custom_ip_allowed_list / 023022211203 / 4

- [rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](resources--cdn_loadbalancer--reference--group-014.md#canonical-1133001332311010-3112323220120313-1220011323101130-1332132211101130-2312133303011233-0221102210200211-1302201101231200-2023100031230022)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1133001332311010-3112323220120313-1220011323101130-1332132211101130-2312133303011233-0221102210200211-1302201101231200-2023100031230022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301131010103032-0113333221200231-2102333013123010-2222101123021013-3023013113233202-0213101121010023-0111023213031202-0002320132003000"></a>

## rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes — rate_limiter_allowed_prefixes / 320232121122 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-3211202232232112-1332230121300022-3332022212230331-2133332330303223-2213132312211233-3323123331302103-2333032000301112-3332003332201331)
- rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-3003302313033001-1313202211310203-1030233213303232-1323133031032320-3201012320003101-0321303031202220-3003003102302322-0011002320000232"></a>

Type: `"object"`. list nested block, Optional.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Upstream description:

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
rate_limiter_allowed_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022101210333311-1233021023011133-3000312221221102-1011031320311320-2211300213231121-3302132300300011-2121202122011221-1213223312332213"></a>

## Direct properties — rate_limiter_allowed_prefixes / 320232121122 / 3

<a id="canonical-0101022010001110-2303231222213232-3321221033002203-2302213112133330-3120032102013122-2302220230110320-3011121322003122-2310011112020233"></a>

<a id="canonical-2013002021322120-2131303122233210-2122233211201232-2213121321201102-0020211030021113-1110100230233310-1310222331032013-2020230113313003"></a>

## name property — rate_limiter_allowed_prefixes / 320232121122 / 4

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

<a id="canonical-3001022021112321-0223003303221323-0300222013010222-3012330322000302-3132333132112313-1320020321010030-3130233003233301-1133202321022133"></a>

<a id="canonical-1220033031113221-0323331102121311-3203100311331211-1100113013100112-0231331230230303-3221200132001322-3301213130102122-0212011302123232"></a>

## namespace property — rate_limiter_allowed_prefixes / 320232121122 / 5

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

<a id="canonical-2101303110223230-2102213211123333-3012322212303122-2032302203303201-3200320000333111-3311030210122231-2130033032313110-2301231202121221"></a>

<a id="canonical-2130131233011121-3112210031112201-1000232213133322-3032333202331110-3202122232123331-0222110132211003-3202313310103300-3102310121023133"></a>

## tenant property — rate_limiter_allowed_prefixes / 320232121122 / 6

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

<a id="canonical-1303120001231132-0223121311131130-3221123102233113-2222033321310210-3301213200112232-3020323232002201-2100302123202022-0212003100312020"></a>

## Next pages — rate_limiter_allowed_prefixes / 320232121122 / 7

- [rate_limit.custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-3211202232232112-1332230121300022-3332022212230331-2133332330303223-2213132312211233-3323123331302103-2333032000301112-3332003332201331)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0131111121111212-3101110311102221-2133123013002303-2210110213111330-1133030303203122-0003130031330133-1021013031213230-3221000033233130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222333021200020-1223221302021213-0203330021220201-1121332032031030-0011220313312320-1132012021230000-3233303021130200-2003203211300333"></a>

## rate_limit.ip_allowed_list — ip_allowed_list / 010332310101 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- rate_limit.ip_allowed_list

<a id="canonical-3322230203002002-3021222312312202-1122032302131230-0230132130333231-2121310202320301-1133200122232321-0330111003213100-1201230232122321"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201120222012002-0012310231333003-1223221003203131-0320212203223300-2222320331211222-0301202023120130-0113322200212303-2230201331310121"></a>

## Direct properties — ip_allowed_list / 010332310101 / 3

<a id="canonical-0311010123110110-1213103003232003-2133230131033120-0021133112301122-3000231202331313-1033220012022103-2213021123002012-3011013010303233"></a>

<a id="canonical-1012033120220313-0322031013212331-3101013102302211-1122222100103222-1332311120122130-1202323110212321-0011232222111002-2013310302202223"></a>

## prefixes property — ip_allowed_list / 010332310101 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2130221101333020-2001303122331201-3031131331323103-3321032301221222-2003013321133001-1313321033101333-0302013021021231-3231230213111033"></a>

## Next pages — ip_allowed_list / 010332310101 / 5

- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2213010213120023-0100103320013021-3122212310121112-1322112200122212-2010212132123210-0012012333133011-2020101110020110-2132200212132230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022323010221003-1021132110233310-2000333232122112-3000110033020001-2322200111101123-3112321131332033-2131313301230231-2303100223212212"></a>

## rate_limit.no_ip_allowed_list — no_ip_allowed_list / 030322203332 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- rate_limit.no_ip_allowed_list

<a id="canonical-0211122133022010-2030330003021210-0111021020303323-3132000021023102-3001020023320322-3301333012113322-3010203122321311-1202101023310321"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
no_ip_allowed_list = {}
```

<a id="canonical-1120023103020020-3312023300033302-3332101002322122-0103010133213121-2112113311013120-1033202021201223-2333333221313030-3231012320130230"></a>

## Direct properties — no_ip_allowed_list / 030322203332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332223332010331-2202210313000102-1310000200113120-0101122112021133-0330020332113133-2323010202310112-1033003010111300-1022100111321332"></a>

## Next pages — no_ip_allowed_list / 030322203332 / 4

- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0132002221322313-3000222211302221-0201132202300003-2010311200322230-2303312220332001-0221100132130231-0122131031023220-0011032210230110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331130110132321-2103122133002303-1213002101300301-1000211013302200-0000333123001302-0232302223330312-3321303030022031-1022233111020121"></a>

## rate_limit.no_policies — no_policies / 331130102020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- rate_limit.no_policies

<a id="canonical-0201120202210222-0020323201213220-2301033103202302-2202321323132201-1133113023120222-0022213321131113-2113223103033102-1023302330100313"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no policies. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_policies = {}
```

<a id="canonical-1113010331333033-3200213033100203-2123202222130213-1120032131031103-0322221312230121-0333332110031220-1200202120320221-0212210331301123"></a>

## Direct properties — no_policies / 331130102020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330003310313001-3330222323030230-2303220203133122-0303103312123101-1113110011111301-2030101121113011-2230000230221132-1300123030222123"></a>

## Next pages — no_policies / 331130102020 / 4

- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0332012013133002-0110210320323201-1003122022020300-0120220113202102-3111330131211212-2231211322332003-0010201011000213-2030331032023213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330230302303320-1113033221020012-0301033323333032-2013321210220001-0203131313023213-2323202000203210-0301102313332200-2122310230311330"></a>

## rate_limit.policies — policies / 232320102103 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- rate_limit.policies

<a id="canonical-3001311301302003-1033320212321121-2012223000213002-2233311231320313-2000213030222022-0022102310212320-2033021111210303-3011333223003122"></a>

Type: `"object"`. single nested block, Optional.

List of rate limiter policies to be applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("policies")}
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
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313213013123101-3003201222132332-0213000032200311-2321233030211120-2011210033310002-3333130012133020-3131313123302133-3030213303203322"></a>

## Direct properties — policies / 232320102103 / 3

- [policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-3031012130111313-0333133110300122-3212021223302331-1010031320331132-3220011102222022-1030211202302321-3011132103101230-2210310200333303): complete subsection reference.

<a id="canonical-0013123311020312-1211031110230023-0102000213030030-3330332102330131-0201201301203231-2232103032033223-0230333301113233-3201200321202021"></a>

## Next pages — policies / 232320102103 / 4

- [rate_limit.policies.policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-3031012130111313-0333133110300122-3212021223302331-1010031320331132-3220011102222022-1030211202302321-3011132103101230-2210310200333303)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3031012130111313-0333133110300122-3212021223302331-1010031320331132-3220011102222022-1030211202302321-3011132103101230-2210310200333303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130130322101101-3033100123321302-0122120313320003-2013023132012023-0233211013001201-0211021233002213-1303002003120201-2103321012330311"></a>

## rate_limit.policies.policies — policies / 123222220202 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0332012013133002-0110210320323201-1003122022020300-0120220113202102-3111330131211212-2231211322332003-0010201011000213-2030331032023213)
- rate_limit.policies.policies

<a id="canonical-0113000321133111-2002020013322023-1001010232211332-3133333113132203-3001203310120031-0030001301030032-1211312033033022-0122002223311101"></a>

Type: `"object"`. list nested block, Optional.

Rate Limiter Policies. Ordered list of rate limiter policies.

Upstream description:

Ordered list of rate limiter policies.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123330000001032-0232113120021012-2131102132120223-3001022112332131-2133010023121110-1231333013203300-1020310213022122-2321011010030121"></a>

## Direct properties — policies / 123222220202 / 3

<a id="canonical-0110320110211230-3103100313112231-0120102030221013-1103011231012330-2333213023220200-0032101233212120-0303110332000310-0232120100321202"></a>

<a id="canonical-3200101130021033-0110201212110220-0222023313002210-1121003101213012-3100231113331021-2130311230321121-0123003023021210-0312210120220213"></a>

## name property — policies / 123222220202 / 4

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

<a id="canonical-1110123333110001-0203120303231123-0222120223131230-2312210010120120-2302022302120203-0312231331010130-0220330022233123-3021230332122033"></a>

<a id="canonical-0002301131210323-3133330312020100-1200111213023320-1221320303100211-0013201323232212-3003100100220132-0213220313211233-1223210020213110"></a>

## namespace property — policies / 123222220202 / 5

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

<a id="canonical-3001121002003321-2230001000330311-2211032032120130-0203110213303322-2103121132131212-0030201122320003-3330221233132232-1033320302300022"></a>

<a id="canonical-2001210332303002-3020321103022113-3322000212023020-2212202123112131-1001023210330010-3223332301200021-2013221110003220-0231202103231211"></a>

## tenant property — policies / 123222220202 / 6

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

<a id="canonical-1321003323133001-1303003122330322-1212313032302130-0213031333020331-3103302003012013-3021320020310213-3021320312221220-0231232022123001"></a>

## Next pages — policies / 123222220202 / 7

- [rate_limit.policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0332012013133002-0110210320323201-1003122022020300-0120220113202102-3111330131211212-2231211322332003-0010201011000213-2030331032023213)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220310321311020-0031030011111222-2122002223130020-1213113003120230-2031213121231203-3203302013210113-1011020023313320-2323010312022323"></a>

## rate_limit.rate_limiter — rate_limiter / 013012100020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- rate_limit.rate_limiter

<a id="canonical-3202232323000300-1302010301323202-2112210310332012-2013111231122303-2323230010231120-2110001103122020-2032132221113323-0222333022323332"></a>

Type: `"object"`. single nested block, Optional.

Tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Upstream description:

A tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("total_number"),
  validators.ConflictingObjectAttributes("action_block",
    "disabled"),
  validators.ConflictingObjectAttributes("leaky_bucket",
    "token_bucket")}
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
  "x-ves-oneof-field-action_choice": "[\"action_block\",\"disabled\"]",
  "x-ves-oneof-field-algorithm": "[\"leaky_bucket\",\"token_bucket\"]"
}
```

Terraform syntax:

```terraform
rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-1312131121103011-0231000011120310-1002113303231201-2012022001023130-3232302200200210-0132213001223322-2021222200001030-3110132021102322"></a>

## Direct properties — rate_limiter / 013012100020 / 3

- [action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131): complete subsection reference.

<a id="canonical-2121003302013021-2121130112031230-3123332123101323-3221123222213223-0330023100113303-2330213013221210-3202103013331021-0300232022011222"></a>

<a id="canonical-2110012021001222-1331102000102322-3322232121111310-3312333222310232-0101321033231331-1222202231002232-2232121023320020-1111032331120321"></a>

## burst_multiplier property — rate_limiter / 013012100020 / 4

Type: `"number"`. Optional.

The maximum burst of requests to accommodate, expressed as a multiple of the rate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [disabled](resources--cdn_loadbalancer--reference--group-014.md#canonical-0121113003211121-3000233231212222-0210322310333131-1232221120110102-0323331121001101-3300320021101330-3211100221101210-3303200202212010): complete subsection reference.

- [leaky_bucket](resources--cdn_loadbalancer--reference--group-014.md#canonical-2121301302012332-0200200003121032-3002102313230233-0221300213312113-3323111113013323-2321301313301030-1201012223223023-2123010120320301): complete subsection reference.

<a id="canonical-0321222321313130-2131301201011131-1101123011120111-0123310301230323-0130032012210031-0121003212300313-1232132202312211-3313202303300210"></a>

<a id="canonical-0132311222220200-1000120210111023-0020223132201232-3220213013123200-2020120210110003-2023012212203333-1323332020202213-2012012231230221"></a>

## period_multiplier property — rate_limiter / 013012100020 / 5

Type: `"number"`. Optional, Computed.

Setting, combined with Per Period units, provides a duration. Server applies default when omitted.

Upstream description:

This setting, combined with Per Period units, provides a duration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  }
}
```

- [token_bucket](resources--cdn_loadbalancer--reference--group-014.md#canonical-3211020022010202-1101201222211102-3010213311221001-2102010330320122-0120301001332101-0332320102011122-0321200121122313-3032003110020000): complete subsection reference.

<a id="canonical-2113133301210033-2133213221021332-1112211323000030-2300100021213032-1131011033001301-2301320222133020-3022112322223110-0323001230233001"></a>

<a id="canonical-0022123123320322-1312312231201013-2321000220323303-0213132102121003-1122120032233020-0310300113303310-0020310120310031-2332203232002323"></a>

## total_number property — rate_limiter / 013012100020 / 6

Type: `"number"`. Optional.

The total number of allowed requests per rate-limiting period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-2223000210211010-2313102202201003-0323033323011331-1201232232320123-3320110200231122-3002100322030212-2101211212310323-2213201030212331"></a>

<a id="canonical-1011101003222231-1231200203102112-0011232221032033-0033120300123021-3020213133123303-0301020131013310-1111102020132201-3131101230000303"></a>

## unit property — rate_limiter / 013012100020 / 7

Type: `"string"`. Optional.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Upstream description:

Unit for the period per which the rate limit is applied.

&#8203;- SECOND: Second

Rate limit period unit is seconds &#8203;- MINUTE: Minute

Rate limit period unit is minutes &#8203;- HOUR: Hour

Rate limit period unit is hours &#8203;- DAY: Day

Rate limit period unit is days.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SECOND",
    "MINUTE",
    "HOUR"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2021002102303112-3323320332323133-3332133123130323-1211021132013022-1213202320223323-0003002302333110-3130133313023322-0001120111000232"></a>

## Next pages — rate_limiter / 013012100020 / 8

- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131)
- [rate_limit.rate_limiter.disabled](resources--cdn_loadbalancer--reference--group-014.md#canonical-0121113003211121-3000233231212222-0210322310333131-1232221120110102-0323331121001101-3300320021101330-3211100221101210-3303200202212010)
- [rate_limit.rate_limiter.leaky_bucket](resources--cdn_loadbalancer--reference--group-014.md#canonical-2121301302012332-0200200003121032-3002102313230233-0221300213312113-3323111113013323-2321301313301030-1201012223223023-2123010120320301)
- [rate_limit.rate_limiter.token_bucket](resources--cdn_loadbalancer--reference--group-014.md#canonical-3211020022010202-1101201222211102-3010213311221001-2102010330320122-0120301001332101-0332320102011122-0321200121122313-3032003110020000)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022323122112022-1133131113332201-3221032210123200-1112100212210330-1300231201110121-2013322011212312-3211330122130010-0100120001310230"></a>

## rate_limit.rate_limiter.action_block — action_block / 023232233113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- rate_limit.rate_limiter.action_block

<a id="canonical-3031132200211133-1023311111200312-2101201003113000-3223033010030221-0122300321113023-0100301000300220-3001220320112013-0021231001133011"></a>

Type: `"object"`. single nested block, Optional.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes"),
  validators.ConflictingObjectAttributes("hours",
    "seconds"),
  validators.ConflictingObjectAttributes("minutes",
    "seconds")}
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
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

Terraform syntax:

```terraform
action_block {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131003130113100-3103233000100021-1213221213003200-3320023021030112-2000330231001123-0020033312202313-0312002122012301-2201011221223200"></a>

## Direct properties — action_block / 023232233113 / 3

- [hours](resources--cdn_loadbalancer--reference--group-014.md#canonical-0001220131200010-1033233213023323-3333300333212200-2133013021233210-0301302322310330-3130330031202120-2310111303100313-0313332031331111): complete subsection reference.

- [minutes](resources--cdn_loadbalancer--reference--group-014.md#canonical-2232102003211200-3310003313003312-0012200212113100-0002323313331101-1020102203000121-1302032211020003-3000302231323112-1301000013201333): complete subsection reference.

- [seconds](resources--cdn_loadbalancer--reference--group-014.md#canonical-1320212332333121-2121103102302021-2000231022122212-1221020333221331-0120202330203000-1000300113331220-0211232131110222-2311323201323002): complete subsection reference.

<a id="canonical-3211020132202230-1222113221201300-2331221202021330-0123223010100322-3332212223022310-1232211122101221-1112012032113110-3000130323202120"></a>

## Next pages — action_block / 023232233113 / 4

- [rate_limit.rate_limiter.action_block.hours](resources--cdn_loadbalancer--reference--group-014.md#canonical-0001220131200010-1033233213023323-3333300333212200-2133013021233210-0301302322310330-3130330031202120-2310111303100313-0313332031331111)
- [rate_limit.rate_limiter.action_block.minutes](resources--cdn_loadbalancer--reference--group-014.md#canonical-2232102003211200-3310003313003312-0012200212113100-0002323313331101-1020102203000121-1302032211020003-3000302231323112-1301000013201333)
- [rate_limit.rate_limiter.action_block.seconds](resources--cdn_loadbalancer--reference--group-014.md#canonical-1320212332333121-2121103102302021-2000231022122212-1221020333221331-0120202330203000-1000300113331220-0211232131110222-2311323201323002)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0001220131200010-1033233213023323-3333300333212200-2133013021233210-0301302322310330-3130330031202120-2310111303100313-0313332031331111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330103111232101-3132233001303331-2222303213302230-3133203101103011-0002013223121020-1200133123211331-1313111122221333-3111101303113100"></a>

## rate_limit.rate_limiter.action_block.hours — hours / 001032330030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131)
- rate_limit.rate_limiter.action_block.hours

<a id="canonical-1020022103303203-0100213122112331-3022212321120202-0112003001222223-3033200113332023-3131300313013123-0003221133123311-2223123221123333"></a>

Type: `"object"`. single nested block, Optional.

Hours. Input Duration Hours.

Upstream description:

Input Duration Hours.

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
hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213103131211203-1203101022020111-2002210223300312-1120312311213130-0312000233203132-0302310222103013-2111320003023300-1102331123123113"></a>

## Direct properties — hours / 001032330030 / 3

<a id="canonical-3132212132133200-1032200323323022-1102313200331323-2333002220222202-2111313231323230-0313212020210323-0233100320002111-3330002311211120"></a>

<a id="canonical-2323121323123030-0002212313310113-1221320003310311-1121110013121302-2101201302323012-0213102103103033-3023002313212013-2221310223032200"></a>

## duration property — hours / 001032330030 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 48),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 48,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  }
}
```

<a id="canonical-2331222312300023-3013322200312130-2101000333032112-1231202032222301-3211322332312030-1311310120331112-2213200223311130-0023203112230210"></a>

## Next pages — hours / 001032330030 / 5

- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2232102003211200-3310003313003312-0012200212113100-0002323313331101-1020102203000121-1302032211020003-3000302231323112-1301000013201333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320001103010013-2123312000020332-0221013030031313-0112133130111222-1221212333123221-3101310031123010-2031220232203020-1321012131321323"></a>

## rate_limit.rate_limiter.action_block.minutes — minutes / 031003120122 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131)
- rate_limit.rate_limiter.action_block.minutes

<a id="canonical-1232303121032132-2302122103131301-0332200103310110-0121121331333213-0110203030123330-2021310203021323-3321031321010313-3202013302022010"></a>

Type: `"object"`. single nested block, Optional.

Minutes. Input Duration Minutes.

Upstream description:

Input Duration Minutes.

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
minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201000330020302-0022223323220311-2310032232111122-2032103221001013-1220322201031001-2133321333020001-3121000020330022-3122333200120202"></a>

## Direct properties — minutes / 031003120122 / 3

<a id="canonical-0011330110031122-2111300321230301-2010003300013220-2321013203003023-1233200032301030-0110003120000232-0233220112012303-0023133110231233"></a>

<a id="canonical-1003122323112032-0110322323120131-1022223002100103-1010023012222200-1111113100122230-2200322012213201-2120002220221010-0120313211233203"></a>

## duration property — minutes / 031003120122 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 60),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  }
}
```

<a id="canonical-0300010101211200-0333013001123102-1131021231112121-1133210233320022-1113220020310333-3110323130030312-0110002211333100-2131010101231330"></a>

## Next pages — minutes / 031003120122 / 5

- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1320212332333121-2121103102302021-2000231022122212-1221020333221331-0120202330203000-1000300113331220-0211232131110222-2311323201323002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112002010011011-0031211333010321-0032300003201331-1200001222200130-2213330022332331-0013212103113003-2231213200332320-2312220113312100"></a>

## rate_limit.rate_limiter.action_block.seconds — seconds / 023121233123 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131)
- rate_limit.rate_limiter.action_block.seconds

<a id="canonical-0011331330000033-2011200332122022-3130122233122122-3102131133231221-1011200333001022-0121000030102123-1031021113230122-3122122231021333"></a>

Type: `"object"`. single nested block, Optional.

Seconds. Input Duration Seconds.

Upstream description:

Input Duration Seconds.

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
seconds {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332211303000033-0133212121100311-1301202002202013-3130223122000303-2300113223203212-0101311331111001-3333310122230312-1222222112020122"></a>

## Direct properties — seconds / 023121233123 / 3

<a id="canonical-1232220201100310-3222033101201002-1113033122230001-3023003203210332-3001000320210310-1112102100100001-1012323033202123-0323303120223132"></a>

<a id="canonical-1210012102203313-2323101012210220-2332013221301203-3212230303332020-1302313110200320-2013322210301221-1011102121031203-1210311331030330"></a>

## duration property — seconds / 023121233123 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 300),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-3110122022010023-0300310311000120-1202012113233120-3100113222011203-0001333102303223-0033321320001113-2201111321001020-1011030123002102"></a>

## Next pages — seconds / 023121233123 / 5

- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0121113003211121-3000233231212222-0210322310333131-1232221120110102-0323331121001101-3300320021101330-3211100221101210-3303200202212010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012102113301212-3202210321300003-2122133213203030-0133023321202210-3232030002332121-0233233302122021-0021013320212320-0330323310112132"></a>

## rate_limit.rate_limiter.disabled — disabled / 002301100332 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- rate_limit.rate_limiter.disabled

<a id="canonical-1121312100131031-2223013120000303-3012122112011231-3000022132131122-3323321221113132-0212210312123320-0211111221020201-3202230211322330"></a>

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

<a id="canonical-1101223012200300-0233002232031303-1221003211301013-1313031322020012-0133003102203320-2030313333303122-1020210011001230-0311111122300022"></a>

## Direct properties — disabled / 002301100332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201113310301230-3012120013000120-3021322021000313-3321203302103333-0112312110013123-1302313001100332-1203233223102103-3103130131303203"></a>

## Next pages — disabled / 002301100332 / 4

- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2121301302012332-0200200003121032-3002102313230233-0221300213312113-3323111113013323-2321301313301030-1201012223223023-2123010120320301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331220203120002-0120002020302221-2312122113132212-2111233022031313-2233030000212222-2103131213300233-1210022322131032-0113220230330212"></a>

## rate_limit.rate_limiter.leaky_bucket — leaky_bucket / 301000032313 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- rate_limit.rate_limiter.leaky_bucket

<a id="canonical-1001223310211010-2211301333033321-0122300320033221-2301010110311101-3010003232121100-0232213013303010-2200103300012030-1211133312003023"></a>

Type: `["object", {}]`. Optional.

Leaky-Bucket is the default rate limiter algorithm for F5.

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
leaky_bucket = {}
```

<a id="canonical-3331323313320323-1301210203331103-0211001130320031-1312132333210203-1231333202322210-1202300330101011-1020233121232131-2220320210002013"></a>

## Direct properties — leaky_bucket / 301000032313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223131232110020-0310121211213300-3222022032030331-3121331213230200-3112321203232100-0313211131202332-3103320202001211-3130120013332101"></a>

## Next pages — leaky_bucket / 301000032313 / 4

- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3211020022010202-1101201222211102-3010213311221001-2102010330320122-0120301001332101-0332320102011122-0321200121122313-3032003110020000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303233201110311-1102030130203310-3030320130301311-2221202313023101-2201000031033122-3330331330020003-0222023132032111-1021213223100112"></a>

## rate_limit.rate_limiter.token_bucket — token_bucket / 010120210200 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- rate_limit.rate_limiter.token_bucket

<a id="canonical-2012231200012321-2301010321102022-1122211011332002-1231202222131112-3130202030223103-0021133201020233-1230312222321023-2132012312022122"></a>

Type: `["object", {}]`. Optional.

Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.

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
token_bucket = {}
```

<a id="canonical-2221030002123000-3203333020202322-2121220101100312-1201032300321310-1130303103002120-1103300320212100-3212111300030113-3012031010323332"></a>

## Direct properties — token_bucket / 010120210200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302302212233233-1001333310131110-0012211300323103-3222221100032323-3032131020121133-1112313013022322-1211002002020112-1320001011003221"></a>

## Next pages — token_bucket / 010120210200 / 4

- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3101301311121220-1022212312223231-0103230132010000-1031303112000210-1003201331112110-2210011211032323-3113203213013300-1111103022320301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032323113320110-2313230222132212-3120030220032000-1220220010022010-2303000220211222-1112102222102130-1312103200122002-3031000100322233"></a>

## sensitive_data_policy — sensitive_data_policy / 031322200023 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- sensitive_data_policy

<a id="canonical-1300020121333330-1303221303112300-0023322213301102-1002303223322122-2130031202131001-1332232121231102-1300111112133313-2022220010223230"></a>

Type: `"object"`. single nested block, Optional.

Policy configuration for this feature.

Upstream description:

Settings for data type policy.

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
sensitive_data_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-1100013112103210-0132000212130230-2200003122101333-0323112130010010-2202202321003303-1022122013032201-1313123332011121-1011102232230021"></a>

## Direct properties — sensitive_data_policy / 031322200023 / 3

- [sensitive_data_policy_ref](resources--cdn_loadbalancer--reference--group-014.md#canonical-0201002102200313-1002330132131011-3132132012332313-3300220302033000-0332032322333323-1213010011231132-0030331210203000-3231223310010012): complete subsection reference.

<a id="canonical-1310001010311331-2001313113303001-2011202313123002-2023303132010012-1223232303202122-0322321130103002-0122330011003201-0221230010302010"></a>

## Next pages — sensitive_data_policy / 031322200023 / 4

- [sensitive_data_policy.sensitive_data_policy_ref](resources--cdn_loadbalancer--reference--group-014.md#canonical-0201002102200313-1002330132131011-3132132012332313-3300220302033000-0332032322333323-1213010011231132-0030331210203000-3231223310010012)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0201002102200313-1002330132131011-3132132012332313-3300220302033000-0332032322333323-1213010011231132-0030331210203000-3231223310010012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122030223133001-2022103302033130-1233230121002121-0030020222322220-3112301320232110-1130111000232021-2303202323213330-0020121231121002"></a>

## sensitive_data_policy.sensitive_data_policy_ref — sensitive_data_policy_ref / 311030320133 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [sensitive_data_policy](resources--cdn_loadbalancer--reference--group-014.md#canonical-3101301311121220-1022212312223231-0103230132010000-1031303112000210-1003201331112110-2210011211032323-3113203213013300-1111103022320301)
- sensitive_data_policy.sensitive_data_policy_ref

<a id="canonical-3210100202211022-1332030101023121-0300010231230001-1202122213333001-1232033103122211-3303322310121230-2112033112212102-1132200333031200"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
sensitive_data_policy_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112000210032202-1310001212303322-1233133002313133-1330013133112110-1022323100030300-0231120203033323-2112012022213013-0011101101301200"></a>

## Direct properties — sensitive_data_policy_ref / 311030320133 / 3

<a id="canonical-3230300102010002-0201110113031320-0022020212113030-0121221121322003-2133231213332121-1001113233222233-2301023100323131-2001011011201110"></a>

<a id="canonical-3200032001301210-3232023103003122-3211020120313101-2022131301013320-1333301011222013-1021320300011112-1302222221303232-0010313030321021"></a>

## name property — sensitive_data_policy_ref / 311030320133 / 4

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

<a id="canonical-0212131002020121-1011132203320132-1000212330230121-1310111023233102-3111303021131202-3122212022313200-1021101222032031-0013210320032301"></a>

<a id="canonical-2300102202033203-1132133001232302-3332103033301031-3320221201111021-0123332100123110-2112221013323232-1110301113132223-1121330000003321"></a>

## namespace property — sensitive_data_policy_ref / 311030320133 / 5

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

<a id="canonical-0330322312332322-0133321130012122-3213100012313221-3212213022003220-0223110100300333-3121202122202330-0003030013300301-0332102323012000"></a>

<a id="canonical-2312203131232130-3200101230130123-3212122223301103-2113001011221303-1013302111010223-0100021131223211-0012110301203130-2232031201012201"></a>

## tenant property — sensitive_data_policy_ref / 311030320133 / 6

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

<a id="canonical-0203002100003201-2011301223321203-1101322332203320-0220100112001020-0103011221223031-0121020000023302-2101032222132311-2201211133310301"></a>

## Next pages — sensitive_data_policy_ref / 311030320133 / 7

- [sensitive_data_policy](resources--cdn_loadbalancer--reference--group-014.md#canonical-3101301311121220-1022212312223231-0103230132010000-1031303112000210-1003201331112110-2210011211032323-3113203213013300-1111103022320301)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0232201322200031-0330122100331311-3221203302302022-3021301122011232-3230010232020130-3332030103113102-0223113230122010-1011332330021133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102021023231203-2322013132303200-3333012112333010-1310012232020323-3210131100121311-1032330023310302-3011033221321211-2200101010033233"></a>

## service_policies_from_namespace — service_policies_from_namespace / 102230303311 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- service_policies_from_namespace

<a id="canonical-0212231232132301-3212113303020322-3222132120230030-1123313213201131-1130023023230203-2203303210021230-2103211113213120-2321022020101213"></a>

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
service_policies_from_namespace = {}
```

<a id="canonical-1212123233020210-1120132330200332-0313330200200223-3102100000232032-3321202323013120-1333203010131231-2003333331231231-0312120202321233"></a>

## Direct properties — service_policies_from_namespace / 102230303311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0330213323221221-1000310033001220-1213321110000220-1111331030321312-2130012302233121-3000130011031223-2010020300210131-0232021221001000"></a>

## Next pages — service_policies_from_namespace / 102230303311 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1233003020202211-1121000220031121-0002002321202011-3200220102103002-3312211001303013-1111001323333210-3232211013001110-3302023302323112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122132022120031-2021300013310121-2103211111333321-2220222032031101-0111213320011131-2233002323312120-0003001310021110-1333332333001333"></a>

## slow_ddos_mitigation — slow_ddos_mitigation / 320333120103 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- slow_ddos_mitigation

<a id="canonical-0112221300332233-1321331031100323-2021100310011011-3232120031232301-3333233030001223-2021312201110121-1201200310231320-3001221030310130"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: slow\_ddos\_mitigation, system\_default\_timeouts; Default: system\_default\_timeouts\]
'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Upstream description:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("request_headers_timeout"),
  validators.ConflictingObjectAttributes("disable_request_timeout",
    "request_timeout")}
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
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

OneOf alternatives in this subsection:

- [slow_ddos_mitigation](resources--cdn_loadbalancer--reference--group-014.md#canonical-0112221300332233-1321331031100323-2021100310011011-3232120031232301-3333233030001223-2021312201110121-1201200310231320-3001221030310130)
- [system_default_timeouts](resources--cdn_loadbalancer--reference--group-015.md#canonical-2003220332131013-2211333310231313-0030002201303021-1102200303220023-3011112311203322-2110130310322231-2112032032222113-2133222202031012)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
slow_ddos_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3323330331003121-0203313020131132-3233322231302213-0201331032132213-3221122131222332-2121321311303020-2023131230331300-2113332120120130"></a>

## Direct properties — slow_ddos_mitigation / 320333120103 / 3

- [disable_request_timeout](resources--cdn_loadbalancer--reference--group-014.md#canonical-3202203332302003-2003031212203232-1211332211112331-3321100223302000-0123021002310131-1231322130131002-0102223001003233-0223223333102113): complete subsection reference.

<a id="canonical-2322123032100230-1212230310211102-2033220310120103-2122100002033212-1023221333000200-3002013000301011-0322121113232330-1002120021000122"></a>

<a id="canonical-0122201220311230-3002002211223033-3002322121032000-0022130303323003-1203120213220113-1133323121032323-3110210312012100-2321222212023302"></a>

## request_headers_timeout property — slow_ddos_mitigation / 320333120103 / 4

Type: `"number"`. Optional.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Upstream description:

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The default value is 10000 milliseconds. This setting provides protection against
Slowloris attacks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2000, 30000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-3321022001022222-2212122111020130-3210110301332333-2103103221230003-3320210211002133-3202311210030301-0033010210211313-2003030101212202"></a>

<a id="canonical-0022010201001101-0313022330220013-2021120210312001-2210103032303121-1102222322020333-3313023200123222-1210202313233120-3231133230010312"></a>

## request_timeout property — slow_ddos_mitigation / 320333120103 / 5

Type: `"number"`. Optional.

Exclusive with \[disable\_request\_timeout\].

Upstream description:

Exclusive with \[disable\_request\_timeout\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2000, 300000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

<a id="canonical-3300033033031200-1303221211021022-0311100000310011-2000010221303200-3322002221121221-1213012031201021-0030210301321101-2132222311100231"></a>

## Next pages — slow_ddos_mitigation / 320333120103 / 6

- [slow_ddos_mitigation.disable_request_timeout](resources--cdn_loadbalancer--reference--group-014.md#canonical-3202203332302003-2003031212203232-1211332211112331-3321100223302000-0123021002310131-1231322130131002-0102223001003233-0223223333102113)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3202203332302003-2003031212203232-1211332211112331-3321100223302000-0123021002310131-1231322130131002-0102223001003233-0223223333102113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031333222211302-1323203121223203-2223020110021310-2120101121230313-0201123323131010-0010032203000013-2300123002330003-3303203033302211"></a>

## slow_ddos_mitigation.disable_request_timeout — disable_request_timeout / 332001303003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [slow_ddos_mitigation](resources--cdn_loadbalancer--reference--group-014.md#canonical-1233003020202211-1121000220031121-0002002321202011-3200220102103002-3312211001303013-1111001323333210-3232211013001110-3302023302323112)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-3332032001123310-1001010123211132-1310032101023223-1023003323313332-0330032123100302-1023220133021202-2311123330010300-2112113320131312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable request timeout.

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
disable_request_timeout = {}
```

<a id="canonical-1331320102223233-0303202300332222-2323302200223201-3020132110221322-0110012031202200-2113311011232123-3010002312133123-3233331213201012"></a>

## Direct properties — disable_request_timeout / 332001303003 / 3

This is an empty object or choice marker. It has no direct properties.
