---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2012110033111201-2323122231313121-3301031330103210-2312132322110233-3323213233311103-0320232202220320-3231001222133133-1333230032132310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- service.volumes

<a id="canonical-2230133101101332-1132230203103133-1000102132330203-1202133132202101-1022333130030303-0002103133013123-2132232023303202-0210311322130032"></a>

Type: `"list"`. Computed.

Volumes. Volumes for the service.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2031001320223330-2210232131333230-2322303312022303-1133103101122103-3030001030011320-2313230320332223-0132123002020030-3302031332130020"></a>

### Direct properties for `service.volumes`

- [empty_dir](data-sources--workload--reference--group-015.md#canonical-0002103122000011-0113223020303230-2332232221203300-3113120213201230-0003320012012003-1033230200333223-0300323232122131-3222113003320322): complete subsection reference.

- [host_path](data-sources--workload--reference--group-015.md#canonical-3220202111330130-0310032103002210-0122302221113022-3022122321110002-1211123003210300-3121213331203021-3230312212312232-2213002303130131): complete subsection reference.

<a id="canonical-1110332012211323-3302121312102133-3002321323300121-0031202313302311-0103022333023112-0232312002032233-1213032311120130-2100212002323332"></a>

<a id="canonical-3211130331223022-0010213033223222-1211331220231013-1130311302332123-0003221002030122-1210322310200120-0032231113011030-1220002333202113"></a>

#### `service.volumes.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [persistent_volume](data-sources--workload--reference--group-015.md#canonical-0323101102310122-2101332202201223-0202032030321130-1222010001103323-1032210101013230-2331003223212333-3103322030312233-3303311003100331): complete subsection reference.

<a id="canonical-0002103122000011-0113223020303230-2332232221203300-3113120213201230-0003320012012003-1033230200333223-0300323232122131-3222113003320322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.empty_dir` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-2012110033111201-2323122231313121-3301031330103210-2312132322110233-3323213233311103-0320232202220320-3231001222133133-1333230032132310)
- service.volumes.empty_dir

<a id="canonical-0313300222010300-2310222311000022-2103320233013002-3113131311211312-1323030111110130-1210320030023231-0322013012003222-3102110223032021"></a>

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

<a id="canonical-0230300232212303-2002120210212011-3032321122230210-1322031332010023-1302312020033002-0233211102113023-1210133233033003-0303123031122331"></a>

### Direct properties for `service.volumes.empty_dir`

- [mount](data-sources--workload--reference--group-015.md#canonical-1312223003211110-1312010301201323-2130111212021113-2203321203023132-1002112000110132-2103312221220023-0311300301013321-0130231320031031): complete subsection reference.

<a id="canonical-2221311232133031-1031231300302031-0300112002102323-1210101031313010-1120003122310131-1221120030200331-0131221210023002-0312213130102002"></a>

<a id="canonical-1332230220321303-1030212220310333-0030210232221323-2010120223121121-1023130001202102-1133022331123113-0010300110010210-3202222031203333"></a>

#### `service.volumes.empty_dir.size_limit` property

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

<a id="canonical-1312223003211110-1312010301201323-2130111212021113-2203321203023132-1002112000110132-2103312221220023-0311300301013321-0130231320031031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.empty_dir.mount` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-2012110033111201-2323122231313121-3301031330103210-2312132322110233-3323213233311103-0320232202220320-3231001222133133-1333230032132310)
- [service.volumes.empty_dir](data-sources--workload--reference--group-015.md#canonical-0002103122000011-0113223020303230-2332232221203300-3113120213201230-0003320012012003-1033230200333223-0300323232122131-3222113003320322)
- service.volumes.empty_dir.mount

<a id="canonical-3033221313222132-3300210210102030-3023122222201031-2021202300203121-2222330132303123-0323101201032001-0302203232200313-2201302133102120"></a>

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

<a id="canonical-1130010132312333-0300032202100100-0330212033331223-3012300200312303-2021013300032333-0010031200301203-0110002202021311-0222001230001110"></a>

### Direct properties for `service.volumes.empty_dir.mount`

<a id="canonical-0113202200320030-0312300323220012-0333333321300311-2122001002101031-3032302221021013-3001321302231121-3130022330313203-3302010103322123"></a>

#### `service.volumes.empty_dir.mount.mode` property

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

<a id="canonical-3210332231121001-1230312012333011-3121020000302102-3103110123100311-2000121301100203-1101302110202330-3002200013101133-0001203231001131"></a>

<a id="canonical-1033313321211203-0230331121001202-3123110233122302-0112000232110320-2010330122011031-3230302331133313-3101230121302000-2121101310201030"></a>

#### `service.volumes.empty_dir.mount.mount_path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3123121031232131-1113111001332321-1133212111202302-3110223211233311-0020033131303023-2210111003311101-1122032223302022-0310211332133002"></a>

<a id="canonical-1101023010112203-1021200110200211-2012123022211203-1100322223103011-3200011132221030-1330201303331003-1331330313031210-2031220030303023"></a>

#### `service.volumes.empty_dir.mount.sub_path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3220202111330130-0310032103002210-0122302221113022-3022122321110002-1211123003210300-3121213331203021-3230312212312232-2213002303130131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.host_path` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-2012110033111201-2323122231313121-3301031330103210-2312132322110233-3323213233311103-0320232202220320-3231001222133133-1333230032132310)
- service.volumes.host_path

<a id="canonical-2302321133301311-2020201200030030-0032112233011110-3010012311320032-3213120300000231-3302112212202301-3033031233220101-2313000130322131"></a>

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

<a id="canonical-2233112321111220-3122000130003123-0103100302113200-0311120233002000-2002102223121211-1313130320133003-0010030033302031-2222033011311300"></a>

### Direct properties for `service.volumes.host_path`

- [mount](data-sources--workload--reference--group-015.md#canonical-0330123032022222-2110122013333230-3312011213232122-2301301122310033-1310033333000122-1120200333202201-0311333311001032-3300022010012002): complete subsection reference.

<a id="canonical-2332201031121213-2021200333013100-1032302332313111-1101011020113321-3221002130203320-2120321321011312-1011331111102331-2113120323233303"></a>

<a id="canonical-1222321023113321-3112011122311112-1030232210033131-1312100230223302-3030033121320330-1201200031221130-1202130001031301-1331032022332313"></a>

#### `service.volumes.host_path.path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0330123032022222-2110122013333230-3312011213232122-2301301122310033-1310033333000122-1120200333202201-0311333311001032-3300022010012002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.host_path.mount` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-2012110033111201-2323122231313121-3301031330103210-2312132322110233-3323213233311103-0320232202220320-3231001222133133-1333230032132310)
- [service.volumes.host_path](data-sources--workload--reference--group-015.md#canonical-3220202111330130-0310032103002210-0122302221113022-3022122321110002-1211123003210300-3121213331203021-3230312212312232-2213002303130131)
- service.volumes.host_path.mount

<a id="canonical-0322032003033032-1102210030320003-3100313232233102-2303112313310032-1013303113022130-3021312221302130-2031301101323000-1311212222323122"></a>

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

<a id="canonical-0203111312232301-3020202010321303-3230330212313132-2203230331001132-1112030320310213-0111001020312320-3110003301212232-0133301223130211"></a>

### Direct properties for `service.volumes.host_path.mount`

<a id="canonical-3000201303233020-1110330102202003-1332131313031101-2302222300230023-3012013123200321-3111122333033301-1001210023102312-1331312012121232"></a>

#### `service.volumes.host_path.mount.mode` property

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

<a id="canonical-1332200311312003-3112332023111003-0313000202322103-1320122103033223-1103101223001020-0123233121130132-3321210230001020-2030210033303032"></a>

<a id="canonical-2211120320211113-0222032001032320-2221023233100203-0113220333003100-2112203022001131-2302322323320210-2233221112000011-2131323122330233"></a>

#### `service.volumes.host_path.mount.mount_path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1202321200103130-0223022120221303-1101313133201023-0331211211131331-3213323113031310-2030011131031330-3020223133031133-1222001123023310"></a>

<a id="canonical-0332101110202330-3123230303032322-2333300002301100-3222213022321131-1103310321213302-0331322030011103-2330121222322113-1323111302322020"></a>

#### `service.volumes.host_path.mount.sub_path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0323101102310122-2101332202201223-0202032030321130-1222010001103323-1032210101013230-2331003223212333-3103322030312233-3303311003100331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.persistent_volume` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-2012110033111201-2323122231313121-3301031330103210-2312132322110233-3323213233311103-0320232202220320-3231001222133133-1333230032132310)
- service.volumes.persistent_volume

<a id="canonical-1131221030003300-2132312121313223-0001210203112201-1311230312010320-1211233122320012-2330222302322113-2302322103301031-2213010220023131"></a>

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

<a id="canonical-2232231211030212-3013121103231133-3012223003001023-0033010220132210-1030010120232020-3330013213211103-0012030330013201-1210100203322312"></a>

### Direct properties for `service.volumes.persistent_volume`

- [mount](data-sources--workload--reference--group-015.md#canonical-3210121222001300-2301030103012321-3010222203301000-0321300310301302-0011222232002102-1231323230101110-1130120332331013-2122031003231222): complete subsection reference.

- [storage](data-sources--workload--reference--group-015.md#canonical-3010310231223102-0122212032212332-3211313303102133-3002130021020032-0202010012201213-1010223002202202-3321202212132223-3221301102030311): complete subsection reference.

<a id="canonical-3210121222001300-2301030103012321-3010222203301000-0321300310301302-0011222232002102-1231323230101110-1130120332331013-2122031003231222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.persistent_volume.mount` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-2012110033111201-2323122231313121-3301031330103210-2312132322110233-3323213233311103-0320232202220320-3231001222133133-1333230032132310)
- [service.volumes.persistent_volume](data-sources--workload--reference--group-015.md#canonical-0323101102310122-2101332202201223-0202032030321130-1222010001103323-1032210101013230-2331003223212333-3103322030312233-3303311003100331)
- service.volumes.persistent_volume.mount

<a id="canonical-1311021323321100-3010100211212103-1301302011323320-0223202103310333-3122332300021232-1223311003022321-2232010312112032-0232010233022233"></a>

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

<a id="canonical-1212122321313021-2023003212222020-0003302122212233-3013313030032033-3003210002300133-2200112110331200-3313032022233331-2132001113002020"></a>

### Direct properties for `service.volumes.persistent_volume.mount`

<a id="canonical-3010101321122023-0110222003320133-2000211131113113-0200200130330313-2203020231301021-2030100220303021-0033320300101011-1121300301120030"></a>

#### `service.volumes.persistent_volume.mount.mode` property

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

<a id="canonical-1231212111231323-2213311223320202-0300110201103221-3233311002131210-0001030102332321-2100033233331321-3012130332323131-1103101112101123"></a>

<a id="canonical-0222311123030003-2022010033000100-3303132222232122-2303313130212323-1330103021320131-2033333031321113-1010302021303133-0311112121220220"></a>

#### `service.volumes.persistent_volume.mount.mount_path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3023321333032220-2333013200202101-0313313120102310-2313201331200133-0302132020321021-0122101320310203-3101231110310000-2121302300010011"></a>

<a id="canonical-1222232212102221-0232211311023020-3100030122001133-3131322210323220-0232103101301302-2003310323330220-1210332000122233-0232133300122122"></a>

#### `service.volumes.persistent_volume.mount.sub_path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3010310231223102-0122212032212332-3211313303102133-3002130021020032-0202010012201213-1010223002202202-3321202212132223-3221301102030311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.persistent_volume.storage` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-2012110033111201-2323122231313121-3301031330103210-2312132322110233-3323213233311103-0320232202220320-3231001222133133-1333230032132310)
- [service.volumes.persistent_volume](data-sources--workload--reference--group-015.md#canonical-0323101102310122-2101332202201223-0202032030321130-1222010001103323-1032210101013230-2331003223212333-3103322030312233-3303311003100331)
- service.volumes.persistent_volume.storage

<a id="canonical-2001231103331012-0131220001000132-1210301033022301-2301220101232213-2032320020322230-3303323130200131-0120110032313111-3100032330010213"></a>

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

<a id="canonical-1200320103313031-3002023213321200-3120313210101302-3320013033133302-2021311111232102-2013201313013123-3003213230102333-0223110220021331"></a>

### Direct properties for `service.volumes.persistent_volume.storage`

<a id="canonical-0011032012301123-2121200230303131-3001233303100321-0331323212122033-0030210302010132-1031120010102020-0233232310101011-1111001111002303"></a>

#### `service.volumes.persistent_volume.storage.access_mode` property

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

<a id="canonical-1211001110203022-2301112012110102-0203001102320303-2231320220123323-3233202123220303-3001312300033213-2200310020021112-0233112002322023"></a>

<a id="canonical-0332221022222222-3311002201023302-3300313233313201-2131220112230013-2301103323001211-2331013101202221-1113003231313223-2100200130211311"></a>

#### `service.volumes.persistent_volume.storage.class_name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [default](data-sources--workload--reference--group-015.md#canonical-1131020312232331-0030132323002220-3333213200023011-2112213012003303-2113230201002021-1222022130023201-3100221133202020-1102310113332130): complete subsection reference.

<a id="canonical-2310300333101131-1222310211233132-0000220010023310-3121301200331320-0331333220120322-0113201000012323-2213123022303310-0021103313011023"></a>

<a id="canonical-3012133032131132-2213120230110111-1102221002200332-3130322123021231-3233021102032001-1022311012211313-2122202020303001-0232331002323031"></a>

#### `service.volumes.persistent_volume.storage.storage_size` property

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

<a id="canonical-1131020312232331-0030132323002220-3333213200023011-2112213012003303-2113230201002021-1222022130023201-3100221133202020-1102310113332130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.volumes.persistent_volume.storage.default` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-2012110033111201-2323122231313121-3301031330103210-2312132322110233-3323213233311103-0320232202220320-3231001222133133-1333230032132310)
- [service.volumes.persistent_volume](data-sources--workload--reference--group-015.md#canonical-0323101102310122-2101332202201223-0202032030321130-1222010001103323-1032210101013230-2331003223212333-3103322030312233-3303311003100331)
- [service.volumes.persistent_volume.storage](data-sources--workload--reference--group-015.md#canonical-3010310231223102-0122212032212332-3211313303102133-3002130021020032-0202010012201213-1010223002202202-3321202212132223-3221301102030311)
- service.volumes.persistent_volume.storage.default

<a id="canonical-3102033111013223-0210023303323110-1301113130302013-2202302003200201-3013300202311102-1132030223112303-3000131331322330-1310100002011011"></a>

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

<a id="canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- simple_service

<a id="canonical-3121301102123212-3203212301222222-0220221301202312-0130031331121220-3113312230331033-0213320230010022-0212211223031222-3200001111111110"></a>

Type: `"single"`. Computed.

SimpleService is a service having one container and one replica that is deployed on all Regional
Edges and advertised on internet via HTTP loadbalancer on default VIP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"do_not_advertise\",\"simple_advertise\"]",
  "x-ves-oneof-field-persistence_choice": "[\"disabled\",\"enabled\"]"
}
```

<a id="canonical-1301211200333100-0301201121312333-2110101003213132-3323230222020031-1013232230222010-3213321002031330-0003223122323330-1303300023012302"></a>

### Direct properties for `simple_service`

- [configuration](data-sources--workload--reference--group-015.md#canonical-2220321213211303-0020122331330313-0211202333012210-0020030122022330-1222021321103030-0003330110113021-1132101131233203-3010122300030122): complete subsection reference.

- [container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223): complete subsection reference.

- [disabled](data-sources--workload--reference--group-016.md#canonical-2112103133232330-2212313300022031-2132030012302003-0033223033113333-2311301323022003-0110122323101031-0101130200003103-2203110002233113): complete subsection reference.

- [do_not_advertise](data-sources--workload--reference--group-016.md#canonical-1210220313300213-3013331300310122-2232001332132003-2302230111202103-0230012033213001-3213221121020312-3303110313013123-1101012020030220): complete subsection reference.

- [enabled](data-sources--workload--reference--group-016.md#canonical-3001213300210321-2011120130303023-2211001300322013-1332022123332020-0232131121303302-0131000112110201-1133003122121203-3212202000010030): complete subsection reference.

<a id="canonical-2301033221233000-2130012113321013-3001023023131131-1202222102300100-3022320011301331-2200112130301032-2232131231231031-3122321320122013"></a>

<a id="canonical-1100210320301200-2013012300201231-2033112200323100-3121121103131100-1232122322303203-1210321323122220-2301121301203202-1221221200221000"></a>

#### `simple_service.scale_to_zero` property

Type: `"bool"`. Computed.

Scale down replicas of the service to zero.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [simple_advertise](data-sources--workload--reference--group-016.md#canonical-1332023020220230-0231312012221313-1132321323222002-2223303033022101-2220302230211001-0323332011222312-2312021312021221-1032021121111230): complete subsection reference.

<a id="canonical-2220321213211303-0020122331330313-0211202333012210-0020030122022330-1222021321103030-0003330110113021-1132101131233203-3010122300030122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.configuration` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- simple_service.configuration

<a id="canonical-0101333202310023-0110121132020333-3120203001333110-0311311020223133-3112332331232120-0020010232220331-3030132200021210-2021220223211011"></a>

Type: `"single"`. Computed.

Configuration parameters of the workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0323031120323321-3020002130231221-3013331302010121-1332303121033033-0023303102221120-3103032101220320-0012112102003123-3102230320301110"></a>

### Direct properties for `simple_service.configuration`

- [parameters](data-sources--workload--reference--group-015.md#canonical-1312302100222221-3310233130230312-3011122312201321-3020123032330102-1223232031312102-2111321113311321-1112011300203202-2330023033123211): complete subsection reference.

<a id="canonical-1312302100222221-3310233130230312-3011122312201321-3020123032330102-1223232031312102-2111321113311321-1112011300203202-2330023033123211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.configuration.parameters` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.configuration](data-sources--workload--reference--group-015.md#canonical-2220321213211303-0020122331330313-0211202333012210-0020030122022330-1222021321103030-0003330110113021-1132101131233203-3010122300030122)
- simple_service.configuration.parameters

<a id="canonical-1023323330020101-0203312032210003-0132002200110313-1123121203111223-2122003121231220-1111000010100102-0012120221220013-0013020230132213"></a>

Type: `"list"`. Computed.

Parameters. Parameters for the workload.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0100021233032013-0230111300111202-1032011001001031-2003310213112032-0310022320321211-1233122211012113-0022230120332112-3001012220302203"></a>

### Direct properties for `simple_service.configuration.parameters`

- [env_var](data-sources--workload--reference--group-015.md#canonical-0311103231201330-2203202132201132-2231032300112002-0131121232112210-3333000331023313-2022022013000310-0031332233300200-3121130302021303): complete subsection reference.

- [file](data-sources--workload--reference--group-015.md#canonical-3002023310331013-0011321210222212-2220122301103122-0200211333123133-1333101122021012-0002032101121130-0023031100212100-1023121300121132): complete subsection reference.

<a id="canonical-0311103231201330-2203202132201132-2231032300112002-0131121232112210-3333000331023313-2022022013000310-0031332233300200-3121130302021303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.configuration.parameters.env_var` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.configuration](data-sources--workload--reference--group-015.md#canonical-2220321213211303-0020122331330313-0211202333012210-0020030122022330-1222021321103030-0003330110113021-1132101131233203-3010122300030122)
- [simple_service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-1312302100222221-3310233130230312-3011122312201321-3020123032330102-1223232031312102-2111321113311321-1112011300203202-2330023033123211)
- simple_service.configuration.parameters.env_var

<a id="canonical-2331232021320110-3012231121013131-2230230123022303-1202103212131320-2332320220110032-2231303112231033-1103133111132203-1320232010112212"></a>

Type: `"single"`. Computed.

Environment Variable. Environment Variable.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0033202111001010-3223332011100030-0133203013230231-0033012021031123-2133323110300312-3010320011230210-0131203330212000-3231321031030002"></a>

### Direct properties for `simple_service.configuration.parameters.env_var`

<a id="canonical-0220111123330011-1103231100331002-0033300103121233-3321013022033310-0000031230232203-0032331302011320-0332103102000112-2031311310022102"></a>

#### `simple_service.configuration.parameters.env_var.name` property

Type: `"string"`. Computed.

Name. Name of Environment Variable.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1000222302323300-3321231122013112-2130100120030333-2211103021300321-0322031203001230-2120232123121200-2200321212221311-0331121203302331"></a>

<a id="canonical-2302111030201212-1030221011301020-3200312203003023-0230200011300011-1033333013310310-2313012211323010-0223032103100320-3300123023012202"></a>

#### `simple_service.configuration.parameters.env_var.value` property

Type: `"string"`. Computed.

Value. Value of Environment Variable.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3002023310331013-0011321210222212-2220122301103122-0200211333123133-1333101122021012-0002032101121130-0023031100212100-1023121300121132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.configuration.parameters.file` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.configuration](data-sources--workload--reference--group-015.md#canonical-2220321213211303-0020122331330313-0211202333012210-0020030122022330-1222021321103030-0003330110113021-1132101131233203-3010122300030122)
- [simple_service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-1312302100222221-3310233130230312-3011122312201321-3020123032330102-1223232031312102-2111321113311321-1112011300203202-2330023033123211)
- simple_service.configuration.parameters.file

<a id="canonical-1203022321320233-0110100012003103-0001133133231320-1131303112310220-0300023331321131-2023003233300133-1102233020110213-3122112112332221"></a>

Type: `"single"`. Computed.

Configuration File. Configuration File for the workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2320031111330213-1320112323111321-0103210120000333-3312213200120001-0130331321230100-3130103302302130-0321020000120122-2301331311222010"></a>

### Direct properties for `simple_service.configuration.parameters.file`

<a id="canonical-3102232302232310-2111033110223303-0030333300313102-0320230210330032-3303112223122031-3322032123120311-1102313302112030-2210232032210103"></a>

#### `simple_service.configuration.parameters.file.data` property

Type: `"string"`. Computed.

Data. File data

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16384,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 16384,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [mount](data-sources--workload--reference--group-015.md#canonical-1210202201113120-1201133233312013-0321131200300300-0021312002200220-3012333303132313-1323113233223003-3212200103222110-2321203211200121): complete subsection reference.

<a id="canonical-2113213023010331-1203133033303302-1300113113001320-0200332101202122-0130022111011001-3302130022121112-0112122011012122-1232030320102320"></a>

<a id="canonical-2320301131311101-3121033332003121-3212203211000111-2331312233232320-3002003121023101-2001113102113333-1220021021110113-1031001221213102"></a>

#### `simple_service.configuration.parameters.file.name` property

Type: `"string"`. Computed.

Name. Name of the file.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3002100123103221-3003331133313112-2302330203120332-3102010302300222-3122303031232123-0222002122110311-3020323022112303-1331001230333312"></a>

<a id="canonical-0211303021311130-3321233033133031-1212320212300110-3223301111330002-1132133301133222-2312312211030321-3033313331111010-1021023001132020"></a>

#### `simple_service.configuration.parameters.file.volume_name` property

Type: `"string"`. Computed.

Volume Name. Name of the Volume.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1210202201113120-1201133233312013-0321131200300300-0021312002200220-3012333303132313-1323113233223003-3212200103222110-2321203211200121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.configuration.parameters.file.mount` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.configuration](data-sources--workload--reference--group-015.md#canonical-2220321213211303-0020122331330313-0211202333012210-0020030122022330-1222021321103030-0003330110113021-1132101131233203-3010122300030122)
- [simple_service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-1312302100222221-3310233130230312-3011122312201321-3020123032330102-1223232031312102-2111321113311321-1112011300203202-2330023033123211)
- [simple_service.configuration.parameters.file](data-sources--workload--reference--group-015.md#canonical-3002023310331013-0011321210222212-2220122301103122-0200211333123133-1333101122021012-0002032101121130-0023031100212100-1023121300121132)
- simple_service.configuration.parameters.file.mount

<a id="canonical-0002210220233120-1032213010112131-0130133102212003-1302021211323101-1200231323311211-1031230222122131-3013123201101213-3113030113313010"></a>

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

<a id="canonical-2312213132110131-3131102231131100-0133203300202220-2100331200212322-0220323231331302-1102031103323212-2032220013211033-0110331023033312"></a>

### Direct properties for `simple_service.configuration.parameters.file.mount`

<a id="canonical-2312302330030212-1332213201011221-2330003202130301-2222212123212103-0013323002132222-1132110202021211-1031333231022113-3120120123220210"></a>

#### `simple_service.configuration.parameters.file.mount.mode` property

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

<a id="canonical-3222113232031022-3013221322113102-0333213300301330-1301313321320233-0222132001020231-2330312311033002-1102131110303002-3211231133110212"></a>

<a id="canonical-0211123311033102-0203213101021011-3032322013130332-2032213303111131-2020312111103020-0303201321303211-0203211120121132-1201300101312300"></a>

#### `simple_service.configuration.parameters.file.mount.mount_path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2100120000313221-1230202011133033-0231212011002031-0320120120331012-3132000123312130-3320213331122333-1210000113202202-2030123330021003"></a>

<a id="canonical-3232123000331110-2131220202200120-3031023021202123-1002021322332233-1120212022213003-0312332120031322-2303112331131023-0021233320100311"></a>

#### `simple_service.configuration.parameters.file.mount.sub_path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- simple_service.container

<a id="canonical-2300323110211023-2132103213210001-0021322100130333-0201021332022122-1223021313322033-1333212113201233-1212121320310231-0011020203300322"></a>

Type: `"single"`. Computed.

ContainerType configures the container information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-flavor_choice": "[\"custom_flavor\",\"default_flavor\",\"flavor\"]"
}
```

<a id="canonical-3101102013001332-3013201012013010-3001013111222212-1010020110231333-1133003102232202-3230121100301031-2110112101213133-0111322333031132"></a>

### Direct properties for `simple_service.container`

<a id="canonical-0030331313211012-3313312031023113-3302023100320312-2322230211100113-3030110133230300-0031022021201331-1201220301202220-0310023103120111"></a>

#### `simple_service.container.args` property

Type: `["list", "string"]`. Computed.

Arguments to the entrypoint. Overrides the Docker image's CMD.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-3100011133123132-1000031003020023-1010211203111131-0112111331112333-2020330110300012-2112012103023223-2213212331032000-1230333310132021"></a>

<a id="canonical-2203303231232002-2002032120122212-3331001122111023-0200222030200231-3131202032222312-1111113003032230-3203121212211311-1010021012321313"></a>

#### `simple_service.container.command` property

Type: `["list", "string"]`. Computed.

Command to execute. Overrides the Docker image's ENTRYPOINT.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [custom_flavor](data-sources--workload--reference--group-015.md#canonical-0322113130021212-0120003213222223-2202021123032322-2111201110233212-1030212320131002-2022121220000203-3012312313011221-2330023333321132): complete subsection reference.

- [default_flavor](data-sources--workload--reference--group-015.md#canonical-3201130223113330-0011130232011320-2022033110110011-2020321202031121-2213303310031030-0332211033301102-1120032023313132-3322030011000001): complete subsection reference.

<a id="canonical-3333213330112222-2313111032213300-0022113122202302-0223303001122020-1201023223303013-3222222003032000-2202301310321110-3112013023012121"></a>

<a id="canonical-1012130202110133-1120213121002121-3003113021213011-0232011111223121-2101210233000013-2002030313300033-2020013212033232-0132312231202222"></a>

#### `simple_service.container.flavor` property

Type: `"string"`. Computed.

\[Enum:
CONTAINER\_FLAVOR\_TYPE\_TINY|CONTAINER\_FLAVOR\_TYPE\_MEDIUM|CONTAINER\_FLAVOR\_TYPE\_LARGE\]
Container Flavor type - CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny Tiny containers have limit of 0.1 vCPU
and 256 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium Medium containers have limit
of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_LARGE: Large Large containers
have.. Possible values are \`CONTAINER\_FLAVOR\_TYPE\_TINY\`, \`CONTAINER\_FLAVOR\_TYPE\_MEDIUM\`,
\`CONTAINER\_FLAVOR\_TYPE\_LARGE\`. Defaults to \`CONTAINER\_FLAVOR\_TYPE\_TINY\`.

Additional upstream details:

Container Flavor type

&#8203;- CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny

Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium

Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_LARGE: Large

Large containers have limit of 1 vCPU and 2048 MiB (mebibyte) memory.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTAINER_FLAVOR_TYPE_TINY",
  "enum": [
    "CONTAINER_FLAVOR_TYPE_TINY",
    "CONTAINER_FLAVOR_TYPE_MEDIUM",
    "CONTAINER_FLAVOR_TYPE_LARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [image](data-sources--workload--reference--group-015.md#canonical-0320303012022222-0102013233200321-3233211002302123-0331020012101021-1013001232102300-0233102021220301-3333320212331010-1031111110300203): complete subsection reference.

<a id="canonical-2211133323131321-2123231302212120-3013120231021121-0322113203331322-0211331313211103-2321122011312322-1120211232312031-2333223201111101"></a>

<a id="canonical-1322012000332020-1322001232003102-3100310201020120-2213220113331012-1032020131031201-3233123131002211-2013211201030322-3001111130120133"></a>

#### `simple_service.container.init_container` property

Type: `"bool"`. Computed.

Specialized container that runs before application container and runs to completion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [liveness_check](data-sources--workload--reference--group-015.md#canonical-1000001320031212-3211122331022310-3122021013232032-3330131020312223-3222210202130232-0222321120213120-1102002210120022-2120333031230103): complete subsection reference.

<a id="canonical-2103032103220112-2102320121033112-0320121223322231-1033113320011120-1000132113132233-3332202330230120-3203221011321132-2300312303230120"></a>

<a id="canonical-3132000001333012-3212330233330112-1111332130332212-1310020200121122-1201101100303302-2200002321210232-0133213120102123-1332111301212133"></a>

#### `simple_service.container.name` property

Type: `"string"`. Computed.

Name. Name of the container.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [readiness_check](data-sources--workload--reference--group-015.md#canonical-0312322121032303-0120231311113002-3232323333022313-0313301223202010-0113120313112233-3202000230211322-1223203103333003-2223013222202203): complete subsection reference.

<a id="canonical-0322113130021212-0120003213222223-2202021123032322-2111201110233212-1030212320131002-2022121220000203-3012312313011221-2330023333321132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.custom_flavor` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- simple_service.container.custom_flavor

<a id="canonical-3323031211211102-2101121230221223-0212031121221330-0103130322320333-3112311132331232-1122012303232123-0100233122202330-3101311233130222"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2301003333303121-2221012111131031-1333300330322311-1322131223032230-0110110201023333-0020213313332011-2223333333013223-2112202121132210"></a>

### Direct properties for `simple_service.container.custom_flavor`

<a id="canonical-2133023120320000-2111221021022133-1202313333322330-3003031333222132-1013130210030022-0321233011303030-2231021221233031-0331321001022122"></a>

#### `simple_service.container.custom_flavor.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2220213103213323-0332223321130131-2122102131112011-2321320100101132-3233013301100100-2322202001311012-0021113231123120-3132130031333311"></a>

<a id="canonical-0101100020132302-0321330011210000-1133012221021003-2300122320122101-1123301001101230-0233030302220323-0013312323030231-0120102030101330"></a>

#### `simple_service.container.custom_flavor.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3230001012032321-3133132133003301-3230320020012010-1103203022222002-2332103120001223-0303311012211231-3113302231322331-0233320230031323"></a>

<a id="canonical-2200032232131330-0002232102031221-1223012010231133-1111101111301303-3203233131313223-0303103303012300-1123332230303220-2300233233103130"></a>

#### `simple_service.container.custom_flavor.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3201130223113330-0011130232011320-2022033110110011-2020321202031121-2213303310031030-0332211033301102-1120032023313132-3322030011000001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.default_flavor` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- simple_service.container.default_flavor

<a id="canonical-2030012220001312-1003201223322021-0230233222000212-0122212331013202-3112100302333333-2131123202233021-2312123232201233-3110113120333210"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default flavor.

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

<a id="canonical-0320303012022222-0102013233200321-3233211002302123-0331020012101021-1013001232102300-0233102021220301-3333320212331010-1031111110300203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.image` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- simple_service.container.image

<a id="canonical-2022033312302231-0303323211111013-3113132230012313-1013112002123021-3013101210320211-1220211013200321-3301120102331031-2000331011211133"></a>

Type: `"single"`. Computed.

ImageType configures the image to use, how to pull the image, and the associated secrets to use if
any.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-registry_choice": "[\"container_registry\",\"public\"]"
}
```

<a id="canonical-0132032300101301-1222133213301101-2101332100031232-1133313030012331-1333102313131010-2230320102210313-2210102122303013-3312110031300200"></a>

### Direct properties for `simple_service.container.image`

- [container_registry](data-sources--workload--reference--group-015.md#canonical-3100133310022331-3113223213212030-2200023302110021-0330323213310003-2211300202313133-3211333211030232-0311031332203211-2202131310112330): complete subsection reference.

<a id="canonical-3030322232032300-1231020233031302-0012001322232300-3000132130130333-1033202221130212-3300203202213023-3013323111320132-1021033020310003"></a>

<a id="canonical-2132022313120132-1202211310200233-1120312231031131-1313231011033311-0030221032130010-0100021000001012-3310302200320110-3133003322013120"></a>

#### `simple_service.container.image.name` property

Type: `"string"`. Computed.

Name is a container image which are usually given a name such as alpine, Ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [public](data-sources--workload--reference--group-015.md#canonical-0023332123233323-2002200231322313-3120210333220312-2322212120312320-2232311303230023-1012200120320020-3021331032111313-3310012121123220): complete subsection reference.

<a id="canonical-0322012000311000-2312121110131313-2311301210023121-3230223210112312-3131331213021022-3122230122013002-2230002231221332-2232202320333032"></a>

<a id="canonical-3210101302113021-1231131300131301-2300132000310303-2232320031123231-2121231311123111-2210110322130012-0210310011320210-0203220001223323"></a>

#### `simple_service.container.image.pull_policy` property

Type: `"string"`. Computed.

\[Enum:
IMAGE\_PULL\_POLICY\_DEFAULT|IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT|IMAGE\_PULL\_POLICY\_ALWAYS|IMAGE\_PULL\_POLICY\_NEVER\]
Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload - IMAGE\_PULL\_POLICY\_DEFAULT: Default Default will always pull image if :latest tag
is specified in image name. If :latest tag is not specified in image name, it will pull image only..
Possible values are \`IMAGE\_PULL\_POLICY\_DEFAULT\`, \`IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT\`,
\`IMAGE\_PULL\_POLICY\_ALWAYS\`, \`IMAGE\_PULL\_POLICY\_NEVER\`. Defaults to
\`IMAGE\_PULL\_POLICY\_DEFAULT\`.

Additional upstream details:

Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload

&#8203;- IMAGE\_PULL\_POLICY\_DEFAULT: Default

Default will always pull image if :latest tag is specified in image name. If :latest tag is not
specified in image name, it will pull image only if it does not already exist on the node &#8203;-
IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT: IfNotPresent

Only pull the image if it does not already exist on the node &#8203;- IMAGE\_PULL\_POLICY\_ALWAYS:
Always

Always pull the image &#8203;- IMAGE\_PULL\_POLICY\_NEVER: Never

Never pull the image.

Receipt-pinned upstream constraints:

```json
{
  "default": "IMAGE_PULL_POLICY_DEFAULT",
  "enum": [
    "IMAGE_PULL_POLICY_DEFAULT",
    "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
    "IMAGE_PULL_POLICY_ALWAYS",
    "IMAGE_PULL_POLICY_NEVER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3100133310022331-3113223213212030-2200023302110021-0330323213310003-2211300202313133-3211333211030232-0311031332203211-2202131310112330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.image.container_registry` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- [simple_service.container.image](data-sources--workload--reference--group-015.md#canonical-0320303012022222-0102013233200321-3233211002302123-0331020012101021-1013001232102300-0233102021220301-3333320212331010-1031111110300203)
- simple_service.container.image.container_registry

<a id="canonical-2221210103123023-3113012102013313-3003112311332012-2300213021133322-0131132230031312-3021121030120031-3233333233201133-0100232031220331"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3301313232300033-0012131311011121-1310030213231131-1310201333300332-0212031031123100-1332321331120311-0101311022202030-1201313221222313"></a>

### Direct properties for `simple_service.container.image.container_registry`

<a id="canonical-3323120122310222-1030102300101210-3132022332322033-3112102012131110-1001300312203223-2212201123130133-3000000222220320-2310112300021132"></a>

#### `simple_service.container.image.container_registry.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1322321002033303-0102212220133313-0021203233201232-2121010222121221-3123122231021301-0121331100313102-2113123333230102-1102203302302333"></a>

<a id="canonical-2333122023301021-1031233320312103-3020111320302122-2012321013133022-0303330132120232-2130213223120010-1001013021300323-3301312303301300"></a>

#### `simple_service.container.image.container_registry.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1332223332323120-2210023323120022-0220103303110231-1121102123130113-3112212202001130-2030122113101000-0300113032112011-3203021220133300"></a>

<a id="canonical-2230123133322201-2101313032212330-1322222000202113-0332222030223121-2232011322320321-2033020230023210-0020213121102011-1003130003200130"></a>

#### `simple_service.container.image.container_registry.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0023332123233323-2002200231322313-3120210333220312-2322212120312320-2232311303230023-1012200120320020-3021331032111313-3310012121123220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.image.public` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- [simple_service.container.image](data-sources--workload--reference--group-015.md#canonical-0320303012022222-0102013233200321-3233211002302123-0331020012101021-1013001232102300-0233102021220301-3333320212331010-1031111110300203)
- simple_service.container.image.public

<a id="canonical-3303310112031032-1213223211102201-3133122022021212-3203131001320133-3103320332312003-2102300321031100-3203130221323302-1033232103011120"></a>

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

<a id="canonical-1000001320031212-3211122331022310-3122021013232032-3330131020312223-3222210202130232-0222321120213120-1102002210120022-2120333031230103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.liveness_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- simple_service.container.liveness_check

<a id="canonical-2121320010323110-2133332201333021-3131031321330033-2032302013233111-2021321200301333-0111003300001230-0312210222113130-2112302233322301"></a>

Type: `"single"`. Computed.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

<a id="canonical-1302111121012111-0022303100031221-0031031330230223-2301212013032112-2001333112333123-2312110033221033-2232000300120302-1110312201112111"></a>

### Direct properties for `simple_service.container.liveness_check`

- [exec_health_check](data-sources--workload--reference--group-015.md#canonical-2133211220011223-2330203100011300-1020021203311113-0212210202331320-1302012013121033-0001320111001130-2132202033022211-0103110213332333): complete subsection reference.

<a id="canonical-2300013111003302-0220031010231133-3233332102011101-2233303322333010-1311111003103132-0311233321021232-2303131113330332-1322120010201201"></a>

<a id="canonical-2332103022321013-1213032220231320-1213121112321331-3130131130030323-1221220302131103-0010203112120001-0121121230112031-0001311232312002"></a>

#### `simple_service.container.liveness_check.healthy_threshold` property

Type: `"number"`. Computed.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](data-sources--workload--reference--group-015.md#canonical-3222103322132103-0310130002332210-0100322111130121-3100012213200222-2013230001020130-0013213201013233-2113031123302120-2021110321101112): complete subsection reference.

<a id="canonical-2001212132333203-2012223132203110-2132113010322333-0223303003003003-3012120221220132-2001202313233103-1012220003202203-3200001202100232"></a>

<a id="canonical-2001123312213131-0011032031233232-3230130210301330-1320001033312130-2323302031031020-0002011220031012-2123330120022323-3100231101022121"></a>

#### `simple_service.container.liveness_check.initial_delay` property

Type: `"number"`. Computed.

Number of seconds after the container has started before health checks are initiated.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-2103113323020202-1200030131220301-0012101123310222-0033213233221310-3020333033331013-1023333322032022-2033310133130231-0100210111311020"></a>

<a id="canonical-1212312223312201-1113231331211003-2323303110003113-3210021021231003-0023123130011013-0103113332302230-2132211010233001-2213113110133212"></a>

#### `simple_service.container.liveness_check.interval` property

Type: `"number"`. Computed.

Time interval in seconds between two health check requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](data-sources--workload--reference--group-015.md#canonical-2032231103032010-3220203132033331-3010332021232021-0030203121323211-2320331001100232-2213121310301211-0101010201120103-1212023213330120): complete subsection reference.

<a id="canonical-2311131320301212-3002022232220102-0233032332331021-1030003012212332-3313133110303231-0233222330322131-3221113023133301-3310311020212312"></a>

<a id="canonical-3002201122013331-2010313023102103-2002221131210213-3122233310030122-1101121332102212-3310311231000333-2230101312330202-1121032032131021"></a>

#### `simple_service.container.liveness_check.timeout` property

Type: `"number"`. Computed.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-2102300101123222-2333130312020223-3233331111130030-1121111113113331-0113300021300200-1111022010312113-0120211302020121-3230212312330000"></a>

<a id="canonical-0221310110013031-3321202330003012-1220310022303210-0320133221021111-0101010212001222-3023010000113122-3022311230220201-0212111103123331"></a>

#### `simple_service.container.liveness_check.unhealthy_threshold` property

Type: `"number"`. Computed.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-2133211220011223-2330203100011300-1020021203311113-0212210202331320-1302012013121033-0001320111001130-2132202033022211-0103110213332333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.liveness_check.exec_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- [simple_service.container.liveness_check](data-sources--workload--reference--group-015.md#canonical-1000001320031212-3211122331022310-3122021013232032-3330131020312223-3222210202130232-0222321120213120-1102002210120022-2120333031230103)
- simple_service.container.liveness_check.exec_health_check

<a id="canonical-2332021131120121-1132302303103231-2022111000231330-0131030111020130-2103130313332010-2310220102321102-2001202220031131-3011001031023233"></a>

Type: `"single"`. Computed.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Additional upstream details:

ExecHealthCheckType describes a health check based on "run in container" action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2301033300011020-1323322012102021-0310102123300312-2030111312101131-1212322010322333-1223100003203202-2212333233020323-3002120100031300"></a>

### Direct properties for `simple_service.container.liveness_check.exec_health_check`

<a id="canonical-2320323300123000-3313121032020103-0221032221330213-0300100032220003-1311011231110320-2201333031311321-3330023012302031-0303123001310232"></a>

#### `simple_service.container.liveness_check.exec_health_check.command` property

Type: `["list", "string"]`. Computed.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3222103322132103-0310130002332210-0100322111130121-3100012213200222-2013230001020130-0013213201013233-2113031123302120-2021110321101112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.liveness_check.http_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- [simple_service.container.liveness_check](data-sources--workload--reference--group-015.md#canonical-1000001320031212-3211122331022310-3122021013232032-3330131020312223-3222210202130232-0222321120213120-1102002210120022-2120333031230103)
- simple_service.container.liveness_check.http_health_check

<a id="canonical-1012113333312133-3200211013322201-2000211322321231-1120231320200132-2003021011232222-0301321003102132-0302003130311022-2021013033332203"></a>

Type: `"single"`. Computed.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0201300311131021-2203021132322223-1302112100230123-3331300133211033-0003131022232211-1223032331122230-0202023230310320-1113202022001222"></a>

### Direct properties for `simple_service.container.liveness_check.http_health_check`

<a id="canonical-3202103100013131-0122333202220322-1120111102200201-3320220111230232-3311101322203103-0032123131003123-1312321030303000-3223223033112100"></a>

#### `simple_service.container.liveness_check.http_health_check.headers` property

Type: `["map", "string"]`. Computed.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

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
    "keys": {
      "maxLength": 256,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "256",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "2048",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 2048,
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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-1310300212003001-1101032011123001-3303123000303102-0232331211312330-1330300331100102-2031203021013132-3233320223101120-0123300333020300"></a>

<a id="canonical-0332132111220001-1320100203321303-1311232112310030-1023133112010321-0131113121313302-1210311302032212-0221331202033102-0211100030223013"></a>

#### `simple_service.container.liveness_check.http_health_check.host_header` property

Type: `"string"`. Computed.

The value of the host header in the HTTP health check request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-1101212130310031-2110323331020002-3021323111001113-2112201010233331-3010313333031112-1022301232200120-2330031332330230-1331202200312212"></a>

<a id="canonical-3212200310203330-2212321303311231-0233202111003231-2112023321022213-1030221032220100-2101322320000330-1311033213203100-3003000233132310"></a>

#### `simple_service.container.liveness_check.http_health_check.path` property

Type: `"string"`. Computed.

Path. Path to access on the HTTP server.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](data-sources--workload--reference--group-015.md#canonical-1133033122131021-0003030201311231-3232231223012312-2031003231122132-3210223321113303-0230200021001000-1001122031212201-2303312022023301): complete subsection reference.

<a id="canonical-1133033122131021-0003030201311231-3232231223012312-2031003231122132-3210223321113303-0230200021001000-1001122031212201-2303312022023301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.liveness_check.http_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- [simple_service.container.liveness_check](data-sources--workload--reference--group-015.md#canonical-1000001320031212-3211122331022310-3122021013232032-3330131020312223-3222210202130232-0222321120213120-1102002210120022-2120333031230103)
- [simple_service.container.liveness_check.http_health_check](data-sources--workload--reference--group-015.md#canonical-3222103322132103-0310130002332210-0100322111130121-3100012213200222-2013230001020130-0013213201013233-2113031123302120-2021110321101112)
- simple_service.container.liveness_check.http_health_check.port

<a id="canonical-1130330010231211-3103020301321131-3313311130110123-3212102020033300-2310313212122301-1223211202212203-3110302001031301-0210220100122311"></a>

Type: `"single"`. Computed.

Port. Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

<a id="canonical-1102130032102133-1011212022113210-1221322331313131-0010000301212121-3100302102020310-1103003032011001-3220133320213331-3122121201131013"></a>

### Direct properties for `simple_service.container.liveness_check.http_health_check.port`

<a id="canonical-3020303003320033-3023033220331222-3103212233312333-0323002331103222-2231120102012121-3302323323201233-1013201320120201-1303121221013201"></a>

#### `simple_service.container.liveness_check.http_health_check.port.name` property

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-0100003001000012-1201112232002003-2302330013033312-3101110320202303-3031211313231200-0000302202210021-2100303021013210-3323011230123300"></a>

<a id="canonical-0121020122223301-0131222002321030-1200121002112101-3100310111023320-3103300302122021-1132130231332320-3213113033312320-0310330103122231"></a>

#### `simple_service.container.liveness_check.http_health_check.port.num` property

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2032231103032010-3220203132033331-3010332021232021-0030203121323211-2320331001100232-2213121310301211-0101010201120103-1212023213330120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.liveness_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- [simple_service.container.liveness_check](data-sources--workload--reference--group-015.md#canonical-1000001320031212-3211122331022310-3122021013232032-3330131020312223-3222210202130232-0222321120213120-1102002210120022-2120333031230103)
- simple_service.container.liveness_check.tcp_health_check

<a id="canonical-2301332223100211-2110033033313233-0310122321102212-2211212310323223-0210231202030023-1221222222302303-3102130212102133-1133123130330030"></a>

Type: `"single"`. Computed.

TCPHealthCheckType describes a health check based on opening a TCP connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2221101103310212-0001102121122103-1031332123000111-1020223221132201-3332330122000012-2120311220021010-3022312102110013-1231201030132120"></a>

### Direct properties for `simple_service.container.liveness_check.tcp_health_check`

- [port](data-sources--workload--reference--group-015.md#canonical-2310131301023121-2313021113230302-2332003223112003-0200131211302123-0333121101203332-1212320300201020-2312103201111121-0333131310121120): complete subsection reference.

<a id="canonical-2310131301023121-2313021113230302-2332003223112003-0200131211302123-0333121101203332-1212320300201020-2312103201111121-0333131310121120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.liveness_check.tcp_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- [simple_service.container.liveness_check](data-sources--workload--reference--group-015.md#canonical-1000001320031212-3211122331022310-3122021013232032-3330131020312223-3222210202130232-0222321120213120-1102002210120022-2120333031230103)
- [simple_service.container.liveness_check.tcp_health_check](data-sources--workload--reference--group-015.md#canonical-2032231103032010-3220203132033331-3010332021232021-0030203121323211-2320331001100232-2213121310301211-0101010201120103-1212023213330120)
- simple_service.container.liveness_check.tcp_health_check.port

<a id="canonical-2200120001020120-3023033012230200-3100223320130130-0321213223023230-3103013000300222-3020301002330212-3312001311330200-3232032232103302"></a>

Type: `"single"`. Computed.

Port. Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

<a id="canonical-3222112310101332-0221102121120203-1102332033210330-3133011310010222-1322133210113000-1311113320000233-1230011003223232-0103002301230322"></a>

### Direct properties for `simple_service.container.liveness_check.tcp_health_check.port`

<a id="canonical-3101212002010000-0321232230230120-2303012131312121-1012332111020020-1313202211101223-1333102200103212-3303311113020003-3210113332320210"></a>

#### `simple_service.container.liveness_check.tcp_health_check.port.name` property

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-0310303121211023-0130013302232231-0113220320113032-3031221020203332-2312312011002322-1132313003212113-0000332212211131-1111031031010232"></a>

<a id="canonical-2301023012100200-1112201312102200-1330333331021113-3111321202030011-0223333003101023-1201201231122322-2012313302023323-2011111003101023"></a>

#### `simple_service.container.liveness_check.tcp_health_check.port.num` property

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0312322121032303-0120231311113002-3232323333022313-0313301223202010-0113120313112233-3202000230211322-1223203103333003-2223013222202203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.readiness_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- simple_service.container.readiness_check

<a id="canonical-2113010201212222-1311222232320221-0130122302011210-0120110331032021-0102330133123032-3220302032233123-0223233120311311-2102302112121032"></a>

Type: `"single"`. Computed.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

<a id="canonical-3332202032010331-1201123123313133-1312233320031123-1020232133010300-3202202200032213-1131230333323302-0023322232331212-3312213131303210"></a>

### Direct properties for `simple_service.container.readiness_check`

- [exec_health_check](data-sources--workload--reference--group-015.md#canonical-3120202100010300-3312202210311013-0101220202010002-0320202110330020-1020133211020023-1323111233201103-0211131131232333-3220230310201221): complete subsection reference.

<a id="canonical-1002233202220103-3111211113113102-3020332213302030-2333112203330200-0022030110133002-3210201122333123-3232212120321033-0233021122303132"></a>

<a id="canonical-1202032331030303-3313331321001003-1333212330012012-3113323213210323-3301121303020030-2312211112120112-3203212122333011-0311201201211013"></a>

#### `simple_service.container.readiness_check.healthy_threshold` property

Type: `"number"`. Computed.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](data-sources--workload--reference--group-016.md#canonical-2020220101001102-2133303210311300-0310110110222130-0100110211223233-2030002322201020-2310231223210200-2213111230102200-3011023323231031): complete subsection reference.

<a id="canonical-1200310223213333-2121301033203120-1101230111031002-2031020220123323-0010033330310030-0303013102211112-2202213112023331-0110102312131302"></a>

<a id="canonical-0313332012213102-2232303000123002-0330330033002100-1223021330333213-2332332102103310-1333102033003132-2333333130330020-2221100021010100"></a>

#### `simple_service.container.readiness_check.initial_delay` property

Type: `"number"`. Computed.

Number of seconds after the container has started before health checks are initiated.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-2001122223020331-3023313021130110-0300033312010211-3220112030120013-1321132000130202-3013121221232333-0213010112020303-1230213211121223"></a>

<a id="canonical-2220320200101311-0311120230121120-1002012023122311-2200201101001023-1221302032033013-3112303100133332-2010311330023303-3332321320122313"></a>

#### `simple_service.container.readiness_check.interval` property

Type: `"number"`. Computed.

Time interval in seconds between two health check requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](data-sources--workload--reference--group-016.md#canonical-2003212223310021-3331301230033133-2110300232002230-1202301130320132-3231030132313333-2123332103111333-2332101310010223-2302122111302123): complete subsection reference.

<a id="canonical-2321120011000112-2113031310202320-3222211211232031-1301131030031202-0010210133101110-2223132111302131-1101201331303201-1200301302011012"></a>

<a id="canonical-2021310003231110-1211203033230101-1113010112233201-3132012023131112-0301110113102100-1223333021322133-1222330322002331-0033303031132100"></a>

#### `simple_service.container.readiness_check.timeout` property

Type: `"number"`. Computed.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-0301330130230213-1320102310003023-1321323321023231-0313112210203133-3103011313331331-3030323110101012-2011200013103100-2202301122110103"></a>

<a id="canonical-1303313200313310-3333321231330230-3010031102301202-0232023210312003-1032010102303123-3010320333022100-2121131012021000-0203000303130103"></a>

#### `simple_service.container.readiness_check.unhealthy_threshold` property

Type: `"number"`. Computed.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-3120202100010300-3312202210311013-0101220202010002-0320202110330020-1020133211020023-1323111233201103-0211131131232333-3220230310201221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.readiness_check.exec_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- [simple_service.container.readiness_check](data-sources--workload--reference--group-015.md#canonical-0312322121032303-0120231311113002-3232323333022313-0313301223202010-0113120313112233-3202000230211322-1223203103333003-2223013222202203)
- simple_service.container.readiness_check.exec_health_check

<a id="canonical-1230011223133202-1003222120230130-2220001031230212-3232123200311331-0202031101322333-1103301020331121-1131312210002332-1013010101002021"></a>

Type: `"single"`. Computed.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Additional upstream details:

ExecHealthCheckType describes a health check based on "run in container" action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2000232001030012-1331100020220322-0220311211132233-2022310131123000-2030133230222200-1021102200120330-1131033232031011-1010103130310201"></a>

### Direct properties for `simple_service.container.readiness_check.exec_health_check`

<a id="canonical-0030022321301223-1010130222133222-0132131002002212-0011201310312213-0220201223132322-2301132221032213-3221113200233122-2023203113302312"></a>

#### `simple_service.container.readiness_check.exec_health_check.command` property

Type: `["list", "string"]`. Computed.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
