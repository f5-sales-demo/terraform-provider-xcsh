---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-1103020122120201-2032021323300321-3130001322133122-3131323200133222-2000133032321033-2013223033213322-3332100131211201-0002320023201231"></a>

## Direct properties for `s3_receiver.batch`

<a id="canonical-0113101013331121-3130111212012123-3321030210031320-0232113101331210-0000212130121331-1221222130101001-0002103012203232-1103033010021130"></a>

### `s3_receiver.batch.max_bytes` property

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

- [max_bytes_disabled](resources--global_log_receiver--reference--group-005.md#canonical-0321130201332121-0213022012232031-2020331310122110-2000223111132213-1030001032023110-2100132111010221-3130020201010130-0332212321311332): complete subsection reference.

<a id="canonical-1000333231232102-1330102030003021-3031303123201210-2121021330221002-3333330332330001-2203213320022233-3223322302231130-3031120003130222"></a>

<a id="canonical-0100331331111222-3013132123123320-3333233132131311-3112310212212333-1100202103210223-0232033013102023-0332110231012332-2200311131230033"></a>

### `s3_receiver.batch.max_events` property

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

- [max_events_disabled](resources--global_log_receiver--reference--group-005.md#canonical-0332011113233200-3103031331111220-1320310112330301-1011331220111200-2003030013103122-0131033023121031-3001330121223023-1020030200022033): complete subsection reference.

<a id="canonical-3233213113101112-0321213132031321-0203110103311031-0021323212303032-2332222221202023-2001001100302322-2213023313003310-1121100323022223"></a>

<a id="canonical-0310203112310112-1203303212233323-1230010101300233-2303331212220300-0012323032023032-1101031011121000-3231210030102323-2130023210000211"></a>

### `s3_receiver.batch.timeout_seconds` property

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

- [timeout_seconds_default](resources--global_log_receiver--reference--group-005.md#canonical-1003000213233232-1120203302312111-3023323111233222-2223122312233020-3023012202202123-2111011202202222-0002200031113010-1213002003130212): complete subsection reference.

<a id="canonical-0321130201332121-0213022012232031-2020331310122110-2000223111132213-1030001032023110-2100132111010221-3130020201010130-0332212321311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.batch.max_bytes_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122)
- [s3_receiver.batch](resources--global_log_receiver--reference--group-004.md#canonical-2212222233301133-1023002210322033-0231130330222102-0011130133000313-3300323323033222-2033232313203232-3013223202013210-1211323111232000)
- s3_receiver.batch.max_bytes_disabled

<a id="canonical-3023230222202322-1322032200312333-2113321022221013-3213100103023331-0013132023000210-1200222021223322-3202103312321032-1112002210102102"></a>

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

<a id="canonical-0332011113233200-3103031331111220-1320310112330301-1011331220111200-2003030013103122-0131033023121031-3001330121223023-1020030200022033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.batch.max_events_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122)
- [s3_receiver.batch](resources--global_log_receiver--reference--group-004.md#canonical-2212222233301133-1023002210322033-0231130330222102-0011130133000313-3300323323033222-2033232313203232-3013223202013210-1211323111232000)
- s3_receiver.batch.max_events_disabled

<a id="canonical-0022222122101311-2200322223103313-1103312312023330-2101103001310133-2301313221301123-1131223330130003-2320331321213220-2222220110200323"></a>

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

<a id="canonical-1003000213233232-1120203302312111-3023323111233222-2223122312233020-3023012202202123-2111011202202222-0002200031113010-1213002003130212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.batch.timeout_seconds_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122)
- [s3_receiver.batch](resources--global_log_receiver--reference--group-004.md#canonical-2212222233301133-1023002210322033-0231130330222102-0011130133000313-3300323323033222-2033232313203232-3013223202013210-1211323111232000)
- s3_receiver.batch.timeout_seconds_default

<a id="canonical-2312332231221303-2130010031321002-1031200133223200-3203221022130201-0003003112002101-2033300200011313-1312323013311332-1132233113123111"></a>

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

<a id="canonical-2111330022112232-1022203012213032-2120111001213003-2321001220303202-0130300211313331-3300322120302121-3012320000202120-1022011011101303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.compression` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122)
- s3_receiver.compression

<a id="canonical-1011210030121103-0232020132001132-3211333102230330-2120031231032323-3321010332020310-0030301013332320-3202233221230300-1133201023001312"></a>

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

<a id="canonical-0322203203330302-3033231211203303-1322213100123200-2221233121113302-3220132000300223-0100330232031101-0310121112310222-2332133333111222"></a>

### Direct properties for `s3_receiver.compression`

- [compression_default](resources--global_log_receiver--reference--group-005.md#canonical-2232000223310212-2310002112201211-1031110231211300-2230132313120112-1213330033213213-3303010112011300-0322020312001203-1230113100111322): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-005.md#canonical-0220003103110012-2222211113033102-2231212202011213-2021133201033311-3032131333321201-3033131012232102-0101111223213122-2212231002313011): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-005.md#canonical-3132203120123222-1012333132232200-3111222120312212-1202230112213103-1332010311313300-3201132003112010-0301232111130122-0301123200130210): complete subsection reference.

<a id="canonical-2232000223310212-2310002112201211-1031110231211300-2230132313120112-1213330033213213-3303010112011300-0322020312001203-1230113100111322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.compression.compression_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122)
- [s3_receiver.compression](resources--global_log_receiver--reference--group-005.md#canonical-2111330022112232-1022203012213032-2120111001213003-2321001220303202-0130300211313331-3300322120302121-3012320000202120-1022011011101303)
- s3_receiver.compression.compression_default

<a id="canonical-3013230132333031-3201003101131320-2303120200203300-0221213333002122-0120223301131203-3211012132000130-1220320230200203-0122111313020201"></a>

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

<a id="canonical-0220003103110012-2222211113033102-2231212202011213-2021133201033311-3032131333321201-3033131012232102-0101111223213122-2212231002313011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.compression.compression_gzip` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122)
- [s3_receiver.compression](resources--global_log_receiver--reference--group-005.md#canonical-2111330022112232-1022203012213032-2120111001213003-2321001220303202-0130300211313331-3300322120302121-3012320000202120-1022011011101303)
- s3_receiver.compression.compression_gzip

<a id="canonical-1032002013300213-0020223002132132-0010200331211310-0110131200213011-0132331330023201-0033201232311330-3003220302131100-3031003210330030"></a>

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

<a id="canonical-3132203120123222-1012333132232200-3111222120312212-1202230112213103-1332010311313300-3201132003112010-0301232111130122-0301123200130210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.compression.compression_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122)
- [s3_receiver.compression](resources--global_log_receiver--reference--group-005.md#canonical-2111330022112232-1022203012213032-2120111001213003-2321001220303202-0130300211313331-3300322120302121-3012320000202120-1022011011101303)
- s3_receiver.compression.compression_none

<a id="canonical-1311201200012203-2332233221333223-2211120102313330-0120213232231312-1010310201210221-2312131300100211-0303211223010312-2221112032002212"></a>

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

<a id="canonical-0310300212301021-1030032312012230-0231233231132003-0310010220111031-0000031333323200-1103211313320310-2031132312223122-1323220303231200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.filename_options` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122)
- s3_receiver.filename_options

<a id="canonical-2012320332233303-2002230101231132-3231323102013020-1000333131330223-0213122213001223-1102203333033013-3000233320132320-3112333220233223"></a>

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

<a id="canonical-2113110213203011-3331201032211221-1023320321031302-1111220212311131-0232233321133312-2030122230301000-3013103302131100-2111320323311033"></a>

### Direct properties for `s3_receiver.filename_options`

<a id="canonical-0022012322321123-2013112133121200-1200102111122211-2010020120103032-1320201012330210-3302130003302201-1030113203030323-0201202200100010"></a>

#### `s3_receiver.filename_options.custom_folder` property

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

- [log_type_folder](resources--global_log_receiver--reference--group-005.md#canonical-3202111322221322-3121013110333310-0231110020021123-0310333132302001-3301232303212210-0320212321230212-0212113012012101-1001312320303212): complete subsection reference.

- [no_folder](resources--global_log_receiver--reference--group-005.md#canonical-1102332322222222-2220303310013101-2221110100220123-2013021201120200-3301330231212102-1103030022200320-3213202213031211-0332021021333302): complete subsection reference.

<a id="canonical-3202111322221322-3121013110333310-0231110020021123-0310333132302001-3301232303212210-0320212321230212-0212113012012101-1001312320303212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.filename_options.log_type_folder` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122)
- [s3_receiver.filename_options](resources--global_log_receiver--reference--group-005.md#canonical-0310300212301021-1030032312012230-0231233231132003-0310010220111031-0000031333323200-1103211313320310-2031132312223122-1323220303231200)
- s3_receiver.filename_options.log_type_folder

<a id="canonical-0320023300210110-1213003231020313-3333011202213100-3023023113201212-3320123010330031-1113012013333031-2122120132230211-3021130033212113"></a>

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

<a id="canonical-1102332322222222-2220303310013101-2221110100220123-2013021201120200-3301330231212102-1103030022200320-3213202213031211-0332021021333302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.filename_options.no_folder` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122)
- [s3_receiver.filename_options](resources--global_log_receiver--reference--group-005.md#canonical-0310300212301021-1030032312012230-0231233231132003-0310010220111031-0000031333323200-1103211313320310-2031132312223122-1323220303231200)
- s3_receiver.filename_options.no_folder

<a id="canonical-3233202022203000-3103123123221132-2210213323000210-1231320213311111-1001332223000103-0113002222201121-1000100101202133-1030030231230131"></a>

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

<a id="canonical-3211022021302113-1000032210211020-3113000022100301-1332120310011110-2113003230211310-2211212131301002-0201111232112322-1122230010302122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `security_events` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- security_events

<a id="canonical-1211211103013320-0312220220033311-0221122101003311-3101230113233111-1103132232203013-3121322010213021-1202332113201032-2002003231210233"></a>

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
security_events = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- splunk_receiver

<a id="canonical-3323012321222032-0301123231022010-0310113310103303-0311323212010001-2321221011222303-3332101132321300-2330331123203201-1202330320122010"></a>

Type: `"object"`. single nested block, Optional.

Configuration for Splunk HEC Logs endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("endpoint"),
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
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
splunk_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303023020122020-1032011331113230-2200010233100013-3012212103022102-3022030223330033-0101103312033202-3331312012132003-2330110220030222"></a>

### Direct properties for `splunk_receiver`

- [batch](resources--global_log_receiver--reference--group-005.md#canonical-3022103112101030-3030322221021011-1100211102020311-3203210031202132-3132332232211032-1120213321203332-0011203200220120-0203203310213313): complete subsection reference.

- [compression](resources--global_log_receiver--reference--group-005.md#canonical-2010123313311221-1030023122002032-3012203213310221-3131203202200000-2221302102003322-0021032303110122-2021131210112022-3211201020320223): complete subsection reference.

<a id="canonical-3211201101231311-2211021131220000-2102311221031133-2202131000020002-1322322112232330-3300120300213333-0300233210202033-2031010000012211"></a>

<a id="canonical-3011121102011120-2001232021312123-1203323121223033-0131331033011333-1331002123323210-3123101122012021-0012133300120230-2202120210302322"></a>

#### `splunk_receiver.endpoint` property

Type: `"string"`. Optional.

Splunk HEC Logs Endpoint. Splunk HEC Logs Endpoint, (Note: must not contain \`/services/collector\`)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
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
    },
    "minLength": 1,
    "pattern": "^https?://[^\\s/$.?#].[^\\s]*$"
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

- [no_tls](resources--global_log_receiver--reference--group-005.md#canonical-2310330213200023-2211313133003121-0113031232323320-0313002113212330-0023230223320230-3232313110110311-2013022321212133-1302202110323132): complete subsection reference.

- [splunk_hec_token](resources--global_log_receiver--reference--group-005.md#canonical-0000330110231330-2023201131210100-3101201332111030-0210121333201332-3000210133232120-1003111101031000-3313010103012022-2103123333112310): complete subsection reference.

- [use_tls](resources--global_log_receiver--reference--group-005.md#canonical-1032311130223211-1300102100202101-0200123213003120-1021210202000030-0002113312200122-1000220323112123-0100032221032310-1111233222310112): complete subsection reference.

<a id="canonical-3022103112101030-3030322221021011-1100211102020311-3203210031202132-3132332232211032-1120213321203332-0011203200220120-0203203310213313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.batch` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- splunk_receiver.batch

<a id="canonical-3001101301220201-2313210130330123-2312220230201203-3201122231333113-1102330203330311-1330030132330233-2133011303111213-0300113201302131"></a>

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

<a id="canonical-0012321011210033-0002201303323212-0133121000112332-1101212231011212-3220003211311220-0022011110031103-2221321012232010-1300221321003033"></a>

### Direct properties for `splunk_receiver.batch`

<a id="canonical-3222330330231301-3221221032030102-2020202001013301-1102202331201002-3022023031220120-2320122212010111-3111221111123000-2300122203023230"></a>

#### `splunk_receiver.batch.max_bytes` property

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

- [max_bytes_disabled](resources--global_log_receiver--reference--group-005.md#canonical-2231210211002131-2100003222320311-2132101201122032-0211120312220210-2200303132122202-3111120030210130-0213221102332330-0003323302322232): complete subsection reference.

<a id="canonical-1302310120203021-2101100020320032-3002102201223002-2013321101211302-0332300013103001-0011110212322213-3111203031030322-1222112123122023"></a>

<a id="canonical-2223010122000323-0111223110102201-0213301032132220-0120030130323022-2223023021331102-0103310223030102-2322013103020132-3003213130010012"></a>

#### `splunk_receiver.batch.max_events` property

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

- [max_events_disabled](resources--global_log_receiver--reference--group-005.md#canonical-2021000121032012-1203321203120000-2301123211221222-1213331111031130-1111031211320300-2320122020210012-2321010003210231-3100022001323231): complete subsection reference.

<a id="canonical-3102120222323323-0202220321211302-1222213031302320-1000312112303112-0213122023223023-0223323122203100-2121132320103203-2122100200332303"></a>

<a id="canonical-3113300002322032-2303333313022220-1102103212302112-3120011303201121-1203021322030301-1130020200123333-1332231032303322-3010321222030222"></a>

#### `splunk_receiver.batch.timeout_seconds` property

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

- [timeout_seconds_default](resources--global_log_receiver--reference--group-005.md#canonical-3320110000122221-2333131323230303-0231000203212230-2200111311011121-1012332022020310-0330232103223221-1301220223103132-3233031032230333): complete subsection reference.

<a id="canonical-2231210211002131-2100003222320311-2132101201122032-0211120312220210-2200303132122202-3111120030210130-0213221102332330-0003323302322232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.batch.max_bytes_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.batch](resources--global_log_receiver--reference--group-005.md#canonical-3022103112101030-3030322221021011-1100211102020311-3203210031202132-3132332232211032-1120213321203332-0011203200220120-0203203310213313)
- splunk_receiver.batch.max_bytes_disabled

<a id="canonical-1130103322320110-0232222220021222-3220233202300102-3322022001013211-1300301222323131-2123300303000123-1333312123300322-0112101120120200"></a>

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

<a id="canonical-2021000121032012-1203321203120000-2301123211221222-1213331111031130-1111031211320300-2320122020210012-2321010003210231-3100022001323231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.batch.max_events_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.batch](resources--global_log_receiver--reference--group-005.md#canonical-3022103112101030-3030322221021011-1100211102020311-3203210031202132-3132332232211032-1120213321203332-0011203200220120-0203203310213313)
- splunk_receiver.batch.max_events_disabled

<a id="canonical-0211320000123013-2301203330032120-2120332301313232-1130032003210130-0113332012111011-1230223122020023-2221310223211030-1120321131123013"></a>

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

<a id="canonical-3320110000122221-2333131323230303-0231000203212230-2200111311011121-1012332022020310-0330232103223221-1301220223103132-3233031032230333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.batch.timeout_seconds_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.batch](resources--global_log_receiver--reference--group-005.md#canonical-3022103112101030-3030322221021011-1100211102020311-3203210031202132-3132332232211032-1120213321203332-0011203200220120-0203203310213313)
- splunk_receiver.batch.timeout_seconds_default

<a id="canonical-0002211212013112-2211313321222322-1230121100003302-0231133331021113-3212330132201133-2233113123301030-0203122320231110-1030302013211012"></a>

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

<a id="canonical-2010123313311221-1030023122002032-3012203213310221-3131203202200000-2221302102003322-0021032303110122-2021131210112022-3211201020320223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.compression` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- splunk_receiver.compression

<a id="canonical-0123102103323303-2301231012333233-0231133100032202-2002200312130003-2212032232220002-1133003113203100-3330120300202211-3113003222130030"></a>

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

<a id="canonical-0302211110201302-1303113232233100-0312322010332202-0021020203030330-0223210221322313-2031223233101032-1002322310122311-3331111021030131"></a>

### Direct properties for `splunk_receiver.compression`

- [compression_default](resources--global_log_receiver--reference--group-005.md#canonical-0210233011102313-3002003003223202-3102223130132102-1333213021021132-2322220101121112-2321011331302032-0013203310001012-3310220111230110): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-005.md#canonical-2311223300303300-3122033133012121-3303021121121232-0120010002232320-1300103332301133-2312333033202121-3131023103223200-2130333313100321): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-005.md#canonical-1022202221213102-3202310322223210-2313031212333312-1233013312001113-3312320030101233-2302003102012330-1220223121030132-1113022021211203): complete subsection reference.

<a id="canonical-0210233011102313-3002003003223202-3102223130132102-1333213021021132-2322220101121112-2321011331302032-0013203310001012-3310220111230110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.compression.compression_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.compression](resources--global_log_receiver--reference--group-005.md#canonical-2010123313311221-1030023122002032-3012203213310221-3131203202200000-2221302102003322-0021032303110122-2021131210112022-3211201020320223)
- splunk_receiver.compression.compression_default

<a id="canonical-3133220301122210-1231223220302230-2321321233020110-2303332111310033-1213303101230010-3312333221330220-2302132230020131-1212321120201311"></a>

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

<a id="canonical-2311223300303300-3122033133012121-3303021121121232-0120010002232320-1300103332301133-2312333033202121-3131023103223200-2130333313100321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.compression.compression_gzip` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.compression](resources--global_log_receiver--reference--group-005.md#canonical-2010123313311221-1030023122002032-3012203213310221-3131203202200000-2221302102003322-0021032303110122-2021131210112022-3211201020320223)
- splunk_receiver.compression.compression_gzip

<a id="canonical-3033320211220110-1032120223122323-0012230130023302-2102101100213211-1233130133323112-0032033320221002-2001212310001111-1231101001310201"></a>

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

<a id="canonical-1022202221213102-3202310322223210-2313031212333312-1233013312001113-3312320030101233-2302003102012330-1220223121030132-1113022021211203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.compression.compression_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.compression](resources--global_log_receiver--reference--group-005.md#canonical-2010123313311221-1030023122002032-3012203213310221-3131203202200000-2221302102003322-0021032303110122-2021131210112022-3211201020320223)
- splunk_receiver.compression.compression_none

<a id="canonical-3213231113033213-0202202113131310-2301300323330003-2322031300200200-1013031223033331-0022221302311003-0310013200310013-2321003231302330"></a>

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

<a id="canonical-2310330213200023-2211313133003121-0113031232323320-0313002113212330-0023230223320230-3232313110110311-2013022321212133-1302202110323132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.no_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- splunk_receiver.no_tls

<a id="canonical-2302200133232232-2210300012333311-3113310120313132-1110312113332100-0203130310212231-0103233233131011-0331032301103010-1123332021222332"></a>

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

<a id="canonical-0000330110231330-2023201131210100-3101201332111030-0210121333201332-3000210133232120-1003111101031000-3313010103012022-2103123333112310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.splunk_hec_token` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- splunk_receiver.splunk_hec_token

<a id="canonical-0212200222200100-1022012001210000-3330020001330201-3302103121101113-2200210320313132-3332120320213030-0000022333333020-3232100021302021"></a>

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
splunk_hec_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310021301321212-3321010232002231-2203202132131231-3002003202003013-3110012000223320-1111002231133311-2202101213202310-2031103202233300"></a>

### Direct properties for `splunk_receiver.splunk_hec_token`

- [blindfold_secret_info](resources--global_log_receiver--reference--group-005.md#canonical-1232203330123312-2130100122003202-3230233312023131-0113001223230220-1030322233122101-0003003221233313-3020110011322333-0221223311103131): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-005.md#canonical-3021122321113120-0112313030231122-3131321002113011-2000320132332000-1312031133332220-2030101220312332-0301003221302330-3031031121302312): complete subsection reference.

<a id="canonical-1232203330123312-2130100122003202-3230233312023131-0113001223230220-1030322233122101-0003003221233313-3020110011322333-0221223311103131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.splunk_hec_token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.splunk_hec_token](resources--global_log_receiver--reference--group-005.md#canonical-0000330110231330-2023201131210100-3101201332111030-0210121333201332-3000210133232120-1003111101031000-3313010103012022-2103123333112310)
- splunk_receiver.splunk_hec_token.blindfold_secret_info

<a id="canonical-1113101201321200-2033223030020321-3132233233232313-0203030302223023-3110223101103201-2230000312103203-1303003201311310-0023033032110222"></a>

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

<a id="canonical-1031001213213202-2003211131203013-0113233023033212-2123302231020322-1113321010120222-2100030313012323-1003012211320310-3120321003101311"></a>

### Direct properties for `splunk_receiver.splunk_hec_token.blindfold_secret_info`

<a id="canonical-2110003321223122-1112231210032310-2132111032000003-2333001331321132-3133030233233223-2033101201211231-3211131030101012-0032230212330020"></a>

#### `splunk_receiver.splunk_hec_token.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-2311012201002231-2131031110222313-3323300101311012-0013310213003303-2300220202023211-0312203303322121-0123101200301000-3102210102310122"></a>

<a id="canonical-2101120220111202-3021300213213212-2320020331113333-0220030333120310-1131123033200202-3133002003332323-2202202201013001-1201131310001332"></a>

#### `splunk_receiver.splunk_hec_token.blindfold_secret_info.location` property

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

<a id="canonical-0102013302202322-1300302033222221-2010201010212030-0232103230331101-3112032131121310-2132003303331311-0200012002200020-3323030100123112"></a>

<a id="canonical-3301330102303120-2132312312333301-2212113211212302-3013101122312132-3001233030301023-3331122203121002-1133003223232030-2103003300333202"></a>

#### `splunk_receiver.splunk_hec_token.blindfold_secret_info.store_provider` property

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

<a id="canonical-3021122321113120-0112313030231122-3131321002113011-2000320132332000-1312031133332220-2030101220312332-0301003221302330-3031031121302312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.splunk_hec_token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.splunk_hec_token](resources--global_log_receiver--reference--group-005.md#canonical-0000330110231330-2023201131210100-3101201332111030-0210121333201332-3000210133232120-1003111101031000-3313010103012022-2103123333112310)
- splunk_receiver.splunk_hec_token.clear_secret_info

<a id="canonical-3220033202123033-1022023311133120-2223110223210102-0132331212312202-1000233121200302-1230223321130002-3323221002110020-0230033301111320"></a>

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

<a id="canonical-3213033203133323-2212232330200130-2330031330110122-3323213131230323-1322120333213130-1003021123303200-3322200312232111-0101211311103200"></a>

### Direct properties for `splunk_receiver.splunk_hec_token.clear_secret_info`

<a id="canonical-1013033223000333-0302113230311233-1031321002001212-2210122102123120-1020102122323023-2013332010132023-1220111312203200-3230030302212011"></a>

#### `splunk_receiver.splunk_hec_token.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2100101001013311-1221320300033021-2132121101011210-0303020111311322-3011231130023100-1122032101223011-0022012202300100-1231133300003032"></a>

<a id="canonical-2111101131202033-1133011200230311-0023032312211023-0133033032211011-2132020312130310-1022203222003132-1123001101200213-0311310223030310"></a>

#### `splunk_receiver.splunk_hec_token.clear_secret_info.url` property

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

<a id="canonical-1032311130223211-1300102100202101-0200123213003120-1021210202000030-0002113312200122-1000220323112123-0100032221032310-1111233222310112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- splunk_receiver.use_tls

<a id="canonical-1301012031221211-3010122330230213-0020032213300201-0103010210220100-0101322202133201-0003333032103312-2210320013011312-3033202212330333"></a>

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

<a id="canonical-1330322121012210-3023222301111003-1323022001020203-1201103200033200-0311313211210023-0133330203331012-3103012222120002-3232320100301123"></a>

### Direct properties for `splunk_receiver.use_tls`

- [disable_verify_certificate](resources--global_log_receiver--reference--group-005.md#canonical-2213212021232300-3202331322322312-3101230031330303-1313002302211122-3332322231213333-0203123200010012-0300303002232020-0230100113203010): complete subsection reference.

- [disable_verify_hostname](resources--global_log_receiver--reference--group-005.md#canonical-2311003200321032-1013121033023113-1013012132323231-1221110032020211-3311000312103230-1321233223300032-1101303120300323-3311003100110100): complete subsection reference.

- [enable_verify_certificate](resources--global_log_receiver--reference--group-005.md#canonical-2313311322031201-3110232231033011-0101233100132222-0302233021012320-3333132000201323-0203103130033100-3211000133220211-2020101112010232): complete subsection reference.

- [enable_verify_hostname](resources--global_log_receiver--reference--group-005.md#canonical-2312113311100230-3101001302023100-1211020310032023-2213012003221010-0000002003330000-2311013211133102-3000111011233111-1311002320030321): complete subsection reference.

- [mtls_disabled](resources--global_log_receiver--reference--group-005.md#canonical-0012011230002113-2012120013030000-1230322202121223-0320201331332032-1010120300131033-1030010120313201-3011112001230333-3003212212220331): complete subsection reference.

- [mtls_enable](resources--global_log_receiver--reference--group-005.md#canonical-0211331211121211-3221211003312312-3003211030203121-0121213303021100-0320111011202003-2100103032101123-3313131301003122-0223111011311131): complete subsection reference.

- [no_ca](resources--global_log_receiver--reference--group-005.md#canonical-3001101302222020-1100211020222002-1133112021111032-2132310333310102-1233111322201101-2313330203233223-1200333202223332-0331001121001000): complete subsection reference.

<a id="canonical-0220011300222101-0330103123330110-0200321332123021-1310222103003100-1112011003023310-3322030003223030-1032211212021120-3312332020303021"></a>

<a id="canonical-0132211332233300-3320112313022322-3323232203310103-0232122321303022-0001220122102132-2131221212221132-2131313212213302-2012113201330302"></a>

#### `splunk_receiver.use_tls.trusted_ca_url` property

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

<a id="canonical-2213212021232300-3202331322322312-3101230031330303-1313002302211122-3332322231213333-0203123200010012-0300303002232020-0230100113203010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.disable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.use_tls](resources--global_log_receiver--reference--group-005.md#canonical-1032311130223211-1300102100202101-0200123213003120-1021210202000030-0002113312200122-1000220323112123-0100032221032310-1111233222310112)
- splunk_receiver.use_tls.disable_verify_certificate

<a id="canonical-1002332031302211-3102311302032101-2303012113321113-2102231201022022-2311203222212111-0020001213313311-2012122023223202-1321123112130300"></a>

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

<a id="canonical-2311003200321032-1013121033023113-1013012132323231-1221110032020211-3311000312103230-1321233223300032-1101303120300323-3311003100110100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.disable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.use_tls](resources--global_log_receiver--reference--group-005.md#canonical-1032311130223211-1300102100202101-0200123213003120-1021210202000030-0002113312200122-1000220323112123-0100032221032310-1111233222310112)
- splunk_receiver.use_tls.disable_verify_hostname

<a id="canonical-0313022131210333-3330113001131031-0030123332113211-2213323010310123-3003211003213203-1110233111030203-1012331020233133-1013211213130231"></a>

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

<a id="canonical-2313311322031201-3110232231033011-0101233100132222-0302233021012320-3333132000201323-0203103130033100-3211000133220211-2020101112010232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.enable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.use_tls](resources--global_log_receiver--reference--group-005.md#canonical-1032311130223211-1300102100202101-0200123213003120-1021210202000030-0002113312200122-1000220323112123-0100032221032310-1111233222310112)
- splunk_receiver.use_tls.enable_verify_certificate

<a id="canonical-2331101230323102-3313131111212102-0311123303021221-1333320233003012-1100012123331123-2300021121330121-3030030221212120-2032222323321011"></a>

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

<a id="canonical-2312113311100230-3101001302023100-1211020310032023-2213012003221010-0000002003330000-2311013211133102-3000111011233111-1311002320030321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.enable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.use_tls](resources--global_log_receiver--reference--group-005.md#canonical-1032311130223211-1300102100202101-0200123213003120-1021210202000030-0002113312200122-1000220323112123-0100032221032310-1111233222310112)
- splunk_receiver.use_tls.enable_verify_hostname

<a id="canonical-3310323221111010-0220132232310030-3001113200031303-2303311001302000-3331320110012131-3331310201222030-3013120321333023-3203211210001003"></a>

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

<a id="canonical-0012011230002113-2012120013030000-1230322202121223-0320201331332032-1010120300131033-1030010120313201-3011112001230333-3003212212220331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.mtls_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.use_tls](resources--global_log_receiver--reference--group-005.md#canonical-1032311130223211-1300102100202101-0200123213003120-1021210202000030-0002113312200122-1000220323112123-0100032221032310-1111233222310112)
- splunk_receiver.use_tls.mtls_disabled

<a id="canonical-0020121012101032-1033120030310110-0003310031003332-3300030332230122-2001011322123200-3013211312112121-0012322001221201-2033213021100221"></a>

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

<a id="canonical-0211331211121211-3221211003312312-3003211030203121-0121213303021100-0320111011202003-2100103032101123-3313131301003122-0223111011311131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.mtls_enable` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.use_tls](resources--global_log_receiver--reference--group-005.md#canonical-1032311130223211-1300102100202101-0200123213003120-1021210202000030-0002113312200122-1000220323112123-0100032221032310-1111233222310112)
- splunk_receiver.use_tls.mtls_enable

<a id="canonical-0210231320311301-3103331023000321-2322122012101022-1230331131201310-1123031222332120-2020003320030133-2033131212021012-2312030113113330"></a>

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

<a id="canonical-1331222101321132-0013301212213310-2020233102021223-1022122222213102-3103312120231223-3303032202202203-3221033130223222-3130020000021000"></a>

### Direct properties for `splunk_receiver.use_tls.mtls_enable`

<a id="canonical-3323030330301310-1323021133013320-1313333010331331-0210113130231031-0203130132110211-3012003012012032-1320302312121320-2321321003031212"></a>

#### `splunk_receiver.use_tls.mtls_enable.certificate` property

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

- [key_url](resources--global_log_receiver--reference--group-005.md#canonical-1231103203102133-1122221303110021-3131210322111210-3000131101103222-0221110310122120-1321133312100310-1000323201221013-2300132321100032): complete subsection reference.

<a id="canonical-1231103203102133-1122221303110021-3131210322111210-3000131101103222-0221110310122120-1321133312100310-1000323201221013-2300132321100032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.mtls_enable.key_url` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.use_tls](resources--global_log_receiver--reference--group-005.md#canonical-1032311130223211-1300102100202101-0200123213003120-1021210202000030-0002113312200122-1000220323112123-0100032221032310-1111233222310112)
- [splunk_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-005.md#canonical-0211331211121211-3221211003312312-3003211030203121-0121213303021100-0320111011202003-2100103032101123-3313131301003122-0223111011311131)
- splunk_receiver.use_tls.mtls_enable.key_url

<a id="canonical-1001303023230033-3213231320333033-3302020010223302-2111223100000303-0300312101100032-0333310202203103-2131121213213221-1320222333230010"></a>

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

<a id="canonical-0320112121322223-2033112112300021-1123211011103012-1131203112132221-0310111122310321-2322302020110012-3120213101012001-1331320002100202"></a>

### Direct properties for `splunk_receiver.use_tls.mtls_enable.key_url`

- [blindfold_secret_info](resources--global_log_receiver--reference--group-005.md#canonical-0133103010210311-1020222322100133-1011330210321111-0311313231302133-2310221333033320-3321112012220032-3322103230121232-1311333311131300): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-005.md#canonical-2102031113102202-3022333231113202-2311032000102101-0203030322203200-2002313122232020-0130323221203223-2122312320320332-3300330300002132): complete subsection reference.

<a id="canonical-0133103010210311-1020222322100133-1011330210321111-0311313231302133-2310221333033320-3321112012220032-3322103230121232-1311333311131300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.use_tls](resources--global_log_receiver--reference--group-005.md#canonical-1032311130223211-1300102100202101-0200123213003120-1021210202000030-0002113312200122-1000220323112123-0100032221032310-1111233222310112)
- [splunk_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-005.md#canonical-0211331211121211-3221211003312312-3003211030203121-0121213303021100-0320111011202003-2100103032101123-3313131301003122-0223111011311131)
- [splunk_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-005.md#canonical-1231103203102133-1122221303110021-3131210322111210-3000131101103222-0221110310122120-1321133312100310-1000323201221013-2300132321100032)
- splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-3021313132232111-1131211332213330-1211123012033123-0000121003220201-3321222303220113-0310213313131312-3000123020233310-3220221003230232"></a>

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

<a id="canonical-0003001021130122-1131120133133203-3323121000003100-2313013131233231-1102210303332222-3033002123102001-3233331202113211-0103202120031203"></a>

### Direct properties for `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info`

<a id="canonical-3003032223302033-2132120033330112-3132321121020120-1022332330001211-2100332310322003-0030221120131202-1133321211200300-2223201012220131"></a>

#### `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3330100222331232-2121301013312310-1331232232332331-2223220312021131-3123223220132232-1113122122221302-1033102222231112-1322133023122022"></a>

<a id="canonical-2200031232323302-3203010022210322-2203011023310111-0121300332311323-3311133231030013-2023333010311203-3023313210313302-0233330122112310"></a>

#### `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` property

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

<a id="canonical-1110132013220002-2331113232102322-2110302222210111-0122012112202201-0110122013212231-2332301000233221-2130232202130303-0313021031312031"></a>

<a id="canonical-0110103023120133-3203111013030103-0113101022331130-2133201010212212-0320122032132123-1221112220013332-0022211312123220-2320310231003203"></a>

#### `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` property

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

<a id="canonical-2102031113102202-3022333231113202-2311032000102101-0203030322203200-2002313122232020-0130323221203223-2122312320320332-3300330300002132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.use_tls](resources--global_log_receiver--reference--group-005.md#canonical-1032311130223211-1300102100202101-0200123213003120-1021210202000030-0002113312200122-1000220323112123-0100032221032310-1111233222310112)
- [splunk_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-005.md#canonical-0211331211121211-3221211003312312-3003211030203121-0121213303021100-0320111011202003-2100103032101123-3313131301003122-0223111011311131)
- [splunk_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-005.md#canonical-1231103203102133-1122221303110021-3131210322111210-3000131101103222-0221110310122120-1321133312100310-1000323201221013-2300132321100032)
- splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-2230332321023310-1001122100210111-0023120323020220-3232023222202111-0100200001012002-2113112111010221-2110112011012111-1230112012123333"></a>

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

<a id="canonical-2123000323220211-0011131221332312-3313322102121033-3022123300333132-1001232001300302-1300301232130032-2032010321232121-1130003331110302"></a>

### Direct properties for `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info`

<a id="canonical-1021311213001100-0111031210121221-3022023321320312-2102013332222230-3110111131130033-2311323230313322-0200222202232202-2230231220132100"></a>

#### `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2210032320313023-1113021232210033-0331330133221130-2321130001130101-1031113123231210-0201321211323310-3033322131130320-2131220231203221"></a>

<a id="canonical-0300202230033200-0001112012112122-2100103310100202-0012322030300313-2312002120322303-2203032032010120-1202322123321001-1113111030231100"></a>

#### `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` property

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

<a id="canonical-3001101302222020-1100211020222002-1133112021111032-2132310333310102-1233111322201101-2313330203233223-1200333202223332-0331001121001000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.no_ca` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [splunk_receiver](resources--global_log_receiver--reference--group-005.md#canonical-1331121330211320-3112333223033312-2132001103010221-1001332023230013-2111133323303322-0033030200222023-3130031113012002-3332101331003130)
- [splunk_receiver.use_tls](resources--global_log_receiver--reference--group-005.md#canonical-1032311130223211-1300102100202101-0200123213003120-1021210202000030-0002113312200122-1000220323112123-0100032221032310-1111233222310112)
- splunk_receiver.use_tls.no_ca

<a id="canonical-0012022201312133-1133210220001032-1300231332103210-3321310303032110-3000222301233320-0000031010232001-3123132223320121-1002122313322021"></a>

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

<a id="canonical-3133110202111331-3110012132230231-1313311200132200-3313312202001233-0330332232223132-0100032000033300-1111310013112123-0303121102330130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sumo_logic_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- sumo_logic_receiver

<a id="canonical-1113220201232312-0211101312212201-3102312333303113-0100111113320322-0002010113322332-0120002221203122-2300101023133211-1222013031111322"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sumo logic receiver.

Additional upstream details:

Configuration for SumoLogic endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
sumo_logic_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303023231111111-1031020221211322-3101100222230101-3110200131211011-3000222210133100-3333101220102332-0133111233312013-3031033133103300"></a>

### Direct properties for `sumo_logic_receiver`

- [URL](resources--global_log_receiver--reference--group-005.md#canonical-1211232133011100-3311220123122001-2231031332221030-2311333131201230-3030311013102130-1320012132103022-0212200323121201-1213232003112030): complete subsection reference.

<a id="canonical-1211232133011100-3311220123122001-2231031332221030-2311333131201230-3030311013102130-1320012132103022-0212200323121201-1213232003112030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sumo_logic_receiver.url` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [sumo_logic_receiver](resources--global_log_receiver--reference--group-005.md#canonical-3133110202111331-3110012132230231-1313311200132200-3313312202001233-0330332232223132-0100032000033300-1111310013112123-0303121102330130)
- sumo_logic_receiver.URL

<a id="canonical-0331212021213032-3020113200223320-1121232303120211-3120023213233311-0301110331132002-3201200333223301-0211122303000021-2020223330322213"></a>

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
url {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201330322133112-0130031313202203-2001031300211223-3312130000121130-0201233001200323-2300210222032312-3232031332133231-3032332021320102"></a>

### Direct properties for `sumo_logic_receiver.url`

- [blindfold_secret_info](resources--global_log_receiver--reference--group-005.md#canonical-1222300100301200-2230113301113031-0203313023321102-2120212003331332-0213310212322030-1211110122330211-2103213230303110-0022223311222122): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-005.md#canonical-3121032322130121-1230231123003111-3311322132332221-3210310302030210-2213011003012100-0311331110013130-2130103210211231-1201203322103000): complete subsection reference.

<a id="canonical-1222300100301200-2230113301113031-0203313023321102-2120212003331332-0213310212322030-1211110122330211-2103213230303110-0022223311222122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sumo_logic_receiver.url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [sumo_logic_receiver](resources--global_log_receiver--reference--group-005.md#canonical-3133110202111331-3110012132230231-1313311200132200-3313312202001233-0330332232223132-0100032000033300-1111310013112123-0303121102330130)
- [sumo_logic_receiver.url](resources--global_log_receiver--reference--group-005.md#canonical-1211232133011100-3311220123122001-2231031332221030-2311333131201230-3030311013102130-1320012132103022-0212200323121201-1213232003112030)
- sumo_logic_receiver.URL.blindfold_secret_info

<a id="canonical-3010203311220001-1002210211100002-1031031113111332-1222121030213131-2030031132322300-2101300312020313-2001313033001301-3322131120011220"></a>

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

<a id="canonical-0122303122213003-1023203220230333-2031132113312213-3100130303211331-3131213020332033-3222023003122131-3211321200231003-0130002311103221"></a>

### Direct properties for `sumo_logic_receiver.url.blindfold_secret_info`

<a id="canonical-1111130001013120-2332213311000231-3313231031102033-1013103302132130-3022133133120030-2000231033231130-2221032212311230-2232130230232233"></a>

#### `sumo_logic_receiver.url.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3030231301113303-0230220031233230-0002003302121332-2122333310131331-1300100021302010-3100333011330020-3333112130122230-1121020023220022"></a>

<a id="canonical-0310002330100010-1322123222231001-2330302301122023-1132011232121013-3130111310223331-3202310121201222-2221200223133121-2223012320331313"></a>

#### `sumo_logic_receiver.url.blindfold_secret_info.location` property

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

<a id="canonical-3222013330321333-2212200112223103-3121022313300311-2110223113000330-0222232202303221-3001313023112003-0330031133001303-3031200232122321"></a>

<a id="canonical-3012112122323302-3203010000033133-2223033000103010-0101301321220220-0122131002002220-1021121332130313-3033030002213331-3110030023322222"></a>

#### `sumo_logic_receiver.url.blindfold_secret_info.store_provider` property

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

<a id="canonical-3121032322130121-1230231123003111-3311322132332221-3210310302030210-2213011003012100-0311331110013130-2130103210211231-1201203322103000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sumo_logic_receiver.url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [sumo_logic_receiver](resources--global_log_receiver--reference--group-005.md#canonical-3133110202111331-3110012132230231-1313311200132200-3313312202001233-0330332232223132-0100032000033300-1111310013112123-0303121102330130)
- [sumo_logic_receiver.url](resources--global_log_receiver--reference--group-005.md#canonical-1211232133011100-3311220123122001-2231031332221030-2311333131201230-3030311013102130-1320012132103022-0212200323121201-1213232003112030)
- sumo_logic_receiver.URL.clear_secret_info

<a id="canonical-3230133031132200-3013323300021221-0112022130230010-3202231331223233-1020232013320220-1302102100022000-3020202200110211-0120121310313130"></a>

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

<a id="canonical-2311230233203022-2120102310031011-2200033311132302-2333300100031011-1123200210033200-2322122011112302-3102222200112121-1132032001002002"></a>

### Direct properties for `sumo_logic_receiver.url.clear_secret_info`

<a id="canonical-2311220113323202-3332332033030300-1330231021010113-0320111133110010-0112013000310000-3010220131131323-1000313020310201-1200201322000011"></a>

#### `sumo_logic_receiver.url.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2321122321002030-1210232310222233-0110123110130110-1302010200301031-1201130100021020-0300011301033012-2000332231330311-3123203312321002"></a>

<a id="canonical-0322011210201100-3031103320332313-0031222030213100-1331323321030102-3022102102100101-2032133103302033-0031202102311312-2222211110312311"></a>

#### `sumo_logic_receiver.url.clear_secret_info.url` property

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

<a id="canonical-2300030330132202-1012220213322303-1300301203131302-1312202110301333-1300333322313010-2203010102023132-0210122130033122-1033121302122132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- timeouts

<a id="canonical-2323101322201020-2201013121231321-3110201100222113-1023233112330212-2113123032321331-2031001100133233-0232113100210011-1212031323321012"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030220010312322-3101120012303230-3312001101200101-2003022113331220-0313121223131001-3113320112331321-0201121031121333-1113320323310102"></a>

### Direct properties for `timeouts`

<a id="canonical-2130100022000302-0233020112231100-1203103031203200-0221232133223120-1010202210130203-1301002211010230-0133223110303203-0330233303323013"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3121121313100001-0210133021230333-2012103322333331-1311011223032023-3311113113131130-2203312331030031-3233013333221131-2103111212303002"></a>

<a id="canonical-3133322001001321-3233130010301033-0213100202110021-0332233013231232-1311011131230133-1222321001131121-1200301200020330-0211320231322111"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1221022123333101-3003311211233102-1033203310301211-3111100313022012-2231100203331300-2033122101323220-1211312002313233-3001222231021020"></a>

<a id="canonical-1313030310320030-1121132132103311-2222332113300230-2302030002121103-0301030322001100-2200101301233232-1211301212231130-3021021213223132"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0313011203010113-3312002121230013-0210100012300332-3111003220030121-3033102121222311-1321312032232010-3002210210020333-2102131202223102"></a>

<a id="canonical-3030130120321110-2310311232211223-3112233012130133-3033132313333211-1231030211003122-0000000211200301-0202011031030313-2022030020123323"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
