---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3201222030031023-0013133222021331-0121210112232103-3003022313223010-3023331320110200-1003221211311002-0220112010311110-0111303311113230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.persistent_volumes.persistent_volume.storage` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.persistent_volumes](data-sources--workload--reference--group-027.md#canonical-1132102002130121-0202013313231301-3123121303133333-3111103021121020-1113123203313322-3332020232010320-2113332300313132-2131322221012303)
- [stateful_service.persistent_volumes.persistent_volume](data-sources--workload--reference--group-027.md#canonical-0313133013332231-3300200130030103-2002302022200030-1003021003231002-1121111231212110-2232311322111203-3202032103100110-0130333321220032)
- stateful_service.persistent_volumes.persistent_volume.storage

<a id="canonical-2313301123333031-3001202212312111-1320201222130313-2110132300233202-3232022331321321-3133113033220233-2200022311330112-2332202102320213"></a>

Type: `"single"`. Computed.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

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

<a id="canonical-2131033112031200-3031022023002120-3330030132123120-2310313030321213-0321333231012120-1321222111233220-3122002202203220-2233023022112013"></a>

### Direct properties for `stateful_service.persistent_volumes.persistent_volume.storage`

<a id="canonical-0023020221003100-2133022301210020-2132013310202312-1102102003103302-0313233221011200-2010223023212213-3132023010322230-2111300000032212"></a>

#### `stateful_service.persistent_volumes.persistent_volume.storage.access_mode` property

Type: `"string"`. Computed.

\[Enum:
ACCESS\_MODE\_READ\_WRITE\_ONCE|ACCESS\_MODE\_READ\_WRITE\_MANY|ACCESS\_MODE\_READ\_ONLY\_MANY\]
Persistence storage access mode is used to configure access mode for persistent storage -
ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once Read Write Once is used to mount persistent storage
in read/write mode to exactly 1 host - ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many Read Write
Many is used.. Possible values are \`ACCESS\_MODE\_READ\_WRITE\_ONCE\`,
\`ACCESS\_MODE\_READ\_WRITE\_MANY\`, \`ACCESS\_MODE\_READ\_ONLY\_MANY\`. Defaults to
\`ACCESS\_MODE\_READ\_WRITE\_ONCE\`.

Additional upstream details:

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

<a id="canonical-3022230303223220-2203023100220013-2113033221330321-1032032202333221-3233222103032110-3001202211231033-0113301202220101-2212323100230300"></a>

<a id="canonical-3133322301202031-1133300330232013-0331212202320133-0330233321033201-1331321133221331-0000120303002130-2233200332222223-2200322102233122"></a>

#### `stateful_service.persistent_volumes.persistent_volume.storage.class_name` property

Type: `"string"`. Computed.

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

- [default](data-sources--workload--reference--group-028.md#canonical-2200011122232231-3230100011203113-2210300302011001-0333322222101332-1313121100223203-2223330031100313-3230233021200001-1222103002113330): complete subsection reference.

<a id="canonical-2110130122111222-1113233221131320-1333102031023230-1212223022331233-3103301331130010-1311202330022303-2033312022301220-3202233102311023"></a>

<a id="canonical-3023132331211000-2021322313001302-0002001120333320-2113120101022111-3212301132121222-3203233310022020-0312123023201010-1020000012020311"></a>

#### `stateful_service.persistent_volumes.persistent_volume.storage.storage_size` property

Type: `"number"`. Computed.

Size (in GiB). Size in GiB of the persistent storage.

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

<a id="canonical-2200011122232231-3230100011203113-2210300302011001-0333322222101332-1313121100223203-2223330031100313-3230233021200001-1222103002113330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.persistent_volumes.persistent_volume.storage.default` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.persistent_volumes](data-sources--workload--reference--group-027.md#canonical-1132102002130121-0202013313231301-3123121303133333-3111103021121020-1113123203313322-3332020232010320-2113332300313132-2131322221012303)
- [stateful_service.persistent_volumes.persistent_volume](data-sources--workload--reference--group-027.md#canonical-0313133013332231-3300200130030103-2002302022200030-1003021003231002-1121111231212110-2232311322111203-3202032103100110-0130333321220032)
- [stateful_service.persistent_volumes.persistent_volume.storage](data-sources--workload--reference--group-028.md#canonical-3201222030031023-0013133222021331-0121210112232103-3003022313223010-3023331320110200-1003221211311002-0220112010311110-0111303311113230)
- stateful_service.persistent_volumes.persistent_volume.storage.default

<a id="canonical-3200101110121110-3032210112133121-0110112102232203-1301123020220312-2021021133130321-0023320031312012-1310230202313023-1310100033101032"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130310200302331-3301113001002211-0231300130002032-0320200012113302-2010020220101022-0221323130323300-0010123130122332-1232023022230312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.scale_to_zero` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- stateful_service.scale_to_zero

<a id="canonical-3022212313232201-0221131102023211-2001332323332133-1222121120032012-2131333220313212-0012233220003333-2233033110323333-2110022220230113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for scale to zero.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011330200003232-3232210131033220-1331201211101011-0113112011320103-3212220101322221-0100131112031303-3331333222200110-3212321201321212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.volumes` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- stateful_service.volumes

<a id="canonical-3231102200003111-1331010312330113-0002220201023130-3112000211210301-3213023010223121-2322211222223200-0021102120100020-3000020110123210"></a>

Type: `"list"`. Computed.

Ephemeral Volumes. Ephemeral volumes for the service.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-3223212213222311-2001330311133310-1233022213233011-2033032303032023-2323132310003010-0031222000110203-0122002300011011-2321130020130112"></a>

### Direct properties for `stateful_service.volumes`

- [empty_dir](data-sources--workload--reference--group-028.md#canonical-0033202132313123-2023132130023310-3332020110210013-3302202202010300-1330032320303130-2033122121133233-3301101000011133-1202220231201113): complete subsection reference.

- [host_path](data-sources--workload--reference--group-028.md#canonical-1333110011012310-2301313013120021-0313322312222210-2211020110313201-1132023021111231-1032231101231302-0002011130231111-1111302030222112): complete subsection reference.

<a id="canonical-3213331220112113-1210322133131001-0022202230010100-1232112233202113-2211320311321201-1003002230333002-1021103033020133-3221030200103033"></a>

<a id="canonical-2120320330023321-1033311111021221-2230013033013312-3122120332033320-2032110231330323-3031010313313111-0322022121101203-3131200230122311"></a>

#### `stateful_service.volumes.name` property

Type: `"string"`. Computed.

Name. Name of the volume.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
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
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

<a id="canonical-0033202132313123-2023132130023310-3332020110210013-3302202202010300-1330032320303130-2033122121133233-3301101000011133-1202220231201113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.volumes.empty_dir` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.volumes](data-sources--workload--reference--group-028.md#canonical-2011330200003232-3232210131033220-1331201211101011-0113112011320103-3212220101322221-0100131112031303-3331333222200110-3212321201321212)
- stateful_service.volumes.empty_dir

<a id="canonical-1310022233201222-0113000110301312-3110030321233012-3012111212321300-0321133210330231-0002102023021000-2331122103013020-1022201103100202"></a>

Type: `"single"`. Computed.

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3211021313021113-0011031302302330-3231202233203121-0112121301020103-0221031223032330-2312021231321122-0030020130310130-3121333330301203"></a>

### Direct properties for `stateful_service.volumes.empty_dir`

- [mount](data-sources--workload--reference--group-028.md#canonical-3311122201122312-2121311112030101-3033332230211121-0222120031101121-0013010111200330-0113023001232331-3200311301311303-1221300023221130): complete subsection reference.

<a id="canonical-1100301033122222-1013210102303212-3132313323131130-3103032323233010-0012132121211302-1111321312033221-0312110101131201-2302103220300021"></a>

<a id="canonical-3200310231100212-1011221033223132-0130222001111313-0332031222210211-2003202303222100-2031112023303210-1112010302330030-1222121120202320"></a>

#### `stateful_service.volumes.empty_dir.size_limit` property

Type: `"number"`. Computed.

Size Limit (in GiB). Configuration parameter for size limit

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
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3311122201122312-2121311112030101-3033332230211121-0222120031101121-0013010111200330-0113023001232331-3200311301311303-1221300023221130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.volumes.empty_dir.mount` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.volumes](data-sources--workload--reference--group-028.md#canonical-2011330200003232-3232210131033220-1331201211101011-0113112011320103-3212220101322221-0100131112031303-3331333222200110-3212321201321212)
- [stateful_service.volumes.empty_dir](data-sources--workload--reference--group-028.md#canonical-0033202132313123-2023132130023310-3332020110210013-3302202202010300-1330032320303130-2033122121133233-3301101000011133-1202220231201113)
- stateful_service.volumes.empty_dir.mount

<a id="canonical-3200011033101221-3320233003331233-2133233113000011-2111012230312300-2022121300113333-0233211203231011-3310121121313112-3332102321103123"></a>

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

<a id="canonical-1123331203332302-2232221123133212-0031123211331122-0120020001102002-2131113001220112-0230112100200001-2013123130010331-3203210123122030"></a>

### Direct properties for `stateful_service.volumes.empty_dir.mount`

<a id="canonical-1232010313313220-3323133310322300-2001023031200021-3112201013001310-1321132322212130-2133131213021123-3000310202313223-1210203020033212"></a>

#### `stateful_service.volumes.empty_dir.mount.mode` property

Type: `"string"`. Computed.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

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

<a id="canonical-1230210110230330-3323003322033131-0012323121223031-1030022011321023-2311221201222120-0121032021223333-1033000130010332-1023020001320030"></a>

<a id="canonical-3221022201101212-2033020203021101-1212202103311002-2022211012321310-3231222212111311-3002313023110012-2001031020130223-0020001203021010"></a>

#### `stateful_service.volumes.empty_dir.mount.mount_path` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0033200013020232-0302100210003212-3031330022131130-3223301011030320-2000110231330223-1203220233203113-3131112031311230-0121100102002101"></a>

<a id="canonical-0100202200200201-0231133310323013-0122311303313012-1200300202310312-3022213021113012-1311133220032031-0231213101030001-0321003001112233"></a>

#### `stateful_service.volumes.empty_dir.mount.sub_path` property

Type: `"string"`. Computed.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Additional upstream details:

Defaults to "" (volume's root).

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

<a id="canonical-1333110011012310-2301313013120021-0313322312222210-2211020110313201-1132023021111231-1032231101231302-0002011130231111-1111302030222112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.volumes.host_path` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.volumes](data-sources--workload--reference--group-028.md#canonical-2011330200003232-3232210131033220-1331201211101011-0113112011320103-3212220101322221-0100131112031303-3331333222200110-3212321201321212)
- stateful_service.volumes.host_path

<a id="canonical-2222002112301222-0032110121200202-0222330222223311-3330323201230101-2321222023111012-0310111122103131-3031132311210022-3130010332013220"></a>

Type: `"single"`. Computed.

Volume containing a host mapped path into the workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0020113311302313-1002232013100331-2321020132233003-1013102220011210-0212311100232201-2101230031012333-3103310112103333-0311022301102313"></a>

### Direct properties for `stateful_service.volumes.host_path`

- [mount](data-sources--workload--reference--group-028.md#canonical-1030303000212200-0031122301101001-2020110013102222-0200030033322200-1112222320333002-1333133303302202-0301130001212300-2011301113100303): complete subsection reference.

<a id="canonical-2121132120230231-3033322122332210-3111133001122033-1103311331102202-2223312101203202-2231200202313122-3121113332013110-2033333212230203"></a>

<a id="canonical-3121132102212222-0133200032301223-2300322030111200-3021121313323102-3233111222023011-3012113123313121-1131213112002020-0322332112132120"></a>

#### `stateful_service.volumes.host_path.path` property

Type: `"string"`. Computed.

Path. Path of the directory on the host.

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

<a id="canonical-1030303000212200-0031122301101001-2020110013102222-0200030033322200-1112222320333002-1333133303302202-0301130001212300-2011301113100303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.volumes.host_path.mount` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.volumes](data-sources--workload--reference--group-028.md#canonical-2011330200003232-3232210131033220-1331201211101011-0113112011320103-3212220101322221-0100131112031303-3331333222200110-3212321201321212)
- [stateful_service.volumes.host_path](data-sources--workload--reference--group-028.md#canonical-1333110011012310-2301313013120021-0313322312222210-2211020110313201-1132023021111231-1032231101231302-0002011130231111-1111302030222112)
- stateful_service.volumes.host_path.mount

<a id="canonical-0000003232031323-0020323302311023-0111330233320231-3211003322113021-0232111302222103-1230230311123320-1320231002130112-0030203331311320"></a>

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

<a id="canonical-1301002023120301-0032303001231132-3121211330301230-0310100322022003-2231330331001200-0010232221011221-0220112103100032-1010032133020331"></a>

### Direct properties for `stateful_service.volumes.host_path.mount`

<a id="canonical-0022103131011233-2331223202001223-3122221203131222-3132001230322013-1301323010220233-0132113122311013-1010303013320012-1333201112222000"></a>

#### `stateful_service.volumes.host_path.mount.mode` property

Type: `"string"`. Computed.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

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

<a id="canonical-3300121211120212-3232020230300203-2311013231220321-2212320212013310-2003331032313233-1032003313312132-2111203322322110-2120111000113113"></a>

<a id="canonical-3301233203323331-2131032031023222-2201011200023212-0013310030321121-3123131320232010-3011113021120212-1333203320131333-1202123120130032"></a>

#### `stateful_service.volumes.host_path.mount.mount_path` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2110301232330233-1331221230132121-3320321301331302-3103200301012330-1103230033022212-3203331212010123-0012321122031301-0213321300221221"></a>

<a id="canonical-2312312110030211-0212031321120000-1133123022322030-1331320133201012-3121032210223312-3220111120230122-0210231223001210-3011212003212330"></a>

#### `stateful_service.volumes.host_path.mount.sub_path` property

Type: `"string"`. Computed.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Additional upstream details:

Defaults to "" (volume's root).

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
