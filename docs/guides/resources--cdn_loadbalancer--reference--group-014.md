---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-2312131312221300-0230011100003021-0021131212211331-0123333322230033-1210310032303322-1320033320200030-2221332222001320-2330310311021003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.http_method` properties

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

<a id="canonical-2000011311220310-3010010100122200-2213332310122103-3010031312312011-2100131103221231-1332113130102301-2020332223002021-0111321302112230"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.http_method`

<a id="canonical-2203232232111313-1221212000123010-3323200113023331-3200321032220321-2201110100201303-2222301123120121-0322202000303033-1213031011010112"></a>

#### `policy_based_challenge.rule_list.rules.spec.http_method.invert_matcher` property

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

<a id="canonical-0010102333021022-1312312211110233-2301231000333301-2033312002002101-2013303130302133-0113220022301202-1011222320210021-1031210211321100"></a>

<a id="canonical-2102313213231110-2313103231123113-3013011013323301-2321011032012130-3210300012101202-0221322112212332-3113102010021201-0221103322113300"></a>

#### `policy_based_challenge.rule_list.rules.spec.http_method.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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

<a id="canonical-1221203300113022-2132333200200013-1003100021200202-1120013111223113-1232003120333102-0311332302111030-2000323030332330-3022032101132020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_matcher` properties

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

<a id="canonical-1021102121333021-3330120312111330-2130202313113013-3311032032313032-0300231230332212-0211103313202313-2122101232301131-3201323100200121"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_matcher`

<a id="canonical-2010112032010113-1330030033132030-1300312212033120-3232102120323323-1030232131012030-3232133302002001-3223122332103033-1021123300301313"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.invert_matcher` property

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

- [prefix_sets](resources--cdn_loadbalancer--reference--group-014.md#canonical-2311012020023332-0201330100311030-3233112032312331-1222302001110020-1110031220112131-3112030003122113-1012032302112032-0312221130221132): complete subsection reference.

<a id="canonical-2311012020023332-0201330100311030-3233112032312331-1222302001110020-1110031220112131-3112030003122113-1012032302112032-0312221130221132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets` properties

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

<a id="canonical-0231213013100021-3330023332021211-1203210100321020-2301231221223102-1300222021131000-0021032223202103-2222100233120220-2233021201321132"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets`

<a id="canonical-0232320111222030-1302203003330311-2331123202101230-3112132310130210-3232231202010102-1230122210230300-0333210333010111-2020101213212102"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.kind` property

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

<a id="canonical-3121132013100330-1220122031222003-2232213001333031-3213320213302331-3201203200212003-0102301321202333-3103310322133321-1101233112000301"></a>

<a id="canonical-3030233132133021-3021023220220203-3011232211222010-3323200132210133-2131322012101122-1131032023200101-2202020311232130-1111321133331001"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.name` property

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

<a id="canonical-0011012202233302-3220331121000213-0120233003300210-1321200212322232-3011023221031233-1113012010233021-3101321232222030-2300331213123313"></a>

<a id="canonical-1010333022100323-2112310011033310-0310003013021231-2232001110001023-2022113213122011-3201203230003321-3220023213131300-2323312312220022"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Optional, Computed.

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
  }
}
```

<a id="canonical-0312000323213231-0111012032202000-3033211033210133-3013230131121200-3213010021120312-1222022132301313-3013202302313023-1002131033012000"></a>

<a id="canonical-1311200113013122-1023220122111011-2320322101211313-0230113102301313-3002202330223003-0031223222003210-0223232010101312-1101231213030212"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.tenant` property

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

<a id="canonical-2300122200321331-3300303111213333-1112202332201133-3110022331232021-1203220200202301-1021203313300132-3102321003101113-3310002112332100"></a>

<a id="canonical-3011210021020201-3102221203112112-2301022223132323-2001130002011302-3302203331211011-0322131230100110-1021100133222121-3110321221201210"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.uid` property

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

<a id="canonical-1012033320010022-2222231020323103-2321000021201310-3031023311331022-0012002023131321-2230011300323100-3130312013222320-2213121223200210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_prefix_list` properties

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

<a id="canonical-0022213012310232-3210021222221233-0330332132332321-2223100031113101-2223113110113123-3013102120021333-2111000233221222-3202033131021100"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_prefix_list`

<a id="canonical-2312232003033113-2302032002221002-3032200012102023-3023122133002223-3011002031300010-2100101111132311-2333320033023322-1002323022212002"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_prefix_list.invert_match` property

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

<a id="canonical-0330231110001002-0332222322223120-2100121122023210-3320332310302010-3233120113221123-0232010201020213-3210122311223300-0000110031131311"></a>

<a id="canonical-1233301313303023-0201011012032333-3021111331110333-2300020203011030-0201132032031113-2000001312002123-2310212121330213-3221020003323201"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

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

<a id="canonical-1323031133100220-3003131300131310-1203112022000311-1230021132323020-2132121211131031-2202311300331212-3111230320020110-3011011131332111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.path` properties

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

<a id="canonical-2113131132203211-1220321131033020-3230103032003001-3123031311130220-0101000203001023-3320310322131300-3111112233310233-3130132222222311"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.path`

<a id="canonical-3310030332322101-3102210022203201-1220320102220233-0013001232011113-1202232320121100-3203202211331303-1333133203321303-2001003203011000"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.encoded_path_matcher` property

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

<a id="canonical-2232233212103003-3233310321133303-3013301322221223-3002001113010110-1211332330201121-3113310021101132-0013003133120322-1321310020132102"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.exact_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-1023110331303112-0102023321222113-0213003120013100-1301323300301220-1112023310321022-0030113121201213-0132123123211033-1231010312223202"></a>

<a id="canonical-0331210032231020-1200123320321223-0100310221032200-0322230031302103-1320311210002222-3223232202231232-2032001220123033-1002121100110323"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.invert_matcher` property

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

<a id="canonical-0311320231223032-2120202323130222-2203102302001202-3220123301312230-0223002331212213-1232312201330221-3231103223100010-2131210333231203"></a>

<a id="canonical-2011101301212203-2011030010130321-1000032021130213-3000212120121301-3331230001332121-3212021311323031-3001133130202303-0233300101031313"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.prefix_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-1333131200321220-0210021133020003-1310122003103122-3312320211030213-1130101133203003-3001321322211311-0033330023013123-0031220001122032"></a>

<a id="canonical-3113302200330222-3121110211310313-1031313330300131-3030132023110201-2022331130333023-2101331030333201-2221333130231332-2003003123010212"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.regex_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-0302133323201220-0320231030111112-1220331012311220-2121311010302022-1212033000301113-3020112222212302-3200123011013011-2020200121131103"></a>

<a id="canonical-3023103333133333-1203022200002101-2032301121001101-0113112123131200-1003010200033111-2313021300003321-0211211231332113-0231102120013133"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.suffix_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-3101323201032123-3130111021201301-2132330011201310-1001121301111111-3030112133003021-0200010300321033-0202010311201222-3220213200013312"></a>

<a id="canonical-2022221022310113-0011221221300231-3222121003330122-3313323033322012-1320321030032132-2200023022322213-3101313011123013-1021333200022213"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.transformers` property

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

<a id="canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params` properties

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

<a id="canonical-1132210231130311-1002030311322103-3230021211120112-3321020321331033-3132303223231231-0322330302023210-0333113213203310-0303033201012100"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.query_params`

- [check_not_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-1312300330202202-0230330032101303-2331213220023113-3312230312301330-3331100020223113-1303131000020002-2031223202300310-1031130011210022): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-0011200213331102-3303001022301002-0122112221113011-1101032132033112-2322010100320111-1203021310120331-0020211133132211-2133230123112201): complete subsection reference.

<a id="canonical-3102010112201132-3222113311210212-1102022201021013-0212100133323203-2302133113021031-2021311301103220-0223311303200101-0312321030212100"></a>

<a id="canonical-2222012103022311-0030322220330112-1102120332311130-0021023322323202-3002021122122121-0332302132022332-2303133301303121-2331030231331102"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.invert_matcher` property

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

- [item](resources--cdn_loadbalancer--reference--group-014.md#canonical-0300132120303011-2222013321322121-1103100003033103-0130000011023130-3102203201303332-3231303212012130-0201320300021112-0303323231120233): complete subsection reference.

<a id="canonical-0121103323123020-0322032200300210-1010000123320232-2213220011002320-0201100133022202-2313102110311320-1011212122101331-3220223033023222"></a>

<a id="canonical-0133030032331312-1233120133000302-3213300120100313-2123023110210031-1332300102200210-3123021122200312-2331021102302123-0300313120210132"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.key` property

Type: `"string"`. Optional.

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

<a id="canonical-1312300330202202-0230330032101303-2331213220023113-3312230312301330-3331100020223113-1303131000020002-2031223202300310-1031130011210022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.check_not_present` properties

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

<a id="canonical-0011200213331102-3303001022301002-0122112221113011-1101032132033112-2322010100320111-1203021310120331-0020211133132211-2133230123112201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.check_present` properties

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

<a id="canonical-0300132120303011-2222013321322121-1103100003033103-0130000011023130-3102203201303332-3231303212012130-0201320300021112-0303323231120233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.item` properties

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

<a id="canonical-0320302312110211-2213212022021231-0002001211221020-1213133231032100-3113200102120021-2112132033230013-1110222310221103-0210313323303010"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.query_params.item`

<a id="canonical-2302023302020011-3110113331121322-2323311121120322-3002132110300213-0301221333230233-3202122303131011-2031213033220032-2023112201313332"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.exact_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-0303000332223232-1232123101211223-1031132013231223-1313113202103030-2031230231323030-0201313012301232-0011021013103222-0022313031313211"></a>

<a id="canonical-2230113320132311-3112331203320201-2202013010133122-1331313012300010-1100330111030201-0113230033203030-3003320213102132-1132100133001000"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.regex_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-1303101121133330-0033031313313021-2302331021320122-3200200211103111-0112311321133021-3332111210220123-2120112213000020-0032123103120022"></a>

<a id="canonical-3020111202023211-1201032333023302-3312132103033210-1000332130003010-1111210023110000-1213010123330302-2120222123221121-1122203303232122"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.transformers` property

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

<a id="canonical-3232022003302213-0032131233110110-3231221110233032-1010300113202110-1012221322120230-1322201202221221-2233333330333200-3030322000303331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher` properties

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

<a id="canonical-1223320300213331-1200031102303232-0230131000301003-1010132103310221-2303123111012030-3203032231221000-2313232221223113-1332031033323203"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher`

<a id="canonical-0332102033203002-1300033000302223-1320111233030302-2011330223122011-0202033001013312-3103112203002020-1113210220023011-1310023101331132"></a>

#### `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher.classes` property

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

<a id="canonical-0232122031310203-1130322311131121-1133103013313321-3121332003023112-1013321200030331-2202211230101021-0232311203201112-2203200311220213"></a>

<a id="canonical-2021320223331302-1012031110000302-1230311310202221-0032321113110212-2013021311003312-2010200121103002-0100131222101020-2201313210333011"></a>

#### `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-2101002303102000-2012230021120100-3120111331223100-1023203333233010-3323322133212220-1101133231221321-1113220103111033-3000333121121203"></a>

<a id="canonical-0311221132103132-3311012203132000-1231022131021130-2121201313312331-1021120010023320-3213121120101101-1222021223000233-0232311230020123"></a>

#### `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-2020113203001122-0131003133101010-3112033112103122-1230132210230022-1110101121230003-3011322320122320-2001301222311323-1030230313310012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.temporary_user_blocking` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.temporary_user_blocking

<a id="canonical-1201223313121033-1103213030333221-1331232332133120-3021210133032030-0032312022010311-2131030221021131-2322001320123223-3002233120103201"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3110003331013331-3130133013301231-3010203210213123-2222111310330220-1203123301030033-3221112233210000-2113132010310313-0201133201120020"></a>

### Direct properties for `policy_based_challenge.temporary_user_blocking`

<a id="canonical-0010103303001310-0213230232201003-3123212123321232-2211122222002222-1330020101323220-1302331232121123-2113123000330012-2233321033223010"></a>

#### `policy_based_challenge.temporary_user_blocking.custom_page` property

Type: `"string"`. Optional.

Custom message is of type . Currently supported URL schemes is . For scheme, message needs to be
encoded in base64 format. You can specify this message as base64 encoded plain text message e.g.
'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '&lt;p&gt;
Blocked..

Additional upstream details:

Custom message is of type \`uri\_ref\`. Currently supported URL schemes is \`string:///\`. For
\`string:///\` scheme, message needs to be encoded in base64 format. You can specify this message as
base64 encoded plain text message e.g. "Blocked.." or it can be HTML paragraph or a body string
encoded as base64 string E.g. "&lt;p&gt; Blocked &lt;/p&gt;". base64 encoded string for this HTML is
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

<a id="canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- protected_cookies

<a id="canonical-2223231201320222-0020301102201100-2031133233311002-1000230332031133-1303122030112310-0231202302013302-3122033213202302-1312113133230002"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
protected_cookies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313333230210103-0130123231011131-3012330332330013-3123011131033220-3133022002233021-3113131000213013-1221201010201311-1330033030030030"></a>

### Direct properties for `protected_cookies`

- [add_httponly](resources--cdn_loadbalancer--reference--group-014.md#canonical-1222133013311211-0333131120023113-3121113303130123-0312213332010220-1233033023120021-0221013100020132-1321230210000300-0210201010300313): complete subsection reference.

- [add_secure](resources--cdn_loadbalancer--reference--group-014.md#canonical-1132000333120232-3302322021300311-2313311210013112-0322232321221301-0212300032233331-1312322121103103-2330002320301013-1320232002031322): complete subsection reference.

- [disable_tampering_protection](resources--cdn_loadbalancer--reference--group-014.md#canonical-3300333222232023-0313031213010202-2102000001032100-2112313322020220-1233220033132310-2111330213212301-3323312202103000-1011031123202323): complete subsection reference.

- [enable_tampering_protection](resources--cdn_loadbalancer--reference--group-014.md#canonical-3022000331032231-1230000032102312-1211131333230321-3212123301010213-3131211011121332-2222331321231220-3023220111001123-1133210011210102): complete subsection reference.

- [ignore_httponly](resources--cdn_loadbalancer--reference--group-014.md#canonical-0322013100232300-3320012232320223-0302321121132122-2022111333203301-1030032201101211-0022020002113133-2123320001003022-0233133312121322): complete subsection reference.

- [ignore_max_age](resources--cdn_loadbalancer--reference--group-014.md#canonical-3331121030110230-3113323312022330-0300013331330211-3211201022230320-3002302301000022-3022011032202013-1012021332121122-1231123101031131): complete subsection reference.

- [ignore_samesite](resources--cdn_loadbalancer--reference--group-014.md#canonical-1001312120203320-0332332333210221-0021331233302310-2213332032130322-0332010223111130-2310201113022012-3033320130302030-0322220300311012): complete subsection reference.

- [ignore_secure](resources--cdn_loadbalancer--reference--group-014.md#canonical-2331301203101122-2122302001110110-2001122133213231-0103100311122302-1323130110210320-3002010000212312-2000232110301312-2130323323010010): complete subsection reference.

<a id="canonical-2012210210032003-3211210200031031-3202023210102201-2031322310333323-3031303112132202-3311331102301002-0121333302031110-1011001302130301"></a>

<a id="canonical-1200331313333211-3310003000030103-2030101213003232-2112111211232101-2131233333210222-1210300311013333-3120320230031331-2300320111311301"></a>

#### `protected_cookies.max_age_value` property

Type: `"number"`. Optional.

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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-0301330121223030-0212231032220331-0330232212023230-0210120333013003-0232030123123222-2323313322023331-1132022332021300-0011333122033301"></a>

<a id="canonical-1200312033123300-3301000312201220-2021303012101001-0123102110203202-3220000312303003-0011010122121102-1113312300013010-0132200230021130"></a>

#### `protected_cookies.name` property

Type: `"string"`. Optional.

Cookie Name. Name of the Cookie.

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

<a id="canonical-1222133013311211-0333131120023113-3121113303130123-0312213332010220-1233033023120021-0221013100020132-1321230210000300-0210201010300313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.add_httponly` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.add_httponly

<a id="canonical-0110023213031302-2102002121002103-2222332030011222-0022122301131021-2132231032111231-1101003022301212-3131302210022012-3130110113212001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132000333120232-3302322021300311-2313311210013112-0322232321221301-0212300032233331-1312322121103103-2330002320301013-1320232002031322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.add_secure` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.add_secure

<a id="canonical-3210033102032123-3311002120120202-3030312001111322-0022332111131033-1231323210133321-2132020133212221-1112121010312212-3301233312110122"></a>

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
add_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300333222232023-0313031213010202-2102000001032100-2112313322020220-1233220033132310-2111330213212301-3323312202103000-1011031123202323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.disable_tampering_protection` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.disable_tampering_protection

<a id="canonical-3033132112112320-1011212120130211-0323321020112001-3032300020033031-3102233200102132-1200000303031213-3312212020301311-2132301230101232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable tampering protection.

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
disable_tampering_protection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022000331032231-1230000032102312-1211131333230321-3212123301010213-3131211011121332-2222331321231220-3023220111001123-1133210011210102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.enable_tampering_protection` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.enable_tampering_protection

<a id="canonical-1011002213133222-1322310130030031-2112321030300220-3002113113003330-3031212122003131-2321031133323203-0111313130332032-0221121313212303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable tampering protection.

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
enable_tampering_protection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322013100232300-3320012232320223-0302321121132122-2022111333203301-1030032201101211-0022020002113133-2123320001003022-0233133312121322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_httponly` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.ignore_httponly

<a id="canonical-0312331323131101-2112322330310133-3031100211220113-3232233321112201-3312013113023300-0110203202223303-2322222230010200-0020231133201202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331121030110230-3113323312022330-0300013331330211-3211201022230320-3002302301000022-3022011032202013-1012021332121122-1231123101031131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_max_age` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.ignore_max_age

<a id="canonical-3203231010202120-3020013211132010-0323322130221002-2131111011212111-2332101200120112-3201311233223303-3202232220310232-3113032122312121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore max age.

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
ignore_max_age = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1001312120203320-0332332333210221-0021331233302310-2213332032130322-0332010223111130-2310201113022012-3033320130302030-0322220300311012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_samesite` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.ignore_samesite

<a id="canonical-0200201121101123-2013330201300120-3033311302212030-1323332021020032-3302221332131130-0232200123110213-3033011232222021-1312301232200120"></a>

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
ignore_samesite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331301203101122-2122302001110110-2001122133213231-0103100311122302-1323130110210320-3002010000212312-2000232110301312-2130323323010010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_secure` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.ignore_secure

<a id="canonical-1301203101223311-1132110222022301-3311320003230132-1311132323112002-0230330320323222-3120123003122011-0330110120021321-2322311220311232"></a>

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
ignore_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101022332031001-0323332310312332-2113121333322000-1313013133120330-0132010202201213-1330301200203231-1021011323232122-0010222230103230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.samesite_lax` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.samesite_lax

<a id="canonical-0112113002012121-3120032321130213-2322222022122032-2122330020000202-2202311210022201-2110313210021313-2233103002322110-1230102332323233"></a>

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
samesite_lax = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232203201132302-1331211313003233-3213032100202223-1223313032122220-2211332102102021-3133010113022030-3321023133303011-2023010011020020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.samesite_none` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.samesite_none

<a id="canonical-1102110013020121-2333000002012003-0112033002010123-1200122231033312-2223103103112103-1200203123110111-2211020221303321-0323220202211120"></a>

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
samesite_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003230130333321-0023020200021211-1100010112302123-3132331221102033-1313232222231120-2233011102220102-0211111032322201-1301311002131023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.samesite_strict` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122)
- protected_cookies.samesite_strict

<a id="canonical-3132232222231122-1202121321110033-0012220121023132-0300320310313331-3020032031003020-3223222100030233-0012220221202212-2102233220113311"></a>

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
samesite_strict = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit` properties

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

<a id="canonical-3023130321310200-1311313112001322-1023223002200032-2201302330030101-3120303121312333-2200121112033212-3121212013120311-1120122230112120"></a>

### Direct properties for `rate_limit`

- [custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-3211202232232112-1332230121300022-3332022212230331-2133332330303223-2213132312211233-3323123331302103-2333032000301112-3332003332201331): complete subsection reference.

- [ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-0131111121111212-3101110311102221-2133123013002303-2210110213111330-1133030303203122-0003130031330133-1021013031213230-3221000033233130): complete subsection reference.

- [no_ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-2213010213120023-0100103320013021-3122212310121112-1322112200122212-2010212132123210-0012012333133011-2020101110020110-2132200212132230): complete subsection reference.

- [no_policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0132002221322313-3000222211302221-0201132202300003-2010311200322230-2303312220332001-0221100132130231-0122131031023220-0011032210230110): complete subsection reference.

- [policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0332012013133002-0110210320323201-1003122022020300-0120220113202102-3111330131211212-2231211322332003-0010201011000213-2030331032023213): complete subsection reference.

- [rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211): complete subsection reference.

<a id="canonical-3211202232232112-1332230121300022-3332022212230331-2133332330303223-2213132312211233-3323123331302103-2333032000301112-3332003332201331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.custom_ip_allowed_list` properties

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

<a id="canonical-3203223300133100-2022022313033223-0002102310010310-3012203100120131-0213023012002232-3231231113132213-2133021001131233-1221223231232113"></a>

### Direct properties for `rate_limit.custom_ip_allowed_list`

- [rate_limiter_allowed_prefixes](resources--cdn_loadbalancer--reference--group-014.md#canonical-1133001332311010-3112323220120313-1220011323101130-1332132211101130-2312133303011233-0221102210200211-1302201101231200-2023100031230022): complete subsection reference.

<a id="canonical-1133001332311010-3112323220120313-1220011323101130-1332132211101130-2312133303011233-0221102210200211-1302201101231200-2023100031230022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes` properties

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1301131010103032-0113333221200231-2102333013123010-2222101123021013-3023013113233202-0213101121010023-0111023213031202-0002320132003000"></a>

### Direct properties for `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes`

<a id="canonical-0101022010001110-2303231222213232-3321221033002203-2302213112133330-3120032102013122-2302220230110320-3011121322003122-2310011112020233"></a>

#### `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.name` property

Type: `"string"`. Optional.

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

<a id="canonical-3001022021112321-0223003303221323-0300222013010222-3012330322000302-3132333132112313-1320020321010030-3130233003233301-1133202321022133"></a>

<a id="canonical-0022101210333311-1233021023011133-3000312221221102-1011031320311320-2211300213231121-3302132300300011-2121202122011221-1213223312332213"></a>

#### `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2101303110223230-2102213211123333-3012322212303122-2032302203303201-3200320000333111-3311030210122231-2130033032313110-2301231202121221"></a>

<a id="canonical-2013002021322120-2131303122233210-2122233211201232-2213121321201102-0020211030021113-1110100230233310-1310222331032013-2020230113313003"></a>

#### `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-0131111121111212-3101110311102221-2133123013002303-2210110213111330-1133030303203122-0003130031330133-1021013031213230-3221000033233130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.ip_allowed_list` properties

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

<a id="canonical-2222333021200020-1223221302021213-0203330021220201-1121332032031030-0011220313312320-1132012021230000-3233303021130200-2003203211300333"></a>

### Direct properties for `rate_limit.ip_allowed_list`

<a id="canonical-0311010123110110-1213103003232003-2133230131033120-0021133112301122-3000231202331313-1033220012022103-2213021123002012-3011013010303233"></a>

#### `rate_limit.ip_allowed_list.prefixes` property

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

<a id="canonical-2213010213120023-0100103320013021-3122212310121112-1322112200122212-2010212132123210-0012012333133011-2020101110020110-2132200212132230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.no_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- rate_limit.no_ip_allowed_list

<a id="canonical-0211122133022010-2030330003021210-0111021020303323-3132000021023102-3001020023320322-3301333012113322-3010203122321311-1202101023310321"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
no_ip_allowed_list = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132002221322313-3000222211302221-0201132202300003-2010311200322230-2303312220332001-0221100132130231-0122131031023220-0011032210230110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.no_policies` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- rate_limit.no_policies

<a id="canonical-0201120202210222-0020323201213220-2301033103202302-2202321323132201-1133113023120222-0022213321131113-2113223103033102-1023302330100313"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no policies. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_policies = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332012013133002-0110210320323201-1003122022020300-0120220113202102-3111330131211212-2231211322332003-0010201011000213-2030331032023213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.policies` properties

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

<a id="canonical-0330230302303320-1113033221020012-0301033323333032-2013321210220001-0203131313023213-2323202000203210-0301102313332200-2122310230311330"></a>

### Direct properties for `rate_limit.policies`

- [policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-3031012130111313-0333133110300122-3212021223302331-1010031320331132-3220011102222022-1030211202302321-3011132103101230-2210310200333303): complete subsection reference.

<a id="canonical-3031012130111313-0333133110300122-3212021223302331-1010031320331132-3220011102222022-1030211202302321-3011132103101230-2210310200333303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.policies.policies` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-0332012013133002-0110210320323201-1003122022020300-0120220113202102-3111330131211212-2231211322332003-0010201011000213-2030331032023213)
- rate_limit.policies.policies

<a id="canonical-0113000321133111-2002020013322023-1001010232211332-3133333113132203-3001203310120031-0030001301030032-1211312033033022-0122002223311101"></a>

Type: `"object"`. list nested block, Optional.

Rate Limiter Policies. Ordered list of rate limiter policies.

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

<a id="canonical-1130130322101101-3033100123321302-0122120313320003-2013023132012023-0233211013001201-0211021233002213-1303002003120201-2103321012330311"></a>

### Direct properties for `rate_limit.policies.policies`

<a id="canonical-0110320110211230-3103100313112231-0120102030221013-1103011231012330-2333213023220200-0032101233212120-0303110332000310-0232120100321202"></a>

#### `rate_limit.policies.policies.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1110123333110001-0203120303231123-0222120223131230-2312210010120120-2302022302120203-0312231331010130-0220330022233123-3021230332122033"></a>

<a id="canonical-0123330000001032-0232113120021012-2131102132120223-3001022112332131-2133010023121110-1231333013203300-1020310213022122-2321011010030121"></a>

#### `rate_limit.policies.policies.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-3001121002003321-2230001000330311-2211032032120130-0203110213303322-2103121132131212-0030201122320003-3330221233132232-1033320302300022"></a>

<a id="canonical-3200101130021033-0110201212110220-0222023313002210-1121003101213012-3100231113331021-2130311230321121-0123003023021210-0312210120220213"></a>

#### `rate_limit.policies.policies.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- rate_limit.rate_limiter

<a id="canonical-3202232323000300-1302010301323202-2112210310332012-2013111231122303-2323230010231120-2110001103122020-2032132221113323-0222333022323332"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2220310321311020-0031030011111222-2122002223130020-1213113003120230-2031213121231203-3203302013210113-1011020023313320-2323010312022323"></a>

### Direct properties for `rate_limit.rate_limiter`

- [action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131): complete subsection reference.

<a id="canonical-2121003302013021-2121130112031230-3123332123101323-3221123222213223-0330023100113303-2330213013221210-3202103013331021-0300232022011222"></a>

<a id="canonical-1312131121103011-0231000011120310-1002113303231201-2012022001023130-3232302200200210-0132213001223322-2021222200001030-3110132021102322"></a>

#### `rate_limit.rate_limiter.burst_multiplier` property

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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [disabled](resources--cdn_loadbalancer--reference--group-015.md#canonical-0121113003211121-3000233231212222-0210322310333131-1232221120110102-0323331121001101-3300320021101330-3211100221101210-3303200202212010): complete subsection reference.

- [leaky_bucket](resources--cdn_loadbalancer--reference--group-015.md#canonical-2121301302012332-0200200003121032-3002102313230233-0221300213312113-3323111113013323-2321301313301030-1201012223223023-2123010120320301): complete subsection reference.

<a id="canonical-0321222321313130-2131301201011131-1101123011120111-0123310301230323-0130032012210031-0121003212300313-1232132202312211-3313202303300210"></a>

<a id="canonical-2110012021001222-1331102000102322-3322232121111310-3312333222310232-0101321033231331-1222202231002232-2232121023320020-1111032331120321"></a>

#### `rate_limit.rate_limiter.period_multiplier` property

Type: `"number"`. Optional, Computed.

Setting, combined with Per Period units, provides a duration. Server applies default when omitted.

Additional upstream details:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [token_bucket](resources--cdn_loadbalancer--reference--group-015.md#canonical-3211020022010202-1101201222211102-3010213311221001-2102010330320122-0120301001332101-0332320102011122-0321200121122313-3032003110020000): complete subsection reference.

<a id="canonical-2113133301210033-2133213221021332-1112211323000030-2300100021213032-1131011033001301-2301320222133020-3022112322223110-0323001230233001"></a>

<a id="canonical-0132311222220200-1000120210111023-0020223132201232-3220213013123200-2020120210110003-2023012212203333-1323332020202213-2012012231230221"></a>

#### `rate_limit.rate_limiter.total_number` property

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

<a id="canonical-0022123123320322-1312312231201013-2321000220323303-0213132102121003-1122120032233020-0310300113303310-0020310120310031-2332203232002323"></a>

#### `rate_limit.rate_limiter.unit` property

Type: `"string"`. Optional.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

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

<a id="canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block` properties

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

<a id="canonical-3022323122112022-1133131113332201-3221032210123200-1112100212210330-1300231201110121-2013322011212312-3211330122130010-0100120001310230"></a>

### Direct properties for `rate_limit.rate_limiter.action_block`

- [hours](resources--cdn_loadbalancer--reference--group-014.md#canonical-0001220131200010-1033233213023323-3333300333212200-2133013021233210-0301302322310330-3130330031202120-2310111303100313-0313332031331111): complete subsection reference.

- [minutes](resources--cdn_loadbalancer--reference--group-014.md#canonical-2232102003211200-3310003313003312-0012200212113100-0002323313331101-1020102203000121-1302032211020003-3000302231323112-1301000013201333): complete subsection reference.

- [seconds](resources--cdn_loadbalancer--reference--group-014.md#canonical-1320212332333121-2121103102302021-2000231022122212-1221020333221331-0120202330203000-1000300113331220-0211232131110222-2311323201323002): complete subsection reference.

<a id="canonical-0001220131200010-1033233213023323-3333300333212200-2133013021233210-0301302322310330-3130330031202120-2310111303100313-0313332031331111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.hours` properties

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

<a id="canonical-3330103111232101-3132233001303331-2222303213302230-3133203101103011-0002013223121020-1200133123211331-1313111122221333-3111101303113100"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.hours`

<a id="canonical-3132212132133200-1032200323323022-1102313200331323-2333002220222202-2111313231323230-0313212020210323-0233100320002111-3330002311211120"></a>

#### `rate_limit.rate_limiter.action_block.hours.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  }
}
```

<a id="canonical-2232102003211200-3310003313003312-0012200212113100-0002323313331101-1020102203000121-1302032211020003-3000302231323112-1301000013201333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.minutes` properties

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

<a id="canonical-3320001103010013-2123312000020332-0221013030031313-0112133130111222-1221212333123221-3101310031123010-2031220232203020-1321012131321323"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.minutes`

<a id="canonical-0011330110031122-2111300321230301-2010003300013220-2321013203003023-1233200032301030-0110003120000232-0233220112012303-0023133110231233"></a>

#### `rate_limit.rate_limiter.action_block.minutes.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  }
}
```

<a id="canonical-1320212332333121-2121103102302021-2000231022122212-1221020333221331-0120202330203000-1000300113331220-0211232131110222-2311323201323002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.seconds` properties

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

<a id="canonical-3112002010011011-0031211333010321-0032300003201331-1200001222200130-2213330022332331-0013212103113003-2231213200332320-2312220113312100"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.seconds`

<a id="canonical-1232220201100310-3222033101201002-1113033122230001-3023003203210332-3001000320210310-1112102100100001-1012323033202123-0323303120223132"></a>

#### `rate_limit.rate_limiter.action_block.seconds.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```
