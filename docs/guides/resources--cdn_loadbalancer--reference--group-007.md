---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-2310021211112122-3111330301033122-1113132301023002-2211100323033223-1030122321200000-1013331231012201-1130211300331032-2102302201211200"></a>

## `blocked_clients.http_header.headers.regex` property

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2200200000300132-1332301203032202-2100200211211313-3111120203213220-3330122020102220-1312102213133013-3200101022303223-1200011131230313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-0101200112203333-2313100322333123-1133301001021123-0132230132312103-2222120211222232-1330312333203021-1213021020200201-2312220003310232)
- blocked_clients.metadata

<a id="canonical-3233123130130120-0000022110120132-1311332230033000-3111220231133123-2122320212130223-2301111301110032-2000133100022100-1112110330133203"></a>

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

<a id="canonical-1123023332302112-2123112113301121-0222132012013123-3023011122323100-0320330313232232-2201033110323033-0032100320123223-3222203112033021"></a>

### Direct properties for `blocked_clients.metadata`

<a id="canonical-0213203010313133-1033010022303031-3123302222203032-1130013310231322-0013211121303303-2310323232230322-2023232230033013-2333321023312032"></a>

#### `blocked_clients.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2132320113103022-3133132302312111-3031310022033222-2023110312313212-1130321123311303-0200032132311120-3232312211003332-1233032201302132"></a>

<a id="canonical-1223320200102201-3021303133003203-2200330010001203-3111101020332003-3122223002210002-0202131321133113-2233211021130203-1102213212113200"></a>

#### `blocked_clients.metadata.name` property

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

<a id="canonical-2202101210222131-0232132212311320-2112012121130132-2133121132210112-1132033031020303-2231231321333131-0211031330112200-2221130212000313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-0101200112203333-2313100322333123-1133301001021123-0132230132312103-2222120211222232-1330312333203021-1213021020200201-2312220003310232)
- blocked_clients.skip_processing

<a id="canonical-3302200001332131-1100010112223102-2122233303001222-2203031132213320-2113003122030312-0310110332102012-3321101323200111-1132203312310303"></a>

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

<a id="canonical-1320031132031333-0322332200020113-1301203133232332-3232323223120032-1000133112100222-2233323333133332-1003310032331111-3012013120333133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.waf_skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-0101200112203333-2313100322333123-1133301001021123-0132230132312103-2222120211222232-1330312333203021-1213021020200201-2312220003310232)
- blocked_clients.waf_skip_processing

<a id="canonical-2123212213310320-0333131113101212-0012232020022321-2031011023320302-0002220200332121-3223033310222233-1320012023032010-3300323022113013"></a>

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
waf_skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- bot_defense

<a id="canonical-2000311131220203-1013013112311122-1031332010101300-3222113321032031-2112031330230210-2033321311130203-3322121100201010-3202000031000203"></a>

Type: `"object"`. single nested block, Optional.

This defines various configuration OPTIONS for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_cors_support",
    "enable_cors_support")}
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
  "x-ves-oneof-field-cors_support_choice": "[\"disable_cors_support\",\"enable_cors_support\"]"
}
```

Terraform syntax:

```terraform
bot_defense {
  # Configure direct properties listed below.
}
```

<a id="canonical-2003200012320130-0221013011001110-2232203102020200-2333303121120331-1000013320331331-2200030230011110-3301020213331021-2222010100211130"></a>

### Direct properties for `bot_defense`

- [disable_cors_support](resources--cdn_loadbalancer--reference--group-007.md#canonical-2110323220222022-1200020013313230-2012021003132203-1202331100201102-1112222110003322-2212332311211210-2030033313013123-1300002003210011): complete subsection reference.

- [enable_cors_support](resources--cdn_loadbalancer--reference--group-007.md#canonical-1130023222310301-2001120103202032-3003321211323123-1123113313100301-1031313022330103-2301233102331112-1310020101212311-2122332220301103): complete subsection reference.

- [policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111): complete subsection reference.

<a id="canonical-2002232020020132-0323030221013223-0202010111023322-2120331311013001-2113003212023301-2122333110101121-1120200011132003-3121102032002100"></a>

<a id="canonical-2333020020130101-3121302003312230-2013301100210332-2000102123203133-0313201021200131-2130231131130002-0231120112221112-2113203121303121"></a>

#### `bot_defense.regional_endpoint` property

Type: `"string"`. Optional.

\[Enum: AUTO|US|EU|ASIA\] Defines a selection for Bot Defense region - AUTO: AUTO Automatic
selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA
Asia region. Possible values are \`AUTO\`, \`US\`, \`EU\`, \`ASIA\`. Defaults to \`AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ASIA","AUTO","EU","US"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("AUTO",
    "US",
    "EU",
    "ASIA"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AUTO",
  "enum": [
    "AUTO",
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2000101121230031-1332003033312230-3323103101322101-1300100130120012-2221032220023033-0303312022233223-3302332021301211-2333211213201010"></a>

<a id="canonical-3130231102101110-2230210103023011-1330132120212203-2231201313332221-3312131132110321-0033312102300133-1300100113313323-1211100110211201"></a>

#### `bot_defense.timeout` property

Type: `"number"`. Optional.

The timeout for the inference check, in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 60000),
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
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-2110323220222022-1200020013313230-2012021003132203-1202331100201102-1112222110003322-2212332311211210-2030033313013123-1300002003210011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.disable_cors_support` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- bot_defense.disable_cors_support

<a id="canonical-1000113332122230-0221302010230000-1030001220302000-1212233221112202-2010101103111303-3121030310002302-0321112310023332-0230323232031231"></a>

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
disable_cors_support = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130023222310301-2001120103202032-3003321211323123-1123113313100301-1031313022330103-2301233102331112-1310020101212311-2122332220301103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.enable_cors_support` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- bot_defense.enable_cors_support

<a id="canonical-0223001222223213-0012200330232312-2211033301003222-2333303023022323-3032330123203130-2131132212221330-3013210323131231-0033023301033232"></a>

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
enable_cors_support = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- bot_defense.policy

<a id="canonical-2101120103201033-1320310231012320-1011310220333012-1330200013111022-1112311320312332-2020223102331030-2313230033112203-0100000320001210"></a>

Type: `"object"`. single nested block, Optional.

This defines various configuration OPTIONS for Bot Defense policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("protected_app_endpoints"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("disable_mobile_sdk",
    "mobile_sdk_config"),
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232200213321030-2313203123231112-1203211030330220-0303010320203312-3332210231032333-1310303300221200-1130023022022231-1322021330213132"></a>

### Direct properties for `bot_defense.policy`

- [disable_js_insert](resources--cdn_loadbalancer--reference--group-007.md#canonical-3113032133121203-2001232210033102-3203122102001200-2032012232202212-1230223221132133-2103202223101322-0013100213101123-3121201130103200): complete subsection reference.

- [disable_mobile_sdk](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120322022300212-0310022310212003-3032313020212132-1021133110312313-3031102102101121-3031010112312113-0133123103012203-0231310000332213): complete subsection reference.

<a id="canonical-0012113233011302-3002222312100303-3313220210303232-3121203022010220-3030311012011012-1123023130301301-1020102011210133-0330230023120312"></a>

<a id="canonical-2122200331221320-0020332223030120-1210203131102310-0103323310210122-3003112122303021-3211100003310223-0232020212333300-3222301110031120"></a>

#### `bot_defense.policy.javascript_mode` property

Type: `"string"`. Optional.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Additional upstream details:

Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is cacheable.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ASYNC_JS_CACHING","ASYNC_JS_NO_CACHING","SYNC_JS_CACHING","SYNC_JS_NO_CACHING"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ASYNC_JS_NO_CACHING",
  "enum": [
    "ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2113232011012300-0032031330323322-1332031032313213-3332022221222123-0230232233332202-1301210333130220-0310301332023322-3023320221020312"></a>

<a id="canonical-2301120220002130-0133201211102312-0311000013000133-3010201213312221-2111112113303230-3020231002122210-0310332212300233-1021200110002330"></a>

#### `bot_defense.policy.js_download_path` property

Type: `"string"`. Optional.

Customize Bot Defense Client JavaScript path. If not specified, default \`/common.js\`

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
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [js_insert_all_pages](resources--cdn_loadbalancer--reference--group-007.md#canonical-0211123212103031-2332332103330000-2013300321210120-1330332123230233-3030322322220130-2123121303221223-1211202023303210-0113321012012302): complete subsection reference.

- [js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-007.md#canonical-2033123012202321-1211101311022113-1201002111032111-3210022010212021-3321310021201313-3202021033020030-1101021222300022-1110131112213111): complete subsection reference.

- [js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-3101332010211112-3031030103123313-2120311030130113-1132001221131033-1103310102132033-0213103111312323-3003202121131230-3323212032330321): complete subsection reference.

- [mobile_sdk_config](resources--cdn_loadbalancer--reference--group-007.md#canonical-0211110332021202-1223322002233230-1010121033102111-2333021322220011-0303302202101022-2020222312011102-1313020332132132-3230110202201322): complete subsection reference.

- [protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013): complete subsection reference.

<a id="canonical-3113032133121203-2001232210033102-3203122102001200-2032012232202212-1230223221132133-2103202223101322-0013100213101123-3121201130103200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.disable_js_insert` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- bot_defense.policy.disable_js_insert

<a id="canonical-2031101020003311-2313303032123003-1122010212122122-2011012202211320-1310003300200021-3131333332201113-3030100233003211-2222332131210001"></a>

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

<a id="canonical-0120322022300212-0310022310212003-3032313020212132-1021133110312313-3031102102101121-3031010112312113-0133123103012203-0231310000332213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.disable_mobile_sdk` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- bot_defense.policy.disable_mobile_sdk

<a id="canonical-2223012123201123-2211003312020332-0321030312010130-1030023102232000-0101331331033303-0322330020102131-0223323032110203-0220121013321321"></a>

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
disable_mobile_sdk = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211123212103031-2332332103330000-2013300321210120-1330332123230233-3030322322220130-2123121303221223-1211202023303210-0113321012012302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- bot_defense.policy.js_insert_all_pages

<a id="canonical-2102110023323330-2003211321210003-0132022131012132-0011100202211231-3020021101010021-0011100323310320-0031200123123323-3202301022003312"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
js_insert_all_pages {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333212132300122-1330022012330202-0202112313312300-2131102100313232-2002230011103133-3133230111321332-3013032302011333-0313133220203322"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages`

<a id="canonical-3212021200302102-2100333220210112-1033111021311232-0013220123110301-3220012021113121-2310302331320012-1210012001123131-1220330330012113"></a>

#### `bot_defense.policy.js_insert_all_pages.javascript_location` property

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

<a id="canonical-2033123012202321-1211101311022113-1201002111032111-3210022010212021-3321310021201313-3202021033020030-1101021222300022-1110131112213111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- bot_defense.policy.js_insert_all_pages_except

<a id="canonical-1202113132300331-2023303003310131-3033211232100310-2311033113122330-1133323030112200-3101300101212033-3302301023012111-0311111031200313"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages with the exceptions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
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

<a id="canonical-2010321230021133-2012023102322132-3110110102222220-1300021032100230-1322302112111112-0013002302203303-3210020232130221-1122321300123111"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except`

- [exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-3321230333023200-0223133222102101-3223130312131313-1330202302222030-3133112013000101-0023222133230111-3321013222321031-3111022101032131): complete subsection reference.

<a id="canonical-1030301122101000-2031330202303022-1120101110001330-1202112031032002-0320013223301002-0221000130120221-2332102303121321-0111221003013332"></a>

<a id="canonical-1331200213222133-1012121311233301-2223220031110333-1221201120200231-1031322320100220-0233001222122200-2202003233110301-1003223203022212"></a>

#### `bot_defense.policy.js_insert_all_pages_except.javascript_location` property

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

<a id="canonical-3321230333023200-0223133222102101-3223130312131313-1330202302222030-3133112013000101-0023222133230111-3321013222321031-3111022101032131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-007.md#canonical-2033123012202321-1211101311022113-1201002111032111-3210022010212021-3321310021201313-3202021033020030-1101021222300022-1110131112213111)
- bot_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-1121201200211030-2302230111200033-2110311302312120-3313313012311301-3130331320312022-0101303113222333-0322032112120113-2312322321220312"></a>

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

<a id="canonical-3222001223112013-1220321302011103-3222100313011333-3110102323323110-0210301113013201-3302220212120202-1220111001201032-2200123031031313"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except.exclude_list`

- [any_domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-3103300313123000-3001011231233213-2022102321010130-2023203022311002-3230313010313231-2211121332002302-3100330221012002-1131103222200303): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-1010123100102101-0032020112331033-3313020122013231-0012010303012120-3031033211223231-0323122110120112-1221333110200223-2232223121122222): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-0310120220000003-0132123103212320-3333010210230232-1301232331111233-1330331033132210-0132112003023112-1331202233331132-3023313113230130): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-007.md#canonical-2103112101023012-1200031003310022-1313210003311110-0212312013022120-0112020103331032-1303030201103203-0101320121023023-1230212001022333): complete subsection reference.

<a id="canonical-3103300313123000-3001011231233213-2022102321010130-2023203022311002-3230313010313231-2211121332002302-3100330221012002-1131103222200303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-007.md#canonical-2033123012202321-1211101311022113-1201002111032111-3210022010212021-3321310021201313-3202021033020030-1101021222300022-1110131112213111)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-3321230333023200-0223133222102101-3223130312131313-1330202302222030-3133112013000101-0023222133230111-3321013222321031-3111022101032131)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-3330001302131120-3022133032323030-3121213333323203-1013030003123232-0002301222000311-0022123100121323-1313110233201300-1131103321011110"></a>

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

<a id="canonical-1010123100102101-0032020112331033-3313020122013231-0012010303012120-3031033211223231-0323122110120112-1221333110200223-2232223121122222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-007.md#canonical-2033123012202321-1211101311022113-1201002111032111-3210022010212021-3321310021201313-3202021033020030-1101021222300022-1110131112213111)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-3321230333023200-0223133222102101-3223130312131313-1330202302222030-3133112013000101-0023222133230111-3321013222321031-3111022101032131)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-2120201003321012-1123210220221023-0331022320323210-0300112112221123-1011031200131210-3300311213100120-3313023232021212-0203133121003232"></a>

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

<a id="canonical-0120023030202132-1000203300103212-3322210012321020-3331011332122201-2031131232113110-0220103000031303-0230232123232132-2020012213211023"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain`

<a id="canonical-3301210011012120-2212001020300310-2233022321100121-0101210310321032-1321132232230231-2011210211213300-0101010211210311-0111133033231113"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain.exact_value` property

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

<a id="canonical-3302101312230031-3122220030211202-0320331201001231-1113130101212133-3212003001130303-3210232230001112-2332010233231230-0330200011332310"></a>

<a id="canonical-3331233020110221-3302012101312020-0301311001103212-3110302201212021-0100031211333303-0201013203323203-0333033130222321-1233322102201212"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain.regex_value` property

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

<a id="canonical-2233021132320122-3222211302102121-2133303222131002-3012300003030201-0013311033103232-2230122303202002-1301200312013230-1202122031331223"></a>

<a id="canonical-3222311131113212-1121300120010230-1331002313111231-1323301222023302-0013123000322030-0301320332002110-3132202110232031-3012131120211010"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.domain.suffix_value` property

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

<a id="canonical-0310120220000003-0132123103212320-3333010210230232-1301232331111233-1330331033132210-0132112003023112-1331202233331132-3023313113230130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-007.md#canonical-2033123012202321-1211101311022113-1201002111032111-3210022010212021-3321310021201313-3202021033020030-1101021222300022-1110131112213111)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-3321230333023200-0223133222102101-3223130312131313-1330202302222030-3133112013000101-0023222133230111-3321013222321031-3111022101032131)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-0101013222001011-1112030112030031-3132221223122003-2220221130131120-0011021210010002-3023020320231110-2101001102332002-1123303303021120"></a>

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

<a id="canonical-0313312032232311-2300223233202213-3223000003303132-2013120313330322-2112301002300213-3020301233303201-2123032120003133-1323222011111222"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata`

<a id="canonical-2312331321332210-1123210131103000-1231111022003232-0113231232322132-2230013220022321-0021230130300033-2133303023303010-2303000000122102"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3113010030331303-3021210111112100-1202033321011100-3003322121002131-3111221321313303-2313220032331320-0100220223101032-2113113032111130"></a>

<a id="canonical-0310133222230200-1123213202010030-3213121230213332-0221120200023102-3011113232200302-0130203032002113-3111111133232022-3203232302330023"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata.name` property

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

<a id="canonical-2103112101023012-1200031003310022-1313210003311110-0212312013022120-0112020103331032-1303030201103203-0101320121023023-1230212001022333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list.path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-007.md#canonical-2033123012202321-1211101311022113-1201002111032111-3210022010212021-3321310021201313-3202021033020030-1101021222300022-1110131112213111)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-3321230333023200-0223133222102101-3223130312131313-1330202302222030-3133112013000101-0023222133230111-3321013222321031-3111022101032131)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-0100233201012100-0311000103310021-2033330131120122-2003312130130013-2032130202312233-1102330213113123-1303033030132221-3310030113321013"></a>

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

<a id="canonical-0313112021013123-2103202130000301-2320130310012013-3100111023331331-1001100230133112-2310221210212103-0110201220232220-2221312033020120"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except.exclude_list.path`

<a id="canonical-0332022302031321-3331210011333212-1020121233323333-3312130321030220-0001331033302132-3102120113130203-0220220330312310-0112010322223123"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.path.path` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1213301220121221-2122331011322023-1121121211103323-2231333203222000-1131213003302202-3321321302322112-3201133210121322-2222323203023332"></a>

<a id="canonical-3113120110100302-1333131133233323-2222331213220322-3002232132320200-3323111221212321-0113021113331232-2033221203202103-0201111103303301"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.path.prefix` property

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

<a id="canonical-0313222311303212-0012203130203320-2002303322120201-3020311232210212-2002332310003002-2300021330002110-3012120331010111-0001223211132211"></a>

<a id="canonical-0030312113320013-2230020331303100-3322203132121233-1303211112302310-2330322202213321-3103022213101010-2103330120022001-0000230122003021"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.path.regex` property

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

<a id="canonical-3101332010211112-3031030103123313-2120311030130113-1132001221131033-1103310102132033-0213103111312323-3003202121131230-3323212032330321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- bot_defense.policy.js_insertion_rules

<a id="canonical-0320333133323313-1023013221203320-0313103000310223-2121313132322122-0010020310331030-0312323112120012-3130213111333230-0210321201301302"></a>

Type: `"object"`. single nested block, Optional.

This defines custom JavaScript insertion rules for Bot Defense Policy.

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

<a id="canonical-2303130232203300-3211131321112223-3102211101210022-3011210302101101-2221132130121021-2223202032013303-2132100031100113-2120213111120100"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules`

- [exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232312133003302-3231123221033012-2220322232002010-0110200313231100-2101100000322313-3201013130133031-2333102320303010-2120301002011010): complete subsection reference.

- [rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-2022211220121310-0223330231203210-2031011233012210-3333320322110323-3000221301202000-2210123032303212-1311133321303303-3023111103101203): complete subsection reference.

<a id="canonical-0232312133003302-3231123221033012-2220322232002010-0110200313231100-2101100000322313-3201013130133031-2333102320303010-2120301002011010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-3101332010211112-3031030103123313-2120311030130113-1132001221131033-1103310102132033-0213103111312323-3003202121131230-3323212032330321)
- bot_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-0231212202303220-2233322122333201-3013301230111302-3322300322211032-1100230211212032-2202032320102332-3223130232132131-3102110113330022"></a>

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

<a id="canonical-2321101232133001-1000200013222022-0031230103133112-2013202322112121-3300032103301231-0321121220223123-0130330311020333-0232213310213000"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list`

- [any_domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-0033313121101023-3303110100203102-2121123012032103-0011120230133100-2130030310230013-2021030030102031-1113110112012130-3121123230212220): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-0123033321321232-1002212212023110-0200223212221210-3023323301231121-2230100133020203-0131223210310001-3223332333232220-1301111013031100): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-1320100323302230-2302223301300030-2030303020102101-1223101232033120-1100002020102011-3103210101012000-1203310310131013-1122123013122033): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-007.md#canonical-3312032321221033-3111032333103203-3133131320203011-3332030223133132-0122212011223132-2213000303000033-0003123021122010-3213303000200030): complete subsection reference.

<a id="canonical-0033313121101023-3303110100203102-2121123012032103-0011120230133100-2130030310230013-2021030030102031-1113110112012130-3121123230212220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-3101332010211112-3031030103123313-2120311030130113-1132001221131033-1103310102132033-0213103111312323-3003202121131230-3323212032330321)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232312133003302-3231123221033012-2220322232002010-0110200313231100-2101100000322313-3201013130133031-2333102320303010-2120301002011010)
- bot_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-0302333320100120-0333020100202302-1301032303300313-3330213302200203-0021003030123133-0311111332231331-2023032232101033-3122131012200232"></a>

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

<a id="canonical-0123033321321232-1002212212023110-0200223212221210-3023323301231121-2230100133020203-0131223210310001-3223332333232220-1301111013031100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-3101332010211112-3031030103123313-2120311030130113-1132001221131033-1103310102132033-0213103111312323-3003202121131230-3323212032330321)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232312133003302-3231123221033012-2220322232002010-0110200313231100-2101100000322313-3201013130133031-2333102320303010-2120301002011010)
- bot_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-1022211111002310-1033021222211030-1201303111133201-2311222113021312-3023000032320112-2200332203112210-1032121003220121-2320023002120322"></a>

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

<a id="canonical-2322031213111231-2020220302111302-1321312321111010-2110302012033332-2000023200322011-0202112211010013-0102001230003320-3013103011323330"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list.domain`

<a id="canonical-3221232013033032-0220311201212323-1011221000331332-0001100012330023-0102212022233000-1112230100000032-1000320312120330-2000131332032113"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.domain.exact_value` property

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

<a id="canonical-1200312123222300-1030332303022131-3102301310200233-2120330323311010-0010101211101001-0201031323330002-3013232200122230-1323231303302312"></a>

<a id="canonical-0211112023203221-0221001030212331-2211301330211131-0233032200130123-2222222201023320-3232103300132232-1301213030123113-1201103003302132"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.domain.regex_value` property

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

<a id="canonical-1213310112030313-3200220311131310-0200232230012001-0220112131122333-1332332023221110-1312331030320133-3033020110122311-2123133322123233"></a>

<a id="canonical-0120300220202022-3320310110102111-2120213011122222-2001032332100200-2103222213023131-2210302111121021-3012330321112333-1110130312303112"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.domain.suffix_value` property

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

<a id="canonical-1320100323302230-2302223301300030-2030303020102101-1223101232033120-1100002020102011-3103210101012000-1203310310131013-1122123013122033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-3101332010211112-3031030103123313-2120311030130113-1132001221131033-1103310102132033-0213103111312323-3003202121131230-3323212032330321)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232312133003302-3231123221033012-2220322232002010-0110200313231100-2101100000322313-3201013130133031-2333102320303010-2120301002011010)
- bot_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-0331130001203010-1012302323022210-2111323213013033-3332201100303100-1321300032013130-1301302033230010-2133000130301220-2333212130010130"></a>

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

<a id="canonical-3100020211313332-0031331120312131-0320103023313200-2232130033102011-0331332122211320-3021203232031013-0100232320223100-2130200203330122"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list.metadata`

<a id="canonical-0101120302123132-1023230003002022-0113103323032003-2311001230122123-0002200232123012-1032031101001303-3203323313201003-3033200033001101"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1221020333013213-0321000231333112-2223110200333301-0133123221232003-0002202001332330-2111320310220110-1220031301200131-1233131202133022"></a>

<a id="canonical-3312112313211221-1110122303311030-3013113300133220-2323122330022022-3012200211332333-0320102323320200-0331231131012010-0230203311121101"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.metadata.name` property

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

<a id="canonical-3312032321221033-3111032333103203-3133131320203011-3332030223133132-0122212011223132-2213000303000033-0003123021122010-3213303000200030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-3101332010211112-3031030103123313-2120311030130113-1132001221131033-1103310102132033-0213103111312323-3003202121131230-3323212032330321)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232312133003302-3231123221033012-2220322232002010-0110200313231100-2101100000322313-3201013130133031-2333102320303010-2120301002011010)
- bot_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-2100032113233133-2222302210331312-0122020230122032-0122322203202313-2023233332202230-2010022121031310-0323131003211101-3130210111030101"></a>

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

<a id="canonical-0103233021313212-0213302102032012-3223001223002002-0321201030302330-0110012303010102-0320131311003322-0230011022232103-3322212012101020"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list.path`

<a id="canonical-1122030033203213-1231133311013123-3023213010210133-0133221133110031-1332202031103110-1232101111130200-3300331003221111-3032223010130301"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.path.path` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1132023130312330-1130320321031100-2021123221211133-1320003020321303-0313200202101022-0213123121200312-3122210031230222-0232130131202111"></a>

<a id="canonical-1312121220020220-3233302002231111-2223103202103331-1123130330000133-1022232032201311-2031211320210110-1233231102111110-3233021302210121"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.path.prefix` property

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

<a id="canonical-0133123021322320-3102012000021323-1302332131322213-0322310310200221-2322302200312301-0302121201110120-3101010333033321-1221232222123213"></a>

<a id="canonical-3200003132202212-3002203320022102-3022232113323222-0322112232030013-0312011230103123-2122121312111212-3233332331233110-1013011203002313"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.path.regex` property

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

<a id="canonical-2022211220121310-0223330231203210-2031011233012210-3333320322110323-3000221301202000-2210123032303212-1311133321303303-3023111103101203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-3101332010211112-3031030103123313-2120311030130113-1132001221131033-1103310102132033-0213103111312323-3003202121131230-3323212032330321)
- bot_defense.policy.js_insertion_rules.rules

<a id="canonical-0222013010023312-3032311320023102-0011210000022233-1002110103133003-0011002122300322-3130302103001223-1202310303331203-3223133211131032"></a>

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

<a id="canonical-3020333111010221-1021220232123212-1321002002212202-0113102320000331-0333033133011220-0230123213331003-3102001213101230-2001201210311332"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules`

- [any_domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102221213221201-2300010323013012-3113033212023203-2213312303112102-3231322212323301-1013030013011032-3002302130302113-0212232110020222): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-0000320233112201-2133103223231313-3231320300211222-2013033221102033-3300222303023020-1310331122331231-0200021232022120-0222320110221300): complete subsection reference.

<a id="canonical-3012322312031032-0211321321133201-2311221101031132-1302203301032202-3201203022103132-1303000223223123-3102303213211301-1131010223222311"></a>

<a id="canonical-2130001002330131-3202213122132301-0210223200312133-1310312212322130-0210031203112212-2303132010311030-3213031022110310-1301020012112031"></a>

#### `bot_defense.policy.js_insertion_rules.rules.javascript_location` property

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

- [metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-0231331130320003-2132321201310320-1103333302213310-1211230312203230-0223100221203112-2230332100233111-1303203123321302-0300032121211121): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-007.md#canonical-1330032131112001-3130321323333133-1000003332003313-2201232223022003-3233300233033120-1210222032101130-1110202113311003-2202023333302002): complete subsection reference.

<a id="canonical-2102221213221201-2300010323013012-3113033212023203-2213312303112102-3231322212323301-1013030013011032-3002302130302113-0212232110020222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-3101332010211112-3031030103123313-2120311030130113-1132001221131033-1103310102132033-0213103111312323-3003202121131230-3323212032330321)
- [bot_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-2022211220121310-0223330231203210-2031011233012210-3333320322110323-3000221301202000-2210123032303212-1311133321303303-3023111103101203)
- bot_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-1313010220000001-3231221300331302-3310220203203330-1032102112320001-0322231302021012-0322031201122310-2022322330100222-1332323203213312"></a>

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

<a id="canonical-0000320233112201-2133103223231313-3231320300211222-2013033221102033-3300222303023020-1310331122331231-0200021232022120-0222320110221300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-3101332010211112-3031030103123313-2120311030130113-1132001221131033-1103310102132033-0213103111312323-3003202121131230-3323212032330321)
- [bot_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-2022211220121310-0223330231203210-2031011233012210-3333320322110323-3000221301202000-2210123032303212-1311133321303303-3023111103101203)
- bot_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-3132121310313113-1202231002030030-3011102201001213-3133010320132220-3213103200002112-3003201031201030-0031323232331233-0010120311202013"></a>

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

<a id="canonical-2012231003002130-1331031012212330-1003202111131302-0113131333331111-3132131033102331-2121221221330112-0300221033211232-2222300013203023"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules.domain`

<a id="canonical-2000112002011013-0000301312211302-3120212202132202-3222321300232033-1330320333323011-3222311300011311-2210010330210031-2021121023213132"></a>

#### `bot_defense.policy.js_insertion_rules.rules.domain.exact_value` property

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

<a id="canonical-2220200123313203-0021321110112102-3001113002022332-1333121022130100-0232123330323231-1222312020313311-1122222222133001-2222203112130203"></a>

<a id="canonical-3122201002022211-2110032200221011-3323213212211131-1103121332233031-1033130213301122-0310301033033310-3230102223203010-3210223101023032"></a>

#### `bot_defense.policy.js_insertion_rules.rules.domain.regex_value` property

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

<a id="canonical-2310013023302133-3031321011131302-2210200121001113-3011133131020021-0312110022303100-2330003232300121-3122321130222322-1331001001011230"></a>

<a id="canonical-3322002313020022-0100132311020023-3201120231313301-1121211010023203-1020233231121033-1011022122210130-0231201012232321-2203210320313302"></a>

#### `bot_defense.policy.js_insertion_rules.rules.domain.suffix_value` property

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

<a id="canonical-0231331130320003-2132321201310320-1103333302213310-1211230312203230-0223100221203112-2230332100233111-1303203123321302-0300032121211121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-3101332010211112-3031030103123313-2120311030130113-1132001221131033-1103310102132033-0213103111312323-3003202121131230-3323212032330321)
- [bot_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-2022211220121310-0223330231203210-2031011233012210-3333320322110323-3000221301202000-2210123032303212-1311133321303303-3023111103101203)
- bot_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-2203312131100022-1233111003222233-3323130300303230-1032303132222012-2023122301133033-2030302033233002-0012311222002220-2332220022303100"></a>

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

<a id="canonical-2333331122112313-2021211102020313-3223231202303101-2330103020332013-1000001203031033-3232230032101002-0331213320112203-1232222022303133"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules.metadata`

<a id="canonical-3333022312202002-1300302032011130-0303330010212313-2232222211322211-2102302133023030-1023111013102033-3333220323213013-3330220002313331"></a>

#### `bot_defense.policy.js_insertion_rules.rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2200113133300223-1222011112232123-1013100330020123-1321311330021310-0313010331011333-1232221032321222-2022101013322301-0021213002203120"></a>

<a id="canonical-0001010212323201-3123101330000130-3111311011323021-1323102010222132-2113113033012000-3131122003111121-0111222321001200-1101333001103311"></a>

#### `bot_defense.policy.js_insertion_rules.rules.metadata.name` property

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

<a id="canonical-1330032131112001-3130321323333133-1000003332003313-2201232223022003-3233300233033120-1210222032101130-1110202113311003-2202023333302002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-3101332010211112-3031030103123313-2120311030130113-1132001221131033-1103310102132033-0213103111312323-3003202121131230-3323212032330321)
- [bot_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-2022211220121310-0223330231203210-2031011233012210-3333320322110323-3000221301202000-2210123032303212-1311133321303303-3023111103101203)
- bot_defense.policy.js_insertion_rules.rules.path

<a id="canonical-2123202310103310-0201302010111103-3120030210332122-0130310200200002-1003233203301133-2232210120110313-3212312002220212-2312122033210202"></a>

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

<a id="canonical-1031222200102201-0200223011201201-1032202312102023-1102221311133223-0200212320202221-3232231100011303-3233002003103130-3032302223110313"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules.path`

<a id="canonical-0220231123223022-2303303310022110-2213003331101322-3210333331121133-0030311221220003-2000301100202122-2113310220330030-1111013230122003"></a>

#### `bot_defense.policy.js_insertion_rules.rules.path.path` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2132030130122023-3001302133013210-3013210201301120-3223031023330131-3023000222103321-1101313003033303-0222311310310131-0231110223332021"></a>

<a id="canonical-2121033310000322-1322222002121212-3002110100032002-1223303300230031-3302012000133003-2021000203103123-0221023031130322-2221312123003320"></a>

#### `bot_defense.policy.js_insertion_rules.rules.path.prefix` property

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

<a id="canonical-1313110200232201-1133311221030321-1200202331122320-0120030122132112-0132031310130123-1211031333200233-3320212120201220-3003310032032123"></a>

<a id="canonical-0232211003120230-0220313203311020-1011131232212020-1101220133231332-1200223303123312-0101320101032303-0003213332031201-2230203213322023"></a>

#### `bot_defense.policy.js_insertion_rules.rules.path.regex` property

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

<a id="canonical-0211110332021202-1223322002233230-1010121033102111-2333021322220011-0303302202101022-2020222312011102-1313020332132132-3230110202201322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- bot_defense.policy.mobile_sdk_config

<a id="canonical-2223002332003002-3212130321102310-3130012231301101-0000003133120301-2121311321300001-1211312002311323-3200223213312303-0011111033110220"></a>

Type: `"object"`. single nested block, Optional.

Mobile SDK Configuration. Mobile SDK configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
mobile_sdk_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013132010122211-3302031203221002-3132100020200020-1331000100231221-0132133021201111-1123210203221310-1111031330321311-1132212120213302"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config`

- [mobile_identifier](resources--cdn_loadbalancer--reference--group-007.md#canonical-3202003010100121-0220322013111031-3221310112201330-1201313332330031-1010323111303122-3113200303311102-0101302332313101-0001011323222120): complete subsection reference.

<a id="canonical-3202003010100121-0220322013111031-3221310112201330-1201313332330031-1010323111303122-3113200303311102-0101302332313101-0001011323222120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.mobile_sdk_config](resources--cdn_loadbalancer--reference--group-007.md#canonical-0211110332021202-1223322002233230-1010121033102111-2333021322220011-0303302202101022-2020222312011102-1313020332132132-3230110202201322)
- bot_defense.policy.mobile_sdk_config.mobile_identifier

<a id="canonical-0211021130232220-0010230120232133-2023333302122020-1133123113231110-3101322232012133-2021302220311010-1222031013213323-2320013231212021"></a>

Type: `"object"`. single nested block, Optional.

Mobile Traffic Identifier. Mobile traffic identifier type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
mobile_identifier {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111100331331301-3200010213301233-3013302103023121-0100103030121033-1230132311021033-3023321001011122-1022332120201133-3033103113100111"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config.mobile_identifier`

- [headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-2003200003011133-0220231003311233-3023122233331231-0100213130111020-3212222133202001-1113311001010103-0333001301300212-1023032002022323): complete subsection reference.

<a id="canonical-2003200003011133-0220231003311233-3023122233331231-0100213130111020-3212222133202001-1113311001010103-0333001301300212-1023032002022323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.mobile_sdk_config](resources--cdn_loadbalancer--reference--group-007.md#canonical-0211110332021202-1223322002233230-1010121033102111-2333021322220011-0303302202101022-2020222312011102-1313020332132132-3230110202201322)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--cdn_loadbalancer--reference--group-007.md#canonical-3202003010100121-0220322013111031-3221310112201330-1201313332330031-1010323111303122-3113200303311102-0101302332313101-0001011323222120)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-1122123331000233-0203110113001321-1112111313330220-3210300233200303-0323310012003231-2322321231203232-2000302032030123-2220032023102002"></a>

Type: `"object"`. list nested block, Optional.

Headers that can be used to identify mobile traffic.

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
    "minItems": 0,
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
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2003213302212103-1233000113210103-0122132202013100-1221000022110021-3310111132120331-2112200120230100-0210203102231012-0023110030300021"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers`

- [check_not_present](resources--cdn_loadbalancer--reference--group-007.md#canonical-2030211311121132-0011030311211213-3220302323232103-3121030121233021-0222321201330021-3320122020201230-3330202211302103-0021233131313011): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102023123233312-1101231210000233-3002323213010311-0332303030210000-3203013131101330-2001001112230112-1221030332023232-1231030213323202): complete subsection reference.

- [item](resources--cdn_loadbalancer--reference--group-007.md#canonical-2200223032130210-3332110303301023-1221313320032303-1202122201030330-3021331333302301-2223202201211312-0312221201110101-1310032121322210): complete subsection reference.

<a id="canonical-1020231000021232-2211023332223100-0322002010311323-0132221121201010-2313201300000132-0010310310121130-2102133002312301-1332202232201302"></a>

<a id="canonical-3120303323122331-3331201123233233-2232232010031030-2221311330113220-0301231321020000-2323313022203320-1023120202122003-1231212132303111"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.name` property

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

<a id="canonical-2030211311121132-0011030311211213-3220302323232103-3121030121233021-0222321201330021-3320122020201230-3330202211302103-0021233131313011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.mobile_sdk_config](resources--cdn_loadbalancer--reference--group-007.md#canonical-0211110332021202-1223322002233230-1010121033102111-2333021322220011-0303302202101022-2020222312011102-1313020332132132-3230110202201322)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--cdn_loadbalancer--reference--group-007.md#canonical-3202003010100121-0220322013111031-3221310112201330-1201313332330031-1010323111303122-3113200303311102-0101302332313101-0001011323222120)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-2003200003011133-0220231003311233-3023122233331231-0100213130111020-3212222133202001-1113311001010103-0333001301300212-1023032002022323)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present

<a id="canonical-0021101112200113-1110230121011000-1200302103223322-1311131301013200-0322303113110002-3012130233222213-0321001131230313-3321003010332330"></a>

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

<a id="canonical-2102023123233312-1101231210000233-3002323213010311-0332303030210000-3203013131101330-2001001112230112-1221030332023232-1231030213323202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.mobile_sdk_config](resources--cdn_loadbalancer--reference--group-007.md#canonical-0211110332021202-1223322002233230-1010121033102111-2333021322220011-0303302202101022-2020222312011102-1313020332132132-3230110202201322)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--cdn_loadbalancer--reference--group-007.md#canonical-3202003010100121-0220322013111031-3221310112201330-1201313332330031-1010323111303122-3113200303311102-0101302332313101-0001011323222120)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-2003200003011133-0220231003311233-3023122233331231-0100213130111020-3212222133202001-1113311001010103-0333001301300212-1023032002022323)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present

<a id="canonical-1322132232301010-0313101211101230-2332321321313210-2122202210201331-2302221310202031-1130010222333121-0322301212010302-1203113110113300"></a>

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

<a id="canonical-2200223032130210-3332110303301023-1221313320032303-1202122201030330-3021331333302301-2223202201211312-0312221201110101-1310032121322210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.mobile_sdk_config](resources--cdn_loadbalancer--reference--group-007.md#canonical-0211110332021202-1223322002233230-1010121033102111-2333021322220011-0303302202101022-2020222312011102-1313020332132132-3230110202201322)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--cdn_loadbalancer--reference--group-007.md#canonical-3202003010100121-0220322013111031-3221310112201330-1201313332330031-1010323111303122-3113200303311102-0101302332313101-0001011323222120)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-2003200003011133-0220231003311233-3023122233331231-0100213130111020-3212222133202001-1113311001010103-0333001301300212-1023032002022323)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item

<a id="canonical-0310200312003130-0210321220203033-3123011000203311-3200223132303130-0210112100300002-0130120203220203-0020300333021322-0320332102222322"></a>

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

<a id="canonical-1311012320300320-1313302310312112-3201221130212302-0111033221121010-3030020320031132-1010012130320201-1133121021022013-1101333123120131"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item`

<a id="canonical-3111203333232312-1312233002002212-1301130211113323-2112302211111320-0212310131212230-3122123212001100-2012222000021133-0220233110202102"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item.exact_values` property

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

<a id="canonical-3321011010201310-0203131222022302-2300221233110323-3111122031001202-1331211021033203-2020233302213103-0221121012103231-1313032233221023"></a>

<a id="canonical-3002102122202322-3222033032333201-0212003013101022-1002012113221123-3333031122233320-0233020021323011-0023223121033200-3020300331213333"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item.regex_values` property

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

<a id="canonical-1211121131323223-1303002210223213-2233301030130132-0111112233102300-0030011011113323-0121210000200301-3223321211321133-0133022302013031"></a>

<a id="canonical-1212030122010202-0013213120213311-3030123133013112-0110030322121202-3111312333313311-1223022120000201-1220202202233121-1303032102300312"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item.transformers` property

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

<a id="canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- bot_defense.policy.protected_app_endpoints

<a id="canonical-3001121211222212-1001222233320001-2330232310311331-0220001022031321-0322232001021113-3211121330132000-1331123221122202-2211021213322020"></a>

Type: `"object"`. list nested block, Optional.

List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32
endpoints per LB' after 4 LBs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("http_methods"),
  validators.ConflictingListObjectAttributes("allow_good_bots",
    "mitigate_good_bots"),
  validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("flow_label",
    "undefined_flow_label"),
  validators.ConflictingListObjectAttributes("mobile",
    "web"),
  validators.ConflictingListObjectAttributes("mobile",
    "web_mobile"),
  validators.ConflictingListObjectAttributes("web",
    "web_mobile")}
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
protected_app_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-1112000102132001-0332220312301323-0213213320211303-3100311221003233-1312121022301232-0201333133222103-3003130120321211-2130221302002121"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints`

- [allow_good_bots](resources--cdn_loadbalancer--reference--group-007.md#canonical-2013331331011110-3220031023121111-2323120010012003-0231332110312013-2013203013202221-1001223312131301-1122011221210220-1321203332203231): complete subsection reference.

- [any_domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-3223301120223113-3030021211211233-1301032101031003-3211212032322101-3113030311112211-3212222012303011-0122002312222113-2010211220332233): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-3010333302130310-1023111103122003-3132303130221230-0120233200120332-1002012322123102-0321022223331301-2312122111133121-2020300233222221): complete subsection reference.

- [flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130): complete subsection reference.

- [headers](resources--cdn_loadbalancer--reference--group-008.md#canonical-1100110031020322-1021330213010003-0313300320311220-0101003231320232-2021332111310001-1330021302111232-3020223001131222-0310111230322302): complete subsection reference.

<a id="canonical-3001311303032002-1211113023122210-2332223202300313-3203000203312021-0230303103203131-2022200330012223-1012031221322033-0301002203123131"></a>

<a id="canonical-2030123123131230-1321313202001032-1233333333100301-3230112112000232-0232322302000011-1212311013033232-2221211022300323-3300100033023223"></a>

#### `bot_defense.policy.protected_app_endpoints.http_methods` property

Type: `["list", "string"]`. Optional.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 5),
}
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](resources--cdn_loadbalancer--reference--group-008.md#canonical-3310130131321203-3331122133211030-1121020112102101-0213323211311213-0102323212323210-3022231033021330-1000312320331032-2330022020312100): complete subsection reference.

- [mitigate_good_bots](resources--cdn_loadbalancer--reference--group-008.md#canonical-0311023021301223-0213131330202333-3031023111101230-0230200120220300-0032012201030313-3202221121302331-3031112300021232-2220312200203022): complete subsection reference.

- [mitigation](resources--cdn_loadbalancer--reference--group-008.md#canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101): complete subsection reference.

- [mobile](resources--cdn_loadbalancer--reference--group-008.md#canonical-2320231113201121-2020300211032022-1102102103223120-0111130001113120-3301232221202322-3003200222210111-3331112021232103-2000212121221231): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-008.md#canonical-2223313223123202-2233212130010311-3312012021323023-3002302323032331-3131221133330101-2202011230010131-3013113021220322-1223300321232230): complete subsection reference.

<a id="canonical-3012133313202021-3323113311213120-1212231113022201-2102121031101231-1223232022322103-1301022330103120-1103301300013330-3021311202022220"></a>

<a id="canonical-2203123311023011-3330101220213131-2030110200103131-0300310131131100-0131221221220221-3121030001320303-1112113000331022-1210203211212010"></a>

#### `bot_defense.policy.protected_app_endpoints.protocol` property

Type: `"string"`. Optional.

\[Enum: BOTH|HTTP|HTTPS\] SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for
HTTPS:// or HTTP://. - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only.
Possible values are \`BOTH\`, \`HTTP\`, \`HTTPS\`. Defaults to \`BOTH\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["BOTH","HTTP","HTTPS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("BOTH",
    "HTTP",
    "HTTPS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BOTH",
  "enum": [
    "BOTH",
    "HTTP",
    "HTTPS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [query_params](resources--cdn_loadbalancer--reference--group-008.md#canonical-2222021120300201-2020223023102321-3221032331210030-3302030312102232-3312302330311213-2130033013233102-3100313111202202-0211201310111123): complete subsection reference.

- [undefined_flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-1103122121102201-2310232002301010-2300221230313032-0023230013011103-2102023331202303-1133033313011223-2120000031313003-1320100210123030): complete subsection reference.

- [web](resources--cdn_loadbalancer--reference--group-008.md#canonical-0113200330202102-0131023200210002-0120011001302123-3003313020231311-0232312022111032-0210130032310120-0320301113013123-0130001112002010): complete subsection reference.

- [web_mobile](resources--cdn_loadbalancer--reference--group-008.md#canonical-1033303310333201-2322202210330231-1303323202223200-0121131311331201-2331312332001013-3100331211333112-3213330112110230-3230321011112100): complete subsection reference.

<a id="canonical-2013331331011110-3220031023121111-2323120010012003-0231332110312013-2013203013202221-1001223312131301-1122011221210220-1321203332203231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.allow_good_bots` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.allow_good_bots

<a id="canonical-2212132022003112-1001221320220223-0120302330120122-2130232133112102-2232223323333031-1033312102121000-2331221132002313-3221102120132200"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for allow good bots.

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
allow_good_bots = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223301120223113-3030021211211233-1301032101031003-3211212032322101-3113030311112211-3212222012303011-0122002312222113-2010211220332233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.any_domain

<a id="canonical-2030123033322101-0231233120033132-3232233322210003-0321230202100211-3230001120000023-1220031302121100-2310133221012321-1211233132323321"></a>

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

<a id="canonical-3010333302130310-1023111103122003-3132303130221230-0120233200120332-1002012322123102-0321022223331301-2312122111133121-2020300233222221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.domain

<a id="canonical-3213332101322212-3312131321230321-0213012013131112-1013103130012102-3010010320333021-0101123201211011-3222201310131220-0323110121323133"></a>

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

<a id="canonical-3020232120313103-0022332313001321-3012130133212111-1323021313332302-2210222213302133-0231312103002221-2003121301311230-2201323311130123"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.domain`

<a id="canonical-3233121012123223-1211120233230212-1001313112123013-3233133213001211-3320013130321310-3203121210331021-1123031233130231-1310121231122210"></a>

#### `bot_defense.policy.protected_app_endpoints.domain.exact_value` property

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

<a id="canonical-3112100332001232-2232003322323011-1233103101100211-2131213230302322-1110020020021133-3111210120221133-3211331131131321-2200232233300101"></a>

<a id="canonical-3021012132103021-1013220203233013-0010200311220223-0120303132301133-0203221202110122-0002103013023220-2101133303323221-3010220203102033"></a>

#### `bot_defense.policy.protected_app_endpoints.domain.regex_value` property

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

<a id="canonical-1113210213111333-0100103100333231-1110121023100323-1221011213212333-2111132103220133-2010233323130002-0202022020303100-3102311022122233"></a>

<a id="canonical-2001330200123333-3101222030033332-2002320010020220-3300323332321112-0331231030310121-2231312333032132-2102321220003320-0200310110030132"></a>

#### `bot_defense.policy.protected_app_endpoints.domain.suffix_value` property

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

<a id="canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.flow_label

<a id="canonical-1011312213232312-3032223231220203-3103001202211310-0121011222230301-2301213231303313-0103210211222322-3021300031203132-0332010231330211"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("account_management",
    "authentication"),
  validators.ConflictingObjectAttributes("account_management",
    "financial_services"),
  validators.ConflictingObjectAttributes("account_management",
    "flight"),
  validators.ConflictingObjectAttributes("account_management",
    "profile_management"),
  validators.ConflictingObjectAttributes("account_management",
    "search"),
  validators.ConflictingObjectAttributes("account_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("authentication",
    "financial_services"),
  validators.ConflictingObjectAttributes("authentication",
    "flight"),
  validators.ConflictingObjectAttributes("authentication",
    "profile_management"),
  validators.ConflictingObjectAttributes("authentication",
    "search"),
  validators.ConflictingObjectAttributes("authentication",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("financial_services",
    "flight"),
  validators.ConflictingObjectAttributes("financial_services",
    "profile_management"),
  validators.ConflictingObjectAttributes("financial_services",
    "search"),
  validators.ConflictingObjectAttributes("financial_services",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("flight",
    "profile_management"),
  validators.ConflictingObjectAttributes("flight",
    "search"),
  validators.ConflictingObjectAttributes("flight",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("profile_management",
    "search"),
  validators.ConflictingObjectAttributes("profile_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("search",
    "shopping_gift_cards")}
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
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

Terraform syntax:

```terraform
flow_label {
  # Configure direct properties listed below.
}
```

<a id="canonical-1100100201210000-2331210332000112-0001233211303011-2200221102301033-3013203202112011-0101101322211332-2031331300302000-1102200303001002"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label`

- [account_management](resources--cdn_loadbalancer--reference--group-007.md#canonical-1301130230221023-0033031011202011-1201121010220133-1111211320102213-1011300322200313-1313011010133122-1001222102120030-2331221311233020): complete subsection reference.

- [authentication](resources--cdn_loadbalancer--reference--group-007.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002): complete subsection reference.

- [financial_services](resources--cdn_loadbalancer--reference--group-007.md#canonical-1321213233222323-2320001331202231-1020211320330323-3221203233311102-2220301031330220-2110321111000220-2013121022213021-2032120302302332): complete subsection reference.

- [flight](resources--cdn_loadbalancer--reference--group-007.md#canonical-0332030030201123-3112022110233212-2303113312031010-2131212031312021-1221203201003101-0002221213221131-1011101101203320-0321231311301230): complete subsection reference.

- [profile_management](resources--cdn_loadbalancer--reference--group-007.md#canonical-2112013023232200-2231110000021121-0231333312321300-1333300003302102-1121033120030020-1133301102321011-3013312100220111-3200310111230032): complete subsection reference.

- [search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013): complete subsection reference.

- [shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102): complete subsection reference.

<a id="canonical-1301130230221023-0033031011202011-1201121010220133-1111211320102213-1011300322200313-1313011010133122-1001222102120030-2331221311233020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.account_management` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management

<a id="canonical-2002100110000300-0333320310121211-1311122301202220-2101311102220102-2113203231332301-2223121322121310-3200031000331310-0112110223132010"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Account Management Category.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "password_reset")}
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
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

Terraform syntax:

```terraform
account_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300111300031102-0331221330022130-1130131023222330-3330111211100132-1121023202023002-2120102133303131-0012130101131311-3100011100031013"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.account_management`

- [create](resources--cdn_loadbalancer--reference--group-007.md#canonical-1332230231313033-0102312302132011-2003030131132321-1121112300332222-3223220031331123-1223323110303332-0300222032310332-3223312120112131): complete subsection reference.

- [password_reset](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102320210330032-0123310121132330-1211333003200101-2003330000101001-1233011310322213-1223031111332323-1301321332202312-3130221202230212): complete subsection reference.

<a id="canonical-1332230231313033-0102312302132011-2003030131132321-1121112300332222-3223220031331123-1223323110303332-0300222032310332-3223312120112131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.account_management.create` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--cdn_loadbalancer--reference--group-007.md#canonical-1301130230221023-0033031011202011-1201121010220133-1111211320102213-1011300322200313-1313011010133122-1001222102120030-2331221311233020)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.create

<a id="canonical-3123003210021212-1222202112013203-0102312121103312-0302121103301021-0023000021311333-3000302130322131-2023223233113032-1322211103222003"></a>

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
create = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102320210330032-0123310121132330-1211333003200101-2003330000101001-1233011310322213-1223031111332323-1301321332202312-3130221202230212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--cdn_loadbalancer--reference--group-007.md#canonical-1301130230221023-0033031011202011-1201121010220133-1111211320102213-1011300322200313-1313011010133122-1001222102120030-2331221311233020)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset

<a id="canonical-3323230032302103-2311120233232030-2123301022132120-3213030313132113-3001122311002012-2203132221212002-3223313011211300-1201121033310222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for password reset.

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
password_reset = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication

<a id="canonical-1020033010322112-2302112101002333-1011123320001330-0311200111000221-3201000133022101-1003033203203320-0211310323323113-1102030200231221"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Authentication Category.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("login",
    "login_mfa"),
  validators.ConflictingObjectAttributes("login",
    "login_partner"),
  validators.ConflictingObjectAttributes("login",
    "logout"),
  validators.ConflictingObjectAttributes("login",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_mfa",
    "login_partner"),
  validators.ConflictingObjectAttributes("login_mfa",
    "logout"),
  validators.ConflictingObjectAttributes("login_mfa",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_partner",
    "logout"),
  validators.ConflictingObjectAttributes("login_partner",
    "token_refresh"),
  validators.ConflictingObjectAttributes("logout",
    "token_refresh")}
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
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

Terraform syntax:

```terraform
authentication {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203232310021331-3013122102300021-1221130232100110-0031110031013131-0232223200312113-0102022030012220-1230233210113103-2120221231323003"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication`

- [login](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012): complete subsection reference.

- [login_mfa](resources--cdn_loadbalancer--reference--group-007.md#canonical-0032323121300012-1122323330311013-3213111031231312-0300210213201032-0133201211210123-0102013130222231-1311213023012122-3101231001132133): complete subsection reference.

- [login_partner](resources--cdn_loadbalancer--reference--group-007.md#canonical-3123033300213313-1220322323020321-0230030023032231-3003332231333212-3033013013011100-0301100001200122-1123130332202002-3010010210232021): complete subsection reference.

- [logout](resources--cdn_loadbalancer--reference--group-007.md#canonical-1111203002230011-0011321031000231-1012130131030312-2210213101320113-0301211130000233-1332230113333313-2022002032132310-3023132020113233): complete subsection reference.

- [token_refresh](resources--cdn_loadbalancer--reference--group-007.md#canonical-0133212203133310-1201032011111112-3320330112303200-1103301310102232-0003031301102113-0233132230001030-0131230002001102-3002331120301112): complete subsection reference.

<a id="canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-007.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

<a id="canonical-1221223320103133-0300303230133101-0021221133331021-1213320233233311-3020211101010002-0332301131113020-0233111300001321-0030231202201221"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_transaction_result",
    "transaction_result")}
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
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

Terraform syntax:

```terraform
login {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011021330222201-2123013002210202-1211303310011333-1100122002302230-0001201030121120-1132010133330120-2333331112223201-2113002320212013"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login`

- [disable_transaction_result](resources--cdn_loadbalancer--reference--group-007.md#canonical-0332003023320000-2122131312310000-2322001111211232-1000121330300013-3123013021031330-3001212011201012-0332320110013332-2323123312101200): complete subsection reference.

- [transaction_result](resources--cdn_loadbalancer--reference--group-007.md#canonical-2022133231331001-1332123030223032-1003322321000121-0103101300111001-1132210031232123-2223311332001221-3222101132230023-0002132232220222): complete subsection reference.

<a id="canonical-0332003023320000-2122131312310000-2322001111211232-1000121330300013-3123013021031330-3001212011201012-0332320110013332-2323123312101200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-007.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-3231231120023111-1311332203130311-2013311030333223-3023020330231101-0031200110021220-3031000202010222-2110020133130332-2331033013330100"></a>

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
disable_transaction_result = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022133231331001-1332123030223032-1003322321000121-0103101300111001-1132210031232123-2223311332001221-3222101132230023-0002132232220222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-007.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-0030011013121023-2331323033110033-3230101202021003-3301212132121331-3232032302022332-3113020113203113-0002222113112002-1112222001121223"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
transaction_result {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131001002300303-0100230032021113-2000033132302330-3011023002123212-0212330021110232-1012231013000230-3032220322002313-0213202132321132"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result`

- [failure_conditions](resources--cdn_loadbalancer--reference--group-007.md#canonical-0310201310333111-3101220000322112-1322133033020100-1200202121200110-1223001131233111-1033333303301331-3333220322130001-0310301010021211): complete subsection reference.

- [success_conditions](resources--cdn_loadbalancer--reference--group-007.md#canonical-0310301023302210-0212002010011003-3011221213211023-2321220120023121-0311320333123230-1232003213303203-0030333130300303-2211012331333222): complete subsection reference.

<a id="canonical-0310201310333111-3101220000322112-1322133033020100-1200202121200110-1223001131233111-1033333303301331-3333220322130001-0310301010021211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-007.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--cdn_loadbalancer--reference--group-007.md#canonical-2022133231331001-1332123030223032-1003322321000121-0103101300111001-1132210031232123-2223311332001221-3222101132230023-0002132232220222)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-1222213221333321-0302000210223332-0312233013323101-2302230021020110-0123130120231221-3221211232101122-3202332321031212-1113112222211013"></a>

Type: `"object"`. list nested block, Optional.

Failure Conditions. Failure Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
failure_conditions {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031212102133312-3013210331330113-2021203233233000-2221300101230112-0012331302210003-1023221031001103-3120220120231313-0221320310202002"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions`

<a id="canonical-3322012200000212-2220202232221102-1310332300223223-1030013230112011-2033032000000102-3310000023003202-0302210332332201-2003301103023213"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name` property

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3300310212301113-2123222230002001-2013310220220232-2200103302333223-0333222320032122-0001220100333021-1211211021203212-1313220232211010"></a>

<a id="canonical-3022103121313321-2101033220320013-1012003202001230-0133010231231120-1000220033303201-0001112110123203-0210001122120010-0310022303111013"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values` property

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
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1200100232100011-0220113010312213-0032102301203223-2201020210221000-2133030211123001-1022332321120321-0122210023333100-3321121303133031"></a>

<a id="canonical-1233311231123201-3130301020023001-3002221000033003-1321003103200131-1310213331030033-1111312223301032-2121332211000103-2032011033213313"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status` property

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

Additional upstream details:

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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["Accepted","AlreadyReported","BadGateway","BadRequest","Conflict","Continue","Created","EmptyStatusCode","ExpectationFailed","FailedDependency","Forbidden","Found","GatewayTimeout","Gone","HTTPVersionNotSupported","IMUsed","InsufficientStorage","InternalServerError","LengthRequired","Locked","LoopDetected","MethodNotAllowed","MisdirectedRequest","MovedPermanently","MultiStatus","MultipleChoices","NetworkAuthenticationRequired","NoContent","NonAuthoritativeInformation","NotAcceptable","NotExtended","NotFound","NotImplemented","NotModified","OK","PartialContent","PayloadTooLarge","PaymentRequired","PermanentRedirect","PreconditionFailed","PreconditionRequired","ProxyAuthenticationRequired","RangeNotSatisfiable","RequestHeaderFieldsTooLarge","RequestTimeout","ResetContent","SeeOther","ServiceUnavailable","TemporaryRedirect","TooManyRequests","URITooLong","Unauthorized","UnprocessableEntity","UnsupportedMediaType","UpgradeRequired","UseProxy","VariantAlsoNegotiates"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-0310301023302210-0212002010011003-3011221213211023-2321220120023121-0311320333123230-1232003213303203-0030333130300303-2211012331333222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-007.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--cdn_loadbalancer--reference--group-007.md#canonical-2022133231331001-1332123030223032-1003322321000121-0103101300111001-1132210031232123-2223311332001221-3222101132230023-0002132232220222)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-0233033321122202-2021113100200120-1202203333300233-3232323032331102-3321211130100201-1321321013123202-1321022333101231-0233220022123300"></a>

Type: `"object"`. list nested block, Optional.

Success Conditions. Success Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
success_conditions {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310001311011220-0312230213010222-0230133323213101-1012113322301310-0013032111200122-0011023121233320-0003111100320312-1020102010301232"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions`

<a id="canonical-3013211113112313-1011133023220230-0300331002122231-1221130312111102-2002003303323032-3200313100003233-2200033203321131-3222202232101030"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name` property

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2132212323211022-1302300321220221-0323130300333112-3111100301320320-3222231100121332-0110120133231130-0223232320313210-1033001221133111"></a>

<a id="canonical-1321021103030311-2020320223301103-3332223322201103-2132031133012121-1210211301230101-1203200103012123-0103131031130011-3333101013320123"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values` property

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
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1310332231223333-1312122031330312-0202302103303212-1030213032110313-1300103333211333-2131210032113001-1112301201300012-0322022133013122"></a>

<a id="canonical-3100221021331233-2230032233121033-3230321321211211-0002313030120100-1003121330031313-0212312332131212-1021210332013132-3220221220033210"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status` property

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

Additional upstream details:

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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["Accepted","AlreadyReported","BadGateway","BadRequest","Conflict","Continue","Created","EmptyStatusCode","ExpectationFailed","FailedDependency","Forbidden","Found","GatewayTimeout","Gone","HTTPVersionNotSupported","IMUsed","InsufficientStorage","InternalServerError","LengthRequired","Locked","LoopDetected","MethodNotAllowed","MisdirectedRequest","MovedPermanently","MultiStatus","MultipleChoices","NetworkAuthenticationRequired","NoContent","NonAuthoritativeInformation","NotAcceptable","NotExtended","NotFound","NotImplemented","NotModified","OK","PartialContent","PayloadTooLarge","PaymentRequired","PermanentRedirect","PreconditionFailed","PreconditionRequired","ProxyAuthenticationRequired","RangeNotSatisfiable","RequestHeaderFieldsTooLarge","RequestTimeout","ResetContent","SeeOther","ServiceUnavailable","TemporaryRedirect","TooManyRequests","URITooLong","Unauthorized","UnprocessableEntity","UnsupportedMediaType","UpgradeRequired","UseProxy","VariantAlsoNegotiates"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-0032323121300012-1122323330311013-3213111031231312-0300210213201032-0133201211210123-0102013130222231-1311213023012122-3101231001132133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-007.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa

<a id="canonical-1211202323231233-0221301312110323-0001012320320103-0222102121312103-0120310303003220-1213213032331132-3323013233010230-3022001113012000"></a>

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
login_mfa = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123033300213313-1220322323020321-0230030023032231-3003332231333212-3033013013011100-0301100001200122-1123130332202002-3010010210232021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-007.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner

<a id="canonical-3130232100323301-3201303210311230-1230120013203102-0231113320232011-1111131303133022-2013301310200333-1333120011232202-0112022030103233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for login partner.

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
login_partner = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111203002230011-0011321031000231-1012130131030312-2210213101320113-0301211130000233-1332230113333313-2022002032132310-3023132020113233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-007.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout

<a id="canonical-0331200230220202-2302122223031221-2103202101213203-1033213000221231-1322210012202220-3022113013031011-2332312030202332-0123021130222021"></a>

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
logout = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133212203133310-1201032011111112-3320330112303200-1103301310102232-0003031301102113-0233132230001030-0131230002001102-3002331120301112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-007.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh

<a id="canonical-0320233020133300-1122220301130032-1021320030033031-0203002013103013-1023322020231012-0211211131132123-3110312131013003-3030212031120110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for token refresh.

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
token_refresh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321213233222323-2320001331202231-1020211320330323-3221203233311102-2220301031330220-2110321111000220-2013121022213021-2032120302302332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.financial_services` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services

<a id="canonical-3213130002330221-0313333132331133-3312303323110112-2121111332212102-2112102131231231-0133212310330013-0322113023302321-0332202302311021"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Financial Services Category.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("apply",
    "money_transfer")}
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
  "x-ves-oneof-field-label_choice": "[\"apply\",\"money_transfer\"]"
}
```

Terraform syntax:

```terraform
financial_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221021202231012-3103211302222011-1001002331222301-0001302113213013-1130031221022121-2303211032012323-0212313132110312-3320101123231021"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.financial_services`

- [apply](resources--cdn_loadbalancer--reference--group-007.md#canonical-3201110323313302-0330112021210310-0022331313100311-0301223313001032-0103232301222313-2331311332030121-1323210130303203-1001012023123012): complete subsection reference.

- [money_transfer](resources--cdn_loadbalancer--reference--group-007.md#canonical-3112132023123100-0321333011000312-3321121030203121-0312303332013131-2020130103321003-2032313321112213-1012212001023201-3013032010213030): complete subsection reference.

<a id="canonical-3201110323313302-0330112021210310-0022331313100311-0301223313001032-0103232301222313-2331311332030121-1323210130303203-1001012023123012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--cdn_loadbalancer--reference--group-007.md#canonical-1321213233222323-2320001331202231-1020211320330323-3221203233311102-2220301031330220-2110321111000220-2013121022213021-2032120302302332)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply

<a id="canonical-1111132212323020-3313200133331132-0312012312012123-3201122302210123-2021312301121210-0102233121111123-3012123020110202-0010212133001011"></a>

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
apply = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112132023123100-0321333011000312-3321121030203121-0312303332013131-2020130103321003-2032313321112213-1012212001023201-3013032010213030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--cdn_loadbalancer--reference--group-007.md#canonical-1321213233222323-2320001331202231-1020211320330323-3221203233311102-2220301031330220-2110321111000220-2013121022213021-2032120302302332)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer

<a id="canonical-0313213020223323-2133202013012220-0201311200000033-3121212010002312-3021102313313011-1101001113312300-1110232003221301-0000311120212220"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for money transfer.

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
money_transfer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332030030201123-3112022110233212-2303113312031010-2131212031312021-1221203201003101-0002221213221131-1011101101203320-0321231311301230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.flight` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- bot_defense.policy.protected_app_endpoints.flow_label.flight

<a id="canonical-2030223102300103-2011312222001213-2123321100131211-0123022033123320-0232013223221202-3310321202212113-0203123300212231-2032011332000021"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"checkin\"]"
}
```

Terraform syntax:

```terraform
flight {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333100120121220-3223033131301232-1330031211022303-3323000330033133-3102121133201100-1200033201223023-2123021332120130-0212012133033103"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.flight`

- [checkin](resources--cdn_loadbalancer--reference--group-007.md#canonical-0001323123302232-2023131101313303-0000202332322123-0201322031333121-3322122102203111-2020021010111313-0110013000200132-1123300203203232): complete subsection reference.

<a id="canonical-0001323123302232-2023131101313303-0000202332322123-0201322031333121-3322122102203111-2020021010111313-0110013000200132-1123300203203232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](resources--cdn_loadbalancer--reference--group-007.md#canonical-0332030030201123-3112022110233212-2303113312031010-2131212031312021-1221203201003101-0002221213221131-1011101101203320-0321231311301230)
- bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin

<a id="canonical-1313321031200131-1321111222230111-1223203333321332-1302332223320030-1221232030031220-2213000321012302-3131210022000313-0030333311131332"></a>

Type: `"object"`. single nested block, Optional.

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
checkin {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112013023232200-2231110000021121-0231333312321300-1333300003302102-1121033120030020-1133301102321011-3013312100220111-3200310111230032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management

<a id="canonical-3213021230003323-3303210330033302-0030221121310001-2312313010320333-3301111022202120-0012112201302023-0223332301213322-0102301303133111"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Profile Management Category.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "update"),
  validators.ConflictingObjectAttributes("create",
    "view"),
  validators.ConflictingObjectAttributes("update",
    "view")}
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
  "x-ves-oneof-field-label_choice": "[\"create\",\"update\",\"view\"]"
}
```

Terraform syntax:

```terraform
profile_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201333103312301-3303120300301101-0021320303130112-0200111222001230-0121130330200202-0221221130110013-2102131233002331-3211213311311301"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.profile_management`

- [create](resources--cdn_loadbalancer--reference--group-007.md#canonical-3330202013011331-2020012321322023-1321223331121013-3300023101203020-3030233303131102-2133001001110230-3311231311200013-1203203022133111): complete subsection reference.

- [update](resources--cdn_loadbalancer--reference--group-007.md#canonical-0033223221211332-2310331312211003-1323221010220133-0103002201031121-1323201113321331-1232331201203220-2230132210031120-3203221022010210): complete subsection reference.

- [view](resources--cdn_loadbalancer--reference--group-008.md#canonical-0321130032220302-2322003033030333-2332311132320213-0310332231001123-2321200010302230-1032011103233021-2333102302033221-1330031000110310): complete subsection reference.

<a id="canonical-3330202013011331-2020012321322023-1321223331121013-3300023101203020-3030233303131102-2133001001110230-3311231311200013-1203203022133111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--cdn_loadbalancer--reference--group-007.md#canonical-2112013023232200-2231110000021121-0231333312321300-1333300003302102-1121033120030020-1133301102321011-3013312100220111-3200310111230032)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create

<a id="canonical-1131320302002201-0133231020013032-1300311020201302-1022213302021313-2012202210311331-2201223223222102-3100121213121132-2221013312001000"></a>

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
create = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033223221211332-2310331312211003-1323221010220133-0103002201031121-1323201113321331-1232331201203220-2230132210031120-3203221022010210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--cdn_loadbalancer--reference--group-007.md#canonical-2112013023232200-2231110000021121-0231333312321300-1333300003302102-1121033120030020-1133301102321011-3013312100220111-3200310111230032)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update

<a id="canonical-0212313111122013-0322310022000100-3320102201322120-2222031013213011-2010220111010320-3321211230233113-2330313013112332-2222313003320330"></a>

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
update = {}
```

This is an empty object or choice marker. It has no direct properties.
