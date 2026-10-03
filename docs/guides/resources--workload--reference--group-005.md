---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2001023102132332-0002021331201030-3013113213230032-1120211100311210-2302302311303112-2310112013122020-2022110123102130-0120220102333312"></a>

## sub_path property — mount / 202000032332 / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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

<a id="canonical-0233203120000012-0002112011202212-3323022223133011-2220003213110132-2012122110120013-2120030133123020-1132120321012232-2220230032030233"></a>

## Next pages — mount / 202000032332 / 7

- [job.volumes.empty_dir](resources--workload--reference--group-004.md#canonical-2103002013003112-0200112230012330-1120110101321210-2130121011113010-1002233201310100-2121020032313303-2220023232010202-1102001023221221)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1223231211123302-0302203312202000-2111232130221103-1230023203102122-1133130211023032-2231232202003100-0213320303000202-3033231213122010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003011111131030-0203233201110122-2010322111122022-3231312011300322-1303321321323103-2032211200113303-1332031300211101-2011033020113133"></a>

## job.volumes.host_path — host_path / 103031212123 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.volumes](resources--workload--reference--group-004.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- job.volumes.host_path

<a id="canonical-2310111133311202-2212321203311312-3031110223112231-1010131011222123-1131122331022323-2021333011110123-2030031230203023-1320021331032131"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a host mapped path into the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
host_path {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232331301222130-1022210323101222-1003023233212021-3131113110311032-1121123300022331-1111203323200032-2132102330202322-3231111321331100"></a>

## Direct properties — host_path / 103031212123 / 3

- [mount](resources--workload--reference--group-005.md#canonical-0120233102003331-3110220213311222-0201010013211013-0313133003021232-2201212223230113-0022201212002011-0130031223322301-2100303001311123): complete subsection reference.

<a id="canonical-3210322300203001-3332300110020211-3202302001301331-2020110321033103-1131203330110223-2232203233012010-0113101000232310-2101233203313032"></a>

<a id="canonical-2223222033133312-0301313203031220-2213232132011311-2021300001011301-1130132002020321-0201223302012312-3001321330121012-3330022121310202"></a>

## path property — host_path / 103031212123 / 4

Type: `"string"`. Optional.

Path. Path of the directory on the host.

Upstream description:

Path of the directory on the host.

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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "[^\\\\0]+"
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
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  }
}
```

<a id="canonical-0302111000113200-3111033210331032-3321222221003000-3003311322232311-3333332332032202-3300203301213332-3031313132233311-3031303003123131"></a>

## Next pages — host_path / 103031212123 / 5

- [job.volumes.host_path.mount](resources--workload--reference--group-005.md#canonical-0120233102003331-3110220213311222-0201010013211013-0313133003021232-2201212223230113-0022201212002011-0130031223322301-2100303001311123)
- [job.volumes](resources--workload--reference--group-004.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0120233102003331-3110220213311222-0201010013211013-0313133003021232-2201212223230113-0022201212002011-0130031223322301-2100303001311123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132100222023233-2203132311320210-0011333311021030-1010223103320003-2301223332333133-3313122001030320-1213332000233103-2022033021310000"></a>

## job.volumes.host_path.mount — mount / 111023303332 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.volumes](resources--workload--reference--group-004.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- [job.volumes.host_path](resources--workload--reference--group-005.md#canonical-1223231211123302-0302203312202000-2111232130221103-1230023203102122-1133130211023032-2231232202003100-0213320303000202-3033231213122010)
- job.volumes.host_path.mount

<a id="canonical-2102110233303021-0321200123311013-0212332300123002-1133113130302101-0010102211232311-2102210033201213-0201230130233111-3130333021100132"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211301321002112-2313203322213312-3132113300322201-1321131021222201-0000223031020233-1131003133000112-3011000322012000-0211013010232212"></a>

## Direct properties — mount / 111023303332 / 3

<a id="canonical-3111210131333202-3001001333102111-1131002302030123-0112301030022030-3103233020133302-3310102010233331-1323002100302332-1031233212220300"></a>

<a id="canonical-3100203023032112-2201112130313122-0103332313202102-3303310223313020-0212221311122001-1102100123212211-1300201232033322-3133212020210220"></a>

## mode property — mount / 111023303332 / 4

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1303121021321231-3320011231302200-2303330211330032-2112122111100111-1122301331332032-1020121301220320-2323332321331110-3112331111230032"></a>

<a id="canonical-2301312003103032-1130323320032320-1000201102033233-1230133302100030-3233020221001031-0123013003211022-1002130033233032-0311321002213030"></a>

## mount_path property — mount / 111023303332 / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
    },
    "pattern": "^[^:]*$"
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
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-1133332103022332-0310032012330101-1200331223110321-1203201232202023-1023302332111202-0331232120033332-0131221213002321-3111011120231200"></a>

<a id="canonical-3130031312021000-2231303031110200-0122010302222210-3222012203121001-0301023010302003-1301202232020211-3212331311003123-3113111212113311"></a>

## sub_path property — mount / 111023303332 / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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

<a id="canonical-2000023123122212-0132313223310113-3103223331230012-2320113312303231-1111112100030030-3230013300312120-3132333331011103-2332122133202300"></a>

## Next pages — mount / 111023303332 / 7

- [job.volumes.host_path](resources--workload--reference--group-005.md#canonical-1223231211123302-0302203312202000-2111232130221103-1230023203102122-1133130211023032-2231232202003100-0213320303000202-3033231213122010)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3203200012213001-3111300330011222-2132220323122220-1123013113230202-1310102202333231-2032301213111230-0312131213202310-3223220232002311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330023211002221-1030313112033122-3320301301230002-3211320320333220-1000131222131002-3121212013321230-1133203131003012-0102330303331222"></a>

## job.volumes.persistent_volume — persistent_volume / 312013222020 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.volumes](resources--workload--reference--group-004.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- job.volumes.persistent_volume

<a id="canonical-3111011201203203-3300310131112033-0232012332322131-0311203232211302-1202021322001233-2302130021231001-0021320232200200-2310113120300103"></a>

Type: `"object"`. single nested block, Optional.

Volume containing the Persistent Storage for the workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
persistent_volume {
  # Configure direct properties listed below.
}
```

<a id="canonical-1120331312301300-3301101110300233-0220232233033230-1032302033220213-0121123303012111-2110312303221311-1231201111110321-1120221011311101"></a>

## Direct properties — persistent_volume / 312013222020 / 3

- [mount](resources--workload--reference--group-005.md#canonical-2023331022012002-2120231221033010-3011232212230232-3001013023321101-2102102031311231-0231022101220013-3212313133022021-0231112012320031): complete subsection reference.

- [storage](resources--workload--reference--group-005.md#canonical-0031100121220303-2023112232312211-2312022203121230-2022312101303030-2310210212123232-0230131333201103-3220011000221100-1223333313213203): complete subsection reference.

<a id="canonical-0121033333332013-2302323003121230-1300322201003233-3231012331300021-1120320321323321-2130011211211200-0013221211233121-3330310301112020"></a>

## Next pages — persistent_volume / 312013222020 / 4

- [job.volumes.persistent_volume.mount](resources--workload--reference--group-005.md#canonical-2023331022012002-2120231221033010-3011232212230232-3001013023321101-2102102031311231-0231022101220013-3212313133022021-0231112012320031)
- [job.volumes.persistent_volume.storage](resources--workload--reference--group-005.md#canonical-0031100121220303-2023112232312211-2312022203121230-2022312101303030-2310210212123232-0230131333201103-3220011000221100-1223333313213203)
- [job.volumes](resources--workload--reference--group-004.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2023331022012002-2120231221033010-3011232212230232-3001013023321101-2102102031311231-0231022101220013-3212313133022021-0231112012320031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023312233213203-3232202133312212-0021023012211001-1231212133310113-2210033323212102-2221301001001331-3322032031333330-1221031010223000"></a>

## job.volumes.persistent_volume.mount — mount / 100001322030 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.volumes](resources--workload--reference--group-004.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- [job.volumes.persistent_volume](resources--workload--reference--group-005.md#canonical-3203200012213001-3111300330011222-2132220323122220-1123013113230202-1310102202333231-2032301213111230-0312131213202310-3223220232002311)
- job.volumes.persistent_volume.mount

<a id="canonical-3212121222330030-2330010331030301-2203101221223201-1013030101111321-2202022232230313-3132010231012332-2002110032302013-1122101123122002"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213132323200213-3203201002130102-0230023103303223-3320020121230320-3032201113011232-0222333131213303-0102210231012301-0132203113323123"></a>

## Direct properties — mount / 100001322030 / 3

<a id="canonical-0203321331200100-3211121023213301-3232222133210202-3212203322130333-1123232323101020-0012132300312101-1223023213002300-1203122211200311"></a>

<a id="canonical-1210232111013232-3023032201220120-2312321113011323-1323003102222213-3130102323112021-0123301311022310-3312101020030100-1301213022001131"></a>

## mode property — mount / 100001322030 / 4

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1111220221131312-0120220311032313-1130132123320231-3001303320332211-0021201223220123-1123332202123221-2001111102033312-2323230131020221"></a>

<a id="canonical-0331302320302210-2333102001112000-0312321033132331-3111000011130103-3120221101330132-2100201103333022-3133222320000000-2033110331233030"></a>

## mount_path property — mount / 100001322030 / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
    },
    "pattern": "^[^:]*$"
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
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-0113300023002022-3031300333321011-3222133011113111-3003230313130323-1321200220330101-1132103333323123-2100131312323201-0002100100203101"></a>

<a id="canonical-3032312001331011-2012112032011320-0332002121320102-1232221330312110-2312122111220302-3121000223322110-0102300330332311-3332113120031033"></a>

## sub_path property — mount / 100001322030 / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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

<a id="canonical-0123000003221111-3301123013012212-2300031203311102-3020031210102002-3313212032323220-0313112223202223-3031300131021021-3220101132103231"></a>

## Next pages — mount / 100001322030 / 7

- [job.volumes.persistent_volume](resources--workload--reference--group-005.md#canonical-3203200012213001-3111300330011222-2132220323122220-1123013113230202-1310102202333231-2032301213111230-0312131213202310-3223220232002311)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0031100121220303-2023112232312211-2312022203121230-2022312101303030-2310210212123232-0230131333201103-3220011000221100-1223333313213203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022111113301002-2100002232131223-0203000123232002-0331333211201203-3000121113331010-2222320131121232-0030230101323003-2302332203222323"></a>

## job.volumes.persistent_volume.storage — storage / 333232121021 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.volumes](resources--workload--reference--group-004.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- [job.volumes.persistent_volume](resources--workload--reference--group-005.md#canonical-3203200012213001-3111300330011222-2132220323122220-1123013113230202-1310102202333231-2032301213111230-0312131213202310-3223220232002311)
- job.volumes.persistent_volume.storage

<a id="canonical-0233303313001032-2212132132321200-3313233231012323-3232002222110131-2000131033121112-1001002232133220-0302301111310012-1001023312122000"></a>

Type: `"object"`. single nested block, Optional.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Upstream description:

Persistent storage configuration is used to configure Persistent Volume Claim (PVC)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_size"),
  validators.ConflictingObjectAttributes("class_name",
    "default")}
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
  "x-ves-oneof-field-class_name_choice": "[\"class_name\",\"default\"]"
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212322110232332-3002121212223103-0111030001130030-1210320133312301-2223121320233133-3230111010321312-0322203012202022-0033210021303132"></a>

## Direct properties — storage / 333232121021 / 3

<a id="canonical-0310011032001330-3312010133030210-2301222333301122-0313132313210233-1211000223203231-1231011110312010-1000333030033032-1023221313212231"></a>

<a id="canonical-0300233231302230-2212011131311002-1323100221012200-0213013032003221-2131203333001203-1103112133030122-2222200121223233-0300120213110132"></a>

## access_mode property — storage / 333232121021 / 4

Type: `"string"`. Optional.

\[Enum:
ACCESS\_MODE\_READ\_WRITE\_ONCE|ACCESS\_MODE\_READ\_WRITE\_MANY|ACCESS\_MODE\_READ\_ONLY\_MANY\]
Persistence storage access mode is used to configure access mode for persistent storage -
ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once Read Write Once is used to mount persistent storage
in read/write mode to exactly 1 host - ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many Read Write
Many is used.. Possible values are \`ACCESS\_MODE\_READ\_WRITE\_ONCE\`,
\`ACCESS\_MODE\_READ\_WRITE\_MANY\`, \`ACCESS\_MODE\_READ\_ONLY\_MANY\`. Defaults to
\`ACCESS\_MODE\_READ\_WRITE\_ONCE\`.

Upstream description:

Persistence storage access mode is used to configure access mode for persistent storage

&#8203;- ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once

Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host &#8203;-
ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many

Read Write Many is used to mount persistent storage in read/write mode to many hosts &#8203;-
ACCESS\_MODE\_READ\_ONLY\_MANY: Read Only Many

Read Only Many is used to mount persistent storage in read-only mode to many hosts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ACCESS_MODE_READ_WRITE_ONCE",
  "enum": [
    "ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1010222220321123-2200101223203123-0223200210321020-2332303210122130-0020203323320011-0300331301201031-2102020121223323-2212032101011211"></a>

<a id="canonical-2103000223232013-2030321212310000-3213232100223123-3002010110203201-1213222100113101-3333310323031022-3123313332110013-2302111312020130"></a>

## class_name property — storage / 333232121021 / 5

Type: `"string"`. Optional.

Exclusive with \[default\] Use the specified class name.

Upstream description:

Exclusive with \[default\] Use the specified class name.

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

- [default](resources--workload--reference--group-005.md#canonical-3212012101330020-1220300220102232-2230312110120333-2202331201031221-1033110101001300-1223031221320100-3203213122300311-3320003322322112): complete subsection reference.

<a id="canonical-2311003321313230-2322101013131003-0321211121010211-1212132100220103-2202030012312002-3313313002013123-1232321331212030-0131010011233313"></a>

<a id="canonical-3221331201232103-3232310001023321-0210330310210023-3313232332121220-2221102012120211-1203103120333312-2210012002231232-2123221232123213"></a>

## storage_size property — storage / 333232121021 / 6

Type: `"number"`. Optional.

Size (in GiB). Size in GiB of the persistent storage.

Upstream description:

Size in GiB of the persistent storage.

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
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0330122200121121-0232233133033012-3131103213031113-1023211222102202-3203301010303133-1103031302310311-2303110020130303-2201002113000231"></a>

## Next pages — storage / 333232121021 / 7

- [job.volumes.persistent_volume.storage.default](resources--workload--reference--group-005.md#canonical-3212012101330020-1220300220102232-2230312110120333-2202331201031221-1033110101001300-1223031221320100-3203213122300311-3320003322322112)
- [job.volumes.persistent_volume](resources--workload--reference--group-005.md#canonical-3203200012213001-3111300330011222-2132220323122220-1123013113230202-1310102202333231-2032301213111230-0312131213202310-3223220232002311)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3212012101330020-1220300220102232-2230312110120333-2202331201031221-1033110101001300-1223031221320100-3203213122300311-3320003322322112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121230310122333-2022121212112320-0233323221021000-1121001211300031-2320011120220332-1011001123122231-3213220102223231-3020003003103033"></a>

## job.volumes.persistent_volume.storage.default — default / 100113010321 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.volumes](resources--workload--reference--group-004.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- [job.volumes.persistent_volume](resources--workload--reference--group-005.md#canonical-3203200012213001-3111300330011222-2132220323122220-1123013113230202-1310102202333231-2032301213111230-0312131213202310-3223220232002311)
- [job.volumes.persistent_volume.storage](resources--workload--reference--group-005.md#canonical-0031100121220303-2023112232312211-2312022203121230-2022312101303030-2310210212123232-0230131333201103-3220011000221100-1223333313213203)
- job.volumes.persistent_volume.storage.default

<a id="canonical-2303131013301203-3201113313311111-0321303231023110-3323102013201003-2012203013311222-0010002203202223-3033231203111103-0030030112131202"></a>

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
default = {}
```

<a id="canonical-2223003203023320-1232333303021132-2103230313102100-2021203122003021-3230213302333123-3232201203221103-3030000203311110-1013201211112320"></a>

## Direct properties — default / 100113010321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010123032322300-0131310021113322-3033110302011222-0031033201102300-2000110222301210-2210000013212023-3302211120212110-0333020031310120"></a>

## Next pages — default / 100113010321 / 4

- [job.volumes.persistent_volume.storage](resources--workload--reference--group-005.md#canonical-0031100121220303-2023112232312211-2312022203121230-2022312101303030-2310210212123232-0230131333201103-3220011000221100-1223333313213203)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330132120323330-3112121100333103-1110013030322003-2032131330100331-1000102112111231-3313012221121221-1101032021013233-3122320011022321"></a>

## service — service / 301321010210 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- service

<a id="canonical-1332310211103303-1212000223320202-2221213223003333-1121213113131100-0301313231102131-2310301013222100-0012222131003032-0011220233300032"></a>

Type: `"object"`. single nested block, Optional.

Service does not maintain per replica state, however it can be configured to use persistent storage
that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable
network identity or storage. Common examples of services are web servers, application servers..

Upstream description:

Service does not maintain per replica state, however it can be configured to use persistent storage
that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable
network identity or storage. Common examples of services are web servers, application servers,
traditional SQL databases, etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("containers"),
  validators.ConflictingObjectAttributes("num_replicas",
    "scale_to_zero")}
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
  "x-ves-oneof-field-scaling_choice": "[\"num_replicas\",\"scale_to_zero\"]"
}
```

Terraform syntax:

```terraform
service {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102122230223311-3232333012320330-2013133232301132-2221112313101222-2312320213301230-0112300120021212-1103221121120131-3301230101233202"></a>

## Direct properties — service / 301321010210 / 3

- [advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130): complete subsection reference.

- [configuration](resources--workload--reference--group-015.md#canonical-0120300131213121-2030323030023013-0301331320133010-0123120231213121-1220222031210011-1001323001213020-0213022223133211-2110030001312333): complete subsection reference.

- [containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122): complete subsection reference.

- [deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003): complete subsection reference.

<a id="canonical-0333212012011110-2130112003031213-0311010210012302-1221223333233322-0201111111323232-0030112320000202-1212020101231213-2123201101310111"></a>

<a id="canonical-3302120123201323-3211313311302031-3312111211210331-3332123031302222-2122301201231321-0103333332010231-3120112022311030-0102023311132033"></a>

## num_replicas property — service / 301321010210 / 4

Type: `"number"`. Optional.

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Upstream description:

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  }
}
```

- [scale_to_zero](resources--workload--reference--group-016.md#canonical-2112203323032023-1022033211220013-2331111010311212-0200221020302323-2123121101110311-2211200012322311-2122132200232301-3022222221010211): complete subsection reference.

- [volumes](resources--workload--reference--group-016.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003): complete subsection reference.

<a id="canonical-2311233120010333-1132233213311221-1302031221323233-0023212001301120-3302121031102202-1332012212011122-1010321232023133-3222100020202002"></a>

## Next pages — service / 301321010210 / 5

- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.configuration](resources--workload--reference--group-015.md#canonical-0120300131213121-2030323030023013-0301331320133010-0123120231213121-1220222031210011-1001323001213020-0213022223133211-2110030001312333)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [service.scale_to_zero](resources--workload--reference--group-016.md#canonical-2112203323032023-1022033211220013-2331111010311212-0200221020302323-2123121101110311-2211200012322311-2122132200232301-3022222221010211)
- [service.volumes](resources--workload--reference--group-016.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100013133221331-2301110112300101-2323131002220001-3330201102301312-2323313032133333-0032003103332332-3321220323033133-3311310100311021"></a>

## service.advertise_options — advertise_options / 132233331103 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- service.advertise_options

<a id="canonical-1330113111333321-0202110223311103-1033021331320313-1330313231121010-1101130003002030-0202122233113330-2031032023320312-1202202212321031"></a>

Type: `"object"`. single nested block, Optional.

Advertise OPTIONS are used to configure how and where to advertise the workload using load
balancers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_in_cluster"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_in_cluster",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_in_cluster",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "do_not_advertise")}
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
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_in_cluster\",\"advertise_on_public\",\"do_not_advertise\"]"
}
```

Terraform syntax:

```terraform
advertise_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000312000003201-2003210132320323-0013332111233232-3101002331023221-0013111100312233-2331101223311022-2221102322030033-2230121103331213"></a>

## Direct properties — advertise_options / 132233331103 / 3

- [advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020): complete subsection reference.

- [advertise_in_cluster](resources--workload--reference--group-008.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213): complete subsection reference.

- [advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101): complete subsection reference.

- [do_not_advertise](resources--workload--reference--group-015.md#canonical-0310302023323212-1013021312131132-1222231111312001-2121122023132203-1332202130321213-0133012220300033-3111032100113201-3031211330022333): complete subsection reference.

<a id="canonical-1201133231131332-2200032210133123-1233231111102032-0201120330200003-1311132333000210-2230232331101310-3313310002003222-2230010223222110"></a>

## Next pages — advertise_options / 132233331103 / 4

- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.do_not_advertise](resources--workload--reference--group-015.md#canonical-0310302023323212-1013021312131132-1222231111312001-2121122023132203-1332202130321213-0133012220300033-3111032100113201-3031211330022333)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132212123130021-1102321130220100-1231230002133302-3333201112300303-2110121332222111-0023032220210231-0121300103033130-2231203120223312"></a>

## service.advertise_options.advertise_custom — advertise_custom / 022220230123 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- service.advertise_options.advertise_custom

<a id="canonical-2211032332323103-1333121332123121-3032103321322003-3331122020211113-1201023212110002-0323233232022301-2102101111213233-3113013033132003"></a>

Type: `"object"`. single nested block, Optional.

Advertise this workload via loadbalancer on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where",
    "ports")}
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
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212331213011310-3111103103223333-2032030001333223-3131320010133002-1310300220221321-3113333300020113-3202011330012021-2003002211310003"></a>

## Direct properties — advertise_custom / 022220230123 / 3

- [advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033): complete subsection reference.

- [ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200): complete subsection reference.

<a id="canonical-0201012222110021-3223010012100013-2311313002311212-0311230221332003-2102002310130302-2001032123310031-2212302230122311-1223132120302110"></a>

## Next pages — advertise_custom / 022220230123 / 4

- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121131022221230-2132033031121131-3221202123102010-0120133302330131-1030231122132033-2232310031023232-3133111003221220-1132123032311301"></a>

## service.advertise_options.advertise_custom.advertise_where — advertise_where / 120003032113 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- service.advertise_options.advertise_custom.advertise_where

<a id="canonical-0002001030223131-0201101132010012-0320131300310320-2223311102012032-3111223311000333-1111223012201302-0211112201230303-2010200302312220"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service")}
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221002132030113-2232000101203000-0121320130101220-0202110111013201-1333332001121300-0301200201023210-1210132110103020-3302232232330310"></a>

## Direct properties — advertise_where / 120003032113 / 3

- [site](resources--workload--reference--group-005.md#canonical-0030333100303013-1333122033000310-2322131031010012-3320311021122031-0010002121033331-2002122020332203-1232011312312311-0000322030003001): complete subsection reference.

- [virtual_site](resources--workload--reference--group-005.md#canonical-0233210232101101-1111132120230132-3012030232123130-2031231011320101-3311201113011302-2113330313323032-0302331031323203-3101303230022131): complete subsection reference.

- [vk8s_service](resources--workload--reference--group-005.md#canonical-3203023030311011-1123233030130010-2121223300322103-0331031333102030-1003323232232011-2122300232331102-3312332311122002-1200323203122001): complete subsection reference.

<a id="canonical-1020121203231233-3310231100113223-3222302032033212-2120221003103011-1002331331011312-1102031333110232-3221033111313323-0000032010231301"></a>

## Next pages — advertise_where / 120003032113 / 4

- [service.advertise_options.advertise_custom.advertise_where.site](resources--workload--reference--group-005.md#canonical-0030333100303013-1333122033000310-2322131031010012-3320311021122031-0010002121033331-2002122020332203-1232011312312311-0000322030003001)
- [service.advertise_options.advertise_custom.advertise_where.virtual_site](resources--workload--reference--group-005.md#canonical-0233210232101101-1111132120230132-3012030232123130-2031231011320101-3311201113011302-2113330313323032-0302331031323203-3101303230022131)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-005.md#canonical-3203023030311011-1123233030130010-2121223300322103-0331031333102030-1003323232232011-2122300232331102-3312332311122002-1200323203122001)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0030333100303013-1333122033000310-2322131031010012-3320311021122031-0010002121033331-2002122020332203-1232011312312311-0000322030003001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102022010222230-2012033031111020-0012021002333210-1112331232231022-0132000321310313-2023131032330102-0231111103113201-0131122313331332"></a>

## service.advertise_options.advertise_custom.advertise_where.site — site / 323122303231 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033)
- service.advertise_options.advertise_custom.advertise_where.site

<a id="canonical-0001312130223102-0011232213223010-1103222321330121-1331102022003330-2112310003313313-3333212302310121-1022021211311223-3012002311030212"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a CE site along with network type and an optional IP address where a load
balancer could be advertised.

Upstream description:

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131330022221233-3223230111021121-2022222101022020-2311221220131031-3330311320100021-0311311201332200-3110003213322303-2203303010123223"></a>

## Direct properties — site / 323122303231 / 3

<a id="canonical-1322022022303123-2223023033133120-0000011031321213-0332211112012000-0213010322021133-2132223201102132-3222132111113200-1202210120120010"></a>

<a id="canonical-3102133200221112-2300330010333312-2303201210330200-3221103021223013-1320000030210322-2220122233312221-2020312032031203-3010110003331210"></a>

## ip property — site / 323122303231 / 4

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

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

<a id="canonical-0330200022012011-1210312131230032-0203103211113130-2101222330100002-1123102132011211-2321103102123232-0000223000101231-0002323012030232"></a>

<a id="canonical-2023013300120113-1323310323003123-0233133303313230-3113010001132033-1203330001032301-3133233001110010-0022203021112032-3131030323120312"></a>

## network property — site / 323122303231 / 5

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](resources--workload--reference--group-005.md#canonical-3320022311110022-3103030003202322-1131012232020133-2320321110203332-0111013210321122-0322313101132331-2131233111312213-1120001332232122): complete subsection reference.

<a id="canonical-1032300220210221-2312301230013010-0323101203012102-3133302102312332-0301002031331331-2312230101102213-2100233312200113-1033312002330130"></a>

## Next pages — site / 323122303231 / 6

- [service.advertise_options.advertise_custom.advertise_where.site.site](resources--workload--reference--group-005.md#canonical-3320022311110022-3103030003202322-1131012232020133-2320321110203332-0111013210321122-0322313101132331-2131233111312213-1120001332232122)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3320022311110022-3103030003202322-1131012232020133-2320321110203332-0111013210321122-0322313101132331-2131233111312213-1120001332232122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120213110202332-1030032021020233-1120130300001031-0120011322313000-2321231220333321-2021323301023323-2310311003100332-2331231023112110"></a>

## service.advertise_options.advertise_custom.advertise_where.site.site — site / 002031300031 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033)
- [service.advertise_options.advertise_custom.advertise_where.site](resources--workload--reference--group-005.md#canonical-0030333100303013-1333122033000310-2322131031010012-3320311021122031-0010002121033331-2002122020332203-1232011312312311-0000322030003001)
- service.advertise_options.advertise_custom.advertise_where.site.site

<a id="canonical-2132121311133003-2132020123320132-3332203131112003-1131100121303200-0130032222331033-0100222032331213-0223032110201120-3321312123022323"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1213022002010013-2332223020301110-0023010010322011-0020102130322123-3212000031300132-0310100131211213-3011231313321300-0221211012030332"></a>

## Direct properties — site / 002031300031 / 3

<a id="canonical-3200323120100221-1013303103220103-1212213103010333-3211133232330113-3311223230210302-0102102230123132-3320023200001002-3103223120133023"></a>

<a id="canonical-3320102030232001-2000110230301030-3122200210331020-3003113110311023-1020202202310023-2321122000003012-3112233212211033-3231003023001232"></a>

## name property — site / 002031300031 / 4

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

<a id="canonical-0022220303010221-3313003230212012-2322222211012013-2300330211123220-0033303222322331-2131032113030022-0101102333023103-1313312020300213"></a>

<a id="canonical-1030120113232032-0231023012300020-1321303320322000-1211022022321130-2322311003302203-3033022102020322-0320031121320300-1121012323233200"></a>

## namespace property — site / 002031300031 / 5

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

<a id="canonical-3023233233311000-2101130031123223-1032311333203101-2112111033231330-3101303002200213-3103103120330010-3103310310032212-0003222122013030"></a>

<a id="canonical-1331130300220113-0002322323130001-2023211023210223-2311321130321322-1003122100230001-2030103231000121-0001220123102333-0220231110321213"></a>

## tenant property — site / 002031300031 / 6

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

<a id="canonical-0002011013211313-1102022011211310-0001000132211323-3311000011002231-1120132033023020-2213301003010000-3220031200323010-0320123220033313"></a>

## Next pages — site / 002031300031 / 7

- [service.advertise_options.advertise_custom.advertise_where.site](resources--workload--reference--group-005.md#canonical-0030333100303013-1333122033000310-2322131031010012-3320311021122031-0010002121033331-2002122020332203-1232011312312311-0000322030003001)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0233210232101101-1111132120230132-3012030232123130-2031231011320101-3311201113011302-2113330313323032-0302331031323203-3101303230022131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310223220102202-1110133102200010-1312001021210201-0210132012013033-3333012101222333-1113132312100022-3333103001030132-3113120003230231"></a>

## service.advertise_options.advertise_custom.advertise_where.virtual_site — virtual_site / 220101000121 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033)
- service.advertise_options.advertise_custom.advertise_where.virtual_site

<a id="canonical-3333313023011111-1100321202102100-2131232110322201-3201021102210023-2301210230301202-1023033022001233-2322030032002232-0333113131003301"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a customer site virtual site along with network type where a load balancer
could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120122221211111-3002203302203121-0001010100331331-3020321032232001-2200321323221110-1002320011203211-1220112022310110-0010111113032211"></a>

## Direct properties — virtual_site / 220101000121 / 3

<a id="canonical-3030021213122133-2211120333111013-1233323112022120-3301203032310133-1202333003200213-1231000002001330-2021001233210123-0120333013230003"></a>

<a id="canonical-3111310031212311-1303102322003103-3313220210303031-2011313301123212-2331111010023130-1213211220130111-3113130222122022-0311300121122331"></a>

## network property — virtual_site / 220101000121 / 4

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](resources--workload--reference--group-005.md#canonical-3111211030302213-0110213100301232-1020233303203221-3032002013303100-2020003131020020-2123313002320022-0102003213012311-0112003323132321): complete subsection reference.

<a id="canonical-2331032030131303-2111211133232003-1102323232131030-1132133133200323-0010133302201130-2113303130112233-3032320232012133-2121220331203333"></a>

## Next pages — virtual_site / 220101000121 / 5

- [service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site](resources--workload--reference--group-005.md#canonical-3111211030302213-0110213100301232-1020233303203221-3032002013303100-2020003131020020-2123313002320022-0102003213012311-0112003323132321)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3111211030302213-0110213100301232-1020233303203221-3032002013303100-2020003131020020-2123313002320022-0102003213012311-0112003323132321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302232200310332-2310200111301010-1120011233223120-2121221103123231-2112331233113100-0122313320033100-2200221021200332-0231213311003120"></a>

## service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site — virtual_site / 213033013312 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033)
- [service.advertise_options.advertise_custom.advertise_where.virtual_site](resources--workload--reference--group-005.md#canonical-0233210232101101-1111132120230132-3012030232123130-2031231011320101-3311201113011302-2113330313323032-0302331031323203-3101303230022131)
- service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-3201031312332103-1033012022130031-3101001212302000-3200330223100230-3121110120121130-3222112031220232-2130102332200201-3301100310001231"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213332003321303-0232200131100212-0211313220121111-2323130003113301-3102020300010310-3312311331311003-2203312023012101-3022312201103133"></a>

## Direct properties — virtual_site / 213033013312 / 3

<a id="canonical-0112222122333301-0121231330022320-1330131003310110-3310133201123331-1010130110133302-0323033001213330-3133112231011203-3003103011211212"></a>

<a id="canonical-1332332002300010-1210212322223031-2031333012221320-1112000322210022-1022213022212101-2232221313111010-3003123213010002-0123130223130131"></a>

## name property — virtual_site / 213033013312 / 4

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

<a id="canonical-1210313233323130-2310311001203201-0110312000110300-2210130113302230-1002122311210033-1101011231332000-3310011112123331-1130221232221032"></a>

<a id="canonical-2002300003112322-1202001121123332-0202333122211033-1103201131310132-1020230021032112-1210030101100300-2103002211230220-0231112000323033"></a>

## namespace property — virtual_site / 213033013312 / 5

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

<a id="canonical-2102301113113013-1012013232330022-1002321223302201-3100331131112201-2220300032121112-3323320002133112-0222131100231230-0021230332000033"></a>

<a id="canonical-1022033023212211-2312232021220300-1032230222223021-1233032102332120-2103210232020123-1001013033203232-2000301200021211-3130110201232100"></a>

## tenant property — virtual_site / 213033013312 / 6

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

<a id="canonical-1100031002101310-2333330122330232-1210202122223101-1301303120100200-0100232110202113-1302332031201302-2322120011022133-1123313023332313"></a>

## Next pages — virtual_site / 213033013312 / 7

- [service.advertise_options.advertise_custom.advertise_where.virtual_site](resources--workload--reference--group-005.md#canonical-0233210232101101-1111132120230132-3012030232123130-2031231011320101-3311201113011302-2113330313323032-0302331031323203-3101303230022131)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3203023030311011-1123233030130010-2121223300322103-0331031333102030-1003323232232011-2122300232331102-3312332311122002-1200323203122001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322111023120300-0022000102000003-1300000002011332-2010333320310101-2300223023031103-1203322310120110-2113121230320331-2230103021231312"></a>

## service.advertise_options.advertise_custom.advertise_where.vk8s_service — vk8s_service / 311130320023 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service

<a id="canonical-0220001131232322-0111013031110002-3121201111203000-2221111131310323-0332023310013020-1123221210123111-0032100312222212-3301200133312332"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
vk8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-1213220033023131-0333003031220313-0322013211233200-0011331030211330-0332213322012312-3311223022012300-1222200203021312-0213002133201302"></a>

## Direct properties — vk8s_service / 311130320023 / 3

- [site](resources--workload--reference--group-005.md#canonical-2122321022022332-3203222311110010-2223232233223303-3221121030033023-1233330023311333-1022303200311320-2331313213202110-0013023001222212): complete subsection reference.

- [virtual_site](resources--workload--reference--group-005.md#canonical-1133023321102002-1320011023312030-1023301022212101-0111011331103112-1103212020122232-3302002323023120-0033111101001203-0103310231110001): complete subsection reference.

<a id="canonical-1220302121203300-2231121223332220-0230011113100301-3222123210200110-2011212211310120-3223300222200010-2011121122203033-0231110330201030"></a>

## Next pages — vk8s_service / 311130320023 / 4

- [service.advertise_options.advertise_custom.advertise_where.vk8s_service.site](resources--workload--reference--group-005.md#canonical-2122321022022332-3203222311110010-2223232233223303-3221121030033023-1233330023311333-1022303200311320-2331313213202110-0013023001222212)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--workload--reference--group-005.md#canonical-1133023321102002-1320011023312030-1023301022212101-0111011331103112-1103212020122232-3302002323023120-0033111101001203-0103310231110001)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2122321022022332-3203222311110010-2223232233223303-3221121030033023-1233330023311333-1022303200311320-2331313213202110-0013023001222212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033200200223212-1001203132133330-0002011321203102-1233230210130220-2231021312100321-0213321033103022-2030023121202002-1223301130322031"></a>

## service.advertise_options.advertise_custom.advertise_where.vk8s_service.site — site / 231021233200 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-005.md#canonical-3203023030311011-1123233030130010-2121223300322103-0331031333102030-1003323232232011-2122300232331102-3312332311122002-1200323203122001)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-0122000002030203-0023100131313132-3031112100301212-0201001011211211-2312011312011211-1012233300223300-3233223330213201-1232123000320331"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300102032131301-2122233133213223-3230112012311033-0201031202300212-3101321222013113-2133111033131232-3133013300313102-3231101322212302"></a>

## Direct properties — site / 231021233200 / 3

<a id="canonical-1310110330320300-3233221132031131-2222201020022331-2002323000020330-3323312111332000-2331233012031123-0212200001020210-2313010311033103"></a>

<a id="canonical-1202031033230032-2103002300032221-2003020021031333-1031222311312101-3102100113221123-1032221333333231-2302001301212001-3311230011221200"></a>

## name property — site / 231021233200 / 4

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

<a id="canonical-2012223332203330-2220312002233231-2031310020131101-3032120313121322-1221001221130102-3020022323130201-2333320133033121-2222130123213033"></a>

<a id="canonical-0323303003321313-3103012312211130-2122003122003201-1231300322011011-3322010312103131-0313132132021033-2102020130233300-1031312222111232"></a>

## namespace property — site / 231021233200 / 5

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

<a id="canonical-1301120211033323-0301031013113311-0023100300233101-1132220100230312-0211001112313032-1330230123220011-0132031210112303-2321303020221133"></a>

<a id="canonical-3021320013030101-3211231013021331-3233222020302232-0221120022312131-1102333221101202-3331002032330231-2103221003321120-2321232212010313"></a>

## tenant property — site / 231021233200 / 6

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

<a id="canonical-0302223003330311-3323130023333230-2001333233231231-0113112202033230-3013332222132033-3013210011211100-2012103030130003-1001111110013223"></a>

## Next pages — site / 231021233200 / 7

- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-005.md#canonical-3203023030311011-1123233030130010-2121223300322103-0331031333102030-1003323232232011-2122300232331102-3312332311122002-1200323203122001)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1133023321102002-1320011023312030-1023301022212101-0111011331103112-1103212020122232-3302002323023120-0033111101001203-0103310231110001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000111213030220-3023031220120211-2000331330111113-3213000031232010-0100101100312133-1000331312323000-0011021310132010-3220123033322003"></a>

## service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site — virtual_site / 321321323222 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-005.md#canonical-3203023030311011-1123233030130010-2121223300322103-0331031333102030-1003323232232011-2122300232331102-3312332311122002-1200323203122001)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-3012020110101220-0233001232110300-2221223201333320-1002012220320130-2330110203313220-2332110221221013-1103313201231031-0332120022121012"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103112101221110-1002301131321323-0223010113203022-2011321010120021-1033023100323100-1310302201012023-0111320030312011-0103220312031223"></a>

## Direct properties — virtual_site / 321321323222 / 3

<a id="canonical-2213333212102010-1301310132010100-3230331130131201-1123222131112332-1323312002030023-2331312003311101-0203323012202011-1202220332100030"></a>

<a id="canonical-0232213301013210-0231310012110321-3003200211330300-0003100023223032-3323202233123212-1103132103111203-3131323010133312-0303232210303101"></a>

## name property — virtual_site / 321321323222 / 4

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

<a id="canonical-3120112102303031-3332103112031301-0333222110221101-3231000332311200-2330130122123221-0313022010011110-3033300303333000-1022313030023200"></a>

<a id="canonical-3032311011303122-2013022332211133-2031322023320133-3110231300230310-1301300131010200-0231232233032313-1023003201132001-0123130211311130"></a>

## namespace property — virtual_site / 321321323222 / 5

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

<a id="canonical-2113210110012013-0213103321322022-1110322231022212-3123333120120213-2023201130200101-2321202223323332-0312231030302020-3131312001110330"></a>

<a id="canonical-1010121320300033-3302133022222333-3231100102213133-0030121322300221-0102030222203332-3202301133003302-2121102300311200-3201011302321002"></a>

## tenant property — virtual_site / 321321323222 / 6

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

<a id="canonical-0210301130133033-2123101012023321-1020123111131320-3311231321020332-3332312331112331-0230023233132221-3330233220013003-3322030030213320"></a>

## Next pages — virtual_site / 321321323222 / 7

- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-005.md#canonical-3203023030311011-1123233030130010-2121223300322103-0331031333102030-1003323232232011-2122300232331102-3312332311122002-1200323203122001)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133211333003222-2301203002022310-2313023220120020-0323010232300323-0132200330231132-3000132031000312-1310002313011013-1101131021210221"></a>

## service.advertise_options.advertise_custom.ports — ports / 203121112123 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- service.advertise_options.advertise_custom.ports

<a id="canonical-2123011100123200-2023222200303001-2110333131132110-3333332120020000-3130213022212011-3113113201333111-3313111233113103-1120113010221233"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("http_loadbalancer",
    "tcp_loadbalancer")}
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
ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223322212333203-1022210321111333-2112000002031110-3313303030033100-1323212112230201-1221202213121332-2333221000101023-3320100301233032"></a>

## Direct properties — ports / 203121112123 / 3

- [http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102): complete subsection reference.

- [port](resources--workload--reference--group-008.md#canonical-2011231233103202-1130133301231221-3112112222130011-0133100320022222-3002023210301231-3323320331313221-3230300303311111-0302223303200123): complete subsection reference.

- [tcp_loadbalancer](resources--workload--reference--group-008.md#canonical-2333000323211321-3320220213111111-3000113231121121-2211220223011110-3303030010002012-3322030301222113-2111223031202010-0320112332022000): complete subsection reference.

<a id="canonical-2101222120203100-2230013200003201-0203033103323012-3033222232211223-3102021203000112-0320230330300101-2100323123010333-1130301330202221"></a>

## Next pages — ports / 203121112123 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-008.md#canonical-2011231233103202-1130133301231221-3112112222130011-0133100320022222-3002023210301231-3323320331313221-3230300303311111-0302223303200123)
- [service.advertise_options.advertise_custom.ports.tcp_loadbalancer](resources--workload--reference--group-008.md#canonical-2333000323211321-3320220213111111-3000113231121121-2211220223011110-3303030010002012-3322030301222113-2111223031202010-0320112332022000)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321323233022110-2100133101001003-1233211300211322-2120020033332330-1021111211212010-1000123021313031-3123100331203130-3123112132132323"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer — http_loadbalancer / 011000133202 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- service.advertise_options.advertise_custom.ports.http_loadbalancer

<a id="canonical-0303230033011100-3031131302312112-3310132312002001-3311232123100021-1011120121133303-0230130321022020-0022200313033211-3123212113011032"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http loadbalancer.

Upstream description:

HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("default_route",
    "specific_routes"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "https_auto_cert"),
  validators.ConflictingObjectAttributes("https",
    "https_auto_cert")}
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
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

Terraform syntax:

```terraform
http_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002101123211201-2310001210231211-0231320120110011-3102033011233020-3123033220311100-3333030202000020-3313020103012222-2031130110333311"></a>

## Direct properties — http_loadbalancer / 011000133202 / 3

- [default_route](resources--workload--reference--group-005.md#canonical-1120332333313120-0322101311023310-1112002320230010-1102031301101300-3001012203013301-2221123112031100-1003333321223200-1331102232222102): complete subsection reference.

<a id="canonical-1213303020331131-3021223120323130-2010323012301130-2130100330023332-1223120221330221-0100111122332222-1302000020111112-2302113212011110"></a>

<a id="canonical-2221111011231330-3330310102213330-3302012030323211-3211002213301101-2231323332131323-2310112331211220-3100110310332312-2132030201101310"></a>

## domains property — http_loadbalancer / 011000133202 / 4

Type: `["list", "string"]`. Optional.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\`
and \`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
Wildcards must match a whole DNS label. E.g. \`\`\*.example.com\`\` and \*.bar.example.com are
valid, however \`\`\*bar.example.com\`\` or \`\`\*-bar.example.com\`\` is invalid

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

Provider validators and defaults (from schema source):

```go
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

- [http](resources--workload--reference--group-005.md#canonical-2130033223303203-2133111123312320-1112100320211003-2023132021300031-0033130333001121-1021321231212232-2003013202213311-2020031200323310): complete subsection reference.

- [https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312): complete subsection reference.

- [specific_routes](resources--workload--reference--group-007.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232): complete subsection reference.

<a id="canonical-2332023322332001-3231332112223222-0112320301230122-2221201003022102-3312322232032030-3302020103112200-3330323300200121-2110032310320302"></a>

## Next pages — http_loadbalancer / 011000133202 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-005.md#canonical-1120332333313120-0322101311023310-1112002320230010-1102031301101300-3001012203013301-2221123112031100-1003333321223200-1331102232222102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.http](resources--workload--reference--group-005.md#canonical-2130033223303203-2133111123312320-1112100320211003-2023132021300031-0033130333001121-1021321231212232-2003013202213311-2020031200323310)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1120332333313120-0322101311023310-1112002320230010-1102031301101300-3001012203013301-2221123112031100-1003333321223200-1331102232222102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233222133133210-2212300211020202-1302122013130112-3302310013211111-2103300322131012-0122003302300300-2213021223332303-3002202322011213"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route — default_route / 310120012221 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route

<a id="canonical-2021023200320110-3113330120122212-2023320013302021-1120131111010133-3320220300113201-0313020003000233-2102232231111331-0011022111122002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

Default route matching all APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
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
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210121301101313-3333303003030102-0201020122332112-3111203222322123-2322010101300121-1233013313230113-2002210001012032-3220210301022030"></a>

## Direct properties — default_route / 310120012221 / 3

- [auto_host_rewrite](resources--workload--reference--group-005.md#canonical-2333210103230010-1112232312000120-2222331322132323-0111001330321031-2001303021310101-2131213003012332-1301130223133223-3010132211313221): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-005.md#canonical-1101321300120022-3232010310223023-3213223030111101-3311002222110012-1121000111102131-0303213020333130-1301200311032201-2300101103020332): complete subsection reference.

<a id="canonical-3020301333000102-0301033032013311-1222313231111230-3312032323330313-1222111101211320-0120110221220201-2201102033032320-0003022320223222"></a>

<a id="canonical-3211133231323233-0202313132221033-2230011020332303-2232113202323011-1233001101310113-0123213230310223-3323302113333232-2103022113303312"></a>

## host_rewrite property — default_route / 310120012221 / 4

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

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

<a id="canonical-1032230211132221-2222010332112011-0303221111231113-2002302112131010-1010120103001300-0222112102322120-1300102120132321-2133102310130013"></a>

## Next pages — default_route / 310120012221 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite](resources--workload--reference--group-005.md#canonical-2333210103230010-1112232312000120-2222331322132323-0111001330321031-2001303021310101-2131213003012332-1301130223133223-3010132211313221)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite](resources--workload--reference--group-005.md#canonical-1101321300120022-3232010310223023-3213223030111101-3311002222110012-1121000111102131-0303213020333130-1301200311032201-2300101103020332)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2333210103230010-1112232312000120-2222331322132323-0111001330321031-2001303021310101-2131213003012332-1301130223133223-3010132211313221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322311002323013-0111220232331100-3222100002120123-3122132220322131-2301001003333221-0311010002322031-1113300321302210-2201010013103130"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite — auto_host_rewrite / 133222231230 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-005.md#canonical-1120332333313120-0322101311023310-1112002320230010-1102031301101300-3001012203013301-2221123112031100-1003333321223200-1331102232222102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-1001221310233112-0230313322302332-1010303020013323-3101103112303123-3300000022330321-0312002001113131-1033121303001033-3033321311333212"></a>

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
auto_host_rewrite = {}
```

<a id="canonical-2330122311213300-2210000232200120-1122211022333033-1211330101013131-1301312012121202-3012011010101131-0001021032122013-1330132230120301"></a>

## Direct properties — auto_host_rewrite / 133222231230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122303302212120-0000221131133121-1031000311212302-3121323110221330-1231321303132323-1000102322120031-0221001111223331-2200133313010101"></a>

## Next pages — auto_host_rewrite / 133222231230 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-005.md#canonical-1120332333313120-0322101311023310-1112002320230010-1102031301101300-3001012203013301-2221123112031100-1003333321223200-1331102232222102)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1101321300120022-3232010310223023-3213223030111101-3311002222110012-1121000111102131-0303213020333130-1301200311032201-2300101103020332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020301201022111-2232232320002120-2111310333303303-1300322120101331-2331232002113003-2200001131202213-2022112332121002-2130310112011123"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite — disable_host_rewrite / 111121200013 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-005.md#canonical-1120332333313120-0322101311023310-1112002320230010-1102031301101300-3001012203013301-2221123112031100-1003333321223200-1331102232222102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-1020023122301101-0130100110302010-3123012231122030-2120110100023132-2301333013233231-1110332333201023-1210023223312311-2003003203301332"></a>

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
disable_host_rewrite = {}
```

<a id="canonical-0210111320210020-2000211312132003-0100013231003220-3033130000121010-0033231002203113-3232230000110203-2023330311201120-2322223233201030"></a>

## Direct properties — disable_host_rewrite / 111121200013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133323310301202-1130211131110023-3330033322333211-0211323103021223-3221221213313333-3001212212320103-1331121131210310-0201033313212213"></a>

## Next pages — disable_host_rewrite / 111121200013 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-005.md#canonical-1120332333313120-0322101311023310-1112002320230010-1102031301101300-3001012203013301-2221123112031100-1003333321223200-1331102232222102)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2130033223303203-2133111123312320-1112100320211003-2023132021300031-0033130333001121-1021321231212232-2003013202213311-2020031200323310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211021113112102-1231121322203131-2223232302213231-3233020322133030-1120100003031103-0301223031033200-0221112332221102-1113300031211123"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.http — http / 102010321330 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.http

<a id="canonical-0222313312123032-3312120003201000-2133133011112323-3113222020001313-2021212023020000-2221021211233313-1223131101113331-3011322203332220"></a>

Type: `"object"`. single nested block, Optional.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103023311332222-3331232311310300-3101101332033102-2332213233320122-3022311333002222-0020022232331200-2122120113010210-2100213233022230"></a>

## Direct properties — http / 102010321330 / 3

<a id="canonical-2101100001023133-3033321012233010-3232012033011331-1201332003011302-3130020321130300-2202020021320333-3200120133131210-3010133200011220"></a>

<a id="canonical-0132130302113000-3121003000013113-0231130230110331-2130000232111300-3320223121332120-2111222221302003-1122003130112130-2222011203022033"></a>

## dns_volterra_managed property — http / 102010321330 / 4

Type: `"bool"`. Optional.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

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

<a id="canonical-3122030101212322-2120313112033321-3111012221310130-0011113333100302-3213022222230131-1223002213020113-2033132023221023-1003300200311213"></a>

<a id="canonical-3021310133021212-1021120111122123-0331102203323012-2113102123302330-2220002321123200-2320321002113031-2113331212311112-1032013300230120"></a>

## port property — http / 102010321330 / 5

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

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

<a id="canonical-2311133332002223-2313230333023321-3023120123133023-0233110120023000-3010331023331233-1232032011313102-0321113220112233-2231313203131323"></a>

<a id="canonical-1100331320331331-3303231023303303-1200110031213101-3021112223300101-1112301011210212-0231223113322012-3232101302203212-1300232032103022"></a>

## port_ranges property — http / 102010321330 / 6

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

<a id="canonical-2033132130230200-1310032111303230-2320331001031200-3211133031213000-1321233203100011-1133013022123221-2121110132111123-0123022332120011"></a>

## Next pages — http / 102010321330 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210333303321113-3130031113211203-2211111302000230-3102002220213202-0212323311313222-0021333310022113-3021230333120030-2330330222133032"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https — https / 033300302223 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https

<a id="canonical-0020303301233223-0233020303230321-3122010003303123-3131101022222002-2011220100123301-3330223003012312-3013212023312010-2332320000100131"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_parameters")}
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
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111101030012210-1313112022112203-0210210212322312-0311330013103230-2120313331013300-2202032212200210-2212320001330210-2201202230022211"></a>

## Direct properties — https / 033300302223 / 3

<a id="canonical-0301303112320222-0131210033001103-2201102203122231-1013003232230010-2302333223131113-2123300200001033-2331313003212222-1322223100312031"></a>

<a id="canonical-3011010120131123-1020103221000233-1331201232133100-2123230100201030-0001333202021000-2212200120033113-0113212323223212-1300012013022230"></a>

## add_hsts property — https / 033300302223 / 4

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

<a id="canonical-0301221322230011-3311330031201222-2012130023201201-0321131112303132-1011312230130333-0013003222302002-2222230131313212-0130030123232110"></a>

<a id="canonical-3201010331101201-3113203221221132-2233013312321221-0230020212300223-1301012222131311-2322113132000212-2233333103202011-1201011313003122"></a>

## append_server_name property — https / 033300302223 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](resources--workload--reference--group-005.md#canonical-0033100331020001-2221313100101220-3310100020233022-2211123323300001-0231323312333323-0032322122221223-0222211200302233-3122211001332213): complete subsection reference.

<a id="canonical-3000123331231110-2123113120320033-1313113323221131-0320002123031033-2210300001010002-0032221311133000-1330022231121220-0023320212320122"></a>

<a id="canonical-2110232013102023-2121301322323020-1232310233310302-2021100233121210-2310013313213033-0222023223100112-2123320331022103-2222112020223130"></a>

## connection_idle_timeout property — https / 033300302223 / 6

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](resources--workload--reference--group-005.md#canonical-2222132001202300-2223100111302031-3221210230001201-1220312313130120-0001232202031003-0302013332210031-1130121020112230-2221102100233321): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-005.md#canonical-3013320033200112-2330211103132210-1111021302202201-1003233032230222-0300000132132211-0201133003111333-1012021102221030-2022022103103133): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-005.md#canonical-2130013022101113-3232010232323220-0210310301112110-1233302313322122-1310311322332312-2201211202201021-3333320313123113-2100323212032033): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-005.md#canonical-3110011313233100-3302030313300031-1211110123230312-1103211030232121-1230111303122131-0000021202212203-0301121032130213-1003013003110302): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-005.md#canonical-1302010300230312-3111312201201113-2013133300112103-0112233323311000-0220233130302132-1102033102333132-1100222323112330-1222302133120000): complete subsection reference.

<a id="canonical-2321031113020221-3220010231133313-0102232012322230-2222202123122013-2323330023012303-2022013100332220-1310122230020232-0320332222302203"></a>

<a id="canonical-2110113301312132-1220231231220110-1021312222221220-0112111222202002-2021221003011001-3032011130202203-3121121302221300-0323231000211320"></a>

## http_redirect property — https / 033300302223 / 7

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [non_default_loadbalancer](resources--workload--reference--group-006.md#canonical-3032202302032213-0303131002200121-0310313203203001-0303223311331203-3110132231230001-3300021313312233-0130110330103111-2132202022101200): complete subsection reference.

- [pass_through](resources--workload--reference--group-006.md#canonical-1020131000000021-0213323132231310-3302312202012121-1100210321023123-0023331100011120-0223121031212023-2031020223031102-0220123120021331): complete subsection reference.

<a id="canonical-2010303330210302-0010210232301303-0221303023302233-3311310203033131-2110101202101122-2220101012011123-3232002321331102-3120312221330331"></a>

<a id="canonical-2211020203023320-0120013312003032-1123330123012030-1211231333200020-3221211203100200-1313020101012022-1200110230021032-1322002110232000"></a>

## port property — https / 033300302223 / 8

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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

<a id="canonical-0232330220302203-1122312010302112-2200032121020220-0030100023111133-3323111333013210-1122131132113223-3200020201133321-2232213320230100"></a>

<a id="canonical-0110002111001113-1222230333120110-0213133111233113-2012231301303330-0030301330232111-1110210031101000-2020100113303011-3203100321001021"></a>

## port_ranges property — https / 033300302223 / 9

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

<a id="canonical-2321322132233112-0131011101123210-0221230331313233-1100230130111103-3311111321110013-0002013131013003-3130021030211220-2311200012230000"></a>

<a id="canonical-3123223032220230-1203323233012212-0330131220333313-0020003223222023-3333231130300001-2211332210212312-1232330122200223-3221230312130103"></a>

## server_name property — https / 033300302223 / 10

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320): complete subsection reference.

<a id="canonical-1010030020213031-0201311102313220-1020022222130222-2000122333203221-3202300002320320-2313320222112032-2320031000322012-3132112200211313"></a>

## Next pages — https / 033300302223 / 11

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-005.md#canonical-0033100331020001-2221313100101220-3310100020233022-2211123323300001-0231323312333323-0032322122221223-0222211200302233-3122211001332213)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header](resources--workload--reference--group-005.md#canonical-2222132001202300-2223100111302031-3221210230001201-1220312313130120-0001232202031003-0302013332210031-1130121020112230-2221102100233321)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer](resources--workload--reference--group-005.md#canonical-3013320033200112-2330211103132210-1111021302202201-1003233032230222-0300000132132211-0201133003111333-1012021102221030-2022022103103133)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize](resources--workload--reference--group-005.md#canonical-2130013022101113-3232010232323220-0210310301112110-1233302313322122-1310311322332312-2201211202201021-3333320313123113-2100323212032033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize](resources--workload--reference--group-005.md#canonical-3110011313233100-3302030313300031-1211110123230312-1103211030232121-1230111303122131-0000021202212203-0301121032130213-1003013003110302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-1302010300230312-3111312201201113-2013133300112103-0112233323311000-0220233130302132-1102033102333132-1100222323112330-1222302133120000)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer](resources--workload--reference--group-006.md#canonical-3032202302032213-0303131002200121-0310313203203001-0303223311331203-3110132231230001-3300021313312233-0130110330103111-2132202022101200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through](resources--workload--reference--group-006.md#canonical-1020131000000021-0213323132231310-3302312202012121-1100210321023123-0023331100011120-0223121031212023-2031020223031102-0220123120021331)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0033100331020001-2221313100101220-3310100020233022-2211123323300001-0231323312333323-0032322122221223-0222211200302233-3122211001332213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022220113031131-0100311311113113-1102332122202303-1101322020230100-3213220121121301-3213100100321221-1131111013112133-0120131110113103"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options — coalescing_options / 333232321131 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-0200200233121101-2312130202023321-3313321320223203-3022001130103001-3213312131212020-2010331303122221-1212303101112010-0313121300221130"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302020123003032-1211020123212123-2233313211000301-2211210310322330-2312111012110211-1001213023332322-1133033112110231-2021131331131003"></a>

## Direct properties — coalescing_options / 333232321131 / 3

- [default_coalescing](resources--workload--reference--group-005.md#canonical-2011103211301323-2322201220200001-0021120322330002-1322032122221322-3130200231031302-0300023322300202-1312311012022131-1022023212322312): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-005.md#canonical-2330233310310011-0212211300031020-3322121312332213-0033223203100100-3201021030200002-3213311222113032-1033131101120123-3133322010021232): complete subsection reference.

<a id="canonical-0010111001103111-0210003003210113-0103021202102021-1200131303310130-3312301000003211-3320301230021233-2123222020002302-3123332301223133"></a>

## Next pages — coalescing_options / 333232321131 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing](resources--workload--reference--group-005.md#canonical-2011103211301323-2322201220200001-0021120322330002-1322032122221322-3130200231031302-0300023322300202-1312311012022131-1022023212322312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing](resources--workload--reference--group-005.md#canonical-2330233310310011-0212211300031020-3322121312332213-0033223203100100-3201021030200002-3213311222113032-1033131101120123-3133322010021232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2011103211301323-2322201220200001-0021120322330002-1322032122221322-3130200231031302-0300023322300202-1312311012022131-1022023212322312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303332200200312-0101022200122211-0103222203231320-3321001120112320-3003020302023202-0230013211330320-3112200312100231-1010220130002010"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing — default_coalescing / 222001312202 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-005.md#canonical-0033100331020001-2221313100101220-3310100020233022-2211123323300001-0231323312333323-0032322122221223-0222211200302233-3122211001332213)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-1130220300322200-2203332331221023-1022203000222132-0201313320012113-0120300100111331-2020211331213002-3232101113020100-3133023231300031"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

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
default_coalescing = {}
```

<a id="canonical-1131032121110012-0301320020311132-1200232100001222-3031000132210000-3321122032310303-3022220233300231-3012103312022010-2000122033231031"></a>

## Direct properties — default_coalescing / 222001312202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310202320211003-2020030113220211-1121203313033121-2313230023030203-0333132220201333-0233003220130131-2123223230212302-3133112320211102"></a>

## Next pages — default_coalescing / 222001312202 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-005.md#canonical-0033100331020001-2221313100101220-3310100020233022-2211123323300001-0231323312333323-0032322122221223-0222211200302233-3122211001332213)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2330233310310011-0212211300031020-3322121312332213-0033223203100100-3201021030200002-3213311222113032-1033131101120123-3133322010021232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203223223201111-0121321031303311-0321023200000310-1110312231120001-2223302011021322-1032011101212312-0323300322321033-0312303012300203"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing — strict_coalescing / 223011331300 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-005.md#canonical-0033100331020001-2221313100101220-3310100020233022-2211123323300001-0231323312333323-0032322122221223-0222211200302233-3122211001332213)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-3322132222322002-3033302320213302-1300103130301030-1323031332313312-0103220020230323-2312203023130332-2232130003001000-3103130112010132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

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
strict_coalescing = {}
```

<a id="canonical-1120303020011023-3312030222331332-2002100302023121-3012121213201221-3103132231110030-0003211213333020-0000311232330300-2300223021000003"></a>

## Direct properties — strict_coalescing / 223011331300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101322030210310-2021330123023311-1121302331301300-2310310211000321-0002022203232212-1212121313223311-3112321211131221-2130031321100012"></a>

## Next pages — strict_coalescing / 223011331300 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-005.md#canonical-0033100331020001-2221313100101220-3310100020233022-2211123323300001-0231323312333323-0032322122221223-0222211200302233-3122211001332213)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2222132001202300-2223100111302031-3221210230001201-1220312313130120-0001232202031003-0302013332210031-1130121020112230-2221102100233321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121332321333330-1213122320200331-0233033031302212-3232030131323323-0020213011202100-2201331111121100-1320301121211030-3033110201000233"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header — default_header / 222002203001 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header

<a id="canonical-1122013321213203-2111123132020220-2033313100203212-3133223323213132-3231221001021032-2300003210132220-0201210300303220-1122101122331021"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default header.

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
default_header = {}
```

<a id="canonical-0102322010223003-1103003330102220-2002331113320200-3312023233331022-3220302200313203-1303122132223031-1330032223001003-2012221113022220"></a>

## Direct properties — default_header / 222002203001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301101113233130-0023012231320330-2221122132103133-2102111333103103-3001221320011100-3102210023230233-2301112203000121-2312233230331030"></a>

## Next pages — default_header / 222002203001 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3013320033200112-2330211103132210-1111021302202201-1003233032230222-0300000132132211-0201133003111333-1012021102221030-2022022103103133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030300313223132-2133312233111010-2303221330111203-1322301222123320-3011123330131310-0302131333330301-0212203003232323-1310003311202132"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer — default_loadbalancer / 213222213301 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-0200211031103332-3223222211230123-2011021032302232-0032223033312032-2103313223130033-3230201232202000-3022112123101311-2312331001203133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default loadbalancer.

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
default_loadbalancer = {}
```

<a id="canonical-3202003131202123-0231322312330130-3211033000332133-1333103101211211-0203031220332002-1231121021230312-3101303103002030-1122031232203112"></a>

## Direct properties — default_loadbalancer / 213222213301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310201203032032-0323212022131120-1200232221301101-3023313001002133-3031001300212213-2023211321303202-2332000213321300-0123312003312032"></a>

## Next pages — default_loadbalancer / 213222213301 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2130013022101113-3232010232323220-0210310301112110-1233302313322122-1310311322332312-2201211202201021-3333320313123113-2100323212032033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002120201332313-0021130212321002-2123300223213000-1303232332300333-0203121210231200-0223310113202013-0211200231212330-2220120022200233"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize — disable_path_normalize / 313322103112 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-0232233210001032-2103030001313001-0003213113222221-0003313222000033-0123210120332202-2033301233330012-2003103303011331-1223133300120123"></a>

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
disable_path_normalize = {}
```

<a id="canonical-1033320021303001-2220120113113323-3132330002300112-3323231310030112-1122323022200130-3313021212330001-0021301201322013-1333130212031010"></a>

## Direct properties — disable_path_normalize / 313322103112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102031030223220-1333111020213020-0230101230123022-1312201020321002-0233313222020331-2123031123132132-2333130120202221-0013122221101212"></a>

## Next pages — disable_path_normalize / 313322103112 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3110011313233100-3302030313300031-1211110123230312-1103211030232121-1230111303122131-0000021202212203-0301121032130213-1003013003110302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311030303001200-0001001103223300-0001133010210123-0222111022321222-0310233332221031-0221110020322201-0011032013302022-3013123211221310"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize — enable_path_normalize / 200330020003 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-0313021301312021-0010332331120301-3330010111113210-1102210112022221-1013012112331233-1112302012302102-1113023132012021-3233310021013211"></a>

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
enable_path_normalize = {}
```

<a id="canonical-2311120111230010-0210103133233331-1020312003012003-3212102120020120-1333321110331201-1113203233013023-2120213223120221-1300033103312112"></a>

## Direct properties — enable_path_normalize / 200330020003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033102112120000-2333130230333312-0003130010201031-3301100213030232-0211011312101000-3203330112203201-0200130302003202-2313211102323303"></a>

## Next pages — enable_path_normalize / 200330020003 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1302010300230312-3111312201201113-2013133300112103-0112233323311000-0220233130302132-1102033102333132-1100222323112330-1222302133120000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210001030123300-2010030130001020-0330122313020013-1112033102111013-3003203302232300-2113120030011203-0100201231112120-3111313212332130"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options — http_protocol_options / 213112303312 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-3330310011100132-2323120100022001-1211223032133202-1200310000201033-3100223013013203-1323133213123131-0330310203012333-3011123031230013"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
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
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110333103331112-2022232232023311-3302200023003010-0133323102032033-3122330210233010-1200221233112310-0311001011113313-0012132013222013"></a>

## Direct properties — http_protocol_options / 213112303312 / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-005.md#canonical-2233203112100231-1112132323232131-3031231033123302-1033010021010011-2320033101303110-3210223311031302-3220021330100000-3201021202130000): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-005.md#canonical-1230101320301000-2331012123121232-1220322002021310-1102100231333020-1333201310232113-0321100121033212-0110220333021011-2230100202300133): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-005.md#canonical-1022010202302222-1122122112032203-3210312021030223-0202032220213211-0030111310202023-3010303133303120-2220001020031113-0011221012021221): complete subsection reference.

<a id="canonical-1103322010113113-2000321132002211-1010031200231120-0011312231130303-1311113020331100-2302332003301220-1022300121233022-1031123323101123"></a>

## Next pages — http_protocol_options / 213112303312 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-005.md#canonical-2233203112100231-1112132323232131-3031231033123302-1033010021010011-2320033101303110-3210223311031302-3220021330100000-3201021202130000)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-005.md#canonical-1230101320301000-2331012123121232-1220322002021310-1102100231333020-1333201310232113-0321100121033212-0110220333021011-2230100202300133)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-005.md#canonical-1022010202302222-1122122112032203-3210312021030223-0202032220213211-0030111310202023-3010303133303120-2220001020031113-0011221012021221)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2233203112100231-1112132323232131-3031231033123302-1033010021010011-2320033101303110-3210223311031302-3220021330100000-3201021202130000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323130131322323-3332200030303113-1320233123300132-3201000311111201-1213012102112112-2132102013221110-3330332102111103-0200002111202001"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 003020131201 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-1302010300230312-3111312201201113-2013133300112103-0112233323311000-0220233130302132-1102033102333132-1100222323112330-1222302133120000)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0101333213030021-1230220102233121-2020303213113033-3103333301333230-3111103021210203-1031013233003102-1133311123313320-0013203222312100"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323302230033310-2032311303330333-3133131011332112-2230133133223123-0332330103121211-0301203332231000-0210300012002323-2322323101103310"></a>

## Direct properties — http_protocol_enable_v1_only / 003020131201 / 3

- [header_transformation](resources--workload--reference--group-005.md#canonical-0323023303133003-0312112223230022-1033310120000130-2201300300121333-1130011332211123-3213320213030310-2332123102322332-3331122132210132): complete subsection reference.

<a id="canonical-3012202211121000-2112210113021121-2103331202020002-1221212201322122-3113311111213223-3312100030331112-0300331030331202-0031321020220033"></a>

## Next pages — http_protocol_enable_v1_only / 003020131201 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-005.md#canonical-0323023303133003-0312112223230022-1033310120000130-2201300300121333-1130011332211123-3213320213030310-2332123102322332-3331122132210132)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-1302010300230312-3111312201201113-2013133300112103-0112233323311000-0220233130302132-1102033102333132-1100222323112330-1222302133120000)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0323023303133003-0312112223230022-1033310120000130-2201300300121333-1130011332211123-3213320213030310-2332123102322332-3331122132210132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012321211322212-3300010322200200-1101010033013013-1230031023330101-3211100222300123-0321000331021333-1100301022131110-0221112030032212"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 233010111320 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-1302010300230312-3111312201201113-2013133300112103-0112233323311000-0220233130302132-1102033102333132-1100222323112330-1222302133120000)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-005.md#canonical-2233203112100231-1112132323232131-3031231033123302-1033010021010011-2320033101303110-3210223311031302-3220021330100000-3201021202130000)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0312021130032011-1101001223203331-1320222313311123-0020001213313023-2203120100131202-3212213233120112-2221201212122332-3322323111033232"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211201131332021-2130210000120320-0302201232220023-0211321121011010-3333310300132002-2130130111020201-3031313132013232-1130033102013222"></a>

## Direct properties — header_transformation / 233010111320 / 3

- [default_header_transformation](resources--workload--reference--group-005.md#canonical-2220122231132211-1023210102120120-0100110222321100-1222220112111000-0320310132100131-3300102322100301-2202303310112122-0221220203301012): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-005.md#canonical-3211133200212211-3310201030031310-3013310012103130-2310201000303213-1122201321020112-0013313110120210-0303110033303303-0212002300123202): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-005.md#canonical-0311113130001213-2112121102313320-0030211011203131-2302330323120210-1111233303233221-2030110201012103-1032211033333213-2031200222210211): complete subsection reference.

<a id="canonical-3223213031102322-0303031311311113-0131023312322333-2112012330102320-3123301303023222-2111030302100223-3123201123010201-0032323000202103"></a>

## Next pages — header_transformation / 233010111320 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-005.md#canonical-2220122231132211-1023210102120120-0100110222321100-1222220112111000-0320310132100131-3300102322100301-2202303310112122-0221220203301012)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-005.md#canonical-3211133200212211-3310201030031310-3013310012103130-2310201000303213-1122201321020112-0013313110120210-0303110033303303-0212002300123202)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-005.md#canonical-0311113130001213-2112121102313320-0030211011203131-2302330323120210-1111233303233221-2030110201012103-1032211033333213-2031200222210211)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-005.md#canonical-2233203112100231-1112132323232131-3031231033123302-1033010021010011-2320033101303110-3210223311031302-3220021330100000-3201021202130000)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2220122231132211-1023210102120120-0100110222321100-1222220112111000-0320310132100131-3300102322100301-2202303310112122-0221220203301012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310233223102311-3001203203323310-0010303333211230-2303010303023303-3230102311330010-1023120131100323-3022031132133131-2212112323122320"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 121331201223 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-1302010300230312-3111312201201113-2013133300112103-0112233323311000-0220233130302132-1102033102333132-1100222323112330-1222302133120000)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-005.md#canonical-2233203112100231-1112132323232131-3031231033123302-1033010021010011-2320033101303110-3210223311031302-3220021330100000-3201021202130000)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-005.md#canonical-0323023303133003-0312112223230022-1033310120000130-2201300300121333-1130011332211123-3213320213030310-2332123102322332-3331122132210132)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-2331213010033110-3320313321012220-2023001030031232-2222232212202312-3121121130002323-1111000312212012-3300122301101003-0011202110201132"></a>

Type: `["object", {}]`. Optional.

Use the platform's current default HTTP header transformation behavior.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
default_header_transformation = {}
```

<a id="canonical-0033111221112200-3211310232232312-1001300000030212-3310211230111331-3112312203002002-3132132203333210-2121012203320223-2333102333111003"></a>

## Direct properties — default_header_transformation / 121331201223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031320303211233-1231311211102211-3331230012013211-2223122102222330-1321203322030000-2331020121322320-3220010133222031-2333202022013321"></a>

## Next pages — default_header_transformation / 121331201223 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-005.md#canonical-0323023303133003-0312112223230022-1033310120000130-2201300300121333-1130011332211123-3213320213030310-2332123102322332-3331122132210132)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3211133200212211-3310201030031310-3013310012103130-2310201000303213-1122201321020112-0013313110120210-0303110033303303-0212002300123202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101303110101033-3002110203010302-2131222231111001-3020210130102130-1003202111303300-1303021320201022-1310202312000032-2020010300131131"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 220230310231 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-1302010300230312-3111312201201113-2013133300112103-0112233323311000-0220233130302132-1102033102333132-1100222323112330-1222302133120000)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-005.md#canonical-2233203112100231-1112132323232131-3031231033123302-1033010021010011-2320033101303110-3210223311031302-3220021330100000-3201021202130000)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-005.md#canonical-0323023303133003-0312112223230022-1033310120000130-2201300300121333-1130011332211123-3213320213030310-2332123102322332-3331122132210132)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2233321021112233-2303123333031322-1002231312233112-3330331203102011-1020321132103000-0033013110011320-1311303230113000-2122210201021312"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
preserve_case_header_transformation = {}
```

<a id="canonical-2032332212013133-2010132201111230-0300200203120331-1133202130110123-3330023312310313-1010131131213010-2233231331101003-3033310200233232"></a>

## Direct properties — preserve_case_header_transformation / 220230310231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113012122010112-0312101313233300-0203112122300002-1101222332322123-1310303201301001-2313031233220203-1333210130032012-2203331000002121"></a>

## Next pages — preserve_case_header_transformation / 220230310231 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-005.md#canonical-0323023303133003-0312112223230022-1033310120000130-2201300300121333-1130011332211123-3213320213030310-2332123102322332-3331122132210132)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0311113130001213-2112121102313320-0030211011203131-2302330323120210-1111233303233221-2030110201012103-1032211033333213-2031200222210211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113333021010001-2301310301110021-3300203310003321-0102132022330332-0200022212302333-3023220231203203-0331213003222211-2032330122032222"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 210200103120 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-1302010300230312-3111312201201113-2013133300112103-0112233323311000-0220233130302132-1102033102333132-1100222323112330-1222302133120000)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-005.md#canonical-2233203112100231-1112132323232131-3031231033123302-1033010021010011-2320033101303110-3210223311031302-3220021330100000-3201021202130000)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-005.md#canonical-0323023303133003-0312112223230022-1033310120000130-2201300300121333-1130011332211123-3213320213030310-2332123102322332-3331122132210132)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-2301020200111123-0000131133030211-0301223131121103-3331213302200220-1130101221123321-1210320030110312-0231022131311300-2031102301001022"></a>

Type: `["object", {}]`. Optional.

Transform HTTP header names to proper case when explicit transformation is required.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
proper_case_header_transformation = {}
```

<a id="canonical-2212313222333230-0222302333232103-0212131300230312-1230212333333011-2032230032110333-3223212003313222-0130131323330013-3120133223101012"></a>

## Direct properties — proper_case_header_transformation / 210200103120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221213022001213-2101101000113110-2131323000121320-0132133113302330-1220320213332302-3013310111013031-0020220331320211-0022110010232132"></a>

## Next pages — proper_case_header_transformation / 210200103120 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-005.md#canonical-0323023303133003-0312112223230022-1033310120000130-2201300300121333-1130011332211123-3213320213030310-2332123102322332-3331122132210132)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1230101320301000-2331012123121232-1220322002021310-1102100231333020-1333201310232113-0321100121033212-0110220333021011-2230100202300133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022313011301202-3010231033313002-3312130112123133-1031211121023313-3003002012312132-1111103013030102-2300003320131131-1001102111030123"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 303013220211 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-1302010300230312-3111312201201113-2013133300112103-0112233323311000-0220233130302132-1102033102333132-1100222323112330-1222302133120000)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-0032121012212201-1312321003130221-0200011123212301-1121120222320110-1321022023121310-1103213122031203-3130122032101322-1312100113202101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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
http_protocol_enable_v1_v2 = {}
```

<a id="canonical-2120120123021001-1113112203233233-0000122112203133-2231103131222201-1021230331133023-1211210022232322-3213121013013211-1202023020311311"></a>

## Direct properties — http_protocol_enable_v1_v2 / 303013220211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201130233023102-0112021121100231-0201233312110231-1333023123000302-1232212333003333-0233213221330200-2003132212200213-2201231101112031"></a>

## Next pages — http_protocol_enable_v1_v2 / 303013220211 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-1302010300230312-3111312201201113-2013133300112103-0112233323311000-0220233130302132-1102033102333132-1100222323112330-1222302133120000)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1022010202302222-1122122112032203-3210312021030223-0202032220213211-0030111310202023-3010303133303120-2220001020031113-0011221012021221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
