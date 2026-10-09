---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-3103031133111310-3103122030131131-2102122001020330-0312320333033100-3221221210323002-0203030002021112-0003003123013030-1302221033203112"></a>

## Direct properties for `azure_receiver.filename_options`

<a id="canonical-2221130031202331-3220332121302101-2101010001123320-2231331201310210-2123033003333300-3211222002110222-2210233313303303-3023123223312330"></a>

### `azure_receiver.filename_options.custom_folder` property

Type: `"string"`. Optional.

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match \`/^\[a-z\_\]\[a-z0-9\\\\-\\\\.\_\]\*$/i\`

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
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  }
}
```

- [log_type_folder](resources--global_log_receiver--reference--group-002.md#canonical-2313133202312310-3003023111113031-1311332220021230-3012030101123323-1312022021301302-3010032220132011-2333133101221310-3331211233213210): complete subsection reference.

- [no_folder](resources--global_log_receiver--reference--group-002.md#canonical-2220212033330103-3110031030300003-1113333313102121-2212122232031001-2102133320130321-2333130321333031-2322122120333310-0022002100200233): complete subsection reference.

<a id="canonical-2313133202312310-3003023111113031-1311332220021230-3012030101123323-1312022021301302-3010032220132011-2333133101221310-3331211233213210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure_receiver.filename_options.log_type_folder` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [azure_receiver.filename_options](resources--global_log_receiver--reference--group-001.md#canonical-3110031023120033-1213230103132203-2110113022103201-0132031102313032-0222333230033201-2320000331022303-3233031321002010-2110103103010001)
- azure_receiver.filename_options.log_type_folder

<a id="canonical-3302233113221300-0012021231120122-3130301002220321-0310331301321303-2113232110123023-3313222301031023-3201212111032130-2122033121002020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for log type folder.

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
log_type_folder = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220212033330103-3110031030300003-1113333313102121-2212122232031001-2102133320130321-2333130321333031-2322122120333310-0022002100200233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure_receiver.filename_options.no_folder` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [azure_receiver.filename_options](resources--global_log_receiver--reference--group-001.md#canonical-3110031023120033-1213230103132203-2110113022103201-0132031102313032-0222333230033201-2320000331022303-3233031321002010-2110103103010001)
- azure_receiver.filename_options.no_folder

<a id="canonical-3001030000033221-2033211000100320-1211012312212131-0023313313201022-0222112323203011-3223331011200320-0100120221321310-1002231001332121"></a>

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
no_folder = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- datadog_receiver

<a id="canonical-0003033200133213-2021130011102033-2321233212202033-0010010332230103-2101001230301213-0313001301111103-3112321322103202-2021022012211332"></a>

Type: `"object"`. single nested block, Optional.

Datadog Configuration. Configuration for Datadog endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("endpoint",
    "site"),
  validators.ConflictingObjectAttributes("no_tls",
    "use_tls")}
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
  "x-ves-oneof-field-endpoint_choice": "[\"endpoint\",\"site\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
datadog_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112313220222300-2331013012000112-3330330233021211-3011121320333100-1212130013130111-1033231031223221-3333012113212321-3310223321330003"></a>

### Direct properties for `datadog_receiver`

- [batch](resources--global_log_receiver--reference--group-002.md#canonical-2223332013002130-2132203020333313-0300211311222212-1321133023230031-3331032131001210-2101021213200223-2320110123000111-0020112013112230): complete subsection reference.

- [compression](resources--global_log_receiver--reference--group-002.md#canonical-3030230332320233-2330011003211132-0112312123001121-1301000021010102-0101130200212121-3200321213211330-0133002130323203-3111221213133220): complete subsection reference.

- [datadog_api_key](resources--global_log_receiver--reference--group-002.md#canonical-3003331032300003-3113111102303300-0022111111131031-2220230330110113-1120122230101110-2030202100202320-3232330113023022-0231001111213310): complete subsection reference.

<a id="canonical-2121132022131301-1021220221122333-3102223320122020-2032212120002232-1301010303210303-2100300311231021-3031223321200231-1202100221321011"></a>

<a id="canonical-2332321022030002-3013323113100321-1322010011011330-1110323101033130-1001203300100231-0003300312232210-1303003011230211-2210233220012302"></a>

#### `datadog_receiver.endpoint` property

Type: `"string"`. Optional.

Exclusive with \[site\] Datadog Endpoint,.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.9,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^https?://[^\\s/$.?#].[^\\s]*$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_tls](resources--global_log_receiver--reference--group-002.md#canonical-3013212312131302-2022221332103200-2120320302301030-0202201232233030-0110233213023302-2330201031221231-3030111002013123-1020323020201132): complete subsection reference.

<a id="canonical-3232330011012121-3222103200321203-3211122001332130-0113313111103100-1020321222033022-0312113313210232-3230303121002033-3110311123303200"></a>

<a id="canonical-2132220101122103-3110020130012211-3031332031103312-1031001112320020-1123220222003121-3203312311100312-1220012230130302-0311030301010311"></a>

#### `datadog_receiver.site` property

Type: `"string"`. Optional.

Exclusive with \[endpoint\] Datadog Site,.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  }
}
```

- [use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323): complete subsection reference.

<a id="canonical-2223332013002130-2132203020333313-0300211311222212-1321133023230031-3331032131001210-2101021213200223-2320110123000111-0020112013112230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.batch` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- datadog_receiver.batch

<a id="canonical-1102323220133001-3102020103031023-1210120111332011-2323100010331012-2221123031201102-0030023002100213-0003223000300332-0233103101122020"></a>

Type: `"object"`. single nested block, Optional.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_bytes",
    "max_bytes_disabled"),
  validators.ConflictingObjectAttributes("max_events",
    "max_events_disabled"),
  validators.ConflictingObjectAttributes("timeout_seconds",
    "timeout_seconds_default")}
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
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

Terraform syntax:

```terraform
batch {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223311210113201-2122030021131022-3312013021331023-0101003210123230-3002021213103313-1211122321313231-0131301103300110-2030322321000300"></a>

### Direct properties for `datadog_receiver.batch`

<a id="canonical-3311123322330313-3203103233122202-0111200331123310-0000122202122003-0322221030210001-2122121311220232-0303321330022213-3132020330330201"></a>

#### `datadog_receiver.batch.max_bytes` property

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(4096, 10485760),
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](resources--global_log_receiver--reference--group-002.md#canonical-3021020103021100-2032012202002331-1221212303102301-0001121023002033-2301310330113130-1200303111330100-3201203110102322-3103201222322102): complete subsection reference.

<a id="canonical-2101230131331303-3133020102030322-0112001230122010-3210033211212112-3021002212121000-1321132010010112-2023321231130001-3133111201133001"></a>

<a id="canonical-2130313001201230-2003210221200003-1320300023221013-0013332002000321-3212023301102201-0200103321112122-2120121011303222-2123211021231200"></a>

#### `datadog_receiver.batch.max_events` property

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(32, 2000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](resources--global_log_receiver--reference--group-002.md#canonical-1333032111100131-3002112200203012-0330002201331112-3031332322030302-2102200233113122-1233021032112102-0200201331233201-2103112313320000): complete subsection reference.

<a id="canonical-1220230221130321-1122130233032220-3033322103020133-1001100103230232-2320332113302010-1000010321213133-0111030112303123-1200013203023223"></a>

<a id="canonical-2110330131111001-2321021011123021-1220321312223233-0120100120011312-3211330233103310-0232333013230110-3231322330133031-2203032330010213"></a>

#### `datadog_receiver.batch.timeout_seconds` property

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](resources--global_log_receiver--reference--group-002.md#canonical-1013330000232023-0233320033100123-0003002111023113-1031001102301313-0303021210033211-2130323213012011-1002002203130120-3232202023312311): complete subsection reference.

<a id="canonical-3021020103021100-2032012202002331-1221212303102301-0001121023002033-2301310330113130-1200303111330100-3201203110102322-3103201222322102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.batch.max_bytes_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-2223332013002130-2132203020333313-0300211311222212-1321133023230031-3331032131001210-2101021213200223-2320110123000111-0020112013112230)
- datadog_receiver.batch.max_bytes_disabled

<a id="canonical-0111213032122002-0212231001232223-2311101322323113-0333031230023033-2130300030102110-1213002130212313-1221312130003332-1322223102223202"></a>

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
max_bytes_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333032111100131-3002112200203012-0330002201331112-3031332322030302-2102200233113122-1233021032112102-0200201331233201-2103112313320000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.batch.max_events_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-2223332013002130-2132203020333313-0300211311222212-1321133023230031-3331032131001210-2101021213200223-2320110123000111-0020112013112230)
- datadog_receiver.batch.max_events_disabled

<a id="canonical-2000031012002030-2103300321130003-0320323122013100-3233102330022232-2232203011122111-0003123212201030-0211222001012103-3211312131332030"></a>

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
max_events_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013330000232023-0233320033100123-0003002111023113-1031001102301313-0303021210033211-2130323213012011-1002002203130120-3232202023312311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.batch.timeout_seconds_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-2223332013002130-2132203020333313-0300211311222212-1321133023230031-3331032131001210-2101021213200223-2320110123000111-0020112013112230)
- datadog_receiver.batch.timeout_seconds_default

<a id="canonical-1222112302032112-2323321012111203-0333311321103031-0120123201103310-3113220231311303-2201331013212030-3330131032032113-3221301032331210"></a>

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
timeout_seconds_default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030230332320233-2330011003211132-0112312123001121-1301000021010102-0101130200212121-3200321213211330-0133002130323203-3111221213133220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.compression` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- datadog_receiver.compression

<a id="canonical-0100020323133330-2021232133002231-1033032121301031-2323111023313303-3312120101223112-0101032021001221-3231210211100323-2102303333030003"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Additional upstream details:

Compression Type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("compression_default",
    "compression_gzip"),
  validators.ConflictingObjectAttributes("compression_default",
    "compression_none"),
  validators.ConflictingObjectAttributes("compression_gzip",
    "compression_none")}
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
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

Terraform syntax:

```terraform
compression {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111201301130100-3101313330330231-0311311122330210-1322320131331230-3003000201332310-3310100023132300-2102300221003311-1203120323110210"></a>

### Direct properties for `datadog_receiver.compression`

- [compression_default](resources--global_log_receiver--reference--group-002.md#canonical-2013020312120222-1032312300013001-3303231303303130-1132303033303310-3102212001301320-1130033031111110-3303233202033313-3202100131003100): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-002.md#canonical-0202203113122102-0211002230030313-2310130203222301-1232200301310130-3301030033130330-1230131311020100-3213302031311030-3001300301020233): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-002.md#canonical-0023110002100010-3130000210023222-0203222110322211-1311121303211120-1303301332022231-0303213130323110-1111230100223000-1101031023002321): complete subsection reference.

<a id="canonical-2013020312120222-1032312300013001-3303231303303130-1132303033303310-3102212001301320-1130033031111110-3303233202033313-3202100131003100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.compression.compression_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-3030230332320233-2330011003211132-0112312123001121-1301000021010102-0101130200212121-3200321213211330-0133002130323203-3111221213133220)
- datadog_receiver.compression.compression_default

<a id="canonical-3200131330223011-2220311302123323-1310310200203210-1202121210122333-2202123113301122-0311002200102130-1032323132010021-2222331132203000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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
compression_default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202203113122102-0211002230030313-2310130203222301-1232200301310130-3301030033130330-1230131311020100-3213302031311030-3001300301020233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.compression.compression_gzip` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-3030230332320233-2330011003211132-0112312123001121-1301000021010102-0101130200212121-3200321213211330-0133002130323203-3111221213133220)
- datadog_receiver.compression.compression_gzip

<a id="canonical-1012031012132023-1230031221212030-0022331120022302-2312301031323102-2332121203010200-0233210203320123-1123000230230201-2022023322012323"></a>

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
compression_gzip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023110002100010-3130000210023222-0203222110322211-1311121303211120-1303301332022231-0303213130323110-1111230100223000-1101031023002321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.compression.compression_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-3030230332320233-2330011003211132-0112312123001121-1301000021010102-0101130200212121-3200321213211330-0133002130323203-3111221213133220)
- datadog_receiver.compression.compression_none

<a id="canonical-1101132011322000-2131323330200310-1221031313131223-1023312231010022-0223232111332313-1112112233301010-0122100211220311-0313300220312013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression none.

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
compression_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003331032300003-3113111102303300-0022111111131031-2220230330110113-1120122230101110-2030202100202320-3232330113023022-0231001111213310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.datadog_api_key` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- datadog_receiver.datadog_api_key

<a id="canonical-0311110001003200-1333323330303023-0213211030202113-3313133213111110-2321301113233211-3023202232300333-2210202333032301-0230233131223232"></a>

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
datadog_api_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013231031013011-0133333201112031-3311023312210031-1033112200112121-0123103302220212-0310331012130120-3231031120101122-0232032121003111"></a>

### Direct properties for `datadog_receiver.datadog_api_key`

- [blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-2103202232221111-0102221310212202-0020012111032131-1121101220022213-1101303001122303-2200231302313332-0200303320002322-2203003332312231): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-3212100202100303-0011101223230130-0321000232102223-3011203321032330-2200303231212210-1011322012032201-2033103022013202-0011322132130313): complete subsection reference.

<a id="canonical-2103202232221111-0102221310212202-0020012111032131-1121101220022213-1101303001122303-2200231302313332-0200303320002322-2203003332312231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.datadog_api_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.datadog_api_key](resources--global_log_receiver--reference--group-002.md#canonical-3003331032300003-3113111102303300-0022111111131031-2220230330110113-1120122230101110-2030202100202320-3232330113023022-0231001111213310)
- datadog_receiver.datadog_api_key.blindfold_secret_info

<a id="canonical-3212121320233333-2011211200232111-3201111031232012-0323123300322202-2121300022211201-2131331210310113-3223210022202301-3021102311320203"></a>

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

<a id="canonical-0003332013301200-2131121011121302-2101300310011231-0223021210021031-0022210002323210-0211132010330302-2220230000313031-2122011101002000"></a>

### Direct properties for `datadog_receiver.datadog_api_key.blindfold_secret_info`

<a id="canonical-2110220013220200-2313232110030021-3131110210130320-0333213010102133-0310030213320013-3331323211301003-2312321231123201-3123202131020320"></a>

#### `datadog_receiver.datadog_api_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3020131131212121-3221322033002313-0022311120231303-1223221133331101-2332021301332202-2331210031332311-0320021223320031-0200012020311320"></a>

<a id="canonical-0301022311333131-3300122120213020-1110031210323233-3032002013301032-2212002322320210-1311233223231201-2032122021020121-3021123020123220"></a>

#### `datadog_receiver.datadog_api_key.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2333021100030001-1112322312113220-0222110000313022-3012022220020203-3330120011111310-3333301100123323-2010212203113101-0002120122201013"></a>

<a id="canonical-1311213220330230-1233021201002222-3132222032123022-0223323331112101-3010311233002102-1131013223121301-1212222112210101-3120022231310111"></a>

#### `datadog_receiver.datadog_api_key.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3212100202100303-0011101223230130-0321000232102223-3011203321032330-2200303231212210-1011322012032201-2033103022013202-0011322132130313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.datadog_api_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.datadog_api_key](resources--global_log_receiver--reference--group-002.md#canonical-3003331032300003-3113111102303300-0022111111131031-2220230330110113-1120122230101110-2030202100202320-3232330113023022-0231001111213310)
- datadog_receiver.datadog_api_key.clear_secret_info

<a id="canonical-0212212011332201-2112331202222210-1000231021300131-0303021231130032-1201111020201302-0210013003232011-2121210330133000-2303320213131033"></a>

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

<a id="canonical-2202231122002033-2212300023112211-3231132133001202-1320322311211210-3211033310301002-0022220222113321-2321120001130133-3131030023113213"></a>

### Direct properties for `datadog_receiver.datadog_api_key.clear_secret_info`

<a id="canonical-2212222301022122-1022000120020000-0131221223313001-0133122110201310-0302113233012032-1100322033001113-2102330113232033-3330031010310232"></a>

#### `datadog_receiver.datadog_api_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3210031323132101-3300021330003110-3300330110031113-3320003322211223-3210111230332033-2133211313002023-1031202301103121-0000213111331222"></a>

<a id="canonical-2033100232030103-0111111133220201-1230122002222102-3331211133220132-0300220101220333-3032000010012301-2111020003110032-2032220203202300"></a>

#### `datadog_receiver.datadog_api_key.clear_secret_info.url` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3013212312131302-2022221332103200-2120320302301030-0202201232233030-0110233213023302-2330201031221231-3030111002013123-1020323020201132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.no_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- datadog_receiver.no_tls

<a id="canonical-2323030200323112-3112102121002321-2003201100321203-2101112010003303-2120203103130300-2210110123231110-0211301022233112-1002102330021000"></a>

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
no_tls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- datadog_receiver.use_tls

<a id="canonical-3220330103021032-0210311031103113-3030233033030230-1223010313131033-1210301320222101-0003300030031001-3031012332221313-3202210033130111"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for client connection to the endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_verify_certificate",
    "enable_verify_certificate"),
  validators.ConflictingObjectAttributes("disable_verify_hostname",
    "enable_verify_hostname"),
  validators.ConflictingObjectAttributes("mtls_disabled",
    "mtls_enable"),
  validators.ConflictingObjectAttributes("no_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133000303320013-0232232231333010-3020103011130231-0302130312012213-2112321122021222-2221211201101110-0013323011221331-3102131100310123"></a>

### Direct properties for `datadog_receiver.use_tls`

- [disable_verify_certificate](resources--global_log_receiver--reference--group-002.md#canonical-2010023220221312-3111030320011000-2110200032003013-3202303233322013-1111312211203111-3220301131302110-1322231223311310-0312203311011213): complete subsection reference.

- [disable_verify_hostname](resources--global_log_receiver--reference--group-002.md#canonical-0110020111321033-2210111001012202-2332123203013331-3322332300222211-0122313323220002-3023320112123321-2023233210311302-1113221013133020): complete subsection reference.

- [enable_verify_certificate](resources--global_log_receiver--reference--group-002.md#canonical-0133011203031313-2230213023320331-1122200103123131-0103021310331220-1002103230211233-1103303031200300-0031110222313301-0230123001021121): complete subsection reference.

- [enable_verify_hostname](resources--global_log_receiver--reference--group-002.md#canonical-0031022130230301-1033031021021113-2330122311023233-0032121330133130-3033101001020231-1012103111232122-0232211322111331-3133222022230000): complete subsection reference.

- [mtls_disabled](resources--global_log_receiver--reference--group-002.md#canonical-2021031131131012-0111031111133000-3113311032202112-3331130002222011-1230231101013232-2320133011101132-1003331033201033-1333121201032332): complete subsection reference.

- [mtls_enable](resources--global_log_receiver--reference--group-002.md#canonical-2032221203123021-3330200312030002-1202131321312300-0101030303001021-2010110012310223-3312320223002023-2320200203320313-3321030202011112): complete subsection reference.

- [no_ca](resources--global_log_receiver--reference--group-002.md#canonical-2222023032230211-3130303313312302-2033020112330031-2320101301203120-2023002023112301-2112323023323233-0021312332023201-0300230332330001): complete subsection reference.

<a id="canonical-3000201022320103-1233121233010210-3123222302301310-0000003213030331-2232230210030111-1230110221113100-3201102031202033-2301202322222320"></a>

<a id="canonical-3311312212223332-1311301313022202-0102202123322031-1222132210221202-1023333212302223-0330223301101103-0002320300322210-3021021233220101"></a>

#### `datadog_receiver.use_tls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-2010023220221312-3111030320011000-2110200032003013-3202303233322013-1111312211203111-3220301131302110-1322231223311310-0312203311011213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.disable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- datadog_receiver.use_tls.disable_verify_certificate

<a id="canonical-3113220300212223-2130021103010123-3201010130113032-0232222212220312-3012200023230331-0022302323000320-3023310310022032-0030313333003310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable verify certificate.

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
disable_verify_certificate = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110020111321033-2210111001012202-2332123203013331-3322332300222211-0122313323220002-3023320112123321-2023233210311302-1113221013133020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.disable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- datadog_receiver.use_tls.disable_verify_hostname

<a id="canonical-3231303032021303-2223300101331120-2113322302333323-2222331300323313-1321022203030120-2320203122202133-3003103003003021-0220202102101003"></a>

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
disable_verify_hostname = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133011203031313-2230213023320331-1122200103123131-0103021310331220-1002103230211233-1103303031200300-0031110222313301-0230123001021121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.enable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- datadog_receiver.use_tls.enable_verify_certificate

<a id="canonical-2031312120033231-1233001020332033-2310103133202101-2111102232201003-3222323301012101-0110000213010323-1313022322232203-2303120200321003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable verify certificate.

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
enable_verify_certificate = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031022130230301-1033031021021113-2330122311023233-0032121330133130-3033101001020231-1012103111232122-0232211322111331-3133222022230000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.enable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- datadog_receiver.use_tls.enable_verify_hostname

<a id="canonical-1330321111112203-1023303310212100-0322310022233212-0323002213203312-1303120331131033-1230310331020001-2120132322003300-3322112300031200"></a>

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
enable_verify_hostname = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021031131131012-0111031111133000-3113311032202112-3331130002222011-1230231101013232-2320133011101132-1003331033201033-1333121201032332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.mtls_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- datadog_receiver.use_tls.mtls_disabled

<a id="canonical-2030320302211130-0232233012123113-2302113030332003-2211303123301032-1312130301011011-3211203022310330-2003012220113221-3210122033022133"></a>

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
mtls_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032221203123021-3330200312030002-1202131321312300-0101030303001021-2010110012310223-3312320223002023-2320200203320313-3321030202011112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.mtls_enable` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- datadog_receiver.use_tls.mtls_enable

<a id="canonical-1211123010030223-3110100232033321-2101231123131001-0320231210220130-2112033121130200-2132301303120200-0233121002301310-0302111111101112"></a>

Type: `"object"`. single nested block, Optional.

MTLS Client config allows configuration of mTLS client OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
mtls_enable {
  # Configure direct properties listed below.
}
```

<a id="canonical-2222112333132331-1132200332133231-1331333023130002-2101303100231031-3311213132002213-0110023332300022-1012023001032322-1102320030133220"></a>

### Direct properties for `datadog_receiver.use_tls.mtls_enable`

<a id="canonical-3231113030112213-2330031211212303-2001113222102323-3313020332233332-1000123010220203-0001021303111332-3132030231001010-1331002001023111"></a>

#### `datadog_receiver.use_tls.mtls_enable.certificate` property

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(100, 131072),
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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](resources--global_log_receiver--reference--group-002.md#canonical-3213223010013312-3011133203210000-0121313222022122-2332233312203221-3102301313232011-0202302130201132-2030121132311023-1012320132110113): complete subsection reference.

<a id="canonical-3213223010013312-3011133203210000-0121313222022122-2332233312203221-3102301313232011-0202302130201132-2030121132311023-1012320132110113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.mtls_enable.key_url` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- [datadog_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-002.md#canonical-2032221203123021-3330200312030002-1202131321312300-0101030303001021-2010110012310223-3312320223002023-2320200203320313-3321030202011112)
- datadog_receiver.use_tls.mtls_enable.key_url

<a id="canonical-0003011213120310-3322300103012000-1033230100230302-2312302122110102-3131222202110123-1010200330333233-3331101332331010-3110230120323000"></a>

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
key_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220210313211303-0321231010230332-0322211100022113-2300000212013100-3313221122112032-1201221310022023-3002233022322012-1203002332133210"></a>

### Direct properties for `datadog_receiver.use_tls.mtls_enable.key_url`

- [blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-0030312032110101-1010022312212103-0101002220200130-2202102333022120-1131013320300313-3100101112310230-3130112312012303-2110233111231301): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-1203031212311331-1121212032000320-1120323231121330-2210223203201320-2011011320132003-2222011212030212-0113112011021211-2133312010301221): complete subsection reference.

<a id="canonical-0030312032110101-1010022312212103-0101002220200130-2202102333022120-1131013320300313-3100101112310230-3130112312012303-2110233111231301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- [datadog_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-002.md#canonical-2032221203123021-3330200312030002-1202131321312300-0101030303001021-2010110012310223-3312320223002023-2320200203320313-3321030202011112)
- [datadog_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-002.md#canonical-3213223010013312-3011133203210000-0121313222022122-2332233312203221-3102301313232011-0202302130201132-2030121132311023-1012320132110113)
- datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-0020311202202033-1333123302213210-0123210123123121-2033120213200203-1333203123012332-2013222130010220-0323001111113313-1202030020000303"></a>

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

<a id="canonical-2301002123030221-1210331000002213-2111303112121211-0023312120012331-1001222312123313-0310230022012012-3011033123333333-2223013300202313"></a>

### Direct properties for `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info`

<a id="canonical-0330121030032012-2210321121010120-3331302222132201-2230020212113103-1003101123211031-1030122232330103-3201103230210020-0313100230110110"></a>

#### `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1230332202301332-3311210220113211-1230331110202003-3020120121220313-0103023131113223-1333220130223033-2100221322212320-0223201203111233"></a>

<a id="canonical-3232211233030312-1101213320223103-3112120331033102-0331202230001130-1110122211310222-2320233310120220-1212132033111232-0122232313320022"></a>

#### `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2313310000211230-1320300112032003-1311032323132230-3032202233323101-3032020021311322-0202210213011333-3303110102312313-3301021203001331"></a>

<a id="canonical-3313122332322010-0102020312333012-0133323230133013-3333101031311323-0303322012031001-2111203202332201-3013031313002113-1103100302001202"></a>

#### `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1203031212311331-1121212032000320-1120323231121330-2210223203201320-2011011320132003-2222011212030212-0113112011021211-2133312010301221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- [datadog_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-002.md#canonical-2032221203123021-3330200312030002-1202131321312300-0101030303001021-2010110012310223-3312320223002023-2320200203320313-3321030202011112)
- [datadog_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-002.md#canonical-3213223010013312-3011133203210000-0121313222022122-2332233312203221-3102301313232011-0202302130201132-2030121132311023-1012320132110113)
- datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-3301013130203001-2002233222321222-3121333300113233-0222011310330120-1111230333122223-1221123210222133-2111122313202332-2211130112110021"></a>

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

<a id="canonical-3031313103212031-0313033223231131-3121223001201220-0222223110130311-1013202331333301-1021011003030323-2332000003030221-1031130320210020"></a>

### Direct properties for `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info`

<a id="canonical-3112303302301022-2200121233221131-0102100031113030-1113312322113220-1123300211002332-0003100332213113-1113221103022302-1002000123121331"></a>

#### `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3201230100332312-2330123312313133-3213022213112301-2001331201320320-3123213311213211-1120320331330001-2203110311012032-1213302331230032"></a>

<a id="canonical-1222330221031313-3232213222211200-1132201203010311-3332223022230312-1322130030310031-2123102322203202-1012203013132322-2212222300301302"></a>

#### `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2222023032230211-3130303313312302-2033020112330031-2320101301203120-2023002023112301-2112323023323233-0021312332023201-0300230332330001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.no_ca` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- datadog_receiver.use_tls.no_ca

<a id="canonical-1233332031103301-0122321320010033-2321000001023003-2210203232012321-2331102211302211-1320113213032022-3303311102020312-3300200211013201"></a>

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
no_ca = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211233331121300-1322331131001220-0102011110001201-1030212021001001-2311133020330123-0200201221122221-1011112033100211-0230123330102313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dns_logs` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- dns_logs

<a id="canonical-2112120022003322-1202120213333133-0010302200210023-3232021030212312-0232123211210022-2203312323223132-3100211030221032-3112323313223031"></a>

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
dns_logs = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- gcp_bucket_receiver

<a id="canonical-1221011012130302-2001213102221021-3130230103112313-1211200120113032-0230233232020101-3331301323120221-0220020210111122-1233101312331111"></a>

Type: `"object"`. single nested block, Optional.

GCP Bucket Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("bucket")}
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
gcp_bucket_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332230002320213-0312110213000110-0223301232022303-0302131120223201-0103230120001100-0232112201201010-1313010330131023-2132221223323120"></a>

### Direct properties for `gcp_bucket_receiver`

- [batch](resources--global_log_receiver--reference--group-002.md#canonical-3133310112221111-3211112311021302-2332002032022330-1302021211131112-1300112202033110-3120133000102323-3301230313032221-3131333233103121): complete subsection reference.

<a id="canonical-3130231201110312-2313032121312001-2133000130000033-3013022311101111-2023011002320122-0010112210322221-2300233020333130-2211200202033133"></a>

<a id="canonical-1232131031031231-1130122303120202-1300303221002012-1121111120031233-1202210032322133-3123111033231022-0312010133133202-2033210313033032"></a>

#### `gcp_bucket_receiver.bucket` property

Type: `"string"`. Optional.

GCP Bucket Name. GCP Bucket Name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-z0-9]+[a-z0-9_\\\\.-]+[a-z0-9]$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9_\\\\.-]+[a-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9_\\\\.-]+[a-z0-9]$"
  }
}
```

- [compression](resources--global_log_receiver--reference--group-002.md#canonical-1232123311312122-3101322333002231-3123130033011200-3303321211110303-0033003121032320-3233320022020311-1333131120222310-1122111012131010): complete subsection reference.

- [filename_options](resources--global_log_receiver--reference--group-002.md#canonical-2021230000023312-3203222230013203-1110123330012120-3023232011030102-0013000110312031-2112130203121003-2200023323233220-2311112121100122): complete subsection reference.

- [gcp_cred](resources--global_log_receiver--reference--group-002.md#canonical-0133101020130210-0230322212213100-1003202332233312-2210312222232001-3031130333211231-2201030332232330-0000203223001032-2100203032233122): complete subsection reference.

<a id="canonical-3133310112221111-3211112311021302-2332002032022330-1302021211131112-1300112202033110-3120133000102323-3301230313032221-3131333233103121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.batch` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- gcp_bucket_receiver.batch

<a id="canonical-2011111023320021-3023300220300200-2222110221112221-0303300000202213-3331330003213113-2301013013130301-0333123232022331-2002203130101121"></a>

Type: `"object"`. single nested block, Optional.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_bytes",
    "max_bytes_disabled"),
  validators.ConflictingObjectAttributes("max_events",
    "max_events_disabled"),
  validators.ConflictingObjectAttributes("timeout_seconds",
    "timeout_seconds_default")}
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
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

Terraform syntax:

```terraform
batch {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030210033232110-0021223301110203-3132100121022311-3201231103213100-2330322233330133-0321121100131112-3321110212020030-0230232231223031"></a>

### Direct properties for `gcp_bucket_receiver.batch`

<a id="canonical-0223013312231200-2103020030302213-3223210031031330-0233012223010022-2110300212232331-1331323103120313-2322130213203320-3032113130130102"></a>

#### `gcp_bucket_receiver.batch.max_bytes` property

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(4096, 10485760),
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](resources--global_log_receiver--reference--group-002.md#canonical-2112232112210321-1320202121330001-2301100002130202-0122031332022220-1221233011030030-0000033021003233-0332132231100200-0212320233132110): complete subsection reference.

<a id="canonical-1030312123132312-1011002010010202-0122120223121122-2122123313211032-0320020131013102-0021211310301011-1110231000311100-0332013213010320"></a>

<a id="canonical-1302031200030302-0313331011020103-3033001020320132-2133100323320013-0032100010233302-3111010233200021-2301332333300121-1212203301130212"></a>

#### `gcp_bucket_receiver.batch.max_events` property

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(32, 2000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](resources--global_log_receiver--reference--group-002.md#canonical-0310031031220202-3201210220000022-0313033002111113-1220312002121310-3223303121231321-2330223002020310-2212023330031221-1311222112021220): complete subsection reference.

<a id="canonical-0302122321111103-2302130231010202-2100233130102131-2010232133300130-1313330131122010-1002303111230201-2122233030331210-2023121113032331"></a>

<a id="canonical-0203312010231123-2030032201013032-3101133312123230-2130131001022122-0023211020111200-3323003012120223-0231311302210110-0220102320131323"></a>

#### `gcp_bucket_receiver.batch.timeout_seconds` property

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](resources--global_log_receiver--reference--group-002.md#canonical-3330333233300320-3222311303113321-0210333230001230-3121321120322212-0033120130222020-0201232021221323-3000020200323013-3100003020230121): complete subsection reference.

<a id="canonical-2112232112210321-1320202121330001-2301100002130202-0122031332022220-1221233011030030-0000033021003233-0332132231100200-0212320233132110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.batch.max_bytes_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-3133310112221111-3211112311021302-2332002032022330-1302021211131112-1300112202033110-3120133000102323-3301230313032221-3131333233103121)
- gcp_bucket_receiver.batch.max_bytes_disabled

<a id="canonical-0302321120031321-1031302022223301-1032231111231213-3220103021120122-0010131102123321-0133213223200223-0133210312022003-0230121201202000"></a>

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
max_bytes_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310031031220202-3201210220000022-0313033002111113-1220312002121310-3223303121231321-2330223002020310-2212023330031221-1311222112021220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.batch.max_events_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-3133310112221111-3211112311021302-2332002032022330-1302021211131112-1300112202033110-3120133000102323-3301230313032221-3131333233103121)
- gcp_bucket_receiver.batch.max_events_disabled

<a id="canonical-0200020232111312-1113302203331230-0011233020301100-1313033200032013-1000021222011203-2130000202032331-3000013121321102-1223313013230120"></a>

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
max_events_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330333233300320-3222311303113321-0210333230001230-3121321120322212-0033120130222020-0201232021221323-3000020200323013-3100003020230121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.batch.timeout_seconds_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-3133310112221111-3211112311021302-2332002032022330-1302021211131112-1300112202033110-3120133000102323-3301230313032221-3131333233103121)
- gcp_bucket_receiver.batch.timeout_seconds_default

<a id="canonical-3133313201003133-3230022032311232-3020001002322113-1130301011331210-0231201033101200-1010302001020302-2212010300103121-2331111233312131"></a>

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
timeout_seconds_default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232123311312122-3101322333002231-3123130033011200-3303321211110303-0033003121032320-3233320022020311-1333131120222310-1122111012131010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.compression` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- gcp_bucket_receiver.compression

<a id="canonical-1113100101333332-3031020200011212-0113011233222121-0210211132002130-1121312010223302-0002133022300102-2020003231021321-2001132123010002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Additional upstream details:

Compression Type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("compression_default",
    "compression_gzip"),
  validators.ConflictingObjectAttributes("compression_default",
    "compression_none"),
  validators.ConflictingObjectAttributes("compression_gzip",
    "compression_none")}
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
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

Terraform syntax:

```terraform
compression {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020312102101022-3131230011003011-0103333003013130-2103333011010032-2120322122100332-2102032033123230-0201212001300300-3032200312222203"></a>

### Direct properties for `gcp_bucket_receiver.compression`

- [compression_default](resources--global_log_receiver--reference--group-002.md#canonical-2223232203330200-2233023000001131-2230121020103101-0320310031010130-3323201130131111-3102202230111210-0133123200100233-3101032223322203): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-002.md#canonical-3313110332301023-2331103333102321-3020213133313302-2121211103001110-0113302022212311-3000001221322111-3223020103322111-0313023032130202): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-002.md#canonical-1221121211130122-1032200200200120-2133321131120020-1330100223232300-3231332322022122-0331021321033303-1131303203322322-1100022103230310): complete subsection reference.

<a id="canonical-2223232203330200-2233023000001131-2230121020103101-0320310031010130-3323201130131111-3102202230111210-0133123200100233-3101032223322203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.compression.compression_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-1232123311312122-3101322333002231-3123130033011200-3303321211110303-0033003121032320-3233320022020311-1333131120222310-1122111012131010)
- gcp_bucket_receiver.compression.compression_default

<a id="canonical-0002023331200331-0103323020333300-3213001300131331-0213202213222012-1012033211301210-2233003110313020-0032003333233231-2321130131313003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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
compression_default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313110332301023-2331103333102321-3020213133313302-2121211103001110-0113302022212311-3000001221322111-3223020103322111-0313023032130202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.compression.compression_gzip` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-1232123311312122-3101322333002231-3123130033011200-3303321211110303-0033003121032320-3233320022020311-1333131120222310-1122111012131010)
- gcp_bucket_receiver.compression.compression_gzip

<a id="canonical-3011211111323330-1101102103301301-3100000032322022-0311102321320103-2323020011231330-1332010300012032-0111031322323303-0111020000311210"></a>

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
compression_gzip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221121211130122-1032200200200120-2133321131120020-1330100223232300-3231332322022122-0331021321033303-1131303203322322-1100022103230310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.compression.compression_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-1232123311312122-3101322333002231-3123130033011200-3303321211110303-0033003121032320-3233320022020311-1333131120222310-1122111012131010)
- gcp_bucket_receiver.compression.compression_none

<a id="canonical-3112321202331322-0322330221231222-1033312200020312-3330213213211201-1200023132213030-3231212201212232-0010113020120231-1103233312223112"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression none.

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
compression_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021230000023312-3203222230013203-1110123330012120-3023232011030102-0013000110312031-2112130203121003-2200023323233220-2311112121100122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.filename_options` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- gcp_bucket_receiver.filename_options

<a id="canonical-2030133222111302-3213332031002101-2200321321200110-2021201112112012-2003202312120122-0331021230330030-1112232032132023-3212200123310222"></a>

Type: `"object"`. single nested block, Optional.

Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint
bucket or file.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_folder",
    "log_type_folder"),
  validators.ConflictingObjectAttributes("custom_folder",
    "no_folder"),
  validators.ConflictingObjectAttributes("log_type_folder",
    "no_folder")}
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
  "x-ves-oneof-field-folder": "[\"custom_folder\",\"log_type_folder\",\"no_folder\"]"
}
```

Terraform syntax:

```terraform
filename_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020201031302112-2222333013222000-1023323300112223-2213133213011133-2213000220030023-1030123130200010-2200311112232121-1030223020112232"></a>

### Direct properties for `gcp_bucket_receiver.filename_options`

<a id="canonical-2233333211102303-2101002211033330-2221010210020112-1332021230321220-3303213300002100-0111330223213213-3220123130311201-0112233212311320"></a>

#### `gcp_bucket_receiver.filename_options.custom_folder` property

Type: `"string"`. Optional.

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match \`/^\[a-z\_\]\[a-z0-9\\\\-\\\\.\_\]\*$/i\`

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
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  }
}
```

- [log_type_folder](resources--global_log_receiver--reference--group-002.md#canonical-3022233221120233-1000320331230333-2010030300302011-1232223000120223-1330323012121211-2311101311131002-1132323023020223-3322333322230223): complete subsection reference.

- [no_folder](resources--global_log_receiver--reference--group-002.md#canonical-3112031102120102-0011222021120121-2303231033002133-1122031210310202-1320320322121122-1212213301012313-3121211132330212-2300110012333033): complete subsection reference.

<a id="canonical-3022233221120233-1000320331230333-2010030300302011-1232223000120223-1330323012121211-2311101311131002-1132323023020223-3322333322230223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.filename_options.log_type_folder` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-2021230000023312-3203222230013203-1110123330012120-3023232011030102-0013000110312031-2112130203121003-2200023323233220-2311112121100122)
- gcp_bucket_receiver.filename_options.log_type_folder

<a id="canonical-0212210220323110-2102232312331320-0231130123323122-2323111200103133-1323113103003030-1031232101303100-3101010020222032-0130111311301230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for log type folder.

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
log_type_folder = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112031102120102-0011222021120121-2303231033002133-1122031210310202-1320320322121122-1212213301012313-3121211132330212-2300110012333033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.filename_options.no_folder` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-2021230000023312-3203222230013203-1110123330012120-3023232011030102-0013000110312031-2112130203121003-2200023323233220-2311112121100122)
- gcp_bucket_receiver.filename_options.no_folder

<a id="canonical-0002110330232213-1110120223121211-0311033212020332-3113023000323223-2301321102121102-3320203110320113-3100310200001333-0013121110133203"></a>

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
no_folder = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133101020130210-0230322212213100-1003202332233312-2210312222232001-3031130333211231-2201030332232330-0000203223001032-2100203032233122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.gcp_cred` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- gcp_bucket_receiver.gcp_cred

<a id="canonical-0230223301232110-1003332030021033-2301330313131021-1310333023303130-2201131303320202-2103101133122332-0133210321101200-3110101212223300"></a>

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
gcp_cred {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312322000213023-1011112103011202-1320103122021112-2200001221021121-0130011310001302-1001130202302131-1120212212033300-2220133210320022"></a>

### Direct properties for `gcp_bucket_receiver.gcp_cred`

<a id="canonical-0030321011133023-3000312230010333-2100310031233221-0132303023002120-3312031320233102-0300333023300000-0202201020111110-3031321232120222"></a>

#### `gcp_bucket_receiver.gcp_cred.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1211202031120213-1031033321001130-0300201111213320-3003332012222233-0113131330220011-3003210230020312-3322101023012001-3212223330332330"></a>

<a id="canonical-1202220231030312-1123231311033011-2323112320033333-0210100303110012-3312321303213311-3031013133122332-1021031123232112-0003022123102103"></a>

#### `gcp_bucket_receiver.gcp_cred.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1000113133002230-3130231000312130-0030032132013100-0201120132023133-3323031302001200-0111303100330012-0221320103103220-3321311022322102"></a>

<a id="canonical-0033121202230201-2021232103003231-0022030012001121-0312002010221322-1200301031232221-3233113031301221-2312013031320012-2331303131003123"></a>

#### `gcp_bucket_receiver.gcp_cred.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- http_receiver

<a id="canonical-1220320110200100-0011203102210333-0202102011023021-3001200203203131-1331100111313102-1233023130201221-3030123213313322-1030001103211002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http receiver.

Additional upstream details:

Configuration for HTTP endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("uri"),
  validators.ConflictingObjectAttributes("auth_basic",
    "auth_none"),
  validators.ConflictingObjectAttributes("auth_basic",
    "auth_token"),
  validators.ConflictingObjectAttributes("auth_none",
    "auth_token"),
  validators.ConflictingObjectAttributes("no_tls",
    "use_tls")}
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
  "x-ves-oneof-field-auth_choice": "[\"auth_basic\",\"auth_none\",\"auth_token\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
http_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232101330300120-2121002322221002-2111310222112013-0022002330021211-0120232102300120-3333003322110003-0100322302130020-0120003322230333"></a>

### Direct properties for `http_receiver`

- [auth_basic](resources--global_log_receiver--reference--group-002.md#canonical-1022130202001203-2123222000302011-0303000113321310-1202322232132113-3001121022210122-3232110303312101-0210032303121020-2131003101302233): complete subsection reference.

- [auth_none](resources--global_log_receiver--reference--group-003.md#canonical-3331330032301033-3102210321202122-3212212113000230-1203031121110232-3301023022023030-1033012130200332-3101222021333330-1303103201031213): complete subsection reference.

- [auth_token](resources--global_log_receiver--reference--group-003.md#canonical-3211110223222312-1122132110030223-3113232133213102-3013031131221331-1303210011032221-3303131311330213-2103031022323113-1102013133333303): complete subsection reference.

- [batch](resources--global_log_receiver--reference--group-003.md#canonical-0102310011303233-1212220212302323-3313000020023020-1100320133210211-2333121320223233-2313312313131202-1200320323030001-1320032012112211): complete subsection reference.

- [compression](resources--global_log_receiver--reference--group-003.md#canonical-3133203112210102-2303223113321220-1030012202132221-1231310113101233-1233330221020123-1232300100103012-1321231320233133-2103012323303331): complete subsection reference.

- [no_tls](resources--global_log_receiver--reference--group-003.md#canonical-3233100233103022-2230123213321331-0131001032100301-2102322211220020-3112230003112133-0301020101101012-1200122011110232-2212133013103130): complete subsection reference.

<a id="canonical-1033020103120002-0123213303012030-3303331000212223-1303000233331233-3000022223011201-2333130132001220-3313100121010130-2122211223322303"></a>

<a id="canonical-3111311311101131-1312013200212131-1132133022232201-1033103313330013-0111310313132101-2320303323103111-0332233332120310-3130001202010332"></a>

#### `http_receiver.uri` property

Type: `"string"`. Optional.

HTTP URI is the URI of the HTTP endpoint to send logs to,.

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
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123): complete subsection reference.

<a id="canonical-1022130202001203-2123222000302011-0303000113321310-1202322232132113-3001121022210122-3232110303312101-0210032303121020-2131003101302233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_basic` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- http_receiver.auth_basic

<a id="canonical-3212310332211222-3232202111020220-3031310313113232-3313212222012200-2003321211113132-1013230300231033-1231033202002022-3310212132331000"></a>

Type: `"object"`. single nested block, Optional.

Authentication parameters to access HTPP Log Receiver Endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
auth_basic {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300320221211102-0222110221300103-1123023032021113-3122021301323032-1123112030213123-1001030000023021-0031021203211122-3202203321033200"></a>

### Direct properties for `http_receiver.auth_basic`

- [password](resources--global_log_receiver--reference--group-003.md#canonical-1033323112022211-3003201200310311-3010321003132310-1213101100232030-1130010321211202-3321121122113013-0110212000322033-2023303300123230): complete subsection reference.

<a id="canonical-3101333321222000-2322120102103323-1302130201121123-2133330211330123-3310302030031022-3110130122022333-0222111330301223-1310103230330213"></a>

<a id="canonical-3232322003232002-1222013123230011-3102011302331102-2002302031200130-2312001303121301-3031103120202300-1202001212333130-2320011331212332"></a>

#### `http_receiver.auth_basic.user_name` property

Type: `"string"`. Optional.

username. HTTP Basic Auth username.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```
