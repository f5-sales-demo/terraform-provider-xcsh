---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2001221201223331-0132231331300033-0200210021233110-1101022212212200-3020010312323102-1323101211210203-2112221302301201-2121301301121220"></a>

## Direct properties — mount / 112220020011 / 3

<a id="canonical-1202111233111003-2012021012123312-2333121312023103-0131110030230022-0310310000323300-2100131020211233-0202312110011232-0222222132113323"></a>

<a id="canonical-3031101120121103-0322012312232210-1211212202311230-3110003330321030-2031323113033003-1121030330031030-2213103233123201-0101133023131333"></a>

## mode property — mount / 112220020011 / 4

Type: `"string"`. Computed.

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

<a id="canonical-3322330232021103-1211030331111320-0031010220303232-2031010333020000-2201101121111201-0313222213211130-0031210033221130-3230302003233323"></a>

<a id="canonical-3223213121132020-2333323303112112-1212010333021101-2102012131233002-1231020303313233-3023131103000132-3312102213233311-2331112100101303"></a>

## mount_path property — mount / 112220020011 / 5

Type: `"string"`. Computed.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3312223013002300-2200103213313212-2020230133110102-2200033301201031-3222131200020031-1223030212203011-3002120303113302-2112200010233213"></a>

<a id="canonical-0200333112003320-0103222010133012-3013133213331231-3030220323012211-3322203211332100-3210130103310322-2333220333333232-1031231200032031"></a>

## sub_path property — mount / 112220020011 / 6

Type: `"string"`. Computed.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1132311130332003-2212202303103311-1023203131312202-2312112211223003-3113130312301222-1033113101330020-3032030111101110-2033120331022320"></a>

## Next pages — mount / 112220020011 / 7

- [job.volumes.host_path](data-sources--workload--reference--group-004.md#canonical-1101123330022122-0312023230320121-2333332030310233-1331312111301103-0120032000312031-2200113123130032-2333123300102003-0212211023011122)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3011233111030203-0033211202203110-3130100112113111-2300013120101211-1010133312301103-0331203321122300-2120131033331030-1130111020230122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300231213311111-2312022103113030-0302030303223203-1233210000103011-2102323123010320-3313000230313131-2001033303123122-3221210313230330"></a>

## job.volumes.persistent_volume — persistent_volume / 012010101212 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.volumes](data-sources--workload--reference--group-004.md#canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011)
- job.volumes.persistent_volume

<a id="canonical-0233313130310321-0131232032110223-1003303130023333-0113210030202110-1333330331002000-0211200300312120-2202002303300002-0021023112233221"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2111333123331212-1323011032103202-0020331131302313-2211131032203320-1030323210123202-3133200302203122-3230200023333031-0233103313102321"></a>

## Direct properties — persistent_volume / 012010101212 / 3

- [mount](data-sources--workload--reference--group-005.md#canonical-2010021302120122-2302212220000131-0331121210311130-1132100012321020-2300301213211223-1103022231030103-3223102120121032-3230310011031333): complete subsection reference.

- [storage](data-sources--workload--reference--group-005.md#canonical-2213202331120023-2231110111323223-1301320301300323-2103321202200000-2113211210103200-1220133202221033-2101020223110213-3203011223010331): complete subsection reference.

<a id="canonical-1112002322221102-2311333223331312-2102113232120023-1332031213331221-1110301113030121-3133320311323032-1022022210302013-1300011301302320"></a>

## Next pages — persistent_volume / 012010101212 / 4

- [job.volumes.persistent_volume.mount](data-sources--workload--reference--group-005.md#canonical-2010021302120122-2302212220000131-0331121210311130-1132100012321020-2300301213211223-1103022231030103-3223102120121032-3230310011031333)
- [job.volumes.persistent_volume.storage](data-sources--workload--reference--group-005.md#canonical-2213202331120023-2231110111323223-1301320301300323-2103321202200000-2113211210103200-1220133202221033-2101020223110213-3203011223010331)
- [job.volumes](data-sources--workload--reference--group-004.md#canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2010021302120122-2302212220000131-0331121210311130-1132100012321020-2300301213211223-1103022231030103-3223102120121032-3230310011031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003123031221212-2011122320322011-2323011310030311-0011031230122122-0023302322303112-0313123200233112-1111210102222202-2203202221103000"></a>

## job.volumes.persistent_volume.mount — mount / 330233332303 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.volumes](data-sources--workload--reference--group-004.md#canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011)
- [job.volumes.persistent_volume](data-sources--workload--reference--group-005.md#canonical-3011233111030203-0033211202203110-3130100112113111-2300013120101211-1010133312301103-0331203321122300-2120131033331030-1130111020230122)
- job.volumes.persistent_volume.mount

<a id="canonical-2301102032332330-0323131113112130-1023320001030031-0023113120002002-0003003303231020-2000333023311123-0003330122321232-2023122103113002"></a>

Type: `"single"`. Computed.

Volume mount describes how volume is mounted inside a workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2012213233202333-0233303120011021-0333330133220230-3303130212123011-3331023121310210-2313221132002033-3021220333000221-1311132302022122"></a>

## Direct properties — mount / 330233332303 / 3

<a id="canonical-1220022010120222-3130130220133103-0333112311020013-1210223021133331-1100302002331010-2010133320212122-0111001130210330-2213020333020230"></a>

<a id="canonical-1303303200033023-1002313123313022-3111122033113220-0333220003200220-3322032213111222-1233212212320133-1022032331203322-0031031220120033"></a>

## mode property — mount / 330233332303 / 4

Type: `"string"`. Computed.

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

<a id="canonical-3201313312213302-3311003103200220-2102223212300011-1211000313310120-1223202313212223-0031223231133003-3020211020323123-1321231331011000"></a>

<a id="canonical-1113312302002021-3000023031022230-2020203111032320-0301233101201233-0200012132210120-2322221133010101-1231023022113012-2133033010013020"></a>

## mount_path property — mount / 330233332303 / 5

Type: `"string"`. Computed.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0202311201231023-2103331210112320-3103300211313330-2323122112023223-1132313033030112-0031302300121302-3200200033120233-1230332212112212"></a>

<a id="canonical-0332200020000023-0313213200020022-3112101031113101-2103213313202103-3101232212201202-2320202220313012-3022301201233123-1103323011033101"></a>

## sub_path property — mount / 330233332303 / 6

Type: `"string"`. Computed.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3320131120023121-2123332300311310-3003113011102210-1022133201132131-2022003330200200-1131113222102213-0113322102220233-3221123333101213"></a>

## Next pages — mount / 330233332303 / 7

- [job.volumes.persistent_volume](data-sources--workload--reference--group-005.md#canonical-3011233111030203-0033211202203110-3130100112113111-2300013120101211-1010133312301103-0331203321122300-2120131033331030-1130111020230122)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2213202331120023-2231110111323223-1301320301300323-2103321202200000-2113211210103200-1220133202221033-2101020223110213-3203011223010331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301231112323020-0113230100000212-2212020111010033-1120232312201323-2312300221021101-2102223203031230-1031001103103030-1213103232001130"></a>

## job.volumes.persistent_volume.storage — storage / 222121112122 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.volumes](data-sources--workload--reference--group-004.md#canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011)
- [job.volumes.persistent_volume](data-sources--workload--reference--group-005.md#canonical-3011233111030203-0033211202203110-3130100112113111-2300013120101211-1010133312301103-0331203321122300-2120131033331030-1130111020230122)
- job.volumes.persistent_volume.storage

<a id="canonical-1133313113333331-0113021203200232-0323101023131132-2222323202131120-0001232203231303-0032031301233302-2330331332213222-0123131112012311"></a>

Type: `"single"`. Computed.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Upstream description:

Persistent storage configuration is used to configure Persistent Volume Claim (PVC)

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

<a id="canonical-0021012032223202-1212011213303021-1213311321131023-2001103211000230-0003122100311030-2322323330310130-2233201233123121-0201122112003131"></a>

## Direct properties — storage / 222121112122 / 3

<a id="canonical-3031230232020210-3332201001000002-1213031313223022-0312220031000013-3023320012320222-2003330222022320-1031033132323323-0011210033111310"></a>

<a id="canonical-0230333132122310-3302120120000213-1021300110313310-1101133231232123-3221202303303223-0221003232031303-2023230112002220-2120232130120032"></a>

## access_mode property — storage / 222121112122 / 4

Type: `"string"`. Computed.

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

<a id="canonical-1030321013112223-0012322202200231-3200300030221000-2002030202020103-3011010021131003-3100321320120302-1212333212220022-0132022221203023"></a>

<a id="canonical-1312321322122330-0302331231322201-2000103230101223-3113303123021322-3002020120210203-3233101102332131-3032130032203002-3331313213323022"></a>

## class_name property — storage / 222121112122 / 5

Type: `"string"`. Computed.

Exclusive with \[default\] Use the specified class name.

Upstream description:

Exclusive with \[default\] Use the specified class name.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [default](data-sources--workload--reference--group-005.md#canonical-3031333332222000-0002301011111123-2020200223121303-0302222312002201-2011203021303011-2011031030112301-1112303033132013-2102231120000012): complete subsection reference.

<a id="canonical-3032030303333030-1131201333223310-0333133101133311-3001120202030013-0110211200203303-3102211022112202-1112232223030103-3023021302212320"></a>

<a id="canonical-2020203002220011-1010223312101122-3101103131203003-0130003302200021-1330003220230031-1321210103320012-2013203320003123-1203023312302300"></a>

## storage_size property — storage / 222121112122 / 6

Type: `"number"`. Computed.

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

<a id="canonical-3133302333213010-3212001031031013-1001203110201312-0111132020132321-2213212132122212-0223121123110220-1122103231032110-3332122212111323"></a>

## Next pages — storage / 222121112122 / 7

- [job.volumes.persistent_volume.storage.default](data-sources--workload--reference--group-005.md#canonical-3031333332222000-0002301011111123-2020200223121303-0302222312002201-2011203021303011-2011031030112301-1112303033132013-2102231120000012)
- [job.volumes.persistent_volume](data-sources--workload--reference--group-005.md#canonical-3011233111030203-0033211202203110-3130100112113111-2300013120101211-1010133312301103-0331203321122300-2120131033331030-1130111020230122)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3031333332222000-0002301011111123-2020200223121303-0302222312002201-2011203021303011-2011031030112301-1112303033132013-2102231120000012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200123220300003-3133011220312221-2100313031020012-0133232030303211-3012312102030003-2111331231110130-2133102321313312-3321111213111012"></a>

## job.volumes.persistent_volume.storage.default — default / 212221221022 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.volumes](data-sources--workload--reference--group-004.md#canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011)
- [job.volumes.persistent_volume](data-sources--workload--reference--group-005.md#canonical-3011233111030203-0033211202203110-3130100112113111-2300013120101211-1010133312301103-0331203321122300-2120131033331030-1130111020230122)
- [job.volumes.persistent_volume.storage](data-sources--workload--reference--group-005.md#canonical-2213202331120023-2231110111323223-1301320301300323-2103321202200000-2113211210103200-1220133202221033-2101020223110213-3203011223010331)
- job.volumes.persistent_volume.storage.default

<a id="canonical-2331303313212213-3033311322132033-3213230131121102-2002321003221021-3033232321311103-1101101232101230-3332033321123302-3323300212022132"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1013023232131113-0130130301011231-2122121120312212-0330321223010330-0030112211212110-0122303002300122-1103121000010032-1123332131332333"></a>

## Direct properties — default / 212221221022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202200010221111-2303230013033111-2313023020030210-3201310112011130-1322013302033230-2201322201312312-2323333032001113-0123132311100212"></a>

## Next pages — default / 212221221022 / 4

- [job.volumes.persistent_volume.storage](data-sources--workload--reference--group-005.md#canonical-2213202331120023-2231110111323223-1301320301300323-2103321202200000-2113211210103200-1220133202221033-2101020223110213-3203011223010331)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213321311112322-0123030330020013-1033031333333121-0122333201200230-1010002000203300-3132202123130323-2110311312211020-1022032002221212"></a>

## service — service / 222001010030 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- service

<a id="canonical-3033313320232002-0000103223202222-2330001311332233-1012320230331333-0131031100203120-3310301322123012-3111111013032000-3120131111111003"></a>

Type: `"single"`. Computed.

Service does not maintain per replica state, however it can be configured to use persistent storage
that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable
network identity or storage. Common examples of services are web servers, application servers..

Upstream description:

Service does not maintain per replica state, however it can be configured to use persistent storage
that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable
network identity or storage. Common examples of services are web servers, application servers,
traditional SQL databases, etc.

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

<a id="canonical-0333113322000120-1031210112210330-1312013301011320-0112130331103100-2203123113112330-3102112213002333-1200021000113011-3333203233311130"></a>

## Direct properties — service / 222001010030 / 3

- [advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333): complete subsection reference.

- [configuration](data-sources--workload--reference--group-015.md#canonical-3123002002331230-0201233111020120-1333221333320003-3110223112030011-2003333010313113-2333211001212031-3333112223203131-0322131002310210): complete subsection reference.

- [containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033): complete subsection reference.

- [deploy_options](data-sources--workload--reference--group-016.md#canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320): complete subsection reference.

<a id="canonical-3033322032211333-1002012201102210-3311003332200001-2230320021303332-1231101322120211-0020012313011102-1312312032131100-3011312122022012"></a>

<a id="canonical-3112102100100302-3321000310200201-2120233300312331-1103221320123012-3120113032121210-1322101210231010-1011101312220001-0021231320333023"></a>

## num_replicas property — service / 222001010030 / 4

Type: `"number"`. Computed.

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Upstream description:

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [scale_to_zero](data-sources--workload--reference--group-016.md#canonical-1232232223133202-3202330012213121-2321010021332212-2131311121001321-2011131113311031-0312303012321112-3333213331000032-1000100101021221): complete subsection reference.

- [volumes](data-sources--workload--reference--group-016.md#canonical-2012110033111201-2323122231313121-3301031330103210-2312132322110233-3323213233311103-0320232202220320-3231001222133133-1333230032132310): complete subsection reference.

<a id="canonical-0212330110211123-2201302201120130-2033301321103301-2032120232011012-3130300031202113-2121333310032301-2320201223210101-2011210031120023"></a>

## Next pages — service / 222001010030 / 5

- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.configuration](data-sources--workload--reference--group-015.md#canonical-3123002002331230-0201233111020120-1333221333320003-3110223112030011-2003333010313113-2333211001212031-3333112223203131-0322131002310210)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.deploy_options](data-sources--workload--reference--group-016.md#canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320)
- [service.scale_to_zero](data-sources--workload--reference--group-016.md#canonical-1232232223133202-3202330012213121-2321010021332212-2131311121001321-2011131113311031-0312303012321112-3333213331000032-1000100101021221)
- [service.volumes](data-sources--workload--reference--group-016.md#canonical-2012110033111201-2323122231313121-3301031330103210-2312132322110233-3323213233311103-0320232202220320-3231001222133133-1333230032132310)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333022032223101-3022322211323203-1330201330032003-1213103133202112-3002121132230101-2223222331021110-1003101203201201-0223213211322211"></a>

## service.advertise_options — advertise_options / 101022100010 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- service.advertise_options

<a id="canonical-2111112122123010-3312212303211300-3011110031232000-2122013331121023-0302003021221312-0100321310010132-2120030223023101-2323320131131121"></a>

Type: `"single"`. Computed.

Advertise OPTIONS are used to configure how and where to advertise the workload using load
balancers.

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

<a id="canonical-2221301022213212-0033203112133302-2113323021132330-0231002023020011-2312003003311133-1021331220103220-3212323201321333-0012302220311130"></a>

## Direct properties — advertise_options / 101022100010 / 3

- [advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202): complete subsection reference.

- [advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-2131020211211101-0213303111310032-0232312121101303-0303013213123313-0112311302213302-2012200033230321-1231231303121203-0121103301113123): complete subsection reference.

- [advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130): complete subsection reference.

- [do_not_advertise](data-sources--workload--reference--group-015.md#canonical-3303332030303022-2112310112031201-1232033023113213-2310022310310212-0130333030130000-2200322102123323-3201131322011030-3120330230033000): complete subsection reference.

<a id="canonical-2100022332233210-1332233100231210-3112311222031021-2131030333102002-3013232213121322-2132332313310210-2130223301111013-1030131123300332"></a>

## Next pages — advertise_options / 101022100010 / 4

- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-2131020211211101-0213303111310032-0232312121101303-0303013213123313-0112311302213302-2012200033230321-1231231303121203-0121103301113123)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.do_not_advertise](data-sources--workload--reference--group-015.md#canonical-3303332030303022-2112310112031201-1232033023113213-2310022310310212-0130333030130000-2200322102123323-3201131322011030-3120330230033000)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232001112121033-1213130020020113-0131213222310210-1001301323103233-2033312221010112-2233201210100222-1310022330320100-0010300300131330"></a>

## service.advertise_options.advertise_custom — advertise_custom / 223332010211 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- service.advertise_options.advertise_custom

<a id="canonical-0320002303303233-1133333031302200-3112000002132103-1002313233333012-0030203231021001-3212311010020303-1222231120010030-2223221301211033"></a>

Type: `"single"`. Computed.

Advertise this workload via loadbalancer on specific sites.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3223301131013200-2111303022101333-0131220213033011-3322323202032231-2322300122110113-0030332230231221-3213110310010020-0111112231130202"></a>

## Direct properties — advertise_custom / 223332010211 / 3

- [advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120): complete subsection reference.

- [ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102): complete subsection reference.

<a id="canonical-0123021202312111-3301332213030011-3221321322123333-0111322330110112-0320202002211230-2131232323210120-2200330220301221-3213311023001113"></a>

## Next pages — advertise_custom / 223332010211 / 4

- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200132002321331-0220011230311103-1212332110133031-1311330123311230-0001212212021301-1020103013200032-3230321201202102-1330013301121220"></a>

## service.advertise_options.advertise_custom.advertise_where — advertise_where / 030011103113 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- service.advertise_options.advertise_custom.advertise_where

<a id="canonical-0132031021332232-3310223001033103-2201030130333220-1302301213231102-3102321210131300-3020223133233323-3212323130333320-3122012330333112"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3113003323201203-1301000213033312-2300331133013310-2123221310310321-1011130021322332-0230010133012012-0212232030033221-3321220313031101"></a>

## Direct properties — advertise_where / 030011103113 / 3

- [site](data-sources--workload--reference--group-005.md#canonical-0210013023032210-1303200212213212-0113330010223002-3121320332310012-2133211001112103-2302323102210031-3211012011131201-0112121032223032): complete subsection reference.

- [virtual_site](data-sources--workload--reference--group-005.md#canonical-2112123100111000-3331023313331020-0233223333000211-0132023311103233-2201330132013110-1212313303103212-1033013200123000-3033112112202011): complete subsection reference.

- [vk8s_service](data-sources--workload--reference--group-005.md#canonical-2232032323312210-1320210021123102-2323123132332020-2021233311111231-1003000211331102-3032322003100330-1011330303201322-1303131022011320): complete subsection reference.

<a id="canonical-2323023220222303-2010003021033123-0103333322312300-1212001130113300-2311101232213013-3010323220023200-3202021001332123-2323320202220333"></a>

## Next pages — advertise_where / 030011103113 / 4

- [service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--reference--group-005.md#canonical-0210013023032210-1303200212213212-0113330010223002-3121320332310012-2133211001112103-2302323102210031-3211012011131201-0112121032223032)
- [service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--reference--group-005.md#canonical-2112123100111000-3331023313331020-0233223333000211-0132023311103233-2201330132013110-1212313303103212-1033013200123000-3033112112202011)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-005.md#canonical-2232032323312210-1320210021123102-2323123132332020-2021233311111231-1003000211331102-3032322003100330-1011330303201322-1303131022011320)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0210013023032210-1303200212213212-0113330010223002-3121320332310012-2133211001112103-2302323102210031-3211012011131201-0112121032223032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211011101013302-3121100301211103-0022131122020311-0022232232002221-2330202301323212-3222013032312002-3232231103322323-1302023211013130"></a>

## service.advertise_options.advertise_custom.advertise_where.site — site / 013033321302 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- service.advertise_options.advertise_custom.advertise_where.site

<a id="canonical-2013031300033220-2222211301013030-0222301231211010-1233101000331333-0221103132121002-1013131033111112-0112030331331011-0012220021221210"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0210122303333321-2230033001312211-0232112212220103-0121030012303120-0311201033311112-2302203333130113-3300311212310112-0200020312111301"></a>

## Direct properties — site / 013033321302 / 3

<a id="canonical-3103122022111023-0130030220320133-1233310022331212-1012010231100031-1013121113130301-0300330003333202-2132332120022320-2011111100011100"></a>

<a id="canonical-0110200120011113-1202321310102211-3030030220013213-0131112123022021-0120102300311303-1320302322121311-2320220101121120-1200231312331223"></a>

## ip property — site / 013033321302 / 4

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1213300211031021-2230312332223021-2210220121001332-1322012213203003-2131202230220001-0230233102012232-1001222030003131-3201030313002012"></a>

<a id="canonical-2131000020003301-1121112023022303-3333112213202323-1330102331023030-2203210231230331-0012332231321110-1132000203313232-0200000100313120"></a>

## network property — site / 013033321302 / 5

Type: `"string"`. Computed.

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

- [site](data-sources--workload--reference--group-005.md#canonical-0130221011231112-0333020231020303-0033321013223220-3100120222201331-2320332021131131-1002223001132122-2230012033013322-3030303323010032): complete subsection reference.

<a id="canonical-1133011213030301-2320012021310323-0123032223110130-3330103010322101-3300222302212202-1202101310330032-1121103012301202-2211230211232333"></a>

## Next pages — site / 013033321302 / 6

- [service.advertise_options.advertise_custom.advertise_where.site.site](data-sources--workload--reference--group-005.md#canonical-0130221011231112-0333020231020303-0033321013223220-3100120222201331-2320332021131131-1002223001132122-2230012033013322-3030303323010032)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0130221011231112-0333020231020303-0033321013223220-3100120222201331-2320332021131131-1002223001132122-2230012033013322-3030303323010032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031310103331320-3022022333110121-0132100213031102-1102311302111323-3200210032322322-3322201131212301-3132033302020033-0200310311302231"></a>

## service.advertise_options.advertise_custom.advertise_where.site.site — site / 332233101033 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- [service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--reference--group-005.md#canonical-0210013023032210-1303200212213212-0113330010223002-3121320332310012-2133211001112103-2302323102210031-3211012011131201-0112121032223032)
- service.advertise_options.advertise_custom.advertise_where.site.site

<a id="canonical-1303030110021132-0212122032100033-3002303022100012-0212311233031211-0200300332112121-2311331233321022-2013112003011130-3102303133011123"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-1113010131323012-1302031222133033-3123200023120332-3020023020102030-0032023032200210-1302312321132023-3110133032221333-1120023320111031"></a>

## Direct properties — site / 332233101033 / 3

<a id="canonical-2020102011302220-3201203331211113-0022210131133110-0103113131210033-3231300223311120-1103323230223303-0321232203330131-2012120110302113"></a>

<a id="canonical-0011000100213202-0302022032000312-1032020032221313-2111002002100132-1320302102110001-1302320333111131-3313013310322332-0100103322213121"></a>

## name property — site / 332233101033 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0131032220022101-2333113332130103-3012203222020021-0212201332120233-0103210333031233-1023112301322231-0203123030303032-3333200013201113"></a>

<a id="canonical-0110233330200312-1133300130220030-1131323020332302-1001022131301302-2020333122333010-2201033331313030-3020030303123222-3232321130203133"></a>

## namespace property — site / 332233101033 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-1010331112001102-3321023021023113-0323223201321113-3230032023010312-3213110000033000-1132003322231330-2111023322200331-3123133031133102"></a>

<a id="canonical-1120302012032223-0022103122311120-3123001301201012-1122232220132333-2212301323033233-1310203033311031-3032123130033222-3200323300330323"></a>

## tenant property — site / 332233101033 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0121031102000321-2110221012102011-3121123033023132-3203210131031133-3132031022231323-3222301013011122-0032201310013013-0311201310033222"></a>

## Next pages — site / 332233101033 / 7

- [service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--reference--group-005.md#canonical-0210013023032210-1303200212213212-0113330010223002-3121320332310012-2133211001112103-2302323102210031-3211012011131201-0112121032223032)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2112123100111000-3331023313331020-0233223333000211-0132023311103233-2201330132013110-1212313303103212-1033013200123000-3033112112202011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201133030002202-0010211122100311-0031121100012030-1000002212013113-0130120210003030-3000011111100322-1332003020010313-0202333122202303"></a>

## service.advertise_options.advertise_custom.advertise_where.virtual_site — virtual_site / 012333201113 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- service.advertise_options.advertise_custom.advertise_where.virtual_site

<a id="canonical-1232231102031020-2031301000002113-3121033311113200-0123213122023222-0102301312110333-1321333012003012-2112121232320320-0133231131310003"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0310001033203213-2132030332022312-2321330232223012-1030002130211001-2212023310101333-3213111112233330-3220331120120102-1110113102330113"></a>

## Direct properties — virtual_site / 012333201113 / 3

<a id="canonical-1113101220321022-3013313331113112-1222211131213230-0132210333301103-1000003213032210-0101010033030321-1111023032021213-1010302131000301"></a>

<a id="canonical-0320230301201220-3023100130101331-1212221222123030-0313101212303110-3020232030221010-3230231330220103-1232210132011300-2312303320132100"></a>

## network property — virtual_site / 012333201113 / 4

Type: `"string"`. Computed.

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

- [virtual_site](data-sources--workload--reference--group-005.md#canonical-1100031201330203-1100112112130230-0030033303022010-3323302131320331-2313203032003003-0120033213031110-2300123303211331-3332102012011213): complete subsection reference.

<a id="canonical-3021200100021300-0310213312300310-0131132201000220-3033133231000300-1103202112120321-3032133330301002-1000132200221010-0310201201001321"></a>

## Next pages — virtual_site / 012333201113 / 5

- [service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--workload--reference--group-005.md#canonical-1100031201330203-1100112112130230-0030033303022010-3323302131320331-2313203032003003-0120033213031110-2300123303211331-3332102012011213)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1100031201330203-1100112112130230-0030033303022010-3323302131320331-2313203032003003-0120033213031110-2300123303211331-3332102012011213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231202123330300-0303331010200020-0321120133010220-3002211220212330-0113202022000120-2213332223030023-3220130211320030-1021232211120303"></a>

## service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site — virtual_site / 023232001013 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- [service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--reference--group-005.md#canonical-2112123100111000-3331023313331020-0233223333000211-0132023311103233-2201330132013110-1212313303103212-1033013200123000-3033112112202011)
- service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-0231221030213202-0102021130103102-0121120330223312-3103231100223130-1222211310221003-2121101320311023-0032100300022231-1313023200202000"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-2301112300102200-3203132113121033-1000233021031202-0121021021110033-1132020303323202-2101130102030123-1220303103210221-1003221211233213"></a>

## Direct properties — virtual_site / 023232001013 / 3

<a id="canonical-3023100231333232-3201322112312011-1033120210033033-0211220013003311-1211220320131322-1110002023100202-2303322021233030-3211100333311222"></a>

<a id="canonical-1330103013202223-0030323330300100-0020220333112230-0321230032213223-2031022121100032-0012211323113023-3202332210102031-0231233233130023"></a>

## name property — virtual_site / 023232001013 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3002101231011210-0213013100202020-2123123331011321-2221233031302111-1323020213113233-3322020213013021-2130211111020132-1221100303023021"></a>

<a id="canonical-1031112110200033-2323103313321013-0311331031021203-3311220222232032-1102320021323021-2003023010023100-0013331100020113-2210031322011332"></a>

## namespace property — virtual_site / 023232001013 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-1013131322221321-2212330131033133-0003201211101313-3132332022320322-1201130132320201-3222300112311221-2323123213313220-3321321321223100"></a>

<a id="canonical-1121132310132002-1001031123110123-2333013311113203-3210100301233331-3102203330023202-3112331000110001-2231311201003313-1133120330302022"></a>

## tenant property — virtual_site / 023232001013 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0211100222201320-1020112202223232-3331210031013313-3132320010120110-0121233101200113-1301113330113120-3121002303122010-3011101211033322"></a>

## Next pages — virtual_site / 023232001013 / 7

- [service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--reference--group-005.md#canonical-2112123100111000-3331023313331020-0233223333000211-0132023311103233-2201330132013110-1212313303103212-1033013200123000-3033112112202011)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2232032323312210-1320210021123102-2323123132332020-2021233311111231-1003000211331102-3032322003100330-1011330303201322-1303131022011320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003303222030103-2013002220220230-3201003323232211-2203030312031323-0102320132033300-2123133323023101-1002111111030233-0200022021231231"></a>

## service.advertise_options.advertise_custom.advertise_where.vk8s_service — vk8s_service / 113130201122 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service

<a id="canonical-2220321101212303-3121021023003030-0013131231220312-0032003002331331-1230130301333031-0030302312303301-2222100110332333-1112100001203002"></a>

Type: `"single"`. Computed.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

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

<a id="canonical-2112110002102310-3200323321301130-1131110230123230-3031302122301023-2122130122231331-0323223133112033-1231123032121103-3021133132221111"></a>

## Direct properties — vk8s_service / 113130201122 / 3

- [site](data-sources--workload--reference--group-005.md#canonical-0323310031021201-3000011021131322-0303113302302030-3213110111310120-1100202322111223-2233303021030113-3011200312221033-1303310313222331): complete subsection reference.

- [virtual_site](data-sources--workload--reference--group-005.md#canonical-1332131311232133-2211310100013233-3200002302333013-3300200321100312-1212031232231123-0120222202220200-2021332303332132-1020033010323130): complete subsection reference.

<a id="canonical-0202313020333013-2100223223233200-0301103131001123-2333213303330010-1210233031331332-2003212121311331-1321330311101033-2013012201223102"></a>

## Next pages — vk8s_service / 113130201122 / 4

- [service.advertise_options.advertise_custom.advertise_where.vk8s_service.site](data-sources--workload--reference--group-005.md#canonical-0323310031021201-3000011021131322-0303113302302030-3213110111310120-1100202322111223-2233303021030113-3011200312221033-1303310313222331)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--workload--reference--group-005.md#canonical-1332131311232133-2211310100013233-3200002302333013-3300200321100312-1212031232231123-0120222202220200-2021332303332132-1020033010323130)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0323310031021201-3000011021131322-0303113302302030-3213110111310120-1100202322111223-2233303021030113-3011200312221033-1303310313222331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221203330300312-2023230321133120-1100031100220330-0022211033023321-3223111320011222-1310212121320121-0023223112010323-1033102000033312"></a>

## service.advertise_options.advertise_custom.advertise_where.vk8s_service.site — site / 203330030330 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-005.md#canonical-2232032323312210-1320210021123102-2323123132332020-2021233311111231-1003000211331102-3032322003100330-1011330303201322-1303131022011320)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-2000021112203331-0030201103032320-2120310311211233-0322011321211300-1121231333202212-3130122022011303-2002022312221102-1231111033230210"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-0211112230220131-3003022010301332-2200202123220110-2030302310011310-2322222003023303-3333313102120232-0132112321131102-0110332001332233"></a>

## Direct properties — site / 203330030330 / 3

<a id="canonical-2232103333322210-0010112101112311-1130113131110221-1220010102021130-0123100023323012-0233010203020211-3233223033203113-2122300010333222"></a>

<a id="canonical-0201231301203000-1330302300020212-2023222310012103-3303313023120122-3120030033322132-0212001113011220-0033030102030130-0021300330310033"></a>

## name property — site / 203330030330 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2222211203013030-3233321323231032-0032031301123010-2012103100312230-1323310030220133-1132101313310322-1220102010023111-0101202002013331"></a>

<a id="canonical-1311202122312210-2001132201130302-1201323031231002-2011101032230301-2030111023231212-2011220320312131-0322332023112222-1310132033132322"></a>

## namespace property — site / 203330030330 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3000213033102032-2012131330232321-0100322112202202-0201002102013232-0012300323301212-1013113012302110-2223010222001310-2033331120023123"></a>

<a id="canonical-3103022101201201-2032101232132212-3301001012212101-2002020301203010-1033302310321110-1212010212322303-1023120322201121-3132201201300112"></a>

## tenant property — site / 203330030330 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0212311201110332-1001000100322020-3201123323123132-0121322303131202-3332212022311202-3133030010322213-3223330322013112-3211313121202011"></a>

## Next pages — site / 203330030330 / 7

- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-005.md#canonical-2232032323312210-1320210021123102-2323123132332020-2021233311111231-1003000211331102-3032322003100330-1011330303201322-1303131022011320)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1332131311232133-2211310100013233-3200002302333013-3300200321100312-1212031232231123-0120222202220200-2021332303332132-1020033010323130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211332031010230-0201303020020333-1212202321212211-3210303120331321-3000011102003310-2010112020201032-1323211102311201-1312002103211202"></a>

## service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site — virtual_site / 111020212002 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-005.md#canonical-2232032323312210-1320210021123102-2323123132332020-2021233311111231-1003000211331102-3032322003100330-1011330303201322-1303131022011320)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-0220200300312030-2300020331323223-0323220131203321-2301232102111212-1030223032333130-0110323110030332-1301320131101212-0300131123112203"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-0122013321221023-1200100002233303-1020202200300203-2000320311331100-2103322331311300-3100332132203323-2133222122222120-0003301231102132"></a>

## Direct properties — virtual_site / 111020212002 / 3

<a id="canonical-3123233221011222-2013202033122313-3112120323131101-1032331322323213-1303222333323300-0223300032132320-0102000122122221-3231022013023033"></a>

<a id="canonical-2032002122320320-2323312022103101-3032312330302120-0032301020013303-2203002303120232-3331133132313003-0200323220221330-3330002021103203"></a>

## name property — virtual_site / 111020212002 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2210010032120200-1011200333101121-2011332222002213-2233201312003030-2001013110200113-3223213010211121-3200113101132333-3331103133132221"></a>

<a id="canonical-1302101302210031-3132033212120000-3313213203030010-2211030121121132-2023033322322323-2330031031131121-2233032123202331-0131123113213213"></a>

## namespace property — virtual_site / 111020212002 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-1032221120333103-2100203110313020-1211323130131012-0022130132121032-0121132322200322-1120320303123231-0332211333132220-1111123220330231"></a>

<a id="canonical-1133221200212233-2000023203123001-1113212132120312-2213322311111303-2312302323001002-1002130220110122-1110202122313020-1312120310210323"></a>

## tenant property — virtual_site / 111020212002 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1102202010033023-3312313132320213-2301322001032333-1123302200133200-0212032322133231-2030201330223112-2321323201220210-1220300110201231"></a>

## Next pages — virtual_site / 111020212002 / 7

- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-005.md#canonical-2232032323312210-1320210021123102-2323123132332020-2021233311111231-1003000211331102-3032322003100330-1011330303201322-1303131022011320)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300232133201010-2300013201021232-2103102132033021-0010013023101323-2122021302331223-2133322300201023-3313200211133001-0220032111130200"></a>

## service.advertise_options.advertise_custom.ports — ports / 021221312002 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- service.advertise_options.advertise_custom.ports

<a id="canonical-0212113103103321-1121323333310332-3033111213033011-3013131003212010-3030102232132112-1133131120031010-2013022103223200-1211210113123023"></a>

Type: `"list"`. Computed.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-0211030221203223-2323102131331113-1121032131202311-1232311030010220-1102212101001330-2101313110003312-1130101300321030-3312200232330112"></a>

## Direct properties — ports / 021221312002 / 3

- [http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130): complete subsection reference.

- [port](data-sources--workload--reference--group-008.md#canonical-1330120103311032-3323002022332313-1230010032111002-0301100332330310-1013300322113000-3020110122120300-1230010001312321-1311000201122222): complete subsection reference.

- [tcp_loadbalancer](data-sources--workload--reference--group-008.md#canonical-1213013231333121-3300100100220230-1131111132003220-3022122220020203-0200132131130103-3220333332311300-2220212322331102-1331120231121131): complete subsection reference.

<a id="canonical-3202331230322330-3031013100200233-0233221000010103-2022023232100221-3333103032002203-3001231230023311-3331110022011203-0020020311130203"></a>

## Next pages — ports / 021221312002 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.port](data-sources--workload--reference--group-008.md#canonical-1330120103311032-3323002022332313-1230010032111002-0301100332330310-1013300322113000-3020110122120300-1230010001312321-1311000201122222)
- [service.advertise_options.advertise_custom.ports.tcp_loadbalancer](data-sources--workload--reference--group-008.md#canonical-1213013231333121-3300100100220230-1131111132003220-3022122220020203-0200132131130103-3220333332311300-2220212322331102-1331120231121131)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100211332010112-0302301232013123-1010312221231113-3130012331022122-1102122103033031-1220320120032221-3122221120032302-2320121312213110"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer — http_loadbalancer / 212003100320 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer

<a id="canonical-2012020330331301-1222323202201221-3101202222211311-0222003102033221-2311022312023002-3223132103021120-2012233201113110-2203311322003223"></a>

Type: `"single"`. Computed.

Configuration parameter for http loadbalancer.

Upstream description:

HTTP/HTTPS Load balancer.

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

<a id="canonical-0113223322032323-1000222202332312-3213300030201021-0300123132010301-1332021003212203-0100332021011200-3222112330111013-0133233132001022"></a>

## Direct properties — http_loadbalancer / 212003100320 / 3

- [default_route](data-sources--workload--reference--group-005.md#canonical-1100310320110100-2302121112222022-1100020033110320-1212331103112120-2131122332132031-3322130323320030-0332231122010303-2011001301220212): complete subsection reference.

<a id="canonical-0330230203303120-2010000210323102-2321322011203331-2013032103203331-3203130322113001-3211023332322130-2302021313200103-1330020100111210"></a>

<a id="canonical-2231212113311333-3301001032322131-1130311211012033-1322202233122223-1311332200232230-2203010032321011-0101323110222100-2010303133110203"></a>

## domains property — http_loadbalancer / 212003100320 / 4

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [http](data-sources--workload--reference--group-005.md#canonical-2131032112032331-1212000113313202-0122031300220010-1121231103011220-1001301011001131-0300103122221021-1102320032333021-1013300012132012): complete subsection reference.

- [https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101): complete subsection reference.

- [https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330): complete subsection reference.

- [specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001): complete subsection reference.

<a id="canonical-1032201021330320-1303012102021113-3032221032032201-3111313300201311-2313323200212201-1221312002233311-1301130033111220-1311333233110111"></a>

## Next pages — http_loadbalancer / 212003100320 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-005.md#canonical-1100310320110100-2302121112222022-1100020033110320-1212331103112120-2131122332132031-3322130323320030-0332231122010303-2011001301220212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.http](data-sources--workload--reference--group-005.md#canonical-2131032112032331-1212000113313202-0122031300220010-1121231103011220-1001301011001131-0300103122221021-1102320032333021-1013300012132012)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1100310320110100-2302121112222022-1100020033110320-1212331103112120-2131122332132031-3322130323320030-0332231122010303-2011001301220212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101022032212313-2120033201000213-0033211331331030-3131120211023132-1220101020323131-2030003333132111-0022103301023012-2222110320200020"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route — default_route / 130221121132 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route

<a id="canonical-3123312200123212-1313111103301001-1011313112112031-3020320211300323-0201133003201202-0312021132220200-3000220221322201-0213222120303023"></a>

Type: `"single"`. Computed.

Configuration parameter for default route.

Upstream description:

Default route matching all APIs.

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

<a id="canonical-1202023310232133-2122220212132322-0223110232232130-2023210012020013-2311322301221130-2201301013011013-0011202131022002-3222012211221311"></a>

## Direct properties — default_route / 130221121132 / 3

- [auto_host_rewrite](data-sources--workload--reference--group-005.md#canonical-3221011020121112-1110331010131210-3320120321132121-2320321032003333-1122102021220030-2010003010000332-2212031113011212-3202321130031232): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-005.md#canonical-3032223222131010-0200213023002312-0303120021110312-3211211210030212-1212003110210123-2101313121003033-1220221301102323-2101031202221311): complete subsection reference.

<a id="canonical-3323131301333122-3302003312120332-2112020311231312-0311123332333321-3331312233320212-1033332123102311-2132322033210202-1322101302031200"></a>

<a id="canonical-2110220211033221-3301022011121032-3233212330203232-0032100103133320-1201020003121100-2313123213223313-0100300100201330-2301111322301013"></a>

## host_rewrite property — default_route / 130221121132 / 4

Type: `"string"`. Computed.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-2110202231022201-2213313100333003-0103103323012003-0120131230132012-1302103013302310-0011301102011021-0031100322233212-2013301030330021"></a>

## Next pages — default_route / 130221121132 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite](data-sources--workload--reference--group-005.md#canonical-3221011020121112-1110331010131210-3320120321132121-2320321032003333-1122102021220030-2010003010000332-2212031113011212-3202321130031232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite](data-sources--workload--reference--group-005.md#canonical-3032223222131010-0200213023002312-0303120021110312-3211211210030212-1212003110210123-2101313121003033-1220221301102323-2101031202221311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3221011020121112-1110331010131210-3320120321132121-2320321032003333-1122102021220030-2010003010000332-2212031113011212-3202321130031232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100302013032111-0111310110232101-2322023111231022-1133211312320021-1120132110133031-0302230003330003-0333230102201300-0012321133302301"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite — auto_host_rewrite / 321022312313 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-005.md#canonical-1100310320110100-2302121112222022-1100020033110320-1212331103112120-2131122332132031-3322130323320030-0332231122010303-2011001301220212)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-0022022210133100-3133010333221232-2011113123031020-3213300210200001-3330012113231132-1130022103303200-0002213100100311-3023223321013113"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1320302002233230-0021003120323322-1200322323130303-3011331200203130-3323210232100131-0303232331122331-3110222020203210-0131302310322012"></a>

## Direct properties — auto_host_rewrite / 321022312313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203212331232211-3300201203000202-3312022231212222-1213232231220303-0300120223000113-2122201033331021-1103033331031320-2321022210323331"></a>

## Next pages — auto_host_rewrite / 321022312313 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-005.md#canonical-1100310320110100-2302121112222022-1100020033110320-1212331103112120-2131122332132031-3322130323320030-0332231122010303-2011001301220212)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3032223222131010-0200213023002312-0303120021110312-3211211210030212-1212003110210123-2101313121003033-1220221301102323-2101031202221311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022333313213301-0100111203320300-2003321312032110-1121331032023000-2220330112132003-1123122112312323-2003033320303232-2020123122013130"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite — disable_host_rewrite / 032113030121 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-005.md#canonical-1100310320110100-2302121112222022-1100020033110320-1212331103112120-2131122332132031-3322130323320030-0332231122010303-2011001301220212)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-1001102332122222-1113301310221211-3300332133302102-0130031231030312-1013101131202200-1031331033001130-3210010112303030-2113310121220320"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2222002111030121-3201100001130333-1023231033130213-2103122101302303-2200101122000231-3031232211132020-0130103112221210-2331011131112320"></a>

## Direct properties — disable_host_rewrite / 032113030121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312233333321000-0012021233333013-0301121220233133-0123223103103023-2010313313022121-3231203123321122-0003333021132122-3110121100333233"></a>

## Next pages — disable_host_rewrite / 032113030121 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-005.md#canonical-1100310320110100-2302121112222022-1100020033110320-1212331103112120-2131122332132031-3322130323320030-0332231122010303-2011001301220212)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2131032112032331-1212000113313202-0122031300220010-1121231103011220-1001301011001131-0300103122221021-1102320032333021-1013300012132012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213030233300323-3122231110001322-3023320331132113-1332100320323102-1233111322200200-2102312311101131-0301022020331313-0011122212133102"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.http — http / 030310101231 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.http

<a id="canonical-0313202020013301-1312220312030313-2221222202032333-1232222032212000-1121202200300031-3303320001022120-0123202022011323-2223210333333030"></a>

Type: `"single"`. Computed.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

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

<a id="canonical-2012332031022212-2100113330031211-2323111202201320-1102323031131322-0221232201211221-1301013003130220-0031001031011011-1022001300220112"></a>

## Direct properties — http / 030310101231 / 3

<a id="canonical-1020303001133103-2130002101230222-1231103220010102-3000311100302210-3233210033031012-0321131300133132-3313301013123312-1002221011322333"></a>

<a id="canonical-3101123332212221-1131220113031012-1231331121321011-3020322312220000-2033001121111201-2030013022102000-1213230313203113-2121000112200023"></a>

## dns_volterra_managed property — http / 030310101231 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-2223210112121221-1232132332012010-3110230102212331-0320012223323330-3003102111121233-0120003322321222-2110200002302300-3030231122300022"></a>

<a id="canonical-3123001232030311-2232201220323123-3033010200220033-3030212113003303-0003122202232121-0333301220202111-3103330230313201-0223232230230213"></a>

## port property — http / 030310101231 / 5

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3200033013023300-0103033221323033-3230003011303231-0003110021110221-3002010222320231-1011330231012022-3123130323231003-2312022222003101"></a>

<a id="canonical-2332212013020201-3130000303121000-3321133231211101-2111233331230213-2321113303220203-0321213031120110-1322332312020101-1211032103213330"></a>

## port_ranges property — http / 030310101231 / 6

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2020211102230012-0221102123211210-0103223100023102-2010221010032001-3321302311121211-0133321030211221-1211011231112202-0033021301300011"></a>

## Next pages — http / 030310101231 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332013303230021-2001132221120201-0313313330121232-1230311123221131-1230313200032102-1202013233201333-1312232211121033-1013111112320120"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https — https / 310001233000 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https

<a id="canonical-2130033003113233-1123101232220001-3222332321011122-0033110020332123-3000121020233023-1212121322221001-3130312030211211-0201232230133230"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

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

<a id="canonical-0200221020132110-2233022002331031-1230312111203011-3103320212330301-0310100211132312-2001333121021301-1212301121020112-1233002021302101"></a>

## Direct properties — https / 310001233000 / 3

<a id="canonical-0000313301013120-2302113222233021-1301302012012020-0201202203223310-3022013212012201-3022123031123133-3232231322223021-0132032333333210"></a>

<a id="canonical-2223302323333020-1133020112322330-2311301121032022-0221320132302310-2310312301032230-0023013221221003-2212330312310120-1110310230313230"></a>

## add_hsts property — https / 310001233000 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-0302322123113332-1233011201331112-0033020030322012-1301031001131132-2113013003131110-0111100331103211-2110031310333212-1232221323003312"></a>

<a id="canonical-1310213011103101-0330010323101133-1211101100133131-2223030110112203-1002101230313332-1113232110121310-2033310011132021-0020331100103021"></a>

## append_server_name property — https / 310001233000 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--workload--reference--group-005.md#canonical-0322223103020101-1000101310121210-3303203301102222-1032002302323231-0301313230112121-3320212103202212-0313213211232132-0212121132202131): complete subsection reference.

<a id="canonical-1023032101112001-1202102132003332-2101132203312312-3211032213131202-2220312200230023-0132033232033220-1021133200010330-2111020133211132"></a>

<a id="canonical-1200223220310303-3201122213221120-1000313230310322-1033301323220132-3013313000101230-3033301102120232-0310300203231211-0320131321330022"></a>

## connection_idle_timeout property — https / 310001233000 / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](data-sources--workload--reference--group-005.md#canonical-0303010113011030-1100101132323222-3110330220230311-1123113022321111-1031023103231312-3231013311313033-1033312311130023-2002010001330011): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-005.md#canonical-2303310130021022-3212020020113312-2221020222202212-0302322200233003-3011213121102333-1313223210230303-2122330312210111-1310123122232023): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-005.md#canonical-1323211121300110-0303211200202133-3221111333202323-3330301123321313-0302000301320102-3020200312111321-2033132321301303-2231213223311302): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-005.md#canonical-1212333322101032-3003211111321311-2300101113120201-0003113002210232-1022311132323312-0000021022012002-2210001213321301-0001202112022012): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-005.md#canonical-3221032210102010-2232003333032200-2020323221120130-3211032312320212-2312333210030113-0332231132032313-1231000003030300-3232321113132302): complete subsection reference.

<a id="canonical-1130023320031011-3033313233321333-0133221023030031-0321031132102303-3331332310320223-3310320221132123-0000002320231203-1122310203212300"></a>

<a id="canonical-3221321031030132-3110233213003332-1133023110211001-3322322131230212-2230212211121021-0313222303330131-0322031203331133-2303030310100112"></a>

## http_redirect property — https / 310001233000 / 7

Type: `"bool"`. Computed.

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

- [non_default_loadbalancer](data-sources--workload--reference--group-005.md#canonical-3311133312013003-2113232313000302-3203223023203103-3233301300210333-0330330332212113-1222232322102122-3230032010233213-3202101013120022): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-005.md#canonical-3200122102123201-1203030300113003-0131010322303212-2033330232203020-3032233110122030-1231103333000002-0210201101023223-1101023203033033): complete subsection reference.

<a id="canonical-2132233121131023-2332013232132133-3030221002220010-3031030020323200-1010330020130001-0300331030310203-0211322313210130-2333110012310333"></a>

<a id="canonical-3210130132213000-0012331122110023-1202222313131022-0131111102100010-0310312332313112-2123222030133232-2103220312232301-0132303022100112"></a>

## port property — https / 310001233000 / 8

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2021223201333322-2112132321230122-2333331300233302-3323202222201121-2003231030100232-3213122220120021-0023312003002121-0210031011321331"></a>

<a id="canonical-1003000231132221-2210233332331212-3113111013211310-3302221221030203-2230232010303011-1203202000301013-3003133212010010-1113313231210221"></a>

## port_ranges property — https / 310001233000 / 9

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3113100233332202-2313213323322130-2303023002010023-2120023213202221-3011303110101212-3202111100121231-2321001330232303-1222123102300201"></a>

<a id="canonical-2311310303301022-3330101022022333-3202300301311301-2132313032121003-2320310011013023-0123011203113011-1001100320030301-0210211120021223"></a>

## server_name property — https / 310001233000 / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033): complete subsection reference.

- [tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232): complete subsection reference.

<a id="canonical-1100333323212101-2112110111231023-0102202212202303-2230023020231313-2202130310030201-2022300100132032-3132100200010133-2233300013232321"></a>

## Next pages — https / 310001233000 / 11

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-005.md#canonical-0322223103020101-1000101310121210-3303203301102222-1032002302323231-0301313230112121-3320212103202212-0313213211232132-0212121132202131)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header](data-sources--workload--reference--group-005.md#canonical-0303010113011030-1100101132323222-3110330220230311-1123113022321111-1031023103231312-3231013311313033-1033312311130023-2002010001330011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer](data-sources--workload--reference--group-005.md#canonical-2303310130021022-3212020020113312-2221020222202212-0302322200233003-3011213121102333-1313223210230303-2122330312210111-1310123122232023)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize](data-sources--workload--reference--group-005.md#canonical-1323211121300110-0303211200202133-3221111333202323-3330301123321313-0302000301320102-3020200312111321-2033132321301303-2231213223311302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize](data-sources--workload--reference--group-005.md#canonical-1212333322101032-3003211111321311-2300101113120201-0003113002210232-1022311132323312-0000021022012002-2210001213321301-0001202112022012)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-3221032210102010-2232003333032200-2020323221120130-3211032312320212-2312333210030113-0332231132032313-1231000003030300-3232321113132302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer](data-sources--workload--reference--group-005.md#canonical-3311133312013003-2113232313000302-3203223023203103-3233301300210333-0330330332212113-1222232322102122-3230032010233213-3202101013120022)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through](data-sources--workload--reference--group-005.md#canonical-3200122102123201-1203030300113003-0131010322303212-2033330232203020-3032233110122030-1231103333000002-0210201101023223-1101023203033033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0322223103020101-1000101310121210-3303203301102222-1032002302323231-0301313230112121-3320212103202212-0313213211232132-0212121132202131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311102331300032-0301131311011103-0010020333110210-3113232000233312-2312232011022321-2010100031121011-2231300213333311-2000330232200311"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options — coalescing_options / 333110322333 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-2021113110310112-2021000013231222-3030232213123321-2301202331211302-3011231121300101-0213013130101123-0030320312000123-1010120333001011"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

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

<a id="canonical-0031200212010001-0031223303030032-0303011010121113-0031333002123212-1223321121003321-0133002200203030-3202301023320323-3123002002232021"></a>

## Direct properties — coalescing_options / 333110322333 / 3

- [default_coalescing](data-sources--workload--reference--group-005.md#canonical-1030112112132330-0022031133202201-0331013233101010-1103122013302332-0213103333222321-0023210100303000-1123302100302321-0122330203212220): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-005.md#canonical-1103133330223201-2123101103303313-0312213000012203-2330302331001123-1032231101201221-3322011000030201-1130120003133312-1332120010002311): complete subsection reference.

<a id="canonical-1120123133031331-1133013131032333-3030210323110001-2010113023001001-1100222222213303-0113031322332203-0313211323033221-0010102030300031"></a>

## Next pages — coalescing_options / 333110322333 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing](data-sources--workload--reference--group-005.md#canonical-1030112112132330-0022031133202201-0331013233101010-1103122013302332-0213103333222321-0023210100303000-1123302100302321-0122330203212220)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing](data-sources--workload--reference--group-005.md#canonical-1103133330223201-2123101103303313-0312213000012203-2330302331001123-1032231101201221-3322011000030201-1130120003133312-1332120010002311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1030112112132330-0022031133202201-0331013233101010-1103122013302332-0213103333222321-0023210100303000-1123302100302321-0122330203212220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111003002311101-0120323103303013-1322220111333102-3031011233212101-0000333322202100-1110221213321232-2112120210301111-0202012312101200"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing — default_coalescing / 330331200112 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-005.md#canonical-0322223103020101-1000101310121210-3303203301102222-1032002302323231-0301313230112121-3320212103202212-0313213211232132-0212121132202131)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-3303032122030013-0222122313110110-2310330012222021-3322201002321313-3211310220033121-3233033310012002-0202333333222021-0111013302010031"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3301122213022323-1010013232021301-3230000011031200-2013121131310320-2132003203111023-3032033132030201-3113202102020110-0103001033100321"></a>

## Direct properties — default_coalescing / 330331200112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023033330221030-1133211332032312-0121200130022300-0301101130212120-0230202022320330-0332023333010133-3201200311232100-3303003123000113"></a>

## Next pages — default_coalescing / 330331200112 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-005.md#canonical-0322223103020101-1000101310121210-3303203301102222-1032002302323231-0301313230112121-3320212103202212-0313213211232132-0212121132202131)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1103133330223201-2123101103303313-0312213000012203-2330302331001123-1032231101201221-3322011000030201-1130120003133312-1332120010002311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133210331020130-0132213232113321-2223012323330321-1203012130010000-3132330000312000-1313203323333230-2133223011113123-1111223013022333"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing — strict_coalescing / 332032322111 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-005.md#canonical-0322223103020101-1000101310121210-3303203301102222-1032002302323231-0301313230112121-3320212103202212-0313213211232132-0212121132202131)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-1113122031212331-2102020300320023-0122032000130222-3310211200022200-0111323333232011-2123200302132103-2302131110230203-2102332020111220"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2223122233022331-1222220011231130-3300121103132310-0110023300221113-3102021031200200-0030330010333330-1231203313321200-2000232030021231"></a>

## Direct properties — strict_coalescing / 332032322111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031330230201330-2003101032231231-3011033301211230-3120300320312333-0021322322133302-3312123212330222-2221323200120033-3230032132031113"></a>

## Next pages — strict_coalescing / 332032322111 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-005.md#canonical-0322223103020101-1000101310121210-3303203301102222-1032002302323231-0301313230112121-3320212103202212-0313213211232132-0212121132202131)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0303010113011030-1100101132323222-3110330220230311-1123113022321111-1031023103231312-3231013311313033-1033312311130023-2002010001330011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021101132130023-1000221221213131-0322031203231333-2122002112021132-1010032311000101-2033120202320332-2111130202100022-3300310230201300"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header — default_header / 020012323013 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header

<a id="canonical-3232212231132231-0331312312030222-3300320031131211-0301213200021110-1120210311201330-3333110002021312-1002223131322101-3220031012133213"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2102321331003201-2201221032220122-3010312201311133-2030011020122212-0210023200321123-3322102320012113-0310031301101002-0220212322113020"></a>

## Direct properties — default_header / 020012323013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000130101220102-0233213211333131-3112331222300223-0013101122001303-2230132003321020-3012310131110221-2121130011032001-2313023212323321"></a>

## Next pages — default_header / 020012323013 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2303310130021022-3212020020113312-2221020222202212-0302322200233003-3011213121102333-1313223210230303-2122330312210111-1310123122232023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133313023300303-1122101210121033-0303003230213020-0221122010331100-3321211021203302-2112012220300000-2302132002311300-3231020303330100"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer — default_loadbalancer / 320330300231 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-0200213302313320-2333120130311102-1311002313113121-0002032232121210-2223212333120233-0230010321232213-3133123030310330-3202331130233120"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3301101231313030-2021001103201323-0220203322111103-2201111311321121-1111220010110021-2000310100311332-2320103332232203-0303133331113332"></a>

## Direct properties — default_loadbalancer / 320330300231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002323132310302-0033121331133002-2333131201120220-0103222302230312-1103102232222022-2002000033010210-2100313312110220-0103132031013230"></a>

## Next pages — default_loadbalancer / 320330300231 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1323211121300110-0303211200202133-3221111333202323-3330301123321313-0302000301320102-3020200312111321-2033132321301303-2231213223311302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130300133202121-1210330113102133-3102033130023313-0123221033021223-3221003020032322-3322333121211211-1002232311323020-1202002203201200"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize — disable_path_normalize / 230111221111 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-2012313023132133-3222011332203300-3310210311010103-2113301123330301-0000320111201031-3031030010010303-0023210231030013-1111223310221111"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3103121003103031-3213213320210333-3000122322310221-0021331211112210-0220011123013021-0101232230231203-2320212030303323-0032213122033131"></a>

## Direct properties — disable_path_normalize / 230111221111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102301311132031-3312002003311021-0111132310001121-0303200300003230-2200321330122022-2023121223013110-1202101112210331-0010133000132301"></a>

## Next pages — disable_path_normalize / 230111221111 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1212333322101032-3003211111321311-2300101113120201-0003113002210232-1022311132323312-0000021022012002-2210001213321301-0001202112022012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133111103132231-1212002201113310-0312330113010330-0101322003213231-2232212010300111-3320111233002332-1333123131112011-2220111303103231"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize — enable_path_normalize / 221220023032 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-0013112123321121-2101001112031101-2201031012331121-0213112120100021-1000200223300211-2211303212210230-2031332111233123-1213232131012202"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3112321002301223-0132020311121131-0011022322122013-3020303013101203-2323230132311203-0020323211203330-1020111013003022-3310032013220113"></a>

## Direct properties — enable_path_normalize / 221220023032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030112111202022-0003231102200213-3332103323001313-3331113310013313-3203210000220211-3120032033220132-1023013301333312-3100220021202013"></a>

## Next pages — enable_path_normalize / 221220023032 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3221032210102010-2232003333032200-2020323221120130-3211032312320212-2312333210030113-0332231132032313-1231000003030300-3232321113132302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133121031020130-0031022313020110-1201203332211002-3102303103001202-2013221031312103-0023120130012012-3100231100230113-3232330030122130"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options — http_protocol_options / 001312200023 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-2230231330012321-2200000131001320-2031011033200012-1110323030020112-2302131111112000-3103131010002122-3021001121123010-0231112111023033"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

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

<a id="canonical-1323332223101000-2130223200312122-2312120011103112-0011002113111203-3011022121020202-2121121010220231-1303000311020331-3122102332202230"></a>

## Direct properties — http_protocol_options / 001312200023 / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-005.md#canonical-3020103321201122-2221220202310013-1332130311322113-1211232202300213-3022112111301013-2330133020112233-3033301123122231-0231011003223320): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-005.md#canonical-1022032023123020-0002303110132000-1231213112033311-1311013230230112-3100011023123230-1322321210323013-0032133220310301-1130302211032230): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-005.md#canonical-0021323322032221-3130302332122312-2112102030111031-2111320300003003-0131212011001202-1213022213001301-3130032331331201-2320001310201112): complete subsection reference.

<a id="canonical-2032133032322103-2313321012132211-3232220212221321-2032021201212103-0333030100113230-1201300230330231-3222023003312231-1222022230110101"></a>

## Next pages — http_protocol_options / 001312200023 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-005.md#canonical-3020103321201122-2221220202310013-1332130311322113-1211232202300213-3022112111301013-2330133020112233-3033301123122231-0231011003223320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-005.md#canonical-1022032023123020-0002303110132000-1231213112033311-1311013230230112-3100011023123230-1322321210323013-0032133220310301-1130302211032230)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-005.md#canonical-0021323322032221-3130302332122312-2112102030111031-2111320300003003-0131212011001202-1213022213001301-3130032331331201-2320001310201112)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3020103321201122-2221220202310013-1332130311322113-1211232202300213-3022112111301013-2330133020112233-3033301123122231-0231011003223320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033012003221103-0110023313221133-2023332233311021-1230221111130300-3322220020300213-3002203111212333-2213220110002120-3001101010313221"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 021001010202 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-3221032210102010-2232003333032200-2020323221120130-3211032312320212-2312333210030113-0332231132032313-1231000003030300-3232321113132302)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-3302301303320300-2023002021222221-0322030123201132-0313313310020331-2120330231030130-2323333311021130-1222123223213013-0111212132303212"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0002333331232213-2122133012011022-0200121003202232-3212130311311003-1011200102220100-0112120003211312-0001020311020112-1110223130230120"></a>

## Direct properties — http_protocol_enable_v1_only / 021001010202 / 3

- [header_transformation](data-sources--workload--reference--group-005.md#canonical-1203333321030233-1003201021322211-2002122002110223-2230233130132232-3013300330300223-2200233331031213-0220300321000012-2113203320230133): complete subsection reference.

<a id="canonical-0033023202213023-3200023131211120-0100202032312313-1033203210002320-1033302010120331-1212110010201020-0001013210202011-0101003021223201"></a>

## Next pages — http_protocol_enable_v1_only / 021001010202 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-005.md#canonical-1203333321030233-1003201021322211-2002122002110223-2230233130132232-3013300330300223-2200233331031213-0220300321000012-2113203320230133)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-3221032210102010-2232003333032200-2020323221120130-3211032312320212-2312333210030113-0332231132032313-1231000003030300-3232321113132302)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1203333321030233-1003201021322211-2002122002110223-2230233130132232-3013300330300223-2200233331031213-0220300321000012-2113203320230133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010103320301103-3010130212120210-2222203121101000-3031301033131022-1203111313030120-3130002112310023-1300320311021030-3013111223002332"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 101213313213 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-3221032210102010-2232003333032200-2020323221120130-3211032312320212-2312333210030113-0332231132032313-1231000003030300-3232321113132302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-005.md#canonical-3020103321201122-2221220202310013-1332130311322113-1211232202300213-3022112111301013-2330133020112233-3033301123122231-0231011003223320)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0201323223121333-0302010230011312-2120130211330230-3203310302022333-1320231123123233-1101033230330030-2021301333032230-1132113313100001"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

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

<a id="canonical-0323131230201221-3010301321133113-3313300012133103-0323320133013112-3203010201021220-1132213232012211-3030323300003230-1101013331113103"></a>

## Direct properties — header_transformation / 101213313213 / 3

- [default_header_transformation](data-sources--workload--reference--group-005.md#canonical-1302022011223121-1102211000222332-1031330310001310-1020112323220313-2232030121002210-2010012002031133-2221303323210032-3332130111012032): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-005.md#canonical-1332021033202202-3331310032112222-3113110003022011-1202021212013033-1120321310000313-1331210310003330-3022232101022301-2131130310013003): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-005.md#canonical-2111203000213202-3222222302302300-0211302300211132-1032202032023121-2120300130323301-3320203123130021-1022222311030200-0121101103203033): complete subsection reference.

<a id="canonical-2232202222103033-1332231201010211-3121030230233332-0212313223331121-3230330011320203-2033011120132311-3120320320202112-0010322012110210"></a>

## Next pages — header_transformation / 101213313213 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--reference--group-005.md#canonical-1302022011223121-1102211000222332-1031330310001310-1020112323220313-2232030121002210-2010012002031133-2221303323210032-3332130111012032)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--reference--group-005.md#canonical-1332021033202202-3331310032112222-3113110003022011-1202021212013033-1120321310000313-1331210310003330-3022232101022301-2131130310013003)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--reference--group-005.md#canonical-2111203000213202-3222222302302300-0211302300211132-1032202032023121-2120300130323301-3320203123130021-1022222311030200-0121101103203033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-005.md#canonical-3020103321201122-2221220202310013-1332130311322113-1211232202300213-3022112111301013-2330133020112233-3033301123122231-0231011003223320)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1302022011223121-1102211000222332-1031330310001310-1020112323220313-2232030121002210-2010012002031133-2221303323210032-3332130111012032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102202301222101-1213301203330200-3002220201230333-0200023012001222-3031212220322013-0201012133000103-3310030331331100-3303022311232111"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 011101230221 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-3221032210102010-2232003333032200-2020323221120130-3211032312320212-2312333210030113-0332231132032313-1231000003030300-3232321113132302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-005.md#canonical-3020103321201122-2221220202310013-1332130311322113-1211232202300213-3022112111301013-2330133020112233-3033301123122231-0231011003223320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-005.md#canonical-1203333321030233-1003201021322211-2002122002110223-2230233130132232-3013300330300223-2200233331031213-0220300321000012-2113203320230133)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-2221330011200100-2022030233322330-1333331330303221-0332331320123303-0120333122233230-3313330322310032-0002322133023321-3011012122323233"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3002032203210133-3001320030022101-3102321230121312-0100202000311333-2011313300333311-0103201101110033-1200030302303012-0002102022212021"></a>

## Direct properties — default_header_transformation / 011101230221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212010303222312-1331101223123203-0000213101102011-1212323003000322-2003302001033202-2013100103200131-0212220302333113-3311332221310222"></a>

## Next pages — default_header_transformation / 011101230221 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-005.md#canonical-1203333321030233-1003201021322211-2002122002110223-2230233130132232-3013300330300223-2200233331031213-0220300321000012-2113203320230133)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1332021033202202-3331310032112222-3113110003022011-1202021212013033-1120321310000313-1331210310003330-3022232101022301-2131130310013003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303001220102333-1102212200331122-1332330221022021-3201021211133032-2210001031022113-1031030131211012-1202330111333121-2033101211131003"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 130222101022 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-3221032210102010-2232003333032200-2020323221120130-3211032312320212-2312333210030113-0332231132032313-1231000003030300-3232321113132302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-005.md#canonical-3020103321201122-2221220202310013-1332130311322113-1211232202300213-3022112111301013-2330133020112233-3033301123122231-0231011003223320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-005.md#canonical-1203333321030233-1003201021322211-2002122002110223-2230233130132232-3013300330300223-2200233331031213-0220300321000012-2113203320230133)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-0133030011332311-2332133302003232-0203332212131310-3003113232001210-3020012311311001-2112122202032131-1013111210123101-3201321121211132"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0221012310132231-0113022330123320-3222133033330222-1223003113111030-0320103232102302-2122030121113100-0021120102031213-1203322231332110"></a>

## Direct properties — preserve_case_header_transformation / 130222101022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311331001021323-2012213222300310-1223322032110023-1031203332212321-1121200101010111-3021203302131312-3223021022032100-0311103022222233"></a>

## Next pages — preserve_case_header_transformation / 130222101022 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-005.md#canonical-1203333321030233-1003201021322211-2002122002110223-2230233130132232-3013300330300223-2200233331031213-0220300321000012-2113203320230133)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2111203000213202-3222222302302300-0211302300211132-1032202032023121-2120300130323301-3320203123130021-1022222311030200-0121101103203033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011313002332111-0203322133323233-0013023201021012-1312000301133011-1200002023111331-3103030320202210-1301103301211122-2222332233212103"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 202023003213 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-3221032210102010-2232003333032200-2020323221120130-3211032312320212-2312333210030113-0332231132032313-1231000003030300-3232321113132302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-005.md#canonical-3020103321201122-2221220202310013-1332130311322113-1211232202300213-3022112111301013-2330133020112233-3033301123122231-0231011003223320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-005.md#canonical-1203333321030233-1003201021322211-2002122002110223-2230233130132232-3013300330300223-2200233331031213-0220300321000012-2113203320230133)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-0220203222122321-3321020302203213-2330131013130123-0300021000332321-3212120312121112-3030132333100020-2230211232220120-0213330313121131"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3121131122112112-2211010213320022-0130101033333003-1123203200300011-3120131102330033-2200320133103112-2122211020002011-1231330221300230"></a>

## Direct properties — proper_case_header_transformation / 202023003213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020322013202123-3311033001112130-2221011000112103-0103212203223211-2000022031122003-3032021110110110-0331131032303112-2212103132212302"></a>

## Next pages — proper_case_header_transformation / 202023003213 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-005.md#canonical-1203333321030233-1003201021322211-2002122002110223-2230233130132232-3013300330300223-2200233331031213-0220300321000012-2113203320230133)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1022032023123020-0002303110132000-1231213112033311-1311013230230112-3100011023123230-1322321210323013-0032133220310301-1130302211032230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330220311322201-0013011233113103-0012112032002102-0303011211330102-1212313302311012-0130330330302123-3200022002232312-2213311311131322"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 030102302033 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-3221032210102010-2232003333032200-2020323221120130-3211032312320212-2312333210030113-0332231132032313-1231000003030300-3232321113132302)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2212221102121230-0301012311310332-0231102132231331-2201002331130212-1300220312201331-2100111112000333-3312021311303111-1023032021310330"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0122310010330112-3221322201121210-0030301130120312-2000323331320221-1323123313203012-2211122323013200-0210131021023011-1310000201312223"></a>

## Direct properties — http_protocol_enable_v1_v2 / 030102302033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130232013201332-3030100233100011-3000220221332020-1102113133320332-3130230330200221-3033112111222313-2312013211202322-0232020200120302"></a>

## Next pages — http_protocol_enable_v1_v2 / 030102302033 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-3221032210102010-2232003333032200-2020323221120130-3211032312320212-2312333210030113-0332231132032313-1231000003030300-3232321113132302)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0021323322032221-3130302332122312-2112102030111031-2111320300003003-0131212011001202-1213022213001301-3130032331331201-2320001310201112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302231203120032-3003210230002313-2321133311022222-1220130011112133-1210320022202100-0313111130303322-1330213110213200-1323123203233113"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 011233112110 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-3221032210102010-2232003333032200-2020323221120130-3211032312320212-2312333210030113-0332231132032313-1231000003030300-3232321113132302)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-1323002301200330-2031331320131131-0102132230203201-3011310210203210-3101303212001012-1110022112332311-2101112132123312-1223221222200231"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-1003010102201123-1203010013212112-0000031322103200-2132313213233332-3232223013103100-1122112200130001-1102232233300310-0300020112232011"></a>

## Direct properties — http_protocol_enable_v2_only / 011233112110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201001310010232-1211031001231102-0333100020120021-3023200003302002-1233021031021001-2210311233131113-0030300001200320-3303003113201220"></a>

## Next pages — http_protocol_enable_v2_only / 011233112110 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-3221032210102010-2232003333032200-2020323221120130-3211032312320212-2312333210030113-0332231132032313-1231000003030300-3232321113132302)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3311133312013003-2113232313000302-3203223023203103-3233301300210333-0330330332212113-1222232322102122-3230032010233213-3202101013120022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202011313320301-1103130303123112-3303220002121010-0002111111201203-2033221322003111-3111223230332100-0301323331200003-3332132222300111"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer — non_default_loadbalancer / 122111311310 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-1313130112212323-3133332033101203-3320023030100221-1111132032000123-1201011003022002-0213000233313222-0130021020132001-2030313021233131"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

<a id="canonical-1130330233331113-1032133302312220-3321131210102303-1231011003111100-1131033323011302-0032311120020232-2331202230310103-2001322320103213"></a>

## Direct properties — non_default_loadbalancer / 122111311310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330120213130110-1010300110210011-0230320102003220-0302131222121332-2122233332320220-3301300210213301-2320322300230032-3130313202101223"></a>

## Next pages — non_default_loadbalancer / 122111311310 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3200122102123201-1203030300113003-0131010322303212-2033330232203020-3032233110122030-1231103333000002-0210201101023223-1101023203033033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233132102030320-3011331303113211-3132213210113210-2013113333101310-1121323012210212-2033330310121132-0022110212310020-3020301032131132"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through — pass_through / 310031311011 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through

<a id="canonical-0113330321121101-2312032033121231-2320302221322333-2203103222022023-3311032210213232-3203221203111012-0220032020203230-2210310112102131"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

<a id="canonical-1303333120001220-1323220231303113-0031001100312103-1130031112313232-0222032102013000-2323330330010203-3221333131303221-1230311223232123"></a>

## Direct properties — pass_through / 310031311011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221321002103123-1310103231001002-1312233303313231-0120031122010213-2210130113032233-3101100232013023-1322031032301330-0213013312021333"></a>

## Next pages — pass_through / 310031311011 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
