---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1230212122322002-2121020302223322-1111021002133302-2021301201301010-1012022203330120-2211330111121130-3201330303120333-1123022211033230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.configuration.parameters.file.mount` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.configuration](data-sources--workload--reference--group-013.md#canonical-3123002002331230-0201233111020120-1333221333320003-3110223112030011-2003333010313113-2333211001212031-3333112223203131-0322131002310210)
- [service.configuration.parameters](data-sources--workload--reference--group-013.md#canonical-0012011320103113-3033302112011313-2110032033231102-1120002333001312-0230133122323022-2322200011110000-1310321323010332-0020233203221233)
- [service.configuration.parameters.file](data-sources--workload--reference--group-013.md#canonical-3332330103203231-2232213211320013-3031002131033012-0231300111321033-0200101300032231-0102031133310322-1202321212212011-3303003200302311)
- service.configuration.parameters.file.mount

<a id="canonical-3331122212302203-3100332222020132-1030023321013303-0010312122333302-0313212303310203-3023231312201112-0022213120300130-1003022300233202"></a>

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

<a id="canonical-2022320123021312-3320202113020120-3101010330002031-1310311011113011-0311000312031303-2313212331001100-1010230123213302-1320112212200321"></a>

### Direct properties for `service.configuration.parameters.file.mount`

<a id="canonical-2231322312222202-0020031311322333-2222223201100200-0101321101122312-0103102222133232-0112010222120222-3132321321120122-1300203020031101"></a>

#### `service.configuration.parameters.file.mount.mode` property

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

<a id="canonical-3011333330032312-1202321101333022-1330232133230100-1211230323002220-3122203123110122-3332302233212313-1011310121011300-3222223203203120"></a>

<a id="canonical-2022101033003102-2231222001231222-2221022022021300-0323113031113312-2131102311033030-2011201331033203-2100031122221121-1120003230320010"></a>

#### `service.configuration.parameters.file.mount.mount_path` property

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

<a id="canonical-3221312222213120-0110120122113320-2112113302111021-2323201202033102-1033232330211333-3012121302123312-0320121330031302-1123013311021223"></a>

<a id="canonical-3201113133333010-2100101112033032-2302130301221000-3313213223312230-3330302330020233-3330233201112133-3230101123311200-2021333311030020"></a>

#### `service.configuration.parameters.file.mount.sub_path` property

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

<a id="canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- service.containers

<a id="canonical-1030103313112111-3203302313130121-1101132111201320-2302232210123111-1200221232312022-0330311210023211-3301330100012202-1223003001221031"></a>

Type: `"list"`. Computed.

Containers. Containers to use for service.

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

<a id="canonical-1232333202302221-0233211020110100-1130133302101020-2312320223211203-0020220031021001-1313002033113332-2002100103023313-1123313222312230"></a>

### Direct properties for `service.containers`

<a id="canonical-1221132310133302-2322320333211033-2210331111120003-2223301101331112-0012100010201223-0203100122130212-2102130210012012-3322013222213302"></a>

#### `service.containers.args` property

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

<a id="canonical-2021131332311132-0230332233301300-2313302012322210-3201012111003131-0030031013322320-2101130021012002-1023212000300332-3232301320103130"></a>

<a id="canonical-0120312101103001-3122102320123001-1021320301212322-2020331322230303-3022130002323300-3211230310022300-3133123220311200-2233332002222020"></a>

#### `service.containers.command` property

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

- [custom_flavor](data-sources--workload--reference--group-014.md#canonical-0220222331120113-2122123100010221-0310120331213320-0123133100313303-1201123030002230-0120232210311131-1210023200022301-1323320310110230): complete subsection reference.

- [default_flavor](data-sources--workload--reference--group-014.md#canonical-2233010122103202-0301112200120010-3023311330130020-0101312220031302-2233021333311333-1011212003200123-3331100212230113-1002313301313231): complete subsection reference.

<a id="canonical-2012003210032020-1023000122022221-3130212032320212-1033331223221312-1032310300313233-1101203302332121-2021033020300320-0002232302332133"></a>

<a id="canonical-0130103001202100-0131010102023011-0223313113212320-2133012321201212-0310010002202122-0001302311210131-1230011320210101-2000221023300003"></a>

#### `service.containers.flavor` property

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

- [image](data-sources--workload--reference--group-014.md#canonical-2031322233230031-1232103120020002-3332210332031001-3331321000310203-0102013211002110-0032032121233323-3230101213231312-3030302330121333): complete subsection reference.

<a id="canonical-0320133032010211-1110202010013310-1023011210133203-2001020311112000-2123110000303332-3022120303211303-2321010323130311-3002031003311023"></a>

<a id="canonical-2001021323023312-0211313210231103-2200111222021203-0301223030000311-3033111200201222-2212003311120310-3300002301302031-2322200012110122"></a>

#### `service.containers.init_container` property

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

- [liveness_check](data-sources--workload--reference--group-014.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120): complete subsection reference.

<a id="canonical-3333133230213323-1012110220310030-1200323213133201-0213012113331021-1132201013130323-3330200023310031-3113321323323313-3031023332000301"></a>

<a id="canonical-1121133222101303-3202023322232113-0011212210123230-2110132210301102-2213113122121032-0332000313232312-2323033322130212-0212220100213103"></a>

#### `service.containers.name` property

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

- [readiness_check](data-sources--workload--reference--group-014.md#canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012): complete subsection reference.

<a id="canonical-0220222331120113-2122123100010221-0310120331213320-0123133100313303-1201123030002230-0120232210311131-1210023200022301-1323320310110230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.custom_flavor` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- service.containers.custom_flavor

<a id="canonical-1023312133103123-1132100121032322-3230011302030323-0220202100313011-2032030111112300-2300323003223212-1232013011330102-2133201133300221"></a>

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

<a id="canonical-0213003312031302-1302032232102113-0203023022130310-2021021021012310-3133031001332110-0122033330231021-1302302010101221-1230022312233332"></a>

### Direct properties for `service.containers.custom_flavor`

<a id="canonical-2002333010332233-2320310121110012-0230012001231120-0013011222233123-1120013033133012-0332032213200222-2023302232033202-2122032330100032"></a>

#### `service.containers.custom_flavor.name` property

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

<a id="canonical-2032310230323001-1321301110020011-2311011200313123-3313021221121312-0120332133102313-3230223100220121-3323000232121200-0113301103122030"></a>

<a id="canonical-0313031021120311-1122300200133200-2030312001012221-0100212202220231-1213003312201300-3022021021231201-2021312312313332-0303332323233013"></a>

#### `service.containers.custom_flavor.namespace` property

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

<a id="canonical-2301330013021333-3333333110120022-2131331301022102-1322221120023320-1112012023310323-3231211330110013-0122312021320301-0212103323211213"></a>

<a id="canonical-2132201132132022-3132321122201120-3310221300033133-1200203103321002-0201211301002131-0122231030110020-1122101010131000-1030010032130030"></a>

#### `service.containers.custom_flavor.tenant` property

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

<a id="canonical-2233010122103202-0301112200120010-3023311330130020-0101312220031302-2233021333311333-1011212003200123-3331100212230113-1002313301313231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.default_flavor` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- service.containers.default_flavor

<a id="canonical-1102131220000011-3030011122332331-3003211220100103-1103333000132313-2232101312012112-3311200312222100-0031033110203220-2133032020320300"></a>

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

<a id="canonical-2031322233230031-1232103120020002-3332210332031001-3331321000310203-0102013211002110-0032032121233323-3230101213231312-3030302330121333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.image` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- service.containers.image

<a id="canonical-2031303113222323-1003133102310220-1123133330203012-3001301132110303-2000003031231002-0213022223213330-1000232323223322-0222321203101133"></a>

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

<a id="canonical-3233223102202201-1232202133121211-2231011132310001-0100102203322311-2331310013020231-3312032201003102-1312221110222233-1202203310311010"></a>

### Direct properties for `service.containers.image`

- [container_registry](data-sources--workload--reference--group-014.md#canonical-0200232232103113-0020033231030001-2112000331001000-1120312203111231-2131221223033001-0112022312302223-1232302032023123-3100021001322300): complete subsection reference.

<a id="canonical-2210332323133001-0213112231321210-0203110200331122-3003011031333220-2302203321231132-3230312013133123-1222223320321200-3023020123322003"></a>

<a id="canonical-1030030323200013-0211323212202010-1320322021030102-1221021201220000-0111110231010301-3133001223012003-2131213111312123-3211200123000303"></a>

#### `service.containers.image.name` property

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

- [public](data-sources--workload--reference--group-014.md#canonical-1013002102221300-3120310221330313-2032123311323001-3102220200300221-3323322030321330-0010100030212200-2220330023130320-1320223112003200): complete subsection reference.

<a id="canonical-3220011220110231-0321232322030131-3010220332033203-1000323230230313-1032032130120001-1212132013131133-0110321221001013-0302321210303102"></a>

<a id="canonical-2301013111333132-1210203103303032-2020202301102230-1011033121201202-2331033131223231-1120200023022203-0211222003112330-1303200321131003"></a>

#### `service.containers.image.pull_policy` property

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

<a id="canonical-0200232232103113-0020033231030001-2112000331001000-1120312203111231-2131221223033001-0112022312302223-1232302032023123-3100021001322300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.image.container_registry` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.image](data-sources--workload--reference--group-014.md#canonical-2031322233230031-1232103120020002-3332210332031001-3331321000310203-0102013211002110-0032032121233323-3230101213231312-3030302330121333)
- service.containers.image.container_registry

<a id="canonical-2101203103200010-1100223010230030-0023320222102321-0001202130230111-0003223001330021-1030221300120201-1012312301110011-3110233012222303"></a>

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

<a id="canonical-2321022301201233-3021223013211211-3221330113322203-0012122321000312-0332032232033000-1123030130212313-1330213211333131-1001121203303301"></a>

### Direct properties for `service.containers.image.container_registry`

<a id="canonical-1202131010300301-0212102223133130-3021020103000020-1202032332023330-2332021013020210-3112203302212100-0223003233200000-0013213003000310"></a>

#### `service.containers.image.container_registry.name` property

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

<a id="canonical-1032111120122122-0222002030220032-0123110130131203-0201133030310222-3302233210111231-3131232312302331-3211232302222322-1002010120332223"></a>

<a id="canonical-3031203331133112-1111210033011102-0132021102303233-3313211220221011-0000302223100301-0212323123123003-0001200122131020-1121121220320311"></a>

#### `service.containers.image.container_registry.namespace` property

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

<a id="canonical-3321223103031210-0011011110121332-3312000302213121-2221312310231110-2133032311233120-0322022020311331-1331021010001022-0322212312202023"></a>

<a id="canonical-2021030012122122-0301122332003101-2121101331330103-1203132320011310-3013011302313321-3122031001210302-0221001202101110-2131311310313330"></a>

#### `service.containers.image.container_registry.tenant` property

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

<a id="canonical-1013002102221300-3120310221330313-2032123311323001-3102220200300221-3323322030321330-0010100030212200-2220330023130320-1320223112003200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.image.public` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.image](data-sources--workload--reference--group-014.md#canonical-2031322233230031-1232103120020002-3332210332031001-3331321000310203-0102013211002110-0032032121233323-3230101213231312-3030302330121333)
- service.containers.image.public

<a id="canonical-3211210123200031-1122231210312111-3123123101313000-3220022113210112-1132210203301033-0010220133233303-2232311303103033-2103110311202021"></a>

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

<a id="canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.liveness_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- service.containers.liveness_check

<a id="canonical-3002022213311203-1223201000312100-2013131001201332-1102111223132332-3013113313313132-1130313000010131-0202002101130312-2322002101010202"></a>

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

<a id="canonical-3312113323331223-2221130130003232-0211221000122001-3332113332303201-2201012303330221-0111002122221121-2112021131123000-2300011102123210"></a>

### Direct properties for `service.containers.liveness_check`

- [exec_health_check](data-sources--workload--reference--group-014.md#canonical-1332232120312312-0111300020001122-2231211200220311-2010220321113020-2313123020123031-0232031230123021-0311021231003023-1203331203000122): complete subsection reference.

<a id="canonical-0222230303310103-2102223013222131-0131031230120323-0010223231123100-2312110330222010-0122123113203210-1222022230102101-2321320232213223"></a>

<a id="canonical-3333323122021311-3223332202221322-3313020200300221-3323213010023111-3033103301301330-2133033132131222-0230132320322003-1223123031302311"></a>

#### `service.containers.liveness_check.healthy_threshold` property

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

- [http_health_check](data-sources--workload--reference--group-014.md#canonical-1123231322133301-2311101311332033-2000320331330302-2030201003323030-0023112021030200-2033100032113223-3133033003133121-3120021323313200): complete subsection reference.

<a id="canonical-0121331330011121-0011011211130303-2002110300203210-0120333133002123-2033022323333121-3231112322211322-0332221122031033-3222321300322023"></a>

<a id="canonical-1320312323003303-0022200323202110-1200203023021231-1010022320011033-1133001233120020-2112003210302321-0302011112031222-2322100223022002"></a>

#### `service.containers.liveness_check.initial_delay` property

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

<a id="canonical-0331210002011332-0201102122331112-0112301032120001-3120131213112222-0231230001011122-2220230120321120-2103113213201013-3130033122213310"></a>

<a id="canonical-3131021233201011-1320333013002011-2013220003213301-1011230033332020-2100233321200001-2110130233201120-3203313213123033-1232300010200132"></a>

#### `service.containers.liveness_check.interval` property

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

- [tcp_health_check](data-sources--workload--reference--group-014.md#canonical-2233021321103233-3232122010233002-1230103200310032-0022111223313331-0020133301023222-3022133030203123-2120022220102022-0001202030200101): complete subsection reference.

<a id="canonical-0210103221121330-3323002101313002-3330312312120012-1301022201300103-3010233200020013-0332022313310203-1303001000101031-2011312232033112"></a>

<a id="canonical-1232001002310300-2122110211313100-1001101113203122-0021300011210100-2130313012312112-2231023133131302-2331320211233210-3001020313210023"></a>

#### `service.containers.liveness_check.timeout` property

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

<a id="canonical-0313210231121022-3200320222210222-1101223031113212-3031030010331012-2321232323332323-1112312220300013-2000312300023312-3313230033101121"></a>

<a id="canonical-1120330200122001-1220112110012301-2223133033031200-1220330013002201-2031312321210012-2100122320230013-1230311131111010-3230320222301303"></a>

#### `service.containers.liveness_check.unhealthy_threshold` property

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

<a id="canonical-1332232120312312-0111300020001122-2231211200220311-2010220321113020-2313123020123031-0232031230123021-0311021231003023-1203331203000122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.liveness_check.exec_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.liveness_check](data-sources--workload--reference--group-014.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120)
- service.containers.liveness_check.exec_health_check

<a id="canonical-0221001223012333-3201303221203101-2010321102331101-1033323223123111-2121132010320011-2003020213220230-2111003103331203-3200121120310220"></a>

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

<a id="canonical-1021131203313101-1103002322323310-0231133201200010-1120011310100102-0310010021321033-2200211202231310-2102321220231122-2311322013200023"></a>

### Direct properties for `service.containers.liveness_check.exec_health_check`

<a id="canonical-1302032013111211-1110313102032321-0132201103311032-1002132010223130-1110311302111102-0321013323202011-0112221213303011-1313221311300210"></a>

#### `service.containers.liveness_check.exec_health_check.command` property

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

<a id="canonical-1123231322133301-2311101311332033-2000320331330302-2030201003323030-0023112021030200-2033100032113223-3133033003133121-3120021323313200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.liveness_check.http_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.liveness_check](data-sources--workload--reference--group-014.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120)
- service.containers.liveness_check.http_health_check

<a id="canonical-0002032311221020-2333203000210031-1131223201102322-1033013000122110-2100311001200031-0303221133330222-2333100030111300-3110011232201232"></a>

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

<a id="canonical-1011312130033012-0232110011012220-1331332123221301-1233013300313133-3230133100023211-3101300000223231-3000232303112303-0012330231232111"></a>

### Direct properties for `service.containers.liveness_check.http_health_check`

<a id="canonical-3211021011310231-0231333321232213-1223220112013213-1220103002132003-0233012003212233-0303113133210102-0033323000233030-0102330223233102"></a>

#### `service.containers.liveness_check.http_health_check.headers` property

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

<a id="canonical-3221221130301320-3011001032303220-2101321310230122-2221021011112002-3312221032013300-3321133130103013-3222123111103112-3103222003202232"></a>

<a id="canonical-1023201032212202-0130220102332001-2231210012310232-3333000332110231-1103020300130031-2003030022332002-2331303312123001-2013233031312302"></a>

#### `service.containers.liveness_check.http_health_check.host_header` property

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

<a id="canonical-0210210020320030-1112023323222313-2000213102102022-3213310031212030-0221013231133321-3003022322211301-1321011021200233-0012101101030133"></a>

<a id="canonical-1322030130332010-1132232220010222-0101210321112133-0203222203300031-0312021122030220-2133000020121132-3003201320330300-1230200122130031"></a>

#### `service.containers.liveness_check.http_health_check.path` property

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

- [port](data-sources--workload--reference--group-014.md#canonical-2323332311033013-0102132333310313-1220303311232030-2323132012203211-0203132131332323-3022101301331332-1333333132132203-0022111011012021): complete subsection reference.

<a id="canonical-2323332311033013-0102132333310313-1220303311232030-2323132012203211-0203132131332323-3022101301331332-1333333132132203-0022111011012021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.liveness_check.http_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.liveness_check](data-sources--workload--reference--group-014.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120)
- [service.containers.liveness_check.http_health_check](data-sources--workload--reference--group-014.md#canonical-1123231322133301-2311101311332033-2000320331330302-2030201003323030-0023112021030200-2033100032113223-3133033003133121-3120021323313200)
- service.containers.liveness_check.http_health_check.port

<a id="canonical-1022310121330110-3101201030312222-2322302100021010-3003112021120110-3010202003102303-0213111120023200-0312010023310310-2002020122323210"></a>

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

<a id="canonical-2202202222112333-2031112112200320-1012320300021203-0232130210300212-3303323201232021-1023103011112300-3013032321230312-2322321113033220"></a>

### Direct properties for `service.containers.liveness_check.http_health_check.port`

<a id="canonical-1331310130322021-0233130120330220-3322310123000010-3321122213212232-2102030230223101-2201101220233232-3230323131113100-2331302222132030"></a>

#### `service.containers.liveness_check.http_health_check.port.name` property

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

<a id="canonical-2032210103013003-0301121102221330-0013022120033010-0321001220313013-0320222110221021-2221233233223300-0103030303033130-0331112223223133"></a>

<a id="canonical-1200131102113301-0110311003101022-2100030123000213-2120320011300312-3001100132312230-0211233030313311-1010301000202213-2321020131202332"></a>

#### `service.containers.liveness_check.http_health_check.port.num` property

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

<a id="canonical-2233021321103233-3232122010233002-1230103200310032-0022111223313331-0020133301023222-3022133030203123-2120022220102022-0001202030200101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.liveness_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.liveness_check](data-sources--workload--reference--group-014.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120)
- service.containers.liveness_check.tcp_health_check

<a id="canonical-3130110123132001-3220123000232233-0021001322321110-1300220323223102-1111231230021320-1032011013030211-0203120002102103-3101132122233303"></a>

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

<a id="canonical-3111020322021332-1010221110003010-0213300101230131-1222200033030310-3020131032102131-0223223201221332-2232301220311123-3303131102310301"></a>

### Direct properties for `service.containers.liveness_check.tcp_health_check`

- [port](data-sources--workload--reference--group-014.md#canonical-2220312333010203-0013130000120301-3310011311323303-0313033321300332-1123031313331012-1221212300313320-3130211310110110-0000301031301023): complete subsection reference.

<a id="canonical-2220312333010203-0013130000120301-3310011311323303-0313033321300332-1123031313331012-1221212300313320-3130211310110110-0000301031301023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.liveness_check.tcp_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.liveness_check](data-sources--workload--reference--group-014.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120)
- [service.containers.liveness_check.tcp_health_check](data-sources--workload--reference--group-014.md#canonical-2233021321103233-3232122010233002-1230103200310032-0022111223313331-0020133301023222-3022133030203123-2120022220102022-0001202030200101)
- service.containers.liveness_check.tcp_health_check.port

<a id="canonical-1321301200221112-1010312321210231-0132113300122332-0220300331203203-3031312321023112-2103110012110313-2313201303200110-0021231011222000"></a>

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

<a id="canonical-0303111022020101-2212113030121200-0311223131011023-1302132113222011-2301312103103330-0330211012112200-0032333222321121-2013130003132030"></a>

### Direct properties for `service.containers.liveness_check.tcp_health_check.port`

<a id="canonical-3220231210112013-2132101210030213-2012310333113103-3212321002232121-1212133003313311-0321200023112113-2112322133301310-0133320300033100"></a>

#### `service.containers.liveness_check.tcp_health_check.port.name` property

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

<a id="canonical-2003230100201220-1120332200022202-3301003133102121-1133320003202232-1232001300100130-1211023000030132-0011031010011003-3022122333012330"></a>

<a id="canonical-0010311031002111-0113021312313220-2310210312210013-3200203031201301-2100232310212233-3310200031211312-3200320023113001-2220013100003221"></a>

#### `service.containers.liveness_check.tcp_health_check.port.num` property

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

<a id="canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.readiness_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- service.containers.readiness_check

<a id="canonical-2123302320010122-3032113300023011-2200330022002310-3113002320003321-3100203203323022-2311131112100322-0010301001131122-2300122111223223"></a>

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

<a id="canonical-2000332221110231-1330120312120222-0200222202231102-2111133012020033-1021113221130333-0231001030232320-3330111312130303-2032223230201112"></a>

### Direct properties for `service.containers.readiness_check`

- [exec_health_check](data-sources--workload--reference--group-014.md#canonical-3203123232102231-0203100231032212-1113300221012032-2002031012232310-3333313011131301-1220231231312223-3221201123120101-2103323030021200): complete subsection reference.

<a id="canonical-3220013222312332-0001311132121203-1233223211001311-1010023130231301-3220210022221203-2323120320233112-1000212030201223-3222300003203332"></a>

<a id="canonical-1120213122211330-1022221032323113-1021220311112031-2231132101031011-3331122201110303-3110111113222311-3233200330200211-2112132332322121"></a>

#### `service.containers.readiness_check.healthy_threshold` property

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

- [http_health_check](data-sources--workload--reference--group-014.md#canonical-1223323322103201-2103302313231120-0103310320002203-2333222312011203-1020003312022100-0001111000211211-1111312111112030-3130210210222130): complete subsection reference.

<a id="canonical-2111222010330032-2121212123301333-2203123233120133-2010122233211031-3120222311000220-3010031310123031-0201310323333230-1120230033332122"></a>

<a id="canonical-2133301312122023-2023312222223220-2213311313213312-0002213323121103-2023123002222131-1021212021032021-1303123023023201-2103133012330210"></a>

#### `service.containers.readiness_check.initial_delay` property

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

<a id="canonical-0333003132311012-1023301110303202-2120001230213222-3331102320011212-2210003002021130-0202022000303332-2110302033130030-0112033031101300"></a>

<a id="canonical-3311302220001021-1220302111100233-1211023132210002-0032330030013011-2122221300122223-1110020330112321-3302211113212022-1213100212222111"></a>

#### `service.containers.readiness_check.interval` property

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

- [tcp_health_check](data-sources--workload--reference--group-014.md#canonical-1212111010322100-1002213020102212-2130303313110301-0233301101332221-2321303033123301-3012032310002132-2003232123020003-0300221002312130): complete subsection reference.

<a id="canonical-0311121113221222-3211330013220332-1212022033030302-2113013100101023-3122333201133032-3031200031322123-0010113312312003-0201311312113031"></a>

<a id="canonical-1331210011332321-2012213330221213-1202301020032130-0322001310313221-0212312022130212-0011113322210211-0133002020333231-1002210112120132"></a>

#### `service.containers.readiness_check.timeout` property

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

<a id="canonical-1302210132200212-1321133102201103-1232301131101202-3030201130212002-1303220223033223-3122111012111000-0220003221010200-3003130130231301"></a>

<a id="canonical-3221321322100220-0202101121013212-1213121301231111-2323213200322213-3130200320323232-2230202121120101-3131000222002330-3203122000323111"></a>

#### `service.containers.readiness_check.unhealthy_threshold` property

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

<a id="canonical-3203123232102231-0203100231032212-1113300221012032-2002031012232310-3333313011131301-1220231231312223-3221201123120101-2103323030021200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.readiness_check.exec_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.readiness_check](data-sources--workload--reference--group-014.md#canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012)
- service.containers.readiness_check.exec_health_check

<a id="canonical-3021311203312313-2300120122230002-3130012232023111-0013010010333313-1200103112310130-1332231122122012-1033200211310101-1322222300030213"></a>

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

<a id="canonical-0120302201310110-1032312322003331-2201200110303301-3203321231111222-3211003013021102-3232112023020011-3212202233213130-2102101211113232"></a>

### Direct properties for `service.containers.readiness_check.exec_health_check`

<a id="canonical-3020211010213321-1310033030213022-3031121002132300-0000220032120221-3230100102221311-2212112210010232-2200321303021220-3102220300100000"></a>

#### `service.containers.readiness_check.exec_health_check.command` property

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

<a id="canonical-1223323322103201-2103302313231120-0103310320002203-2333222312011203-1020003312022100-0001111000211211-1111312111112030-3130210210222130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.readiness_check.http_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.readiness_check](data-sources--workload--reference--group-014.md#canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012)
- service.containers.readiness_check.http_health_check

<a id="canonical-2100303330220011-3321311300123220-2000020210010130-2020221221113332-0030332103303112-0311032112212203-2312212132112131-3012302102330333"></a>

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

<a id="canonical-1300020133302210-2201000300000333-1313131321002120-2223132000000321-0010222110031222-3200223112313103-1310321031130330-2133222201000233"></a>

### Direct properties for `service.containers.readiness_check.http_health_check`

<a id="canonical-0223100133010033-0332103330232110-3212030131132133-2023113232311120-0200201030030231-0133132123133101-1010021313023020-1103032102331223"></a>

#### `service.containers.readiness_check.http_health_check.headers` property

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

<a id="canonical-0023332301222323-0033332303111020-2212300331321221-0332310113201313-0030322230222313-2012332222132023-2103023113101111-0101233221120033"></a>

<a id="canonical-2110301030231102-1223031233201203-2010223100011313-3311012110332013-2103300232233213-2202033332200133-3231011210230203-3101310003301023"></a>

#### `service.containers.readiness_check.http_health_check.host_header` property

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

<a id="canonical-0312210320012002-0310333110303232-1123330111203322-0332112321301211-1301230320120221-0333302220220201-1120331120221010-1121002011213222"></a>

<a id="canonical-2033123102121311-3130012131131013-0101020201302312-0302333112103121-1233123110232210-3121222322100101-3230200112210012-2230331103321221"></a>

#### `service.containers.readiness_check.http_health_check.path` property

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

- [port](data-sources--workload--reference--group-014.md#canonical-2331301302232210-3223012223110031-1001322111123023-3230303200030321-2033220311223231-2321111202130233-2003303303100321-1230033120132123): complete subsection reference.

<a id="canonical-2331301302232210-3223012223110031-1001322111123023-3230303200030321-2033220311223231-2321111202130233-2003303303100321-1230033120132123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.readiness_check.http_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.readiness_check](data-sources--workload--reference--group-014.md#canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012)
- [service.containers.readiness_check.http_health_check](data-sources--workload--reference--group-014.md#canonical-1223323322103201-2103302313231120-0103310320002203-2333222312011203-1020003312022100-0001111000211211-1111312111112030-3130210210222130)
- service.containers.readiness_check.http_health_check.port

<a id="canonical-1110322020311323-2102302110323013-1033221202230303-1221022101001233-2100021313023123-2330030333102103-2030103001320011-2132321202100230"></a>

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

<a id="canonical-1313012001110331-1130022311332113-3130021211130133-1020330011023013-3102020030133121-0322212302131031-3203103212312121-1331323233101133"></a>

### Direct properties for `service.containers.readiness_check.http_health_check.port`

<a id="canonical-2213323113130122-3320301122330223-1303322311200122-2012322010000011-3213003133120131-1000323300310020-0102212202232231-2330230223013213"></a>

#### `service.containers.readiness_check.http_health_check.port.name` property

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

<a id="canonical-3231012302200300-2212302213313010-2012003031302203-3232231323031221-1000100320300132-0023323111012300-3113120210311301-2131222012333031"></a>

<a id="canonical-2133220301103203-0021300030213232-2201000301200332-3111031023111302-0333220221122220-1231033212321111-2231033012130332-0233101301030110"></a>

#### `service.containers.readiness_check.http_health_check.port.num` property

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

<a id="canonical-1212111010322100-1002213020102212-2130303313110301-0233301101332221-2321303033123301-3012032310002132-2003232123020003-0300221002312130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.readiness_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.readiness_check](data-sources--workload--reference--group-014.md#canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012)
- service.containers.readiness_check.tcp_health_check

<a id="canonical-0212201202111330-3321231020133011-0221220330223303-0222001212301010-0101130031032300-0221110300001202-2220300222100111-3333331133103221"></a>

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

<a id="canonical-0110231303013123-0111000231201333-1123133333213023-1101320001133200-0030002100122002-3010130100310103-3221100322123020-0333111203000313"></a>

### Direct properties for `service.containers.readiness_check.tcp_health_check`

- [port](data-sources--workload--reference--group-014.md#canonical-3002222033122000-1102232133123211-3303201331122313-3002333110123022-0101013032120220-0322022133001313-1020023323111230-3222311132013121): complete subsection reference.

<a id="canonical-3002222033122000-1102232133123211-3303201331122313-3002333110123022-0101013032120220-0322022133001313-1020023323111230-3222311132013121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.readiness_check.tcp_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.readiness_check](data-sources--workload--reference--group-014.md#canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012)
- [service.containers.readiness_check.tcp_health_check](data-sources--workload--reference--group-014.md#canonical-1212111010322100-1002213020102212-2130303313110301-0233301101332221-2321303033123301-3012032310002132-2003232123020003-0300221002312130)
- service.containers.readiness_check.tcp_health_check.port

<a id="canonical-3012300313010303-1023200203310003-1011201212033023-0330121200331311-2101021233200220-3312301121021122-3200323103322132-2030222313031122"></a>

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

<a id="canonical-1300323200232211-3232320321221003-2331131313133313-2032132101000333-3323312120110232-2001130020201131-3201301132201031-3332301211303132"></a>

### Direct properties for `service.containers.readiness_check.tcp_health_check.port`

<a id="canonical-0121121322022003-3103013013101033-1212333322122200-2131001012321112-3003020011010213-2231102012220131-1303330010033232-2313022110001331"></a>

#### `service.containers.readiness_check.tcp_health_check.port.name` property

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

<a id="canonical-0212210320232131-2203203213100133-3122033210103323-0000133213201312-2003303211202020-1201013211210122-0232030203133003-3230021310213103"></a>

<a id="canonical-0223132103323012-2022102330302313-0010022200303200-1113131332112133-1313223033020311-1030012311203003-3310111210120122-0212010212330132"></a>

#### `service.containers.readiness_check.tcp_health_check.port.num` property

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

<a id="canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- service.deploy_options

<a id="canonical-0231300112311113-1011102300323233-1222132313331310-1112213322032131-3120201131032002-3023022221220031-2303020001320122-0000133231103323"></a>

Type: `"single"`. Computed.

Deploy OPTIONS are used to configure the workload deployment OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-deploy_choice": "[\"all_res\",\"default_virtual_sites\",\"deploy_ce_sites\",\"deploy_ce_virtual_sites\",\"deploy_re_sites\",\"deploy_re_virtual_sites\"]"
}
```

<a id="canonical-0300221023031112-3322322212333122-2233210222230300-3021020231232113-0021021320223102-2221122220211130-1111323210220111-2130210022102133"></a>

### Direct properties for `service.deploy_options`

- [all_res](data-sources--workload--reference--group-014.md#canonical-0202000322003301-0203221223032321-3222210310333311-0221213122223133-2331203020001121-2230320220020303-2121302032302212-0203321012000132): complete subsection reference.

- [default_virtual_sites](data-sources--workload--reference--group-014.md#canonical-0112333223300320-0222131002221301-0213301111220311-3132213322213222-3130320003112100-1300220013020330-1110222212020022-1222123322310132): complete subsection reference.

- [deploy_ce_sites](data-sources--workload--reference--group-014.md#canonical-1220333231032303-1013123311222301-3132310033223123-2021321322301132-3321010000211301-1210212301002230-1231223332331321-1120333231133203): complete subsection reference.

- [deploy_ce_virtual_sites](data-sources--workload--reference--group-014.md#canonical-3032332130302300-1030222322333200-3212210321133210-3221313101010100-2023133330132301-1212113110012202-2220322131332100-3003030330130003): complete subsection reference.

- [deploy_re_sites](data-sources--workload--reference--group-014.md#canonical-1303232022011320-1200101023133223-1211210313033103-2030201303303322-0231212030222200-3113013231131010-0012330122032110-2201123321132232): complete subsection reference.

- [deploy_re_virtual_sites](data-sources--workload--reference--group-014.md#canonical-1032333221231212-0012100102110032-2232330111130330-0002233211230220-3011231233013001-2331331003213001-3221223033110011-3211122303201310): complete subsection reference.

<a id="canonical-0202000322003301-0203221223032321-3222210310333311-0221213122223133-2331203020001121-2230320220020303-2121302032302212-0203321012000132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.all_res` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.deploy_options](data-sources--workload--reference--group-014.md#canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320)
- service.deploy_options.all_res

<a id="canonical-2201211010232210-2010321232132112-2102131102110111-3000220212032212-0232100102013213-1121112222332121-0202102311120033-0032200031021323"></a>

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

<a id="canonical-0112333223300320-0222131002221301-0213301111220311-3132213322213222-3130320003112100-1300220013020330-1110222212020022-1222123322310132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.default_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.deploy_options](data-sources--workload--reference--group-014.md#canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320)
- service.deploy_options.default_virtual_sites

<a id="canonical-3220302330311311-3130030100221210-3302013200320023-1111131030112123-3121302330030311-1230310303223101-0230102302323000-3031103100111313"></a>

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

<a id="canonical-1220333231032303-1013123311222301-3132310033223123-2021321322301132-3321010000211301-1210212301002230-1231223332331321-1120333231133203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_ce_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.deploy_options](data-sources--workload--reference--group-014.md#canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320)
- service.deploy_options.deploy_ce_sites

<a id="canonical-2300233110333133-3222120332221200-2302320321100201-0121112332212210-0223201100200122-2111022203032032-2110222331030021-0011020330300222"></a>

Type: `"single"`. Computed.

This defines a way to deploy a workload on specific Customer sites.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0321212100103133-2133202313003331-2210232203012233-0233123312133101-0313231133320302-1301210212213303-3132331233222323-0210333220013110"></a>

### Direct properties for `service.deploy_options.deploy_ce_sites`

- [site](data-sources--workload--reference--group-014.md#canonical-1301031121133121-2123221120122231-2101110230231332-3213001210010203-3100012201333322-1202332001133312-0322133211311122-2031230232333103): complete subsection reference.

<a id="canonical-1301031121133121-2123221120122231-2101110230231332-3213001210010203-3100012201333322-1202332001133312-0322133211311122-2031230232333103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_ce_sites.site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.deploy_options](data-sources--workload--reference--group-014.md#canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320)
- [service.deploy_options.deploy_ce_sites](data-sources--workload--reference--group-014.md#canonical-1220333231032303-1013123311222301-3132310033223123-2021321322301132-3321010000211301-1210212301002230-1231223332331321-1120333231133203)
- service.deploy_options.deploy_ce_sites.site

<a id="canonical-3011012000311111-1233303210232212-0032010331312330-1120102033001132-1332011030110030-0002210331102332-3311323132132020-2100223123013110"></a>

Type: `"list"`. Computed.

Which customer sites should this workload be deployed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0000210103321303-2221003200020111-1023120303212332-1120010012022010-3022111203320200-2113201111202322-2201003312022033-1023312222302201"></a>

### Direct properties for `service.deploy_options.deploy_ce_sites.site`

<a id="canonical-3302000031330201-2212102300020231-1013333200310011-3233201003011122-1223110230311201-3220132033102132-2132310321033032-2120201110112100"></a>

#### `service.deploy_options.deploy_ce_sites.site.name` property

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

<a id="canonical-0133020220003221-1132122101031001-0330121231222131-1213130132222330-2031303333130321-0230302300331202-2313320230002111-0102332103323020"></a>

<a id="canonical-2031120031023213-2130001220220113-0223220132310012-1033132023303010-1211113113110001-1203032302012331-1230111013200301-3321331102130000"></a>

#### `service.deploy_options.deploy_ce_sites.site.namespace` property

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

<a id="canonical-0003111302220230-3220331100233132-2301120110213302-3302001030203130-1313012022133303-0321330201012031-3011231230030112-0311220023012131"></a>

<a id="canonical-0001321001022120-1010033112010101-3011311011211133-0020223231300132-0011000231301211-1211333230133302-0112301202112101-2102330022320203"></a>

#### `service.deploy_options.deploy_ce_sites.site.tenant` property

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

<a id="canonical-3032332130302300-1030222322333200-3212210321133210-3221313101010100-2023133330132301-1212113110012202-2220322131332100-3003030330130003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_ce_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.deploy_options](data-sources--workload--reference--group-014.md#canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320)
- service.deploy_options.deploy_ce_virtual_sites

<a id="canonical-1331001031212211-1222130201100130-3232030312013102-1130013200232303-0032001322030121-1123323310322110-1112313230312120-1131233203303330"></a>

Type: `"single"`. Computed.

This defines a way to deploy a workload on specific Customer virtual sites.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2001100321022303-1200023213211222-2230221302022003-1332022100200332-3003101220201323-3233001220221001-2100010023032010-0331022221100310"></a>

### Direct properties for `service.deploy_options.deploy_ce_virtual_sites`

- [virtual_site](data-sources--workload--reference--group-014.md#canonical-2002002023000013-1331113311220311-0012330131213201-2303111000332323-0220213110222113-1222331103023022-3330012322011200-0212010213032300): complete subsection reference.

<a id="canonical-2002002023000013-1331113311220311-0012330131213201-2303111000332323-0220213110222113-1222331103023022-3330012322011200-0212010213032300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_ce_virtual_sites.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.deploy_options](data-sources--workload--reference--group-014.md#canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320)
- [service.deploy_options.deploy_ce_virtual_sites](data-sources--workload--reference--group-014.md#canonical-3032332130302300-1030222322333200-3212210321133210-3221313101010100-2023133330132301-1212113110012202-2220322131332100-3003030330130003)
- service.deploy_options.deploy_ce_virtual_sites.virtual_site

<a id="canonical-3103220011000310-2123102320101322-0301012323132102-0212111133011201-1300233333121013-2202013220132130-0122010132112020-0230232210020122"></a>

Type: `"list"`. Computed.

Which customer virtual sites should this workload be deployed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0111211303032300-1022021330312333-2313010003102100-1133230312012213-2201031021003020-1112300311223313-2101303332203003-3231120001023131"></a>

### Direct properties for `service.deploy_options.deploy_ce_virtual_sites.virtual_site`

<a id="canonical-3301231123102313-2122211223323112-3303200101310210-0213201311110010-3023112120322303-1322321212100323-3011200022020310-2131330110001211"></a>

#### `service.deploy_options.deploy_ce_virtual_sites.virtual_site.name` property

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

<a id="canonical-3302103113311002-2300022222302232-0102201330103020-3012333100110331-1213110001303111-2101123012012331-0001333110010201-3230320210032231"></a>

<a id="canonical-1203321013132331-0232312323300231-0331203332103303-3100333300230021-1121211103112313-2222200232311130-0113003210010013-3013313132111202"></a>

#### `service.deploy_options.deploy_ce_virtual_sites.virtual_site.namespace` property

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

<a id="canonical-0313333221310200-0021120332300112-3230113112010303-2211231112020200-2300120221102032-0332233320101100-3011302320031003-0301331001113300"></a>

<a id="canonical-3211203121000302-2010330230001231-3110321310102110-2103203122321000-1220221220002021-0033003012211103-2303213333201323-2110133032322101"></a>

#### `service.deploy_options.deploy_ce_virtual_sites.virtual_site.tenant` property

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

<a id="canonical-1303232022011320-1200101023133223-1211210313033103-2030201303303322-0231212030222200-3113013231131010-0012330122032110-2201123321132232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_re_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.deploy_options](data-sources--workload--reference--group-014.md#canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320)
- service.deploy_options.deploy_re_sites

<a id="canonical-2203120220220203-0123302202213322-3101331303333000-2200023202021120-2022021210232301-2011211110032312-0023032112013002-2121220103212123"></a>

Type: `"single"`. Computed.

This defines a way to deploy a workload on specific Regional Edge sites.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3323030113002121-1213032013301113-3333301002221322-3233310110230111-0313312332000032-0233213112220002-3210011231032012-2232020302312133"></a>

### Direct properties for `service.deploy_options.deploy_re_sites`

- [site](data-sources--workload--reference--group-014.md#canonical-1111201233321320-0113101012103223-0131202121022103-1133312210302320-2013202012233333-2133030333010132-1132103003003213-3111201111221313): complete subsection reference.

<a id="canonical-1111201233321320-0113101012103223-0131202121022103-1133312210302320-2013202012233333-2133030333010132-1132103003003213-3111201111221313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_re_sites.site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.deploy_options](data-sources--workload--reference--group-014.md#canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320)
- [service.deploy_options.deploy_re_sites](data-sources--workload--reference--group-014.md#canonical-1303232022011320-1200101023133223-1211210313033103-2030201303303322-0231212030222200-3113013231131010-0012330122032110-2201123321132232)
- service.deploy_options.deploy_re_sites.site

<a id="canonical-3210111111020303-3223202111333202-3131130230330322-0113230303231303-1112011132332323-2130131013123002-0133110003233130-0122323011113231"></a>

Type: `"list"`. Computed.

Which regional edge sites should this workload be deployed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1221333023312001-3230232122032023-0211323012110320-1100300002111100-0203323101323202-0132301011223012-0000232313132231-3223023332032211"></a>

### Direct properties for `service.deploy_options.deploy_re_sites.site`

<a id="canonical-0032313121100201-3001310332020023-3020003022031110-0233321213100020-1003133020131223-1232020230311312-0103223323212102-0100011110301222"></a>

#### `service.deploy_options.deploy_re_sites.site.name` property

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

<a id="canonical-0322101113302322-0311030212213021-1232233121122002-1013020132133101-3021002011112300-0201112021333223-1133230213313311-0133111112203003"></a>

<a id="canonical-0102122001203211-0131332013300023-0213323103202320-2033201301303133-3013031213102313-3001132203331012-2032122323210321-1210223111030322"></a>

#### `service.deploy_options.deploy_re_sites.site.namespace` property

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

<a id="canonical-0101130001331211-1312021011320203-0211201130200220-2302103221220122-3110013101103233-2003030111130030-0313300200022100-2300322200201213"></a>

<a id="canonical-0212002320131300-0221122131100223-3221021122103003-0101330330132332-3001212123210210-0001330230031321-3012321322233331-2132112002112003"></a>

#### `service.deploy_options.deploy_re_sites.site.tenant` property

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

<a id="canonical-1032333221231212-0012100102110032-2232330111130330-0002233211230220-3011231233013001-2331331003213001-3221223033110011-3211122303201310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_re_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.deploy_options](data-sources--workload--reference--group-014.md#canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320)
- service.deploy_options.deploy_re_virtual_sites

<a id="canonical-1201312021033111-2312230313020322-0232320230113200-0331112013301012-0022212322232230-0313302301103321-0101213321131033-2231121101100202"></a>

Type: `"single"`. Computed.

This defines a way to deploy a workload on specific Regional Edge virtual sites.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1113022113001210-2031220101133220-0103111222310311-2223012030200111-1212211001023321-0002032300112330-2111322320323103-1233121013001322"></a>

### Direct properties for `service.deploy_options.deploy_re_virtual_sites`

- [virtual_site](data-sources--workload--reference--group-014.md#canonical-1232322202202201-0120301103011021-2201212300110333-2102020122131132-0323102112333123-0032111212013322-2122330130013101-1113213021312132): complete subsection reference.

<a id="canonical-1232322202202201-0120301103011021-2201212300110333-2102020122131132-0323102112333123-0032111212013322-2122330130013101-1113213021312132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_re_virtual_sites.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.deploy_options](data-sources--workload--reference--group-014.md#canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320)
- [service.deploy_options.deploy_re_virtual_sites](data-sources--workload--reference--group-014.md#canonical-1032333221231212-0012100102110032-2232330111130330-0002233211230220-3011231233013001-2331331003213001-3221223033110011-3211122303201310)
- service.deploy_options.deploy_re_virtual_sites.virtual_site

<a id="canonical-3133213000320123-0220320230233223-3333001020020313-2110313201103112-1223201101300201-3103320203230032-3132202110321312-0232211311301333"></a>

Type: `"list"`. Computed.

Which regional edge virtual sites should this workload be deployed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3000300323210303-2010211231330212-3221222130023031-3112022012003001-2211022333021020-3032333230033333-0120202311102012-2311302332002211"></a>

### Direct properties for `service.deploy_options.deploy_re_virtual_sites.virtual_site`

<a id="canonical-2003130003212220-2122123320203001-2010203331111100-1313113310310020-2023202330023301-3121132001113233-0203002002301303-2230221230333202"></a>

#### `service.deploy_options.deploy_re_virtual_sites.virtual_site.name` property

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

<a id="canonical-0323033103033220-2010033021300223-0300130113102213-1313203220221023-1123322331031223-2203220303201302-2233303123112333-2313323212020103"></a>

<a id="canonical-0300230033222333-1221033212101311-1330032023223032-3031320003312030-3120000103220023-0201232023101123-0113233202322030-1032202202103330"></a>

#### `service.deploy_options.deploy_re_virtual_sites.virtual_site.namespace` property

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

<a id="canonical-1223231213000102-1232033313233011-3102001202101210-1223020303202231-2011223120012222-0133102332323201-0310011210313103-0130302033212202"></a>

<a id="canonical-2022113203131300-2002113310302102-1100033030300003-2130120311011021-0202012122201230-1001022101323013-3111010002000202-0103021210133131"></a>

#### `service.deploy_options.deploy_re_virtual_sites.virtual_site.tenant` property

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

<a id="canonical-1232232223133202-3202330012213121-2321010021332212-2131311121001321-2011131113311031-0312303012321112-3333213331000032-1000100101021221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.scale_to_zero` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- service.scale_to_zero

<a id="canonical-1022210201030331-3021111313023011-1103332123113023-1211313221323012-2132220201232212-2232121031331303-0223211022122113-3333133210322312"></a>

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
