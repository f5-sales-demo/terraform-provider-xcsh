---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2210031331220231-0003300122101220-1122220232123210-2313022300122320-2103112111233303-1230222210232131-1210313323210111-3033132132213210"></a>

## `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-1112201310023113-1302120212211300-1210321011312300-2032300112230112-1003233022203030-2331212302112213-2313130121113332-2211112020112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path

<a id="canonical-3130112331122121-2210112233222001-0010021322203002-2032300030023003-3003120200121130-1022132103133120-1103003231123101-3212330321120102"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2100102221322223-0031103123330013-3310322212030322-0010031023130230-3332323332103132-2233213330133031-3030013130231313-0011220303330233"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path`

<a id="canonical-0202322332130001-2011001222001022-3000021211301210-2000233300302123-2202010222112313-0200133030323000-0211323003230330-0312220313032210"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1201322320211211-1013010300031201-2210230202332013-3320021211223022-2101103301123033-2032130321303122-3313221032320222-2233101001311100"></a>

<a id="canonical-1013233013221232-0320131332322301-0001032202132021-1030001212121313-2331122311313320-0230022321230101-3012001010330311-1222033302131110"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-1231032203230303-2201013122303103-2302213333210320-3222003121021322-0023110330320133-2332220331013030-0102222102202233-1132122302200113"></a>

<a id="canonical-2103200330033233-3020033213302021-2031223020203130-1001203000012303-0113030030313021-2000311332231302-0031200321122130-0022132002323303"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules

<a id="canonical-0300001222100111-2001221002103222-3303100133210330-2131312101101032-2301131012011013-2331323300033121-1000033001103220-3211301003222132"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1322003031301323-1110311210230010-2332323310012331-1021021201031130-1031031231320201-3333032010030103-3020201000010200-0210200202011322"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.rules`

- [any_domain](resources--http_loadbalancer--reference--group-013.md#canonical-1200331032112302-0133002312021221-0013103132010212-3331330232303022-0333122302131133-2023201122233322-1101010330031001-0321133001110013): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-013.md#canonical-1301233102201233-0233012311103202-1023210002020011-2233110231333123-1300111222303013-1031033030030211-0220012203013020-3110201010213100): complete subsection reference.

<a id="canonical-0032102021221300-1233311231220310-1211323321213221-0021310100213020-2120330103113022-0201111311133222-3032123133012203-3332233313003212"></a>

<a id="canonical-3312132330310303-2223221102111203-0031333332100200-2333001010033023-3220330023330032-2010013200232222-1011130002113023-3112033313203130"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.javascript_location` property

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["AFTER_HEAD","AFTER_TITLE_END","BEFORE_SCRIPT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-013.md#canonical-1002301111101310-2233213103220330-3230330101321010-2132222110030233-2122312323031000-1131103001033230-0320321102312220-3111312232211001): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-013.md#canonical-3303333331111031-2310213001310120-1222102313203113-0003112123311101-2133101020201311-0222300312012013-0230312311102330-1021223130331333): complete subsection reference.

<a id="canonical-1200331032112302-0133002312021221-0013103132010212-3331330232303022-0333122302131133-2023201122233322-1101010330031001-0321133001110013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-013.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain

<a id="canonical-2101330300232023-2302033033302313-3030300212032320-0232121012022330-2010222000322100-3102110220300003-2332310313311303-2110100012013022"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301233102201233-0233012311103202-1023210002020011-2233110231333123-1300111222303013-1031033030030211-0220012203013020-3110201010213100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-013.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain

<a id="canonical-3100222123222003-3131313010022211-2321112222201103-0333010121231000-3213313203021211-3133113113310202-1111223101331011-0231232310301013"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2210113030312122-1203122332001031-3021311330003230-0323302313010020-3200012323100010-3312203311102132-3302101201100330-2300031323221133"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain`

<a id="canonical-2301020021133000-1311331003200221-1021123003121313-0333033000113011-2111332011011201-0120110021230131-3211203331213101-0231202213300132"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3120323332312303-0023133121301200-1020312121101302-0231020221323111-0333330222001100-3220023330020011-3212322333101133-0213020023032100"></a>

<a id="canonical-1312131322200111-3310312210131133-0333332232121003-3211113002032222-0121121231231102-0001030020021232-1110110231210222-1111121132212131"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0112002321111110-2033300100331322-3101000332013221-1131002222112000-2031311130332331-0321010133331032-3001220220013320-3032111232021333"></a>

<a id="canonical-3132020012213022-2131211202031232-1112111331323103-1112110331301233-1301212203131123-3201333101032203-3103101203332130-0012221020113310"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1002301111101310-2233213103220330-3230330101321010-2132222110030233-2122312323031000-1131103001033230-0320321102312220-3111312232211001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-013.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata

<a id="canonical-3000030203002212-0102322000113032-2211112113101102-3100023001122112-1232121300320020-1032100013032332-0011222110103012-0222201313223332"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1001032102200011-2203300003332231-1010013123201100-3333023303000120-3022100231113301-1231333332123013-2013030020133311-0322301202032203"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata`

<a id="canonical-2132101122101101-3102021221230003-3303220312132020-0330301312102103-0230302023002022-2102213111203231-1120200323130130-3101030022123201"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3030301101112321-0112122101310022-3032010211120313-0022300201021322-1230010023013033-3220130332003222-2122212320321230-0003011032300222"></a>

<a id="canonical-1333302020123222-1011113200122110-3333111310321032-1202301322322230-0022332013222032-2113301002223120-2231333113320030-1131022000031321"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-3303333331111031-2310213001310120-1222102313203113-0003112123311101-2133101020201311-0222300312012013-0230312311102330-1021223130331333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-013.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path

<a id="canonical-0203203313211031-3310330102131011-0330120301012022-0120002330032330-1001100103301023-3300302021222331-2032323313311020-0020113010030201"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3110000032133001-2001121000331023-1031233200023202-1201321102030221-3330331121310320-1033211212211213-2223331201002211-3011201131230220"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path`

<a id="canonical-2010231332120320-3231002102112133-1012003033122200-0330102123222131-1120010010110312-2131121201220102-2213133123323301-2113222130000101"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3302333312213312-3111222031131202-3332110232232030-3302022303222312-3011311133210302-0301223211303120-3112212201221001-1011020011110121"></a>

<a id="canonical-0102333021112113-0310033113333210-3333310021102213-1023013200300122-2231330312313013-0020313002130321-3210002302313121-2313312020022021"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-0321013030320331-2231312200002103-0112102231032021-2330032132322113-2020332320031311-1113311133012010-0323223122101203-2012220101323122"></a>

<a id="canonical-2301201222223113-2033201011231021-3023212303212331-0233102112033233-1001123313202322-2113011332231321-1113022130123120-0112113220132202"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0032332312200101-3012322220030112-1320110202102312-0220201232102020-2202130130221013-0301003122332030-0331121213303131-0133333302230210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.web` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- bot_defense_advanced_protection.web_only.web

<a id="canonical-3202100120323111-2013333220112231-3133232100133331-2320300122131320-2120011112112320-3000233012101323-1022301103101100-0221220221101013"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
web {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213121332130200-3301003331022010-3211120131001333-0211221013230000-3201302311010032-1010210220203001-1011130201223232-0311332132013002"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.web`

<a id="canonical-2120213121111301-1003213333120030-3023233120030002-0001022120030022-2222033223320000-2131101220003030-1331220330012112-2321100310223113"></a>

#### `bot_defense_advanced_protection.web_only.web.name` property

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

<a id="canonical-0132330101223222-0211111122322110-2323103203103202-2110133311330101-2302132101030310-3212303210220231-1002133313101122-3323013102220222"></a>

<a id="canonical-2201111202123322-0131133032333012-0031100331012023-3033100213011103-3310110030132220-1010300323323323-3201313022133303-2313122330303222"></a>

#### `bot_defense_advanced_protection.web_only.web.namespace` property

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

<a id="canonical-2313221322230111-3031232211012123-3223121130213320-1001233320223133-1220120211001103-3222332102131211-3310202103321321-0112021031323020"></a>

<a id="canonical-3322322023201121-2120011233323222-0123022200312333-1132032121103101-0322230011310212-1202323301231320-3200202003010021-2202110313330101"></a>

#### `bot_defense_advanced_protection.web_only.web.tenant` property

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

<a id="canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `caching_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- caching_policy

<a id="canonical-3203201312330020-1323212233312020-3213030330312302-3201133202232001-1012201201233023-0001210000312110-2333103330300332-3113013310222320"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: caching\_policy, disable\_caching; Default: disable\_caching\] Policy configuration for
this feature.

Additional upstream details:

Caching Policies for the CDN.

Receipt-pinned upstream constraints:

```json
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

- [caching_policy](resources--http_loadbalancer--reference--group-013.md#canonical-3203201312330020-1323212233312020-3213030330312302-3201133202232001-1012201201233023-0001210000312110-2333103330300332-3113013310222320)
- [disable_caching](resources--http_loadbalancer--reference--group-017.md#canonical-2310010330303022-3212102030230223-3311321133320300-3300013001030002-1102021010321320-0321310020220321-0011203222103312-1122212300121322)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
caching_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2331130101300201-2212310330121222-1230313121111023-3202320001101212-3122300113130323-3103333220130132-0013212011102003-2112132133122232"></a>

### Direct properties for `caching_policy`

- [custom_cache_rule](resources--http_loadbalancer--reference--group-013.md#canonical-1112323230330001-3210133301302130-1001103132120102-1113111200333202-2010211133110132-3203313202133021-1030000033033020-1212102233212101): complete subsection reference.

- [default_cache_action](resources--http_loadbalancer--reference--group-013.md#canonical-2102303123010113-2230132023113110-0301132233211000-0123111331333002-0133122303200233-0021301123001210-2212032330201311-1313032131321121): complete subsection reference.

<a id="canonical-1112323230330001-3210133301302130-1001103132120102-1113111200333202-2010211133110132-3203313202133021-1030000033033020-1212102233212101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `caching_policy.custom_cache_rule` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [caching_policy](resources--http_loadbalancer--reference--group-013.md#canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331)
- caching_policy.custom_cache_rule

<a id="canonical-1030232001203231-1120322313102322-1310033221321020-0311213221031110-3031322212233113-0200000033211102-0333322022102132-2020203303122311"></a>

Type: `"object"`. single nested block, Optional.

Custom Cache Rules. Caching policies for CDN.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1310310201300323-2201120202321021-3212123023303113-0101221111301332-0121232103001013-0332333221001210-3222323123021233-0230301101022302"></a>

### Direct properties for `caching_policy.custom_cache_rule`

- [cdn_cache_rules](resources--http_loadbalancer--reference--group-013.md#canonical-0320220112122003-3200020033122332-1301313302000102-1122133000313102-1020022313003231-1322102210203301-3010230133213232-3003132230022032): complete subsection reference.

<a id="canonical-0320220112122003-3200020033122332-1301313302000102-1122133000313102-1020022313003231-1322102210203301-3010230133213232-3003132230022032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `caching_policy.custom_cache_rule.cdn_cache_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [caching_policy](resources--http_loadbalancer--reference--group-013.md#canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331)
- [caching_policy.custom_cache_rule](resources--http_loadbalancer--reference--group-013.md#canonical-1112323230330001-3210133301302130-1001103132120102-1113111200333202-2010211133110132-3203313202133021-1030000033033020-1212102233212101)
- caching_policy.custom_cache_rule.cdn_cache_rules

<a id="canonical-1322101121113022-0022112212033033-2230311211313132-0323131213121103-0321231323313210-2310033332000210-3222013302001212-3122113031132212"></a>

Type: `"object"`. list nested block, Optional.

Reference to CDN Cache Rule configuration object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0112310103300000-2022001001033331-3211030222211120-1210012022023010-0122030321303100-2133110122213101-0102323011030021-2132111013300011"></a>

### Direct properties for `caching_policy.custom_cache_rule.cdn_cache_rules`

<a id="canonical-3221120010133022-3232223032202133-2331030122231203-0330331202132213-1322020223203130-3112310002010122-3110022232000222-2331131223302123"></a>

#### `caching_policy.custom_cache_rule.cdn_cache_rules.name` property

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

<a id="canonical-3132201313211200-3032331220011120-1111310031112021-1002001101321202-1030122332302323-1212331000023103-0111311222012303-2222311020223323"></a>

<a id="canonical-0232203033002223-0130221312223302-0113323132110030-2030203321223020-1133110221200321-2321301122210113-2321010023031013-3221011112331013"></a>

#### `caching_policy.custom_cache_rule.cdn_cache_rules.namespace` property

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

<a id="canonical-0002223230122003-0012203300221322-2113332020033120-0031113232223031-0123110012231033-0121221120313310-0333203012322332-3101222332330011"></a>

<a id="canonical-3120301020020103-1322111302202210-0033321233323321-2301332221223101-3022311033312323-0113001221132111-0003300021230213-1003131231230131"></a>

#### `caching_policy.custom_cache_rule.cdn_cache_rules.tenant` property

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

<a id="canonical-2102303123010113-2230132023113110-0301132233211000-0123111331333002-0133122303200233-0021301123001210-2212032330201311-1313032131321121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `caching_policy.default_cache_action` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [caching_policy](resources--http_loadbalancer--reference--group-013.md#canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331)
- caching_policy.default_cache_action

<a id="canonical-0231101131123021-0111200113000033-1303222201222132-0310210233310203-2212213300120232-3002021013033102-1302310322310221-0213303022202231"></a>

Type: `"object"`. single nested block, Optional.

Default Cache Behaviour. This defines a Default Cache Action.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("cache_disabled",
    "cache_ttl_default"),
  validators.ConflictingObjectAttributes("cache_disabled",
    "cache_ttl_override"),
  validators.ConflictingObjectAttributes("cache_ttl_default",
    "cache_ttl_override")}
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
  "x-ves-oneof-field-cache_actions": "[\"cache_disabled\",\"cache_ttl_default\",\"cache_ttl_override\"]"
}
```

Terraform syntax:

```terraform
default_cache_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123222302110103-1112333021101000-3020321013120002-2103320322230022-3301013202220211-3130020131022323-2202212023310101-3200231212322130"></a>

### Direct properties for `caching_policy.default_cache_action`

- [cache_disabled](resources--http_loadbalancer--reference--group-013.md#canonical-3210112322012122-2202223201233230-0231212203120000-1333302200310022-3030002313231033-1230310201012301-2221012202100321-3012003200320101): complete subsection reference.

<a id="canonical-2203031003231100-3022333131332021-2221023111303010-1111210311031131-2002200031103213-2100310230301010-0031011003133032-1020222022003100"></a>

<a id="canonical-1201012323331313-0212220020030130-0032031323023103-2303202230232130-2322123102232320-0210203202331122-2210210111002130-3112033303312301"></a>

#### `caching_policy.default_cache_action.cache_ttl_default` property

Type: `"string"`. Optional.

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-1103331303022230-2030311020010200-2312003102022223-3331033001011201-1132002101322200-1331310323132103-2000001013111112-0330012000223031"></a>

<a id="canonical-1213133333201102-0010020300101200-0121100222013200-1331202320221031-1213311332033021-3311332003032322-3021001230111233-2222201320112333"></a>

#### `caching_policy.default_cache_action.cache_ttl_override` property

Type: `"string"`. Optional.

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-3210112322012122-2202223201233230-0231212203120000-1333302200310022-3030002313231033-1230310201012301-2221012202100321-3012003200320101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `caching_policy.default_cache_action.cache_disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [caching_policy](resources--http_loadbalancer--reference--group-013.md#canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331)
- [caching_policy.default_cache_action](resources--http_loadbalancer--reference--group-013.md#canonical-2102303123010113-2230132023113110-0301132233211000-0123111331333002-0133122303200233-0021301123001210-2212032330201311-1313032131321121)
- caching_policy.default_cache_action.cache_disabled

<a id="canonical-2321332220331300-0021013010322012-2122023211002010-0313210333111321-1232132323022201-0002130102301131-1212223230022133-3320332113331132"></a>

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
cache_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331001122010103-1330222312203012-2102022330310300-3121202020223232-3202122212132233-0220222003113103-2301213310311323-1232121323232302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `captcha_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- captcha_challenge

<a id="canonical-1232032332023102-1010001000122202-0112003231022230-3233031333012222-3331333220233000-3300331300233031-1213033101103311-0120000120323113"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: captcha\_challenge, enable\_challenge, js\_challenge, no\_challenge,
policy\_based\_challenge; Default: no\_challenge\] Enables loadbalancer to perform captcha challenge
Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that
pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is
configured to do Captcha Challenge, it will redirect..

Additional upstream details:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha. When loadbalancer is configured to do Captcha
Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page
will have captcha challenge embedded in it. Client will be allowed to make the request only if the
captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid
Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only
human users can pass through. Generally, computers or bots are not capable of solving a captcha. You
can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [captcha_challenge](resources--http_loadbalancer--reference--group-013.md#canonical-1232032332023102-1010001000122202-0112003231022230-3233031333012222-3331333220233000-3300331300233031-1213033101103311-0120000120323113)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-0331332220112013-2112002310330311-0000103101120120-1022133131132112-2221133100010123-1020333222003102-1303303322000302-2321231220030231)
- [js_challenge](resources--http_loadbalancer--reference--group-019.md#canonical-1011111331300221-3213330020020313-1033022332302021-1111331213030331-0032003020022230-3000023321331133-2002032233231133-0233011032133100)
- [no_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-3132021103233010-0102003303001330-3110103203321112-0211303310001001-0000320020121320-3033100021101033-2220023223112313-3310221213303023)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-0321223103102130-0121132313031223-1220210002011320-2133231330312121-2131132110220002-1313302332130023-1110033030130013-0300230120302232)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
captcha_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133303122032321-1102031001103232-0130323202211110-0131132123321032-2100221023312112-0033303023201301-0131020332231231-1300311201002102"></a>

### Direct properties for `captcha_challenge`

<a id="canonical-0310033331102021-0301312112130232-0113033213310010-1330111223022001-3300201112031022-2101233333130013-3212321221102111-1232102220300020"></a>

#### `captcha_challenge.cookie_expiry` property

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3223102002230332-3000031121132103-2320100232201302-0121230003321322-0322212333031120-1013133331211202-0030000220132103-3023210211231202"></a>

<a id="canonical-1311102022130030-0332013011230103-1010201013020033-3313023121011013-2320202323123200-2312300211301121-1201210202222000-1102302033132122"></a>

#### `captcha_challenge.custom_page` property

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- client_side_defense

<a id="canonical-0033332020201100-0231102120012232-1311202303221021-2103323133031333-2103310211223013-1103120330033200-3332001302000030-0233201201222132"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: client\_side\_defense, disable\_client\_side\_defense; Default:
disable\_client\_side\_defense\] Defines various configuration OPTIONS for Client-Side Defense
Policy.

Additional upstream details:

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

- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0033332020201100-0231102120012232-1311202303221021-2103323133031333-2103310211223013-1103120330033200-3332001302000030-0233201201222132)
- [disable_client_side_defense](resources--http_loadbalancer--reference--group-017.md#canonical-1313000133203032-2310331201223313-3311201202200022-2303113031201211-1010211132031221-2313112333213303-2001211021211300-1211030023323302)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
client_side_defense {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032032301023113-3100231103301321-1230030131223003-3203230231233011-2010013330021023-1103010031300331-0113202302310320-3211231113232121"></a>

### Direct properties for `client_side_defense`

- [policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310): complete subsection reference.

<a id="canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- client_side_defense.policy

<a id="canonical-1213221003211122-1132011111321033-2013113123130032-3111322002231211-2103303031223230-0201123200123032-1231100122322021-1031031101000230"></a>

Type: `"object"`. single nested block, Optional.

This defines various configuration OPTIONS for Client-Side Defense policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3202032233130320-0223132112300033-1300101303310010-2032110011101312-2320332112120003-2031203333331230-3113131323032323-3023123121131203"></a>

### Direct properties for `client_side_defense.policy`

- [disable_js_insert](resources--http_loadbalancer--reference--group-013.md#canonical-0113131101113001-3321120322331133-2303222210330300-0300223210012311-3322133012200132-1110000111313120-0313300310231031-0121202103222232): complete subsection reference.

- [js_insert_all_pages](resources--http_loadbalancer--reference--group-013.md#canonical-3321121013323313-0201322201321222-2200102111023133-2312132222201301-2223320202220022-2332133133322110-3112313330120313-2011323100003200): complete subsection reference.

- [js_insert_all_pages_except](resources--http_loadbalancer--reference--group-013.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231): complete subsection reference.

- [js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023): complete subsection reference.

<a id="canonical-0113131101113001-3321120322331133-2303222210330300-0300223210012311-3322133012200132-1110000111313120-0313300310231031-0121202103222232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.disable_js_insert` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- client_side_defense.policy.disable_js_insert

<a id="canonical-1031311331321213-0203202131310323-0022230000002002-2110313003310132-1123031023301020-3202320021223300-1200010203322230-3133223321133001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321121013323313-0201322201321222-2200102111023133-2312132222201301-2223320202220022-2332133133322110-3112313330120313-2011323100003200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- client_side_defense.policy.js_insert_all_pages

<a id="canonical-2032311313000003-2300322321221032-1121303333310212-1022020201223230-1330022331302200-3131001310202301-3301210131010012-3021022311111020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for js insert all pages.

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
js_insert_all_pages = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- client_side_defense.policy.js_insert_all_pages_except

<a id="canonical-3133230212120011-3032310011031001-3310330000123011-0132000110133330-0231122131212303-1011210322202321-0303132232100232-3202233212302013"></a>

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

<a id="canonical-1121222102322232-0131210203230223-2001213203333322-3010312302311321-1113302121200201-3231323223230231-0213212022101131-2323321002331122"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except`

- [exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313): complete subsection reference.

<a id="canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-013.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-3131000322120213-2323121100223113-3311003130303011-0321021323110200-3102323022110122-3311132002033203-2210013203313331-0232300310121003"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2131100021112022-1201131320231132-3332312203031331-1132312032002331-3331331103311203-2131301221212213-3310333301103100-0031121220113011"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except.exclude_list`

- [any_domain](resources--http_loadbalancer--reference--group-013.md#canonical-3201331123201221-0313230333212000-0132203201121120-2032221301322010-0010132300133000-3301113012010321-3133133012001123-3101022113302312): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-013.md#canonical-2002333202310111-1022200100012001-1223122121131230-2223310031202011-2320100231113213-0302310310331300-3311113032031312-0230213130322133): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-013.md#canonical-2213132032212100-1030112202020323-1032113021332311-3203323121302113-2000021221311113-1231013120322233-0023111100301111-0302322221113211): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-013.md#canonical-0313113200033031-0213201113202133-0212000000101012-1220101330223002-0233201331122321-2012200102121121-0232212210302121-2023222022330020): complete subsection reference.

<a id="canonical-3201331123201221-0313230333212000-0132203201121120-2032221301322010-0010132300133000-3301113012010321-3133133012001123-3101022113302312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-013.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-0121132311002212-3331021223313100-1211230321231113-1122123021020211-2133213130203232-1113132200032031-0101310111131001-1010323311230130"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002333202310111-1022200100012001-1223122121131230-2223310031202011-2320100231113213-0302310310331300-3311113032031312-0230213130322133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-013.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-3133220201201202-2020302121212302-0320112130302213-1132132332021112-1031213003110003-0312120111323001-0023031003103323-2223130113110000"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1003310102001121-1231210222031000-3313112100010122-1322003300020023-0033133320231031-3133220312020321-3320010301123320-3133023120113113"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain`

<a id="canonical-3331123031010323-3201030310221110-2012230332013133-0100013231131123-2330200122113223-1231002220011131-2112133332200201-3222020321302120"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1120123030013013-2103001112303113-1221220320231112-2113111111032021-2210301322032202-2010232231011300-2131010333111220-3312111222303320"></a>

<a id="canonical-3201332333203113-2123021103300030-1310033121010303-3331012231013221-0223310300002132-3200202002321121-1211220100210011-2331232302110321"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3100203031221001-3112130201333213-1021213233002302-3211011131203000-2032231022233231-1101032103321031-1202311211101302-1302111010000122"></a>

<a id="canonical-3221013111332320-0000232222332301-0220002032303101-1202013212322203-0310021330030030-3302110131300310-1000003213130312-3030130100202123"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2213132032212100-1030112202020323-1032113021332311-3203323121302113-2000021221311113-1231013120322233-0023111100301111-0302322221113211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-013.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-2021001212312321-1010322200001202-0121320212022300-0002321010133133-3301130130120200-0301023031113213-2201101002121233-1321003313310122"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2013202322001220-3001222231112132-3211031121303200-1300033303322320-1213200011012021-3112101320231121-0302232233013111-0323210333002112"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata`

<a id="canonical-1112013112202133-1020230302001013-3113111303130133-2301101030010031-3311123001103320-3200121303322103-1231202302202232-0102333010021220"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1111032031200311-1322302222301132-0310203000121310-0302313032311031-2011030320112110-2302311110023022-2001100323303232-1202132002110133"></a>

<a id="canonical-3202323330000322-1033101010210302-3210323011122131-1212133001021001-2332112012233312-3212021112013121-1132322010100211-2020220310311121"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-0313113200033031-0213201113202133-0212000000101012-1220101330223002-0233201331122321-2012200102121121-0232212210302121-2023222022330020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-013.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-2221000100332222-2303223321131020-0111013311213213-3310323233333123-3000031113100001-2233302121020221-0021123012001131-3212012021021311"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2013013022300321-1010002231031230-3332111021202231-3320210302101230-1330130130132300-3303201122002001-0311301130213223-3303331101011113"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path`

<a id="canonical-1111212032200320-3011301031023323-3221330302210210-0000311030221221-3101013120030130-0131130212220210-2003302020012202-0011111131132110"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0333303301330320-0312103102302023-0331033022033101-0333203001001100-3131333013111302-2203201101121311-1221102022233120-3101023111201021"></a>

<a id="canonical-0221310312133230-2203133213302202-0020223031132123-0231133100132111-1201120300302223-2202230012333232-1101130013022312-2111221223021233"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-2100031031030130-2012301112332302-0223131222120022-0302012310202000-2333023001003312-2300202331010122-3103322232202213-0221210311001322"></a>

<a id="canonical-3021031230112032-2013110101300330-1331023221232211-2002131021310023-1320011032311203-1313111213012323-2013232103211230-3032001113321031"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- client_side_defense.policy.js_insertion_rules

<a id="canonical-2223330120010033-3330220112301033-0301003123001311-3320220321313233-1132332333110202-3003123131132332-1130231203233122-0213223013101130"></a>

Type: `"object"`. single nested block, Optional.

This defines custom JavaScript insertion rules for Client-Side Defense Policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2213031210023023-3000011212202323-2211002013132210-3303212311102221-2211101103303023-0100112202221020-3300301203030310-1232122001033000"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules`

- [exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222): complete subsection reference.

- [rules](resources--http_loadbalancer--reference--group-014.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321): complete subsection reference.

<a id="canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- client_side_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-0302200211013330-3120012120312110-1023212310321322-1101203233113201-1201102222121132-1220100321322210-3122330033303312-2232322300130320"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1203022302023012-2230320022001313-1023130211300120-1133002013231222-3222103000111031-0013311233103333-1221030113213231-0113332321000311"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list`

- [any_domain](resources--http_loadbalancer--reference--group-013.md#canonical-1211032113213031-1312033220303030-2000131032203030-0210101023300331-1302000301020223-3021320231203002-3032322113012022-3023223132032202): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-013.md#canonical-1302133302131301-0302203010332010-3032202302312331-1322312023023320-0121210331002113-3331322031031103-3212132120101113-3210001000000231): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-013.md#canonical-0323321301322131-3031120031101112-2203233330200212-1222233131100003-0310002222021030-1211312130102001-2320312203103213-1012210013102312): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-013.md#canonical-3211010123012220-1110123231023113-1023333033223033-3122203312012010-3321230002200121-0033210201023331-2132301132120200-3021313130322321): complete subsection reference.

<a id="canonical-1211032113213031-1312033220303030-2000131032203030-0210101023300331-1302000301020223-3021320231203002-3032322113012022-3023223132032202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- client_side_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-2303001012113301-0023200320010131-3220013113232303-1133220220310320-1100231331323113-2322033121231203-0333110011112320-1003012332200223"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302133302131301-0302203010332010-3032202302312331-1322312023023320-0121210331002113-3331322031031103-3212132120101113-3210001000000231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- client_side_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-1221121033332311-2012000301122123-2113031000012302-3012301133313221-1222101310211300-1002231102321103-0332023012300313-1300132120312113"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2331210213101210-0233313220303113-1312100230102302-1333331111112322-1323331011232333-1222230220312330-1101113220213220-2210233131333303"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.domain`

<a id="canonical-1122100220003232-1022313203011322-1120031301310230-0200223222310232-0313110232203130-3330022323330233-0022210020231223-3220001021202211"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0031021330022122-3320020301030213-1210022001202101-1201123202331132-2300123000232203-2310232003013201-0313101321213111-1022210011321013"></a>

<a id="canonical-1221013130132331-3210012323033230-0130021000201002-0300310311300230-2330321230023201-3013322211111202-2213310002113103-0103301032223331"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3212121011203331-1120101130313031-3333310303102103-1212023331231311-0330021011331013-2123220030212130-3332100020230211-2020120312022212"></a>

<a id="canonical-2311131310311013-2122030102002211-1012313100101003-3312122013103102-2201322132131320-0022130200231100-1112112332322302-3221032100201013"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0323321301322131-3031120031101112-2203233330200212-1222233131100003-0310002222021030-1211312130102001-2320312203103213-1012210013102312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- client_side_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-0231121201131001-1310103223323133-2210210103100121-1030332233103330-2000032120122311-0100322132313332-0010323223301031-2200323033223010"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0031013303003322-0230323310131112-2001112220310212-2320033213003102-0301232103132312-0301121033231230-3232030010200023-3233312112222321"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.metadata`

<a id="canonical-1331203223102003-2212021133333302-0230023021300001-1002231133223310-1133101023113320-3121010031120313-1303202302033313-0100200130002003"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3123123003033212-0003203322121032-0201303233321321-0312022301223301-1113011003002333-1003302330123101-1313333201231231-3311000312231121"></a>

<a id="canonical-0313233020301113-2001110200222000-2101113031131232-2322002321210321-0010010220311130-1330302220211332-2113110311123303-3303100210112323"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-3211010123012220-1110123231023113-1023333033223033-3122203312012010-3321230002200121-0033210201023331-2132301132120200-3021313130322321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-013.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- client_side_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-0213001220012101-0132033101322101-3331002303112111-2033122100033133-1122023211110301-0211031010131320-0213233130013023-0321231231020222"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0321032121132301-0233333031223330-0021233020133222-0030233330201011-3133121213011033-2311211013031233-3231330101202033-2021022022332300"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.path`

<a id="canonical-3001213122213233-0302112032230113-0010120333220232-0320330321112320-3112023321132120-1112031030202213-2333300332023000-0013331211122322"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0220323202003033-3100302023111320-1003020021320121-0033121013322003-0001222002002010-1120300023210333-1011321033310100-0111300031320233"></a>

<a id="canonical-1201203013131010-0133233210121020-0011010032312300-3230201331230000-0330332203200121-2103020231132210-1222010212211122-0201231130223022"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-0211332330211200-2300310012010133-3311012200330220-1213210131320213-1331001012213033-1030113001111020-2331010103230113-2320330323223130"></a>

<a id="canonical-2332213022301011-2101022322121012-1222200200113332-1121333003221301-1030210101313110-3212230310101002-0133030222331210-3032112202123212"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
