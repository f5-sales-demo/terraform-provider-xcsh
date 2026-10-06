---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1103203303113331-2302123302001100-1121020032121231-2222122210210133-2002002122010003-2232003301121213-3201221312313321-0202203023112121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.bot_defense_javascript_injection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.bot_defense_javascript_injection

<a id="canonical-3112222331022130-1332123020020302-2300331200133303-1200121303103320-3010212320333220-1301323113301132-2333032212133323-1212203233121122"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense JavaScript Injection Configuration for inline bot defense deployments.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("javascript_tags")}
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
bot_defense_javascript_injection {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011121133213003-2003233313002233-1320110321300012-2313213330110300-3130030002120232-0221212333210120-2221200011310302-3221100031313131"></a>

### Direct properties for `routes.simple_route.advanced_options.bot_defense_javascript_injection`

<a id="canonical-2010121222013120-0130321033322012-1003222230200300-3131022000222030-3202012002031333-1103223332030021-1031311110031301-0031221300113021"></a>

#### `routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_location` property

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

- [javascript_tags](resources--http_loadbalancer--reference--group-026.md#canonical-0230323012012233-1331313230201303-2013102231310211-2300022200010131-0013100032031013-1323313021311010-2301021330031011-3003103033102121): complete subsection reference.

<a id="canonical-0230323012012233-1331313230201303-2013102231310211-2300022200010131-0013100032031013-1323313021311010-2301021330031011-3003103033102121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-026.md#canonical-1103203303113331-2302123302001100-1121020032121231-2222122210210133-2002002122010003-2232003301121213-3201221312313321-0202203023112121)
- routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags

<a id="canonical-3211033200233212-2313002301301112-2303030310230002-1210212112123000-2303001121000111-1213321320221211-2030132022310021-2302001311333213"></a>

Type: `"object"`. list nested block, Optional.

Select Add item to configure your JavaScript tag. If adding both Bot Adv and Fraud, the Bot
JavaScript should be added first.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("javascript_url")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
javascript_tags {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222233101211003-1121221212013303-2021021231310322-2030221330320022-2130012111123300-3323202023201300-2333220033200131-0011201112001032"></a>

### Direct properties for `routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags`

<a id="canonical-1303311102311322-2132132101111330-1010012220010001-2011233130120010-0300310013022233-3230022020011021-3132203320111113-3112123311103300"></a>

#### `routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.javascript_url` property

Type: `"string"`. Optional.

Please enter the full URL (include domain and path), or relative path.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 2048,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [tag_attributes](resources--http_loadbalancer--reference--group-026.md#canonical-2113101232122210-0112022221131010-2302331003003113-1033122321301121-2100233130331013-1003333011321300-2220033023013303-1121111302131110): complete subsection reference.

<a id="canonical-2113101232122210-0112022221131010-2302331003003113-1033122321301121-2100233130331013-1003333011321300-2220033023013303-1121111302131110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-026.md#canonical-1103203303113331-2302123302001100-1121020032121231-2222122210210133-2002002122010003-2232003301121213-3201221312313321-0202203023112121)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags](resources--http_loadbalancer--reference--group-026.md#canonical-0230323012012233-1331313230201303-2013102231310211-2300022200010131-0013100032031013-1323313021311010-2301021330031011-3003103033102121)
- routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes

<a id="canonical-3121300113022303-1010023332312023-2330202003302102-2103331113311202-2102131223301021-0100232102101131-0122221010002021-3202020102013203"></a>

Type: `"object"`. list nested block, Optional.

Add the tag attributes you want to include in your JavaScript tag.

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

Terraform syntax:

```terraform
tag_attributes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022220000230303-2320222123002122-0303112220202301-1203113230332223-0323010011023312-3013313203312231-0013100132210222-0230222332331121"></a>

### Direct properties for `routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes`

<a id="canonical-0232100331022001-2332333203130221-0033203231203123-2213323100011233-0321101233331320-2133213032123012-3121221123210203-3010013030330332"></a>

#### `routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag` property

Type: `"string"`. Optional.

\[Enum:
JS\_ATTR\_ID|JS\_ATTR\_CID|JS\_ATTR\_CN|JS\_ATTR\_API\_DOMAIN|JS\_ATTR\_API\_URL|JS\_ATTR\_API\_PATH|JS\_ATTR\_ASYNC|JS\_ATTR\_DEFER\]
Select from one of the predefined tag attributes. Possible values are \`JS\_ATTR\_ID\`,
\`JS\_ATTR\_CID\`, \`JS\_ATTR\_CN\`, \`JS\_ATTR\_API\_DOMAIN\`, \`JS\_ATTR\_API\_URL\`,
\`JS\_ATTR\_API\_PATH\`, \`JS\_ATTR\_ASYNC\`, \`JS\_ATTR\_DEFER\`. Defaults to \`JS\_ATTR\_ID\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["JS_ATTR_API_DOMAIN","JS_ATTR_API_PATH","JS_ATTR_API_URL","JS_ATTR_ASYNC","JS_ATTR_CID","JS_ATTR_CN","JS_ATTR_DEFER","JS_ATTR_ID"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("JS_ATTR_ID",
    "JS_ATTR_CID",
    "JS_ATTR_CN",
    "JS_ATTR_API_DOMAIN",
    "JS_ATTR_API_URL",
    "JS_ATTR_API_PATH",
    "JS_ATTR_ASYNC",
    "JS_ATTR_DEFER"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "JS_ATTR_ID",
  "enum": [
    "JS_ATTR_ID",
    "JS_ATTR_CID",
    "JS_ATTR_CN",
    "JS_ATTR_API_DOMAIN",
    "JS_ATTR_API_URL",
    "JS_ATTR_API_PATH",
    "JS_ATTR_ASYNC",
    "JS_ATTR_DEFER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3101202021202120-1022113031203323-3022020220313102-3030210203020111-1121331131331202-3222230210133333-3100111310120011-0022310133012310"></a>

<a id="canonical-3103332200311222-0010133221101320-1333313031011211-3012223301031302-0331120030031210-1211003103213200-3220100312312013-1011333333013032"></a>

#### `routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value` property

Type: `"string"`. Optional.

Value. Add the tag attribute value.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-1211001130101123-0202021333221102-1013312002223231-0033202013223031-2322012302201213-0221112223001103-3230112102023130-3102232202213203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.buffer_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.buffer_policy

<a id="canonical-0100003330320311-1130013313132330-1332211133101210-0112202221322202-2112023010123231-3132013111213022-1220001332310301-0223131121330331"></a>

Type: `"object"`. single nested block, Optional.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer
configuration takes precedence.

Receipt-pinned upstream constraints:

```json
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
buffer_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133013003112221-2112101211021333-3332332112021231-1231330011030333-3310210023103300-0201133303223320-2302302103211133-0221221120002303"></a>

### Direct properties for `routes.simple_route.advanced_options.buffer_policy`

<a id="canonical-1111222330201102-0010202011110122-0232230023100020-3203133201002332-0303223103122202-0032231023123021-2221231130023213-2232120232022021"></a>

#### `routes.simple_route.advanced_options.buffer_policy.disabled` property

Type: `"bool"`. Optional.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3210223022200001-0032230303021123-3201122300320012-1110002103233101-1300323231130023-3102100102033300-0331232303103001-0132232103233300"></a>

<a id="canonical-1010012232101312-0212203313301302-0003020030111100-2302330201232231-0112330032030010-0310303120312232-2010233302310332-0112302212123332"></a>

#### `routes.simple_route.advanced_options.buffer_policy.max_request_bytes` property

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(10485760),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

<a id="canonical-3020133322310122-3120022303032302-1333013033231033-0203223021012000-1333112131131133-1101022331131132-0202323320200021-2220033102132032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.common_buffering` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.common_buffering

<a id="canonical-3201331033013213-1132133200002311-2213131321030011-2213222003212323-0000231132330102-0313011232100201-2323110323302011-1012311002020020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for common buffering.

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
common_buffering = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220202221213121-2003330033111010-3112031032321133-0331322330011032-2100203210203032-1312312203321022-3331311132203101-1323310311023321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.common_hash_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.common_hash_policy

<a id="canonical-0023222113102030-2133033113012302-0333322313321103-3221333321200100-1232100113213201-1011310221311332-1330002200013002-3032002111312231"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
common_hash_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222022011013211-0231322200331030-3003132300320323-0223133302201212-3011220203110302-1311320103313321-2022323330221202-0131123122131320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.cors_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.cors_policy

<a id="canonical-2023331023001023-2020203132023101-2331310220020032-0123020311022201-2110332123100303-3223032113133301-3030112112230200-3020302212332000"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2103133001033303-2311032132311131-2200232202330122-3131300230233322-0133233023020232-0021221101101111-1021103210222322-2121100212222323"></a>

### Direct properties for `routes.simple_route.advanced_options.cors_policy`

<a id="canonical-0222001133330130-3202233330130000-3303120213020120-1201312231203313-3322233010300130-0112031233222202-2110311230110102-1130333311230223"></a>

#### `routes.simple_route.advanced_options.cors_policy.allow_credentials` property

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

<a id="canonical-3330102120203133-0130302222230202-2210300303011101-3130210131211322-1233011230103003-1212210133212021-3301303002131032-1020310002013111"></a>

<a id="canonical-3003310122030303-1011311030032320-0032222213302012-2112233100311332-2330213133303012-3132200322100313-1122221322112200-3311012113100200"></a>

#### `routes.simple_route.advanced_options.cors_policy.allow_headers` property

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

<a id="canonical-3121313321023303-0203101121221200-3232021031311223-1231202313101120-0200310023233100-1321200121210031-2110031032101102-0100233201122213"></a>

<a id="canonical-0230011133100001-1210011102221133-0133020200001203-3031212222330301-0312102301330012-0303221023303102-1023232101302021-3201011220003100"></a>

#### `routes.simple_route.advanced_options.cors_policy.allow_methods` property

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
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-0200322113021211-2201113221130210-1102113331011023-3021013102111312-0120213023333232-0022011101200121-3133102110303010-3002323131233233"></a>

<a id="canonical-0203100221310222-0231223021021223-3130012122223110-2323103311103232-3030131012110302-1013231323320000-0312223021203131-3122020012212132"></a>

#### `routes.simple_route.advanced_options.cors_policy.allow_origin` property

Type: `["list", "string"]`. Optional.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

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

<a id="canonical-0232110312022111-1331021301213022-0331131320102020-0011213311222311-1200120201230233-3220200201213200-1111312202000120-3300120030101313"></a>

<a id="canonical-3112322323332120-3201303113330320-0300330102113333-2031010230232212-1031102210022202-1122023213100320-3312000300300330-0321331133112003"></a>

#### `routes.simple_route.advanced_options.cors_policy.allow_origin_regex` property

Type: `["list", "string"]`. Optional.

Specifies regular expression patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

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

<a id="canonical-2321130231121322-0123311030333212-0210001130213311-0210110222102012-1111222002110302-2012203220332003-1120132301003310-1310021000002000"></a>

<a id="canonical-3220100113130021-1203001020003121-3120002131223030-1020223111211001-0231323230122013-3302302300203311-0030012233120131-2222213200022000"></a>

#### `routes.simple_route.advanced_options.cors_policy.disabled` property

Type: `"bool"`. Optional.

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

<a id="canonical-1231130321323201-0022232333300330-0300000312230221-0230322023133332-1100203332013130-1110031230231303-2202313012133100-3233210002300101"></a>

<a id="canonical-0231323002010212-0303022223211110-2313311100311312-0032213331322032-1220132202310022-1233201013011202-1202102100201033-2111332000213233"></a>

#### `routes.simple_route.advanced_options.cors_policy.expose_headers` property

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

<a id="canonical-1311210313132213-3210132133203231-1032111210313312-3102310001303111-3133121130031012-1213201310133232-0103023321132210-3010313111030113"></a>

<a id="canonical-1011301312020212-1203111121310333-3331033003222231-3001223333320023-0311120231133200-0011032332100103-3103121311301313-0111211000311302"></a>

#### `routes.simple_route.advanced_options.cors_policy.maximum_age` property

Type: `"number"`. Optional.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.csrf_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.csrf_policy

<a id="canonical-3210133010213111-1312130212033003-3022113113101331-1010102020220033-1010220233110232-3223132020131032-3310132330002033-2113303103211212"></a>

Type: `"object"`. single nested block, Optional.

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
EnumExtractionComplete: false
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

<a id="canonical-3302332133031330-3310210020103232-2203033023231331-3331100010012101-3213023113221102-0210332210200212-3120322231121132-3001200011110231"></a>

### Direct properties for `routes.simple_route.advanced_options.csrf_policy`

- [all_load_balancer_domains](resources--http_loadbalancer--reference--group-026.md#canonical-3102211303130002-0311302231123112-3021133021121320-2300112000101322-0022213122203220-3321203021313221-1030122323103302-1211130133103221): complete subsection reference.

- [custom_domain_list](resources--http_loadbalancer--reference--group-026.md#canonical-2232023010120200-1330112210123120-2122301123331132-3212302133120122-0002301100303000-2120003300300223-2211300033013032-0310121310232122): complete subsection reference.

- [disabled](resources--http_loadbalancer--reference--group-026.md#canonical-2213021333110233-0321023022220030-3221132311032333-0103000100120312-3322200031221311-2301022200220203-1302310322301221-1201100321332030): complete subsection reference.

<a id="canonical-3102211303130002-0311302231123112-3021133021121320-2300112000101322-0022213122203220-3321203021313221-1030122323103302-1211130133103221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311)
- routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains

<a id="canonical-3330010002123331-0022223323002002-3212032222132113-3133203012310321-3231113310301302-2003312122222010-0122012202020021-0030113232111333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all load balancer domains.

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
all_load_balancer_domains = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232023010120200-1330112210123120-2122301123331132-3212302133120122-0002301100303000-2120003300300223-2211300033013032-0310121310232122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.csrf_policy.custom_domain_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311)
- routes.simple_route.advanced_options.csrf_policy.custom_domain_list

<a id="canonical-2111121030123202-0103312021020101-3012203203132300-0111131221313003-1323333033023330-0212311110023113-2011120333300333-3001022323033110"></a>

Type: `"object"`. single nested block, Optional.

List of domain names used for Host header matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3023230132321033-0101132302132103-0132322300300001-1302011001111303-1023311003231001-2123130210122232-3112110031123300-1223321232322103"></a>

### Direct properties for `routes.simple_route.advanced_options.csrf_policy.custom_domain_list`

<a id="canonical-2302133323021231-1023123012230012-3220033331021322-2033012111133121-3110212122013320-1320011012302003-3303101220133023-3201221033033110"></a>

#### `routes.simple_route.advanced_options.csrf_policy.custom_domain_list.domains` property

Type: `["list", "string"]`. Optional.

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2213021333110233-0321023022220030-3221132311032333-0103000100120312-3322200031221311-2301022200220203-1302310322301221-1201100321332030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.csrf_policy.disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311)
- routes.simple_route.advanced_options.csrf_policy.disabled

<a id="canonical-2200132233200213-3323322130131233-2112201110103322-0331210302310203-0321233032120302-3311231133210010-0111212232201221-1300003232330211"></a>

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
disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013002323211230-2000321010322031-0202220333302003-3320223120203023-3311023301310130-1123220212320222-2303233031330210-0222030031022311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.default_retry_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.default_retry_policy

<a id="canonical-3103011220122302-3321311323030103-1032101303211301-3103011110101031-0300022110300232-0103123021013002-3201132121210023-2302110330111320"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
default_retry_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223210220132003-1032031311131022-0221333010313232-3300232111321333-3020023100102303-3000311333203311-2323022130322120-1301313300201232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.disable_mirroring` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_mirroring

<a id="canonical-1021031031132110-2033021030110230-0231111133321222-3131103331311021-3332130220020323-0222201323200231-0312010103333220-0200212111112210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable mirroring.

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
disable_mirroring = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101222121331231-0002101210102120-1000231102102220-2110231230123212-2111001231312132-2121211202233201-1300220103223003-2023201222120123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.disable_prefix_rewrite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_prefix_rewrite

<a id="canonical-1222222120001310-3020213133313220-0021000331131010-3113133202110330-1221201321122132-2331133300020222-2220221223203102-2221233223111201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable prefix rewrite.

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
disable_prefix_rewrite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333112331031332-3110230102112221-1123102212221222-1320333331030110-2102300032332222-3010102120200200-0200212202113120-3221312122220012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.disable_spdy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_spdy

<a id="canonical-0301033111111300-1200133121222020-0233012113003210-0010312030213210-0330310232113003-0103132321001201-1102321013201122-1022131111003311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable spdy.

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
disable_spdy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012322210102001-1000202101032101-0301012012020133-2222012131120302-2033130323220023-0131111302321232-1103220102022322-0130330011322111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.disable_waf` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_waf

<a id="canonical-3303013011100012-0200023313012302-0002023003301032-3111022310002230-0112220121131100-1213103021322322-2210322122321100-3322103003022132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable waf.

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
disable_waf = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123013100222203-3012313301032221-2333300120200121-2031023303323312-2330102330202121-3313212031211202-1300121232122232-3230133031002030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.disable_web_socket_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_web_socket_config

<a id="canonical-0013020202203110-3312230002300003-2330231002010231-1111020021200000-1032312100120321-0201310310023232-3223000123310102-0213233200030033"></a>

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
disable_web_socket_config = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013100112201010-3320010213003210-2110103020232131-1233302023023333-1320111303223010-1220201100133030-2322330330102010-1111303301333333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.do_not_retract_cluster` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.do_not_retract_cluster

<a id="canonical-3032010210001000-0132301310112113-3231312222323112-0320100231221110-1313022301112231-3320211323123121-3202333203021331-1011211031002103"></a>

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
do_not_retract_cluster = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220221103013030-1033023211020212-2313000220302030-0112212123203222-2223110003210120-1330110301122333-0120013212320200-2011130012212031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.enable_spdy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.enable_spdy

<a id="canonical-1012101022332202-2310312301302112-1301110311103223-0312211212223123-0320331111303232-3230130213121033-0020320312112021-2022113200301210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable spdy.

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
enable_spdy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123231212310000-0130220212030233-2332222131111223-3330003103011023-0211122003332012-0200233023011101-3232021110002213-0302122311123131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.endpoint_subsets

<a id="canonical-1221121322322213-0221232101101003-3131212031322320-0000323232312133-3201011032032323-3200100001302133-1113231122011333-1332133101210333"></a>

Type: `"object"`. single nested block, Optional.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101301002211102-3101130203012322-2333001122113100-1010021103332111-0203220210230210-1230033102313122-1102112111223121-2210120101222000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection

<a id="canonical-1122201311111012-3332331300213133-0331001223212103-2210100300201122-0021211020120201-0011310112302132-0031001122302210-3233322011310332"></a>

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
inherited_bot_defense_javascript_injection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213201113012322-1102122130221322-3332120321331122-3312320023133031-1013323333130130-0033111303102122-3133232312300201-0023120200310323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.inherited_waf` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.inherited_waf

<a id="canonical-1213202313123210-3132132322133331-1121203130032333-1001231202200300-2211032020131300-0311123000323133-0312302203222331-0301233211313322"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherited waf.

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
inherited_waf = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002012001013000-3310311312101011-3313033113321211-2131033121023113-1002001131311300-2001101231211002-3100031020332322-2132012212030231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.inherited_waf_exclusion` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.inherited_waf_exclusion

<a id="canonical-3223112102232310-3321103232021123-3331233121022220-0013021212303220-3230121322333000-2030221101321013-1031220232231203-0200203001323331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherited waf exclusion.

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
inherited_waf_exclusion = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013001221331220-1013310121201113-3023002030203320-3030001023331313-3101303133113133-2112000032110210-2021222012332022-3303330132233023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.mirror_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.mirror_policy

<a id="canonical-3203310310232132-1110311230130323-3100032331000310-1332003132203113-2013321032131232-1221311332102131-1210021022223002-1210211131322011"></a>

Type: `"object"`. single nested block, Optional.

MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is
'fire and forget', meaning it will not wait for the shadow origin pool to respond before returning
the response from the primary origin pool. All normal statistics are collected for the shadow
origin..

Additional upstream details:

The approach used is "fire and forget", meaning it will not wait for the shadow origin pool to
respond before returning the response from the primary origin pool. All normal statistics are
collected for the shadow origin pool making this feature useful for testing and troubleshooting.

Receipt-pinned upstream constraints:

```json
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
mirror_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100332221110030-3001201303011101-2231212231032300-0232132330301212-2001322222130131-3333201200212012-3131303302000201-1233321032303130"></a>

### Direct properties for `routes.simple_route.advanced_options.mirror_policy`

- [origin_pool](resources--http_loadbalancer--reference--group-026.md#canonical-1103223210103233-3323023333223103-3202033211110230-3123111220212303-1220320130301312-2303110030230323-3323221001100011-2221131001203201): complete subsection reference.

- [percent](resources--http_loadbalancer--reference--group-026.md#canonical-2313121201002301-1112103031031030-2213123011002321-0330220012031022-3302303212211102-3121203313102011-2121121230313002-3102200032303031): complete subsection reference.

<a id="canonical-1103223210103233-3323023333223103-3202033211110230-3123111220212303-1220320130301312-2303110030230323-3323221001100011-2221131001203201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.mirror_policy.origin_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.mirror_policy](resources--http_loadbalancer--reference--group-026.md#canonical-1013001221331220-1013310121201113-3023002030203320-3030001023331313-3101303133113133-2112000032110210-2021222012332022-3303330132233023)
- routes.simple_route.advanced_options.mirror_policy.origin_pool

<a id="canonical-0330312013233303-3302331312302211-3130203301310003-3323332322132021-3113122133301332-3100000312203131-1220023302033213-1312100321202331"></a>

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
origin_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310301012321302-3012000020211213-3001323230130101-0023110001010020-1020100213112230-0330310133132321-3110231003032231-0020233233010001"></a>

### Direct properties for `routes.simple_route.advanced_options.mirror_policy.origin_pool`

<a id="canonical-3211113123010223-1023002120123110-2303113110200102-2013031032121301-0323023200310200-3013001323231213-0011303123330111-1222312130011333"></a>

#### `routes.simple_route.advanced_options.mirror_policy.origin_pool.name` property

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

<a id="canonical-1101100332220010-1331313023021002-2302102000023023-1130311310012022-0033212210212011-0201013130300232-3233203310222213-3011030320301022"></a>

<a id="canonical-3233013230300010-0231122103112002-2032232212110212-2330322210032221-2203001021201130-0120111310001002-0301100221220311-0300003210020201"></a>

#### `routes.simple_route.advanced_options.mirror_policy.origin_pool.namespace` property

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

<a id="canonical-3100000232102232-1213203100330213-1030310333213110-2223112310100221-2332311033311321-2133220321111002-3133233001111313-0232003011123032"></a>

<a id="canonical-0302003331102210-1031302132212133-0210320322331310-1203223233300232-1130200132102133-0323020301210110-0333333123020333-1113013302132012"></a>

#### `routes.simple_route.advanced_options.mirror_policy.origin_pool.tenant` property

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

<a id="canonical-2313121201002301-1112103031031030-2213123011002321-0330220012031022-3302303212211102-3121203313102011-2121121230313002-3102200032303031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.mirror_policy.percent` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.mirror_policy](resources--http_loadbalancer--reference--group-026.md#canonical-1013001221331220-1013310121201113-3023002030203320-3030001023331313-3101303133113133-2112000032110210-2021222012332022-3303330132233023)
- routes.simple_route.advanced_options.mirror_policy.percent

<a id="canonical-3333100001321201-2123213121321233-0212200010330223-1300330210212023-2031031221302133-2331330023102203-0332000030211210-2000032223012111"></a>

Type: `"object"`. single nested block, Optional.

Fraction used where sampling percentages are needed. Example sampled requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("numerator")}
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
percent {
  # Configure direct properties listed below.
}
```

<a id="canonical-1211221302013022-0010303233000013-3213203023312110-2331011030110300-0122231003123101-0221311031203020-2213013002103310-0220300220313203"></a>

### Direct properties for `routes.simple_route.advanced_options.mirror_policy.percent`

<a id="canonical-1002101303031213-0323213000330212-3331000001300302-2212110102211110-0123023102130113-2021212232100200-3322212213202130-1020301213033210"></a>

#### `routes.simple_route.advanced_options.mirror_policy.percent.denominator` property

Type: `"string"`. Optional.

\[Enum: HUNDRED|TEN\_THOUSAND|MILLION\] Denominator used in fraction where sampling percentages are
needed. Example sampled requests Use hundred as denominator Use ten thousand as denominator Use
million as denominator. Possible values are \`HUNDRED\`, \`TEN\_THOUSAND\`, \`MILLION\`. Defaults to
\`HUNDRED\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["HUNDRED","MILLION","TEN_THOUSAND"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("HUNDRED",
    "TEN_THOUSAND",
    "MILLION"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "HUNDRED",
  "enum": [
    "HUNDRED",
    "TEN_THOUSAND",
    "MILLION"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3011323121203331-1131130302322021-3101022232210103-0111322310012121-1201100121201102-2031000312000101-0232010002110301-1102212200002221"></a>

<a id="canonical-3022203002213123-0303331021233320-2312120103223133-2132120323211000-0130211003032312-1330313002023310-0010323321110310-3230133212323313"></a>

#### `routes.simple_route.advanced_options.mirror_policy.percent.numerator` property

Type: `"number"`. Optional.

Sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000.

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

<a id="canonical-2220333121303102-3132022333332033-2111322310103130-0230222232313212-1102113200030202-3202032303023333-2120330013220013-0120312220300111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.no_retry_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.no_retry_policy

<a id="canonical-3223110100331210-0311003333323320-0032321010031313-2301320120121302-0333220301202313-2203311000223033-3333032123130023-3022213322223331"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_retry_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320010111202212-1023320033013103-0030311122023131-1323130311311223-3130123101023222-0210202202123330-0003303022231213-3303032312010332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.regex_rewrite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.regex_rewrite

<a id="canonical-1332213201131233-1311103113232313-0003110113230122-0032003101331030-1322220113200113-0122021121110230-1031201222131010-3132223232111312"></a>

Type: `"object"`. single nested block, Optional.

RegexMatchRewrite describes how to match a string and then produce a new string using a regular
expression and a substitution string.

Receipt-pinned upstream constraints:

```json
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
regex_rewrite {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310131301033321-3100012113110200-0203100220010203-0103101010121210-3300130331322132-1020021213013303-3123223111231110-2022303212131032"></a>

### Direct properties for `routes.simple_route.advanced_options.regex_rewrite`

<a id="canonical-2312210120331321-1032020300102000-3332010231222123-0100121301312303-3013033311120332-0102013213132310-0313331222302330-2023212230212100"></a>

#### `routes.simple_route.advanced_options.regex_rewrite.pattern` property

Type: `"string"`. Optional.

The regular expression used to find portions of a string that should be replaced.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3100331213020210-3122202001023210-0132103301222210-2000123200322211-0231021313202230-1131000122220200-1332122311032222-0303122322211300"></a>

<a id="canonical-1103023313300101-1133133011221330-3121121322111110-2100302312000330-3030232303331310-1103110231033312-3120213322120031-2203230220303133"></a>

#### `routes.simple_route.advanced_options.regex_rewrite.substitution` property

Type: `"string"`. Optional.

The string that should be substituted into matching portions of the subject string during a
substitution operation to produce a new string.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_cookies_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.request_cookies_to_add

<a id="canonical-2023123122330012-1011112221023022-2310012330221032-2230313222233031-0033302110203301-3331032233211112-2302113000113003-3330003113232023"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022200000101330-2030003120211121-0101301210330201-0233102013101111-3330323233323310-3313001111230230-3132222101302102-2202230331200110"></a>

### Direct properties for `routes.simple_route.advanced_options.request_cookies_to_add`

<a id="canonical-1131032303302203-1213301003023000-2111120020000022-0210222233211112-2310210113113112-0112122003003220-1032131033121202-1211101022032010"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

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

<a id="canonical-0012300203220231-3111002121230131-3311031230030213-2022110323132220-2110100103123333-1030332010102111-3101212200322201-1113032212003000"></a>

<a id="canonical-0221022211010232-0232210033000222-2112133100332110-2200312303020133-0100201223120102-3111111031022302-3210132121000110-3003000301020233"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.overwrite` property

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Additional upstream details:

If true, the value is overwritten to existing values. Default value is do not overwrite.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121): complete subsection reference.

<a id="canonical-2012233030000330-3220233123110232-2333031310332021-1133012323021311-1301222201131211-1301211302100110-3313123010113220-1103211303333312"></a>

<a id="canonical-1133100103113331-3301200103200002-2101033012101123-3002330033101020-0023313210212113-1010101300102212-1001333321021233-3122302312022213"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value

<a id="canonical-2221112133101200-2321000221313010-3233000123000132-3333220220332012-2220301120332123-0303030011102322-0221103012033122-3313033011201220"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120133230213313-0101030300023122-0121011332003121-3221213100211133-2322330001202303-1020313230020312-1322211310013310-3001031033220331"></a>

### Direct properties for `routes.simple_route.advanced_options.request_cookies_to_add.secret_value`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-0212231213230312-3102112110221211-1100022003031232-3320033210021133-3132120212202021-1320121122013221-2330103113101111-0102012112312132): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-2100100321020011-0120311233112012-3121323322113103-2310100301231321-3030210112013001-2301002311012321-0210103011301002-1211023132203232): complete subsection reference.

<a id="canonical-0212231213230312-3102112110221211-1100022003031232-3320033210021133-3132120212202021-1320121122013221-2330103113101111-0102012112312132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321)
- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-2321123312320130-2111000211032220-1333121322132313-1131131203330213-2101113330320112-0233030331133212-3213000133011220-1311302102031003"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200202220202113-1233210102200032-0303032123211230-0202111233000330-3233000010032113-2202001232222102-2310112122300132-2203201132101302"></a>

### Direct properties for `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1302213220323032-2213302113221312-1100321203333312-0211030210031110-1133330101223231-1001331110210202-0232121203212031-1331123031220033"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-3031231230232321-2202211213301132-2010033002031122-1121212201212003-0321300232233112-2103202301121321-3231211211211130-3211031200233223"></a>

<a id="canonical-0000000100231011-2301212013033322-3103002001323233-0121332300321130-3233211120021030-0133200101222223-3312130210323111-2010022220200232"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2231012030000121-1030312220211210-0102130001311301-2120021100022123-2023000313121213-2003320333302013-2203312112021223-2213121000121211"></a>

<a id="canonical-0012223231311133-0333210333132131-1120323223221312-2113220023112201-0113233223033021-2302022002001110-2112130202121022-0211310221230221"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-2100100321020011-0120311233112012-3121323322113103-2310100301231321-3030210112013001-2301002311012321-0210103011301002-1211023132203232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321)
- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2110221023333101-2333331023011323-3300301203303210-1212033110112011-1020130110030311-1033313332230122-2311312010322022-0013101011333132"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0300110310302021-1310323313121210-0120011111012211-1332132311312110-3233030302033030-1010003002131212-0303030201203101-1010103032310320"></a>

### Direct properties for `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-1203202002331030-1302322301331032-1232033122211020-3313102011212222-0130033230303220-1200233231100122-2101310011003312-0122111123003003"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0320030231133333-1122012032301220-1000332032033100-1030203132012230-1212130021001100-3110022110103210-3203210323000102-0003130022123020"></a>

<a id="canonical-0111102033300320-2001111311230313-0021030220122302-3020222222002003-2201100300312200-2102312211023323-3303232030310333-1000122123233322"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_headers_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.request_headers_to_add

<a id="canonical-2022011310031200-0013303231220310-0010311312321030-2121221330003233-2132210232231103-0223031231323331-0303001302312200-3001132120123023"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301023222200221-1213303102301331-3110203022303312-3013221121210030-1322012123013133-0120221102313122-1303003213002132-0220121300321021"></a>

### Direct properties for `routes.simple_route.advanced_options.request_headers_to_add`

<a id="canonical-2223222320001202-0311023120022130-3202111032220231-0010132132011203-0132331000121303-3301001213331220-2211111301330213-3220222211200023"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.append` property

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Additional upstream details:

If true, the value is appended to existing values. Default value is do not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3210020222301333-3302221031223113-1232211003232012-3102323322000220-3332021022213113-0221303313110200-2002022111210112-3110322113303120"></a>

<a id="canonical-1300302012033230-3300001333212232-3213302012011221-0231111020002122-1322321122020023-2313020112200010-0001130322130030-2022330231302003"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the HTTP header.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120): complete subsection reference.

<a id="canonical-2311012021002332-2130231220030200-1113301213210230-3202032333030013-2111111303033303-0232310320232303-2113012011312131-0010212130033231"></a>

<a id="canonical-1322323303122211-1111330021310333-1320313002003232-0310323023021133-2321302312022231-2312233231003222-2310303312000323-2100323220331030"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value

<a id="canonical-0103233013032131-3102300312132100-3232201300122033-0013003323321312-0223302212101113-1333111120303012-2231331233310322-1301201022002320"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101133230113233-0302320303220001-1210002210013132-2321123100302033-2102113301222213-2200323210030200-0133323132231102-2333300103101013"></a>

### Direct properties for `routes.simple_route.advanced_options.request_headers_to_add.secret_value`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-2010300000023322-0213320200330302-1131301130213121-2121133311202221-1220022210031330-0132020101323311-0330010132222303-1023120132222101): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-3333203031110332-1013301011031232-1122112311011033-3030310211111231-2320012202112313-1233233131230221-0102310210020010-1113102130122120): complete subsection reference.

<a id="canonical-2010300000023322-0213320200330302-1131301130213121-2121133311202221-1220022210031330-0132020101323311-0330010132222303-1023120132222101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313)
- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0331110020333233-1120231003031120-1001123130102012-1012232131300200-0212201213113003-2220123333233300-3323333101010322-2022232123233221"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203231113230310-1331231300211100-3111033032310323-2030300130322012-0301033112202002-2332131332033301-2313100230231023-0222332311001123"></a>

### Direct properties for `routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-0012123013230032-1331120113202120-2230313323302000-3321110220111020-1101012322133111-0030131201030303-1113132132321122-1203223202313010"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-0033333121002323-3000030022322320-3221033131123210-1230123230120210-0301013312123322-3130300033323003-3000330032310201-0321203011111010"></a>

<a id="canonical-1111100202110020-1202202111210001-0232201330121200-1011110232010222-1312300033333213-0333310231123111-1113013201130123-2221302113202022"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1210311213320121-1332222211221220-0130313311221332-0131023303121203-2200133203202300-3031101120302302-1012000022211133-0113003331232011"></a>

<a id="canonical-0321210121020233-1220231303131033-0302022220232022-0110332021122320-1200212032012203-0230001323221213-2233201200222021-0023310222030021"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-3333203031110332-1013301011031232-1122112311011033-3030310211111231-2320012202112313-1233233131230221-0102310210020010-1113102130122120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313)
- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-1101111300232320-0212013002313233-3301002021322303-0320032200233011-0212202323331121-2031332022330031-1333131002310033-0110031110103331"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012121223223110-3220132031031200-3233321011112031-0300023223103003-1232100032201223-3210322331330110-2113212001011003-3313331113332022"></a>

### Direct properties for `routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-1310303211021213-1331123110201200-2003332020310022-0300000312033110-1013020023332312-2101320002111002-0003001231231221-1332322230323133"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2230012223313332-1201321331230010-3030332322300321-2313111022202022-2002222333331103-3112033210322113-0010023022003013-3033021003132002"></a>

<a id="canonical-2221323101033333-3311301013113113-0122002111310200-3330033320223033-0322020230032133-0223230113002232-3321101331230120-3003320112003102"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.response_cookies_to_add

<a id="canonical-2202302222030131-2310030001022313-0132123230331302-2333213311020302-1102132333301132-1301021330220120-2122002133211030-1221303312023122"></a>

Type: `"object"`. list nested block, Optional.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_domain",
    "ignore_domain"),
  validators.ConflictingListObjectAttributes("add_expiry",
    "ignore_expiry"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_partitioned",
    "ignore_partitioned"),
  validators.ConflictingListObjectAttributes("add_path",
    "ignore_path"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "secret_value"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "value"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
response_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113103010301100-2233121001213031-1023011100311220-1110021210202332-2322032220031100-3001030100303301-0312031112232311-0102311101313102"></a>

### Direct properties for `routes.simple_route.advanced_options.response_cookies_to_add`

<a id="canonical-2223300310031313-2033311311202321-3131200332101212-1233110331233031-3021122100332331-3120011311333201-3123133322300223-2231330233332020"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.add_domain` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2311130213121103-2330332121000222-3123021003013100-2021233200323003-3031303223312220-0322212320021001-2021211031211121-1003001212331003"></a>

<a id="canonical-2122100303213132-0122332032230320-1310331203101202-0312033233313103-3122231313310233-1120233102220322-1021221120211201-0100122301003122"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.add_expiry` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_httponly](resources--http_loadbalancer--reference--group-026.md#canonical-2313030222311331-0310301120323200-3212110122231121-2111000013022300-2100300222033231-2110122110010311-3021221230021320-0223003232223323): complete subsection reference.

- [add_partitioned](resources--http_loadbalancer--reference--group-026.md#canonical-1323232012020231-1323232033333100-0022200303021233-1132201023301022-1311100300203202-1032201201220220-2212131302003003-2302012011131123): complete subsection reference.

<a id="canonical-2122310233111221-0023021131031103-3202000210110231-2000103032001030-1022202202213222-0220130330120031-0022100101323231-2330213030323230"></a>

<a id="canonical-0201133003000312-3201210031122103-1131212311112333-0120210233222032-0132222100110131-3221330103002301-2302302213200323-2101332312221312"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.add_path` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_secure](resources--http_loadbalancer--reference--group-026.md#canonical-0213002321303031-3200322021113130-3120313220233100-0313331231203332-0301220311230132-0033000312233230-0221310333001103-3012223000332102): complete subsection reference.

- [ignore_domain](resources--http_loadbalancer--reference--group-026.md#canonical-1231230310231311-3031012003130120-0203331001131021-3303121233230332-0000320021103021-2122031210201132-0221212102002000-1020101120000013): complete subsection reference.

- [ignore_expiry](resources--http_loadbalancer--reference--group-026.md#canonical-3000330231110221-3010013022021133-0111331201233132-2023302321123223-1133033122331331-1310300003223121-1301302023322230-0130102220120122): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-026.md#canonical-3033003021303110-2221320232111022-3210213120013012-1202002313131101-0111030110111110-3121230031330010-1320203302233301-2101101020133230): complete subsection reference.

- [ignore_max_age](resources--http_loadbalancer--reference--group-026.md#canonical-3332100132012110-1201031310330221-3300001231012323-2123023123313110-1320002032122101-2110030113013200-3000320302313022-0313132123300010): complete subsection reference.

- [ignore_partitioned](resources--http_loadbalancer--reference--group-026.md#canonical-0011012023201321-2203022033203003-2130131111312003-1202010321023003-2031320331333121-1133130221220202-3130312111111113-3112011213003301): complete subsection reference.

- [ignore_path](resources--http_loadbalancer--reference--group-026.md#canonical-3130303020100023-1032211010221013-0011013101110003-2100102032021202-2002232131231101-1112031231033121-3231013030113202-1311212333001023): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-026.md#canonical-1113003101311201-1030323301021312-3320212231220300-2131010001120201-1300010123101211-1021011203220320-2223002133203022-3330320022311103): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-026.md#canonical-0121100310301220-3302022132310000-0002023211103222-1331230311320323-0101322131033012-3212012122033212-0102321221032303-1231203220020120): complete subsection reference.

- [ignore_value](resources--http_loadbalancer--reference--group-026.md#canonical-0330032312013200-2213310211110302-3311111113310303-1020202111102302-2331212002020232-1032313002232013-1211000230200012-1203301033112132): complete subsection reference.

<a id="canonical-2331330333332022-3313133111012021-2221303223013301-2013323220323210-3101023033300033-2022202103103233-3120012221103003-3333110313110321"></a>

<a id="canonical-0223222311201000-0333333111113200-1023331221133033-0102201031012330-0313310010310201-2301000000012330-2003003033301233-1101311231102222"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.max_age_value` property

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-0022013303312003-0001333102103022-0033131321220310-2321221001202003-1131023330011210-3133003132313012-3111210123001223-3232103332003212"></a>

<a id="canonical-2313100113221130-0112330031101310-3233333010220102-0101232220013003-3031113011101301-3202313121210312-0132003223023220-2302001110310212"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

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

<a id="canonical-3323321121121000-0220301213210022-0331212230231233-1203113332232120-1312312323231101-2310231303113030-0321310211202210-2321122012002332"></a>

<a id="canonical-0212222210031222-1212231220011123-2201322302023120-3321230213303303-1100232031030221-1222322103203311-3100300010310112-1122011101232313"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.overwrite` property

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Additional upstream details:

If true, the value is overwritten to existing values. Default value is do not overwrite.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](resources--http_loadbalancer--reference--group-026.md#canonical-2122100130200302-2010033202110221-2101223332201321-1312310321132322-1313033110030110-2323000311312221-3030331110000103-0032112303312122): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-026.md#canonical-3213021033222202-3012330311220203-1130031023201103-1302120001111331-3030131223120232-2212010231121301-0013330311320112-1233330332212232): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-026.md#canonical-0111123301032121-3302113132300132-1031023023123223-1322221302020311-0103222130111210-2302010031223231-2032223221033320-2033012013022313): complete subsection reference.

- [secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-0203212013131210-3112233202311012-1211021022312223-2301111332221021-3300232131111021-2003332002013101-3123100331013113-1132013121310321): complete subsection reference.

<a id="canonical-2021312331201033-3233002311300333-0122132332121013-2302323033100220-0302020022001300-1203033020002030-0312031321222123-2131021211120333"></a>

<a id="canonical-3120323232301130-3203221230333300-1301122110112203-3313102032202233-1320110321201232-3101320322232210-1130230332010301-2232022000131113"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-2313030222311331-0310301120323200-3212110122231121-2111000013022300-2100300222033231-2110122110010311-3021221230021320-0223003232223323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.add_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.add_httponly

<a id="canonical-1213203130131311-0320303331203300-1100201221200310-1121213002001330-1220321233222002-3033311021113201-2102200302022331-2310223231010123"></a>

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

<a id="canonical-1323232012020231-1323232033333100-0022200303021233-1132201023301022-1311100300203202-1032201201220220-2212131302003003-2302012011131123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned

<a id="canonical-0331010332023100-1232113001202213-0310220321001301-3121003003230311-1113112012213221-2331021120003033-3023232020302212-3121000320022130"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add partitioned.

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
add_partitioned = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213002321303031-3200322021113130-3120313220233100-0313331231203332-0301220311230132-0033000312233230-0221310333001103-3012223000332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.add_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.add_secure

<a id="canonical-2210203101310221-1121133030010130-2131212201233021-0233312101110333-2112131320112130-2000032233312100-2300013011212310-0121232132031011"></a>

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

<a id="canonical-1231230310231311-3031012003130120-0203331001131021-3303121233230332-0000320021103021-2122031210201132-0221212102002000-1020101120000013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain

<a id="canonical-1321232010010302-0301323230010302-0331121101130230-3330011201212113-2123000011311232-2211313311311320-3212122022100031-0110201222002321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore domain.

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
ignore_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000330231110221-3010013022021133-0111331201233132-2023302321123223-1133033122331331-1310300003223121-1301302023322230-0130102220120122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry

<a id="canonical-0201022101233001-1333131213020323-3333312100013310-3311113312220202-3231020310103120-2331311302010033-0023302123323102-0223031313221300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore expiry.

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
ignore_expiry = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033003021303110-2221320232111022-3210213120013012-1202002313131101-0111030110111110-3121230031330010-1320203302233301-2101101020133230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.ignore_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_httponly

<a id="canonical-2013012112312012-0010012333132023-0330113013023311-2121211230123303-0213212332021002-2102030013022211-1001310000312030-0133011101313321"></a>

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

<a id="canonical-3332100132012110-1201031310330221-3300001231012323-2123023123313110-1320002032122101-2110030113013200-3000320302313022-0313132123300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.ignore_max_age` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_max_age

<a id="canonical-0232310221211123-2322202222123001-0022002313313131-0010323032030030-0010330320202320-1321033132220303-1330301302020211-3332203003130301"></a>

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

<a id="canonical-0011012023201321-2203022033203003-2130131111312003-1202010321023003-2031320331333121-1133130221220202-3130312111111113-3112011213003301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.ignore_partitioned` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_partitioned

<a id="canonical-0321022301331033-0010222123332202-2202200320110023-2313003312020222-1123102233300003-2321023200123232-0131032011131222-1013033011202002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore partitioned.

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
ignore_partitioned = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130303020100023-1032211010221013-0011013101110003-2100102032021202-2002232131231101-1112031231033121-3231013030113202-1311212333001023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.ignore_path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_path

<a id="canonical-1132221220132031-2203322002022021-3203032023232331-0132313021132102-3200300221100132-3332311220311223-1300303210201330-1202013202100223"></a>

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
ignore_path = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113003101311201-1030323301021312-3320212231220300-2131010001120201-1300010123101211-1021011203220320-2223002133203022-3330320022311103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.ignore_samesite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_samesite

<a id="canonical-2020232020222332-1010030113130321-2131222102010120-1022220300223032-2031000121001333-0212331221331111-2031301022223211-2131013232121302"></a>

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

<a id="canonical-0121100310301220-3302022132310000-0002023211103222-1331230311320323-0101322131033012-3212012122033212-0102321221032303-1231203220020120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.ignore_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_secure

<a id="canonical-1103212002112230-3033121222223213-1212110202321302-0021121323023122-0232230030021202-2233122001103300-1031011202210212-0303203020131013"></a>

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

<a id="canonical-0330032312013200-2213310211110302-3311111113310303-1020202111102302-2331212002020232-1032313002232013-1211000230200012-1203301033112132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.ignore_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_value

<a id="canonical-3021113133012033-0100300103233230-2332312233222012-2133313311111320-1031131103121021-2133020300103221-1200000021331021-0313112123011321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore value.

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
ignore_value = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122100130200302-2010033202110221-2101223332201321-1312310321132322-1313033110030110-2323000311312221-3030331110000103-0032112303312122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.samesite_lax` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.samesite_lax

<a id="canonical-2102130012003111-2212232132223221-1021310112321301-1121223001321102-1202102322213013-2230000303011330-3220301030113311-0232200212301331"></a>

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

<a id="canonical-3213021033222202-3012330311220203-1130031023201103-1302120001111331-3030131223120232-2212010231121301-0013330311320112-1233330332212232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.samesite_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.samesite_none

<a id="canonical-2301332010021220-2222111301221210-1021303213323233-3311033311022031-1332301110032123-2113323100111013-0113011031311213-0232021012200022"></a>

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

<a id="canonical-0111123301032121-3302113132300132-1031023023123223-1322221302020311-0103222130111210-2302010031223231-2032223221033320-2033012013022313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict

<a id="canonical-3020200210230222-1131210121203310-0303213223322102-0201312333033303-1333032031133201-0310220310222021-3331122001002131-2111231330123232"></a>

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

<a id="canonical-0203212013131210-3112233202311012-1211021022312223-2301111332221021-3300232131111021-2003332002013101-3123100331013113-1132013121310321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.secret_value

<a id="canonical-2222031300212311-0302012131210311-1331232011232313-2310202320302311-0230300223110113-1020031113033101-1202320213312300-0302112132022210"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030003112200033-2112112130210230-1100300201222203-0202002221032001-3020312310030022-0201111032020311-1202311310001201-1121222132013200"></a>

### Direct properties for `routes.simple_route.advanced_options.response_cookies_to_add.secret_value`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-3121311132202321-1012200332320102-0112302311202113-1021020201113030-2220332302330000-2102212211202123-1221230222202001-0301323002311200): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-3021003020203031-2030001320211233-2313222110312133-2001120233112102-2032233010022330-0201031121033211-2220201010331010-3232333233200320): complete subsection reference.

<a id="canonical-3121311132202321-1012200332320102-0112302311202113-1021020201113030-2220332302330000-2102212211202123-1221230222202001-0301323002311200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-0203212013131210-3112233202311012-1211021022312223-2301111332221021-3300232131111021-2003332002013101-3123100331013113-1132013121310321)
- routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-2033311232331212-3020212021321330-2133302131201203-0233212311203003-2113312112013002-0100333110333203-3110020000232232-1110031102030012"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1213313110021201-1332331222223312-3333301101300010-3013220123100303-1201123010311001-3210311221121101-1230232333232023-1112000022320001"></a>

### Direct properties for `routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1232110030203333-3202230020110212-1331323202012100-3112110220103310-1033213111210030-3212313203231102-0110313023022331-2232111102311003"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-1221221133210021-1322011120132231-3132332203103321-3102322212211210-0323020213223332-3320212320213332-1021132111232232-1122121133131220"></a>

<a id="canonical-1200103201211121-3121302033130021-2021303021231300-0222112102220031-3021111123110311-0211202303301031-0003102202222110-3223031203201032"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2303132022221212-1100010300020020-3133300320003021-1001212311001011-1010103233030111-1331022233230232-2113301102230210-0312321011113010"></a>

<a id="canonical-3100310230131313-1033111110300033-2203222130003011-3310013220023003-3121112030013210-0011302121221030-1022200011102300-1130033230113021"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-3021003020203031-2030001320211233-2313222110312133-2001120233112102-2032233010022330-0201031121033211-2220201010331010-3232333233200320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-0203212013131210-3112233202311012-1211021022312223-2301111332221021-3300232131111021-2003332002013101-3123100331013113-1132013121310321)
- routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2300220123011312-1330010301322330-0101020200122310-1033103213120102-0231233033331332-3013100010311303-1111102001230233-0130200030000321"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333020233133120-3232230121120123-0113231100302031-3013102233030131-0111131103201010-0022200001102230-0212101121102030-1010321032131202"></a>

### Direct properties for `routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-2330203311321033-1113211323302013-0100030302223121-1023322322233302-2231212322333221-1331120012302230-0100101300212103-1030302221002313"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0000110333102001-2230222312121101-3012123131210133-3112131123200322-1000311313002213-0021023022301101-1132133111313100-2322303222101231"></a>

<a id="canonical-2001003122231323-1212013032330101-3302203030303330-1201113333210331-1121231110111222-3221130130122220-2020302011032100-2311232123323023"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1302033232022313-1021121312312133-3011110120013233-0011303102032230-1312323112002300-2310303030311201-2003333031333222-2030332033023230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_headers_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.response_headers_to_add

<a id="canonical-1210022120322001-0033112023320210-3230033333202001-3022033033302222-0231123002010212-0311311201113133-0133010231311123-3311323322202311"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
response_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231031212110121-2200230201331132-0212223200203222-3001301300231313-2113303013220000-2133310330222120-0130222222112330-1022111131022102"></a>

### Direct properties for `routes.simple_route.advanced_options.response_headers_to_add`

<a id="canonical-0100222312213003-0121330313122333-2113320012122003-1323100332013230-1232321000213230-0231320203200223-3222201213031222-1312002112233010"></a>

#### `routes.simple_route.advanced_options.response_headers_to_add.append` property

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Additional upstream details:

If true, the value is appended to existing values. Default value is do not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2213223122223002-3012320021321113-3133311232010013-2020131111112223-3212032003001133-0002323332230131-3012111131111102-1311120111303221"></a>

<a id="canonical-1131131312310213-2103211000001312-3333310331333220-2123102233021303-2021213010011112-3000211032331233-3201120133111233-1210320200312123"></a>

#### `routes.simple_route.advanced_options.response_headers_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the HTTP header.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-0321020001320130-1220312103300231-0231232211203031-0310102123320200-2112203130120213-2332322311001113-2210220322200030-2020022202121322): complete subsection reference.

<a id="canonical-2302331321000210-3300100321023013-3300222111033013-1020013001031032-0213003203110302-2011102310103101-0100011210200001-1301030220111311"></a>

<a id="canonical-0212210213021213-2201330010131123-0231120313103103-3132220330022122-1020231202230011-1022003200200010-3211310000010013-3100133303030112"></a>

#### `routes.simple_route.advanced_options.response_headers_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-0321020001320130-1220312103300231-0231232211203031-0310102123320200-2112203130120213-2332322311001113-2210220322200030-2020022202121322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_headers_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-1302033232022313-1021121312312133-3011110120013233-0011303102032230-1312323112002300-2310303030311201-2003333031333222-2030332033023230)
- routes.simple_route.advanced_options.response_headers_to_add.secret_value

<a id="canonical-2100321003221100-0201101313121020-3130122011223130-1012210103321200-1230212233000301-0101030020121320-2103312113100123-2122302223123300"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030000312330101-1232022210302212-2132120333012303-1022220110202111-1202021332133320-3322031110331320-2321013203021112-0132021112301322"></a>

### Direct properties for `routes.simple_route.advanced_options.response_headers_to_add.secret_value`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-2330333011323210-1222031310002303-1011112333101220-2230300132022003-3332203203332122-1330313330011232-2111322300331110-2320103330121110): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-1212122203311012-1201112221300103-3321131203133111-1010133130221001-2000111213013213-1200201211301310-0132013230032012-2112112311221200): complete subsection reference.

<a id="canonical-2330333011323210-1222031310002303-1011112333101220-2230300132022003-3332203203332122-1330313330011232-2111322300331110-2320103330121110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_headers_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-1302033232022313-1021121312312133-3011110120013233-0011303102032230-1312323112002300-2310303030311201-2003333031333222-2030332033023230)
- [routes.simple_route.advanced_options.response_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-0321020001320130-1220312103300231-0231232211203031-0310102123320200-2112203130120213-2332322311001113-2210220322200030-2020022202121322)
- routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-2112222123121200-2311313331303021-0102303121222103-0232323300001313-0310313302212133-0332220121311232-3302000010112301-1103320312130210"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320131002232221-3323033322001011-0023113320322302-3010213200010132-1112102311003233-1211122221112013-3033113110331222-3021201212121123"></a>

### Direct properties for `routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3013130221030122-3132000221133113-3330212010122133-2312311101122103-2000232311020323-0110110231310003-1213223101330220-1011021132331011"></a>

#### `routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-0300011033223132-3003310002200101-2031112303232322-0303301033122302-2310002113230113-3312112013321333-2220111121321320-0302321220120111"></a>

<a id="canonical-1000302112331323-0230223312031300-3203021121322212-1003230111200330-1102201322312023-3100121123220000-1112103011310032-0220322101113231"></a>

#### `routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3012132010100301-1000300230113013-1000111302013030-3132002100012221-2231210001122301-1223201330331201-3330022030213002-0320000011321332"></a>

<a id="canonical-1002301130021223-0213031211022010-1303033100302033-1332222220201021-2033013212020131-1100320331332310-0130201132132132-0313132102120122"></a>

#### `routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-1212122203311012-1201112221300103-3321131203133111-1010133130221001-2000111213013213-1200201211301310-0132013230032012-2112112311221200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_headers_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-1302033232022313-1021121312312133-3011110120013233-0011303102032230-1312323112002300-2310303030311201-2003333031333222-2030332033023230)
- [routes.simple_route.advanced_options.response_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-0321020001320130-1220312103300231-0231232211203031-0310102123320200-2112203130120213-2332322311001113-2210220322200030-2020022202121322)
- routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-0122321311331100-1113311232203030-2132033123310011-0032301020030031-1203303331103013-0322330031312002-1000103132022122-0333021002002021"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1330123122011132-0021102321220230-3122211130310332-2113232210333000-0032313210120313-1032133120112132-2202120200033330-0031011030202212"></a>

### Direct properties for `routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-1211211330133022-1120023123023332-1022222020130221-0110020213120122-0110330311212302-2302013201122003-3003220311203322-3112211331331230"></a>

#### `routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1320122123010310-2221013010201111-1213231020120322-1000032122303231-3212131000013333-3200310232033131-1223212223101331-3012320101111211"></a>

<a id="canonical-0331120100330310-1123331000031003-0300021311101022-3023230002120212-0012011131220201-3133311213332020-2001132113120101-2001030221300312"></a>

#### `routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1001202133030320-0013120123332230-2010322030001110-0213023101000331-3211331232101131-1213100001010311-0221112031110323-2311021231132111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.retract_cluster` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.retract_cluster

<a id="canonical-1021132220311000-2011230130033233-2113001221010303-0012301322331011-1220102230321013-2122220020030332-1320000132013113-2122022130032201"></a>

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
retract_cluster = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031213021312332-2113132313130100-3222333013111133-0122010300022110-1323302212201331-0133200321222002-3221100221013001-0011102330211230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.retry_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.retry_policy

<a id="canonical-3013032133312203-3302203013023013-3223203121120302-3001303323102100-0201322010033331-3210113310110013-3122310310313213-2301011130100222"></a>

Type: `"object"`. single nested block, Optional.

Retry policy configuration for route destination.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("retry_condition")}
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
retry_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0112133331100200-2202211013101200-0233102011200132-0021031022031203-3221203330032111-0000310332120013-2112303011122312-2102331333233112"></a>

### Direct properties for `routes.simple_route.advanced_options.retry_policy`

- [back_off](resources--http_loadbalancer--reference--group-026.md#canonical-0021013312320103-1311301210231102-2021303000233200-1302233022111221-2330133211121310-2112222031310030-1000321202203303-3102302031102203): complete subsection reference.

<a id="canonical-2110131320331220-2301201120203020-1132220331221203-1322110100221103-2210032033311201-1001030110222202-1213030131211313-1230313013131030"></a>

<a id="canonical-2023101333230331-3303022011230232-0313001021100031-2021312113313201-2011332000323132-1322233110131100-2301320210322021-1002320231310221"></a>

#### `routes.simple_route.advanced_options.retry_policy.num_retries` property

Type: `"number"`. Optional.

Specifies the allowed number of retries. Retries can be done any number of times. An exponential
back-off algorithm is used between each retry. Defaults to \`1\`.

Additional upstream details:

Defaults to 1.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  }
}
```

<a id="canonical-1233011122000313-1031233033233002-2013313132010133-1133113101303003-3000130023022113-3202220020332302-1311120221320211-3122302012332231"></a>

<a id="canonical-3310101011203203-3032121000323001-0202331032003021-1122133121032131-3112022113022330-0123022010220323-2022011112113000-3201132222221113"></a>

#### `routes.simple_route.advanced_options.retry_policy.per_try_timeout` property

Type: `"number"`. Optional.

Specifies a non-zero timeout per retry attempt. In milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-2010123123301132-2103103321123001-2212200231033112-3302312300021333-2023232123001301-1212200201201030-2001112101031210-3002113021311310"></a>

<a id="canonical-1322203320202302-1031221121110021-1101001323001100-0102132102103021-1032133023300222-0110203001022100-1112330321011032-0202223001123003"></a>

#### `routes.simple_route.advanced_options.retry_policy.retriable_status_codes` property

Type: `["list", "number"]`. Optional.

HTTP status codes that should trigger a retry in addition to those specified by retry\_on.

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

<a id="canonical-3010230002323300-2100330012300030-2113301113223211-2000323302122211-1200113213003300-0210333320102111-1020211020020102-2032301012030203"></a>

<a id="canonical-1021223000330211-0031121013220220-1202033223303112-3110002012232022-3131311320030131-3012320002032200-1321221011223103-1032110221122230"></a>

#### `routes.simple_route.advanced_options.retry_policy.retry_condition` property

Type: `["list", "string"]`. Optional.

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if
the..

Additional upstream details:

For example, network failure, all 5xx response codes, idempotent 4xx response codes, etc

The possible values are

"5xx" : Retry will be done if the upstream server responds with any 5xx response code, or does not
respond at all (disconnect/reset/read timeout).

"gateway-error" : Retry will be done only if the upstream server responds with 502, 503 or 504
responses (Included in 5xx)

"connect-failure" : Retry will be done if the request fails because of a connection failure to the
upstream server (connect timeout, etc.). (Included in 5xx)

"refused-stream" : Retry is done if the upstream server resets the stream with a REFUSED\_STREAM
error code (Included in 5xx)

"retriable-4xx" : Retry is done if the upstream server responds with a retriable 4xx response code.
The only response code in this category is HTTP CONFLICT (409)

"retriable-status-codes" : Retry is done if the upstream server responds with any response code
matching one defined in retriable\_status\_codes field

"reset" : Retry is done if the upstream server does not respond at all (disconnect/reset/read
timeout.)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 7),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 7,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 7,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0021013312320103-1311301210231102-2021303000233200-1302233022111221-2330133211121310-2112222031310030-1000321202203303-3102302031102203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.retry_policy.back_off` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.retry_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2031213021312332-2113132313130100-3222333013111133-0122010300022110-1323302212201331-0133200321222002-3221100221013001-0011102330211230)
- routes.simple_route.advanced_options.retry_policy.back_off

<a id="canonical-1100032210033302-2203220332130211-2202223220232330-0122302210010302-2203200232313200-0030213021110012-1230101102301302-3131230331102133"></a>

Type: `"object"`. single nested block, Optional.

Specifies parameters that control retry back off.

Receipt-pinned upstream constraints:

```json
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
back_off {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132332100320121-2222112231330013-2301001212310301-2011021201112123-0020301303322203-2120131021120032-2110202000110320-3130302212031331"></a>

### Direct properties for `routes.simple_route.advanced_options.retry_policy.back_off`

<a id="canonical-3030130121133020-0111033021102320-1320123110303222-2302111320202033-2233302202123113-3223012232313213-2211330300110211-2320011121111332"></a>

#### `routes.simple_route.advanced_options.retry_policy.back_off.base_interval` property

Type: `"number"`. Optional.

Specifies the base interval between retries in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(1),
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
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

<a id="canonical-0013233200111213-3200003212033100-3313312223230010-3031221212301220-2332132021101200-0130012113221213-0313102133311302-3112301012031200"></a>

<a id="canonical-0123300023110100-0101131133033002-0231322021111111-0110322122302102-1233303010311022-2022322210031012-3101111110330011-2320301300303300"></a>

#### `routes.simple_route.advanced_options.retry_policy.back_off.max_interval` property

Type: `"number"`. Optional.

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The times the base\_interval. Defaults to
\`10\`.

Additional upstream details:

The default is 10 times the base\_interval.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2332103311111231-3001202122300101-0003322101013201-2300220101202301-0330222321301302-3233120102200011-0001011001021100-2330220212112132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.specific_hash_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.specific_hash_policy

<a id="canonical-0022223130113010-3111110220112001-3003031002130320-3113100010002320-0310001102212232-2222330223032012-0103213231120203-1100032121013322"></a>

Type: `"object"`. single nested block, Optional.

Policy configuration for this feature.

Additional upstream details:

List of hash policy rules.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_policy")}
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
specific_hash_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310221102103100-2113312213023231-0020312333323030-3310013130330232-3230320122031110-2033302211120221-1322330120232200-1123303201301113"></a>

### Direct properties for `routes.simple_route.advanced_options.specific_hash_policy`

- [hash_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2333120110011121-1313332102102032-3201233223100103-3211320101012332-3032321222123131-2001321033010020-3333231101303312-3110032111130110): complete subsection reference.

<a id="canonical-2333120110011121-1313332102102032-3201233223100103-3211320101012332-3032321222123131-2001321033010020-3333231101303312-3110032111130110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.specific_hash_policy.hash_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.specific_hash_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2332103311111231-3001202122300101-0003322101013201-2300220101202301-0330222321301302-3233120102200011-0001011001021100-2330220212112132)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy

<a id="canonical-2021212011023320-0013311121031020-0301203022111231-2201320230211313-2203202012332320-3211000332311013-2120101233111003-0331321003312103"></a>

Type: `"object"`. list nested block, Optional.

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("cookie",
    "header_name"),
  validators.ConflictingListObjectAttributes("cookie",
    "source_ip"),
  validators.ConflictingListObjectAttributes("header_name",
    "source_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
hash_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313220231210312-3003330301300210-1013332332032133-1010332222020233-0032221023221001-0222220230102221-1332233201220120-3023002300323311"></a>

### Direct properties for `routes.simple_route.advanced_options.specific_hash_policy.hash_policy`

- [cookie](resources--http_loadbalancer--reference--group-027.md#canonical-2130110203221012-1121230202302212-2012030120310003-0211322323023221-1302222213221201-1132121033023133-3210311223230130-1300120322313020): complete subsection reference.

<a id="canonical-3322221002112220-0310221322020102-2022233300121202-2021121212202201-3121030212232002-3331313313131112-2003112321021000-0033030232213101"></a>

<a id="canonical-3111312011302203-1110302210230131-0202110022302212-3302103021003000-0002011220022133-0320333300120332-1030023221033330-0122110330211101"></a>

#### `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.header_name` property

Type: `"string"`. Optional.

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1230000303012311-2120013321231100-0312003231232102-1312222203221023-3310123212320331-2033002313020321-2132233112102013-0212302030231112"></a>

<a id="canonical-3120331002200202-3102012300203210-3030001333101033-3001331002110330-0132223000303132-0322112230220112-0303010021200013-1112223000322111"></a>

#### `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.source_ip` property

Type: `"bool"`. Optional.

Exclusive with \[cookie header\_name\] Hash based on source IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0210103011122330-2121103322003203-0312121012331202-3333122233120032-2032101313022311-2212221303231320-0320210210331000-0203301222323020"></a>

<a id="canonical-2020221012121020-1331333211101213-1133101123120112-0332230213003103-2033222031012100-1012301320302201-0332323211020200-0023013301332010"></a>

#### `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.terminal` property

Type: `"bool"`. Optional.

Terminal. Specify if its a terminal policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
