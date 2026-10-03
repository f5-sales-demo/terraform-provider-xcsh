---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-0222323033110022-1133012123122002-3012003321110222-3220333100201333-1202212010033202-2322230222011030-3200013221312010-0321101220101110"></a>

## storage_class_list.storage_classes.pure_service_orchestrator — pure_service_orchestrator / 132013020320 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023)
- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- storage_class_list.storage_classes.pure_service_orchestrator

<a id="canonical-2102032013030302-3313002330122221-3211033121302013-0310311210011323-0200201031123010-2233100223203002-1333301133030033-2001102003233021"></a>

Type: `"object"`. single nested block, Optional.

Storage class Device configuration for Pure Service Orchestrator.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
pure_service_orchestrator {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310200120132201-1212201313020323-0320221100102111-3022133113210230-2322031202301210-3312123202000120-0303003303031233-0232022112223122"></a>

## Direct properties — pure_service_orchestrator / 132013020320 / 3

<a id="canonical-0210010221120302-3002023220030311-0113300021331233-1333323112113322-2210123130321230-1122320202312310-3212230103230113-3032301202032312"></a>

<a id="canonical-2121003331203213-0031011212032003-0123321323212132-2010120132331302-0122323311123202-1110032323323222-2311202010220133-2211013000000112"></a>

## backend property — pure_service_orchestrator / 132013020320 / 4

Type: `"string"`. Optional.

\[Enum: block|file\] Defines type of Pure storage backend block or file. The volume will have the
aspects defined in the chosen virtual pool. Possible values are \`block\`, \`file\`.

Upstream description:

Defines type of Pure storage backend block or file. The volume will have the aspects defined in the
chosen virtual pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("block",
    "file"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "block",
    "file"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  }
}
```

<a id="canonical-1230013231223200-0211121221313133-1122130202323311-0030031320222111-3200232032013311-2022311232030123-1000230010020201-0003222313232120"></a>

<a id="canonical-0310013011201302-2033200202021020-0022223012133120-2303012013333130-3201032022311320-1313231200122131-3121202221200031-2100210120212301"></a>

## bandwidth_limit property — pure_service_orchestrator / 132013020320 / 5

Type: `"string"`. Optional.

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Upstream description:

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(12),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 12,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 12,
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
    "ves.io.schema.rules.string.max_len": "12"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "12"
  }
}
```

<a id="canonical-0112023103302112-1303020021023310-2211221203122210-1211032313103023-0130211133203211-2302013020100201-3121003232131030-0311012310302131"></a>

<a id="canonical-1231332332320011-0011133120323130-2312320201201200-1121103201210010-3033310121323112-0030323202332131-0232112230222000-0100312032232302"></a>

## iops_limit property — pure_service_orchestrator / 132013020320 / 6

Type: `"number"`. Optional.

Enable IOPS limitation. It must be between 100 and 100 million. If value is 0, IOPS limit is not
defined.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 100, Maximum: 100000000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100000000,
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
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  }
}
```

<a id="canonical-3311012222011320-1213131123331231-2220220300111201-3122302021003020-2213202111300333-0102111203202210-2121331003031132-0321301313210003"></a>

## Next pages — pure_service_orchestrator / 132013020320 / 7

- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221323220213322-2233203232120131-0131232202302312-3113311032101333-3021010311033212-2102320011021000-2030210121121022-2303120022323122"></a>

## storage_device_list — storage_device_list / 310032303000 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- storage_device_list

<a id="canonical-1220331113303212-2010311203330203-1123010023313010-2333311011310323-2013213130203030-0232022330203123-2033201030211020-3221200223133131"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
storage_device_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1221310000033330-0133202331301103-1202202321123230-3200323100322032-2133203030022130-0300320121312120-3202000133100111-3210130011133032"></a>

## Direct properties — storage_device_list / 310032303000 / 3

- [storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221): complete subsection reference.

<a id="canonical-1220331120220320-2101201032123000-1001112222333012-0003011212320323-3312000003122222-1002033301110201-2312222220113312-0322021221000130"></a>

## Next pages — storage_device_list / 310032303000 / 4

- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133021332010120-1300030312220330-1332101010220030-0101201131132102-2100210220102300-1332003312131330-3330233202302221-3120020021131203"></a>

## storage_device_list.storage_devices — storage_devices / 030013002221 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- storage_device_list.storage_devices

<a id="canonical-2303022122102011-3003130203301312-2102021200031130-1013321101232010-3022110212312002-1300222232122012-3110113302112132-1103221033211211"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Devices. List of custom storage devices.

Upstream description:

List of custom storage devices.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_device"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "hpe_storage"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("netapp_trident",
    "pure_service_orchestrator")}
```

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_devices {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313212233001302-0032101212223121-2110131210102123-3013310002203110-1220330310322230-3023201022011113-0212222321133023-1103103001212002"></a>

## Direct properties — storage_devices / 030013002221 / 3

<a id="canonical-3310031021310201-1130031213203122-2222103030112210-0221123102103022-1031001031020313-1321333202030101-2223223023022022-0123002131313001"></a>

<a id="canonical-0013223212301002-1002221102222223-2121133102231102-2112233321310100-1231113120221022-1212321100130121-0102103122130102-3313301212332023"></a>

## advanced_advanced_parameters property — storage_devices / 030013002221 / 4

Type: `["map", "string"]`. Optional.

Advanced Parameters. Map of parameter name and string value.

Upstream description:

Map of parameter name and string value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [custom_storage](resources--fleet--reference--group-003.md#canonical-1201010313101013-1022332222103330-0131311131312023-1221012021200303-3022303010333232-3021021003111313-2030220111213022-0221030232013033): complete subsection reference.

- [hpe_storage](resources--fleet--reference--group-003.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313): complete subsection reference.

- [netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023): complete subsection reference.

- [pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010): complete subsection reference.

<a id="canonical-1303002102023230-1212021212311133-3012312033032013-2102222101303330-1220300031023302-0213130200100133-0320110322222301-1232210033213310"></a>

<a id="canonical-3303313201002131-2012320332012223-2133210120310303-1121030313001030-1203023201300022-3201331211002021-2303220013200220-0113033220131223"></a>

## storage_device property — storage_devices / 030013002221 / 5

Type: `"string"`. Optional.

Storage Device. Storage device and device unit.

Upstream description:

Storage device and device unit.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2021110301020010-2123211102112300-0012230123030332-1210132311333003-3013103203221003-2011220023221200-0113331213201113-1200123302020000"></a>

## Next pages — storage_devices / 030013002221 / 6

- [storage_device_list.storage_devices.custom_storage](resources--fleet--reference--group-003.md#canonical-1201010313101013-1022332222103330-0131311131312023-1221012021200303-3022303010333232-3021021003111313-2030220111213022-0221030232013033)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-003.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1201010313101013-1022332222103330-0131311131312023-1221012021200303-3022303010333232-3021021003111313-2030220111213022-0221030232013033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322230031110302-3211221213130311-0331023032030130-0013321020132323-3112002213113130-2001120323233313-3210332231033323-0313002121021031"></a>

## storage_device_list.storage_devices.custom_storage — custom_storage / 222223112131 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- storage_device_list.storage_devices.custom_storage

<a id="canonical-0100120110322223-0112330323012102-0130122221222112-2000033202332103-2211220000120230-1202321200321112-1220322023101100-1233121231010130"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for custom storage.

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
custom_storage = {}
```

<a id="canonical-3113032001103010-1033032103130033-3221003320222121-3020231133221200-2100320201221030-0213200111033112-0333202013323222-0223333113102321"></a>

## Direct properties — custom_storage / 222223112131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322013220101112-3003300233100010-0013310203210110-2123202001302323-2222332113212000-1101021303231032-2233032221112033-3323103103133123"></a>

## Next pages — custom_storage / 222223112131 / 4

- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320331322033121-2310013021100013-3310020223221131-0311131100301323-3012320221300222-2113133232130032-3320233011111322-3212132110103100"></a>

## storage_device_list.storage_devices.hpe_storage — hpe_storage / 302120300030 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- storage_device_list.storage_devices.hpe_storage

<a id="canonical-1300123020201210-1210323131311333-3010122233320022-0003003322221212-2133201201131013-2033320331211322-0332300022020123-2323012132313022"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for hpe storage.

Upstream description:

Device configuration for HPE Storage.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_server_port",
    "username")}
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
hpe_storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211203012311032-2232102002303033-1000311031322102-1310001110323032-3010023303000223-2330003120000021-3122221100033130-1010213331231022"></a>

## Direct properties — hpe_storage / 302120300030 / 3

<a id="canonical-0100121002203112-3202321112103103-1303012230123123-1032323131130213-0012031022031020-0232110000113021-3232201211102332-3120202220033001"></a>

<a id="canonical-3202003000103010-3222112321203320-3311333002232322-0133220331113212-1233132031320111-2010210021223320-0222131132010330-3303300330211002"></a>

## api_server_port property — hpe_storage / 302120300030 / 4

Type: `"number"`. Optional.

Storage server Port. Enter Storage Server Port.

Upstream description:

Enter Storage Server Port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [iscsi_chap_password](resources--fleet--reference--group-003.md#canonical-0220213210310012-3110112220202121-0102302122303013-3211121021300211-2332032122311201-0303302300223213-2311000121321111-3230211321012132): complete subsection reference.

<a id="canonical-1113023311332300-1111230300333131-2323221112111110-0103131033112032-2210122222300322-1023223200313311-3010102233111331-2020111222300303"></a>

<a id="canonical-1301221011330011-0321103110303323-0100001032023311-1313202120212110-3321000011110000-3301012122111332-2130233331312203-3202223322203210"></a>

## iscsi_chap_user property — hpe_storage / 302120300030 / 5

Type: `"string"`. Optional.

Chap Username to connect to the HPE storage.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [password](resources--fleet--reference--group-003.md#canonical-1212101202211001-2200100120130110-1332213110330321-2213301011202321-1111330332131000-2222221012033230-1010112133212033-3012203021132102): complete subsection reference.

<a id="canonical-3130131021220103-2310130023223033-1212220103022231-1112133102113030-2332213330212211-0313231133112020-2322303022023331-1002122021033120"></a>

<a id="canonical-3032123030300110-1130223212311001-2110311130300113-1332110100200220-1132103201323010-1222011223032101-2112111230220320-3330133102121133"></a>

## storage_server_ip_address property — hpe_storage / 302120300030 / 6

Type: `"string"`. Optional.

Storage Server IP address. Enter storage server IP address.

Upstream description:

Enter storage server IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1032221020101332-1220231120222122-1123312121002001-1022131000123000-3021310100101310-1101203103020213-2020210021223303-3300013310231100"></a>

<a id="canonical-2211103132012212-1131232130232221-1101131023110120-0322102022231233-0323001303111211-1230001200232023-0331230101111230-1320103302331213"></a>

## storage_server_name property — hpe_storage / 302120300030 / 7

Type: `"string"`. Optional.

Storage Server Name. Enter storage server Name.

Upstream description:

Enter storage server Name.

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
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-2121200030301312-3111212101221202-2323001030322113-3033230130121230-3323321113002213-2200201021313103-3303013001302212-1220303033211111"></a>

<a id="canonical-1312131310110231-3223012223230221-0113000001122333-3232303032021001-2123332021111130-2021131030002300-2210023313210332-1333211130132000"></a>

## username property — hpe_storage / 302120300030 / 8

Type: `"string"`. Optional.

Username to connect to the HPE storage management IP.

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
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0110230310030021-0302230320303132-3330120330312022-3320200211131013-2030332103220003-2103202121000211-1120323133233200-1201201333223003"></a>

## Next pages — hpe_storage / 302120300030 / 9

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--fleet--reference--group-003.md#canonical-0220213210310012-3110112220202121-0102302122303013-3211121021300211-2332032122311201-0303302300223213-2311000121321111-3230211321012132)
- [storage_device_list.storage_devices.hpe_storage.password](resources--fleet--reference--group-003.md#canonical-1212101202211001-2200100120130110-1332213110330321-2213301011202321-1111330332131000-2222221012033230-1010112133212033-3012203021132102)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0220213210310012-3110112220202121-0102302122303013-3211121021300211-2332032122311201-0303302300223213-2311000121321111-3230211321012132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233333212033103-3222332222313011-3112103202331232-0223311000012221-1032303201032221-1223220111130011-1303010020313002-1132303321021221"></a>

## storage_device_list.storage_devices.hpe_storage.iscsi_chap_password — iscsi_chap_password / 202302222032 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-003.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password

<a id="canonical-1333200000001010-1022131030213331-0211121200013033-1001011211103100-2300121021230232-2202133313323302-1103002233203000-1032302020222310"></a>

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
iscsi_chap_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112212303310230-3233220300122000-2031332303011102-3320020030220130-3030231320221021-3210202112321230-0202323023132000-1130213232200030"></a>

## Direct properties — iscsi_chap_password / 202302222032 / 3

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-2110222221121130-3323000022312202-1222221312330331-2220310022003331-2023131232330003-0302030033031223-2332203030132000-0330011232232203): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-1321102021232213-1301130012100310-0102103120131110-2203033023233011-0120110210130302-2103133301000030-1210101310333003-3120322303012013): complete subsection reference.

<a id="canonical-1323121023230203-2131223313301211-2222103121133333-3002310133330032-0221331012203333-3303303012113023-2032101321333100-3100210330122223"></a>

## Next pages — iscsi_chap_password / 202302222032 / 4

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-2110222221121130-3323000022312202-1222221312330331-2220310022003331-2023131232330003-0302030033031223-2332203030132000-0330011232232203)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info](resources--fleet--reference--group-003.md#canonical-1321102021232213-1301130012100310-0102103120131110-2203033023233011-0120110210130302-2103133301000030-1210101310333003-3120322303012013)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-003.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2110222221121130-3323000022312202-1222221312330331-2220310022003331-2023131232330003-0302030033031223-2332203030132000-0330011232232203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130031120010031-1120332103211220-1022312222103131-2211322300311100-3021212312201222-1322122033102330-2102031321303332-1021002021301330"></a>

## storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info — blindfold_secret_info / 232230100302 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-003.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--fleet--reference--group-003.md#canonical-0220213210310012-3110112220202121-0102302122303013-3211121021300211-2332032122311201-0303302300223213-2311000121321111-3230211321012132)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info

<a id="canonical-2011203201120113-1011203012002231-3321012033222312-3302200210000033-0101320221330033-1310220100320112-3313103113220210-3121102031033020"></a>

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

<a id="canonical-2232230323300000-0332210220323123-1101122323333103-0301032032300103-2232100131323101-1200213100201203-2121012020031201-2131331211121100"></a>

## Direct properties — blindfold_secret_info / 232230100302 / 3

<a id="canonical-0310223111332103-3122221332110000-1101112020123021-1003212223101213-2222012210003310-1121120200303201-2031233002322332-3013330032133300"></a>

<a id="canonical-3122101330312231-3002213322212032-3302032010230320-2013303230032202-0122330121123000-2213233223130113-1311220100030331-2213132210020030"></a>

## decryption_provider property — blindfold_secret_info / 232230100302 / 4

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

<a id="canonical-0221310303202313-2121011001201312-0323032021231120-1231303310202303-0000212233133322-2133230331030213-3030000310233012-2130200100331231"></a>

<a id="canonical-0313302331232101-2231301003232313-3111130130330330-3010233012300101-0031011133233213-3120132122103002-0011012232112223-0022233020033020"></a>

## location property — blindfold_secret_info / 232230100302 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3311001322212321-2333002021021222-3322120012331121-0300200100032020-3102113012132020-2302233233033331-0112031130113223-2031302102120003"></a>

<a id="canonical-3122232221112022-0313210102311230-1320120230303332-3213132002323122-0301123023210203-1101210020220020-1310211303202320-3333302231321021"></a>

## store_provider property — blindfold_secret_info / 232230100302 / 6

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

<a id="canonical-0003322310130302-0010311220203101-3013002313012212-2011013130010223-1332213310212101-0232332123132001-2231220100113213-2132000211222010"></a>

## Next pages — blindfold_secret_info / 232230100302 / 7

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--fleet--reference--group-003.md#canonical-0220213210310012-3110112220202121-0102302122303013-3211121021300211-2332032122311201-0303302300223213-2311000121321111-3230211321012132)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1321102021232213-1301130012100310-0102103120131110-2203033023233011-0120110210130302-2103133301000030-1210101310333003-3120322303012013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013322110310211-0102203310221100-2223331001012133-3320221110133213-2110311032012320-0110011301112032-3330213231310003-1131022113210233"></a>

## storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info — clear_secret_info / 121311112030 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-003.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--fleet--reference--group-003.md#canonical-0220213210310012-3110112220202121-0102302122303013-3211121021300211-2332032122311201-0303302300223213-2311000121321111-3230211321012132)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info

<a id="canonical-2031303221102222-0030111031002030-2333101200112203-2000220011111223-1013231103320302-2220212302122131-0020002230132303-0210001210303021"></a>

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

<a id="canonical-0232100203200101-0332311322313233-1133332001333323-0311020133230203-1003033222101211-0211110003320303-3120223002211022-1013122323300033"></a>

## Direct properties — clear_secret_info / 121311112030 / 3

<a id="canonical-1311031210011103-1312322211013031-2110122020221302-1303332022002303-2222122113131032-1123321300103331-2320031132310001-1302302232112203"></a>

<a id="canonical-1010000330021202-3003023220010231-2232110211332112-1313021101301200-1301311231003133-0112330310301032-0310013022020010-1233221032203132"></a>

## provider_ref property — clear_secret_info / 121311112030 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0023011130010020-0210202013100310-2302312101333331-3012210122330131-3330211000123200-3131033000332313-1331211013220023-2212023302123100"></a>

<a id="canonical-1001301001003031-0133120332132321-1033123200132002-0321120033321230-0110021032121001-0022311222122223-1302111220231310-1331321212000033"></a>

## URL property — clear_secret_info / 121311112030 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1330002310111102-0020023102030213-0233303312201010-2330212012203312-0033332221003121-2333123202002223-0121100132211133-2333002201213330"></a>

## Next pages — clear_secret_info / 121311112030 / 6

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--fleet--reference--group-003.md#canonical-0220213210310012-3110112220202121-0102302122303013-3211121021300211-2332032122311201-0303302300223213-2311000121321111-3230211321012132)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1212101202211001-2200100120130110-1332213110330321-2213301011202321-1111330332131000-2222221012033230-1010112133212033-3012203021132102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021233212112130-1120302323030002-0213033232033223-3232220233200012-2033331000313302-3113100321000021-2303303213132011-3022131031210232"></a>

## storage_device_list.storage_devices.hpe_storage.password — password / 320113101222 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-003.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- storage_device_list.storage_devices.hpe_storage.password

<a id="canonical-2131003020012222-1312020222122031-3033303322002020-2203010302123011-0230321123321121-3033331222100311-1330300022110313-0030033002310200"></a>

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

<a id="canonical-1331232032301033-1012211233330330-0301103122303310-3231003311032120-3231011200020130-3212203303010003-3020312111002212-3122301322001211"></a>

## Direct properties — password / 320113101222 / 3

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-0010013002333102-3332313022212223-1012222103020322-1300001011331221-2103322322102322-0130010331210030-3102131110222030-1313111002102211): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-1133000301131223-1230300011311221-0101232113102020-0200321001220323-0203213311200330-3102102112332231-2222313113210133-3102322100230123): complete subsection reference.

<a id="canonical-2300130320020131-3132202012020022-0012331131201012-1200131213301030-1231031201112101-0333232103013023-3330300033001312-0032312201322031"></a>

## Next pages — password / 320113101222 / 4

- [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-0010013002333102-3332313022212223-1012222103020322-1300001011331221-2103322322102322-0130010331210030-3102131110222030-1313111002102211)
- [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info](resources--fleet--reference--group-003.md#canonical-1133000301131223-1230300011311221-0101232113102020-0200321001220323-0203213311200330-3102102112332231-2222313113210133-3102322100230123)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-003.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0010013002333102-3332313022212223-1012222103020322-1300001011331221-2103322322102322-0130010331210030-3102131110222030-1313111002102211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103210012013333-2331011200323022-1033312120201332-0202000030003302-3111020321221322-3331223310322133-1002110030010212-3112310302113013"></a>

## storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info — blindfold_secret_info / 313200020103 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-003.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- [storage_device_list.storage_devices.hpe_storage.password](resources--fleet--reference--group-003.md#canonical-1212101202211001-2200100120130110-1332213110330321-2213301011202321-1111330332131000-2222221012033230-1010112133212033-3012203021132102)
- storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info

<a id="canonical-1311211111133200-1013113321000123-1121102322020233-1202023132103002-0211322231023231-3112300311030220-3000321101121103-0111312132002321"></a>

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

<a id="canonical-3001312011103323-2121112003301220-1123332202310313-2332111003213013-3100012231103000-3130100033020332-2132300213130122-0201301321031302"></a>

## Direct properties — blindfold_secret_info / 313200020103 / 3

<a id="canonical-1230200211213203-0023103220022320-2012121111222312-3202203202002331-3213332032113330-3112200111232331-0323033022111020-1122232100033103"></a>

<a id="canonical-0102213212220322-2303302131031331-0323301122023232-3123002331302110-3100302213122232-1220223011120300-2312133102031333-0010332131223300"></a>

## decryption_provider property — blindfold_secret_info / 313200020103 / 4

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

<a id="canonical-2010013323212311-0312100003202302-2313221011112310-3221130010001122-3331321232001000-0013002322330120-3311002013321112-0332130332112010"></a>

<a id="canonical-2132203331132100-1020130213303310-2010220320203120-0132020320133220-1031121103023320-1101231012302311-2013101300121203-1030223020222112"></a>

## location property — blindfold_secret_info / 313200020103 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1331332100110312-0032030030111020-2021232333003123-0233203113223031-1333032232123212-2020111333300310-0020130112231033-1230102311132033"></a>

<a id="canonical-0321313221113012-1003110322002000-0310231123010223-3231233003211323-2130002310220030-1310103321020330-1101212322111101-3012232100003223"></a>

## store_provider property — blindfold_secret_info / 313200020103 / 6

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

<a id="canonical-1200011032322201-1210202233110320-1332122233020031-0222201321020211-0303312310323331-1111103033323312-2002211221113231-1201102032200011"></a>

## Next pages — blindfold_secret_info / 313200020103 / 7

- [storage_device_list.storage_devices.hpe_storage.password](resources--fleet--reference--group-003.md#canonical-1212101202211001-2200100120130110-1332213110330321-2213301011202321-1111330332131000-2222221012033230-1010112133212033-3012203021132102)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1133000301131223-1230300011311221-0101232113102020-0200321001220323-0203213311200330-3102102112332231-2222313113210133-3102322100230123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210213303033223-1223123010120100-2322310103211103-3103131000221330-0322033213313331-1100102213310203-1113223300212001-1313321112020121"></a>

## storage_device_list.storage_devices.hpe_storage.password.clear_secret_info — clear_secret_info / 110120020203 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-003.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- [storage_device_list.storage_devices.hpe_storage.password](resources--fleet--reference--group-003.md#canonical-1212101202211001-2200100120130110-1332213110330321-2213301011202321-1111330332131000-2222221012033230-1010112133212033-3012203021132102)
- storage_device_list.storage_devices.hpe_storage.password.clear_secret_info

<a id="canonical-3312223300230222-2221102330132003-3020131323013102-0132313003022123-1232103332312031-0003013222032311-1012223300320022-2031312120310132"></a>

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

<a id="canonical-0122320322320233-2310200130132103-0211213220211232-1120202232221302-1322323223122201-3120032131001211-0122103103123202-1211033211333301"></a>

## Direct properties — clear_secret_info / 110120020203 / 3

<a id="canonical-3122220201313000-0131330200201321-0302111212222202-0301312121203231-2002330300211103-2030011333102333-3101122221223110-1111212300012332"></a>

<a id="canonical-2201312023130303-2331223002320331-1123212012033301-2103332312223131-1221022223312103-2200300312323011-0223131312231223-0211313323001123"></a>

## provider_ref property — clear_secret_info / 110120020203 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1323030021203330-1133231111120013-3020313012220331-3330122100131201-1303332322133232-2220220002322022-2230322100011010-1000202301231102"></a>

<a id="canonical-2110122030130132-0112112312333233-2111213312221122-1030210312300133-2003120003120133-0121023001320312-2322200333131112-0033000313013320"></a>

## URL property — clear_secret_info / 110120020203 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2322100100332221-0032111331131121-1210013013211013-1331322323221122-1200223033330312-2010221333012331-0313033113200012-0123023302221121"></a>

## Next pages — clear_secret_info / 110120020203 / 6

- [storage_device_list.storage_devices.hpe_storage.password](resources--fleet--reference--group-003.md#canonical-1212101202211001-2200100120130110-1332213110330321-2213301011202321-1111330332131000-2222221012033230-1010112133212033-3012203021132102)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231333321313111-2333201130233302-2201230010300203-0320331320020212-2023320311231202-3122331302200020-2132233022020112-0112303202110213"></a>

## storage_device_list.storage_devices.netapp_trident — netapp_trident / 200023320013 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- storage_device_list.storage_devices.netapp_trident

<a id="canonical-3101021010223311-2100233103222002-2031222103120330-0032130300130100-1032231200121110-3313101222102111-2301311133330031-1001330210033033"></a>

Type: `"object"`. single nested block, Optional.

Device configuration for NetApp Trident Storage.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("netapp_backend_ontap_nas",
    "netapp_backend_ontap_san")}
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
  "x-ves-oneof-field-backend_choice": "[\"netapp_backend_ontap_nas\",\"netapp_backend_ontap_san\"]"
}
```

Terraform syntax:

```terraform
netapp_trident {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231202211300011-1322002020302200-1321111123100331-0133001322212331-1132100002301032-3003021113031110-2103222300102032-2222312020201133"></a>

## Direct properties — netapp_trident / 200023320013 / 3

- [netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131): complete subsection reference.

- [netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101): complete subsection reference.

<a id="canonical-0111111203221102-2231321331020101-3002032130230002-3001332220320032-1002210023220132-2002222032100313-2320332123211321-3001203023012321"></a>

## Next pages — netapp_trident / 200023320013 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300323203020332-0210001200332133-1220330312013203-1333312033323011-1010101133110200-2013222032211002-3331330133331300-0231211023313331"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas — netapp_backend_ontap_nas / 032330200123 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas

<a id="canonical-1300011332320212-0313000331011031-0302220301103130-2120310231022210-3113120211122333-2233132200230103-0133012320023111-1003312123221132"></a>

Type: `"object"`. single nested block, Optional.

Configuration of storage backend for NetApp ONTAP NAS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_driver_name",
    "username"),
  validators.ConflictingObjectAttributes("data_lif_dns_name",
    "data_lif_ip"),
  validators.ConflictingObjectAttributes("management_lif_dns_name",
    "management_lif_ip")}
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
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

Terraform syntax:

```terraform
netapp_backend_ontap_nas {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302211121101212-3132220222212110-0322231113131213-1033100121102132-1030323221121121-3002021322231203-2210200212221211-1120013310000123"></a>

## Direct properties — netapp_backend_ontap_nas / 032330200123 / 3

- [auto_export_cidrs](resources--fleet--reference--group-003.md#canonical-0011311003202231-1223000031133112-0130132012021312-2120200332320230-3122232220212010-0122313230210323-1103223303212302-0100333200330320): complete subsection reference.

<a id="canonical-1231313321302020-0000021011303123-2330330331320010-1223023221101211-0133213231123203-1131200031300320-3033310031323100-0112203230012230"></a>

<a id="canonical-0022030101301030-3021320100030033-1333210111302031-0302133123323233-1330330201002130-0103303321131333-1111300313231201-0201031310310102"></a>

## auto_export_policy property — netapp_backend_ontap_nas / 032330200123 / 4

Type: `"bool"`. Optional.

Policy configuration for this feature.

Upstream description:

Enable automatic export policy creation and updating.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1210031232131302-3211232322012102-0312312001221211-1303202120233112-0220230011320213-3122111223212303-3210013310230200-2023311101200233"></a>

<a id="canonical-3003132112211212-1202232320220331-0111023132202230-0030302230213222-3031320002132312-0311212333012002-0031222333002131-3021221132013303"></a>

## backend_name property — netapp_backend_ontap_nas / 032330200123 / 5

Type: `"string"`. Optional.

Configuration of Backend Name. Driver is name + '\_' + dataLIF.

Upstream description:

Configuration of Backend Name. Driver is name + "\_" + dataLIF.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 50,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 50,
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
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3001312332233233-1101010101222020-2220131133121111-0220001230130021-1011023110132210-2331213312311031-0033300300221311-0000100202000220"></a>

<a id="canonical-3130201230312012-2032330210221013-1022103131132031-3313221313201220-1211112213213100-2313031120330323-0323012303320312-2130311320212022"></a>

## client_certificate property — netapp_backend_ontap_nas / 032330200123 / 6

Type: `"string"`. Optional.

Please Enter base64-encoded value of client certificate. Used for certificate-based auth.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](resources--fleet--reference--group-003.md#canonical-2021123321033031-2131220302333132-2013232223201100-1211122103213311-1123311311323332-1333300211323100-3332210130000322-1110303130220221): complete subsection reference.

<a id="canonical-3210222313333120-1212302323213103-3133313032220232-3301321131112311-0131030311013101-2201212122100130-0301203031102202-0010102220002212"></a>

<a id="canonical-2301101103002233-3300321023203313-2110101010112031-2233031003020020-0322010000110320-1012032030131331-3311012030200032-2232122230132211"></a>

## data_lif_dns_name property — netapp_backend_ontap_nas / 032330200123 / 7

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0011310011100021-1012023130231030-2300322111100111-2311310221223212-1232220332133103-2331331030102111-0301313200131123-3302320021212020"></a>

<a id="canonical-2332113321201213-2320220302200010-3321032003231000-3312111131222233-3223212301132132-3023003312002230-0332130021333120-0331032003111200"></a>

## data_lif_ip property — netapp_backend_ontap_nas / 032330200123 / 8

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-1131321231021330-0321000230001110-2123300312321330-1302102112312221-0131313231010120-0020130003000031-1010310322112033-0132022123002232"></a>

<a id="canonical-0301001203111202-0223202113332010-1132030223223122-2120120032111212-0020003120122221-0313200133213023-3122301012033230-1323020210000130"></a>

## labels property — netapp_backend_ontap_nas / 032330200123 / 9

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":20},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"20\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-1212012020302331-1100101013232022-0233323020333023-3322001031010031-2322102212323002-2030032231220110-3101022310322303-3320002213012332"></a>

<a id="canonical-1230002221002101-2111123321012113-0331333222202312-0123011221220231-1111302323000212-1132301103331010-1320323031000133-0223012223331023"></a>

## limit_aggregate_usage property — netapp_backend_ontap_nas / 032330200123 / 10

Type: `"string"`. Optional.

Fail provisioning if usage is above this percentage. Not enforced by default.

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

<a id="canonical-2012300211032020-3100103233101013-1112312020230110-2232202023230000-3331130312300221-1031222231120030-1122333202112311-1311303202112303"></a>

<a id="canonical-0213101123211033-2301123211020022-2333232312121212-0102311201310111-3323302321200311-0303020030312200-1010220013223310-0113200223131222"></a>

## limit_volume_size property — netapp_backend_ontap_nas / 032330200123 / 11

Type: `"string"`. Optional.

Fail provisioning if requested volume size is above this value. Not enforced by default.

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

<a id="canonical-3301201200012332-1032312121230011-3010112210201220-1002102302101201-2122102002110231-1013031101220221-1011311032313100-3022100332012303"></a>

<a id="canonical-2223232032212033-3312022203320311-2003202302223102-3010031003213232-1120111232110210-1202030330320123-2103331112002213-3123032200132301"></a>

## management_lif_dns_name property — netapp_backend_ontap_nas / 032330200123 / 12

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0130113221022031-0231200122210132-2030110123103212-2221320021223333-3023300223020032-1213220212331122-3020101001201101-0110330111110102"></a>

<a id="canonical-0001021101113300-3223220211210301-0323221232310321-2120012121003031-1012000033111233-2130230021132233-3123310303200300-3332213222211200"></a>

## management_lif_ip property — netapp_backend_ontap_nas / 032330200123 / 13

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2322331212203101-2032212003322321-2303333321222121-3213010111101213-3103221220001131-1310221300121132-3131331323311222-1000323120210310"></a>

<a id="canonical-2001201312031221-3111232021010323-1003123012323302-2010000010202322-2201200200131101-2111000332211003-3111313330300002-1222330130000130"></a>

## nfs_mount_options property — netapp_backend_ontap_nas / 032330200123 / 14

Type: `"string"`. Optional.

Comma-separated list of NFS mount OPTIONS. Not enforced by default.

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

- [password](resources--fleet--reference--group-003.md#canonical-1300320120013220-1111311020033332-2110201300332223-1330230012212101-3132130210000210-0310202102232330-2230320110022220-0211021223113033): complete subsection reference.

<a id="canonical-3331213221131132-2023230121203212-2003000331121231-1131000100303122-0203332230231322-2211313103200331-1012110002220312-0000233210031123"></a>

<a id="canonical-2110333203331223-3323233030330301-2023101112010321-0000130010023220-0222311301123032-1003130330211230-3333121232121221-1030103022033212"></a>

## region property — netapp_backend_ontap_nas / 032330200123 / 15

Type: `"string"`. Optional.

Backend Region. Virtual Pool Region.

Upstream description:

Virtual Pool Region.

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

- [storage](resources--fleet--reference--group-003.md#canonical-0232211313301321-0212033213020113-3120021131102001-2231330003231332-2031221231130011-2302223013200320-3030323232233300-2202032103012300): complete subsection reference.

<a id="canonical-3030322130013021-2111301300020323-2133223002133323-3201132312233323-3213023110301130-2011320110230132-2031332112331310-1311031120333330"></a>

<a id="canonical-0120111131021121-2011003112323121-2121021223223203-0201032323133130-2122132003223003-3231310220230231-3312213101012230-0031003023030023"></a>

## storage_driver_name property — netapp_backend_ontap_nas / 032330200123 / 16

Type: `"string"`. Optional.

\[Enum: ontap-nas|ontap-nas-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-nas\`, \`ontap-nas-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-0232321133002011-1032223232113020-2012130130032112-3231232020120031-1021011101313130-1332203011210311-1333233012303100-3331303030013213"></a>

<a id="canonical-1130021230121202-3302332133310233-3310301201102130-1202301102210332-0102131022022232-2300110103212130-1103012011120131-2310101022333110"></a>

## storage_prefix property — netapp_backend_ontap_nas / 032330200123 / 17

Type: `"string"`. Optional.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

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

<a id="canonical-2100310330303223-1020310133303103-0313303300201020-1332121130232322-1111111122320223-3321120302111311-2111232030101112-2002202020311133"></a>

<a id="canonical-0202301100202132-0101300123113003-3201310121011211-0110331022101201-0322031113120010-1130312032203133-0100110132230000-2111111200311012"></a>

## svm property — netapp_backend_ontap_nas / 032330200123 / 18

Type: `"string"`. Optional.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3320120122022020-3010323023113330-1330001220230111-3323222113112202-2330103233020310-3103030300323011-3012010200313232-2330213320122230"></a>

<a id="canonical-3333311221132103-3030332123103122-3020311201032311-0231212212203020-1131013223330201-0110320321202311-3023220131023231-3011010112023131"></a>

## trusted_ca_certificate property — netapp_backend_ontap_nas / 032330200123 / 19

Type: `"string"`. Optional.

Please Enter base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

<a id="canonical-3012113013132330-1331210322023311-2031013220130030-1123311013230221-3210313012301333-1020213220003201-3030032222232121-3000022333031310"></a>

<a id="canonical-3212302030211000-2303232023021230-2100011032212221-1013221322210302-2302212303013123-2133030303221011-2333000311112303-2211300113231312"></a>

## username property — netapp_backend_ontap_nas / 032330200123 / 20

Type: `"string"`. Optional.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

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
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--fleet--reference--group-003.md#canonical-2331201210221021-2100203101103023-3231010232113332-2331001322002230-3311121030003303-0210021021322130-2210111331110130-1031002103213331): complete subsection reference.

<a id="canonical-0313101010120110-1210003230032122-2300121310302230-1000003033322101-2102021202231003-1333330312310220-2103212110001130-1131033001121112"></a>

## Next pages — netapp_backend_ontap_nas / 032330200123 / 21

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](resources--fleet--reference--group-003.md#canonical-0011311003202231-1223000031133112-0130132012021312-2120200332320230-3122232220212010-0122313230210323-1103223303212302-0100333200330320)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--fleet--reference--group-003.md#canonical-2021123321033031-2131220302333132-2013232223201100-1211122103213311-1123311311323332-1333300211323100-3332210130000322-1110303130220221)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--fleet--reference--group-003.md#canonical-1300320120013220-1111311020033332-2110201300332223-1330230012212101-3132130210000210-0310202102232330-2230320110022220-0211021223113033)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--fleet--reference--group-003.md#canonical-0232211313301321-0212033213020113-3120021131102001-2231330003231332-2031221231130011-2302223013200320-3030323232233300-2202032103012300)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](resources--fleet--reference--group-003.md#canonical-2331201210221021-2100203101103023-3231010232113332-2331001322002230-3311121030003303-0210021021322130-2210111331110130-1031002103213331)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0011311003202231-1223000031133112-0130132012021312-2120200332320230-3122232220212010-0122313230210323-1103223303212302-0100333200330320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220103113233210-1111011010010113-0302222232313101-1113322232303312-0102122122303013-0011130133321130-3121033333200010-2223331211313023"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs — auto_export_cidrs / 320211310131 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs

<a id="canonical-1021110303213110-3302031121103320-1332312323011213-2212202132032211-0321032102312102-2311210101332031-1000233032220111-3130222013121311"></a>

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
auto_export_cidrs {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311303130330322-0313122230101100-0233103100313103-3231033303220302-3331133333003130-0203231033212030-3121122302023033-1130032133221312"></a>

## Direct properties — auto_export_cidrs / 320211310131 / 3

<a id="canonical-0122131302011200-2233200123103300-2303003303103203-1112113321200021-3323333011212210-1030320002323023-3103333102210313-0231120222201103"></a>

<a id="canonical-0032311132202131-2101021201233032-3230100023033221-0101000130121020-1203211101231303-3110220212023123-1202113033333300-0123300233211103"></a>

## prefixes property — auto_export_cidrs / 320211310131 / 4

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

<a id="canonical-2331221021320231-3201131220212031-0223211122103322-2301213221032101-1013210331131320-2103332123123200-0033021133130122-2131222123221013"></a>

## Next pages — auto_export_cidrs / 320211310131 / 5

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2021123321033031-2131220302333132-2013232223201100-1211122103213311-1123311311323332-1333300211323100-3332210130000322-1110303130220221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201321001300232-2331122110011232-1100312103313012-2110223321123002-3213320222231311-0232122013130132-3322231331110203-2111131223223131"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key — client_private_key / 031111111222 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key

<a id="canonical-1230113112310231-3223032113113331-1001113333123330-1010323232000312-1233301023032323-3102323330023311-2310212310031311-1033113201201300"></a>

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
client_private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110332222313102-1130130003202301-2233101001013202-3213310303233102-2012103220000112-0203332033322202-1320202103130331-0133122120023122"></a>

## Direct properties — client_private_key / 031111111222 / 3

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-0331023132232112-0331320023011112-0333332032213010-0300022123213023-1213202320223202-3303002001220113-0003321131232310-3000320032023032): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-1321101120300120-2212323121331132-3201213113030033-2101133022330202-0110030321310301-1333313310331212-2011120330113131-2030020001022003): complete subsection reference.

<a id="canonical-1110201021311133-3300230103101203-2021331221202222-0331302110233202-0113312001032002-2312312101223030-3001323332012121-3313220030123212"></a>

## Next pages — client_private_key / 031111111222 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-0331023132232112-0331320023011112-0333332032213010-0300022123213023-1213202320223202-3303002001220113-0003321131232310-3000320032023032)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](resources--fleet--reference--group-003.md#canonical-1321101120300120-2212323121331132-3201213113030033-2101133022330202-0110030321310301-1333313310331212-2011120330113131-2030020001022003)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0331023132232112-0331320023011112-0333332032213010-0300022123213023-1213202320223202-3303002001220113-0003321131232310-3000320032023032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030331110321100-3113112131321011-2020300200030020-2332223120212120-2022203210111031-0333133213023211-0310211012002311-3310200331002130"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info — blindfold_secret_info / 233203112010 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--fleet--reference--group-003.md#canonical-2021123321033031-2131220302333132-2013232223201100-1211122103213311-1123311311323332-1333300211323100-3332210130000322-1110303130220221)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info

<a id="canonical-2110132010000223-1001012310202032-1013231232032230-1202010011133322-1020200101303101-0222211103111022-3123010101021331-0200321000110323"></a>

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

<a id="canonical-3031120231320210-2330212111301323-3112130201132010-1321111130112102-3321013312021130-3233233012132200-3310301103332032-0310020211231302"></a>

## Direct properties — blindfold_secret_info / 233203112010 / 3

<a id="canonical-0200303202212032-2320201022103332-2310031120131210-2113231302023112-1222302000123023-1101022303131223-3002011223210222-0101310212332331"></a>

<a id="canonical-3303133132032331-1031232223122312-2331023200303202-3333103032333220-1003321111130333-2022303302120101-1213110310323203-1113022302100100"></a>

## decryption_provider property — blindfold_secret_info / 233203112010 / 4

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

<a id="canonical-1003221012303232-1222013221030303-2020113122130211-2322212002002133-1131111002032003-0302123331322131-3033212330231330-0322113020131300"></a>

<a id="canonical-2133311313120231-2103210012120321-0131223111311210-2130113010122020-1322200213133120-1231131122120301-2322222122011232-1332122303221031"></a>

## location property — blindfold_secret_info / 233203112010 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3110212323130123-3121320311010202-1201103020021211-2333121323312300-2102201223211113-2103331002020330-2003123210121313-1300101120112022"></a>

<a id="canonical-1022003002203020-1220111203221120-2033010002330131-2121113230111210-0121003303010010-2132003213103321-2010023300230012-1122220203022221"></a>

## store_provider property — blindfold_secret_info / 233203112010 / 6

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

<a id="canonical-0021332030023122-2230220323302101-0312120131230221-1121332233313102-1122233221332332-3021221033301311-0112030302331103-2003110132321122"></a>

## Next pages — blindfold_secret_info / 233203112010 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--fleet--reference--group-003.md#canonical-2021123321033031-2131220302333132-2013232223201100-1211122103213311-1123311311323332-1333300211323100-3332210130000322-1110303130220221)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1321101120300120-2212323121331132-3201213113030033-2101133022330202-0110030321310301-1333313310331212-2011120330113131-2030020001022003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213303123203202-3033122032013031-2201010021200023-0111233231311322-0033023013232130-1322020101121331-1102123130033030-2213011103313103"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info — clear_secret_info / 202030312211 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--fleet--reference--group-003.md#canonical-2021123321033031-2131220302333132-2013232223201100-1211122103213311-1123311311323332-1333300211323100-3332210130000322-1110303130220221)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info

<a id="canonical-1301333012002230-1211121232220303-1001132330013210-1110312130033102-2111003110212300-2213310312211100-1210130322013100-1303032200332220"></a>

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

<a id="canonical-3020300310102131-0112201230203132-1223100231203113-0012133223221130-2330133103313202-0313030102032102-2333222323002012-2312130233313323"></a>

## Direct properties — clear_secret_info / 202030312211 / 3

<a id="canonical-2212230131002122-2032211110233122-0111313302333023-1012210200303320-3113023123123120-0221220202333122-0132330021020022-2221110303232020"></a>

<a id="canonical-2212231111221011-2203303111223203-0023121122132102-3220330001233022-2111330202331311-2202111030130311-2131211001123220-3320132301002133"></a>

## provider_ref property — clear_secret_info / 202030312211 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3303021030021130-0302233013003223-2020021202013333-2222030310200212-3311333212013013-2032302133233233-0132212230300331-2310000023023131"></a>

<a id="canonical-3100031120321213-2033301002101230-3301031322310130-2333100001121200-0123120313031022-3211102221121223-0201233323001301-2213202022300310"></a>

## URL property — clear_secret_info / 202030312211 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3033012330033023-2220010022333023-2210222310201212-2320031310010332-2013322100331111-3223200033110312-1301002310221101-0203031220331203"></a>

## Next pages — clear_secret_info / 202030312211 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--fleet--reference--group-003.md#canonical-2021123321033031-2131220302333132-2013232223201100-1211122103213311-1123311311323332-1333300211323100-3332210130000322-1110303130220221)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1300320120013220-1111311020033332-2110201300332223-1330230012212101-3132130210000210-0310202102232330-2230320110022220-0211021223113033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212222021213321-0021333320132123-1033312323221231-1211331212213122-0321221213203113-2001231321221323-3220330320213210-0223213031321303"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password — password / 212301200031 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password

<a id="canonical-2310201301301300-3100121212331033-2232033100203311-1330302310003123-0232233021333003-2200330123100132-3300122313002212-2100311033121231"></a>

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

<a id="canonical-2231133121011120-3303331021201232-3100212103111321-1200102311210332-2001312021302211-3100100122231123-1102000121230311-3013033020331221"></a>

## Direct properties — password / 212301200031 / 3

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-0313031320220001-1022121101320112-3101002202203311-0230200323122201-0100310223000021-2112000302223133-3231231321013323-1120320102230330): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-0103103020330320-2211332020322103-0111232101323212-2320200313030001-1201233210331002-3331223222012101-2133203032230302-1232113003313111): complete subsection reference.

<a id="canonical-1011112133002211-2310112302300000-1121211022100021-2032000233210032-3223131011133202-2002021223320112-1002131321123121-3331111230103213"></a>

## Next pages — password / 212301200031 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-0313031320220001-1022121101320112-3101002202203311-0230200323122201-0100310223000021-2112000302223133-3231231321013323-1120320102230330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info](resources--fleet--reference--group-003.md#canonical-0103103020330320-2211332020322103-0111232101323212-2320200313030001-1201233210331002-3331223222012101-2133203032230302-1232113003313111)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0313031320220001-1022121101320112-3101002202203311-0230200323122201-0100310223000021-2112000302223133-3231231321013323-1120320102230330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230212230222022-2312020221012010-1021101031330101-2013110211230213-0223110233302201-2332001101012030-3230021313302212-3010223201220333"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info — blindfold_secret_info / 101103302312 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--fleet--reference--group-003.md#canonical-1300320120013220-1111311020033332-2110201300332223-1330230012212101-3132130210000210-0310202102232330-2230320110022220-0211021223113033)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info

<a id="canonical-1320123220002003-2203232200231032-3301031122212302-3122003300332103-3303331233110203-1302223022222111-0030311012311010-0120002332101312"></a>

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

<a id="canonical-2010230201013011-1210311131131210-2023021211132303-1002210023320000-2103323000133301-1032230000020230-3310320032223323-3013001013212011"></a>

## Direct properties — blindfold_secret_info / 101103302312 / 3

<a id="canonical-0012222303202103-1030022203232022-1031300000330331-3322130231130130-1011212201110012-3311100003331202-1323301110333310-0112212110323021"></a>

<a id="canonical-0032201133011022-2210203012030202-3101030313201211-1312231232132022-3023210310110200-3203210332210310-1130122002020013-0202221003330322"></a>

## decryption_provider property — blindfold_secret_info / 101103302312 / 4

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

<a id="canonical-2001103321320003-2233120032302001-0120133101032131-3032020323303122-3122021213013330-1202203022020033-3102100120030100-2211333311110001"></a>

<a id="canonical-2102112132020121-0322113230320111-0330012003330032-1311300220323202-2301131113310020-0311031323332323-0132003131032100-3012321300331220"></a>

## location property — blindfold_secret_info / 101103302312 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1321301321301330-0020122313331331-1032310130023223-0110301023103323-1133013313221230-3212331020012300-1311313200223303-3303202300011212"></a>

<a id="canonical-1132122001132130-0300201013200312-3121331032330333-0023222303320001-2233010032121100-1133321303331220-3121213112233230-3031010030311020"></a>

## store_provider property — blindfold_secret_info / 101103302312 / 6

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

<a id="canonical-3300022220310002-2131110221110002-1021030303022323-1002110121322100-3031202011112330-2123111102233321-3332031010101133-3220003331233223"></a>

## Next pages — blindfold_secret_info / 101103302312 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--fleet--reference--group-003.md#canonical-1300320120013220-1111311020033332-2110201300332223-1330230012212101-3132130210000210-0310202102232330-2230320110022220-0211021223113033)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0103103020330320-2211332020322103-0111232101323212-2320200313030001-1201233210331002-3331223222012101-2133203032230302-1232113003313111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230030001200103-1222333322233332-1332023020323020-0113233110130023-2022031220231033-1300012321233333-3133301211112200-1130101311121310"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info — clear_secret_info / 120230123222 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--fleet--reference--group-003.md#canonical-1300320120013220-1111311020033332-2110201300332223-1330230012212101-3132130210000210-0310202102232330-2230320110022220-0211021223113033)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info

<a id="canonical-3233010122021111-2130130130011012-1321232312010131-3002100011220120-2133022111032011-3021203301111201-0103110110303311-0130302113321132"></a>

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

<a id="canonical-3112213231320111-1120012320302111-2313212131022223-0233220312320303-1220131013002031-1101302023301023-0220113112330213-3012031100232110"></a>

## Direct properties — clear_secret_info / 120230123222 / 3

<a id="canonical-0023330102210212-2031110120200133-2030201223100222-0230332112020120-1013120322220231-1033220312320023-0103000200203101-1131230322332032"></a>

<a id="canonical-0122210003023032-2131021310303130-2100310110033030-3212331230112233-0130023103300123-2032333213322001-0131121333010203-1312033312330003"></a>

## provider_ref property — clear_secret_info / 120230123222 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0133202232200021-2122231020032100-3321302331221321-3211120221131301-3301322333210300-3332203210210032-0313102013122232-1332032021323233"></a>

<a id="canonical-3112321110301223-3103031210202303-0111331101010200-2022312011010310-2301310232310300-1302221231031200-1311133002130003-0000021111300230"></a>

## URL property — clear_secret_info / 120230123222 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0033310020111302-0003001010301132-0011201110103121-2010303222323313-0222202111112323-1103220310010131-2212303231123021-2031112313122021"></a>

## Next pages — clear_secret_info / 120230123222 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--fleet--reference--group-003.md#canonical-1300320120013220-1111311020033332-2110201300332223-1330230012212101-3132130210000210-0310202102232330-2230320110022220-0211021223113033)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0232211313301321-0212033213020113-3120021131102001-2231330003231332-2031221231130011-2302223013200320-3030323232233300-2202032103012300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122113201220021-1102201320331231-3012101112322200-3222030111000313-3001012030033013-1222320300220311-0310203033021102-3133033310203333"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage — storage / 223101113200 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage

<a id="canonical-0323300303111030-2320210301123301-3031331000201323-2102132131230001-1220023300211123-0002211023213130-1220311130031020-3230300211333321"></a>

Type: `"object"`. list nested block, Optional.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221303032333333-3330012310232231-0122230310230022-0131112232333021-3132321200032010-1302011322310320-0330201011100313-2011010220300311"></a>

## Direct properties — storage / 223101113200 / 3

<a id="canonical-3113032131030330-3211313202123032-3130332011123121-3123100200302130-0300133332323112-1321320130121110-3022110213032121-3133322210232131"></a>

<a id="canonical-0000220012101023-0133133112110111-3233111003133322-1003120203020300-0033120120223103-2212213321300320-0000320132310302-1002110312301322"></a>

## labels property — storage / 223101113200 / 4

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":20},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"20\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--fleet--reference--group-003.md#canonical-2113320110201301-2101220231002013-3002022330010220-2021032231233120-3223323201302332-0302333022030323-1320122213233212-2012033130320012): complete subsection reference.

<a id="canonical-1111112113201000-1303200213203101-2032032301333011-1311133030021232-2323101133113011-3032022201201201-2111130221100323-3020103122303103"></a>

<a id="canonical-0231331202020301-3001202101103200-1322120211312211-0000221212323300-0022311020200023-1011202232301000-1123122232230210-1300230310221003"></a>

## zone property — storage / 223101113200 / 5

Type: `"string"`. Optional.

Virtual Pool Zone. Virtual Storage Pool zone definition.

Upstream description:

Virtual Storage Pool zone definition.

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

<a id="canonical-1012223123300233-2011233212113131-3320201132121000-0331113132310330-1221032230201323-0210211233331123-1203232102031022-0223303010122032"></a>

## Next pages — storage / 223101113200 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-2113320110201301-2101220231002013-3002022330010220-2021032231233120-3223323201302332-0302333022030323-1320122213233212-2012033130320012)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2113320110201301-2101220231002013-3002022330010220-2021032231233120-3223323201302332-0302333022030323-1320122213233212-2012033130320012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120000013111132-1320200020120213-2223123010002011-0103231010110211-0212210112130002-1010312103211131-3212002103311013-1232223211213120"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults — volume_defaults / 001130330130 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--fleet--reference--group-003.md#canonical-0232211313301321-0212033213020113-3120021131102001-2231330003231332-2031221231130011-2302223013200320-3030323232233300-2202032103012300)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults

<a id="canonical-0000132022102330-0231030023212231-2210011210103300-3232131333032210-3120303013311202-3023012311313331-2013301002010120-0032033232033201"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
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
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020313200212122-0332302301110001-3012210320122021-1331210012022001-2000100231211102-1333303312021102-2230023003313213-0320022221112230"></a>

## Direct properties — volume_defaults / 001130330130 / 3

<a id="canonical-1123313322201132-3201103312320230-2112321023323332-3010202112020231-2230023320323011-0200032231311022-1313231223302002-0201323330301231"></a>

<a id="canonical-3033113223312212-1330223201222120-2213310021301310-2210100211110021-1021102200323103-1101221301022102-0213312131100130-3330021013020322"></a>

## adaptive_qos_policy property — volume_defaults / 001130330130 / 4

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0330202233301311-1330201101233130-2220301010032200-1320002200310331-2313012333031101-0332031113033333-3321200312131200-2211311300002132"></a>

<a id="canonical-1320330323010133-2231202111213120-1200231231310233-0033213232002120-2311220323211230-1223321301301003-0100012003101330-3231310332310101"></a>

## encryption property — volume_defaults / 001130330130 / 5

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2113330322231332-0002120111033202-0232012320313313-1320030220012233-0010110320120211-1033203000212023-3003031202332313-3112331100022131"></a>

<a id="canonical-1330311131302032-0233032213313202-0131123330121310-0100312302313203-0130221033122110-3223210021033300-3310002133033032-0110311021010011"></a>

## export_policy property — volume_defaults / 001130330130 / 6

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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

- [no_qos](resources--fleet--reference--group-003.md#canonical-2313213003113030-0311202130033012-1012330231303212-2313221100211133-3311200223200332-3013012330201112-1030203113303131-2010233320301020): complete subsection reference.

<a id="canonical-3100311033130012-1330212011302313-0200201110132231-1220133130000032-3130332230222112-3310301121023133-1220300332232203-2231122220320232"></a>

<a id="canonical-1100203323102132-2021030310230001-3331033133130203-1122233032302013-2133321131220313-1033231132201013-2320230303102323-2330302100000212"></a>

## qos_policy property — volume_defaults / 001130330130 / 7

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3003310223321330-1332322213010222-1020201223033002-3233003103223001-0120020231233311-0202101321002001-1020020233221312-1302123333131101"></a>

<a id="canonical-1132302112100013-3332311121300132-0000320100020323-3232200133222100-0213101022113312-0302011301010201-2303122031132132-3120013120010102"></a>

## security_style property — volume_defaults / 001130330130 / 8

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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

<a id="canonical-1320011202312013-2113302320312302-2010221001130331-0033120100033000-1130231311321000-3102210112231002-2012032212022130-0233312020303213"></a>

<a id="canonical-3021203230113233-3131122022112011-2000330232201101-2130313101111320-0021332231323230-3303013333311030-1021020223120033-2031011300320233"></a>

## snapshot_dir property — volume_defaults / 001130330130 / 9

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1311320330221233-0121331102122101-1333233233330300-2230132222032330-3320023113301320-2003201101103030-0122102131313332-1232222032330132"></a>

<a id="canonical-3331000301132000-3133302313130303-1120100033230001-2022213312100333-1022033220312113-1123013302003321-2103202230111203-1200232232330213"></a>

## snapshot_policy property — volume_defaults / 001130330130 / 10

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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

<a id="canonical-0321101321233220-2333320230121301-0113301002332022-2101121111033022-0203110011202231-3110330303303123-2022122131222231-0003022301022130"></a>

<a id="canonical-1323311330012121-0322301131202223-1120111111203011-3301311133231333-2301122130032020-3312212102202330-0130323100113010-0011103220221010"></a>

## snapshot_reserve property — volume_defaults / 001130330130 / 11

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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

<a id="canonical-1021021101001221-0200103022202113-3302100233210031-3302013022322301-3131231330323130-3112200233020000-2002202232320211-3210230303322003"></a>

<a id="canonical-1100000031020333-1303021130332110-3213132111002203-1101200020213223-1133200200001230-1122003033330133-3032300220032033-1002201110002120"></a>

## space_reserve property — volume_defaults / 001130330130 / 12

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-2323101023330100-0201201012330331-1022232122101132-2133130323121332-2321330031110203-1330332310211000-3012102123232202-1213022303011320"></a>

<a id="canonical-2233330103211121-0220013212203220-0002310322310331-0303111001101321-0022220302023110-1322100010023320-0112323222133033-2001113202331313"></a>

## split_on_clone property — volume_defaults / 001130330130 / 13

Type: `"bool"`. Optional.

Split a clone from its parent upon creation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3100021310223013-1023112233120302-2210203030020022-3030200301112000-1222010313313022-3333130133030001-3312101130231011-3332300300113000"></a>

<a id="canonical-3001212121301310-2231012100103001-2112223200000133-3001012231102132-2332031003223323-0100220023331333-1202121112220321-1020201232313012"></a>

## tiering_policy property — volume_defaults / 001130330130 / 14

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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

<a id="canonical-2322020110010321-0033032323201203-2002323010021303-2330031222120100-1300330033230020-3102212133332221-0302310011312332-2331113333130203"></a>

<a id="canonical-3120013121211313-1023212332331233-3101001011021113-0311103303123203-0000332312223131-3112302233213102-2212302323020211-2222103220000231"></a>

## unix_permissions property — volume_defaults / 001130330130 / 15

Type: `"number"`. Optional.

Unix permission mode for new volumes. All allowed 777.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1300020123011212-0132121100312132-2000333131330222-0302210002233101-0122033001120031-1300313002200302-2322202110033323-3033011020022112"></a>

## Next pages — volume_defaults / 001130330130 / 16

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos](resources--fleet--reference--group-003.md#canonical-2313213003113030-0311202130033012-1012330231303212-2313221100211133-3311200223200332-3013012330201112-1030203113303131-2010233320301020)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--fleet--reference--group-003.md#canonical-0232211313301321-0212033213020113-3120021131102001-2231330003231332-2031221231130011-2302223013200320-3030323232233300-2202032103012300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2313213003113030-0311202130033012-1012330231303212-2313221100211133-3311200223200332-3013012330201112-1030203113303131-2010233320301020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221232213112223-2012130022322120-1332020000033200-0302033330212220-1122123231313211-2111233012110013-3000030010313210-1331222312031120"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos — no_qos / 133301100320 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--fleet--reference--group-003.md#canonical-0232211313301321-0212033213020113-3120021131102001-2231330003231332-2031221231130011-2302223013200320-3030323232233300-2202032103012300)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-2113320110201301-2101220231002013-3002022330010220-2021032231233120-3223323201302332-0302333022030323-1320122213233212-2012033130320012)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos

<a id="canonical-1111301331220233-3312003223022320-1323212032312332-0221033022323022-2221102121311121-3100031312312222-1233221313001301-0113103201310023"></a>

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
no_qos = {}
```

<a id="canonical-3013032023101310-1023222212331103-1330213333121001-0203010330301012-2100222120122030-3300131332133312-0132003033300313-2011331330113012"></a>

## Direct properties — no_qos / 133301100320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123220210332103-0132012200031232-3121121122222210-1013003211321312-0103100303123003-3030200132022001-0222013203331232-3023122030311300"></a>

## Next pages — no_qos / 133301100320 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-2113320110201301-2101220231002013-3002022330010220-2021032231233120-3223323201302332-0302333022030323-1320122213233212-2012033130320012)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2331201210221021-2100203101103023-3231010232113332-2331001322002230-3311121030003303-0210021021322130-2210111331110130-1031002103213331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003110021311010-1323102113020201-2122122130300100-3002331311003232-2121323010002100-0033212003030222-3231222311023111-1130301023113120"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults — volume_defaults / 121112230302 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults

<a id="canonical-0232310013002023-1113101303132333-1223201012331133-2300100203120113-2120003011000200-0120111033032303-3213002011031220-0003323310003012"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
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
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301203210022330-2332233113313322-1323313333002003-2022123021201133-2313101100221331-3223111020311330-3000300111133201-2232120111210211"></a>

## Direct properties — volume_defaults / 121112230302 / 3

<a id="canonical-1301310112321123-1103101102100223-3100202111102313-0202312230332100-2131202310321331-2222012023131302-1031000111121200-3300130223310233"></a>

<a id="canonical-3311211231212030-1301333332010011-0121132123331220-0310220320132010-2132100131330111-0202100233032230-0313322202200003-1313222122110222"></a>

## adaptive_qos_policy property — volume_defaults / 121112230302 / 4

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3221220112110001-3210323312300332-3302333312120201-2002213222212033-3212132112031003-2232002033333010-1312213200011310-2331002003110322"></a>

<a id="canonical-3200202200202002-1013001210303012-1230311331222023-2211321233030133-3031000301223013-2131201310211233-1210103221121000-2110003122113013"></a>

## encryption property — volume_defaults / 121112230302 / 5

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3313032321030203-1012101202311213-3012030233011233-1122321112330110-1130311331233320-0030300011320130-0001333122033201-2312300201123302"></a>

<a id="canonical-1323312112310020-2031301111123120-0321300002213110-0233200231112300-3130233221300210-2211222202301213-1201303332230003-0003023232113033"></a>

## export_policy property — volume_defaults / 121112230302 / 6

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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

- [no_qos](resources--fleet--reference--group-003.md#canonical-2233101330131022-2300110220203011-1230120022130033-2311302133323021-2033232112230323-3120331121123302-0131113211011002-0020113231123232): complete subsection reference.

<a id="canonical-1323211230012022-3210030032230001-0121213131022133-0331222030031312-1210123012231023-0220320201230323-1230232023110013-3231022112022120"></a>

<a id="canonical-0220001101333031-1102211101300032-1332102110312213-1333102331302122-2133131220003010-3321222013110103-1303011121122303-0230003320110303"></a>

## qos_policy property — volume_defaults / 121112230302 / 7

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3000210100233320-3030321021200212-2101320103302230-1330101123200012-2003021123001123-2211111001330002-0130300331332130-1223122010322122"></a>

<a id="canonical-1102111232110331-1002030321011003-0203002110012003-3220131021300112-2033320012022001-0313310211313011-0211201010233333-0321032232121312"></a>

## security_style property — volume_defaults / 121112230302 / 8

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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

<a id="canonical-1211301000102130-2231000200120211-2303313003120220-1120021301230112-0032110300130313-1031332120123030-3132010230120000-2332203302323122"></a>

<a id="canonical-1232323300030302-0112233223221131-3132301122322112-1331023011123300-1312223233022012-3022321101301222-1231102333121303-1312203000011101"></a>

## snapshot_dir property — volume_defaults / 121112230302 / 9

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1121311121212002-1200120212002203-2011102120031103-3003230220210022-3031132232221101-0330323010110301-0031333103330131-2123212111322232"></a>

<a id="canonical-0022300111003121-1123203003120103-2030113131311333-1003001203032310-2330331122233200-3312111123231330-1322300313003200-0002002220102112"></a>

## snapshot_policy property — volume_defaults / 121112230302 / 10

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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

<a id="canonical-1120023203110310-3021011102210320-2102213100322100-3330210232013310-3312112123230333-3332313203220000-3201202133233133-0203332021312010"></a>

<a id="canonical-3000212000132301-1131211111001103-0102223031111002-1201303110303220-1202120213321111-1332230132222332-1302113200201011-1120010300002022"></a>

## snapshot_reserve property — volume_defaults / 121112230302 / 11

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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

<a id="canonical-0322200131020031-2002313022011012-3303323230201123-2210023303232110-3000230003033033-2312233331230302-1032230031232220-3101130200313311"></a>

<a id="canonical-3101031312003313-3113130133320110-2112300232011221-3013220231323332-3313013122223003-1011203202132312-2210303133310001-2101333202001110"></a>

## space_reserve property — volume_defaults / 121112230302 / 12

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-1132321301130121-2103311321210130-1032011211333013-1002211123312233-0302122032010210-2210313211023200-2000232111103331-0311122212123232"></a>

<a id="canonical-1333200221031132-1120200121010203-2211102231032300-3210110310320200-2231013001113320-1132030323311203-1321330302120302-2031312230320323"></a>

## split_on_clone property — volume_defaults / 121112230302 / 13

Type: `"bool"`. Optional.

Split a clone from its parent upon creation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1301212333212322-2102120103213022-3201002132131123-2120133302011010-2132100320233101-2100202331112211-3123011111112122-2212213303320211"></a>

<a id="canonical-2322000311211301-0211121032333002-1120203030320011-0202213332030310-3211113022331232-0123000331110133-3210113302032031-1300013032232302"></a>

## tiering_policy property — volume_defaults / 121112230302 / 14

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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

<a id="canonical-0223111201001002-0002122121223333-0020002021122221-0312300022013313-3301031030332211-3212010200321132-3200130031300213-3001322112303330"></a>

<a id="canonical-3301033023322302-1301210112303302-0003323131033211-2210300212032211-2230120001120031-0302122233221332-3111133000220020-0233123301321312"></a>

## unix_permissions property — volume_defaults / 121112230302 / 15

Type: `"number"`. Optional.

Unix permission mode for new volumes. All allowed 777.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3001000010300321-0133121020132020-3312011202103321-1112303221303012-3203222130122231-3332022103110202-0120233101020130-1213221123022130"></a>

## Next pages — volume_defaults / 121112230302 / 16

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos](resources--fleet--reference--group-003.md#canonical-2233101330131022-2300110220203011-1230120022130033-2311302133323021-2033232112230323-3120331121123302-0131113211011002-0020113231123232)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2233101330131022-2300110220203011-1230120022130033-2311302133323021-2033232112230323-3120331121123302-0131113211011002-0020113231123232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110223131023221-1211111300002033-3030131111223102-1200310120011313-3020032321110321-1212000332130021-0030103323020212-3013332122320001"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos — no_qos / 001110103210 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](resources--fleet--reference--group-003.md#canonical-2331201210221021-2100203101103023-3231010232113332-2331001322002230-3311121030003303-0210021021322130-2210111331110130-1031002103213331)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos

<a id="canonical-3002123320300132-0223331021300220-1210103031121331-2123031323213111-0011011022012122-0212310103013013-1233320313131233-0232112111213103"></a>

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
no_qos = {}
```

<a id="canonical-2300113222112010-2101211111313203-0221302100103322-3133312113010231-2112013212111112-0211100102122023-1033022011120303-3221111131303310"></a>

## Direct properties — no_qos / 001110103210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111203222131032-3110132030223310-1313133133233331-3110231122133310-1003221222133113-1230032231011321-0211322323223001-0231110112201310"></a>

## Next pages — no_qos / 001110103210 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](resources--fleet--reference--group-003.md#canonical-2331201210221021-2100203101103023-3231010232113332-2331001322002230-3311121030003303-0210021021322130-2210111331110130-1031002103213331)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103132203202303-1232222020312202-3200020233210213-3122202011111122-1200223220300202-0220003122223111-3010120111323330-1230001020222133"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san — netapp_backend_ontap_san / 301033021111 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san

<a id="canonical-3323210230230230-2010001013033130-1322132130110013-2001110133212220-1213311130133212-3130321202013113-0323121301331013-0123230022003303"></a>

Type: `"object"`. single nested block, Optional.

Configuration of storage backend for NetApp ONTAP SAN.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_driver_name",
    "username"),
  validators.ConflictingObjectAttributes("data_lif_dns_name",
    "data_lif_ip"),
  validators.ConflictingObjectAttributes("management_lif_dns_name",
    "management_lif_ip"),
  validators.ConflictingObjectAttributes("no_chap",
    "use_chap")}
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
  "x-ves-oneof-field-chap_choice": "[\"no_chap\",\"use_chap\"]",
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

Terraform syntax:

```terraform
netapp_backend_ontap_san {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123312221310201-3011302002112201-2122000102122112-1213320321233130-0311112220010320-0030013110030332-3010300301030230-2200021031120311"></a>

## Direct properties — netapp_backend_ontap_san / 301033021111 / 3

<a id="canonical-1333220200332203-1102231001200321-2000221232122022-0220130302101023-2200231001302100-0232000312111132-2323131321132010-2132133033011210"></a>

<a id="canonical-0123233233021012-3223210300020220-2223333331320023-3223101130020132-2033303112232301-3301123111331131-3013333230231231-2332321020310312"></a>

## client_certificate property — netapp_backend_ontap_san / 301033021111 / 4

Type: `"string"`. Optional.

Please Enter base64-encoded value of client certificate. Used for certificate-based auth.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](resources--fleet--reference--group-003.md#canonical-2232002211332213-2313133232203222-2021113333022113-1101102033001221-0113320231122221-0233101120302221-1123133323322012-0101112100013321): complete subsection reference.

<a id="canonical-2002111031200132-3311123011032112-0110322030021311-0022332123323013-0011331321131221-2231102003123131-2122230212233012-0110011212102121"></a>

<a id="canonical-2121130322010302-2111320232320233-0101103200220313-0303100001003033-2230110223132012-0130213120011131-2210030103330230-3133332312100100"></a>

## data_lif_dns_name property — netapp_backend_ontap_san / 301033021111 / 5

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2012111013231313-2100331100002022-1320213312331231-3222113322021232-2110300200223101-1333331120312331-3313231120310200-1322001011321300"></a>

<a id="canonical-1301033303102323-3001302021111221-1230323122113002-3020320222200011-0212211200001102-1200223222021030-0030133312130331-0013022300232122"></a>

## data_lif_ip property — netapp_backend_ontap_san / 301033021111 / 6

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2012011322130221-1123031013002011-2000133311221101-2123111023203221-1323332001233002-2223010010032112-3021000003310000-2111333002013232"></a>

<a id="canonical-1200011333003332-1233313231100023-2031102123221131-0202311301032101-1023112312313001-2110132202001231-0132020032230310-1123331231022122"></a>

## igroup_name property — netapp_backend_ontap_san / 301033021111 / 7

Type: `"string"`. Optional.

Name of the igroup for SAN volumes to use.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1233203030100212-0313212330132230-2231021101303301-2310213112123011-0033133021101101-2300112212003123-0003120112302133-0101311231131222"></a>

<a id="canonical-0233320101213231-0132312101102323-3033003013122313-2332333232012201-0212102113322132-2002003113022111-2320132313023113-2100310013030122"></a>

## labels property — netapp_backend_ontap_san / 301033021111 / 8

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":20},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"20\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0131320010213210-1220000033003302-2312203332003003-0310033320002123-1000001013000200-3230032223120022-3033321322222013-1013300112103200"></a>

<a id="canonical-1210213312120031-1302111112200131-2103000000211303-1210113231332221-1011020031021323-3110110022033100-3103312013220021-1102123220000332"></a>

## limit_aggregate_usage property — netapp_backend_ontap_san / 301033021111 / 9

Type: `"number"`. Optional.

Fail provisioning if usage is above this percentage. Not enforced by default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 100),
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
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-3122211201223003-3021303220123210-0331301012310100-3103110210011103-2133220222313213-1300331312220101-2311131102010323-2330002322220121"></a>

<a id="canonical-0020023032311103-1211022332023231-3000202010211311-2111002033102023-1123123000321303-0012132113301010-3003203330033021-0033000213023213"></a>

## limit_volume_size property — netapp_backend_ontap_san / 301033021111 / 10

Type: `"number"`. Optional.

Fail provisioning if requested volume size in GBi is above this value. Not enforced by default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0213102023320321-3313120133230223-0212213103112300-2103011130203320-2022033012100132-3122321200202302-2221130122322220-0323302313330002"></a>

<a id="canonical-3213012231310323-2301033323222101-0232211221210023-3001311021323201-3031231002200333-1303112120212123-3331030022201203-1312030331232121"></a>

## management_lif_dns_name property — netapp_backend_ontap_san / 301033021111 / 11

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3321302332003231-2020020100012231-3102013012200321-3200033233110313-2202220203213011-0031320333333023-2202001110113300-2210223223310110"></a>

<a id="canonical-3333132213110001-1211231131311231-1000213300103130-2001121323033023-3220330303203112-3210332133132120-2203331033232322-3121123210202233"></a>

## management_lif_ip property — netapp_backend_ontap_san / 301033021111 / 12

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

- [no_chap](resources--fleet--reference--group-003.md#canonical-2023212231132001-0233233303010331-1322332333230000-3022312201100210-0013202332122032-1001132113122332-0313002113102131-2232330213122212): complete subsection reference.

- [password](resources--fleet--reference--group-003.md#canonical-0332110131332013-0220121311102123-0210223010320022-2100123103202201-0131201101033031-0221312322121031-1323200232310111-3120331210320131): complete subsection reference.

<a id="canonical-3113021301202103-1120203201022013-3110131321022332-2321010111201202-0220123202230130-3133102021113222-2333233013003331-3321202212303110"></a>

<a id="canonical-2203001020003013-0233220031221212-1202110030311232-1031011102033321-0002221210121311-0322131032031111-1001002333110202-2103133103030321"></a>

## region property — netapp_backend_ontap_san / 301033021111 / 13

Type: `"string"`. Optional.

Backend Region. Virtual Pool Region.

Upstream description:

Virtual Pool Region.

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

- [storage](resources--fleet--reference--group-003.md#canonical-0132123032123033-1300031111020020-0122303313210220-0301213233313102-3000331003132200-0232211022231113-3313132300030110-0123333132302302): complete subsection reference.

<a id="canonical-1302220010102231-2313311230120312-0112100020331100-0031021323322033-2010110021211021-2302020303002012-3330323302221112-1300212220332022"></a>

<a id="canonical-0133001210331130-2013010330311230-1213213232210221-1322322212133123-0300212221221300-3102302312101223-0111013202231320-2302300011023303"></a>

## storage_driver_name property — netapp_backend_ontap_san / 301033021111 / 14

Type: `"string"`. Optional.

\[Enum: ontap-san|ontap-san-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-san\`, \`ontap-san-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ontap-san",
    "ontap-san-economy",
    "ontap-nas-flexgroup"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-san",
    "ontap-san-economy",
    "ontap-nas-flexgroup"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-2201002330330333-2002100130001202-1132110133003112-3313102201121221-1002332003103013-3111001332130233-3321230200111013-3100303203111031"></a>

<a id="canonical-2213032000303103-2320203211100211-3210023102100103-0100011103303033-2130111103100012-0230320022001121-1333232000311220-0101323303233331"></a>

## storage_prefix property — netapp_backend_ontap_san / 301033021111 / 15

Type: `"string"`. Optional.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 80),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 80,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 80,
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
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1112213220332103-2331030233202121-3221310322030302-2203033311101133-2103220223012231-3222232222000313-2230203211003112-2022220020030332"></a>

<a id="canonical-1132301303111332-3221313002031123-0020323023231201-0100020101103223-3330330112121331-1011101030200023-1002013001100202-0200111321133313"></a>

## svm property — netapp_backend_ontap_san / 301033021111 / 16

Type: `"string"`. Optional.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1030122212211303-1033122101021312-1012213300312023-0001002122203202-3300123202301131-3100003011133013-3321313131100310-2022013213311032"></a>

<a id="canonical-1113212210130203-1311212000223311-3220103101212002-0013322012200102-2200131102020113-0211211212010211-2220020112321203-3201220303211010"></a>

## trusted_ca_certificate property — netapp_backend_ontap_san / 301033021111 / 17

Type: `"string"`. Optional.

Please Enter base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [use_chap](resources--fleet--reference--group-004.md#canonical-0210223030221013-1133322303123113-0112201303102112-2213010000130321-2133111120133330-2233132011002112-3313333201032133-3311001320131013): complete subsection reference.

<a id="canonical-1022321112232130-0311002130232210-1330122220010222-1333111013033232-3120012303123320-3131322011310001-3021330000220020-3022323320101300"></a>

<a id="canonical-3231020310201213-1011001301211132-0310223303232012-3101030201012330-3311322012032221-3202000302203003-0300123211023001-2112203210003011"></a>

## username property — netapp_backend_ontap_san / 301033021111 / 18

Type: `"string"`. Optional.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

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
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--fleet--reference--group-004.md#canonical-1023300110000030-3322320012222203-3312211121033330-0233123000012211-0313012003121021-3303103100101231-0022013323332110-0101320001330221): complete subsection reference.

<a id="canonical-3001221213223131-3001200102222330-1300231333101102-1231131003223202-2122321213000000-3323221133333311-2313200033100101-1112332302031233"></a>

## Next pages — netapp_backend_ontap_san / 301033021111 / 19

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--fleet--reference--group-003.md#canonical-2232002211332213-2313133232203222-2021113333022113-1101102033001221-0113320231122221-0233101120302221-1123133323322012-0101112100013321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap](resources--fleet--reference--group-003.md#canonical-2023212231132001-0233233303010331-1322332333230000-3022312201100210-0013202332122032-1001132113122332-0313002113102131-2232330213122212)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--fleet--reference--group-003.md#canonical-0332110131332013-0220121311102123-0210223010320022-2100123103202201-0131201101033031-0221312322121031-1323200232310111-3120331210320131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--fleet--reference--group-003.md#canonical-0132123032123033-1300031111020020-0122303313210220-0301213233313102-3000331003132200-0232211022231113-3313132300030110-0123333132302302)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-004.md#canonical-0210223030221013-1133322303123113-0112201303102112-2213010000130321-2133111120133330-2233132011002112-3313333201032133-3311001320131013)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](resources--fleet--reference--group-004.md#canonical-1023300110000030-3322320012222203-3312211121033330-0233123000012211-0313012003121021-3303103100101231-0022013323332110-0101320001330221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2232002211332213-2313133232203222-2021113333022113-1101102033001221-0113320231122221-0233101120302221-1123133323322012-0101112100013321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311301232221000-1233223210212311-0112023001133311-2122330302122110-0112023332100223-1100223133033332-1320233302222203-3320012033330003"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key — client_private_key / 203121332033 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key

<a id="canonical-1333023203233113-2220031121323220-1231121203300023-0122013010232003-1220202332213131-1103300033133200-1102230102012002-2033131210130003"></a>

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
client_private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011222010021031-3131222212232111-2010030320313221-3130000210203310-2333110032222301-2332212333220302-2100112101013002-3132212032231011"></a>

## Direct properties — client_private_key / 203121332033 / 3

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-1220302221311032-2203010123221321-3213220210333022-3013321121313100-3230321211022333-2133120012002310-3012233113312320-3131330320030233): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-2133201323002230-0312012232123303-3112010321130320-1212102021103100-0300033231111230-0202002231132002-3332202023210003-3333313101202132): complete subsection reference.

<a id="canonical-3331120121222323-3211230221233313-3102022223311301-0202132100312110-0332313123103021-1312000311030200-1201133003000322-1010322212111110"></a>

## Next pages — client_private_key / 203121332033 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-1220302221311032-2203010123221321-3213220210333022-3013321121313100-3230321211022333-2133120012002310-3012233113312320-3131330320030233)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info](resources--fleet--reference--group-003.md#canonical-2133201323002230-0312012232123303-3112010321130320-1212102021103100-0300033231111230-0202002231132002-3332202023210003-3333313101202132)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1220302221311032-2203010123221321-3213220210333022-3013321121313100-3230321211022333-2133120012002310-3012233113312320-3131330320030233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301302100310122-2023311332313120-0131022212013221-3121322332203032-1103203130331020-0230020210010210-0120023202310012-3103323232331312"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info — blindfold_secret_info / 222001121011 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--fleet--reference--group-003.md#canonical-2232002211332213-2313133232203222-2021113333022113-1101102033001221-0113320231122221-0233101120302221-1123133323322012-0101112100013321)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info

<a id="canonical-1032101211332033-0323320121220133-0030003312332200-2302232033012231-1211320312213223-0122333220020121-1120010231313222-2211123111103123"></a>

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

<a id="canonical-1131211330002100-2301031103320113-0103100010220232-0202111123023201-0002023020031233-1202121001102303-0213313003231133-3320000202331200"></a>

## Direct properties — blindfold_secret_info / 222001121011 / 3

<a id="canonical-0002330000000333-2100323201331121-3002321101203313-3110102123200333-0020000313303130-3202021233232233-1320003211001021-1200110222210303"></a>

<a id="canonical-1131212210312223-3032230032022232-0010122030121313-1202213033111132-1332111302123200-1222223123201033-2323223001233100-2232010211322312"></a>

## decryption_provider property — blindfold_secret_info / 222001121011 / 4

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

<a id="canonical-2032213310223333-0013311123223100-3111132102303203-0003322311320032-2030321112030100-1320320130323200-2331310121031331-3330013320122121"></a>

<a id="canonical-1100100211001011-2021311022223132-0031121120332001-0321032132112110-0002113311232320-2132103323302201-2133112001331233-3313020131310320"></a>

## location property — blindfold_secret_info / 222001121011 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3011122020112110-3020031131323020-0202000100122231-0233121000100330-0133310023323323-2133232223202010-0331030103132223-1231200221330320"></a>

<a id="canonical-0010302000333001-0232031303210321-3321003113231102-1002113303212130-2111101332102111-3231223012331020-0311223010320131-0133212031113000"></a>

## store_provider property — blindfold_secret_info / 222001121011 / 6

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

<a id="canonical-3033121001110233-2012010020122133-3333230020001021-3302012010031221-0010331031123100-3033111003301112-0103223332103030-2120110021231300"></a>

## Next pages — blindfold_secret_info / 222001121011 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--fleet--reference--group-003.md#canonical-2232002211332213-2313133232203222-2021113333022113-1101102033001221-0113320231122221-0233101120302221-1123133323322012-0101112100013321)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2133201323002230-0312012232123303-3112010321130320-1212102021103100-0300033231111230-0202002231132002-3332202023210003-3333313101202132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211101332201222-1213032302032012-0311012231331003-0123232003031000-1020302220332030-1230001302132322-0313133010110201-3121220323233010"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info — clear_secret_info / 223232312221 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--fleet--reference--group-003.md#canonical-2232002211332213-2313133232203222-2021113333022113-1101102033001221-0113320231122221-0233101120302221-1123133323322012-0101112100013321)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info

<a id="canonical-2111210331213021-1030012213212013-0130210111021323-0011121310030321-0303010302333230-0100302222202033-3102121113230231-1100000232103210"></a>

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

<a id="canonical-0303300130131312-0002123202113010-2202213223022211-1203021222110030-2032311032030100-2223113120322103-1103032131032032-0311321111313211"></a>

## Direct properties — clear_secret_info / 223232312221 / 3

<a id="canonical-3131103210331203-2220030230322322-1121010313023300-2012302011020301-3012223303213311-0020331021021002-0213202202010123-2102302102301012"></a>

<a id="canonical-1022002303233120-1001123332233303-1000132023233311-3201012030333230-3122003021303321-0313001101310110-1332031023002000-1323220100131002"></a>

## provider_ref property — clear_secret_info / 223232312221 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1332321023113022-2321210211203213-0313312231212203-2230333131200033-1303333013320213-3113020002212121-2313230130000301-1302213313331211"></a>

<a id="canonical-3202320201011122-3011230030021322-0012010133300112-3121232220213231-1323123313233030-3032332322101123-0232220122231323-2122022212220001"></a>

## URL property — clear_secret_info / 223232312221 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3111121333021112-2020232311101221-1303103201120201-0003002122313322-3220300000220100-1302113033100330-0212302023332131-2003000230113012"></a>

## Next pages — clear_secret_info / 223232312221 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--fleet--reference--group-003.md#canonical-2232002211332213-2313133232203222-2021113333022113-1101102033001221-0113320231122221-0233101120302221-1123133323322012-0101112100013321)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2023212231132001-0233233303010331-1322332333230000-3022312201100210-0013202332122032-1001132113122332-0313002113102131-2232330213122212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311231333312011-2323200101301322-2330221213322010-2313221003201123-1110031131122222-3313112233312012-2103000221210111-2001000003011123"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap — no_chap / 123320033020 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap

<a id="canonical-2312231103301222-2303013202220233-2012103312203311-0302300223213112-1322111133033332-0310121011022233-1331102122201223-0120313021103321"></a>

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
no_chap = {}
```

<a id="canonical-0321120003003110-3210201331123332-3112003030202312-1000202220331232-1201011031323311-3131233320133130-0221122131112201-0111132230112000"></a>

## Direct properties — no_chap / 123320033020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122002212022212-0221133033133101-0000200120022000-3102322030133120-2302330320200132-0003130023331013-0301311320111023-2320103313120302"></a>

## Next pages — no_chap / 123320033020 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0332110131332013-0220121311102123-0210223010320022-2100123103202201-0131201101033031-0221312322121031-1323200232310111-3120331210320131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110011311121122-3201300032331321-1210110001321121-2323003032032210-0201101010233111-2032203012101311-0023220202223120-1223213202031020"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password — password / 102211010113 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password

<a id="canonical-0311220320002233-2031131013333113-1230200132010013-0120222121311103-3212313300023203-0320022012312101-2320233010131023-2113321121332123"></a>

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

<a id="canonical-2203321323033111-1001031030231331-3032323012102312-2220101213102230-0210122002211000-0012210013003031-1321310221312321-1002311222111233"></a>

## Direct properties — password / 102211010113 / 3

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-0303321000302203-3000100313100310-2330132312013223-1102030110232131-0000101330011313-2231230031010101-0232312030313121-1102002233112220): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-2202220231210203-2000203110211122-1331231222330313-3123032313032231-3022303121212123-1121112330122201-0022010131023200-0032300013101100): complete subsection reference.

<a id="canonical-1331120231202031-1302132211021121-2131232211303211-2200312303132031-1200021022112303-2012112031222231-0323321230300101-0102033133221010"></a>

## Next pages — password / 102211010113 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-0303321000302203-3000100313100310-2330132312013223-1102030110232131-0000101330011313-2231230031010101-0232312030313121-1102002233112220)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info](resources--fleet--reference--group-003.md#canonical-2202220231210203-2000203110211122-1331231222330313-3123032313032231-3022303121212123-1121112330122201-0022010131023200-0032300013101100)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0303321000302203-3000100313100310-2330132312013223-1102030110232131-0000101330011313-2231230031010101-0232312030313121-1102002233112220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233101120232002-1120132101310132-2030322000021113-3213123012021101-3213011303333012-3122003030112113-1313203023213013-2312010221123302"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info — blindfold_secret_info / 021130210031 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--fleet--reference--group-003.md#canonical-0332110131332013-0220121311102123-0210223010320022-2100123103202201-0131201101033031-0221312322121031-1323200232310111-3120331210320131)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info

<a id="canonical-0003322203023300-3012202302303003-1011322203212103-3310323103322303-0113121312302131-3113310101110330-1310000101320233-2103300022331221"></a>

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

<a id="canonical-1030012011221311-1233300023321310-3100310220303031-3113230312001100-3232013021333211-0133100103102031-1023110202231302-2030202012120332"></a>

## Direct properties — blindfold_secret_info / 021130210031 / 3

<a id="canonical-3030322023203110-1012123230012320-1323010223002002-3311300301131233-0310313202312322-1221112030332123-0020311202111011-0020321231023301"></a>

<a id="canonical-2313133222021012-0302330221012131-0002323203033021-1133203312302200-0111010002210223-2100111220332023-3133313200032003-3122320020133123"></a>

## decryption_provider property — blindfold_secret_info / 021130210031 / 4

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

<a id="canonical-3123003120331232-2033301223111112-0321030133113031-2231212012322030-3001101302101123-1103033210230203-2012230323123230-0310203333230223"></a>

<a id="canonical-0203323331023021-0031323122032000-3101100020232201-0013030113323110-1000021332311011-0002312333122223-2223201001020110-0312112033320300"></a>

## location property — blindfold_secret_info / 021130210031 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1300131113120232-0222132210030312-3023012320330301-3111101130300101-1022202121022013-3213032110331221-0011203330223113-0030301212123312"></a>

<a id="canonical-3320010121111003-1312030112311023-1323321010021302-3123320012121213-0021203303131312-3330203133233010-1131233200213100-2300222331113011"></a>

## store_provider property — blindfold_secret_info / 021130210031 / 6

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

<a id="canonical-1121102012001222-1233323211133011-0213001300101131-3123233320221300-3321101331212010-1231230101201033-0230100023120333-3231203321303103"></a>

## Next pages — blindfold_secret_info / 021130210031 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--fleet--reference--group-003.md#canonical-0332110131332013-0220121311102123-0210223010320022-2100123103202201-0131201101033031-0221312322121031-1323200232310111-3120331210320131)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2202220231210203-2000203110211122-1331231222330313-3123032313032231-3022303121212123-1121112330122201-0022010131023200-0032300013101100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302032112231113-3102310122221030-0223011121313221-3123032230232033-2102223311111121-2312121130230332-1031303001101322-0121033312021121"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info — clear_secret_info / 213230303232 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--fleet--reference--group-003.md#canonical-0332110131332013-0220121311102123-0210223010320022-2100123103202201-0131201101033031-0221312322121031-1323200232310111-3120331210320131)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info

<a id="canonical-3231023312310010-0111102221002032-2121103001212000-2333130132120032-1223001032131223-2330230220233221-3021000010120311-2003030221203101"></a>

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

<a id="canonical-3003100131030203-1313020102022202-0012111310200122-1002012223213221-0111102000330201-0303133211200030-3133031101002120-2310120122113131"></a>

## Direct properties — clear_secret_info / 213230303232 / 3

<a id="canonical-0130220311301111-3201001110102323-3311300222113302-2322132000223120-2330111012311320-3230003321112012-1301322313221213-2231333232022303"></a>

<a id="canonical-2132222303222310-0023101212130022-0233322013233331-3210100330102100-1013200131131213-3123121020311010-2302101300321303-0313310323203330"></a>

## provider_ref property — clear_secret_info / 213230303232 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0001230103223201-3103313212232313-3303202101333031-1332331332003120-2231101233330301-2332102301202003-3113223013332031-0323003011302313"></a>

<a id="canonical-3020031231322201-3011131210312202-2313033210231323-3110022211110011-1312103312110200-2201302102112222-2223131201120030-0123223011320000"></a>

## URL property — clear_secret_info / 213230303232 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2133011203111001-2223100333003001-0111313232003000-3032213132222221-0312220112020030-0023302020202030-3312212210210013-2201102201003021"></a>

## Next pages — clear_secret_info / 213230303232 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--fleet--reference--group-003.md#canonical-0332110131332013-0220121311102123-0210223010320022-2100123103202201-0131201101033031-0221312322121031-1323200232310111-3120331210320131)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0132123032123033-1300031111020020-0122303313210220-0301213233313102-3000331003132200-0232211022231113-3313132300030110-0123333132302302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232030311311221-3111320121321131-0301221111201203-3023330130013300-0121302323111321-0211333123010123-3103011120023112-2112222122311230"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage — storage / 231320003100 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage

<a id="canonical-0211032110321131-3312003313011133-1233331030301322-3312031133133231-0213230312110110-0032012133303202-2321030123303101-2333231131032301"></a>

Type: `"object"`. list nested block, Optional.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111322322231121-1310032223220110-3222110133211331-1332011010020220-1330002220303222-3032201220012213-2023120303001111-0133203023303120"></a>

## Direct properties — storage / 231320003100 / 3

<a id="canonical-1031323230331001-1101112001113021-3003130101130201-3212102133311110-3211331013210001-3131231313220311-3011323000230200-2031233023110202"></a>

<a id="canonical-3301102103131123-1110300333113012-1322321323033232-0132113123032323-2011202330332102-1133011000332001-3203130001202121-1202313032122123"></a>

## labels property — storage / 231320003100 / 4

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":20},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"20\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--fleet--reference--group-003.md#canonical-1331203201033220-3033100023220013-2010010023203222-2210330212120313-2310202223112112-2033230330001120-3220221121000221-1210120221222223): complete subsection reference.

<a id="canonical-2223300130001123-1020221130012330-2221203113311211-2110033211130301-1022321230301010-1032100301011002-0030213103210201-1220030110231322"></a>

<a id="canonical-3121102331300332-1310230112001023-1002111300112123-1223101322131222-0232102222023220-2020132021113201-1310203303121001-3203211330211010"></a>

## zone property — storage / 231320003100 / 5

Type: `"string"`. Optional.

Virtual Pool Zone. Virtual Storage Pool zone definition.

Upstream description:

Virtual Storage Pool zone definition.

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

<a id="canonical-0013013102210200-1132110120031300-2332120021001200-3132100232203102-1221333210210133-1231213213113303-0130213232122003-0330000121122222"></a>

## Next pages — storage / 231320003100 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-1331203201033220-3033100023220013-2010010023203222-2210330212120313-2310202223112112-2033230330001120-3220221121000221-1210120221222223)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1331203201033220-3033100023220013-2010010023203222-2210330212120313-2310202223112112-2033230330001120-3220221121000221-1210120221222223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313001022132122-1113121033300002-1220110332122001-0033212033033132-1033303222113212-0220202210030212-1102000230110230-0033301221012210"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults — volume_defaults / 130110113011 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-003.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--fleet--reference--group-003.md#canonical-0132123032123033-1300031111020020-0122303313210220-0301213233313102-3000331003132200-0232211022231113-3313132300030110-0123333132302302)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults

<a id="canonical-2301022122323310-3321013323230120-2303310023223021-0120321200233130-2001120302001031-2010011333030111-2013212312311212-3000313000023301"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
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
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013332323020033-3123121212110113-1213110310312221-2033203003112021-0023311301032212-2203123013112101-1303101033021112-1331123103001300"></a>

## Direct properties — volume_defaults / 130110113011 / 3

<a id="canonical-1022212002223031-2221112020301210-1303300102113000-1020303122020010-3100003102012010-1031233301131313-1110033022022003-2130313101303201"></a>

<a id="canonical-3021313212030223-3121213213302312-3100030120221130-2120012223222101-1221313003133101-0102103023312023-3332212131122030-3010011032021300"></a>

## adaptive_qos_policy property — volume_defaults / 130110113011 / 4

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0330330022322210-3301101000302032-2132023210033101-0223111032030301-1030330213312101-2232021313233320-2330102021023110-3021223023131222"></a>

<a id="canonical-0021322001333321-0022331022233211-1022221023012122-3213321133302121-0311320100322031-1322301211003212-1300203133330002-2232310200220032"></a>

## encryption property — volume_defaults / 130110113011 / 5

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2213021120201112-2003101113230331-0303213102200212-3002333313221013-1123212231202012-2122202023122230-2331112010023211-1210230113103131"></a>

<a id="canonical-2033202030030212-1222221122121313-1312331233221113-2100032112120310-2123113111131323-2031323003023023-0131311002013020-1320330323100211"></a>

## export_policy property — volume_defaults / 130110113011 / 6

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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

- [no_qos](resources--fleet--reference--group-004.md#canonical-0230132003233011-1302200101120032-3302332032133012-3132303201232302-1223221123003013-2101021331312201-1332131002010030-2012322020112120): complete subsection reference.

<a id="canonical-3013310231230110-0213300133130022-2011011333202031-0123013132133012-1223333022121122-1210300330233101-0323113022030300-2312011313303321"></a>

<a id="canonical-1112312221221002-1300222331201003-2312201310131110-3102003102112100-0023021101013310-3022121131201102-3132320101233132-1311322212310331"></a>

## qos_policy property — volume_defaults / 130110113011 / 7

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1310101232210213-1000222222131132-0133121323130303-0033233011121323-0210133022201023-2201322302232202-2303131002101211-2212100123203302"></a>

<a id="canonical-0002002221013302-3302110231222211-1110212111332210-1200133002003300-0311202133002212-3301330220033103-1322302310122202-1030312313312022"></a>

## security_style property — volume_defaults / 130110113011 / 8

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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

<a id="canonical-1303130223112010-2020311230033012-1010200302212311-1132131223113010-1000002311211003-0230333331023233-1010210321021131-3321031221021130"></a>

<a id="canonical-0110222310113022-0303310003033121-0230102030003022-3120033312221122-0223302032031212-3202030023330222-2220121022000010-1231020333003103"></a>

## snapshot_dir property — volume_defaults / 130110113011 / 9

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3122312112000331-2013032020102312-3301120311111221-0323122213002000-0331132232122123-2003222330221131-2333313311213132-3110321122023323"></a>

<a id="canonical-2223031211133302-3312321010021333-0322300202300303-1220013220022013-0101221201112111-3032222012112001-1100333130122131-1010220132223311"></a>

## snapshot_policy property — volume_defaults / 130110113011 / 10

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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

<a id="canonical-1110102000233101-1223012221010300-1222313202211233-1031320301000003-3213331313022303-3122201313201001-1023112303020202-2221100201010332"></a>

<a id="canonical-2112330222330300-1333210230313220-2333103311200000-3011230011200322-1201202012022031-2311300013031313-1330100201111203-0033213310122312"></a>

## snapshot_reserve property — volume_defaults / 130110113011 / 11

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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

<a id="canonical-3112103233200230-0313213130201003-2202123221301131-3121223010212031-3201233333013003-2123223212023203-0111113312033301-1012201103313230"></a>

<a id="canonical-1101002032020002-0020312133201323-3213010120302233-2211113221201332-3003130312332003-2310131301001101-2303200321130111-1122321010010312"></a>

## space_reserve property — volume_defaults / 130110113011 / 12

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-1103301002033212-2022222310112100-2131310122213301-1032320212232132-0220220222311223-0132112112313211-2132021310211011-1203222302313101"></a>

<a id="canonical-2022310121220110-1212022302210020-3201020013310312-3212230321001023-3111123230311211-0021022130223023-0231013011211103-0001203102211203"></a>

## split_on_clone property — volume_defaults / 130110113011 / 13

Type: `"bool"`. Optional.

Split a clone from its parent upon creation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1320221323223103-1110012330110123-1023130132003333-2223102212222231-2100010222003201-1332213203222210-1202023032120332-1203033030030230"></a>

<a id="canonical-3012222201301112-1233001021202110-0303233313333012-3133213323022321-3200230202123330-2110200003133211-3201110311021321-1301232010320332"></a>

## tiering_policy property — volume_defaults / 130110113011 / 14

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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

<a id="canonical-1012021000321233-1321020100222120-2001333133203313-3231332221020110-1300313202021020-0101210330030221-2323312100112121-2111332012010101"></a>

<a id="canonical-0222221300122110-2120232000111321-0133230131331333-3113200001332313-2010200312313033-1222001003333332-3003110320200300-3222201303212101"></a>

## unix_permissions property — volume_defaults / 130110113011 / 15

Type: `"number"`. Optional.

Unix permission mode for new volumes. All allowed 777.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
