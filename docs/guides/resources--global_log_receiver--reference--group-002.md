---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-1022231331213223-0131003133300123-1131010013321220-3032003300303333-1131133201110313-1222302120301023-3012032113200333-1103132120000012"></a>

## Direct properties — batch / 322202131020 / 3

<a id="canonical-3233300230003111-1002232321120212-1120131320330230-2230011200000220-0310011031311213-1032213212031012-3032301302222211-1033011200102003"></a>

<a id="canonical-3333012302023231-2121211211102132-2012221302013013-2300203322231030-1232220123100312-1010003200013113-1122031130212312-0312221020010110"></a>

## max_bytes property — batch / 322202131020 / 4

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_bytes_disabled](resources--global_log_receiver--reference--group-002.md#canonical-2323232230013213-2102123301122133-0010201111031123-0111301103031030-1031223110311200-3302030203321120-0033113310231123-0123231211021010): complete subsection reference.

<a id="canonical-2023133201123030-2200212013213212-3122333300012111-2220100300033212-0031303300101121-3112302011102000-2312023330311320-0213010330111331"></a>

<a id="canonical-0003231232230310-1330233212213233-0001132311101111-0012013310222210-1320123132301122-1301101231221020-0203102302011201-2121013011231330"></a>

## max_events property — batch / 322202131020 / 5

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_events_disabled](resources--global_log_receiver--reference--group-002.md#canonical-3202130000101301-0130300012332222-0100131320011200-2110203321001303-3001130100020302-0330131233123020-3030110031002102-1333211333132133): complete subsection reference.

<a id="canonical-1133200331220331-2322000033231301-2102112113223033-2122100231002113-3333303112202123-2212003103122112-3231213312012110-1020120211321321"></a>

<a id="canonical-0201302130320123-2233101322310311-3202231213130202-2032232023202030-3211031312122203-2022102200033110-2302230022001321-1001313331012103"></a>

## timeout_seconds property — batch / 322202131020 / 6

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](resources--global_log_receiver--reference--group-002.md#canonical-2332203312010110-3110223100311230-2012010223131123-1313133310123212-3330330133312030-0232001132321132-3130311013011021-2221000211210332): complete subsection reference.

<a id="canonical-3132120020332211-2223331011003233-2323123313332230-1322311310310020-0313111032303121-2231220301223121-1202332232011312-0003132032231203"></a>

## Next pages — batch / 322202131020 / 7

- [azure_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-002.md#canonical-2323232230013213-2102123301122133-0010201111031123-0111301103031030-1031223110311200-3302030203321120-0033113310231123-0123231211021010)
- [azure_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-002.md#canonical-3202130000101301-0130300012332222-0100131320011200-2110203321001303-3001130100020302-0330131233123020-3030110031002102-1333211333132133)
- [azure_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-002.md#canonical-2332203312010110-3110223100311230-2012010223131123-1313133310123212-3330330133312030-0232001132321132-3130311013011021-2221000211210332)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2323232230013213-2102123301122133-0010201111031123-0111301103031030-1031223110311200-3302030203321120-0033113310231123-0123231211021010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202121110021022-3231210322223111-3000232010203003-3312120001011233-3101313232120311-1203233002001032-1200030300310123-1010111030300030"></a>

## azure_receiver.batch.max_bytes_disabled — max_bytes_disabled / 321210220111 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [azure_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-0312232333311231-3011031131000331-3023312102030310-0300331121300221-2331302021110232-1331023203212322-0133100023022312-2002133312101210)
- azure_receiver.batch.max_bytes_disabled

<a id="canonical-2201121010212221-3333130012202230-0310323003321221-3302103331010130-3001122022013112-1231031202112300-2300223100212021-0023330003303031"></a>

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
max_bytes_disabled = {}
```

<a id="canonical-3121210100001011-3021122011101312-1210003112120203-0313220313331032-1113121112222303-1130100230333002-2100022131330312-3301103003013200"></a>

## Direct properties — max_bytes_disabled / 321210220111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233012003011001-2212202110323320-3132001033113303-3202101313300203-0312312123132121-0311201021123122-0121223320203102-1200133221101003"></a>

## Next pages — max_bytes_disabled / 321210220111 / 4

- [azure_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-0312232333311231-3011031131000331-3023312102030310-0300331121300221-2331302021110232-1331023203212322-0133100023022312-2002133312101210)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3202130000101301-0130300012332222-0100131320011200-2110203321001303-3001130100020302-0330131233123020-3030110031002102-1333211333132133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300123230221021-3033000233110221-1323231012312121-3113132332233200-1010103303303202-0221032212331030-2301020313320021-0321000122300331"></a>

## azure_receiver.batch.max_events_disabled — max_events_disabled / 133103113120 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [azure_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-0312232333311231-3011031131000331-3023312102030310-0300331121300221-2331302021110232-1331023203212322-0133100023022312-2002133312101210)
- azure_receiver.batch.max_events_disabled

<a id="canonical-2331113323131313-2212320102112212-2321332232232321-0011112310321133-2003233220210202-3312301323330221-1230002211112211-0222200111013310"></a>

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
max_events_disabled = {}
```

<a id="canonical-2122121223311112-1313220133003213-2221101332220010-1020220120012313-0120213223121321-0331010012002203-2202112010232333-1033300300102221"></a>

## Direct properties — max_events_disabled / 133103113120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103123131101113-0001000311112310-0120030303023330-1013200200121001-1013222213030300-0031220123030300-1222313102212312-0320032331321233"></a>

## Next pages — max_events_disabled / 133103113120 / 4

- [azure_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-0312232333311231-3011031131000331-3023312102030310-0300331121300221-2331302021110232-1331023203212322-0133100023022312-2002133312101210)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2332203312010110-3110223100311230-2012010223131123-1313133310123212-3330330133312030-0232001132321132-3130311013011021-2221000211210332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021212331322211-2033001331031311-2233231112220100-1331100221222202-3230023001323123-2310223032031322-2033031333331003-1102231023223220"></a>

## azure_receiver.batch.timeout_seconds_default — timeout_seconds_default / 321233231110 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [azure_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-0312232333311231-3011031131000331-3023312102030310-0300331121300221-2331302021110232-1331023203212322-0133100023022312-2002133312101210)
- azure_receiver.batch.timeout_seconds_default

<a id="canonical-3010231230212120-3233010321201322-3122222301100211-3332231100012323-3032031103303113-1220231231032311-1030333112211301-2210133100010011"></a>

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
timeout_seconds_default = {}
```

<a id="canonical-3323120130220220-3223202012002021-3110321321100012-3331210320010213-2021013000222322-0130330003210101-3020012101200233-3013331332303131"></a>

## Direct properties — timeout_seconds_default / 321233231110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132332020131300-1123223030201020-3232311130301021-1310101113032210-2123021200133120-1103020322031311-3011302333012200-0300312331203032"></a>

## Next pages — timeout_seconds_default / 321233231110 / 4

- [azure_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-0312232333311231-3011031131000331-3023312102030310-0300331121300221-2331302021110232-1331023203212322-0133100023022312-2002133312101210)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2212231022022020-2232331232100113-0300233311201203-3300213121033120-1133003222102120-1310021311213323-1001131113123113-1002030321000210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311131111322102-1302023133001213-3022301011123113-3323312212001102-0100113122213311-3330100201112202-1123222101313230-2022212320230213"></a>

## azure_receiver.compression — compression / 330213013110 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- azure_receiver.compression

<a id="canonical-1333001023030232-3033021230020012-2200212023020012-1123203330001131-0311311201310213-2113311230120332-3223302210130103-2113032301112023"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Upstream description:

Compression Type.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1203320012320323-0013202233332211-3313323000320312-1311103321332221-3122032331002210-0333133022103310-0333231200330002-0331131130113330"></a>

## Direct properties — compression / 330213013110 / 3

- [compression_default](resources--global_log_receiver--reference--group-002.md#canonical-3232331330012133-0123122303230003-3013223100122003-3121201332122311-3000311003113012-1310010220130000-0223232130321023-0202110010213033): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-002.md#canonical-2030313311323220-0311111231121231-2102300203330013-2121330003000032-2222103331022012-0202020213202220-2111211123033021-2133300111020333): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-002.md#canonical-2231210233121311-3020230230003001-3013013213303130-1213320232232030-3132132023332102-3231232103222210-2013320022030013-1133220333202200): complete subsection reference.

<a id="canonical-0201211112213102-0023003311100012-1221023112003111-3012310211130111-0003101312010201-3321333102120303-0200003330332013-0000012033132100"></a>

## Next pages — compression / 330213013110 / 4

- [azure_receiver.compression.compression_default](resources--global_log_receiver--reference--group-002.md#canonical-3232331330012133-0123122303230003-3013223100122003-3121201332122311-3000311003113012-1310010220130000-0223232130321023-0202110010213033)
- [azure_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-002.md#canonical-2030313311323220-0311111231121231-2102300203330013-2121330003000032-2222103331022012-0202020213202220-2111211123033021-2133300111020333)
- [azure_receiver.compression.compression_none](resources--global_log_receiver--reference--group-002.md#canonical-2231210233121311-3020230230003001-3013013213303130-1213320232232030-3132132023332102-3231232103222210-2013320022030013-1133220333202200)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3232331330012133-0123122303230003-3013223100122003-3121201332122311-3000311003113012-1310010220130000-0223232130321023-0202110010213033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212103211221321-0202200220321013-1200031002320333-2013311311022111-2030112113310132-3200333102021122-1311203022020021-3120232120303311"></a>

## azure_receiver.compression.compression_default — compression_default / 331113102220 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [azure_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-2212231022022020-2232331232100113-0300233311201203-3300213121033120-1133003222102120-1310021311213323-1001131113123113-1002030321000210)
- azure_receiver.compression.compression_default

<a id="canonical-2003030013011233-1011012032210321-1321013022131323-3302031212113022-2120311232021332-2230110333211000-2330031110223031-2302033233123002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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
compression_default = {}
```

<a id="canonical-3332132321230303-3121300101133121-1220111120212101-0130111221333322-3300313221222312-0230003202103231-1013231223011112-0330131202331010"></a>

## Direct properties — compression_default / 331113102220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231021021003333-2221221203132321-1220323130003112-0021001133001233-0121320100000130-0110132233011330-1111213022322321-2022031010113132"></a>

## Next pages — compression_default / 331113102220 / 4

- [azure_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-2212231022022020-2232331232100113-0300233311201203-3300213121033120-1133003222102120-1310021311213323-1001131113123113-1002030321000210)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2030313311323220-0311111231121231-2102300203330013-2121330003000032-2222103331022012-0202020213202220-2111211123033021-2133300111020333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200203302312311-0330210333012230-3102131001332332-1210231021112220-0022323221103132-3131110110022120-3301132023120203-3212033213202020"></a>

## azure_receiver.compression.compression_gzip — compression_gzip / 203032101120 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [azure_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-2212231022022020-2232331232100113-0300233311201203-3300213121033120-1133003222102120-1310021311213323-1001131113123113-1002030321000210)
- azure_receiver.compression.compression_gzip

<a id="canonical-1223021002132132-1111323330133130-0332023101121001-0031210213211222-2133100211013312-0120222320302030-3231203303212120-3113003032130130"></a>

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
compression_gzip = {}
```

<a id="canonical-1233103032322233-0002021003110331-3123023232013121-3103232023222320-1230123002213221-2233003131022100-0131232022330323-0210202210000313"></a>

## Direct properties — compression_gzip / 203032101120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211303133221332-2130100322001313-1122232110213222-3011233330000232-2101000130320222-2320001201210232-3333221312312103-0230300100332333"></a>

## Next pages — compression_gzip / 203032101120 / 4

- [azure_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-2212231022022020-2232331232100113-0300233311201203-3300213121033120-1133003222102120-1310021311213323-1001131113123113-1002030321000210)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2231210233121311-3020230230003001-3013013213303130-1213320232232030-3132132023332102-3231232103222210-2013320022030013-1133220333202200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212232123321202-2332333311303223-0221123003222023-3110122202313223-3101301310011003-0132001022322032-1213012121111332-2132333332320022"></a>

## azure_receiver.compression.compression_none — compression_none / 203030211123 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [azure_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-2212231022022020-2232331232100113-0300233311201203-3300213121033120-1133003222102120-1310021311213323-1001131113123113-1002030321000210)
- azure_receiver.compression.compression_none

<a id="canonical-0322011332110320-3020030301102330-0300111213313323-1020300133023301-0032102330303013-3302131221110333-2102130313301233-0131310100110012"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression none.

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
compression_none = {}
```

<a id="canonical-1320312300101233-1230002132123121-1202310200131132-1323013201103301-3322301201123222-2123112111010103-3200103002202033-2123311231102112"></a>

## Direct properties — compression_none / 203030211123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302132122113200-2331321032200100-2310133303110002-1131130122130021-2301302203030002-2201000000322321-3013313201331311-2030102321311310"></a>

## Next pages — compression_none / 203030211123 / 4

- [azure_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-2212231022022020-2232331232100113-0300233311201203-3300213121033120-1133003222102120-1310021311213323-1001131113123113-1002030321000210)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3321122021101233-0323300002001002-2212120302013221-2231302210203201-3121310031113020-0111200312123321-2120030012123003-0212312323022131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320202032033110-0132113111100120-2333031203300221-2233002320033212-3330332103113221-2332003311030013-1113211020013230-2103030023123131"></a>

## azure_receiver.connection_string — connection_string / 102331032301 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- azure_receiver.connection_string

<a id="canonical-3310012233002333-0122120112212320-2011103110201122-3222023032132113-1220113331113220-3331120211330113-1010121013010132-2012031002301020"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
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
connection_string {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132113331213020-3322131232133123-3123032330030311-0113101211202011-1010320011121020-1311321232331101-2020322031100313-3110110031331312"></a>

## Direct properties — connection_string / 102331032301 / 3

- [blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-3301301211103100-3023131033302123-2100020223130301-1320000300200223-3312021100331312-2333202300231031-2023202001022032-1232020103330131): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-0133302123223211-2203020213301031-1000122113023013-1233312131032003-0301112011103022-0002313213131020-3113031231233023-2313102210222302): complete subsection reference.

<a id="canonical-2310013203110032-0130333100111300-0330313112132302-0221200131311010-1030013002323022-0133230201331230-1100313331002020-2220031010132311"></a>

## Next pages — connection_string / 102331032301 / 4

- [azure_receiver.connection_string.blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-3301301211103100-3023131033302123-2100020223130301-1320000300200223-3312021100331312-2333202300231031-2023202001022032-1232020103330131)
- [azure_receiver.connection_string.clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-0133302123223211-2203020213301031-1000122113023013-1233312131032003-0301112011103022-0002313213131020-3113031231233023-2313102210222302)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3301301211103100-3023131033302123-2100020223130301-1320000300200223-3312021100331312-2333202300231031-2023202001022032-1232020103330131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130302132123010-0213301313112113-2231301131113131-3110023300001211-2102000020011133-1030030301013002-0320001232112210-0132032030300332"></a>

## azure_receiver.connection_string.blindfold_secret_info — blindfold_secret_info / 320323003313 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [azure_receiver.connection_string](resources--global_log_receiver--reference--group-002.md#canonical-3321122021101233-0323300002001002-2212120302013221-2231302210203201-3121310031113020-0111200312123321-2120030012123003-0212312323022131)
- azure_receiver.connection_string.blindfold_secret_info

<a id="canonical-2221332330013210-2212303002000303-3033121011303301-1030333213220221-2100000312303010-1003332212223322-1302031311322332-2312133133321220"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2311031311030022-0122301231031320-2333300011331233-1232201101123213-0031010301230112-3113110113223303-0312233322032313-1211200202030231"></a>

## Direct properties — blindfold_secret_info / 320323003313 / 3

<a id="canonical-1332113000102211-1030032212232202-3300011232333200-0130320230323323-3031023020221231-0012031222222101-0331030200231213-1010020213100013"></a>

<a id="canonical-2123212212023112-0002323201322022-1001130113123320-2331223001221321-3023201311221300-2213233112113302-3012023322030233-3002122301302113"></a>

## decryption_provider property — blindfold_secret_info / 320323003313 / 4

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

<a id="canonical-0201102231220210-3230032021131022-3330011203002113-1323322331300101-0310122203112311-0033210302001121-1301030032121201-2331301330031031"></a>

<a id="canonical-2232220331021131-0111031103010313-0230310101110012-0032200231031003-1233033023033003-1010223102122312-3011031112230000-3023331033113031"></a>

## location property — blindfold_secret_info / 320323003313 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3202200210213230-2110103323333303-1000121213312333-0111310303030101-1220221212211023-0301120322000213-2001300023233111-3233211313322320"></a>

<a id="canonical-2100331230010021-0122331031330022-1023321212011120-2023000213311332-0233200331001313-0203201103210123-1002301203022120-0111011130011121"></a>

## store_provider property — blindfold_secret_info / 320323003313 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-0130113110111022-1002221103011303-2213031313121120-2120001222231002-2023313212033200-0110132123133022-2203101101002332-3210323230121301"></a>

## Next pages — blindfold_secret_info / 320323003313 / 7

- [azure_receiver.connection_string](resources--global_log_receiver--reference--group-002.md#canonical-3321122021101233-0323300002001002-2212120302013221-2231302210203201-3121310031113020-0111200312123321-2120030012123003-0212312323022131)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0133302123223211-2203020213301031-1000122113023013-1233312131032003-0301112011103022-0002313213131020-3113031231233023-2313102210222302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030111021010203-3133000311013332-1233121012100121-3112122012113131-3200000323121021-3311311101203121-3020023103122130-2011303001230130"></a>

## azure_receiver.connection_string.clear_secret_info — clear_secret_info / 102311030321 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [azure_receiver.connection_string](resources--global_log_receiver--reference--group-002.md#canonical-3321122021101233-0323300002001002-2212120302013221-2231302210203201-3121310031113020-0111200312123321-2120030012123003-0212312323022131)
- azure_receiver.connection_string.clear_secret_info

<a id="canonical-1120321123300223-0331213203033300-3111301311121132-2022213223112030-1330231331001130-0131023233102013-3111201200010021-0130103023213220"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2030301013020301-0302120010101032-2020131302232322-1322130231330203-0000213313010011-3332103323033211-3322013200000310-3023022333002130"></a>

## Direct properties — clear_secret_info / 102311030321 / 3

<a id="canonical-3232212030133111-0201122001203111-2111331210031313-0320102000130232-3003002301231302-0011121031203110-1312203112132302-0320103301311212"></a>

<a id="canonical-1131110300010310-3322131120003030-1121230331031121-3123032302012332-1310311301031031-3222233210103333-2021233301320203-1123121311123033"></a>

## provider_ref property — clear_secret_info / 102311030321 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2123113022133202-3203332330103031-0030232122232132-3303330012123310-1232102122001310-2233313111101102-2303211332022101-2201233322101303"></a>

<a id="canonical-2331312130200120-2323031332001202-0223310332310011-3133231122222213-0120301201130331-3001302033223123-3113201300031323-1010100003322230"></a>

## URL property — clear_secret_info / 102311030321 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2221123231011022-0230121232111112-2230213111123303-1122001012203312-1103221020202003-1321303231021212-0122331221113023-2211203030123212"></a>

## Next pages — clear_secret_info / 102311030321 / 6

- [azure_receiver.connection_string](resources--global_log_receiver--reference--group-002.md#canonical-3321122021101233-0323300002001002-2212120302013221-2231302210203201-3121310031113020-0111200312123321-2120030012123003-0212312323022131)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3110031023120033-1213230103132203-2110113022103201-0132031102313032-0222333230033201-2320000331022303-3233031321002010-2110103103010001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103031133111310-3103122030131131-2102122001020330-0312320333033100-3221221210323002-0203030002021112-0003003123013030-1302221033203112"></a>

## azure_receiver.filename_options — filename_options / 101201203133 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- azure_receiver.filename_options

<a id="canonical-0311023201301101-1132312022231232-2103120011132102-0102011221233030-2112000210012131-3100333013323222-1231123302313012-2032013220322111"></a>

Type: `"object"`. single nested block, Optional.

Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint
bucket or file.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3132213331332033-1012330232000313-3110002322211301-2331023132131311-0211133223333233-0013100223202103-2013123321332131-2220302000121332"></a>

## Direct properties — filename_options / 101201203133 / 3

<a id="canonical-2221130031202331-3220332121302101-2101010001123320-2231331201310210-2123033003333300-3211222002110222-2210233313303303-3023123223312330"></a>

<a id="canonical-2303233201320301-2303330123022113-0313100013010033-1221301211321113-1122121103301102-1302010030333132-1312111122211221-0221333310311210"></a>

## custom_folder property — filename_options / 101201203133 / 4

Type: `"string"`. Optional.

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match.

Upstream description:

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match \`/^\[a-z\_\]\[a-z0-9\\\\-\\\\.\_\]\*$/i\`

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1231321113001123-1221332000120111-1332120101313223-3101203332013311-0222200001230203-0332011133222113-1303232221102130-1320123303012133"></a>

## Next pages — filename_options / 101201203133 / 5

- [azure_receiver.filename_options.log_type_folder](resources--global_log_receiver--reference--group-002.md#canonical-2313133202312310-3003023111113031-1311332220021230-3012030101123323-1312022021301302-3010032220132011-2333133101221310-3331211233213210)
- [azure_receiver.filename_options.no_folder](resources--global_log_receiver--reference--group-002.md#canonical-2220212033330103-3110031030300003-1113333313102121-2212122232031001-2102133320130321-2333130321333031-2322122120333310-0022002100200233)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2313133202312310-3003023111113031-1311332220021230-3012030101123323-1312022021301302-3010032220132011-2333133101221310-3331211233213210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301333130212011-1320333120213020-3211122011232021-1133103210123312-0230310003113210-3021002310201321-0111211300300020-3001022030101012"></a>

## azure_receiver.filename_options.log_type_folder — log_type_folder / 122111123201 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [azure_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-3110031023120033-1213230103132203-2110113022103201-0132031102313032-0222333230033201-2320000331022303-3233031321002010-2110103103010001)
- azure_receiver.filename_options.log_type_folder

<a id="canonical-3302233113221300-0012021231120122-3130301002220321-0310331301321303-2113232110123023-3313222301031023-3201212111032130-2122033121002020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for log type folder.

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
log_type_folder = {}
```

<a id="canonical-0000233333331102-1103133000313313-0332310103110001-2002313032102113-0023330113221311-3300023103022101-2220130130233230-0031213303021301"></a>

## Direct properties — log_type_folder / 122111123201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220212100223232-3133302303220203-1323211203232230-1223003331303132-1233102331001011-3003031032032020-3222002211223201-0021232332221013"></a>

## Next pages — log_type_folder / 122111123201 / 4

- [azure_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-3110031023120033-1213230103132203-2110113022103201-0132031102313032-0222333230033201-2320000331022303-3233031321002010-2110103103010001)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2220212033330103-3110031030300003-1113333313102121-2212122232031001-2102133320130321-2333130321333031-2322122120333310-0022002100200233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132331011132320-0211212001302000-2002222130313331-3022111002222333-1313313130212223-1100030232120021-2230121133121223-1103111233220032"></a>

## azure_receiver.filename_options.no_folder — no_folder / 013332031022 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031)
- [azure_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-3110031023120033-1213230103132203-2110113022103201-0132031102313032-0222333230033201-2320000331022303-3233031321002010-2110103103010001)
- azure_receiver.filename_options.no_folder

<a id="canonical-3001030000033221-2033211000100320-1211012312212131-0023313313201022-0222112323203011-3223331011200320-0100120221321310-1002231001332121"></a>

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
no_folder = {}
```

<a id="canonical-0222213021002101-2311202113331133-1331311202312033-0210311232203030-3023101122212322-1321110330012121-3111102310001203-0332302123113010"></a>

## Direct properties — no_folder / 013332031022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011322131330301-1111201211021303-3002301313131121-0010300120231110-2333210122012200-2212112001001133-0332033220021002-3222103200231021"></a>

## Next pages — no_folder / 013332031022 / 4

- [azure_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-3110031023120033-1213230103132203-2110113022103201-0132031102313032-0222333230033201-2320000331022303-3233031321002010-2110103103010001)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112313220222300-2331013012000112-3330330233021211-3011121320333100-1212130013130111-1033231031223221-3333012113212321-3310223321330003"></a>

## datadog_receiver — datadog_receiver / 300320213200 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- datadog_receiver

<a id="canonical-0003033200133213-2021130011102033-2321233212202033-0010010332230103-2101001230301213-0313001301111103-3112321322103202-2021022012211332"></a>

Type: `"object"`. single nested block, Optional.

Datadog Configuration. Configuration for Datadog endpoint.

Upstream description:

Configuration for Datadog endpoint.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2332321022030002-3013323113100321-1322010011011330-1110323101033130-1001203300100231-0003300312232210-1303003011230211-2210233220012302"></a>

## Direct properties — datadog_receiver / 300320213200 / 3

- [batch](resources--global_log_receiver--reference--group-002.md#canonical-2223332013002130-2132203020333313-0300211311222212-1321133023230031-3331032131001210-2101021213200223-2320110123000111-0020112013112230): complete subsection reference.

- [compression](resources--global_log_receiver--reference--group-002.md#canonical-3030230332320233-2330011003211132-0112312123001121-1301000021010102-0101130200212121-3200321213211330-0133002130323203-3111221213133220): complete subsection reference.

- [datadog_api_key](resources--global_log_receiver--reference--group-002.md#canonical-3003331032300003-3113111102303300-0022111111131031-2220230330110113-1120122230101110-2030202100202320-3232330113023022-0231001111213310): complete subsection reference.

<a id="canonical-2121132022131301-1021220221122333-3102223320122020-2032212120002232-1301010303210303-2100300311231021-3031223321200231-1202100221321011"></a>

<a id="canonical-2132220101122103-3110020130012211-3031332031103312-1031001112320020-1123220222003121-3203312311100312-1220012230130302-0311030301010311"></a>

## endpoint property — datadog_receiver / 300320213200 / 4

Type: `"string"`. Optional.

Exclusive with \[site\] Datadog Endpoint,.

Upstream description:

Exclusive with \[site\] Datadog Endpoint,.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2330132020310013-0111213331032001-3103033013010031-1010212210021221-2212020333131133-3202031121030121-0011310121320031-1212201203213011"></a>

## site property — datadog_receiver / 300320213200 / 5

Type: `"string"`. Optional.

Exclusive with \[endpoint\] Datadog Site,.

Upstream description:

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
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  }
}
```

- [use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323): complete subsection reference.

<a id="canonical-3101212022131230-0102223231312210-3123133121030221-3232301112302310-0002103331222011-0212230220101310-2113231210230012-1130200010333303"></a>

## Next pages — datadog_receiver / 300320213200 / 6

- [datadog_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-2223332013002130-2132203020333313-0300211311222212-1321133023230031-3331032131001210-2101021213200223-2320110123000111-0020112013112230)
- [datadog_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-3030230332320233-2330011003211132-0112312123001121-1301000021010102-0101130200212121-3200321213211330-0133002130323203-3111221213133220)
- [datadog_receiver.datadog_api_key](resources--global_log_receiver--reference--group-002.md#canonical-3003331032300003-3113111102303300-0022111111131031-2220230330110113-1120122230101110-2030202100202320-3232330113023022-0231001111213310)
- [datadog_receiver.no_tls](resources--global_log_receiver--reference--group-002.md#canonical-3013212312131302-2022221332103200-2120320302301030-0202201232233030-0110233213023302-2330201031221231-3030111002013123-1020323020201132)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2223332013002130-2132203020333313-0300211311222212-1321133023230031-3331032131001210-2101021213200223-2320110123000111-0020112013112230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223311210113201-2122030021131022-3312013021331023-0101003210123230-3002021213103313-1211122321313231-0131301103300110-2030322321000300"></a>

## datadog_receiver.batch — batch / 030133030301 / 2

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

<a id="canonical-2130313001201230-2003210221200003-1320300023221013-0013332002000321-3212023301102201-0200103321112122-2120121011303222-2123211021231200"></a>

## Direct properties — batch / 030133030301 / 3

<a id="canonical-3311123322330313-3203103233122202-0111200331123310-0000122202122003-0322221030210001-2122121311220232-0303321330022213-3132020330330201"></a>

<a id="canonical-2110330131111001-2321021011123021-1220321312223233-0120100120011312-3211330233103310-0232333013230110-3231322330133031-2203032330010213"></a>

## max_bytes property — batch / 030133030301 / 4

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0121302120113220-3330103101110021-0311213332112021-2330113133223212-3002013232233330-3223220233100322-3101202233103213-2132023331231033"></a>

## max_events property — batch / 030133030301 / 5

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0100103003113133-1212100201223211-1130013332310102-3030020323133302-3102030111223120-2203212310213003-2313033101003033-0000212123320120"></a>

## timeout_seconds property — batch / 030133030301 / 6

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

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

<a id="canonical-3202310331231102-3201120123221223-3300103110332210-3103111322232201-0203032312131331-0012122313320013-3332111320033003-1121311103231200"></a>

## Next pages — batch / 030133030301 / 7

- [datadog_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-002.md#canonical-3021020103021100-2032012202002331-1221212303102301-0001121023002033-2301310330113130-1200303111330100-3201203110102322-3103201222322102)
- [datadog_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-002.md#canonical-1333032111100131-3002112200203012-0330002201331112-3031332322030302-2102200233113122-1233021032112102-0200201331233201-2103112313320000)
- [datadog_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-002.md#canonical-1013330000232023-0233320033100123-0003002111023113-1031001102301313-0303021210033211-2130323213012011-1002002203130120-3232202023312311)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3021020103021100-2032012202002331-1221212303102301-0001121023002033-2301310330113130-1200303111330100-3201203110102322-3103201222322102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011030202100011-2213220210122112-2112212100213202-0103020022212230-1030320222213101-0001313103200102-1022001022322110-0212303321131020"></a>

## datadog_receiver.batch.max_bytes_disabled — max_bytes_disabled / 221210311213 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-2223332013002130-2132203020333313-0300211311222212-1321133023230031-3331032131001210-2101021213200223-2320110123000111-0020112013112230)
- datadog_receiver.batch.max_bytes_disabled

<a id="canonical-0111213032122002-0212231001232223-2311101322323113-0333031230023033-2130300030102110-1213002130212313-1221312130003332-1322223102223202"></a>

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
max_bytes_disabled = {}
```

<a id="canonical-2332001133311100-0210321301300230-0313011323000321-3321220001322332-1220121321333220-3022313311122023-2100201200123201-3101321311122011"></a>

## Direct properties — max_bytes_disabled / 221210311213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303100120311032-2000031122113310-0322030011100211-3201001032113303-0013322021323311-2002211312310322-2212200331030123-2333123303212313"></a>

## Next pages — max_bytes_disabled / 221210311213 / 4

- [datadog_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-2223332013002130-2132203020333313-0300211311222212-1321133023230031-3331032131001210-2101021213200223-2320110123000111-0020112013112230)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1333032111100131-3002112200203012-0330002201331112-3031332322030302-2102200233113122-1233021032112102-0200201331233201-2103112313320000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220110132003213-1032313222331122-3003020311202012-2232111330211130-0112220031231022-1313210100213132-1130321102313030-1322213233112102"></a>

## datadog_receiver.batch.max_events_disabled — max_events_disabled / 211221322302 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-2223332013002130-2132203020333313-0300211311222212-1321133023230031-3331032131001210-2101021213200223-2320110123000111-0020112013112230)
- datadog_receiver.batch.max_events_disabled

<a id="canonical-2000031012002030-2103300321130003-0320323122013100-3233102330022232-2232203011122111-0003123212201030-0211222001012103-3211312131332030"></a>

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
max_events_disabled = {}
```

<a id="canonical-1311013100131331-2011022102303003-2031231131323002-1222311101002032-1321121323200322-0011132221222222-0022131102202031-2102322102202013"></a>

## Direct properties — max_events_disabled / 211221322302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121030001211132-2113322031323023-0233203130123031-3110332030330333-0123033321000123-2013322301132300-0220231001230310-1220232023223132"></a>

## Next pages — max_events_disabled / 211221322302 / 4

- [datadog_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-2223332013002130-2132203020333313-0300211311222212-1321133023230031-3331032131001210-2101021213200223-2320110123000111-0020112013112230)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1013330000232023-0233320033100123-0003002111023113-1031001102301313-0303021210033211-2130323213012011-1002002203130120-3232202023312311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201310211110122-3222110122023323-0330133132233331-1211122220101312-0133321221312023-2112031331330323-3023022232012133-1300113221323113"></a>

## datadog_receiver.batch.timeout_seconds_default — timeout_seconds_default / 322000300030 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-2223332013002130-2132203020333313-0300211311222212-1321133023230031-3331032131001210-2101021213200223-2320110123000111-0020112013112230)
- datadog_receiver.batch.timeout_seconds_default

<a id="canonical-1222112302032112-2323321012111203-0333311321103031-0120123201103310-3113220231311303-2201331013212030-3330131032032113-3221301032331210"></a>

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
timeout_seconds_default = {}
```

<a id="canonical-0011030322201103-1133100230211311-3221211021320133-3030202321213212-2213023330232131-3310013100123033-2121033113132212-1333102321123102"></a>

## Direct properties — timeout_seconds_default / 322000300030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121203202032013-2011133120001233-2223202223320000-3020121022200033-1232200022122031-1220303321322232-3110231022311300-1223012001301313"></a>

## Next pages — timeout_seconds_default / 322000300030 / 4

- [datadog_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-2223332013002130-2132203020333313-0300211311222212-1321133023230031-3331032131001210-2101021213200223-2320110123000111-0020112013112230)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3030230332320233-2330011003211132-0112312123001121-1301000021010102-0101130200212121-3200321213211330-0133002130323203-3111221213133220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111201301130100-3101313330330231-0311311122330210-1322320131331230-3003000201332310-3310100023132300-2102300221003311-1203120323110210"></a>

## datadog_receiver.compression — compression / 212322210312 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- datadog_receiver.compression

<a id="canonical-0100020323133330-2021232133002231-1033032121301031-2323111023313303-3312120101223112-0101032021001221-3231210211100323-2102303333030003"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Upstream description:

Compression Type.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2220231200223112-2232221010200212-1012101002211110-0201122011323210-1022233030010331-3303301113202312-0230330230321002-1333020111120112"></a>

## Direct properties — compression / 212322210312 / 3

- [compression_default](resources--global_log_receiver--reference--group-002.md#canonical-2013020312120222-1032312300013001-3303231303303130-1132303033303310-3102212001301320-1130033031111110-3303233202033313-3202100131003100): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-002.md#canonical-0202203113122102-0211002230030313-2310130203222301-1232200301310130-3301030033130330-1230131311020100-3213302031311030-3001300301020233): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-002.md#canonical-0023110002100010-3130000210023222-0203222110322211-1311121303211120-1303301332022231-0303213130323110-1111230100223000-1101031023002321): complete subsection reference.

<a id="canonical-3120320130101321-0022110232223312-0331302010323021-0101200023223030-2231031001132111-0130010012021022-2201032202210011-1032210333230133"></a>

## Next pages — compression / 212322210312 / 4

- [datadog_receiver.compression.compression_default](resources--global_log_receiver--reference--group-002.md#canonical-2013020312120222-1032312300013001-3303231303303130-1132303033303310-3102212001301320-1130033031111110-3303233202033313-3202100131003100)
- [datadog_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-002.md#canonical-0202203113122102-0211002230030313-2310130203222301-1232200301310130-3301030033130330-1230131311020100-3213302031311030-3001300301020233)
- [datadog_receiver.compression.compression_none](resources--global_log_receiver--reference--group-002.md#canonical-0023110002100010-3130000210023222-0203222110322211-1311121303211120-1303301332022231-0303213130323110-1111230100223000-1101031023002321)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2013020312120222-1032312300013001-3303231303303130-1132303033303310-3102212001301320-1130033031111110-3303233202033313-3202100131003100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123322222021320-2213203112021230-3022303101022213-1031233111100110-0233210323313230-1113100303223123-3103202310203031-0120000302332313"></a>

## datadog_receiver.compression.compression_default — compression_default / 013023100021 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-3030230332320233-2330011003211132-0112312123001121-1301000021010102-0101130200212121-3200321213211330-0133002130323203-3111221213133220)
- datadog_receiver.compression.compression_default

<a id="canonical-3200131330223011-2220311302123323-1310310200203210-1202121210122333-2202123113301122-0311002200102130-1032323132010021-2222331132203000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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
compression_default = {}
```

<a id="canonical-1231102213023200-0232101130222312-3110010221310010-1132032121000202-1102102102123112-3021231123322201-1301121131111333-0132310223100322"></a>

## Direct properties — compression_default / 013023100021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213131332113032-3030211303333132-2323013032221001-3123220303323023-3003312311100023-1012022002020023-0102212203312133-1121321130233010"></a>

## Next pages — compression_default / 013023100021 / 4

- [datadog_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-3030230332320233-2330011003211132-0112312123001121-1301000021010102-0101130200212121-3200321213211330-0133002130323203-3111221213133220)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0202203113122102-0211002230030313-2310130203222301-1232200301310130-3301030033130330-1230131311020100-3213302031311030-3001300301020233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302023102323301-2012310220131003-2332003320322003-0103331103231220-1300232131123022-2132113203022133-0033103321333110-0122213011103130"></a>

## datadog_receiver.compression.compression_gzip — compression_gzip / 212113023131 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-3030230332320233-2330011003211132-0112312123001121-1301000021010102-0101130200212121-3200321213211330-0133002130323203-3111221213133220)
- datadog_receiver.compression.compression_gzip

<a id="canonical-1012031012132023-1230031221212030-0022331120022302-2312301031323102-2332121203010200-0233210203320123-1123000230230201-2022023322012323"></a>

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
compression_gzip = {}
```

<a id="canonical-3010330012223011-1332213112323202-0030000010222210-3303201122310321-0233113203003001-3201220200121121-1131020213130132-2233201333323012"></a>

## Direct properties — compression_gzip / 212113023131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310101203333103-2330031223332022-3021230131123232-2021321101300112-1132222131003222-1130222021312000-0132330102303330-2102103301332300"></a>

## Next pages — compression_gzip / 212113023131 / 4

- [datadog_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-3030230332320233-2330011003211132-0112312123001121-1301000021010102-0101130200212121-3200321213211330-0133002130323203-3111221213133220)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0023110002100010-3130000210023222-0203222110322211-1311121303211120-1303301332022231-0303213130323110-1111230100223000-1101031023002321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300011122123030-3032212232022022-1003313032322201-1311233302000203-0220220301331301-0210220212130022-0133010331130313-0300322101223202"></a>

## datadog_receiver.compression.compression_none — compression_none / 203131203300 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-3030230332320233-2330011003211132-0112312123001121-1301000021010102-0101130200212121-3200321213211330-0133002130323203-3111221213133220)
- datadog_receiver.compression.compression_none

<a id="canonical-1101132011322000-2131323330200310-1221031313131223-1023312231010022-0223232111332313-1112112233301010-0122100211220311-0313300220312013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression none.

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
compression_none = {}
```

<a id="canonical-3100330123213120-0101320111001033-3311131001233332-3300331310212212-2121220301302033-2231232021111233-1000101231203222-0132133301112303"></a>

## Direct properties — compression_none / 203131203300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003003320012131-0331320020011221-2330102023302322-3303321303300312-2332211131222001-0233313111112303-3232322333330123-1103001111120211"></a>

## Next pages — compression_none / 203131203300 / 4

- [datadog_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-3030230332320233-2330011003211132-0112312123001121-1301000021010102-0101130200212121-3200321213211330-0133002130323203-3111221213133220)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3003331032300003-3113111102303300-0022111111131031-2220230330110113-1120122230101110-2030202100202320-3232330113023022-0231001111213310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013231031013011-0133333201112031-3311023312210031-1033112200112121-0123103302220212-0310331012130120-3231031120101122-0232032121003111"></a>

## datadog_receiver.datadog_api_key — datadog_api_key / 012200303100 / 2

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

<a id="canonical-1330132321211311-1231022131033233-2303201213003013-1123221332001002-0131030303323013-3232200102121112-1001011303002122-1020203030110020"></a>

## Direct properties — datadog_api_key / 012200303100 / 3

- [blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-2103202232221111-0102221310212202-0020012111032131-1121101220022213-1101303001122303-2200231302313332-0200303320002322-2203003332312231): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-3212100202100303-0011101223230130-0321000232102223-3011203321032330-2200303231212210-1011322012032201-2033103022013202-0011322132130313): complete subsection reference.

<a id="canonical-3210113001021310-3301203233313322-0030212033133311-3203110132300301-0221332330100221-1132322001330011-3331221011101010-1010122110122011"></a>

## Next pages — datadog_api_key / 012200303100 / 4

- [datadog_receiver.datadog_api_key.blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-2103202232221111-0102221310212202-0020012111032131-1121101220022213-1101303001122303-2200231302313332-0200303320002322-2203003332312231)
- [datadog_receiver.datadog_api_key.clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-3212100202100303-0011101223230130-0321000232102223-3011203321032330-2200303231212210-1011322012032201-2033103022013202-0011322132130313)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2103202232221111-0102221310212202-0020012111032131-1121101220022213-1101303001122303-2200231302313332-0200303320002322-2203003332312231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003332013301200-2131121011121302-2101300310011231-0223021210021031-0022210002323210-0211132010330302-2220230000313031-2122011101002000"></a>

## datadog_receiver.datadog_api_key.blindfold_secret_info — blindfold_secret_info / 223010030233 / 2

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

<a id="canonical-0301022311333131-3300122120213020-1110031210323233-3032002013301032-2212002322320210-1311233223231201-2032122021020121-3021123020123220"></a>

## Direct properties — blindfold_secret_info / 223010030233 / 3

<a id="canonical-2110220013220200-2313232110030021-3131110210130320-0333213010102133-0310030213320013-3331323211301003-2312321231123201-3123202131020320"></a>

<a id="canonical-1311213220330230-1233021201002222-3132222032123022-0223323331112101-3010311233002102-1131013223121301-1212222112210101-3120022231310111"></a>

## decryption_provider property — blindfold_secret_info / 223010030233 / 4

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

<a id="canonical-3020131131212121-3221322033002313-0022311120231303-1223221133331101-2332021301332202-2331210031332311-0320021223320031-0200012020311320"></a>

<a id="canonical-2133123223100122-0110133013332312-2011010103031012-3021122231123312-0032102032022031-1030213332201102-3131323202000210-1032311123202111"></a>

## location property — blindfold_secret_info / 223010030233 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0321030202313000-0003020210112113-3230333012032130-1131332110011100-1131221210322001-1302013333232113-2111213102113303-2312132033202300"></a>

## store_provider property — blindfold_secret_info / 223010030233 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-3002121110210122-0133132021113033-3031203332233133-3103201322122023-0121230213210313-3032031120302111-0123200131330200-2122313113010200"></a>

## Next pages — blindfold_secret_info / 223010030233 / 7

- [datadog_receiver.datadog_api_key](resources--global_log_receiver--reference--group-002.md#canonical-3003331032300003-3113111102303300-0022111111131031-2220230330110113-1120122230101110-2030202100202320-3232330113023022-0231001111213310)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3212100202100303-0011101223230130-0321000232102223-3011203321032330-2200303231212210-1011322012032201-2033103022013202-0011322132130313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202231122002033-2212300023112211-3231132133001202-1320322311211210-3211033310301002-0022220222113321-2321120001130133-3131030023113213"></a>

## datadog_receiver.datadog_api_key.clear_secret_info — clear_secret_info / 312323122023 / 2

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

<a id="canonical-2033100232030103-0111111133220201-1230122002222102-3331211133220132-0300220101220333-3032000010012301-2111020003110032-2032220203202300"></a>

## Direct properties — clear_secret_info / 312323122023 / 3

<a id="canonical-2212222301022122-1022000120020000-0131221223313001-0133122110201310-0302113233012032-1100322033001113-2102330113232033-3330031010310232"></a>

<a id="canonical-2312221321130323-2122232313132220-0110223232020031-3221031333310232-1333230310322210-3303030100021003-3302310102012223-2321320331122032"></a>

## provider_ref property — clear_secret_info / 312323122023 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3210031323132101-3300021330003110-3300330110031113-3320003322211223-3210111230332033-2133211313002023-1031202301103121-0000213111331222"></a>

<a id="canonical-1213223300320312-3300213123113201-2000313333031000-1101121233001132-0100213300101330-3101233232331310-3100021122132110-2222311323222011"></a>

## URL property — clear_secret_info / 312323122023 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1121331033112102-1003021011011312-0000201101301120-0102231221022031-1322230110031031-0230303330311232-1023223222130013-0112010033022311"></a>

## Next pages — clear_secret_info / 312323122023 / 6

- [datadog_receiver.datadog_api_key](resources--global_log_receiver--reference--group-002.md#canonical-3003331032300003-3113111102303300-0022111111131031-2220230330110113-1120122230101110-2030202100202320-3232330113023022-0231001111213310)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3013212312131302-2022221332103200-2120320302301030-0202201232233030-0110233213023302-2330201031221231-3030111002013123-1020323020201132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313202122010022-3110233312211232-2121323033202231-0213113011011002-0331210222332210-0020123212221230-3320300333030010-0200301221223303"></a>

## datadog_receiver.no_tls — no_tls / 010030013233 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- datadog_receiver.no_tls

<a id="canonical-2323030200323112-3112102121002321-2003201100321203-2101112010003303-2120203103130300-2210110123231110-0211301022233112-1002102330021000"></a>

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
no_tls = {}
```

<a id="canonical-2231011230220112-1010302213231113-0310330222012331-3231202232311330-2320031032233333-2202232210021121-3032113220001111-1221233311210023"></a>

## Direct properties — no_tls / 010030013233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202221100312233-1233012023020032-3201300013323020-3320331232021021-0312113302110113-1111130122030201-2212122230213200-1331231033001221"></a>

## Next pages — no_tls / 010030013233 / 4

- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133000303320013-0232232231333010-3020103011130231-0302130312012213-2112321122021222-2221211201101110-0013323011221331-3102131100310123"></a>

## datadog_receiver.use_tls — use_tls / 301313330111 / 2

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

<a id="canonical-3311312212223332-1311301313022202-0102202123322031-1222132210221202-1023333212302223-0330223301101103-0002320300322210-3021021233220101"></a>

## Direct properties — use_tls / 301313330111 / 3

- [disable_verify_certificate](resources--global_log_receiver--reference--group-002.md#canonical-2010023220221312-3111030320011000-2110200032003013-3202303233322013-1111312211203111-3220301131302110-1322231223311310-0312203311011213): complete subsection reference.

- [disable_verify_hostname](resources--global_log_receiver--reference--group-002.md#canonical-0110020111321033-2210111001012202-2332123203013331-3322332300222211-0122313323220002-3023320112123321-2023233210311302-1113221013133020): complete subsection reference.

- [enable_verify_certificate](resources--global_log_receiver--reference--group-002.md#canonical-0133011203031313-2230213023320331-1122200103123131-0103021310331220-1002103230211233-1103303031200300-0031110222313301-0230123001021121): complete subsection reference.

- [enable_verify_hostname](resources--global_log_receiver--reference--group-002.md#canonical-0031022130230301-1033031021021113-2330122311023233-0032121330133130-3033101001020231-1012103111232122-0232211322111331-3133222022230000): complete subsection reference.

- [mtls_disabled](resources--global_log_receiver--reference--group-002.md#canonical-2021031131131012-0111031111133000-3113311032202112-3331130002222011-1230231101013232-2320133011101132-1003331033201033-1333121201032332): complete subsection reference.

- [mtls_enable](resources--global_log_receiver--reference--group-002.md#canonical-2032221203123021-3330200312030002-1202131321312300-0101030303001021-2010110012310223-3312320223002023-2320200203320313-3321030202011112): complete subsection reference.

- [no_ca](resources--global_log_receiver--reference--group-002.md#canonical-2222023032230211-3130303313312302-2033020112330031-2320101301203120-2023002023112301-2112323023323233-0021312332023201-0300230332330001): complete subsection reference.

<a id="canonical-3000201022320103-1233121233010210-3123222302301310-0000003213030331-2232230210030111-1230110221113100-3201102031202033-2301202322222320"></a>

<a id="canonical-3002213221221203-0222312103022010-0031003000201013-2301121030121203-1323100200201001-2121113111123122-2221230022211131-3213123000232111"></a>

## trusted_ca_url property — use_tls / 301313330111 / 4

Type: `"string"`. Optional.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-2101012112002333-3332220231103332-3232110303033320-1031011333113212-3302203300203031-3210213111123003-2200333221332203-2220010101111002"></a>

## Next pages — use_tls / 301313330111 / 5

- [datadog_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-002.md#canonical-2010023220221312-3111030320011000-2110200032003013-3202303233322013-1111312211203111-3220301131302110-1322231223311310-0312203311011213)
- [datadog_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-002.md#canonical-0110020111321033-2210111001012202-2332123203013331-3322332300222211-0122313323220002-3023320112123321-2023233210311302-1113221013133020)
- [datadog_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-002.md#canonical-0133011203031313-2230213023320331-1122200103123131-0103021310331220-1002103230211233-1103303031200300-0031110222313301-0230123001021121)
- [datadog_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-002.md#canonical-0031022130230301-1033031021021113-2330122311023233-0032121330133130-3033101001020231-1012103111232122-0232211322111331-3133222022230000)
- [datadog_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-002.md#canonical-2021031131131012-0111031111133000-3113311032202112-3331130002222011-1230231101013232-2320133011101132-1003331033201033-1333121201032332)
- [datadog_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-002.md#canonical-2032221203123021-3330200312030002-1202131321312300-0101030303001021-2010110012310223-3312320223002023-2320200203320313-3321030202011112)
- [datadog_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-002.md#canonical-2222023032230211-3130303313312302-2033020112330031-2320101301203120-2023002023112301-2112323023323233-0021312332023201-0300230332330001)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2010023220221312-3111030320011000-2110200032003013-3202303233322013-1111312211203111-3220301131302110-1322231223311310-0312203311011213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211303211311023-3023213022201131-2310131013023033-2231323200210312-1321203123213131-3220313112213232-2023210110021200-3133011132201231"></a>

## datadog_receiver.use_tls.disable_verify_certificate — disable_verify_certificate / 130232112213 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- datadog_receiver.use_tls.disable_verify_certificate

<a id="canonical-3113220300212223-2130021103010123-3201010130113032-0232222212220312-3012200023230331-0022302323000320-3023310310022032-0030313333003310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable verify certificate.

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
disable_verify_certificate = {}
```

<a id="canonical-0322132030110003-2313011302110310-1232233112203212-0102000132312133-3322031301212100-1002102103100023-0212232122121332-0030022032200232"></a>

## Direct properties — disable_verify_certificate / 130232112213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321302002131111-2003333020210020-3221021213221303-1331320023021232-1030310031311232-0022003210311120-1201313001220010-1023202311211101"></a>

## Next pages — disable_verify_certificate / 130232112213 / 4

- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0110020111321033-2210111001012202-2332123203013331-3322332300222211-0122313323220002-3023320112123321-2023233210311302-1113221013133020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132303300232113-1113133002312303-2013322231310033-0000332212313201-1311002101011111-0123211000021121-2302011030133011-1113222013130230"></a>

## datadog_receiver.use_tls.disable_verify_hostname — disable_verify_hostname / 303302201100 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- datadog_receiver.use_tls.disable_verify_hostname

<a id="canonical-3231303032021303-2223300101331120-2113322302333323-2222331300323313-1321022203030120-2320203122202133-3003103003003021-0220202102101003"></a>

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
disable_verify_hostname = {}
```

<a id="canonical-1212111003313131-2333212331132131-0011321033011232-0300320110022021-1212123213212112-3323203303230121-3113221031022233-3211330223331300"></a>

## Direct properties — disable_verify_hostname / 303302201100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011230112201131-3120323211303202-2313101221000100-2032101302033121-1133201313102222-2111302020312003-2331213233013102-1010310100212023"></a>

## Next pages — disable_verify_hostname / 303302201100 / 4

- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0133011203031313-2230213023320331-1122200103123131-0103021310331220-1002103230211233-1103303031200300-0031110222313301-0230123001021121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301032100110101-2313203120220210-0102233333010211-1003020312011312-2202330202120212-0030210211013133-3030100302321201-2123300223001210"></a>

## datadog_receiver.use_tls.enable_verify_certificate — enable_verify_certificate / 003101102011 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- datadog_receiver.use_tls.enable_verify_certificate

<a id="canonical-2031312120033231-1233001020332033-2310103133202101-2111102232201003-3222323301012101-0110000213010323-1313022322232203-2303120200321003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable verify certificate.

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
enable_verify_certificate = {}
```

<a id="canonical-2221002123230313-1102131220122221-3102013221000130-0023333120321313-0023113103223303-1123313210030210-1232222111102223-0202011200122200"></a>

## Direct properties — enable_verify_certificate / 003101102011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133130133202221-1203303321020221-1011203022033003-3112003020023133-0000233120113112-2102232230330002-0333210103131223-0323303331100322"></a>

## Next pages — enable_verify_certificate / 003101102011 / 4

- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0031022130230301-1033031021021113-2330122311023233-0032121330133130-3033101001020231-1012103111232122-0232211322111331-3133222022230000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210320003101023-3133020332110203-0211133201202000-0233232123231212-1013000022301002-1003021232122212-3100232231222100-3103021021113000"></a>

## datadog_receiver.use_tls.enable_verify_hostname — enable_verify_hostname / 001212133101 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- datadog_receiver.use_tls.enable_verify_hostname

<a id="canonical-1330321111112203-1023303310212100-0322310022233212-0323002213203312-1303120331131033-1230310331020001-2120132322003300-3322112300031200"></a>

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
enable_verify_hostname = {}
```

<a id="canonical-3102310122300320-1233123033230121-2311321031301221-0332012130032120-1010023001332111-1020121323002020-1303320130131030-2002301120011001"></a>

## Direct properties — enable_verify_hostname / 001212133101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311232032023111-2033123012132011-2032332232032322-0310201021123133-2111222210313123-3331333303311011-3202000230321032-2121211113103012"></a>

## Next pages — enable_verify_hostname / 001212133101 / 4

- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2021031131131012-0111031111133000-3113311032202112-3331130002222011-1230231101013232-2320133011101132-1003331033201033-1333121201032332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210313111021330-0021113233122300-1323101201100010-1202222333120121-0023233213232121-2120331132020032-2223210311231232-2103221230223313"></a>

## datadog_receiver.use_tls.mtls_disabled — mtls_disabled / 011322212222 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- datadog_receiver.use_tls.mtls_disabled

<a id="canonical-2030320302211130-0232233012123113-2302113030332003-2211303123301032-1312130301011011-3211203022310330-2003012220113221-3210122033022133"></a>

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
mtls_disabled = {}
```

<a id="canonical-2113031102122210-2132021022321332-2112102203032333-2311321100200200-1122123222321033-0023120033331333-0220012110130323-3101011321303133"></a>

## Direct properties — mtls_disabled / 011322212222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011332010013132-3222332131300013-2331102130101220-0030212313021310-0301231111311303-3311011210221203-3313221012011031-1031231202322311"></a>

## Next pages — mtls_disabled / 011322212222 / 4

- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2032221203123021-3330200312030002-1202131321312300-0101030303001021-2010110012310223-3312320223002023-2320200203320313-3321030202011112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222112333132331-1132200332133231-1331333023130002-2101303100231031-3311213132002213-0110023332300022-1012023001032322-1102320030133220"></a>

## datadog_receiver.use_tls.mtls_enable — mtls_enable / 031030210000 / 2

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

<a id="canonical-3022123021232100-1002121101220232-2211301033320130-1333332101120312-3132122011000321-2001331011112113-3200323030021030-1231021030200121"></a>

## Direct properties — mtls_enable / 031030210000 / 3

<a id="canonical-3231113030112213-2330031211212303-2001113222102323-3313020332233332-1000123010220203-0001021303111332-3132030231001010-1331002001023111"></a>

<a id="canonical-3333002133231132-2103210110030131-1201031002120203-3302322310223233-3013213123310211-3201003103110211-1031220033000010-2110202303310323"></a>

## certificate property — mtls_enable / 031030210000 / 4

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1122130123120331-2120331331102321-3111230130310031-2330021332333022-3132112010310203-3002132111130001-2201323230321100-2321031311320312"></a>

## Next pages — mtls_enable / 031030210000 / 5

- [datadog_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-002.md#canonical-3213223010013312-3011133203210000-0121313222022122-2332233312203221-3102301313232011-0202302130201132-2030121132311023-1012320132110113)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3213223010013312-3011133203210000-0121313222022122-2332233312203221-3102301313232011-0202302130201132-2030121132311023-1012320132110113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220210313211303-0321231010230332-0322211100022113-2300000212013100-3313221122112032-1201221310022023-3002233022322012-1203002332133210"></a>

## datadog_receiver.use_tls.mtls_enable.key_url — key_url / 232332313210 / 2

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

<a id="canonical-1002030110030021-2120130320310030-3101321033203011-1120203203013302-2031301312320210-0331102010030223-3312001203203023-3101320231021333"></a>

## Direct properties — key_url / 232332313210 / 3

- [blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-0030312032110101-1010022312212103-0101002220200130-2202102333022120-1131013320300313-3100101112310230-3130112312012303-2110233111231301): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-1203031212311331-1121212032000320-1120323231121330-2210223203201320-2011011320132003-2222011212030212-0113112011021211-2133312010301221): complete subsection reference.

<a id="canonical-3131230013030330-1001320230303002-3322222110312302-3222100013310303-3223212230022023-1020230223010011-0011332011333113-3122311002033010"></a>

## Next pages — key_url / 232332313210 / 4

- [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-0030312032110101-1010022312212103-0101002220200130-2202102333022120-1131013320300313-3100101112310230-3130112312012303-2110233111231301)
- [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-1203031212311331-1121212032000320-1120323231121330-2210223203201320-2011011320132003-2222011212030212-0113112011021211-2133312010301221)
- [datadog_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-002.md#canonical-2032221203123021-3330200312030002-1202131321312300-0101030303001021-2010110012310223-3312320223002023-2320200203320313-3321030202011112)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0030312032110101-1010022312212103-0101002220200130-2202102333022120-1131013320300313-3100101112310230-3130112312012303-2110233111231301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301002123030221-1210331000002213-2111303112121211-0023312120012331-1001222312123313-0310230022012012-3011033123333333-2223013300202313"></a>

## datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info — blindfold_secret_info / 130200230030 / 2

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

<a id="canonical-3232211233030312-1101213320223103-3112120331033102-0331202230001130-1110122211310222-2320233310120220-1212132033111232-0122232313320022"></a>

## Direct properties — blindfold_secret_info / 130200230030 / 3

<a id="canonical-0330121030032012-2210321121010120-3331302222132201-2230020212113103-1003101123211031-1030122232330103-3201103230210020-0313100230110110"></a>

<a id="canonical-3313122332322010-0102020312333012-0133323230133013-3333101031311323-0303322012031001-2111203202332201-3013031313002113-1103100302001202"></a>

## decryption_provider property — blindfold_secret_info / 130200230030 / 4

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

<a id="canonical-1230332202301332-3311210220113211-1230331110202003-3020120121220313-0103023131113223-1333220130223033-2100221322212320-0223201203111233"></a>

<a id="canonical-3113231233332300-3230301021323130-1013101203132103-1220203320102021-1320020000103120-2220222202133220-2223302103103113-3132113100331333"></a>

## location property — blindfold_secret_info / 130200230030 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2031312003333110-3032023230212221-1112113032013230-1101332312010301-2302003110310000-1120013231032001-1123202222333300-1223300103202011"></a>

## store_provider property — blindfold_secret_info / 130200230030 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-2101020202223333-2000011221110220-3202131111013323-3031120203000001-3022110302310102-0301312032231201-0102222102010010-2021130021221012"></a>

## Next pages — blindfold_secret_info / 130200230030 / 7

- [datadog_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-002.md#canonical-3213223010013312-3011133203210000-0121313222022122-2332233312203221-3102301313232011-0202302130201132-2030121132311023-1012320132110113)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1203031212311331-1121212032000320-1120323231121330-2210223203201320-2011011320132003-2222011212030212-0113112011021211-2133312010301221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031313103212031-0313033223231131-3121223001201220-0222223110130311-1013202331333301-1021011003030323-2332000003030221-1031130320210020"></a>

## datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info — clear_secret_info / 323012332113 / 2

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

<a id="canonical-1222330221031313-3232213222211200-1132201203010311-3332223022230312-1322130030310031-2123102322203202-1012203013132322-2212222300301302"></a>

## Direct properties — clear_secret_info / 323012332113 / 3

<a id="canonical-3112303302301022-2200121233221131-0102100031113030-1113312322113220-1123300211002332-0003100332213113-1113221103022302-1002000123121331"></a>

<a id="canonical-1230120023103320-2111123022023212-2123031220223320-0223103023312120-0313223221221323-2120331022101322-0202230100003000-1300132320211112"></a>

## provider_ref property — clear_secret_info / 323012332113 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3201230100332312-2330123312313133-3213022213112301-2001331201320320-3123213311213211-1120320331330001-2203110311012032-1213302331230032"></a>

<a id="canonical-3032222300111233-3212303202333232-0120223230123301-0113011301010122-1022001030121112-0310332212222313-2130010232122200-1021320223030101"></a>

## URL property — clear_secret_info / 323012332113 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2222323013002312-3003103231321232-0033223313023220-1013202220202011-1203301310232022-3232130010312131-0332121321202333-2123232230002321"></a>

## Next pages — clear_secret_info / 323012332113 / 6

- [datadog_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-002.md#canonical-3213223010013312-3011133203210000-0121313222022122-2332233312203221-3102301313232011-0202302130201132-2030121132311023-1012320132110113)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2222023032230211-3130303313312302-2033020112330031-2320101301203120-2023002023112301-2112323023323233-0021312332023201-0300230332330001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013000211133231-3120030101300122-3101230103103231-1100211332310312-0033110121312111-1203101030123332-0120110322302110-2210303020112011"></a>

## datadog_receiver.use_tls.no_ca — no_ca / 020032231033 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-3123003223013222-3130032303302330-2101330200200000-3020320030020201-1131303120030011-3132113220103112-3110012203132100-0130002012332133)
- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- datadog_receiver.use_tls.no_ca

<a id="canonical-1233332031103301-0122321320010033-2321000001023003-2210203232012321-2331102211302211-1320113213032022-3303311102020312-3300200211013201"></a>

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
no_ca = {}
```

<a id="canonical-0120220332110120-2330230021212102-2301330200020111-1031130000210313-1002231302323100-2020130100032221-3323120213330200-2012002112202320"></a>

## Direct properties — no_ca / 020032231033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031220200002330-1321201231202120-3012232302102000-2131130200212211-2111331332230203-2213202221303022-1221232312333013-1301000101323310"></a>

## Next pages — no_ca / 020032231033 / 4

- [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-0303321220323123-3020213201320312-3010202023323302-3110210001301110-0323320233100300-2323230020221020-0021023013200000-0032222203031323)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3211233331121300-1322331131001220-0102011110001201-1030212021001001-2311133020330123-0200201221122221-1011112033100211-0230123330102313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223233210331122-1132212120100001-3220210303003201-1320132020221332-0002203132130000-0031211033213033-1331310302111120-3123001322333021"></a>

## dns_logs — dns_logs / 200203102131 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- dns_logs

<a id="canonical-2112120022003322-1202120213333133-0010302200210023-3232021030212312-0232123211210022-2203312323223132-3100211030221032-3112323313223031"></a>

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
dns_logs = {}
```

<a id="canonical-0021303301112021-3312312223022023-2020300130121330-2131212031001213-1202323001112311-0233100330002001-3302310132323030-0113021312332013"></a>

## Direct properties — dns_logs / 200203102131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032003222333211-0213021020301330-1030303132022301-0102221321313311-3033332131101221-3313222131132022-1103313301012303-1300010011003103"></a>

## Next pages — dns_logs / 200203102131 / 4

- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332230002320213-0312110213000110-0223301232022303-0302131120223201-0103230120001100-0232112201201010-1313010330131023-2132221223323120"></a>

## gcp_bucket_receiver — gcp_bucket_receiver / 121120000302 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- gcp_bucket_receiver

<a id="canonical-1221011012130302-2001213102221021-3130230103112313-1211200120113032-0230233232020101-3331301323120221-0220020210111122-1233101312331111"></a>

Type: `"object"`. single nested block, Optional.

GCP Bucket Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1232131031031231-1130122303120202-1300303221002012-1121111120031233-1202210032322133-3123111033231022-0312010133133202-2033210313033032"></a>

## Direct properties — gcp_bucket_receiver / 121120000302 / 3

- [batch](resources--global_log_receiver--reference--group-002.md#canonical-3133310112221111-3211112311021302-2332002032022330-1302021211131112-1300112202033110-3120133000102323-3301230313032221-3131333233103121): complete subsection reference.

<a id="canonical-3130231201110312-2313032121312001-2133000130000033-3013022311101111-2023011002320122-0010112210322221-2300233020333130-2211200202033133"></a>

<a id="canonical-0300301320103131-1330033213230200-3031302321101001-3330331023312330-2212323320213300-1200310000330020-3203200320010030-1103033333211011"></a>

## bucket property — gcp_bucket_receiver / 121120000302 / 4

Type: `"string"`. Optional.

GCP Bucket Name. GCP Bucket Name.

Upstream description:

GCP Bucket Name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3213033210202132-1012332033300200-2212010300100302-2003111321220033-1231321012323233-1202322132121002-3322301021322013-3322232120002010"></a>

## Next pages — gcp_bucket_receiver / 121120000302 / 5

- [gcp_bucket_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-3133310112221111-3211112311021302-2332002032022330-1302021211131112-1300112202033110-3120133000102323-3301230313032221-3131333233103121)
- [gcp_bucket_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-1232123311312122-3101322333002231-3123130033011200-3303321211110303-0033003121032320-3233320022020311-1333131120222310-1122111012131010)
- [gcp_bucket_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-2021230000023312-3203222230013203-1110123330012120-3023232011030102-0013000110312031-2112130203121003-2200023323233220-2311112121100122)
- [gcp_bucket_receiver.gcp_cred](resources--global_log_receiver--reference--group-002.md#canonical-0133101020130210-0230322212213100-1003202332233312-2210312222232001-3031130333211231-2201030332232330-0000203223001032-2100203032233122)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3133310112221111-3211112311021302-2332002032022330-1302021211131112-1300112202033110-3120133000102323-3301230313032221-3131333233103121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030210033232110-0021223301110203-3132100121022311-3201231103213100-2330322233330133-0321121100131112-3321110212020030-0230232231223031"></a>

## gcp_bucket_receiver.batch — batch / 302002322331 / 2

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

<a id="canonical-1302031200030302-0313331011020103-3033001020320132-2133100323320013-0032100010233302-3111010233200021-2301332333300121-1212203301130212"></a>

## Direct properties — batch / 302002322331 / 3

<a id="canonical-0223013312231200-2103020030302213-3223210031031330-0233012223010022-2110300212232331-1331323103120313-2322130213203320-3032113130130102"></a>

<a id="canonical-0203312010231123-2030032201013032-3101133312123230-2130131001022122-0023211020111200-3323003012120223-0231311302210110-0220102320131323"></a>

## max_bytes property — batch / 302002322331 / 4

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2233222220001031-2030033311103132-3312130310110332-2223213001203123-3232031112130331-1132123201203333-1303200302222133-2302321032200333"></a>

## max_events property — batch / 302002322331 / 5

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0221210230013210-2023302131012303-2100102323201001-0122220012123021-2101202230031220-0131013300130322-1103211332220301-2022033310123022"></a>

## timeout_seconds property — batch / 302002322331 / 6

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

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

<a id="canonical-3023202031030021-0113221031223013-2310021212111202-3330002313332331-0220320123121133-3210001132013100-3133310013302032-1003330012000220"></a>

## Next pages — batch / 302002322331 / 7

- [gcp_bucket_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-002.md#canonical-2112232112210321-1320202121330001-2301100002130202-0122031332022220-1221233011030030-0000033021003233-0332132231100200-0212320233132110)
- [gcp_bucket_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-002.md#canonical-0310031031220202-3201210220000022-0313033002111113-1220312002121310-3223303121231321-2330223002020310-2212023330031221-1311222112021220)
- [gcp_bucket_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-002.md#canonical-3330333233300320-3222311303113321-0210333230001230-3121321120322212-0033120130222020-0201232021221323-3000020200323013-3100003020230121)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2112232112210321-1320202121330001-2301100002130202-0122031332022220-1221233011030030-0000033021003233-0332132231100200-0212320233132110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022300302021111-2300122113010013-0212101331111020-2123120121200233-0133110211101333-1320003011230132-0121102103110223-1223311003133332"></a>

## gcp_bucket_receiver.batch.max_bytes_disabled — max_bytes_disabled / 020331010300 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-3133310112221111-3211112311021302-2332002032022330-1302021211131112-1300112202033110-3120133000102323-3301230313032221-3131333233103121)
- gcp_bucket_receiver.batch.max_bytes_disabled

<a id="canonical-0302321120031321-1031302022223301-1032231111231213-3220103021120122-0010131102123321-0133213223200223-0133210312022003-0230121201202000"></a>

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
max_bytes_disabled = {}
```

<a id="canonical-2313313032011332-1130233122021133-3211013010333033-2133013232211020-3301231302132331-1033111310113123-2022133311022200-0303021210233000"></a>

## Direct properties — max_bytes_disabled / 020331010300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002113131300211-1322000311130332-1300120202331112-2030223303331232-1222230030302311-3032103113122010-0111311312320203-0021312222233130"></a>

## Next pages — max_bytes_disabled / 020331010300 / 4

- [gcp_bucket_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-3133310112221111-3211112311021302-2332002032022330-1302021211131112-1300112202033110-3120133000102323-3301230313032221-3131333233103121)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0310031031220202-3201210220000022-0313033002111113-1220312002121310-3223303121231321-2330223002020310-2212023330031221-1311222112021220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233132033202311-2033100220320330-1211232113030313-2300330131313013-0302323012332211-0003220332232301-0223100332231303-1010312112101023"></a>

## gcp_bucket_receiver.batch.max_events_disabled — max_events_disabled / 232112111312 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-3133310112221111-3211112311021302-2332002032022330-1302021211131112-1300112202033110-3120133000102323-3301230313032221-3131333233103121)
- gcp_bucket_receiver.batch.max_events_disabled

<a id="canonical-0200020232111312-1113302203331230-0011233020301100-1313033200032013-1000021222011203-2130000202032331-3000013121321102-1223313013230120"></a>

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
max_events_disabled = {}
```

<a id="canonical-3030300312203102-1011321120321230-1023120002102312-0123212312231230-3100302031110233-2131132210102213-1023122202033110-2331110102313133"></a>

## Direct properties — max_events_disabled / 232112111312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223210330302121-2333302112310203-0102123103133131-2031223330310100-0030032322220203-1200030233223011-1000111032221322-0330321031320320"></a>

## Next pages — max_events_disabled / 232112111312 / 4

- [gcp_bucket_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-3133310112221111-3211112311021302-2332002032022330-1302021211131112-1300112202033110-3120133000102323-3301230313032221-3131333233103121)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3330333233300320-3222311303113321-0210333230001230-3121321120322212-0033120130222020-0201232021221323-3000020200323013-3100003020230121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320112033030300-1211031110033133-0302122133011012-1322131110101223-1303212103113210-2232133212123021-0130202023122231-1110212322301101"></a>

## gcp_bucket_receiver.batch.timeout_seconds_default — timeout_seconds_default / 202221210322 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-3133310112221111-3211112311021302-2332002032022330-1302021211131112-1300112202033110-3120133000102323-3301230313032221-3131333233103121)
- gcp_bucket_receiver.batch.timeout_seconds_default

<a id="canonical-3133313201003133-3230022032311232-3020001002322113-1130301011331210-0231201033101200-1010302001020302-2212010300103121-2331111233312131"></a>

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
timeout_seconds_default = {}
```

<a id="canonical-1310231103202031-1020323231212333-1130210331201300-3002233332101201-2103303223010220-0211312023332310-1220211211312110-0330220101003311"></a>

## Direct properties — timeout_seconds_default / 202221210322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133201021323210-1232212133100032-0332211000313011-0222203132101301-1003311013232303-3133112333303200-0331102323111121-3130333312231100"></a>

## Next pages — timeout_seconds_default / 202221210322 / 4

- [gcp_bucket_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-3133310112221111-3211112311021302-2332002032022330-1302021211131112-1300112202033110-3120133000102323-3301230313032221-3131333233103121)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1232123311312122-3101322333002231-3123130033011200-3303321211110303-0033003121032320-3233320022020311-1333131120222310-1122111012131010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020312102101022-3131230011003011-0103333003013130-2103333011010032-2120322122100332-2102032033123230-0201212001300300-3032200312222203"></a>

## gcp_bucket_receiver.compression — compression / 113201323111 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- gcp_bucket_receiver.compression

<a id="canonical-1113100101333332-3031020200011212-0113011233222121-0210211132002130-1121312010223302-0002133022300102-2020003231021321-2001132123010002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Upstream description:

Compression Type.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2032133022100031-3212030023202211-0130121331211230-3122310303003202-0221302102110211-0302120130302120-0121001111203130-0320331011032103"></a>

## Direct properties — compression / 113201323111 / 3

- [compression_default](resources--global_log_receiver--reference--group-002.md#canonical-2223232203330200-2233023000001131-2230121020103101-0320310031010130-3323201130131111-3102202230111210-0133123200100233-3101032223322203): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-002.md#canonical-3313110332301023-2331103333102321-3020213133313302-2121211103001110-0113302022212311-3000001221322111-3223020103322111-0313023032130202): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-002.md#canonical-1221121211130122-1032200200200120-2133321131120020-1330100223232300-3231332322022122-0331021321033303-1131303203322322-1100022103230310): complete subsection reference.

<a id="canonical-2301013221221002-0231302133012201-3302302323313101-2113331222000230-3313303222133202-2221313331330222-0310110102030200-2002133301013203"></a>

## Next pages — compression / 113201323111 / 4

- [gcp_bucket_receiver.compression.compression_default](resources--global_log_receiver--reference--group-002.md#canonical-2223232203330200-2233023000001131-2230121020103101-0320310031010130-3323201130131111-3102202230111210-0133123200100233-3101032223322203)
- [gcp_bucket_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-002.md#canonical-3313110332301023-2331103333102321-3020213133313302-2121211103001110-0113302022212311-3000001221322111-3223020103322111-0313023032130202)
- [gcp_bucket_receiver.compression.compression_none](resources--global_log_receiver--reference--group-002.md#canonical-1221121211130122-1032200200200120-2133321131120020-1330100223232300-3231332322022122-0331021321033303-1131303203322322-1100022103230310)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2223232203330200-2233023000001131-2230121020103101-0320310031010130-3323201130131111-3102202230111210-0133123200100233-3101032223322203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032000031121011-0101102023313322-0123033330112120-1111302010121032-3321333231021013-0311112323223110-3023130102221333-0030321223011023"></a>

## gcp_bucket_receiver.compression.compression_default — compression_default / 011320111321 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-1232123311312122-3101322333002231-3123130033011200-3303321211110303-0033003121032320-3233320022020311-1333131120222310-1122111012131010)
- gcp_bucket_receiver.compression.compression_default

<a id="canonical-0002023331200331-0103323020333300-3213001300131331-0213202213222012-1012033211301210-2233003110313020-0032003333233231-2321130131313003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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
compression_default = {}
```

<a id="canonical-2232120133033322-0330002312331030-0233020322002023-0103113032211123-3000301000003122-0313132131212101-0000312203210303-3220223333132210"></a>

## Direct properties — compression_default / 011320111321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302120102011333-0000131213030332-3101302222232003-2232220321220322-0022231232031030-2120133320013202-1003021220231010-3320123001101021"></a>

## Next pages — compression_default / 011320111321 / 4

- [gcp_bucket_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-1232123311312122-3101322333002231-3123130033011200-3303321211110303-0033003121032320-3233320022020311-1333131120222310-1122111012131010)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3313110332301023-2331103333102321-3020213133313302-2121211103001110-0113302022212311-3000001221322111-3223020103322111-0313023032130202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322302023100132-2221011203023220-0111220011130010-1211122131201031-1003010312110321-3121321322113003-0332030320210231-1113330112212220"></a>

## gcp_bucket_receiver.compression.compression_gzip — compression_gzip / 212331033213 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-1232123311312122-3101322333002231-3123130033011200-3303321211110303-0033003121032320-3233320022020311-1333131120222310-1122111012131010)
- gcp_bucket_receiver.compression.compression_gzip

<a id="canonical-3011211111323330-1101102103301301-3100000032322022-0311102321320103-2323020011231330-1332010300012032-0111031322323303-0111020000311210"></a>

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
compression_gzip = {}
```

<a id="canonical-3221232330211330-1313211212122222-2310132202031031-0020111030230213-3023202323023231-1332211130333231-2003032031022220-1121303130011302"></a>

## Direct properties — compression_gzip / 212331033213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301213212313130-0210110031102023-2312001022010122-0222300300313122-3312122232303212-1122200312311211-0103313021213010-0011312103132331"></a>

## Next pages — compression_gzip / 212331033213 / 4

- [gcp_bucket_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-1232123311312122-3101322333002231-3123130033011200-3303321211110303-0033003121032320-3233320022020311-1333131120222310-1122111012131010)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1221121211130122-1032200200200120-2133321131120020-1330100223232300-3231332322022122-0331021321033303-1131303203322322-1100022103230310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310122302300011-0212323311223120-3310232330030232-2130012130311213-1003311200130221-1212000022333123-2022320311201013-3031220132231212"></a>

## gcp_bucket_receiver.compression.compression_none — compression_none / 333233032230 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-1232123311312122-3101322333002231-3123130033011200-3303321211110303-0033003121032320-3233320022020311-1333131120222310-1122111012131010)
- gcp_bucket_receiver.compression.compression_none

<a id="canonical-3112321202331322-0322330221231222-1033312200020312-3330213213211201-1200023132213030-3231212201212232-0010113020120231-1103233312223112"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression none.

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
compression_none = {}
```

<a id="canonical-1030322120021322-1212212203321231-0121301122202110-2321201331230202-2322030130132331-3033113231133121-0201031300330320-1221310013211000"></a>

## Direct properties — compression_none / 333233032230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321311211032221-0202101303213002-3223131000231233-0301222102221221-3100330202302222-1332121222010001-1311220223203011-1302332131032022"></a>

## Next pages — compression_none / 333233032230 / 4

- [gcp_bucket_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-1232123311312122-3101322333002231-3123130033011200-3303321211110303-0033003121032320-3233320022020311-1333131120222310-1122111012131010)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2021230000023312-3203222230013203-1110123330012120-3023232011030102-0013000110312031-2112130203121003-2200023323233220-2311112121100122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020201031302112-2222333013222000-1023323300112223-2213133213011133-2213000220030023-1030123130200010-2200311112232121-1030223020112232"></a>

## gcp_bucket_receiver.filename_options — filename_options / 200001030302 / 2

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

<a id="canonical-2132131012330213-3211030201122030-0321012113001302-0311223023131303-3332233020210223-1003232332220213-3233221130121101-2332131213322123"></a>

## Direct properties — filename_options / 200001030302 / 3

<a id="canonical-2233333211102303-2101002211033330-2221010210020112-1332021230321220-3303213300002100-0111330223213213-3220123130311201-0112233212311320"></a>

<a id="canonical-2201023031122101-1032311112231303-1010203221302331-3213312211121011-0221230231201110-0213022303333001-3111121011011222-3222130233311231"></a>

## custom_folder property — filename_options / 200001030302 / 4

Type: `"string"`. Optional.

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match.

Upstream description:

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match \`/^\[a-z\_\]\[a-z0-9\\\\-\\\\.\_\]\*$/i\`

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0033023332111311-2321303322310033-2023302100201131-0302031320131231-1210011230110321-2011020202220202-2211112212031321-3100300001320210"></a>

## Next pages — filename_options / 200001030302 / 5

- [gcp_bucket_receiver.filename_options.log_type_folder](resources--global_log_receiver--reference--group-002.md#canonical-3022233221120233-1000320331230333-2010030300302011-1232223000120223-1330323012121211-2311101311131002-1132323023020223-3322333322230223)
- [gcp_bucket_receiver.filename_options.no_folder](resources--global_log_receiver--reference--group-002.md#canonical-3112031102120102-0011222021120121-2303231033002133-1122031210310202-1320320322121122-1212213301012313-3121211132330212-2300110012333033)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3022233221120233-1000320331230333-2010030300302011-1232223000120223-1330323012121211-2311101311131002-1132323023020223-3322333322230223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222103122032133-1213321020203331-0212010213103322-3233010001200312-1022020220022321-2122012223122010-2311233222023231-2232110122132302"></a>

## gcp_bucket_receiver.filename_options.log_type_folder — log_type_folder / 132033030230 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-2021230000023312-3203222230013203-1110123330012120-3023232011030102-0013000110312031-2112130203121003-2200023323233220-2311112121100122)
- gcp_bucket_receiver.filename_options.log_type_folder

<a id="canonical-0212210220323110-2102232312331320-0231130123323122-2323111200103133-1323113103003030-1031232101303100-3101010020222032-0130111311301230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for log type folder.

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
log_type_folder = {}
```

<a id="canonical-2120322322003001-2312210002130130-0232320202122310-3022131011031121-0012312111333302-3232103033333212-0200001321201101-2210320210010002"></a>

## Direct properties — log_type_folder / 132033030230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032323232102313-1030333203231020-1312032231013232-0230133120301233-3031211132311123-2310220203131320-0333202212231230-3220300110202303"></a>

## Next pages — log_type_folder / 132033030230 / 4

- [gcp_bucket_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-2021230000023312-3203222230013203-1110123330012120-3023232011030102-0013000110312031-2112130203121003-2200023323233220-2311112121100122)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3112031102120102-0011222021120121-2303231033002133-1122031210310202-1320320322121122-1212213301012313-3121211132330212-2300110012333033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130230231021023-0230303313330011-1131020303303321-3320023211301131-2111212122311132-2322313323232123-1303201312330022-2103200030222231"></a>

## gcp_bucket_receiver.filename_options.no_folder — no_folder / 133203002021 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [gcp_bucket_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-2021230000023312-3203222230013203-1110123330012120-3023232011030102-0013000110312031-2112130203121003-2200023323233220-2311112121100122)
- gcp_bucket_receiver.filename_options.no_folder

<a id="canonical-0002110330232213-1110120223121211-0311033212020332-3113023000323223-2301321102121102-3320203110320113-3100310200001333-0013121110133203"></a>

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
no_folder = {}
```

<a id="canonical-3023321010103002-3332310302111021-0311221003000100-0323122323233121-0330200000320033-1101020221221003-2030120203311002-1023001333301102"></a>

## Direct properties — no_folder / 133203002021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133132130302101-3001211030113211-3222100111110203-1230101103312322-0123131020011011-2003020033120222-2011212000113301-2222020111210231"></a>

## Next pages — no_folder / 133203002021 / 4

- [gcp_bucket_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-2021230000023312-3203222230013203-1110123330012120-3023232011030102-0013000110312031-2112130203121003-2200023323233220-2311112121100122)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0133101020130210-0230322212213100-1003202332233312-2210312222232001-3031130333211231-2201030332232330-0000203223001032-2100203032233122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312322000213023-1011112103011202-1320103122021112-2200001221021121-0130011310001302-1001130202302131-1120212212033300-2220133210320022"></a>

## gcp_bucket_receiver.gcp_cred — gcp_cred / 131111030203 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- gcp_bucket_receiver.gcp_cred

<a id="canonical-0230223301232110-1003332030021033-2301330313131021-1310333023303130-2201131303320202-2103101133122332-0133210321101200-3110101212223300"></a>

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
gcp_cred {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202220231030312-1123231311033011-2323112320033333-0210100303110012-3312321303213311-3031013133122332-1021031123232112-0003022123102103"></a>

## Direct properties — gcp_cred / 131111030203 / 3

<a id="canonical-0030321011133023-3000312230010333-2100310031233221-0132303023002120-3312031320233102-0300333023300000-0202201020111110-3031321232120222"></a>

<a id="canonical-0033121202230201-2021232103003231-0022030012001121-0312002010221322-1200301031232221-3233113031301221-2312013031320012-2331303131003123"></a>

## name property — gcp_cred / 131111030203 / 4

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

<a id="canonical-1211202031120213-1031033321001130-0300201111213320-3003332012222233-0113131330220011-3003210230020312-3322101023012001-3212223330332330"></a>

<a id="canonical-2012003001200223-2003331222300202-1001133233213210-2101122322031233-2223021203030211-0120220222323001-1232230311000122-1012301233133323"></a>

## namespace property — gcp_cred / 131111030203 / 5

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

<a id="canonical-1000113133002230-3130231000312130-0030032132013100-0201120132023133-3323031302001200-0111303100330012-0221320103103220-3321311022322102"></a>

<a id="canonical-0200232020332110-2230033103321023-1202200333110030-1120213013223033-2000222221033320-0302130332011310-2133212221120200-3213322201130220"></a>

## tenant property — gcp_cred / 131111030203 / 6

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

<a id="canonical-2130011311231010-1033012122301230-1031120312123001-3031322120111322-3030332133321121-0320202033212020-1001001110131221-0330133111212222"></a>

## Next pages — gcp_cred / 131111030203 / 7

- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232101330300120-2121002322221002-2111310222112013-0022002330021211-0120232102300120-3333003322110003-0100322302130020-0120003322230333"></a>

## http_receiver — http_receiver / 133331002101 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- http_receiver

<a id="canonical-1220320110200100-0011203102210333-0202102011023021-3001200203203131-1331100111313102-1233023130201221-3030123213313322-1030001103211002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http receiver.

Upstream description:

Configuration for HTTP endpoint.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3111311311101131-1312013200212131-1132133022232201-1033103313330013-0111310313132101-2320303323103111-0332233332120310-3130001202010332"></a>

## Direct properties — http_receiver / 133331002101 / 3

- [auth_basic](resources--global_log_receiver--reference--group-002.md#canonical-1022130202001203-2123222000302011-0303000113321310-1202322232132113-3001121022210122-3232110303312101-0210032303121020-2131003101302233): complete subsection reference.

- [auth_none](resources--global_log_receiver--reference--group-002.md#canonical-3331330032301033-3102210321202122-3212212113000230-1203031121110232-3301023022023030-1033012130200332-3101222021333330-1303103201031213): complete subsection reference.

- [auth_token](resources--global_log_receiver--reference--group-002.md#canonical-3211110223222312-1122132110030223-3113232133213102-3013031131221331-1303210011032221-3303131311330213-2103031022323113-1102013133333303): complete subsection reference.

- [batch](resources--global_log_receiver--reference--group-003.md#canonical-0102310011303233-1212220212302323-3313000020023020-1100320133210211-2333121320223233-2313312313131202-1200320323030001-1320032012112211): complete subsection reference.

- [compression](resources--global_log_receiver--reference--group-003.md#canonical-3133203112210102-2303223113321220-1030012202132221-1231310113101233-1233330221020123-1232300100103012-1321231320233133-2103012323303331): complete subsection reference.

- [no_tls](resources--global_log_receiver--reference--group-003.md#canonical-3233100233103022-2230123213321331-0131001032100301-2102322211220020-3112230003112133-0301020101101012-1200122011110232-2212133013103130): complete subsection reference.

<a id="canonical-1033020103120002-0123213303012030-3303331000212223-1303000233331233-3000022223011201-2333130132001220-3313100121010130-2122211223322303"></a>

<a id="canonical-3030000012221011-3223312302223301-1333300022222231-3202211132030330-0302023101310221-3120303020200031-3320031230121320-3332021001323211"></a>

## URI property — http_receiver / 133331002101 / 4

Type: `"string"`. Optional.

HTTP URI is the URI of the HTTP endpoint to send logs to,.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123): complete subsection reference.

<a id="canonical-1212203022232223-3010102013100232-3132012111221313-1213300130031332-1200101323322032-1311013030113120-2021121321201203-3220133320030101"></a>

## Next pages — http_receiver / 133331002101 / 5

- [http_receiver.auth_basic](resources--global_log_receiver--reference--group-002.md#canonical-1022130202001203-2123222000302011-0303000113321310-1202322232132113-3001121022210122-3232110303312101-0210032303121020-2131003101302233)
- [http_receiver.auth_none](resources--global_log_receiver--reference--group-002.md#canonical-3331330032301033-3102210321202122-3212212113000230-1203031121110232-3301023022023030-1033012130200332-3101222021333330-1303103201031213)
- [http_receiver.auth_token](resources--global_log_receiver--reference--group-002.md#canonical-3211110223222312-1122132110030223-3113232133213102-3013031131221331-1303210011032221-3303131311330213-2103031022323113-1102013133333303)
- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-0102310011303233-1212220212302323-3313000020023020-1100320133210211-2333121320223233-2313312313131202-1200320323030001-1320032012112211)
- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-3133203112210102-2303223113321220-1030012202132221-1231310113101233-1233330221020123-1232300100103012-1321231320233133-2103012323303331)
- [http_receiver.no_tls](resources--global_log_receiver--reference--group-003.md#canonical-3233100233103022-2230123213321331-0131001032100301-2102322211220020-3112230003112133-0301020101101012-1200122011110232-2212133013103130)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1022130202001203-2123222000302011-0303000113321310-1202322232132113-3001121022210122-3232110303312101-0210032303121020-2131003101302233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300320221211102-0222110221300103-1123023032021113-3122021301323032-1123112030213123-1001030000023021-0031021203211122-3202203321033200"></a>

## http_receiver.auth_basic — auth_basic / 200023202023 / 2

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

<a id="canonical-3232322003232002-1222013123230011-3102011302331102-2002302031200130-2312001303121301-3031103120202300-1202001212333130-2320011331212332"></a>

## Direct properties — auth_basic / 200023202023 / 3

- [password](resources--global_log_receiver--reference--group-002.md#canonical-1033323112022211-3003201200310311-3010321003132310-1213101100232030-1130010321211202-3321121122113013-0110212000322033-2023303300123230): complete subsection reference.

<a id="canonical-3101333321222000-2322120102103323-1302130201121123-2133330211330123-3310302030031022-3110130122022333-0222111330301223-1310103230330213"></a>

<a id="canonical-0013023200322002-2331001003330010-2123103011122132-3120010122223120-3101021330130032-3233303111203303-3002013123111223-0123020210233122"></a>

## user_name property — auth_basic / 200023202023 / 4

Type: `"string"`. Optional.

username. HTTP Basic Auth username.

Upstream description:

HTTP Basic Auth username.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2123220303301012-2111312331003302-2310100100330212-1123322113103332-3022313233002033-2120130030113011-0000133032220200-0010133133102002"></a>

## Next pages — auth_basic / 200023202023 / 5

- [http_receiver.auth_basic.password](resources--global_log_receiver--reference--group-002.md#canonical-1033323112022211-3003201200310311-3010321003132310-1213101100232030-1130010321211202-3321121122113013-0110212000322033-2023303300123230)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1033323112022211-3003201200310311-3010321003132310-1213101100232030-1130010321211202-3321121122113013-0110212000322033-2023303300123230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203220313323012-1033203012123232-1332303202313330-3020132313130112-0010213032213103-2323030113210330-3301132103213200-1211331233033123"></a>

## http_receiver.auth_basic.password — password / 301100110311 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.auth_basic](resources--global_log_receiver--reference--group-002.md#canonical-1022130202001203-2123222000302011-0303000113321310-1202322232132113-3001121022210122-3232110303312101-0210032303121020-2131003101302233)
- http_receiver.auth_basic.password

<a id="canonical-2230311112201102-3323031123201210-3131022111310113-0230111200302000-0102332121021221-3033302311313132-2103000101122122-0302230223110221"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
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
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331101202003101-3300100233222122-2002100022021033-0022300022033002-3211300233011222-0013001213210101-1130333202103103-1132201222020220"></a>

## Direct properties — password / 301100110311 / 3

- [blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-0113112210111023-0200001333020232-3103313213132311-1113313220011133-2001121131112113-0123330111113002-3121122212223132-2331311211102223): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-2033102031312132-3201120320220222-1001232213300230-1220223033333221-3130212312220010-0000222323110231-1201311013023031-1133302032233013): complete subsection reference.

<a id="canonical-2120020203003021-2011332022332331-0032111100201130-3130212210033212-0301023311200023-0113031332112021-1301211302231311-3003012023312000"></a>

## Next pages — password / 301100110311 / 4

- [http_receiver.auth_basic.password.blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-0113112210111023-0200001333020232-3103313213132311-1113313220011133-2001121131112113-0123330111113002-3121122212223132-2331311211102223)
- [http_receiver.auth_basic.password.clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-2033102031312132-3201120320220222-1001232213300230-1220223033333221-3130212312220010-0000222323110231-1201311013023031-1133302032233013)
- [http_receiver.auth_basic](resources--global_log_receiver--reference--group-002.md#canonical-1022130202001203-2123222000302011-0303000113321310-1202322232132113-3001121022210122-3232110303312101-0210032303121020-2131003101302233)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0113112210111023-0200001333020232-3103313213132311-1113313220011133-2001121131112113-0123330111113002-3121122212223132-2331311211102223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003313111201220-0201323210320001-1003011123113302-3331200203100323-1101002131201003-0002321011030312-0022222120111300-0021323032221210"></a>

## http_receiver.auth_basic.password.blindfold_secret_info — blindfold_secret_info / 130002130301 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.auth_basic](resources--global_log_receiver--reference--group-002.md#canonical-1022130202001203-2123222000302011-0303000113321310-1202322232132113-3001121022210122-3232110303312101-0210032303121020-2131003101302233)
- [http_receiver.auth_basic.password](resources--global_log_receiver--reference--group-002.md#canonical-1033323112022211-3003201200310311-3010321003132310-1213101100232030-1130010321211202-3321121122113013-0110212000322033-2023303300123230)
- http_receiver.auth_basic.password.blindfold_secret_info

<a id="canonical-2021002230020210-3213022213201321-3122200021310323-1300123213003330-1031230133320031-1113102101322022-3311320011101222-2201201031200102"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3303211301030111-2010002103320113-1310110323000113-2220001120012112-2002200213003320-1320301131012102-2302101000001031-3033123113131001"></a>

## Direct properties — blindfold_secret_info / 130002130301 / 3

<a id="canonical-0131101320302333-2033321030133302-2000122020012211-3203313102003303-0231332112302220-1202002000111002-0003132033201322-3123331203122103"></a>

<a id="canonical-0131302300200133-2133232033002020-1021213313331200-3330000330201330-1202113331210101-0112003031123330-1201130010120302-3303223203002032"></a>

## decryption_provider property — blindfold_secret_info / 130002130301 / 4

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

<a id="canonical-1210110320220213-3323201211311201-2310103123313231-3212301312020313-2303121021122321-2223203320012320-0033111313020231-0023112302212023"></a>

<a id="canonical-0332313031202300-0310021212313310-0032130033222112-3300332013033033-2221112123321311-1130203121303203-3321323230200132-2002221322001020"></a>

## location property — blindfold_secret_info / 130002130301 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1311131111222112-0111123320012221-1331021102010331-1023121223022232-3013023233223233-3102233033202330-3332010321333311-3213020003323220"></a>

<a id="canonical-2032302312101310-0021133332301020-2200322021010112-3233013223332301-0120222102330113-3132111033020310-2133310133010331-2331230113300130"></a>

## store_provider property — blindfold_secret_info / 130002130301 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-1001203003032000-3232303323200122-1132030230121012-0302233121113302-1211102111202120-0231021001313222-3202213200312101-1200121100313300"></a>

## Next pages — blindfold_secret_info / 130002130301 / 7

- [http_receiver.auth_basic.password](resources--global_log_receiver--reference--group-002.md#canonical-1033323112022211-3003201200310311-3010321003132310-1213101100232030-1130010321211202-3321121122113013-0110212000322033-2023303300123230)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2033102031312132-3201120320220222-1001232213300230-1220223033333221-3130212312220010-0000222323110231-1201311013023031-1133302032233013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123133033020131-0320003331100201-0110211220001303-3023000212322112-3121130102310003-1121232121311223-2032300022013001-2113012313111120"></a>

## http_receiver.auth_basic.password.clear_secret_info — clear_secret_info / 102332102310 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.auth_basic](resources--global_log_receiver--reference--group-002.md#canonical-1022130202001203-2123222000302011-0303000113321310-1202322232132113-3001121022210122-3232110303312101-0210032303121020-2131003101302233)
- [http_receiver.auth_basic.password](resources--global_log_receiver--reference--group-002.md#canonical-1033323112022211-3003201200310311-3010321003132310-1213101100232030-1130010321211202-3321121122113013-0110212000322033-2023303300123230)
- http_receiver.auth_basic.password.clear_secret_info

<a id="canonical-1211313223030302-2022231333303013-0313323321330023-1133333313300103-1203022230012210-0333200030333101-1131012131020200-2020310322000222"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3031112021030222-3100323133111113-0010202220222330-1201132123332212-1022023310333030-1330013222010012-0102310202032313-2111023111323102"></a>

## Direct properties — clear_secret_info / 102332102310 / 3

<a id="canonical-2132111303200032-0332221123110312-0012220203030100-2310230010032203-0202311203302321-2132123331113100-2332320211112320-2311111032131001"></a>

<a id="canonical-3123133333032010-1103023101123023-2211311111033110-1120232121201131-0030033003320232-2000232302201133-0211331222001313-2111100210200003"></a>

## provider_ref property — clear_secret_info / 102332102310 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1002101213231003-0333103100122231-0020001021011201-0303001003133230-3122132211032330-2132213113132321-3122221003313332-0211012210113013"></a>

<a id="canonical-2020203210032011-3223020003210212-1333110321212130-2212311003233011-0130003221022003-0223220022101322-3200200220202132-3221133013220310"></a>

## URL property — clear_secret_info / 102332102310 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1002002223111112-2212230101203222-0012122220033020-0002333310133303-0203303311022121-3003003211232020-1012323023310220-1003332302112303"></a>

## Next pages — clear_secret_info / 102332102310 / 6

- [http_receiver.auth_basic.password](resources--global_log_receiver--reference--group-002.md#canonical-1033323112022211-3003201200310311-3010321003132310-1213101100232030-1130010321211202-3321121122113013-0110212000322033-2023303300123230)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3331330032301033-3102210321202122-3212212113000230-1203031121110232-3301023022023030-1033012130200332-3101222021333330-1303103201031213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302202211111110-0321123120030133-1330321201023020-3002233332300331-2311213201312233-3320122012123233-1300023030122111-1310301132110321"></a>

## http_receiver.auth_none — auth_none / 131331203131 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- http_receiver.auth_none

<a id="canonical-0330311131133321-1321020021102231-3100320111330201-3120000200123211-2311302122121303-0123202330213223-1101203311333222-2323220133111231"></a>

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
auth_none = {}
```

<a id="canonical-2230130202203213-1313132030310102-2023110122002230-1303201332113202-3332301310323311-1220320320013232-1112110113332313-3203300023313221"></a>

## Direct properties — auth_none / 131331203131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112210310132221-0011201303320102-2221303221323130-2023322301302122-0132012310301301-3121312301103021-0003230033300112-2323303313232212"></a>

## Next pages — auth_none / 131331203131 / 4

- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3211110223222312-1122132110030223-3113232133213102-3013031131221331-1303210011032221-3303131311330213-2103031022323113-1102013133333303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
