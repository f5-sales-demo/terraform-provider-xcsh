---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- graphql_rules

<a id="canonical-3033113222332331-3132301221000332-2212230233231332-2330333210220301-0322131001130303-1131320103233222-1330301211123330-2232113032230103"></a>

Type: `"object"`. list nested block, Optional.

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
graphql_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222132010310333-3320233301012221-2020133023031332-0102123222032003-2220313210230110-2102203012301223-3033102130122113-0311302333213312"></a>

### Direct properties for `graphql_rules`

- [any_domain](resources--cdn_loadbalancer--reference--group-011.md#canonical-0310312333200300-1001133112201301-2231012213003033-1211231123311203-0233000022332111-0012232012120121-1100033002132011-3333220110033032): complete subsection reference.

<a id="canonical-2112013200300110-3201131023132300-0102210030000203-0033322220111131-1213310331123121-0000301132321131-1232033122011130-3020003032010123"></a>

<a id="canonical-0032322322232311-1121101200302211-3120110320323313-3233031121122133-1102213020323321-1322333312032123-0021201210002113-3102230313032030"></a>

#### `graphql_rules.exact_path` property

Type: `"string"`. Optional.

Specifies the exact path to GraphQL endpoint. Defaults to \`/graphql\`.

Additional upstream details:

Default value is /GraphQL.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1100232233221102-3302033233011201-2230220112220103-2301130031231210-3322031300000320-0300022123221231-0333320021011211-3210112013322302"></a>

<a id="canonical-0102021211022121-2100310030312022-2302231322120111-1220112231223012-2111110103102221-1232212332121222-2301313032223300-1102221313203320"></a>

#### `graphql_rules.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [graphql_settings](resources--cdn_loadbalancer--reference--group-011.md#canonical-0303202121321220-1213302012121122-3010013331001221-0330202211132123-3330213012212101-3031001203002213-1113001320220013-0230201221030331): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-011.md#canonical-2111103032113221-2300330331032303-3021202023221012-1300020310021031-1001113031002312-0122031211203333-2111112020000120-2203030003322211): complete subsection reference.

- [method_get](resources--cdn_loadbalancer--reference--group-011.md#canonical-3112311302031310-2013010311020001-1012213233002331-3203312232102320-3332222023031121-2113023320012123-3003232203122311-3020001112232321): complete subsection reference.

- [method_post](resources--cdn_loadbalancer--reference--group-011.md#canonical-3101201020133222-2123332031201300-1233023232013013-3010301220131123-0000011132001021-1031001333131330-3303102132212030-1201311020130133): complete subsection reference.

<a id="canonical-3332221223013123-2033322221131331-3312213031321211-1013303313313001-2210013302100121-3132133122013122-2103110032030032-0321001312300003"></a>

<a id="canonical-0333313022001313-2023033012100312-3220102223121123-2001030323213122-0210330102231133-2021012012121320-2121210332123200-3330222312300211"></a>

#### `graphql_rules.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0310312333200300-1001133112201301-2231012213003033-1211231123311203-0233000022332111-0012232012120121-1100033002132011-3333220110033032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-011.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223)
- graphql_rules.any_domain

<a id="canonical-3001123003101213-2300323003102000-3130023130112210-2021233022011000-0110300003013222-3130322203032302-0231113220032230-2021130123001122"></a>

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

<a id="canonical-0303202121321220-1213302012121122-3010013331001221-0330202211132123-3330213012212101-3031001203002213-1113001320220013-0230201221030331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-011.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223)
- graphql_rules.graphql_settings

<a id="canonical-0011113113220331-3333032022312102-1203133231233131-0132130130020232-2033011100101010-1313123013001231-3113103022230233-3130311100311331"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for GraphQL settings.

Additional upstream details:

GraphQL configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allow_introspection_queries_choice": "[\"disable_introspection\",\"enable_introspection\"]"
}
```

Terraform syntax:

```terraform
graphql_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202123010132002-3230320101122131-2021220211031013-3032201201231013-0231110210130131-0313222033300302-2102122032320221-1313331220210000"></a>

### Direct properties for `graphql_rules.graphql_settings`

- [disable_introspection](resources--cdn_loadbalancer--reference--group-011.md#canonical-3011202003003032-0001310103133331-0111311023220230-3202133313001031-2011303203032233-1303320103110121-1232112230112301-2301230212301233): complete subsection reference.

- [enable_introspection](resources--cdn_loadbalancer--reference--group-011.md#canonical-3323021120022333-1030001100012121-1030112033113022-0232222010022320-2122131332132002-3333121333003323-1202022302200200-1122220120032310): complete subsection reference.

<a id="canonical-3211220113030133-1200213123221232-0301220313213310-1310132222230210-0301101031200303-0102321203230220-0111320320121122-3312303312121320"></a>

<a id="canonical-3013101313013111-3332131033012111-3202322200212003-2211203230331033-3312131123113002-0122302031023113-1332330211111111-2003323021002323"></a>

#### `graphql_rules.graphql_settings.max_batched_queries` property

Type: `"number"`. Optional.

Specify maximum number of queries in a single batched request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-3213210311113222-2022312021003221-1300232213113101-2121303333001020-0320233232310022-0312101122203023-3131001222133123-1220211033100111"></a>

<a id="canonical-2002020312110111-0233123220103301-2022313022010013-1300101330033221-2022300223203200-0203302231103213-2103012222031110-2132123300003202"></a>

#### `graphql_rules.graphql_settings.max_depth` property

Type: `"number"`. Optional.

Specify maximum depth for the GraphQL query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-0130230321231113-0331011323202233-2300320012123112-3101221221232012-0322003031303111-0330322010211212-0231010312200011-3122323313231011"></a>

<a id="canonical-1113310200103310-1202032301311101-0010032301022310-3020010322033132-0113121120011213-0003031203312223-1100101022123021-1123002211310211"></a>

#### `graphql_rules.graphql_settings.max_total_length` property

Type: `"number"`. Optional.

Specify maximum length in bytes for the GraphQL query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16386,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "16386"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  }
}
```

<a id="canonical-3011202003003032-0001310103133331-0111311023220230-3202133313001031-2011303203032233-1303320103110121-1232112230112301-2301230212301233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings.disable_introspection` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-011.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223)
- [graphql_rules.graphql_settings](resources--cdn_loadbalancer--reference--group-011.md#canonical-0303202121321220-1213302012121122-3010013331001221-0330202211132123-3330213012212101-3031001203002213-1113001320220013-0230201221030331)
- graphql_rules.graphql_settings.disable_introspection

<a id="canonical-3101003223213131-0302023333100020-3330000313333233-1101213203002303-3202311323123010-3213133122001112-0332000000230023-2230222303323323"></a>

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
disable_introspection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323021120022333-1030001100012121-1030112033113022-0232222010022320-2122131332132002-3333121333003323-1202022302200200-1122220120032310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings.enable_introspection` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-011.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223)
- [graphql_rules.graphql_settings](resources--cdn_loadbalancer--reference--group-011.md#canonical-0303202121321220-1213302012121122-3010013331001221-0330202211132123-3330213012212101-3031001203002213-1113001320220013-0230201221030331)
- graphql_rules.graphql_settings.enable_introspection

<a id="canonical-0112123030000133-3321231320123131-2303313110122003-3212213110203222-0101330120102132-0323102103232000-1113023320101333-1320031313001232"></a>

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
enable_introspection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111103032113221-2300330331032303-3021202023221012-1300020310021031-1001113031002312-0122031211203333-2111112020000120-2203030003322211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-011.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223)
- graphql_rules.metadata

<a id="canonical-2330312010203330-1012110013030332-1210331323020011-2231112203202213-0313123323013212-0130110311103332-2323332300321223-2022120113003213"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
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

<a id="canonical-0003123030230220-3012131020211133-0021000130002000-3300233302222322-3103221233131303-0030330023233123-2313322011011221-0103330023133010"></a>

### Direct properties for `graphql_rules.metadata`

<a id="canonical-2100133331032012-1032110132320011-3010120332330000-1121010332300111-0000103310321313-3111031131201030-1133213013323323-1012303231321321"></a>

#### `graphql_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-2302112022103202-0000133103102303-3133221132130110-0100331212332000-3231332122020321-1322131111211312-1032133022010021-1102101120111303"></a>

<a id="canonical-1223222310133223-0121121213331201-0133332230330200-3100011002213010-0003110122013021-2221030133132132-1222111031322032-2002122132322133"></a>

#### `graphql_rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3112311302031310-2013010311020001-1012213233002331-3203312232102320-3332222023031121-2113023320012123-3003232203122311-3020001112232321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.method_get` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-011.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223)
- graphql_rules.method_get

<a id="canonical-3222233221101112-0202132120001201-1201222112323102-1202103122220113-3331222200313322-3003110310113022-2122010001233211-0321112111030303"></a>

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
method_get = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101201020133222-2123332031201300-1233023232013013-3010301220131123-0000011132001021-1031001333131330-3303102132212030-1201311020130133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.method_post` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-011.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223)
- graphql_rules.method_post

<a id="canonical-0013230202323121-1013332121001311-1321301023201110-1233301102312201-0202210313231333-2020101200110101-3203000113012012-3323102110321022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for method post.

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
method_post = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320012221332001-3023010200311133-0012202111211311-2312132120132210-3032223212022123-1212221020223222-2203232202220232-2203232310011031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- http

<a id="canonical-1000121033311203-2210131322032220-3212323011131021-0001030112113101-0100303223210201-2212000012230011-3130012011101013-1133213221020323"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: http, https, https\_auto\_cert; Default: https\_auto\_cert\] HTTP Choice. Choice for
selecting HTTP proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

OneOf alternatives in this subsection:

- [http](resources--cdn_loadbalancer--reference--group-011.md#canonical-1000121033311203-2210131322032220-3212323011131021-0001030112113101-0100303223210201-2212000012230011-3130012011101013-1133213221020323)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-0030210123302230-2132331033300110-1110303122211311-0023313300030000-1100221102102222-3331022302312113-0311203102112002-0232201201033012)
- [https_auto_cert](resources--cdn_loadbalancer--reference--group-012.md#canonical-1011001132000320-2032211231022312-1103022230201012-2202332311121101-0230203310001320-2210023110030310-2011102022223002-0133123021300313)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110203032210122-1230103020310123-3322223110021200-3033232232110221-0312003020212010-1110130300322120-0230020222021330-2230331323000312"></a>

### Direct properties for `http`

<a id="canonical-0303010032021311-0113113123111312-1302003130011322-1230013312022021-1112001233233123-2210102012323213-3121011010311200-2212220131131323"></a>

#### `http.dns_volterra_managed` property

Type: `"bool"`. Optional.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1311213121120331-0202110020301001-3231322230232210-0000100021323121-0031311021322332-0122010200112113-2311321303230103-2332030002221203"></a>

<a id="canonical-0200301131222331-2203030001230130-1113031030201120-2222112301002211-2333220033012001-1200211013133213-1233003011330030-0330223130010011"></a>

#### `http.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2010120000113220-1211130202013023-0201023113222223-3200110231021211-0210331202213222-1001011201001233-3202201031300000-0002023202212321"></a>

<a id="canonical-2211313223213231-0111232221031322-3003203212233110-0013203023002021-1002023222112203-3211200222332323-3031202100223022-3220023123120110"></a>

#### `http.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- https

<a id="canonical-0030210123302230-2132331033300110-1110303122211311-0023313300030000-1100221102102222-3331022302312113-0311203102112002-0232201201033012"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting CDN Distribution with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131222220220130-0003323101031201-1010302311020032-3320001010113301-1302311313223323-1313301103000221-0000101230220210-1303232120303122"></a>

### Direct properties for `https`

<a id="canonical-0133031133212130-1321032323001130-2233001313313310-2021301213211221-3010013102130311-0232123212130322-1200123303330023-2121223101130300"></a>

#### `https.add_hsts` property

Type: `"bool"`. Optional.

Add HTTP Strict-Transport-Security response header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2331220202202202-3200110211012320-0122210103123013-1320121002030232-2130332230313130-0210020122313110-1323332113220011-1212031013210021"></a>

<a id="canonical-2003023323121320-3020311021020013-2110221002131212-3012232133013110-3023222210311113-3312022202202311-1212222211011002-0020100333202110"></a>

#### `https.http_redirect` property

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132): complete subsection reference.

<a id="canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- https.tls_cert_options

<a id="canonical-2002111312101220-1033231311000323-0002301330212330-1030113130123120-3220222031022023-1230233233211020-1022030031103100-0203200201210312"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert options.

Additional upstream details:

TLS Certificate OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_inline_params\"]"
}
```

Terraform syntax:

```terraform
tls_cert_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0002311210321321-1333231312332311-3210001222030203-0021213111231210-1232130013113323-3220013230223330-2222323310231330-3221311111121113"></a>

### Direct properties for `https.tls_cert_options`

- [tls_cert_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101): complete subsection reference.

- [tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100): complete subsection reference.

<a id="canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- https.tls_cert_options.tls_cert_params

<a id="canonical-3313331331032322-0103010213220221-3330121302020311-2110202321230310-3330110111331131-1030310222121120-2322013030000303-1111032312133023"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Additional upstream details:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003113022313323-2210031312003010-3303000021003131-0013332230110220-0312113203133120-2322303101313301-0313222233132310-2111123302300333"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params`

- [certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-1020021333030211-3011200000212023-3210031201010321-0210213122231010-2022223303302133-3132210020202113-3223321030111210-3022032111122322): complete subsection reference.

- [no_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-3021200213311112-0032113010312301-0030303030133012-1233012221331023-2133330311113211-2022022021310210-1112303023200211-2203030130332233): complete subsection reference.

- [tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-1031331122100021-3131111330111203-2333333132031233-2130331021100120-1232131312311320-0120031130123220-1013301010302110-3333102021313120): complete subsection reference.

- [use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-2013031123320130-3023100003003001-0130010131231023-3100210112220023-0231230020210122-1000131313031002-0111001332022033-0102232331322320): complete subsection reference.

<a id="canonical-1020021333030211-3011200000212023-3210031201010321-0210213122231010-2022223303302133-3132210020202113-3223321030111210-3022032111122322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- https.tls_cert_options.tls_cert_params.certificates

<a id="canonical-2000100222122210-2202232030131201-1000203201201132-3003012121312033-1222211223203113-0031332231032230-3113323121212131-0001030300122310"></a>

Type: `"object"`. list nested block, Optional.

Select one or more certificates with any domain names.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102032231110311-1212313012113130-2333121102013233-1300123330233021-0332002123102020-2321210132322002-3312011112222132-2123302023202203"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.certificates`

<a id="canonical-1012312010133120-2211131110100222-0013010023101020-3301211020023311-3033231002200100-0131123321312132-3123303221130223-3112012312023121"></a>

#### `https.tls_cert_options.tls_cert_params.certificates.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0300303022131203-1030100000300200-3220023200103311-0011123101303320-0012123200030330-0012302102020020-2310023012213131-1321023300123332"></a>

<a id="canonical-1223201000333233-1011023330012300-2302110212100302-3221101323121110-1003330203210332-3303012120213221-0132001221223230-2302110012232020"></a>

#### `https.tls_cert_options.tls_cert_params.certificates.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1111033010223022-3030333331231113-1013202231202022-2320201220323310-0220223121210233-2002302101313301-1011030111130210-0011132121022213"></a>

<a id="canonical-1020131101221220-1113012032022301-2122221020120001-3221323121113102-0330013110213111-1311212102103330-1303220313100323-2321012201033103"></a>

#### `https.tls_cert_options.tls_cert_params.certificates.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3021200213311112-0032113010312301-0030303030133012-1233012221331023-2133330311113211-2022022021310210-1112303023200211-2203030130332233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- https.tls_cert_options.tls_cert_params.no_mtls

<a id="canonical-3020222203232310-2012203131200230-1101313203221113-2012023111133011-1301113110223212-1223211001030323-1333310023003202-0331222302311012"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031331122100021-3131111330111203-2333333132031233-2130331021100120-1232131312311320-0120031130123220-1013301010302110-3333102021313120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- https.tls_cert_options.tls_cert_params.tls_config

<a id="canonical-0100013130032010-0121010333130320-2133133111121102-3222011102303210-0000002031010032-0222303110302223-0121303010131223-3212231332331300"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202020102313333-1312111002302332-0133203323100123-0032100020201012-3000333130311113-0121311102000031-3331223332102300-2010322312201201"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.tls_config`

- [custom_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-0323322323121333-1303311220212232-0223201133021022-1020011012312020-3312210303211202-3032220122202331-2312021023002111-1113232330001033): complete subsection reference.

- [default_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-0120001133011022-3023333001330231-2022102023023132-0211132313213033-3031200311200322-0030122132223101-3202221322232221-3200120021201321): complete subsection reference.

- [low_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-3201130202012103-1300201123010231-0032201013002213-0220332212233302-3312112202322023-1002010112222303-1333323330101230-0111100031002332): complete subsection reference.

- [medium_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-2330012003201103-1033011111321030-3002023213331032-0223023310011123-1000213320323321-3203323130313132-2200211113332020-3200331311112110): complete subsection reference.

<a id="canonical-0323322323121333-1303311220212232-0223201133021022-1020011012312020-3312210303211202-3032220122202331-2312021023002111-1113232330001033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-1031331122100021-3131111330111203-2333333132031233-2130331021100120-1232131312311320-0120031130123220-1013301010302110-3333102021313120)
- https.tls_cert_options.tls_cert_params.tls_config.custom_security

<a id="canonical-3112223131321101-1032002121311233-2310202330300110-3133303211203100-0211330033321103-1321333200012320-3132320321232001-1003223301120222"></a>

Type: `"object"`. single nested block, Optional.

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303120023011001-2333320110111233-1212010313213023-2232023311121030-1032302100302130-3312231112110200-0230000221010031-3113102312013112"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.tls_config.custom_security`

<a id="canonical-0100332021321112-2330303211101120-1232302123021023-2212301301002122-2212111331331233-2312201300111333-0021233210200030-2002000131203221"></a>

#### `https.tls_cert_options.tls_cert_params.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1230300131320233-2311222221022223-2021320003010033-1331333003212101-0223232113002231-3111312332313212-0222322101111112-2211020210330132"></a>

<a id="canonical-3011000300130100-2212232001133310-0331203321221221-3213303033120011-0333330012202201-0020012000100230-2001202332210130-0003130132100103"></a>

#### `https.tls_cert_options.tls_cert_params.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0102032011131330-1013200231030311-2333302221222331-1000311020020320-0020032031131333-1210211322012020-1111322020320120-0300021110230201"></a>

<a id="canonical-1123113302302222-3103010111212002-0112112211221022-1320302213300311-0222031012112202-3112333320123233-0200033101323001-2230311011031030"></a>

#### `https.tls_cert_options.tls_cert_params.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0120001133011022-3023333001330231-2022102023023132-0211132313213033-3031200311200322-0030122132223101-3202221322232221-3200120021201321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-1031331122100021-3131111330111203-2333333132031233-2130331021100120-1232131312311320-0120031130123220-1013301010302110-3333102021313120)
- https.tls_cert_options.tls_cert_params.tls_config.default_security

<a id="canonical-0312012022311011-3132222122030120-2230200111130023-2112023102112003-1022220122320020-2011003210122210-3010211202221120-2112020110101211"></a>

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
default_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201130202012103-1300201123010231-0032201013002213-0220332212233302-3312112202322023-1002010112222303-1333323330101230-0111100031002332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-1031331122100021-3131111330111203-2333333132031233-2130331021100120-1232131312311320-0120031130123220-1013301010302110-3333102021313120)
- https.tls_cert_options.tls_cert_params.tls_config.low_security

<a id="canonical-2310000222120033-1200223031001233-1002012110030330-0133322200210133-3100321203031001-0301211122221103-3023230201332032-2111221122333123"></a>

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
low_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330012003201103-1033011111321030-3002023213331032-0223023310011123-1000213320323321-3203323130313132-2200211113332020-3200331311112110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-1031331122100021-3131111330111203-2333333132031233-2130331021100120-1232131312311320-0120031130123220-1013301010302110-3333102021313120)
- https.tls_cert_options.tls_cert_params.tls_config.medium_security

<a id="canonical-0300203323211332-1010313031002121-2010300013311332-3223213210122113-3023002303133301-0103232221322021-3032212333323013-2330310230110331"></a>

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
medium_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2013031123320130-3023100003003001-0130010131231023-3100210112220023-0231230020210122-1000131313031002-0111001332022033-0102232331322320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- https.tls_cert_options.tls_cert_params.use_mtls

<a id="canonical-2000303230303111-1302000200312220-2312113313133123-1333312033102003-2231233212020120-3323121022330132-0103311311231301-0312112213020231"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313202121112202-0133003233312110-2313132022303332-3033103201012320-1100102103331021-0130202132003313-1311211000032210-2322020112203021"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.use_mtls`

<a id="canonical-0220310022022213-1122000011103111-1102133032331110-2332312222301133-2002233031101020-0223333010200113-0013010323330210-1113320111032002"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.client_certificate_optional` property

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](resources--cdn_loadbalancer--reference--group-011.md#canonical-3113103200232003-3121110302230001-2102113220103120-1020203112331320-3020132213010203-3233121012222010-3322022131120001-3031022320111211): complete subsection reference.

- [no_crl](resources--cdn_loadbalancer--reference--group-011.md#canonical-1033003031200310-0023022222211010-3031100020132002-2312030031011333-0001031000121330-1122013022210313-0213223131323323-0133332202031302): complete subsection reference.

- [trusted_ca](resources--cdn_loadbalancer--reference--group-011.md#canonical-3310013132310332-0232123320300021-1111201303230022-2332312122112131-0320331122212112-0011312120010003-2101313322231003-2003300212233230): complete subsection reference.

<a id="canonical-0233211230033310-1123000010030200-1103321231112333-2201313333030331-2131320211313301-1331202210133203-3021033310323123-3200132210020223"></a>

<a id="canonical-2112102232301330-3210010201123310-1121211003203311-3111221123011230-1333203220100311-3200012301133133-3002122330021313-1231312232233313"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](resources--cdn_loadbalancer--reference--group-011.md#canonical-1131202213031230-2101230201231111-2113013302223330-2212230020323323-3102210320010133-1202112233133333-2120000323131211-0323203313013133): complete subsection reference.

- [xfcc_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-1032303320022232-0031123332300121-1102131212233111-3203313302330222-2001011112113330-2013002230000101-0122301021023122-1013231302121001): complete subsection reference.

<a id="canonical-3113103200232003-3121110302230001-2102113220103120-1020203112331320-3020132213010203-3233121012222010-3322022131120001-3031022320111211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-2013031123320130-3023100003003001-0130010131231023-3100210112220023-0231230020210122-1000131313031002-0111001332022033-0102232331322320)
- https.tls_cert_options.tls_cert_params.use_mtls.crl

<a id="canonical-2303310210203330-0030302320230111-1222233213121301-3133103100132302-2032100231323231-0100123322003001-3001103302201031-2022032033311123"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011033103220101-0222201212022012-3123012121110220-1323313002233131-2223032003103320-2023132310103130-0120012213323030-3232223000311023"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.use_mtls.crl`

<a id="canonical-2323011202131321-2011333221200030-0121320312023330-2312102132303203-0333133312133202-3303103231131232-1023232312302021-1011301301300032"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.crl.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2221012001222000-1003320122132131-3330121311121130-1102112300122000-1132102201322223-0303001131232101-1112213213231023-1130003233233010"></a>

<a id="canonical-1033133221121302-0103123113132203-2132312222302302-2001333101002230-1030131220230311-0110331123302022-1122230021001003-2301313302222100"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.crl.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0121102300020011-0021012213002100-0322223213122330-0131221110123110-1311002103212111-0001321031132201-0313302213022232-3310012122303020"></a>

<a id="canonical-3103320021320003-3122002112120301-3310311203133202-3031022212122021-3312223230110211-3022031300013303-2331202032303201-2032202002013313"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.crl.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1033003031200310-0023022222211010-3031100020132002-2312030031011333-0001031000121330-1122013022210313-0213223131323323-0133332202031302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-2013031123320130-3023100003003001-0130010131231023-3100210112220023-0231230020210122-1000131313031002-0111001332022033-0102232331322320)
- https.tls_cert_options.tls_cert_params.use_mtls.no_crl

<a id="canonical-2022203123032322-1111003312021130-2201330022130002-1310301300120030-0322212120330112-0010132002032012-2011230321111103-0210302112131312"></a>

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
no_crl = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310013132310332-0232123320300021-1111201303230022-2332312122112131-0320331122212112-0011312120010003-2101313322231003-2003300212233230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-2013031123320130-3023100003003001-0130010131231023-3100210112220023-0231230020210122-1000131313031002-0111001332022033-0102232331322320)
- https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-0000122110111301-2011230031232033-2210302333002203-1120113010332131-2012233133213131-3023231003233012-0321303301221321-3220212231121203"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322002002202031-1131230301003103-0202103003023023-2010003330111033-0023031321210031-2212232022132133-1331201233102113-2223012030102233"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-2101110321002213-0220100111211322-3012011230020310-0211033110302033-2123303002102000-3023003223112130-3022131130033110-3012330003130121"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0330301102002112-3331221310302003-2003210023333000-3003230313020122-0200232133112003-0031020311300320-2323013220103133-1201102301122232"></a>

<a id="canonical-0101203132012012-3130013132320232-3130211002122032-0021230302333110-3220110013002301-0112000313023121-0121221210110100-0003321002331310"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2011203003302022-3120030301220111-0210202330333323-3101110023110101-3323030213310011-0211202213313101-3101302302000011-3223230300311200"></a>

<a id="canonical-2133022022102313-0031112311201230-0313102222012220-3221301103311221-3031031303023300-3122033031211000-0003320123011012-1021132311223222"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1131202213031230-2101230201231111-2113013302223330-2212230020323323-3102210320010133-1202112233133333-2120000323131211-0323203313013133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-2013031123320130-3023100003003001-0130010131231023-3100210112220023-0231230020210122-1000131313031002-0111001332022033-0102232331322320)
- https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-0123301113211230-1223033101113202-2030032111223033-1032033323331100-2011110211131200-2212232010001110-3332313212122213-0132321213032020"></a>

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
xfcc_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032303320022232-0031123332300121-1102131212233111-3203313302330222-2001011112113330-2013002230000101-0122301021023122-1013231302121001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-2013031123320130-3023100003003001-0130010131231023-3100210112220023-0231230020210122-1000131313031002-0111001332022033-0102232331322320)
- https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-3233001020233312-0023113013121130-3333112030013011-0201112011102220-2230210030320220-2121332111002320-1123132210233010-0322122021301301"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3000323231113023-1102133013130313-2233233023302220-2303120201312232-0023132030132223-2331113023312332-3112330123220331-3101220300020011"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-1323320033211020-1221221132230123-2303031301233021-3211331312033330-2130321100300121-3321201201233233-1301220101001320-1200300200302101"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- https.tls_cert_options.tls_inline_params

<a id="canonical-3031011130003123-0113333013121023-0222101333232332-3213021213320212-1002310200203223-1001202000122210-0101023300000013-2113000210213002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls inline params.

Additional upstream details:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_inline_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130211233320202-1312020213011200-0131313022002023-1100023200202031-1323331313332203-2132213320321020-0120233123222002-0320102032222310"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params`

- [no_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-1230031113013103-0322011010001133-3303112023213133-2023213011230121-1020131032231103-3011332311212201-1323302302210310-2123032230002110): complete subsection reference.

- [tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-1130022333232022-0132213131032021-0100011203130203-0113132332122320-0012102233213010-2202111010210112-1221012130332321-3100331121222122): complete subsection reference.

- [tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-0332223200111113-1202013221133021-2232223301232031-2130111333302310-1133210303301203-3102230200001331-1130001201210021-3230033231321233): complete subsection reference.

- [use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-3212211212021233-3313330312310302-3113222102222002-1333002032022131-2111303321313330-3230131120312200-2300230201110012-2103013032022211): complete subsection reference.

<a id="canonical-1230031113013103-0322011010001133-3303112023213133-2023213011230121-1020131032231103-3011332311212201-1323302302210310-2123032230002110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.no_mtls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- https.tls_cert_options.tls_inline_params.no_mtls

<a id="canonical-2011312132120301-1001311023113001-2222111110022103-0103102222020321-3210221133333133-1223022312110003-3132013013122110-3020223031101023"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130022333232022-0132213131032021-0100011203130203-0113132332122320-0012102233213010-2202111010210112-1221012130332321-3100331121222122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- https.tls_cert_options.tls_inline_params.tls_certificates

<a id="canonical-1233321222200010-0111321210000233-0131030120332102-3032300022020031-0001120120122123-3230022113033321-0023113213112331-0003322013203103"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300133101232112-3302221102202032-0231332020010111-3203202131010203-0020133312131301-1333321012222301-0121002300000010-0310211311021213"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.tls_certificates`

- [blindfold](resources--cdn_loadbalancer--reference--group-011.md#canonical-1233023123201333-2301311131132022-3203033101013120-0212230222010232-2300330222120200-1023211013232212-1103230213332210-1220330002002112): complete subsection reference.

<a id="canonical-3020301203012013-3201021131213021-3010332030033020-0322112311030133-3202102101100223-3030112211130100-3030233021021313-2210320022012023"></a>

<a id="canonical-1111320000020112-3001330103300020-2213231201133002-1121331331103211-3221023000230023-1222330011111032-3101003100111232-0203113032331233"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.certificate_url` property

Type: `"string"`. Optional, Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--cdn_loadbalancer--reference--group-011.md#canonical-3123300311013321-3310101222312223-2012330321301012-2221121101220232-0220020123130113-0213021322303323-1321322022102332-2121133112010300): complete subsection reference.

<a id="canonical-3011212321122103-0131110011131022-3212201012023011-2012230111101303-3211111200312113-1211333331212102-1200101333330121-1202011212210212"></a>

<a id="canonical-3133011313321122-0321320122310323-3321210333113131-1210103313033333-0202131131211013-2032001032031133-3003300331011122-1221123231220330"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--cdn_loadbalancer--reference--group-011.md#canonical-1212012013031110-3132010110102102-1133003322122133-2022030102112201-1011010333113300-1103111211011200-2131302303212021-2122331120313133): complete subsection reference.

- [private_key](resources--cdn_loadbalancer--reference--group-011.md#canonical-3312122210323110-2022033001322111-0210020213112221-3111210310130011-1323323323032313-3320331113331030-1103323323113210-3000312012021303): complete subsection reference.

- [use_system_defaults](resources--cdn_loadbalancer--reference--group-012.md#canonical-0020003122102012-0323213211021111-3103300212000331-2133023113001032-3323011311000002-1300013002231113-3133302101120331-3121200113232223): complete subsection reference.

<a id="canonical-1233023123201333-2301311131132022-3203033101013120-0212230222010232-2300330222120200-1023211013232212-1103230213332210-1220330002002112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-1130022333232022-0132213131032021-0100011203130203-0113132332122320-0012102233213010-2202111010210112-1221012130332321-3100331121222122)
- https.tls_cert_options.tls_inline_params.tls_certificates.blindfold

<a id="canonical-0201210013133033-0303010302231000-0331231033211033-2123033031012222-2000313122212220-2003230312011200-1023303312032213-1231300103113002"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-3010102010233303-3021002133301322-1330122013021310-2233102331310010-3121013120303002-1133332101231023-1003210332010003-3320122212001113"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold`

<a id="canonical-3023132120232130-0302210222200030-2001030321302231-0112211213010033-0012130120102313-3010322032011032-3333321103201013-2222223013131203"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-2022322301200122-0101211231101100-2220000323300123-3231202110313222-0010122222112103-0012101210101102-1112213321100023-1333213321101300"></a>

<a id="canonical-1332000211132012-2200010213221000-0201311311223111-0032233120001113-2132220021102123-2232000031132013-2312222332013311-3010132133220200"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-3311132201333122-2332132313113212-0021103233012331-0313230013303010-3303102010013002-2001320200033022-3131001011201012-3321310010030311"></a>

<a id="canonical-0302030223022121-3010113031300212-1231132013221300-2232112231202213-3030232210002100-1032211311312300-1101332311100032-3332221010101023"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-2322223013123111-2310031120212011-1021210110121330-2233103010322021-3123032011111302-2013302233212130-2210022223130031-2201201013312121"></a>

<a id="canonical-1000212003311113-0322112133300231-1322103303030011-2101012220222332-2000222321212213-0232232021310320-3230101232222200-0302102033103121"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-0031103100131133-2031200323001302-0131121132100333-2121003111202112-1202012233023113-1020311220130030-0330111013131221-1313002103111010"></a>

<a id="canonical-1332112120120021-0010003203031021-2303333230023232-1201303220211203-0231133131313130-2310001331333012-3231212203033121-0321122010030200"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-1020132230002003-3200330202310013-0120133201111032-2302301320200003-1300333333113120-1211003211033111-2312100021310100-3121011021301100"></a>

<a id="canonical-2300303100011300-3033123012112303-2000120100033100-1202031332321022-3011101311312033-2223320113221213-0333122031102123-0120020221033310"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-0222122201310123-1130100313201220-1023033230212201-2131012312210200-3333301111131313-3303111102231300-0302131202103310-0130300320223031"></a>

<a id="canonical-3003013330111203-0301220023000303-0232220123310132-3113011132003202-0033221220021312-1010311022030132-0013032021310212-0011102212302330"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-1033010302122012-1312111222032003-0011231333213232-3311211231103121-0330233103200333-2203013023211333-0030101200003001-3233210210312230"></a>

<a id="canonical-1300112333102110-2121130313201333-3312022200201010-0323013311303013-2132100033100330-2320233212022030-1011002223200131-0233102221221332"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-3220231032013201-2230122001231222-1212332002031130-1102211013033012-3301110132333213-0320230032110002-0033211333313201-1302212100201133"></a>

<a id="canonical-0031131003233311-1201330313202012-1312321102331202-2120322110333301-2220032222223301-0301002230300320-1133220131023133-0022003022310103"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-1202222120030321-0203222233132322-1001001122102120-2223030103130100-0321002123111303-0301223313033211-1321300231102002-0231312210101023"></a>

<a id="canonical-1221013332203022-2203330222200121-3103112103002013-3112011212030032-3320211312211031-0201312103022131-3312022301311010-1223212003011323"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-3331223013223303-1321230220022332-3021330121311310-3031201301212231-0003330011230123-0103200101013121-1221201302033302-3201030101103030"></a>

<a id="canonical-0120133030110022-2301203133330022-2123013300000230-3330011032003321-1312311030022023-3033231032212310-1130132103200200-1102332222032100"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-3332132332333020-1220112232220020-3130102002113033-1101012123211223-1031021321010132-1322033211103121-2233313301001220-1221232330112212"></a>

<a id="canonical-1300230301323121-3013323210133112-0222310121323203-1301201330300032-3021011030130301-3010021012203333-2023213022302321-1300221223011113"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-3100310112232300-3000202312133202-3122320211233203-1031113110301122-1001210200331313-1131212301333011-3130001230211302-0313202022220313"></a>

<a id="canonical-1300111021203231-0222102311123101-3302132021212110-1000313121101202-1303021111321221-1002030321032111-3221100101232003-2023012103220213"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-1021313332102222-3122120021221013-2011300002203010-3113110012231110-1121013303030033-0133310231301200-2022121322103230-3123200302313201"></a>

<a id="canonical-0320032313020320-0203302113230210-3221012023333212-3022030103232131-2302322032222231-3102113220312201-2300110303033200-1202103023320220"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-3000232002212030-3300210000010232-2102333230121131-3032213320101030-0012322332030223-3200222000331121-3003013022101321-3211113032033313"></a>

<a id="canonical-0213013110203300-1230201131202211-2222133222322220-2112001010310232-1012123032203301-2200320213010201-0313313131331013-3322020231210331"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-0112000021013332-1200220321330003-1230022302233030-1223103010123330-1131221012101020-0303310101002030-3131301232011032-0110122032123032"></a>

<a id="canonical-1202332003122232-3131320021202023-1232111201032111-3302132332222303-1123003000323301-1131012110201223-2132231132310323-3220003000210123"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-1302332331320102-3322201102002033-0033133113113331-0203331103332323-2232011102021320-2211312230213332-3010130120201010-3200221011012331"></a>

<a id="canonical-0132310220321011-0001220211112221-2130300120223113-0211133122321321-1323013112232332-1120213021030312-2233121113131321-0023323311022230"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-0320123113023033-3222031311122023-1131201131033231-2102100122021110-3330130231021132-2011112012333112-2010033332210112-0301122202120310"></a>

<a id="canonical-0003300301023320-2312121100220011-3330103323103132-2321130110101001-3332012332330023-1223302331323003-3100013032302112-1101221032020301"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-2010303332120222-1212030001023321-3130102032101120-3112120001001211-3133330331003123-1101100000020230-1102033022221303-1232332132010112"></a>

<a id="canonical-2310330323303130-0312303120002313-0010302130111203-0133031130310300-1002121100312230-2203112033230013-0133231102330112-1012221331131020"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.blindfold.spki_identity` property

Type: `"string"`. Computed.

<a id="canonical-3123300311013321-3310101222312223-2012330321301012-2221121101220232-0220020123130113-0213021322303323-1321322022102332-2121133112010300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-1130022333232022-0132213131032021-0100011203130203-0113132332122320-0012102233213010-2202111010210112-1221012130332321-3100331121222122)
- https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms

<a id="canonical-1001020201300311-3002132131233131-1321101201311022-1202003120113113-1313113302001222-1210200301123031-2123303010212330-1101301300203001"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323001021331200-3203010112021203-2300231130130010-3301103222101333-3100132020010000-3010102013010300-1320100120030010-0133330221312302"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms`

<a id="canonical-3022232321010102-1010000011110200-0111101132033031-2211032123130120-1333220200333202-3103231321011132-0001022011233220-1332230203220332"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1212012013031110-3132010110102102-1133003322122133-2022030102112201-1011010333113300-1103111211011200-2131302303212021-2122331120313133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-1130022333232022-0132213131032021-0100011203130203-0113132332122320-0012102233213010-2202111010210112-1221012130332321-3100331121222122)
- https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-0131211320321120-1312222312020102-0222233002231120-3010000010012031-0333233210011002-1020320032232120-3302011312333003-3021321110030313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312122210323110-2022033001322111-0210020213112221-3111210310130011-1323323323032313-3320331113331030-1103323323113210-3000312012021303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-1130022333232022-0132213131032021-0100011203130203-0113132332122320-0012102233213010-2202111010210112-1221012130332321-3100331121222122)
- https.tls_cert_options.tls_inline_params.tls_certificates.private_key

<a id="canonical-2023302330231031-2203332313001201-3020003130022231-0221023310012220-1331333322223023-1132010332323023-2233313201132113-2113012233203332"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3031301021131203-0310001233030333-1233232033003122-0230220021032213-3102300331000310-0113110120110133-1201022223111123-2033220313031333"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.tls_certificates.private_key`

- [blindfold_secret_info](resources--cdn_loadbalancer--reference--group-011.md#canonical-2103121202132010-3000313031332203-1310303323010033-3030311231323302-3322310023230200-1333201333302012-1101210333303303-1313133211031013): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-2132112002132230-1001001332012103-0232001312331023-3021210022233330-2313331132001312-0311010013230320-1310033023203233-0023310012010200): complete subsection reference.

<a id="canonical-2103121202132010-3000313031332203-1310303323010033-3030311231323302-3322310023230200-1333201333302012-1101210333303303-1313133211031013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-011.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-1130022333232022-0132213131032021-0100011203130203-0113132332122320-0012102233213010-2202111010210112-1221012130332321-3100331121222122)
- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-011.md#canonical-3312122210323110-2022033001322111-0210020213112221-3111210310130011-1323323323032313-3320331113331030-1103323323113210-3000312012021303)
- https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0233132012313311-3333211221030103-2032332130113111-2302002201023321-1311133020010200-1321302211321012-2321013023320313-3212002000210120"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
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
